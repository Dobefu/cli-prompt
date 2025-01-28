package utils

func GetOSIcon() string {
	os := GetOS()

	switch os {
	case "darwin":
		return ""

	case "linux":
		return ""

	default:
		return "?"
	}
}
