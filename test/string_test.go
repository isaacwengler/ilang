package test

import (
	"ilang/interpreter"
	"testing"
)

func TestStringAssign(t *testing.T) {
	input := `let var = "expected";
    var;`

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, `"expected"`)
}

func TestEqualsComparisonString(t *testing.T) {
	input := `"hello" == "hello";`

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, "true")
}

func TestEqualsComparisonStringDifferent(t *testing.T) {
	input := `"hello" == "hello2";`

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, "false")
}

func TestStringConcat(t *testing.T) {
	input := `
    let a = "hi " + "world";
    a;
    `

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, `"hi world"`)
}

func TestStringLength(t *testing.T) {
	input := `
    "hello".length();
    `

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, `5`)
}

func TestStringSplit(t *testing.T) {
	input := `
    "hello world".split(" ");
    `

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, `["hello","world"]`)
}

func TestStringSubstring(t *testing.T) {
	input := `
    "hello world".substring(0, 5);
    `

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, `"hello"`)
}

func TestStringIndex(t *testing.T) {
	input := `
    let s = "hello there";
    s[2];
    `

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, `"l"`)
}
