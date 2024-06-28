package home

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/just-umyt/cube-store/internal/database"
	"github.com/just-umyt/cube-store/internal/models"
	"github.com/just-umyt/cube-store/pkg/middleware"
)

func Home(c *fiber.Ctx) error {
	// Cart y almaly
	cart, err := middleware.ParseCartToken(c.Cookies("shop_cart"))
	if err != nil {
		fmt.Println("ShopCart intak yok")
	}

	// cart, err := middleware.Cart(c)

	var cartProducts []models.CartProduct
	if cart.ID != 0 {
		database.DB.Where("cart_id = ?", cart.ID).Find(&cartProducts)
	}

	var user models.User
	userLocal := c.Locals("user")
	if userLocal != nil {
		user = userLocal.(models.User)
	}

	var products []models.Product
	var categorie []models.Category

	database.DB.Find(&products)
	database.DB.Find(&categorie)

	return c.Render("index", fiber.Map{
		"Name":       user.Name,
		"Categories": categorie,
		"Products":   products,
		"Cart":       len(cartProducts),
	})
}
