package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		data := data{title: "hola"}
		return c.Render("base", data)
	})

	app.Get("/contacts", func(c *fiber.Ctx) error {
		data := data{title: "hola"}
		return c.Render("base", data)
	})

	app.Post("/contacts", func(c *fiber.Ctx) error {
		c.Status(201)
		data := data{title: "aoeuaoeu"}
		return c.Render("base", data)
	})

	log.Fatal(app.Listen(":3000"))
}
