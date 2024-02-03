package db

// docker exec -it postgres-container bash -> interctive terminal with postgres container

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DNS = "host=localhost user=test password=changemeinprod! dbname=gorm port=5432"

var DB *gorm.DB // DB is a variable of type *gorm.DB (pointer to gorm.DB)

func DBConnection() {
	var error error // var error error -> error is a variable of type error
	DB, error = gorm.Open(
		postgres.Open(DNS), &gorm.Config{})
	// = is assignment operator := is declaration and assignment operator

	if error != nil {
		// panic(
		// 	"Could not connect to the database",
		// ) // panic is a function that stops the execution of the program

		log.Fatal("Could not connect to the database")
	} else {
		log.Println("Connected to the database")
	}
}
