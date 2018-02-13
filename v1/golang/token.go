package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// TokenIdentity represents the Identity of the object.
var TokenIdentity = elemental.Identity{
	Name:     "token",
	Category: "tokens",
	Private:  true,
}

// TokensList represents a list of Tokens
type TokensList []*Token

// ContentIdentity returns the identity of the objects in the list.
func (o TokensList) ContentIdentity() elemental.Identity {

	return TokenIdentity
}

// Copy returns a pointer to a copy the TokensList.
func (o TokensList) Copy() elemental.ContentIdentifiable {

	copy := append(TokensList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TokensList.
func (o TokensList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(TokensList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Token))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TokensList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TokensList) DefaultOrder() []string {

	return []string{}
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

	sync.Mutex
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

// Validate valides the current information stored into the structure.
func (o *Token) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("certificate", o.Certificate); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("certificate", o.Certificate); err != nil {
		errors = append(errors, err)
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
		Format:         "free",
		Name:           "certificate",
		Required:       true,
		Type:           "string",
	},
	"SigningKeyID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SigningKeyID",
		Description:    `SigningKeyID holds the ID of the custom CA to use to sign the token.`,
		Exposed:        true,
		Format:         "free",
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
		Format:         "free",
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
		Format:         "free",
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
		Format:         "free",
		Name:           "certificate",
		Required:       true,
		Type:           "string",
	},
	"signingkeyid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SigningKeyID",
		Description:    `SigningKeyID holds the ID of the custom CA to use to sign the token.`,
		Exposed:        true,
		Format:         "free",
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
		Format:         "free",
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
		Format:         "free",
		Name:           "validity",
		Type:           "string",
	},
}
