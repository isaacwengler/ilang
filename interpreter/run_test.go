package interpreter

import (
	"testing"
)

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func TestStringAssign(t *testing.T) {
	input := `let var = "expected";
    var;`

	res := RunIlang(input).PrintValue()
	if res != `"expected"` {
		t.Errorf("unable to assign string literal")
	}
}

func TestIntAssign(t *testing.T) {
	input := `let var = 12;
    var;`

	res := RunIlang(input).PrintValue()
	if res != "12" {
		t.Errorf("unable to assign int literal")
	}
}

func TestNegativeIntAssign(t *testing.T) {
	input := `let var = -12;
    var;`

	res := RunIlang(input).PrintValue()
	if res != "-12" {
		t.Errorf("unable to assign int literal")
	}
}

func TestFloatAssign(t *testing.T) {
	input := `let var = 121.243;
    var;`

	res := RunIlang(input).PrintValue()
	if res != "121.243" {
		t.Errorf("unable to assign float literal")
	}
}

func TestBoolAssign(t *testing.T) {
	input := `let var = true;
    var;`

	res := RunIlang(input).PrintValue()
	if res != "true" {
		t.Errorf("unable to assign bool literal")
	}
}

func TestNullAssign(t *testing.T) {
	input := `let var = null;
    var;`

	res := RunIlang(input).PrintValue()
	if res != "null" {
		t.Errorf("unable to assign null literal")
	}
}

func TestNot(t *testing.T) {
	input := `let hi = !true;
    hi;`

	res := RunIlang(input).PrintValue()
	if res != "false" {
		t.Errorf("unable to use not operator")
	}
}

func TestInvalidNot(t *testing.T) {
	input := `let hi = !22;
    hi;`

	assertPanic(t, func() { RunIlang(input) })
}

func TestUndefinedValPanic(t *testing.T) {
	input := "var;"

	assertPanic(t, func() { RunIlang(input) })
}

func TestReassigningUndefinedValPanic(t *testing.T) {
	input := "var = 1;"

	assertPanic(t, func() { RunIlang(input) })
}

func TestAssigningAlreadyDefinedValPanic(t *testing.T) {
	input := `let hi = 22;
    let hi = 1;
    hi;`

	assertPanic(t, func() { RunIlang(input) })
}

func TestReassigning(t *testing.T) {
	input := `let hi = 22;
    hi = 1;
    hi;`

	res := RunIlang(input).PrintValue()
	if res != "1" {
		t.Errorf("unable to reassign variable")
	}
}

func TestReassigningNewType(t *testing.T) {
	input := `let hi = 22;
    hi = "hello";
    hi;`

	res := RunIlang(input).PrintValue()
	if res != `"hello"` {
		t.Errorf("unable to reassign variable to a new type")
	}
}

func TestReturn(t *testing.T) {
	input := `return "hello";`

	res := RunIlang(input).PrintValue()
	if res != `"hello"` {
		t.Errorf("unable to return a value")
	}
}

func TestEarlyReturn(t *testing.T) {
	input := `return "hello";
    "hi";`

	res := RunIlang(input).PrintValue()
	if res != `"hello"` {
		t.Errorf("returned value had code run after it")
	}
}

func TestReturnVar(t *testing.T) {
	input := `let val = 3;
    return val;`

	res := RunIlang(input).PrintValue()
	if res != "3" {
		t.Errorf("returned value had code run after it")
	}
}

func TestArrayLiteral(t *testing.T) {
	input := `let val = [1, 2, 3, 4];
    return val;`

	res := RunIlang(input).PrintValue()
	if res != "[1,2,3,4]" {
		t.Errorf("unable to assign array literal")
	}
}

func TestArrayDifferentTypes(t *testing.T) {
	input := `let val = [1, "hi", 3];
    return val;`

	res := RunIlang(input).PrintValue()
	if res != `[1,"hi",3]` {
		t.Errorf("unable to assign array literal")
	}
}

func TestMap(t *testing.T) {
	input := `let val = {
        [1]: 2,
        hi: "hello"
    };
    return val;`

	res := RunIlang(input).PrintValue()
    if res != `{1:2,"hi":"hello"}` {
		t.Errorf("unable to assign array literal")
	}
}

func TestMapInvalidKey(t *testing.T) {
	input := `let val = {
        [1.2]: "hi"
    };
    return val;`

	assertPanic(t, func() { RunIlang(input) })
}
