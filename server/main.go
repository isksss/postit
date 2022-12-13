package main

import (
	"log"

	"github.com/isksss/postit/server/database"
	"github.com/isksss/postit/server/model"
)

func main() {
	// database init
	log.Default().Print("Auto Migrate.")
	db := database.Migrate()
	log.Default().Print("End Migrate")
	if database.HasTable(db, &model.Users{}) {
		log.Println("User Table is Exists.")
	}
}
