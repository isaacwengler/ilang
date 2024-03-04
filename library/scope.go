package library

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

func (s *Scope) ReassignVariable(name string, val WrappedValue, children []WrappedValue) {
	curr, ok := s.GetVar(name)
	if ok {
		if len(children) == 0 {
			s.SetVar(name, val)
			return
		}

        for i, c := range children {
            if i == len(children) - 1 {
                curr.SetChild(c, val)
                return
            }
            curr = curr.GetChild(c)
        }
	}

	if s.ParentScope == nil {
		err := errors.New("Attempting to reassign undefined variable '" + name + "'")
		panic(err)
	}
	s.ParentScope.ReassignVariable(name, val, children)
}
