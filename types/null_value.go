package types

type NullValue struct {
    children *map[string]WrappedValue
}

func (s NullValue) IsWrappedValue() bool {
	return true
}

func (s NullValue) GetChildren() *map[string]WrappedValue {
	return s.children
}

func (s NullValue) PrintValue() any {
    return "null" 
}

func NewNullValue() *NullValue {
	children := make(map[string]WrappedValue)
	return &NullValue{&children}
}
