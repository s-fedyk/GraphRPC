package schema

import (
	"fmt"
	"strings"
	"graphrpc.com/ast/node/base"
	"graphrpc.com/ast/node/type"
	"graphrpc.com/ast/node/value"
)

// ScalarTypeDefinition represents a scalar type definition
type ScalarTypeDefinition struct {
	Description *value.StringValue
	Name        *base.Name
	Directives  []*base.Directive
	Location    *base.Location
}

func (s *ScalarTypeDefinition) String() string {
	parts := []string{"scalar", s.Name.Value}
	if len(s.Directives) > 0 {
		for _, d := range s.Directives {
			parts = append(parts, d.String())
		}
	}
	return strings.Join(parts, " ")
}

func (s *ScalarTypeDefinition) GetLocation() *base.Location    { return s.Location }
func (s *ScalarTypeDefinition) definitionNode()           {}
func (s *ScalarTypeDefinition) typeSystemDefinitionNode() {}

// ObjectTypeDefinition represents an object type definition
type ObjectTypeDefinition struct {
	Description *value.StringValue
	Name        *base.Name
	Interfaces  []*typ.NamedType
	Directives  []*base.Directive
	Fields      []*FieldDefinition
	Location    *base.Location
}

func (o *ObjectTypeDefinition) String() string {
	parts := []string{"type", o.Name.Value}
	if len(o.Interfaces) > 0 {
		var interfaces []string
		for _, iface := range o.Interfaces {
			interfaces = append(interfaces, iface.String())
		}
		parts = append(parts, "implements", strings.Join(interfaces, " & "))
	}
	if len(o.Directives) > 0 {
		for _, d := range o.Directives {
			parts = append(parts, d.String())
		}
	}
	if len(o.Fields) > 0 {
		var fields []string
		for _, field := range o.Fields {
			fields = append(fields, field.String())
		}
		parts = append(parts, fmt.Sprintf("{ %s }", strings.Join(fields, " ")))
	}
	return strings.Join(parts, " ")
}

func (o *ObjectTypeDefinition) GetLocation() *base.Location    { return o.Location }
func (o *ObjectTypeDefinition) definitionNode()           {}
func (o *ObjectTypeDefinition) typeSystemDefinitionNode() {}

// InterfaceTypeDefinition represents an interface type definition
type InterfaceTypeDefinition struct {
	Description *value.StringValue
	Name        *base.Name
	Interfaces  []*typ.NamedType
	Directives  []*base.Directive
	Fields      []*FieldDefinition
	Location    *base.Location
}

func (i *InterfaceTypeDefinition) String() string {
	parts := []string{"interface", i.Name.Value}
	if len(i.Interfaces) > 0 {
		var interfaces []string
		for _, iface := range i.Interfaces {
			interfaces = append(interfaces, iface.String())
		}
		parts = append(parts, "implements", strings.Join(interfaces, " & "))
	}
	if len(i.Directives) > 0 {
		for _, d := range i.Directives {
			parts = append(parts, d.String())
		}
	}
	if len(i.Fields) > 0 {
		var fields []string
		for _, field := range i.Fields {
			fields = append(fields, field.String())
		}
		parts = append(parts, fmt.Sprintf("{ %s }", strings.Join(fields, " ")))
	}
	return strings.Join(parts, " ")
}

func (i *InterfaceTypeDefinition) GetLocation() *base.Location    { return i.Location }
func (i *InterfaceTypeDefinition) definitionNode()           {}
func (i *InterfaceTypeDefinition) typeSystemDefinitionNode() {}

// UnionTypeDefinition represents a union type definition
type UnionTypeDefinition struct {
	Description *value.StringValue
	Name        *base.Name
	Directives  []*base.Directive
	Types       []*typ.NamedType
	Location    *base.Location
}

func (u *UnionTypeDefinition) String() string {
	parts := []string{"union", u.Name.Value}
	if len(u.Directives) > 0 {
		for _, d := range u.Directives {
			parts = append(parts, d.String())
		}
	}
	if len(u.Types) > 0 {
		var types []string
		for _, t := range u.Types {
			types = append(types, t.String())
		}
		parts = append(parts, "=", strings.Join(types, " | "))
	}
	return strings.Join(parts, " ")
}

func (u *UnionTypeDefinition) GetLocation() *base.Location    { return u.Location }
func (u *UnionTypeDefinition) definitionNode()           {}
func (u *UnionTypeDefinition) typeSystemDefinitionNode() {}

// EnumTypeDefinition represents an enum type definition
type EnumTypeDefinition struct {
	Description *value.StringValue
	Name        *base.Name
	Directives  []*base.Directive
	Values      []*EnumValueDefinition
	Location    *base.Location
}

func (e *EnumTypeDefinition) String() string {
	parts := []string{"enum", e.Name.Value}
	if len(e.Directives) > 0 {
		for _, d := range e.Directives {
			parts = append(parts, d.String())
		}
	}
	if len(e.Values) > 0 {
		var values []string
		for _, val := range e.Values {
			values = append(values, val.String())
		}
		parts = append(parts, fmt.Sprintf("{ %s }", strings.Join(values, " ")))
	}
	return strings.Join(parts, " ")
}

func (e *EnumTypeDefinition) GetLocation() *base.Location    { return e.Location }
func (e *EnumTypeDefinition) definitionNode()           {}
func (e *EnumTypeDefinition) typeSystemDefinitionNode() {}

// InputObjectTypeDefinition represents an input object type definition
type InputObjectTypeDefinition struct {
	Description *value.StringValue
	Name        *base.Name
	Directives  []*base.Directive
	Fields      []*InputValueDefinition
	Location    *base.Location
}

func (i *InputObjectTypeDefinition) String() string {
	parts := []string{"input", i.Name.Value}
	if len(i.Directives) > 0 {
		for _, d := range i.Directives {
			parts = append(parts, d.String())
		}
	}
	if len(i.Fields) > 0 {
		var fields []string
		for _, field := range i.Fields {
			fields = append(fields, field.String())
		}
		parts = append(parts, fmt.Sprintf("{ %s }", strings.Join(fields, " ")))
	}
	return strings.Join(parts, " ")
}

func (i *InputObjectTypeDefinition) GetLocation() *base.Location    { return i.Location }
func (i *InputObjectTypeDefinition) definitionNode()           {}
func (i *InputObjectTypeDefinition) typeSystemDefinitionNode() {}