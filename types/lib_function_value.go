package types

import (
	"errors"
)

type LibFunctionValue struct {
	f    func(args []WrappedValue) WrappedValue
	name string
}

func (s LibFunctionValue) IsWrappedValue() bool {
	return true
}

func (s LibFunctionValue) PrintValue() string {
	return "func " + s.name
}

func (s LibFunctionValue) Comparison(op string, other WrappedValue) *BooleanValue {
	switch op {
	default:
		err := errors.New("operator '" + op + "' not supported for library function type")
		panic(err)
	}
}

func (s LibFunctionValue) Call(args []WrappedValue) WrappedValue {
	return s.f(args)
}

func NewLibFunctionValue(f func(args []WrappedValue) WrappedValue, name string) *LibFunctionValue {
	return &LibFunctionValue{f, name}
}
