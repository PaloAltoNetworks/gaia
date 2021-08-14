package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// NetworkLabelRule represents the model of a networklabelrule
type NetworkLabelRule struct {
	// A user defined name to keep track of the rule in the reporting.
	Name string `json:"name,omitempty" msgpack:"name,omitempty" bson:"-" mapstructure:"name,omitempty"`

	// A list of IP CIDRS or FQDNS that identify remote endpoints.
	Networks []*NetworkRuleNet `json:"networks,omitempty" msgpack:"networks,omitempty" bson:"-" mapstructure:"networks,omitempty"`

	// Identifies the set of remote workloads that the rule relates to. The selector
	// will identify both processing units as well as external networks that match the
	// selector.
	Object [][]string `json:"object" msgpack:"object" bson:"-" mapstructure:"object,omitempty"`

	// Represents the ports and protocols this policy applies to. Protocol/ports are
	// defined as tcp/80, udp/22. For protocols that do not have ports, the port
	// designation
	// is not allowed.
	ProtocolPorts []string `json:"protocolPorts" msgpack:"protocolPorts" bson:"-" mapstructure:"protocolPorts,omitempty"`

	// Defines the SECMARK label to apply to packets.
	SecmarkLabel string `json:"secmarkLabel" msgpack:"secmarkLabel" bson:"-" mapstructure:"secmarkLabel,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewNetworkLabelRule returns a new *NetworkLabelRule
func NewNetworkLabelRule() *NetworkLabelRule {

	return &NetworkLabelRule{
		ModelVersion:  1,
		Networks:      []*NetworkRuleNet{},
		Object:        [][]string{},
		ProtocolPorts: []string{},
		SecmarkLabel:  "",
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NetworkLabelRule) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesNetworkLabelRule{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NetworkLabelRule) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesNetworkLabelRule{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *NetworkLabelRule) BleveType() string {

	return "networklabelrule"
}

// DeepCopy returns a deep copy if the NetworkLabelRule.
func (o *NetworkLabelRule) DeepCopy() *NetworkLabelRule {

	if o == nil {
		return nil
	}

	out := &NetworkLabelRule{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *NetworkLabelRule.
func (o *NetworkLabelRule) DeepCopyInto(out *NetworkLabelRule) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy NetworkLabelRule: %s", err))
	}

	*out = *target.(*NetworkLabelRule)
}

// Validate valides the current information stored into the structure.
func (o *NetworkLabelRule) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateMaximumLength("name", o.Name, 16, false); err != nil {
		errors = errors.Append(err)
	}

	for _, sub := range o.Networks {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := ValidateEachSubExpressionHasNoDuplicateTags("object", o.Object); err != nil {
		errors = errors.Append(err)
	}
	if err := ValidateExpressionNotEmpty("object", o.Object); err != nil {
		errors = errors.Append(err)
	}
	if err := ValidateNoDuplicateSubExpressions("object", o.Object); err != nil {
		errors = errors.Append(err)
	}
	if err := ValidateSubExpressionsNotEmpty("object", o.Object); err != nil {
		errors = errors.Append(err)
	}
	if err := ValidateTagsExpression("object", o.Object); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateServicePorts("protocolPorts", o.ProtocolPorts); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("secmarkLabel", o.SecmarkLabel); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

type mongoAttributesNetworkLabelRule struct {
}
