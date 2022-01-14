package database

import (
	"log"
	"time"

	"github.com/PedroPaiao/rest-api-go/database/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	str := "host=127.0.0.1 port=5432 dbname=books_development user=postgres password=postgres sslmode=disable"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		log.Fatal("ERROR!: ", err)
	}

	db = database

	migrations.RunMigrations(db)

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetConnMaxIdleTime(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
}

func GetDatabase() *gorm.DB {
	return db
}
