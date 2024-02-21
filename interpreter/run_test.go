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

func TestEqualsComparisonSameInt(t *testing.T) {
	input := `1 == 1;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "true")
}

func TestEqualsComparisonDifferentInt(t *testing.T) {
	input := `1 == 2;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "false")
}

func TestEqualsComparisonString(t *testing.T) {
	input := `"hello" == "hello";`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "true")
}

func TestEqualsComparisonStringDifferent(t *testing.T) {
	input := `"hello" == "hello2";`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "false")
}

func TestNotEqualsComparisonSameInt(t *testing.T) {
	input := `1 != 1;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "false")
}

func TestEqualsArrayPtr(t *testing.T) {
	input := `let arr = [1, 3];
    arr == arr;`

	test.AssertPanic(t, func() { RunIlang(input) })
}

func TestInvalidOp(t *testing.T) {
	input := `let arr = [1, 3];
    3 <= arr;`

	test.AssertPanic(t, func() { RunIlang(input) })
}

func TestInvalidOpReversed(t *testing.T) {
	input := `let arr = [1, 3];
    arr <= 4;`

	test.AssertPanic(t, func() { RunIlang(input) })
}

func TestNotEqualsComparisonDifferentInt(t *testing.T) {
	input := `1 != 2;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "true")
}

func TestNullEquals(t *testing.T) {
	input := `let hi = null; hi == null;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "true")
}

func TestComplexComparison(t *testing.T) {
	input := `1 < 2 == true;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "true")
}

func TestComplexComparisonWithParens(t *testing.T) {
	input := `true == (1 < 2);`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "true")
}

func TestBooleanAnd(t *testing.T) {
	input := `true && true;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "true")
}

func TestBooleanAndFalse(t *testing.T) {
	input := `true && false;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "false")
}

func TestBooleanOr(t *testing.T) {
	input := `false || true;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "true")
}

func TestBooleanOpInvalid(t *testing.T) {
	// this is not javascript
	input := `true && 1;`

	test.AssertPanic(t, func() { RunIlang(input) })
}

func TestAdd(t *testing.T) {
	input := `3 + 28;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "31")
}

func TestAddFloat(t *testing.T) {
	input := `3.3 + 28;`

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "31.3")
}

func TestAddInvalid(t *testing.T) {
	input := `true + 28;`

	test.AssertPanic(t, func() { RunIlang(input) })
}

func TestIfPanic(t *testing.T) {
	input := `if ("hi") {
        let hi = 5;
    }`

	test.AssertPanic(t, func() { RunIlang(input) })
}

func TestIf(t *testing.T) {
	input := `
    let num = 6;
    let next = null;

    if (num < 0) {
        next = -1;
    } else if (num == 0) {
        next = 0;
    } else {
        next = 1;
    }

    next;
    `

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "1")
}

func TestWhilePanic(t *testing.T) {
	input := `while (1) {
        let hi = 5;
    }`

	test.AssertPanic(t, func() { RunIlang(input) })
}

func TestWhile(t *testing.T) {
	input := `
    let num = 0;

    while (num < 5) {
        num = num + 1;
    }

    num;
    `

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "5")
}

func TestFor(t *testing.T) {
	input := `
    let total = 0;

    for (let i = 0; i < 5; i = i + 1) {
        total = total + i;
    }

    total;
    `

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "10")
}

func TestForEach(t *testing.T) {
	input := `
    let total = 0;
    let items = [1, 2, 3, 4];

    for (item in items) {
        total = total + item;
    }

    total;
    `

	res := RunIlang(input).PrintValue()
	test.AssertStringsEqual(t, res, "10")
}

func TestForEachInvalid(t *testing.T) {
	input := `
    let total = 0;
    let items = {a: 1};

    for (item in items) {
        total = total + item;
    }

    total;
    `

	test.AssertPanic(t, func() { RunIlang(input) })
}

