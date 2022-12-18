package database

import (
	"log"

	"github.com/isksss/postit/server/model"
)

func IsExistsSession(userid uint) bool {
	log.Print("is exists session...")
	db, _ := GetGormDb()

	var SessionKey model.SessionKey
	db.Where(model.SessionKey{UserId: userid}).Find(&SessionKey)
	log.Print("userid:", userid)
	log.Print("sessionkey.ID:", SessionKey.UserId)
	if SessionKey.UserId == userid {
		log.Print("あったよー！")
		return false
	}
	log.Print("なかったよー！")
	return true

}
