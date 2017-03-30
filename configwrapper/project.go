package configwrapper

import "github.com/CoachApplication/project"

// ConfigProject a Project object that
type ConfigProject struct {
	N string            `yaml:"name"`
	L string            `yaml:"label"`
	E map[string]string `yaml:"env"`
}

func (cp *ConfigProject) Project() project.Project {
	return project.Project(cp)
}

func (cp *ConfigProject) Merge(m ConfigProject) {
	if m.Name() != "" {
		cp.N = m.Name()
	}
	if m.Label() != "" {
		cp.L = m.Label()
	}
	for ek, ev := range m.Env() {
		cp.E[ek] = ev
	}
}

func (cp *ConfigProject) Name() string {
	return cp.N
}

func (cp *ConfigProject) Label() string {
	return cp.L
}

func (cp *ConfigProject) Env() map[string]string {
	return cp.E
}
