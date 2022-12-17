package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/isksss/postit/server/controller"
	"github.com/isksss/postit/server/database"
	"github.com/isksss/postit/server/middleware"
	"github.com/isksss/postit/server/model"
	"go.uber.org/zap"
)

func main() {
	// database init
	log.Default().Print("Auto Migrate.")
	db := database.Migrate()
	log.Default().Print("End Migrate")
	if database.HasTable(db, &model.User{}) {
		log.Println("User Table is Exists.")
	}

	// logger
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	// gin
	r := gin.Default()
	v1 := r.Group("v1")

	//set loger
	v1.Use(middleware.Logger(logger))
	v1.POST("/register", controller.Register)
	v1.POST("/login", controller.Login)

	r.Run(":8000")
}
