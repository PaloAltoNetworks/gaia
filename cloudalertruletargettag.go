// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudAlertRuleTargetTag represents the model of a cloudalertruletargettag
type CloudAlertRuleTargetTag struct {
	// Alert Rule target tag key.
	Key string `json:"key" msgpack:"key" bson:"key" mapstructure:"key,omitempty"`

	// List of Alert Rule target tag values.
	Values []string `json:"values" msgpack:"values" bson:"values" mapstructure:"values,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudAlertRuleTargetTag returns a new *CloudAlertRuleTargetTag
func NewCloudAlertRuleTargetTag() *CloudAlertRuleTargetTag {

	return &CloudAlertRuleTargetTag{
		ModelVersion: 1,
		Values:       []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudAlertRuleTargetTag) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudAlertRuleTargetTag{}

	s.Key = o.Key
	s.Values = o.Values

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudAlertRuleTargetTag) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudAlertRuleTargetTag{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Key = s.Key
	o.Values = s.Values

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudAlertRuleTargetTag) BleveType() string {

	return "cloudalertruletargettag"
}

// DeepCopy returns a deep copy if the CloudAlertRuleTargetTag.
func (o *CloudAlertRuleTargetTag) DeepCopy() *CloudAlertRuleTargetTag {

	if o == nil {
		return nil
	}

	out := &CloudAlertRuleTargetTag{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudAlertRuleTargetTag.
func (o *CloudAlertRuleTargetTag) DeepCopyInto(out *CloudAlertRuleTargetTag) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudAlertRuleTargetTag: %s", err))
	}

	*out = *target.(*CloudAlertRuleTargetTag)
}

// Validate valides the current information stored into the structure.
func (o *CloudAlertRuleTargetTag) Validate() error {

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
func (*CloudAlertRuleTargetTag) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudAlertRuleTargetTagAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudAlertRuleTargetTagLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudAlertRuleTargetTag) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudAlertRuleTargetTagAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudAlertRuleTargetTag) ValueForAttribute(name string) interface{} {

	switch name {
	case "key":
		return o.Key
	case "values":
		return o.Values
	}

	return nil
}

// CloudAlertRuleTargetTagAttributesMap represents the map of attribute for CloudAlertRuleTargetTag.
var CloudAlertRuleTargetTagAttributesMap = map[string]elemental.AttributeSpecification{
	"Key": {
		AllowedChoices: []string{},
		BSONFieldName:  "key",
		ConvertedName:  "Key",
		Description:    `Alert Rule target tag key.`,
		Exposed:        true,
		Name:           "key",
		Stored:         true,
		SubType:        "string",
		Type:           "string",
	},
	"Values": {
		AllowedChoices: []string{},
		BSONFieldName:  "values",
		ConvertedName:  "Values",
		Description:    `List of Alert Rule target tag values.`,
		Exposed:        true,
		Name:           "values",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
}

// CloudAlertRuleTargetTagLowerCaseAttributesMap represents the map of attribute for CloudAlertRuleTargetTag.
var CloudAlertRuleTargetTagLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"key": {
		AllowedChoices: []string{},
		BSONFieldName:  "key",
		ConvertedName:  "Key",
		Description:    `Alert Rule target tag key.`,
		Exposed:        true,
		Name:           "key",
		Stored:         true,
		SubType:        "string",
		Type:           "string",
	},
	"values": {
		AllowedChoices: []string{},
		BSONFieldName:  "values",
		ConvertedName:  "Values",
		Description:    `List of Alert Rule target tag values.`,
		Exposed:        true,
		Name:           "values",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
}

type mongoAttributesCloudAlertRuleTargetTag struct {
	Key    string   `bson:"key"`
	Values []string `bson:"values"`
}
