package gaia

import (
	"fmt"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// UIStep represents the model of a uistep
type UIStep struct {
	// Defines if the step is an advanced one.
	Advanced bool `json:"advanced" msgpack:"advanced" bson:"advanced" mapstructure:"advanced,omitempty"`

	// Description of the step.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Name of the step.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// List of parameters for this step.
	Parameters []*UIParameter `json:"parameters" msgpack:"parameters" bson:"parameters" mapstructure:"parameters,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewUIStep returns a new *UIStep
func NewUIStep() *UIStep {

	return &UIStep{
		ModelVersion: 1,
		Parameters:   []*UIParameter{},
	}
}

// DeepCopy returns a deep copy if the UIStep.
func (o *UIStep) DeepCopy() *UIStep {

	if o == nil {
		return nil
	}

	out := &UIStep{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *UIStep.
func (o *UIStep) DeepCopyInto(out *UIStep) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy UIStep: %s", err))
	}

	*out = *target.(*UIStep)
}

// Validate valides the current information stored into the structure.
func (o *UIStep) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	for _, sub := range o.Parameters {
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}
