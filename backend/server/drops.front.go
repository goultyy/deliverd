package deliverd

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func h_get_all_classifications(c fiber.Ctx) error {
	classifc, err := get_all_classifications()
	if err != nil {
		return handle_error(c, err)
	}
	return c.JSON(fiber.Map{
		"status":          "ok",
		"classifications": classifc,
	})
}

func h_get_classification_by_id(c fiber.Ctx) error {
	classification_id, err := strconv.Atoi(c.Params("classification_id"))
	if err != nil {
		return handle_error(c, err)
	}
	classification, err := get_classification_by_id(classification_id)
	if err != nil {
		return handle_error(c, err)
	}
	return c.JSON(fiber.Map{
		"status":         "ok",
		"classification": classification,
	})
}
