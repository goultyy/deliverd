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
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}
