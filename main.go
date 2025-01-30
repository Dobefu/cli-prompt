package main

import (
	"fmt"

	"github.com/Dobefu/cli-prompt/cmd/structs"
	"github.com/Dobefu/cli-prompt/cmd/utils"
	"golang.org/x/term"
)

type section struct {
	content string
	fg      structs.ColorRGB
	bg      structs.ColorRGB
}

func main() {
	termWidth, _, err := term.GetSize(0)

	if err != nil {
		return
	}

	charsRemaining := termWidth

	var sectionsLeft = []section{
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

	numSectionsLeft := len(sectionsLeft)

	for i := 0; i < numSectionsLeft; i++ {
		section, sectionLen := utils.SprintSection(
			sectionsLeft[i].content,
			sectionsLeft[i].fg,
			sectionsLeft[i].bg,
		)

		charsRemaining -= sectionLen

		if i == numSectionsLeft-1 {
			section = fmt.Sprintf("%s\n", section)
		} else {
			section = fmt.Sprintf("%s ", section)
			charsRemaining -= 1
		}

		fmt.Print(section)
	}
}
