package main

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestServer interface {
	Test(request *http.Request) (*http.Response, error)
}

func TestGETMyAccount(t *testing.T) {
	t.Run("계좌 정보 가져오기", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "http://localhost:8080/my-account", nil)

		server := NewFiberServer()
		response, _ := server.Test(request)

		got, _ := io.ReadAll(response.Body)

		assert.Equal(t, response.StatusCode, 200)
		assert.Equal(t, string(got), "0")
	})
}
