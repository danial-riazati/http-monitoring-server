package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON("hello")
	})

	if err := app.Listen("127.0.0.1:1307"); err != nil {
		log.Fatal(err)
	}
}
