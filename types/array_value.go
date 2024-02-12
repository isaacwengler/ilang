package types

import (
	"errors"
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
    // TODO: array == comparison? by ptr value or value?
    // ptr value was not as straightforward as I though...
	default:
		err := errors.New("operator '" + op + "' not supported for array type")
		panic(err)
	}
}

func NewArrayValue(value []WrappedValue) *ArrayValue {
	children := make(map[string]WrappedValue)
	return &ArrayValue{value, &children}
}
