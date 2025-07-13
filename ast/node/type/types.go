package typ

import (
	"fmt"
	"graphrpc.com/ast/node/base"
)

// NamedType represents a named type reference
type NamedType struct {
	Name     *base.Name
	Location *base.Location
}

func (n *NamedType) String() string {
	return n.Name.Value
}

func (n *NamedType) GetLocation() *base.Location { return n.Location }
func (n *NamedType) typeNode()              {}

// ListType represents a list type
type ListType struct {
	Type     base.Type
	Location *base.Location
}

func (l *ListType) String() string {
	return fmt.Sprintf("[%s]", l.Type.String())
}

func (l *ListType) GetLocation() *base.Location { return l.Location }
func (l *ListType) typeNode()              {}

// NonNullType represents a non-null type
type NonNullType struct {
	Type     base.Type
	Location *base.Location
}

func (n *NonNullType) String() string {
	return fmt.Sprintf("%s!", n.Type.String())
}

func (n *NonNullType) GetLocation() *base.Location { return n.Location }
func (n *NonNullType) typeNode()              {}