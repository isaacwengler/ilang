package visitor

import (
	"ilang/library"
)

func (v *Visitor) registerGlobalFunctions() {
    // TODO create these when used, not on init
	v.scope.Set("print", library.NewLibFunctionValue(library.Print, "print"))
	v.scope.Set("println", library.NewLibFunctionValue(library.Println, "println"))
}
