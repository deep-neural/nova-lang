// Code generated from CustomLanguage.g4 by ANTLR 4.13.1. DO NOT EDIT.

package main

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

type CustomLanguageLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CustomLanguageLexerLexerStaticData struct {
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

func customlanguagelexerLexerInit() {
	staticData := &CustomLanguageLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'func'", "'int'", "'float'", "'string'", "'bool'", "'return'",
		"'var'", "", "'->'", "'='", "'('", "')'", "'{'", "'}'", "','", "';'",
		"':'", "'+'", "'-'", "'*'", "'/'",
	}
	staticData.SymbolicNames = []string{
		"", "FUNC", "INT", "FLOAT", "STRING", "BOOL", "RETURN", "VAR", "BOOL_LITERAL",
		"ARROW", "ASSIGN", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "COMMA",
		"SEMI", "COLON", "PLUS", "MINUS", "MULT", "DIV", "STRING_LITERAL", "FLOAT_LITERAL",
		"INT_LITERAL", "ID", "WS", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"FUNC", "INT", "FLOAT", "STRING", "BOOL", "RETURN", "VAR", "BOOL_LITERAL",
		"ARROW", "ASSIGN", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "COMMA",
		"SEMI", "COLON", "PLUS", "MINUS", "MULT", "DIV", "STRING_LITERAL", "ESC",
		"FLOAT_LITERAL", "INT_LITERAL", "ID", "WS", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 28, 212, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 107, 8, 7, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 5, 21, 139, 8, 21, 10, 21, 12,
		21, 142, 9, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 4, 23, 150, 8,
		23, 11, 23, 12, 23, 151, 1, 23, 1, 23, 5, 23, 156, 8, 23, 10, 23, 12, 23,
		159, 9, 23, 1, 23, 1, 23, 4, 23, 163, 8, 23, 11, 23, 12, 23, 164, 3, 23,
		167, 8, 23, 1, 24, 4, 24, 170, 8, 24, 11, 24, 12, 24, 171, 1, 25, 1, 25,
		5, 25, 176, 8, 25, 10, 25, 12, 25, 179, 9, 25, 1, 26, 4, 26, 182, 8, 26,
		11, 26, 12, 26, 183, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 192,
		8, 27, 10, 27, 12, 27, 195, 9, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1,
		28, 5, 28, 203, 8, 28, 10, 28, 12, 28, 206, 9, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 204, 0, 29, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33,
		17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 0, 47, 23, 49, 24, 51,
		25, 53, 26, 55, 27, 57, 28, 1, 0, 7, 2, 0, 34, 34, 92, 92, 8, 0, 34, 34,
		39, 39, 92, 92, 98, 98, 102, 102, 110, 110, 114, 114, 116, 116, 1, 0, 48,
		57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122,
		3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 222, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33,
		1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0,
		41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0,
		0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0,
		0, 1, 59, 1, 0, 0, 0, 3, 64, 1, 0, 0, 0, 5, 68, 1, 0, 0, 0, 7, 74, 1, 0,
		0, 0, 9, 81, 1, 0, 0, 0, 11, 86, 1, 0, 0, 0, 13, 93, 1, 0, 0, 0, 15, 106,
		1, 0, 0, 0, 17, 108, 1, 0, 0, 0, 19, 111, 1, 0, 0, 0, 21, 113, 1, 0, 0,
		0, 23, 115, 1, 0, 0, 0, 25, 117, 1, 0, 0, 0, 27, 119, 1, 0, 0, 0, 29, 121,
		1, 0, 0, 0, 31, 123, 1, 0, 0, 0, 33, 125, 1, 0, 0, 0, 35, 127, 1, 0, 0,
		0, 37, 129, 1, 0, 0, 0, 39, 131, 1, 0, 0, 0, 41, 133, 1, 0, 0, 0, 43, 135,
		1, 0, 0, 0, 45, 145, 1, 0, 0, 0, 47, 166, 1, 0, 0, 0, 49, 169, 1, 0, 0,
		0, 51, 173, 1, 0, 0, 0, 53, 181, 1, 0, 0, 0, 55, 187, 1, 0, 0, 0, 57, 198,
		1, 0, 0, 0, 59, 60, 5, 102, 0, 0, 60, 61, 5, 117, 0, 0, 61, 62, 5, 110,
		0, 0, 62, 63, 5, 99, 0, 0, 63, 2, 1, 0, 0, 0, 64, 65, 5, 105, 0, 0, 65,
		66, 5, 110, 0, 0, 66, 67, 5, 116, 0, 0, 67, 4, 1, 0, 0, 0, 68, 69, 5, 102,
		0, 0, 69, 70, 5, 108, 0, 0, 70, 71, 5, 111, 0, 0, 71, 72, 5, 97, 0, 0,
		72, 73, 5, 116, 0, 0, 73, 6, 1, 0, 0, 0, 74, 75, 5, 115, 0, 0, 75, 76,
		5, 116, 0, 0, 76, 77, 5, 114, 0, 0, 77, 78, 5, 105, 0, 0, 78, 79, 5, 110,
		0, 0, 79, 80, 5, 103, 0, 0, 80, 8, 1, 0, 0, 0, 81, 82, 5, 98, 0, 0, 82,
		83, 5, 111, 0, 0, 83, 84, 5, 111, 0, 0, 84, 85, 5, 108, 0, 0, 85, 10, 1,
		0, 0, 0, 86, 87, 5, 114, 0, 0, 87, 88, 5, 101, 0, 0, 88, 89, 5, 116, 0,
		0, 89, 90, 5, 117, 0, 0, 90, 91, 5, 114, 0, 0, 91, 92, 5, 110, 0, 0, 92,
		12, 1, 0, 0, 0, 93, 94, 5, 118, 0, 0, 94, 95, 5, 97, 0, 0, 95, 96, 5, 114,
		0, 0, 96, 14, 1, 0, 0, 0, 97, 98, 5, 116, 0, 0, 98, 99, 5, 114, 0, 0, 99,
		100, 5, 117, 0, 0, 100, 107, 5, 101, 0, 0, 101, 102, 5, 102, 0, 0, 102,
		103, 5, 97, 0, 0, 103, 104, 5, 108, 0, 0, 104, 105, 5, 115, 0, 0, 105,
		107, 5, 101, 0, 0, 106, 97, 1, 0, 0, 0, 106, 101, 1, 0, 0, 0, 107, 16,
		1, 0, 0, 0, 108, 109, 5, 45, 0, 0, 109, 110, 5, 62, 0, 0, 110, 18, 1, 0,
		0, 0, 111, 112, 5, 61, 0, 0, 112, 20, 1, 0, 0, 0, 113, 114, 5, 40, 0, 0,
		114, 22, 1, 0, 0, 0, 115, 116, 5, 41, 0, 0, 116, 24, 1, 0, 0, 0, 117, 118,
		5, 123, 0, 0, 118, 26, 1, 0, 0, 0, 119, 120, 5, 125, 0, 0, 120, 28, 1,
		0, 0, 0, 121, 122, 5, 44, 0, 0, 122, 30, 1, 0, 0, 0, 123, 124, 5, 59, 0,
		0, 124, 32, 1, 0, 0, 0, 125, 126, 5, 58, 0, 0, 126, 34, 1, 0, 0, 0, 127,
		128, 5, 43, 0, 0, 128, 36, 1, 0, 0, 0, 129, 130, 5, 45, 0, 0, 130, 38,
		1, 0, 0, 0, 131, 132, 5, 42, 0, 0, 132, 40, 1, 0, 0, 0, 133, 134, 5, 47,
		0, 0, 134, 42, 1, 0, 0, 0, 135, 140, 5, 34, 0, 0, 136, 139, 3, 45, 22,
		0, 137, 139, 8, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 137, 1, 0, 0, 0, 139,
		142, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 143,
		1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 143, 144, 5, 34, 0, 0, 144, 44, 1, 0,
		0, 0, 145, 146, 5, 92, 0, 0, 146, 147, 7, 1, 0, 0, 147, 46, 1, 0, 0, 0,
		148, 150, 7, 2, 0, 0, 149, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151,
		149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 157,
		5, 46, 0, 0, 154, 156, 7, 2, 0, 0, 155, 154, 1, 0, 0, 0, 156, 159, 1, 0,
		0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 167, 1, 0, 0, 0,
		159, 157, 1, 0, 0, 0, 160, 162, 5, 46, 0, 0, 161, 163, 7, 2, 0, 0, 162,
		161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165,
		1, 0, 0, 0, 165, 167, 1, 0, 0, 0, 166, 149, 1, 0, 0, 0, 166, 160, 1, 0,
		0, 0, 167, 48, 1, 0, 0, 0, 168, 170, 7, 2, 0, 0, 169, 168, 1, 0, 0, 0,
		170, 171, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172,
		50, 1, 0, 0, 0, 173, 177, 7, 3, 0, 0, 174, 176, 7, 4, 0, 0, 175, 174, 1,
		0, 0, 0, 176, 179, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0,
		0, 178, 52, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 180, 182, 7, 5, 0, 0, 181,
		180, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184,
		1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 186, 6, 26, 0, 0, 186, 54, 1, 0,
		0, 0, 187, 188, 5, 47, 0, 0, 188, 189, 5, 47, 0, 0, 189, 193, 1, 0, 0,
		0, 190, 192, 8, 6, 0, 0, 191, 190, 1, 0, 0, 0, 192, 195, 1, 0, 0, 0, 193,
		191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 196, 1, 0, 0, 0, 195, 193,
		1, 0, 0, 0, 196, 197, 6, 27, 0, 0, 197, 56, 1, 0, 0, 0, 198, 199, 5, 47,
		0, 0, 199, 200, 5, 42, 0, 0, 200, 204, 1, 0, 0, 0, 201, 203, 9, 0, 0, 0,
		202, 201, 1, 0, 0, 0, 203, 206, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 204,
		202, 1, 0, 0, 0, 205, 207, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 207, 208,
		5, 42, 0, 0, 208, 209, 5, 47, 0, 0, 209, 210, 1, 0, 0, 0, 210, 211, 6,
		28, 0, 0, 211, 58, 1, 0, 0, 0, 13, 0, 106, 138, 140, 151, 157, 164, 166,
		171, 177, 183, 193, 204, 1, 6, 0, 0,
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

// CustomLanguageLexerInit initializes any static state used to implement CustomLanguageLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCustomLanguageLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CustomLanguageLexerInit() {
	staticData := &CustomLanguageLexerLexerStaticData
	staticData.once.Do(customlanguagelexerLexerInit)
}

// NewCustomLanguageLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCustomLanguageLexer(input antlr.CharStream) *CustomLanguageLexer {
	CustomLanguageLexerInit()
	l := new(CustomLanguageLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CustomLanguageLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "CustomLanguage.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CustomLanguageLexer tokens.
const (
	CustomLanguageLexerFUNC           = 1
	CustomLanguageLexerINT            = 2
	CustomLanguageLexerFLOAT          = 3
	CustomLanguageLexerSTRING         = 4
	CustomLanguageLexerBOOL           = 5
	CustomLanguageLexerRETURN         = 6
	CustomLanguageLexerVAR            = 7
	CustomLanguageLexerBOOL_LITERAL   = 8
	CustomLanguageLexerARROW          = 9
	CustomLanguageLexerASSIGN         = 10
	CustomLanguageLexerLPAREN         = 11
	CustomLanguageLexerRPAREN         = 12
	CustomLanguageLexerLBRACE         = 13
	CustomLanguageLexerRBRACE         = 14
	CustomLanguageLexerCOMMA          = 15
	CustomLanguageLexerSEMI           = 16
	CustomLanguageLexerCOLON          = 17
	CustomLanguageLexerPLUS           = 18
	CustomLanguageLexerMINUS          = 19
	CustomLanguageLexerMULT           = 20
	CustomLanguageLexerDIV            = 21
	CustomLanguageLexerSTRING_LITERAL = 22
	CustomLanguageLexerFLOAT_LITERAL  = 23
	CustomLanguageLexerINT_LITERAL    = 24
	CustomLanguageLexerID             = 25
	CustomLanguageLexerWS             = 26
	CustomLanguageLexerLINE_COMMENT   = 27
	CustomLanguageLexerBLOCK_COMMENT  = 28
)
