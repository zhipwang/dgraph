// Generated from GraphQLPM.g4 by ANTLR 4.6.

package parser // GraphQLPM

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by GraphQLPMParser.
type GraphQLPMVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GraphQLPMParser#document.
	VisitDocument(ctx *DocumentContext) interface{}

	// Visit a parse tree produced by GraphQLPMParser#definition.
	VisitDefinition(ctx *DefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLPMParser#operationDefinition.
	VisitOperationDefinition(ctx *OperationDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLPMParser#operationType.
	VisitOperationType(ctx *OperationTypeContext) interface{}

	// Visit a parse tree produced by GraphQLPMParser#selectionSet.
	VisitSelectionSet(ctx *SelectionSetContext) interface{}

	// Visit a parse tree produced by GraphQLPMParser#selection.
	VisitSelection(ctx *SelectionContext) interface{}

	// Visit a parse tree produced by GraphQLPMParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by GraphQLPMParser#fieldName.
	VisitFieldName(ctx *FieldNameContext) interface{}

	// Visit a parse tree produced by GraphQLPMParser#alias.
	VisitAlias(ctx *AliasContext) interface{}

	// Visit a parse tree produced by GraphQLPMParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by GraphQLPMParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by GraphQLPMParser#stringValue.
	VisitStringValue(ctx *StringValueContext) interface{}
}
