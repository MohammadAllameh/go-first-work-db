package db

import (
	"log"
	"os"

	"github.com/mohammadallameh/go-first-work-db/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func ConnectDB() {
	var err error

	dns := os.Getenv("DB_URL")
	// connect to database
	dbClient, err = gorm.Open(mysql.Open(dns), &gorm.Config{}) // تغییر این خط
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	sqlDB, _ := dbClient.DB()
	err = sqlDB.Ping()
	if err != nil {
		dbClose()
		log.Fatal("Failed to ping database")
	}

	// log success message
	log.Println("Connected to database")
}

func AutoMigrate() {
	dbClient.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	return dbClient
}

func dbClose() {
	sqldb, _ := dbClient.DB()
	sqldb.Close()
}
