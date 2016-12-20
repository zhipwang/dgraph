// Generated from GraphQLPM.g4 by ANTLR 4.6.

package parser // GraphQLPM

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 3, 17, 128, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 3, 2,
	3, 3, 3, 3, 5, 3, 39, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 56, 10, 6, 12, 6,
	14, 6, 59, 11, 6, 3, 6, 5, 6, 62, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7,
	68, 10, 7, 12, 7, 14, 7, 71, 11, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 13, 5, 13, 91, 10, 13, 3, 13, 7, 13, 94, 10, 13, 12, 13, 14, 13,
	97, 11, 13, 3, 13, 3, 13, 3, 14, 3, 14, 5, 14, 103, 10, 14, 3, 14, 5, 14,
	106, 10, 14, 3, 14, 5, 14, 109, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 7,
	15, 115, 10, 15, 12, 15, 14, 15, 118, 11, 15, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 2, 2, 18, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 24, 26, 28, 30, 32, 2, 4, 3, 2, 6, 7, 3, 2, 8, 9, 121, 2, 34,
	3, 2, 2, 2, 4, 38, 3, 2, 2, 2, 6, 40, 3, 2, 2, 2, 8, 44, 3, 2, 2, 2, 10,
	61, 3, 2, 2, 2, 12, 63, 3, 2, 2, 2, 14, 72, 3, 2, 2, 2, 16, 74, 3, 2, 2,
	2, 18, 76, 3, 2, 2, 2, 20, 83, 3, 2, 2, 2, 22, 85, 3, 2, 2, 2, 24, 87,
	3, 2, 2, 2, 26, 100, 3, 2, 2, 2, 28, 110, 3, 2, 2, 2, 30, 121, 3, 2, 2,
	2, 32, 125, 3, 2, 2, 2, 34, 35, 5, 4, 3, 2, 35, 3, 3, 2, 2, 2, 36, 39,
	5, 24, 13, 2, 37, 39, 5, 6, 4, 2, 38, 36, 3, 2, 2, 2, 38, 37, 3, 2, 2,
	2, 39, 5, 3, 2, 2, 2, 40, 41, 5, 22, 12, 2, 41, 42, 7, 15, 2, 2, 42, 43,
	5, 24, 13, 2, 43, 7, 3, 2, 2, 2, 44, 45, 7, 3, 2, 2, 45, 46, 7, 4, 2, 2,
	46, 47, 5, 10, 6, 2, 47, 48, 7, 5, 2, 2, 48, 9, 3, 2, 2, 2, 49, 50, 7,
	4, 2, 2, 50, 51, 5, 10, 6, 2, 51, 57, 7, 5, 2, 2, 52, 53, 5, 14, 8, 2,
	53, 54, 5, 10, 6, 2, 54, 56, 3, 2, 2, 2, 55, 52, 3, 2, 2, 2, 56, 59, 3,
	2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 62, 3, 2, 2, 2, 59,
	57, 3, 2, 2, 2, 60, 62, 5, 12, 7, 2, 61, 49, 3, 2, 2, 2, 61, 60, 3, 2,
	2, 2, 62, 11, 3, 2, 2, 2, 63, 69, 5, 18, 10, 2, 64, 65, 5, 14, 8, 2, 65,
	66, 5, 18, 10, 2, 66, 68, 3, 2, 2, 2, 67, 64, 3, 2, 2, 2, 68, 71, 3, 2,
	2, 2, 69, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 13, 3, 2, 2, 2, 71, 69,
	3, 2, 2, 2, 72, 73, 9, 2, 2, 2, 73, 15, 3, 2, 2, 2, 74, 75, 9, 3, 2, 2,
	75, 17, 3, 2, 2, 2, 76, 77, 5, 16, 9, 2, 77, 78, 7, 4, 2, 2, 78, 79, 5,
	20, 11, 2, 79, 80, 7, 10, 2, 2, 80, 81, 5, 32, 17, 2, 81, 82, 7, 5, 2,
	2, 82, 19, 3, 2, 2, 2, 83, 84, 7, 16, 2, 2, 84, 21, 3, 2, 2, 2, 85, 86,
	7, 11, 2, 2, 86, 23, 3, 2, 2, 2, 87, 88, 7, 12, 2, 2, 88, 95, 5, 26, 14,
	2, 89, 91, 7, 10, 2, 2, 90, 89, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 92,
	3, 2, 2, 2, 92, 94, 5, 26, 14, 2, 93, 90, 3, 2, 2, 2, 94, 97, 3, 2, 2,
	2, 95, 93, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 98, 3, 2, 2, 2, 97, 95,
	3, 2, 2, 2, 98, 99, 7, 13, 2, 2, 99, 25, 3, 2, 2, 2, 100, 102, 7, 15, 2,
	2, 101, 103, 5, 28, 15, 2, 102, 101, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2,
	103, 105, 3, 2, 2, 2, 104, 106, 5, 8, 5, 2, 105, 104, 3, 2, 2, 2, 105,
	106, 3, 2, 2, 2, 106, 108, 3, 2, 2, 2, 107, 109, 5, 24, 13, 2, 108, 107,
	3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 27, 3, 2, 2, 2, 110, 111, 7, 4,
	2, 2, 111, 116, 5, 30, 16, 2, 112, 113, 7, 10, 2, 2, 113, 115, 5, 30, 16,
	2, 114, 112, 3, 2, 2, 2, 115, 118, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 116,
	117, 3, 2, 2, 2, 117, 119, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 119, 120,
	7, 5, 2, 2, 120, 29, 3, 2, 2, 2, 121, 122, 7, 15, 2, 2, 122, 123, 7, 14,
	2, 2, 123, 124, 5, 32, 17, 2, 124, 31, 3, 2, 2, 2, 125, 126, 7, 16, 2,
	2, 126, 33, 3, 2, 2, 2, 12, 38, 57, 61, 69, 90, 95, 102, 105, 108, 116,
}

var deserializer = antlr.NewATNDeserializer(nil)

var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'@filter'", "'('", "')'", "'||'", "'&&'", "'anyof'", "'allof'", "','",
	"'query'", "'{'", "'}'", "':'",
}

var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "NAME", "STRING", "WS",
}

var ruleNames = []string{
	"document", "definition", "operationDefinition", "filters", "pairsNested",
	"pairs", "filterOperation", "funcName", "pair", "fieldName", "operationType",
	"selectionSet", "field", "arguments", "argument", "value",
}

type GraphQLPMParser struct {
	*antlr.BaseParser
}

func NewGraphQLPMParser(input antlr.TokenStream) *GraphQLPMParser {
	var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	var sharedContextCache = antlr.NewPredictionContextCache()

	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}

	this := new(GraphQLPMParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, sharedContextCache)
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "GraphQLPM.g4"

	return this
}

// GraphQLPMParser tokens.
const (
	GraphQLPMParserEOF    = antlr.TokenEOF
	GraphQLPMParserT__0   = 1
	GraphQLPMParserT__1   = 2
	GraphQLPMParserT__2   = 3
	GraphQLPMParserT__3   = 4
	GraphQLPMParserT__4   = 5
	GraphQLPMParserT__5   = 6
	GraphQLPMParserT__6   = 7
	GraphQLPMParserT__7   = 8
	GraphQLPMParserT__8   = 9
	GraphQLPMParserT__9   = 10
	GraphQLPMParserT__10  = 11
	GraphQLPMParserT__11  = 12
	GraphQLPMParserNAME   = 13
	GraphQLPMParserSTRING = 14
	GraphQLPMParserWS     = 15
)

// GraphQLPMParser rules.
const (
	GraphQLPMParserRULE_document            = 0
	GraphQLPMParserRULE_definition          = 1
	GraphQLPMParserRULE_operationDefinition = 2
	GraphQLPMParserRULE_filters             = 3
	GraphQLPMParserRULE_pairsNested         = 4
	GraphQLPMParserRULE_pairs               = 5
	GraphQLPMParserRULE_filterOperation     = 6
	GraphQLPMParserRULE_funcName            = 7
	GraphQLPMParserRULE_pair                = 8
	GraphQLPMParserRULE_fieldName           = 9
	GraphQLPMParserRULE_operationType       = 10
	GraphQLPMParserRULE_selectionSet        = 11
	GraphQLPMParserRULE_field               = 12
	GraphQLPMParserRULE_arguments           = 13
	GraphQLPMParserRULE_argument            = 14
	GraphQLPMParserRULE_value               = 15
)

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLPMParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLPMParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) Definition() IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *GraphQLPMParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GraphQLPMParserRULE_document)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Definition()
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLPMParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLPMParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *DefinitionContext) OperationDefinition() IOperationDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationDefinitionContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *GraphQLPMParser) Definition() (localctx IDefinitionContext) {
	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GraphQLPMParserRULE_definition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(36)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLPMParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.SelectionSet()
		}

	case GraphQLPMParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.OperationDefinition()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOperationDefinitionContext is an interface to support dynamic dispatch.
type IOperationDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationDefinitionContext differentiates from other interfaces.
	IsOperationDefinitionContext()
}

type OperationDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationDefinitionContext() *OperationDefinitionContext {
	var p = new(OperationDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLPMParserRULE_operationDefinition
	return p
}

func (*OperationDefinitionContext) IsOperationDefinitionContext() {}

func NewOperationDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationDefinitionContext {
	var p = new(OperationDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLPMParserRULE_operationDefinition

	return p
}

func (s *OperationDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationDefinitionContext) OperationType() IOperationTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationTypeContext)
}

func (s *OperationDefinitionContext) NAME() antlr.TerminalNode {
	return s.GetToken(GraphQLPMParserNAME, 0)
}

func (s *OperationDefinitionContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *OperationDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.EnterOperationDefinition(s)
	}
}

func (s *OperationDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.ExitOperationDefinition(s)
	}
}

func (p *GraphQLPMParser) OperationDefinition() (localctx IOperationDefinitionContext) {
	localctx = NewOperationDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GraphQLPMParserRULE_operationDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.OperationType()
	}
	{
		p.SetState(39)
		p.Match(GraphQLPMParserNAME)
	}
	{
		p.SetState(40)
		p.SelectionSet()
	}

	return localctx
}

// IFiltersContext is an interface to support dynamic dispatch.
type IFiltersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFiltersContext differentiates from other interfaces.
	IsFiltersContext()
}

type FiltersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFiltersContext() *FiltersContext {
	var p = new(FiltersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLPMParserRULE_filters
	return p
}

func (*FiltersContext) IsFiltersContext() {}

func NewFiltersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FiltersContext {
	var p = new(FiltersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLPMParserRULE_filters

	return p
}

func (s *FiltersContext) GetParser() antlr.Parser { return s.parser }

func (s *FiltersContext) PairsNested() IPairsNestedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairsNestedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPairsNestedContext)
}

func (s *FiltersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FiltersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FiltersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.EnterFilters(s)
	}
}

func (s *FiltersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.ExitFilters(s)
	}
}

func (p *GraphQLPMParser) Filters() (localctx IFiltersContext) {
	localctx = NewFiltersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GraphQLPMParserRULE_filters)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Match(GraphQLPMParserT__0)
	}
	{
		p.SetState(43)
		p.Match(GraphQLPMParserT__1)
	}
	{
		p.SetState(44)
		p.PairsNested()
	}
	{
		p.SetState(45)
		p.Match(GraphQLPMParserT__2)
	}

	return localctx
}

// IPairsNestedContext is an interface to support dynamic dispatch.
type IPairsNestedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPairsNestedContext differentiates from other interfaces.
	IsPairsNestedContext()
}

type PairsNestedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairsNestedContext() *PairsNestedContext {
	var p = new(PairsNestedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLPMParserRULE_pairsNested
	return p
}

func (*PairsNestedContext) IsPairsNestedContext() {}

func NewPairsNestedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairsNestedContext {
	var p = new(PairsNestedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLPMParserRULE_pairsNested

	return p
}

func (s *PairsNestedContext) GetParser() antlr.Parser { return s.parser }

func (s *PairsNestedContext) AllPairsNested() []IPairsNestedContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPairsNestedContext)(nil)).Elem())
	var tst = make([]IPairsNestedContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPairsNestedContext)
		}
	}

	return tst
}

func (s *PairsNestedContext) PairsNested(i int) IPairsNestedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairsNestedContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPairsNestedContext)
}

func (s *PairsNestedContext) AllFilterOperation() []IFilterOperationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFilterOperationContext)(nil)).Elem())
	var tst = make([]IFilterOperationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFilterOperationContext)
		}
	}

	return tst
}

func (s *PairsNestedContext) FilterOperation(i int) IFilterOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterOperationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFilterOperationContext)
}

func (s *PairsNestedContext) Pairs() IPairsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPairsContext)
}

func (s *PairsNestedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairsNestedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairsNestedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.EnterPairsNested(s)
	}
}

func (s *PairsNestedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.ExitPairsNested(s)
	}
}

func (p *GraphQLPMParser) PairsNested() (localctx IPairsNestedContext) {
	localctx = NewPairsNestedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GraphQLPMParserRULE_pairsNested)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(59)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLPMParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)
			p.Match(GraphQLPMParserT__1)
		}
		{
			p.SetState(48)
			p.PairsNested()
		}
		{
			p.SetState(49)
			p.Match(GraphQLPMParserT__2)
		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(50)
					p.FilterOperation()
				}
				{
					p.SetState(51)
					p.PairsNested()
				}

			}
			p.SetState(57)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
		}

	case GraphQLPMParserT__5, GraphQLPMParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(58)
			p.Pairs()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPairsContext is an interface to support dynamic dispatch.
type IPairsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPairsContext differentiates from other interfaces.
	IsPairsContext()
}

type PairsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairsContext() *PairsContext {
	var p = new(PairsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLPMParserRULE_pairs
	return p
}

func (*PairsContext) IsPairsContext() {}

func NewPairsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairsContext {
	var p = new(PairsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLPMParserRULE_pairs

	return p
}

func (s *PairsContext) GetParser() antlr.Parser { return s.parser }

func (s *PairsContext) AllPair() []IPairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPairContext)(nil)).Elem())
	var tst = make([]IPairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPairContext)
		}
	}

	return tst
}

func (s *PairsContext) Pair(i int) IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *PairsContext) AllFilterOperation() []IFilterOperationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFilterOperationContext)(nil)).Elem())
	var tst = make([]IFilterOperationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFilterOperationContext)
		}
	}

	return tst
}

func (s *PairsContext) FilterOperation(i int) IFilterOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterOperationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFilterOperationContext)
}

func (s *PairsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.EnterPairs(s)
	}
}

func (s *PairsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.ExitPairs(s)
	}
}

func (p *GraphQLPMParser) Pairs() (localctx IPairsContext) {
	localctx = NewPairsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GraphQLPMParserRULE_pairs)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Pair()
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(62)
				p.FilterOperation()
			}
			{
				p.SetState(63)
				p.Pair()
			}

		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IFilterOperationContext is an interface to support dynamic dispatch.
type IFilterOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterOperationContext differentiates from other interfaces.
	IsFilterOperationContext()
}

type FilterOperationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterOperationContext() *FilterOperationContext {
	var p = new(FilterOperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLPMParserRULE_filterOperation
	return p
}

func (*FilterOperationContext) IsFilterOperationContext() {}

func NewFilterOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterOperationContext {
	var p = new(FilterOperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLPMParserRULE_filterOperation

	return p
}

func (s *FilterOperationContext) GetParser() antlr.Parser { return s.parser }
func (s *FilterOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.EnterFilterOperation(s)
	}
}

func (s *FilterOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.ExitFilterOperation(s)
	}
}

func (p *GraphQLPMParser) FilterOperation() (localctx IFilterOperationContext) {
	localctx = NewFilterOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GraphQLPMParserRULE_filterOperation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(70)
	_la = p.GetTokenStream().LA(1)

	if !(_la == GraphQLPMParserT__3 || _la == GraphQLPMParserT__4) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IFuncNameContext is an interface to support dynamic dispatch.
type IFuncNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncNameContext differentiates from other interfaces.
	IsFuncNameContext()
}

type FuncNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncNameContext() *FuncNameContext {
	var p = new(FuncNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLPMParserRULE_funcName
	return p
}

func (*FuncNameContext) IsFuncNameContext() {}

func NewFuncNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncNameContext {
	var p = new(FuncNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLPMParserRULE_funcName

	return p
}

func (s *FuncNameContext) GetParser() antlr.Parser { return s.parser }
func (s *FuncNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.EnterFuncName(s)
	}
}

func (s *FuncNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.ExitFuncName(s)
	}
}

func (p *GraphQLPMParser) FuncName() (localctx IFuncNameContext) {
	localctx = NewFuncNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GraphQLPMParserRULE_funcName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(72)
	_la = p.GetTokenStream().LA(1)

	if !(_la == GraphQLPMParserT__5 || _la == GraphQLPMParserT__6) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLPMParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLPMParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) FuncName() IFuncNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncNameContext)
}

func (s *PairContext) FieldName() IFieldNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *PairContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.ExitPair(s)
	}
}

func (p *GraphQLPMParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GraphQLPMParserRULE_pair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.FuncName()
	}
	{
		p.SetState(75)
		p.Match(GraphQLPMParserT__1)
	}
	{
		p.SetState(76)
		p.FieldName()
	}
	{
		p.SetState(77)
		p.Match(GraphQLPMParserT__7)
	}
	{
		p.SetState(78)
		p.Value()
	}
	{
		p.SetState(79)
		p.Match(GraphQLPMParserT__2)
	}

	return localctx
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLPMParserRULE_fieldName
	return p
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLPMParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) CopyFrom(ctx *FieldNameContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FieldNameValueContext struct {
	*FieldNameContext
}

func NewFieldNameValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldNameValueContext {
	var p = new(FieldNameValueContext)

	p.FieldNameContext = NewEmptyFieldNameContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldNameContext))

	return p
}

func (s *FieldNameValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(GraphQLPMParserSTRING, 0)
}

func (s *FieldNameValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.EnterFieldNameValue(s)
	}
}

func (s *FieldNameValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.ExitFieldNameValue(s)
	}
}

func (p *GraphQLPMParser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GraphQLPMParserRULE_fieldName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewFieldNameValueContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(GraphQLPMParserSTRING)
	}

	return localctx
}

// IOperationTypeContext is an interface to support dynamic dispatch.
type IOperationTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationTypeContext differentiates from other interfaces.
	IsOperationTypeContext()
}

type OperationTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationTypeContext() *OperationTypeContext {
	var p = new(OperationTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLPMParserRULE_operationType
	return p
}

func (*OperationTypeContext) IsOperationTypeContext() {}

func NewOperationTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationTypeContext {
	var p = new(OperationTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLPMParserRULE_operationType

	return p
}

func (s *OperationTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *OperationTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.EnterOperationType(s)
	}
}

func (s *OperationTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.ExitOperationType(s)
	}
}

func (p *GraphQLPMParser) OperationType() (localctx IOperationTypeContext) {
	localctx = NewOperationTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GraphQLPMParserRULE_operationType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Match(GraphQLPMParserT__8)
	}

	return localctx
}

// ISelectionSetContext is an interface to support dynamic dispatch.
type ISelectionSetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectionSetContext differentiates from other interfaces.
	IsSelectionSetContext()
}

type SelectionSetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionSetContext() *SelectionSetContext {
	var p = new(SelectionSetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLPMParserRULE_selectionSet
	return p
}

func (*SelectionSetContext) IsSelectionSetContext() {}

func NewSelectionSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionSetContext {
	var p = new(SelectionSetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLPMParserRULE_selectionSet

	return p
}

func (s *SelectionSetContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionSetContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *SelectionSetContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *SelectionSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionSetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionSetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.EnterSelectionSet(s)
	}
}

func (s *SelectionSetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.ExitSelectionSet(s)
	}
}

func (p *GraphQLPMParser) SelectionSet() (localctx ISelectionSetContext) {
	localctx = NewSelectionSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GraphQLPMParserRULE_selectionSet)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(GraphQLPMParserT__9)
	}
	{
		p.SetState(86)
		p.Field()
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GraphQLPMParserT__7 || _la == GraphQLPMParserNAME {
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLPMParserT__7 {
			{
				p.SetState(87)
				p.Match(GraphQLPMParserT__7)
			}

		}
		{
			p.SetState(90)
			p.Field()
		}

		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(96)
		p.Match(GraphQLPMParserT__10)
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLPMParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLPMParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) NAME() antlr.TerminalNode {
	return s.GetToken(GraphQLPMParserNAME, 0)
}

func (s *FieldContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FieldContext) Filters() IFiltersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFiltersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFiltersContext)
}

func (s *FieldContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *GraphQLPMParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GraphQLPMParserRULE_field)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(GraphQLPMParserNAME)
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLPMParserT__1 {
		{
			p.SetState(99)
			p.Arguments()
		}

	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLPMParserT__0 {
		{
			p.SetState(102)
			p.Filters()
		}

	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLPMParserT__9 {
		{
			p.SetState(105)
			p.SelectionSet()
		}

	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLPMParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLPMParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *ArgumentsContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *GraphQLPMParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GraphQLPMParserRULE_arguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(GraphQLPMParserT__1)
	}
	{
		p.SetState(109)
		p.Argument()
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GraphQLPMParserT__7 {
		{
			p.SetState(110)
			p.Match(GraphQLPMParserT__7)
		}
		{
			p.SetState(111)
			p.Argument()
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(117)
		p.Match(GraphQLPMParserT__2)
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLPMParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLPMParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) NAME() antlr.TerminalNode {
	return s.GetToken(GraphQLPMParserNAME, 0)
}

func (s *ArgumentContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *GraphQLPMParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GraphQLPMParserRULE_argument)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(GraphQLPMParserNAME)
	}
	{
		p.SetState(120)
		p.Match(GraphQLPMParserT__11)
	}
	{
		p.SetState(121)
		p.Value()
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLPMParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLPMParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyFrom(ctx *ValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringValueContext struct {
	*ValueContext
}

func NewStringValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringValueContext {
	var p = new(StringValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(GraphQLPMParserSTRING, 0)
}

func (s *StringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.EnterStringValue(s)
	}
}

func (s *StringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.ExitStringValue(s)
	}
}

func (p *GraphQLPMParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GraphQLPMParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewStringValueContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(GraphQLPMParserSTRING)
	}

	return localctx
}
