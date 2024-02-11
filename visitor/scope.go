package visitor

import (
	"errors"
)

func newScope(parentScope *scope) *scope {
	m := make(map[string]any)
	return &scope{parentScope, &m}
}

type scope struct {
	parentScope *scope
	variables   *map[string]any
}

func (s *scope) setVar(name string, val any) {
	(*s.variables)[name] = val
}

func (s *scope) getVar(name string) (any, bool) {
	val, ok := (*s.variables)[name]
	return val, ok
}

func (s *scope) resolveVariable(name string) any {
    val, ok := s.getVar(name)
    if ok {
        return val
    }

    if s.parentScope == nil {
        err := errors.New("Undefined variable " + name) 
        panic(err)
    }
    
    return s.parentScope.resolveVariable(name)
}

