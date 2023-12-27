package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	app.Get("/",func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})


	port := os.Getenv("PORT")
	if port == ""{
		port = "3000"
	}

	app.Listen(":" + port)
}