package project_test

import (
	"github.com/CoachApplication/project"
	"testing"
)

func TestNameOperation_Exec(t *testing.T) {
	b := project.NewFactoryOperationBase(MakeTestFactory(t))

	no := project.NewNameOperation(*b)

	res := no.Exec(no.Properties())
	<-res.Finished()

	if !res.Success() {
		t.Error("NameOperation failed to exec: ", res.Errors())
	} else if p, err := res.Properties().Get(project.PROPERTY_ID_NAME); err != nil {
		t.Error("NameOperation returned no project name Property")
	} else if n, good := p.Get().(string); !good {
		t.Error("NameOperation returned an invalid project name")
	} else if n != TESTPROJECT_NAME {
		t.Error("NameOperation returned the wrong project name")
	}
}
