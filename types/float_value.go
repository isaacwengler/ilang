package types

import "strconv"

type FloatValue struct {
	value float64
    children *map[string]WrappedValue
}

func (s FloatValue) IsWrappedValue() bool {
	return true
}

func (s FloatValue) GetChildren() *map[string]WrappedValue {
	return s.children
}

func (s FloatValue) PrintValue() string {
    return strconv.FormatFloat(s.value, 'f', -1, 64)
}

func NewFloatValue(value float64) *FloatValue {
	children := make(map[string]WrappedValue)
	return &FloatValue{value, &children}
}
