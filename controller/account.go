package controller

import (
	"fmt"
	"strconv"

	"github.com/dev-diver/gongmo/domain"
	"github.com/dev-diver/gongmo/service"
	"github.com/gofiber/fiber/v2"
)

type AccountService interface {
	GetAccount(id domain.AccountId) (int, error)
	StoreAccount(id domain.AccountId, amount int) error
}

type AccountController struct {
	service AccountService
}

func NewAccountController(store service.AccountStore) *AccountController {
	return &AccountController{
		service: service.NewAccountService(store),
	}
}

func (a *AccountController) Register(app *fiber.App) {
	app.Get("/account/:id", a.AccountHandler)
	app.Post("/account/:id", a.PostAccountHandler)
}

func (a *AccountController) AccountHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	account, err := a.service.GetAccount(domain.AccountId(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	return c.SendString(fmt.Sprintf("%d", account))
}

func (a *AccountController) PostAccountHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	amount := c.Body()
	amountInt, err := strconv.Atoi(string(amount))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid amount")
	}
	a.service.StoreAccount(domain.AccountId(id), amountInt)
	return c.SendStatus(fiber.StatusAccepted)
}
