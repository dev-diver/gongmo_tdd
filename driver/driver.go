package driver

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/dev-diver/gongmo/domain"
)

type Driver struct {
	BaseURL string
}

func (d Driver) GetAccount(id domain.AccountId) (int, error) {
	response, err := http.Get(fmt.Sprintf("%s/account/%s", d.BaseURL, id))
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	if response.StatusCode == http.StatusNotFound {
		return 0, errors.New(string(body))
	}

	return strconv.Atoi(string(body))
}

func (d Driver) StoreAccount(id domain.AccountId, amount int) error {
	response, err := http.Post(fmt.Sprintf("%s/account/%s", d.BaseURL, id), "application/json", bytes.NewBuffer([]byte(fmt.Sprintf(`{"amount": %d}`, amount))))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("expected status accepted, got %d", response.StatusCode)
	}

	return nil
}
