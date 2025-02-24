package main

import (
	"log"
)

type Server interface {
	ListenAndServe(port string) error
}

type InMemoryAccountStore struct {
	store map[AccountId]int
}

func (i *InMemoryAccountStore) GetAccount(id AccountId) (int, error) {
	return i.store[id], nil
}

func main() {
	server := NewFiberServer()
	accountStore := &InMemoryAccountStore{}
	accountController := NewAccountController(accountStore)
	server.Register(accountController)
	ListenForGracefulShutdown(server, "8080")
}

func ListenForGracefulShutdown(server Server, port string) {
	if err := server.ListenAndServe(port); err != nil {
		log.Fatalf("graceful shutdown 실패, 응답이 전달되지 않았을 수 있음 %v", err)
	}
	log.Println("graceful shutdown 성공, 모든 응답이 전달됨")
}
