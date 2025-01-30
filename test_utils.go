package main

import (
	"runtime"

	"golang.org/x/term"
)

var runtimeGOOS = runtime.GOOS
var termGetSize = term.GetSize
