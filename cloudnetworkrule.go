package gaia

import (
	"fmt"
	"net"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudNetworkRuleActionValue represents the possible values for attribute "action".
type CloudNetworkRuleActionValue string

const (
	// CloudNetworkRuleActionAllow represents the value Allow.
	CloudNetworkRuleActionAllow CloudNetworkRuleActionValue = "Allow"

	// CloudNetworkRuleActionReject represents the value Reject.
	CloudNetworkRuleActionReject CloudNetworkRuleActionValue = "Reject"
)

// CloudNetworkRule represents the model of a cloudnetworkrule
type CloudNetworkRule struct {
	// Defines the action to apply to a flow.
	// - `Allow`: allows the defined traffic.
	// - `Reject`: rejects the defined traffic; useful in conjunction with an allow all
	// policy.
	Action CloudNetworkRuleActionValue `json:"action" msgpack:"action" bson:"action" mapstructure:"action,omitempty"`

	// A list of IP CIDRS that identify remote endpoints.
	Networks []string `json:"networks,omitempty" msgpack:"networks,omitempty" bson:"networks" mapstructure:"networks,omitempty"`

	// Identifies the set of remote workloads that the rule relates to. The selector
	// will identify both processing units as well as external networks that match the
	// selector.
	Object [][]string `json:"object" msgpack:"object" bson:"object" mapstructure:"object,omitempty"`

	// Priority of the rule. Available only for cloud ACLs.
	Priority int `json:"priority,omitempty" msgpack:"priority,omitempty" bson:"priority" mapstructure:"priority,omitempty"`

	// Represents the ports and protocols this policy applies to. Protocol/ports are
	// defined as tcp/80, udp/22. For protocols that do not have ports, the port
	// designation
	// is not allowed.
	ProtocolPorts []string `json:"protocolPorts" msgpack:"protocolPorts" bson:"protocolports" mapstructure:"protocolPorts,omitempty"`

	// An internal representation of the networks to increase performance. Not visible
	// to end users.
	StoredNetworks []*net.IPNet `json:"storedNetworks,omitempty" msgpack:"storedNetworks,omitempty" bson:"storednetworks" mapstructure:"storedNetworks,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudNetworkRule returns a new *CloudNetworkRule
func NewCloudNetworkRule() *CloudNetworkRule {

	return &CloudNetworkRule{
		ModelVersion:   1,
		Action:         CloudNetworkRuleActionAllow,
		Networks:       []string{},
		Object:         [][]string{},
		ProtocolPorts:  []string{},
		StoredNetworks: []*net.IPNet{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudNetworkRule) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudNetworkRule{}

	s.Action = o.Action
	s.Networks = o.Networks
	s.Object = o.Object
	s.Priority = o.Priority
	s.ProtocolPorts = o.ProtocolPorts
	s.StoredNetworks = o.StoredNetworks

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudNetworkRule) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudNetworkRule{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Action = s.Action
	o.Networks = s.Networks
	o.Object = s.Object
	o.Priority = s.Priority
	o.ProtocolPorts = s.ProtocolPorts
	o.StoredNetworks = s.StoredNetworks

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudNetworkRule) BleveType() string {

	return "cloudnetworkrule"
}

// DeepCopy returns a deep copy if the CloudNetworkRule.
func (o *CloudNetworkRule) DeepCopy() *CloudNetworkRule {

	if o == nil {
		return nil
	}

	out := &CloudNetworkRule{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudNetworkRule.
func (o *CloudNetworkRule) DeepCopyInto(out *CloudNetworkRule) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudNetworkRule: %s", err))
	}

	*out = *target.(*CloudNetworkRule)
}

// Validate valides the current information stored into the structure.
func (o *CloudNetworkRule) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("action", string(o.Action)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("action", string(o.Action), []string{"Allow", "Reject"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateOptionalCIDRorIPList("networks", o.Networks); err != nil {
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

type mongoAttributesCloudNetworkRule struct {
	Action         CloudNetworkRuleActionValue `bson:"action"`
	Networks       []string                    `bson:"networks,omitempty"`
	Object         [][]string                  `bson:"object"`
	Priority       int                         `bson:"priority,omitempty"`
	ProtocolPorts  []string                    `bson:"protocolports"`
	StoredNetworks []*net.IPNet                `bson:"storednetworks,omitempty"`
}
