generate: clean
	antlr4 -Dlanguage=Go -visitor -o gen GraphQL.g4 
clean:
	rm -rf gen/*
