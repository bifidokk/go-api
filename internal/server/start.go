package server

import (
	"log"

	"github.com/bifidokk/go-api/internal/config"
	"github.com/gin-gonic/gin"
)

func Start(conf *config.Config) {
	router := gin.Default()
	if err := router.SetTrustedProxies(nil); err != nil {
		log.Printf("server: %s", err)
	}

	RegisterValidators(conf)
	RegisterRoutes(router, conf)

	router.Run("localhost:8081")
}
