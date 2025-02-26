package main

import (
	"io"
	"net/http"
	"testing"

	"github.com/dev-diver/gongmo/controller"
	"github.com/dev-diver/gongmo/domain"
	"github.com/dev-diver/gongmo/store"
	"github.com/stretchr/testify/assert"
)

func TestStoreAccountAndGetThem(t *testing.T) {
	server := NewFiberServer()
	store := store.NewInMemoryAccountStore()
	accountController := controller.NewAccountController(store)
	server.Register(accountController)

	server.Test(newPostAccountRequest(domain.AccountId("3"), 100))
	server.Test(newPostAccountRequest(domain.AccountId("1"), 300))

	response, _ := server.Test(newGetAccountRequest(domain.AccountId("3")))
	got, _ := io.ReadAll(response.Body)

	assert.Equal(t, response.StatusCode, http.StatusOK)
	assert.Equal(t, "100", string(got))

	response, _ = server.Test(newGetAccountRequest(domain.AccountId("1")))
	got, _ = io.ReadAll(response.Body)

	assert.Equal(t, response.StatusCode, http.StatusOK)
	assert.Equal(t, "300", string(got))
}
