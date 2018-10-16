package gaia

import (
	"fmt"
	"sync"

	"go.aporeto.io/elemental"
)

// TokenIdentity represents the Identity of the object.
var TokenIdentity = elemental.Identity{
	Name:     "token",
	Category: "tokens",
	Package:  "barret",
	Private:  true,
}

// TokensList represents a list of Tokens
type TokensList []*Token

// Identity returns the identity of the objects in the list.
func (o TokensList) Identity() elemental.Identity {

	return TokenIdentity
}

// Copy returns a pointer to a copy the TokensList.
func (o TokensList) Copy() elemental.Identifiables {

	copy := append(TokensList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TokensList.
func (o TokensList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(TokensList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Token))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TokensList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TokensList) DefaultOrder() []string {

	return []string{}
}

// ToFull returns the TokensList converted to SparseTokensList.
func (o TokensList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o TokensList) Version() int {

	return 1
}

// Token represents the model of a token
type Token struct {
	// Certificate contains the client certificate to use to create a token.
	Certificate string `json:"certificate" bson:"-" mapstructure:"certificate,omitempty"`

	// SigningKeyID holds the ID of the custom CA to use to sign the token.
	SigningKeyID string `json:"signingKeyID" bson:"signingkeyid" mapstructure:"signingKeyID,omitempty"`

	// Token contains the generated token.
	Token string `json:"token" bson:"-" mapstructure:"token,omitempty"`

	// Validity contains the token validity duration.
	Validity string `json:"validity" bson:"-" mapstructure:"validity,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewToken returns a new *Token
func NewToken() *Token {

	return &Token{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Token) Identity() elemental.Identity {

	return TokenIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Token) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Token) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *Token) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Token) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Token) Doc() string {
	return `This api issue signed token from the given certificate.`
}

func (o *Token) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
func (o *Token) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		return &SparseToken{
			Certificate:  &o.Certificate,
			SigningKeyID: &o.SigningKeyID,
			Token:        &o.Token,
			Validity:     &o.Validity,
		}
	}

	sp := &SparseToken{}
	for _, f := range fields {
		switch f {
		case "certificate":
			sp.Certificate = &(o.Certificate)
		case "signingKeyID":
			sp.SigningKeyID = &(o.SigningKeyID)
		case "token":
			sp.Token = &(o.Token)
		case "validity":
			sp.Validity = &(o.Validity)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseToken to the object.
func (o *Token) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseToken)
	if so.Certificate != nil {
		o.Certificate = *so.Certificate
	}
	if so.SigningKeyID != nil {
		o.SigningKeyID = *so.SigningKeyID
	}
	if so.Token != nil {
		o.Token = *so.Token
	}
	if so.Validity != nil {
		o.Validity = *so.Validity
	}
}

// Validate valides the current information stored into the structure.
func (o *Token) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("certificate", o.Certificate); err != nil {
		requiredErrors = append(requiredErrors, err)
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
func (*Token) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TokenAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TokenLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Token) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TokenAttributesMap
}

// TokenAttributesMap represents the map of attribute for Token.
var TokenAttributesMap = map[string]elemental.AttributeSpecification{
	"Certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Certificate",
		CreationOnly:   true,
		Description:    `Certificate contains the client certificate to use to create a token.`,
		Exposed:        true,
		Name:           "certificate",
		Required:       true,
		Type:           "string",
	},
	"SigningKeyID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SigningKeyID",
		Description:    `SigningKeyID holds the ID of the custom CA to use to sign the token.`,
		Exposed:        true,
		Name:           "signingKeyID",
		Stored:         true,
		Type:           "string",
	},
	"Token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Token",
		Description:    `Token contains the generated token.`,
		Exposed:        true,
		Name:           "token",
		ReadOnly:       true,
		Type:           "string",
	},
	"Validity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Validity",
		CreationOnly:   true,
		Description:    `Validity contains the token validity duration.`,
		Exposed:        true,
		Name:           "validity",
		Type:           "string",
	},
}

// TokenLowerCaseAttributesMap represents the map of attribute for Token.
var TokenLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Certificate",
		CreationOnly:   true,
		Description:    `Certificate contains the client certificate to use to create a token.`,
		Exposed:        true,
		Name:           "certificate",
		Required:       true,
		Type:           "string",
	},
	"signingkeyid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SigningKeyID",
		Description:    `SigningKeyID holds the ID of the custom CA to use to sign the token.`,
		Exposed:        true,
		Name:           "signingKeyID",
		Stored:         true,
		Type:           "string",
	},
	"token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Token",
		Description:    `Token contains the generated token.`,
		Exposed:        true,
		Name:           "token",
		ReadOnly:       true,
		Type:           "string",
	},
	"validity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Validity",
		CreationOnly:   true,
		Description:    `Validity contains the token validity duration.`,
		Exposed:        true,
		Name:           "validity",
		Type:           "string",
	},
}

// SparseTokensList represents a list of SparseTokens
type SparseTokensList []*SparseToken

// Identity returns the identity of the objects in the list.
func (o SparseTokensList) Identity() elemental.Identity {

	return TokenIdentity
}

// Copy returns a pointer to a copy the SparseTokensList.
func (o SparseTokensList) Copy() elemental.Identifiables {

	copy := append(SparseTokensList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseTokensList.
func (o SparseTokensList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseTokensList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseToken))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseTokensList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseTokensList) DefaultOrder() []string {

	return []string{}
}

// ToFull returns the SparseTokensList converted to TokensList.
func (o SparseTokensList) ToFull() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToFull()
	}

	return out
}

// Version returns the version of the content.
func (o SparseTokensList) Version() int {

	return 1
}

// SparseToken represents the sparse version of a token.
type SparseToken struct {
	// Certificate contains the client certificate to use to create a token.
	Certificate *string `json:"certificate,omitempty" bson:"-" mapstructure:"certificate,omitempty"`

	// SigningKeyID holds the ID of the custom CA to use to sign the token.
	SigningKeyID *string `json:"signingKeyID,omitempty" bson:"signingkeyid" mapstructure:"signingKeyID,omitempty"`

	// Token contains the generated token.
	Token *string `json:"token,omitempty" bson:"-" mapstructure:"token,omitempty"`

	// Validity contains the token validity duration.
	Validity *string `json:"validity,omitempty" bson:"-" mapstructure:"validity,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseToken returns a new  SparseToken.
func NewSparseToken() *SparseToken {
	return &SparseToken{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseToken) Identity() elemental.Identity {

	return TokenIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseToken) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseToken) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseToken) Version() int {

	return 1
}

// ToFull returns a full version of the sparse model.
func (o *SparseToken) ToFull() elemental.FullIdentifiable {

	out := NewToken()
	if o.Certificate != nil {
		out.Certificate = *o.Certificate
	}
	if o.SigningKeyID != nil {
		out.SigningKeyID = *o.SigningKeyID
	}
	if o.Token != nil {
		out.Token = *o.Token
	}
	if o.Validity != nil {
		out.Validity = *o.Validity
	}

	return out
}
