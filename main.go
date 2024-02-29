package main

import (
	"bufio"
	"flag"
	"fmt"
	"ilang/interpreter"
	"ilang/logger"
	"os"
)

func main() {
	debugMode := flag.Bool("debug", false, "enable verbose debug logging")
	file := flag.String("file", "", "Optional ilang file to run. If not provided, will enter interactive mode")
	flag.Parse()

	logger.SetDebugMode(*debugMode)
	logger.Debug("Starting ilang")

	if *file == "" {
		runInteractive()
	} else {
		runFromFile(*file)
	}
}

func runFromFile(file string) {
	logger.Debug("Opening file", file)
	contentsAsBytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	contents := string(contentsAsBytes)
	logger.Debug("File contents:\n", contents)

	res := interpreter.RunIlang(contents)
	logger.Debug("Result =", res.PrintValue())
}

func runInteractive() {
	logger.Debug("Starting interactive mode")
	buffer := ""

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		buffer += input

		res := interpreter.RunIlang(buffer)
		if res != nil {
			fmt.Println(res.PrintValue())
		}
	}
}
