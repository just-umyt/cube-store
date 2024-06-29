package cart

import (
	"github.com/gofiber/fiber/v2"
	"github.com/just-umyt/cube-store/internal/database"
	"github.com/just-umyt/cube-store/internal/models"
)

func Cart(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	//cart Products y Local dan alyas
	shopCart := c.Locals("shop_cart").(models.Cart)

	var cartProducts []models.CartProduct
	database.DB.Where("cart_id = ?", shopCart.ID).Find(&cartProducts)

	//Cartyn icindaki productlary DB dan alyas
	var products []models.Product
	for _, prd := range cartProducts {
		var product models.Product
		database.DB.First(&product, prd.ProductId)
		products = append(products, product)
	}
 
	return c.Render("cart", fiber.Map{
		"Name":     user.Name,
		"Products": products,
		"Carts":    cartProducts,
		"Cart":     len(cartProducts),
	})

}
