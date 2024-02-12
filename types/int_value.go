package types

import (
	"errors"
	"strconv"
)

type IntValue struct {
	value    int64
	children *map[string]WrappedValue
}

func (s IntValue) IsWrappedValue() bool {
	return true
}

func (s IntValue) GetChildren() *map[string]WrappedValue {
	return s.children
}

func (s IntValue) PrintValue() string {
	return strconv.FormatInt(s.value, 10)
}

func (s IntValue) GetValue() int64 {
	return s.value
}

func (s IntValue) Comparison(op string, other WrappedValue) *BooleanValue {
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

func (s IntValue) Equals(other WrappedValue) *BooleanValue {
	switch other.(type) {
	case *IntValue:
		return NewBooleanValue(s.GetValue() == other.(*IntValue).GetValue())
	case *FloatValue:
		return NewBooleanValue(float64(s.GetValue()) == other.(*FloatValue).GetValue())
	default:
		return NewBooleanValue(false)
	}
}

func (s IntValue) LessThan(other WrappedValue) *BooleanValue {
	switch other.(type) {
	case *IntValue:
		return NewBooleanValue(s.GetValue() < other.(*IntValue).GetValue())
	case *FloatValue:
		return NewBooleanValue(float64(s.GetValue()) < other.(*FloatValue).GetValue())
	default:
		err := errors.New("Comparison not supported between int and non-number type")
		panic(err)
	}
}

func (s IntValue) GreaterThan(other WrappedValue) *BooleanValue {
	switch other.(type) {
	case *IntValue:
		return NewBooleanValue(s.GetValue() > other.(*IntValue).GetValue())
	case *FloatValue:
		return NewBooleanValue(float64(s.GetValue()) > other.(*FloatValue).GetValue())
	default:
		err := errors.New("Comparison not supported between int and non-number type")
		panic(err)
	}
}

func NewIntValue(value int64) *IntValue {
	children := make(map[string]WrappedValue)
	return &IntValue{value, &children}
}
