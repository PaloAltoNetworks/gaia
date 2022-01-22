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
func (o *GraphEdgeFlowDetails) GetBSON() (interface{}, error) {

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

type mongoAttributesGraphEdgeFlowDetails struct {
}
