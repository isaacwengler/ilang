package types

import (
	"errors"
	"strconv"
)

type BooleanValue struct {
	value bool
}

func (s BooleanValue) IsWrappedValue() bool {
	return true
}

func (s BooleanValue) PrintValue() string {
	return strconv.FormatBool(s.value)
}

func (s BooleanValue) GetValue() bool {
	return s.value
}

func (s BooleanValue) Comparison(op string, other WrappedValue) *BooleanValue {
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

func (s BooleanValue) Equals(other WrappedValue) *BooleanValue {
	switch other.(type) {
	case *BooleanValue:
		return NewBooleanValue(s.GetValue() == other.(*BooleanValue).GetValue())
	default:
		return NewBooleanValue(false)
	}
}

func NewBooleanValue(value bool) *BooleanValue {
	return &BooleanValue{value}
}
