package cwl

import (
	"regexp"
	"strings"
)

// Eval represents [output]Eval
type Eval struct {
	Raw string
}

// ToJavaScriptString ,
// Because "outputEval" is **NOT** pure JavaScript!
// What The Fuck!!!!!!!!!!!!!!!
func (eval Eval) ToJavaScriptString() (string, error) {
	s := eval.Raw
	s = strings.TrimLeft(s, "$(")
	s = regexp.MustCompile("\\)$").ReplaceAllString(s, "")
	return s, nil
}
