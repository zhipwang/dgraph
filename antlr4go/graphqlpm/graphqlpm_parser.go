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
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 3, 12, 70, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 25, 10,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 5, 6, 36, 10,
	6, 3, 6, 7, 6, 39, 10, 6, 12, 6, 14, 6, 42, 11, 6, 3, 6, 3, 6, 3, 7, 3,
	7, 5, 7, 48, 10, 7, 3, 7, 5, 7, 51, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8,
	57, 10, 8, 12, 8, 14, 8, 60, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 2, 2, 11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 2, 66,
	2, 20, 3, 2, 2, 2, 4, 24, 3, 2, 2, 2, 6, 26, 3, 2, 2, 2, 8, 30, 3, 2, 2,
	2, 10, 32, 3, 2, 2, 2, 12, 45, 3, 2, 2, 2, 14, 52, 3, 2, 2, 2, 16, 63,
	3, 2, 2, 2, 18, 67, 3, 2, 2, 2, 20, 21, 5, 4, 3, 2, 21, 3, 3, 2, 2, 2,
	22, 25, 5, 10, 6, 2, 23, 25, 5, 6, 4, 2, 24, 22, 3, 2, 2, 2, 24, 23, 3,
	2, 2, 2, 25, 5, 3, 2, 2, 2, 26, 27, 5, 8, 5, 2, 27, 28, 7, 10, 2, 2, 28,
	29, 5, 10, 6, 2, 29, 7, 3, 2, 2, 2, 30, 31, 7, 3, 2, 2, 31, 9, 3, 2, 2,
	2, 32, 33, 7, 4, 2, 2, 33, 40, 5, 12, 7, 2, 34, 36, 7, 5, 2, 2, 35, 34,
	3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 39, 5, 12, 7, 2,
	38, 35, 3, 2, 2, 2, 39, 42, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 41, 3,
	2, 2, 2, 41, 43, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 43, 44, 7, 6, 2, 2, 44,
	11, 3, 2, 2, 2, 45, 47, 7, 10, 2, 2, 46, 48, 5, 14, 8, 2, 47, 46, 3, 2,
	2, 2, 47, 48, 3, 2, 2, 2, 48, 50, 3, 2, 2, 2, 49, 51, 5, 10, 6, 2, 50,
	49, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 13, 3, 2, 2, 2, 52, 53, 7, 7, 2,
	2, 53, 58, 5, 16, 9, 2, 54, 55, 7, 5, 2, 2, 55, 57, 5, 16, 9, 2, 56, 54,
	3, 2, 2, 2, 57, 60, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2,
	59, 61, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 61, 62, 7, 8, 2, 2, 62, 15, 3,
	2, 2, 2, 63, 64, 7, 10, 2, 2, 64, 65, 7, 9, 2, 2, 65, 66, 5, 18, 10, 2,
	66, 17, 3, 2, 2, 2, 67, 68, 7, 11, 2, 2, 68, 19, 3, 2, 2, 2, 8, 24, 35,
	40, 47, 50, 58,
}

var deserializer = antlr.NewATNDeserializer(nil)

var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'query'", "'{'", "','", "'}'", "'('", "')'", "':'",
}

var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "NAME", "STRING", "WS",
}

var ruleNames = []string{
	"document", "definition", "operationDefinition", "operationType", "selectionSet",
	"field", "arguments", "argument", "value",
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
	GraphQLPMParserRULE_field               = 5
	GraphQLPMParserRULE_arguments           = 6
	GraphQLPMParserRULE_argument            = 7
	GraphQLPMParserRULE_value               = 8
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

func (s *DocumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLPMVisitor:
		return t.VisitDocument(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(18)
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

func (s *DefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLPMVisitor:
		return t.VisitDefinition(s)

	default:
		return t.VisitChildren(s)
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

	p.SetState(22)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLPMParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)
			p.SelectionSet()
		}

	case GraphQLPMParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(21)
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

func (s *OperationDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLPMVisitor:
		return t.VisitOperationDefinition(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(24)
		p.OperationType()
	}
	{
		p.SetState(25)
		p.Match(GraphQLPMParserNAME)
	}
	{
		p.SetState(26)
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

func (s *OperationTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLPMVisitor:
		return t.VisitOperationType(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(28)
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

func (s *SelectionSetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLPMVisitor:
		return t.VisitSelectionSet(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(30)
		p.Match(GraphQLPMParserT__1)
	}
	{
		p.SetState(31)
		p.Field()
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GraphQLPMParserT__2 || _la == GraphQLPMParserNAME {
		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLPMParserT__2 {
			{
				p.SetState(32)
				p.Match(GraphQLPMParserT__2)
			}

		}
		{
			p.SetState(35)
			p.Field()
		}

		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(41)
		p.Match(GraphQLPMParserT__3)
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

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLPMVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLPMParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GraphQLPMParserRULE_field)
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
		p.SetState(43)
		p.Match(GraphQLPMParserNAME)
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLPMParserT__4 {
		{
			p.SetState(44)
			p.Arguments()
		}

	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLPMParserT__1 {
		{
			p.SetState(47)
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

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLPMVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLPMParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GraphQLPMParserRULE_arguments)
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
		p.SetState(50)
		p.Match(GraphQLPMParserT__4)
	}
	{
		p.SetState(51)
		p.Argument()
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GraphQLPMParserT__2 {
		{
			p.SetState(52)
			p.Match(GraphQLPMParserT__2)
		}
		{
			p.SetState(53)
			p.Argument()
		}

		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(59)
		p.Match(GraphQLPMParserT__5)
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

func (s *ArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLPMVisitor:
		return t.VisitArgument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLPMParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GraphQLPMParserRULE_argument)

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
		p.SetState(61)
		p.Match(GraphQLPMParserNAME)
	}
	{
		p.SetState(62)
		p.Match(GraphQLPMParserT__6)
	}
	{
		p.SetState(63)
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

func (s *StringValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLPMVisitor:
		return t.VisitStringValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLPMParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GraphQLPMParserRULE_value)

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
		p.SetState(65)
		p.Match(GraphQLPMParserSTRING)
	}

	return localctx
}
