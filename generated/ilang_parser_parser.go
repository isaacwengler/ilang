// Code generated from ./ilang_parser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ilang_parser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ilang_parser struct {
	*antlr.BaseParser
}

var Ilang_parserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ilang_parserParserInit() {
	staticData := &Ilang_parserParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'{'", "'}'", "'['", "']'", "'='", "", "", "", "'!'",
		"'if'", "'else'", "'while'", "'for'", "'in'", "'return'", "'let'", "",
		"'true'", "'false'", "'null'", "'?'", "';'", "':'", "','", "'func'",
	}
	staticData.SymbolicNames = []string{
		"", "OPEN_PAREN", "CLOSE_PAREN", "OPEN_BRACE", "CLOSE_BRACE", "OPEN_BRACKET",
		"CLOSE_BRACKET", "EQUALS", "ARITHMETIC_OP", "CONDITIONAL_OP", "BOOLEAN_OP",
		"NOT", "IF", "ELSE", "WHILE", "FOR", "IN", "RETURN", "LET", "DOT", "TRUE",
		"FALSE", "NULL", "QUESTION", "SEMICOLON", "COLON", "COMMA", "FUNC",
		"INT", "FLOAT", "STRING", "COMMENT", "SYMBOL", "WHITE_SPACE",
	}
	staticData.RuleNames = []string{
		"start", "block", "statement", "expr", "functionDef", "functionDefArgs",
		"functionArgs", "assignment", "reassignment", "ifStatement", "whileLoop",
		"foreachLoop", "forLoop", "return", "elseifStatement", "elseStatement",
		"scopeBody", "conditionBody", "not", "symbol", "symbolChild", "stringLiteral",
		"intLiteral", "floatLiteral", "nullLiteral", "booleanLiteral", "arrayLiteral",
		"mapLiteral", "mapLiteralItem", "mapKey", "grouping",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 33, 291, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 71, 8, 1, 11, 1, 12, 1, 72,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 85, 8,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 3, 3, 99, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 120,
		8, 3, 10, 3, 12, 3, 123, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		5, 5, 132, 8, 5, 10, 5, 12, 5, 135, 9, 5, 1, 5, 3, 5, 138, 8, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 146, 8, 6, 10, 6, 12, 6, 149, 9, 6,
		1, 6, 3, 6, 152, 8, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 3, 8, 163, 8, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9,
		172, 8, 9, 10, 9, 12, 9, 175, 9, 9, 1, 9, 3, 9, 178, 8, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 233,
		8, 20, 1, 20, 3, 20, 236, 8, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1,
		23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 252,
		8, 26, 10, 26, 12, 26, 255, 9, 26, 1, 26, 3, 26, 258, 8, 26, 1, 26, 1,
		26, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 266, 8, 27, 10, 27, 12, 27, 269,
		9, 27, 1, 27, 3, 27, 272, 8, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 285, 8, 29, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 0, 1, 6, 31, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
		58, 60, 0, 1, 1, 0, 20, 21, 297, 0, 62, 1, 0, 0, 0, 2, 70, 1, 0, 0, 0,
		4, 84, 1, 0, 0, 0, 6, 98, 1, 0, 0, 0, 8, 124, 1, 0, 0, 0, 10, 128, 1, 0,
		0, 0, 12, 141, 1, 0, 0, 0, 14, 155, 1, 0, 0, 0, 16, 160, 1, 0, 0, 0, 18,
		167, 1, 0, 0, 0, 20, 179, 1, 0, 0, 0, 22, 183, 1, 0, 0, 0, 24, 191, 1,
		0, 0, 0, 26, 201, 1, 0, 0, 0, 28, 205, 1, 0, 0, 0, 30, 210, 1, 0, 0, 0,
		32, 213, 1, 0, 0, 0, 34, 217, 1, 0, 0, 0, 36, 221, 1, 0, 0, 0, 38, 224,
		1, 0, 0, 0, 40, 232, 1, 0, 0, 0, 42, 237, 1, 0, 0, 0, 44, 239, 1, 0, 0,
		0, 46, 241, 1, 0, 0, 0, 48, 243, 1, 0, 0, 0, 50, 245, 1, 0, 0, 0, 52, 247,
		1, 0, 0, 0, 54, 261, 1, 0, 0, 0, 56, 275, 1, 0, 0, 0, 58, 284, 1, 0, 0,
		0, 60, 286, 1, 0, 0, 0, 62, 63, 3, 2, 1, 0, 63, 64, 5, 0, 0, 1, 64, 1,
		1, 0, 0, 0, 65, 71, 3, 26, 13, 0, 66, 71, 3, 4, 2, 0, 67, 68, 3, 6, 3,
		0, 68, 69, 5, 24, 0, 0, 69, 71, 1, 0, 0, 0, 70, 65, 1, 0, 0, 0, 70, 66,
		1, 0, 0, 0, 70, 67, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0,
		72, 73, 1, 0, 0, 0, 73, 3, 1, 0, 0, 0, 74, 75, 3, 14, 7, 0, 75, 76, 5,
		24, 0, 0, 76, 85, 1, 0, 0, 0, 77, 78, 3, 16, 8, 0, 78, 79, 5, 24, 0, 0,
		79, 85, 1, 0, 0, 0, 80, 85, 3, 18, 9, 0, 81, 85, 3, 20, 10, 0, 82, 85,
		3, 24, 12, 0, 83, 85, 3, 22, 11, 0, 84, 74, 1, 0, 0, 0, 84, 77, 1, 0, 0,
		0, 84, 80, 1, 0, 0, 0, 84, 81, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 83,
		1, 0, 0, 0, 85, 5, 1, 0, 0, 0, 86, 87, 6, 3, -1, 0, 87, 99, 3, 36, 18,
		0, 88, 99, 3, 38, 19, 0, 89, 99, 3, 42, 21, 0, 90, 99, 3, 44, 22, 0, 91,
		99, 3, 46, 23, 0, 92, 99, 3, 48, 24, 0, 93, 99, 3, 50, 25, 0, 94, 99, 3,
		52, 26, 0, 95, 99, 3, 54, 27, 0, 96, 99, 3, 60, 30, 0, 97, 99, 3, 8, 4,
		0, 98, 86, 1, 0, 0, 0, 98, 88, 1, 0, 0, 0, 98, 89, 1, 0, 0, 0, 98, 90,
		1, 0, 0, 0, 98, 91, 1, 0, 0, 0, 98, 92, 1, 0, 0, 0, 98, 93, 1, 0, 0, 0,
		98, 94, 1, 0, 0, 0, 98, 95, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 97, 1,
		0, 0, 0, 99, 121, 1, 0, 0, 0, 100, 101, 10, 6, 0, 0, 101, 102, 5, 9, 0,
		0, 102, 120, 3, 6, 3, 7, 103, 104, 10, 5, 0, 0, 104, 105, 5, 8, 0, 0, 105,
		120, 3, 6, 3, 6, 106, 107, 10, 4, 0, 0, 107, 108, 5, 10, 0, 0, 108, 120,
		3, 6, 3, 5, 109, 110, 10, 3, 0, 0, 110, 120, 3, 12, 6, 0, 111, 112, 10,
		2, 0, 0, 112, 113, 5, 19, 0, 0, 113, 120, 5, 32, 0, 0, 114, 115, 10, 1,
		0, 0, 115, 116, 5, 5, 0, 0, 116, 117, 3, 6, 3, 0, 117, 118, 5, 6, 0, 0,
		118, 120, 1, 0, 0, 0, 119, 100, 1, 0, 0, 0, 119, 103, 1, 0, 0, 0, 119,
		106, 1, 0, 0, 0, 119, 109, 1, 0, 0, 0, 119, 111, 1, 0, 0, 0, 119, 114,
		1, 0, 0, 0, 120, 123, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 121, 122, 1, 0,
		0, 0, 122, 7, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 124, 125, 5, 27, 0, 0,
		125, 126, 3, 10, 5, 0, 126, 127, 3, 32, 16, 0, 127, 9, 1, 0, 0, 0, 128,
		137, 5, 1, 0, 0, 129, 130, 5, 32, 0, 0, 130, 132, 5, 26, 0, 0, 131, 129,
		1, 0, 0, 0, 132, 135, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0,
		0, 0, 134, 136, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 136, 138, 5, 32, 0, 0,
		137, 133, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139,
		140, 5, 2, 0, 0, 140, 11, 1, 0, 0, 0, 141, 151, 5, 1, 0, 0, 142, 143, 3,
		6, 3, 0, 143, 144, 5, 26, 0, 0, 144, 146, 1, 0, 0, 0, 145, 142, 1, 0, 0,
		0, 146, 149, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148,
		150, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 150, 152, 3, 6, 3, 0, 151, 147,
		1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 154, 5, 2,
		0, 0, 154, 13, 1, 0, 0, 0, 155, 156, 5, 18, 0, 0, 156, 157, 5, 32, 0, 0,
		157, 158, 5, 7, 0, 0, 158, 159, 3, 6, 3, 0, 159, 15, 1, 0, 0, 0, 160, 162,
		5, 32, 0, 0, 161, 163, 3, 40, 20, 0, 162, 161, 1, 0, 0, 0, 162, 163, 1,
		0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 165, 5, 7, 0, 0, 165, 166, 3, 6, 3,
		0, 166, 17, 1, 0, 0, 0, 167, 168, 5, 12, 0, 0, 168, 169, 3, 34, 17, 0,
		169, 173, 3, 32, 16, 0, 170, 172, 3, 28, 14, 0, 171, 170, 1, 0, 0, 0, 172,
		175, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 177,
		1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 176, 178, 3, 30, 15, 0, 177, 176, 1,
		0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 19, 1, 0, 0, 0, 179, 180, 5, 14, 0,
		0, 180, 181, 3, 34, 17, 0, 181, 182, 3, 32, 16, 0, 182, 21, 1, 0, 0, 0,
		183, 184, 5, 15, 0, 0, 184, 185, 5, 1, 0, 0, 185, 186, 5, 32, 0, 0, 186,
		187, 5, 16, 0, 0, 187, 188, 3, 6, 3, 0, 188, 189, 5, 2, 0, 0, 189, 190,
		3, 32, 16, 0, 190, 23, 1, 0, 0, 0, 191, 192, 5, 15, 0, 0, 192, 193, 5,
		1, 0, 0, 193, 194, 3, 14, 7, 0, 194, 195, 5, 24, 0, 0, 195, 196, 3, 6,
		3, 0, 196, 197, 5, 24, 0, 0, 197, 198, 3, 16, 8, 0, 198, 199, 5, 2, 0,
		0, 199, 200, 3, 32, 16, 0, 200, 25, 1, 0, 0, 0, 201, 202, 5, 17, 0, 0,
		202, 203, 3, 6, 3, 0, 203, 204, 5, 24, 0, 0, 204, 27, 1, 0, 0, 0, 205,
		206, 5, 13, 0, 0, 206, 207, 5, 12, 0, 0, 207, 208, 3, 34, 17, 0, 208, 209,
		3, 32, 16, 0, 209, 29, 1, 0, 0, 0, 210, 211, 5, 13, 0, 0, 211, 212, 3,
		32, 16, 0, 212, 31, 1, 0, 0, 0, 213, 214, 5, 3, 0, 0, 214, 215, 3, 2, 1,
		0, 215, 216, 5, 4, 0, 0, 216, 33, 1, 0, 0, 0, 217, 218, 5, 1, 0, 0, 218,
		219, 3, 6, 3, 0, 219, 220, 5, 2, 0, 0, 220, 35, 1, 0, 0, 0, 221, 222, 5,
		11, 0, 0, 222, 223, 3, 6, 3, 0, 223, 37, 1, 0, 0, 0, 224, 225, 5, 32, 0,
		0, 225, 39, 1, 0, 0, 0, 226, 227, 5, 19, 0, 0, 227, 233, 5, 32, 0, 0, 228,
		229, 5, 5, 0, 0, 229, 230, 3, 6, 3, 0, 230, 231, 5, 6, 0, 0, 231, 233,
		1, 0, 0, 0, 232, 226, 1, 0, 0, 0, 232, 228, 1, 0, 0, 0, 233, 235, 1, 0,
		0, 0, 234, 236, 3, 40, 20, 0, 235, 234, 1, 0, 0, 0, 235, 236, 1, 0, 0,
		0, 236, 41, 1, 0, 0, 0, 237, 238, 5, 30, 0, 0, 238, 43, 1, 0, 0, 0, 239,
		240, 5, 28, 0, 0, 240, 45, 1, 0, 0, 0, 241, 242, 5, 29, 0, 0, 242, 47,
		1, 0, 0, 0, 243, 244, 5, 22, 0, 0, 244, 49, 1, 0, 0, 0, 245, 246, 7, 0,
		0, 0, 246, 51, 1, 0, 0, 0, 247, 257, 5, 5, 0, 0, 248, 249, 3, 6, 3, 0,
		249, 250, 5, 26, 0, 0, 250, 252, 1, 0, 0, 0, 251, 248, 1, 0, 0, 0, 252,
		255, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 256,
		1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 256, 258, 3, 6, 3, 0, 257, 253, 1, 0,
		0, 0, 257, 258, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 260, 5, 6, 0, 0,
		260, 53, 1, 0, 0, 0, 261, 271, 5, 3, 0, 0, 262, 263, 3, 56, 28, 0, 263,
		264, 5, 26, 0, 0, 264, 266, 1, 0, 0, 0, 265, 262, 1, 0, 0, 0, 266, 269,
		1, 0, 0, 0, 267, 265, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 270, 1, 0,
		0, 0, 269, 267, 1, 0, 0, 0, 270, 272, 3, 56, 28, 0, 271, 267, 1, 0, 0,
		0, 271, 272, 1, 0, 0, 0, 272, 273, 1, 0, 0, 0, 273, 274, 5, 4, 0, 0, 274,
		55, 1, 0, 0, 0, 275, 276, 3, 58, 29, 0, 276, 277, 5, 25, 0, 0, 277, 278,
		3, 6, 3, 0, 278, 57, 1, 0, 0, 0, 279, 285, 5, 32, 0, 0, 280, 281, 5, 5,
		0, 0, 281, 282, 3, 6, 3, 0, 282, 283, 5, 6, 0, 0, 283, 285, 1, 0, 0, 0,
		284, 279, 1, 0, 0, 0, 284, 280, 1, 0, 0, 0, 285, 59, 1, 0, 0, 0, 286, 287,
		5, 1, 0, 0, 287, 288, 3, 6, 3, 0, 288, 289, 5, 2, 0, 0, 289, 61, 1, 0,
		0, 0, 20, 70, 72, 84, 98, 119, 121, 133, 137, 147, 151, 162, 173, 177,
		232, 235, 253, 257, 267, 271, 284,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ilang_parserInit initializes any static state used to implement ilang_parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// Newilang_parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Ilang_parserInit() {
	staticData := &Ilang_parserParserStaticData
	staticData.once.Do(ilang_parserParserInit)
}

// Newilang_parser produces a new parser instance for the optional input antlr.TokenStream.
func Newilang_parser(input antlr.TokenStream) *ilang_parser {
	Ilang_parserInit()
	this := new(ilang_parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &Ilang_parserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ilang_parser.g4"

	return this
}

// ilang_parser tokens.
const (
	ilang_parserEOF            = antlr.TokenEOF
	ilang_parserOPEN_PAREN     = 1
	ilang_parserCLOSE_PAREN    = 2
	ilang_parserOPEN_BRACE     = 3
	ilang_parserCLOSE_BRACE    = 4
	ilang_parserOPEN_BRACKET   = 5
	ilang_parserCLOSE_BRACKET  = 6
	ilang_parserEQUALS         = 7
	ilang_parserARITHMETIC_OP  = 8
	ilang_parserCONDITIONAL_OP = 9
	ilang_parserBOOLEAN_OP     = 10
	ilang_parserNOT            = 11
	ilang_parserIF             = 12
	ilang_parserELSE           = 13
	ilang_parserWHILE          = 14
	ilang_parserFOR            = 15
	ilang_parserIN             = 16
	ilang_parserRETURN         = 17
	ilang_parserLET            = 18
	ilang_parserDOT            = 19
	ilang_parserTRUE           = 20
	ilang_parserFALSE          = 21
	ilang_parserNULL           = 22
	ilang_parserQUESTION       = 23
	ilang_parserSEMICOLON      = 24
	ilang_parserCOLON          = 25
	ilang_parserCOMMA          = 26
	ilang_parserFUNC           = 27
	ilang_parserINT            = 28
	ilang_parserFLOAT          = 29
	ilang_parserSTRING         = 30
	ilang_parserCOMMENT        = 31
	ilang_parserSYMBOL         = 32
	ilang_parserWHITE_SPACE    = 33
)

// ilang_parser rules.
const (
	ilang_parserRULE_start           = 0
	ilang_parserRULE_block           = 1
	ilang_parserRULE_statement       = 2
	ilang_parserRULE_expr            = 3
	ilang_parserRULE_functionDef     = 4
	ilang_parserRULE_functionDefArgs = 5
	ilang_parserRULE_functionArgs    = 6
	ilang_parserRULE_assignment      = 7
	ilang_parserRULE_reassignment    = 8
	ilang_parserRULE_ifStatement     = 9
	ilang_parserRULE_whileLoop       = 10
	ilang_parserRULE_foreachLoop     = 11
	ilang_parserRULE_forLoop         = 12
	ilang_parserRULE_return          = 13
	ilang_parserRULE_elseifStatement = 14
	ilang_parserRULE_elseStatement   = 15
	ilang_parserRULE_scopeBody       = 16
	ilang_parserRULE_conditionBody   = 17
	ilang_parserRULE_not             = 18
	ilang_parserRULE_symbol          = 19
	ilang_parserRULE_symbolChild     = 20
	ilang_parserRULE_stringLiteral   = 21
	ilang_parserRULE_intLiteral      = 22
	ilang_parserRULE_floatLiteral    = 23
	ilang_parserRULE_nullLiteral     = 24
	ilang_parserRULE_booleanLiteral  = 25
	ilang_parserRULE_arrayLiteral    = 26
	ilang_parserRULE_mapLiteral      = 27
	ilang_parserRULE_mapLiteralItem  = 28
	ilang_parserRULE_mapKey          = 29
	ilang_parserRULE_grouping        = 30
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(ilang_parserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ilang_parserRULE_start)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Block()
	}
	{
		p.SetState(63)
		p.Match(ilang_parserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllReturn_() []IReturnContext
	Return_(i int) IReturnContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllReturn_() []IReturnContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IReturnContext); ok {
			len++
		}
	}

	tst := make([]IReturnContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IReturnContext); ok {
			tst[i] = t.(IReturnContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Return_(i int) IReturnContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnContext)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BlockContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(ilang_parserSEMICOLON)
}

func (s *BlockContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(ilang_parserSEMICOLON, i)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ilang_parserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6316021802) != 0) {
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(65)
				p.Return_()
			}

		case 2:
			{
				p.SetState(66)
				p.Statement()
			}

		case 3:
			{
				p.SetState(67)
				p.expr(0)
			}
			{
				p.SetState(68)
				p.Match(ilang_parserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	SEMICOLON() antlr.TerminalNode
	Reassignment() IReassignmentContext
	IfStatement() IIfStatementContext
	WhileLoop() IWhileLoopContext
	ForLoop() IForLoopContext
	ForeachLoop() IForeachLoopContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ilang_parserSEMICOLON, 0)
}

func (s *StatementContext) Reassignment() IReassignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReassignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReassignmentContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) WhileLoop() IWhileLoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileLoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileLoopContext)
}

func (s *StatementContext) ForLoop() IForLoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForLoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForLoopContext)
}

func (s *StatementContext) ForeachLoop() IForeachLoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForeachLoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForeachLoopContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ilang_parserRULE_statement)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(74)
			p.Assignment()
		}
		{
			p.SetState(75)
			p.Match(ilang_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(77)
			p.Reassignment()
		}
		{
			p.SetState(78)
			p.Match(ilang_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(80)
			p.IfStatement()
		}

	case 4:
		{
			p.SetState(81)
			p.WhileLoop()
		}

	case 5:
		{
			p.SetState(82)
			p.ForLoop()
		}

	case 6:
		{
			p.SetState(83)
			p.ForeachLoop()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IntExprContext struct {
	ExprContext
}

func NewIntExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntExprContext {
	var p = new(IntExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IntExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntExprContext) IntLiteral() IIntLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntLiteralContext)
}

func (s *IntExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitIntExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullExprContext struct {
	ExprContext
}

func NewNullExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullExprContext {
	var p = new(NullExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NullExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullExprContext) NullLiteral() INullLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INullLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INullLiteralContext)
}

func (s *NullExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitNullExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayExprContext struct {
	ExprContext
}

func NewArrayExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayExprContext {
	var p = new(ArrayExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArrayExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayExprContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *ArrayExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitArrayExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanExprContext struct {
	ExprContext
}

func NewBooleanExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanExprContext {
	var p = new(BooleanExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BooleanExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanExprContext) BooleanLiteral() IBooleanLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *BooleanExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitBooleanExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MapExprContext struct {
	ExprContext
}

func NewMapExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapExprContext {
	var p = new(MapExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MapExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapExprContext) MapLiteral() IMapLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapLiteralContext)
}

func (s *MapExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitMapExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArithmeticContext struct {
	ExprContext
}

func NewArithmeticContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArithmeticContext {
	var p = new(ArithmeticContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArithmeticContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArithmeticContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArithmeticContext) ARITHMETIC_OP() antlr.TerminalNode {
	return s.GetToken(ilang_parserARITHMETIC_OP, 0)
}

func (s *ArithmeticContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitArithmetic(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringExprContext struct {
	ExprContext
}

func NewStringExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringExprContext {
	var p = new(StringExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StringExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringExprContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *StringExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitStringExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatExprContext struct {
	ExprContext
}

func NewFloatExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatExprContext {
	var p = new(FloatExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FloatExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatExprContext) FloatLiteral() IFloatLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatLiteralContext)
}

func (s *FloatExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitFloatExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConditionContext struct {
	ExprContext
}

func NewConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConditionContext {
	var p = new(ConditionContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ConditionContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConditionContext) CONDITIONAL_OP() antlr.TerminalNode {
	return s.GetToken(ilang_parserCONDITIONAL_OP, 0)
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	ExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) Not() INotContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotContext)
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionDefExprContext struct {
	ExprContext
}

func NewFunctionDefExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionDefExprContext {
	var p = new(FunctionDefExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FunctionDefExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefExprContext) FunctionDef() IFunctionDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDefContext)
}

func (s *FunctionDefExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitFunctionDefExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallContext struct {
	ExprContext
}

func NewFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallContext {
	var p = new(FunctionCallContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunctionCallContext) FunctionArgs() IFunctionArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type PropertyContext struct {
	ExprContext
}

func NewPropertyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropertyContext {
	var p = new(PropertyContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PropertyContext) DOT() antlr.TerminalNode {
	return s.GetToken(ilang_parserDOT, 0)
}

func (s *PropertyContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(ilang_parserSYMBOL, 0)
}

func (s *PropertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitProperty(s)

	default:
		return t.VisitChildren(s)
	}
}

type SymbolExprContext struct {
	ExprContext
}

func NewSymbolExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SymbolExprContext {
	var p = new(SymbolExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *SymbolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolExprContext) Symbol() ISymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *SymbolExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitSymbolExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type GroupingExprContext struct {
	ExprContext
}

func NewGroupingExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupingExprContext {
	var p = new(GroupingExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *GroupingExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupingExprContext) Grouping() IGroupingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupingContext)
}

func (s *GroupingExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitGroupingExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanAlgebraContext struct {
	ExprContext
}

func NewBooleanAlgebraContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanAlgebraContext {
	var p = new(BooleanAlgebraContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BooleanAlgebraContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanAlgebraContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *BooleanAlgebraContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BooleanAlgebraContext) BOOLEAN_OP() antlr.TerminalNode {
	return s.GetToken(ilang_parserBOOLEAN_OP, 0)
}

func (s *BooleanAlgebraContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitBooleanAlgebra(s)

	default:
		return t.VisitChildren(s)
	}
}

type ComputedPropertyContext struct {
	ExprContext
}

func NewComputedPropertyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComputedPropertyContext {
	var p = new(ComputedPropertyContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ComputedPropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComputedPropertyContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ComputedPropertyContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ComputedPropertyContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(ilang_parserOPEN_BRACKET, 0)
}

func (s *ComputedPropertyContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(ilang_parserCLOSE_BRACKET, 0)
}

func (s *ComputedPropertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitComputedProperty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *ilang_parser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, ilang_parserRULE_expr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ilang_parserNOT:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(87)
			p.Not()
		}

	case ilang_parserSYMBOL:
		localctx = NewSymbolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(88)
			p.Symbol()
		}

	case ilang_parserSTRING:
		localctx = NewStringExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(89)
			p.StringLiteral()
		}

	case ilang_parserINT:
		localctx = NewIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(90)
			p.IntLiteral()
		}

	case ilang_parserFLOAT:
		localctx = NewFloatExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(91)
			p.FloatLiteral()
		}

	case ilang_parserNULL:
		localctx = NewNullExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(92)
			p.NullLiteral()
		}

	case ilang_parserTRUE, ilang_parserFALSE:
		localctx = NewBooleanExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(93)
			p.BooleanLiteral()
		}

	case ilang_parserOPEN_BRACKET:
		localctx = NewArrayExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(94)
			p.ArrayLiteral()
		}

	case ilang_parserOPEN_BRACE:
		localctx = NewMapExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(95)
			p.MapLiteral()
		}

	case ilang_parserOPEN_PAREN:
		localctx = NewGroupingExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(96)
			p.Grouping()
		}

	case ilang_parserFUNC:
		localctx = NewFunctionDefExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(97)
			p.FunctionDef()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewConditionContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ilang_parserRULE_expr)
				p.SetState(100)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(101)
					p.Match(ilang_parserCONDITIONAL_OP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(102)
					p.expr(7)
				}

			case 2:
				localctx = NewArithmeticContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ilang_parserRULE_expr)
				p.SetState(103)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(104)
					p.Match(ilang_parserARITHMETIC_OP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(105)
					p.expr(6)
				}

			case 3:
				localctx = NewBooleanAlgebraContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ilang_parserRULE_expr)
				p.SetState(106)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(107)
					p.Match(ilang_parserBOOLEAN_OP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(108)
					p.expr(5)
				}

			case 4:
				localctx = NewFunctionCallContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ilang_parserRULE_expr)
				p.SetState(109)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(110)
					p.FunctionArgs()
				}

			case 5:
				localctx = NewPropertyContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ilang_parserRULE_expr)
				p.SetState(111)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(112)
					p.Match(ilang_parserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(113)
					p.Match(ilang_parserSYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 6:
				localctx = NewComputedPropertyContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ilang_parserRULE_expr)
				p.SetState(114)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(115)
					p.Match(ilang_parserOPEN_BRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(116)
					p.expr(0)
				}
				{
					p.SetState(117)
					p.Match(ilang_parserCLOSE_BRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionDefContext is an interface to support dynamic dispatch.
type IFunctionDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	FunctionDefArgs() IFunctionDefArgsContext
	ScopeBody() IScopeBodyContext

	// IsFunctionDefContext differentiates from other interfaces.
	IsFunctionDefContext()
}

type FunctionDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefContext() *FunctionDefContext {
	var p = new(FunctionDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_functionDef
	return p
}

func InitEmptyFunctionDefContext(p *FunctionDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_functionDef
}

func (*FunctionDefContext) IsFunctionDefContext() {}

func NewFunctionDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefContext {
	var p = new(FunctionDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_functionDef

	return p
}

func (s *FunctionDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefContext) FUNC() antlr.TerminalNode {
	return s.GetToken(ilang_parserFUNC, 0)
}

func (s *FunctionDefContext) FunctionDefArgs() IFunctionDefArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDefArgsContext)
}

func (s *FunctionDefContext) ScopeBody() IScopeBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeBodyContext)
}

func (s *FunctionDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitFunctionDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) FunctionDef() (localctx IFunctionDefContext) {
	localctx = NewFunctionDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ilang_parserRULE_functionDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(ilang_parserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.FunctionDefArgs()
	}
	{
		p.SetState(126)
		p.ScopeBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionDefArgsContext is an interface to support dynamic dispatch.
type IFunctionDefArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_PAREN() antlr.TerminalNode
	CLOSE_PAREN() antlr.TerminalNode
	AllSYMBOL() []antlr.TerminalNode
	SYMBOL(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunctionDefArgsContext differentiates from other interfaces.
	IsFunctionDefArgsContext()
}

type FunctionDefArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefArgsContext() *FunctionDefArgsContext {
	var p = new(FunctionDefArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_functionDefArgs
	return p
}

func InitEmptyFunctionDefArgsContext(p *FunctionDefArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_functionDefArgs
}

func (*FunctionDefArgsContext) IsFunctionDefArgsContext() {}

func NewFunctionDefArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefArgsContext {
	var p = new(FunctionDefArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_functionDefArgs

	return p
}

func (s *FunctionDefArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefArgsContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(ilang_parserOPEN_PAREN, 0)
}

func (s *FunctionDefArgsContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(ilang_parserCLOSE_PAREN, 0)
}

func (s *FunctionDefArgsContext) AllSYMBOL() []antlr.TerminalNode {
	return s.GetTokens(ilang_parserSYMBOL)
}

func (s *FunctionDefArgsContext) SYMBOL(i int) antlr.TerminalNode {
	return s.GetToken(ilang_parserSYMBOL, i)
}

func (s *FunctionDefArgsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ilang_parserCOMMA)
}

func (s *FunctionDefArgsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ilang_parserCOMMA, i)
}

func (s *FunctionDefArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitFunctionDefArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) FunctionDefArgs() (localctx IFunctionDefArgsContext) {
	localctx = NewFunctionDefArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ilang_parserRULE_functionDefArgs)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(ilang_parserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ilang_parserSYMBOL {
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(129)
					p.Match(ilang_parserSYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(130)
					p.Match(ilang_parserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(135)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Match(ilang_parserSYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(139)
		p.Match(ilang_parserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionArgsContext is an interface to support dynamic dispatch.
type IFunctionArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetArgs returns the args rule context list.
	GetArgs() []IExprContext

	// SetArgs sets the args rule context list.
	SetArgs([]IExprContext)

	// Getter signatures
	OPEN_PAREN() antlr.TerminalNode
	CLOSE_PAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunctionArgsContext differentiates from other interfaces.
	IsFunctionArgsContext()
}

type FunctionArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	_expr  IExprContext
	args   []IExprContext
}

func NewEmptyFunctionArgsContext() *FunctionArgsContext {
	var p = new(FunctionArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_functionArgs
	return p
}

func InitEmptyFunctionArgsContext(p *FunctionArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_functionArgs
}

func (*FunctionArgsContext) IsFunctionArgsContext() {}

func NewFunctionArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgsContext {
	var p = new(FunctionArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_functionArgs

	return p
}

func (s *FunctionArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgsContext) Get_expr() IExprContext { return s._expr }

func (s *FunctionArgsContext) Set_expr(v IExprContext) { s._expr = v }

func (s *FunctionArgsContext) GetArgs() []IExprContext { return s.args }

func (s *FunctionArgsContext) SetArgs(v []IExprContext) { s.args = v }

func (s *FunctionArgsContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(ilang_parserOPEN_PAREN, 0)
}

func (s *FunctionArgsContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(ilang_parserCLOSE_PAREN, 0)
}

func (s *FunctionArgsContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *FunctionArgsContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunctionArgsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ilang_parserCOMMA)
}

func (s *FunctionArgsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ilang_parserCOMMA, i)
}

func (s *FunctionArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitFunctionArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) FunctionArgs() (localctx IFunctionArgsContext) {
	localctx = NewFunctionArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ilang_parserRULE_functionArgs)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(ilang_parserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6315575338) != 0 {
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(142)

					var _x = p.expr(0)

					localctx.(*FunctionArgsContext)._expr = _x
				}
				localctx.(*FunctionArgsContext).args = append(localctx.(*FunctionArgsContext).args, localctx.(*FunctionArgsContext)._expr)
				{
					p.SetState(143)
					p.Match(ilang_parserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(149)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(150)

			var _x = p.expr(0)

			localctx.(*FunctionArgsContext)._expr = _x
		}
		localctx.(*FunctionArgsContext).args = append(localctx.(*FunctionArgsContext).args, localctx.(*FunctionArgsContext)._expr)

	}
	{
		p.SetState(153)
		p.Match(ilang_parserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LET() antlr.TerminalNode
	SYMBOL() antlr.TerminalNode
	EQUALS() antlr.TerminalNode
	Expr() IExprContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) LET() antlr.TerminalNode {
	return s.GetToken(ilang_parserLET, 0)
}

func (s *AssignmentContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(ilang_parserSYMBOL, 0)
}

func (s *AssignmentContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(ilang_parserEQUALS, 0)
}

func (s *AssignmentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ilang_parserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(ilang_parserLET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Match(ilang_parserSYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.Match(ilang_parserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReassignmentContext is an interface to support dynamic dispatch.
type IReassignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYMBOL() antlr.TerminalNode
	EQUALS() antlr.TerminalNode
	Expr() IExprContext
	SymbolChild() ISymbolChildContext

	// IsReassignmentContext differentiates from other interfaces.
	IsReassignmentContext()
}

type ReassignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReassignmentContext() *ReassignmentContext {
	var p = new(ReassignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_reassignment
	return p
}

func InitEmptyReassignmentContext(p *ReassignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_reassignment
}

func (*ReassignmentContext) IsReassignmentContext() {}

func NewReassignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReassignmentContext {
	var p = new(ReassignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_reassignment

	return p
}

func (s *ReassignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *ReassignmentContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(ilang_parserSYMBOL, 0)
}

func (s *ReassignmentContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(ilang_parserEQUALS, 0)
}

func (s *ReassignmentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReassignmentContext) SymbolChild() ISymbolChildContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolChildContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolChildContext)
}

func (s *ReassignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReassignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReassignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitReassignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) Reassignment() (localctx IReassignmentContext) {
	localctx = NewReassignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ilang_parserRULE_reassignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(ilang_parserSYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ilang_parserOPEN_BRACKET || _la == ilang_parserDOT {
		{
			p.SetState(161)
			p.SymbolChild()
		}

	}
	{
		p.SetState(164)
		p.Match(ilang_parserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	ConditionBody() IConditionBodyContext
	ScopeBody() IScopeBodyContext
	AllElseifStatement() []IElseifStatementContext
	ElseifStatement(i int) IElseifStatementContext
	ElseStatement() IElseStatementContext

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(ilang_parserIF, 0)
}

func (s *IfStatementContext) ConditionBody() IConditionBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionBodyContext)
}

func (s *IfStatementContext) ScopeBody() IScopeBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeBodyContext)
}

func (s *IfStatementContext) AllElseifStatement() []IElseifStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseifStatementContext); ok {
			len++
		}
	}

	tst := make([]IElseifStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseifStatementContext); ok {
			tst[i] = t.(IElseifStatementContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) ElseifStatement(i int) IElseifStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseifStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseifStatementContext)
}

func (s *IfStatementContext) ElseStatement() IElseStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseStatementContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ilang_parserRULE_ifStatement)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(ilang_parserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.ConditionBody()
	}
	{
		p.SetState(169)
		p.ScopeBody()
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(170)
				p.ElseifStatement()
			}

		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ilang_parserELSE {
		{
			p.SetState(176)
			p.ElseStatement()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileLoopContext is an interface to support dynamic dispatch.
type IWhileLoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	ConditionBody() IConditionBodyContext
	ScopeBody() IScopeBodyContext

	// IsWhileLoopContext differentiates from other interfaces.
	IsWhileLoopContext()
}

type WhileLoopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileLoopContext() *WhileLoopContext {
	var p = new(WhileLoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_whileLoop
	return p
}

func InitEmptyWhileLoopContext(p *WhileLoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_whileLoop
}

func (*WhileLoopContext) IsWhileLoopContext() {}

func NewWhileLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileLoopContext {
	var p = new(WhileLoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_whileLoop

	return p
}

func (s *WhileLoopContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileLoopContext) WHILE() antlr.TerminalNode {
	return s.GetToken(ilang_parserWHILE, 0)
}

func (s *WhileLoopContext) ConditionBody() IConditionBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionBodyContext)
}

func (s *WhileLoopContext) ScopeBody() IScopeBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeBodyContext)
}

func (s *WhileLoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileLoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileLoopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitWhileLoop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) WhileLoop() (localctx IWhileLoopContext) {
	localctx = NewWhileLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ilang_parserRULE_whileLoop)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(ilang_parserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
		p.ConditionBody()
	}
	{
		p.SetState(181)
		p.ScopeBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForeachLoopContext is an interface to support dynamic dispatch.
type IForeachLoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	OPEN_PAREN() antlr.TerminalNode
	SYMBOL() antlr.TerminalNode
	IN() antlr.TerminalNode
	Expr() IExprContext
	CLOSE_PAREN() antlr.TerminalNode
	ScopeBody() IScopeBodyContext

	// IsForeachLoopContext differentiates from other interfaces.
	IsForeachLoopContext()
}

type ForeachLoopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForeachLoopContext() *ForeachLoopContext {
	var p = new(ForeachLoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_foreachLoop
	return p
}

func InitEmptyForeachLoopContext(p *ForeachLoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_foreachLoop
}

func (*ForeachLoopContext) IsForeachLoopContext() {}

func NewForeachLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForeachLoopContext {
	var p = new(ForeachLoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_foreachLoop

	return p
}

func (s *ForeachLoopContext) GetParser() antlr.Parser { return s.parser }

func (s *ForeachLoopContext) FOR() antlr.TerminalNode {
	return s.GetToken(ilang_parserFOR, 0)
}

func (s *ForeachLoopContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(ilang_parserOPEN_PAREN, 0)
}

func (s *ForeachLoopContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(ilang_parserSYMBOL, 0)
}

func (s *ForeachLoopContext) IN() antlr.TerminalNode {
	return s.GetToken(ilang_parserIN, 0)
}

func (s *ForeachLoopContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForeachLoopContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(ilang_parserCLOSE_PAREN, 0)
}

func (s *ForeachLoopContext) ScopeBody() IScopeBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeBodyContext)
}

func (s *ForeachLoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForeachLoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForeachLoopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitForeachLoop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) ForeachLoop() (localctx IForeachLoopContext) {
	localctx = NewForeachLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ilang_parserRULE_foreachLoop)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(ilang_parserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.Match(ilang_parserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
		p.Match(ilang_parserSYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.Match(ilang_parserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		p.expr(0)
	}
	{
		p.SetState(188)
		p.Match(ilang_parserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(189)
		p.ScopeBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForLoopContext is an interface to support dynamic dispatch.
type IForLoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInit returns the init rule contexts.
	GetInit() IAssignmentContext

	// GetCond returns the cond rule contexts.
	GetCond() IExprContext

	// GetStep returns the step rule contexts.
	GetStep() IReassignmentContext

	// SetInit sets the init rule contexts.
	SetInit(IAssignmentContext)

	// SetCond sets the cond rule contexts.
	SetCond(IExprContext)

	// SetStep sets the step rule contexts.
	SetStep(IReassignmentContext)

	// Getter signatures
	FOR() antlr.TerminalNode
	OPEN_PAREN() antlr.TerminalNode
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode
	CLOSE_PAREN() antlr.TerminalNode
	ScopeBody() IScopeBodyContext
	Assignment() IAssignmentContext
	Expr() IExprContext
	Reassignment() IReassignmentContext

	// IsForLoopContext differentiates from other interfaces.
	IsForLoopContext()
}

type ForLoopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	init   IAssignmentContext
	cond   IExprContext
	step   IReassignmentContext
}

func NewEmptyForLoopContext() *ForLoopContext {
	var p = new(ForLoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_forLoop
	return p
}

func InitEmptyForLoopContext(p *ForLoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_forLoop
}

func (*ForLoopContext) IsForLoopContext() {}

func NewForLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForLoopContext {
	var p = new(ForLoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_forLoop

	return p
}

func (s *ForLoopContext) GetParser() antlr.Parser { return s.parser }

func (s *ForLoopContext) GetInit() IAssignmentContext { return s.init }

func (s *ForLoopContext) GetCond() IExprContext { return s.cond }

func (s *ForLoopContext) GetStep() IReassignmentContext { return s.step }

func (s *ForLoopContext) SetInit(v IAssignmentContext) { s.init = v }

func (s *ForLoopContext) SetCond(v IExprContext) { s.cond = v }

func (s *ForLoopContext) SetStep(v IReassignmentContext) { s.step = v }

func (s *ForLoopContext) FOR() antlr.TerminalNode {
	return s.GetToken(ilang_parserFOR, 0)
}

func (s *ForLoopContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(ilang_parserOPEN_PAREN, 0)
}

func (s *ForLoopContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(ilang_parserSEMICOLON)
}

func (s *ForLoopContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(ilang_parserSEMICOLON, i)
}

func (s *ForLoopContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(ilang_parserCLOSE_PAREN, 0)
}

func (s *ForLoopContext) ScopeBody() IScopeBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeBodyContext)
}

func (s *ForLoopContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ForLoopContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForLoopContext) Reassignment() IReassignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReassignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReassignmentContext)
}

func (s *ForLoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForLoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForLoopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitForLoop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) ForLoop() (localctx IForLoopContext) {
	localctx = NewForLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ilang_parserRULE_forLoop)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.Match(ilang_parserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(192)
		p.Match(ilang_parserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)

		var _x = p.Assignment()

		localctx.(*ForLoopContext).init = _x
	}
	{
		p.SetState(194)
		p.Match(ilang_parserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(195)

		var _x = p.expr(0)

		localctx.(*ForLoopContext).cond = _x
	}
	{
		p.SetState(196)
		p.Match(ilang_parserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)

		var _x = p.Reassignment()

		localctx.(*ForLoopContext).step = _x
	}
	{
		p.SetState(198)
		p.Match(ilang_parserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.ScopeBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnContext is an interface to support dynamic dispatch.
type IReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expr() IExprContext
	SEMICOLON() antlr.TerminalNode

	// IsReturnContext differentiates from other interfaces.
	IsReturnContext()
}

type ReturnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnContext() *ReturnContext {
	var p = new(ReturnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_return
	return p
}

func InitEmptyReturnContext(p *ReturnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_return
}

func (*ReturnContext) IsReturnContext() {}

func NewReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnContext {
	var p = new(ReturnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_return

	return p
}

func (s *ReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(ilang_parserRETURN, 0)
}

func (s *ReturnContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ilang_parserSEMICOLON, 0)
}

func (s *ReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) Return_() (localctx IReturnContext) {
	localctx = NewReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ilang_parserRULE_return)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(ilang_parserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.expr(0)
	}
	{
		p.SetState(203)
		p.Match(ilang_parserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseifStatementContext is an interface to support dynamic dispatch.
type IElseifStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	IF() antlr.TerminalNode
	ConditionBody() IConditionBodyContext
	ScopeBody() IScopeBodyContext

	// IsElseifStatementContext differentiates from other interfaces.
	IsElseifStatementContext()
}

type ElseifStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseifStatementContext() *ElseifStatementContext {
	var p = new(ElseifStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_elseifStatement
	return p
}

func InitEmptyElseifStatementContext(p *ElseifStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_elseifStatement
}

func (*ElseifStatementContext) IsElseifStatementContext() {}

func NewElseifStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseifStatementContext {
	var p = new(ElseifStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_elseifStatement

	return p
}

func (s *ElseifStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseifStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(ilang_parserELSE, 0)
}

func (s *ElseifStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(ilang_parserIF, 0)
}

func (s *ElseifStatementContext) ConditionBody() IConditionBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionBodyContext)
}

func (s *ElseifStatementContext) ScopeBody() IScopeBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeBodyContext)
}

func (s *ElseifStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseifStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseifStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitElseifStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) ElseifStatement() (localctx IElseifStatementContext) {
	localctx = NewElseifStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ilang_parserRULE_elseifStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(ilang_parserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Match(ilang_parserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
		p.ConditionBody()
	}
	{
		p.SetState(208)
		p.ScopeBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseStatementContext is an interface to support dynamic dispatch.
type IElseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	ScopeBody() IScopeBodyContext

	// IsElseStatementContext differentiates from other interfaces.
	IsElseStatementContext()
}

type ElseStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStatementContext() *ElseStatementContext {
	var p = new(ElseStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_elseStatement
	return p
}

func InitEmptyElseStatementContext(p *ElseStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_elseStatement
}

func (*ElseStatementContext) IsElseStatementContext() {}

func NewElseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStatementContext {
	var p = new(ElseStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_elseStatement

	return p
}

func (s *ElseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(ilang_parserELSE, 0)
}

func (s *ElseStatementContext) ScopeBody() IScopeBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeBodyContext)
}

func (s *ElseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitElseStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) ElseStatement() (localctx IElseStatementContext) {
	localctx = NewElseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ilang_parserRULE_elseStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Match(ilang_parserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)
		p.ScopeBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IScopeBodyContext is an interface to support dynamic dispatch.
type IScopeBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_BRACE() antlr.TerminalNode
	Block() IBlockContext
	CLOSE_BRACE() antlr.TerminalNode

	// IsScopeBodyContext differentiates from other interfaces.
	IsScopeBodyContext()
}

type ScopeBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScopeBodyContext() *ScopeBodyContext {
	var p = new(ScopeBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_scopeBody
	return p
}

func InitEmptyScopeBodyContext(p *ScopeBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_scopeBody
}

func (*ScopeBodyContext) IsScopeBodyContext() {}

func NewScopeBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScopeBodyContext {
	var p = new(ScopeBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_scopeBody

	return p
}

func (s *ScopeBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ScopeBodyContext) OPEN_BRACE() antlr.TerminalNode {
	return s.GetToken(ilang_parserOPEN_BRACE, 0)
}

func (s *ScopeBodyContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ScopeBodyContext) CLOSE_BRACE() antlr.TerminalNode {
	return s.GetToken(ilang_parserCLOSE_BRACE, 0)
}

func (s *ScopeBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScopeBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScopeBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitScopeBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) ScopeBody() (localctx IScopeBodyContext) {
	localctx = NewScopeBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ilang_parserRULE_scopeBody)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Match(ilang_parserOPEN_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		p.Block()
	}
	{
		p.SetState(215)
		p.Match(ilang_parserCLOSE_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionBodyContext is an interface to support dynamic dispatch.
type IConditionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_PAREN() antlr.TerminalNode
	Expr() IExprContext
	CLOSE_PAREN() antlr.TerminalNode

	// IsConditionBodyContext differentiates from other interfaces.
	IsConditionBodyContext()
}

type ConditionBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionBodyContext() *ConditionBodyContext {
	var p = new(ConditionBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_conditionBody
	return p
}

func InitEmptyConditionBodyContext(p *ConditionBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_conditionBody
}

func (*ConditionBodyContext) IsConditionBodyContext() {}

func NewConditionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionBodyContext {
	var p = new(ConditionBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_conditionBody

	return p
}

func (s *ConditionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionBodyContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(ilang_parserOPEN_PAREN, 0)
}

func (s *ConditionBodyContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConditionBodyContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(ilang_parserCLOSE_PAREN, 0)
}

func (s *ConditionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitConditionBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) ConditionBody() (localctx IConditionBodyContext) {
	localctx = NewConditionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ilang_parserRULE_conditionBody)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		p.Match(ilang_parserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(218)
		p.expr(0)
	}
	{
		p.SetState(219)
		p.Match(ilang_parserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INotContext is an interface to support dynamic dispatch.
type INotContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NOT() antlr.TerminalNode
	Expr() IExprContext

	// IsNotContext differentiates from other interfaces.
	IsNotContext()
}

type NotContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotContext() *NotContext {
	var p = new(NotContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_not
	return p
}

func InitEmptyNotContext(p *NotContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_not
}

func (*NotContext) IsNotContext() {}

func NewNotContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotContext {
	var p = new(NotContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_not

	return p
}

func (s *NotContext) GetParser() antlr.Parser { return s.parser }

func (s *NotContext) NOT() antlr.TerminalNode {
	return s.GetToken(ilang_parserNOT, 0)
}

func (s *NotContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitNot(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) Not() (localctx INotContext) {
	localctx = NewNotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ilang_parserRULE_not)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Match(ilang_parserNOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISymbolContext is an interface to support dynamic dispatch.
type ISymbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYMBOL() antlr.TerminalNode

	// IsSymbolContext differentiates from other interfaces.
	IsSymbolContext()
}

type SymbolContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolContext() *SymbolContext {
	var p = new(SymbolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_symbol
	return p
}

func InitEmptySymbolContext(p *SymbolContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_symbol
}

func (*SymbolContext) IsSymbolContext() {}

func NewSymbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolContext {
	var p = new(SymbolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_symbol

	return p
}

func (s *SymbolContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(ilang_parserSYMBOL, 0)
}

func (s *SymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitSymbol(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) Symbol() (localctx ISymbolContext) {
	localctx = NewSymbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ilang_parserRULE_symbol)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Match(ilang_parserSYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISymbolChildContext is an interface to support dynamic dispatch.
type ISymbolChildContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOT() antlr.TerminalNode
	SYMBOL() antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	Expr() IExprContext
	CLOSE_BRACKET() antlr.TerminalNode
	SymbolChild() ISymbolChildContext

	// IsSymbolChildContext differentiates from other interfaces.
	IsSymbolChildContext()
}

type SymbolChildContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolChildContext() *SymbolChildContext {
	var p = new(SymbolChildContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_symbolChild
	return p
}

func InitEmptySymbolChildContext(p *SymbolChildContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_symbolChild
}

func (*SymbolChildContext) IsSymbolChildContext() {}

func NewSymbolChildContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolChildContext {
	var p = new(SymbolChildContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_symbolChild

	return p
}

func (s *SymbolChildContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolChildContext) DOT() antlr.TerminalNode {
	return s.GetToken(ilang_parserDOT, 0)
}

func (s *SymbolChildContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(ilang_parserSYMBOL, 0)
}

func (s *SymbolChildContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(ilang_parserOPEN_BRACKET, 0)
}

func (s *SymbolChildContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SymbolChildContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(ilang_parserCLOSE_BRACKET, 0)
}

func (s *SymbolChildContext) SymbolChild() ISymbolChildContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolChildContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolChildContext)
}

func (s *SymbolChildContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolChildContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolChildContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitSymbolChild(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) SymbolChild() (localctx ISymbolChildContext) {
	localctx = NewSymbolChildContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ilang_parserRULE_symbolChild)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ilang_parserDOT:
		{
			p.SetState(226)
			p.Match(ilang_parserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)
			p.Match(ilang_parserSYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ilang_parserOPEN_BRACKET:
		{
			p.SetState(228)
			p.Match(ilang_parserOPEN_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.expr(0)
		}
		{
			p.SetState(230)
			p.Match(ilang_parserCLOSE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ilang_parserOPEN_BRACKET || _la == ilang_parserDOT {
		{
			p.SetState(234)
			p.SymbolChild()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_stringLiteral
	return p
}

func InitEmptyStringLiteralContext(p *StringLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_stringLiteral
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(ilang_parserSTRING, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) StringLiteral() (localctx IStringLiteralContext) {
	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ilang_parserRULE_stringLiteral)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Match(ilang_parserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntLiteralContext is an interface to support dynamic dispatch.
type IIntLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode

	// IsIntLiteralContext differentiates from other interfaces.
	IsIntLiteralContext()
}

type IntLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntLiteralContext() *IntLiteralContext {
	var p = new(IntLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_intLiteral
	return p
}

func InitEmptyIntLiteralContext(p *IntLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_intLiteral
}

func (*IntLiteralContext) IsIntLiteralContext() {}

func NewIntLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntLiteralContext {
	var p = new(IntLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_intLiteral

	return p
}

func (s *IntLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntLiteralContext) INT() antlr.TerminalNode {
	return s.GetToken(ilang_parserINT, 0)
}

func (s *IntLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitIntLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) IntLiteral() (localctx IIntLiteralContext) {
	localctx = NewIntLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ilang_parserRULE_intLiteral)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(ilang_parserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFloatLiteralContext is an interface to support dynamic dispatch.
type IFloatLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FLOAT() antlr.TerminalNode

	// IsFloatLiteralContext differentiates from other interfaces.
	IsFloatLiteralContext()
}

type FloatLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatLiteralContext() *FloatLiteralContext {
	var p = new(FloatLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_floatLiteral
	return p
}

func InitEmptyFloatLiteralContext(p *FloatLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_floatLiteral
}

func (*FloatLiteralContext) IsFloatLiteralContext() {}

func NewFloatLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_floatLiteral

	return p
}

func (s *FloatLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLiteralContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ilang_parserFLOAT, 0)
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) FloatLiteral() (localctx IFloatLiteralContext) {
	localctx = NewFloatLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ilang_parserRULE_floatLiteral)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Match(ilang_parserFLOAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INullLiteralContext is an interface to support dynamic dispatch.
type INullLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NULL() antlr.TerminalNode

	// IsNullLiteralContext differentiates from other interfaces.
	IsNullLiteralContext()
}

type NullLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNullLiteralContext() *NullLiteralContext {
	var p = new(NullLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_nullLiteral
	return p
}

func InitEmptyNullLiteralContext(p *NullLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_nullLiteral
}

func (*NullLiteralContext) IsNullLiteralContext() {}

func NewNullLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NullLiteralContext {
	var p = new(NullLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_nullLiteral

	return p
}

func (s *NullLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NullLiteralContext) NULL() antlr.TerminalNode {
	return s.GetToken(ilang_parserNULL, 0)
}

func (s *NullLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NullLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitNullLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) NullLiteral() (localctx INullLiteralContext) {
	localctx = NewNullLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ilang_parserRULE_nullLiteral)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(ilang_parserNULL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_booleanLiteral
	return p
}

func InitEmptyBooleanLiteralContext(p *BooleanLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_booleanLiteral
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(ilang_parserTRUE, 0)
}

func (s *BooleanLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(ilang_parserFALSE, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitBooleanLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ilang_parserRULE_booleanLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ilang_parserTRUE || _la == ilang_parserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetItems returns the items rule context list.
	GetItems() []IExprContext

	// SetItems sets the items rule context list.
	SetItems([]IExprContext)

	// Getter signatures
	OPEN_BRACKET() antlr.TerminalNode
	CLOSE_BRACKET() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	_expr  IExprContext
	items  []IExprContext
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_arrayLiteral
	return p
}

func InitEmptyArrayLiteralContext(p *ArrayLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_arrayLiteral
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) Get_expr() IExprContext { return s._expr }

func (s *ArrayLiteralContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ArrayLiteralContext) GetItems() []IExprContext { return s.items }

func (s *ArrayLiteralContext) SetItems(v []IExprContext) { s.items = v }

func (s *ArrayLiteralContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(ilang_parserOPEN_BRACKET, 0)
}

func (s *ArrayLiteralContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(ilang_parserCLOSE_BRACKET, 0)
}

func (s *ArrayLiteralContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArrayLiteralContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrayLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ilang_parserCOMMA)
}

func (s *ArrayLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ilang_parserCOMMA, i)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitArrayLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ilang_parserRULE_arrayLiteral)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(ilang_parserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6315575338) != 0 {
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(248)

					var _x = p.expr(0)

					localctx.(*ArrayLiteralContext)._expr = _x
				}
				localctx.(*ArrayLiteralContext).items = append(localctx.(*ArrayLiteralContext).items, localctx.(*ArrayLiteralContext)._expr)
				{
					p.SetState(249)
					p.Match(ilang_parserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(255)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(256)

			var _x = p.expr(0)

			localctx.(*ArrayLiteralContext)._expr = _x
		}
		localctx.(*ArrayLiteralContext).items = append(localctx.(*ArrayLiteralContext).items, localctx.(*ArrayLiteralContext)._expr)

	}
	{
		p.SetState(259)
		p.Match(ilang_parserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMapLiteralContext is an interface to support dynamic dispatch.
type IMapLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_mapLiteralItem returns the _mapLiteralItem rule contexts.
	Get_mapLiteralItem() IMapLiteralItemContext

	// Set_mapLiteralItem sets the _mapLiteralItem rule contexts.
	Set_mapLiteralItem(IMapLiteralItemContext)

	// GetItems returns the items rule context list.
	GetItems() []IMapLiteralItemContext

	// SetItems sets the items rule context list.
	SetItems([]IMapLiteralItemContext)

	// Getter signatures
	OPEN_BRACE() antlr.TerminalNode
	CLOSE_BRACE() antlr.TerminalNode
	AllMapLiteralItem() []IMapLiteralItemContext
	MapLiteralItem(i int) IMapLiteralItemContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsMapLiteralContext differentiates from other interfaces.
	IsMapLiteralContext()
}

type MapLiteralContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	_mapLiteralItem IMapLiteralItemContext
	items           []IMapLiteralItemContext
}

func NewEmptyMapLiteralContext() *MapLiteralContext {
	var p = new(MapLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_mapLiteral
	return p
}

func InitEmptyMapLiteralContext(p *MapLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_mapLiteral
}

func (*MapLiteralContext) IsMapLiteralContext() {}

func NewMapLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapLiteralContext {
	var p = new(MapLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_mapLiteral

	return p
}

func (s *MapLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *MapLiteralContext) Get_mapLiteralItem() IMapLiteralItemContext { return s._mapLiteralItem }

func (s *MapLiteralContext) Set_mapLiteralItem(v IMapLiteralItemContext) { s._mapLiteralItem = v }

func (s *MapLiteralContext) GetItems() []IMapLiteralItemContext { return s.items }

func (s *MapLiteralContext) SetItems(v []IMapLiteralItemContext) { s.items = v }

func (s *MapLiteralContext) OPEN_BRACE() antlr.TerminalNode {
	return s.GetToken(ilang_parserOPEN_BRACE, 0)
}

func (s *MapLiteralContext) CLOSE_BRACE() antlr.TerminalNode {
	return s.GetToken(ilang_parserCLOSE_BRACE, 0)
}

func (s *MapLiteralContext) AllMapLiteralItem() []IMapLiteralItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMapLiteralItemContext); ok {
			len++
		}
	}

	tst := make([]IMapLiteralItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMapLiteralItemContext); ok {
			tst[i] = t.(IMapLiteralItemContext)
			i++
		}
	}

	return tst
}

func (s *MapLiteralContext) MapLiteralItem(i int) IMapLiteralItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapLiteralItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapLiteralItemContext)
}

func (s *MapLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ilang_parserCOMMA)
}

func (s *MapLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ilang_parserCOMMA, i)
}

func (s *MapLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitMapLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) MapLiteral() (localctx IMapLiteralContext) {
	localctx = NewMapLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ilang_parserRULE_mapLiteral)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.Match(ilang_parserOPEN_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ilang_parserOPEN_BRACKET || _la == ilang_parserSYMBOL {
		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(262)

					var _x = p.MapLiteralItem()

					localctx.(*MapLiteralContext)._mapLiteralItem = _x
				}
				localctx.(*MapLiteralContext).items = append(localctx.(*MapLiteralContext).items, localctx.(*MapLiteralContext)._mapLiteralItem)
				{
					p.SetState(263)
					p.Match(ilang_parserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(269)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(270)

			var _x = p.MapLiteralItem()

			localctx.(*MapLiteralContext)._mapLiteralItem = _x
		}
		localctx.(*MapLiteralContext).items = append(localctx.(*MapLiteralContext).items, localctx.(*MapLiteralContext)._mapLiteralItem)

	}
	{
		p.SetState(273)
		p.Match(ilang_parserCLOSE_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMapLiteralItemContext is an interface to support dynamic dispatch.
type IMapLiteralItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MapKey() IMapKeyContext
	COLON() antlr.TerminalNode
	Expr() IExprContext

	// IsMapLiteralItemContext differentiates from other interfaces.
	IsMapLiteralItemContext()
}

type MapLiteralItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapLiteralItemContext() *MapLiteralItemContext {
	var p = new(MapLiteralItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_mapLiteralItem
	return p
}

func InitEmptyMapLiteralItemContext(p *MapLiteralItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_mapLiteralItem
}

func (*MapLiteralItemContext) IsMapLiteralItemContext() {}

func NewMapLiteralItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapLiteralItemContext {
	var p = new(MapLiteralItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_mapLiteralItem

	return p
}

func (s *MapLiteralItemContext) GetParser() antlr.Parser { return s.parser }

func (s *MapLiteralItemContext) MapKey() IMapKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapKeyContext)
}

func (s *MapLiteralItemContext) COLON() antlr.TerminalNode {
	return s.GetToken(ilang_parserCOLON, 0)
}

func (s *MapLiteralItemContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MapLiteralItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapLiteralItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapLiteralItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitMapLiteralItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) MapLiteralItem() (localctx IMapLiteralItemContext) {
	localctx = NewMapLiteralItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ilang_parserRULE_mapLiteralItem)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(275)
		p.MapKey()
	}
	{
		p.SetState(276)
		p.Match(ilang_parserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(277)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMapKeyContext is an interface to support dynamic dispatch.
type IMapKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYMBOL() antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	Expr() IExprContext
	CLOSE_BRACKET() antlr.TerminalNode

	// IsMapKeyContext differentiates from other interfaces.
	IsMapKeyContext()
}

type MapKeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapKeyContext() *MapKeyContext {
	var p = new(MapKeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_mapKey
	return p
}

func InitEmptyMapKeyContext(p *MapKeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_mapKey
}

func (*MapKeyContext) IsMapKeyContext() {}

func NewMapKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapKeyContext {
	var p = new(MapKeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_mapKey

	return p
}

func (s *MapKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *MapKeyContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(ilang_parserSYMBOL, 0)
}

func (s *MapKeyContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(ilang_parserOPEN_BRACKET, 0)
}

func (s *MapKeyContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MapKeyContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(ilang_parserCLOSE_BRACKET, 0)
}

func (s *MapKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitMapKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) MapKey() (localctx IMapKeyContext) {
	localctx = NewMapKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ilang_parserRULE_mapKey)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ilang_parserSYMBOL:
		{
			p.SetState(279)
			p.Match(ilang_parserSYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ilang_parserOPEN_BRACKET:
		{
			p.SetState(280)
			p.Match(ilang_parserOPEN_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(281)
			p.expr(0)
		}
		{
			p.SetState(282)
			p.Match(ilang_parserCLOSE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGroupingContext is an interface to support dynamic dispatch.
type IGroupingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_PAREN() antlr.TerminalNode
	Expr() IExprContext
	CLOSE_PAREN() antlr.TerminalNode

	// IsGroupingContext differentiates from other interfaces.
	IsGroupingContext()
}

type GroupingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupingContext() *GroupingContext {
	var p = new(GroupingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_grouping
	return p
}

func InitEmptyGroupingContext(p *GroupingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ilang_parserRULE_grouping
}

func (*GroupingContext) IsGroupingContext() {}

func NewGroupingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupingContext {
	var p = new(GroupingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ilang_parserRULE_grouping

	return p
}

func (s *GroupingContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupingContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(ilang_parserOPEN_PAREN, 0)
}

func (s *GroupingContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GroupingContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(ilang_parserCLOSE_PAREN, 0)
}

func (s *GroupingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ilang_parserVisitor:
		return t.VisitGrouping(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ilang_parser) Grouping() (localctx IGroupingContext) {
	localctx = NewGroupingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ilang_parserRULE_grouping)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.Match(ilang_parserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(287)
		p.expr(0)
	}
	{
		p.SetState(288)
		p.Match(ilang_parserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *ilang_parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ilang_parser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
