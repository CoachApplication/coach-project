package project_test

import (
	"github.com/CoachApplication/project"
	"testing"
)

func TestLabelOperation_Exec(t *testing.T) {
	b := project.NewFactoryOperationBase(MakeTestFactory(t))

	lo := project.NewLabelOperation(*b)

	res := lo.Exec(lo.Properties())
	<-res.Finished()

	if !res.Success() {
		t.Error("LabelOperation failed to exec: ", res.Errors())
	} else if p, err := res.Properties().Get(project.PROPERTY_ID_LABEL); err != nil {
		t.Error("LabelOperation returned no project label Property")
	} else if l, good := p.Get().(string); !good {
		t.Error("LabelOperation returned an invalid project label")
	} else if l != TESTPROJECT_LABEL {
		t.Error("LabelOperation returned the wrong project label: ", l)
	}
}
