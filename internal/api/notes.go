package api

import (
	"net/http"

	"github.com/bifidokk/go-api/internal/config"
	"github.com/bifidokk/go-api/internal/entity"
	"github.com/bifidokk/go-api/internal/repository"
	"github.com/gin-gonic/gin"
)

func GetNotes(router *gin.RouterGroup, conf *config.Config) {
	var noteRepository = repository.NewNoteRepository(conf.Db())

	router.GET("/notes", func(c *gin.Context) {
		user, _ := c.Get("user")

		if notes, err := noteRepository.FindByUser(user.(*entity.User)); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, nil)
			return
		} else {
			c.JSON(http.StatusOK, notes)
		}
	})
}
