package types

import "strconv"

type BooleanValue struct {
	value    bool
	children *map[string]WrappedValue
}

func (s BooleanValue) IsWrappedValue() bool {
	return true
}

func (s BooleanValue) GetChildren() *map[string]WrappedValue {
	return s.children
}

func (s BooleanValue) PrintValue() string {
	return strconv.FormatBool(s.value)
}

func (s BooleanValue) GetValue() bool {
    return s.value
}

func NewBooleanValue(value bool) *BooleanValue {
	children := make(map[string]WrappedValue)
	return &BooleanValue{value, &children}
}
