package library

import (
	"fmt"
	"ilang/model"
)

func Print(args []model.WrappedValue) model.WrappedValue {
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

func Println(args []model.WrappedValue) model.WrappedValue {
	Print(args)
	fmt.Println()
	return NewNullValue()
}
