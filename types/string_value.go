package types

import "errors"

type StringValue struct {
	value    string
}

func (s StringValue) IsWrappedValue() bool {
	return true
}

func (s StringValue) PrintValue() string {
	return `"` + s.value + `"`
}

func (s StringValue) GetValue() string {
	return s.value
}

func (s StringValue) Comparison(op string, other WrappedValue) *BooleanValue {
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

func (s StringValue) Equals(other WrappedValue) *BooleanValue {
	switch other.(type) {
	case *StringValue:
		return NewBooleanValue(s.GetValue() == other.(*StringValue).GetValue())
	default:
		return NewBooleanValue(false)
	}
}

func (s StringValue) LessThan(other WrappedValue) *BooleanValue {
	switch other.(type) {
	case *StringValue:
		return NewBooleanValue(s.GetValue() < other.(*StringValue).GetValue())
	default:
		err := errors.New("Comparison not supported between string and other type")
		panic(err)
	}
}

func (s StringValue) GreaterThan(other WrappedValue) *BooleanValue {
	switch other.(type) {
	case *StringValue:
		return NewBooleanValue(s.GetValue() > other.(*StringValue).GetValue())
	default:
		err := errors.New("Comparison not supported between string and other type")
		panic(err)
	}
}

func NewStringValue(value string) *StringValue {
	return &StringValue{value}
}
