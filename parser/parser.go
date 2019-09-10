package parser

import (
	"github.com/dsuarez4/fgrep/tokenizer"
	"github.com/dsuarez4/fgrep/query"
)

type parser struct {
	tokenizer *tokenizer.Tokenizer
	current *tokenizer.Token
	expected tokenizer.TokenType
}

func (p *parser) parse(input string) (*query.Query, error) {

}