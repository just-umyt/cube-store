package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/just-umyt/cube-store/handlers"
	"github.com/just-umyt/cube-store/internal/database"
	"github.com/just-umyt/cube-store/internal/env"
)

func init() {
	env.LoadEnv()
	database.InitDatabase()

	// category1 := models.Category{
	// 	Name: "2x2 cubes", Products: []models.Product{
	// 		{Name: "1st 2x2", About: "lorem", Price: 2.3, Bought: 3},
	// 		{Name: "2nd 2x2", About: "lorem", Price: 21.3, Bought: 2},
	// 		{Name: "3rd 2x2", About: "lorem", Price: 3.4, Bought: 1},
	// 		{Name: "4th 2x2", About: "lorem", Price: 6.7, Bought: 0},
	// 	},
	// }
	// database.DB.Create(&category1)

	// category2 := models.Category{
	// 	Name: "3x3 cubes", Products: []models.Product{
	// 		{Name: "1st 3x3", About: "lorem", Price: 3.3, Bought: 3},
	// 		{Name: "2nd 3x3", About: "lorem", Price: 31.3, Bought: 3},
	// 		{Name: "3rd 3x3", About: "lorem", Price: 3.4, Bought: 1},
	// 		{Name: "4th 3x3", About: "lorem", Price: 6.7, Bought: 0},
	// 	},
	// }

	// database.DB.Create(&category2)

	// category3 := models.Category{
	// 	Name: "4x4 cubes", Products: []models.Product{
	// 		{Name: "1st 4x4", About: "lorem", Price: 3.3, Bought: 3},
	// 		{Name: "2nd 4x4", About: "lorem", Price: 31.3, Bought: 3},
	// 		{Name: "3rd 4x4", About: "lorem", Price: 3.4, Bought: 1},
	// 		{Name: "4th 4x4", About: "lorem", Price: 6.7, Bought: 0},
	// 	},
	// }

	// database.DB.Create(&category3)
	// category4 := models.Category{
	// 	Name: "5x5", Products: []models.Product{
	// 		{Name: "1st 5x5", About: "lorem", Price: 3.3, Bought: 3},
	// 		{Name: "2nd 5x5", About: "lorem", Price: 31.3, Bought: 3},
	// 		{Name: "3rd 5x5", About: "lorem", Price: 3.4, Bought: 1},
	// 		{Name: "4th 5x5", About: "lorem", Price: 6.7, Bought: 0},
	// 	},
	// }

	// database.DB.Create(&category4)

}

func main() {

	// html files
	engine := html.New("./internal/common/views", ".html")

	//New fiber router
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//Static files
	app.Static("/", "./internal/common/assets")

	//Handlers
	handlers.Handlers(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
