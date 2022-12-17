package database

import (
	"github.com/isksss/postit/server/model"
	"gorm.io/gorm"
)

// ユーザ登録処理
func Register(u *model.User) *gorm.DB {
	//todo: エラーチェック
	db, _ := GetGormDb()
	r := db.Create(u)
	return r
}

// 既存登録者の情報をemailから持ってくる
func CheckEmail(email string) *model.User {
	db, _ := GetGormDb()

	u := model.User{}

	db.Where(&model.User{Email: email}).Last(&u)

	return &u
}
