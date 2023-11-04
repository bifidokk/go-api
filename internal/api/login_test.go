package api

import "testing"

func TestLogin(t *testing.T) {
	t.Run("successful login", func(t *testing.T) {
		_, router, conf := NewApiTest()

		Login(router, conf)
	})
}