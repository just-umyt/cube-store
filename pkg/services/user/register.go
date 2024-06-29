package user

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/just-umyt/cube-store/internal/database"
	"github.com/just-umyt/cube-store/internal/models"
	"github.com/just-umyt/cube-store/pkg/middleware"
	"github.com/just-umyt/cube-store/pkg/services/session"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	//userden gelyan infolary parsit etmek ucin struct
	var body struct {
		Name     string
		Phone    string
		Email    string
		Adress   string
		Password string
	}

	//so Struct a parsit edyas we barlayas
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Failed to read body")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Failed to hash password")
	}

	//DB ucin taze user kopiyasy doretmek we taze userin infolaryny icine salmak
	user := models.User{Name: body.Name, Phone: body.Phone, Email: body.Email, Password: string(hash)}

	//DB a taze usery gosmak we barlamak
	result := database.DB.Create(&user)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).SendString("Email already sign up")
	}

	database.DB.Where("email = ?", user.Email).First(&user)

	sessionToken := c.Cookies("session_id")
	session.Update(sessionToken, &user.ID)

	tokenStr, err := middleware.CreateUserToken(user)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Failed to creat token")
	}

	//Taze usere 30 gunlik cookie bermek
	c.Cookie(&fiber.Cookie{
		Name:     "user",
		Value:    tokenStr,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HTTPOnly: true,
		// Secure:   true,
	})

	c.Locals("user", user)
	return c.Next()
}
