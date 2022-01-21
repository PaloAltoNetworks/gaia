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
	Networks []string `json:"networks,omitempty" msgpack:"networks,omitempty" bson:"networks,omitempty" mapstructure:"networks,omitempty"`

	// Identifies the set of remote workloads that the rule relates to. The selector
	// will identify both processing units as well as external networks that match the
	// selector.
	Object [][]string `json:"object" msgpack:"object" bson:"object" mapstructure:"object,omitempty"`

	// Priority of the rule. Available only for cloud ACLs.
	Priority int `json:"priority,omitempty" msgpack:"priority,omitempty" bson:"priority,omitempty" mapstructure:"priority,omitempty"`

	// Represents the ports and protocols this policy applies to. Protocol/ports are
	// defined as tcp/80, udp/22. For protocols that do not have ports, the port
	// designation
	// is not allowed.
	ProtocolPorts []string `json:"protocolPorts" msgpack:"protocolPorts" bson:"protocolports" mapstructure:"protocolPorts,omitempty"`

	// An internal representation of the networks to increase performance. Not visible
	// to end users.
	StoredNetworks []*net.IPNet `json:"storedNetworks,omitempty" msgpack:"storedNetworks,omitempty" bson:"storednetworks,omitempty" mapstructure:"storedNetworks,omitempty"`

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

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*CloudNetworkRule) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudNetworkRuleAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudNetworkRuleLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudNetworkRule) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudNetworkRuleAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudNetworkRule) ValueForAttribute(name string) interface{} {

	switch name {
	case "action":
		return o.Action
	case "networks":
		return o.Networks
	case "object":
		return o.Object
	case "priority":
		return o.Priority
	case "protocolPorts":
		return o.ProtocolPorts
	case "storedNetworks":
		return o.StoredNetworks
	}

	return nil
}

// CloudNetworkRuleAttributesMap represents the map of attribute for CloudNetworkRule.
var CloudNetworkRuleAttributesMap = map[string]elemental.AttributeSpecification{
	"Action": {
		AllowedChoices: []string{"Allow", "Reject"},
		BSONFieldName:  "action",
		ConvertedName:  "Action",
		DefaultValue:   CloudNetworkRuleActionAllow,
		Description: `Defines the action to apply to a flow.
- ` + "`" + `Allow` + "`" + `: allows the defined traffic.
- ` + "`" + `Reject` + "`" + `: rejects the defined traffic; useful in conjunction with an allow all
policy.`,
		Exposed:  true,
		Name:     "action",
		Required: true,
		Stored:   true,
		Type:     "enum",
	},
	"Networks": {
		AllowedChoices: []string{},
		BSONFieldName:  "networks",
		ConvertedName:  "Networks",
		Description:    `A list of IP CIDRS that identify remote endpoints.`,
		Exposed:        true,
		Name:           "networks",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Object": {
		AllowedChoices: []string{},
		BSONFieldName:  "object",
		ConvertedName:  "Object",
		Description: `Identifies the set of remote workloads that the rule relates to. The selector
will identify both processing units as well as external networks that match the
selector.`,
		Exposed:   true,
		Name:      "object",
		Orderable: true,
		Stored:    true,
		SubType:   "[][]string",
		Type:      "external",
	},
	"Priority": {
		AllowedChoices: []string{},
		BSONFieldName:  "priority",
		ConvertedName:  "Priority",
		Description:    `Priority of the rule. Available only for cloud ACLs.`,
		Exposed:        true,
		Name:           "priority",
		Stored:         true,
		Type:           "integer",
	},
	"ProtocolPorts": {
		AllowedChoices: []string{},
		BSONFieldName:  "protocolports",
		ConvertedName:  "ProtocolPorts",
		Description: `Represents the ports and protocols this policy applies to. Protocol/ports are
defined as tcp/80, udp/22. For protocols that do not have ports, the port
designation
is not allowed.`,
		Exposed: true,
		Name:    "protocolPorts",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"StoredNetworks": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "storednetworks",
		ConvertedName:  "StoredNetworks",
		Description: `An internal representation of the networks to increase performance. Not visible
to end users.`,
		Exposed:  true,
		Name:     "storedNetworks",
		ReadOnly: true,
		Stored:   true,
		SubType:  "networklist",
		Type:     "external",
	},
}

// CloudNetworkRuleLowerCaseAttributesMap represents the map of attribute for CloudNetworkRule.
var CloudNetworkRuleLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"action": {
		AllowedChoices: []string{"Allow", "Reject"},
		BSONFieldName:  "action",
		ConvertedName:  "Action",
		DefaultValue:   CloudNetworkRuleActionAllow,
		Description: `Defines the action to apply to a flow.
- ` + "`" + `Allow` + "`" + `: allows the defined traffic.
- ` + "`" + `Reject` + "`" + `: rejects the defined traffic; useful in conjunction with an allow all
policy.`,
		Exposed:  true,
		Name:     "action",
		Required: true,
		Stored:   true,
		Type:     "enum",
	},
	"networks": {
		AllowedChoices: []string{},
		BSONFieldName:  "networks",
		ConvertedName:  "Networks",
		Description:    `A list of IP CIDRS that identify remote endpoints.`,
		Exposed:        true,
		Name:           "networks",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"object": {
		AllowedChoices: []string{},
		BSONFieldName:  "object",
		ConvertedName:  "Object",
		Description: `Identifies the set of remote workloads that the rule relates to. The selector
will identify both processing units as well as external networks that match the
selector.`,
		Exposed:   true,
		Name:      "object",
		Orderable: true,
		Stored:    true,
		SubType:   "[][]string",
		Type:      "external",
	},
	"priority": {
		AllowedChoices: []string{},
		BSONFieldName:  "priority",
		ConvertedName:  "Priority",
		Description:    `Priority of the rule. Available only for cloud ACLs.`,
		Exposed:        true,
		Name:           "priority",
		Stored:         true,
		Type:           "integer",
	},
	"protocolports": {
		AllowedChoices: []string{},
		BSONFieldName:  "protocolports",
		ConvertedName:  "ProtocolPorts",
		Description: `Represents the ports and protocols this policy applies to. Protocol/ports are
defined as tcp/80, udp/22. For protocols that do not have ports, the port
designation
is not allowed.`,
		Exposed: true,
		Name:    "protocolPorts",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"storednetworks": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "storednetworks",
		ConvertedName:  "StoredNetworks",
		Description: `An internal representation of the networks to increase performance. Not visible
to end users.`,
		Exposed:  true,
		Name:     "storedNetworks",
		ReadOnly: true,
		Stored:   true,
		SubType:  "networklist",
		Type:     "external",
	},
}

type mongoAttributesCloudNetworkRule struct {
	Action         CloudNetworkRuleActionValue `bson:"action"`
	Networks       []string                    `bson:"networks,omitempty"`
	Object         [][]string                  `bson:"object"`
	Priority       int                         `bson:"priority,omitempty"`
	ProtocolPorts  []string                    `bson:"protocolports"`
	StoredNetworks []*net.IPNet                `bson:"storednetworks,omitempty"`
}
