package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

var q1 = `
{
	al(xid: "alice") {
		status
		_xid_
		follows {
			status
			_xid_
			follows {
				status
				_xid_
				follows {
					_xid_
					status
				}
			}
		}
		status
		_xid_
	}
}
`

var q2 = `query queryName {
		me(uid : "0x0a") {
			friends {
				name
			}
			gender,age
			hometown
		}
	}
`

func TestQueryParse(t *testing.T) {
	input := antlr.NewInputStream(q2)
	lexer := NewGraphQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewGraphQLParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	_ = p.Document()
	// antlr.ParseTreeWalkerDefault.Walk(newMyListener(), tree)
}

func runParser(q string, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := antlr.NewInputStream(q)
		lexer := NewGraphQLLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := NewGraphQLParser(stream)
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		p.BuildParseTrees = true
		// uptill here we have a cost of : 19000 for q1
		// next call makes it 100 times more costly to : 1800000
		_ = p.Document()
	}
}

func BenchmarkQuery(b *testing.B) {
	b.Run("q1", func(b *testing.B) { runParser(q1, b) })
	b.Run("q2", func(b *testing.B) { runParser(q2, b) })
}
