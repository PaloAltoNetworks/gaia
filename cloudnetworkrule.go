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

	// A list of tags to group resources.
	GroupTags []string `json:"groupTags,omitempty" msgpack:"groupTags,omitempty" bson:"grouptags,omitempty" mapstructure:"groupTags,omitempty"`

	// A list of ips range.
	IpRanges []*CloudIPRange `json:"ipRanges,omitempty" msgpack:"ipRanges,omitempty" bson:"ipranges,omitempty" mapstructure:"ipRanges,omitempty"`

	// A list of IP CIDRS that identify local endpoints.
	LocalNetworks []string `json:"localNetworks,omitempty" msgpack:"localNetworks,omitempty" bson:"localnetworks,omitempty" mapstructure:"localNetworks,omitempty"`

	// A list of Service Tags provided by the platform.
	LocalServiceTags []string `json:"localServiceTags,omitempty" msgpack:"localServiceTags,omitempty" bson:"localservicetags,omitempty" mapstructure:"localServiceTags,omitempty"`

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

	// A list of Service Tags provided by the platform.
	ServiceTags []string `json:"serviceTags,omitempty" msgpack:"serviceTags,omitempty" bson:"servicetags,omitempty" mapstructure:"serviceTags,omitempty"`

	// A list of ips range for internal use. Not visible to end users.
	StoredIPRanges []*CloudStoredIPRange `json:"storedIPRanges,omitempty" msgpack:"storedIPRanges,omitempty" bson:"storedipranges,omitempty" mapstructure:"storedIPRanges,omitempty"`

	// An internal representation of the local networks to increase performance. Not
	// visible
	// to end users.
	StoredLocalNetworks []*net.IPNet `json:"storedLocalNetworks,omitempty" msgpack:"storedLocalNetworks,omitempty" bson:"storedlocalnetworks,omitempty" mapstructure:"storedLocalNetworks,omitempty"`

	// An internal representation of the networks to increase performance. Not visible
	// to end users.
	StoredNetworks []*net.IPNet `json:"storedNetworks,omitempty" msgpack:"storedNetworks,omitempty" bson:"storednetworks,omitempty" mapstructure:"storedNetworks,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudNetworkRule returns a new *CloudNetworkRule
func NewCloudNetworkRule() *CloudNetworkRule {

	return &CloudNetworkRule{
		ModelVersion:        1,
		Object:              [][]string{},
		GroupTags:           []string{},
		IpRanges:            []*CloudIPRange{},
		LocalNetworks:       []string{},
		LocalServiceTags:    []string{},
		Networks:            []string{},
		Action:              CloudNetworkRuleActionAllow,
		ProtocolPorts:       []string{},
		ServiceTags:         []string{},
		StoredIPRanges:      []*CloudStoredIPRange{},
		StoredLocalNetworks: []*net.IPNet{},
		StoredNetworks:      []*net.IPNet{},
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
	s.GroupTags = o.GroupTags
	s.IpRanges = o.IpRanges
	s.LocalNetworks = o.LocalNetworks
	s.LocalServiceTags = o.LocalServiceTags
	s.Networks = o.Networks
	s.Object = o.Object
	s.Priority = o.Priority
	s.ProtocolPorts = o.ProtocolPorts
	s.ServiceTags = o.ServiceTags
	s.StoredIPRanges = o.StoredIPRanges
	s.StoredLocalNetworks = o.StoredLocalNetworks
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
	o.GroupTags = s.GroupTags
	o.IpRanges = s.IpRanges
	o.LocalNetworks = s.LocalNetworks
	o.LocalServiceTags = s.LocalServiceTags
	o.Networks = s.Networks
	o.Object = s.Object
	o.Priority = s.Priority
	o.ProtocolPorts = s.ProtocolPorts
	o.ServiceTags = s.ServiceTags
	o.StoredIPRanges = s.StoredIPRanges
	o.StoredLocalNetworks = s.StoredLocalNetworks
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

	for _, sub := range o.IpRanges {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := ValidateOptionalCIDRorIPList("localNetworks", o.LocalNetworks); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateOptionalCIDRorIPList("networks", o.Networks); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateServicePorts("protocolPorts", o.ProtocolPorts); err != nil {
		errors = errors.Append(err)
	}

	for _, sub := range o.StoredIPRanges {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
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
	case "groupTags":
		return o.GroupTags
	case "ipRanges":
		return o.IpRanges
	case "localNetworks":
		return o.LocalNetworks
	case "localServiceTags":
		return o.LocalServiceTags
	case "networks":
		return o.Networks
	case "object":
		return o.Object
	case "priority":
		return o.Priority
	case "protocolPorts":
		return o.ProtocolPorts
	case "serviceTags":
		return o.ServiceTags
	case "storedIPRanges":
		return o.StoredIPRanges
	case "storedLocalNetworks":
		return o.StoredLocalNetworks
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
	"GroupTags": {
		AllowedChoices: []string{},
		BSONFieldName:  "grouptags",
		ConvertedName:  "GroupTags",
		Description:    `A list of tags to group resources.`,
		Exposed:        true,
		Name:           "groupTags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"IpRanges": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "ipranges",
		ConvertedName:  "IpRanges",
		Description:    `A list of ips range.`,
		Exposed:        true,
		Name:           "ipRanges",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "cloudiprange",
		Type:           "refList",
	},
	"LocalNetworks": {
		AllowedChoices: []string{},
		BSONFieldName:  "localnetworks",
		ConvertedName:  "LocalNetworks",
		Description:    `A list of IP CIDRS that identify local endpoints.`,
		Exposed:        true,
		Name:           "localNetworks",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"LocalServiceTags": {
		AllowedChoices: []string{},
		BSONFieldName:  "localservicetags",
		ConvertedName:  "LocalServiceTags",
		Description:    `A list of Service Tags provided by the platform.`,
		Exposed:        true,
		Name:           "localServiceTags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"ServiceTags": {
		AllowedChoices: []string{},
		BSONFieldName:  "servicetags",
		ConvertedName:  "ServiceTags",
		Description:    `A list of Service Tags provided by the platform.`,
		Exposed:        true,
		Name:           "serviceTags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"StoredIPRanges": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "storedipranges",
		ConvertedName:  "StoredIPRanges",
		Description:    `A list of ips range for internal use. Not visible to end users.`,
		Exposed:        true,
		Name:           "storedIPRanges",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "cloudstorediprange",
		Type:           "refList",
	},
	"StoredLocalNetworks": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "storedlocalnetworks",
		ConvertedName:  "StoredLocalNetworks",
		Description: `An internal representation of the local networks to increase performance. Not
visible
to end users.`,
		Exposed:  true,
		Name:     "storedLocalNetworks",
		ReadOnly: true,
		Stored:   true,
		SubType:  "networklist",
		Type:     "external",
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
	"grouptags": {
		AllowedChoices: []string{},
		BSONFieldName:  "grouptags",
		ConvertedName:  "GroupTags",
		Description:    `A list of tags to group resources.`,
		Exposed:        true,
		Name:           "groupTags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"ipranges": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "ipranges",
		ConvertedName:  "IpRanges",
		Description:    `A list of ips range.`,
		Exposed:        true,
		Name:           "ipRanges",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "cloudiprange",
		Type:           "refList",
	},
	"localnetworks": {
		AllowedChoices: []string{},
		BSONFieldName:  "localnetworks",
		ConvertedName:  "LocalNetworks",
		Description:    `A list of IP CIDRS that identify local endpoints.`,
		Exposed:        true,
		Name:           "localNetworks",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"localservicetags": {
		AllowedChoices: []string{},
		BSONFieldName:  "localservicetags",
		ConvertedName:  "LocalServiceTags",
		Description:    `A list of Service Tags provided by the platform.`,
		Exposed:        true,
		Name:           "localServiceTags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"servicetags": {
		AllowedChoices: []string{},
		BSONFieldName:  "servicetags",
		ConvertedName:  "ServiceTags",
		Description:    `A list of Service Tags provided by the platform.`,
		Exposed:        true,
		Name:           "serviceTags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"storedipranges": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "storedipranges",
		ConvertedName:  "StoredIPRanges",
		Description:    `A list of ips range for internal use. Not visible to end users.`,
		Exposed:        true,
		Name:           "storedIPRanges",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "cloudstorediprange",
		Type:           "refList",
	},
	"storedlocalnetworks": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "storedlocalnetworks",
		ConvertedName:  "StoredLocalNetworks",
		Description: `An internal representation of the local networks to increase performance. Not
visible
to end users.`,
		Exposed:  true,
		Name:     "storedLocalNetworks",
		ReadOnly: true,
		Stored:   true,
		SubType:  "networklist",
		Type:     "external",
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
	Action              CloudNetworkRuleActionValue `bson:"action"`
	GroupTags           []string                    `bson:"grouptags,omitempty"`
	IpRanges            []*CloudIPRange             `bson:"ipranges,omitempty"`
	LocalNetworks       []string                    `bson:"localnetworks,omitempty"`
	LocalServiceTags    []string                    `bson:"localservicetags,omitempty"`
	Networks            []string                    `bson:"networks,omitempty"`
	Object              [][]string                  `bson:"object"`
	Priority            int                         `bson:"priority,omitempty"`
	ProtocolPorts       []string                    `bson:"protocolports"`
	ServiceTags         []string                    `bson:"servicetags,omitempty"`
	StoredIPRanges      []*CloudStoredIPRange       `bson:"storedipranges,omitempty"`
	StoredLocalNetworks []*net.IPNet                `bson:"storedlocalnetworks,omitempty"`
	StoredNetworks      []*net.IPNet                `bson:"storednetworks,omitempty"`
}
