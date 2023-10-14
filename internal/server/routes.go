package server

import (
	"github.com/bifidokk/go-api/internal/api"
	"github.com/bifidokk/go-api/internal/config"
	"github.com/bifidokk/go-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, conf *config.Config) {
	apiV1Group := router.Group("/api")
	publicGroup := router.Group("/public")

	api.Ping(publicGroup)

	apiV1Group.Use(middleware.JwtAuthMiddleware())
	api.GetNotes(apiV1Group, conf)
}
