package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)
var(
	port string
)

func main() {
	err := godotenv.Load(".env")
    if err != nil {
        fmt.Println(err)
    }
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Post("/", func(c *fiber.Ctx) {
		// Get first file(image or video) from form field "document":
		file, err := c.FormFile("file")

		// Check for errors:
		if err == nil {
			// 👷 Save file to root directory:
			c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
			// 👷 Save file inside uploads folder under current working directory:
			c.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
			// 👷 Save file using a relative path:
			c.SaveFile(file, fmt.Sprintf("/tmp/uploads_relative/%s", file.Filename))
		}
	})

	// Start server
	log.Fatal(app.Listen(port))
}