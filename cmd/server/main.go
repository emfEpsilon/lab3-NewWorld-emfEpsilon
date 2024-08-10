package main

import "github.com/gofiber/fiber/v2"

func main(){
	// Create new fiber server
	app := fiber.New()

	app.Get("/",func(c *fiber.Ctx) error {
		c.SendString("Hello, World!")
		return nil
	})

	app.Listen(":3000")
}