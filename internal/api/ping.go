package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		data := gin.H {
			"message": "pong",
			
		}

		c.JSON(http.StatusOK, data)
	})
}
