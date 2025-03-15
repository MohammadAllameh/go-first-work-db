package seeders

import (
	"log"

	db "github.com/mohammadallameh/go-first-work-db/database"
	"github.com/mohammadallameh/go-first-work-db/models"
)

func SeedUser() {
	dbClient := db.GetDB()

	// method 1
	// var user models.User
	// user.Username = "John Doe"
	// user.Email = "john@doe.com"
	// user.Password = "password"

	// method 2
	user := models.User{
		Username: "John Doe",
		Email:    "john@doe.com",
		Password: "password",
	}

	err := dbClient.Create(&user).Error
	if err != nil {
		log.Fatal("User cannat be added to the database")
	}
}
