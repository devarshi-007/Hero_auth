package routes

import (
	"fmt"

	"github.com/devarshitrivedi01/hero_auth/controllers"
	"github.com/devarshitrivedi01/hero_auth/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Startup() {
	app := fiber.New()
	app.Get("/bad", controllers.Message)
	app.Post("/login", controllers.Login)
	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		// var u models.User
		cookie := c.Cookies("session_id")
		valid := utils.Valid(cookie)
		fmt.Println(valid)
		if valid {
			return c.Next()
		}
		return c.Redirect("/bad")
	})
	app.Get("/users", controllers.UserDetail)
	app.Listen(":4000")
}
