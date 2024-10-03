package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

type data struct {
	Tasks []task
}

type task struct {
	Name string
}

func newData() *data {
	tasks := make([]task, 0, 5)
	return &data{Tasks: tasks}
}

func main() {
	engine := html.New("./src/templates", ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "base",
	})
	app.Use(logger.New())
	data := newData()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})
	app.Get("/tasks", func(c *fiber.Ctx) error {
		return c.Render("tasks", data)
	})

	app.Post("/tasks", func(c *fiber.Ctx) error {
		c.Status(201)
		task := task{
			Name: c.FormValue("name"),
		}
		data.Tasks = append(data.Tasks, task)
		if len(data.Tasks) == 1 {
			return c.Render("tasks", data, "")
		}
		return c.Render("task", task, "")
	})

	log.Fatal(app.Listen(":3000"))
}
