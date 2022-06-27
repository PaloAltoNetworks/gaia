package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// SSHIdentityIdentity represents the Identity of the object.
var SSHIdentityIdentity = elemental.Identity{
	Name:     "sshidentity",
	Category: "sshidentities",
	Package:  "cactuar",
	Private:  false,
}

// SSHIdentitiesList represents a list of SSHIdentities
type SSHIdentitiesList []*SSHIdentity

// Identity returns the identity of the objects in the list.
func (o SSHIdentitiesList) Identity() elemental.Identity {

	return SSHIdentityIdentity
}

// Copy returns a pointer to a copy the SSHIdentitiesList.
func (o SSHIdentitiesList) Copy() elemental.Identifiables {

	copy := append(SSHIdentitiesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SSHIdentitiesList.
func (o SSHIdentitiesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SSHIdentitiesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SSHIdentity))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SSHIdentitiesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SSHIdentitiesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the SSHIdentitiesList converted to SparseSSHIdentitiesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o SSHIdentitiesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseSSHIdentitiesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseSSHIdentity)
	}

	return out
}

// Version returns the version of the content.
func (o SSHIdentitiesList) Version() int {

	return 1
}

// SSHIdentity represents the model of a sshidentity
type SSHIdentity struct {
	// Contains the signed SSH certificate in OpenSSH format.
	Certificate string `json:"certificate" msgpack:"certificate" bson:"-" mapstructure:"certificate,omitempty"`

	// Contains the public key to sign in OpenSSH format. You can generate an SSH
	// public key with the standard `ssh-keygen` tool.
	PublicKey string `json:"publicKey" msgpack:"publicKey" bson:"-" mapstructure:"publicKey,omitempty"`

	// Define the targeted system account name.
	SystemAccount string `json:"systemAccount" msgpack:"systemAccount" bson:"-" mapstructure:"systemAccount,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSSHIdentity returns a new *SSHIdentity
func NewSSHIdentity() *SSHIdentity {

	return &SSHIdentity{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *SSHIdentity) Identity() elemental.Identity {

	return SSHIdentityIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SSHIdentity) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SSHIdentity) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SSHIdentity) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSSHIdentity{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SSHIdentity) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSSHIdentity{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SSHIdentity) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *SSHIdentity) BleveType() string {

	return "sshidentity"
}

// DefaultOrder returns the list of default ordering fields.
func (o *SSHIdentity) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *SSHIdentity) Doc() string {

	return `Returns an SSH certificate containing the bearer claims. This SSH certificate
can
be used to connect to a node where the enforcer is protecting SSH sessions.`
}

func (o *SSHIdentity) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *SSHIdentity) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseSSHIdentity{
			Certificate:   &o.Certificate,
			PublicKey:     &o.PublicKey,
			SystemAccount: &o.SystemAccount,
		}
	}

	sp := &SparseSSHIdentity{}
	for _, f := range fields {
		switch f {
		case "certificate":
			sp.Certificate = &(o.Certificate)
		case "publicKey":
			sp.PublicKey = &(o.PublicKey)
		case "systemAccount":
			sp.SystemAccount = &(o.SystemAccount)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseSSHIdentity to the object.
func (o *SSHIdentity) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseSSHIdentity)
	if so.Certificate != nil {
		o.Certificate = *so.Certificate
	}
	if so.PublicKey != nil {
		o.PublicKey = *so.PublicKey
	}
	if so.SystemAccount != nil {
		o.SystemAccount = *so.SystemAccount
	}
}

// DeepCopy returns a deep copy if the SSHIdentity.
func (o *SSHIdentity) DeepCopy() *SSHIdentity {

	if o == nil {
		return nil
	}

	out := &SSHIdentity{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SSHIdentity.
func (o *SSHIdentity) DeepCopyInto(out *SSHIdentity) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SSHIdentity: %s", err))
	}

	*out = *target.(*SSHIdentity)
}

// Validate valides the current information stored into the structure.
func (o *SSHIdentity) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("publicKey", o.PublicKey); err != nil {
		requiredErrors = requiredErrors.Append(err)
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
func (*SSHIdentity) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := SSHIdentityAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return SSHIdentityLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*SSHIdentity) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return SSHIdentityAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *SSHIdentity) ValueForAttribute(name string) interface{} {

	switch name {
	case "certificate":
		return o.Certificate
	case "publicKey":
		return o.PublicKey
	case "systemAccount":
		return o.SystemAccount
	}

	return nil
}

// SSHIdentityAttributesMap represents the map of attribute for SSHIdentity.
var SSHIdentityAttributesMap = map[string]elemental.AttributeSpecification{
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
	"PublicKey": {
		AllowedChoices: []string{},
		ConvertedName:  "PublicKey",
		Description: `Contains the public key to sign in OpenSSH format. You can generate an SSH
public key with the standard ` + "`" + `ssh-keygen` + "`" + ` tool.`,
		Exposed:  true,
		Name:     "publicKey",
		Required: true,
		Type:     "string",
	},
	"SystemAccount": {
		AllowedChoices: []string{},
		ConvertedName:  "SystemAccount",
		Description:    `Define the targeted system account name.`,
		Exposed:        true,
		Name:           "systemAccount",
		Type:           "string",
	},
}

// SSHIdentityLowerCaseAttributesMap represents the map of attribute for SSHIdentity.
var SSHIdentityLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"publickey": {
		AllowedChoices: []string{},
		ConvertedName:  "PublicKey",
		Description: `Contains the public key to sign in OpenSSH format. You can generate an SSH
public key with the standard ` + "`" + `ssh-keygen` + "`" + ` tool.`,
		Exposed:  true,
		Name:     "publicKey",
		Required: true,
		Type:     "string",
	},
	"systemaccount": {
		AllowedChoices: []string{},
		ConvertedName:  "SystemAccount",
		Description:    `Define the targeted system account name.`,
		Exposed:        true,
		Name:           "systemAccount",
		Type:           "string",
	},
}

// SparseSSHIdentitiesList represents a list of SparseSSHIdentities
type SparseSSHIdentitiesList []*SparseSSHIdentity

// Identity returns the identity of the objects in the list.
func (o SparseSSHIdentitiesList) Identity() elemental.Identity {

	return SSHIdentityIdentity
}

// Copy returns a pointer to a copy the SparseSSHIdentitiesList.
func (o SparseSSHIdentitiesList) Copy() elemental.Identifiables {

	copy := append(SparseSSHIdentitiesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseSSHIdentitiesList.
func (o SparseSSHIdentitiesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseSSHIdentitiesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseSSHIdentity))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseSSHIdentitiesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseSSHIdentitiesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseSSHIdentitiesList converted to SSHIdentitiesList.
func (o SparseSSHIdentitiesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseSSHIdentitiesList) Version() int {

	return 1
}

// SparseSSHIdentity represents the sparse version of a sshidentity.
type SparseSSHIdentity struct {
	// Contains the signed SSH certificate in OpenSSH format.
	Certificate *string `json:"certificate,omitempty" msgpack:"certificate,omitempty" bson:"-" mapstructure:"certificate,omitempty"`

	// Contains the public key to sign in OpenSSH format. You can generate an SSH
	// public key with the standard `ssh-keygen` tool.
	PublicKey *string `json:"publicKey,omitempty" msgpack:"publicKey,omitempty" bson:"-" mapstructure:"publicKey,omitempty"`

	// Define the targeted system account name.
	SystemAccount *string `json:"systemAccount,omitempty" msgpack:"systemAccount,omitempty" bson:"-" mapstructure:"systemAccount,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseSSHIdentity returns a new  SparseSSHIdentity.
func NewSparseSSHIdentity() *SparseSSHIdentity {
	return &SparseSSHIdentity{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseSSHIdentity) Identity() elemental.Identity {

	return SSHIdentityIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseSSHIdentity) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseSSHIdentity) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseSSHIdentity) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseSSHIdentity{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseSSHIdentity) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseSSHIdentity{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseSSHIdentity) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseSSHIdentity) ToPlain() elemental.PlainIdentifiable {

	out := NewSSHIdentity()
	if o.Certificate != nil {
		out.Certificate = *o.Certificate
	}
	if o.PublicKey != nil {
		out.PublicKey = *o.PublicKey
	}
	if o.SystemAccount != nil {
		out.SystemAccount = *o.SystemAccount
	}

	return out
}

// DeepCopy returns a deep copy if the SparseSSHIdentity.
func (o *SparseSSHIdentity) DeepCopy() *SparseSSHIdentity {

	if o == nil {
		return nil
	}

	out := &SparseSSHIdentity{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseSSHIdentity.
func (o *SparseSSHIdentity) DeepCopyInto(out *SparseSSHIdentity) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseSSHIdentity: %s", err))
	}

	*out = *target.(*SparseSSHIdentity)
}

type mongoAttributesSSHIdentity struct {
}
type mongoAttributesSparseSSHIdentity struct {
}
