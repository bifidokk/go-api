package api

import (
	"net/http"
	"net/http/httptest"

	"github.com/bifidokk/go-api/internal/config"
	"github.com/gin-gonic/gin"
)

func NewApiTest() (app *gin.Engine, router *gin.RouterGroup, conf *config.Config) {
	gin.SetMode(gin.TestMode)

	app = gin.New()
	router = app.Group("/public")

	conf = config.NewConfig()
	config.Init(conf)

	return app, router, conf
}

func PerformRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}
