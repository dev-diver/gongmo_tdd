package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestServer interface {
	Test(request *http.Request) (*http.Response, error)
	Register(controller Controller)
}

func NewTestServer() TestServer {
	return NewFiberServer() // 다른 서버를 사용할 경우 이 함수를 수정
}

type StubStore struct {
	store map[AccountId]int
}

func (s *StubStore) GetAccount(id AccountId) (int, error) {
	account, ok := s.store[id]
	if !ok {
		return 0, fmt.Errorf("account not found: %s", id)
	}
	return account, nil
}

func TestGETMyAccount(t *testing.T) {
	server := NewTestServer()

	stubStore := &StubStore{
		store: map[AccountId]int{
			"1": 0,
			"2": 1,
		},
	}

	accountController := NewAccountController(stubStore)
	server.Register(accountController)

	t.Run("계좌 정보 가져오기", func(t *testing.T) {

		id := AccountId("1")
		request := newGetAccountRequest(id)
		response, _ := server.Test(request)

		got, _ := io.ReadAll(response.Body)

		assert.Equal(t, response.StatusCode, http.StatusOK)
		assert.Equal(t, string(got), "0")
	})

	t.Run("계좌 정보 가져오기2", func(t *testing.T) {
		id := AccountId("2")
		request := newGetAccountRequest(id)
		response, _ := server.Test(request)

		got, _ := io.ReadAll(response.Body)

		assert.Equal(t, response.StatusCode, http.StatusOK)
		assert.Equal(t, string(got), "1")
	})

	t.Run("없는 계좌 정보 가져오기", func(t *testing.T) {
		id := AccountId("3")
		request := newGetAccountRequest(id)
		response, _ := server.Test(request)

		assert.Equal(t, response.StatusCode, http.StatusNotFound)
	})
}

func newGetAccountRequest(id AccountId) *http.Request {
	request, _ := http.NewRequest("GET", fmt.Sprintf("/account/%s", id), nil)
	return request
}

func TestStoreAccount(t *testing.T) {
	server := NewTestServer()

	stubStore := &StubStore{
		store: map[AccountId]int{},
	}
	accountController := NewAccountController(stubStore)
	server.Register(accountController)

	t.Run("계좌 정보 저장하기", func(t *testing.T) {
		request := newPostAccountRequest("1", 100)
		response, _ := server.Test(request)

		assert.Equal(t, response.StatusCode, http.StatusAccepted)
	})
}

func newPostAccountRequest(id AccountId, amount int) *http.Request {
	request, _ := http.NewRequest("POST", fmt.Sprintf("/account/%s", id), bytes.NewBuffer([]byte(fmt.Sprintf("%d", amount))))
	return request
}
