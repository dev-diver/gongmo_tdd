package main

import "github.com/gofiber/fiber/v2"

type FiberServer struct {
	app *fiber.App
}

func NewFiberServer() *FiberServer {
	fastServer := &FiberServer{
		app: fiber.New(),
	}

	fastServer.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	return fastServer
}

func (s *FiberServer) ListenAndServe(port string) error {
	return s.app.Listen(":" + port)
}

func (s *FiberServer) Shutdown() error {
	return s.app.Shutdown()
}
