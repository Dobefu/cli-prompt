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
			utils.GetCwd(),
			structs.ColorRGB{R: 255, G: 255, B: 255},
			structs.ColorRGB{R: 0, G: 0, B: 200},
		},
	}

	numSections := len(sections)

	for i := 0; i < numSections; i++ {
		section := utils.SprintSection(sections[i].content, sections[i].fg, sections[i].bg)

		if i == numSections-1 {
			section = fmt.Sprintf("%s\n", section)
		} else {
			section = fmt.Sprintf("%s ", section)
		}

		fmt.Print(section)
	}
}
