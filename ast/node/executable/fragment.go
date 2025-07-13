package executable

import (
	"strings"
	"graphrpc.com/ast/node/base"
	"graphrpc.com/ast/node/type"
)

// FragmentDefinition represents a named fragment definition
type FragmentDefinition struct {
	Name          *base.Name
	TypeCondition *typ.NamedType
	Directives    []*base.Directive
	SelectionSet  *SelectionSet
	Location      *base.Location
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

func (f *FragmentDefinition) GetLocation() *base.Location     { return f.Location }
func (f *FragmentDefinition) definitionNode()           {}
func (f *FragmentDefinition) executableDefinitionNode() {}

// InlineFragment represents an inline fragment
type InlineFragment struct {
	TypeCondition *typ.NamedType
	Directives    []*base.Directive
	SelectionSet  *SelectionSet
	Location      *base.Location
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

func (i *InlineFragment) GetLocation() *base.Location { return i.Location }
func (i *InlineFragment) selectionNode()         {}

// FragmentSpread represents a fragment spread
type FragmentSpread struct {
	Name       *base.Name
	Directives []*base.Directive
	Location   *base.Location
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

func (f *FragmentSpread) GetLocation() *base.Location { return f.Location }
func (f *FragmentSpread) selectionNode()         {}