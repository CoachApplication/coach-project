package project_test

import (
	"github.com/CoachApplication/project"
	"testing"
)

const (
	TESTPROJECT_NAME         = "name"
	TESTPROJECT_LABEL        = "Test Project"
	TESTPROJECT_MAP_FIRSTKEY = "one"
	TESTPROJECT_MAP_FIRSTVAL = "1e"
)

func MakeTestProject(t *testing.T) project.Project {
	return NewTestProject(
		t,
		TESTPROJECT_NAME,
		TESTPROJECT_LABEL,
		map[string]string{
			TESTPROJECT_MAP_FIRSTKEY: TESTPROJECT_MAP_FIRSTVAL,
			"two":   "2",
			"three": "3",
		},
	).Project()
}

func Test_TestProject_Name(t *testing.T) {
	p := MakeTestProject(t)

	if p.Name() != TESTPROJECT_NAME {
		t.Error("TestProject returned incorrect Name: ", p.Name())
	}
}

func Test_TestProject_Label(t *testing.T) {
	p := MakeTestProject(t)

	if p.Label() != TESTPROJECT_LABEL {
		t.Error("TestProject returned incorrect Label: ", p.Label())
	}
}
func Test_TestProject_Env(t *testing.T) {
	p := MakeTestProject(t)

	e := p.Env()

	if _, good := e["no"]; good {
		t.Error("TestProject Env returned a value on an invalid key")
	}

	if one, good := e[TESTPROJECT_MAP_FIRSTKEY]; !good {
		t.Error("TestProject Env is missing a valid value")
	} else if one != TESTPROJECT_MAP_FIRSTVAL {
		t.Error("TestProject returned incorrect env value")
	}
}

type TestProject struct {
	t     *testing.T
	name  string
	label string
	env   map[string]string
}

func NewTestProject(t *testing.T, name, label string, env map[string]string) *TestProject {
	return &TestProject{
		t:     t,
		name:  name,
		label: label,
		env:   env,
	}
}

func (tp *TestProject) Project() project.Project {
	return project.Project(tp)
}

func (tp *TestProject) Name() string {
	return tp.name
}
func (tp *TestProject) Label() string {
	return tp.label
}
func (tp *TestProject) Env() map[string]string {
	return tp.env
}
