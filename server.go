package main

import (
	"fmt"
	"net/http"
	"strconv"

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
	StoreAccount(id AccountId, amount int) error
}

type AccountService struct {
	store AccountStore
}

type AccountController struct {
	service AccountService
}

func (a *AccountController) Register(app *fiber.App) {
	app.Get("/account/:id", a.AccountHandler)
	app.Post("/account/:id", a.PostAccountHandler)
}

func NewAccountController(store AccountStore) Controller {
	return &AccountController{
		service: AccountService{
			store: store,
		},
	}
}

func (a *AccountService) GetAccount(id AccountId) (int, error) {
	account, err := a.store.GetAccount(id)
	if err != nil {
		return 0, fmt.Errorf("account not found: %w", err)
	}
	return account, nil
}

func (a *AccountService) StoreAccount(id AccountId, amount int) error {
	return a.store.StoreAccount(id, amount)
}

func (a *AccountController) AccountHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	account, err := a.service.GetAccount(AccountId(id))
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.SendString(fmt.Sprintf("%d", account))
}

func (a *AccountController) PostAccountHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	amount := c.Body()
	amountInt, err := strconv.Atoi(string(amount))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	a.service.StoreAccount(AccountId(id), amountInt)
	return c.SendStatus(fiber.StatusAccepted)
}

func (s *FiberServer) ListenAndServe(port string) error {
	return s.app.Listen(":" + port)
}

func (s *FiberServer) Shutdown() error {
	return s.app.Shutdown()
}
