package http

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	"log"
)

type IndexHandler struct {
}

func (h *IndexHandler) InitRoutes(r *routing.Router) error {
	r.Get("/", h.index)
	log.Printf("setting routes for index.go")
	return nil
}

func (h *IndexHandler) index(c *routing.Context) error {
	fmt.Fprintf(c, "Welcome")
	return nil
}
