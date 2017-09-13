package cwl

// SecondaryFile represents an element of "secondaryFiles".
type SecondaryFile struct {
	Entry string
}

// NewList constructs list of "SecondaryFile".
func (_ SecondaryFile) NewList(i interface{}) []SecondaryFile {
	dest := []SecondaryFile{}
	switch x := i.(type) {
	case []interface{}:
		for _, v := range x {
			dest = append(dest, SecondaryFile{Entry: v.(string)})
		}
	}
	return dest
}
