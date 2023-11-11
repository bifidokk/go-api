package api

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestPing(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		app, routers, _ := NewApiTest()

		Ping(routers.publicRouter)

		r := PerformRequest(app, "GET", "/public/ping", map[string]string{})

		assert.Equal(t, http.StatusOK, r.Code)

		message := gjson.Get(r.Body.String(), "message")
		assert.Equal(t, "pong", message.String())
	})
}
