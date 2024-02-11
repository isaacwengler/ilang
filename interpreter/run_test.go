package interpreter

import (
    "testing"
)

func TestStringAssign(t *testing.T) {
    input := 
    `let var = "expected";
    var;`

    res := RunIlang(input)
    if res != "expected" {
        t.Fatalf("unable to assign string literal")
    }
}
