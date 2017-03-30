package project

/**
 * Create a Factory/Project pair that work from values assigned in the constructor.
 *
 * This gives a way of providing a static Project in cases where there is no valid source of project details, or
 * for testing.
 */

// BufferedFactory an SSA Factory that creates a Project based on passed values
type BufferedFactory struct {
	p Project
}

func NewBufferedFactory(name, label string, env map[string]string) *BufferedFactory {
	p := NewBufferedProject(name, label, env).Project()
	return &BufferedFactory{
		p: p,
	}
}

func (bf *BufferedFactory) Factory() Factory {
	return Factory(bf)
}

func (bf *BufferedFactory) Project() Project {
	return bf.p
}

// BUfferedProject an SSA Project that returned Project properties passed in during creation
type BufferedProject struct {
	name  string
	label string
	env   map[string]string
}

func NewBufferedProject(name, label string, env map[string]string) *BufferedProject {
	return &BufferedProject{
		name:  name,
		label: label,
		env:   env,
	}
}

func (bp *BufferedProject) Project() Project {
	return Project(bp)
}

func (bp *BufferedProject) Name() string {
	return bp.name
}

func (bp *BufferedProject) Label() string {
	return bp.label
}

func (bp *BufferedProject) Env() map[string]string {
	return bp.env
}
