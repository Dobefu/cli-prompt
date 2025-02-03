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
		{
			utils.GetTime(),
			structs.ColorRGB{R: 0, G: 0, B: 0},
			structs.ColorRGB{R: 200, G: 200, B: 200},
		},
	}

	for i := 0; i < len(sectionsLeft); i++ {
		section, sectionLen := utils.SprintSection(
			sectionsLeft[i].content,
			sectionsLeft[i].fg,
			sectionsLeft[i].bg,
		)

		charsRemaining -= sectionLen

		if i != len(sectionsLeft)-1 {
			section = fmt.Sprintf("%s ", section)
			charsRemaining -= 1
		}

		fmt.Print(section)
	}

	alignCursorForSectionsRight(charsRemaining, sectionsRight)

	for i := 0; i < len(sectionsRight); i++ {
		section, sectionLen := utils.SprintSection(
			sectionsRight[i].content,
			sectionsRight[i].fg,
			sectionsRight[i].bg,
		)

		charsRemaining -= sectionLen

		if i != len(sectionsRight)-1 {
			section = fmt.Sprintf("%s ", section)
			charsRemaining -= 1
		}

		fmt.Print(section)
	}

	fmt.Println("")

	hostname, _ := utils.SprintSection(
		utils.GetHostname(),
		structs.ColorRGB{R: 0, G: 0, B: 0},
		structs.ColorRGB{R: 0, G: 200, B: 0},
	)

	fmt.Print(hostname)
	fmt.Print(" ")
}

func getsectionsNumChars(sections []section) (numChars int) {
	for i := 0; i < len(sections); i++ {
		_, sectionLen := utils.SprintSection(
			sections[i].content,
			sections[i].fg,
			sections[i].bg,
		)

		numChars += sectionLen

		if i != len(sections)-1 {
			numChars += 1
		}
	}

	return numChars
}

func alignCursorForSectionsRight(charsRemaining int, sectionsRight []section) {
	numCharsSectionsRight := getsectionsNumChars(sectionsRight)

	if charsRemaining-numCharsSectionsRight < 0 {
		for i := 0; i < (charsRemaining-numCharsSectionsRight)*-1; i++ {
			fmt.Print("\b")
		}
	} else {
		for i := 0; i < (charsRemaining - numCharsSectionsRight); i++ {
			fmt.Print(" ")
		}
	}

}
