package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudGraphNodeAction represents the model of a cloudgraphnodeaction
type CloudGraphNodeAction struct {
	// The action that is been applied for the particular destination.
	Action string `json:"action" msgpack:"action" bson:"-" mapstructure:"action,omitempty"`

	// The ID of the policies that were used in the path.
	PolicyID string `json:"policyID" msgpack:"policyID" bson:"-" mapstructure:"policyID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudGraphNodeAction returns a new *CloudGraphNodeAction
func NewCloudGraphNodeAction() *CloudGraphNodeAction {

	return &CloudGraphNodeAction{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudGraphNodeAction) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudGraphNodeAction{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudGraphNodeAction) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudGraphNodeAction{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudGraphNodeAction) BleveType() string {

	return "cloudgraphnodeaction"
}

// DeepCopy returns a deep copy if the CloudGraphNodeAction.
func (o *CloudGraphNodeAction) DeepCopy() *CloudGraphNodeAction {

	if o == nil {
		return nil
	}

	out := &CloudGraphNodeAction{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudGraphNodeAction.
func (o *CloudGraphNodeAction) DeepCopyInto(out *CloudGraphNodeAction) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudGraphNodeAction: %s", err))
	}

	*out = *target.(*CloudGraphNodeAction)
}

// Validate valides the current information stored into the structure.
func (o *CloudGraphNodeAction) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

type mongoAttributesCloudGraphNodeAction struct {
}
