package utils

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupGetHostnameTest() (cleanup func()) {
	oldLogname := os.Getenv("LOGNAME")

	osHostname = func() (name string, err error) {
		return "test-hostname", nil
	}

	return func() {
		os.Setenv("LOGNAME", oldLogname)
		osHostname = os.Hostname
	}
}

func TestGetHostnameSuccess(t *testing.T) {
	cleanup := setupGetHostnameTest()
	defer cleanup()

	os.Setenv("LOGNAME", "test-user")

	os := GetHostname()
	assert.Equal(t, "test-user@test-hostname", os)
}

func TestGetHostnameErrHostname(t *testing.T) {
	cleanup := setupGetHostnameTest()
	defer cleanup()

	os.Setenv("LOGNAME", "test-user")
	osHostname = func() (name string, err error) {
		return "", errors.New("cannot get the hostname")
	}

	os := GetHostname()
	assert.Equal(t, "test-user", os)
}
