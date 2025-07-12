package ast

import (
	"fmt"
	"strings"
)

type Node interface {
	String() string
	GetLocation() *Location
}

type Location struct {
	Line   int
	Column int
}

type Document struct {
	Definitions []Definition
	Location    *Location
}

func (d *Document) String() string {
	var parts []string
	for _, def := range d.Definitions {
		parts = append(parts, def.String())
	}
	return strings.Join(parts, "\n\n")
}

func (d *Document) GetLocation() *Location {
	return d.Location
}

type Definition interface {
	Node
	definitionNode()
}

type ExecutableDefinition interface {
	Definition
	executableDefinitionNode()
}

type TypeSystemDefinition interface {
	Definition
	typeSystemDefinitionNode()
}

type OperationDefinition struct {
	Type                OperationType
	Name                *Name
	VariableDefinitions []*VariableDefinition
	Directives          []*Directive
	SelectionSet        *SelectionSet
	Location            *Location
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

func (o *OperationDefinition) GetLocation() *Location         { return o.Location }
func (o *OperationDefinition) definitionNode()               {}
func (o *OperationDefinition) executableDefinitionNode()     {}

type OperationType string

const (
	Query        OperationType = "query"
	Mutation     OperationType = "mutation"
	Subscription OperationType = "subscription"
)

type FragmentDefinition struct {
	Name         *Name
	TypeCondition *NamedType
	Directives   []*Directive
	SelectionSet *SelectionSet
	Location     *Location
}

func (f *FragmentDefinition) String() string {
	parts := []string{"fragment", f.Name.Value, "on", f.TypeCondition.String()}
	if len(f.Directives) > 0 {
		for _, d := range f.Directives {
			parts = append(parts, d.String())
		}
	}
	parts = append(parts, f.SelectionSet.String())
	return strings.Join(parts, " ")
}

func (f *FragmentDefinition) GetLocation() *Location     { return f.Location }
func (f *FragmentDefinition) definitionNode()           {}
func (f *FragmentDefinition) executableDefinitionNode() {}

type SelectionSet struct {
	Selections []Selection
	Location   *Location
}

func (s *SelectionSet) String() string {
	if len(s.Selections) == 0 {
		return "{}"
	}
	var parts []string
	for _, sel := range s.Selections {
		parts = append(parts, sel.String())
	}
	return fmt.Sprintf("{ %s }", strings.Join(parts, " "))
}

func (s *SelectionSet) GetLocation() *Location { return s.Location }

type Selection interface {
	Node
	selectionNode()
}

type Field struct {
	Alias        *Name
	Name         *Name
	Arguments    []*Argument
	Directives   []*Directive
	SelectionSet *SelectionSet
	Location     *Location
}

func (f *Field) String() string {
	var parts []string
	if f.Alias != nil {
		parts = append(parts, fmt.Sprintf("%s:", f.Alias.Value))
	}
	parts = append(parts, f.Name.Value)
	if len(f.Arguments) > 0 {
		var args []string
		for _, arg := range f.Arguments {
			args = append(args, arg.String())
		}
		parts = append(parts, fmt.Sprintf("(%s)", strings.Join(args, ", ")))
	}
	if len(f.Directives) > 0 {
		for _, d := range f.Directives {
			parts = append(parts, d.String())
		}
	}
	if f.SelectionSet != nil {
		parts = append(parts, f.SelectionSet.String())
	}
	return strings.Join(parts, " ")
}

func (f *Field) GetLocation() *Location { return f.Location }
func (f *Field) selectionNode()         {}

type InlineFragment struct {
	TypeCondition *NamedType
	Directives    []*Directive
	SelectionSet  *SelectionSet
	Location      *Location
}

func (i *InlineFragment) String() string {
	parts := []string{"..."}
	if i.TypeCondition != nil {
		parts = append(parts, "on", i.TypeCondition.String())
	}
	if len(i.Directives) > 0 {
		for _, d := range i.Directives {
			parts = append(parts, d.String())
		}
	}
	parts = append(parts, i.SelectionSet.String())
	return strings.Join(parts, " ")
}

func (i *InlineFragment) GetLocation() *Location { return i.Location }
func (i *InlineFragment) selectionNode()         {}

type FragmentSpread struct {
	Name       *Name
	Directives []*Directive
	Location   *Location
}

func (f *FragmentSpread) String() string {
	parts := []string{"...", f.Name.Value}
	if len(f.Directives) > 0 {
		for _, d := range f.Directives {
			parts = append(parts, d.String())
		}
	}
	return strings.Join(parts, " ")
}

func (f *FragmentSpread) GetLocation() *Location { return f.Location }
func (f *FragmentSpread) selectionNode()         {}

type Argument struct {
	Name     *Name
	Value    Value
	Location *Location
}

func (a *Argument) String() string {
	return fmt.Sprintf("%s: %s", a.Name.Value, a.Value.String())
}

func (a *Argument) GetLocation() *Location { return a.Location }

type Name struct {
	Value    string
	Location *Location
}

func (n *Name) String() string {
	return n.Value
}

func (n *Name) GetLocation() *Location { return n.Location }

type Value interface {
	Node
	valueNode()
}

type Variable struct {
	Name     *Name
	Location *Location
}

func (v *Variable) String() string {
	return fmt.Sprintf("$%s", v.Name.Value)
}

func (v *Variable) GetLocation() *Location { return v.Location }
func (v *Variable) valueNode()             {}

type IntValue struct {
	Value    string
	Location *Location
}

func (i *IntValue) String() string {
	return i.Value
}

func (i *IntValue) GetLocation() *Location { return i.Location }
func (i *IntValue) valueNode()             {}

type FloatValue struct {
	Value    string
	Location *Location
}

func (f *FloatValue) String() string {
	return f.Value
}

func (f *FloatValue) GetLocation() *Location { return f.Location }
func (f *FloatValue) valueNode()             {}

type StringValue struct {
	Value    string
	Block    bool
	Location *Location
}

func (s *StringValue) String() string {
	if s.Block {
		return fmt.Sprintf(`"""%s"""`, s.Value)
	}
	return fmt.Sprintf(`"%s"`, s.Value)
}

func (s *StringValue) GetLocation() *Location { return s.Location }
func (s *StringValue) valueNode()             {}

type BooleanValue struct {
	Value    bool
	Location *Location
}

func (b *BooleanValue) String() string {
	if b.Value {
		return "true"
	}
	return "false"
}

func (b *BooleanValue) GetLocation() *Location { return b.Location }
func (b *BooleanValue) valueNode()             {}

type NullValue struct {
	Location *Location
}

func (n *NullValue) String() string {
	return "null"
}

func (n *NullValue) GetLocation() *Location { return n.Location }
func (n *NullValue) valueNode()             {}

type EnumValue struct {
	Value    string
	Location *Location
}

func (e *EnumValue) String() string {
	return e.Value
}

func (e *EnumValue) GetLocation() *Location { return e.Location }
func (e *EnumValue) valueNode()             {}

type ListValue struct {
	Values   []Value
	Location *Location
}

func (l *ListValue) String() string {
	var values []string
	for _, v := range l.Values {
		values = append(values, v.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(values, ", "))
}

func (l *ListValue) GetLocation() *Location { return l.Location }
func (l *ListValue) valueNode()             {}

type ObjectValue struct {
	Fields   []*ObjectField
	Location *Location
}

func (o *ObjectValue) String() string {
	var fields []string
	for _, f := range o.Fields {
		fields = append(fields, f.String())
	}
	return fmt.Sprintf("{%s}", strings.Join(fields, ", "))
}

func (o *ObjectValue) GetLocation() *Location { return o.Location }
func (o *ObjectValue) valueNode()             {}

type ObjectField struct {
	Name     *Name
	Value    Value
	Location *Location
}

func (o *ObjectField) String() string {
	return fmt.Sprintf("%s: %s", o.Name.Value, o.Value.String())
}

func (o *ObjectField) GetLocation() *Location { return o.Location }

type VariableDefinition struct {
	Variable     *Variable
	Type         Type
	DefaultValue Value
	Directives   []*Directive
	Location     *Location
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

func (v *VariableDefinition) GetLocation() *Location { return v.Location }

type Type interface {
	Node
	typeNode()
}

type NamedType struct {
	Name     *Name
	Location *Location
}

func (n *NamedType) String() string {
	return n.Name.Value
}

func (n *NamedType) GetLocation() *Location { return n.Location }
func (n *NamedType) typeNode()              {}

type ListType struct {
	Type     Type
	Location *Location
}

func (l *ListType) String() string {
	return fmt.Sprintf("[%s]", l.Type.String())
}

func (l *ListType) GetLocation() *Location { return l.Location }
func (l *ListType) typeNode()              {}

type NonNullType struct {
	Type     Type
	Location *Location
}

func (n *NonNullType) String() string {
	return fmt.Sprintf("%s!", n.Type.String())
}

func (n *NonNullType) GetLocation() *Location { return n.Location }
func (n *NonNullType) typeNode()              {}

type Directive struct {
	Name      *Name
	Arguments []*Argument
	Location  *Location
}

func (d *Directive) String() string {
	parts := []string{"@" + d.Name.Value}
	if len(d.Arguments) > 0 {
		var args []string
		for _, arg := range d.Arguments {
			args = append(args, arg.String())
		}
		parts = append(parts, fmt.Sprintf("(%s)", strings.Join(args, ", ")))
	}
	return strings.Join(parts, "")
}

func (d *Directive) GetLocation() *Location { return d.Location }