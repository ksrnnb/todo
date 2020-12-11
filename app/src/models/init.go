package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// gorm.DB
var Db *gorm.DB

// initialize for database migration
func Initialize() {
	dsn := "user=root password=root dbname=todo host=psql port=5432 sslmode=disable"

	Db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	Db.AutoMigrate(Todo{}, Item{})
}