package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudInterfaceDataAttachmentTypeValue represents the possible values for attribute "attachmentType".
type CloudInterfaceDataAttachmentTypeValue string

const (
	// CloudInterfaceDataAttachmentTypeAPIGatewayManaged represents the value APIGatewayManaged.
	CloudInterfaceDataAttachmentTypeAPIGatewayManaged CloudInterfaceDataAttachmentTypeValue = "APIGatewayManaged"

	// CloudInterfaceDataAttachmentTypeEFA represents the value EFA.
	CloudInterfaceDataAttachmentTypeEFA CloudInterfaceDataAttachmentTypeValue = "EFA"

	// CloudInterfaceDataAttachmentTypeGateway represents the value Gateway.
	CloudInterfaceDataAttachmentTypeGateway CloudInterfaceDataAttachmentTypeValue = "Gateway"

	// CloudInterfaceDataAttachmentTypeGatewayLoadBalancer represents the value GatewayLoadBalancer.
	CloudInterfaceDataAttachmentTypeGatewayLoadBalancer CloudInterfaceDataAttachmentTypeValue = "GatewayLoadBalancer"

	// CloudInterfaceDataAttachmentTypeGatewayLoadBalancerEndpoint represents the value GatewayLoadBalancerEndpoint.
	CloudInterfaceDataAttachmentTypeGatewayLoadBalancerEndpoint CloudInterfaceDataAttachmentTypeValue = "GatewayLoadBalancerEndpoint"

	// CloudInterfaceDataAttachmentTypeInstance represents the value Instance.
	CloudInterfaceDataAttachmentTypeInstance CloudInterfaceDataAttachmentTypeValue = "Instance"

	// CloudInterfaceDataAttachmentTypeLambda represents the value Lambda.
	CloudInterfaceDataAttachmentTypeLambda CloudInterfaceDataAttachmentTypeValue = "Lambda"

	// CloudInterfaceDataAttachmentTypeLoadBalancer represents the value LoadBalancer.
	CloudInterfaceDataAttachmentTypeLoadBalancer CloudInterfaceDataAttachmentTypeValue = "LoadBalancer"

	// CloudInterfaceDataAttachmentTypeNetworkLoadBalancer represents the value NetworkLoadBalancer.
	CloudInterfaceDataAttachmentTypeNetworkLoadBalancer CloudInterfaceDataAttachmentTypeValue = "NetworkLoadBalancer"

	// CloudInterfaceDataAttachmentTypeService represents the value Service.
	CloudInterfaceDataAttachmentTypeService CloudInterfaceDataAttachmentTypeValue = "Service"

	// CloudInterfaceDataAttachmentTypeTransitGatewayVPCAttachment represents the value TransitGatewayVPCAttachment.
	CloudInterfaceDataAttachmentTypeTransitGatewayVPCAttachment CloudInterfaceDataAttachmentTypeValue = "TransitGatewayVPCAttachment"

	// CloudInterfaceDataAttachmentTypeUnsupportedService represents the value UnsupportedService.
	CloudInterfaceDataAttachmentTypeUnsupportedService CloudInterfaceDataAttachmentTypeValue = "UnsupportedService"

	// CloudInterfaceDataAttachmentTypeVPCEndpoint represents the value VPCEndpoint.
	CloudInterfaceDataAttachmentTypeVPCEndpoint CloudInterfaceDataAttachmentTypeValue = "VPCEndpoint"
)

// CloudInterfaceDataResourceStatusValue represents the possible values for attribute "resourceStatus".
type CloudInterfaceDataResourceStatusValue string

const (
	// CloudInterfaceDataResourceStatusActive represents the value Active.
	CloudInterfaceDataResourceStatusActive CloudInterfaceDataResourceStatusValue = "Active"

	// CloudInterfaceDataResourceStatusInactive represents the value Inactive.
	CloudInterfaceDataResourceStatusInactive CloudInterfaceDataResourceStatusValue = "Inactive"
)

// CloudInterfaceData represents the model of a cloudinterfacedata
type CloudInterfaceData struct {
	// List of IP addresses/subnets (IPv4 or IPv6) associated with the
	// interface.
	Addresses []*CloudAddress `json:"addresses" msgpack:"addresses" bson:"addresses" mapstructure:"addresses,omitempty"`

	// ID of associated objects with this interface.
	AttachedEntities []string `json:"attachedEntities" msgpack:"attachedEntities" bson:"attachedentities" mapstructure:"attachedEntities,omitempty"`

	// Attachment type describes where this interface is attached to (Instance, Load
	// Balancer, Gateway, etc).
	AttachmentType CloudInterfaceDataAttachmentTypeValue `json:"attachmentType" msgpack:"attachmentType" bson:"attachmenttype" mapstructure:"attachmentType,omitempty"`

	// Availability zone of the interface.
	AvailabilityZone string `json:"availabilityZone" msgpack:"availabilityZone" bson:"availabilityzone" mapstructure:"availabilityZone,omitempty"`

	// Group tags associated with the interface. In Azure, its Application Security
	// Group.
	GroupTags []string `json:"groupTags" msgpack:"groupTags" bson:"grouptags" mapstructure:"groupTags,omitempty"`

	// If the interface has a public IP in one of its IP address.
	HasPublicIP bool `json:"hasPublicIP" msgpack:"hasPublicIP" bson:"haspublicip" mapstructure:"hasPublicIP,omitempty"`

	// If the interface is attached to Load Balancer, the parentID identifies the
	// related Load Balancer.
	ParentID string `json:"parentID" msgpack:"parentID" bson:"parentid" mapstructure:"parentID,omitempty"`

	// If the interface is of type or external, the relatedObjectID identifies the
	// related service or gateway.
	RelatedObjectID string `json:"relatedObjectID" msgpack:"relatedObjectID" bson:"relatedobjectid" mapstructure:"relatedObjectID,omitempty"`

	// The status of the resource.
	ResourceStatus CloudInterfaceDataResourceStatusValue `json:"resourceStatus" msgpack:"resourceStatus" bson:"resourcestatus" mapstructure:"resourceStatus,omitempty"`

	// The route table that must be used for this interface. Applies to Transit
	// Gateways and other special types.
	RouteTableID string `json:"routeTableID" msgpack:"routeTableID" bson:"routetableid" mapstructure:"routeTableID,omitempty"`

	// Security tags associated with the instance.
	SecurityTags []string `json:"securityTags" msgpack:"securityTags" bson:"securitytags" mapstructure:"securityTags,omitempty"`

	// ID of subnet associated with this interface.
	Subnets []string `json:"subnets" msgpack:"subnets" bson:"subnets" mapstructure:"subnets,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudInterfaceData returns a new *CloudInterfaceData
func NewCloudInterfaceData() *CloudInterfaceData {

	return &CloudInterfaceData{
		ModelVersion:     1,
		Addresses:        []*CloudAddress{},
		AttachedEntities: []string{},
		GroupTags:        []string{},
		ResourceStatus:   CloudInterfaceDataResourceStatusActive,
		SecurityTags:     []string{},
		Subnets:          []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudInterfaceData) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudInterfaceData{}

	s.Addresses = o.Addresses
	s.AttachedEntities = o.AttachedEntities
	s.AttachmentType = o.AttachmentType
	s.AvailabilityZone = o.AvailabilityZone
	s.GroupTags = o.GroupTags
	s.HasPublicIP = o.HasPublicIP
	s.ParentID = o.ParentID
	s.RelatedObjectID = o.RelatedObjectID
	s.ResourceStatus = o.ResourceStatus
	s.RouteTableID = o.RouteTableID
	s.SecurityTags = o.SecurityTags
	s.Subnets = o.Subnets

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudInterfaceData) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudInterfaceData{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Addresses = s.Addresses
	o.AttachedEntities = s.AttachedEntities
	o.AttachmentType = s.AttachmentType
	o.AvailabilityZone = s.AvailabilityZone
	o.GroupTags = s.GroupTags
	o.HasPublicIP = s.HasPublicIP
	o.ParentID = s.ParentID
	o.RelatedObjectID = s.RelatedObjectID
	o.ResourceStatus = s.ResourceStatus
	o.RouteTableID = s.RouteTableID
	o.SecurityTags = s.SecurityTags
	o.Subnets = s.Subnets

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudInterfaceData) BleveType() string {

	return "cloudinterfacedata"
}

// DeepCopy returns a deep copy if the CloudInterfaceData.
func (o *CloudInterfaceData) DeepCopy() *CloudInterfaceData {

	if o == nil {
		return nil
	}

	out := &CloudInterfaceData{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudInterfaceData.
func (o *CloudInterfaceData) DeepCopyInto(out *CloudInterfaceData) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudInterfaceData: %s", err))
	}

	*out = *target.(*CloudInterfaceData)
}

// Validate valides the current information stored into the structure.
func (o *CloudInterfaceData) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.Addresses {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateRequiredString("attachmentType", string(o.AttachmentType)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("attachmentType", string(o.AttachmentType), []string{"Instance", "LoadBalancer", "Gateway", "Service", "TransitGatewayVPCAttachment", "NetworkLoadBalancer", "Lambda", "GatewayLoadBalancer", "GatewayLoadBalancerEndpoint", "VPCEndpoint", "APIGatewayManaged", "EFA", "UnsupportedService"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("resourceStatus", string(o.ResourceStatus), []string{"Active", "Inactive"}, false); err != nil {
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
func (*CloudInterfaceData) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudInterfaceDataAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudInterfaceDataLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudInterfaceData) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudInterfaceDataAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudInterfaceData) ValueForAttribute(name string) interface{} {

	switch name {
	case "addresses":
		return o.Addresses
	case "attachedEntities":
		return o.AttachedEntities
	case "attachmentType":
		return o.AttachmentType
	case "availabilityZone":
		return o.AvailabilityZone
	case "groupTags":
		return o.GroupTags
	case "hasPublicIP":
		return o.HasPublicIP
	case "parentID":
		return o.ParentID
	case "relatedObjectID":
		return o.RelatedObjectID
	case "resourceStatus":
		return o.ResourceStatus
	case "routeTableID":
		return o.RouteTableID
	case "securityTags":
		return o.SecurityTags
	case "subnets":
		return o.Subnets
	}

	return nil
}

// CloudInterfaceDataAttributesMap represents the map of attribute for CloudInterfaceData.
var CloudInterfaceDataAttributesMap = map[string]elemental.AttributeSpecification{
	"Addresses": {
		AllowedChoices: []string{},
		BSONFieldName:  "addresses",
		ConvertedName:  "Addresses",
		Description: `List of IP addresses/subnets (IPv4 or IPv6) associated with the
interface.`,
		Exposed: true,
		Name:    "addresses",
		Stored:  true,
		SubType: "cloudaddress",
		Type:    "refList",
	},
	"AttachedEntities": {
		AllowedChoices: []string{},
		BSONFieldName:  "attachedentities",
		ConvertedName:  "AttachedEntities",
		Description:    `ID of associated objects with this interface.`,
		Exposed:        true,
		Name:           "attachedEntities",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"AttachmentType": {
		AllowedChoices: []string{"Instance", "LoadBalancer", "Gateway", "Service", "TransitGatewayVPCAttachment", "NetworkLoadBalancer", "Lambda", "GatewayLoadBalancer", "GatewayLoadBalancerEndpoint", "VPCEndpoint", "APIGatewayManaged", "EFA", "UnsupportedService"},
		BSONFieldName:  "attachmenttype",
		ConvertedName:  "AttachmentType",
		Description: `Attachment type describes where this interface is attached to (Instance, Load
Balancer, Gateway, etc).`,
		Exposed:  true,
		Name:     "attachmentType",
		Required: true,
		Stored:   true,
		Type:     "enum",
	},
	"AvailabilityZone": {
		AllowedChoices: []string{},
		BSONFieldName:  "availabilityzone",
		ConvertedName:  "AvailabilityZone",
		Description:    `Availability zone of the interface.`,
		Exposed:        true,
		Name:           "availabilityZone",
		Stored:         true,
		Type:           "string",
	},
	"GroupTags": {
		AllowedChoices: []string{},
		BSONFieldName:  "grouptags",
		ConvertedName:  "GroupTags",
		Description: `Group tags associated with the interface. In Azure, its Application Security
Group.`,
		Exposed: true,
		Name:    "groupTags",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"HasPublicIP": {
		AllowedChoices: []string{},
		BSONFieldName:  "haspublicip",
		ConvertedName:  "HasPublicIP",
		Description:    `If the interface has a public IP in one of its IP address.`,
		Exposed:        true,
		Name:           "hasPublicIP",
		Stored:         true,
		Type:           "boolean",
	},
	"ParentID": {
		AllowedChoices: []string{},
		BSONFieldName:  "parentid",
		ConvertedName:  "ParentID",
		Description: `If the interface is attached to Load Balancer, the parentID identifies the
related Load Balancer.`,
		Exposed: true,
		Name:    "parentID",
		Stored:  true,
		Type:    "string",
	},
	"RelatedObjectID": {
		AllowedChoices: []string{},
		BSONFieldName:  "relatedobjectid",
		ConvertedName:  "RelatedObjectID",
		Description: `If the interface is of type or external, the relatedObjectID identifies the
related service or gateway.`,
		Exposed: true,
		Name:    "relatedObjectID",
		Stored:  true,
		Type:    "string",
	},
	"ResourceStatus": {
		AllowedChoices: []string{"Active", "Inactive"},
		BSONFieldName:  "resourcestatus",
		ConvertedName:  "ResourceStatus",
		DefaultValue:   CloudInterfaceDataResourceStatusActive,
		Description:    `The status of the resource.`,
		Exposed:        true,
		Name:           "resourceStatus",
		Stored:         true,
		Type:           "enum",
	},
	"RouteTableID": {
		AllowedChoices: []string{},
		BSONFieldName:  "routetableid",
		ConvertedName:  "RouteTableID",
		Description: `The route table that must be used for this interface. Applies to Transit
Gateways and other special types.`,
		Exposed: true,
		Name:    "routeTableID",
		Stored:  true,
		Type:    "string",
	},
	"SecurityTags": {
		AllowedChoices: []string{},
		BSONFieldName:  "securitytags",
		ConvertedName:  "SecurityTags",
		Description:    `Security tags associated with the instance.`,
		Exposed:        true,
		Name:           "securityTags",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Subnets": {
		AllowedChoices: []string{},
		BSONFieldName:  "subnets",
		ConvertedName:  "Subnets",
		Description:    `ID of subnet associated with this interface.`,
		Exposed:        true,
		Name:           "subnets",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
}

// CloudInterfaceDataLowerCaseAttributesMap represents the map of attribute for CloudInterfaceData.
var CloudInterfaceDataLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"addresses": {
		AllowedChoices: []string{},
		BSONFieldName:  "addresses",
		ConvertedName:  "Addresses",
		Description: `List of IP addresses/subnets (IPv4 or IPv6) associated with the
interface.`,
		Exposed: true,
		Name:    "addresses",
		Stored:  true,
		SubType: "cloudaddress",
		Type:    "refList",
	},
	"attachedentities": {
		AllowedChoices: []string{},
		BSONFieldName:  "attachedentities",
		ConvertedName:  "AttachedEntities",
		Description:    `ID of associated objects with this interface.`,
		Exposed:        true,
		Name:           "attachedEntities",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"attachmenttype": {
		AllowedChoices: []string{"Instance", "LoadBalancer", "Gateway", "Service", "TransitGatewayVPCAttachment", "NetworkLoadBalancer", "Lambda", "GatewayLoadBalancer", "GatewayLoadBalancerEndpoint", "VPCEndpoint", "APIGatewayManaged", "EFA", "UnsupportedService"},
		BSONFieldName:  "attachmenttype",
		ConvertedName:  "AttachmentType",
		Description: `Attachment type describes where this interface is attached to (Instance, Load
Balancer, Gateway, etc).`,
		Exposed:  true,
		Name:     "attachmentType",
		Required: true,
		Stored:   true,
		Type:     "enum",
	},
	"availabilityzone": {
		AllowedChoices: []string{},
		BSONFieldName:  "availabilityzone",
		ConvertedName:  "AvailabilityZone",
		Description:    `Availability zone of the interface.`,
		Exposed:        true,
		Name:           "availabilityZone",
		Stored:         true,
		Type:           "string",
	},
	"grouptags": {
		AllowedChoices: []string{},
		BSONFieldName:  "grouptags",
		ConvertedName:  "GroupTags",
		Description: `Group tags associated with the interface. In Azure, its Application Security
Group.`,
		Exposed: true,
		Name:    "groupTags",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"haspublicip": {
		AllowedChoices: []string{},
		BSONFieldName:  "haspublicip",
		ConvertedName:  "HasPublicIP",
		Description:    `If the interface has a public IP in one of its IP address.`,
		Exposed:        true,
		Name:           "hasPublicIP",
		Stored:         true,
		Type:           "boolean",
	},
	"parentid": {
		AllowedChoices: []string{},
		BSONFieldName:  "parentid",
		ConvertedName:  "ParentID",
		Description: `If the interface is attached to Load Balancer, the parentID identifies the
related Load Balancer.`,
		Exposed: true,
		Name:    "parentID",
		Stored:  true,
		Type:    "string",
	},
	"relatedobjectid": {
		AllowedChoices: []string{},
		BSONFieldName:  "relatedobjectid",
		ConvertedName:  "RelatedObjectID",
		Description: `If the interface is of type or external, the relatedObjectID identifies the
related service or gateway.`,
		Exposed: true,
		Name:    "relatedObjectID",
		Stored:  true,
		Type:    "string",
	},
	"resourcestatus": {
		AllowedChoices: []string{"Active", "Inactive"},
		BSONFieldName:  "resourcestatus",
		ConvertedName:  "ResourceStatus",
		DefaultValue:   CloudInterfaceDataResourceStatusActive,
		Description:    `The status of the resource.`,
		Exposed:        true,
		Name:           "resourceStatus",
		Stored:         true,
		Type:           "enum",
	},
	"routetableid": {
		AllowedChoices: []string{},
		BSONFieldName:  "routetableid",
		ConvertedName:  "RouteTableID",
		Description: `The route table that must be used for this interface. Applies to Transit
Gateways and other special types.`,
		Exposed: true,
		Name:    "routeTableID",
		Stored:  true,
		Type:    "string",
	},
	"securitytags": {
		AllowedChoices: []string{},
		BSONFieldName:  "securitytags",
		ConvertedName:  "SecurityTags",
		Description:    `Security tags associated with the instance.`,
		Exposed:        true,
		Name:           "securityTags",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"subnets": {
		AllowedChoices: []string{},
		BSONFieldName:  "subnets",
		ConvertedName:  "Subnets",
		Description:    `ID of subnet associated with this interface.`,
		Exposed:        true,
		Name:           "subnets",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
}

type mongoAttributesCloudInterfaceData struct {
	Addresses        []*CloudAddress                       `bson:"addresses"`
	AttachedEntities []string                              `bson:"attachedentities"`
	AttachmentType   CloudInterfaceDataAttachmentTypeValue `bson:"attachmenttype"`
	AvailabilityZone string                                `bson:"availabilityzone"`
	GroupTags        []string                              `bson:"grouptags"`
	HasPublicIP      bool                                  `bson:"haspublicip"`
	ParentID         string                                `bson:"parentid"`
	RelatedObjectID  string                                `bson:"relatedobjectid"`
	ResourceStatus   CloudInterfaceDataResourceStatusValue `bson:"resourcestatus"`
	RouteTableID     string                                `bson:"routetableid"`
	SecurityTags     []string                              `bson:"securitytags"`
	Subnets          []string                              `bson:"subnets"`
}
