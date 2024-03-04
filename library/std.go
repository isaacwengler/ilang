package library

import (
	"fmt"
)

func Print(args []WrappedValue) WrappedValue {
	for i, arg := range args {
		switch arg.(type) {
		case *StringValue:
			fmt.Print(arg.(*StringValue).GetValue())
		default:
			fmt.Print(arg.PrintValue())
		}

		if i != len(args)-1 {
			fmt.Print(" ")
		}
	}
	return NewNullValue()
}

func Println(args []WrappedValue) WrappedValue {
	Print(args)
	fmt.Println()
	return NewNullValue()
}
