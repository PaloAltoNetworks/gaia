package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// UIParameterVisibilityOperatorValue represents the possible values for attribute "operator".
type UIParameterVisibilityOperatorValue string

const (
	// UIParameterVisibilityOperatorDefined represents the value Defined.
	UIParameterVisibilityOperatorDefined UIParameterVisibilityOperatorValue = "Defined"

	// UIParameterVisibilityOperatorEqual represents the value Equal.
	UIParameterVisibilityOperatorEqual UIParameterVisibilityOperatorValue = "Equal"

	// UIParameterVisibilityOperatorGreaterThan represents the value GreaterThan.
	UIParameterVisibilityOperatorGreaterThan UIParameterVisibilityOperatorValue = "GreaterThan"

	// UIParameterVisibilityOperatorLesserThan represents the value LesserThan.
	UIParameterVisibilityOperatorLesserThan UIParameterVisibilityOperatorValue = "LesserThan"

	// UIParameterVisibilityOperatorMatch represents the value Match.
	UIParameterVisibilityOperatorMatch UIParameterVisibilityOperatorValue = "Match"

	// UIParameterVisibilityOperatorNotEqual represents the value NotEqual.
	UIParameterVisibilityOperatorNotEqual UIParameterVisibilityOperatorValue = "NotEqual"

	// UIParameterVisibilityOperatorNotMatch represents the value NotMatch.
	UIParameterVisibilityOperatorNotMatch UIParameterVisibilityOperatorValue = "NotMatch"

	// UIParameterVisibilityOperatorUndefined represents the value Undefined.
	UIParameterVisibilityOperatorUndefined UIParameterVisibilityOperatorValue = "Undefined"
)

// UIParameterVisibility represents the model of a uiparametervisibility
type UIParameterVisibility struct {
	// Key holding the value to compare.
	Key string `json:"key" msgpack:"key" bson:"key" mapstructure:"key,omitempty"`

	// Operator to apply.
	Operator UIParameterVisibilityOperatorValue `json:"operator" msgpack:"operator" bson:"operator" mapstructure:"operator,omitempty"`

	// Values that must match the key.
	Value interface{} `json:"value" msgpack:"value" bson:"value" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewUIParameterVisibility returns a new *UIParameterVisibility
func NewUIParameterVisibility() *UIParameterVisibility {

	return &UIParameterVisibility{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *UIParameterVisibility) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesUIParameterVisibility{}

	s.Key = o.Key
	s.Operator = o.Operator
	s.Value = o.Value

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *UIParameterVisibility) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesUIParameterVisibility{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Key = s.Key
	o.Operator = s.Operator
	o.Value = s.Value

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *UIParameterVisibility) BleveType() string {

	return "uiparametervisibility"
}

// DeepCopy returns a deep copy if the UIParameterVisibility.
func (o *UIParameterVisibility) DeepCopy() *UIParameterVisibility {

	if o == nil {
		return nil
	}

	out := &UIParameterVisibility{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *UIParameterVisibility.
func (o *UIParameterVisibility) DeepCopyInto(out *UIParameterVisibility) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy UIParameterVisibility: %s", err))
	}

	*out = *target.(*UIParameterVisibility)
}

// Validate valides the current information stored into the structure.
func (o *UIParameterVisibility) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("key", o.Key); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("operator", string(o.Operator), []string{"Equal", "NotEqual", "GreaterThan", "LesserThan", "Defined", "Undefined", "Match", "NotMatch"}, false); err != nil {
		errors = errors.Append(err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

type mongoAttributesUIParameterVisibility struct {
	Key      string                             `bson:"key"`
	Operator UIParameterVisibilityOperatorValue `bson:"operator"`
	Value    interface{}                        `bson:"value"`
}
