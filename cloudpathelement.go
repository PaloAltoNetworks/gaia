// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

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
func (o *CloudPathElement) GetBSON() (any, error) {

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

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*CloudPathElement) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudPathElementAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudPathElementLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudPathElement) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudPathElementAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudPathElement) ValueForAttribute(name string) any {

	switch name {
	case "nativeID":
		return o.NativeID
	case "policy":
		return o.Policy
	case "routeTables":
		return o.RouteTables
	}

	return nil
}

// CloudPathElementAttributesMap represents the map of attribute for CloudPathElement.
var CloudPathElementAttributesMap = map[string]elemental.AttributeSpecification{
	"NativeID": {
		AllowedChoices: []string{},
		ConvertedName:  "NativeID",
		Description:    `The native ID of the node.`,
		Exposed:        true,
		Name:           "nativeID",
		ReadOnly:       true,
		Type:           "string",
	},
	"Policy": {
		AllowedChoices: []string{},
		ConvertedName:  "Policy",
		Description:    `The policy ID used at this node of the path.`,
		Exposed:        true,
		Name:           "policy",
		ReadOnly:       true,
		SubType:        "cloudgraphnodeaction",
		Type:           "ref",
	},
	"RouteTables": {
		AllowedChoices: []string{},
		ConvertedName:  "RouteTables",
		Description:    `The route table ID used for the route calculation.`,
		Exposed:        true,
		Name:           "routeTables",
		ReadOnly:       true,
		SubType:        "map[string]string",
		Type:           "external",
	},
}

// CloudPathElementLowerCaseAttributesMap represents the map of attribute for CloudPathElement.
var CloudPathElementLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"nativeid": {
		AllowedChoices: []string{},
		ConvertedName:  "NativeID",
		Description:    `The native ID of the node.`,
		Exposed:        true,
		Name:           "nativeID",
		ReadOnly:       true,
		Type:           "string",
	},
	"policy": {
		AllowedChoices: []string{},
		ConvertedName:  "Policy",
		Description:    `The policy ID used at this node of the path.`,
		Exposed:        true,
		Name:           "policy",
		ReadOnly:       true,
		SubType:        "cloudgraphnodeaction",
		Type:           "ref",
	},
	"routetables": {
		AllowedChoices: []string{},
		ConvertedName:  "RouteTables",
		Description:    `The route table ID used for the route calculation.`,
		Exposed:        true,
		Name:           "routeTables",
		ReadOnly:       true,
		SubType:        "map[string]string",
		Type:           "external",
	},
}

type mongoAttributesCloudPathElement struct {
}
