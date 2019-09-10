package tokenizer

import "testing"

func TestToken_TokenTypeString(t *testing.T) {
	type Case struct {
		tt TokenType
		expected string
	}

	cases := []Case {
		{tt: Identifier, expected: "identifier"},
		{tt: Subquery, expected: "subquery"},
		{tt: From, expected: "from"},
		{tt: As, expected: "as"},
		{tt: Where, expected: "where"},
		{tt: Or, expected: "or"},
		{tt: And, expected: "and"},
		{tt: Not, expected: "not"},
	}

	for _, c := range cases {
		actual := c.tt.String()
		if c.expected != actual {
			t.Fatalf("\nExpected %v\n	Got %v", c.expected, actual)
		}
	}
}

func TestToken_String(t *testing.T) {
	input := "select all FROM ."

	tokenLength := 0

	tokenizer := NewTokenizer(input)

	if len(tokenizer.tokens) != tokenLength {
		t.Fatalf("len(tokenizer.tokens)\nExpected %v\n     Got %v", tokenLength,
			len(tokenizer.tokens))
	}
}