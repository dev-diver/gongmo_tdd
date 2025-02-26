package specifications_test

import (
	"errors"
	"testing"

	"github.com/dev-diver/gongmo/domain"
	"github.com/dev-diver/gongmo/service"
	"github.com/dev-diver/gongmo/specifications"
	"github.com/dev-diver/gongmo/store"
)

func TestAccountRetrievalSpec(t *testing.T) {

	store := store.NewInMemoryAccountStore()
	svc := service.NewAccountService(store)

	specifications.AccountRetrievalSpec(t,
		specifications.GetAccountAdapter(svc.GetAccount),
		domain.AccountId("1"),
		0,
		errors.New("failed to get account"))
}
