// Generated from GraphQLPM.g4 by ANTLR 4.6.

package parser // GraphQLPM

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseGraphQLPMVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGraphQLPMVisitor) VisitDocument(ctx *DocumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLPMVisitor) VisitDefinition(ctx *DefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLPMVisitor) VisitOperationDefinition(ctx *OperationDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLPMVisitor) VisitOperationType(ctx *OperationTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLPMVisitor) VisitSelectionSet(ctx *SelectionSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLPMVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLPMVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLPMVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLPMVisitor) VisitStringValue(ctx *StringValueContext) interface{} {
	return v.VisitChildren(ctx)
}
