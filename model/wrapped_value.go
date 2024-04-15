package model

type WrappedValue interface {
	IsWrappedValue() bool
	PrintValue() string
	Comparison(op string, other WrappedValue) WrappedValue
	GetChild(key WrappedValue) WrappedValue
	SetChild(key WrappedValue, value WrappedValue)
    GetState() ExecutionState
    SetState(s ExecutionState)
}
