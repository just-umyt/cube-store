package cartServices

import (
	"github.com/gofiber/fiber/v2"
	"github.com/just-umyt/cube-store/internal/database"
	"github.com/just-umyt/cube-store/internal/models"
)

func CartAdd(c *fiber.Ctx) error {
	//id syny almaly
	productId := c.Params("id")

	//producty DB dan almaly
	var product models.Product
	database.DB.First(&product, productId)

	//cookiedaki shop_carty parse etmeli
	shopCart := c.Locals("shop_cart").(models.Cart)

	//Taze cartProduct
	var cartProduct models.CartProduct
	cartProduct.ProductId = product.ID
	cartProduct.CartId = shopCart.ID

	//On yok bolsa taze product / bar bolsa Quantity++
	if database.DB.Where("cart_id = ? AND product_id = ?", cartProduct.CartId, cartProduct.ProductId).First(&cartProduct); cartProduct.ID == 0 {
		database.DB.Create(&cartProduct)
	} else {
		database.DB.First(&cartProduct)
		cartProduct.Quantity++
		database.DB.Save(&cartProduct)
	}

	c.Locals("shop_cart", shopCart)
	return c.Next()
}
