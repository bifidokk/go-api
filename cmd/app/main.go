package main

import (
	"github.com/bifidokk/go-api/internal/config"
	"github.com/bifidokk/go-api/internal/server"
	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.NewConfig()
	config.Init(conf)

	router := gin.Default()

	server.APIv1 = router.Group("/")
	server.RegisterRoutes(router)

	router.Run("localhost:8081")
}
