package library

import (
	"errors"
	"ilang/utils"
)

func length(arr *ArrayValue, args []WrappedValue) WrappedValue {
	utils.ValidateArgs("array.length", 0, len(args))
	return NewIntValue(int64(len(arr.GetValue())))
}

func makeArrayFunction(a *ArrayValue, name string) WrappedValue {
	displayName := "array." + name
	switch name {
	case "length":
		return NewLibFunctionValue(func(args []WrappedValue) WrappedValue {
			return length(a, args)
		}, displayName)
	default:
		err := errors.New("Array value does not have property: " + displayName)
		panic(err)
	}
}
