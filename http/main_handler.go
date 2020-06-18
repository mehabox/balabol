package http

import (
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
)

type IndexHandler struct {
}

func (h *IndexHandler) InitRoutes(r *routing.Router) error {
	r.Get("/", h.index)
	return nil
}

func (h *IndexHandler) index(c *routing.Context) error {
	fmt.Fprintf(c, "Welcome")
	return nil
}
