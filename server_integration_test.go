package main

import (
	"io"
	"net/http"
	"testing"

	"github.com/dev-diver/gongmo/controller"
	"github.com/dev-diver/gongmo/domain"
	"github.com/stretchr/testify/assert"
)

func TestStoreAccountAndGetThem(t *testing.T) {
	server := NewTestServer()
	store := &InMemoryAccountStore{}
	accountController := controller.NewAccountController(store)
	server.Register(accountController)

	server.Test(newPostAccountRequest(domain.AccountId("2"), 200))
	server.Test(newPostAccountRequest(domain.AccountId("1"), 100))

	response, _ := server.Test(newGetAccountRequest(domain.AccountId("1")))
	got, _ := io.ReadAll(response.Body)

	assert.Equal(t, response.StatusCode, http.StatusOK)
	assert.Equal(t, string(got), "100")

	response, _ = server.Test(newGetAccountRequest(domain.AccountId("2")))
	got, _ = io.ReadAll(response.Body)

	assert.Equal(t, response.StatusCode, http.StatusOK)
	assert.Equal(t, string(got), "200")
}
