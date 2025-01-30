package main

import (
	"errors"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/term"
)

func setupMainTest() (cleanup func()) {
	runtimeGOOS = "test-os"
	termGetSize = func(fd int) (width, height int, err error) {
		return 100, 10, nil
	}

	return func() {
		runtimeGOOS = runtime.GOOS
		termGetSize = term.GetSize
	}
}

func TestMainSuccess(t *testing.T) {
	cleanup := setupMainTest()
	defer cleanup()

	assert.NotPanics(t, func() { main() })
}

func TestMainErrNoTerm(t *testing.T) {
	cleanup := setupMainTest()
	defer cleanup()

	termGetSize = func(fd int) (width, height int, err error) {
		return 0, 0, errors.New("cannot get term size")
	}

	assert.NotPanics(t, func() { main() })
}
