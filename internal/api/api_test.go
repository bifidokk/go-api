package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/bifidokk/go-api/internal/config"
	"github.com/bifidokk/go-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

type apiRouters struct {
	publicRouter *gin.RouterGroup
	apiRouter    *gin.RouterGroup
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func NewApiTest() (app *gin.Engine, routers *apiRouters, conf *config.Config) {
	conf = config.NewTestConfig()
	config.InitTest(conf)

	gin.SetMode(gin.TestMode)

	app = gin.New()

	publicRouter := app.Group("/public")
	apiRouter := app.Group("/api")
	apiRouter.Use(middleware.JwtAuthMiddleware(conf.Env.JwtSecret))

	routers = &apiRouters{
		publicRouter: publicRouter,
		apiRouter:    apiRouter,
	}

	config.RegisterRepositories(conf)
	config.RegisterValidators(conf)

	return app, routers, conf
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

func HTTPBodyToMap(body *bytes.Buffer) map[string]any {
	response := map[string]any{}
	_ = json.NewDecoder(body).Decode(&response)

	return response
}
