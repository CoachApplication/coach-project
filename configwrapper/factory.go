package configwrapper

import (
	"github.com/CoachApplication/config"
	"github.com/CoachApplication/project"
)

const (
	CONFIG_ID_PROJECT = "project"
)

type Factory struct {
	wr config.Wrapper
	p  ConfigProject
}

func NewFactory(wr config.Wrapper) *Factory {
	return &Factory{
		wr: wr,
	}
}

func (f *Factory) Project() project.Project {
	if &f.p == nil {
		f.p = ConfigProject{
			N: "default",
			L: "Default",
			E: map[string]string{},
		}
		if sc, err := f.wr.Get(CONFIG_ID_PROJECT); err == nil {
			for _, s := range sc.Order() {
				c, _ := sc.Get(s)

				res := c.HasValue()
				<-res.Finished()

				if res.Success() {
					var m ConfigProject
					c.Get(&m)
					f.p.Merge(m)
				}
			}
		}
	}
	return f.p
}
