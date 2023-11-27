package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/bifidokk/go-api/internal/config"
	"github.com/bifidokk/go-api/internal/service/token"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(conf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var authService = conf.Services.AuthService

		authHeader := c.Request.Header.Get("Authorization")
		authHeaderParts := strings.Split(authHeader, " ")

		if len(authHeaderParts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})

			c.Abort()
			return
		}

		claims, err := token.ValidateToken(authHeaderParts[1], conf.Env.JwtSecret)

		if err != nil {
			abortUnauthorized(c, err)
			return
		}

		user, err := authService.GetUserByEmail(claims.Email)

		if err != nil {
			abortUnauthorized(c, err)
			return
		}

		c.Set("user", user)

		c.Next()
	}
}

func abortUnauthorized(c *gin.Context, err error) {
	log.Println("JwtAuthMiddleware authorization error: " + err.Error())

	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "Unauthorized",
	})

	c.Abort()
}
