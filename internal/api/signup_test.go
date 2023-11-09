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
		app, routers, conf := NewApiTest()
		Signup(routers.publicRouter, conf)

		body, _ := json.Marshal(signup.SignupRequest{
			Email:    "new-user@test.com",
			Password: "123456!",
		})

		r := PerformRequestWithBody(app, "POST", "/public/signup", string(body))

		assert.Equal(t, http.StatusCreated, r.Code)
	})

	t.Run("unsuccessful signup because of validation error", func(t *testing.T) {
		app, routers, conf := NewApiTest()
		Signup(routers.publicRouter, conf)

		body, _ := json.Marshal(signup.SignupRequest{
			Email:    "new-user",
			Password: "123456!",
		})

		r := PerformRequestWithBody(app, "POST", "/public/signup", string(body))

		assert.Equal(t, http.StatusUnprocessableEntity, r.Code)
	})

	t.Run("unsuccessful signup because of existing email", func(t *testing.T) {
		app, routers, conf := NewApiTest()
		Signup(routers.publicRouter, conf)

		body, _ := json.Marshal(signup.SignupRequest{
			Email:    "user2@test.com",
			Password: "123456!",
		})

		r := PerformRequestWithBody(app, "POST", "/public/signup", string(body))

		assert.Equal(t, http.StatusUnprocessableEntity, r.Code)
	})
}
