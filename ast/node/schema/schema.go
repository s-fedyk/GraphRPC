package schema

import (
	"fmt"
	"strings"
	"graphrpc.com/ast/node/base"
	"graphrpc.com/ast/node/type"
	"graphrpc.com/ast/node/value"
)

// SchemaDefinition represents a schema definition
type SchemaDefinition struct {
	Description    *value.StringValue
	Directives     []*base.Directive
	OperationTypes []*OperationTypeDefinition
	Location       *base.Location
}

func (s *SchemaDefinition) String() string {
	parts := []string{"schema"}
	if len(s.Directives) > 0 {
		for _, d := range s.Directives {
			parts = append(parts, d.String())
		}
	}
	var ops []string
	for _, op := range s.OperationTypes {
		ops = append(ops, op.String())
	}
	parts = append(parts, fmt.Sprintf("{ %s }", strings.Join(ops, " ")))
	return strings.Join(parts, " ")
}

func (s *SchemaDefinition) GetLocation() *base.Location    { return s.Location }
func (s *SchemaDefinition) definitionNode()           {}
func (s *SchemaDefinition) typeSystemDefinitionNode() {}

// OperationTypeDefinition represents an operation type definition in a schema
type OperationTypeDefinition struct {
	Operation base.OperationType
	Type      *typ.NamedType
	Location  *base.Location
}

func (o *OperationTypeDefinition) String() string {
	return fmt.Sprintf("%s: %s", string(o.Operation), o.Type.String())
}

func (o *OperationTypeDefinition) GetLocation() *base.Location { return o.Location }