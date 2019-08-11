package http

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	"log"
)

type AdminHandler struct {
}

func (h *AdminHandler) InitRoutes(r *routing.Router) error {
	r.Get("/admin", h.index)
	r.Get("/admin/db", h.db)
	log.Printf("setting routes for admin.go")
	return nil
}

func (h *AdminHandler) index(c *routing.Context) error {
	fmt.Fprintf(c, "Welcome to admin!")
	return nil
}

func (h *AdminHandler) db(c *routing.Context) error {
	fmt.Fprintf(c, "Welcome to db!")
	return nil
}
