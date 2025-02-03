package utils

import (
	"fmt"
	"os"
	"strings"
)

func GetHostname() string {
	username := os.Getenv("LOGNAME")
	hostname, err := osHostname()

	if err != nil {
		return username
	}

	hostname = strings.Replace(hostname, ".local", "", 1)

	return fmt.Sprintf("%s@%s", username, hostname)
}
