package items

import (
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
)

type Handler struct{}

// InitRoutes initializes this handler's routes :/.
func (h *Handler) InitRoutes(r *routing.Router) error {
	r.Get(`/api/v<version:\d+><path:.*>`, h.getValue)
	return nil
}

func (h *Handler) getValue(c *routing.Context) error {
	fmt.Fprintf(c, "Dynamic route, version:%s, key path: %s",
		c.Param("version"), c.Param("path"))
	return nil
}
