package types

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

func (s BooleanValue) PrintValue() any {
	return s.value
}

func NewBooleanValue(value bool) *BooleanValue {
	children := make(map[string]WrappedValue)
	return &BooleanValue{value, &children}
}