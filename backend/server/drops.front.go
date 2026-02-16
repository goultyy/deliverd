package deliverd

import (
	"errors"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/jinzhu/copier"
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

// get packages by status (admin, partner)
func h_get_packages_by_status(c fiber.Ctx) error {
	err := validate_keys_with_type(c, []APIKeyType{KEY_ADMIN, KEY_PARTNER})
	if err != nil {
		return handle_error_special(c, 401, err)
	}
	status, err := strconv.Atoi(c.Params("status"))
	packages, err := get_packages_by_status(UpdateID(status))
	if err != nil {
		return handle_error(c, err)
	}
	return c.JSON(fiber.Map{
		"packages": packages,
		"status":   "ok",
	})
}

// get package (requires validating postcode)
func h_get_package_by_id(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return handle_error(c, err)
	}
	packages, err := get_package_by_id(id)
	if err != nil {
		return handle_error(c, err)
	}
	if packages.DestinationPostcode != c.Query("postcode") {
		return handle_error(c, errors.New("invalid postcode"))
	}
	return c.JSON(fiber.Map{
		"package": packages,
		"status":  "ok",
	})
}

// create new package
func h_new_package(c fiber.Ctx) error {
	err := validate_keys_with_type(c, []APIKeyType{KEY_ADMIN, KEY_COURIER})
	if err != nil {
		return handle_error_special(c, 401, err)
	}
	var packages B_Package
	err = c.Bind().JSON(&packages)
	if err != nil {
		return handle_error(c, err)
	}
	if err := validator.New().Struct(packages); err != nil {
		return handle_error(c, err)
	}
	var new Packages
	copier.Copy(&new, &packages)
	err = new.new_package()
	if err != nil {
		return handle_error(c, err)
	}
	return c.JSON(fiber.Map{
		"status": "ok",
	})

}
