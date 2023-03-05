package main

import (
	db"github.com/silabig1294/customer/database"
	"github.com/silabig1294/customer/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Connect()
	app := fiber.New()

	routes.Setup(app)
	// app.Get("/",func(c *fiber.Ctx)error{
	// 	return c.SendString("Hello Work!!!")
	// })

	app.Listen(":8050")
}
