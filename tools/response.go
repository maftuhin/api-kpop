package tools

import "github.com/gofiber/fiber/v2"

func ErrorResponse(msg string, c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"message": msg,
	})
}
