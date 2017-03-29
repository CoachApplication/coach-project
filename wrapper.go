package project

type Wrapper interface {
	Name() string
	Label() string
	Env() map[string]string
}
