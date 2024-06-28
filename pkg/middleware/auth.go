package middleware

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/just-umyt/cube-store/internal/database"
	"github.com/just-umyt/cube-store/internal/models"
)

func CreateUserToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseUserToken(tokenString string) (models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return models.Session{}, errors.New("unexpected signing method")
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return models.User{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return models.User{}, errors.New("session is expired")
		}

		var user models.User

		database.DB.First(&user, claims["sub"])

		return user, nil
	}

	return models.User{}, nil
}
