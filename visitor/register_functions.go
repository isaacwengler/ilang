package visitor

import (
	"ilang/library"
)

func (v *Visitor) registerGlobalFunctions() {
    // TODO create these when used, not on init
	v.scope.SetVar("print", library.NewLibFunctionValue(library.Print, "print"))
	v.scope.SetVar("println", library.NewLibFunctionValue(library.Println, "println"))
}
