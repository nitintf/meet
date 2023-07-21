package main

import (
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nitintf/meet/internal/config"
	"github.com/nitintf/meet/internal/logger"
	"go.uber.org/zap"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	cfg, err := config.New()

	log := logger.New()

	if err != nil {
		log.Panic("Unable to initialize config")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	log.Info("Server started", zap.Int("port", cfg.Port))
	if err = app.Listen(":3000"); err != nil {
		log.Error("error starting server", zap.Error(err))
	}
}
