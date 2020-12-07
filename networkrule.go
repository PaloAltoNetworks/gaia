package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// NetworkRuleActionValue represents the possible values for attribute "action".
type NetworkRuleActionValue string

const (
	// NetworkRuleActionAllow represents the value Allow.
	NetworkRuleActionAllow NetworkRuleActionValue = "Allow"

	// NetworkRuleActionReject represents the value Reject.
	NetworkRuleActionReject NetworkRuleActionValue = "Reject"
)

// NetworkRule represents the model of a networkrule
type NetworkRule struct {
	// Defines the action to apply to a flow.
	// - `Allow`: allows the defined traffic.
	// - `Reject`: rejects the defined traffic; useful in conjunction with an allow all
	// policy.
	Action NetworkRuleActionValue `json:"action" msgpack:"action" bson:"-" mapstructure:"action,omitempty"`

	// A list of IP CIDRS or FQDNS that identify remote endpoints.
	Networks []string `json:"networks" msgpack:"networks" bson:"-" mapstructure:"networks,omitempty"`

	// If set to `true`, the flow will be in observation mode.
	ObservationEnabled bool `json:"observationEnabled" msgpack:"observationEnabled" bson:"-" mapstructure:"observationEnabled,omitempty"`

	// Represents the ports and protocols this policy applies to. Protocol/ports are
	// defined as tcp/80, udp/22. For protocols that do not have ports, the port
	// designation
	// is not allowed.
	ProtocolPorts []string `json:"protocolPorts" msgpack:"protocolPorts" bson:"-" mapstructure:"protocolPorts,omitempty"`

	// Identifies the set of remote workloads that the rule relates to. The selector
	// will identify both processing units as well as external networks that match the
	// selector.
	Selector [][]string `json:"selector" msgpack:"selector" bson:"-" mapstructure:"selector,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewNetworkRule returns a new *NetworkRule
func NewNetworkRule() *NetworkRule {

	return &NetworkRule{
		ModelVersion:       1,
		Action:             NetworkRuleActionAllow,
		Networks:           []string{},
		ObservationEnabled: false,
		ProtocolPorts:      []string{},
		Selector:           [][]string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NetworkRule) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesNetworkRule{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NetworkRule) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesNetworkRule{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *NetworkRule) BleveType() string {

	return "networkrule"
}

// DeepCopy returns a deep copy if the NetworkRule.
func (o *NetworkRule) DeepCopy() *NetworkRule {

	if o == nil {
		return nil
	}

	out := &NetworkRule{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *NetworkRule.
func (o *NetworkRule) DeepCopyInto(out *NetworkRule) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy NetworkRule: %s", err))
	}

	*out = *target.(*NetworkRule)
}

// Validate valides the current information stored into the structure.
func (o *NetworkRule) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("action", string(o.Action)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("action", string(o.Action), []string{"Allow", "Reject"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateServicePorts("protocolPorts", o.ProtocolPorts); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsExpression("selector", o.Selector); err != nil {
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

type mongoAttributesNetworkRule struct {
}