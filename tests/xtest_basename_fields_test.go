package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestCWL_basename_fields_test_test(t *testing.T) {
	f := load("basename-fields-test.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)

	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("Workflow")

	Expect(t, root.Requirements[0].Class).ToBe("StepInputExpressionRequirement")
	Expect(t, root.Requirements[0])
}
