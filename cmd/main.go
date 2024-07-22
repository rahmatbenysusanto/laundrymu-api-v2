package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"laundrymu-api/internal/route"
	"time"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	app.Use(logger.New())
	route.PublicAPI(app)
	route.PrivateAPI(app)

	//go kafka.Consumer()

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
