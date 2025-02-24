package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/adaptor/v2"

	"github.com/dev-diver/gongmo/controller"
	"github.com/dev-diver/gongmo/domain"
	"github.com/stretchr/testify/assert"
)

func TestStoreAccountAndGetThem(t *testing.T) {
	server := NewFiberServer()
	store := NewInMemoryAccountStore()
	accountController := controller.NewAccountController(store)
	server.Register(accountController)

	handler := adaptor.FiberApp(server.app)

	recorder1 := httptest.NewRecorder()
	req1 := newPostAccountRequest(domain.AccountId("3b3a"), 100)
	handler.ServeHTTP(recorder1, req1)

	recorder2 := httptest.NewRecorder()
	req2 := newPostAccountRequest(domain.AccountId("30"), 300)
	store.GetAccount(domain.AccountId("3b3a"))
	handler.ServeHTTP(recorder2, req2)

	response := httptest.NewRecorder()
	handler.ServeHTTP(response, newGetAccountRequest(domain.AccountId("3b3a")))
	got, _ := io.ReadAll(response.Body)

	assert.Equal(t, response.Code, http.StatusOK)
	assert.Equal(t, "100", string(got))

	response = httptest.NewRecorder()
	handler.ServeHTTP(response, newGetAccountRequest(domain.AccountId("30")))
	got, _ = io.ReadAll(response.Body)

	assert.Equal(t, response.Code, http.StatusOK)
	assert.Equal(t, "300", string(got))
}
