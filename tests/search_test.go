package cwlgotest

import (
	"sort"
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_search(t *testing.T) {
	f := load("search.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(err)

	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Graphs[0].ID).ToBe("index")
	Expect(t, root.Graphs[0].Class).ToBe("CommandLineTool")
	Expect(t, root.Graphs[0].BaseCommands[0]).ToBe("python")
	Expect(t, root.Graphs[0].Arguments[0].Binding.ValueFrom.Key()).ToBe("input.txt")
	Expect(t, root.Graphs[0].Arguments[0].Binding.Position).ToBe(1)
	Expect(t, root.Graphs[0].Requirements[0].Class).ToBe("InitialWorkDirRequirement")
	Expect(t, root.Graphs[0].Requirements[0].Listing[0].EntryName).ToBe("input.txt")
	Expect(t, root.Graphs[0].Requirements[0].Listing[0].Entry).ToBe("$(inputs.file)")
	Expect(t, root.Graphs[0].Requirements[1].Class).ToBe("InlineJavascriptRequirement")
	Expect(t, root.Graphs[0].Hints[0].Class).ToBe("DockerRequirement")
	Expect(t, root.Graphs[0].Hints[0].DockerPull).ToBe("python:2-slim")
	// skip Default value check
	count := 0
	for _, input := range root.Graphs[0].Inputs {
		id := input.ID
		switch id {
		case "file":
			Expect(t, input.Types[0].Type).ToBe("File")
			count = count + 1
		case "secondfile":
			Expect(t, input.Types[0].Type).ToBe("File")
			count = count + 1
		case "index.py":
			Expect(t, input.Types[0].Type).ToBe("File")
			Expect(t, input.Binding.Position).ToBe(0)
			count = count + 1
		}
	}
	Expect(t, count).ToBe(3)
	Expect(t, root.Graphs[0].Outputs[0].ID).ToBe("result")
	Expect(t, root.Graphs[0].Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Graphs[0].Outputs[0].Binding.Glob[0]).ToBe("input.txt")
	Expect(t, root.Graphs[0].Outputs[0].SecondaryFiles[0].Entry).ToBe(".idx1")
	Expect(t, root.Graphs[0].Outputs[0].SecondaryFiles[1].Entry).ToBe("^.idx2")
	Expect(t, root.Graphs[0].Outputs[0].SecondaryFiles[2].Entry).ToBe(`$(self.basename).idx3`)
	Expect(t, root.Graphs[0].Outputs[0].SecondaryFiles[3].Entry).ToBe(`${ return self.basename+".idx4"; }`)
	Expect(t, root.Graphs[0].Outputs[0].SecondaryFiles[4].Entry).ToBe(`$({"path": self.path+".idx5", "class": "File"})`)
	Expect(t, root.Graphs[0].Outputs[0].SecondaryFiles[5].Entry).ToBe(`$(self.nameroot).idx6$(self.nameext)`)
	Expect(t, root.Graphs[0].Outputs[0].SecondaryFiles[6].Entry).ToBe(`${ return [self.basename+".idx7", inputs.secondfile]; }`)
	Expect(t, root.Graphs[0].Outputs[0].SecondaryFiles[7].Entry).ToBe("_idx8")

	Expect(t, root.Graphs[1].ID).ToBe("search")
	Expect(t, root.Graphs[1].Class).ToBe("CommandLineTool")
	Expect(t, root.Graphs[1].BaseCommands[0]).ToBe("python")
	Expect(t, root.Graphs[1].Requirements[0].Class).ToBe("InlineJavascriptRequirement")
	Expect(t, root.Graphs[1].Hints[0].Class).ToBe("DockerRequirement")
	Expect(t, root.Graphs[1].Hints[0].DockerPull).ToBe("python:2-slim")
	Expect(t, root.Graphs[1].Inputs[0].ID).ToBe("file")
	Expect(t, root.Graphs[1].Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Graphs[1].Inputs[0].Binding.Position).ToBe(1)
	Expect(t, root.Graphs[1].Inputs[0].SecondaryFiles[0].Entry).ToBe(".idx1")
	Expect(t, root.Graphs[1].Inputs[0].SecondaryFiles[1].Entry).ToBe("^.idx2")
	Expect(t, root.Graphs[1].Inputs[0].SecondaryFiles[2].Entry).ToBe(`$(self.basename).idx3`)
	Expect(t, root.Graphs[1].Inputs[0].SecondaryFiles[3].Entry).ToBe(`${ return self.basename+".idx4"; }`)
	Expect(t, root.Graphs[1].Inputs[0].SecondaryFiles[4].Entry).ToBe(`$(self.nameroot).idx6$(self.nameext)`)
	Expect(t, root.Graphs[1].Inputs[0].SecondaryFiles[5].Entry).ToBe(`${ return [self.basename+".idx7"]; }`)
	Expect(t, root.Graphs[1].Inputs[0].SecondaryFiles[6].Entry).ToBe("_idx8")
	Expect(t, root.Graphs[1].Inputs[1].ID).ToBe("search.py")
	Expect(t, root.Graphs[1].Inputs[1].Types[0].Type).ToBe("File")
	Expect(t, root.Graphs[1].Inputs[1].Binding.Position).ToBe(0)
	Expect(t, root.Graphs[1].Inputs[2].ID).ToBe("term")
	Expect(t, root.Graphs[1].Inputs[2].Types[0].Type).ToBe("string")
	Expect(t, root.Graphs[1].Inputs[2].Binding.Position).ToBe(2)
	Expect(t, root.Graphs[1].Outputs[0].ID).ToBe("result")
	Expect(t, root.Graphs[1].Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Graphs[1].Outputs[0].Binding.Glob[0]).ToBe("result.txt")

	Expect(t, root.Graphs[2].ID).ToBe("main")
	Expect(t, root.Graphs[2].Class).ToBe("Workflow")
	count2 := 0
	for _, input := range root.Graphs[2].Inputs {
		id := input.ID
		switch id {
		case "infile":
			Expect(t, input.Types[0].Type).ToBe("File")
			count2 = count2 + 1
		case "secondfile":
			Expect(t, input.Types[0].Type).ToBe("File")
			count2 = count2 + 1
		case "term":
			Expect(t, input.Types[0].Type).ToBe("string")
			count2 = count2 + 1
		}
	}
	Expect(t, count2).ToBe(3)
	sort.Sort(root.Graphs[2].Outputs)
	Expect(t, root.Graphs[2].Outputs[0].ID).ToBe("indexedfile")
	Expect(t, root.Graphs[2].Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Graphs[2].Outputs[0].Source[0]).ToBe("index/result")
	Expect(t, root.Graphs[2].Outputs[1].ID).ToBe("outfile")
	Expect(t, root.Graphs[2].Outputs[1].Types[0].Type).ToBe("File")
	Expect(t, root.Graphs[2].Outputs[1].Source[0]).ToBe("search/result")
	Expect(t, root.Graphs[2].Steps[0].ID).ToBe("index")
	Expect(t, root.Graphs[2].Steps[0].Run.Value).ToBe("#index")
	Expect(t, root.Graphs[2].Steps[0].In[0].ID).ToBe("file")
	Expect(t, root.Graphs[2].Steps[0].In[0].Source[0]).ToBe("infile")
	Expect(t, root.Graphs[2].Steps[0].In[1].ID).ToBe("secondfile")
	Expect(t, root.Graphs[2].Steps[0].In[1].Source[0]).ToBe("secondfile")
	Expect(t, root.Graphs[2].Steps[0].Out[0].ID).ToBe("result")
	Expect(t, root.Graphs[2].Steps[1].ID).ToBe("search")
	Expect(t, root.Graphs[2].Steps[1].Run.Value).ToBe("#search")
	Expect(t, root.Graphs[2].Steps[1].In[0].ID).ToBe("file")
	Expect(t, root.Graphs[2].Steps[1].In[0].Source[0]).ToBe("index/result")
	Expect(t, root.Graphs[2].Steps[1].In[1].ID).ToBe("term")
	Expect(t, root.Graphs[2].Steps[1].In[1].Source[0]).ToBe("term")
	Expect(t, root.Graphs[2].Steps[1].Out[0].ID).ToBe("result")
}
