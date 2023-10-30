package controllers

import (
	"fmt"
	"time"

	"github.com/devarshitrivedi01/hero_auth/models"
	"github.com/devarshitrivedi01/hero_auth/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UserDetail(c *fiber.Ctx) error {
	users := models.GetUserDetail()
	return c.JSON(users)
}

func Message(c *fiber.Ctx) error {
	return c.JSON("Please Logged in")
}

func Login(c *fiber.Ctx) error {
	var u models.User
	err := c.BodyParser(&u)
	if err != nil && len(u.Username) != 0 {
		panic(err)
	}
	exist := models.CheckUser(u)
	if !exist {
		return c.JSON(false)
	}
	var s utils.SessionDetail

	sID := uuid.New()
	expiry := time.Now().Add(3 * time.Hour)
	cookie := new(fiber.Cookie)
	cookie.Name = "session_id"
	cookie.Value = sID.String()
	cookie.Expires = expiry
	cookie.HTTPOnly = false
	cookie.Secure = false
	cookie.SameSite = "None"
	cookie.Domain = "localhost"

	s.Session_id = sID.String()
	s.Username = u.Username
	s.Password = u.Password
	s.Expires = expiry
	c.Cookie(cookie)
	err = models.AddSession(s)
	
	if err != nil {
		panic(err)
	}


	return c.JSON(true)
}

func Ping(ctx *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "demo",
		Value:    "abcd",
		Expires:  time.Now().Add(time.Hour * 10),
		SameSite: "None",
		Domain:   "localhost",
		HTTPOnly: false,
		Secure:   false,
	}

	ctx.Cookie(&cookie)
	return ctx.SendString("done")
}

func Get(ctx *fiber.Ctx) error {
	fmt.Println(ctx.Cookies("demo"))
	return ctx.SendString("dive!")
}