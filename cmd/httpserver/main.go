package main

import (
	"log"

	"github.com/dev-diver/gongmo/controller"
	"github.com/dev-diver/gongmo/store"
	"github.com/gofiber/fiber/v2"
)

type Server interface {
	ListenAndServe(port string) error
}

func main() {
	server := NewFiberServer()
	server.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
	accountStore := store.NewInMemoryAccountStore()
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
