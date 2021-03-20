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

	// CloudInterfaceDataAttachmentTypeVPCEndpoint represents the value VPCEndpoint.
	CloudInterfaceDataAttachmentTypeVPCEndpoint CloudInterfaceDataAttachmentTypeValue = "VPCEndpoint"
)

// CloudInterfaceData represents the model of a cloudinterfacedata
type CloudInterfaceData struct {
	// List of IP addresses/subnets (IPv4 or IPv6) associated with the
	// interface.
	Addresses CloudAddressList `json:"addresses" msgpack:"addresses" bson:"addresses" mapstructure:"addresses,omitempty"`

	// Attachment type describes where this interface is attached to (Instance, Load
	// Balancer, Gateway, etc).
	AttachmentType CloudInterfaceDataAttachmentTypeValue `json:"attachmentType" msgpack:"attachmentType" bson:"attachmenttype" mapstructure:"attachmentType,omitempty"`

	// If the interface is of type or external, the relatedObjectID identifies the
	// related service or gateway.
	RelatedObjectID string `json:"relatedObjectID" msgpack:"relatedObjectID" bson:"relatedobjectid" mapstructure:"relatedObjectID,omitempty"`

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
		ModelVersion: 1,
		Addresses:    CloudAddressList{},
		SecurityTags: []string{},
		Subnets:      []string{},
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
	s.AttachmentType = o.AttachmentType
	s.RelatedObjectID = o.RelatedObjectID
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
	o.AttachmentType = s.AttachmentType
	o.RelatedObjectID = s.RelatedObjectID
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

	if err := elemental.ValidateStringInList("attachmentType", string(o.AttachmentType), []string{"Instance", "LoadBalancer", "Gateway", "Service", "TransitGatewayVPCAttachment", "NetworkLoadBalancer", "Lambda", "GatewayLoadBalancer", "GatewayLoadBalancerEndpoint", "VPCEndpoint", "APIGatewayManaged", "EFA"}, false); err != nil {
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

type mongoAttributesCloudInterfaceData struct {
	Addresses       CloudAddressList                      `bson:"addresses"`
	AttachmentType  CloudInterfaceDataAttachmentTypeValue `bson:"attachmenttype"`
	RelatedObjectID string                                `bson:"relatedobjectid"`
	RouteTableID    string                                `bson:"routetableid"`
	SecurityTags    []string                              `bson:"securitytags"`
	Subnets         []string                              `bson:"subnets"`
}
