package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudPathElement represents the model of a cloudpathelement
type CloudPathElement struct {
	// The native ID of the node.
	NativeID string `json:"nativeID" msgpack:"nativeID" bson:"-" mapstructure:"nativeID,omitempty"`

	// The policy ID used at this node of the path.
	Policy *CloudGraphNodeAction `json:"policy,omitempty" msgpack:"policy,omitempty" bson:"-" mapstructure:"policy,omitempty"`

	// The route table ID used for the route calculation.
	RouteTables map[string]string `json:"routeTables,omitempty" msgpack:"routeTables,omitempty" bson:"-" mapstructure:"routeTables,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudPathElement returns a new *CloudPathElement
func NewCloudPathElement() *CloudPathElement {

	return &CloudPathElement{
		ModelVersion: 1,
		Policy:       NewCloudGraphNodeAction(),
		RouteTables:  map[string]string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudPathElement) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudPathElement{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudPathElement) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudPathElement{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudPathElement) BleveType() string {

	return "cloudpathelement"
}

// DeepCopy returns a deep copy if the CloudPathElement.
func (o *CloudPathElement) DeepCopy() *CloudPathElement {

	if o == nil {
		return nil
	}

	out := &CloudPathElement{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudPathElement.
func (o *CloudPathElement) DeepCopyInto(out *CloudPathElement) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudPathElement: %s", err))
	}

	*out = *target.(*CloudPathElement)
}

// Validate valides the current information stored into the structure.
func (o *CloudPathElement) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if o.Policy != nil {
		elemental.ResetDefaultForZeroValues(o.Policy)
		if err := o.Policy.Validate(); err != nil {
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

type mongoAttributesCloudPathElement struct {
}
