#!/bin/bash

# antlr generated code
rm -rf generated
cd grammar
antlr4 -Dlanguage=Go -no-listener -visitor ./ilang_lexer.g4 -o ../generated
antlr4 -Dlanguage=Go -no-listener -visitor ./ilang_parser.g4 -o ../generated
cd ..

# build go project
go build -o ./build
