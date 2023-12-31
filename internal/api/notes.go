package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/bifidokk/go-api/internal/config"
	"github.com/bifidokk/go-api/internal/entity"
	"github.com/bifidokk/go-api/internal/service/note"
	"github.com/gin-gonic/gin"
)

func GetNotes(router *gin.RouterGroup, conf *config.Config) {
	var noteRepository = conf.Repositories.NoteRepository

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

func CreateNote(router *gin.RouterGroup, conf *config.Config) {
	var noteService = conf.Services.NoteService

	router.POST("/notes", func(c *gin.Context) {
		user, _ := c.Get("user")

		var request note.CreateRequest
		err := c.ShouldBind(&request)

		if err != nil {
			log.Println("Validation error: " + err.Error())
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		if note, err := noteService.CreateNote(request, user.(*entity.User)); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
			return
		} else {
			c.JSON(http.StatusCreated, note)
		}
	})
}

func UpdateNote(router *gin.RouterGroup, conf *config.Config) {
	var noteService = conf.Services.NoteService
	var noteRepository = conf.Repositories.NoteRepository

	router.PATCH("/notes/:id", func(c *gin.Context) {
		user, _ := c.Get("user")
		noteId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, nil)
			return
		}

		userNote, err := noteRepository.FindUserNoteById(user.(*entity.User), noteId)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, nil)
			return
		}

		var request note.UpdateRequest
		err = c.ShouldBind(&request)

		if err != nil {
			log.Println("Validation error: " + err.Error())
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		if note, err := noteService.UpdateNote(request, userNote); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
			return
		} else {
			c.JSON(http.StatusOK, note)
		}
	})
}

func DeleteNote(router *gin.RouterGroup, conf *config.Config) {
	var noteService = conf.Services.NoteService
	var noteRepository = conf.Repositories.NoteRepository

	router.DELETE("/notes/:id", func(c *gin.Context) {
		user, _ := c.Get("user")
		noteId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, nil)
			return
		}

		userNote, err := noteRepository.FindUserNoteById(user.(*entity.User), noteId)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, nil)
			return
		}

		if err := noteService.DeleteNote(userNote); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
			return
		} else {
			c.JSON(http.StatusNoContent, nil)
		}
	})
}
