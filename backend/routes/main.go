package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/devarshitrivedi01/hero_auth/controllers"
)

func Startup() {
	app := fiber.New()
	app.Get("/users", controllers.UserDetail)
	app.Listen(":3000")
}