package main_test

import (
	"testing"

	"github.com/dev-diver/gongmo/domain"
	"github.com/dev-diver/gongmo/driver"
	"github.com/dev-diver/gongmo/specifications"
)

func TestAccountServer(t *testing.T) {
	driver := driver.Driver{BaseURL: "http://localhost:8080"}
	specifications.AccountRetrievalSpec(t, driver, domain.AccountId("1"), 100, nil)
}
