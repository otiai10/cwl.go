package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_recursive_input_directory(t *testing.T) {
	f := load("recursive-input-directory.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, len(root.Requirements)).ToBe(2)
	Expect(t, root.Requirements[0].Class).ToBe("InitialWorkDirRequirement")
	// Expect(t, root.Requirements[1].Listing[0].Entry).ToBe("$(inputs.input_dir)")
	// Expect(t, root.Requirements[1].Listing[0].EntryName).ToBe("work_dir")
	// Expect(t, root.Requirements[1].Listing[0].Writable).ToBe(true)
	Expect(t, root.Requirements[1].Class).ToBe("ShellCommandRequirement")
	Expect(t, root.Stdout).ToBe("output.txt")
	Expect(t, root.Arguments[0].Binding.ShellQuote).ToBe(false)
	Expect(t, root.Arguments[0].Binding.ValueFrom.Key()).ToBe(`touch work_dir/e;
if [ ! -w work_dir ]; then echo work_dir not writable; fi;
if [ -L work_dir ]; then echo work_dir is a symlink; fi;
if [ ! -w work_dir/a ]; then echo work_dir/a not writable; fi;
if [ -L work_dir/a ]; then echo work_dir/a is a symlink; fi;
if [ ! -w work_dir/c ]; then echo work_dir/c not writable; fi;
if [ -L work_dir/c ]; then echo work_dir/c is a symlink; fi;
if [ ! -w work_dir/c/d ]; then echo work_dir/c/d not writable; fi;
if [ -L work_dir/c/d ]; then echo work_dir/c/d is a symlink; fi;
if [ ! -w work_dir/e ]; then echo work_dir/e not writable; fi;
if [ -L work_dir/e ]; then echo work_dir/e is a symlink ; fi;
`)
	Expect(t, root.Inputs[0].ID).ToBe("input_dir")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("Directory")
	Expect(t, len(root.Outputs)).ToBe(2)
	Expect(t, root.Outputs[0].ID).ToBe("output_dir")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("Directory")
	Expect(t, root.Outputs[0].Binding.Glob[0]).ToBe("work_dir")
	Expect(t, root.Outputs[1].ID).ToBe("test_result")
	Expect(t, root.Outputs[1].Types[0].Type).ToBe("stdout")
}
