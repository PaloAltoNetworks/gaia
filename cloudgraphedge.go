package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudGraphEdge represents the model of a cloudgraphedge
type CloudGraphEdge struct {
	// Provides the level of the tree that this edge belongs in order to assist with
	// ordering.
	Level int `json:"level" msgpack:"level" bson:"level" mapstructure:"level,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudGraphEdge returns a new *CloudGraphEdge
func NewCloudGraphEdge() *CloudGraphEdge {

	return &CloudGraphEdge{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudGraphEdge) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudGraphEdge{}

	s.Level = o.Level

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudGraphEdge) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudGraphEdge{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Level = s.Level

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudGraphEdge) BleveType() string {

	return "cloudgraphedge"
}

// DeepCopy returns a deep copy if the CloudGraphEdge.
func (o *CloudGraphEdge) DeepCopy() *CloudGraphEdge {

	if o == nil {
		return nil
	}

	out := &CloudGraphEdge{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudGraphEdge.
func (o *CloudGraphEdge) DeepCopyInto(out *CloudGraphEdge) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudGraphEdge: %s", err))
	}

	*out = *target.(*CloudGraphEdge)
}

// Validate valides the current information stored into the structure.
func (o *CloudGraphEdge) Validate() error {

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

type mongoAttributesCloudGraphEdge struct {
	Level int `bson:"level"`
}
