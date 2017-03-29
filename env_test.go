package project_test

import (
	"github.com/CoachApplication/project"
	"testing"
)

func TestMapOperation_Exec(t *testing.T) {
	b := project.NewFactoryOperationBase(MakeTestFactory(t))

	mo := project.NewEnvironmentMapOperation(*b)

	res := mo.Exec(mo.Properties())
	<-res.Finished()

	if !res.Success() {
		t.Error("EnvironmentMapOperation failed to exec: ", res.Errors())
	} else if p, err := res.Properties().Get(project.PROPERTY_ID_ENV_MAP); err != nil {
		t.Error("EnvironmentMapOperation returned no project Map Property")
	} else if m, good := p.Get().(map[string]string); !good {
		t.Error("EnvironmentMapOperation returned an invalid project label")
	} else if one, good := m[TESTPROJECT_MAP_FIRSTKEY]; !good {
		t.Error("EnvironmentMapOperation missing valid key: ", TESTPROJECT_MAP_FIRSTKEY)
	} else if one != TESTPROJECT_MAP_FIRSTVAL {
		t.Error("EnvironmentMapOperation map returned the wrong value for a valid key: ", TESTPROJECT_MAP_FIRSTKEY, one)
	}
}

func TestMapListOperation_Exec(t *testing.T) {
	b := project.NewFactoryOperationBase(MakeTestFactory(t))

	lo := project.NewEnvironmentListOperation(*b)

	res := lo.Exec(lo.Properties())
	<-res.Finished()

	if !res.Success() {
		t.Error("EnvironmentListOperation failed to exec: ", res.Errors())
	} else if p, err := res.Properties().Get(project.PROPERTY_ID_ENV_KEYS); err != nil {
		t.Error("EnvironmentListOperation returned no project Map Keys Property")
	} else if l, good := p.Get().([]string); !good {
		t.Error("EnvironmentListOperation returned an invalid project keys list")
	} else if len(l) == 0 {
		t.Error("EnvironmentListOperation map key list is empty")
	} else if len(l) != 3 {
		t.Error("EnvironmentListOperation map key list has the wrong number of values:", l)
	} else if !(l[0] == TESTPROJECT_MAP_FIRSTKEY || l[1] != TESTPROJECT_MAP_FIRSTKEY || l[2] != TESTPROJECT_MAP_FIRSTKEY) {
		t.Error("EnvironmentListOperation map map key list was missng the first expected key: ", TESTPROJECT_MAP_FIRSTKEY, l)
	}
}

func TestMapGettOperation_Exec(t *testing.T) {
	b := project.NewFactoryOperationBase(MakeTestFactory(t))

	lo := project.NewEnvironmentGetOperation(*b)
	ps := lo.Properties()

	if kp, err := ps.Get(project.PROPERTY_ID_ENV_KEY); err != nil {
		t.Error("EnvironmentGetOperation did not provide Key property")
	} else if err := kp.Set(TESTPROJECT_MAP_FIRSTKEY); err != nil {
		t.Error("EnvironmentGetOperation Key operation returned an error in Set()")
	} else {
		res := lo.Exec(ps)
		<-res.Finished()

		if !res.Success() {
			t.Error("EnvironmentGetOperation failed to exec: ", res.Errors())
		} else if p, err := res.Properties().Get(project.PROPERTY_ID_ENV_VALUE); err != nil {
			t.Error("EnvironmentGetOperation returned no project Map Keys Property")
		} else if v, good := p.Get().(string); !good {
			t.Error("EnvironmentGetOperation returned an invalid project value")
		} else if v != TESTPROJECT_MAP_FIRSTVAL {
			t.Error("EnvironmentGetOperation returned the wrong value for expected key: ", TESTPROJECT_MAP_FIRSTKEY, v)
		}

	}
}
