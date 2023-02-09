// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"
	"net"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudAddressIPVersionValue represents the possible values for attribute "IPVersion".
type CloudAddressIPVersionValue string

const (
	// CloudAddressIPVersionIPv4 represents the value IPv4.
	CloudAddressIPVersionIPv4 CloudAddressIPVersionValue = "IPv4"

	// CloudAddressIPVersionIPv6 represents the value IPv6.
	CloudAddressIPVersionIPv6 CloudAddressIPVersionValue = "IPv6"
)

// CloudAddress represents the model of a cloudaddress
type CloudAddress struct {
	// Designates IPv4 or IPv6.
	IPVersion CloudAddressIPVersionValue `json:"IPVersion" msgpack:"IPVersion" bson:"ipversion" mapstructure:"IPVersion,omitempty"`

	// Designates the IP address as the primary IP address.
	Primary bool `json:"primary" msgpack:"primary" bson:"primary" mapstructure:"primary,omitempty"`

	// The private DNS name associated with the address.
	PrivateDNSName string `json:"privateDNSName" msgpack:"privateDNSName" bson:"privatednsname" mapstructure:"privateDNSName,omitempty"`

	// The private IP address value.
	PrivateIP string `json:"privateIP" msgpack:"privateIP" bson:"privateip" mapstructure:"privateIP,omitempty"`

	// Internal representation of the private IP to accelerate operations. Not exposed
	// to users.
	PrivateNetwork *net.IPNet `json:"privateNetwork,omitempty" msgpack:"privateNetwork,omitempty" bson:"privatenetwork,omitempty" mapstructure:"privateNetwork,omitempty"`

	// The private DNS name associated with the address.
	PublicDNSName string `json:"publicDNSName" msgpack:"publicDNSName" bson:"publicdnsname" mapstructure:"publicDNSName,omitempty"`

	// The public IP address value.
	PublicIP string `json:"publicIP" msgpack:"publicIP" bson:"publicip" mapstructure:"publicIP,omitempty"`

	// The reference to a public IP address.
	PublicIPRef string `json:"publicIPRef" msgpack:"publicIPRef" bson:"publicipref" mapstructure:"publicIPRef,omitempty"`

	// Internal representation of public IP addresses to accelerate operations. Not
	// exposed to users.
	PublicNetwork *net.IPNet `json:"publicNetwork,omitempty" msgpack:"publicNetwork,omitempty" bson:"publicnetwork,omitempty" mapstructure:"publicNetwork,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudAddress returns a new *CloudAddress
func NewCloudAddress() *CloudAddress {

	return &CloudAddress{
		ModelVersion:   1,
		PrivateNetwork: &net.IPNet{},
		PublicNetwork:  &net.IPNet{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudAddress) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudAddress{}

	s.IPVersion = o.IPVersion
	s.Primary = o.Primary
	s.PrivateDNSName = o.PrivateDNSName
	s.PrivateIP = o.PrivateIP
	s.PrivateNetwork = o.PrivateNetwork
	s.PublicDNSName = o.PublicDNSName
	s.PublicIP = o.PublicIP
	s.PublicIPRef = o.PublicIPRef
	s.PublicNetwork = o.PublicNetwork

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudAddress) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudAddress{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.IPVersion = s.IPVersion
	o.Primary = s.Primary
	o.PrivateDNSName = s.PrivateDNSName
	o.PrivateIP = s.PrivateIP
	o.PrivateNetwork = s.PrivateNetwork
	o.PublicDNSName = s.PublicDNSName
	o.PublicIP = s.PublicIP
	o.PublicIPRef = s.PublicIPRef
	o.PublicNetwork = s.PublicNetwork

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudAddress) BleveType() string {

	return "cloudaddress"
}

// DeepCopy returns a deep copy if the CloudAddress.
func (o *CloudAddress) DeepCopy() *CloudAddress {

	if o == nil {
		return nil
	}

	out := &CloudAddress{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudAddress.
func (o *CloudAddress) DeepCopyInto(out *CloudAddress) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudAddress: %s", err))
	}

	*out = *target.(*CloudAddress)
}

// Validate valides the current information stored into the structure.
func (o *CloudAddress) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("IPVersion", string(o.IPVersion)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("IPVersion", string(o.IPVersion), []string{"IPv4", "IPv6"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateOptionalCIDRorIP("privateIP", o.PrivateIP); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateOptionalCIDRorIP("publicIP", o.PublicIP); err != nil {
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
func (*CloudAddress) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudAddressAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudAddressLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudAddress) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudAddressAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudAddress) ValueForAttribute(name string) any {

	switch name {
	case "IPVersion":
		return o.IPVersion
	case "primary":
		return o.Primary
	case "privateDNSName":
		return o.PrivateDNSName
	case "privateIP":
		return o.PrivateIP
	case "privateNetwork":
		return o.PrivateNetwork
	case "publicDNSName":
		return o.PublicDNSName
	case "publicIP":
		return o.PublicIP
	case "publicIPRef":
		return o.PublicIPRef
	case "publicNetwork":
		return o.PublicNetwork
	}

	return nil
}

// CloudAddressAttributesMap represents the map of attribute for CloudAddress.
var CloudAddressAttributesMap = map[string]elemental.AttributeSpecification{
	"IPVersion": {
		AllowedChoices: []string{"IPv4", "IPv6"},
		BSONFieldName:  "ipversion",
		ConvertedName:  "IPVersion",
		Description:    `Designates IPv4 or IPv6.`,
		Exposed:        true,
		Name:           "IPVersion",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"Primary": {
		AllowedChoices: []string{},
		BSONFieldName:  "primary",
		ConvertedName:  "Primary",
		Description:    `Designates the IP address as the primary IP address.`,
		Exposed:        true,
		Name:           "primary",
		Stored:         true,
		Type:           "boolean",
	},
	"PrivateDNSName": {
		AllowedChoices: []string{},
		BSONFieldName:  "privatednsname",
		ConvertedName:  "PrivateDNSName",
		Description:    `The private DNS name associated with the address.`,
		Exposed:        true,
		Name:           "privateDNSName",
		Stored:         true,
		Type:           "string",
	},
	"PrivateIP": {
		AllowedChoices: []string{},
		BSONFieldName:  "privateip",
		ConvertedName:  "PrivateIP",
		Description:    `The private IP address value.`,
		Exposed:        true,
		Name:           "privateIP",
		Stored:         true,
		Type:           "string",
	},
	"PrivateNetwork": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "privatenetwork",
		ConvertedName:  "PrivateNetwork",
		Description: `Internal representation of the private IP to accelerate operations. Not exposed
to users.`,
		Exposed:  true,
		Name:     "privateNetwork",
		ReadOnly: true,
		Stored:   true,
		SubType:  "network",
		Type:     "external",
	},
	"PublicDNSName": {
		AllowedChoices: []string{},
		BSONFieldName:  "publicdnsname",
		ConvertedName:  "PublicDNSName",
		Description:    `The private DNS name associated with the address.`,
		Exposed:        true,
		Name:           "publicDNSName",
		Stored:         true,
		Type:           "string",
	},
	"PublicIP": {
		AllowedChoices: []string{},
		BSONFieldName:  "publicip",
		ConvertedName:  "PublicIP",
		Description:    `The public IP address value.`,
		Exposed:        true,
		Name:           "publicIP",
		Stored:         true,
		Type:           "string",
	},
	"PublicIPRef": {
		AllowedChoices: []string{},
		BSONFieldName:  "publicipref",
		ConvertedName:  "PublicIPRef",
		Description:    `The reference to a public IP address.`,
		Exposed:        true,
		Name:           "publicIPRef",
		Stored:         true,
		Type:           "string",
	},
	"PublicNetwork": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "publicnetwork",
		ConvertedName:  "PublicNetwork",
		Description: `Internal representation of public IP addresses to accelerate operations. Not
exposed to users.`,
		Exposed:  true,
		Name:     "publicNetwork",
		ReadOnly: true,
		Stored:   true,
		SubType:  "network",
		Type:     "external",
	},
}

// CloudAddressLowerCaseAttributesMap represents the map of attribute for CloudAddress.
var CloudAddressLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"ipversion": {
		AllowedChoices: []string{"IPv4", "IPv6"},
		BSONFieldName:  "ipversion",
		ConvertedName:  "IPVersion",
		Description:    `Designates IPv4 or IPv6.`,
		Exposed:        true,
		Name:           "IPVersion",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"primary": {
		AllowedChoices: []string{},
		BSONFieldName:  "primary",
		ConvertedName:  "Primary",
		Description:    `Designates the IP address as the primary IP address.`,
		Exposed:        true,
		Name:           "primary",
		Stored:         true,
		Type:           "boolean",
	},
	"privatednsname": {
		AllowedChoices: []string{},
		BSONFieldName:  "privatednsname",
		ConvertedName:  "PrivateDNSName",
		Description:    `The private DNS name associated with the address.`,
		Exposed:        true,
		Name:           "privateDNSName",
		Stored:         true,
		Type:           "string",
	},
	"privateip": {
		AllowedChoices: []string{},
		BSONFieldName:  "privateip",
		ConvertedName:  "PrivateIP",
		Description:    `The private IP address value.`,
		Exposed:        true,
		Name:           "privateIP",
		Stored:         true,
		Type:           "string",
	},
	"privatenetwork": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "privatenetwork",
		ConvertedName:  "PrivateNetwork",
		Description: `Internal representation of the private IP to accelerate operations. Not exposed
to users.`,
		Exposed:  true,
		Name:     "privateNetwork",
		ReadOnly: true,
		Stored:   true,
		SubType:  "network",
		Type:     "external",
	},
	"publicdnsname": {
		AllowedChoices: []string{},
		BSONFieldName:  "publicdnsname",
		ConvertedName:  "PublicDNSName",
		Description:    `The private DNS name associated with the address.`,
		Exposed:        true,
		Name:           "publicDNSName",
		Stored:         true,
		Type:           "string",
	},
	"publicip": {
		AllowedChoices: []string{},
		BSONFieldName:  "publicip",
		ConvertedName:  "PublicIP",
		Description:    `The public IP address value.`,
		Exposed:        true,
		Name:           "publicIP",
		Stored:         true,
		Type:           "string",
	},
	"publicipref": {
		AllowedChoices: []string{},
		BSONFieldName:  "publicipref",
		ConvertedName:  "PublicIPRef",
		Description:    `The reference to a public IP address.`,
		Exposed:        true,
		Name:           "publicIPRef",
		Stored:         true,
		Type:           "string",
	},
	"publicnetwork": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "publicnetwork",
		ConvertedName:  "PublicNetwork",
		Description: `Internal representation of public IP addresses to accelerate operations. Not
exposed to users.`,
		Exposed:  true,
		Name:     "publicNetwork",
		ReadOnly: true,
		Stored:   true,
		SubType:  "network",
		Type:     "external",
	},
}

type mongoAttributesCloudAddress struct {
	IPVersion      CloudAddressIPVersionValue `bson:"ipversion"`
	Primary        bool                       `bson:"primary"`
	PrivateDNSName string                     `bson:"privatednsname"`
	PrivateIP      string                     `bson:"privateip"`
	PrivateNetwork *net.IPNet                 `bson:"privatenetwork,omitempty"`
	PublicDNSName  string                     `bson:"publicdnsname"`
	PublicIP       string                     `bson:"publicip"`
	PublicIPRef    string                     `bson:"publicipref"`
	PublicNetwork  *net.IPNet                 `bson:"publicnetwork,omitempty"`
}
