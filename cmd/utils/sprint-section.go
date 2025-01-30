package utils

import (
	"fmt"
	"unicode/utf8"

	"github.com/Dobefu/cli-prompt/cmd/structs"
	"github.com/fatih/color"
)

const bracketOpen = ""
const bracketClose = ""

func SprintSection(content string, fg structs.ColorRGB, bg structs.ColorRGB) (section string, sectionLen int) {
	section = fmt.Sprintf(
		"%s%s%s",
		color.RGB(bg.R, bg.G, bg.B).Sprint(bracketOpen),
		color.RGB(fg.R, fg.G, fg.B).AddBgRGB(bg.R, bg.G, bg.B).Sprint(content),
		color.RGB(bg.R, bg.G, bg.B).Sprint(bracketClose),
	)

	sectionLen = utf8.RuneCountInString(fmt.Sprintf(
		"%s%s%s",
		bracketOpen,
		content,
		bracketClose,
	))

	return section, sectionLen
}
