package middleware

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/just-umyt/cube-store/internal/database"
	"github.com/just-umyt/cube-store/internal/models"
)

func Session(c *fiber.Ctx) error {
	var session models.Session
	sessionCookie := c.Cookies("session_id")
	if sessionCookie == "" {
		sessionToken, err := CreatSessionToken(session)
		if err != nil {
			fmt.Println(err)
		}

		c.Cookie(&fiber.Cookie{
			Name:     "session_id",
			Value:    sessionToken,
			Expires:  time.Now().Add(time.Hour * 24 * 30),
			HTTPOnly: true,
			// Secure:   true,
		})

		var user models.User

		database.DB.First(&user, session.UserId)

		c.Locals("user", user)
		c.Locals("session", session)
		return c.Next()

	} else {
		var err error
		session, err = ParseSessionToken(sessionCookie)
		if err != nil {
			fmt.Println("Failed to parse session token")
		}

		var user models.User

		database.DB.First(&user, session.UserId)

		c.Locals("user", user)
		c.Locals("session", session)
		return c.Next()
	}

}

func CreatSessionToken(session models.Session) (string, error) {

	database.DB.Create(&session)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": session.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseSessionToken(tokenString string) (models.Session, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return models.Session{}, errors.New("unexpected signing method")
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return models.Session{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return models.Session{}, errors.New("session is expired")
		}

		var session models.Session

		database.DB.First(&session, claims["sub"])

		return session, nil
	}

	return models.Session{}, nil
}
