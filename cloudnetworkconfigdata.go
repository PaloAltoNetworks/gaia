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
	IPConfigurations []*CloudIPConfiguration `json:"IPConfigurations" msgpack:"IPConfigurations" bson:"ipconfigurations" mapstructure:"IPConfigurations,omitempty"`

	// Network configuration name of the NIC associated in the Scale Set.
	NetworkConfigName string `json:"networkConfigName" msgpack:"networkConfigName" bson:"networkconfigname" mapstructure:"networkConfigName,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudNetworkConfigData returns a new *CloudNetworkConfigData
func NewCloudNetworkConfigData() *CloudNetworkConfigData {

	return &CloudNetworkConfigData{
		ModelVersion:     1,
		IPConfigurations: []*CloudIPConfiguration{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudNetworkConfigData) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudNetworkConfigData{}

	s.IPConfigurations = o.IPConfigurations
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

	o.IPConfigurations = s.IPConfigurations
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

	for _, sub := range o.IPConfigurations {
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
	case "IPConfigurations":
		return o.IPConfigurations
	case "networkConfigName":
		return o.NetworkConfigName
	}

	return nil
}

// CloudNetworkConfigDataAttributesMap represents the map of attribute for CloudNetworkConfigData.
var CloudNetworkConfigDataAttributesMap = map[string]elemental.AttributeSpecification{
	"IPConfigurations": {
		AllowedChoices: []string{},
		BSONFieldName:  "ipconfigurations",
		ConvertedName:  "IPConfigurations",
		Description:    `IP configurations of the NICs in the Scale Set.`,
		Exposed:        true,
		Name:           "IPConfigurations",
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
		ConvertedName:  "IPConfigurations",
		Description:    `IP configurations of the NICs in the Scale Set.`,
		Exposed:        true,
		Name:           "IPConfigurations",
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
	IPConfigurations  []*CloudIPConfiguration `bson:"ipconfigurations"`
	NetworkConfigName string                  `bson:"networkconfigname"`
}
