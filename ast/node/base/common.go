package base

import (
	"fmt"
	"strings"
)

// Name represents a GraphQL name
type Name struct {
	Value    string
	Location *Location
}

func (n *Name) String() string {
	return n.Value
}

func (n *Name) GetLocation() *Location {
	return n.Location
}

// OperationType represents the type of GraphQL operation
type OperationType string

const (
	Query        OperationType = "query"
	Mutation     OperationType = "mutation"
	Subscription OperationType = "subscription"
)

// Argument represents a field or directive argument
type Argument struct {
	Name     *Name
	Value    Value
	Location *Location
}

func (a *Argument) String() string {
	return fmt.Sprintf("%s: %s", a.Name.Value, a.Value.String())
}

func (a *Argument) GetLocation() *Location {
	return a.Location
}

// Directive represents a GraphQL directive
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

func (d *Directive) GetLocation() *Location {
	return d.Location
}