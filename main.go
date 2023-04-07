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
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	database.Connect(database.Config{
		Host:     "db",            //viper.GetString("db.host"),
		Port:     "5432",          //viper.GetString("db.port"),
		Username: "postgres",      //viper.GetString("db.username"),
		DBName:   "assignment3go", //viper.GetString("db.dbname"),
		SSLMode:  "disable",       //viper.GetString("db.sslmode"),
		Password: "12345",         //os.Getenv("DB_PASSWORD"),
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
