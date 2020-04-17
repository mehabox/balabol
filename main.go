package main

import (
	"github.com/mehabox/balabol/http"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

var version = "dev"

func main() {
	waitForShutdown()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	handlers := http.HandlerList{
		&http.IndexHandler{},
		&http.AdminHandler{},
		&http.ApiHandler{},
	}

	log.Info().Msgf("balabol version %s is starting", version)

	log.Info().Msg("starting web server")

	router := routing.New()

	app := http.NewApplication(handlers, router)
	app.Start()

	log.Info().Msg("done!")

	select {}
}

func waitForShutdown() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		shutdown()
		os.Exit(1)
	}()

}
func shutdown() {
	log.Info().Msg("Terminating")
}
