package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_template_tool(t *testing.T) {
	f := load("template-tool.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)

	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")

	Expect(t, root.Requirements[0].Class).ToBe("InlineJavascriptRequirement")
	Expect(t, root.Requirements[0].ExpressionLib[0].Value).ToBe("underscore.js")
	Expect(t, root.Requirements[0].ExpressionLib[1].Value).ToBe("var t = function(s) { return _.template(s)({'inputs': inputs}); };")

	Expect(t, root.Requirements[1].Class).ToBe("InitialWorkDirRequirement")
	Expect(t, root.Requirements[1].Listing[0].EntryName).ToBe("foo.txt")
	Expect(t, root.Requirements[1].Listing[0].Entry).ToBe(`$(t("The file is <%= inputs.file1.path.split('/').slice(-1)[0] %>\n"))`)

	Expect(t, root.Hints[0].Class).ToBe("DockerRequirement")
	Expect(t, root.Hints[0].DockerPull).ToBe("debian:8")

	Expect(t, root.Inputs[0].ID).ToBe("file1")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")

	Expect(t, root.Outputs[0].ID).ToBe("foo")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].Binding.Glob).ToBe([]string{"foo.txt"})

	Expect(t, root.BaseCommands[0]).ToBe("cat")
	Expect(t, root.BaseCommands[1]).ToBe("foo.txt")
}
