// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// EnforcerSecretIdentity represents the Identity of the object.
var EnforcerSecretIdentity = elemental.Identity{
	Name:     "enforcersecret",
	Category: "enforcersecrets",
	Package:  "squall",
	Private:  false,
}

// EnforcerSecretsList represents a list of EnforcerSecrets
type EnforcerSecretsList []*EnforcerSecret

// Identity returns the identity of the objects in the list.
func (o EnforcerSecretsList) Identity() elemental.Identity {

	return EnforcerSecretIdentity
}

// Copy returns a pointer to a copy the EnforcerSecretsList.
func (o EnforcerSecretsList) Copy() elemental.Identifiables {

	out := append(EnforcerSecretsList{}, o...)
	return &out
}

// Append appends the objects to the a new copy of the EnforcerSecretsList.
func (o EnforcerSecretsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(EnforcerSecretsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*EnforcerSecret))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o EnforcerSecretsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EnforcerSecretsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the EnforcerSecretsList converted to SparseEnforcerSecretsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o EnforcerSecretsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseEnforcerSecretsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseEnforcerSecret)
	}

	return out
}

// Version returns the version of the content.
func (o EnforcerSecretsList) Version() int {

	return 1
}

// EnforcerSecret represents the model of a enforcersecret
type EnforcerSecret struct {
	// Syslog public key in PEM format.
	SyslogCertificate string `json:"syslogCertificate" msgpack:"syslogCertificate" bson:"-" mapstructure:"syslogCertificate,omitempty"`

	// Syslog private key in PEM format.
	SyslogCertificateKey string `json:"syslogCertificateKey" msgpack:"syslogCertificateKey" bson:"-" mapstructure:"syslogCertificateKey,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewEnforcerSecret returns a new *EnforcerSecret
func NewEnforcerSecret() *EnforcerSecret {

	return &EnforcerSecret{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *EnforcerSecret) Identity() elemental.Identity {

	return EnforcerSecretIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EnforcerSecret) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EnforcerSecret) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *EnforcerSecret) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesEnforcerSecret{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *EnforcerSecret) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesEnforcerSecret{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *EnforcerSecret) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *EnforcerSecret) BleveType() string {

	return "enforcersecret"
}

// DefaultOrder returns the list of default ordering fields.
func (o *EnforcerSecret) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *EnforcerSecret) Doc() string {

	return `Manages private keys.`
}

func (o *EnforcerSecret) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *EnforcerSecret) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseEnforcerSecret{
			SyslogCertificate:    &o.SyslogCertificate,
			SyslogCertificateKey: &o.SyslogCertificateKey,
		}
	}

	sp := &SparseEnforcerSecret{}
	for _, f := range fields {
		switch f {
		case "syslogCertificate":
			sp.SyslogCertificate = &(o.SyslogCertificate)
		case "syslogCertificateKey":
			sp.SyslogCertificateKey = &(o.SyslogCertificateKey)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseEnforcerSecret to the object.
func (o *EnforcerSecret) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseEnforcerSecret)
	if so.SyslogCertificate != nil {
		o.SyslogCertificate = *so.SyslogCertificate
	}
	if so.SyslogCertificateKey != nil {
		o.SyslogCertificateKey = *so.SyslogCertificateKey
	}
}

// DeepCopy returns a deep copy if the EnforcerSecret.
func (o *EnforcerSecret) DeepCopy() *EnforcerSecret {

	if o == nil {
		return nil
	}

	out := &EnforcerSecret{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *EnforcerSecret.
func (o *EnforcerSecret) DeepCopyInto(out *EnforcerSecret) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy EnforcerSecret: %s", err))
	}

	*out = *target.(*EnforcerSecret)
}

// Validate valides the current information stored into the structure.
func (o *EnforcerSecret) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("syslogCertificate", o.SyslogCertificate); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := ValidatePEM("syslogCertificate", o.SyslogCertificate); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("syslogCertificateKey", o.SyslogCertificateKey); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := ValidatePEM("syslogCertificateKey", o.SyslogCertificateKey); err != nil {
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
func (*EnforcerSecret) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := EnforcerSecretAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return EnforcerSecretLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*EnforcerSecret) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return EnforcerSecretAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *EnforcerSecret) ValueForAttribute(name string) any {

	switch name {
	case "syslogCertificate":
		return o.SyslogCertificate
	case "syslogCertificateKey":
		return o.SyslogCertificateKey
	}

	return nil
}

// EnforcerSecretAttributesMap represents the map of attribute for EnforcerSecret.
var EnforcerSecretAttributesMap = map[string]elemental.AttributeSpecification{
	"SyslogCertificate": {
		AllowedChoices: []string{},
		ConvertedName:  "SyslogCertificate",
		Description:    `Syslog public key in PEM format.`,
		Exposed:        true,
		Name:           "syslogCertificate",
		Required:       true,
		Type:           "string",
	},
	"SyslogCertificateKey": {
		AllowedChoices: []string{},
		ConvertedName:  "SyslogCertificateKey",
		Description:    `Syslog private key in PEM format.`,
		Exposed:        true,
		Name:           "syslogCertificateKey",
		Required:       true,
		Type:           "string",
	},
}

// EnforcerSecretLowerCaseAttributesMap represents the map of attribute for EnforcerSecret.
var EnforcerSecretLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"syslogcertificate": {
		AllowedChoices: []string{},
		ConvertedName:  "SyslogCertificate",
		Description:    `Syslog public key in PEM format.`,
		Exposed:        true,
		Name:           "syslogCertificate",
		Required:       true,
		Type:           "string",
	},
	"syslogcertificatekey": {
		AllowedChoices: []string{},
		ConvertedName:  "SyslogCertificateKey",
		Description:    `Syslog private key in PEM format.`,
		Exposed:        true,
		Name:           "syslogCertificateKey",
		Required:       true,
		Type:           "string",
	},
}

// SparseEnforcerSecretsList represents a list of SparseEnforcerSecrets
type SparseEnforcerSecretsList []*SparseEnforcerSecret

// Identity returns the identity of the objects in the list.
func (o SparseEnforcerSecretsList) Identity() elemental.Identity {

	return EnforcerSecretIdentity
}

// Copy returns a pointer to a copy the SparseEnforcerSecretsList.
func (o SparseEnforcerSecretsList) Copy() elemental.Identifiables {

	copy := append(SparseEnforcerSecretsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseEnforcerSecretsList.
func (o SparseEnforcerSecretsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseEnforcerSecretsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseEnforcerSecret))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseEnforcerSecretsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseEnforcerSecretsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseEnforcerSecretsList converted to EnforcerSecretsList.
func (o SparseEnforcerSecretsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseEnforcerSecretsList) Version() int {

	return 1
}

// SparseEnforcerSecret represents the sparse version of a enforcersecret.
type SparseEnforcerSecret struct {
	// Syslog public key in PEM format.
	SyslogCertificate *string `json:"syslogCertificate,omitempty" msgpack:"syslogCertificate,omitempty" bson:"-" mapstructure:"syslogCertificate,omitempty"`

	// Syslog private key in PEM format.
	SyslogCertificateKey *string `json:"syslogCertificateKey,omitempty" msgpack:"syslogCertificateKey,omitempty" bson:"-" mapstructure:"syslogCertificateKey,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseEnforcerSecret returns a new  SparseEnforcerSecret.
func NewSparseEnforcerSecret() *SparseEnforcerSecret {
	return &SparseEnforcerSecret{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseEnforcerSecret) Identity() elemental.Identity {

	return EnforcerSecretIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseEnforcerSecret) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseEnforcerSecret) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseEnforcerSecret) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseEnforcerSecret{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseEnforcerSecret) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseEnforcerSecret{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseEnforcerSecret) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseEnforcerSecret) ToPlain() elemental.PlainIdentifiable {

	out := NewEnforcerSecret()
	if o.SyslogCertificate != nil {
		out.SyslogCertificate = *o.SyslogCertificate
	}
	if o.SyslogCertificateKey != nil {
		out.SyslogCertificateKey = *o.SyslogCertificateKey
	}

	return out
}

// DeepCopy returns a deep copy if the SparseEnforcerSecret.
func (o *SparseEnforcerSecret) DeepCopy() *SparseEnforcerSecret {

	if o == nil {
		return nil
	}

	out := &SparseEnforcerSecret{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseEnforcerSecret.
func (o *SparseEnforcerSecret) DeepCopyInto(out *SparseEnforcerSecret) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseEnforcerSecret: %s", err))
	}

	*out = *target.(*SparseEnforcerSecret)
}

type mongoAttributesEnforcerSecret struct {
}
type mongoAttributesSparseEnforcerSecret struct {
}
