package main

import (
	parser "ilang/generated"
	"ilang/logger"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
    logger.Debug("Starting ilang")

    inputStream := antlr.NewInputStream("let hi = \"hello\";")
    lexer := parser.Newilang_lexer(inputStream)
    tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
    
    p := parser.Newilang_parser(tokenStream)

    antlr.ParseTreeWalkerDefault.Walk(&antlr.BaseParseTreeListener{}, p.Start_())
}
