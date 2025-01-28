package utils

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupGetOSTest() (cleanup func()) {
	runtimeGOOS = "test-os"

	return func() {
		runtimeGOOS = runtime.GOOS
	}
}

func TestGetOSSuccess(t *testing.T) {
	cleanup := setupGetOSTest()
	defer cleanup()

	os := GetOS()
	assert.Equal(t, "test-os", os)
}
