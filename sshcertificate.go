package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// SSHCertificateTypeValue represents the possible values for attribute "type".
type SSHCertificateTypeValue string

const (
	// SSHCertificateTypeHost represents the value Host.
	SSHCertificateTypeHost SSHCertificateTypeValue = "Host"

	// SSHCertificateTypeUser represents the value User.
	SSHCertificateTypeUser SSHCertificateTypeValue = "User"
)

// SSHCertificateIdentity represents the Identity of the object.
var SSHCertificateIdentity = elemental.Identity{
	Name:     "sshcertificate",
	Category: "sshcertificates",
	Package:  "barret",
	Private:  true,
}

// SSHCertificatesList represents a list of SSHCertificates
type SSHCertificatesList []*SSHCertificate

// Identity returns the identity of the objects in the list.
func (o SSHCertificatesList) Identity() elemental.Identity {

	return SSHCertificateIdentity
}

// Copy returns a pointer to a copy the SSHCertificatesList.
func (o SSHCertificatesList) Copy() elemental.Identifiables {

	copy := append(SSHCertificatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SSHCertificatesList.
func (o SSHCertificatesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SSHCertificatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SSHCertificate))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SSHCertificatesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SSHCertificatesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the SSHCertificatesList converted to SparseSSHCertificatesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o SSHCertificatesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseSSHCertificatesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseSSHCertificate)
	}

	return out
}

// Version returns the version of the content.
func (o SSHCertificatesList) Version() int {

	return 1
}

// SSHCertificate represents the model of a sshcertificate
type SSHCertificate struct {
	// Contains the signed SSH certificate in OpenSSH format.
	Certificate string `json:"certificate" msgpack:"certificate" bson:"-" mapstructure:"certificate,omitempty"`

	// List of extensions to set in the SSH certificate.
	Extensions map[string]string `json:"extensions" msgpack:"extensions" bson:"-" mapstructure:"extensions,omitempty"`

	// List of options to set in the SSH certificate.
	Options map[string]string `json:"options" msgpack:"options" bson:"-" mapstructure:"options,omitempty"`

	// List of principals to set in the SSH certificate.
	Principals []string `json:"principals" msgpack:"principals" bson:"-" mapstructure:"principals,omitempty"`

	// Contains the public key to sign in OpenSSH format.
	PublicKey string `json:"publicKey" msgpack:"publicKey" bson:"-" mapstructure:"publicKey,omitempty"`

	// The identifier of the CA to use to sign the certificate.
	SignerID string `json:"signerID" msgpack:"signerID" bson:"-" mapstructure:"signerID,omitempty"`

	// Type of SSH certificate.
	Type SSHCertificateTypeValue `json:"type" msgpack:"type" bson:"-" mapstructure:"type,omitempty"`

	// Set the validity of the SSH certificate.
	Validity string `json:"validity" msgpack:"validity" bson:"-" mapstructure:"validity,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSSHCertificate returns a new *SSHCertificate
func NewSSHCertificate() *SSHCertificate {

	return &SSHCertificate{
		ModelVersion: 1,
		Extensions:   map[string]string{},
		Options:      map[string]string{},
		Principals:   []string{},
		Type:         SSHCertificateTypeUser,
		Validity:     "1h",
	}
}

// Identity returns the Identity of the object.
func (o *SSHCertificate) Identity() elemental.Identity {

	return SSHCertificateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SSHCertificate) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SSHCertificate) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SSHCertificate) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSSHCertificate{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SSHCertificate) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSSHCertificate{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SSHCertificate) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *SSHCertificate) BleveType() string {

	return "sshcertificate"
}

// DefaultOrder returns the list of default ordering fields.
func (o *SSHCertificate) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *SSHCertificate) Doc() string {

	return `Internal api to deliver SSH certificate.`
}

func (o *SSHCertificate) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *SSHCertificate) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseSSHCertificate{
			Certificate: &o.Certificate,
			Extensions:  &o.Extensions,
			Options:     &o.Options,
			Principals:  &o.Principals,
			PublicKey:   &o.PublicKey,
			SignerID:    &o.SignerID,
			Type:        &o.Type,
			Validity:    &o.Validity,
		}
	}

	sp := &SparseSSHCertificate{}
	for _, f := range fields {
		switch f {
		case "certificate":
			sp.Certificate = &(o.Certificate)
		case "extensions":
			sp.Extensions = &(o.Extensions)
		case "options":
			sp.Options = &(o.Options)
		case "principals":
			sp.Principals = &(o.Principals)
		case "publicKey":
			sp.PublicKey = &(o.PublicKey)
		case "signerID":
			sp.SignerID = &(o.SignerID)
		case "type":
			sp.Type = &(o.Type)
		case "validity":
			sp.Validity = &(o.Validity)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseSSHCertificate to the object.
func (o *SSHCertificate) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseSSHCertificate)
	if so.Certificate != nil {
		o.Certificate = *so.Certificate
	}
	if so.Extensions != nil {
		o.Extensions = *so.Extensions
	}
	if so.Options != nil {
		o.Options = *so.Options
	}
	if so.Principals != nil {
		o.Principals = *so.Principals
	}
	if so.PublicKey != nil {
		o.PublicKey = *so.PublicKey
	}
	if so.SignerID != nil {
		o.SignerID = *so.SignerID
	}
	if so.Type != nil {
		o.Type = *so.Type
	}
	if so.Validity != nil {
		o.Validity = *so.Validity
	}
}

// DeepCopy returns a deep copy if the SSHCertificate.
func (o *SSHCertificate) DeepCopy() *SSHCertificate {

	if o == nil {
		return nil
	}

	out := &SSHCertificate{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SSHCertificate.
func (o *SSHCertificate) DeepCopyInto(out *SSHCertificate) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SSHCertificate: %s", err))
	}

	*out = *target.(*SSHCertificate)
}

// Validate valides the current information stored into the structure.
func (o *SSHCertificate) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("publicKey", o.PublicKey); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("signerID", o.SignerID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"User", "Host"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTimeDuration("validity", o.Validity); err != nil {
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
func (*SSHCertificate) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := SSHCertificateAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return SSHCertificateLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*SSHCertificate) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return SSHCertificateAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *SSHCertificate) ValueForAttribute(name string) interface{} {

	switch name {
	case "certificate":
		return o.Certificate
	case "extensions":
		return o.Extensions
	case "options":
		return o.Options
	case "principals":
		return o.Principals
	case "publicKey":
		return o.PublicKey
	case "signerID":
		return o.SignerID
	case "type":
		return o.Type
	case "validity":
		return o.Validity
	}

	return nil
}

// SSHCertificateAttributesMap represents the map of attribute for SSHCertificate.
var SSHCertificateAttributesMap = map[string]elemental.AttributeSpecification{
	"Certificate": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Certificate",
		Description:    `Contains the signed SSH certificate in OpenSSH format.`,
		Exposed:        true,
		Name:           "certificate",
		ReadOnly:       true,
		Type:           "string",
	},
	"Extensions": {
		AllowedChoices: []string{},
		ConvertedName:  "Extensions",
		Description:    `List of extensions to set in the SSH certificate.`,
		Exposed:        true,
		Name:           "extensions",
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Options": {
		AllowedChoices: []string{},
		ConvertedName:  "Options",
		Description:    `List of options to set in the SSH certificate.`,
		Exposed:        true,
		Name:           "options",
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Principals": {
		AllowedChoices: []string{},
		ConvertedName:  "Principals",
		Description:    `List of principals to set in the SSH certificate.`,
		Exposed:        true,
		Name:           "principals",
		SubType:        "string",
		Type:           "list",
	},
	"PublicKey": {
		AllowedChoices: []string{},
		ConvertedName:  "PublicKey",
		Description:    `Contains the public key to sign in OpenSSH format.`,
		Exposed:        true,
		Name:           "publicKey",
		Required:       true,
		Type:           "string",
	},
	"SignerID": {
		AllowedChoices: []string{},
		ConvertedName:  "SignerID",
		Description:    `The identifier of the CA to use to sign the certificate.`,
		Exposed:        true,
		Name:           "signerID",
		Required:       true,
		Type:           "string",
	},
	"Type": {
		AllowedChoices: []string{"User", "Host"},
		ConvertedName:  "Type",
		DefaultValue:   SSHCertificateTypeUser,
		Description:    `Type of SSH certificate.`,
		Exposed:        true,
		Name:           "type",
		Type:           "enum",
	},
	"Validity": {
		AllowedChoices: []string{},
		ConvertedName:  "Validity",
		DefaultValue:   "1h",
		Description:    `Set the validity of the SSH certificate.`,
		Exposed:        true,
		Name:           "validity",
		Type:           "string",
	},
}

// SSHCertificateLowerCaseAttributesMap represents the map of attribute for SSHCertificate.
var SSHCertificateLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"certificate": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Certificate",
		Description:    `Contains the signed SSH certificate in OpenSSH format.`,
		Exposed:        true,
		Name:           "certificate",
		ReadOnly:       true,
		Type:           "string",
	},
	"extensions": {
		AllowedChoices: []string{},
		ConvertedName:  "Extensions",
		Description:    `List of extensions to set in the SSH certificate.`,
		Exposed:        true,
		Name:           "extensions",
		SubType:        "map[string]string",
		Type:           "external",
	},
	"options": {
		AllowedChoices: []string{},
		ConvertedName:  "Options",
		Description:    `List of options to set in the SSH certificate.`,
		Exposed:        true,
		Name:           "options",
		SubType:        "map[string]string",
		Type:           "external",
	},
	"principals": {
		AllowedChoices: []string{},
		ConvertedName:  "Principals",
		Description:    `List of principals to set in the SSH certificate.`,
		Exposed:        true,
		Name:           "principals",
		SubType:        "string",
		Type:           "list",
	},
	"publickey": {
		AllowedChoices: []string{},
		ConvertedName:  "PublicKey",
		Description:    `Contains the public key to sign in OpenSSH format.`,
		Exposed:        true,
		Name:           "publicKey",
		Required:       true,
		Type:           "string",
	},
	"signerid": {
		AllowedChoices: []string{},
		ConvertedName:  "SignerID",
		Description:    `The identifier of the CA to use to sign the certificate.`,
		Exposed:        true,
		Name:           "signerID",
		Required:       true,
		Type:           "string",
	},
	"type": {
		AllowedChoices: []string{"User", "Host"},
		ConvertedName:  "Type",
		DefaultValue:   SSHCertificateTypeUser,
		Description:    `Type of SSH certificate.`,
		Exposed:        true,
		Name:           "type",
		Type:           "enum",
	},
	"validity": {
		AllowedChoices: []string{},
		ConvertedName:  "Validity",
		DefaultValue:   "1h",
		Description:    `Set the validity of the SSH certificate.`,
		Exposed:        true,
		Name:           "validity",
		Type:           "string",
	},
}

// SparseSSHCertificatesList represents a list of SparseSSHCertificates
type SparseSSHCertificatesList []*SparseSSHCertificate

// Identity returns the identity of the objects in the list.
func (o SparseSSHCertificatesList) Identity() elemental.Identity {

	return SSHCertificateIdentity
}

// Copy returns a pointer to a copy the SparseSSHCertificatesList.
func (o SparseSSHCertificatesList) Copy() elemental.Identifiables {

	copy := append(SparseSSHCertificatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseSSHCertificatesList.
func (o SparseSSHCertificatesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseSSHCertificatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseSSHCertificate))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseSSHCertificatesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseSSHCertificatesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseSSHCertificatesList converted to SSHCertificatesList.
func (o SparseSSHCertificatesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseSSHCertificatesList) Version() int {

	return 1
}

// SparseSSHCertificate represents the sparse version of a sshcertificate.
type SparseSSHCertificate struct {
	// Contains the signed SSH certificate in OpenSSH format.
	Certificate *string `json:"certificate,omitempty" msgpack:"certificate,omitempty" bson:"-" mapstructure:"certificate,omitempty"`

	// List of extensions to set in the SSH certificate.
	Extensions *map[string]string `json:"extensions,omitempty" msgpack:"extensions,omitempty" bson:"-" mapstructure:"extensions,omitempty"`

	// List of options to set in the SSH certificate.
	Options *map[string]string `json:"options,omitempty" msgpack:"options,omitempty" bson:"-" mapstructure:"options,omitempty"`

	// List of principals to set in the SSH certificate.
	Principals *[]string `json:"principals,omitempty" msgpack:"principals,omitempty" bson:"-" mapstructure:"principals,omitempty"`

	// Contains the public key to sign in OpenSSH format.
	PublicKey *string `json:"publicKey,omitempty" msgpack:"publicKey,omitempty" bson:"-" mapstructure:"publicKey,omitempty"`

	// The identifier of the CA to use to sign the certificate.
	SignerID *string `json:"signerID,omitempty" msgpack:"signerID,omitempty" bson:"-" mapstructure:"signerID,omitempty"`

	// Type of SSH certificate.
	Type *SSHCertificateTypeValue `json:"type,omitempty" msgpack:"type,omitempty" bson:"-" mapstructure:"type,omitempty"`

	// Set the validity of the SSH certificate.
	Validity *string `json:"validity,omitempty" msgpack:"validity,omitempty" bson:"-" mapstructure:"validity,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseSSHCertificate returns a new  SparseSSHCertificate.
func NewSparseSSHCertificate() *SparseSSHCertificate {
	return &SparseSSHCertificate{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseSSHCertificate) Identity() elemental.Identity {

	return SSHCertificateIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseSSHCertificate) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseSSHCertificate) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseSSHCertificate) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseSSHCertificate{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseSSHCertificate) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseSSHCertificate{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseSSHCertificate) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseSSHCertificate) ToPlain() elemental.PlainIdentifiable {

	out := NewSSHCertificate()
	if o.Certificate != nil {
		out.Certificate = *o.Certificate
	}
	if o.Extensions != nil {
		out.Extensions = *o.Extensions
	}
	if o.Options != nil {
		out.Options = *o.Options
	}
	if o.Principals != nil {
		out.Principals = *o.Principals
	}
	if o.PublicKey != nil {
		out.PublicKey = *o.PublicKey
	}
	if o.SignerID != nil {
		out.SignerID = *o.SignerID
	}
	if o.Type != nil {
		out.Type = *o.Type
	}
	if o.Validity != nil {
		out.Validity = *o.Validity
	}

	return out
}

// DeepCopy returns a deep copy if the SparseSSHCertificate.
func (o *SparseSSHCertificate) DeepCopy() *SparseSSHCertificate {

	if o == nil {
		return nil
	}

	out := &SparseSSHCertificate{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseSSHCertificate.
func (o *SparseSSHCertificate) DeepCopyInto(out *SparseSSHCertificate) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseSSHCertificate: %s", err))
	}

	*out = *target.(*SparseSSHCertificate)
}

type mongoAttributesSSHCertificate struct {
}
type mongoAttributesSparseSSHCertificate struct {
}
