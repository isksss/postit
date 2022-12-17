package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/isksss/postit/server/database"
	"github.com/isksss/postit/server/model"
)

func Login(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		//パースしてエラーが起きた時
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// session key
	var SessionKey model.SessionKey
	user := database.CheckEmail(u.Email) // ユーザ情報をゲット

	// セッションがあるか確認
	db, _ := database.GetGormDb()
	if database.IsExistsSession(user.ID) {
		db.Delete(&SessionKey) //セッション削除
	}

	// Session作成
	key, _ := uuid.NewRandom()
	SessionKey.SessionKey = key.String()

}
