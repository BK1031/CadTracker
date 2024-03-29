package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server/service"
	"strings"
)

func InitializeHoods(router *gin.Engine) {
	router.GET("/ping", Ping)
	router.GET("/users", GetAllUsers)
	router.GET("/users/:userID", GetUserByID)
	router.POST("/users/:userID", CreateUser)
	router.GET("/users/:userID/events", GetAllEventsForUser)
	router.GET("/users/:userID/events/latest", GetLatestEventForUser)
	router.GET("/users/:userID/events/recent/year", GetLatestYearEventsForUser)
	router.GET("/users/:userID/events/recent/day", GetLatestDayEventsForUser)
	router.POST("/users/:userID/events/start", StartEvent)
	router.POST("/users/:userID/events/stop", StopEvent)
	router.GET("/users/:userID/events/current", GetCurrentEventForUser)
	router.POST("/users/:userID/events/:eventID", EditEventForUser)
	router.GET("/events/:eventID", GetEventByID)
	router.DELETE("/events/:eventID", DeleteEvent)
}

func HoodLogger() gin.HandlerFunc {
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
		if strings.HasPrefix(c.FullPath(), "/users/:userID") {
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
