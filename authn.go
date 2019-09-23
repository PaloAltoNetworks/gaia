package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
	"go.aporeto.io/gaia/types"
)

// AuthnIdentity represents the Identity of the object.
var AuthnIdentity = elemental.Identity{
	Name:     "authn",
	Category: "authn",
	Package:  "midgard",
	Private:  false,
}

// AuthnsList represents a list of Authns
type AuthnsList []*Authn

// Identity returns the identity of the objects in the list.
func (o AuthnsList) Identity() elemental.Identity {

	return AuthnIdentity
}

// Copy returns a pointer to a copy the AuthnsList.
func (o AuthnsList) Copy() elemental.Identifiables {

	copy := append(AuthnsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AuthnsList.
func (o AuthnsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AuthnsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Authn))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AuthnsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AuthnsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the AuthnsList converted to SparseAuthnsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AuthnsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAuthnsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAuthn)
	}

	return out
}

// Version returns the version of the content.
func (o AuthnsList) Version() int {

	return 1
}

// Authn represents the model of a authn
type Authn struct {
	// The claims in the token.
	Claims *types.MidgardClaims `json:"claims" msgpack:"claims" bson:"-" mapstructure:"claims,omitempty"`

	// The token to verify. This is only used is a POST request is used.
	Token string `json:"token" msgpack:"token" bson:"-" mapstructure:"token,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAuthn returns a new *Authn
func NewAuthn() *Authn {

	return &Authn{
		ModelVersion: 1,
		Claims:       types.NewMidgardClaims(),
	}
}

// Identity returns the Identity of the object.
func (o *Authn) Identity() elemental.Identity {

	return AuthnIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Authn) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Authn) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Authn) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesAuthn{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Authn) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesAuthn{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Authn) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Authn) BleveType() string {

	return "authn"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Authn) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Authn) Doc() string {

	return `Verifies if the given token is valid or not. If it is valid it will
return the claims of the token.`
}

func (o *Authn) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Authn) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAuthn{
			Claims: o.Claims,
			Token:  &o.Token,
		}
	}

	sp := &SparseAuthn{}
	for _, f := range fields {
		switch f {
		case "claims":
			sp.Claims = o.Claims
		case "token":
			sp.Token = &(o.Token)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseAuthn to the object.
func (o *Authn) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAuthn)
	if so.Claims != nil {
		o.Claims = so.Claims
	}
	if so.Token != nil {
		o.Token = *so.Token
	}
}

// DeepCopy returns a deep copy if the Authn.
func (o *Authn) DeepCopy() *Authn {

	if o == nil {
		return nil
	}

	out := &Authn{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Authn.
func (o *Authn) DeepCopyInto(out *Authn) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Authn: %s", err))
	}

	*out = *target.(*Authn)
}

// Validate valides the current information stored into the structure.
func (o *Authn) Validate() error {

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
func (*Authn) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AuthnAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AuthnLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Authn) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AuthnAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Authn) ValueForAttribute(name string) interface{} {

	switch name {
	case "claims":
		return o.Claims
	case "token":
		return o.Token
	}

	return nil
}

// AuthnAttributesMap represents the map of attribute for Authn.
var AuthnAttributesMap = map[string]elemental.AttributeSpecification{
	"Claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Claims",
		Description:    `The claims in the token.`,
		Exposed:        true,
		Name:           "claims",
		ReadOnly:       true,
		SubType:        "_claims",
		Type:           "external",
	},
	"Token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Token",
		Description:    `The token to verify. This is only used is a POST request is used.`,
		Exposed:        true,
		Name:           "token",
		Type:           "string",
	},
}

// AuthnLowerCaseAttributesMap represents the map of attribute for Authn.
var AuthnLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Claims",
		Description:    `The claims in the token.`,
		Exposed:        true,
		Name:           "claims",
		ReadOnly:       true,
		SubType:        "_claims",
		Type:           "external",
	},
	"token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Token",
		Description:    `The token to verify. This is only used is a POST request is used.`,
		Exposed:        true,
		Name:           "token",
		Type:           "string",
	},
}

// SparseAuthnsList represents a list of SparseAuthns
type SparseAuthnsList []*SparseAuthn

// Identity returns the identity of the objects in the list.
func (o SparseAuthnsList) Identity() elemental.Identity {

	return AuthnIdentity
}

// Copy returns a pointer to a copy the SparseAuthnsList.
func (o SparseAuthnsList) Copy() elemental.Identifiables {

	copy := append(SparseAuthnsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAuthnsList.
func (o SparseAuthnsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAuthnsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAuthn))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAuthnsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAuthnsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseAuthnsList converted to AuthnsList.
func (o SparseAuthnsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAuthnsList) Version() int {

	return 1
}

// SparseAuthn represents the sparse version of a authn.
type SparseAuthn struct {
	// The claims in the token.
	Claims *types.MidgardClaims `json:"claims,omitempty" msgpack:"claims,omitempty" bson:"-" mapstructure:"claims,omitempty"`

	// The token to verify. This is only used is a POST request is used.
	Token *string `json:"token,omitempty" msgpack:"token,omitempty" bson:"-" mapstructure:"token,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAuthn returns a new  SparseAuthn.
func NewSparseAuthn() *SparseAuthn {
	return &SparseAuthn{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAuthn) Identity() elemental.Identity {

	return AuthnIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAuthn) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAuthn) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAuthn) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseAuthn{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAuthn) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseAuthn{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseAuthn) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAuthn) ToPlain() elemental.PlainIdentifiable {

	out := NewAuthn()
	if o.Claims != nil {
		out.Claims = o.Claims
	}
	if o.Token != nil {
		out.Token = *o.Token
	}

	return out
}

// DeepCopy returns a deep copy if the SparseAuthn.
func (o *SparseAuthn) DeepCopy() *SparseAuthn {

	if o == nil {
		return nil
	}

	out := &SparseAuthn{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAuthn.
func (o *SparseAuthn) DeepCopyInto(out *SparseAuthn) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAuthn: %s", err))
	}

	*out = *target.(*SparseAuthn)
}

type mongoAttributesAuthn struct {
}
type mongoAttributesSparseAuthn struct {
}
