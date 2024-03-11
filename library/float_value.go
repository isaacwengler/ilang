package library

import (
	"errors"
	"math"
	"strconv"
    "ilang/model"
)

type FloatValue struct {
	value float64
}

func (s *FloatValue) IsWrappedValue() bool {
	return true
}

func (s *FloatValue) PrintValue() string {
	return strconv.FormatFloat(s.value, 'f', -1, 64)
}

func (s *FloatValue) GetValue() float64 {
	return s.value
}

func (s *FloatValue) Comparison(op string, other model.WrappedValue) model.WrappedValue {
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

func (s *FloatValue) Equals(other model.WrappedValue) *BooleanValue {
	switch other.(type) {
	case *IntValue:
		return NewBooleanValue(s.GetValue() == float64(other.(*IntValue).GetValue()))
	case *FloatValue:
		return NewBooleanValue(s.GetValue() == other.(*FloatValue).GetValue())
	default:
		return NewBooleanValue(false)
	}
}

func (s *FloatValue) LessThan(other model.WrappedValue) *BooleanValue {
	switch other.(type) {
	case *IntValue:
		return NewBooleanValue(s.GetValue() < float64(other.(*IntValue).GetValue()))
	case *FloatValue:
		return NewBooleanValue(s.GetValue() < other.(*FloatValue).GetValue())
	default:
		err := errors.New("Comparison not supported between float and non-number")
		panic(err)
	}
}

func (s *FloatValue) GreaterThan(other model.WrappedValue) *BooleanValue {
	switch other.(type) {
	case *IntValue:
		return NewBooleanValue(s.GetValue() > float64(other.(*IntValue).GetValue()))
	case *FloatValue:
		return NewBooleanValue(s.GetValue() > other.(*FloatValue).GetValue())
	default:
		err := errors.New("Comparison not supported between float and non-number type")
		panic(err)
	}
}

func (s *FloatValue) Arithmetic(op string, other model.WrappedValue) model.WrappedValue {
	switch other.(type) {
	case *IntValue:
		otherVal := other.(*IntValue).GetValue()
		switch op {
		case "+":
			return NewFloatValue(s.value + float64(otherVal))
		case "-":
			return NewFloatValue(s.value - float64(otherVal))
		case "*":
			return NewFloatValue(s.value * float64(otherVal))
		case "/":
			return NewFloatValue(s.value / float64(otherVal))
		case "**":
			return NewFloatValue(math.Pow(s.value, float64(otherVal)))
		default:
			err := errors.New("Unsupported arithmetic operator '" + op + "' for type float and int")
			panic(err)
		}
	case *FloatValue:
		otherVal := other.(*FloatValue).GetValue()
		switch op {
		case "+":
			return NewFloatValue(s.value + otherVal)
		case "-":
			return NewFloatValue(s.value - otherVal)
		case "*":
			return NewFloatValue(s.value * otherVal)
		case "/":
			return NewFloatValue(s.value / otherVal)
		case "**":
			return NewFloatValue(math.Pow(s.value, otherVal))
		default:
			err := errors.New("Unsupported arithmetic operator '" + op + "' for type float and float")
			panic(err)
		}
	default:
		err := errors.New("Arithmetic not supported between float and non-number type")
		panic(err)
	}
}

func (s *FloatValue) GetChild(key model.WrappedValue) model.WrappedValue {
	err := errors.New("float value does not have children")
	panic(err)
}

func (s *FloatValue) SetChild(key model.WrappedValue, value model.WrappedValue) {
	err := errors.New("Cannot set child property on float value")
	panic(err)
}

func NewFloatValue(value float64) *FloatValue {
	return &FloatValue{value}
}
