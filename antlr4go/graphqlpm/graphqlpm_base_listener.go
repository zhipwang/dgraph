// Generated from GraphQLPM.g4 by ANTLR 4.6.

package parser // GraphQLPM

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGraphQLPMListener is a complete listener for a parse tree produced by GraphQLPMParser.
type BaseGraphQLPMListener struct{}

var _ GraphQLPMListener = &BaseGraphQLPMListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGraphQLPMListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGraphQLPMListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGraphQLPMListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGraphQLPMListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseGraphQLPMListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseGraphQLPMListener) ExitDocument(ctx *DocumentContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseGraphQLPMListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseGraphQLPMListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterOperationDefinition is called when production operationDefinition is entered.
func (s *BaseGraphQLPMListener) EnterOperationDefinition(ctx *OperationDefinitionContext) {}

// ExitOperationDefinition is called when production operationDefinition is exited.
func (s *BaseGraphQLPMListener) ExitOperationDefinition(ctx *OperationDefinitionContext) {}

// EnterFilters is called when production filters is entered.
func (s *BaseGraphQLPMListener) EnterFilters(ctx *FiltersContext) {}

// ExitFilters is called when production filters is exited.
func (s *BaseGraphQLPMListener) ExitFilters(ctx *FiltersContext) {}

// EnterPairsNested is called when production pairsNested is entered.
func (s *BaseGraphQLPMListener) EnterPairsNested(ctx *PairsNestedContext) {}

// ExitPairsNested is called when production pairsNested is exited.
func (s *BaseGraphQLPMListener) ExitPairsNested(ctx *PairsNestedContext) {}

// EnterPairs is called when production pairs is entered.
func (s *BaseGraphQLPMListener) EnterPairs(ctx *PairsContext) {}

// ExitPairs is called when production pairs is exited.
func (s *BaseGraphQLPMListener) ExitPairs(ctx *PairsContext) {}

// EnterFilterOperation is called when production filterOperation is entered.
func (s *BaseGraphQLPMListener) EnterFilterOperation(ctx *FilterOperationContext) {}

// ExitFilterOperation is called when production filterOperation is exited.
func (s *BaseGraphQLPMListener) ExitFilterOperation(ctx *FilterOperationContext) {}

// EnterFuncName is called when production funcName is entered.
func (s *BaseGraphQLPMListener) EnterFuncName(ctx *FuncNameContext) {}

// ExitFuncName is called when production funcName is exited.
func (s *BaseGraphQLPMListener) ExitFuncName(ctx *FuncNameContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseGraphQLPMListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseGraphQLPMListener) ExitPair(ctx *PairContext) {}

// EnterFieldNameValue is called when production fieldNameValue is entered.
func (s *BaseGraphQLPMListener) EnterFieldNameValue(ctx *FieldNameValueContext) {}

// ExitFieldNameValue is called when production fieldNameValue is exited.
func (s *BaseGraphQLPMListener) ExitFieldNameValue(ctx *FieldNameValueContext) {}

// EnterOperationType is called when production operationType is entered.
func (s *BaseGraphQLPMListener) EnterOperationType(ctx *OperationTypeContext) {}

// ExitOperationType is called when production operationType is exited.
func (s *BaseGraphQLPMListener) ExitOperationType(ctx *OperationTypeContext) {}

// EnterSelectionSet is called when production selectionSet is entered.
func (s *BaseGraphQLPMListener) EnterSelectionSet(ctx *SelectionSetContext) {}

// ExitSelectionSet is called when production selectionSet is exited.
func (s *BaseGraphQLPMListener) ExitSelectionSet(ctx *SelectionSetContext) {}

// EnterField is called when production field is entered.
func (s *BaseGraphQLPMListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseGraphQLPMListener) ExitField(ctx *FieldContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseGraphQLPMListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseGraphQLPMListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseGraphQLPMListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseGraphQLPMListener) ExitArgument(ctx *ArgumentContext) {}

// EnterStringValue is called when production stringValue is entered.
func (s *BaseGraphQLPMListener) EnterStringValue(ctx *StringValueContext) {}

// ExitStringValue is called when production stringValue is exited.
func (s *BaseGraphQLPMListener) ExitStringValue(ctx *StringValueContext) {}
