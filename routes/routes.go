package routes

import (
	"github.com/silabig1294/customer/controllers"
	"github.com/gofiber/fiber/v2"
)
func Setup(app *fiber.App){
	app.Get("/",controllers.Hello)
}