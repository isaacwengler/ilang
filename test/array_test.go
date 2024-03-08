package test

import (
	"ilang/interpreter"
	"testing"
)

func TestArrayLiteral(t *testing.T) {
	input := `let val = [1, 2, 3, 4];
    return val;`

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, "[1,2,3,4]")
}

func TestArrayDifferentTypes(t *testing.T) {
	input := `let val = [1, "hi", 3];
    return val;`

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, `[1,"hi",3]`)
}

func TestEqualsArrayPtr(t *testing.T) {
	input := `let arr = [1, 3];
    arr == arr;`

	assertPanic(t, func() { interpreter.RunIlang(input) })
}

func TestArrayIndex(t *testing.T) {
	input := `let val = [1, 2, 3, 4];
    val[2];`

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, "3")
}

func TestArrayLength(t *testing.T) {
	input := `let val = [1, 2, 3, 4];
    val.length();`

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, "4")
}

func TestArrayLengthComputed(t *testing.T) {
	input := `[1, 2, 3, 4]["length"]();`

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, "4")
}

func TestArrayReassignment(t *testing.T) {
	input := `let arr = [1, 2, 3, 4];
    arr[2] = 15;
    arr;`

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, "[1,2,15,4]")
}

func TestPush(t *testing.T) {
	input := `let arr = [1, 2, 3, 4];
    arr.push(5);
    arr;`

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, "[1,2,3,4,5]")
}

func TestPop(t *testing.T) {
	input := `let arr = [1, 2, 3, 4];
    arr.pop();
    arr;`

	res := interpreter.RunIlang(input).PrintValue()
	assertStringEquals(t, res, "[1,2,3]")
}
