package utils

import (
	"errors"
	"fmt"
)

func ValidateArgs(name string, expected int, recieved int) {
	if expected != recieved {
		err := errors.New("Function " + name + " expected " + fmt.Sprint(expected) + " arguments but got " + fmt.Sprint(recieved))
		panic(err)
	}
}
