package cwl

// StringArrayable converts "xxx" to ["xxx"] if it's not slice.
func StringArrayable(i interface{}) []string {
	dest := []string{}
	switch x := i.(type) {
	case []interface{}:
		for _, s := range x {
			dest = append(dest, s.(string))
		}
	case string:
		dest = append(dest, x)
	}
	return dest
}
