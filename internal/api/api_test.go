package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/bifidokk/go-api/internal/config"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func NewApiTest() (app *gin.Engine, router *gin.RouterGroup, conf *config.Config) {
	conf = config.NewTestConfig()
	config.InitTest(conf)

	gin.SetMode(gin.TestMode)

	app = gin.New()
	router = app.Group("/public")
	config.RegisterRepositories(conf)

	return app, router, conf
}

func PerformRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}

func PerformRequestWithBody(r http.Handler, method, path, body string) *httptest.ResponseRecorder {
	reader := strings.NewReader(body)
	req, _ := http.NewRequest(method, path, reader)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}
