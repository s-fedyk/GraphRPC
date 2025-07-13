package main

import (
	"fmt"
	"os"

	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <graphql-file>")
		fmt.Println("Example: go run main.go parser/test_queries.graphql")
		os.Exit(1)
	}

	filepath := os.Args[1]

	content, err := os.ReadFile(filepath)
	if err != nil {
		panic("Cannot parse file")
	}

	source := ast.Source{
		Name:  "test",
		Input: string(content),
	}

	schema, err := gqlparser.LoadSchema(&source)
	if err != nil {
		panic("Cannot parse schema")
	}

	fmt.Printf("schema: %v\n", schema)
}
