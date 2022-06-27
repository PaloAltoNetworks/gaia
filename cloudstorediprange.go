package gaia

import (
	"fmt"
	"net"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudStoredIPRange represents the model of a cloudstorediprange
type CloudStoredIPRange struct {
	// End IP for the range.
	EndIP *net.IPNet `json:"endIP" msgpack:"endIP" bson:"endip" mapstructure:"endIP,omitempty"`

	// Starting IP for the range.
	StartIP *net.IPNet `json:"startIP" msgpack:"startIP" bson:"startip" mapstructure:"startIP,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudStoredIPRange returns a new *CloudStoredIPRange
func NewCloudStoredIPRange() *CloudStoredIPRange {

	return &CloudStoredIPRange{
		ModelVersion: 1,
		EndIP:        &net.IPNet{},
		StartIP:      &net.IPNet{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudStoredIPRange) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudStoredIPRange{}

	s.EndIP = o.EndIP
	s.StartIP = o.StartIP

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudStoredIPRange) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudStoredIPRange{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.EndIP = s.EndIP
	o.StartIP = s.StartIP

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudStoredIPRange) BleveType() string {

	return "cloudstorediprange"
}

// DeepCopy returns a deep copy if the CloudStoredIPRange.
func (o *CloudStoredIPRange) DeepCopy() *CloudStoredIPRange {

	if o == nil {
		return nil
	}

	out := &CloudStoredIPRange{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudStoredIPRange.
func (o *CloudStoredIPRange) DeepCopyInto(out *CloudStoredIPRange) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudStoredIPRange: %s", err))
	}

	*out = *target.(*CloudStoredIPRange)
}

// Validate valides the current information stored into the structure.
func (o *CloudStoredIPRange) Validate() error {

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
func (*CloudStoredIPRange) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudStoredIPRangeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudStoredIPRangeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudStoredIPRange) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudStoredIPRangeAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudStoredIPRange) ValueForAttribute(name string) interface{} {

	switch name {
	case "endIP":
		return o.EndIP
	case "startIP":
		return o.StartIP
	}

	return nil
}

// CloudStoredIPRangeAttributesMap represents the map of attribute for CloudStoredIPRange.
var CloudStoredIPRangeAttributesMap = map[string]elemental.AttributeSpecification{
	"EndIP": {
		AllowedChoices: []string{},
		BSONFieldName:  "endip",
		ConvertedName:  "EndIP",
		Description:    `End IP for the range.`,
		Exposed:        true,
		Name:           "endIP",
		Stored:         true,
		SubType:        "network",
		Type:           "external",
	},
	"StartIP": {
		AllowedChoices: []string{},
		BSONFieldName:  "startip",
		ConvertedName:  "StartIP",
		Description:    `Starting IP for the range.`,
		Exposed:        true,
		Name:           "startIP",
		Stored:         true,
		SubType:        "network",
		Type:           "external",
	},
}

// CloudStoredIPRangeLowerCaseAttributesMap represents the map of attribute for CloudStoredIPRange.
var CloudStoredIPRangeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"endip": {
		AllowedChoices: []string{},
		BSONFieldName:  "endip",
		ConvertedName:  "EndIP",
		Description:    `End IP for the range.`,
		Exposed:        true,
		Name:           "endIP",
		Stored:         true,
		SubType:        "network",
		Type:           "external",
	},
	"startip": {
		AllowedChoices: []string{},
		BSONFieldName:  "startip",
		ConvertedName:  "StartIP",
		Description:    `Starting IP for the range.`,
		Exposed:        true,
		Name:           "startIP",
		Stored:         true,
		SubType:        "network",
		Type:           "external",
	},
}

type mongoAttributesCloudStoredIPRange struct {
	EndIP   *net.IPNet `bson:"endip"`
	StartIP *net.IPNet `bson:"startip"`
}
