package main

import (
	"context"
	"log"
	"net/http"
)

func main() {
	var (
		ctx        = context.Background()
		httpServer = &http.Server{Addr: ":8080", Handler: http.HandlerFunc(Handler)}
		server     = NewServer(httpServer)
	)

	if err := server.ListenAndServe(ctx); err != nil {
		log.Fatalf("graceful shutdown 실패, 응답이 전달되지 않았을 수 있음 %v", err)
	}

	log.Println("graceful shutdown 성공, 모든 응답이 전달됨")
}
