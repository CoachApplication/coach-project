package project

import (
	"github.com/CoachApplication/api"
	"github.com/CoachApplication/base"
)

type Factory interface {
	Project() Project
}

type FactoryOperationBase struct {
	fac Factory
}

func NewFactoryOperationBase(fac Factory) *FactoryOperationBase {
	return &FactoryOperationBase{
		fac: fac,
	}
}

func (fob *FactoryOperationBase) Factory() Factory {
	return fob.fac
}

func (fob *FactoryOperationBase) Validate(props api.Properties) api.Result {
	if fob.fac == nil {
		return base.MakeFailedResult()
	} else {
		return base.MakeSuccessfulResult()
	}
}
