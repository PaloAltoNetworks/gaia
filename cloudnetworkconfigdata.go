package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudNetworkConfigData represents the model of a cloudnetworkconfigdata
type CloudNetworkConfigData struct {
	// IP configurations of the NICs in the Scale Set.
	IpConfigurations []*CloudIPConfiguration `json:"ipConfigurations" msgpack:"ipConfigurations" bson:"ipconfigurations" mapstructure:"ipConfigurations,omitempty"`

	// Network configuration name of the NIC associated in the Scale Set.
	NetworkConfigName string `json:"networkConfigName" msgpack:"networkConfigName" bson:"networkconfigname" mapstructure:"networkConfigName,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudNetworkConfigData returns a new *CloudNetworkConfigData
func NewCloudNetworkConfigData() *CloudNetworkConfigData {

	return &CloudNetworkConfigData{
		ModelVersion:     1,
		IpConfigurations: []*CloudIPConfiguration{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudNetworkConfigData) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudNetworkConfigData{}

	s.IpConfigurations = o.IpConfigurations
	s.NetworkConfigName = o.NetworkConfigName

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudNetworkConfigData) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudNetworkConfigData{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.IpConfigurations = s.IpConfigurations
	o.NetworkConfigName = s.NetworkConfigName

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudNetworkConfigData) BleveType() string {

	return "cloudnetworkconfigdata"
}

// DeepCopy returns a deep copy if the CloudNetworkConfigData.
func (o *CloudNetworkConfigData) DeepCopy() *CloudNetworkConfigData {

	if o == nil {
		return nil
	}

	out := &CloudNetworkConfigData{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudNetworkConfigData.
func (o *CloudNetworkConfigData) DeepCopyInto(out *CloudNetworkConfigData) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudNetworkConfigData: %s", err))
	}

	*out = *target.(*CloudNetworkConfigData)
}

// Validate valides the current information stored into the structure.
func (o *CloudNetworkConfigData) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.IpConfigurations {
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
func (*CloudNetworkConfigData) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudNetworkConfigDataAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudNetworkConfigDataLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudNetworkConfigData) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudNetworkConfigDataAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudNetworkConfigData) ValueForAttribute(name string) interface{} {

	switch name {
	case "ipConfigurations":
		return o.IpConfigurations
	case "networkConfigName":
		return o.NetworkConfigName
	}

	return nil
}

// CloudNetworkConfigDataAttributesMap represents the map of attribute for CloudNetworkConfigData.
var CloudNetworkConfigDataAttributesMap = map[string]elemental.AttributeSpecification{
	"IpConfigurations": {
		AllowedChoices: []string{},
		BSONFieldName:  "ipconfigurations",
		ConvertedName:  "IpConfigurations",
		Description:    `IP configurations of the NICs in the Scale Set.`,
		Exposed:        true,
		Name:           "ipConfigurations",
		Stored:         true,
		SubType:        "cloudipconfiguration",
		Type:           "refList",
	},
	"NetworkConfigName": {
		AllowedChoices: []string{},
		BSONFieldName:  "networkconfigname",
		ConvertedName:  "NetworkConfigName",
		Description:    `Network configuration name of the NIC associated in the Scale Set.`,
		Exposed:        true,
		Name:           "networkConfigName",
		Stored:         true,
		Type:           "string",
	},
}

// CloudNetworkConfigDataLowerCaseAttributesMap represents the map of attribute for CloudNetworkConfigData.
var CloudNetworkConfigDataLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"ipconfigurations": {
		AllowedChoices: []string{},
		BSONFieldName:  "ipconfigurations",
		ConvertedName:  "IpConfigurations",
		Description:    `IP configurations of the NICs in the Scale Set.`,
		Exposed:        true,
		Name:           "ipConfigurations",
		Stored:         true,
		SubType:        "cloudipconfiguration",
		Type:           "refList",
	},
	"networkconfigname": {
		AllowedChoices: []string{},
		BSONFieldName:  "networkconfigname",
		ConvertedName:  "NetworkConfigName",
		Description:    `Network configuration name of the NIC associated in the Scale Set.`,
		Exposed:        true,
		Name:           "networkConfigName",
		Stored:         true,
		Type:           "string",
	},
}

type mongoAttributesCloudNetworkConfigData struct {
	IpConfigurations  []*CloudIPConfiguration `bson:"ipconfigurations"`
	NetworkConfigName string                  `bson:"networkconfigname"`
}
