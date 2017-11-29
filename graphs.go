package cwl

// Graphs represents "$graph" field in CWL.
type Graphs []*Root

// Graph represents an element of "steps"
type Graph struct {
	Run *Root
}

// New constructs "Graphs" from interface.
func (_ Graphs) New(i interface{}) Graphs {
	dest := Graphs{}
	switch x := i.(type) {
	case []interface{}:
		for _, v := range x {
			g := new(Root)
			g = g.AsStep(v)
			dest = append(dest, g)
		}
	}
	return dest
}

// Len for sorting
func (g Graphs) Len() int {
	return len(g)
}

// Less for sorting
func (g Graphs) Less(i, j int) bool {
	return g[i].ID < g[j].ID
}

// Swap for sorting
func (g Graphs) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}
