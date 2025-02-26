package specifications

import (
	"testing"

	"github.com/dev-diver/gongmo/domain"
	"github.com/stretchr/testify/assert"
)

type Accounter interface {
	GetAccount(id domain.AccountId) (int, error)
	StoreAccount(id domain.AccountId, amount int) error
}

func AccountRetrievalSpec(t testing.TB, accounter Accounter, id domain.AccountId, expectedAmount int, expectedErr error) {
	got, err := accounter.GetAccount(id)
	if expectedErr != nil {
		assert.Equal(t, expectedErr, err)
	} else {
		assert.NoError(t, err)
		assert.Equal(t, expectedAmount, got)
	}
}

func AccountStorageSpec(t testing.TB, service Accounter, id domain.AccountId, amount int, expectedErr error) {
	err := service.StoreAccount(id, amount)
	if expectedErr != nil {
		assert.Equal(t, expectedErr, err)
	} else {
		assert.NoError(t, err)
	}
}
