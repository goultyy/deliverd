package deliverd

import "github.com/gofiber/fiber/v3"

func handle_error(c fiber.Ctx, err error) error {
	return c.JSON(fiber.Map{
		"status": "error",
		"error":  err.Error(),
	})
}

func handle_error_special(c fiber.Ctx, code int, err error) error {
	return c.Status(code).JSON(fiber.Map{
		"status": "error",
		"error":  err.Error(),
	})
}
