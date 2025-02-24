package main

import (
	"log"

	"github.com/dev-diver/gongmo/controller"
	"github.com/dev-diver/gongmo/domain"
)

type Server interface {
	ListenAndServe(port string) error
}

type InMemoryAccountStore struct {
	store map[domain.AccountId]int
}

func (i *InMemoryAccountStore) GetAccount(id domain.AccountId) (int, error) {
	return i.store[id], nil
}

func (i *InMemoryAccountStore) StoreAccount(id domain.AccountId, amount int) error {
	i.store[id] = amount
	return nil
}

func main() {
	server := NewFiberServer()
	accountStore := &InMemoryAccountStore{}
	accountController := controller.NewAccountController(accountStore)
	server.Register(accountController)
	ListenForGracefulShutdown(server, "8080")
}

func ListenForGracefulShutdown(server Server, port string) {
	if err := server.ListenAndServe(port); err != nil {
		log.Fatalf("graceful shutdown 실패, 응답이 전달되지 않았을 수 있음 %v", err)
	}
	log.Println("graceful shutdown 성공, 모든 응답이 전달됨")
}
