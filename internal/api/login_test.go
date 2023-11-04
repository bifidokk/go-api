package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bifidokk/go-api/internal/service/auth"
	"github.com/stretchr/testify/assert"
)

type InvalidLoginRequest struct {
    Email string `json:"email"`
}

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

	t.Run("failed login with wrong credentials", func(t *testing.T) {
		app, router, conf := NewApiTest()
		Login(router, conf)

		body, _ := json.Marshal(auth.LoginRequest{
			Email:    "non-existing-user@test.com",
			Password: "123456!",
		})

		r := PerformRequestWithBody(app, "POST", "/public/login", string(body))

		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})

	t.Run("failed login with wrong password", func(t *testing.T) {
		app, router, conf := NewApiTest()
		Login(router, conf)

		body, _ := json.Marshal(auth.LoginRequest{
			Email:    "user@test.com",
			Password: "wrong password",
		})

		r := PerformRequestWithBody(app, "POST", "/public/login", string(body))

		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})

	t.Run("failed login without required fields", func(t *testing.T) {
		app, router, conf := NewApiTest()
		Login(router, conf)

		body, _ := json.Marshal(InvalidLoginRequest{
			Email:    "non-existing-user@test.com",
		})

		r := PerformRequestWithBody(app, "POST", "/public/login", string(body))

		assert.Equal(t, http.StatusUnprocessableEntity, r.Code)
	})
}
