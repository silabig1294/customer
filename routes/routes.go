package routes

import (
	"github.com/silabig1294/customer/controllers"
	"github.com/gofiber/fiber/v2"
)
func Setup(app *fiber.App){

	api := app.Group("/api")

	// group user
	v1  := api.Group("v1/user")
	v1.Post("/register",controllers.Register)
	v1.Post("/login",controllers.Login)
	v1.Get("/welcome",controllers.User)
	v1.Delete("/logout",controllers.Logout) //or Post  
}