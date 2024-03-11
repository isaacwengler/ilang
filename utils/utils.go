package utils

import (
	"errors"
	"fmt"
)

func ValidateArgsBetween(name string, minArgs int, maxArgs int, recieved int) {
	if recieved < minArgs || recieved > maxArgs {
		var args string
		if minArgs == maxArgs {
			args = fmt.Sprint(minArgs)
		} else {
			args = fmt.Sprint(minArgs) + "-" + fmt.Sprint(maxArgs)
		}
		err := errors.New("Function " + name + " expected " + args + " arguments but got " + fmt.Sprint(recieved))
		panic(err)
	}
}

