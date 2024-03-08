package test

import (
	"ilang/interpreter"
	"testing"
)

func TestMap(t *testing.T) {
	input := `let val = {
        hi: "hello"
    };
    return val;`

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, `{"hi":"hello"}`)
}

func TestMapComputedProperty(t *testing.T) {
	input := `let val = {
        [1]: 2
    };
    return val;`

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, `{1:2}`)
}

func TestMapInvalidKey(t *testing.T) {
	input := `let val = {
        [1.2]: "hi"
    };
    return val;`

	assertPanic(t, func() { interpreter.RunIlang(input) })
}
