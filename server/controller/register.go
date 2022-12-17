package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isksss/postit/server/database"
	"github.com/isksss/postit/server/model"
)

func Register(c *gin.Context) {
	// ここにユーザ登録処理を記述
	// log.Println("register user...")
	// json parse

	var uinfo model.User
	if err := c.ShouldBindJSON(&uinfo); err != nil {
		//パースしてエラーが起きた時
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//emailに被りがないかチェック
	u := database.CheckEmail(uinfo.Email)

	//emailがすでに登録されている場合
	if u.Email == uinfo.Email {
		c.JSON(http.StatusOK, gin.H{"error": "email is already register."})
		return
	}

	// ユーザ登録
	uinfo.ToHash()
	err := database.Register(&uinfo)
	if err.Error != nil {
		//登録失敗した時
		c.JSON(http.StatusOK, gin.H{"error": err.Error.Error()})
		return
	}
	// log.Println("log")
	u = database.CheckEmail(uinfo.Email)
	c.JSON(http.StatusOK, gin.H{"id": u.ID})

}
