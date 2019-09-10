package evaluate

import (
	"github.com/dsuarez4/fgrep/tokenizer"
	"fmt"
)

type ErrUnsupportedType struct {
	Attribute 	string
	Value 		interface{}
}

func (e *ErrUnsupportedType) Error() string {
	return fmt.Sprintf("unsupported type %T for attribute %s",
		e.Value, e.Attribute)
}

type ErrUnsupportedOperator struct {
	Attribute string
	Operator tokenizer.TokenType
}

func (e *ErrUnsupportedOperator) Error() string {
	return fmt.Sprintf("unsupported type %T for attribute %s",
		e.Operator.String(), e.Attribute)
}
