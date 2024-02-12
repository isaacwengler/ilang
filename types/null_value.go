package types

import "errors"

type NullValue struct {
    children *map[string]WrappedValue
}

func (s NullValue) IsWrappedValue() bool {
	return true
}

func (s NullValue) GetChildren() *map[string]WrappedValue {
	return s.children
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

func NewNullValue() *NullValue {
	children := make(map[string]WrappedValue)
	return &NullValue{&children}
}
