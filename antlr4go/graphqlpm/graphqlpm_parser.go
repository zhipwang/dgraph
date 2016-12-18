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
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 3, 12, 89, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 3, 2, 6, 2, 28, 10, 2, 13, 2, 14, 2, 29, 3, 3, 3, 3, 5, 3, 34, 10,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 5, 6, 45, 10,
	6, 3, 6, 7, 6, 48, 10, 6, 12, 6, 14, 6, 51, 11, 6, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 8, 3, 8, 5, 8, 59, 10, 8, 3, 8, 5, 8, 62, 10, 8, 3, 9, 3, 9, 5, 9,
	66, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11,
	76, 10, 11, 12, 11, 14, 11, 79, 11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 13, 3, 13, 3, 13, 2, 2, 14, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 2, 2, 84, 2, 27, 3, 2, 2, 2, 4, 33, 3, 2, 2, 2, 6, 35, 3, 2,
	2, 2, 8, 39, 3, 2, 2, 2, 10, 41, 3, 2, 2, 2, 12, 54, 3, 2, 2, 2, 14, 56,
	3, 2, 2, 2, 16, 65, 3, 2, 2, 2, 18, 67, 3, 2, 2, 2, 20, 71, 3, 2, 2, 2,
	22, 82, 3, 2, 2, 2, 24, 86, 3, 2, 2, 2, 26, 28, 5, 4, 3, 2, 27, 26, 3,
	2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30,
	3, 3, 2, 2, 2, 31, 34, 5, 10, 6, 2, 32, 34, 5, 6, 4, 2, 33, 31, 3, 2, 2,
	2, 33, 32, 3, 2, 2, 2, 34, 5, 3, 2, 2, 2, 35, 36, 5, 8, 5, 2, 36, 37, 7,
	10, 2, 2, 37, 38, 5, 10, 6, 2, 38, 7, 3, 2, 2, 2, 39, 40, 7, 3, 2, 2, 40,
	9, 3, 2, 2, 2, 41, 42, 7, 4, 2, 2, 42, 49, 5, 12, 7, 2, 43, 45, 7, 5, 2,
	2, 44, 43, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 48,
	5, 12, 7, 2, 47, 44, 3, 2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2,
	49, 50, 3, 2, 2, 2, 50, 52, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 52, 53, 7,
	6, 2, 2, 53, 11, 3, 2, 2, 2, 54, 55, 5, 14, 8, 2, 55, 13, 3, 2, 2, 2, 56,
	58, 5, 16, 9, 2, 57, 59, 5, 20, 11, 2, 58, 57, 3, 2, 2, 2, 58, 59, 3, 2,
	2, 2, 59, 61, 3, 2, 2, 2, 60, 62, 5, 10, 6, 2, 61, 60, 3, 2, 2, 2, 61,
	62, 3, 2, 2, 2, 62, 15, 3, 2, 2, 2, 63, 66, 5, 18, 10, 2, 64, 66, 7, 10,
	2, 2, 65, 63, 3, 2, 2, 2, 65, 64, 3, 2, 2, 2, 66, 17, 3, 2, 2, 2, 67, 68,
	7, 10, 2, 2, 68, 69, 7, 7, 2, 2, 69, 70, 7, 10, 2, 2, 70, 19, 3, 2, 2,
	2, 71, 72, 7, 8, 2, 2, 72, 77, 5, 22, 12, 2, 73, 74, 7, 5, 2, 2, 74, 76,
	5, 22, 12, 2, 75, 73, 3, 2, 2, 2, 76, 79, 3, 2, 2, 2, 77, 75, 3, 2, 2,
	2, 77, 78, 3, 2, 2, 2, 78, 80, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 80, 81,
	7, 9, 2, 2, 81, 21, 3, 2, 2, 2, 82, 83, 7, 10, 2, 2, 83, 84, 7, 7, 2, 2,
	84, 85, 5, 24, 13, 2, 85, 23, 3, 2, 2, 2, 86, 87, 7, 11, 2, 2, 87, 25,
	3, 2, 2, 2, 10, 29, 33, 44, 49, 58, 61, 65, 77,
}

var deserializer = antlr.NewATNDeserializer(nil)

var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'query'", "'{'", "','", "'}'", "':'", "'('", "')'",
}

var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "NAME", "STRING", "WS",
}

var ruleNames = []string{
	"document", "definition", "operationDefinition", "operationType", "selectionSet",
	"selection", "field", "fieldName", "alias", "arguments", "argument", "value",
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
	GraphQLPMParserNAME   = 8
	GraphQLPMParserSTRING = 9
	GraphQLPMParserWS     = 10
)

// GraphQLPMParser rules.
const (
	GraphQLPMParserRULE_document            = 0
	GraphQLPMParserRULE_definition          = 1
	GraphQLPMParserRULE_operationDefinition = 2
	GraphQLPMParserRULE_operationType       = 3
	GraphQLPMParserRULE_selectionSet        = 4
	GraphQLPMParserRULE_selection           = 5
	GraphQLPMParserRULE_field               = 6
	GraphQLPMParserRULE_fieldName           = 7
	GraphQLPMParserRULE_alias               = 8
	GraphQLPMParserRULE_arguments           = 9
	GraphQLPMParserRULE_argument            = 10
	GraphQLPMParserRULE_value               = 11
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

func (s *DocumentContext) AllDefinition() []IDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefinitionContext)(nil)).Elem())
	var tst = make([]IDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefinitionContext)
		}
	}

	return tst
}

func (s *DocumentContext) Definition(i int) IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), i)

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
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GraphQLPMParserT__0 || _la == GraphQLPMParserT__1 {
		{
			p.SetState(24)
			p.Definition()
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

	p.SetState(31)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLPMParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.SelectionSet()
		}

	case GraphQLPMParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(30)
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
		p.SetState(33)
		p.OperationType()
	}
	{
		p.SetState(34)
		p.Match(GraphQLPMParserNAME)
	}
	{
		p.SetState(35)
		p.SelectionSet()
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
	p.EnterRule(localctx, 6, GraphQLPMParserRULE_operationType)

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
		p.SetState(37)
		p.Match(GraphQLPMParserT__0)
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

func (s *SelectionSetContext) AllSelection() []ISelectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectionContext)(nil)).Elem())
	var tst = make([]ISelectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectionContext)
		}
	}

	return tst
}

func (s *SelectionSetContext) Selection(i int) ISelectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectionContext)
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
	p.EnterRule(localctx, 8, GraphQLPMParserRULE_selectionSet)
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
		p.SetState(39)
		p.Match(GraphQLPMParserT__1)
	}
	{
		p.SetState(40)
		p.Selection()
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GraphQLPMParserT__2 || _la == GraphQLPMParserNAME {
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLPMParserT__2 {
			{
				p.SetState(41)
				p.Match(GraphQLPMParserT__2)
			}

		}
		{
			p.SetState(44)
			p.Selection()
		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(50)
		p.Match(GraphQLPMParserT__3)
	}

	return localctx
}

// ISelectionContext is an interface to support dynamic dispatch.
type ISelectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectionContext differentiates from other interfaces.
	IsSelectionContext()
}

type SelectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionContext() *SelectionContext {
	var p = new(SelectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLPMParserRULE_selection
	return p
}

func (*SelectionContext) IsSelectionContext() {}

func NewSelectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionContext {
	var p = new(SelectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLPMParserRULE_selection

	return p
}

func (s *SelectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *SelectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.EnterSelection(s)
	}
}

func (s *SelectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.ExitSelection(s)
	}
}

func (p *GraphQLPMParser) Selection() (localctx ISelectionContext) {
	localctx = NewSelectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GraphQLPMParserRULE_selection)

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
		p.SetState(52)
		p.Field()
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

func (s *FieldContext) FieldName() IFieldNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
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
	p.EnterRule(localctx, 12, GraphQLPMParserRULE_field)
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
		p.SetState(54)
		p.FieldName()
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLPMParserT__5 {
		{
			p.SetState(55)
			p.Arguments()
		}

	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLPMParserT__1 {
		{
			p.SetState(58)
			p.SelectionSet()
		}

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

func (s *FieldNameContext) Alias() IAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *FieldNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(GraphQLPMParserNAME, 0)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (p *GraphQLPMParser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GraphQLPMParserRULE_fieldName)

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

	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.Alias()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			p.Match(GraphQLPMParserNAME)
		}

	}

	return localctx
}

// IAliasContext is an interface to support dynamic dispatch.
type IAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasContext() *AliasContext {
	var p = new(AliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLPMParserRULE_alias
	return p
}

func (*AliasContext) IsAliasContext() {}

func NewAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasContext {
	var p = new(AliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLPMParserRULE_alias

	return p
}

func (s *AliasContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(GraphQLPMParserNAME)
}

func (s *AliasContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(GraphQLPMParserNAME, i)
}

func (s *AliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.EnterAlias(s)
	}
}

func (s *AliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLPMListener); ok {
		listenerT.ExitAlias(s)
	}
}

func (p *GraphQLPMParser) Alias() (localctx IAliasContext) {
	localctx = NewAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GraphQLPMParserRULE_alias)

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
		p.SetState(65)
		p.Match(GraphQLPMParserNAME)
	}
	{
		p.SetState(66)
		p.Match(GraphQLPMParserT__4)
	}
	{
		p.SetState(67)
		p.Match(GraphQLPMParserNAME)
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
	p.EnterRule(localctx, 18, GraphQLPMParserRULE_arguments)
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
		p.SetState(69)
		p.Match(GraphQLPMParserT__5)
	}
	{
		p.SetState(70)
		p.Argument()
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GraphQLPMParserT__2 {
		{
			p.SetState(71)
			p.Match(GraphQLPMParserT__2)
		}
		{
			p.SetState(72)
			p.Argument()
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(78)
		p.Match(GraphQLPMParserT__6)
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
	p.EnterRule(localctx, 20, GraphQLPMParserRULE_argument)

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
		p.SetState(80)
		p.Match(GraphQLPMParserNAME)
	}
	{
		p.SetState(81)
		p.Match(GraphQLPMParserT__4)
	}
	{
		p.SetState(82)
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
	p.EnterRule(localctx, 22, GraphQLPMParserRULE_value)

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
		p.SetState(84)
		p.Match(GraphQLPMParserSTRING)
	}

	return localctx
}
