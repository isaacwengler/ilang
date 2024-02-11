package visitor

import (
	"errors"
	"ilang/types"
)

func newScope(parentScope *scope) *scope {
	m := make(map[string]types.WrappedValue)
	return &scope{parentScope, &m}
}

type scope struct {
	parentScope *scope
	variables   *map[string]types.WrappedValue
}

func (s *scope) setVar(name string, val types.WrappedValue) {
	(*s.variables)[name] = val
}

func (s *scope) getVar(name string) (types.WrappedValue, bool) {
	val, ok := (*s.variables)[name]
	return val, ok
}

func (s *scope) resolveVariable(name string) types.WrappedValue {
	val, ok := s.getVar(name)
	if ok {
		return val
	}

	if s.parentScope == nil {
		err := errors.New("Undefined reference to variable '" + name + "'")
		panic(err)
	}

	return s.parentScope.resolveVariable(name)
}

func (s *scope) reassignVariable(name string, val types.WrappedValue) {
	_, ok := s.getVar(name)
	if ok {
		s.setVar(name, val)
		return
	}

	if s.parentScope == nil {
		err := errors.New("Attempting to reassign undefined variable '" + name + "'")
		panic(err)
	}
	s.parentScope.reassignVariable(name, val)
}
