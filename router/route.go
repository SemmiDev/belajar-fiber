package router

import (
	"LearnFiber/handler"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes (app *fiber.App) {

	// Middleware
	api := app.Group("/api")

	// routes
	api.Get("/", handler.GetAllProducts)
	api.Get("/:id", handler.GetSingleProduct)
	api.Post("/", handler.CreateProduct)
	api.Delete("/:id", handler.DeleteProduct)
	api.Put("/:id", handler.UpdateProduct)
}