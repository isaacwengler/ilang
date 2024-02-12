package types

type WrappedValue interface {
	IsWrappedValue() bool
	GetChildren() *map[string]WrappedValue
	PrintValue() string
	Comparison(op string, other WrappedValue) *BooleanValue
}
