package jsonindent

import (
	"encoding/json"
	"io"
)

const (
	// Prefix for json.Encoder.SetIndent
	Prefix = ""
	// Indent for json.Encoder.SetIndent
	Indent = "\t"
)

// NewEncoder ...
func NewEncoder(w io.Writer, options ...string) *json.Encoder {
	encoder := json.NewEncoder(w)
	options = append(options, "", "\t")
	encoder.SetIndent(options[0], options[1])
	return encoder
}
