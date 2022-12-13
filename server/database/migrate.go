package database

import (
	"log"

	"github.com/isksss/postit/server/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// データベースMigration
func Migrate() *gorm.DB {
	//todo: DB情報などを環境変数から持ってくる
	dsn := "user:mysql@tcp(db:3306)/postit"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("")
		panic(err)
	}

	//AutoMigrate
	log.Print("start auto migrate")
	log.Print("=== Users")
	if err := db.AutoMigrate(&model.Users{}); err != nil {
		panic(err)
	}
	log.Print("=== Users End ===")
	// log.Print("=== Postit")
	// if err := db.AutoMigrate(&model.Postit{}); err != nil {
	// 	panic(err)
	// }
	// log.Print("=== Postit End ===")
	// log.Print("=== Key")
	// if err := db.AutoMigrate(&model.Key{}); err != nil {
	// 	panic(err)
	// }
	// log.Print("=== Key End ===")

	return db
}
