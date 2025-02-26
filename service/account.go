package service

import (
	"fmt"

	"github.com/dev-diver/gongmo/domain"
)

type AccountStore interface {
	GetAccount(id domain.AccountId) (int, error)
	StoreAccount(id domain.AccountId, amount int) error
}

type AccountService struct {
	store AccountStore
}

func NewAccountService(store AccountStore) *AccountService {
	return &AccountService{
		store: store,
	}
}

func (a *AccountService) GetAccount(id domain.AccountId) (int, error) {
	account, err := a.store.GetAccount(id)
	if err != nil {
		return 0, fmt.Errorf("failed to get account: %w", err)
	}
	return account, nil
}

func (a *AccountService) StoreAccount(id domain.AccountId, amount int) error {
	return a.store.StoreAccount(id, amount)
}
