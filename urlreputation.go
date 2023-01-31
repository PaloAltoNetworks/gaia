// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// URLReputation represents the model of a urlreputation
type URLReputation struct {
	// The IP address or FQDN.
	URL string `json:"URL,omitempty" msgpack:"URL,omitempty" bson:"-" mapstructure:"URL,omitempty"`

	// The categories codes related to the risk level.
	Categories []int `json:"categories,omitempty" msgpack:"categories,omitempty" bson:"-" mapstructure:"categories,omitempty"`

	// The evidence related to the risk level.
	Evidence []string `json:"evidence,omitempty" msgpack:"evidence,omitempty" bson:"-" mapstructure:"evidence,omitempty"`

	// The associated IP/FQDN risk level.
	RiskLevel string `json:"riskLevel,omitempty" msgpack:"riskLevel,omitempty" bson:"-" mapstructure:"riskLevel,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewURLReputation returns a new *URLReputation
func NewURLReputation() *URLReputation {

	return &URLReputation{
		ModelVersion: 1,
		Categories:   []int{},
		Evidence:     []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *URLReputation) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesURLReputation{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *URLReputation) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesURLReputation{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *URLReputation) BleveType() string {

	return "urlreputation"
}

// DeepCopy returns a deep copy if the URLReputation.
func (o *URLReputation) DeepCopy() *URLReputation {

	if o == nil {
		return nil
	}

	out := &URLReputation{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *URLReputation.
func (o *URLReputation) DeepCopyInto(out *URLReputation) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy URLReputation: %s", err))
	}

	*out = *target.(*URLReputation)
}

// Validate valides the current information stored into the structure.
func (o *URLReputation) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*URLReputation) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := URLReputationAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return URLReputationLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*URLReputation) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return URLReputationAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *URLReputation) ValueForAttribute(name string) any {

	switch name {
	case "URL":
		return o.URL
	case "categories":
		return o.Categories
	case "evidence":
		return o.Evidence
	case "riskLevel":
		return o.RiskLevel
	}

	return nil
}

// URLReputationAttributesMap represents the map of attribute for URLReputation.
var URLReputationAttributesMap = map[string]elemental.AttributeSpecification{
	"URL": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "URL",
		Description:    `The IP address or FQDN.`,
		Exposed:        true,
		Name:           "URL",
		ReadOnly:       true,
		Type:           "string",
	},
	"Categories": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Categories",
		Description:    `The categories codes related to the risk level.`,
		Exposed:        true,
		Name:           "categories",
		ReadOnly:       true,
		SubType:        "integer",
		Type:           "list",
	},
	"Evidence": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Evidence",
		Description:    `The evidence related to the risk level.`,
		Exposed:        true,
		Name:           "evidence",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"RiskLevel": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "RiskLevel",
		Description:    `The associated IP/FQDN risk level.`,
		Exposed:        true,
		Name:           "riskLevel",
		ReadOnly:       true,
		Type:           "string",
	},
}

// URLReputationLowerCaseAttributesMap represents the map of attribute for URLReputation.
var URLReputationLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"url": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "URL",
		Description:    `The IP address or FQDN.`,
		Exposed:        true,
		Name:           "URL",
		ReadOnly:       true,
		Type:           "string",
	},
	"categories": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Categories",
		Description:    `The categories codes related to the risk level.`,
		Exposed:        true,
		Name:           "categories",
		ReadOnly:       true,
		SubType:        "integer",
		Type:           "list",
	},
	"evidence": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Evidence",
		Description:    `The evidence related to the risk level.`,
		Exposed:        true,
		Name:           "evidence",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"risklevel": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "RiskLevel",
		Description:    `The associated IP/FQDN risk level.`,
		Exposed:        true,
		Name:           "riskLevel",
		ReadOnly:       true,
		Type:           "string",
	},
}

type mongoAttributesURLReputation struct {
}
