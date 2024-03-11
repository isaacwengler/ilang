package library

import (
	"errors"
	"ilang/model"
)

type NullValue struct {
}

func (s *NullValue) IsWrappedValue() bool {
	return true
}

func (s *NullValue) PrintValue() string {
	return "null"
}

func (s *NullValue) Comparison(op string, other model.WrappedValue) model.WrappedValue {
	switch op {
	case "==":
		return s.Equals(other)
	case "!=":
		return NewBooleanValue(!s.Equals(other).GetValue())
	default:
		err := errors.New("operator '" + op + "' not supported for null type")
		panic(err)
	}
}

func (s *NullValue) Equals(other model.WrappedValue) *BooleanValue {
	switch other.(type) {
	case *NullValue:
		return NewBooleanValue(true)
	default:
		return NewBooleanValue(false)
	}
}

func (s *NullValue) GetChild(key model.WrappedValue) model.WrappedValue {
	err := errors.New("null value does not have children")
	panic(err)
}

func (s *NullValue) SetChild(key model.WrappedValue, value model.WrappedValue) {
	err := errors.New("Cannot set child property on null value")
	panic(err)
}

func NewNullValue() *NullValue {
	return &NullValue{}
}
