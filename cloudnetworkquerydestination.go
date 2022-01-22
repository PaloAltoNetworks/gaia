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

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*CloudNetworkQueryDestination) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudNetworkQueryDestinationAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudNetworkQueryDestinationLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudNetworkQueryDestination) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudNetworkQueryDestinationAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudNetworkQueryDestination) ValueForAttribute(name string) interface{} {

	switch name {
	case "indirectNodeID":
		return o.IndirectNodeID
	case "isIndirect":
		return o.IsIndirect
	case "reachable":
		return o.Reachable
	case "type":
		return o.Type
	case "verdict":
		return o.Verdict
	}

	return nil
}

// CloudNetworkQueryDestinationAttributesMap represents the map of attribute for CloudNetworkQueryDestination.
var CloudNetworkQueryDestinationAttributesMap = map[string]elemental.AttributeSpecification{
	"IndirectNodeID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "IndirectNodeID",
		Description:    `Returns the native ID of the indirect node.`,
		Exposed:        true,
		Name:           "indirectNodeID",
		ReadOnly:       true,
		Type:           "string",
	},
	"IsIndirect": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "IsIndirect",
		Description:    `Returns true if this is an indirect path through an forwarding entities.`,
		Exposed:        true,
		Name:           "isIndirect",
		ReadOnly:       true,
		Type:           "boolean",
	},
	"Reachable": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Reachable",
		Description:    `Returns true if the destination is reachable through routing.`,
		Exposed:        true,
		Name:           "reachable",
		ReadOnly:       true,
		Type:           "boolean",
	},
	"Type": {
		AllowedChoices: []string{"Interface", "Instance", "LoadBalancer", "PublicIP"},
		Autogenerated:  true,
		ConvertedName:  "Type",
		Description:    `Returns the type of the destination.`,
		Exposed:        true,
		Name:           "type",
		ReadOnly:       true,
		Type:           "enum",
	},
	"Verdict": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Verdict",
		Description:    `Returns the network security verdict for the destination.`,
		Exposed:        true,
		Name:           "verdict",
		ReadOnly:       true,
		Type:           "string",
	},
}

// CloudNetworkQueryDestinationLowerCaseAttributesMap represents the map of attribute for CloudNetworkQueryDestination.
var CloudNetworkQueryDestinationLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"indirectnodeid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "IndirectNodeID",
		Description:    `Returns the native ID of the indirect node.`,
		Exposed:        true,
		Name:           "indirectNodeID",
		ReadOnly:       true,
		Type:           "string",
	},
	"isindirect": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "IsIndirect",
		Description:    `Returns true if this is an indirect path through an forwarding entities.`,
		Exposed:        true,
		Name:           "isIndirect",
		ReadOnly:       true,
		Type:           "boolean",
	},
	"reachable": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Reachable",
		Description:    `Returns true if the destination is reachable through routing.`,
		Exposed:        true,
		Name:           "reachable",
		ReadOnly:       true,
		Type:           "boolean",
	},
	"type": {
		AllowedChoices: []string{"Interface", "Instance", "LoadBalancer", "PublicIP"},
		Autogenerated:  true,
		ConvertedName:  "Type",
		Description:    `Returns the type of the destination.`,
		Exposed:        true,
		Name:           "type",
		ReadOnly:       true,
		Type:           "enum",
	},
	"verdict": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Verdict",
		Description:    `Returns the network security verdict for the destination.`,
		Exposed:        true,
		Name:           "verdict",
		ReadOnly:       true,
		Type:           "string",
	},
}

type mongoAttributesCloudNetworkQueryDestination struct {
}
