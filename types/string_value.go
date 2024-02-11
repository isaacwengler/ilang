package types

type StringValue struct {
	value    string
	children *map[string]WrappedValue
}

func (s StringValue) IsWrappedValue() bool {
	return true
}

func (s StringValue) GetChildren() *map[string]WrappedValue {
	return s.children
}

func (s StringValue) PrintValue() any {
	return s.value
}

func NewStringValue(value string) *StringValue {
	children := make(map[string]WrappedValue)
	return &StringValue{value, &children}
}
