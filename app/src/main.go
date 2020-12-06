package main

import (
	"models"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// gorm.DB
var db *gorm.DB

// initialize for database migration
func init() {
	dsn := "user=root password=root dbname=todo host=psql port=5432 sslmode=disable"

	db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.Todo{}, &models.Item{})
}

func main() {
	server := http.Server{
		Addr: ":8000",
	}

	http.HandleFunc("/", handleRequest)
	server.ListenAndServe()
}
