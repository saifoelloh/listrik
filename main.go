package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/saifoelloh/listrik/api/routes"
	"github.com/saifoelloh/listrik/initializers"
	"github.com/saifoelloh/listrik/pkg/pln"
)

func init() {
	initializers.GoDotEnvVariable()
	err := initializers.DatabaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")
}

func main() {
	databaseName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("DB_COLLECTION_PLN")

	db := initializers.Client.Database(databaseName)
	plnCollection := db.Collection(collectionName)
	plnRepo := pln.NewRepo(plnCollection)
	plnService := pln.NewService(plnRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(limiter.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/readiness", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	api := app.Group("/api")
	routes.RoutePLN(api, plnService)

	defer initializers.Cancel()
	log.Fatal(app.Listen(":8080"))
}
