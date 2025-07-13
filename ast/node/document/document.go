package document

import (
	"strings"
	"graphrpc.com/ast/node/base"
)

// Document represents the root of a GraphQL document
type Document struct {
	Definitions []base.Definition
	Location    *base.Location
}

func (d *Document) String() string {
	var parts []string
	for _, def := range d.Definitions {
		parts = append(parts, def.String())
	}
	return strings.Join(parts, "\n\n")
}

func (d *Document) GetLocation() *base.Location {
	return d.Location
}