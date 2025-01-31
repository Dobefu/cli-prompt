package main

import (
	"fmt"

	"github.com/Dobefu/cli-prompt/cmd/structs"
	"github.com/Dobefu/cli-prompt/cmd/utils"
	"github.com/fatih/color"
)

type section struct {
	content string
	fg      structs.ColorRGB
	bg      structs.ColorRGB
}

func main() {
	// Force the use of colours.
	color.NoColor = false

	termWidth, _, err := termGetSize(0)

	if err != nil {
		return
	}

	exitCode, exitCodeFg := utils.GetExitCode()

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

	var sectionsRight = []section{
		{
			exitCode,
			exitCodeFg,
			structs.ColorRGB{},
		},
	}

	numSectionsLeft := len(sectionsLeft)
	numCharsSectionsLeft := 0
	numSectionsRight := len(sectionsRight)
	numCharsSectionsRight := 0

	for i := 0; i < numSectionsLeft; i++ {
		_, sectionLen := utils.SprintSection(
			sectionsLeft[i].content,
			sectionsLeft[i].fg,
			sectionsLeft[i].bg,
		)

		numCharsSectionsLeft += sectionLen
	}

	for i := 0; i < numSectionsRight; i++ {
		_, sectionLen := utils.SprintSection(
			sectionsRight[i].content,
			sectionsRight[i].fg,
			sectionsRight[i].bg,
		)

		numCharsSectionsRight += sectionLen
	}

	for i := 0; i < numSectionsLeft; i++ {
		section, sectionLen := utils.SprintSection(
			sectionsLeft[i].content,
			sectionsLeft[i].fg,
			sectionsLeft[i].bg,
		)

		charsRemaining -= sectionLen

		if i != numSectionsLeft-1 {
			section = fmt.Sprintf("%s ", section)
			charsRemaining -= 1
		}

		fmt.Print(section)
	}

	if charsRemaining-numCharsSectionsRight < 0 {
		for i := 0; i < (charsRemaining-numCharsSectionsRight)*-1; i++ {
			fmt.Print("\b")
		}
	} else {
		for i := 0; i < (charsRemaining - numCharsSectionsRight); i++ {
			fmt.Print(" ")
		}
	}

	for i := 0; i < numSectionsRight; i++ {
		section, sectionLen := utils.SprintSection(
			sectionsRight[i].content,
			sectionsRight[i].fg,
			sectionsRight[i].bg,
		)

		charsRemaining -= sectionLen

		if i != numSectionsRight-1 {
			section = fmt.Sprintf("%s ", section)
			charsRemaining -= 1
		}

		fmt.Print(section)
	}

	fmt.Println("$ ")
}
