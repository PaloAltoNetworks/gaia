package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudNetworkRuleSetDataTypeValue represents the possible values for attribute "type".
type CloudNetworkRuleSetDataTypeValue string

const (
	// CloudNetworkRuleSetDataTypeACL represents the value ACL.
	CloudNetworkRuleSetDataTypeACL CloudNetworkRuleSetDataTypeValue = "ACL"

	// CloudNetworkRuleSetDataTypeEffectiveSecurityGroup represents the value EffectiveSecurityGroup.
	CloudNetworkRuleSetDataTypeEffectiveSecurityGroup CloudNetworkRuleSetDataTypeValue = "EffectiveSecurityGroup"

	// CloudNetworkRuleSetDataTypeSecurityGroup represents the value SecurityGroup.
	CloudNetworkRuleSetDataTypeSecurityGroup CloudNetworkRuleSetDataTypeValue = "SecurityGroup"
)

// CloudNetworkRuleSetData represents the model of a cloudnetworkrulesetdata
type CloudNetworkRuleSetData struct {
	// Internal field storing all the subject tags.
	AllSubjectTags []string `json:"-" msgpack:"-" bson:"allsubjecttags" mapstructure:"-,omitempty"`

	// The set of rules to apply to incoming traffic (traffic coming to the Processing
	// Unit matching the subject).
	IncomingRules []*CloudNetworkRule `json:"incomingRules" msgpack:"incomingRules" bson:"incomingrules" mapstructure:"incomingRules,omitempty"`

	// The set of rules to apply to outgoing traffic (traffic coming from the
	// Processing Unit matching the subject).
	OutgoingRules []*CloudNetworkRule `json:"outgoingRules" msgpack:"outgoingRules" bson:"outgoingrules" mapstructure:"outgoingRules,omitempty"`

	// A tag expression identifying used to match processing units to which this policy
	// applies to.
	Subject [][]string `json:"subject" msgpack:"subject" bson:"subject" mapstructure:"subject,omitempty"`

	// Type identifies if this is a security group rule set or ACL.
	Type CloudNetworkRuleSetDataTypeValue `json:"type" msgpack:"type" bson:"type" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudNetworkRuleSetData returns a new *CloudNetworkRuleSetData
func NewCloudNetworkRuleSetData() *CloudNetworkRuleSetData {

	return &CloudNetworkRuleSetData{
		ModelVersion:   1,
		AllSubjectTags: []string{},
		IncomingRules:  []*CloudNetworkRule{},
		OutgoingRules:  []*CloudNetworkRule{},
		Subject:        [][]string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudNetworkRuleSetData) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudNetworkRuleSetData{}

	s.AllSubjectTags = o.AllSubjectTags
	s.IncomingRules = o.IncomingRules
	s.OutgoingRules = o.OutgoingRules
	s.Subject = o.Subject
	s.Type = o.Type

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudNetworkRuleSetData) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudNetworkRuleSetData{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.AllSubjectTags = s.AllSubjectTags
	o.IncomingRules = s.IncomingRules
	o.OutgoingRules = s.OutgoingRules
	o.Subject = s.Subject
	o.Type = s.Type

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudNetworkRuleSetData) BleveType() string {

	return "cloudnetworkrulesetdata"
}

// DeepCopy returns a deep copy if the CloudNetworkRuleSetData.
func (o *CloudNetworkRuleSetData) DeepCopy() *CloudNetworkRuleSetData {

	if o == nil {
		return nil
	}

	out := &CloudNetworkRuleSetData{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudNetworkRuleSetData.
func (o *CloudNetworkRuleSetData) DeepCopyInto(out *CloudNetworkRuleSetData) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudNetworkRuleSetData: %s", err))
	}

	*out = *target.(*CloudNetworkRuleSetData)
}

// Validate valides the current information stored into the structure.
func (o *CloudNetworkRuleSetData) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.IncomingRules {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	for _, sub := range o.OutgoingRules {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := ValidateCloudTagsExpression("subject", o.Subject); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("type", string(o.Type)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"EffectiveSecurityGroup", "SecurityGroup", "ACL"}, false); err != nil {
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
func (*CloudNetworkRuleSetData) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudNetworkRuleSetDataAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudNetworkRuleSetDataLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudNetworkRuleSetData) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudNetworkRuleSetDataAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudNetworkRuleSetData) ValueForAttribute(name string) interface{} {

	switch name {
	case "allSubjectTags":
		return o.AllSubjectTags
	case "incomingRules":
		return o.IncomingRules
	case "outgoingRules":
		return o.OutgoingRules
	case "subject":
		return o.Subject
	case "type":
		return o.Type
	}

	return nil
}

// CloudNetworkRuleSetDataAttributesMap represents the map of attribute for CloudNetworkRuleSetData.
var CloudNetworkRuleSetDataAttributesMap = map[string]elemental.AttributeSpecification{
	"AllSubjectTags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "allsubjecttags",
		ConvertedName:  "AllSubjectTags",
		Description:    `Internal field storing all the subject tags.`,
		Name:           "allSubjectTags",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"IncomingRules": {
		AllowedChoices: []string{},
		BSONFieldName:  "incomingrules",
		ConvertedName:  "IncomingRules",
		Description: `The set of rules to apply to incoming traffic (traffic coming to the Processing
Unit matching the subject).`,
		Exposed: true,
		Name:    "incomingRules",
		Stored:  true,
		SubType: "cloudnetworkrule",
		Type:    "refList",
	},
	"OutgoingRules": {
		AllowedChoices: []string{},
		BSONFieldName:  "outgoingrules",
		ConvertedName:  "OutgoingRules",
		Description: `The set of rules to apply to outgoing traffic (traffic coming from the
Processing Unit matching the subject).`,
		Exposed: true,
		Name:    "outgoingRules",
		Stored:  true,
		SubType: "cloudnetworkrule",
		Type:    "refList",
	},
	"Subject": {
		AllowedChoices: []string{},
		BSONFieldName:  "subject",
		ConvertedName:  "Subject",
		Description: `A tag expression identifying used to match processing units to which this policy
applies to.`,
		Exposed: true,
		Name:    "subject",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"Type": {
		AllowedChoices: []string{"EffectiveSecurityGroup", "SecurityGroup", "ACL"},
		BSONFieldName:  "type",
		ConvertedName:  "Type",
		Description:    `Type identifies if this is a security group rule set or ACL.`,
		Exposed:        true,
		Name:           "type",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
}

// CloudNetworkRuleSetDataLowerCaseAttributesMap represents the map of attribute for CloudNetworkRuleSetData.
var CloudNetworkRuleSetDataLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"allsubjecttags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "allsubjecttags",
		ConvertedName:  "AllSubjectTags",
		Description:    `Internal field storing all the subject tags.`,
		Name:           "allSubjectTags",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"incomingrules": {
		AllowedChoices: []string{},
		BSONFieldName:  "incomingrules",
		ConvertedName:  "IncomingRules",
		Description: `The set of rules to apply to incoming traffic (traffic coming to the Processing
Unit matching the subject).`,
		Exposed: true,
		Name:    "incomingRules",
		Stored:  true,
		SubType: "cloudnetworkrule",
		Type:    "refList",
	},
	"outgoingrules": {
		AllowedChoices: []string{},
		BSONFieldName:  "outgoingrules",
		ConvertedName:  "OutgoingRules",
		Description: `The set of rules to apply to outgoing traffic (traffic coming from the
Processing Unit matching the subject).`,
		Exposed: true,
		Name:    "outgoingRules",
		Stored:  true,
		SubType: "cloudnetworkrule",
		Type:    "refList",
	},
	"subject": {
		AllowedChoices: []string{},
		BSONFieldName:  "subject",
		ConvertedName:  "Subject",
		Description: `A tag expression identifying used to match processing units to which this policy
applies to.`,
		Exposed: true,
		Name:    "subject",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"type": {
		AllowedChoices: []string{"EffectiveSecurityGroup", "SecurityGroup", "ACL"},
		BSONFieldName:  "type",
		ConvertedName:  "Type",
		Description:    `Type identifies if this is a security group rule set or ACL.`,
		Exposed:        true,
		Name:           "type",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
}

type mongoAttributesCloudNetworkRuleSetData struct {
	AllSubjectTags []string                         `bson:"allsubjecttags"`
	IncomingRules  []*CloudNetworkRule              `bson:"incomingrules"`
	OutgoingRules  []*CloudNetworkRule              `bson:"outgoingrules"`
	Subject        [][]string                       `bson:"subject"`
	Type           CloudNetworkRuleSetDataTypeValue `bson:"type"`
}
