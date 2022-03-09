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

	// Network Security Group associated with the subnet.
	NsgID string `json:"nsgID" msgpack:"nsgID" bson:"nsgid" mapstructure:"nsgID,omitempty"`

	// Route Table associated with the subnet.
	RouteTableID string `json:"routeTableID" msgpack:"routeTableID" bson:"routetableid" mapstructure:"routeTableID,omitempty"`

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
	s.NsgID = o.NsgID
	s.RouteTableID = o.RouteTableID
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
	o.NsgID = s.NsgID
	o.RouteTableID = s.RouteTableID
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

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*CloudSubnetData) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudSubnetDataAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudSubnetDataLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudSubnetData) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudSubnetDataAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudSubnetData) ValueForAttribute(name string) interface{} {

	switch name {
	case "address":
		return o.Address
	case "nsgID":
		return o.NsgID
	case "routeTableID":
		return o.RouteTableID
	case "zoneID":
		return o.ZoneID
	case "zoneName":
		return o.ZoneName
	}

	return nil
}

// CloudSubnetDataAttributesMap represents the map of attribute for CloudSubnetData.
var CloudSubnetDataAttributesMap = map[string]elemental.AttributeSpecification{
	"Address": {
		AllowedChoices: []string{},
		BSONFieldName:  "address",
		ConvertedName:  "Address",
		Description:    `Address CIDR of the Subnet.`,
		Exposed:        true,
		Name:           "address",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"NsgID": {
		AllowedChoices: []string{},
		BSONFieldName:  "nsgid",
		ConvertedName:  "NsgID",
		Description:    `Network Security Group associated with the subnet.`,
		Exposed:        true,
		Name:           "nsgID",
		Stored:         true,
		Type:           "string",
	},
	"RouteTableID": {
		AllowedChoices: []string{},
		BSONFieldName:  "routetableid",
		ConvertedName:  "RouteTableID",
		Description:    `Route Table associated with the subnet.`,
		Exposed:        true,
		Name:           "routeTableID",
		Stored:         true,
		Type:           "string",
	},
	"ZoneID": {
		AllowedChoices: []string{},
		BSONFieldName:  "zoneid",
		ConvertedName:  "ZoneID",
		Description:    `The availability zone ID of the subnet.`,
		Exposed:        true,
		Name:           "zoneID",
		Stored:         true,
		Type:           "string",
	},
	"ZoneName": {
		AllowedChoices: []string{},
		BSONFieldName:  "zonename",
		ConvertedName:  "ZoneName",
		Description:    `The availability zone of the subnet.`,
		Exposed:        true,
		Name:           "zoneName",
		Stored:         true,
		Type:           "string",
	},
}

// CloudSubnetDataLowerCaseAttributesMap represents the map of attribute for CloudSubnetData.
var CloudSubnetDataLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"address": {
		AllowedChoices: []string{},
		BSONFieldName:  "address",
		ConvertedName:  "Address",
		Description:    `Address CIDR of the Subnet.`,
		Exposed:        true,
		Name:           "address",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"nsgid": {
		AllowedChoices: []string{},
		BSONFieldName:  "nsgid",
		ConvertedName:  "NsgID",
		Description:    `Network Security Group associated with the subnet.`,
		Exposed:        true,
		Name:           "nsgID",
		Stored:         true,
		Type:           "string",
	},
	"routetableid": {
		AllowedChoices: []string{},
		BSONFieldName:  "routetableid",
		ConvertedName:  "RouteTableID",
		Description:    `Route Table associated with the subnet.`,
		Exposed:        true,
		Name:           "routeTableID",
		Stored:         true,
		Type:           "string",
	},
	"zoneid": {
		AllowedChoices: []string{},
		BSONFieldName:  "zoneid",
		ConvertedName:  "ZoneID",
		Description:    `The availability zone ID of the subnet.`,
		Exposed:        true,
		Name:           "zoneID",
		Stored:         true,
		Type:           "string",
	},
	"zonename": {
		AllowedChoices: []string{},
		BSONFieldName:  "zonename",
		ConvertedName:  "ZoneName",
		Description:    `The availability zone of the subnet.`,
		Exposed:        true,
		Name:           "zoneName",
		Stored:         true,
		Type:           "string",
	},
}

type mongoAttributesCloudSubnetData struct {
	Address      string `bson:"address"`
	NsgID        string `bson:"nsgid"`
	RouteTableID string `bson:"routetableid"`
	ZoneID       string `bson:"zoneid"`
	ZoneName     string `bson:"zonename"`
}
