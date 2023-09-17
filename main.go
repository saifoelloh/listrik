package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/saifoelloh/listrik/api/routes"
	"github.com/saifoelloh/listrik/helper"
	"github.com/saifoelloh/listrik/pkg/pln"
)

func main() {
	db, cancel, err := helper.DatabaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")

	plnCollection := db.Collection("Users")
	plnRepo := pln.NewRepo(plnCollection)
	plnService := pln.NewService(plnRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")
	routes.RoutePLN(api, plnService)

	defer cancel()
	log.Fatal(app.Listen(":8080"))
}
