package api

import (
	"log"
	"net/http"

	"github.com/bifidokk/go-api/internal/config"
	"github.com/bifidokk/go-api/internal/repository"
	"github.com/bifidokk/go-api/internal/service/auth"
	"github.com/gin-gonic/gin"
)

func Login(router *gin.RouterGroup, conf *config.Config) {
	var userRepository = repository.NewUserRepository(conf.Db())
	var authService = auth.NewAuth(userRepository, conf.Env.JwtSecret, int(conf.Env.JwtTtl))

	router.POST("/login", func(c *gin.Context) {
		var request auth.LoginRequest
		err := c.ShouldBind(&request)

		if err != nil {
			log.Println("Validation error: " + err.Error())
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, nil)
			return
		}

		token, err := authService.Authenticate(request)

		if err != nil {
			log.Println("Authentication error:" + err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid credentials",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	})
}
