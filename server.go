package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	app *fiber.App
}

func NewFiberServer() *FiberServer {
	fastServer := &FiberServer{
		app: fiber.New(),
	}

	fastServer.app.Get("/account/:id", AccountHandler)

	return fastServer
}

func (s *FiberServer) Test(request *http.Request) (*http.Response, error) {
	return s.app.Test(request)
}

func AccountHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "1" {
		return c.SendString("0")
	} else if id == "2" {
		return c.SendString("1")
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func (s *FiberServer) ListenAndServe(port string) error {
	return s.app.Listen(":" + port)
}

func (s *FiberServer) Shutdown() error {
	return s.app.Shutdown()
}
