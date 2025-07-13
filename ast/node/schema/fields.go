package schema

import (
	"fmt"
	"graphrpc.com/ast/node/base"
	"graphrpc.com/ast/node/value"
	"strings"
)

// FieldDefinition represents a field definition in an object or interface
type FieldDefinition struct {
	Description *value.StringValue
	Name        *base.Name
	Arguments   []*InputValueDefinition
	Type        base.Type
	Directives  []*base.Directive
	Location    *base.Location
}

func (f *FieldDefinition) String() string {
	parts := []string{f.Name.Value}
	if len(f.Arguments) > 0 {
		var args []string
		for _, arg := range f.Arguments {
			args = append(args, arg.String())
		}
		parts = append(parts, fmt.Sprintf("(%s)", strings.Join(args, ", ")))
	}
	parts = append(parts, ":", f.Type.String())
	if len(f.Directives) > 0 {
		for _, d := range f.Directives {
			parts = append(parts, d.String())
		}
	}
	return strings.Join(parts, " ")
}

func (f *FieldDefinition) GetLocation() *base.Location { return f.Location }

// InputValueDefinition represents an input value definition (argument or input field)
type InputValueDefinition struct {
	Description  *value.StringValue
	Name         *base.Name
	Type         base.Type
	DefaultValue base.Value
	Directives   []*base.Directive
	Location     *base.Location
}

func (i *InputValueDefinition) String() string {
	parts := []string{i.Name.Value + ":", i.Type.String()}
	if i.DefaultValue != nil {
		parts = append(parts, "=", i.DefaultValue.String())
	}
	if len(i.Directives) > 0 {
		for _, d := range i.Directives {
			parts = append(parts, d.String())
		}
	}
	return strings.Join(parts, " ")
}

func (i *InputValueDefinition) GetLocation() *base.Location { return i.Location }

// EnumValueDefinition represents an enum value definition
type EnumValueDefinition struct {
	Description *value.StringValue
	Name        *base.Name
	Directives  []*base.Directive
	Location    *base.Location
}

func (e *EnumValueDefinition) String() string {
	parts := []string{e.Name.Value}
	if len(e.Directives) > 0 {
		for _, d := range e.Directives {
			parts = append(parts, d.String())
		}
	}
	return strings.Join(parts, " ")
}

func (e *EnumValueDefinition) GetLocation() *base.Location { return e.Location }

// DirectiveDefinition represents a directive definition
type DirectiveDefinition struct {
	Description *value.StringValue
	Name        *base.Name
	Arguments   []*InputValueDefinition
	Repeatable  bool
	Locations   []*base.Name
	Location    *base.Location
}

func (d *DirectiveDefinition) String() string {
	parts := []string{"directive", "@" + d.Name.Value}
	if len(d.Arguments) > 0 {
		var args []string
		for _, arg := range d.Arguments {
			args = append(args, arg.String())
		}
		parts = append(parts, fmt.Sprintf("(%s)", strings.Join(args, ", ")))
	}
	if d.Repeatable {
		parts = append(parts, "repeatable")
	}
	parts = append(parts, "on")
	var locations []string
	for _, loc := range d.Locations {
		locations = append(locations, loc.Value)
	}
	parts = append(parts, strings.Join(locations, " | "))
	return strings.Join(parts, " ")
}

func (d *DirectiveDefinition) GetLocation() *base.Location { return d.Location }
func (d *DirectiveDefinition) definitionNode()             {}
func (d *DirectiveDefinition) typeSystemDefinitionNode()   {}

