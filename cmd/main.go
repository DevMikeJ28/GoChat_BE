package main 

import (
	"log";
	"github.com/gofiber/fiber/v2";
)

func main() {

	app := fiber.New()

	//Router:
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	log.Println("Server running at:: 8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}

}