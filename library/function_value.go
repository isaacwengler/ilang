package library

import (
	"errors"
	parser "ilang/generated"
	"ilang/model"
)

type FunctionValue struct {
	Args         []string
	Def          parser.IScopeBodyContext
	ClosureScope *Scope
}

func (s *FunctionValue) IsWrappedValue() bool {
	return true
}

func (s *FunctionValue) PrintValue() string {
	res := "func("
	for i, arg := range s.Args {
		res += arg
		if i != len(s.Args)-1 {
			res += ","
		}
	}
	res += ")"
	res += ";"
	return res
}

func (s *FunctionValue) Comparison(op string, other model.WrappedValue) model.WrappedValue {
	switch op {
	default:
		err := errors.New("operator '" + op + "' not supported for function type")
		panic(err)
	}
}

func (s *FunctionValue) GetChild(key model.WrappedValue) model.WrappedValue {
	err := errors.New("function value does not have children")
	panic(err)
}

func (s *FunctionValue) SetChild(key model.WrappedValue, value model.WrappedValue) {
	err := errors.New("Cannot set child property on function value")
	panic(err)
}

func NewFunctionValue(args []string, def parser.IScopeBodyContext, outsideScope *Scope) *FunctionValue {
	return &FunctionValue{args, def, outsideScope}
}
