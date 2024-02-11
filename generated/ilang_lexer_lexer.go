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
		"", "'('", "')'", "'{'", "'}'", "'='", "", "", "", "'!'", "'if'", "'else'",
		"'while'", "'for'", "'in'", "'let'", "", "'true'", "'false'", "'null'",
		"'?'", "';'", "','", "'func'",
	}
	staticData.SymbolicNames = []string{
		"", "OPEN_PAREN", "CLOSE_PAREN", "OPEN_BRACE", "CLOSE_BRACE", "EQUALS",
		"ARITHMETIC_OP", "CONDITIONAL_OP", "BOOLEAN_OP", "NOT", "IF", "ELSE",
		"WHILE", "FOR", "IN", "LET", "DOT", "TRUE", "FALSE", "NULL", "QUESTION",
		"SEMICOLON", "COMMA", "FUNC", "INT", "FLOAT", "STRING", "COMMENT", "SYMBOL",
		"WHITE_SPACE",
	}
	staticData.RuleNames = []string{
		"F_LETTER", "F_NUMBER", "F_QUOTE", "F_UNDERSCORE", "F_DOT", "F_STRING_INSIDE",
		"F_NEGATIVE", "OPEN_PAREN", "CLOSE_PAREN", "OPEN_BRACE", "CLOSE_BRACE",
		"EQUALS", "ARITHMETIC_OP", "CONDITIONAL_OP", "BOOLEAN_OP", "NOT", "IF",
		"ELSE", "WHILE", "FOR", "IN", "LET", "DOT", "TRUE", "FALSE", "NULL",
		"QUESTION", "SEMICOLON", "COMMA", "FUNC", "INT", "FLOAT", "STRING",
		"COMMENT", "SYMBOL", "WHITE_SPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 29, 243, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		3, 5, 87, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 106, 8,
		12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13,
		117, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 123, 8, 14, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 30, 3, 30, 182, 8, 30, 1, 30, 4, 30, 185, 8, 30, 11, 30,
		12, 30, 186, 1, 31, 3, 31, 190, 8, 31, 1, 31, 4, 31, 193, 8, 31, 11, 31,
		12, 31, 194, 3, 31, 197, 8, 31, 1, 31, 1, 31, 4, 31, 201, 8, 31, 11, 31,
		12, 31, 202, 1, 32, 1, 32, 5, 32, 207, 8, 32, 10, 32, 12, 32, 210, 9, 32,
		1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 5, 33, 218, 8, 33, 10, 33, 12,
		33, 221, 9, 33, 1, 33, 1, 33, 1, 34, 1, 34, 3, 34, 227, 8, 34, 1, 34, 1,
		34, 1, 34, 5, 34, 232, 8, 34, 10, 34, 12, 34, 235, 9, 34, 1, 35, 4, 35,
		238, 8, 35, 11, 35, 12, 35, 239, 1, 35, 1, 35, 0, 0, 36, 1, 0, 3, 0, 5,
		0, 7, 0, 9, 0, 11, 0, 13, 0, 15, 1, 17, 2, 19, 3, 21, 4, 23, 5, 25, 6,
		27, 7, 29, 8, 31, 9, 33, 10, 35, 11, 37, 12, 39, 13, 41, 14, 43, 15, 45,
		16, 47, 17, 49, 18, 51, 19, 53, 20, 55, 21, 57, 22, 59, 23, 61, 24, 63,
		25, 65, 26, 67, 27, 69, 28, 71, 29, 1, 0, 7, 2, 0, 65, 90, 97, 122, 1,
		0, 48, 57, 3, 0, 10, 10, 13, 13, 34, 34, 4, 0, 37, 37, 42, 43, 45, 45,
		47, 47, 2, 0, 60, 60, 62, 62, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13,
		32, 32, 256, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1,
		0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59,
		1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0,
		67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 1, 73, 1, 0, 0, 0,
		3, 75, 1, 0, 0, 0, 5, 77, 1, 0, 0, 0, 7, 79, 1, 0, 0, 0, 9, 81, 1, 0, 0,
		0, 11, 86, 1, 0, 0, 0, 13, 88, 1, 0, 0, 0, 15, 90, 1, 0, 0, 0, 17, 92,
		1, 0, 0, 0, 19, 94, 1, 0, 0, 0, 21, 96, 1, 0, 0, 0, 23, 98, 1, 0, 0, 0,
		25, 105, 1, 0, 0, 0, 27, 116, 1, 0, 0, 0, 29, 122, 1, 0, 0, 0, 31, 124,
		1, 0, 0, 0, 33, 126, 1, 0, 0, 0, 35, 129, 1, 0, 0, 0, 37, 134, 1, 0, 0,
		0, 39, 140, 1, 0, 0, 0, 41, 144, 1, 0, 0, 0, 43, 147, 1, 0, 0, 0, 45, 151,
		1, 0, 0, 0, 47, 153, 1, 0, 0, 0, 49, 158, 1, 0, 0, 0, 51, 164, 1, 0, 0,
		0, 53, 169, 1, 0, 0, 0, 55, 171, 1, 0, 0, 0, 57, 173, 1, 0, 0, 0, 59, 175,
		1, 0, 0, 0, 61, 181, 1, 0, 0, 0, 63, 189, 1, 0, 0, 0, 65, 204, 1, 0, 0,
		0, 67, 213, 1, 0, 0, 0, 69, 226, 1, 0, 0, 0, 71, 237, 1, 0, 0, 0, 73, 74,
		7, 0, 0, 0, 74, 2, 1, 0, 0, 0, 75, 76, 7, 1, 0, 0, 76, 4, 1, 0, 0, 0, 77,
		78, 5, 34, 0, 0, 78, 6, 1, 0, 0, 0, 79, 80, 5, 95, 0, 0, 80, 8, 1, 0, 0,
		0, 81, 82, 5, 46, 0, 0, 82, 10, 1, 0, 0, 0, 83, 84, 5, 92, 0, 0, 84, 87,
		5, 34, 0, 0, 85, 87, 8, 2, 0, 0, 86, 83, 1, 0, 0, 0, 86, 85, 1, 0, 0, 0,
		87, 12, 1, 0, 0, 0, 88, 89, 5, 45, 0, 0, 89, 14, 1, 0, 0, 0, 90, 91, 5,
		40, 0, 0, 91, 16, 1, 0, 0, 0, 92, 93, 5, 41, 0, 0, 93, 18, 1, 0, 0, 0,
		94, 95, 5, 123, 0, 0, 95, 20, 1, 0, 0, 0, 96, 97, 5, 125, 0, 0, 97, 22,
		1, 0, 0, 0, 98, 99, 5, 61, 0, 0, 99, 24, 1, 0, 0, 0, 100, 106, 7, 3, 0,
		0, 101, 102, 5, 47, 0, 0, 102, 106, 5, 47, 0, 0, 103, 104, 5, 42, 0, 0,
		104, 106, 5, 42, 0, 0, 105, 100, 1, 0, 0, 0, 105, 101, 1, 0, 0, 0, 105,
		103, 1, 0, 0, 0, 106, 26, 1, 0, 0, 0, 107, 117, 7, 4, 0, 0, 108, 109, 5,
		62, 0, 0, 109, 117, 5, 61, 0, 0, 110, 111, 5, 60, 0, 0, 111, 117, 5, 61,
		0, 0, 112, 113, 5, 61, 0, 0, 113, 117, 5, 61, 0, 0, 114, 115, 5, 33, 0,
		0, 115, 117, 5, 61, 0, 0, 116, 107, 1, 0, 0, 0, 116, 108, 1, 0, 0, 0, 116,
		110, 1, 0, 0, 0, 116, 112, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117, 28, 1,
		0, 0, 0, 118, 119, 5, 38, 0, 0, 119, 123, 5, 38, 0, 0, 120, 121, 5, 124,
		0, 0, 121, 123, 5, 124, 0, 0, 122, 118, 1, 0, 0, 0, 122, 120, 1, 0, 0,
		0, 123, 30, 1, 0, 0, 0, 124, 125, 5, 33, 0, 0, 125, 32, 1, 0, 0, 0, 126,
		127, 5, 105, 0, 0, 127, 128, 5, 102, 0, 0, 128, 34, 1, 0, 0, 0, 129, 130,
		5, 101, 0, 0, 130, 131, 5, 108, 0, 0, 131, 132, 5, 115, 0, 0, 132, 133,
		5, 101, 0, 0, 133, 36, 1, 0, 0, 0, 134, 135, 5, 119, 0, 0, 135, 136, 5,
		104, 0, 0, 136, 137, 5, 105, 0, 0, 137, 138, 5, 108, 0, 0, 138, 139, 5,
		101, 0, 0, 139, 38, 1, 0, 0, 0, 140, 141, 5, 102, 0, 0, 141, 142, 5, 111,
		0, 0, 142, 143, 5, 114, 0, 0, 143, 40, 1, 0, 0, 0, 144, 145, 5, 105, 0,
		0, 145, 146, 5, 110, 0, 0, 146, 42, 1, 0, 0, 0, 147, 148, 5, 108, 0, 0,
		148, 149, 5, 101, 0, 0, 149, 150, 5, 116, 0, 0, 150, 44, 1, 0, 0, 0, 151,
		152, 3, 9, 4, 0, 152, 46, 1, 0, 0, 0, 153, 154, 5, 116, 0, 0, 154, 155,
		5, 114, 0, 0, 155, 156, 5, 117, 0, 0, 156, 157, 5, 101, 0, 0, 157, 48,
		1, 0, 0, 0, 158, 159, 5, 102, 0, 0, 159, 160, 5, 97, 0, 0, 160, 161, 5,
		108, 0, 0, 161, 162, 5, 115, 0, 0, 162, 163, 5, 101, 0, 0, 163, 50, 1,
		0, 0, 0, 164, 165, 5, 110, 0, 0, 165, 166, 5, 117, 0, 0, 166, 167, 5, 108,
		0, 0, 167, 168, 5, 108, 0, 0, 168, 52, 1, 0, 0, 0, 169, 170, 5, 63, 0,
		0, 170, 54, 1, 0, 0, 0, 171, 172, 5, 59, 0, 0, 172, 56, 1, 0, 0, 0, 173,
		174, 5, 44, 0, 0, 174, 58, 1, 0, 0, 0, 175, 176, 5, 102, 0, 0, 176, 177,
		5, 117, 0, 0, 177, 178, 5, 110, 0, 0, 178, 179, 5, 99, 0, 0, 179, 60, 1,
		0, 0, 0, 180, 182, 3, 13, 6, 0, 181, 180, 1, 0, 0, 0, 181, 182, 1, 0, 0,
		0, 182, 184, 1, 0, 0, 0, 183, 185, 3, 3, 1, 0, 184, 183, 1, 0, 0, 0, 185,
		186, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 62, 1,
		0, 0, 0, 188, 190, 3, 13, 6, 0, 189, 188, 1, 0, 0, 0, 189, 190, 1, 0, 0,
		0, 190, 196, 1, 0, 0, 0, 191, 193, 3, 3, 1, 0, 192, 191, 1, 0, 0, 0, 193,
		194, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 197,
		1, 0, 0, 0, 196, 192, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 198, 1, 0,
		0, 0, 198, 200, 3, 9, 4, 0, 199, 201, 3, 3, 1, 0, 200, 199, 1, 0, 0, 0,
		201, 202, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203,
		64, 1, 0, 0, 0, 204, 208, 3, 5, 2, 0, 205, 207, 3, 11, 5, 0, 206, 205,
		1, 0, 0, 0, 207, 210, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 208, 209, 1, 0,
		0, 0, 209, 211, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 211, 212, 3, 5, 2, 0,
		212, 66, 1, 0, 0, 0, 213, 214, 5, 47, 0, 0, 214, 215, 5, 47, 0, 0, 215,
		219, 1, 0, 0, 0, 216, 218, 8, 5, 0, 0, 217, 216, 1, 0, 0, 0, 218, 221,
		1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 222, 1, 0,
		0, 0, 221, 219, 1, 0, 0, 0, 222, 223, 6, 33, 0, 0, 223, 68, 1, 0, 0, 0,
		224, 227, 3, 1, 0, 0, 225, 227, 3, 7, 3, 0, 226, 224, 1, 0, 0, 0, 226,
		225, 1, 0, 0, 0, 227, 233, 1, 0, 0, 0, 228, 232, 3, 1, 0, 0, 229, 232,
		3, 3, 1, 0, 230, 232, 3, 7, 3, 0, 231, 228, 1, 0, 0, 0, 231, 229, 1, 0,
		0, 0, 231, 230, 1, 0, 0, 0, 232, 235, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0,
		233, 234, 1, 0, 0, 0, 234, 70, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 236, 238,
		7, 6, 0, 0, 237, 236, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 237, 1, 0,
		0, 0, 239, 240, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 242, 6, 35, 0, 0,
		242, 72, 1, 0, 0, 0, 17, 0, 86, 105, 116, 122, 181, 186, 189, 194, 196,
		202, 208, 219, 226, 231, 233, 239, 1, 6, 0, 0,
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
	ilang_lexerEQUALS         = 5
	ilang_lexerARITHMETIC_OP  = 6
	ilang_lexerCONDITIONAL_OP = 7
	ilang_lexerBOOLEAN_OP     = 8
	ilang_lexerNOT            = 9
	ilang_lexerIF             = 10
	ilang_lexerELSE           = 11
	ilang_lexerWHILE          = 12
	ilang_lexerFOR            = 13
	ilang_lexerIN             = 14
	ilang_lexerLET            = 15
	ilang_lexerDOT            = 16
	ilang_lexerTRUE           = 17
	ilang_lexerFALSE          = 18
	ilang_lexerNULL           = 19
	ilang_lexerQUESTION       = 20
	ilang_lexerSEMICOLON      = 21
	ilang_lexerCOMMA          = 22
	ilang_lexerFUNC           = 23
	ilang_lexerINT            = 24
	ilang_lexerFLOAT          = 25
	ilang_lexerSTRING         = 26
	ilang_lexerCOMMENT        = 27
	ilang_lexerSYMBOL         = 28
	ilang_lexerWHITE_SPACE    = 29
)
