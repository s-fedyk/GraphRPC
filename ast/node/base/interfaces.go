package base

// Node is the base interface for all AST nodes
type Node interface {
	String() string
	GetLocation() *Location
}

// Location represents the source location of a node
type Location struct {
	Line   int
	Column int
}

// Definition is the base interface for all top-level definitions
type Definition interface {
	Node
	definitionNode()
}

// ExecutableDefinition represents executable definitions (operations, fragments)
type ExecutableDefinition interface {
	Definition
	executableDefinitionNode()
}

// TypeSystemDefinition represents type system definitions (schema, types, directives)
type TypeSystemDefinition interface {
	Definition
	typeSystemDefinitionNode()
}

// Selection represents a selection in a selection set
type Selection interface {
	Node
	selectionNode()
}

// Value represents a GraphQL value
type Value interface {
	Node
	valueNode()
}

// Type represents a GraphQL type reference
type Type interface {
	Node
	typeNode()
}