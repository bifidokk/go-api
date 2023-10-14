package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("JwtAuthMiddleware checking")
		c.Next()
	}
}
