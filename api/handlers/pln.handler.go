package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/saifoelloh/listrik/api/presenter"
	"github.com/saifoelloh/listrik/pkg/pln"
)

// GetBooks is handler/controller which lists all Books from the BookShop
func GetBooks(service pln.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchUsers(pln.QueryReadUser{
			Name:    c.Query("name"),
			SortBy:  c.Query("sort-by"),
			OrderBy: c.Query("order-by"),
			Page:    uint16(c.QueryInt("page")),
			Show:    uint16(c.QueryInt("show")),
		})
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserPLNErrorResponse(err))
		}
		return c.JSON(presenter.UsersPLNSuccessResponse(fetched))
	}
}
