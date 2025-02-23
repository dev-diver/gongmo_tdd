package main

import (
	"context"
	"net/http"
)

type Server struct {
	ctx        context.Context
	httpServer *http.Server
}

func NewServer(port string) *Server {
	ctx := context.Background()
	httpServer := &http.Server{Addr: ":" + port, Handler: http.HandlerFunc(Handler)}

	return &Server{
		ctx:        ctx,
		httpServer: httpServer,
	}
}

func (s *Server) ListenAndServe() error {
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
