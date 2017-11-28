package cwl

import "strings"

// Alias represents an expression with "$(xxx.yyy)"
type Alias struct {
	string
}

// Key extract the exact (and flattened) name of an expression.
func (a *Alias) Key() string {
	return strings.Trim(a.string, "$()")
}
