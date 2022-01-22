package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudRouteData represents the model of a cloudroutedata
type CloudRouteData struct {
	// The gateway that this route table is associated with.
	GatewayID string `json:"gatewayID" msgpack:"gatewayID" bson:"gatewayid" mapstructure:"gatewayID,omitempty"`

	// Indicates that this is the default route table for the VPC.
	MainTable bool `json:"mainTable" msgpack:"mainTable" bson:"maintable" mapstructure:"mainTable,omitempty"`

	// Routes associated with this route table.
	Routelist []*CloudRoute `json:"routelist" msgpack:"routelist" bson:"routelist" mapstructure:"routelist,omitempty"`

	// The list of subnets that this route table is associated with.
	SubnetAssociations []string `json:"subnetAssociations" msgpack:"subnetAssociations" bson:"subnetassociations" mapstructure:"subnetAssociations,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudRouteData returns a new *CloudRouteData
func NewCloudRouteData() *CloudRouteData {

	return &CloudRouteData{
		ModelVersion:       1,
		Routelist:          []*CloudRoute{},
		SubnetAssociations: []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudRouteData) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudRouteData{}

	s.GatewayID = o.GatewayID
	s.MainTable = o.MainTable
	s.Routelist = o.Routelist
	s.SubnetAssociations = o.SubnetAssociations

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudRouteData) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudRouteData{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.GatewayID = s.GatewayID
	o.MainTable = s.MainTable
	o.Routelist = s.Routelist
	o.SubnetAssociations = s.SubnetAssociations

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudRouteData) BleveType() string {

	return "cloudroutedata"
}

// DeepCopy returns a deep copy if the CloudRouteData.
func (o *CloudRouteData) DeepCopy() *CloudRouteData {

	if o == nil {
		return nil
	}

	out := &CloudRouteData{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudRouteData.
func (o *CloudRouteData) DeepCopyInto(out *CloudRouteData) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudRouteData: %s", err))
	}

	*out = *target.(*CloudRouteData)
}

// Validate valides the current information stored into the structure.
func (o *CloudRouteData) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.Routelist {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
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

type mongoAttributesCloudRouteData struct {
	GatewayID          string        `bson:"gatewayid"`
	MainTable          bool          `bson:"maintable"`
	Routelist          []*CloudRoute `bson:"routelist"`
	SubnetAssociations []string      `bson:"subnetassociations"`
}
