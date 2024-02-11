package types

type WrappedValue interface {
	IsWrappedValue() bool
	GetChildren() *map[string]WrappedValue
	PrintValue() string
}
