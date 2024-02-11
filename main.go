package main

import (
	"ilang/interpreter"
	"ilang/logger"
)

func main() {
	logger.Debug("Starting ilang")

	input :=
		`let var = 12;
    return var;
    2389;`
	logger.Debug("input displayed below\n", input)
	res := interpreter.RunIlang(input)
	logger.Debug("Result =", res.PrintValue())
}
