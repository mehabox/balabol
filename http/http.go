package http

import (
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

// Handler is an http handler with a defined list
// of routes and actions associated with them
type Handler interface {
	// InitRoutes attaches Handler's methods
	InitRoutes(r *routing.Router) error
}

// HandlerList is a list of Handlers
type HandlerList []Handler

// Application is a collection of Handlers that initializes them
type Application struct {
	handlers HandlerList
	router   *routing.Router
	server   *fasthttp.Server
}

// NewApplication todo what it does
func NewApplication(handlers HandlerList, router *routing.Router) *Application {
	server := &fasthttp.Server{
		Name:    "balabol",
		Handler: router.HandleRequest,
	}
	return &Application{handlers: handlers, router: router, server: server}
}

// Start initializes all of the handlers and starts listening on the designated port
func (app *Application) Start() error {
	for _, handler := range app.handlers {
		err := handler.InitRoutes(app.router)

		if err != nil {
			return err
		}
	}

	go app.server.ListenAndServe(":8080")
	return nil
}
