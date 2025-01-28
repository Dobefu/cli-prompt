package utils

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/Dobefu/cli-prompt/cmd/structs"
	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func setupPrintSectionTest() (r *os.File, w *os.File, cleanup func()) {
	orig := os.Stdout
	r, w, _ = os.Pipe()
	os.Stdout = w

	// Force the use of colours.
	color.NoColor = false

	return r, w, func() {
		os.Stdout = orig
		w.Close()
	}
}

func TestPrintSectionSuccess(t *testing.T) {
	r, w, cleanup := setupPrintSectionTest()
	defer cleanup()

	PrintSection(
		"test",
		structs.ColorRGB{R: 0, G: 0, B: 100},
		structs.ColorRGB{R: 200, G: 50, B: 0},
	)

	w.Close()
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	assert.NoError(t, err)

	assert.Contains(t, buf.String(), bracketOpen)
	assert.Contains(t, buf.String(), bracketClose)
	assert.Contains(t, buf.String(), "test")
}
