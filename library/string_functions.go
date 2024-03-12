package library

import (
	"errors"
	"ilang/model"
	"strings"
)

func sLength(s *StringValue, args []model.WrappedValue) model.WrappedValue {
	validateArgsBetween("string.length", 0, 0, len(args))
	return NewIntValue(int64(len(s.GetValue())))
}

func split(s *StringValue, args []model.WrappedValue) model.WrappedValue {
	validateArgsBetween("string.split", 1, 1, len(args))
	delim := extractStringValue("string.split", args[0])
	arr := strings.Split(s.value, delim)
	wrappedArr := make([]model.WrappedValue, len(arr))
	for i, s := range arr {
		wrappedArr[i] = NewStringValue(s)
	}
	return NewArrayValue(wrappedArr)
}

func substring(s *StringValue, args []model.WrappedValue) model.WrappedValue {
	name := "string.substring"
	validateArgsBetween(name, 2, 2, len(args))
	first := extractIntValue(name, args[0])
	second := extractIntValue(name, args[1])
	if first < 0 || first >= int64(len(s.value)) || second < 0 || second > int64(len(s.value)) || first > second {
		err := errors.New("String index out of bounds")
		panic(err)
	}
	return NewStringValue(s.value[first:second])
}

func makeStringFunction(s *StringValue, name string) model.WrappedValue {
	displayName := "array." + name
	switch name {
	case "length":
		return NewLibFunctionValue(func(args []model.WrappedValue) model.WrappedValue {
			return sLength(s, args)
		}, displayName)
	case "split":
		return NewLibFunctionValue(func(args []model.WrappedValue) model.WrappedValue {
			return split(s, args)
		}, displayName)
	case "substring":
		return NewLibFunctionValue(func(args []model.WrappedValue) model.WrappedValue {
			return substring(s, args)
		}, displayName)
	default:
		err := errors.New("String value does not have function: " + displayName)
		panic(err)
	}
}
