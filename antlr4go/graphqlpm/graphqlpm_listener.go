// Generated from GraphQLPM.g4 by ANTLR 4.6.

package parser // GraphQLPM

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GraphQLPMListener is a complete listener for a parse tree produced by GraphQLPMParser.
type GraphQLPMListener interface {
	antlr.ParseTreeListener

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterOperationDefinition is called when entering the operationDefinition production.
	EnterOperationDefinition(c *OperationDefinitionContext)

	// EnterFilters is called when entering the filters production.
	EnterFilters(c *FiltersContext)

	// EnterFilterOperation is called when entering the filterOperation production.
	EnterFilterOperation(c *FilterOperationContext)

	// EnterFuncName is called when entering the funcName production.
	EnterFuncName(c *FuncNameContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterFieldNameValue is called when entering the fieldNameValue production.
	EnterFieldNameValue(c *FieldNameValueContext)

	// EnterOperationType is called when entering the operationType production.
	EnterOperationType(c *OperationTypeContext)

	// EnterSelectionSet is called when entering the selectionSet production.
	EnterSelectionSet(c *SelectionSetContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterStringValue is called when entering the stringValue production.
	EnterStringValue(c *StringValueContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitOperationDefinition is called when exiting the operationDefinition production.
	ExitOperationDefinition(c *OperationDefinitionContext)

	// ExitFilters is called when exiting the filters production.
	ExitFilters(c *FiltersContext)

	// ExitFilterOperation is called when exiting the filterOperation production.
	ExitFilterOperation(c *FilterOperationContext)

	// ExitFuncName is called when exiting the funcName production.
	ExitFuncName(c *FuncNameContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitFieldNameValue is called when exiting the fieldNameValue production.
	ExitFieldNameValue(c *FieldNameValueContext)

	// ExitOperationType is called when exiting the operationType production.
	ExitOperationType(c *OperationTypeContext)

	// ExitSelectionSet is called when exiting the selectionSet production.
	ExitSelectionSet(c *SelectionSetContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitStringValue is called when exiting the stringValue production.
	ExitStringValue(c *StringValueContext)
}
