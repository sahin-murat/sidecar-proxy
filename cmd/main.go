package main

import (
	"github.com/gofiber/fiber/v2"
	"sidecar-proxy/global"
	"sidecar-proxy/internal/routes"
)

func main() {
	sideCarApplication := fiber.New()


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
