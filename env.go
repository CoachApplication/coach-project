package project

import (
	"errors"
	"github.com/CoachApplication/api"
	"github.com/CoachApplication/base"
)

const (
	OPERATION_ID_ENV_MAP  = "project.env"
	OPERATION_ID_ENV_LIST = "project.env.list"
	OPERATION_ID_ENV_GET  = "project.env.get"
)

type EnvironmentMapOperation struct {
	FactoryOperationBase
}

func NewEnvironmentMapOperation(base FactoryOperationBase) *EnvironmentMapOperation {
	return &EnvironmentMapOperation{
		FactoryOperationBase: base,
	}
}

func (emo *EnvironmentMapOperation) Operation() api.Operation {
	return api.Operation(emo)
}

func (emo *EnvironmentMapOperation) Id() string {
	return OPERATION_ID_ENV_MAP
}

func (emo *EnvironmentMapOperation) Ui() api.Ui {
	return base.NewUi(
		emo.Id(),
		"Project environment map",
		"Key value string map of environment variables for the project",
		"",
	)
}

func (emo *EnvironmentMapOperation) Usage() api.Usage {
	return base.ExternalOperationUsage{}
}

func (emo *EnvironmentMapOperation) Properties() api.Properties {
	return base.NewProperties().Properties()
}

func (emo *EnvironmentMapOperation) Validate(props api.Properties) api.Result {
	return emo.FactoryOperationBase.Validate(props)
}

func (emo *EnvironmentMapOperation) Exec(props api.Properties) api.Result {
	res := base.NewResult()

	go func(fac Factory) {
		p := fac.Project()
		m := p.Env()

		mp := (&EnvironmentMapProperty{}).Property()
		mp.Set(m)
		res.AddProperty(mp)

		res.MarkSucceeded()
		res.MarkFinished()
	}(emo.Factory())

	return res.Result()
}

type EnvironmentGetOperation struct {
	FactoryOperationBase
}

func NewEnvironmentGetOperation(base FactoryOperationBase) *EnvironmentGetOperation {
	return &EnvironmentGetOperation{
		FactoryOperationBase: base,
	}
}

func (ego *EnvironmentGetOperation) Operation() api.Operation {
	return api.Operation(ego)
}

func (ego *EnvironmentGetOperation) Id() string {
	return OPERATION_ID_ENV_GET
}

func (ego *EnvironmentGetOperation) Ui() api.Ui {
	return base.NewUi(
		ego.Id(),
		"Project environment variable",
		"Single environment variable for the project",
		"",
	)
}

func (ego *EnvironmentGetOperation) Usage() api.Usage {
	return base.ExternalOperationUsage{}
}

func (ego *EnvironmentGetOperation) Properties() api.Properties {
	props := base.NewProperties()

	// PROPERTY_ID_ENV_KEY Property
	props.Add((&EnvironmentMapKeyProperty{}).Property())

	return props.Properties()
}

func (ego *EnvironmentGetOperation) Validate(props api.Properties) api.Result {
	res := base.NewResult()

	res.Merge(ego.FactoryOperationBase.Validate(props))

	propRes := base.NewResult()
	go func(props api.Properties) {
		if kp, err := props.Get(PROPERTY_ID_ENV_KEY); err != nil {
			res.AddError(errors.New("MIssing required env key property: " + PROPERTY_ID_ENV_KEY))
			res.MarkFailed()
		} else if key, good := kp.Get().(string); !good {
			res.AddError(errors.New("Invalid env key value provided"))
			res.MarkFailed()
		} else if key == "" {
			res.AddError(errors.New("Empty env key value provided"))
			res.MarkFailed()
		}
		res.MarkFinished()
	}(props)
	res.Merge(propRes)

	return res.Result()
}

func (ego *EnvironmentGetOperation) Exec(props api.Properties) api.Result {
	res := base.NewResult()
	res.Merge(ego.Validate(props))

	go func(fac Factory, props api.Properties) {
		if kp, err := props.Get(PROPERTY_ID_ENV_KEY); err != nil {
			res.AddError(err)
			res.MarkFailed()
		} else if key, good := kp.Get().(string); !good {
			res.AddError(errors.New("Missing required property: " + PROPERTY_ID_ENV_KEY))
		} else if key == "" {
			res.AddError(errors.New("Empty required property: " + PROPERTY_ID_ENV_KEY))
		} else {
			p := fac.Project()
			m := p.Env()

			valProp := (&EnvironmentMapValueProperty{}).Property()
			if val, good := m[key]; good {
				valProp.Set(val)
				res.MarkSucceeded()
			} else {
				res.AddError(errors.New("Environment key was not found: " + key))
				res.MarkFailed()
			}
			res.AddProperty(valProp)
		}

		res.MarkFinished()
	}(ego.Factory(), props)

	return res.Result()
}

type EnvironmentListOperation struct {
	FactoryOperationBase
}

func NewEnvironmentListOperation(base FactoryOperationBase) *EnvironmentListOperation {
	return &EnvironmentListOperation{
		FactoryOperationBase: base,
	}
}

func (elo *EnvironmentListOperation) Operation() api.Operation {
	return api.Operation(elo)
}

func (elo *EnvironmentListOperation) Id() string {
	return OPERATION_ID_ENV_LIST
}

func (elo *EnvironmentListOperation) Ui() api.Ui {
	return base.NewUi(
		elo.Id(),
		"Project name",
		"Machine name for the project",
		"",
	)
}

func (elo *EnvironmentListOperation) Usage() api.Usage {
	return base.ExternalOperationUsage{}
}

func (elo *EnvironmentListOperation) Properties() api.Properties {
	return base.NewProperties().Properties()
}

func (elo *EnvironmentListOperation) Validate(props api.Properties) api.Result {
	return elo.FactoryOperationBase.Validate(props)
}

func (elo *EnvironmentListOperation) Exec(props api.Properties) api.Result {
	res := base.NewResult()

	go func(fac Factory) {
		p := fac.Project()
		m := p.Env()

		ks := []string{}
		for k, _ := range m {
			ks = append(ks, k)
		}

		kp := (&EnvironmentMapKeysProperty{}).Property()
		kp.Set(ks)
		res.AddProperty(kp)

		res.MarkSucceeded()
		res.MarkFinished()
	}(elo.Factory())

	return res.Result()
}
