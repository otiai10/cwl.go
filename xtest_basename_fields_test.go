package cwl

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestCWL_basename_fields_test_test(t *testing.T) {
	f := cwl("basename-fields-test.cwl")
	root := NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)

	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("Workflow")

	Expect(t, root.Requirements[0].Class).ToBe("StepInputExpressionRequirement")
	Expect(t, root.Requirements[0])
}
