// Code generated from ./ilang_lexer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ilang_lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var Ilang_lexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ilang_lexerLexerInit() {
	staticData := &Ilang_lexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"F_LETTER", "F_NUMBER", "F_QUOTE", "F_UNDERSCORE", "F_DOT", "F_STRING_INSIDE",
		"F_NEGATIVE", "OPEN_PAREN", "CLOSE_PAREN", "OPEN_BRACE", "CLOSE_BRACE",
		"OPEN_BRACKET", "CLOSE_BRACKET", "EQUALS", "ARITHMETIC_OP", "CONDITIONAL_OP",
		"BOOLEAN_OP", "NOT", "IF", "ELSE", "WHILE", "FOR", "IN", "RETURN", "LET",
		"DOT", "TRUE", "FALSE", "NULL", "QUESTION", "SEMICOLON", "COLON", "COMMA",
		"FUNC", "INT", "FLOAT", "STRING", "COMMENT", "SYMBOL", "WHITE_SPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 33, 264, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 3, 5, 95, 8, 5, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 118,
		8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3,
		15, 129, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 135, 8, 16, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29,
		1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 34, 3, 34, 203, 8, 34, 1, 34, 4, 34, 206, 8, 34, 11, 34, 12, 34,
		207, 1, 35, 3, 35, 211, 8, 35, 1, 35, 4, 35, 214, 8, 35, 11, 35, 12, 35,
		215, 3, 35, 218, 8, 35, 1, 35, 1, 35, 4, 35, 222, 8, 35, 11, 35, 12, 35,
		223, 1, 36, 1, 36, 5, 36, 228, 8, 36, 10, 36, 12, 36, 231, 9, 36, 1, 36,
		1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 5, 37, 239, 8, 37, 10, 37, 12, 37, 242,
		9, 37, 1, 37, 1, 37, 1, 38, 1, 38, 3, 38, 248, 8, 38, 1, 38, 1, 38, 1,
		38, 5, 38, 253, 8, 38, 10, 38, 12, 38, 256, 9, 38, 1, 39, 4, 39, 259, 8,
		39, 11, 39, 12, 39, 260, 1, 39, 1, 39, 0, 0, 40, 1, 0, 3, 0, 5, 0, 7, 0,
		9, 0, 11, 0, 13, 0, 15, 1, 17, 2, 19, 3, 21, 4, 23, 5, 25, 6, 27, 7, 29,
		8, 31, 9, 33, 10, 35, 11, 37, 12, 39, 13, 41, 14, 43, 15, 45, 16, 47, 17,
		49, 18, 51, 19, 53, 20, 55, 21, 57, 22, 59, 23, 61, 24, 63, 25, 65, 26,
		67, 27, 69, 28, 71, 29, 73, 30, 75, 31, 77, 32, 79, 33, 1, 0, 7, 2, 0,
		65, 90, 97, 122, 1, 0, 48, 57, 3, 0, 10, 10, 13, 13, 34, 34, 4, 0, 37,
		37, 42, 43, 45, 45, 47, 47, 2, 0, 60, 60, 62, 62, 2, 0, 10, 10, 13, 13,
		3, 0, 9, 10, 13, 13, 32, 32, 277, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0,
		0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0,
		0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0,
		0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1,
		0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49,
		1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0,
		57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0,
		0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0,
		0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0,
		0, 0, 1, 81, 1, 0, 0, 0, 3, 83, 1, 0, 0, 0, 5, 85, 1, 0, 0, 0, 7, 87, 1,
		0, 0, 0, 9, 89, 1, 0, 0, 0, 11, 94, 1, 0, 0, 0, 13, 96, 1, 0, 0, 0, 15,
		98, 1, 0, 0, 0, 17, 100, 1, 0, 0, 0, 19, 102, 1, 0, 0, 0, 21, 104, 1, 0,
		0, 0, 23, 106, 1, 0, 0, 0, 25, 108, 1, 0, 0, 0, 27, 110, 1, 0, 0, 0, 29,
		117, 1, 0, 0, 0, 31, 128, 1, 0, 0, 0, 33, 134, 1, 0, 0, 0, 35, 136, 1,
		0, 0, 0, 37, 138, 1, 0, 0, 0, 39, 141, 1, 0, 0, 0, 41, 146, 1, 0, 0, 0,
		43, 152, 1, 0, 0, 0, 45, 156, 1, 0, 0, 0, 47, 159, 1, 0, 0, 0, 49, 166,
		1, 0, 0, 0, 51, 170, 1, 0, 0, 0, 53, 172, 1, 0, 0, 0, 55, 177, 1, 0, 0,
		0, 57, 183, 1, 0, 0, 0, 59, 188, 1, 0, 0, 0, 61, 190, 1, 0, 0, 0, 63, 192,
		1, 0, 0, 0, 65, 194, 1, 0, 0, 0, 67, 196, 1, 0, 0, 0, 69, 202, 1, 0, 0,
		0, 71, 210, 1, 0, 0, 0, 73, 225, 1, 0, 0, 0, 75, 234, 1, 0, 0, 0, 77, 247,
		1, 0, 0, 0, 79, 258, 1, 0, 0, 0, 81, 82, 7, 0, 0, 0, 82, 2, 1, 0, 0, 0,
		83, 84, 7, 1, 0, 0, 84, 4, 1, 0, 0, 0, 85, 86, 5, 34, 0, 0, 86, 6, 1, 0,
		0, 0, 87, 88, 5, 95, 0, 0, 88, 8, 1, 0, 0, 0, 89, 90, 5, 46, 0, 0, 90,
		10, 1, 0, 0, 0, 91, 92, 5, 92, 0, 0, 92, 95, 5, 34, 0, 0, 93, 95, 8, 2,
		0, 0, 94, 91, 1, 0, 0, 0, 94, 93, 1, 0, 0, 0, 95, 12, 1, 0, 0, 0, 96, 97,
		5, 45, 0, 0, 97, 14, 1, 0, 0, 0, 98, 99, 5, 40, 0, 0, 99, 16, 1, 0, 0,
		0, 100, 101, 5, 41, 0, 0, 101, 18, 1, 0, 0, 0, 102, 103, 5, 123, 0, 0,
		103, 20, 1, 0, 0, 0, 104, 105, 5, 125, 0, 0, 105, 22, 1, 0, 0, 0, 106,
		107, 5, 91, 0, 0, 107, 24, 1, 0, 0, 0, 108, 109, 5, 93, 0, 0, 109, 26,
		1, 0, 0, 0, 110, 111, 5, 61, 0, 0, 111, 28, 1, 0, 0, 0, 112, 118, 7, 3,
		0, 0, 113, 114, 5, 47, 0, 0, 114, 118, 5, 47, 0, 0, 115, 116, 5, 42, 0,
		0, 116, 118, 5, 42, 0, 0, 117, 112, 1, 0, 0, 0, 117, 113, 1, 0, 0, 0, 117,
		115, 1, 0, 0, 0, 118, 30, 1, 0, 0, 0, 119, 129, 7, 4, 0, 0, 120, 121, 5,
		62, 0, 0, 121, 129, 5, 61, 0, 0, 122, 123, 5, 60, 0, 0, 123, 129, 5, 61,
		0, 0, 124, 125, 5, 61, 0, 0, 125, 129, 5, 61, 0, 0, 126, 127, 5, 33, 0,
		0, 127, 129, 5, 61, 0, 0, 128, 119, 1, 0, 0, 0, 128, 120, 1, 0, 0, 0, 128,
		122, 1, 0, 0, 0, 128, 124, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 129, 32, 1,
		0, 0, 0, 130, 131, 5, 38, 0, 0, 131, 135, 5, 38, 0, 0, 132, 133, 5, 124,
		0, 0, 133, 135, 5, 124, 0, 0, 134, 130, 1, 0, 0, 0, 134, 132, 1, 0, 0,
		0, 135, 34, 1, 0, 0, 0, 136, 137, 5, 33, 0, 0, 137, 36, 1, 0, 0, 0, 138,
		139, 5, 105, 0, 0, 139, 140, 5, 102, 0, 0, 140, 38, 1, 0, 0, 0, 141, 142,
		5, 101, 0, 0, 142, 143, 5, 108, 0, 0, 143, 144, 5, 115, 0, 0, 144, 145,
		5, 101, 0, 0, 145, 40, 1, 0, 0, 0, 146, 147, 5, 119, 0, 0, 147, 148, 5,
		104, 0, 0, 148, 149, 5, 105, 0, 0, 149, 150, 5, 108, 0, 0, 150, 151, 5,
		101, 0, 0, 151, 42, 1, 0, 0, 0, 152, 153, 5, 102, 0, 0, 153, 154, 5, 111,
		0, 0, 154, 155, 5, 114, 0, 0, 155, 44, 1, 0, 0, 0, 156, 157, 5, 105, 0,
		0, 157, 158, 5, 110, 0, 0, 158, 46, 1, 0, 0, 0, 159, 160, 5, 114, 0, 0,
		160, 161, 5, 101, 0, 0, 161, 162, 5, 116, 0, 0, 162, 163, 5, 117, 0, 0,
		163, 164, 5, 114, 0, 0, 164, 165, 5, 110, 0, 0, 165, 48, 1, 0, 0, 0, 166,
		167, 5, 108, 0, 0, 167, 168, 5, 101, 0, 0, 168, 169, 5, 116, 0, 0, 169,
		50, 1, 0, 0, 0, 170, 171, 3, 9, 4, 0, 171, 52, 1, 0, 0, 0, 172, 173, 5,
		116, 0, 0, 173, 174, 5, 114, 0, 0, 174, 175, 5, 117, 0, 0, 175, 176, 5,
		101, 0, 0, 176, 54, 1, 0, 0, 0, 177, 178, 5, 102, 0, 0, 178, 179, 5, 97,
		0, 0, 179, 180, 5, 108, 0, 0, 180, 181, 5, 115, 0, 0, 181, 182, 5, 101,
		0, 0, 182, 56, 1, 0, 0, 0, 183, 184, 5, 110, 0, 0, 184, 185, 5, 117, 0,
		0, 185, 186, 5, 108, 0, 0, 186, 187, 5, 108, 0, 0, 187, 58, 1, 0, 0, 0,
		188, 189, 5, 63, 0, 0, 189, 60, 1, 0, 0, 0, 190, 191, 5, 59, 0, 0, 191,
		62, 1, 0, 0, 0, 192, 193, 5, 58, 0, 0, 193, 64, 1, 0, 0, 0, 194, 195, 5,
		44, 0, 0, 195, 66, 1, 0, 0, 0, 196, 197, 5, 102, 0, 0, 197, 198, 5, 117,
		0, 0, 198, 199, 5, 110, 0, 0, 199, 200, 5, 99, 0, 0, 200, 68, 1, 0, 0,
		0, 201, 203, 3, 13, 6, 0, 202, 201, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203,
		205, 1, 0, 0, 0, 204, 206, 3, 3, 1, 0, 205, 204, 1, 0, 0, 0, 206, 207,
		1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 70, 1, 0,
		0, 0, 209, 211, 3, 13, 6, 0, 210, 209, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0,
		211, 217, 1, 0, 0, 0, 212, 214, 3, 3, 1, 0, 213, 212, 1, 0, 0, 0, 214,
		215, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 218,
		1, 0, 0, 0, 217, 213, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 219, 1, 0,
		0, 0, 219, 221, 3, 9, 4, 0, 220, 222, 3, 3, 1, 0, 221, 220, 1, 0, 0, 0,
		222, 223, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224,
		72, 1, 0, 0, 0, 225, 229, 3, 5, 2, 0, 226, 228, 3, 11, 5, 0, 227, 226,
		1, 0, 0, 0, 228, 231, 1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 229, 230, 1, 0,
		0, 0, 230, 232, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 232, 233, 3, 5, 2, 0,
		233, 74, 1, 0, 0, 0, 234, 235, 5, 47, 0, 0, 235, 236, 5, 47, 0, 0, 236,
		240, 1, 0, 0, 0, 237, 239, 8, 5, 0, 0, 238, 237, 1, 0, 0, 0, 239, 242,
		1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 243, 1, 0,
		0, 0, 242, 240, 1, 0, 0, 0, 243, 244, 6, 37, 0, 0, 244, 76, 1, 0, 0, 0,
		245, 248, 3, 1, 0, 0, 246, 248, 3, 7, 3, 0, 247, 245, 1, 0, 0, 0, 247,
		246, 1, 0, 0, 0, 248, 254, 1, 0, 0, 0, 249, 253, 3, 1, 0, 0, 250, 253,
		3, 3, 1, 0, 251, 253, 3, 7, 3, 0, 252, 249, 1, 0, 0, 0, 252, 250, 1, 0,
		0, 0, 252, 251, 1, 0, 0, 0, 253, 256, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0,
		254, 255, 1, 0, 0, 0, 255, 78, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0, 257, 259,
		7, 6, 0, 0, 258, 257, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 258, 1, 0,
		0, 0, 260, 261, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 263, 6, 39, 0, 0,
		263, 80, 1, 0, 0, 0, 17, 0, 94, 117, 128, 134, 202, 207, 210, 215, 217,
		223, 229, 240, 247, 252, 254, 260, 1, 6, 0, 0,
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

// ilang_lexerInit initializes any static state used to implement ilang_lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// Newilang_lexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Ilang_lexerInit() {
	staticData := &Ilang_lexerLexerStaticData
	staticData.once.Do(ilang_lexerLexerInit)
}

// Newilang_lexer produces a new lexer instance for the optional input antlr.CharStream.
func Newilang_lexer(input antlr.CharStream) *ilang_lexer {
	Ilang_lexerInit()
	l := new(ilang_lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &Ilang_lexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "ilang_lexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ilang_lexer tokens.
const (
	ilang_lexerOPEN_PAREN     = 1
	ilang_lexerCLOSE_PAREN    = 2
	ilang_lexerOPEN_BRACE     = 3
	ilang_lexerCLOSE_BRACE    = 4
	ilang_lexerOPEN_BRACKET   = 5
	ilang_lexerCLOSE_BRACKET  = 6
	ilang_lexerEQUALS         = 7
	ilang_lexerARITHMETIC_OP  = 8
	ilang_lexerCONDITIONAL_OP = 9
	ilang_lexerBOOLEAN_OP     = 10
	ilang_lexerNOT            = 11
	ilang_lexerIF             = 12
	ilang_lexerELSE           = 13
	ilang_lexerWHILE          = 14
	ilang_lexerFOR            = 15
	ilang_lexerIN             = 16
	ilang_lexerRETURN         = 17
	ilang_lexerLET            = 18
	ilang_lexerDOT            = 19
	ilang_lexerTRUE           = 20
	ilang_lexerFALSE          = 21
	ilang_lexerNULL           = 22
	ilang_lexerQUESTION       = 23
	ilang_lexerSEMICOLON      = 24
	ilang_lexerCOLON          = 25
	ilang_lexerCOMMA          = 26
	ilang_lexerFUNC           = 27
	ilang_lexerINT            = 28
	ilang_lexerFLOAT          = 29
	ilang_lexerSTRING         = 30
	ilang_lexerCOMMENT        = 31
	ilang_lexerSYMBOL         = 32
	ilang_lexerWHITE_SPACE    = 33
)
