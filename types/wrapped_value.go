package types

type WrappedValue interface {
	IsWrappedValue() bool
	PrintValue() string
	Comparison(op string, other WrappedValue) *BooleanValue
}
