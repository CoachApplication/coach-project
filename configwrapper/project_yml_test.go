package configwrapper_test

import (
	"context"
	"github.com/CoachApplication/config"
	"github.com/CoachApplication/config/provider/buffered"
	"github.com/CoachApplication/config/provider/yaml"
	"github.com/CoachApplication/project/configwrapper"
	"testing"
	"time"
)

/**
 * TESTING: Here we test the ConfigProject by using a buffered connector connected to some inline string, and the
 * YamlConfig from the config/provider/yaml repository.  These tests only test whether or not the ConfigProject
 * properly marshals/unmarshalls from yml.
 */

var pTest configwrapper.ConfigProject = configwrapper.ConfigProject{
	N: "test",
	L: "Test Project",
	E: map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
	},
}

var pTestYaml []byte = []byte(`
name: test
label: Test Project

env:
  one: 1
  two: 2
  three: 3
`)

// Generate a YamlConfig with a buffered string connector, so that we can test the marshalling on it.
func MakeTestYamlConfig(t *testing.T, b []byte) config.Config {
	return yaml.NewConfig("key", "scope", buffered.NewSingle("key", "scope", b)).Config()
}

func TestProject_Get_Yaml(t *testing.T) {
	dur, _ := time.ParseDuration("2s")
	ctx, _ := context.WithTimeout(context.Background(), dur)
	c := MakeTestYamlConfig(t, pTestYaml)

	var p configwrapper.ConfigProject

	res := c.Get(&p)

	select {
	case <-res.Finished():

		if !res.Success() {
			t.Error("ConfigProject marshalling using Yaml config failed: ", res.Errors())
		} else if p.Name() != "test" {
			t.Error("ConfigProject marshalling using Yaml gave the wrong Name: ", p.Name())
		}

	case <-ctx.Done():
		t.Error("Config Get timed out: ", ctx.Err().Error())
	}
}

func TestProject_Set_Yaml(t *testing.T) {
	dur, _ := time.ParseDuration("2s")
	ctx, _ := context.WithTimeout(context.Background(), dur)
	c := MakeTestYamlConfig(t, []byte{})

	var p configwrapper.ConfigProject

	// first test it when it is empty
	res := c.Get(&p)

	select {
	case <-res.Finished():

		if !res.Success() {
			t.Error("ConfigProject marshalling using Yaml config failed: ", res.Errors())
		} else if p.Name() != "" {
			t.Error("ConfigProject marshalling using Yaml gave the a name when it is supposed to be empty: ", p.Name())
		}

	case <-ctx.Done():
		t.Error("Config Get timed out: ", ctx.Err().Error())
	}

	// now fill it and test it again.
	res = c.Set(pTest)

	select {
	case <-res.Finished():

		if !res.Success() {
			t.Error("ConfigProject marshalling using Yaml config failed: ", res.Errors())
		} else {

			res = c.Get(&p)

			select {
			case <-res.Finished():

				if !res.Success() {
					t.Error("ConfigProject marshalling using Yaml config failed: ", res.Errors())
				} else if p.Name() != "test" {
					t.Error("ConfigProject marshalling using Yaml gave the wrong Name: ", p.Name())
				}

			case <-ctx.Done():
				t.Error("Config Get timed out: ", ctx.Err().Error())
			}
		}

	case <-ctx.Done():
		t.Error("Config Get timed out: ", ctx.Err().Error())
	}
}
