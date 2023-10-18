package api

import (
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

}
