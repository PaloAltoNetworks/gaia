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

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"SecurityGroup", "ACL"}, false); err != nil {
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

type mongoAttributesCloudNetworkRuleSetData struct {
	AllSubjectTags []string                         `bson:"allsubjecttags"`
	IncomingRules  []*CloudNetworkRule              `bson:"incomingrules"`
	OutgoingRules  []*CloudNetworkRule              `bson:"outgoingrules"`
	Subject        [][]string                       `bson:"subject"`
	Type           CloudNetworkRuleSetDataTypeValue `bson:"type"`
}
