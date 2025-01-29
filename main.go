package main

import (
	"fmt"

	"github.com/Dobefu/cli-prompt/cmd/structs"
	"github.com/Dobefu/cli-prompt/cmd/utils"
)

type section struct {
	content string
	fg      structs.ColorRGB
	bg      structs.ColorRGB
}

func main() {
	var sections = []section{
		{
			utils.GetOSIcon(),
			structs.ColorRGB{R: 0, G: 0, B: 255},
			structs.ColorRGB{R: 200, G: 200, B: 200},
		},
		{
			utils.GetPath(),
			structs.ColorRGB{R: 255, G: 255, B: 255},
			structs.ColorRGB{R: 0, G: 0, B: 200},
		},
	}

	numSections := len(sections)

	for i := 0; i < numSections; i++ {
		utils.PrintSection(sections[i].content, sections[i].fg, sections[i].bg)

		if i == numSections-1 {
			fmt.Println("")
		} else {
			fmt.Print(" ")
		}
	}
}
