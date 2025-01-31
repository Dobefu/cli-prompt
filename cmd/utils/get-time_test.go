package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setupGetTimeTest() (cleanup func()) {
	return func() {
		// No cleanup needed.
	}
}

func TestGetTimeSuccess(t *testing.T) {
	cleanup := setupGetTimeTest()
	defer cleanup()

	timeString := GetTime()
	assert.Equal(t, time.Now().Format("15:04:05 "), timeString)
}
