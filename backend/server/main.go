package deliverd

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

// environment file, from project directory root.
const ENV_FILE = ".env"

// external function to start server
func StartServer() {
	godotenv.Load(ENV_FILE)

	_, err := get_db()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("* database connection went well")
	}

	app := fiber.New()

	app.Get("/auth/check", h_check_auth)                                    // Check auth header
	app.Get("/auth/keys/get_all", h_get_keys)                               // Get all keys
	app.Get("/auth/keys/get_owner/:owner_id", h_get_key_by_owner)           // Get keys by owner id
	app.Get("/auth/keys/status/:key_id/:status", h_update_key_status)       // Update key status
	app.Get("/class/get_all", h_get_all_classifications)                    // Get all classifications
	app.Get("/class/get/:classification_id", h_get_classification_by_id)    // Get classification by id
	app.Get("/courier/get_all", h_get_all_couriers)                         // Get all couriers
	app.Get("/courier/get/:courier_id", h_get_courier_by_id)                // Get courier by id
	app.Post("/courier/new", h_new_courier)                                 // Create a new courier
	app.Get("/courier/mark/:courier_id/active", h_mark_courier_active)      // Mark a courier as active
	app.Get("/courier/mark/:courier_id/leave", h_mark_courier_leave)        // Mark a courier as inactive
	app.Get("/courier/mark/:courier_id/left", h_mark_courier_left)          // Mark a courier as left
	app.Get("/partner/get_all", h_get_all_partners)                         // Get all partners
	app.Get("/partner/get/:partner_id", h_get_partner_by_id)                // Get partner by id
	app.Post("/partner/new", h_new_partner)                                 // Create a new partner
	app.Get("/partner/status/:partner_id/:status", h_update_partner_status) // Update partner status

	app.Get("/", func(c fiber.Ctx) error {
		c.Type("html")
		return c.SendString("<center><h2>deliverd</h2><p>a project by ccl, see <a href=\"https://github.com/goultyy/deliverd\">github</a>")
	})

	app.Listen(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}
