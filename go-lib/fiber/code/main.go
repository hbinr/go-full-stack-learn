package main

import (
	"github.com/gofiber/fiber/v2"
	"hb.study/go-lib/fiber/code/rest"
)

func main() {
	app := fiber.New()
	app.Get("/", rest.GetStr)
	app.Get("/json", rest.GetJSON)

	app.Listen(":3001")
}
