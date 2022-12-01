package database

import (
	"github.com/escalopa/goblog/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open db")
	}
	log.Println("Opened DB successfully")
}

func Migrate() {
	err := Instance.AutoMigrate(&entities.Post{})
	if err != nil {
		log.Fatal("Failed to migrate db")
	}
	log.Println("DB migrated successfully")
}
