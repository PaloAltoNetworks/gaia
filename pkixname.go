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

type mongoAttributesPKIXName struct {
}
