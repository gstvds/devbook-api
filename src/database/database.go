package database

import (
	"api/src/models"
	"api/src/utils/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func Setup() {
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.DATABASE.Host,
		config.DATABASE.User,
		config.DATABASE.Password,
		config.DATABASE.Name,
		config.DATABASE.Port,
	)

	openDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN:  dsn,
	}))

	if err != nil {
		fmt.Println("Failed to open db connection")
		log.Fatal(err)
	}

	if err = openDB.AutoMigrate(&models.User{}, &models.Follower{}); err != nil {
		fmt.Println("Failed to auto migrate User model")
		log.Fatal(err)
	}

	db = openDB
}

func GetDB() *gorm.DB {
	return db
}
