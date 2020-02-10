package main

import (
	"log"

	"github.com/jinzhu/gorm"
	// gorm postgres connection
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	dbSource := "user=postgres password=postgres dbname=postgres sslmode=disable host=localhost"
	db, err := gorm.Open("postgres", dbSource)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(db)
}
