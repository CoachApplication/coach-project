package project_test

import (
	"github.com/CoachApplication/project"
	"testing"
)

func Test_TestFactory_Project(t *testing.T) {
	f := MakeTestFactory(t)

	p := f.Project()
	if p == nil {
		t.Error("TestFactory returned a null property")
	} else if p.Name() != TESTPROJECT_NAME {
		t.Error("TestFactory Project has the wrong name")
	}
}

type TestFactory struct {
	t *testing.T
}

func MakeTestFactory(t *testing.T) project.Factory {
	return (&TestFactory{t: t}).Factory()
}

func (tf *TestFactory) Factory() project.Factory {
	return project.Factory(tf)
}

func (tf *TestFactory) Project() project.Project {
	return MakeTestProject(tf.t)
}
