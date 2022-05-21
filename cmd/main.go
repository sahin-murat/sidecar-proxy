package main

import "github.com/gofiber/fiber/v2"

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

	err := sideCarApplication.Listen(":333")
	if err != nil {
		panic(err)
	}
}

func ValidatePath(path string) bool {
	// Some magic here...
	return false
}
