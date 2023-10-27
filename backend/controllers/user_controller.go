package controllers

import (
	"time"
	"fmt"
	"github.com/devarshitrivedi01/hero_auth/models"
	"github.com/devarshitrivedi01/hero_auth/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UserDetail(c *fiber.Ctx) error{
	users := models.GetUserDetail()
	return c.JSON(users)
}

func Message(c *fiber.Ctx) error{
	return c.JSON("Please Logged in")
}

func Login(c *fiber.Ctx) error {
	fmt.Println("Ahoya")
	var u models.User
	fmt.Println("Hello, ")
	err := c.BodyParser(&u)
	fmt.Println(u)
	if err != nil && len(u.Username) != 0 {
		panic(err)
	}
	exist := models.CheckUser(u)
	if exist {
		sID := uuid.New()
		expiry := time.Now().Add(3 * time.Minute)
		cookie := new(fiber.Cookie)
		cookie.Name = "session_id"
		cookie.Value = sID.String()
		cookie.Expires = expiry

		c.Cookie(cookie)
		var s utils.SessionDetail
		s.Session_id = sID.String()
		s.Username = u.Username
		s.Password = u.Password
		s.Expires = expiry
		err := models.AddSession(s)
		if err!= nil {
			panic(err)
		}
		return c.JSON(true)
	}
	return c.JSON(false)
}