package http

import (
	"time"

	"github.com/mehabox/balabol"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

// Handler is an http handler with a defined list
// of routes and actions associated with them.
type Handler interface {
	// InitRoutes attaches Handler's methods.
	InitRoutes(r *routing.Router) error
}

// HandlerList is a list of Handlers.
type HandlerList []Handler

// Application is a collection of Handlers that initializes them.
type Application struct {
	handlers HandlerList
	router   *routing.Router
	server   *fasthttp.Server
	logger   balabol.Logger
	config   balabol.AppConfig
}

// NewApplication initializes the app, settings routes and making it ready to start.
func NewApplication(config balabol.AppConfig, handlers HandlerList, router *routing.Router, logger balabol.Logger) *Application {
	server := &fasthttp.Server{
		Handler:      router.HandleRequest,
		Name:         "balabol",
		ReadTimeout:  config.ServerConfig.ReadTimeout,
		WriteTimeout: config.ServerConfig.WriteTimeout,
		IdleTimeout:  config.ServerConfig.IdleTimeout,
		Logger:       logger,
	}

	return &Application{handlers: handlers, router: router, server: server, logger: logger, config: config}
}

// Start initializes all of the handlers and starts listening on the designated port.
func (app *Application) Start() error {
	for _, handler := range app.handlers {
		err := handler.InitRoutes(app.router)
		if err != nil {
			return err
		}
	}

	err := app.server.ListenAndServe(app.config.ServerConfig.Listen)
	if err != nil {
		return err
	}

	return nil
}

// Shutdown gracefully terminates http server.
func (app *Application) Shutdown() error {
	t1 := time.Now()

	app.logger.Printf("shutting down")
	err := app.server.Shutdown()

	elapsed := time.Since(t1)
	app.logger.Printf("shut down in %s", elapsed)

	return err
}
