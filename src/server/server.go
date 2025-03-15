package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mohammadallameh/go-first-work-db/seeders"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func main() {
	// create new fiber app
	app := fiber.New()

	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dns := os.Getenv("DB_URL")

	// connect to database
	dbClient, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	sqlDB, _ := dbClient.DB()
	defer sqlDB.Close()
	if err != nil {
		panic("failed to connect database")
	}

	// log success message
	log.Println("Connected to database")

	// define routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// seed database
	seeders.SeedUser()

	// start server
	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
