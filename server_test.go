package main

import (
	"fmt"
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

		id := AccountId("1")
		request := newGetAccountRequest(id)
		server := NewTestServer()
		response, _ := server.Test(request)

		got, _ := io.ReadAll(response.Body)

		assert.Equal(t, response.StatusCode, 200)
		assert.Equal(t, string(got), "0")
	})

	t.Run("계좌 정보 가져오기2", func(t *testing.T) {
		id := AccountId("2")
		request := newGetAccountRequest(id)
		server := NewTestServer()
		response, _ := server.Test(request)

		got, _ := io.ReadAll(response.Body)

		assert.Equal(t, response.StatusCode, 200)
		assert.Equal(t, string(got), "1")
	})
}

func newGetAccountRequest(id AccountId) *http.Request {
	request, _ := http.NewRequest("GET", fmt.Sprintf("/account/%s", id), nil)
	return request
}
