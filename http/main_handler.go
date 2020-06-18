package http

import (
	"fmt"

	"github.com/mehabox/balabol"

	routing "github.com/qiangxue/fasthttp-routing"
)

type MainHandler struct {
	Logger balabol.Logger
}

func (h *MainHandler) InitRoutes(r *routing.Router) error {
	r.Get("/", logging(slow(h.index), h.Logger))
	return nil
}

func (h *MainHandler) index(c *routing.Context) error {
	fmt.Fprintf(c, "Welcome")
	return nil
}
