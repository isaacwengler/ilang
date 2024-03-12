package library

import (
	"errors"
    "ilang/model"
)

func get(m *MapValue, args []model.WrappedValue) model.WrappedValue {
	validateArgsBetween("map.get", 1, 2, len(args))
	property := args[0]
	var fallback model.WrappedValue
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

func set(m *MapValue, args []model.WrappedValue) {
	validateArgsBetween("map.set", 2, 2, len(args))
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

func keys(m *MapValue, args []model.WrappedValue) model.WrappedValue {
	validateArgsBetween("map.keys", 0, 0, len(args))
	keys := make([]model.WrappedValue, len(m.value))
	i := 0
	for k := range m.value {
		switch k.(type) {
		case int64:
			keys[i] = NewIntValue(k.(int64)) 
		case string:
			keys[i] = NewStringValue(k.(string)) 
		default:
			err := errors.New("Map indexing only supported with int and string values. Found key with different type")
			panic(err)
		}
		i++
	}
	return NewArrayValue(keys)
}

func values(m *MapValue, args []model.WrappedValue) model.WrappedValue {
	validateArgsBetween("map.values", 0, 0, len(args))
	values := make([]model.WrappedValue, len(m.value))
	i := 0
	for _, v := range m.value {
		values[i] = v
		i++
	}
	return NewArrayValue(values)
}

func makeMapFunction(m *MapValue, name string) (model.WrappedValue, bool) {
	displayName := "map." + name
	switch name {
	case "get":
		return NewLibFunctionValue(func(args []model.WrappedValue) model.WrappedValue {
			return get(m, args)
		}, displayName), true
	case "set":
		return NewLibFunctionValue(func(args []model.WrappedValue) model.WrappedValue {
			set(m, args)
            return nil
		}, displayName), true
	case "keys":
		return NewLibFunctionValue(func(args []model.WrappedValue) model.WrappedValue {
			return keys(m, args)
		}, displayName), true
	case "values":
		return NewLibFunctionValue(func(args []model.WrappedValue) model.WrappedValue {
			return values(m, args)
		}, displayName), true
	default:
		return nil, false
	}
}
