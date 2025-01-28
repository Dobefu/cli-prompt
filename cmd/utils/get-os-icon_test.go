package utils

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupGetOSIconTest() (cleanup func()) {
	runtimeGOOS = "test-os"

	return func() {
		runtimeGOOS = runtime.GOOS
	}
}

func TestGetOSIconUnknown(t *testing.T) {
	cleanup := setupGetOSIconTest()
	defer cleanup()

	icon := GetOSIcon()
	assert.Equal(t, "?", icon)
}

func TestGetOSIconMacos(t *testing.T) {
	cleanup := setupGetOSIconTest()
	defer cleanup()

	runtimeGOOS = "darwin"

	icon := GetOSIcon()
	assert.Equal(t, "", icon)
}

func TestGetOSIconLinux(t *testing.T) {
	cleanup := setupGetOSIconTest()
	defer cleanup()

	runtimeGOOS = "linux"

	icon := GetOSIcon()
	assert.Equal(t, "", icon)
}
