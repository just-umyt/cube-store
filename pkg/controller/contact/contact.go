package contact

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/just-umyt/cube-store/internal/models"
	"github.com/just-umyt/cube-store/pkg/middleware"
)

func Contact(c *fiber.Ctx) error {
	var user models.User
	userToken := c.Cookies("user")
	if userToken != "" {
		var err error
		user, err = middleware.ParseUserToken(userToken)
		if err != nil {
			fmt.Println(err)
		}
	}

	
	return c.Render("contact", fiber.Map{
		"Name": user.Name,
	})
}
