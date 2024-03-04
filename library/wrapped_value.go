package library

type WrappedValue interface {
	IsWrappedValue() bool
	PrintValue() string
	Comparison(op string, other WrappedValue) *BooleanValue
	GetChild(key WrappedValue) WrappedValue
	SetChild(key WrappedValue, value WrappedValue)
}
