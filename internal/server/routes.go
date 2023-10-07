package server

import (
	"github.com/bifidokk/go-api/internal/api"
	"github.com/gin-gonic/gin"
)

var APIv1 *gin.RouterGroup

func RegisterRoutes(router *gin.Engine) {
	api.Ping(APIv1)
}
