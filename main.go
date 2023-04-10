package main

import (
	"Assignment3Go/database"
	"Assignment3Go/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	"log"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	//app.Use(cors.New())
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	database.Connect(database.Config{
		Host:     "localhost", //localhost - for local, db for docker         //viper.GetString("db.host"),
		Port:     "5432",
		Username: "postgres",
		DBName:   "assignment3go",
		SSLMode:  "disable",
		Password: "12345",
	})

	//AllowCredentials is important because
	//front-end can get a cookie and send it back
	//if we want to authenticate using httponly cookies

	routes.Setup(app)

	app.Listen(":8000")
}
func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
