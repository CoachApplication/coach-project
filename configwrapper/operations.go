package configwrapper

import (
	"github.com/CoachApplication/api"
	"github.com/CoachApplication/base"
	"github.com/CoachApplication/config"
	"github.com/CoachApplication/project"
)

func MakeOperations(wr config.Wrapper) api.Operations {
	ops := base.NewOperations()

	base := project.NewFactoryOperationBase(NewFactory(wr))

	ops.Add(project.NewNameOperation(*base).Operation())
	ops.Add(project.NewLabelOperation(*base).Operation())

	ops.Add(project.NewEnvironmentGetOperation(*base).Operation())
	ops.Add(project.NewEnvironmentListOperation(*base).Operation())
	ops.Add(project.NewEnvironmentMapOperation(*base).Operation())

	return ops.Operations()
}
