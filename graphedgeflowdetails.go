// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// GraphEdgeFlowDetails represents the model of a graphedgeflowdetails
type GraphEdgeFlowDetails struct {
	// Indicates whether the flow was accepted.
	Accepted bool `json:"accepted" msgpack:"accepted" bson:"-" mapstructure:"accepted,omitempty"`

	// The protocol for this edge. If it is tcp or udp, the port is also included.
	ProtoPort string `json:"protoPort" msgpack:"protoPort" bson:"-" mapstructure:"protoPort,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewGraphEdgeFlowDetails returns a new *GraphEdgeFlowDetails
func NewGraphEdgeFlowDetails() *GraphEdgeFlowDetails {

	return &GraphEdgeFlowDetails{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *GraphEdgeFlowDetails) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesGraphEdgeFlowDetails{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *GraphEdgeFlowDetails) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesGraphEdgeFlowDetails{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *GraphEdgeFlowDetails) BleveType() string {

	return "graphedgeflowdetails"
}

// DeepCopy returns a deep copy if the GraphEdgeFlowDetails.
func (o *GraphEdgeFlowDetails) DeepCopy() *GraphEdgeFlowDetails {

	if o == nil {
		return nil
	}

	out := &GraphEdgeFlowDetails{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *GraphEdgeFlowDetails.
func (o *GraphEdgeFlowDetails) DeepCopyInto(out *GraphEdgeFlowDetails) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy GraphEdgeFlowDetails: %s", err))
	}

	*out = *target.(*GraphEdgeFlowDetails)
}

// Validate valides the current information stored into the structure.
func (o *GraphEdgeFlowDetails) Validate() error {

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

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*GraphEdgeFlowDetails) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := GraphEdgeFlowDetailsAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return GraphEdgeFlowDetailsLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*GraphEdgeFlowDetails) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return GraphEdgeFlowDetailsAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *GraphEdgeFlowDetails) ValueForAttribute(name string) any {

	switch name {
	case "accepted":
		return o.Accepted
	case "protoPort":
		return o.ProtoPort
	}

	return nil
}

// GraphEdgeFlowDetailsAttributesMap represents the map of attribute for GraphEdgeFlowDetails.
var GraphEdgeFlowDetailsAttributesMap = map[string]elemental.AttributeSpecification{
	"Accepted": {
		AllowedChoices: []string{},
		ConvertedName:  "Accepted",
		Description:    `Indicates whether the flow was accepted.`,
		Exposed:        true,
		Name:           "accepted",
		Type:           "boolean",
	},
	"ProtoPort": {
		AllowedChoices: []string{},
		ConvertedName:  "ProtoPort",
		Description:    `The protocol for this edge. If it is tcp or udp, the port is also included.`,
		Exposed:        true,
		Name:           "protoPort",
		Type:           "string",
	},
}

// GraphEdgeFlowDetailsLowerCaseAttributesMap represents the map of attribute for GraphEdgeFlowDetails.
var GraphEdgeFlowDetailsLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"accepted": {
		AllowedChoices: []string{},
		ConvertedName:  "Accepted",
		Description:    `Indicates whether the flow was accepted.`,
		Exposed:        true,
		Name:           "accepted",
		Type:           "boolean",
	},
	"protoport": {
		AllowedChoices: []string{},
		ConvertedName:  "ProtoPort",
		Description:    `The protocol for this edge. If it is tcp or udp, the port is also included.`,
		Exposed:        true,
		Name:           "protoPort",
		Type:           "string",
	},
}

type mongoAttributesGraphEdgeFlowDetails struct {
}
