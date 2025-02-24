package main

import (
	"fmt"
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

	return fastServer
}

type Controller interface {
	Register(app *fiber.App)
}

func (f *FiberServer) Register(controller Controller) {
	controller.Register(f.app)
}

func (s *FiberServer) Test(request *http.Request) (*http.Response, error) {
	return s.app.Test(request)
}

type AccountId string

type AccountStore interface {
	GetAccount(id AccountId) (int, error)
}

type AStore struct {
	store map[AccountId]int
}

func NewAccountStore() *AStore {
	store := &AStore{
		store: make(map[AccountId]int),
	}
	store.store[AccountId("1")] = 0
	store.store[AccountId("2")] = 1
	return store
}

func (s *AStore) GetAccount(id AccountId) (int, error) {
	return s.store[id], nil
}

type AccountService struct {
	store AccountStore
}

type AccountController struct {
	service AccountService
}

func (a *AccountController) Register(app *fiber.App) {
	app.Get("/account/:id", a.AccountHandler)
}

func NewAccountController(store AccountStore) Controller {
	return &AccountController{
		service: AccountService{
			store: store,
		},
	}
}

func (a *AccountService) GetAccount(id AccountId) (int, error) {
	return a.store.GetAccount(id)
}

func (a *AccountController) AccountHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	account, err := a.service.GetAccount(AccountId(id))
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.SendString(fmt.Sprintf("%d", account))
}

func (s *FiberServer) ListenAndServe(port string) error {
	return s.app.Listen(":" + port)
}

func (s *FiberServer) Shutdown() error {
	return s.app.Shutdown()
}
