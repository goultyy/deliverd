/*
	Frontend files for auth
*/

package deliverd

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func h_check_auth(c fiber.Ctx) error {
	err := validate_key(c)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":    "error",
			"given_key": c.Get("Authorization"),
			"error":     err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

func h_get_keys(c fiber.Ctx) error {
	err := validate_key_with_type(c, KEY_ADMIN)
	if err != nil {
		return handle_error_special(c, 401, err)
	}
	keys, err := get_all_keys()
	if err != nil {
		return handle_error(c, err)
	}
	return c.JSON(fiber.Map{
		"status": "ok",
		"keys":   keys,
	})
}

func h_get_key_by_owner(c fiber.Ctx) error {
	err := validate_key_with_type(c, KEY_ADMIN)
	if err != nil {
		return handle_error_special(c, 401, err)
	}
	owner_id, err := strconv.Atoi(c.Params("owner_id"))
	if err != nil {
		return handle_error_special(c, 400, err)
	}
	keys, err := get_keys_by_owner(owner_id)
	if err != nil {
		return handle_error(c, err)
	}
	return c.JSON(fiber.Map{
		"status": "ok",
		"keys":   keys,
	})
}

func h_update_key_status(c fiber.Ctx) error {
	err := validate_key_with_type(c, KEY_ADMIN)
	if err != nil {
		return handle_error_special(c, 401, err)
	}
	key_id, err := strconv.Atoi(c.Params("key_id"))
	if err != nil {
		return handle_error_special(c, 400, err)
	}
	status, err := strconv.Atoi(c.Params("status"))
	if err != nil {
		return handle_error_special(c, 400, err)
	}
	err = change_key_status(key_id, APIKeyStatus(status))
	if err != nil {
		return handle_error(c, err)
	}
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
