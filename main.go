package main

import (
	"log"
)

type Server interface {
	ListenAndServe(port ...string) error
}

func main() {
	server := NewHttpServer()
	ListenForGracefulShutdown(server, "8080")
}

func ListenForGracefulShutdown(server Server, port string) {
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("graceful shutdown 실패, 응답이 전달되지 않았을 수 있음 %v", err)
	}
	log.Println("graceful shutdown 성공, 모든 응답이 전달됨")
}
