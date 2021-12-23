package rest

import "github.com/gofiber/fiber/v2"

func GetStr(c *fiber.Ctx) error {
	return c.SendString("hello world")
}

func GetJSON(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": "hello world",
	})
}
