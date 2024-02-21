package types

import (
	"errors"
)

func NewScope(parentScope *Scope) *Scope {
	m := make(map[string]WrappedValue)
	return &Scope{parentScope, &m}
}

type Scope struct {
	ParentScope *Scope
	variables   *map[string]WrappedValue
}

func (s *Scope) SetVar(name string, val WrappedValue) {
	(*s.variables)[name] = val
}

func (s *Scope) GetVar(name string) (WrappedValue, bool) {
	val, ok := (*s.variables)[name]
	return val, ok
}

func (s *Scope) ResolveVariable(name string) WrappedValue {
	val, ok := s.GetVar(name)
	if ok {
		return val
	}

	if s.ParentScope == nil {
		err := errors.New("Undefined reference to variable '" + name + "'")
		panic(err)
	}

	return s.ParentScope.ResolveVariable(name)
}

func (s *Scope) ReassignVariable(name string, val WrappedValue) {
	_, ok := s.GetVar(name)
	if ok {
		s.SetVar(name, val)
		return
	}

	if s.ParentScope == nil {
		err := errors.New("Attempting to reassign undefined variable '" + name + "'")
		panic(err)
	}
	s.ParentScope.ReassignVariable(name, val)
}
