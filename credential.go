package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// Credential represents the model of a credential
type Credential struct {
	// The URL of the Microsegmentation Console API.
	APIURL string `json:"APIURL" msgpack:"APIURL" bson:"-" mapstructure:"APIURL,omitempty"`

	// The ID of the app credential.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// The base64-encoded certificate.
	Certificate string `json:"certificate" msgpack:"certificate" bson:"-" mapstructure:"certificate,omitempty"`

	// The base64-encoded certificate authority.
	CertificateAuthority string `json:"certificateAuthority" msgpack:"certificateAuthority" bson:"-" mapstructure:"certificateAuthority,omitempty"`

	// The base64-encoded certificate key.
	CertificateKey string `json:"certificateKey" msgpack:"certificateKey" bson:"-" mapstructure:"certificateKey,omitempty"`

	// The name of the app credential.
	Name string `json:"name" msgpack:"name" bson:"-" mapstructure:"name,omitempty"`

	// The namespace of the app credential.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCredential returns a new *Credential
func NewCredential() *Credential {

	return &Credential{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Credential) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCredential{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Credential) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCredential{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *Credential) BleveType() string {

	return "credential"
}

// DeepCopy returns a deep copy if the Credential.
func (o *Credential) DeepCopy() *Credential {

	if o == nil {
		return nil
	}

	out := &Credential{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Credential.
func (o *Credential) DeepCopyInto(out *Credential) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Credential: %s", err))
	}

	*out = *target.(*Credential)
}

// Validate valides the current information stored into the structure.
func (o *Credential) Validate() error {

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
func (*Credential) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CredentialAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CredentialLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Credential) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CredentialAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Credential) ValueForAttribute(name string) interface{} {

	switch name {
	case "APIURL":
		return o.APIURL
	case "ID":
		return o.ID
	case "certificate":
		return o.Certificate
	case "certificateAuthority":
		return o.CertificateAuthority
	case "certificateKey":
		return o.CertificateKey
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	}

	return nil
}

// CredentialAttributesMap represents the map of attribute for Credential.
var CredentialAttributesMap = map[string]elemental.AttributeSpecification{
	"APIURL": {
		AllowedChoices: []string{},
		ConvertedName:  "APIURL",
		Description:    `The URL of the Microsegmentation Console API.`,
		Exposed:        true,
		Name:           "APIURL",
		Type:           "string",
	},
	"ID": {
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `The ID of the app credential.`,
		Exposed:        true,
		Name:           "ID",
		Type:           "string",
	},
	"Certificate": {
		AllowedChoices: []string{},
		ConvertedName:  "Certificate",
		Description:    `The base64-encoded certificate.`,
		Exposed:        true,
		Name:           "certificate",
		Type:           "string",
	},
	"CertificateAuthority": {
		AllowedChoices: []string{},
		ConvertedName:  "CertificateAuthority",
		Description:    `The base64-encoded certificate authority.`,
		Exposed:        true,
		Name:           "certificateAuthority",
		Type:           "string",
	},
	"CertificateKey": {
		AllowedChoices: []string{},
		ConvertedName:  "CertificateKey",
		Description:    `The base64-encoded certificate key.`,
		Exposed:        true,
		Name:           "certificateKey",
		Type:           "string",
	},
	"Name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `The name of the app credential.`,
		Exposed:        true,
		Name:           "name",
		Type:           "string",
	},
	"Namespace": {
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `The namespace of the app credential.`,
		Exposed:        true,
		Name:           "namespace",
		Type:           "string",
	},
}

// CredentialLowerCaseAttributesMap represents the map of attribute for Credential.
var CredentialLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"apiurl": {
		AllowedChoices: []string{},
		ConvertedName:  "APIURL",
		Description:    `The URL of the Microsegmentation Console API.`,
		Exposed:        true,
		Name:           "APIURL",
		Type:           "string",
	},
	"id": {
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `The ID of the app credential.`,
		Exposed:        true,
		Name:           "ID",
		Type:           "string",
	},
	"certificate": {
		AllowedChoices: []string{},
		ConvertedName:  "Certificate",
		Description:    `The base64-encoded certificate.`,
		Exposed:        true,
		Name:           "certificate",
		Type:           "string",
	},
	"certificateauthority": {
		AllowedChoices: []string{},
		ConvertedName:  "CertificateAuthority",
		Description:    `The base64-encoded certificate authority.`,
		Exposed:        true,
		Name:           "certificateAuthority",
		Type:           "string",
	},
	"certificatekey": {
		AllowedChoices: []string{},
		ConvertedName:  "CertificateKey",
		Description:    `The base64-encoded certificate key.`,
		Exposed:        true,
		Name:           "certificateKey",
		Type:           "string",
	},
	"name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `The name of the app credential.`,
		Exposed:        true,
		Name:           "name",
		Type:           "string",
	},
	"namespace": {
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `The namespace of the app credential.`,
		Exposed:        true,
		Name:           "namespace",
		Type:           "string",
	},
}

type mongoAttributesCredential struct {
}
