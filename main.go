package main

import (
	"github.com/Dobefu/cli-prompt/cmd/structs"
	"github.com/Dobefu/cli-prompt/cmd/utils"
)

func main() {
	utils.PrintSection(
		utils.GetOSIcon(),
		structs.ColorRGB{R: 0, G: 0, B: 255},
		structs.ColorRGB{R: 200, G: 200, B: 200},
	)
}
