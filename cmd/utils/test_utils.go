package utils

import (
	"os"
	"runtime"
)

var runtimeGOOS = runtime.GOOS
var osGetwd = os.Getwd
var osHostname = os.Hostname
