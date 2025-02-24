package main

import (
	"encoding/json"
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

	fastServer.app.Get("/account", AccountHandler)

	return fastServer
}

func (s *FiberServer) Test(request *http.Request) (*http.Response, error) {
	return s.app.Test(request)
}

func AccountHandler(c *fiber.Ctx) error {
	body := c.Body()

	var data map[string]interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	id := data["id"]

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
