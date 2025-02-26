package specifications

import "github.com/dev-diver/gongmo/domain"

type GetAccountAdapter func(id domain.AccountId) (int, error)

func (a GetAccountAdapter) GetAccount(id domain.AccountId) (int, error) {
	return a(id)
}
