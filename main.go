package main

import (
	"LearnFiber/database"
	"LearnFiber/router"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	router.SetupRoutes(app)

	app.Listen(":3000")
}