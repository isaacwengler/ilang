package main

import (
	parser "ilang/generated"
	"ilang/logger"
	"ilang/visitor"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	logger.Debug("Starting ilang")

	input := "let hi = \"hello\";hi;"
	logger.Debug("input displayed below\n", input)
	inputStream := antlr.NewInputStream(input)

	lexer := parser.Newilang_lexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.Newilang_parser(tokenStream)

	logger.Debug("Initialized lexer and parser")
	logger.Debug("Parsing and visiting parse tree")
	v := visitor.NewVisitor()

    res := v.Visit(p.Start_())
    logger.Debug("Result =", res)
}
