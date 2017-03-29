package configwrapper

import (
	api "github.com/CoachApplication/api"
	base "github.com/CoachApplication/base"
	config "github.com/CoachApplication/config"
)

func MakeOperations(cw config.Wrapper) api.Operations {
	ops := base.NewOperations()

	return ops.Operations()
}
