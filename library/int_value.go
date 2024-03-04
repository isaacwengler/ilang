package library

import (
	"errors"
	"math"
	"strconv"
)

type IntValue struct {
	value int64
}

func (s IntValue) IsWrappedValue() bool {
	return true
}

func (s IntValue) PrintValue() string {
	return strconv.FormatInt(s.value, 10)
}

func (s IntValue) GetValue() int64 {
	return s.value
}

func (s IntValue) Comparison(op string, other WrappedValue) *BooleanValue {
	switch op {
	case "==":
		return s.Equals(other)
	case "!=":
		return NewBooleanValue(!s.Equals(other).GetValue())
	case "<":
		return s.LessThan(other)
	case ">":
		return s.GreaterThan(other)
	case "<=":
		return NewBooleanValue(s.LessThan(other).GetValue() || s.Equals(other).GetValue())
	case ">=":
		return NewBooleanValue(s.GreaterThan(other).GetValue() || s.Equals(other).GetValue())
	default:
		err := errors.New("Unsupported comparison operator '" + op + "'")
		panic(err)
	}
}

func (s IntValue) Equals(other WrappedValue) *BooleanValue {
	switch other.(type) {
	case *IntValue:
		return NewBooleanValue(s.GetValue() == other.(*IntValue).GetValue())
	case *FloatValue:
		return NewBooleanValue(float64(s.GetValue()) == other.(*FloatValue).GetValue())
	default:
		return NewBooleanValue(false)
	}
}

func (s IntValue) LessThan(other WrappedValue) *BooleanValue {
	switch other.(type) {
	case *IntValue:
		return NewBooleanValue(s.GetValue() < other.(*IntValue).GetValue())
	case *FloatValue:
		return NewBooleanValue(float64(s.GetValue()) < other.(*FloatValue).GetValue())
	default:
		err := errors.New("Comparison not supported between int and non-number type")
		panic(err)
	}
}

func (s IntValue) GreaterThan(other WrappedValue) *BooleanValue {
	switch other.(type) {
	case *IntValue:
		return NewBooleanValue(s.GetValue() > other.(*IntValue).GetValue())
	case *FloatValue:
		return NewBooleanValue(float64(s.GetValue()) > other.(*FloatValue).GetValue())
	default:
		err := errors.New("Comparison not supported between int and non-number type")
		panic(err)
	}
}

func (s IntValue) Arithmetic(op string, other WrappedValue) WrappedValue {
	switch other.(type) {
	case *IntValue:
		otherVal := other.(*IntValue).GetValue()
		switch op {
		case "+":
			return NewIntValue(s.value + otherVal)
		case "-":
			return NewIntValue(s.value - otherVal)
		case "*":
			return NewIntValue(s.value * otherVal)
		case "/":
			return NewFloatValue(float64(s.value) / float64(otherVal))
		case "%":
			return NewIntValue(s.value % otherVal)
		case "//":
			return NewIntValue(s.value / otherVal)
		case "**":
			return NewFloatValue(math.Pow(float64(s.value), float64(otherVal)))
		default:
			err := errors.New("Unsupported arithmetic operator '" + op + "' for type int and int")
			panic(err)
		}
	case *FloatValue:
		otherVal := other.(*FloatValue).GetValue()
		switch op {
		case "+":
			return NewFloatValue(float64(s.value) + otherVal)
		case "-":
			return NewFloatValue(float64(s.value) - otherVal)
		case "*":
			return NewFloatValue(float64(s.value) * otherVal)
		case "/":
			return NewFloatValue(float64(s.value) / otherVal)
		case "**":
			return NewFloatValue(math.Pow(float64(s.value), otherVal))
		default:
			err := errors.New("Unsupported arithmetic operator '" + op + "' for type int and float")
			panic(err)
		}
	default:
		err := errors.New("Arithmetic not supported between int and non-number type")
		panic(err)
	}
}

func (s IntValue) GetChild(key WrappedValue) WrappedValue {
    return nil
}

func NewIntValue(value int64) *IntValue {
	return &IntValue{value}
}
