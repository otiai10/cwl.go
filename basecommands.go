package cwl

// BaseCommands represents "baseCommand" of CWL CommandLineTool.
// @see http://www.commonwl.org/v1.0/CommandLineTool.html#CommandLineTool
type BaseCommands []string

// New constructs "BaseCommands" struct.
func (baseCommands BaseCommands) New(i interface{}) BaseCommands {
	return StringArrayable(i)
}
