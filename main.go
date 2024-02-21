package main

import (
	"ilang/interpreter"
	"ilang/logger"
)

func main() {
	logger.Debug("Starting ilang")

	input := `
    let num = 0;
    let done = null;

    if (num < 0) {
       done = -1; 
    } else if (num > 0) {
        done = 1;
    }

    done;
    `
	logger.Debug("input displayed below\n", input)
	res := interpreter.RunIlang(input)
	logger.Debug("Result =", res.PrintValue())
}
