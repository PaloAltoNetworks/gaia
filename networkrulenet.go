package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// NetworkRuleNet represents the model of a networkrulenet
type NetworkRuleNet struct {
	// The ID of the external network.
	ID string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// List of CIDRs or domain name.
	Entries []string `json:"entries,omitempty" msgpack:"entries,omitempty" bson:"-" mapstructure:"entries,omitempty"`

	// The namespace of the external network.
	Namespace string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"-" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewNetworkRuleNet returns a new *NetworkRuleNet
func NewNetworkRuleNet() *NetworkRuleNet {

	return &NetworkRuleNet{
		ModelVersion: 1,
		Entries:      []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NetworkRuleNet) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesNetworkRuleNet{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NetworkRuleNet) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesNetworkRuleNet{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *NetworkRuleNet) BleveType() string {

	return "networkrulenet"
}

// DeepCopy returns a deep copy if the NetworkRuleNet.
func (o *NetworkRuleNet) DeepCopy() *NetworkRuleNet {

	if o == nil {
		return nil
	}

	out := &NetworkRuleNet{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *NetworkRuleNet.
func (o *NetworkRuleNet) DeepCopyInto(out *NetworkRuleNet) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy NetworkRuleNet: %s", err))
	}

	*out = *target.(*NetworkRuleNet)
}

// Validate valides the current information stored into the structure.
func (o *NetworkRuleNet) Validate() error {

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

type mongoAttributesNetworkRuleNet struct {
}
