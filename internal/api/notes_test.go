package api

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNotes(t *testing.T) {
	t.Run("failed unauthorized get notes", func(t *testing.T) {
		app, routers, conf := NewApiTest()
		GetNotes(routers.apiRouter, conf)

		r := PerformRequest(app, "GET", "/api/notes")

		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})
}
