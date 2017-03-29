package project

import (
	"github.com/CoachApplication/api"
	"github.com/CoachApplication/base"
)

const (
	OPERATION_ID_LABEL = "project.label"
)

type LabelOperation struct {
	FactoryOperationBase
}

func NewLabelOperation(base FactoryOperationBase) *LabelOperation {
	return &LabelOperation{
		FactoryOperationBase: base,
	}
}

func (lo *LabelOperation) Operation() api.Operation {
	return api.Operation(lo)
}

func (lo *LabelOperation) Id() string {
	return OPERATION_ID_LABEL
}

func (lo *LabelOperation) Ui() api.Ui {
	return base.NewUi(
		lo.Id(),
		"Project label",
		"Interface label for the project",
		"",
	)
}

func (lo *LabelOperation) Usage() api.Usage {
	return base.ReadonlyPropertyUsage{}
}

func (lo *LabelOperation) Properties() api.Properties {
	return base.NewProperties().Properties()
}

func (lo *LabelOperation) Validate(props api.Properties) api.Result {
	return lo.FactoryOperationBase.Validate(props)
}

func (lo *LabelOperation) Exec(props api.Properties) api.Result {
	res := base.NewResult()

	go func(fac Factory) {
		p := fac.Project()
		n := p.Name()

		lp := (&LabelProperty{}).Property()
		lp.Set(n)
		res.AddProperty(lp)

		res.MarkSucceeded()
		res.MarkFinished()
	}(lo.Factory())

	return res.Result()
}
