package cwl

// Provided represents the provided input value
// by parameter files.
type Provided struct {
	Raw interface{}
}

// New constructs new "Provided" struct.
func (provided Provided) New(i interface{}) Provided {
	dest := Provided{Raw: i}
	return dest
}
