package main

import (
	"ilang/interpreter"
	"ilang/logger"
)

func main() {
	logger.Debug("Starting ilang")

	input := `
        let hi = "hi ";
        
        let sayHi = func(name) {
            println(hi + name);
        };
        
        sayHi("ilang");
    `

	logger.Debug("input displayed below\n", input)
	res := interpreter.RunIlang(input)
	logger.Debug("Result =", res.PrintValue())
}
