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
	}
	return dest
}

// Requirement represent an element of "requirements".
type Requirement struct {
	Class string
	Types []RequirementType
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
