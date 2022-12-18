package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/isksss/postit/server/database"
	"github.com/isksss/postit/server/model"
)

func Login(c *gin.Context) {
	/*
	 * {"email":"email@email.com", "password":"passw0rd"}
	 */
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		//パースしてエラーが起きた時
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// session key
	var session model.SessionKey
	log.Print("check email")
	user := database.CheckEmail(u.Email) // ユーザ情報をゲット
	if user.Name == "" {
		c.JSON(http.StatusOK, gin.H{
			"error": "not exists user.",
		})
		return
	}

	u.ToHash()
	log.Print("check password")
	if user.Password != u.Password || user.Email != u.Email {
		c.JSON(http.StatusOK, gin.H{
			"error": "email or password.",
		})
		return
	}
	// セッションがあるか確認
	log.Print("check session.")
	db, _ := database.GetGormDb()
	session.UserId = user.ID
	if database.IsExistsSession(user.ID) {
		log.Print("create sessionkey colum.")
		db.Create(&session) //セッション
	}

	// Session作成
	log.Print("make session key.")
	var session2 model.SessionKey
	key, _ := uuid.NewRandom()
	session2.SessionKey = key.String()
	session2.UserId = user.ID

	log.Print("insert session key.:", session2.SessionKey)
	// db.Create(&session2)
	db.Model(&session2).Where("user_id", user.ID).Updates(session2)
	log.Print("return sessionkey")
	c.JSON(http.StatusOK, gin.H{"session_key": session2.SessionKey})
}
