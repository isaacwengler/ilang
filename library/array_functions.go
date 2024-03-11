package library

import (
	"errors"
	"ilang/model"
	"ilang/utils"
)

func length(arr *ArrayValue, args []model.WrappedValue) model.WrappedValue {
	utils.ValidateArgsBetween("array.length", 0, 0, len(args))
	return NewIntValue(int64(len(arr.GetValue())))
}

func push(arr *ArrayValue, args []model.WrappedValue) model.WrappedValue {
	arr.value = append(arr.GetValue(), args...)
	return arr
}

func pop(arr *ArrayValue, args []model.WrappedValue) model.WrappedValue {
	utils.ValidateArgsBetween("array.pop", 0, 0, len(args))
	l := len(arr.value)
	popped := arr.value[l-1]
	arr.value = arr.value[:l-1]
	return popped
}

func makeArrayFunction(a *ArrayValue, name string) model.WrappedValue {
	displayName := "array." + name
	switch name {
	case "length":
		return NewLibFunctionValue(func(args []model.WrappedValue) model.WrappedValue {
			return length(a, args)
		}, displayName)
	case "push":
		return NewLibFunctionValue(func(args []model.WrappedValue) model.WrappedValue {
			return push(a, args)
		}, displayName)
	case "pop":
		return NewLibFunctionValue(func(args []model.WrappedValue) model.WrappedValue {
			return pop(a, args)
		}, displayName)
	default:
		err := errors.New("Array value does not have property: " + displayName)
		panic(err)
	}
}
