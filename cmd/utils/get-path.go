package utils

import (
	"fmt"
)

const pathIcon = ""

func GetPath() string {
	path, err := osGetwd()

	if err != nil {
		return "‼"
	}

	return fmt.Sprintf("%s %s", pathIcon, path)
}
