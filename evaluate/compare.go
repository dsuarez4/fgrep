package evaluate

import "github.com/dsuarez4/fgrep/tokenizer"

func cmpAlpha(o *Opts, a, b interface{}) (res bool, err error) {

	switch o.Operator {
	case tokenizer.Equals:
		res = a.(string) == b.(string)

	default:
		err = &ErrUnsupportedOperator{}
	}

	return res, err
}

// cmpNumeric performs numeric comparison on a and b.
func cmpNumeric(o *Opts, a, b interface{}) (res bool, err error) {
	switch o.Operator {
	case tokenizer.Equals:
		res = a.(int64) == b.(int64)
	default:
		err = &ErrUnsupportedOperator{o.Attribute, o.Operator}
	}
	return res, err
}