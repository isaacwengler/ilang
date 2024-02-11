package types

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

func (s IntValue) PrintValue() any {
	return s.value
}

func NewIntValue(value int64) *IntValue {
	children := make(map[string]WrappedValue)
	return &IntValue{value, &children}
}
