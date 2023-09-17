package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saifoelloh/listrik/api/handlers"
	"github.com/saifoelloh/listrik/pkg/pln"
)

func RoutePLN(app fiber.Router, service pln.Service) {
	plnGroup := app.Group("/pln")
	plnGroup.Get("/user", handlers.GetBooks(service))
}
