package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bifidokk/go-api/internal/service/signup"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	t.Run("successful signup", func(t *testing.T) {
		app, router, conf := NewApiTest()
		Signup(router, conf)

		body, _ := json.Marshal(signup.SignupRequest{
			Email:    "new-user@test.com",
			Password: "123456!",
		})

		r := PerformRequestWithBody(app, "POST", "/public/signup", string(body))

		assert.Equal(t, http.StatusCreated, r.Code)
	})
}
