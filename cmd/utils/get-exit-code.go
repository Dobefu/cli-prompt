package utils

import (
	"fmt"
	"os"

	"github.com/Dobefu/cli-prompt/cmd/structs"
)

func GetExitCode() (string, structs.ColorRGB) {
	if len(os.Args) < 2 {
		return "✘ -", structs.ColorRGB{R: 255, G: 0, B: 0}
	}

	if os.Args[1] != "0" {
		return fmt.Sprintf("✘ %s", os.Args[1]), structs.ColorRGB{R: 255, G: 0, B: 0}
	}

	return "✔", structs.ColorRGB{R: 0, G: 200, B: 0}
}
