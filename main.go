package main

import (
	"github.com/gofiber/fiber"
	//"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	
	app.Static("/", "view/homepage.html")
	app.Listen(":"+port)
}