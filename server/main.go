package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"server/config"
	"server/controller"
	"server/service"
	"time"
)

var router *gin.Engine

func setupRouter() *gin.Engine {
	if config.Env == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		MaxAge:           12 * time.Hour,
		AllowCredentials: true,
	}))
	r.Use(controller.RequestLogger())
	r.Use(controller.AuthChecker())
	return r
}

func main() {
	router = setupRouter()
	service.InitializeDB()
	service.InitializeFirebase()
	service.ConnectDiscord()
	controller.InitializeDiscordBot()
	controller.InitializeRoutes(router)
	router.Run(":" + config.Port)
}
