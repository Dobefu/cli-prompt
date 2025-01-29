package utils

import (
	"fmt"
)

const dirIcon = ""

func GetCwd() string {
	path, err := osGetwd()

	if err != nil {
		return "‼"
	}

	return fmt.Sprintf("%s %s", dirIcon, path)
}
