package main

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	app *fiber.App
}

type AccountId int

func NewFiberServer() *FiberServer {
	fastServer := &FiberServer{
		app: fiber.New(),
	}
	return fastServer
}

type Controller interface {
	Register(app *fiber.App)
}

func (f *FiberServer) Register(controller Controller) {
	controller.Register(f.app)
}

func (s *FiberServer) Test(request *http.Request) (*http.Response, error) {
	clonedRequest := request.Clone(context.Background())
	return s.app.Test(clonedRequest)
}

func (s *FiberServer) ListenAndServe(port string) error {
	return s.app.Listen(":" + port)
}

func (s *FiberServer) Shutdown() error {
	return s.app.Shutdown()
}
