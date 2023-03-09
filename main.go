package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"Assignment3Go/database"
	"Assignment3Go/routes"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	database.Connect()

	//AllowCredentials is important because
	//front-end can get a cookie and send it back
	//if we want to authenticate using httponly cookies

	routes.Setup(app)

	app.Listen(":8000")
}
