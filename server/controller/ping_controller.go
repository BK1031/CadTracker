package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/config"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "CadTracker Server v" + config.Version + " is online!"})
}
