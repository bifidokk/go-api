package server

import (
	"net/http"

	"github.com/bifidokk/go-api/internal/api"
	"github.com/bifidokk/go-api/internal/config"
	"github.com/bifidokk/go-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, conf *config.Config) {
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "404 Not found",
		})
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "405 Method not allowed",
		})
	})

	apiV1Group := router.Group("/api")
	publicGroup := router.Group("/public")

	api.Ping(publicGroup)

	apiV1Group.Use(middleware.JwtAuthMiddleware())
	api.GetNotes(apiV1Group, conf)
}
