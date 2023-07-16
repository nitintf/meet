package main

import (
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	app.Listen(":3000")
}
