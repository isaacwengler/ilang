package main

import (
	"ilang/interpreter"
	"ilang/logger"
)

func main() {
	logger.Debug("Starting ilang")

	input := "let hi = -1;hi;"
	logger.Debug("input displayed below\n", input)
	res := interpreter.RunIlang(input)
	logger.Debug("Result =", res)
}
