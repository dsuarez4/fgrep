package evaluate

import (
	"github.com/dsuarez4/fgrep/tokenizer"
	"os"
)

type Opts struct {
	Path 		string
	File 		os.FileInfo
	Attribute 	string
	Operator 	tokenizer.TokenType
	Value 		interface{}
}

type Modifier struct {
	Name string
	Arguments []string
}