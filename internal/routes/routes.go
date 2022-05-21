package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

// Initialize initializing all the routes for sidecar-proxy
func Initialize(app *fiber.App) {
	//api := app.Group("/api")

	// TODO: Version group will be handled later
	//api.Group("/v1")

	//I usually move methods inside their individual route file for example
	// for /company route I create company_route.go and inside it create 'company' group and group methods get,put,post etc
	//Also creating route handlers to the downstream app
	app.Get("/company",BasicSuccess)
	app.Get("/company/:id",BasicSuccess)
	app.Get("/company/account",BasicSuccess)
	app.Get("/account",BasicSuccess)
	app.Get("/account/:id",BasicSuccess)
	app.Get("/:id",BasicSuccess)
	app.Get("/account/:id/user",BasicSuccess)
	app.Get("/tenant/account/blocked",BasicSuccess)
}

func BasicSuccess(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("Request successful! - Path: %s",c.Path()))
}
