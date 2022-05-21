package main

import (
	"github.com/gofiber/fiber/v2"
	"sidecar-proxy/global"
	"sidecar-proxy/internal/routes"
)

func main() {
	sideCarApplication := fiber.New()

	//Middleware that checks if the http request has allowed path
	//If not allowed returns 406 - Not Acceptable status
	sideCarApplication.Use(func(c *fiber.Ctx) error {
		if !ValidatePath(c.Path()){
			return c.SendStatus(406)
		}

		return c.Next()
	})

	routes.Initialize(sideCarApplication)

	err := sideCarApplication.Listen(":3333")
	if err != nil {

		panic(err)
	}
}

func ValidatePath(path string) bool {

	//Check if given path is allowed, if yes  return true
	for _, allowedPath := range global.AllowedRoutes {
		if _,ok := allowedPath.Match(path); ok {
			return true
		}
	}

	return false
}
