package admin

import (
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
)

type Handler struct {
}

func (h *Handler) InitRoutes(r *routing.Router) error {
	r.Get("/admin", h.index)
	r.Get("/admin/db", h.db)

	return nil
}

func (h *Handler) index(c *routing.Context) error {
	fmt.Fprintf(c, "Welcome to admin!")
	return nil
}

func (h *Handler) db(c *routing.Context) error {
	fmt.Fprintf(c, "Welcome to db!")
	return nil
}
