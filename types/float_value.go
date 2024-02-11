package types

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

func (s FloatValue) PrintValue() any {
	return s.value
}

func NewFloatValue(value float64) *FloatValue {
	children := make(map[string]WrappedValue)
	return &FloatValue{value, &children}
}
