package models

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/midgard-lib/claims"
)

// AuthIdentity represents the Identity of the object.
var AuthIdentity = elemental.Identity{
	Name:     "auth",
	Category: "auth",
}

// AuthsList represents a list of Auths
type AuthsList []*Auth

// ContentIdentity returns the identity of the objects in the list.
func (o AuthsList) ContentIdentity() elemental.Identity {

	return AuthIdentity
}

// Copy returns a pointer to a copy the AuthsList.
func (o AuthsList) Copy() elemental.ContentIdentifiable {

	copy := append(AuthsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AuthsList.
func (o AuthsList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(AuthsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Auth))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AuthsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AuthsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o AuthsList) Version() int {

	return 1
}

// Auth represents the model of a auth
type Auth struct {
	// Claims are the claims.
	Claims *claims.MidgardClaims `json:"claims" bson:"-" mapstructure:"claims,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewAuth returns a new *Auth
func NewAuth() *Auth {

	return &Auth{
		ModelVersion: 1,
		Claims:       claims.NewMidgardClaims(),
	}
}

// Identity returns the Identity of the object.
func (o *Auth) Identity() elemental.Identity {

	return AuthIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Auth) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Auth) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *Auth) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Auth) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Auth) Doc() string {
	return `This API verifies if the given token is valid or not.`
}

func (o *Auth) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Auth) Validate() error {

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
func (*Auth) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AuthAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AuthLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Auth) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AuthAttributesMap
}

// AuthAttributesMap represents the map of attribute for Auth.
var AuthAttributesMap = map[string]elemental.AttributeSpecification{
	"Claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Claims",
		Description:    `Claims are the claims.`,
		Exposed:        true,
		Name:           "claims",
		ReadOnly:       true,
		SubType:        "claims",
		Type:           "external",
	},
}

// AuthLowerCaseAttributesMap represents the map of attribute for Auth.
var AuthLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Claims",
		Description:    `Claims are the claims.`,
		Exposed:        true,
		Name:           "claims",
		ReadOnly:       true,
		SubType:        "claims",
		Type:           "external",
	},
}
