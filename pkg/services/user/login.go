package user

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/just-umyt/cube-store/internal/database"
	"github.com/just-umyt/cube-store/internal/models"
	"github.com/just-umyt/cube-store/pkg/middleware"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	//userden gelyan infolary parsit etmek ucin struct
	var body struct {
		Email    string
		Password string
	}

	//so Struct a parsit edyas we barlayas
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Failed to read body")
	}

	var user models.User
	database.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		return c.Status(http.StatusBadRequest).SendString("Invalid email")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid password")
	}

	tokenStr, err := middleware.CreateUserToken(user)
	if err != nil {

		return c.Status(http.StatusBadRequest).SendString("Failed to creat token")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "user",
		Value:    tokenStr,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HTTPOnly: true,
	})

	c.Locals("user", user)

	return c.Next()
}
