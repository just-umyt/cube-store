package auth

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
	return c.Render("login", nil)
}

func Register(c *fiber.Ctx) error {
	return c.Render("register", nil)
}
