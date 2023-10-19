package main

import (
	_ "github.com/lib/pq"
	"github.com/gofiber/fiber/v2"
	"github.com/devarshitrivedi01/hero_auth/controllers"
)


func main() {
	app := fiber.New()
	app.Get("/users", controllers.UserDetail)
	app.Listen(":3000")
}
