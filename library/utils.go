package library

import (
	"errors"
	"fmt"
	"ilang/model"
)

func extractIntValue(name string, v model.WrappedValue) int64 {
	switch v.(type) {
	case *IntValue:
		return v.(*IntValue).GetValue()
	default:
		err := errors.New("Function " + name + " expected int value argument")
		panic(err)
	}
}

func extractStringValue(name string, v model.WrappedValue) string {
	switch v.(type) {
	case *StringValue:
		return v.(*StringValue).GetValue()
	default:
		err := errors.New("Function " + name + " expected string value argument")
		panic(err)
	}
}

func validateArgsBetween(name string, minArgs int, maxArgs int, recieved int) {
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
