package middleware

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/just-umyt/cube-store/internal/database"
	"github.com/just-umyt/cube-store/internal/models"
)

func Cart(c *fiber.Ctx) error {

	var shopCart models.Cart

	shopCartCookie := c.Cookies("shop_cart")

	if shopCartCookie == "" {

		session := c.Locals("session").(models.Session)

		shopCart.SessionId = session.ID

		shopCartToken, err := CreateCartToken(shopCart)
		if err != nil {
			fmt.Println(err)
		}
		c.Cookie(&fiber.Cookie{
			Name:     "shop_cart",
			Value:    shopCartToken,
			Expires:  time.Now().Add(time.Hour * 24 * 30),
			HTTPOnly: true,
			// Secure:   true,
		})

		c.Locals("shop_cart", shopCart)
		return c.Next()
	} else {
		var err error
		shopCart, err = ParseCartToken(shopCartCookie)

		if err != nil {
			log.Fatal("Error", err)
		}

		c.Locals("shop_cart", shopCart)
		return c.Next()
	}
}

func CreateCartToken(shopCart models.Cart) (string, error) {
	database.DB.Create(&shopCart)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": shopCart.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func ParseCartToken(tokenString string) (models.Cart, error) {
	//Simple parsing
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return models.Session{}, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return models.Cart{}, err
	}

	//Claimden cart a salnan productlary almaly
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return models.Cart{}, errors.New("session is expired")
		}

		var cart models.Cart

		database.DB.First(&cart, claims["sub"])

		return cart, nil
	}

	return models.Cart{}, nil

}
