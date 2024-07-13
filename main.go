package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"slices"
	"strings"
)

var data = make([]string, 0)

type Body struct {
	Name string `json:"name"`
}

func removeElement(slice []string, item string) []string {
	for i, value := range slice {
		if value == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}

	return slice
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		return c.SendString("Hello, " + c.Params("name"))
	})

	app.Post("/register", func(c *fiber.Ctx) error {
		p := new(Body)

		if err := c.BodyParser(p); err != nil {
			return err
		}

		name := p.Name

		fmt.Println("Request to register " + name)
		data = append(data, name)
		return c.SendString("Registered " + name + ". Hello, " + strings.Join(data, ", ") + ".")
	})

	app.Delete("/deregister", func(c *fiber.Ctx) error {
		p := new(Body)

		if err := c.BodyParser(p); err != nil {
			return err
		}

		name := p.Name

		fmt.Println("Request to deregister " + name)

		if !slices.Contains(data, name) {
			return c.SendString("404: Cannot deregister " + name)
		}

		removeElement(data, name)
		return c.SendString("Deregistered " + name + ". Hello, " + strings.Join(data, ", ") + ".")
	})

	err := app.Listen(":3000")

	if err != nil {
		fmt.Println("Error:", err)
	}
}
