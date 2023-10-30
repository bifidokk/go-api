package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bifidokk/go-api/internal/config"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	conf := config.NewConfig()
	config.Init(conf)

	code := m.Run()
	os.Exit(code)
}

func NewApiTest() (app *gin.Engine, router *gin.RouterGroup) {
	gin.SetMode(gin.TestMode)

	app = gin.New()
	router = app.Group("/public")

	return app, router
}

func PerformRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}
