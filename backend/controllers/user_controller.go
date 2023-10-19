package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/devarshitrivedi01/hero_auth/models"
)

func UserDetail(c *fiber.Ctx) error{
	users := models.GetUserDetail()
	return c.JSON(users)
}