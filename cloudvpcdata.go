package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudVPCData represents the model of a cloudvpcdata
type CloudVPCData struct {
	// Address CIDR of the VPC.
	Address string `json:"address" msgpack:"address" bson:"address" mapstructure:"address,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudVPCData returns a new *CloudVPCData
func NewCloudVPCData() *CloudVPCData {

	return &CloudVPCData{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudVPCData) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudVPCData{}

	s.Address = o.Address

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudVPCData) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudVPCData{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Address = s.Address

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudVPCData) BleveType() string {

	return "cloudvpcdata"
}

// DeepCopy returns a deep copy if the CloudVPCData.
func (o *CloudVPCData) DeepCopy() *CloudVPCData {

	if o == nil {
		return nil
	}

	out := &CloudVPCData{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudVPCData.
func (o *CloudVPCData) DeepCopyInto(out *CloudVPCData) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudVPCData: %s", err))
	}

	*out = *target.(*CloudVPCData)
}

// Validate valides the current information stored into the structure.
func (o *CloudVPCData) Validate() error {

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
func (*CloudVPCData) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudVPCDataAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudVPCDataLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudVPCData) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudVPCDataAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudVPCData) ValueForAttribute(name string) interface{} {

	switch name {
	case "address":
		return o.Address
	}

	return nil
}

// CloudVPCDataAttributesMap represents the map of attribute for CloudVPCData.
var CloudVPCDataAttributesMap = map[string]elemental.AttributeSpecification{
	"Address": {
		AllowedChoices: []string{},
		BSONFieldName:  "address",
		ConvertedName:  "Address",
		Description:    `Address CIDR of the VPC.`,
		Exposed:        true,
		Name:           "address",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}

// CloudVPCDataLowerCaseAttributesMap represents the map of attribute for CloudVPCData.
var CloudVPCDataLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"address": {
		AllowedChoices: []string{},
		BSONFieldName:  "address",
		ConvertedName:  "Address",
		Description:    `Address CIDR of the VPC.`,
		Exposed:        true,
		Name:           "address",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}

type mongoAttributesCloudVPCData struct {
	Address string `bson:"address"`
}
