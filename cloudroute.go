package gaia

import (
	"fmt"
	"net"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudRouteNextHopTypeValue represents the possible values for attribute "nextHopType".
type CloudRouteNextHopTypeValue string

const (
	// CloudRouteNextHopTypeEgressOnlyGateway represents the value EgressOnlyGateway.
	CloudRouteNextHopTypeEgressOnlyGateway CloudRouteNextHopTypeValue = "EgressOnlyGateway"

	// CloudRouteNextHopTypeGateway represents the value Gateway.
	CloudRouteNextHopTypeGateway CloudRouteNextHopTypeValue = "Gateway"

	// CloudRouteNextHopTypeInstance represents the value Instance.
	CloudRouteNextHopTypeInstance CloudRouteNextHopTypeValue = "Instance"

	// CloudRouteNextHopTypeLocalGateway represents the value LocalGateway.
	CloudRouteNextHopTypeLocalGateway CloudRouteNextHopTypeValue = "LocalGateway"

	// CloudRouteNextHopTypeNATGateway represents the value NATGateway.
	CloudRouteNextHopTypeNATGateway CloudRouteNextHopTypeValue = "NATGateway"

	// CloudRouteNextHopTypeNetworkInterface represents the value NetworkInterface.
	CloudRouteNextHopTypeNetworkInterface CloudRouteNextHopTypeValue = "NetworkInterface"

	// CloudRouteNextHopTypeTransitGateway represents the value TransitGateway.
	CloudRouteNextHopTypeTransitGateway CloudRouteNextHopTypeValue = "TransitGateway"

	// CloudRouteNextHopTypeTransitGatewayAttachment represents the value TransitGatewayAttachment.
	CloudRouteNextHopTypeTransitGatewayAttachment CloudRouteNextHopTypeValue = "TransitGatewayAttachment"

	// CloudRouteNextHopTypeVPCPeeringConnection represents the value VPCPeeringConnection.
	CloudRouteNextHopTypeVPCPeeringConnection CloudRouteNextHopTypeValue = "VPCPeeringConnection"
)

// CloudRoute represents the model of a cloudroute
type CloudRoute struct {
	// The Destination CIDR for the route.
	DestinationIPv4CIDR string `json:"destinationIPv4CIDR" msgpack:"destinationIPv4CIDR" bson:"destinationipv4cidr" mapstructure:"destinationIPv4CIDR,omitempty"`

	// The destination IPV6 CIDR for the route.
	DestinationIPv6CIDR string `json:"destinationIPv6CIDR" msgpack:"destinationIPv6CIDR" bson:"destinationipv6cidr" mapstructure:"destinationIPv6CIDR,omitempty"`

	// The destination is identified as a prefix list ID.
	DestinationPrefixListID string `json:"destinationPrefixListID" msgpack:"destinationPrefixListID" bson:"destinationprefixlistid" mapstructure:"destinationPrefixListID,omitempty"`

	// The ID of the next hop object.
	NextHopID string `json:"nextHopID" msgpack:"nextHopID" bson:"nexthopid" mapstructure:"nextHopID,omitempty"`

	// The type of the next hop.
	NextHopType CloudRouteNextHopTypeValue `json:"nextHopType" msgpack:"nextHopType" bson:"nexthoptype" mapstructure:"nextHopType,omitempty"`

	// Internal representation of IPv4 networks.
	StoredDestinationIPv4CIDR *net.IPNet `json:"storedDestinationIPv4CIDR,omitempty" msgpack:"storedDestinationIPv4CIDR,omitempty" bson:"storeddestinationipv4cidr,omitempty" mapstructure:"storedDestinationIPv4CIDR,omitempty"`

	// Internal representation of IPv6 networks.
	StoredDestinationIPv6CIDR *net.IPNet `json:"storedDestinationIPv6CIDR,omitempty" msgpack:"storedDestinationIPv6CIDR,omitempty" bson:"storeddestinationipv6cidr,omitempty" mapstructure:"storedDestinationIPv6CIDR,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudRoute returns a new *CloudRoute
func NewCloudRoute() *CloudRoute {

	return &CloudRoute{
		ModelVersion:              1,
		StoredDestinationIPv4CIDR: &net.IPNet{},
		StoredDestinationIPv6CIDR: &net.IPNet{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudRoute) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudRoute{}

	s.DestinationIPv4CIDR = o.DestinationIPv4CIDR
	s.DestinationIPv6CIDR = o.DestinationIPv6CIDR
	s.DestinationPrefixListID = o.DestinationPrefixListID
	s.NextHopID = o.NextHopID
	s.NextHopType = o.NextHopType
	s.StoredDestinationIPv4CIDR = o.StoredDestinationIPv4CIDR
	s.StoredDestinationIPv6CIDR = o.StoredDestinationIPv6CIDR

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudRoute) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudRoute{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.DestinationIPv4CIDR = s.DestinationIPv4CIDR
	o.DestinationIPv6CIDR = s.DestinationIPv6CIDR
	o.DestinationPrefixListID = s.DestinationPrefixListID
	o.NextHopID = s.NextHopID
	o.NextHopType = s.NextHopType
	o.StoredDestinationIPv4CIDR = s.StoredDestinationIPv4CIDR
	o.StoredDestinationIPv6CIDR = s.StoredDestinationIPv6CIDR

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudRoute) BleveType() string {

	return "cloudroute"
}

// DeepCopy returns a deep copy if the CloudRoute.
func (o *CloudRoute) DeepCopy() *CloudRoute {

	if o == nil {
		return nil
	}

	out := &CloudRoute{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudRoute.
func (o *CloudRoute) DeepCopyInto(out *CloudRoute) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudRoute: %s", err))
	}

	*out = *target.(*CloudRoute)
}

// Validate valides the current information stored into the structure.
func (o *CloudRoute) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateOptionalCIDRorIP("destinationIPv4CIDR", o.DestinationIPv4CIDR); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateOptionalCIDRorIP("destinationIPv6CIDR", o.DestinationIPv6CIDR); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("nextHopType", string(o.NextHopType)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("nextHopType", string(o.NextHopType), []string{"EgressOnlyGateway", "Gateway", "Instance", "LocalGateway", "NATGateway", "NetworkInterface", "TransitGateway", "VPCPeeringConnection", "TransitGatewayAttachment"}, false); err != nil {
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

type mongoAttributesCloudRoute struct {
	DestinationIPv4CIDR       string                     `bson:"destinationipv4cidr"`
	DestinationIPv6CIDR       string                     `bson:"destinationipv6cidr"`
	DestinationPrefixListID   string                     `bson:"destinationprefixlistid"`
	NextHopID                 string                     `bson:"nexthopid"`
	NextHopType               CloudRouteNextHopTypeValue `bson:"nexthoptype"`
	StoredDestinationIPv4CIDR *net.IPNet                 `bson:"storeddestinationipv4cidr,omitempty"`
	StoredDestinationIPv6CIDR *net.IPNet                 `bson:"storeddestinationipv6cidr,omitempty"`
}
