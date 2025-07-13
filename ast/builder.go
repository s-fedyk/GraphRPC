package ast

import (
	"github.com/antlr4-go/antlr/v4"
	parser "graphrpc.com/parser/gen"
)

type ASTBuilder struct {
	parser.GraphQLVisitor
}

func NewASTBuilder() *ASTBuilder {
	return &ASTBuilder{
		&parser.BaseGraphQLVisitor{},
	}
}

func (b *ASTBuilder) VisitDocument(ctx *parser.DocumentContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitDefinition(ctx *parser.DefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitExecutableDocument(ctx *parser.ExecutableDocumentContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitExecutableDefinition(ctx *parser.ExecutableDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitOperationDefinition(ctx *parser.OperationDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitOperationType(ctx *parser.OperationTypeContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitSelectionSet(ctx *parser.SelectionSetContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitSelection(ctx *parser.SelectionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitField(ctx *parser.FieldContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitArguments(ctx *parser.ArgumentsContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitArgument(ctx *parser.ArgumentContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitAlias(ctx *parser.AliasContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitFragmentSpread(ctx *parser.FragmentSpreadContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitFragmentDefinition(ctx *parser.FragmentDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitFragmentName(ctx *parser.FragmentNameContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitTypeCondition(ctx *parser.TypeConditionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitInlineFragment(ctx *parser.InlineFragmentContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitValue(ctx *parser.ValueContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitIntValue(ctx *parser.IntValueContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitFloatValue(ctx *parser.FloatValueContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitBooleanValue(ctx *parser.BooleanValueContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitStringValue(ctx *parser.StringValueContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitNullValue(ctx *parser.NullValueContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitEnumValue(ctx *parser.EnumValueContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitListValue(ctx *parser.ListValueContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitObjectValue(ctx *parser.ObjectValueContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitObjectField(ctx *parser.ObjectFieldContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitVariable(ctx *parser.VariableContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitVariableDefinitions(ctx *parser.VariableDefinitionsContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitVariableDefinition(ctx *parser.VariableDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitDefaultValue(ctx *parser.DefaultValueContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitType_(ctx *parser.Type_Context) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitNamedType(ctx *parser.NamedTypeContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitListType(ctx *parser.ListTypeContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitDirectives(ctx *parser.DirectivesContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitDirective(ctx *parser.DirectiveContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitTypeSystemDocument(ctx *parser.TypeSystemDocumentContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitTypeSystemDefinition(ctx *parser.TypeSystemDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitTypeSystemExtensionDocument(ctx *parser.TypeSystemExtensionDocumentContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitTypeSystemExtension(ctx *parser.TypeSystemExtensionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitSchemaDefinition(ctx *parser.SchemaDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitRootOperationTypeDefinition(ctx *parser.RootOperationTypeDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitSchemaExtension(ctx *parser.SchemaExtensionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitDescription(ctx *parser.DescriptionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitTypeDefinition(ctx *parser.TypeDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitTypeExtension(ctx *parser.TypeExtensionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitScalarTypeDefinition(ctx *parser.ScalarTypeDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitScalarTypeExtension(ctx *parser.ScalarTypeExtensionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitObjectTypeDefinition(ctx *parser.ObjectTypeDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitImplementsInterfaces(ctx *parser.ImplementsInterfacesContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitFieldsDefinition(ctx *parser.FieldsDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitFieldDefinition(ctx *parser.FieldDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitArgumentsDefinition(ctx *parser.ArgumentsDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitInputValueDefinition(ctx *parser.InputValueDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitObjectTypeExtension(ctx *parser.ObjectTypeExtensionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitInterfaceTypeDefinition(ctx *parser.InterfaceTypeDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitInterfaceTypeExtension(ctx *parser.InterfaceTypeExtensionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitUnionTypeDefinition(ctx *parser.UnionTypeDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitUnionMemberTypes(ctx *parser.UnionMemberTypesContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitUnionTypeExtension(ctx *parser.UnionTypeExtensionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitEnumTypeDefinition(ctx *parser.EnumTypeDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitEnumValuesDefinition(ctx *parser.EnumValuesDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitEnumValueDefinition(ctx *parser.EnumValueDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitEnumTypeExtension(ctx *parser.EnumTypeExtensionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitInputObjectTypeDefinition(ctx *parser.InputObjectTypeDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitInputFieldsDefinition(ctx *parser.InputFieldsDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitInputObjectTypeExtension(ctx *parser.InputObjectTypeExtensionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitDirectiveDefinition(ctx *parser.DirectiveDefinitionContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitDirectiveLocations(ctx *parser.DirectiveLocationsContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitDirectiveLocation(ctx *parser.DirectiveLocationContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitExecutableDirectiveLocation(ctx *parser.ExecutableDirectiveLocationContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitTypeSystemDirectiveLocation(ctx *parser.TypeSystemDirectiveLocationContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitName(ctx *parser.NameContext) interface{} {
	return b.VisitChildren(ctx)
}

func (b *ASTBuilder) VisitChildren(ctx antlr.RuleNode) interface{} {

	for _, child := range ctx.GetChildren() {
		child.(antlr.ParseTree).Accept(b)
	}

	return nil
}
