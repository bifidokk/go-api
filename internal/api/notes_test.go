package api

import (
	"net/http"
	"testing"

	"github.com/bifidokk/go-api/internal/fixtures"
	"github.com/bifidokk/go-api/internal/service/token"
	"github.com/stretchr/testify/assert"
)

func TestGetNotes(t *testing.T) {
	t.Run("failed unauthorized get notes", func(t *testing.T) {
		app, routers, conf := NewApiTest()
		GetNotes(routers.apiRouter, conf)

		r := PerformRequest(app, "GET", "/api/notes", map[string]string{})

		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})

	t.Run("succesfull get notes", func(t *testing.T) {
		app, routers, conf := NewApiTest()
		GetNotes(routers.apiRouter, conf)

		user := fixtures.UserFixtures["user@test.com"]

		accessToken, _ := token.CreatAccessToken(
			&user,
			conf.Env.JwtSecret,
			int(conf.Env.JwtTtl),
		)

		headers := map[string]string{
			"Authorization": "Bearer " + accessToken,
		}

		r := PerformRequest(app, "GET", "/api/notes", headers)

		assert.Equal(t, http.StatusOK, r.Code)
	})
}
