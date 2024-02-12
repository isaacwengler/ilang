package types

import (
	"errors"
	"ilang/logger"
)

type ArrayValue struct {
	value    []WrappedValue
	children *map[string]WrappedValue
}

func (s ArrayValue) IsWrappedValue() bool {
	return true
}

func (s ArrayValue) GetChildren() *map[string]WrappedValue {
	return s.children
}

func (s ArrayValue) PrintValue() string {
	res := "["
	for i, item := range s.value {
		res += item.PrintValue()
		if i != len(s.value)-1 {
			res += ","
		}
	}
	res += "]"
	return res
}

func (s ArrayValue) Comparison(op string, other WrappedValue) *BooleanValue {
	switch op {
	case "==":
		return s.Equals(other)
	case "!=":
		return NewBooleanValue(!s.Equals(other).GetValue())
	default:
		err := errors.New("operator '" + op + "' not supported for array type")
		panic(err)
	}
}

func (s ArrayValue) Equals(other WrappedValue) *BooleanValue {
	return NewBooleanValue(&s == other)
}

func NewArrayValue(value []WrappedValue) *ArrayValue {
	logger.Debug("initialized array value")
	children := make(map[string]WrappedValue)
	return &ArrayValue{value, &children}
}
