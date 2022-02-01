package main

import (
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/mehabox/balabol/infra/storage"
	"github.com/mehabox/balabol/items"

	"github.com/mehabox/balabol"
	// "github.com/mehabox/balabol/admin"
	"github.com/mehabox/balabol/http"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v2"
)

var version = "dev"

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	log := zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02 15:04:05",
	}).With().Timestamp().Logger()

	logger := balabol.NewAppLogger(&log)
	logger.Printf("balabol version %s is starting", version)

	repo := items.NewItemsRepository(storage.NewMemoryMapStorage())

	handlers := http.HandlerList{
		&http.MainHandler{Logger: logger},
		items.NewItemsHandler(repo),
	}
	router := routing.New()

	file, err := ioutil.ReadFile("./cmd/balabol/config.yml")
	if err != nil {
		logger.Error("can't read config file " + err.Error())
		os.Exit(1)
	}

	var ymlConfig config

	err = yaml.Unmarshal(file, &ymlConfig)
	if err != nil {
		logger.Error("can't parse config file " + err.Error())
		os.Exit(1)
	}

	cfg := getConfig(ymlConfig)
	logger.Printf("starting web server at %s", cfg.ServerConfig.Listen)
	app := http.NewApplication(cfg, handlers, router, logger)

	go func() {
		err = app.Start()
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	}()

	logger.Println("done!")
	waitForShutdown(c, app)
}

// getConfig transforms yaml config to app config.
func getConfig(config config) balabol.AppConfig {
	return balabol.AppConfig{
		ServerConfig: balabol.ServerConfig{
			ReadTimeout: config.ServerConfig.ReadTimeout,
			Listen:      config.ServerConfig.Listen,
		},
	}
}

// waitForShutdown performs graceful shutdown.
func waitForShutdown(c <-chan os.Signal, app *http.Application) {
	<-c

	err := app.Shutdown()
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
