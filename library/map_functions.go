package library

import (
	"errors"
	"ilang/utils"
)

func get(m *MapValue, args []WrappedValue) WrappedValue {
	utils.ValidateArgsBetween("map.get", 1, 2, len(args))
	property := args[0]
	var fallback WrappedValue
	if len(args) == 1 {
		fallback = NewNullValue()
	} else {
		fallback = args[1]
	}

	switch property.(type) {
	case *IntValue:
		index := property.(*IntValue).GetValue()
		val, ok := m.value[index]
		if !ok {
			return fallback
		}
		return val
	case *StringValue:
		index := property.(*StringValue).GetValue()
		val, ok := m.value[index]
		if !ok {
			return fallback
		}
		return val
	default:
		err := errors.New("Map indexing only supported with int and string values")
		panic(err)
	}
}

func set(m *MapValue, args []WrappedValue) {
	utils.ValidateArgsBetween("map.set", 2, 2, len(args))
	property := args[0]
	value := args[1]

	switch property.(type) {
	case *IntValue:
		index := property.(*IntValue).GetValue()
		m.value[index] = value
	case *StringValue:
		index := property.(*StringValue).GetValue()
		m.value[index] = value
	default:
		err := errors.New("Map indexing only supported with int and string values")
		panic(err)
	}
}

func keys(m *MapValue, args []WrappedValue) WrappedValue {
	utils.ValidateArgsBetween("map.keys", 0, 0, len(args))
	keys := make([]WrappedValue, len(m.value))
	i := 0
	for k := range m.value {
		switch k.(type) {
		case *IntValue:
			keys[i] = k.(*IntValue)
		case *StringValue:
			keys[i] = k.(*StringValue)
		default:
			err := errors.New("Map indexing only supported with int and string values. Found key with different type")
			panic(err)
		}
		i++
	}
	return NewArrayValue(keys)
}

func values(m *MapValue, args []WrappedValue) WrappedValue {
	utils.ValidateArgsBetween("map.values", 0, 0, len(args))
	values := make([]WrappedValue, len(m.value))
	i := 0
	for _, v := range m.value {
		values[i] = v
		i++
	}
	return NewArrayValue(values)
}

func makeMapFunction(m *MapValue, name string) (WrappedValue, bool) {
	displayName := "map." + name
	switch name {
	case "get":
		return NewLibFunctionValue(func(args []WrappedValue) WrappedValue {
			return get(m, args)
		}, displayName), true
	case "set":
		return NewLibFunctionValue(func(args []WrappedValue) WrappedValue {
			set(m, args)
            return nil
		}, displayName), true
	case "keys":
		return NewLibFunctionValue(func(args []WrappedValue) WrappedValue {
			return keys(m, args)
		}, displayName), true
	case "values":
		return NewLibFunctionValue(func(args []WrappedValue) WrappedValue {
			return values(m, args)
		}, displayName), true
	default:
		return nil, false
	}
}
