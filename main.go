package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/saifoelloh/listrik/api/routes"
	"github.com/saifoelloh/listrik/helper"
	"github.com/saifoelloh/listrik/pkg/pln"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	user := helper.GoDotEnvVariable("DB_USERNAME")
	pass := helper.GoDotEnvVariable("DB_PASSWORD")
	host := helper.GoDotEnvVariable("DB_HOST")
	port := helper.GoDotEnvVariable("DB_PORT")
	databaseName := helper.GoDotEnvVariable("DB_NAME")
	mongoURI := "mongodb://" + user + ":" + pass + "@" + host + ":" + port + "/lokal"
	fmt.Println("url: ", mongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	mongoOption := options.Client().ApplyURI("mongodb://admin:User%231st@localhost:27017/?authSource=admin").SetServerSelectionTimeout(5 * time.Second)
	client, err := mongo.Connect(ctx, mongoOption)
	if err != nil {
		cancel()
		return nil, nil, err
	}

	db := client.Database(databaseName)
	return db, cancel, nil
}

func main() {
	db, cancel, err := databaseConnection()
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
