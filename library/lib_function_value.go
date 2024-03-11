package library

import (
	"errors"
	"ilang/model"
)

type LibFunctionValue struct {
	f    func(args []model.WrappedValue) model.WrappedValue
	name string
}

func (s *LibFunctionValue) IsWrappedValue() bool {
	return true
}

func (s *LibFunctionValue) PrintValue() string {
	return "func " + s.name
}

func (s *LibFunctionValue) Comparison(op string, other model.WrappedValue) model.WrappedValue {
	switch op {
	default:
		err := errors.New("operator '" + op + "' not supported for library function type")
		panic(err)
	}
}

func (s *LibFunctionValue) Call(args []model.WrappedValue) model.WrappedValue {
	return s.f(args)
}

func (s *LibFunctionValue) GetChild(key model.WrappedValue) model.WrappedValue {
	err := errors.New("function value does not have children")
	panic(err)
}

func (s *LibFunctionValue) SetChild(key model.WrappedValue, value model.WrappedValue) {
	err := errors.New("Cannot set child property on function value")
	panic(err)
}

func NewLibFunctionValue(f func(args []model.WrappedValue) model.WrappedValue, name string) *LibFunctionValue {
	return &LibFunctionValue{f, name}
}
