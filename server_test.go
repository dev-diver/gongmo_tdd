package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/dev-diver/gongmo/controller"
	"github.com/dev-diver/gongmo/domain"
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
	store      map[domain.AccountId]int
	storeCalls []struct {
		Id     domain.AccountId
		Amount int
	}
}

func (s *StubStore) GetAccount(id domain.AccountId) (int, error) {
	account, ok := s.store[id]
	if !ok {
		return 0, fmt.Errorf("account not found: %s", id)
	}
	return account, nil
}

func (s *StubStore) StoreAccount(id domain.AccountId, amount int) error {
	s.storeCalls = append(s.storeCalls, struct {
		Id     domain.AccountId
		Amount int
	}{Id: id, Amount: amount})
	return nil
}

func TestGETMyAccount(t *testing.T) {
	server := NewTestServer()

	stubStore := &StubStore{
		store: map[domain.AccountId]int{
			"1": 0,
			"2": 1,
		},
	}

	accountController := controller.NewAccountController(stubStore)
	server.Register(accountController)

	t.Run("계좌 정보 가져오기", func(t *testing.T) {

		id := domain.AccountId("1")
		request := newGetAccountRequest(id)
		response, _ := server.Test(request)

		got, _ := io.ReadAll(response.Body)

		assert.Equal(t, response.StatusCode, http.StatusOK)
		assert.Equal(t, string(got), "0")
	})

	t.Run("계좌 정보 가져오기2", func(t *testing.T) {
		id := domain.AccountId("2")
		request := newGetAccountRequest(id)
		response, _ := server.Test(request)

		got, _ := io.ReadAll(response.Body)

		assert.Equal(t, response.StatusCode, http.StatusOK)
		assert.Equal(t, string(got), "1")
	})

	t.Run("없는 계좌 정보 가져오기", func(t *testing.T) {
		id := domain.AccountId("3")
		request := newGetAccountRequest(id)
		response, _ := server.Test(request)

		assert.Equal(t, response.StatusCode, http.StatusNotFound)
	})
}

func newGetAccountRequest(id domain.AccountId) *http.Request {
	request, _ := http.NewRequest("GET", fmt.Sprintf("/account/%s", id), nil)
	return request
}

func TestStoreAccount(t *testing.T) {
	server := NewTestServer()

	stubStore := &StubStore{
		store: map[domain.AccountId]int{},
	}
	accountController := controller.NewAccountController(stubStore)
	server.Register(accountController)

	t.Run("계좌 정보 저장하기", func(t *testing.T) {
		request := newPostAccountRequest(domain.AccountId("1"), 100)
		response, _ := server.Test(request)

		assert.Equal(t, response.StatusCode, http.StatusAccepted)
		assert.Equal(t, stubStore.storeCalls[0], struct {
			Id     domain.AccountId
			Amount int
		}{Id: domain.AccountId("1"), Amount: 100})
	})
}

func newPostAccountRequest(id domain.AccountId, amount int) *http.Request {
	request, _ := http.NewRequest("POST", fmt.Sprintf("/account/%s", id), bytes.NewBuffer([]byte(fmt.Sprintf("%d", amount))))
	return request
}
