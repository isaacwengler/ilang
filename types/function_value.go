package types

import (
	"errors"
	parser "ilang/generated"
)

type FunctionValue struct {
	Args         []string
	Def          parser.IScopeBodyContext
	ClosureScope *Scope
}

func (s FunctionValue) IsWrappedValue() bool {
	return true
}

func (s FunctionValue) PrintValue() string {
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

func (s FunctionValue) Comparison(op string, other WrappedValue) *BooleanValue {
	switch op {
	default:
		err := errors.New("operator '" + op + "' not supported for function type")
		panic(err)
	}
}

func NewFunctionValue(args []string, def parser.IScopeBodyContext, outsideScope *Scope) *FunctionValue {
	return &FunctionValue{args, def, outsideScope}
}
