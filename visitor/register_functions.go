package visitor

import (
	"ilang/library"
	"ilang/types"
)

func (v *Visitor) registerFunctions() {
	v.scope.SetVar("print", types.NewLibFunctionValue(library.Print, "print"))
	v.scope.SetVar("println", types.NewLibFunctionValue(library.Println, "println"))
}
