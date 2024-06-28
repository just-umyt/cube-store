package session

import (
	"fmt"

	"github.com/just-umyt/cube-store/internal/database"
	"github.com/just-umyt/cube-store/pkg/middleware"
)

func Update(sessionToken string, id *uint) {
	session, err := middleware.ParseSessionToken(sessionToken)
	if err != nil {
		fmt.Println(err)
	}

	database.DB.First(&session)
	session.UserId = id
	database.DB.Save(&session)
}
