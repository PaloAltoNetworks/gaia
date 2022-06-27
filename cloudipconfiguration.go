package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudIPConfiguration represents the model of a cloudipconfiguration
type CloudIPConfiguration struct {
	// List of IP addresses/subnets (IPv4 or IPv6) associated with the
	// interface.
	Addresses []*CloudAddress `json:"addresses" msgpack:"addresses" bson:"addresses" mapstructure:"addresses,omitempty"`

	// Backend pool of the load balancer (if any) fronting the Scale Set.
	BackendPool []string `json:"backendPool" msgpack:"backendPool" bson:"backendpool" mapstructure:"backendPool,omitempty"`

	// If the IP Configuration has a public IP.
	HasPublicIP bool `json:"hasPublicIP" msgpack:"hasPublicIP" bson:"haspublicip" mapstructure:"hasPublicIP,omitempty"`

	// IP configuration of the NICs associated to the Scale Set.
	IpConfigName string `json:"ipConfigName" msgpack:"ipConfigName" bson:"ipconfigname" mapstructure:"ipConfigName,omitempty"`

	// Subnet of the NIC associated to the Scale Set.
	Subnet string `json:"subnet" msgpack:"subnet" bson:"subnet" mapstructure:"subnet,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudIPConfiguration returns a new *CloudIPConfiguration
func NewCloudIPConfiguration() *CloudIPConfiguration {

	return &CloudIPConfiguration{
		ModelVersion: 1,
		Addresses:    []*CloudAddress{},
		BackendPool:  []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudIPConfiguration) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudIPConfiguration{}

	s.Addresses = o.Addresses
	s.BackendPool = o.BackendPool
	s.HasPublicIP = o.HasPublicIP
	s.IpConfigName = o.IpConfigName
	s.Subnet = o.Subnet

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudIPConfiguration) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudIPConfiguration{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Addresses = s.Addresses
	o.BackendPool = s.BackendPool
	o.HasPublicIP = s.HasPublicIP
	o.IpConfigName = s.IpConfigName
	o.Subnet = s.Subnet

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudIPConfiguration) BleveType() string {

	return "cloudipconfiguration"
}

// DeepCopy returns a deep copy if the CloudIPConfiguration.
func (o *CloudIPConfiguration) DeepCopy() *CloudIPConfiguration {

	if o == nil {
		return nil
	}

	out := &CloudIPConfiguration{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudIPConfiguration.
func (o *CloudIPConfiguration) DeepCopyInto(out *CloudIPConfiguration) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudIPConfiguration: %s", err))
	}

	*out = *target.(*CloudIPConfiguration)
}

// Validate valides the current information stored into the structure.
func (o *CloudIPConfiguration) Validate() error {

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

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*CloudIPConfiguration) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudIPConfigurationAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudIPConfigurationLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudIPConfiguration) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudIPConfigurationAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudIPConfiguration) ValueForAttribute(name string) interface{} {

	switch name {
	case "addresses":
		return o.Addresses
	case "backendPool":
		return o.BackendPool
	case "hasPublicIP":
		return o.HasPublicIP
	case "ipConfigName":
		return o.IpConfigName
	case "subnet":
		return o.Subnet
	}

	return nil
}

// CloudIPConfigurationAttributesMap represents the map of attribute for CloudIPConfiguration.
var CloudIPConfigurationAttributesMap = map[string]elemental.AttributeSpecification{
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
	"BackendPool": {
		AllowedChoices: []string{},
		BSONFieldName:  "backendpool",
		ConvertedName:  "BackendPool",
		Description:    `Backend pool of the load balancer (if any) fronting the Scale Set.`,
		Exposed:        true,
		Name:           "backendPool",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"HasPublicIP": {
		AllowedChoices: []string{},
		BSONFieldName:  "haspublicip",
		ConvertedName:  "HasPublicIP",
		Description:    `If the IP Configuration has a public IP.`,
		Exposed:        true,
		Name:           "hasPublicIP",
		Stored:         true,
		Type:           "boolean",
	},
	"IpConfigName": {
		AllowedChoices: []string{},
		BSONFieldName:  "ipconfigname",
		ConvertedName:  "IpConfigName",
		Description:    `IP configuration of the NICs associated to the Scale Set.`,
		Exposed:        true,
		Name:           "ipConfigName",
		Stored:         true,
		Type:           "string",
	},
	"Subnet": {
		AllowedChoices: []string{},
		BSONFieldName:  "subnet",
		ConvertedName:  "Subnet",
		Description:    `Subnet of the NIC associated to the Scale Set.`,
		Exposed:        true,
		Name:           "subnet",
		Stored:         true,
		Type:           "string",
	},
}

// CloudIPConfigurationLowerCaseAttributesMap represents the map of attribute for CloudIPConfiguration.
var CloudIPConfigurationLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"backendpool": {
		AllowedChoices: []string{},
		BSONFieldName:  "backendpool",
		ConvertedName:  "BackendPool",
		Description:    `Backend pool of the load balancer (if any) fronting the Scale Set.`,
		Exposed:        true,
		Name:           "backendPool",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"haspublicip": {
		AllowedChoices: []string{},
		BSONFieldName:  "haspublicip",
		ConvertedName:  "HasPublicIP",
		Description:    `If the IP Configuration has a public IP.`,
		Exposed:        true,
		Name:           "hasPublicIP",
		Stored:         true,
		Type:           "boolean",
	},
	"ipconfigname": {
		AllowedChoices: []string{},
		BSONFieldName:  "ipconfigname",
		ConvertedName:  "IpConfigName",
		Description:    `IP configuration of the NICs associated to the Scale Set.`,
		Exposed:        true,
		Name:           "ipConfigName",
		Stored:         true,
		Type:           "string",
	},
	"subnet": {
		AllowedChoices: []string{},
		BSONFieldName:  "subnet",
		ConvertedName:  "Subnet",
		Description:    `Subnet of the NIC associated to the Scale Set.`,
		Exposed:        true,
		Name:           "subnet",
		Stored:         true,
		Type:           "string",
	},
}

type mongoAttributesCloudIPConfiguration struct {
	Addresses    []*CloudAddress `bson:"addresses"`
	BackendPool  []string        `bson:"backendpool"`
	HasPublicIP  bool            `bson:"haspublicip"`
	IpConfigName string          `bson:"ipconfigname"`
	Subnet       string          `bson:"subnet"`
}
