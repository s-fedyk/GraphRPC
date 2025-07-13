package main

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"graphrpc.com/ast"
	parser "graphrpc.com/parser/gen"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <graphql-file>")
		fmt.Println("Example: go run main.go parser/test_queries.graphql")
		os.Exit(1)
	}

	filename := os.Args[1]

	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}

	input := antlr.NewInputStream(string(content))
	lexer := parser.NewGraphQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewGraphQLParser(stream)

	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	tree := p.Document()

	var builder = ast.NewASTBuilder()
	tree.Accept(builder)
}
