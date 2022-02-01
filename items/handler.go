package items

import (
	"fmt"
	"strings"

	"github.com/mehabox/balabol"

	routing "github.com/qiangxue/fasthttp-routing"
)

type Handler struct {
	repo balabol.ItemsRepository
}

// NewItemsHandler returns an initialized items handler.
func NewItemsHandler(repo balabol.ItemsRepository) *Handler {
	// populate storage with test values (api/v1/test1, api/v1/texts/hidden).
	presetItem1 := balabol.NewItem("test1", balabol.NewStringValue("TEST ME"))
	_ = repo.Save(presetItem1)

	presetItem2 := balabol.NewItem("texts/hidden", balabol.NewStringValue("I am another value you can get!"))
	_ = repo.Save(presetItem2)

	return &Handler{repo: repo}
}

// InitRoutes initializes this handler's routes :/.
func (h *Handler) InitRoutes(r *routing.Router) error {
	r.Get(`/api/v<version:\d+><path:.*>`, h.getValue)
	return nil
}

// getValue is a GET handler (gets values from the repository, if they exist).
func (h *Handler) getValue(c *routing.Context) error {
	path := strings.Trim(c.Param("path"), "/ \\")

	item, err := h.repo.GetByPath(path)
	if err != nil {
		c.SetStatusCode(500)
		_, _ = fmt.Fprintf(c, err.Error())
		return err
	}

	if item == nil {
		c.SetStatusCode(404)
		_, _ = fmt.Fprintf(c, "not found")
		return nil
	}

	_, _ = fmt.Fprintf(c, "%s", item.Value().Bytes())
	return nil
}
