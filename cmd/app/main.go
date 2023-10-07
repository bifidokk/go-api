package main

import (
	"github.com/bifidokk/go-api/internal/server"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	server.APIv1 = router.Group("/")
	server.RegisterRoutes(router)

	router.Run("localhost:8081")
}
