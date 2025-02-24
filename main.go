package main

import (
	"fmt"
	"log"

	"github.com/dev-diver/gongmo/domain"
)

type Server interface {
	ListenAndServe(port string) error
}

type InMemoryAccountStore struct {
	accounts map[domain.AccountId]int
}

func NewInMemoryAccountStore() *InMemoryAccountStore {
	return &InMemoryAccountStore{
		accounts: make(map[domain.AccountId]int),
	}
}

func (i *InMemoryAccountStore) GetAccount(id domain.AccountId) (int, error) {
	fmt.Printf("Retrieving ID: %v\n", id)
	fmt.Printf("Get Current map state: %+v\n", i.accounts)
	return i.accounts[id], nil
}

func (i *InMemoryAccountStore) StoreAccount(id domain.AccountId, amount int) error {
	fmt.Printf("Before Current map state: %+v\n", i.accounts)
	fmt.Printf("Storing ID: %v, Amount: %d\n", id, amount)
	i.accounts[id] = amount
	fmt.Printf("After Current map state: %+v\n", i.accounts)
	return nil
}

func main() {
	server := NewFiberServer()
	// accountStore := NewInMemoryAccountStore()
	// accountController := controller.NewAccountController(accountStore)
	// server.Register(accountController)
	ListenForGracefulShutdown(server, "8080")
}

func ListenForGracefulShutdown(server Server, port string) {
	if err := server.ListenAndServe(port); err != nil {
		log.Fatalf("graceful shutdown 실패, 응답이 전달되지 않았을 수 있음 %v", err)
	}
	log.Println("graceful shutdown 성공, 모든 응답이 전달됨")
}
