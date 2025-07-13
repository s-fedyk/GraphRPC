package node

// This file re-exports all node types for convenience

// Base interfaces and types
import (
	"graphrpc.com/ast/node/base"
	"graphrpc.com/ast/node/document"
	"graphrpc.com/ast/node/executable"
	"graphrpc.com/ast/node/value"
	"graphrpc.com/ast/node/type"
	"graphrpc.com/ast/node/schema"
	"graphrpc.com/ast/node/extension"
)

// Base interfaces
type (
	Node                   = base.Node
	Location              = base.Location
	Definition            = base.Definition
	ExecutableDefinition  = base.ExecutableDefinition
	TypeSystemDefinition  = base.TypeSystemDefinition
	Selection             = base.Selection
	Value                 = base.Value
	Type                  = base.Type
)

// Common types
type (
	Name          = base.Name
	OperationType = base.OperationType
	Argument      = base.Argument
	Directive     = base.Directive
)

// Document types
type (
	Document = document.Document
)

// Executable definition types
type (
	OperationDefinition   = executable.OperationDefinition
	VariableDefinition    = executable.VariableDefinition
	FragmentDefinition    = executable.FragmentDefinition
	InlineFragment        = executable.InlineFragment
	FragmentSpread        = executable.FragmentSpread
	SelectionSet          = executable.SelectionSet
	Field                 = executable.Field
)

// Value types
type (
	Variable      = value.Variable
	IntValue      = value.IntValue
	FloatValue    = value.FloatValue
	StringValue   = value.StringValue
	BooleanValue  = value.BooleanValue
	NullValue     = value.NullValue
	EnumValue     = value.EnumValue
	ListValue     = value.ListValue
	ObjectValue   = value.ObjectValue
	ObjectField   = value.ObjectField
)

// Type system types
type (
	NamedType    = typ.NamedType
	ListType     = typ.ListType
	NonNullType  = typ.NonNullType
)

// Schema definition types
type (
	SchemaDefinition           = schema.SchemaDefinition
	OperationTypeDefinition    = schema.OperationTypeDefinition
	ScalarTypeDefinition       = schema.ScalarTypeDefinition
	ObjectTypeDefinition       = schema.ObjectTypeDefinition
	InterfaceTypeDefinition    = schema.InterfaceTypeDefinition
	UnionTypeDefinition        = schema.UnionTypeDefinition
	EnumTypeDefinition         = schema.EnumTypeDefinition
	InputObjectTypeDefinition  = schema.InputObjectTypeDefinition
	FieldDefinition            = schema.FieldDefinition
	InputValueDefinition       = schema.InputValueDefinition
	EnumValueDefinition        = schema.EnumValueDefinition
	DirectiveDefinition        = schema.DirectiveDefinition
)

// Extension types
type (
	SchemaExtension              = extension.SchemaExtension
	ScalarTypeExtension          = extension.ScalarTypeExtension
	ObjectTypeExtension          = extension.ObjectTypeExtension
	InterfaceTypeExtension       = extension.InterfaceTypeExtension
	UnionTypeExtension           = extension.UnionTypeExtension
	EnumTypeExtension            = extension.EnumTypeExtension
	InputObjectTypeExtension     = extension.InputObjectTypeExtension
)

// Constants
const (
	Query        = base.Query
	Mutation     = base.Mutation
	Subscription = base.Subscription
)