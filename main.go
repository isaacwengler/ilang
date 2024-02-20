package main

import (
	"ilang/interpreter"
	"ilang/logger"
)

func main() {
	logger.Debug("Starting ilang")

	input := `
    let val = 3;
    if (true) {
        let val = 2;
    }
    val;
    `
	logger.Debug("input displayed below\n", input)
	res := interpreter.RunIlang(input)
	logger.Debug("Result =", res.PrintValue())
}
