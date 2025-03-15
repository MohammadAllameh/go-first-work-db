package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	db "github.com/mohammadallameh/go-first-work-db/database"
	"github.com/mohammadallameh/go-first-work-db/models"
	"gorm.io/gorm"
)

func main() {
	// create new fiber app
	app := fiber.New()

	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to database
	db.ConnectDB()

	// auto migrate
	db.AutoMigrate()

	// seed database
	// seeders.SeedUser()

	// define routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/add-user", func(c *fiber.Ctx) error {
		user := models.User{
			Username: "super-admin",
			Password: "superadmin@123",
			Email:    "superadmin@gmail.com",
		}
		dbclient := db.GetDB()

		return dbclient.Create(&user).Error
	})
	app.Get("/delete-user/:id", func(c *fiber.Ctx) error {

		id := c.Params("id")

		userId, _ := strconv.Atoi(id)

		user := models.User{
			Model: gorm.Model{
				ID: uint(userId),
			},
		}

		dbclient := db.GetDB()
		// dbclient.Model(&user).Updates(nil)
		// dbclient.Model(&user).Transaction()

		return dbclient.Delete(&user).Error
	})

	// start server
	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
