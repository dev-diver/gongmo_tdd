package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGETMyAccount(t *testing.T) {
	t.Run("계좌 정보 가져오기", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "http://localhost:8080/my-account", nil)
		response := httptest.NewRecorder()
		AccountServer(response, request)

		got := response.Body.String()

		assert.Equal(t, response.Code, 200)
		assert.Equal(t, got, "0")
	})
}
