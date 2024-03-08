package library

import (
	"errors"
	"fmt"
	"ilang/logger"
	"strconv"
)

type MapValue struct {
	value map[any]WrappedValue
}

func (s *MapValue) IsWrappedValue() bool {
	return true
}

func (s *MapValue) PrintValue() string {
	res := "{"
	for key, item := range s.value {
		switch key.(type) {
		case int64:
			res += strconv.FormatInt(key.(int64), 10)
		case string:
			res += `"`
			res += key.(string)
			res += `"`
		default:
			logger.Warn("Unexpected key type '", fmt.Sprintf("%T", key), "' when calling PrintValue on a Map")
		}
		res += ":"
		res += item.PrintValue()
		res += ","
	}

	// cut off trailing comma
	res = res[:len(res)-1]
	res += "}"
	return res
}

func (s *MapValue) Comparison(op string, other WrappedValue) *BooleanValue {
	switch op {
	default:
		err := errors.New("operator '" + op + "' not supported for map type")
		panic(err)
	}
}

func (s *MapValue) GetChild(key WrappedValue) WrappedValue {
	switch key.(type) {
	case *IntValue:
		index := key.(*IntValue).GetValue()
		val, ok := s.value[index]
		if !ok {
			return NewNullValue()
		}
		return val
	case *StringValue:
		f, ok := makeMapFunction(s, key.(*StringValue).GetValue())
		if ok {
			return f
		}

		index := key.(*StringValue).GetValue()
		val, ok := s.value[index]
		if !ok {
			return NewNullValue()
		}
		return val
	default:
		err := errors.New("Map indexing only supported with int and string values")
		panic(err)
	}
}

func (s *MapValue) SetChild(key WrappedValue, value WrappedValue) {
	set(s, []WrappedValue{key, value})
}

func NewMapValue(value map[any]WrappedValue) *MapValue {
	return &MapValue{value}
}
