package project

import (
	"github.com/CoachApplication/api"
	"github.com/CoachApplication/base"
)

const (
	OPERATION_ID_NAME = "project.name"
)

type NameOperation struct {
	FactoryOperationBase
}

func NewNameOperation(base FactoryOperationBase) *NameOperation {
	return &NameOperation{
		FactoryOperationBase: base,
	}
}

func (no *NameOperation) Operation() api.Operation {
	return api.Operation(no)
}

func (no *NameOperation) Id() string {
	return OPERATION_ID_NAME
}

func (no *NameOperation) Ui() api.Ui {
	return base.NewUi(
		no.Id(),
		"Project name",
		"Machine name for the project",
		"",
	)
}

func (no *NameOperation) Usage() api.Usage {
	return base.ReadonlyPropertyUsage{}
}

func (no *NameOperation) Properties() api.Properties {
	return base.NewProperties().Properties()
}

func (no *NameOperation) Validate(props api.Properties) api.Result {
	return no.FactoryOperationBase.Validate(props)
}

func (no *NameOperation) Exec(props api.Properties) api.Result {
	res := base.NewResult()

	go func(fac Factory) {
		p := fac.Project()
		n := p.Name()

		np := (&NameProperty{}).Property()
		np.Set(n)
		res.AddProperty(np)

		res.MarkSucceeded()
		res.MarkFinished()
	}(no.Factory())

	return res.Result()
}
