package utils

import (
	"testing"

	"github.com/Dobefu/cli-prompt/cmd/structs"
	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func setupSprintSectionTest() (cleanup func()) {
	// Force the use of colours.
	color.NoColor = false

	return func() {
		// No cleanup required.
	}
}

func TestSprintSectionSuccess(t *testing.T) {
	cleanup := setupSprintSectionTest()
	defer cleanup()

	section := SprintSection(
		"test",
		structs.ColorRGB{R: 0, G: 0, B: 100},
		structs.ColorRGB{R: 200, G: 50, B: 0},
	)

	assert.Contains(t, section, bracketOpen)
	assert.Contains(t, section, bracketClose)
	assert.Contains(t, section, "test")
}
