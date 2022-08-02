package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var database *gorm.DB

func Connect() {
	d, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=AnjomanShits password=09116903138a sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	database = d
}

func GetDB() *gorm.DB {
	return database
}
