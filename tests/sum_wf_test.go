package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_sum_wf(t *testing.T) {
	f := load("sum-wf.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(err)

	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("Workflow")
	Expect(t, root.Requirements[0].Class).ToBe("StepInputExpressionRequirement")
	Expect(t, root.Requirements[1].Class).ToBe("MultipleInputFeatureRequirement")
	Expect(t, root.Requirements[2].Class).ToBe("InlineJavascriptRequirement")
	Expect(t, len(root.Inputs)).ToBe(2)
	Expect(t, root.Inputs[0].ID).ToBe("int_1")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("int")
	Expect(t, root.Inputs[0].Types[1].Type).ToBe("string")
	Expect(t, root.Inputs[1].ID).ToBe("int_2")
	Expect(t, root.Inputs[1].Types[0].Type).ToBe("int")
	Expect(t, root.Inputs[1].Types[1].Type).ToBe("string")
	Expect(t, root.Outputs[0].ID).ToBe("result")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("int")
	Expect(t, root.Outputs[0].Source[0]).ToBe("sum/result")
	Expect(t, root.Steps[0].ID).ToBe("sum")
	Expect(t, root.Steps[0].In[0].ID).ToBe("data")
	Expect(t, root.Steps[0].In[0].Source[0]).ToBe("int_1")
	Expect(t, root.Steps[0].In[0].Source[1]).ToBe("int_2")
	Expect(t, root.Steps[0].In[0].ValueFrom).ToBe(`${
  var sum = 0;
  for (var i = 0; i < self.length; i++){
    sum += self[i];
  };
  return sum;
}
`)
	Expect(t, root.Steps[0].Out[0].ID).ToBe("result")
	Expect(t, root.Steps[0].Run.Workflow.Class).ToBe("ExpressionTool")
	Expect(t, root.Steps[0].Run.Workflow.Inputs[0].ID).ToBe("data")
	Expect(t, root.Steps[0].Run.Workflow.Inputs[0].Types[0].Type).ToBe("int")
	Expect(t, root.Steps[0].Run.Workflow.Outputs[0].ID).ToBe("result")
	Expect(t, root.Steps[0].Run.Workflow.Outputs[0].Types[0].Type).ToBe("int")
	Expect(t, root.Steps[0].Run.Workflow.Expression).ToBe(`${
  return {"result": inputs.data};
}
`)
}
