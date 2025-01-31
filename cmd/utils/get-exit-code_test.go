package utils

import (
	"os"
	"testing"

	"github.com/Dobefu/cli-prompt/cmd/structs"
	"github.com/stretchr/testify/assert"
)

func setupGetExitCodeTest() (cleanup func()) {
	oldOsArgs := os.Args

	return func() {
		os.Args = oldOsArgs
	}
}

func TestGetExitCodeSuccess(t *testing.T) {
	cleanup := setupGetExitCodeTest()
	defer cleanup()

	os.Args = []string{os.Args[0], "0"}

	exitCode, exitCodeFg := GetExitCode()

	assert.Equal(t, "✔", exitCode)
	assert.Equal(t, structs.ColorRGB{R: 0, G: 200, B: 0}, exitCodeFg)
}

func TestGetExitCodeErrNonZero(t *testing.T) {
	cleanup := setupGetExitCodeTest()
	defer cleanup()

	os.Args = []string{os.Args[0], "1"}

	exitCode, exitCodeFg := GetExitCode()

	assert.Equal(t, "✘ 1", exitCode)
	assert.Equal(t, structs.ColorRGB{R: 255, G: 0, B: 0}, exitCodeFg)
}

func TestGetExitCodeErrNoArg(t *testing.T) {
	cleanup := setupGetExitCodeTest()
	defer cleanup()

	os.Args = []string{os.Args[0]}

	exitCode, exitCodeFg := GetExitCode()

	assert.Equal(t, "✘ -", exitCode)
	assert.Equal(t, structs.ColorRGB{R: 255, G: 0, B: 0}, exitCodeFg)
}
