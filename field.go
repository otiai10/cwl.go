package cwl

// Field represents CommandInputRecordField
// @see http://www.commonwl.org/v1.0/CommandLineTool.html#CommandInputRecordField
type Field struct {
	Name    string
	Doc     string
	Types   []Type
	Binding *Binding
	Label   string
}
