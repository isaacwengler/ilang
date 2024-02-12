package types

import (
	"errors"
	"fmt"
	"ilang/logger"
	"strconv"
)

type MapValue struct {
	value    map[any]WrappedValue
	children *map[string]WrappedValue
}

func (s MapValue) IsWrappedValue() bool {
	return true
}

func (s MapValue) GetChildren() *map[string]WrappedValue {
	return s.children
}

func (s MapValue) PrintValue() string {
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

func (s MapValue) Comparison(op string, other WrappedValue) *BooleanValue {
	switch op {
	case "==":
		return s.Equals(other)
	case "!=":
		return NewBooleanValue(!s.Equals(other).GetValue())
	default:
		err := errors.New("operator '" + op + "' not supported for map type")
		panic(err)
	}
}

func (s MapValue) Equals(other WrappedValue) *BooleanValue {
	return NewBooleanValue(&s == other)
}

func NewMapValue(value map[any]WrappedValue) *MapValue {
	children := make(map[string]WrappedValue)
	return &MapValue{value, &children}
}
