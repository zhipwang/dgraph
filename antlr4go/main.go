package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"parser"
)

type myListener struct {
	*parser.BaseGraphQLListener
}

func newMyListener() *myListener {
	return new(myListener)
}

func (this *myListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewGraphQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewGraphQLParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Document()
	antlr.ParseTreeWalkerDefault.Walk(newMyListener(), tree)
}
