package main

import (
	"errors"
	"log"
	"strings"

	"github.com/dev-diver/gongmo/controller"
	"github.com/dev-diver/gongmo/domain"
	"github.com/gofiber/fiber/v2"
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
	idCopy := domain.AccountId(strings.Clone(string(id)))
	if _, ok := i.accounts[idCopy]; !ok {
		return 0, errors.New("account not found: " + string(id))
	}
	return i.accounts[idCopy], nil
}

func (i *InMemoryAccountStore) StoreAccount(id domain.AccountId, amount int) error {
	idCopy := domain.AccountId(strings.Clone(string(id)))
	i.accounts[idCopy] = amount
	return nil
}

func main() {
	server := NewFiberServer()
	server.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
	accountStore := NewInMemoryAccountStore()
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
