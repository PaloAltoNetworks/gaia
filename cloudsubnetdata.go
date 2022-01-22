package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudSubnetData represents the model of a cloudsubnetdata
type CloudSubnetData struct {
	// Address CIDR of the Subnet.
	Address string `json:"address" msgpack:"address" bson:"address" mapstructure:"address,omitempty"`

	// The availability zone ID of the subnet.
	ZoneID string `json:"zoneID" msgpack:"zoneID" bson:"zoneid" mapstructure:"zoneID,omitempty"`

	// The availability zone of the subnet.
	ZoneName string `json:"zoneName" msgpack:"zoneName" bson:"zonename" mapstructure:"zoneName,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudSubnetData returns a new *CloudSubnetData
func NewCloudSubnetData() *CloudSubnetData {

	return &CloudSubnetData{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudSubnetData) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudSubnetData{}

	s.Address = o.Address
	s.ZoneID = o.ZoneID
	s.ZoneName = o.ZoneName

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudSubnetData) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudSubnetData{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Address = s.Address
	o.ZoneID = s.ZoneID
	o.ZoneName = s.ZoneName

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudSubnetData) BleveType() string {

	return "cloudsubnetdata"
}

// DeepCopy returns a deep copy if the CloudSubnetData.
func (o *CloudSubnetData) DeepCopy() *CloudSubnetData {

	if o == nil {
		return nil
	}

	out := &CloudSubnetData{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudSubnetData.
func (o *CloudSubnetData) DeepCopyInto(out *CloudSubnetData) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudSubnetData: %s", err))
	}

	*out = *target.(*CloudSubnetData)
}

// Validate valides the current information stored into the structure.
func (o *CloudSubnetData) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("address", o.Address); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := ValidateCIDR("address", o.Address); err != nil {
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

type mongoAttributesCloudSubnetData struct {
	Address  string `bson:"address"`
	ZoneID   string `bson:"zoneid"`
	ZoneName string `bson:"zonename"`
}
