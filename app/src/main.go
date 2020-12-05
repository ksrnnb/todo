package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

type Todo struct {
	Id   int    `gorm:"primaryKey"`
	Uuid string `gorm:"type:varchar(36);unique;not null"`
}

type Item struct {
	Id     int    `gorm:"primaryKey"`
	TodoId int    `gorm:"not null"`
	Name   string `gorm:"not null"`
	Done   bool   `gorm:"default:false;not null"`
}

func init() {
	dsn := "user=root password=root dbname=todo host=psql port=5432 sslmode=disable"
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Db.AutoMigrate(&Todo{}, &Item{})

	// test := Todo{
	// 	Uuid: "a272270a-34f7-11eb-a0cf-0242ac120003",
	// }

	// Db.Create(&test)

	// fmt.Println(test)
}

func main() {

	server := http.Server{
		Addr: ":8000",
	}

	http.HandleFunc("/", handleRequest)
	server.ListenAndServe()
}
