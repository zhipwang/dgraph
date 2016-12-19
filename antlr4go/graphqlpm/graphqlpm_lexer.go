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
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 2, 12, 64, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 7, 9, 44, 10, 9, 12, 9, 14, 9, 47, 11, 9,
	3, 10, 3, 10, 7, 10, 51, 10, 10, 12, 10, 14, 10, 54, 11, 10, 3, 10, 3,
	10, 3, 11, 6, 11, 59, 10, 11, 13, 11, 14, 11, 60, 3, 11, 3, 11, 2, 2, 12,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 3,
	2, 6, 5, 2, 67, 92, 97, 97, 99, 124, 7, 2, 48, 48, 50, 59, 67, 92, 97,
	97, 99, 124, 6, 2, 48, 48, 50, 59, 67, 92, 99, 124, 5, 2, 11, 12, 15, 15,
	34, 34, 66, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2,
	17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 3, 23, 3, 2, 2, 2,
	5, 29, 3, 2, 2, 2, 7, 31, 3, 2, 2, 2, 9, 33, 3, 2, 2, 2, 11, 35, 3, 2,
	2, 2, 13, 37, 3, 2, 2, 2, 15, 39, 3, 2, 2, 2, 17, 41, 3, 2, 2, 2, 19, 48,
	3, 2, 2, 2, 21, 58, 3, 2, 2, 2, 23, 24, 7, 115, 2, 2, 24, 25, 7, 119, 2,
	2, 25, 26, 7, 103, 2, 2, 26, 27, 7, 116, 2, 2, 27, 28, 7, 123, 2, 2, 28,
	4, 3, 2, 2, 2, 29, 30, 7, 125, 2, 2, 30, 6, 3, 2, 2, 2, 31, 32, 7, 46,
	2, 2, 32, 8, 3, 2, 2, 2, 33, 34, 7, 127, 2, 2, 34, 10, 3, 2, 2, 2, 35,
	36, 7, 42, 2, 2, 36, 12, 3, 2, 2, 2, 37, 38, 7, 43, 2, 2, 38, 14, 3, 2,
	2, 2, 39, 40, 7, 60, 2, 2, 40, 16, 3, 2, 2, 2, 41, 45, 9, 2, 2, 2, 42,
	44, 9, 3, 2, 2, 43, 42, 3, 2, 2, 2, 44, 47, 3, 2, 2, 2, 45, 43, 3, 2, 2,
	2, 45, 46, 3, 2, 2, 2, 46, 18, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 48, 52,
	7, 36, 2, 2, 49, 51, 9, 4, 2, 2, 50, 49, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2,
	52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 55, 3, 2, 2, 2, 54, 52, 3,
	2, 2, 2, 55, 56, 7, 36, 2, 2, 56, 20, 3, 2, 2, 2, 57, 59, 9, 5, 2, 2, 58,
	57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2,
	2, 61, 62, 3, 2, 2, 2, 62, 63, 8, 11, 2, 2, 63, 22, 3, 2, 2, 2, 6, 2, 45,
	52, 60, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'query'", "'{'", "','", "'}'", "'('", "')'", "':'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "NAME", "STRING", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "NAME", "STRING",
	"WS",
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
	GraphQLPMLexerNAME   = 8
	GraphQLPMLexerSTRING = 9
	GraphQLPMLexerWS     = 10
)
