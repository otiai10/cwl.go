package cwl

// Binding represents and combines "CommandLineBinding" and "CommandOutputBinding"
// @see
// - http://www.commonwl.org/v1.0/CommandLineTool.html#CommandLineBinding
// - http://www.commonwl.org/v1.0/CommandLineTool.html#CommandOutputBinding
type Binding struct {
	// Common
	LoadContents bool
	// CommandLineBinding
	Position   int    `json:"position"`
	Prefix     string `json:"prefix"`
	Separate   bool   `json:"separate"`
	Separator  string `json:"separator"`
	ShellQuote bool   `json:"shellQuote"`
	ValueFrom  string `json:"valueFrom"`
	// CommandOutputBinding
	Glob []string `json:"glob"`
	Eval string   `json:"eval"`
}

// New constructs new "Binding".
func (binding Binding) New(i interface{}) *Binding {
	dest := new(Binding)
	switch x := i.(type) {
	case map[string]interface{}:
		for key, v := range x {
			switch key {
			case "position":
				dest.Position = int(v.(float64))
			case "prefix":
				dest.Prefix = v.(string)
			case "itemSeparator":
				dest.Separator = v.(string)
			case "loadContents":
				dest.LoadContents = v.(bool)
			case "glob":
				dest.Glob = StringArrayable(v)
			}
		}
	}
	return dest
}
