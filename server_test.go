package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestServer interface {
	Test(request *http.Request) (*http.Response, error)
}

func NewTestServer() TestServer {
	return NewFiberServer() // 다른 서버를 사용할 경우 이 함수를 수정
}

func TestGETMyAccount(t *testing.T) {

	t.Run("계좌 정보 가져오기", func(t *testing.T) {

		data := map[string]interface{}{
			"id": "1",
		}
		request := newGetAccountRequest(data)
		server := NewTestServer()
		response, _ := server.Test(request)

		got, _ := io.ReadAll(response.Body)

		assert.Equal(t, response.StatusCode, 200)
		assert.Equal(t, string(got), "0")
	})
}

func newGetAccountRequest(data map[string]interface{}) *http.Request {
	json, _ := json.Marshal(data)
	request, _ := http.NewRequest("GET", "/account", bytes.NewBuffer(json))
	return request
}
