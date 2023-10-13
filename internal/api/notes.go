package api

import (
	"net/http"

	"github.com/bifidokk/go-api/internal/config"
	"github.com/bifidokk/go-api/internal/repository"
	"github.com/gin-gonic/gin"
)

func GetNotes(router *gin.RouterGroup, conf *config.Config) {
	var noteRepository = repository.NewNoteRepository(conf.Db())

	router.GET("/notes", func(c *gin.Context) {
		if notes, err := noteRepository.FindAll(); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, nil)
			return
		} else {
			c.JSON(http.StatusOK, notes)
		}
	})
}
