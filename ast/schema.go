package ast

import (
	"fmt"
	"strings"
)

type SchemaDefinition struct {
	Description    *StringValue
	Directives     []*Directive
	OperationTypes []*OperationTypeDefinition
	Location       *Location
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

func (s *SchemaDefinition) GetLocation() *Location    { return s.Location }
func (s *SchemaDefinition) definitionNode()           {}
func (s *SchemaDefinition) typeSystemDefinitionNode() {}

type OperationTypeDefinition struct {
	Operation OperationType
	Type      *NamedType
	Location  *Location
}

func (o *OperationTypeDefinition) String() string {
	return fmt.Sprintf("%s: %s", string(o.Operation), o.Type.String())
}

func (o *OperationTypeDefinition) GetLocation() *Location { return o.Location }

type ScalarTypeDefinition struct {
	Description *StringValue
	Name        *Name
	Directives  []*Directive
	Location    *Location
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

func (s *ScalarTypeDefinition) GetLocation() *Location    { return s.Location }
func (s *ScalarTypeDefinition) definitionNode()           {}
func (s *ScalarTypeDefinition) typeSystemDefinitionNode() {}

type ObjectTypeDefinition struct {
	Description *StringValue
	Name        *Name
	Interfaces  []*NamedType
	Directives  []*Directive
	Fields      []*FieldDefinition
	Location    *Location
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

func (o *ObjectTypeDefinition) GetLocation() *Location    { return o.Location }
func (o *ObjectTypeDefinition) definitionNode()           {}
func (o *ObjectTypeDefinition) typeSystemDefinitionNode() {}

type FieldDefinition struct {
	Description *StringValue
	Name        *Name
	Arguments   []*InputValueDefinition
	Type        Type
	Directives  []*Directive
	Location    *Location
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

func (f *FieldDefinition) GetLocation() *Location { return f.Location }

type InputValueDefinition struct {
	Description  *StringValue
	Name         *Name
	Type         Type
	DefaultValue Value
	Directives   []*Directive
	Location     *Location
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

func (i *InputValueDefinition) GetLocation() *Location { return i.Location }

type InterfaceTypeDefinition struct {
	Description *StringValue
	Name        *Name
	Interfaces  []*NamedType
	Directives  []*Directive
	Fields      []*FieldDefinition
	Location    *Location
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

func (i *InterfaceTypeDefinition) GetLocation() *Location    { return i.Location }
func (i *InterfaceTypeDefinition) definitionNode()           {}
func (i *InterfaceTypeDefinition) typeSystemDefinitionNode() {}

type UnionTypeDefinition struct {
	Description *StringValue
	Name        *Name
	Directives  []*Directive
	Types       []*NamedType
	Location    *Location
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

func (u *UnionTypeDefinition) GetLocation() *Location    { return u.Location }
func (u *UnionTypeDefinition) definitionNode()           {}
func (u *UnionTypeDefinition) typeSystemDefinitionNode() {}

type EnumTypeDefinition struct {
	Description *StringValue
	Name        *Name
	Directives  []*Directive
	Values      []*EnumValueDefinition
	Location    *Location
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

func (e *EnumTypeDefinition) GetLocation() *Location    { return e.Location }
func (e *EnumTypeDefinition) definitionNode()           {}
func (e *EnumTypeDefinition) typeSystemDefinitionNode() {}

type EnumValueDefinition struct {
	Description *StringValue
	Name        *Name
	Directives  []*Directive
	Location    *Location
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

func (e *EnumValueDefinition) GetLocation() *Location { return e.Location }

type InputObjectTypeDefinition struct {
	Description *StringValue
	Name        *Name
	Directives  []*Directive
	Fields      []*InputValueDefinition
	Location    *Location
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

func (i *InputObjectTypeDefinition) GetLocation() *Location    { return i.Location }
func (i *InputObjectTypeDefinition) definitionNode()           {}
func (i *InputObjectTypeDefinition) typeSystemDefinitionNode() {}

type DirectiveDefinition struct {
	Description *StringValue
	Name        *Name
	Arguments   []*InputValueDefinition
	Repeatable  bool
	Locations   []*Name
	Location    *Location
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

func (d *DirectiveDefinition) GetLocation() *Location    { return d.Location }
func (d *DirectiveDefinition) definitionNode()           {}
func (d *DirectiveDefinition) typeSystemDefinitionNode() {}

type SchemaExtension struct {
	Directives     []*Directive
	OperationTypes []*OperationTypeDefinition
	Location       *Location
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

func (s *SchemaExtension) GetLocation() *Location    { return s.Location }
func (s *SchemaExtension) definitionNode()           {}
func (s *SchemaExtension) typeSystemDefinitionNode() {}

type ScalarTypeExtension struct {
	Name       *Name
	Directives []*Directive
	Location   *Location
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

func (s *ScalarTypeExtension) GetLocation() *Location    { return s.Location }
func (s *ScalarTypeExtension) definitionNode()           {}
func (s *ScalarTypeExtension) typeSystemDefinitionNode() {}

type ObjectTypeExtension struct {
	Name       *Name
	Interfaces []*NamedType
	Directives []*Directive
	Fields     []*FieldDefinition
	Location   *Location
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

func (o *ObjectTypeExtension) GetLocation() *Location    { return o.Location }
func (o *ObjectTypeExtension) definitionNode()           {}
func (o *ObjectTypeExtension) typeSystemDefinitionNode() {}

