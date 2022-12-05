package main

import (
	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"server/config"
	"server/controller"
	"server/service"
	"time"
)

var router *gin.Engine
var dgrouter *exrouter.Route

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
	dgrouter = exrouter.New()
	service.InitializeDB()
	service.InitializeFirebase()
	service.ConnectDiscord()
	controller.InitializeRoutes(router)
	controller.InitializeDiscordBot(dgrouter)
	router.Run(":" + config.Port)
}
