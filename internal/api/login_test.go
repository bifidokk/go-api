package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bifidokk/go-api/internal/service/auth"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	t.Run("successful login", func(t *testing.T) {
		app, router, conf := NewApiTest()
		Login(router, conf)

		body, _ := json.Marshal(auth.LoginRequest{
			Email:    "user@test.com",
			Password: "123456!",
		})

		r := PerformRequestWithBody(app, "POST", "/public/login", string(body))

		assert.Equal(t, http.StatusOK, r.Code)
	})
}
