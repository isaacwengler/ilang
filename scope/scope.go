package scope

import (
	"errors"
    "ilang/model"
)

func New(parentScope *Scope) *Scope {
	m := make(map[string]model.WrappedValue)
	return &Scope{parentScope, &m}
}

type Scope struct {
	ParentScope *Scope
	variables   *map[string]model.WrappedValue
}

func (s *Scope) Set(name string, val model.WrappedValue) {
	(*s.variables)[name] = val
}

func (s *Scope) Get(name string) (model.WrappedValue, bool) {
	val, ok := (*s.variables)[name]
	return val, ok
}

func (s *Scope) Resolve(name string) model.WrappedValue {
	val, ok := s.Get(name)
	if ok {
		return val
	}

	if s.ParentScope == nil {
		err := errors.New("Undefined reference to variable '" + name + "'")
		panic(err)
	}

	return s.ParentScope.Resolve(name)
}

func (s *Scope) Reassign(name string, val model.WrappedValue, children []model.WrappedValue) {
	curr, ok := s.Get(name)
	if ok {
		if len(children) == 0 {
			s.Set(name, val)
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
	s.ParentScope.Reassign(name, val, children)
}
