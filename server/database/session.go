package database

import "github.com/isksss/postit/server/model"

func IsExistsSession(userid uint) bool {
	db, _ := GetGormDb()

	var SessionKey model.SessionKey
	db.Where(model.SessionKey{UserId: userid}).Find(&SessionKey)

	if SessionKey.ID == userid {
		return true
	}

	return false

}
