// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudEndpointDataResourceStatusValue represents the possible values for attribute "resourceStatus".
type CloudEndpointDataResourceStatusValue string

const (
	// CloudEndpointDataResourceStatusActive represents the value Active.
	CloudEndpointDataResourceStatusActive CloudEndpointDataResourceStatusValue = "Active"

	// CloudEndpointDataResourceStatusInactive represents the value Inactive.
	CloudEndpointDataResourceStatusInactive CloudEndpointDataResourceStatusValue = "Inactive"
)

// CloudEndpointDataServiceTypeValue represents the possible values for attribute "serviceType".
type CloudEndpointDataServiceTypeValue string

const (
	// CloudEndpointDataServiceTypeGateway represents the value Gateway.
	CloudEndpointDataServiceTypeGateway CloudEndpointDataServiceTypeValue = "Gateway"

	// CloudEndpointDataServiceTypeGatewayLoadBalancer represents the value GatewayLoadBalancer.
	CloudEndpointDataServiceTypeGatewayLoadBalancer CloudEndpointDataServiceTypeValue = "GatewayLoadBalancer"

	// CloudEndpointDataServiceTypeInterface represents the value Interface.
	CloudEndpointDataServiceTypeInterface CloudEndpointDataServiceTypeValue = "Interface"

	// CloudEndpointDataServiceTypeManagedService represents the value ManagedService.
	CloudEndpointDataServiceTypeManagedService CloudEndpointDataServiceTypeValue = "ManagedService"

	// CloudEndpointDataServiceTypeNotApplicable represents the value NotApplicable.
	CloudEndpointDataServiceTypeNotApplicable CloudEndpointDataServiceTypeValue = "NotApplicable"
)

// CloudEndpointDataTypeValue represents the possible values for attribute "type".
type CloudEndpointDataTypeValue string

const (
	// CloudEndpointDataTypeGateway represents the value Gateway.
	CloudEndpointDataTypeGateway CloudEndpointDataTypeValue = "Gateway"

	// CloudEndpointDataTypeInstance represents the value Instance.
	CloudEndpointDataTypeInstance CloudEndpointDataTypeValue = "Instance"

	// CloudEndpointDataTypeLoadBalancer represents the value LoadBalancer.
	CloudEndpointDataTypeLoadBalancer CloudEndpointDataTypeValue = "LoadBalancer"

	// CloudEndpointDataTypeNATGateway represents the value NATGateway.
	CloudEndpointDataTypeNATGateway CloudEndpointDataTypeValue = "NATGateway"

	// CloudEndpointDataTypePeeringConnection represents the value PeeringConnection.
	CloudEndpointDataTypePeeringConnection CloudEndpointDataTypeValue = "PeeringConnection"

	// CloudEndpointDataTypePublicIPAddress represents the value PublicIPAddress.
	CloudEndpointDataTypePublicIPAddress CloudEndpointDataTypeValue = "PublicIPAddress"

	// CloudEndpointDataTypeService represents the value Service.
	CloudEndpointDataTypeService CloudEndpointDataTypeValue = "Service"

	// CloudEndpointDataTypeTransitGateway represents the value TransitGateway.
	CloudEndpointDataTypeTransitGateway CloudEndpointDataTypeValue = "TransitGateway"
)

// CloudEndpointData represents the model of a cloudendpointdata
type CloudEndpointData struct {
	// Indicates that the endpoint is directly attached to the VPC. In this case the
	// attachedInterfaces is empty. In general this is only valid for endpoint type
	// Gateway and Peering Connection.
	VPCAttached bool `json:"VPCAttached" msgpack:"VPCAttached" bson:"vpcattached" mapstructure:"VPCAttached,omitempty"`

	// The list of VPCs that this endpoint is directly attached to.
	VPCAttachments []string `json:"VPCAttachments" msgpack:"VPCAttachments" bson:"vpcattachments" mapstructure:"VPCAttachments,omitempty"`

	// List of route tables associated with this endpoint. Depending on cloud provider
	// it can apply in some gateways.
	AssociatedRouteTables []string `json:"associatedRouteTables" msgpack:"associatedRouteTables" bson:"associatedroutetables" mapstructure:"associatedRouteTables,omitempty"`

	// A list of entities that are associated to a given endpoint.
	AttachedEntities []string `json:"attachedEntities" msgpack:"attachedEntities" bson:"attachedentities" mapstructure:"attachedEntities,omitempty"`

	// A list of interfaces attached with the endpoint. In some cases endpoints can
	// have more than one interface.
	AttachedInterfaces []string `json:"attachedInterfaces" msgpack:"attachedInterfaces" bson:"attachedinterfaces" mapstructure:"attachedInterfaces,omitempty"`

	// If the endpoint has multiple connections and forwarding can be enabled between
	// them.
	ForwardingEnabled bool `json:"forwardingEnabled" msgpack:"forwardingEnabled" bson:"forwardingenabled" mapstructure:"forwardingEnabled,omitempty"`

	// If the endpoint has a public IP or is directly exposed.
	HasPublicIP bool `json:"hasPublicIP,omitempty" msgpack:"hasPublicIP,omitempty" bson:"haspublicip,omitempty" mapstructure:"hasPublicIP,omitempty"`

	// The imageID of running in the endpoint. Available for instances and
	// potentially other 3rd parties. This can be the AMI ID in AWS or corresponding
	// instance imageID in other clouds.
	ImageID string `json:"imageID,omitempty" msgpack:"imageID,omitempty" bson:"imageid,omitempty" mapstructure:"imageID,omitempty"`

	// Product related metadata associated with this endpoint.
	ProductInfo []*CloudEndpointDataProductInfo `json:"productInfo,omitempty" msgpack:"productInfo,omitempty" bson:"productinfo,omitempty" mapstructure:"productInfo,omitempty"`

	// if the endpoint has a public IP we store the IP address in this field.
	PublicIPAddresses []string `json:"publicIPAddresses" msgpack:"publicIPAddresses" bson:"publicipaddresses" mapstructure:"publicIPAddresses,omitempty"`

	// The status of the resource.
	ResourceStatus CloudEndpointDataResourceStatusValue `json:"resourceStatus" msgpack:"resourceStatus" bson:"resourcestatus" mapstructure:"resourceStatus,omitempty"`

	// Identifies the name of the service for service endpoints.
	ServiceName string `json:"serviceName,omitempty" msgpack:"serviceName,omitempty" bson:"servicename,omitempty" mapstructure:"serviceName,omitempty"`

	// Identifies the service type that this endpoint represents (example Gateway Load
	// Balancer).
	ServiceType CloudEndpointDataServiceTypeValue `json:"serviceType" msgpack:"serviceType" bson:"servicetype" mapstructure:"serviceType,omitempty"`

	// Type of the endpoint.
	Type CloudEndpointDataTypeValue `json:"type" msgpack:"type" bson:"type" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudEndpointData returns a new *CloudEndpointData
func NewCloudEndpointData() *CloudEndpointData {

	return &CloudEndpointData{
		ModelVersion:          1,
		AssociatedRouteTables: []string{},
		AttachedEntities:      []string{},
		AttachedInterfaces:    []string{},
		ResourceStatus:        CloudEndpointDataResourceStatusActive,
		ProductInfo:           []*CloudEndpointDataProductInfo{},
		PublicIPAddresses:     []string{},
		ServiceType:           CloudEndpointDataServiceTypeNotApplicable,
		VPCAttachments:        []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudEndpointData) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudEndpointData{}

	s.VPCAttached = o.VPCAttached
	s.VPCAttachments = o.VPCAttachments
	s.AssociatedRouteTables = o.AssociatedRouteTables
	s.AttachedEntities = o.AttachedEntities
	s.AttachedInterfaces = o.AttachedInterfaces
	s.ForwardingEnabled = o.ForwardingEnabled
	s.HasPublicIP = o.HasPublicIP
	s.ImageID = o.ImageID
	s.ProductInfo = o.ProductInfo
	s.PublicIPAddresses = o.PublicIPAddresses
	s.ResourceStatus = o.ResourceStatus
	s.ServiceName = o.ServiceName
	s.ServiceType = o.ServiceType
	s.Type = o.Type

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudEndpointData) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudEndpointData{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.VPCAttached = s.VPCAttached
	o.VPCAttachments = s.VPCAttachments
	o.AssociatedRouteTables = s.AssociatedRouteTables
	o.AttachedEntities = s.AttachedEntities
	o.AttachedInterfaces = s.AttachedInterfaces
	o.ForwardingEnabled = s.ForwardingEnabled
	o.HasPublicIP = s.HasPublicIP
	o.ImageID = s.ImageID
	o.ProductInfo = s.ProductInfo
	o.PublicIPAddresses = s.PublicIPAddresses
	o.ResourceStatus = s.ResourceStatus
	o.ServiceName = s.ServiceName
	o.ServiceType = s.ServiceType
	o.Type = s.Type

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudEndpointData) BleveType() string {

	return "cloudendpointdata"
}

// DeepCopy returns a deep copy if the CloudEndpointData.
func (o *CloudEndpointData) DeepCopy() *CloudEndpointData {

	if o == nil {
		return nil
	}

	out := &CloudEndpointData{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudEndpointData.
func (o *CloudEndpointData) DeepCopyInto(out *CloudEndpointData) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudEndpointData: %s", err))
	}

	*out = *target.(*CloudEndpointData)
}

// Validate valides the current information stored into the structure.
func (o *CloudEndpointData) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.ProductInfo {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateStringInList("resourceStatus", string(o.ResourceStatus), []string{"Active", "Inactive"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("serviceType", string(o.ServiceType), []string{"Interface", "Gateway", "GatewayLoadBalancer", "ManagedService", "NotApplicable"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("type", string(o.Type)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Instance", "LoadBalancer", "PeeringConnection", "Service", "Gateway", "TransitGateway", "NATGateway", "PublicIPAddress"}, false); err != nil {
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
func (*CloudEndpointData) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudEndpointDataAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudEndpointDataLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudEndpointData) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudEndpointDataAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudEndpointData) ValueForAttribute(name string) interface{} {

	switch name {
	case "VPCAttached":
		return o.VPCAttached
	case "VPCAttachments":
		return o.VPCAttachments
	case "associatedRouteTables":
		return o.AssociatedRouteTables
	case "attachedEntities":
		return o.AttachedEntities
	case "attachedInterfaces":
		return o.AttachedInterfaces
	case "forwardingEnabled":
		return o.ForwardingEnabled
	case "hasPublicIP":
		return o.HasPublicIP
	case "imageID":
		return o.ImageID
	case "productInfo":
		return o.ProductInfo
	case "publicIPAddresses":
		return o.PublicIPAddresses
	case "resourceStatus":
		return o.ResourceStatus
	case "serviceName":
		return o.ServiceName
	case "serviceType":
		return o.ServiceType
	case "type":
		return o.Type
	}

	return nil
}

// CloudEndpointDataAttributesMap represents the map of attribute for CloudEndpointData.
var CloudEndpointDataAttributesMap = map[string]elemental.AttributeSpecification{
	"VPCAttached": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcattached",
		ConvertedName:  "VPCAttached",
		Description: `Indicates that the endpoint is directly attached to the VPC. In this case the
attachedInterfaces is empty. In general this is only valid for endpoint type
Gateway and Peering Connection.`,
		Exposed: true,
		Name:    "VPCAttached",
		Stored:  true,
		Type:    "boolean",
	},
	"VPCAttachments": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcattachments",
		ConvertedName:  "VPCAttachments",
		Description:    `The list of VPCs that this endpoint is directly attached to.`,
		Exposed:        true,
		Name:           "VPCAttachments",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"AssociatedRouteTables": {
		AllowedChoices: []string{},
		BSONFieldName:  "associatedroutetables",
		ConvertedName:  "AssociatedRouteTables",
		Description: `List of route tables associated with this endpoint. Depending on cloud provider
it can apply in some gateways.`,
		Exposed: true,
		Name:    "associatedRouteTables",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"AttachedEntities": {
		AllowedChoices: []string{},
		BSONFieldName:  "attachedentities",
		ConvertedName:  "AttachedEntities",
		Description:    `A list of entities that are associated to a given endpoint.`,
		Exposed:        true,
		Name:           "attachedEntities",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"AttachedInterfaces": {
		AllowedChoices: []string{},
		BSONFieldName:  "attachedinterfaces",
		ConvertedName:  "AttachedInterfaces",
		Description: `A list of interfaces attached with the endpoint. In some cases endpoints can
have more than one interface.`,
		Exposed: true,
		Name:    "attachedInterfaces",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"ForwardingEnabled": {
		AllowedChoices: []string{},
		BSONFieldName:  "forwardingenabled",
		ConvertedName:  "ForwardingEnabled",
		Description: `If the endpoint has multiple connections and forwarding can be enabled between
them.`,
		Exposed: true,
		Name:    "forwardingEnabled",
		Stored:  true,
		Type:    "boolean",
	},
	"HasPublicIP": {
		AllowedChoices: []string{},
		BSONFieldName:  "haspublicip",
		ConvertedName:  "HasPublicIP",
		Description:    `If the endpoint has a public IP or is directly exposed.`,
		Exposed:        true,
		Name:           "hasPublicIP",
		Stored:         true,
		Type:           "boolean",
	},
	"ImageID": {
		AllowedChoices: []string{},
		BSONFieldName:  "imageid",
		ConvertedName:  "ImageID",
		Description: `The imageID of running in the endpoint. Available for instances and
potentially other 3rd parties. This can be the AMI ID in AWS or corresponding
instance imageID in other clouds.`,
		Exposed: true,
		Name:    "imageID",
		Stored:  true,
		Type:    "string",
	},
	"ProductInfo": {
		AllowedChoices: []string{},
		BSONFieldName:  "productinfo",
		ConvertedName:  "ProductInfo",
		Description:    `Product related metadata associated with this endpoint.`,
		Exposed:        true,
		Name:           "productInfo",
		Stored:         true,
		SubType:        "cloudendpointdataproductinfo",
		Type:           "refList",
	},
	"PublicIPAddresses": {
		AllowedChoices: []string{},
		BSONFieldName:  "publicipaddresses",
		ConvertedName:  "PublicIPAddresses",
		Description:    `if the endpoint has a public IP we store the IP address in this field.`,
		Exposed:        true,
		Name:           "publicIPAddresses",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"ResourceStatus": {
		AllowedChoices: []string{"Active", "Inactive"},
		BSONFieldName:  "resourcestatus",
		ConvertedName:  "ResourceStatus",
		DefaultValue:   CloudEndpointDataResourceStatusActive,
		Description:    `The status of the resource.`,
		Exposed:        true,
		Name:           "resourceStatus",
		Stored:         true,
		Type:           "enum",
	},
	"ServiceName": {
		AllowedChoices: []string{},
		BSONFieldName:  "servicename",
		ConvertedName:  "ServiceName",
		Description:    `Identifies the name of the service for service endpoints.`,
		Exposed:        true,
		Name:           "serviceName",
		Stored:         true,
		Type:           "string",
	},
	"ServiceType": {
		AllowedChoices: []string{"Interface", "Gateway", "GatewayLoadBalancer", "ManagedService", "NotApplicable"},
		BSONFieldName:  "servicetype",
		ConvertedName:  "ServiceType",
		DefaultValue:   CloudEndpointDataServiceTypeNotApplicable,
		Description: `Identifies the service type that this endpoint represents (example Gateway Load
Balancer).`,
		Exposed: true,
		Name:    "serviceType",
		Stored:  true,
		Type:    "enum",
	},
	"Type": {
		AllowedChoices: []string{"Instance", "LoadBalancer", "PeeringConnection", "Service", "Gateway", "TransitGateway", "NATGateway", "PublicIPAddress"},
		BSONFieldName:  "type",
		ConvertedName:  "Type",
		Description:    `Type of the endpoint.`,
		Exposed:        true,
		Name:           "type",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
}

// CloudEndpointDataLowerCaseAttributesMap represents the map of attribute for CloudEndpointData.
var CloudEndpointDataLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"vpcattached": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcattached",
		ConvertedName:  "VPCAttached",
		Description: `Indicates that the endpoint is directly attached to the VPC. In this case the
attachedInterfaces is empty. In general this is only valid for endpoint type
Gateway and Peering Connection.`,
		Exposed: true,
		Name:    "VPCAttached",
		Stored:  true,
		Type:    "boolean",
	},
	"vpcattachments": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcattachments",
		ConvertedName:  "VPCAttachments",
		Description:    `The list of VPCs that this endpoint is directly attached to.`,
		Exposed:        true,
		Name:           "VPCAttachments",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"associatedroutetables": {
		AllowedChoices: []string{},
		BSONFieldName:  "associatedroutetables",
		ConvertedName:  "AssociatedRouteTables",
		Description: `List of route tables associated with this endpoint. Depending on cloud provider
it can apply in some gateways.`,
		Exposed: true,
		Name:    "associatedRouteTables",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"attachedentities": {
		AllowedChoices: []string{},
		BSONFieldName:  "attachedentities",
		ConvertedName:  "AttachedEntities",
		Description:    `A list of entities that are associated to a given endpoint.`,
		Exposed:        true,
		Name:           "attachedEntities",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"attachedinterfaces": {
		AllowedChoices: []string{},
		BSONFieldName:  "attachedinterfaces",
		ConvertedName:  "AttachedInterfaces",
		Description: `A list of interfaces attached with the endpoint. In some cases endpoints can
have more than one interface.`,
		Exposed: true,
		Name:    "attachedInterfaces",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"forwardingenabled": {
		AllowedChoices: []string{},
		BSONFieldName:  "forwardingenabled",
		ConvertedName:  "ForwardingEnabled",
		Description: `If the endpoint has multiple connections and forwarding can be enabled between
them.`,
		Exposed: true,
		Name:    "forwardingEnabled",
		Stored:  true,
		Type:    "boolean",
	},
	"haspublicip": {
		AllowedChoices: []string{},
		BSONFieldName:  "haspublicip",
		ConvertedName:  "HasPublicIP",
		Description:    `If the endpoint has a public IP or is directly exposed.`,
		Exposed:        true,
		Name:           "hasPublicIP",
		Stored:         true,
		Type:           "boolean",
	},
	"imageid": {
		AllowedChoices: []string{},
		BSONFieldName:  "imageid",
		ConvertedName:  "ImageID",
		Description: `The imageID of running in the endpoint. Available for instances and
potentially other 3rd parties. This can be the AMI ID in AWS or corresponding
instance imageID in other clouds.`,
		Exposed: true,
		Name:    "imageID",
		Stored:  true,
		Type:    "string",
	},
	"productinfo": {
		AllowedChoices: []string{},
		BSONFieldName:  "productinfo",
		ConvertedName:  "ProductInfo",
		Description:    `Product related metadata associated with this endpoint.`,
		Exposed:        true,
		Name:           "productInfo",
		Stored:         true,
		SubType:        "cloudendpointdataproductinfo",
		Type:           "refList",
	},
	"publicipaddresses": {
		AllowedChoices: []string{},
		BSONFieldName:  "publicipaddresses",
		ConvertedName:  "PublicIPAddresses",
		Description:    `if the endpoint has a public IP we store the IP address in this field.`,
		Exposed:        true,
		Name:           "publicIPAddresses",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"resourcestatus": {
		AllowedChoices: []string{"Active", "Inactive"},
		BSONFieldName:  "resourcestatus",
		ConvertedName:  "ResourceStatus",
		DefaultValue:   CloudEndpointDataResourceStatusActive,
		Description:    `The status of the resource.`,
		Exposed:        true,
		Name:           "resourceStatus",
		Stored:         true,
		Type:           "enum",
	},
	"servicename": {
		AllowedChoices: []string{},
		BSONFieldName:  "servicename",
		ConvertedName:  "ServiceName",
		Description:    `Identifies the name of the service for service endpoints.`,
		Exposed:        true,
		Name:           "serviceName",
		Stored:         true,
		Type:           "string",
	},
	"servicetype": {
		AllowedChoices: []string{"Interface", "Gateway", "GatewayLoadBalancer", "ManagedService", "NotApplicable"},
		BSONFieldName:  "servicetype",
		ConvertedName:  "ServiceType",
		DefaultValue:   CloudEndpointDataServiceTypeNotApplicable,
		Description: `Identifies the service type that this endpoint represents (example Gateway Load
Balancer).`,
		Exposed: true,
		Name:    "serviceType",
		Stored:  true,
		Type:    "enum",
	},
	"type": {
		AllowedChoices: []string{"Instance", "LoadBalancer", "PeeringConnection", "Service", "Gateway", "TransitGateway", "NATGateway", "PublicIPAddress"},
		BSONFieldName:  "type",
		ConvertedName:  "Type",
		Description:    `Type of the endpoint.`,
		Exposed:        true,
		Name:           "type",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
}

type mongoAttributesCloudEndpointData struct {
	VPCAttached           bool                                 `bson:"vpcattached"`
	VPCAttachments        []string                             `bson:"vpcattachments"`
	AssociatedRouteTables []string                             `bson:"associatedroutetables"`
	AttachedEntities      []string                             `bson:"attachedentities"`
	AttachedInterfaces    []string                             `bson:"attachedinterfaces"`
	ForwardingEnabled     bool                                 `bson:"forwardingenabled"`
	HasPublicIP           bool                                 `bson:"haspublicip,omitempty"`
	ImageID               string                               `bson:"imageid,omitempty"`
	ProductInfo           []*CloudEndpointDataProductInfo      `bson:"productinfo,omitempty"`
	PublicIPAddresses     []string                             `bson:"publicipaddresses"`
	ResourceStatus        CloudEndpointDataResourceStatusValue `bson:"resourcestatus"`
	ServiceName           string                               `bson:"servicename,omitempty"`
	ServiceType           CloudEndpointDataServiceTypeValue    `bson:"servicetype"`
	Type                  CloudEndpointDataTypeValue           `bson:"type"`
}
