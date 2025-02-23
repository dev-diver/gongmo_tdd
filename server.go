package main

import (
	"context"
	"net/http"
)

type HttpServer struct {
	ctx        context.Context
	httpServer *http.Server
}

var defaultPort = "8080"

func NewHttpServer() *HttpServer {
	ctx := context.Background()
	httpServer := &http.Server{Addr: ":" + defaultPort, Handler: http.HandlerFunc(Handler)}

	return &HttpServer{
		ctx:        ctx,
		httpServer: httpServer,
	}
}

func (s *HttpServer) ListenAndServe(port ...string) error {
	if len(port) > 0 {
		s.httpServer.Addr = ":" + port[0]
	}
	return s.httpServer.ListenAndServe()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "안녕하세요"}`))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "지원하지 않는 메서드입니다"}`))
	}
}
