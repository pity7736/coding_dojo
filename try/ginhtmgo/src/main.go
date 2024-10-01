package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type data struct {
	title string
}

func (d data) GetTitle() string {
	return d.title
}

func main() {
	engine := html.New("./src/templates", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/t", func(c *fiber.Ctx) error {
		// data := make(map[string]string, 1)
		// data["title"] = "whatever"
		data := data{title: "hola"}
		return c.Render("index", data)
	})

	app.Listen(":3000")
}
