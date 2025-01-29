package main

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupMainTest() (cleanup func()) {
	runtimeGOOS = "test-os"

	return func() {
		runtimeGOOS = runtime.GOOS
	}
}

func TestMain(t *testing.T) {
	cleanup := setupMainTest()
	defer cleanup()

	assert.NotPanics(t, func() { main() })
}
