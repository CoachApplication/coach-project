package project_test

import (
	"github.com/CoachApplication/project"
	"testing"
)

func MakeTestBufferedFactory(t *testing.T) project.Factory {
	return project.NewBufferedFactory(
		"test",
		"Test project",
		map[string]string{
			"one":   "1",
			"two":   "2",
			"three": "3",
		},
	).Factory()
}

func TestBufferedFactory_Project(t *testing.T) {
	f := MakeTestBufferedFactory(t)

	p := f.Project()
	if p == nil {
		t.Error("BufferedFactory created a nil Project.")
	}
}

func TestBufferedProject_Name(t *testing.T) {
	f := MakeTestBufferedFactory(t)
	p := f.Project()
	n := p.Name()

	if n == "" {
		t.Error("BufferedProject gave an empty name.")
	} else if n != "test" {
		t.Error("BufferedProject gave the wrong name: ", n)
	}
}

func TestBufferedProject_Label(t *testing.T) {
	f := MakeTestBufferedFactory(t)
	p := f.Project()
	l := p.Label()

	if l == "" {
		t.Error("BufferedProject gave an empty label.")
	} else if l != "Test project" {
		t.Error("BufferedProject gave the wrong label: ", l)
	}
}

func TestBufferedProject_Env(t *testing.T) {
	f := MakeTestBufferedFactory(t)
	p := f.Project()
	e := p.Env()

	if len(e) == 0 {
		t.Error("BufferedProject gave an empty Env map.")
	} else if v, good := e["one"]; !good {
		t.Error("BufferedProject env is missing a valid key")
	} else if v != "1" {
		t.Error("BufferedProject env gave the wrong value for a valid key: ", v)
	}
}
