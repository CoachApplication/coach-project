package project

import (
	api "github.com/CoachApplication/api"
	base "github.com/CoachApplication/base"
	base_props "github.com/CoachApplication/base/property"
)

const (
	PROPERTY_ID_NAME      = "project.name"
	PROPERTY_ID_LABEL     = "project.label"
	PROPERTY_ID_ENV_MAP   = "project.environment"
	PROPERTY_ID_ENV_KEYS  = "project.environment.keys"
	PROPERTY_ID_ENV_KEY   = "project.environment.key"
	PROPERTY_ID_ENV_VALUE = "project.environment.value"
)

type NameProperty struct {
	base_props.StringProperty
}

func (np *NameProperty) Property() api.Property {
	return api.Property(np)
}

func (np *NameProperty) Id() string {
	return PROPERTY_ID_NAME
}

func (np *NameProperty) Ui() api.Ui {
	return base.NewUi(
		np.Id(),
		"Project machine name",
		"A string name for the project that can be used to for machine labeling of the project relaed elements.",
		"This property contains a string value that is machine safe usable for naming related assets.  The string should contain only alpha-numeric characters to be safe.",
	)
}

func (np *NameProperty) Usage() api.Usage {
	return base.ReadonlyPropertyUsage{}
}

type LabelProperty struct {
	base_props.StringProperty
}

func (lp *LabelProperty) Property() api.Property {
	return api.Property(lp)
}

func (lp *LabelProperty) Id() string {
	return PROPERTY_ID_LABEL
}

func (lp *LabelProperty) Ui() api.Ui {
	return base.NewUi(
		lp.Id(),
		"Project label",
		"A string label for the project that can be used to describe the project",
		"",
	)
}

func (lp *LabelProperty) Usage() api.Usage {
	return base.ReadonlyPropertyUsage{}
}

type EnvironmentMapProperty struct {
	base_props.StringMapProperty
}

func (emmp *EnvironmentMapProperty) Property() api.Property {
	return api.Property(emmp)
}

func (emmp *EnvironmentMapProperty) Id() string {
	return PROPERTY_ID_ENV_MAP
}

func (emmp *EnvironmentMapProperty) Ui() api.Ui {
	return base.NewUi(
		emmp.Id(),
		"Project environment map",
		"An environment map in the form of a string map, used for string settings of the project",
		"",
	)
}

func (emmp *EnvironmentMapProperty) Usage() api.Usage {
	return base.ReadonlyPropertyUsage{}
}

type EnvironmentMapKeyProperty struct {
	base_props.StringProperty
}

func (emkp *EnvironmentMapKeyProperty) Property() api.Property {
	return api.Property(emkp)
}

func (emkp *EnvironmentMapKeyProperty) Id() string {
	return PROPERTY_ID_ENV_KEY
}

func (emkp *EnvironmentMapKeyProperty) Ui() api.Ui {
	return base.NewUi(
		emkp.Id(),
		"Project environment key",
		"A single key from the environment map",
		"",
	)
}

func (emkp *EnvironmentMapKeyProperty) Usage() api.Usage {
	return base.RequiredPropertyUsage{}
}

type EnvironmentMapKeysProperty struct {
	base_props.StringSliceProperty
}

func (emkp *EnvironmentMapKeysProperty) Property() api.Property {
	return api.Property(emkp)
}

func (emkp *EnvironmentMapKeysProperty) Id() string {
	return PROPERTY_ID_ENV_KEYS
}

func (emkp *EnvironmentMapKeysProperty) Ui() api.Ui {
	return base.NewUi(
		emkp.Id(),
		"Project environment keys",
		"A list of keys from the environment map",
		"",
	)
}

func (emkp *EnvironmentMapKeysProperty) Usage() api.Usage {
	return base.ReadonlyPropertyUsage{}
}

type EnvironmentMapValueProperty struct {
	base_props.StringProperty
}

func (emvp *EnvironmentMapValueProperty) Property() api.Property {
	return api.Property(emvp)
}

func (emvp *EnvironmentMapValueProperty) Id() string {
	return PROPERTY_ID_ENV_VALUE
}

func (emvp *EnvironmentMapValueProperty) Ui() api.Ui {
	return base.NewUi(
		emvp.Id(),
		"Project environment value",
		"A single environment map value form the environment map",
		"",
	)
}

func (emvp *EnvironmentMapValueProperty) Usage() api.Usage {
	return base.ReadonlyPropertyUsage{}
}
