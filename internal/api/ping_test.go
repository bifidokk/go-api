package api

import (
	"fmt"
	"testing"
)

func TestGetAlbum(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		app, router, _ := NewApiTest()

		Ping(router)

		r := PerformRequest(app, "GET", "/api/ping")
		fmt.Println(r)
	})
}
