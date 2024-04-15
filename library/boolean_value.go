package library

import (
	"errors"
	"ilang/model"
	"strconv"
)

type BooleanValue struct {
	value bool
	state model.ExecutionState
}

func (s *BooleanValue) IsWrappedValue() bool {
	return true
}

func (s *BooleanValue) PrintValue() string {
	return strconv.FormatBool(s.value)
}

func (s *BooleanValue) GetValue() bool {
	return s.value
}

func (s *BooleanValue) GetState() model.ExecutionState {
	return s.state
}

func (s *BooleanValue) SetState(state model.ExecutionState) {
	s.state = state
}

func (s *BooleanValue) Comparison(op string, other model.WrappedValue) model.WrappedValue {
	switch op {
	case "==":
		return s.Equals(other)
	case "!=":
		return NewBooleanValue(!s.Equals(other).GetValue())
	default:
		err := errors.New("operator '" + op + "' not supported for boolean type")
		panic(err)
	}
}

func (s *BooleanValue) Equals(other model.WrappedValue) *BooleanValue {
	switch other.(type) {
	case *BooleanValue:
		return NewBooleanValue(s.GetValue() == other.(*BooleanValue).GetValue())
	default:
		return NewBooleanValue(false)
	}
}

func (s *BooleanValue) GetChild(key model.WrappedValue) model.WrappedValue {
	err := errors.New("boolean value does not have children")
	panic(err)
}

func (s *BooleanValue) SetChild(key model.WrappedValue, value model.WrappedValue) {
	err := errors.New("Cannot set child property on boolean value")
	panic(err)
}

func NewBooleanValue(value bool) *BooleanValue {
	return &BooleanValue{value, model.DEFAULT}
}
