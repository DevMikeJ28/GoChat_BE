package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	//Router:
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})


	log.Println("Server running at:: 8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}

}