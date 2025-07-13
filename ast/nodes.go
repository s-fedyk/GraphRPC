package ast

type GraphQLType int

const (
	StringType GraphQLType = iota
	IntType
	FloatType
	BooleanType
	IDType
	ObjectType
)

type Node struct {
	Name string
}

// contains typing information (everything on lhs)
type TypeInfo struct {
	Nullable bool
	Type     string
}

// A value is a type:value combination.
type ValueNode struct {
	Node
	TypeInfo TypeInfo
}

// A leaf node is collected off some object.
type LeafNode struct {
	ValueNode
}

type ListNode struct {
	ValueNode
	Contents ValueNode
}

// Every time we hit an ObjectNode, we can write a resolver for it.
type ObjectNode struct {
	ValueNode
	Children map[string]ValueNode
}

type QueryNode struct {
	Node
}

type MutationNode struct {
	Node
}

type DefinitionNode struct {
	Node
}

type Document struct {
	Operations []*Node
}
