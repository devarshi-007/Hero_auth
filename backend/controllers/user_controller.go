package controllers

import (
	"fmt"
	"time"

	"github.com/devarshitrivedi01/hero_auth/config"
	"github.com/devarshitrivedi01/hero_auth/models"
	"github.com/devarshitrivedi01/hero_auth/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	jtoken "github.com/golang-jwt/jwt/v4"
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

func Ljwt(c *fiber.Ctx) error {
	var u models.User
	err := c.BodyParser(&u)
	if err != nil && len(u.Username) != 0 {
		panic(err)
	}
	exist := models.CheckUser(u)
	if !exist {
		return c.JSON(false)
	}
	
	claims := jtoken.MapClaims{
		"ID":  u.Username,
		"fav": u.Password,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})
	}
	// Return the token
	return c.JSON(t)
}

func Protected(c *fiber.Ctx) error {
	// Get the user from the context and return it
	
	tokenString := c.Get("Auth")
	if tokenString == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "No token provided",
        })
    }

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Add your secret key for token validation here
        return []byte(config.Secret), nil
    })

	if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "Invalid token",
        })
    }
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
        })
    }
	
    // Store the claims in the context for later use in controllers
    c.Locals("user", claims)
	
	user := c.Locals("user").(jwt.MapClaims)
	favPhrase := user["fav"].(string)
	return c.SendString("Welcome 👋  " + favPhrase)
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
