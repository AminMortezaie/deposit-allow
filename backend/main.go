package main

import (
	"backend/entity"
	pg "backend/repository/db/postgres"
	"log"
)

func main() {
	db, err := pg.ConnectToPostgreSQL()
	if err != nil {
		log.Fatal(err)
	}

	// Perform database migration
	err = db.AutoMigrate(&entity.Wallet{})
	err = db.AutoMigrate(&entity.Transaction{})
	err = db.AutoMigrate(&entity.SubTransaction{})

	if err != nil {
		log.Fatal(err)
	}

	// Your CRUD operations go here
}
