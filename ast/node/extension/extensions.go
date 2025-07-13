package extension

import (
	"fmt"
	"strings"
	"graphrpc.com/ast/node/base"
	"graphrpc.com/ast/node/type"
	"graphrpc.com/ast/node/schema"
)

// SchemaExtension represents a schema extension
type SchemaExtension struct {
	Directives     []*base.Directive
	OperationTypes []*schema.OperationTypeDefinition
	Location       *base.Location
}

func (s *SchemaExtension) String() string {
	parts := []string{"extend", "schema"}
	if len(s.Directives) > 0 {
		for _, d := range s.Directives {
			parts = append(parts, d.String())
		}
	}
	if len(s.OperationTypes) > 0 {
		var ops []string
		for _, op := range s.OperationTypes {
			ops = append(ops, op.String())
		}
		parts = append(parts, fmt.Sprintf("{ %s }", strings.Join(ops, " ")))
	}
	return strings.Join(parts, " ")
}

func (s *SchemaExtension) GetLocation() *base.Location    { return s.Location }
func (s *SchemaExtension) definitionNode()           {}
func (s *SchemaExtension) typeSystemDefinitionNode() {}

// ScalarTypeExtension represents a scalar type extension
type ScalarTypeExtension struct {
	Name       *base.Name
	Directives []*base.Directive
	Location   *base.Location
}

func (s *ScalarTypeExtension) String() string {
	parts := []string{"extend", "scalar", s.Name.Value}
	if len(s.Directives) > 0 {
		for _, d := range s.Directives {
			parts = append(parts, d.String())
		}
	}
	return strings.Join(parts, " ")
}

func (s *ScalarTypeExtension) GetLocation() *base.Location    { return s.Location }
func (s *ScalarTypeExtension) definitionNode()           {}
func (s *ScalarTypeExtension) typeSystemDefinitionNode() {}

// ObjectTypeExtension represents an object type extension
type ObjectTypeExtension struct {
	Name       *base.Name
	Interfaces []*typ.NamedType
	Directives []*base.Directive
	Fields     []*schema.FieldDefinition
	Location   *base.Location
}

func (o *ObjectTypeExtension) String() string {
	parts := []string{"extend", "type", o.Name.Value}
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

func (o *ObjectTypeExtension) GetLocation() *base.Location    { return o.Location }
func (o *ObjectTypeExtension) definitionNode()           {}
func (o *ObjectTypeExtension) typeSystemDefinitionNode() {}

// InterfaceTypeExtension represents an interface type extension
type InterfaceTypeExtension struct {
	Name       *base.Name
	Interfaces []*typ.NamedType
	Directives []*base.Directive
	Fields     []*schema.FieldDefinition
	Location   *base.Location
}

func (i *InterfaceTypeExtension) String() string {
	parts := []string{"extend", "interface", i.Name.Value}
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

func (i *InterfaceTypeExtension) GetLocation() *base.Location    { return i.Location }
func (i *InterfaceTypeExtension) definitionNode()           {}
func (i *InterfaceTypeExtension) typeSystemDefinitionNode() {}

// UnionTypeExtension represents a union type extension
type UnionTypeExtension struct {
	Name       *base.Name
	Directives []*base.Directive
	Types      []*typ.NamedType
	Location   *base.Location
}

func (u *UnionTypeExtension) String() string {
	parts := []string{"extend", "union", u.Name.Value}
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

func (u *UnionTypeExtension) GetLocation() *base.Location    { return u.Location }
func (u *UnionTypeExtension) definitionNode()           {}
func (u *UnionTypeExtension) typeSystemDefinitionNode() {}

// EnumTypeExtension represents an enum type extension
type EnumTypeExtension struct {
	Name       *base.Name
	Directives []*base.Directive
	Values     []*schema.EnumValueDefinition
	Location   *base.Location
}

func (e *EnumTypeExtension) String() string {
	parts := []string{"extend", "enum", e.Name.Value}
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

func (e *EnumTypeExtension) GetLocation() *base.Location    { return e.Location }
func (e *EnumTypeExtension) definitionNode()           {}
func (e *EnumTypeExtension) typeSystemDefinitionNode() {}

// InputObjectTypeExtension represents an input object type extension
type InputObjectTypeExtension struct {
	Name       *base.Name
	Directives []*base.Directive
	Fields     []*schema.InputValueDefinition
	Location   *base.Location
}

func (i *InputObjectTypeExtension) String() string {
	parts := []string{"extend", "input", i.Name.Value}
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

func (i *InputObjectTypeExtension) GetLocation() *base.Location    { return i.Location }
func (i *InputObjectTypeExtension) definitionNode()           {}
func (i *InputObjectTypeExtension) typeSystemDefinitionNode() {}