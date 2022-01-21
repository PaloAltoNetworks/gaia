package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
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

	// A list of interfaces attached with the endpoint. In some cases endpoints can
	// have more than one interface.
	AttachedInterfaces []string `json:"attachedInterfaces" msgpack:"attachedInterfaces" bson:"attachedinterfaces" mapstructure:"attachedInterfaces,omitempty"`

	// If the endpoint has multiple connections and forwarding can be enabled between
	// them.
	ForwardingEnabled bool `json:"forwardingEnabled" msgpack:"forwardingEnabled" bson:"forwardingenabled" mapstructure:"forwardingEnabled,omitempty"`

	// The imageID of running in the endpoint. Available for instances and
	// potentially other 3rd parties. This can be the AMI ID in AWS or corresponding
	// instance imageID in other clouds.
	ImageID string `json:"imageID,omitempty" msgpack:"imageID,omitempty" bson:"imageid" mapstructure:"imageID,omitempty"`

	// Product related metadata associated with this endpoint.
	ProductInfo []*CloudEndpointDataProductInfo `json:"productInfo,omitempty" msgpack:"productInfo,omitempty" bson:"productinfo" mapstructure:"productInfo,omitempty"`

	// if the endpoint has a public IP we store the IP address in this field.
	PublicIPAddresses []string `json:"publicIPAddresses" msgpack:"publicIPAddresses" bson:"publicipaddresses" mapstructure:"publicIPAddresses,omitempty"`

	// Identifies the name of the service for service endpoints.
	ServiceName string `json:"serviceName,omitempty" msgpack:"serviceName,omitempty" bson:"servicename" mapstructure:"serviceName,omitempty"`

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
		ProductInfo:           []*CloudEndpointDataProductInfo{},
		AssociatedRouteTables: []string{},
		PublicIPAddresses:     []string{},
		AttachedInterfaces:    []string{},
		VPCAttachments:        []string{},
		ServiceType:           CloudEndpointDataServiceTypeNotApplicable,
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
	s.AttachedInterfaces = o.AttachedInterfaces
	s.ForwardingEnabled = o.ForwardingEnabled
	s.ImageID = o.ImageID
	s.ProductInfo = o.ProductInfo
	s.PublicIPAddresses = o.PublicIPAddresses
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
	o.AttachedInterfaces = s.AttachedInterfaces
	o.ForwardingEnabled = s.ForwardingEnabled
	o.ImageID = s.ImageID
	o.ProductInfo = s.ProductInfo
	o.PublicIPAddresses = s.PublicIPAddresses
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

	if err := elemental.ValidateStringInList("serviceType", string(o.ServiceType), []string{"Interface", "Gateway", "GatewayLoadBalancer", "NotApplicable"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("type", string(o.Type)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Instance", "LoadBalancer", "PeeringConnection", "Service", "Gateway", "TransitGateway", "NATGateway"}, false); err != nil {
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

type mongoAttributesCloudEndpointData struct {
	VPCAttached           bool                              `bson:"vpcattached"`
	VPCAttachments        []string                          `bson:"vpcattachments"`
	AssociatedRouteTables []string                          `bson:"associatedroutetables"`
	AttachedInterfaces    []string                          `bson:"attachedinterfaces"`
	ForwardingEnabled     bool                              `bson:"forwardingenabled"`
	ImageID               string                            `bson:"imageid,omitempty"`
	ProductInfo           []*CloudEndpointDataProductInfo   `bson:"productinfo,omitempty"`
	PublicIPAddresses     []string                          `bson:"publicipaddresses"`
	ServiceName           string                            `bson:"servicename,omitempty"`
	ServiceType           CloudEndpointDataServiceTypeValue `bson:"servicetype"`
	Type                  CloudEndpointDataTypeValue        `bson:"type"`
}
