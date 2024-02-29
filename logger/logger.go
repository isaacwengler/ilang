package logger

import (
	"fmt"
)

var debugLog = false

func SetDebugMode(flag bool) {
	debugLog = flag
}

func Print(v ...any) {
	fmt.Println(v...)
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
