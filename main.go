package main

import (
	"log"
)

func main() {
	var (
		port   = "8080"
		server = NewServer(port)
	)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("graceful shutdown 실패, 응답이 전달되지 않았을 수 있음 %v", err)
	}

	log.Println("graceful shutdown 성공, 모든 응답이 전달됨")
}
