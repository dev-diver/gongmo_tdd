package store

import (
	"errors"
	"strings"

	"github.com/dev-diver/gongmo/domain"
)

type InMemoryAccountStore struct {
	accounts map[domain.AccountId]int
}

func NewInMemoryAccountStore() *InMemoryAccountStore {
	return &InMemoryAccountStore{
		accounts: make(map[domain.AccountId]int),
	}
}

func (i *InMemoryAccountStore) GetAccount(id domain.AccountId) (int, error) {
	idCopy := domain.AccountId(strings.Clone(string(id)))
	if _, ok := i.accounts[idCopy]; !ok {
		return 0, errors.New("account not found: " + string(id))
	}
	return i.accounts[idCopy], nil
}

func (i *InMemoryAccountStore) StoreAccount(id domain.AccountId, amount int) error {
	idCopy := domain.AccountId(strings.Clone(string(id)))
	i.accounts[idCopy] = amount
	return nil
}
