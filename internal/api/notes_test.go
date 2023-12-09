package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/bifidokk/go-api/internal/fixtures"
	"github.com/bifidokk/go-api/internal/service/note"
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

	t.Run("succesfull get user notes", func(t *testing.T) {
		app, routers, conf := NewApiTest()
		GetNotes(routers.apiRouter, conf)

		user := fixtures.UserFixtures[0]

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

func TestCreateNote(t *testing.T) {
	t.Run("succesfull create a note", func(t *testing.T) {
		app, routers, conf := NewApiTest()
		CreateNote(routers.apiRouter, conf)

		user := fixtures.UserFixtures[0]

		accessToken, _ := token.CreatAccessToken(
			&user,
			conf.Env.JwtSecret,
			int(conf.Env.JwtTtl),
		)

		headers := map[string]string{
			"Authorization": "Bearer " + accessToken,
		}

		body, _ := json.Marshal(note.CreateRequest{
			Title:       "Note title",
			Description: "Note description",
		})

		r := PerformRequestWithBody(app, "POST", "/api/notes", string(body), headers)

		assert.Equal(t, http.StatusCreated, r.Code)

		responseContent := HTTPBodyToMap(r.Body)

		assert.Contains(t, responseContent, "id")

		assert.Contains(t, responseContent, "title")
		assert.Equal(t, "Note title", responseContent["title"])

		assert.Contains(t, responseContent, "description")
		assert.Equal(t, "Note description", responseContent["description"])

		assert.Contains(t, responseContent, "user_id")
		assert.Equal(t, 1, int(responseContent["user_id"].(float64)))
	})
}
func TestFailedCreateNote(t *testing.T) {
	t.Run("failed create a note", func(t *testing.T) {
		app, routers, conf := NewApiTest()
		CreateNote(routers.apiRouter, conf)

		user := fixtures.UserFixtures[0]

		accessToken, _ := token.CreatAccessToken(
			&user,
			conf.Env.JwtSecret,
			int(conf.Env.JwtTtl),
		)

		headers := map[string]string{
			"Authorization": "Bearer " + accessToken,
		}

		body, _ := json.Marshal(note.CreateRequest{
			Title:       "Note title",
		})

		r := PerformRequestWithBody(app, "POST", "/api/notes", string(body), headers)

		assert.Equal(t, http.StatusUnprocessableEntity, r.Code)

		responseContent := HTTPBodyToMap(r.Body)
		fmt.Println(responseContent)
	})
}

func TestUpdateNote(t *testing.T) {
	t.Run("succesfull edit a note", func(t *testing.T) {
		app, routers, conf := NewApiTest()
		UpdateNote(routers.apiRouter, conf)

		user := fixtures.UserFixtures[0]

		accessToken, _ := token.CreatAccessToken(
			&user,
			conf.Env.JwtSecret,
			int(conf.Env.JwtTtl),
		)

		headers := map[string]string{
			"Authorization": "Bearer " + accessToken,
		}

		body, _ := json.Marshal(note.CreateRequest{
			Title:       "New note title",
			Description: "New note description",
		})

		r := PerformRequestWithBody(app, "PATCH", "/api/notes/1", string(body), headers)

		assert.Equal(t, http.StatusOK, r.Code)

		responseContent := HTTPBodyToMap(r.Body)

		assert.Contains(t, responseContent, "id")

		assert.Contains(t, responseContent, "title")
		assert.Equal(t, "New note title", responseContent["title"])

		assert.Contains(t, responseContent, "description")
		assert.Equal(t, "New note description", responseContent["description"])
	})

	t.Run("failed edit a note", func(t *testing.T) {
		app, routers, conf := NewApiTest()
		UpdateNote(routers.apiRouter, conf)

		user := fixtures.UserFixtures[0]

		accessToken, _ := token.CreatAccessToken(
			&user,
			conf.Env.JwtSecret,
			int(conf.Env.JwtTtl),
		)

		headers := map[string]string{
			"Authorization": "Bearer " + accessToken,
		}

		body, _ := json.Marshal(note.CreateRequest{
			Description: "New note description",
		})

		r := PerformRequestWithBody(app, "PATCH", "/api/notes/1", string(body), headers)

		assert.Equal(t, http.StatusUnprocessableEntity, r.Code)
	})
}
