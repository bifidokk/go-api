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
	var authService = auth.NewAuth(userRepository)

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
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	})
}