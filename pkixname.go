package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PKIXName represents the model of a pkixname
type PKIXName struct {
	// Represents the Common Name field.
	CommonName string `json:"commonName" msgpack:"commonName" bson:"-" mapstructure:"commonName,omitempty"`

	// Represents the Country field.
	Country []string `json:"country" msgpack:"country" bson:"-" mapstructure:"country,omitempty"`

	// Represents the Locality field.
	Locality []string `json:"locality" msgpack:"locality" bson:"-" mapstructure:"locality,omitempty"`

	// Represents the Organization field.
	Organization []string `json:"organization" msgpack:"organization" bson:"-" mapstructure:"organization,omitempty"`

	// Represents the Organizational Unit field.
	OrganizationalUnit []string `json:"organizationalUnit" msgpack:"organizationalUnit" bson:"-" mapstructure:"organizationalUnit,omitempty"`

	// Represents the Postal Code field.
	PostalCode []string `json:"postalCode" msgpack:"postalCode" bson:"-" mapstructure:"postalCode,omitempty"`

	// Represents the Province field.
	Province []string `json:"province" msgpack:"province" bson:"-" mapstructure:"province,omitempty"`

	// Represents the Street Address field.
	StreetAddress []string `json:"streetAddress" msgpack:"streetAddress" bson:"-" mapstructure:"streetAddress,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPKIXName returns a new *PKIXName
func NewPKIXName() *PKIXName {

	return &PKIXName{
		ModelVersion:       1,
		Country:            []string{},
		Locality:           []string{},
		Organization:       []string{},
		OrganizationalUnit: []string{},
		PostalCode:         []string{},
		Province:           []string{},
		StreetAddress:      []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PKIXName) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPKIXName{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PKIXName) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPKIXName{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *PKIXName) BleveType() string {

	return "pkixname"
}

// DeepCopy returns a deep copy if the PKIXName.
func (o *PKIXName) DeepCopy() *PKIXName {

	if o == nil {
		return nil
	}

	out := &PKIXName{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PKIXName.
func (o *PKIXName) DeepCopyInto(out *PKIXName) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PKIXName: %s", err))
	}

	*out = *target.(*PKIXName)
}

// Validate valides the current information stored into the structure.
func (o *PKIXName) Validate() error {

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
func (*PKIXName) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PKIXNameAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PKIXNameLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PKIXName) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PKIXNameAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PKIXName) ValueForAttribute(name string) interface{} {

	switch name {
	case "commonName":
		return o.CommonName
	case "country":
		return o.Country
	case "locality":
		return o.Locality
	case "organization":
		return o.Organization
	case "organizationalUnit":
		return o.OrganizationalUnit
	case "postalCode":
		return o.PostalCode
	case "province":
		return o.Province
	case "streetAddress":
		return o.StreetAddress
	}

	return nil
}

// PKIXNameAttributesMap represents the map of attribute for PKIXName.
var PKIXNameAttributesMap = map[string]elemental.AttributeSpecification{
	"CommonName": {
		AllowedChoices: []string{},
		ConvertedName:  "CommonName",
		Description:    `Represents the Common Name field.`,
		Exposed:        true,
		Name:           "commonName",
		Type:           "string",
	},
	"Country": {
		AllowedChoices: []string{},
		ConvertedName:  "Country",
		Description:    `Represents the Country field.`,
		Exposed:        true,
		Name:           "country",
		SubType:        "string",
		Type:           "list",
	},
	"Locality": {
		AllowedChoices: []string{},
		ConvertedName:  "Locality",
		Description:    `Represents the Locality field.`,
		Exposed:        true,
		Name:           "locality",
		SubType:        "string",
		Type:           "list",
	},
	"Organization": {
		AllowedChoices: []string{},
		ConvertedName:  "Organization",
		Description:    `Represents the Organization field.`,
		Exposed:        true,
		Name:           "organization",
		SubType:        "string",
		Type:           "list",
	},
	"OrganizationalUnit": {
		AllowedChoices: []string{},
		ConvertedName:  "OrganizationalUnit",
		Description:    `Represents the Organizational Unit field.`,
		Exposed:        true,
		Name:           "organizationalUnit",
		SubType:        "string",
		Type:           "list",
	},
	"PostalCode": {
		AllowedChoices: []string{},
		ConvertedName:  "PostalCode",
		Description:    `Represents the Postal Code field.`,
		Exposed:        true,
		Name:           "postalCode",
		SubType:        "string",
		Type:           "list",
	},
	"Province": {
		AllowedChoices: []string{},
		ConvertedName:  "Province",
		Description:    `Represents the Province field.`,
		Exposed:        true,
		Name:           "province",
		SubType:        "string",
		Type:           "list",
	},
	"StreetAddress": {
		AllowedChoices: []string{},
		ConvertedName:  "StreetAddress",
		Description:    `Represents the Street Address field.`,
		Exposed:        true,
		Name:           "streetAddress",
		SubType:        "string",
		Type:           "list",
	},
}

// PKIXNameLowerCaseAttributesMap represents the map of attribute for PKIXName.
var PKIXNameLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"commonname": {
		AllowedChoices: []string{},
		ConvertedName:  "CommonName",
		Description:    `Represents the Common Name field.`,
		Exposed:        true,
		Name:           "commonName",
		Type:           "string",
	},
	"country": {
		AllowedChoices: []string{},
		ConvertedName:  "Country",
		Description:    `Represents the Country field.`,
		Exposed:        true,
		Name:           "country",
		SubType:        "string",
		Type:           "list",
	},
	"locality": {
		AllowedChoices: []string{},
		ConvertedName:  "Locality",
		Description:    `Represents the Locality field.`,
		Exposed:        true,
		Name:           "locality",
		SubType:        "string",
		Type:           "list",
	},
	"organization": {
		AllowedChoices: []string{},
		ConvertedName:  "Organization",
		Description:    `Represents the Organization field.`,
		Exposed:        true,
		Name:           "organization",
		SubType:        "string",
		Type:           "list",
	},
	"organizationalunit": {
		AllowedChoices: []string{},
		ConvertedName:  "OrganizationalUnit",
		Description:    `Represents the Organizational Unit field.`,
		Exposed:        true,
		Name:           "organizationalUnit",
		SubType:        "string",
		Type:           "list",
	},
	"postalcode": {
		AllowedChoices: []string{},
		ConvertedName:  "PostalCode",
		Description:    `Represents the Postal Code field.`,
		Exposed:        true,
		Name:           "postalCode",
		SubType:        "string",
		Type:           "list",
	},
	"province": {
		AllowedChoices: []string{},
		ConvertedName:  "Province",
		Description:    `Represents the Province field.`,
		Exposed:        true,
		Name:           "province",
		SubType:        "string",
		Type:           "list",
	},
	"streetaddress": {
		AllowedChoices: []string{},
		ConvertedName:  "StreetAddress",
		Description:    `Represents the Street Address field.`,
		Exposed:        true,
		Name:           "streetAddress",
		SubType:        "string",
		Type:           "list",
	},
}

type mongoAttributesPKIXName struct {
}
