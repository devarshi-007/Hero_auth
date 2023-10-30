package routes

import (
	"github.com/devarshitrivedi01/hero_auth/controllers"
	"github.com/devarshitrivedi01/hero_auth/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Startup() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowCredentials: true,
	}))
	app.Get("/ping", controllers.Ping)
	app.Get("/demo", controllers.Get)
	app.Get("/bad", controllers.Message)
	app.Post("/login", controllers.Login)
	app.Use(func(c *fiber.Ctx) error {
		cookie := c.Cookies("session_id")
		valid := utils.Valid(cookie)
		if !valid {
			return c.Redirect("/bad")
		}
		return c.Next()
	})
	app.Get("/users", controllers.UserDetail)
	app.Listen(":4000")
}
