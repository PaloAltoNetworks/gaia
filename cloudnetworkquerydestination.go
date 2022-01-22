package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudNetworkQueryDestinationTypeValue represents the possible values for attribute "type".
type CloudNetworkQueryDestinationTypeValue string

const (
	// CloudNetworkQueryDestinationTypeInstance represents the value Instance.
	CloudNetworkQueryDestinationTypeInstance CloudNetworkQueryDestinationTypeValue = "Instance"

	// CloudNetworkQueryDestinationTypeInterface represents the value Interface.
	CloudNetworkQueryDestinationTypeInterface CloudNetworkQueryDestinationTypeValue = "Interface"

	// CloudNetworkQueryDestinationTypeLoadBalancer represents the value LoadBalancer.
	CloudNetworkQueryDestinationTypeLoadBalancer CloudNetworkQueryDestinationTypeValue = "LoadBalancer"

	// CloudNetworkQueryDestinationTypePublicIP represents the value PublicIP.
	CloudNetworkQueryDestinationTypePublicIP CloudNetworkQueryDestinationTypeValue = "PublicIP"
)

// CloudNetworkQueryDestination represents the model of a cloudnetworkquerydestination
type CloudNetworkQueryDestination struct {
	// Returns the native ID of the indirect node.
	IndirectNodeID string `json:"indirectNodeID,omitempty" msgpack:"indirectNodeID,omitempty" bson:"-" mapstructure:"indirectNodeID,omitempty"`

	// Returns true if this is an indirect path through an forwarding entities.
	IsIndirect bool `json:"isIndirect" msgpack:"isIndirect" bson:"-" mapstructure:"isIndirect,omitempty"`

	// Returns true if the destination is reachable through routing.
	Reachable bool `json:"reachable" msgpack:"reachable" bson:"-" mapstructure:"reachable,omitempty"`

	// Returns the type of the destination.
	Type CloudNetworkQueryDestinationTypeValue `json:"type" msgpack:"type" bson:"-" mapstructure:"type,omitempty"`

	// Returns the network security verdict for the destination.
	Verdict string `json:"verdict" msgpack:"verdict" bson:"-" mapstructure:"verdict,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudNetworkQueryDestination returns a new *CloudNetworkQueryDestination
func NewCloudNetworkQueryDestination() *CloudNetworkQueryDestination {

	return &CloudNetworkQueryDestination{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudNetworkQueryDestination) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudNetworkQueryDestination{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudNetworkQueryDestination) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudNetworkQueryDestination{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudNetworkQueryDestination) BleveType() string {

	return "cloudnetworkquerydestination"
}

// DeepCopy returns a deep copy if the CloudNetworkQueryDestination.
func (o *CloudNetworkQueryDestination) DeepCopy() *CloudNetworkQueryDestination {

	if o == nil {
		return nil
	}

	out := &CloudNetworkQueryDestination{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudNetworkQueryDestination.
func (o *CloudNetworkQueryDestination) DeepCopyInto(out *CloudNetworkQueryDestination) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudNetworkQueryDestination: %s", err))
	}

	*out = *target.(*CloudNetworkQueryDestination)
}

// Validate valides the current information stored into the structure.
func (o *CloudNetworkQueryDestination) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Interface", "Instance", "LoadBalancer", "PublicIP"}, true); err != nil {
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

type mongoAttributesCloudNetworkQueryDestination struct {
}
