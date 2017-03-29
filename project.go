package project

type Project interface {
	Name() string
	Label() string
	Env() map[string]string
}
