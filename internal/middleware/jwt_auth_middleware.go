package middleware

import (
	"log"
	"net/http"
	"strings"

	jwtauth "github.com/bifidokk/go-api/internal/service"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		authHeaderParts := strings.Split(authHeader, " ")

		if len(authHeaderParts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})

			c.Abort()
			return
		}

		authorized, err := jwtauth.IsAuthorized(authHeaderParts[1], secret)

		if !authorized || err != nil {
			log.Println("JwtAuthMiddleware authorization error: " + err.Error())

			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
