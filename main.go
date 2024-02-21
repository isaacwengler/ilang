package main

import (
	"ilang/interpreter"
	"ilang/logger"
)

func main() {
	logger.Debug("Starting ilang")

	input := `
    let getAdder = func(num) {
        return func(n) {
            return num + n;
        };
    };

    let adder = getAdder(3);
    let result = 0;

    if (true) {
        result = adder(20);
    }
    result;
    `
	logger.Debug("input displayed below\n", input)
	res := interpreter.RunIlang(input)
	logger.Debug("Result =", res.PrintValue())
}
