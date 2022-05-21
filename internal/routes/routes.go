package routes

import "github.com/gofiber/fiber/v2"

// Initialize initializing all the routes for sidecar-proxy
func Initialize(app *fiber.App) {
	api := app.Group("/api")

	// Version group will be handled later
	api.Group("/v1")
}