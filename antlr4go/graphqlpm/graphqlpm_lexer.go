// Generated from GraphQLPM.g4 by ANTLR 4.6.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 2, 17, 101, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 7, 14, 81,
	10, 14, 12, 14, 14, 14, 84, 11, 14, 3, 15, 3, 15, 7, 15, 88, 10, 15, 12,
	15, 14, 15, 91, 11, 15, 3, 15, 3, 15, 3, 16, 6, 16, 96, 10, 16, 13, 16,
	14, 16, 97, 3, 16, 3, 16, 2, 2, 17, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	3, 2, 6, 5, 2, 67, 92, 97, 97, 99, 124, 7, 2, 48, 48, 50, 59, 67, 92, 97,
	97, 99, 124, 7, 2, 34, 34, 48, 48, 50, 59, 67, 92, 99, 124, 5, 2, 11, 12,
	15, 15, 34, 34, 103, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2,
	2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2,
	2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3,
	2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31,
	3, 2, 2, 2, 3, 33, 3, 2, 2, 2, 5, 42, 3, 2, 2, 2, 7, 44, 3, 2, 2, 2, 9,
	47, 3, 2, 2, 2, 11, 50, 3, 2, 2, 2, 13, 56, 3, 2, 2, 2, 15, 62, 3, 2, 2,
	2, 17, 64, 3, 2, 2, 2, 19, 66, 3, 2, 2, 2, 21, 72, 3, 2, 2, 2, 23, 74,
	3, 2, 2, 2, 25, 76, 3, 2, 2, 2, 27, 78, 3, 2, 2, 2, 29, 85, 3, 2, 2, 2,
	31, 95, 3, 2, 2, 2, 33, 34, 7, 66, 2, 2, 34, 35, 7, 104, 2, 2, 35, 36,
	7, 107, 2, 2, 36, 37, 7, 110, 2, 2, 37, 38, 7, 118, 2, 2, 38, 39, 7, 103,
	2, 2, 39, 40, 7, 116, 2, 2, 40, 41, 7, 42, 2, 2, 41, 4, 3, 2, 2, 2, 42,
	43, 7, 43, 2, 2, 43, 6, 3, 2, 2, 2, 44, 45, 7, 126, 2, 2, 45, 46, 7, 126,
	2, 2, 46, 8, 3, 2, 2, 2, 47, 48, 7, 40, 2, 2, 48, 49, 7, 40, 2, 2, 49,
	10, 3, 2, 2, 2, 50, 51, 7, 99, 2, 2, 51, 52, 7, 112, 2, 2, 52, 53, 7, 123,
	2, 2, 53, 54, 7, 113, 2, 2, 54, 55, 7, 104, 2, 2, 55, 12, 3, 2, 2, 2, 56,
	57, 7, 99, 2, 2, 57, 58, 7, 110, 2, 2, 58, 59, 7, 110, 2, 2, 59, 60, 7,
	113, 2, 2, 60, 61, 7, 104, 2, 2, 61, 14, 3, 2, 2, 2, 62, 63, 7, 42, 2,
	2, 63, 16, 3, 2, 2, 2, 64, 65, 7, 46, 2, 2, 65, 18, 3, 2, 2, 2, 66, 67,
	7, 115, 2, 2, 67, 68, 7, 119, 2, 2, 68, 69, 7, 103, 2, 2, 69, 70, 7, 116,
	2, 2, 70, 71, 7, 123, 2, 2, 71, 20, 3, 2, 2, 2, 72, 73, 7, 125, 2, 2, 73,
	22, 3, 2, 2, 2, 74, 75, 7, 127, 2, 2, 75, 24, 3, 2, 2, 2, 76, 77, 7, 60,
	2, 2, 77, 26, 3, 2, 2, 2, 78, 82, 9, 2, 2, 2, 79, 81, 9, 3, 2, 2, 80, 79,
	3, 2, 2, 2, 81, 84, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2,
	83, 28, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 85, 89, 7, 36, 2, 2, 86, 88, 9,
	4, 2, 2, 87, 86, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89,
	90, 3, 2, 2, 2, 90, 92, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 93, 7, 36,
	2, 2, 93, 30, 3, 2, 2, 2, 94, 96, 9, 5, 2, 2, 95, 94, 3, 2, 2, 2, 96, 97,
	3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2,
	99, 100, 8, 16, 2, 2, 100, 32, 3, 2, 2, 2, 6, 2, 82, 89, 97, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'@filter('", "')'", "'||'", "'&&'", "'anyof'", "'allof'", "'('", "','",
	"'query'", "'{'", "'}'", "':'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "NAME", "STRING", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "NAME", "STRING", "WS",
}

type GraphQLPMLexer struct {
	*antlr.BaseLexer
	modeNames []string
	// TODO: EOF string
}

func NewGraphQLPMLexer(input antlr.CharStream) *GraphQLPMLexer {
	var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}

	l := new(GraphQLPMLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "GraphQLPM.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GraphQLPMLexer tokens.
const (
	GraphQLPMLexerT__0   = 1
	GraphQLPMLexerT__1   = 2
	GraphQLPMLexerT__2   = 3
	GraphQLPMLexerT__3   = 4
	GraphQLPMLexerT__4   = 5
	GraphQLPMLexerT__5   = 6
	GraphQLPMLexerT__6   = 7
	GraphQLPMLexerT__7   = 8
	GraphQLPMLexerT__8   = 9
	GraphQLPMLexerT__9   = 10
	GraphQLPMLexerT__10  = 11
	GraphQLPMLexerT__11  = 12
	GraphQLPMLexerNAME   = 13
	GraphQLPMLexerSTRING = 14
	GraphQLPMLexerWS     = 15
)
