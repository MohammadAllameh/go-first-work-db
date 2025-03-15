package seeders

import "github.com/mohammadallameh/go-first-work-db/models"

func SeedUser() {
	User := models.User{
		Name:     "John Doe",
		Email:    "john@doe.com",
		Password: "password",
	}
}
