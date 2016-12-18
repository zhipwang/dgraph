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

	// EnterOperationType is called when entering the operationType production.
	EnterOperationType(c *OperationTypeContext)

	// EnterSelectionSet is called when entering the selectionSet production.
	EnterSelectionSet(c *SelectionSetContext)

	// EnterSelection is called when entering the selection production.
	EnterSelection(c *SelectionContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

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

	// ExitOperationType is called when exiting the operationType production.
	ExitOperationType(c *OperationTypeContext)

	// ExitSelectionSet is called when exiting the selectionSet production.
	ExitSelectionSet(c *SelectionSetContext)

	// ExitSelection is called when exiting the selection production.
	ExitSelection(c *SelectionContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitStringValue is called when exiting the stringValue production.
	ExitStringValue(c *StringValueContext)
}
