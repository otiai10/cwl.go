package cwl

// Requirements represents "requirements" field in CWL.
type Requirements []Requirement

// New constructs "Requirements" struct from interface.
func (requirements Requirements) New(i interface{}) Requirements {
	dest := Requirements{}
	switch x := i.(type) {
	case []interface{}:
		for _, r := range x {
			dest = append(dest, Requirement{}.New(r))
		}
	case map[string]interface{}:
		for key, v := range x {
			r := Requirement{}.New(v)
			r.Class = key
			dest = append(dest, r)
		}
	}
	return dest
}

// Requirement represent an element of "requirements".
type Requirement struct {
	Class         string
	Types         []RequirementType
	ExpressionLib []RequirementExpression // For InlineJavascriptRequirement
	Listing       []RequirementEntry      // For InitialWorkDirRequirement
	EnvDef        []RequirementEnvDef      // For EnvVarRequirement
}

// New constructs "Requirement" struct from interface.
func (requirement Requirement) New(i interface{}) Requirement {
	dest := Requirement{}
	switch x := i.(type) {
	case map[string]interface{}:
		for key, val := range x {
			switch key {
			case "class":
				dest.Class = val.(string)
			case "types":
				dest.Types = RequirementType{}.NewList(val)
			case "expressionLib":
				dest.ExpressionLib = RequirementExpression{}.NewList(val)
			case "envDef":
				dest.EnvDef = RequirementEnvDef{}.NewList(val)
			case "listing":
				dest.Listing = RequirementEntry{}.NewList(val)
			}
		}
	}
	return dest
}

// RequirementType represents "types" field in "requirements".
type RequirementType struct {
	Name   string
	Type   string
	Fields []RequirementTypeField
}

// NewList constructs a list of "RequirementType" from interface.
func (typ RequirementType) NewList(i interface{}) []RequirementType {
	dest := []RequirementType{}
	switch x := i.(type) {
	case []interface{}:
		for _, v := range x {
			dest = append(dest, RequirementType{}.New(v))
		}
	}
	return dest
}

// New constructs a "RequirementType" from interface.
func (typ RequirementType) New(i interface{}) RequirementType {
	dest := RequirementType{}
	switch x := i.(type) {
	case map[string]interface{}:
		for key, val := range x {
			switch key {
			case "name":
				dest.Name = val.(string)
			case "type":
				dest.Type = val.(string)
			case "fields":
				dest.Fields = RequirementTypeField{}.NewList(val)
			}
		}
	}
	return dest
}

// RequirementTypeField represents an element of "requirements[0].types[0].fields"
type RequirementTypeField struct {
	Name    string
	Types   []RequirementTypeFieldType  // WTF!
	Binding RequirementTypeFieldBinding // WTF!!
}

// NewList constructs a list of "RequirementTypeField" from interface.
func (field RequirementTypeField) NewList(i interface{}) []RequirementTypeField {
	dest := []RequirementTypeField{}
	switch x := i.(type) {
	case []interface{}:
		for _, v := range x {
			dest = append(dest, RequirementTypeField{}.New(v))
		}
	}
	return dest
}

// New constructs a "RequirementTypeField" from interface.
func (field RequirementTypeField) New(i interface{}) RequirementTypeField {
	dest := RequirementTypeField{}
	switch x := i.(type) {
	case map[string]interface{}:
		for key, val := range x {
			switch key {
			case "name":
				dest.Name = val.(string)
			case "type":
				dest.Types = RequirementTypeFieldType{}.NewList(val)
			case "inputBinding":
				dest.Binding = RequirementTypeFieldBinding{}.New(val)
			}
		}
	}
	return dest
}

// RequirementTypeFieldType represents "requirements[0].types[0].fields[0].type"
type RequirementTypeFieldType struct {
	Type    string
	Name    string
	Symbols []string
	Items   []string
}

// NewList constructs a list of "RequirementTypeFieldType" from interface.
func (rtft RequirementTypeFieldType) NewList(i interface{}) []RequirementTypeFieldType {
	dest := []RequirementTypeFieldType{}
	switch x := i.(type) {
	case map[string]interface{}:
		t := RequirementTypeFieldType{}
		for key, v := range x {
			switch key {
			case "type":
				t.Type = v.(string)
			case "name":
				t.Name = v.(string)
			case "symbols":
				if list, ok := v.([]interface{}); ok {
					for _, s := range list {
						t.Symbols = append(t.Symbols, s.(string))
					}
				}
			case "items":
				if list, ok := v.([]interface{}); ok {
					for _, s := range list {
						t.Items = append(t.Items, s.(string))
					}
				}
			}
		}
		dest = append(dest, t)
	case []interface{}:
		for _, v := range x {
			dest = append(dest, RequirementTypeFieldType{Type: v.(string)})
		}
	}
	return dest
}

// RequirementTypeFieldBinding represents "requirements[0].types[0].fields[0].inputBinding"
type RequirementTypeFieldBinding struct {
	Prefix   string
	Position int
	Separate bool
}

// New constructs RequirementTypeFieldBinding from interface.
func (rtfb RequirementTypeFieldBinding) New(i interface{}) RequirementTypeFieldBinding {
	dest := RequirementTypeFieldBinding{}
	switch x := i.(type) {
	case map[string]interface{}:
		for key, val := range x {
			switch key {
			case "position":
				dest.Position = int(val.(float64))
			case "prefix":
				dest.Prefix = val.(string)
			case "separate":
				dest.Separate = val.(bool)
			}
		}
	}
	return dest
}

// RequirementExpression only appears if requirement class is "InlineJavascriptRequirement"
type RequirementExpression struct {
	Include string
	Execute string
}

// NewList constructs a list of RequirementExpression from interface.
func (expr RequirementExpression) NewList(i interface{}) []RequirementExpression {
	dest := []RequirementExpression{}
	switch x := i.(type) {
	case []interface{}:
		for _, v := range x {
			switch e := v.(type) {
			case map[string]interface{}:
				exp := RequirementExpression{}
				if include, ok := e["$include"]; ok {
					exp.Include = include.(string)
				}
				dest = append(dest, exp)
			case string:
				dest = append(dest, RequirementExpression{Execute: e})
			}
		}
	}
	return dest
}

// RequirementEntry only appears if requirement class is "InitialWorkDirRequirement"
type RequirementEntry struct {
	Name  string // WTF naming!?
	Entry string // WTF naming!?
}

// NewList constructs a list of RequirementEntry from interface
func (list RequirementEntry) NewList(i interface{}) []RequirementEntry {
	dest := []RequirementEntry{}
	switch x := i.(type) {
	case []interface{}:
		for _, v := range x {
			dest = append(dest, RequirementEntry{}.New(v))
		}
	}
	return dest
}

// New constructs a RequirementEntry from interface
func (list RequirementEntry) New(i interface{}) RequirementEntry {
	dest := RequirementEntry{}
	switch x := i.(type) {
	case map[string]interface{}:
		for key, v := range x {
			switch key {
			case "entryname":
				dest.Name = v.(string)
			case "entry":
				dest.Entry = v.(string)
			}
		}
	}
	return dest
}
