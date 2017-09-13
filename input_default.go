package cwl

import "reflect"

// InputDefault represents "default" field in an element of "inputs".
type InputDefault struct {
	Self interface{}
	Kind reflect.Kind
}

// New constructs new "InputDefault".
func (_ InputDefault) New(i interface{}) *InputDefault {
	dest := &InputDefault{Self: i, Kind: reflect.TypeOf(i).Kind()}
	return dest
}
