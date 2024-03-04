package interpreter

import (
	parser "ilang/generated"
	"ilang/logger"
	"ilang/library"
	"ilang/visitor"

	"github.com/antlr4-go/antlr/v4"
)

func RunIlang(input string) library.WrappedValue {
	inputStream := antlr.NewInputStream(input)

	lexer := parser.Newilang_lexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.Newilang_parser(tokenStream)

	logger.Debug("Initialized lexer and parser")
	logger.Debug("Parsing and visiting parse tree")
	v := visitor.NewVisitor()

	return v.VisitWrapped(p.Start_())
}
