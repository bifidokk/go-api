package api

import (
	"log"
	"net/http"

	"github.com/bifidokk/go-api/internal/config"
	"github.com/bifidokk/go-api/internal/repository"
	"github.com/bifidokk/go-api/internal/service/auth"
	"github.com/bifidokk/go-api/internal/service/signup"
	"github.com/gin-gonic/gin"
)

func Signup(router *gin.RouterGroup, conf *config.Config) {
	var userRepository = repository.NewUserRepository(conf.Db())
	var authService = auth.NewAuth(userRepository, conf.Env.JwtSecret, int(conf.Env.JwtTtl))
	var signupService = signup.NewSignup(userRepository, authService)

	router.POST("/signup", func(c *gin.Context) {
		var request signup.SignupRequest
		err := c.ShouldBind(&request)

		if err != nil {
			log.Println("Validation error: " + err.Error())
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		_, err = signupService.CreateUser(request)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, nil)
			return
		}

		c.JSON(http.StatusCreated, nil)
	})
}
