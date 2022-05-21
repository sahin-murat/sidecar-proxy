package main

import (
	"github.com/gofiber/fiber/v2"
	"sidecar-proxy/internal/routes"
)

var allowedList []string

func main() {
	allowedList = []string{
		"/company/",
		"/company/{id}",
		"/company/account",
		"/account",
		"/account/{id}",
		"/{id}",
		"/account/{id}/user",
		"/tenant/account/blocked",
	}

	sideCarApplication := fiber.New()

	routes.Initialize(sideCarApplication)

	err := sideCarApplication.Listen(":3333")
	if err != nil {

		panic(err)
	}
}

func ValidatePath(path string) bool {
	// Some magic here...
	return false
}
