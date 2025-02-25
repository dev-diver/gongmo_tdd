package main_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dev-diver/gongmo/domain"
	"github.com/dev-diver/gongmo/driver"
	"github.com/dev-diver/gongmo/specifications"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestAccountServer(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "../../.",
			Dockerfile: "./cmd/httpserver/Dockerfile",
			// set to false if you want less spam, but this is helpful if you're having troubles
			PrintBuildLog: true,
		},
		ExposedPorts: []string{"8080:8080"},
		WaitingFor:   wait.ForHTTP("/").WithPort("8080"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})

	driver := driver.Driver{BaseURL: "http://localhost:8080"}
	specifications.AccountRetrievalSpec(t, driver, domain.AccountId("1"), 0, errors.New("failed to get account: account not found: 1"))
}
