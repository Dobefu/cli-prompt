package utils

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupGetCwdTest() (cleanup func()) {
	osGetwd = func() (dir string, err error) { return "test-dir", nil }

	return func() {
		osGetwd = os.Getwd
	}
}

func TestGetCwdSuccess(t *testing.T) {
	cleanup := setupGetCwdTest()
	defer cleanup()

	path := GetCwd()
	assert.Equal(t, fmt.Sprintf("%s test-dir", dirIcon), path)
}

func TestGetCwdErr(t *testing.T) {
	cleanup := setupGetCwdTest()
	defer cleanup()

	osGetwd = func() (dir string, err error) { return "", errors.New("") }

	path := GetCwd()
	assert.Equal(t, "‼", path)
}
