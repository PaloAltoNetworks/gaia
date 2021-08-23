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
// +k8s:openapi-gen=true
type NetworkRule struct {
	// Defines the action to apply to a flow.
	// - `Allow`: allows the defined traffic.
	// - `Reject`: rejects the defined traffic; useful in conjunction with an allow all
	// policy.
	Action NetworkRuleActionValue `json:"action" msgpack:"action" bson:"-" mapstructure:"action,omitempty"`

	// If `true`, the relevant flows will not be reported to the Microsegmentation
	// Console.
	// Under some advanced scenarios you may wish to set this to `true`, such as to
	// save space or improve performance.
	LogsDisabled bool `json:"logsDisabled" msgpack:"logsDisabled" bson:"-" mapstructure:"logsDisabled,omitempty"`

	// A user defined name to keep track of the rule in the reporting.
	Name string `json:"name,omitempty" msgpack:"name,omitempty" bson:"-" mapstructure:"name,omitempty"`

	// A list of IP CIDRS or FQDNS that identify remote endpoints.
	Networks []*NetworkRuleNet `json:"networks,omitempty" msgpack:"networks,omitempty" bson:"-" mapstructure:"networks,omitempty"`

	// Identifies the set of remote workloads that the rule relates to. The selector
	// will identify both processing units as well as external networks that match the
	// selector.
	Object [][]string `json:"object" msgpack:"object" bson:"-" mapstructure:"object,omitempty"`

	// If set to `true`, the flow will be in observation mode.
	ObservationEnabled bool `json:"observationEnabled" msgpack:"observationEnabled" bson:"-" mapstructure:"observationEnabled,omitempty"`

	// Represents the ports and protocols this policy applies to. Protocol/ports are
	// defined as tcp/80, udp/22. For protocols that do not have ports, the port
	// designation
	// is not allowed.
	ProtocolPorts []string `json:"protocolPorts" msgpack:"protocolPorts" bson:"-" mapstructure:"protocolPorts,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewNetworkRule returns a new *NetworkRule
func NewNetworkRule() *NetworkRule {

	return &NetworkRule{
		ModelVersion:       1,
		Action:             NetworkRuleActionAllow,
		Networks:           []*NetworkRuleNet{},
		Object:             [][]string{},
		ObservationEnabled: false,
		ProtocolPorts:      []string{},
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
