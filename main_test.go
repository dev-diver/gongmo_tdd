package main

import (
	"testing"
	"time"

	"github.com/quii/go-graceful-shutdown/acceptancetests"
	"github.com/quii/go-graceful-shutdown/assert"
)

const (
	port = "8080"
	url  = "http://localhost:" + port
)

func TestGracefulShutdown(t *testing.T) {
	cleanup, sendInterrupt, err := acceptancetests.LaunchTestProgram(port)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(cleanup)

	// shutdown 전에 서버가 작동하는지 확인
	assert.CanGet(t, url)

	// 요청을 보내고 응답을 기다리기 전에 SIGTERM을 보냄.
	time.AfterFunc(10*time.Millisecond, func() {
		assert.NoError(t, sendInterrupt())
	})
	// graceful shutdown이 없으면 이 테스트가 실패함.
	assert.CanGet(t, url)

	time.Sleep(20 * time.Millisecond)
	// 인터럽트 후, 서버가 종료되고 더 이상 요청이 작동하지 않음.
	assert.CantGet(t, url)
}
