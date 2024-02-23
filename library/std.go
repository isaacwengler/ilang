package library

import (
	"fmt"
	"ilang/types"
)

func Print(args []types.WrappedValue) types.WrappedValue {
	for i, arg := range args {
		switch arg.(type) {
		case *types.StringValue:
			fmt.Print(arg.(*types.StringValue).GetValue())
		default:
			fmt.Print(arg.PrintValue())
		}

		if i != len(args)-1 {
			fmt.Print(" ")
		}
	}
	return types.NewNullValue()
}

func Println(args []types.WrappedValue) types.WrappedValue {
	Print(args)
	fmt.Println()
	return types.NewNullValue()
}
