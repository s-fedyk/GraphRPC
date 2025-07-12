package ast

import (
	"fmt"
	parser "graphrpc.com/parser/gen"
)

type ASTBuilder struct {
	*parser.BaseGraphQLVisitor
}

func NewASTBuilder() *ASTBuilder {
	return &ASTBuilder{
		BaseGraphQLVisitor: &parser.BaseGraphQLVisitor{},
	}
}

func (b *ASTBuilder) BuildDocument(ctx parser.IDocumentContext) *Document {
	fmt.Printf("Hello world!")
	return &Document{}
}
