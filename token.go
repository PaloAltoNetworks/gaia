package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
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

// ToSparse returns the TokensList converted to SparseTokensList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o TokensList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseTokensList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseToken)
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
	Certificate string `json:"certificate" msgpack:"certificate" bson:"-" mapstructure:"certificate,omitempty"`

	// SigningKeyID holds the ID of the custom CA to use to sign the token.
	SigningKeyID string `json:"signingKeyID" msgpack:"signingKeyID" bson:"signingkeyid" mapstructure:"signingKeyID,omitempty"`

	// Tags includes a list of tags that must be added to the token.
	Tags []string `json:"tags" msgpack:"tags" bson:"-" mapstructure:"tags,omitempty"`

	// Token contains the generated token.
	Token string `json:"token" msgpack:"token" bson:"-" mapstructure:"token,omitempty"`

	// Validity contains the token validity duration.
	Validity string `json:"validity" msgpack:"validity" bson:"-" mapstructure:"validity,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewToken returns a new *Token
func NewToken() *Token {

	return &Token{
		ModelVersion: 1,
		Tags:         []string{},
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

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Token) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesToken{}

	s.SigningKeyID = o.SigningKeyID

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Token) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesToken{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.SigningKeyID = s.SigningKeyID

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Token) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Token) BleveType() string {

	return "token"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Token) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Token) Doc() string {

	return `Issues a signed token from the given certificate.`
}

func (o *Token) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Token) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseToken{
			Certificate:  &o.Certificate,
			SigningKeyID: &o.SigningKeyID,
			Tags:         &o.Tags,
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
		case "tags":
			sp.Tags = &(o.Tags)
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
	if so.Tags != nil {
		o.Tags = *so.Tags
	}
	if so.Token != nil {
		o.Token = *so.Token
	}
	if so.Validity != nil {
		o.Validity = *so.Validity
	}
}

// DeepCopy returns a deep copy if the Token.
func (o *Token) DeepCopy() *Token {

	if o == nil {
		return nil
	}

	out := &Token{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Token.
func (o *Token) DeepCopyInto(out *Token) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Token: %s", err))
	}

	*out = *target.(*Token)
}

// Validate valides the current information stored into the structure.
func (o *Token) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("certificate", o.Certificate); err != nil {
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

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Token) ValueForAttribute(name string) interface{} {

	switch name {
	case "certificate":
		return o.Certificate
	case "signingKeyID":
		return o.SigningKeyID
	case "tags":
		return o.Tags
	case "token":
		return o.Token
	case "validity":
		return o.Validity
	}

	return nil
}

// TokenAttributesMap represents the map of attribute for Token.
var TokenAttributesMap = map[string]elemental.AttributeSpecification{
	"Certificate": {
		AllowedChoices: []string{},
		ConvertedName:  "Certificate",
		CreationOnly:   true,
		Description:    `Certificate contains the client certificate to use to create a token.`,
		Exposed:        true,
		Name:           "certificate",
		Required:       true,
		Type:           "string",
	},
	"SigningKeyID": {
		AllowedChoices: []string{},
		ConvertedName:  "SigningKeyID",
		Description:    `SigningKeyID holds the ID of the custom CA to use to sign the token.`,
		Exposed:        true,
		Name:           "signingKeyID",
		Stored:         true,
		Type:           "string",
	},
	"Tags": {
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		CreationOnly:   true,
		Description:    `Tags includes a list of tags that must be added to the token.`,
		Exposed:        true,
		Name:           "tags",
		SubType:        "string",
		Type:           "list",
	},
	"Token": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Token",
		Description:    `Token contains the generated token.`,
		Exposed:        true,
		Name:           "token",
		ReadOnly:       true,
		Type:           "string",
	},
	"Validity": {
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
	"certificate": {
		AllowedChoices: []string{},
		ConvertedName:  "Certificate",
		CreationOnly:   true,
		Description:    `Certificate contains the client certificate to use to create a token.`,
		Exposed:        true,
		Name:           "certificate",
		Required:       true,
		Type:           "string",
	},
	"signingkeyid": {
		AllowedChoices: []string{},
		ConvertedName:  "SigningKeyID",
		Description:    `SigningKeyID holds the ID of the custom CA to use to sign the token.`,
		Exposed:        true,
		Name:           "signingKeyID",
		Stored:         true,
		Type:           "string",
	},
	"tags": {
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		CreationOnly:   true,
		Description:    `Tags includes a list of tags that must be added to the token.`,
		Exposed:        true,
		Name:           "tags",
		SubType:        "string",
		Type:           "list",
	},
	"token": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Token",
		Description:    `Token contains the generated token.`,
		Exposed:        true,
		Name:           "token",
		ReadOnly:       true,
		Type:           "string",
	},
	"validity": {
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

// ToPlain returns the SparseTokensList converted to TokensList.
func (o SparseTokensList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
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
	Certificate *string `json:"certificate,omitempty" msgpack:"certificate,omitempty" bson:"-" mapstructure:"certificate,omitempty"`

	// SigningKeyID holds the ID of the custom CA to use to sign the token.
	SigningKeyID *string `json:"signingKeyID,omitempty" msgpack:"signingKeyID,omitempty" bson:"signingkeyid,omitempty" mapstructure:"signingKeyID,omitempty"`

	// Tags includes a list of tags that must be added to the token.
	Tags *[]string `json:"tags,omitempty" msgpack:"tags,omitempty" bson:"-" mapstructure:"tags,omitempty"`

	// Token contains the generated token.
	Token *string `json:"token,omitempty" msgpack:"token,omitempty" bson:"-" mapstructure:"token,omitempty"`

	// Validity contains the token validity duration.
	Validity *string `json:"validity,omitempty" msgpack:"validity,omitempty" bson:"-" mapstructure:"validity,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
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

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseToken) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseToken{}

	if o.SigningKeyID != nil {
		s.SigningKeyID = o.SigningKeyID
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseToken) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseToken{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.SigningKeyID != nil {
		o.SigningKeyID = s.SigningKeyID
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseToken) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseToken) ToPlain() elemental.PlainIdentifiable {

	out := NewToken()
	if o.Certificate != nil {
		out.Certificate = *o.Certificate
	}
	if o.SigningKeyID != nil {
		out.SigningKeyID = *o.SigningKeyID
	}
	if o.Tags != nil {
		out.Tags = *o.Tags
	}
	if o.Token != nil {
		out.Token = *o.Token
	}
	if o.Validity != nil {
		out.Validity = *o.Validity
	}

	return out
}

// DeepCopy returns a deep copy if the SparseToken.
func (o *SparseToken) DeepCopy() *SparseToken {

	if o == nil {
		return nil
	}

	out := &SparseToken{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseToken.
func (o *SparseToken) DeepCopyInto(out *SparseToken) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseToken: %s", err))
	}

	*out = *target.(*SparseToken)
}

type mongoAttributesToken struct {
	SigningKeyID string `bson:"signingkeyid"`
}
type mongoAttributesSparseToken struct {
	SigningKeyID *string `bson:"signingkeyid,omitempty"`
}
