package library

import "errors"

type NullValue struct {
}

func (s NullValue) IsWrappedValue() bool {
	return true
}

func (s NullValue) PrintValue() string {
	return "null"
}

func (s NullValue) Comparison(op string, other WrappedValue) *BooleanValue {
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

func (s NullValue) Equals(other WrappedValue) *BooleanValue {
	switch other.(type) {
	case *NullValue:
		return NewBooleanValue(true)
	default:
		return NewBooleanValue(false)
	}
}

func (s NullValue) GetChild(key WrappedValue) WrappedValue {
    return nil
}

func NewNullValue() *NullValue {
	return &NullValue{}
}