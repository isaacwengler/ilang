package library

import (
	"errors"
	"ilang/model"
)

type ArrayValue struct {
	value []model.WrappedValue
	state model.ExecutionState
}

func (s *ArrayValue) IsWrappedValue() bool {
	return true
}

func (s *ArrayValue) GetValue() []model.WrappedValue {
	return s.value
}

func (s *ArrayValue) GetState() model.ExecutionState {
	return s.state
}

func (s *ArrayValue) SetState(state model.ExecutionState) {
	s.state = state
}

func (s *ArrayValue) PrintValue() string {
	res := "["
	for i, item := range s.value {
		res += item.PrintValue()
		if i != len(s.value)-1 {
			res += ","
		}
	}
	res += "]"
	return res
}

func (s *ArrayValue) Comparison(op string, other model.WrappedValue) model.WrappedValue {
	switch op {
	// TODO: array == comparison? by ptr value or value?
	// ptr value was not as straightforward as I though...
	default:
		err := errors.New("operator '" + op + "' not supported for array type")
		panic(err)
	}
}

func (s *ArrayValue) GetChild(key model.WrappedValue) model.WrappedValue {
	switch key.(type) {
	case *IntValue:
		index := key.(*IntValue).GetValue()
		if index < 0 || index >= int64(len(s.value)) {
			err := errors.New("Array index out of bounds")
			panic(err)
		}
		return s.value[index]
	case *StringValue:
		return makeArrayFunction(s, key.(*StringValue).GetValue())
	default:
		err := errors.New("Array indexing only supported with int and string values")
		panic(err)
	}
}

func (s *ArrayValue) SetChild(key model.WrappedValue, value model.WrappedValue) {
	switch key.(type) {
	case *IntValue:
		index := key.(*IntValue).GetValue()
		if index < 0 || index >= int64(len(s.value)) {
			err := errors.New("Array index out of bounds")
			panic(err)
		}
		s.value[index] = value
	default:
		err := errors.New("Array property assignment only supported with int value")
		panic(err)
	}
}

func NewArrayValue(value []model.WrappedValue) *ArrayValue {
	return &ArrayValue{value, model.DEFAULT}
}
