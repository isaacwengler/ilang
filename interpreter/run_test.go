package interpreter

import (
	"testing"
)

func TestStringAssign(t *testing.T) {
	input := `let var = "expected";
    var;`

	res := RunIlang(input).PrintValue()
	if res != "expected" {
		t.Errorf("unable to assign string literal")
	}
}

func TestIntAssign(t *testing.T) {
	input := `let var = 12;
    var;`

	res := RunIlang(input).PrintValue()
	if res != int64(12) {
		t.Errorf("unable to assign int literal")
	}
}

func TestNegativeIntAssign(t *testing.T) {
	input := `let var = -12;
    var;`

	res := RunIlang(input).PrintValue()
	if res != int64(-12) {
		t.Errorf("unable to assign int literal")
	}
}

func TestFloatAssign(t *testing.T) {
	input := `let var = 121.243;
    var;`

	res := RunIlang(input).PrintValue()
	if res != 121.243 {
		t.Errorf("unable to assign float literal")
	}
}

func TestBoolAssign(t *testing.T) {
	input := `let var = true;
    var;`

	res := RunIlang(input).PrintValue()
	if res != true {
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
	if res != false {
		t.Errorf("unable to use not operator")
	}
}

func TestInvalidNot(t *testing.T) {
	input := `let hi = !22;
    hi;`

	assertPanic(t, func() { RunIlang(input) })
}

func TestFakeValPanic(t *testing.T) {
	input := "var;"

	assertPanic(t, func() { RunIlang(input) })
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}
