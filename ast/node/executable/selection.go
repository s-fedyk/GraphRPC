package executable

import (
	"fmt"
	"strings"
	"graphrpc.com/ast/node/base"
)

// SelectionSet represents a set of selections
type SelectionSet struct {
	Selections []base.Selection
	Location   *base.Location
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

func (s *SelectionSet) GetLocation() *base.Location { return s.Location }

// Field represents a field selection
type Field struct {
	Alias        *base.Name
	Name         *base.Name
	Arguments    []*base.Argument
	Directives   []*base.Directive
	SelectionSet *SelectionSet
	Location     *base.Location
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

func (f *Field) GetLocation() *base.Location { return f.Location }
func (f *Field) selectionNode()         {}