package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/pusher/pusher-http-go"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("failed to load: %v", err)
	}

	app := fiber.New()

	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID:   os.Getenv("PUSHER_APP_ID"),
		Key:     os.Getenv("PUSHER_KEY"),
		Secret:  os.Getenv("PUSHER_SECRET"),
		Cluster: "ap3",
		Secure:  true,
	}

	app.Post("/api/messages", func(c *fiber.Ctx) error {
		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return err
		}

		pusherClient.Trigger("go-live-chat", "message", data)

		return c.JSON([]string{})
	})

	app.Listen(":8000")
}
