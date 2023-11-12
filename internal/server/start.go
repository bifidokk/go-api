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

	config.RegisterRepositories(conf)
	config.RegisterValidators(conf)
	RegisterRoutes(router, conf)

	router.Run(":8080")
}
