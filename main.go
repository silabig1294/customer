package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	db "github.com/silabig1294/customer/database"
	"github.com/silabig1294/customer/routes"
)

func main() {
	db.Connect()

	app := fiber.New()
	app.Use(cors.New(
		cors.Config{
		 	AllowCredentials: true,
		},
		))
	routes.Setup(app)
	// app.Get("/",func(c *fiber.Ctx)error{
	// 	return c.SendString("Hello Work!!!")
	// })
	app.Listen(":8050")
}
