package library

import (
	"errors"
	"ilang/model"
)

type StringValue struct {
	value string
    state model.ExecutionState
}

func (s *StringValue) IsWrappedValue() bool {
	return true
}

func (s *StringValue) PrintValue() string {
	return `"` + s.value + `"`
}

func (s *StringValue) GetValue() string {
	return s.value
}

func (s *StringValue) GetState() model.ExecutionState {
	return s.state
}

func (s *StringValue) SetState(state model.ExecutionState) {
	s.state = state
}

func (s *StringValue) Comparison(op string, other model.WrappedValue) model.WrappedValue {
	switch op {
	case "==":
		return s.Equals(other)
	case "!=":
		return NewBooleanValue(!s.Equals(other).GetValue())
	case "<":
		return s.LessThan(other)
	case ">":
		return s.GreaterThan(other)
	case "<=":
		return NewBooleanValue(s.LessThan(other).GetValue() || s.Equals(other).GetValue())
	case ">=":
		return NewBooleanValue(s.GreaterThan(other).GetValue() || s.Equals(other).GetValue())
	default:
		err := errors.New("Unsupported comparison operator '" + op + "'")
		panic(err)
	}
}

func (s *StringValue) Equals(other model.WrappedValue) *BooleanValue {
	switch other.(type) {
	case *StringValue:
		return NewBooleanValue(s.GetValue() == other.(*StringValue).GetValue())
	default:
		return NewBooleanValue(false)
	}
}

func (s *StringValue) LessThan(other model.WrappedValue) *BooleanValue {
	switch other.(type) {
	case *StringValue:
		return NewBooleanValue(s.GetValue() < other.(*StringValue).GetValue())
	default:
		err := errors.New("Comparison not supported between string and other type")
		panic(err)
	}
}

func (s *StringValue) GreaterThan(other model.WrappedValue) *BooleanValue {
	switch other.(type) {
	case *StringValue:
		return NewBooleanValue(s.GetValue() > other.(*StringValue).GetValue())
	default:
		err := errors.New("Comparison not supported between string and other type")
		panic(err)
	}
}

func (s *StringValue) Arithmetic(op string, other model.WrappedValue) model.WrappedValue {
	switch other.(type) {
	case *StringValue:
		otherVal := other.(*StringValue).GetValue()
		switch op {
		case "+":
			return NewStringValue(s.value + otherVal)
		default:
			err := errors.New("Unsupported arithmetic operator '" + op + "' for type string")
			panic(err)
		}
	default:
		err := errors.New("Arithmetic not supported between string other than +")
		panic(err)
	}
}

func (s *StringValue) GetChild(key model.WrappedValue) model.WrappedValue {
	switch key.(type) {
	case *IntValue:
		index := key.(*IntValue).GetValue()
		if index < 0 || index >= int64(len(s.value)) {
			err := errors.New("String index out of bounds")
			panic(err)
		}
		return NewStringValue(string(s.value[index]))
	case *StringValue:
		return makeStringFunction(s, key.(*StringValue).GetValue())
	default:
		err := errors.New("string indexing only supported with int and string values")
		panic(err)
	}
}

func (s *StringValue) SetChild(key model.WrappedValue, value model.WrappedValue) {
	err := errors.New("Cannot set child property on string value")
	panic(err)
}

func NewStringValue(value string) *StringValue {
	return &StringValue{value, model.DEFAULT}
}
