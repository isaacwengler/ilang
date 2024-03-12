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

func TestMapIndex(t *testing.T) {
	input := `let val = {
        hi: "hello"
    };
    val.hi;`

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, `"hello"`)
}

func TestMapComputedIndex(t *testing.T) {
	input := `let val = {
        hi: "hello"
    };
    let index = "hi";
    val[index];`

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, `"hello"`)
}

func TestMaxIntIndex(t *testing.T) {
	input := `let val = {
        [1]: 2
    };
    val[1];`

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, `2`)
}

func TestMapIndexInvalidKey(t *testing.T) {
	input := `let val = {
        [1]: "hi"
    };
    return val[1.0];`

	assertPanic(t, func() { interpreter.RunIlang(input) })
}

func TestMapGet(t *testing.T) {
	input := `let val = {
        hi: "hello"
    };
    val.get("hi");`

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, `"hello"`)
}

func TestMapGetDefault(t *testing.T) {
	input := `let val = {
        hi: "hello"
    };
    val.get("hii", "oink");`

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, `"oink"`)
}

func TestMapSet(t *testing.T) {
	input := `let val = {
        hi: "hello"
    };
    val.set("hi", "oink");
    val.hi;`

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, `"oink"`)
}

func TestMapSetIndex(t *testing.T) {
	input := `let val = {
        hi: "hello"
    };
    val.hi = "oink";
    val.hi;`

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, `"oink"`)
}

func TestMapSetComputed(t *testing.T) {
	input := `let val = {
        hi: "hello"
    };
    let p = "hi";
    val[p] = "oink";
    val.hi;`

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, `"oink"`)
}

func TestMapKeys(t *testing.T) {
	input := `let val = {
        hi: "hello"
    };
    val.keys(); `

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, `["hi"]`)
}

func TestMapValues(t *testing.T) {
	input := `let val = {
        hi: "hello"
    };
    val.values(); `

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, `["hello"]`)
}
