package configwrapper

import (
	api "github.com/CoachApplication/coach-api"
	base "github.com/CoachApplication/coach-base"
	config "github.com/CoachApplication/coach-config"
)

func MakeOperations(cw config.Wrapper) api.Operations {
	ops := base.NewOperations()

	return ops.Operations()
}
