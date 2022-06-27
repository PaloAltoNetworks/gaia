package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudScaleSetData represents the model of a cloudscalesetdata
type CloudScaleSetData struct {
	// ID of associated objects with this interface configuration.
	AttachedEntities []string `json:"attachedEntities" msgpack:"attachedEntities" bson:"attachedentities" mapstructure:"attachedEntities,omitempty"`

	// Scale set related parameters.
	NetworkConfigs []*CloudNetworkConfigData `json:"networkConfigs" msgpack:"networkConfigs" bson:"networkconfigs" mapstructure:"networkConfigs,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudScaleSetData returns a new *CloudScaleSetData
func NewCloudScaleSetData() *CloudScaleSetData {

	return &CloudScaleSetData{
		ModelVersion:     1,
		AttachedEntities: []string{},
		NetworkConfigs:   []*CloudNetworkConfigData{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudScaleSetData) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudScaleSetData{}

	s.AttachedEntities = o.AttachedEntities
	s.NetworkConfigs = o.NetworkConfigs

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudScaleSetData) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudScaleSetData{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.AttachedEntities = s.AttachedEntities
	o.NetworkConfigs = s.NetworkConfigs

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudScaleSetData) BleveType() string {

	return "cloudscalesetdata"
}

// DeepCopy returns a deep copy if the CloudScaleSetData.
func (o *CloudScaleSetData) DeepCopy() *CloudScaleSetData {

	if o == nil {
		return nil
	}

	out := &CloudScaleSetData{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudScaleSetData.
func (o *CloudScaleSetData) DeepCopyInto(out *CloudScaleSetData) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudScaleSetData: %s", err))
	}

	*out = *target.(*CloudScaleSetData)
}

// Validate valides the current information stored into the structure.
func (o *CloudScaleSetData) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.NetworkConfigs {
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
func (*CloudScaleSetData) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudScaleSetDataAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudScaleSetDataLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudScaleSetData) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudScaleSetDataAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudScaleSetData) ValueForAttribute(name string) interface{} {

	switch name {
	case "attachedEntities":
		return o.AttachedEntities
	case "networkConfigs":
		return o.NetworkConfigs
	}

	return nil
}

// CloudScaleSetDataAttributesMap represents the map of attribute for CloudScaleSetData.
var CloudScaleSetDataAttributesMap = map[string]elemental.AttributeSpecification{
	"AttachedEntities": {
		AllowedChoices: []string{},
		BSONFieldName:  "attachedentities",
		ConvertedName:  "AttachedEntities",
		Description:    `ID of associated objects with this interface configuration.`,
		Exposed:        true,
		Name:           "attachedEntities",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"NetworkConfigs": {
		AllowedChoices: []string{},
		BSONFieldName:  "networkconfigs",
		ConvertedName:  "NetworkConfigs",
		Description:    `Scale set related parameters.`,
		Exposed:        true,
		Name:           "networkConfigs",
		Stored:         true,
		SubType:        "cloudnetworkconfigdata",
		Type:           "refList",
	},
}

// CloudScaleSetDataLowerCaseAttributesMap represents the map of attribute for CloudScaleSetData.
var CloudScaleSetDataLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"attachedentities": {
		AllowedChoices: []string{},
		BSONFieldName:  "attachedentities",
		ConvertedName:  "AttachedEntities",
		Description:    `ID of associated objects with this interface configuration.`,
		Exposed:        true,
		Name:           "attachedEntities",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"networkconfigs": {
		AllowedChoices: []string{},
		BSONFieldName:  "networkconfigs",
		ConvertedName:  "NetworkConfigs",
		Description:    `Scale set related parameters.`,
		Exposed:        true,
		Name:           "networkConfigs",
		Stored:         true,
		SubType:        "cloudnetworkconfigdata",
		Type:           "refList",
	},
}

type mongoAttributesCloudScaleSetData struct {
	AttachedEntities []string                  `bson:"attachedentities"`
	NetworkConfigs   []*CloudNetworkConfigData `bson:"networkconfigs"`
}
