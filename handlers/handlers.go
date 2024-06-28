package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/just-umyt/cube-store/pkg/controller/auth"
	"github.com/just-umyt/cube-store/pkg/controller/cart"
	"github.com/just-umyt/cube-store/pkg/controller/contact"
	"github.com/just-umyt/cube-store/pkg/controller/home"
	"github.com/just-umyt/cube-store/pkg/middleware"
	cartServices "github.com/just-umyt/cube-store/pkg/services/cartSevices"
	"github.com/just-umyt/cube-store/pkg/services/user"
)

func Handlers(app *fiber.App) {
	//home page
	app.Get("/", middleware.Session, home.Home)
	app.Get("/add/:id", middleware.Cart, cartServices.CartAdd, home.Home)

	//Cart Page
	cartApi := app.Group("/cart", middleware.Cart)
	cartApi.Get("/", cart.Cart)
	cartApi.Get("/add/:id", cartServices.CartAdd, cart.Cart)
	cartApi.Get("/less/:id", cartServices.CartLess, cart.Cart)
	cartApi.Get("/delete/:id", cartServices.CartDelete, cart.Cart)

	//Auth
	app.Get("/login", auth.Login)
	app.Get("/register", auth.Register)
	app.Post("/login", user.Login, home.Home)
	app.Post("/register", user.Register, home.Home)

	//contact
	app.Get("/contact", contact.Contact)
}
