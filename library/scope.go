package library

import (
	"errors"
    "ilang/model"
)

func NewScope(parentScope *Scope) *Scope {
	m := make(map[string]model.WrappedValue)
	return &Scope{parentScope, &m}
}

type Scope struct {
	ParentScope *Scope
	variables   *map[string]model.WrappedValue
}

func (s *Scope) SetVar(name string, val model.WrappedValue) {
	(*s.variables)[name] = val
}

func (s *Scope) GetVar(name string) (model.WrappedValue, bool) {
	val, ok := (*s.variables)[name]
	return val, ok
}

func (s *Scope) ResolveVariable(name string) model.WrappedValue {
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

func (s *Scope) ReassignVariable(name string, val model.WrappedValue, children []model.WrappedValue) {
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
