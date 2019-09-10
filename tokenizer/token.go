package tokenizer

import "fmt"

type TokenType int8

const (

	Unknown TokenType = iota

	Identifier
	Subquery

	Select
	From
	Where

	As
	Or
	And
	Not

	Equals
	NotEquals
)

func (t TokenType) String() string {
	switch t {
	case Identifier:
		return "identifier"
	case Subquery:
		return "subquery"
	case Select:
		return "select"
	case From:
		return "from"
	case As:
		return "as"
	case Where:
		return "where"
	case Or:
		return "or"
	case And:
		return "and"
	case Not:
		return "not"

	// Not finished-
	case Equals:
		return "equal"
	case NotEquals:
		return "not-equal"

	default:
		return "unknown"


	}
}

// Represents single token
type Token struct {
	Type TokenType
	Raw string
}

func (t *Token) String() string {
	return fmt.Sprintf("{type: %s, raw: \"%s\"}",
		t.Type.String(), t.Raw)
}

type Tokenizer struct {
	input []rune
	tokens []*Token
}

func NewTokenizer(input string) *Tokenizer {
	return &Tokenizer {
		input: []rune(input),
		tokens: make([]*Token, 0),
	}
}



