package interpreter

import (
	"ilang/test"
	"testing"
)

func TestStringAssign(t *testing.T) {
	input := `let var = "expected";
    var;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, `"expected"`)
}

func TestIntAssign(t *testing.T) {
	input := `let var = 12;
    var;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "12")
}

func TestNegativeIntAssign(t *testing.T) {
	input := `let var = -12;
    var;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "-12")
}

func TestFloatAssign(t *testing.T) {
	input := `let var = 121.243;
    var;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "121.243")
}

func TestBoolAssign(t *testing.T) {
	input := `let var = true;
    var;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "true")
}

func TestNullAssign(t *testing.T) {
	input := `let var = null;
    var;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "null")
}

func TestNot(t *testing.T) {
	input := `let hi = !true;
    hi;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "false")
}

func TestInvalidNot(t *testing.T) {
	input := `let hi = !22;
    hi;`

	test.AssertPanic(t, func() { RunIlang(input) })
}

func TestUndefinedValPanic(t *testing.T) {
	input := "var;"

	test.AssertPanic(t, func() { RunIlang(input) })
}

func TestReassigningUndefinedValPanic(t *testing.T) {
	input := "var = 1;"

	test.AssertPanic(t, func() { RunIlang(input) })
}

func TestAssigningAlreadyDefinedValPanic(t *testing.T) {
	input := `let hi = 22;
    let hi = 1;
    hi;`

	test.AssertPanic(t, func() { RunIlang(input) })
}

func TestReassigning(t *testing.T) {
	input := `let hi = 22;
    hi = 1;
    hi;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "1")
}

func TestReassigningNewType(t *testing.T) {
	input := `let hi = 22;
    hi = "hello";
    hi;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, `"hello"`)
}

func TestReturn(t *testing.T) {
	input := `return "hello";`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, `"hello"`)
}

func TestEarlyReturn(t *testing.T) {
	input := `return "hello";
    "hi";`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, `"hello"`)
}

func TestReturnVar(t *testing.T) {
	input := `let val = 3;
    return val;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "3")
}

func TestArrayLiteral(t *testing.T) {
	input := `let val = [1, 2, 3, 4];
    return val;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "[1,2,3,4]")
}

func TestArrayDifferentTypes(t *testing.T) {
	input := `let val = [1, "hi", 3];
    return val;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, `[1,"hi",3]`)
}

func TestMap(t *testing.T) {
	input := `let val = {
        hi: "hello"
    };
    return val;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, `{"hi":"hello"}`)
}

func TestMapComputedProperty(t *testing.T) {
	input := `let val = {
        [1]: 2
    };
    return val;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, `{1:2}`)
}

func TestMapInvalidKey(t *testing.T) {
	input := `let val = {
        [1.2]: "hi"
    };
    return val;`

	test.AssertPanic(t, func() { RunIlang(input) })
}
