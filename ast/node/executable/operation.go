package executable

import (
	"fmt"
	"strings"
	"graphrpc.com/ast/node/base"
	"graphrpc.com/ast/node/value"
)

// OperationDefinition represents a GraphQL operation (query, mutation, subscription)
type OperationDefinition struct {
	Type                base.OperationType
	Name                *base.Name
	VariableDefinitions []*VariableDefinition
	Directives          []*base.Directive
	SelectionSet        *SelectionSet
	Location            *base.Location
}

func (o *OperationDefinition) String() string {
	var parts []string
	parts = append(parts, string(o.Type))
	if o.Name != nil {
		parts = append(parts, o.Name.Value)
	}
	if len(o.VariableDefinitions) > 0 {
		var vars []string
		for _, v := range o.VariableDefinitions {
			vars = append(vars, v.String())
		}
		parts = append(parts, fmt.Sprintf("(%s)", strings.Join(vars, ", ")))
	}
	if len(o.Directives) > 0 {
		for _, d := range o.Directives {
			parts = append(parts, d.String())
		}
	}
	parts = append(parts, o.SelectionSet.String())
	return strings.Join(parts, " ")
}

func (o *OperationDefinition) GetLocation() *base.Location         { return o.Location }
func (o *OperationDefinition) definitionNode()               {}
func (o *OperationDefinition) executableDefinitionNode()     {}

// VariableDefinition represents a variable definition in an operation
type VariableDefinition struct {
	Variable     *value.Variable
	Type         base.Type
	DefaultValue base.Value
	Directives   []*base.Directive
	Location     *base.Location
}

func (v *VariableDefinition) String() string {
	parts := []string{v.Variable.String() + ":", v.Type.String()}
	if v.DefaultValue != nil {
		parts = append(parts, "=", v.DefaultValue.String())
	}
	if len(v.Directives) > 0 {
		for _, d := range v.Directives {
			parts = append(parts, d.String())
		}
	}
	return strings.Join(parts, " ")
}

func (v *VariableDefinition) GetLocation() *base.Location { return v.Location }