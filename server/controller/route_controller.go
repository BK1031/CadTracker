package controller

import (
	"context"
	"flag"
	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server/config"
	"server/service"
	"strings"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/ping", Ping)
	router.GET("/users", GetAllUsers)
	router.GET("/users/:userID", GetUserByID)
	router.POST("/users/:userID", CreateUser)
}

func InitializeDiscordBot(dgrouter *exrouter.Route) {
	dgrouter.On("ping", func(ctx *exrouter.Context) {
		ctx.Reply("pong")
	}).Desc("status check")
	dgrouter.On("avatar", func(ctx *exrouter.Context) {
		ctx.Reply(ctx.Msg.Author.AvatarURL("2048"))
	}).Desc("returns the user's avatar")

	dgrouter.Default = dgrouter.On("help", func(ctx *exrouter.Context) {
		var text = ""
		for _, v := range dgrouter.Routes {
			text += v.Name + " : \t" + v.Description + "\n"
		}
		ctx.Reply("```" + text + "```")
	}).Desc("prints this help menu")
	service.Discord.AddHandler(func(_ *discordgo.Session, m *discordgo.MessageCreate) {
		dgrouter.FindAndExecute(service.Discord, *flag.String("p", config.DiscordPrefix, "bot prefix"), service.Discord.State.User.ID, m.Message)
	})
	err := service.Discord.Open()
	if err != nil {
		log.Fatal(err)
	}
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func AuthChecker() gin.HandlerFunc {
	return func(c *gin.Context) {

		var requestUserID string

		ctx := context.Background()
		client, err := service.FirebaseAdmin.Auth(ctx)
		if err != nil {
			log.Fatalf("error getting Auth client: %v\n", err)
		}
		if c.GetHeader("Authorization") != "" {
			token, err := client.VerifyIDToken(ctx, strings.Split(c.GetHeader("Authorization"), "Bearer ")[1])
			if err != nil {
				println("error verifying ID token")
				requestUserID = "null"
			} else {
				println("Decoded User ID: " + token.UID)
				requestUserID = token.UID
			}
		} else {
			println("No user token provided")
			requestUserID = "null"
		}

		// The main authentication gateway per request path
		// The requesting user's ID and roles are pulled and used below
		// Any path can also be quickly halted if not ready for prod
		if c.FullPath() == "/users/:userID" {
			// Creating or modifying a user requires the requesting user to have a matching user ID
			if c.Request.Method == "POST" {
				if requestUserID != c.Param("userID") {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You do not have permission to edit this resource"})
				}
			}
		}
		c.Next()
	}
}
