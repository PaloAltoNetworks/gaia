package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudIPRange represents the model of a cloudiprange
type CloudIPRange struct {
	// End IP for the range.
	EndIP string `json:"endIP" msgpack:"endIP" bson:"endip" mapstructure:"endIP,omitempty"`

	// Starting IP for the range.
	StartIP string `json:"startIP" msgpack:"startIP" bson:"startip" mapstructure:"startIP,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudIPRange returns a new *CloudIPRange
func NewCloudIPRange() *CloudIPRange {

	return &CloudIPRange{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudIPRange) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudIPRange{}

	s.EndIP = o.EndIP
	s.StartIP = o.StartIP

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudIPRange) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudIPRange{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.EndIP = s.EndIP
	o.StartIP = s.StartIP

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudIPRange) BleveType() string {

	return "cloudiprange"
}

// DeepCopy returns a deep copy if the CloudIPRange.
func (o *CloudIPRange) DeepCopy() *CloudIPRange {

	if o == nil {
		return nil
	}

	out := &CloudIPRange{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudIPRange.
func (o *CloudIPRange) DeepCopyInto(out *CloudIPRange) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudIPRange: %s", err))
	}

	*out = *target.(*CloudIPRange)
}

// Validate valides the current information stored into the structure.
func (o *CloudIPRange) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateOptionalCIDRorIP("endIP", o.EndIP); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateOptionalCIDRorIP("startIP", o.StartIP); err != nil {
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
func (*CloudIPRange) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudIPRangeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudIPRangeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudIPRange) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudIPRangeAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudIPRange) ValueForAttribute(name string) interface{} {

	switch name {
	case "endIP":
		return o.EndIP
	case "startIP":
		return o.StartIP
	}

	return nil
}

// CloudIPRangeAttributesMap represents the map of attribute for CloudIPRange.
var CloudIPRangeAttributesMap = map[string]elemental.AttributeSpecification{
	"EndIP": {
		AllowedChoices: []string{},
		BSONFieldName:  "endip",
		ConvertedName:  "EndIP",
		Description:    `End IP for the range.`,
		Exposed:        true,
		Name:           "endIP",
		Stored:         true,
		Type:           "string",
	},
	"StartIP": {
		AllowedChoices: []string{},
		BSONFieldName:  "startip",
		ConvertedName:  "StartIP",
		Description:    `Starting IP for the range.`,
		Exposed:        true,
		Name:           "startIP",
		Stored:         true,
		Type:           "string",
	},
}

// CloudIPRangeLowerCaseAttributesMap represents the map of attribute for CloudIPRange.
var CloudIPRangeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"endip": {
		AllowedChoices: []string{},
		BSONFieldName:  "endip",
		ConvertedName:  "EndIP",
		Description:    `End IP for the range.`,
		Exposed:        true,
		Name:           "endIP",
		Stored:         true,
		Type:           "string",
	},
	"startip": {
		AllowedChoices: []string{},
		BSONFieldName:  "startip",
		ConvertedName:  "StartIP",
		Description:    `Starting IP for the range.`,
		Exposed:        true,
		Name:           "startIP",
		Stored:         true,
		Type:           "string",
	},
}

type mongoAttributesCloudIPRange struct {
	EndIP   string `bson:"endip"`
	StartIP string `bson:"startip"`
}
