package server

import (
	"github.com/bifidokk/go-api/internal/api"
	"github.com/bifidokk/go-api/internal/config"
	"github.com/gin-gonic/gin"
)

var APIv1 *gin.RouterGroup

func RegisterRoutes(router *gin.Engine, conf *config.Config) {
	api.Ping(APIv1)
	api.GetNotes(APIv1, conf)
}
