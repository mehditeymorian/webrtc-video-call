package main

import (
	"log"

	"hermes/internal/http/handler"
	"hermes/internal/store"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	st := store.Create()

	handler.Room{Store: st}.Register(app)

	log.Fatal(app.Listen(":8080"))
}
