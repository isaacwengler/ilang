package visitor

import (
	"ilang/library"
	library1 "ilang/library"
)

func (v *Visitor) registerGlobalFunctions() {
	v.scope.SetVar("print", library1.NewLibFunctionValue(library.Print, "print"))
	v.scope.SetVar("println", library1.NewLibFunctionValue(library.Println, "println"))
}
