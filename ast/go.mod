module graphrpc.com/ast

go 1.22.3

require (
	github.com/antlr4-go/antlr/v4 v4.13.1
	graphrpc.com/parser v0.0.0
)

replace graphrpc.com/parser => ../parser

require golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
