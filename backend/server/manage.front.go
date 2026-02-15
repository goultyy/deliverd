package deliverd

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/jinzhu/copier"
)

// return all couriers
func h_get_all_couriers(c fiber.Ctx) error {
	err := validate_key_with_type(c, KEY_ADMIN)
	if err != nil {
		return handle_error_special(c, 401, err)
	}
	couriers, err := get_all_couriers()
	if err != nil {
		return handle_error(c, err)
	}
	return c.JSON(fiber.Map{
		"status":   "ok",
		"couriers": couriers,
	})
}

// get courier by id
func h_get_courier_by_id(c fiber.Ctx) error {
	err := validate_key_with_type(c, KEY_ADMIN)
	if err != nil {
		return handle_error_special(c, 401, err)
	}
	courier_id, err := strconv.Atoi(c.Params("courier_id"))
	if err != nil {
		return handle_error(c, err)
	}
	courier, err := get_courier_by_id(courier_id)
	if err != nil {
		return handle_error(c, err)
	}
	return c.JSON(fiber.Map{
		"status":  "ok",
		"courier": courier,
	})
}

// create a courier
func h_new_courier(c fiber.Ctx) error {
	err := validate_key_with_type(c, KEY_ADMIN)
	if err != nil {
		return handle_error_special(c, 401, err)
	}
	var courier B_Courier
	err = c.Bind().JSON(&courier)
	if err != nil {
		return handle_error(c, err)
	}
	if err := validator.New().Struct(courier); err != nil {
		return handle_error(c, err)
	}
	var new Courier
	copier.Copy(&new, &courier)
	err = new.new_courier()
	if err != nil {
		return handle_error(c, err)
	}
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

// changing courier type also disables/enables any previous api keys
// mark courier on leave
func h_mark_courier_leave(c fiber.Ctx) error {
	err := validate_key_with_type(c, KEY_ADMIN)
	if err != nil {
		return handle_error_special(c, 401, err)
	}
	courier_id, err := strconv.Atoi(c.Params("courier_id"))
	if err != nil {
		return handle_error(c, err)
	}
	courier, err := get_courier_by_id(courier_id)
	if err != nil {
		return handle_error(c, err)
	}
	err = courier.mark_leave()
	if err != nil {
		return handle_error(c, err)
	}
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

// mark courier active
func h_mark_courier_active(c fiber.Ctx) error {
	err := validate_key_with_type(c, KEY_ADMIN)
	if err != nil {
		return handle_error_special(c, 401, err)
	}
	courier_id, err := strconv.Atoi(c.Params("courier_id"))
	if err != nil {
		return handle_error(c, err)
	}
	courier, err := get_courier_by_id(courier_id)
	if err != nil {
		return handle_error(c, err)
	}
	err = courier.mark_active()
	if err != nil {
		return handle_error(c, err)
	}
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

// mark courier perma left
func h_mark_courier_left(c fiber.Ctx) error {
	err := validate_key_with_type(c, KEY_ADMIN)
	if err != nil {
		return handle_error_special(c, 401, err)
	}
	courier_id, err := strconv.Atoi(c.Params("courier_id"))
	if err != nil {
		return handle_error(c, err)
	}
	courier, err := get_courier_by_id(courier_id)
	if err != nil {
		return handle_error(c, err)
	}
	err = courier.mark_left()
	if err != nil {
		return handle_error(c, err)
	}
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

// get all partners
func h_get_all_partners(c fiber.Ctx) error {
	err := validate_key_with_type(c, KEY_ADMIN)
	if err != nil {
		return handle_error_special(c, 401, err)
	}
	partners, err := get_all_partners()
	if err != nil {
		return handle_error(c, err)
	}
	return c.JSON(fiber.Map{
		"status":   "ok",
		"partners": partners,
	})
}

// get partner by id
func h_get_partner_by_id(c fiber.Ctx) error {
	err := validate_key_with_type(c, KEY_ADMIN)
	if err != nil {
		return handle_error_special(c, 401, err)
	}
	partner_id, err := strconv.Atoi(c.Params("partner_id"))
	if err != nil {
		return handle_error(c, err)
	}
	partner, err := get_partner_by_id(partner_id)
	if err != nil {
		return handle_error(c, err)
	}
	return c.JSON(fiber.Map{
		"status":  "ok",
		"partner": partner,
	})
}

// create partner
func h_new_partner(c fiber.Ctx) error {
	err := validate_key_with_type(c, KEY_ADMIN)
	if err != nil {
		return handle_error_special(c, 401, err)
	}
	var partner B_Partner
	err = c.Bind().JSON(&partner)
	if err != nil {
		return handle_error(c, err)
	}
	if err := validator.New().Struct(partner); err != nil {
		return handle_error(c, err)
	}
	var new Partner
	copier.Copy(&new, &partner)
	err = new.new_partner()
	if err != nil {
		return handle_error(c, err)
	}
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

// change partner status, this will also change any related API keys
func h_update_partner_status(c fiber.Ctx) error {
	err := validate_key_with_type(c, KEY_ADMIN)
	if err != nil {
		return handle_error_special(c, 401, err)
	}
	partner_id, err := strconv.Atoi(c.Params("partner_id"))
	if err != nil {
		return handle_error(c, err)
	}
	status_int, err := strconv.Atoi(c.Params("status"))
	if err != nil {
		return handle_error(c, err)
	}
	err = change_partner_status(partner_id, PartnerStatus(status_int))
	if err != nil {
		return handle_error(c, err)
	}
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
