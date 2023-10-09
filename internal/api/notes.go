package api

import (
	"net/http"

	"github.com/bifidokk/go-api/internal/query"
	"github.com/gin-gonic/gin"
)

func GetNotes(router *gin.RouterGroup) {
	router.GET("/notes", func(c *gin.Context) {
		if notes, err := query.Notes(); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, nil)
			return
		} else {
			c.JSON(http.StatusOK, notes)
		}
	})
}
