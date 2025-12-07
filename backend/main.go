package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "POST",
	}))

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello Fiber!")
	// })

	// app.Listen(":3000")

	app.Post("/upload", func(c *fiber.Ctx) error {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			return c.Status(400).SendString("No file uploaded")
			// return fiber.NewError(782, "Custom error message")
		}

		err = c.SaveFile(fileHeader, "./uploads/"+fileHeader.Filename)
		if err != nil {
			return c.Status(500).SendString("failed to save file")
		}

		return c.SendString("File uploaded successfully")
	})
}
