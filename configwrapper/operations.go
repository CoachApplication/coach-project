package configwrapper

import (
	"github.com/CoachApplication/api"
	"github.com/CoachApplication/base"
	"github.com/CoachApplication/config"
	"github.com/CoachApplication/project"
)

/**
 * MakeOperations Make a set of Project Operations based on a config wrapper.
 *
 * The config wrapper is used to create a new Factory, which is added to a base Operation.  The base is then used
 * for a series of /project Operation structs.
 *
 * @note this is not necessarily used, but is an example of how to add all of the operations when you have a
 * config wrapper.
 */
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
