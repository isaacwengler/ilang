package logger

import (
	"fmt"
	"os"
)

var debugLog = isDebugFlagOn()

func isDebugFlagOn() bool {
	for _, flag := range os.Args[1:] {
		if flag == "-d" || flag == "--debug" {
			return true
		}
	}
	return false
}

func Debug(v ...any) {
	if debugLog {
		fmt.Print("DEBUG: ")
		fmt.Println(v...)
	}
}

func Warn(v ...any) {
	fmt.Print("WARN: ")
	fmt.Println(v...)
}

func Error(v ...any) {
	fmt.Print("ERROR: ")
	fmt.Println(v...)
}
