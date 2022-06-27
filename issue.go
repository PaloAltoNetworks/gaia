package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
	"go.aporeto.io/gaia/types"
)

// IssueRealmValue represents the possible values for attribute "realm".
type IssueRealmValue string

const (
	// IssueRealmAWSSecurityToken represents the value AWSSecurityToken.
	IssueRealmAWSSecurityToken IssueRealmValue = "AWSSecurityToken"

	// IssueRealmAporetoIdentityToken represents the value AporetoIdentityToken.
	IssueRealmAporetoIdentityToken IssueRealmValue = "AporetoIdentityToken"

	// IssueRealmAzureIdentityToken represents the value AzureIdentityToken.
	IssueRealmAzureIdentityToken IssueRealmValue = "AzureIdentityToken"

	// IssueRealmCertificate represents the value Certificate.
	IssueRealmCertificate IssueRealmValue = "Certificate"

	// IssueRealmGCPIdentityToken represents the value GCPIdentityToken.
	IssueRealmGCPIdentityToken IssueRealmValue = "GCPIdentityToken"

	// IssueRealmGoogle represents the value Google.
	IssueRealmGoogle IssueRealmValue = "Google"

	// IssueRealmLDAP represents the value LDAP.
	IssueRealmLDAP IssueRealmValue = "LDAP"

	// IssueRealmOIDC represents the value OIDC.
	IssueRealmOIDC IssueRealmValue = "OIDC"

	// IssueRealmPCIdentityToken represents the value PCIdentityToken.
	IssueRealmPCIdentityToken IssueRealmValue = "PCIdentityToken"

	// IssueRealmSAML represents the value SAML.
	IssueRealmSAML IssueRealmValue = "SAML"

	// IssueRealmVince represents the value Vince.
	IssueRealmVince IssueRealmValue = "Vince"
)

// IssueIdentity represents the Identity of the object.
var IssueIdentity = elemental.Identity{
	Name:     "issue",
	Category: "issue",
	Package:  "midgard",
	Private:  false,
}

// IssuesList represents a list of Issues
type IssuesList []*Issue

// Identity returns the identity of the objects in the list.
func (o IssuesList) Identity() elemental.Identity {

	return IssueIdentity
}

// Copy returns a pointer to a copy the IssuesList.
func (o IssuesList) Copy() elemental.Identifiables {

	copy := append(IssuesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the IssuesList.
func (o IssuesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(IssuesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Issue))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o IssuesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o IssuesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the IssuesList converted to SparseIssuesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o IssuesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseIssuesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseIssue)
	}

	return out
}

// Version returns the version of the content.
func (o IssuesList) Version() int {

	return 1
}

// Issue represents the model of a issue
type Issue struct {
	// If given, the issued token will only be valid for the specified namespace.
	// Refer to [JSON Web Token (JWT)RFC
	// 7519](https://tools.ietf.org/html/rfc7519#section-4.1.3).
	// for further information.
	Audience string `json:"audience" msgpack:"audience" bson:"-" mapstructure:"audience,omitempty"`

	// The claims in the token. It is only set is the parameter `asCookie` is given.
	Claims *types.MidgardClaims `json:"claims,omitempty" msgpack:"claims,omitempty" bson:"-" mapstructure:"claims,omitempty"`

	// Contains additional data. The value depends on the issuer type.
	Data string `json:"data" msgpack:"data" bson:"-" mapstructure:"data,omitempty"`

	// Contains various additional information. Meaning depends on the `realm`.
	Metadata map[string]interface{} `json:"metadata" msgpack:"metadata" bson:"-" mapstructure:"metadata,omitempty"`

	// Opaque data that will be included in the issued token.
	Opaque map[string]string `json:"opaque" msgpack:"opaque" bson:"-" mapstructure:"opaque,omitempty"`

	// Restricts the number of times the issued token can be used.
	Quota int `json:"quota" msgpack:"quota" bson:"-" mapstructure:"quota,omitempty"`

	// The authentication realm. This will define how to verify
	// credentials from internal or external source of authentication.
	Realm IssueRealmValue `json:"realm" msgpack:"realm" bson:"-" mapstructure:"realm,omitempty"`

	// Restricts the namespace where the token can be used.
	//
	// For instance, if you have have access to `/namespace` and below, you can
	// tell the policy engine that it should restrict further more to
	// `/namespace/child`.
	//
	// Restricting to a namespace you don't have initially access according to the
	// policy engine has no effect and may end up making the token unusable.
	RestrictedNamespace string `json:"restrictedNamespace" msgpack:"restrictedNamespace" bson:"-" mapstructure:"restrictedNamespace,omitempty"`

	// Restricts the networks from where the token can be used. This will reduce the
	// existing set of authorized networks that normally apply to the token according
	// to the policy engine.
	//
	// For instance, If you have authorized access from `0.0.0.0/0` (by default) or
	// from
	// `10.0.0.0/8`, you can ask for a token that will only be valid if used from
	// `10.1.0.0/16`.
	//
	// Restricting to a network that is not initially authorized by the policy
	// engine has no effect and may end up making the token unusable.
	RestrictedNetworks []string `json:"restrictedNetworks" msgpack:"restrictedNetworks" bson:"-" mapstructure:"restrictedNetworks,omitempty"`

	// Restricts the permissions of token. This will reduce the existing permissions
	// that normally apply to the token according to the policy engine.
	//
	// For instance, if you have administrative role, you can ask for a token that will
	// tell the policy engine to reduce the permission it would have granted to what is
	// given defined in the token.
	//
	// Restricting to some permissions you don't initially have according to the policy
	// engine has no effect and may end up making the token unusable.
	RestrictedPermissions []string `json:"restrictedPermissions" msgpack:"restrictedPermissions" bson:"-" mapstructure:"restrictedPermissions,omitempty"`

	// The token to use for the registration.
	Token string `json:"token" msgpack:"token" bson:"-" mapstructure:"token,omitempty"`

	// Configures the maximum length of validity for a token, using
	// [Golang duration syntax](https://golang.org/pkg/time/#example_Duration). If it
	// is
	// bigger than the configured max validity, it will be capped. Default: `24h`.
	Validity string `json:"validity" msgpack:"validity" bson:"-" mapstructure:"validity,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewIssue returns a new *Issue
func NewIssue() *Issue {

	return &Issue{
		ModelVersion:          1,
		Claims:                types.NewMidgardClaims(),
		Metadata:              map[string]interface{}{},
		Opaque:                map[string]string{},
		RestrictedNetworks:    []string{},
		RestrictedPermissions: []string{},
		Validity:              "24h",
	}
}

// Identity returns the Identity of the object.
func (o *Issue) Identity() elemental.Identity {

	return IssueIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Issue) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Issue) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Issue) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesIssue{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Issue) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesIssue{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Issue) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Issue) BleveType() string {

	return "issue"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Issue) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Issue) Doc() string {

	return `Issues a new Microsegmentation token according to given data.`
}

func (o *Issue) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Issue) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseIssue{
			Audience:              &o.Audience,
			Claims:                o.Claims,
			Data:                  &o.Data,
			Metadata:              &o.Metadata,
			Opaque:                &o.Opaque,
			Quota:                 &o.Quota,
			Realm:                 &o.Realm,
			RestrictedNamespace:   &o.RestrictedNamespace,
			RestrictedNetworks:    &o.RestrictedNetworks,
			RestrictedPermissions: &o.RestrictedPermissions,
			Token:                 &o.Token,
			Validity:              &o.Validity,
		}
	}

	sp := &SparseIssue{}
	for _, f := range fields {
		switch f {
		case "audience":
			sp.Audience = &(o.Audience)
		case "claims":
			sp.Claims = o.Claims
		case "data":
			sp.Data = &(o.Data)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "opaque":
			sp.Opaque = &(o.Opaque)
		case "quota":
			sp.Quota = &(o.Quota)
		case "realm":
			sp.Realm = &(o.Realm)
		case "restrictedNamespace":
			sp.RestrictedNamespace = &(o.RestrictedNamespace)
		case "restrictedNetworks":
			sp.RestrictedNetworks = &(o.RestrictedNetworks)
		case "restrictedPermissions":
			sp.RestrictedPermissions = &(o.RestrictedPermissions)
		case "token":
			sp.Token = &(o.Token)
		case "validity":
			sp.Validity = &(o.Validity)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseIssue to the object.
func (o *Issue) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseIssue)
	if so.Audience != nil {
		o.Audience = *so.Audience
	}
	if so.Claims != nil {
		o.Claims = so.Claims
	}
	if so.Data != nil {
		o.Data = *so.Data
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
	}
	if so.Opaque != nil {
		o.Opaque = *so.Opaque
	}
	if so.Quota != nil {
		o.Quota = *so.Quota
	}
	if so.Realm != nil {
		o.Realm = *so.Realm
	}
	if so.RestrictedNamespace != nil {
		o.RestrictedNamespace = *so.RestrictedNamespace
	}
	if so.RestrictedNetworks != nil {
		o.RestrictedNetworks = *so.RestrictedNetworks
	}
	if so.RestrictedPermissions != nil {
		o.RestrictedPermissions = *so.RestrictedPermissions
	}
	if so.Token != nil {
		o.Token = *so.Token
	}
	if so.Validity != nil {
		o.Validity = *so.Validity
	}
}

// DeepCopy returns a deep copy if the Issue.
func (o *Issue) DeepCopy() *Issue {

	if o == nil {
		return nil
	}

	out := &Issue{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Issue.
func (o *Issue) DeepCopyInto(out *Issue) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Issue: %s", err))
	}

	*out = *target.(*Issue)
}

// Validate valides the current information stored into the structure.
func (o *Issue) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateAudience("audience", o.Audience); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("realm", string(o.Realm)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("realm", string(o.Realm), []string{"AWSSecurityToken", "Certificate", "Google", "LDAP", "Vince", "GCPIdentityToken", "AzureIdentityToken", "OIDC", "SAML", "AporetoIdentityToken", "PCIdentityToken"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateOptionalCIDRList("restrictedNetworks", o.RestrictedNetworks); err != nil {
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
func (*Issue) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := IssueAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return IssueLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Issue) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return IssueAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Issue) ValueForAttribute(name string) interface{} {

	switch name {
	case "audience":
		return o.Audience
	case "claims":
		return o.Claims
	case "data":
		return o.Data
	case "metadata":
		return o.Metadata
	case "opaque":
		return o.Opaque
	case "quota":
		return o.Quota
	case "realm":
		return o.Realm
	case "restrictedNamespace":
		return o.RestrictedNamespace
	case "restrictedNetworks":
		return o.RestrictedNetworks
	case "restrictedPermissions":
		return o.RestrictedPermissions
	case "token":
		return o.Token
	case "validity":
		return o.Validity
	}

	return nil
}

// IssueAttributesMap represents the map of attribute for Issue.
var IssueAttributesMap = map[string]elemental.AttributeSpecification{
	"Audience": {
		AllowedChoices: []string{},
		ConvertedName:  "Audience",
		Description: `If given, the issued token will only be valid for the specified namespace.
Refer to [JSON Web Token (JWT)RFC
7519](https://tools.ietf.org/html/rfc7519#section-4.1.3).
for further information.`,
		Exposed: true,
		Name:    "audience",
		Type:    "string",
	},
	"Claims": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Claims",
		Description:    `The claims in the token. It is only set is the parameter ` + "`" + `asCookie` + "`" + ` is given.`,
		Exposed:        true,
		Name:           "claims",
		ReadOnly:       true,
		SubType:        "_claims",
		Type:           "external",
	},
	"Data": {
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Deprecated:     true,
		Description:    `Contains additional data. The value depends on the issuer type.`,
		Exposed:        true,
		Name:           "data",
		Orderable:      true,
		Type:           "string",
	},
	"Metadata": {
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		Description:    `Contains various additional information. Meaning depends on the ` + "`" + `realm` + "`" + `.`,
		Exposed:        true,
		Name:           "metadata",
		Orderable:      true,
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"Opaque": {
		AllowedChoices: []string{},
		ConvertedName:  "Opaque",
		Description:    `Opaque data that will be included in the issued token.`,
		Exposed:        true,
		Name:           "opaque",
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Quota": {
		AllowedChoices: []string{},
		ConvertedName:  "Quota",
		Description:    `Restricts the number of times the issued token can be used.`,
		Exposed:        true,
		Name:           "quota",
		Type:           "integer",
	},
	"Realm": {
		AllowedChoices: []string{"AWSSecurityToken", "Certificate", "Google", "LDAP", "Vince", "GCPIdentityToken", "AzureIdentityToken", "OIDC", "SAML", "AporetoIdentityToken", "PCIdentityToken"},
		ConvertedName:  "Realm",
		Description: `The authentication realm. This will define how to verify
credentials from internal or external source of authentication.`,
		Exposed:  true,
		Name:     "realm",
		Required: true,
		Type:     "enum",
	},
	"RestrictedNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "RestrictedNamespace",
		Description: `Restricts the namespace where the token can be used.

For instance, if you have have access to ` + "`" + `/namespace` + "`" + ` and below, you can
tell the policy engine that it should restrict further more to
` + "`" + `/namespace/child` + "`" + `.

Restricting to a namespace you don't have initially access according to the
policy engine has no effect and may end up making the token unusable.`,
		Exposed: true,
		Name:    "restrictedNamespace",
		Type:    "string",
	},
	"RestrictedNetworks": {
		AllowedChoices: []string{},
		ConvertedName:  "RestrictedNetworks",
		Description: `Restricts the networks from where the token can be used. This will reduce the
existing set of authorized networks that normally apply to the token according
to the policy engine.

For instance, If you have authorized access from ` + "`" + `0.0.0.0/0` + "`" + ` (by default) or
from
` + "`" + `10.0.0.0/8` + "`" + `, you can ask for a token that will only be valid if used from
` + "`" + `10.1.0.0/16` + "`" + `.

Restricting to a network that is not initially authorized by the policy
engine has no effect and may end up making the token unusable.`,
		Exposed: true,
		Name:    "restrictedNetworks",
		SubType: "string",
		Type:    "list",
	},
	"RestrictedPermissions": {
		AllowedChoices: []string{},
		ConvertedName:  "RestrictedPermissions",
		Description: `Restricts the permissions of token. This will reduce the existing permissions
that normally apply to the token according to the policy engine.

For instance, if you have administrative role, you can ask for a token that will
tell the policy engine to reduce the permission it would have granted to what is
given defined in the token.

Restricting to some permissions you don't initially have according to the policy
engine has no effect and may end up making the token unusable.`,
		Exposed: true,
		Name:    "restrictedPermissions",
		SubType: "string",
		Type:    "list",
	},
	"Token": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Token",
		Description:    `The token to use for the registration.`,
		Exposed:        true,
		Name:           "token",
		ReadOnly:       true,
		Type:           "string",
	},
	"Validity": {
		AllowedChoices: []string{},
		ConvertedName:  "Validity",
		DefaultValue:   "24h",
		Description: `Configures the maximum length of validity for a token, using
[Golang duration syntax](https://golang.org/pkg/time/#example_Duration). If it
is
bigger than the configured max validity, it will be capped. Default: ` + "`" + `24h` + "`" + `.`,
		Exposed:   true,
		Name:      "validity",
		Orderable: true,
		Type:      "string",
	},
}

// IssueLowerCaseAttributesMap represents the map of attribute for Issue.
var IssueLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"audience": {
		AllowedChoices: []string{},
		ConvertedName:  "Audience",
		Description: `If given, the issued token will only be valid for the specified namespace.
Refer to [JSON Web Token (JWT)RFC
7519](https://tools.ietf.org/html/rfc7519#section-4.1.3).
for further information.`,
		Exposed: true,
		Name:    "audience",
		Type:    "string",
	},
	"claims": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Claims",
		Description:    `The claims in the token. It is only set is the parameter ` + "`" + `asCookie` + "`" + ` is given.`,
		Exposed:        true,
		Name:           "claims",
		ReadOnly:       true,
		SubType:        "_claims",
		Type:           "external",
	},
	"data": {
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Deprecated:     true,
		Description:    `Contains additional data. The value depends on the issuer type.`,
		Exposed:        true,
		Name:           "data",
		Orderable:      true,
		Type:           "string",
	},
	"metadata": {
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		Description:    `Contains various additional information. Meaning depends on the ` + "`" + `realm` + "`" + `.`,
		Exposed:        true,
		Name:           "metadata",
		Orderable:      true,
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"opaque": {
		AllowedChoices: []string{},
		ConvertedName:  "Opaque",
		Description:    `Opaque data that will be included in the issued token.`,
		Exposed:        true,
		Name:           "opaque",
		SubType:        "map[string]string",
		Type:           "external",
	},
	"quota": {
		AllowedChoices: []string{},
		ConvertedName:  "Quota",
		Description:    `Restricts the number of times the issued token can be used.`,
		Exposed:        true,
		Name:           "quota",
		Type:           "integer",
	},
	"realm": {
		AllowedChoices: []string{"AWSSecurityToken", "Certificate", "Google", "LDAP", "Vince", "GCPIdentityToken", "AzureIdentityToken", "OIDC", "SAML", "AporetoIdentityToken", "PCIdentityToken"},
		ConvertedName:  "Realm",
		Description: `The authentication realm. This will define how to verify
credentials from internal or external source of authentication.`,
		Exposed:  true,
		Name:     "realm",
		Required: true,
		Type:     "enum",
	},
	"restrictednamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "RestrictedNamespace",
		Description: `Restricts the namespace where the token can be used.

For instance, if you have have access to ` + "`" + `/namespace` + "`" + ` and below, you can
tell the policy engine that it should restrict further more to
` + "`" + `/namespace/child` + "`" + `.

Restricting to a namespace you don't have initially access according to the
policy engine has no effect and may end up making the token unusable.`,
		Exposed: true,
		Name:    "restrictedNamespace",
		Type:    "string",
	},
	"restrictednetworks": {
		AllowedChoices: []string{},
		ConvertedName:  "RestrictedNetworks",
		Description: `Restricts the networks from where the token can be used. This will reduce the
existing set of authorized networks that normally apply to the token according
to the policy engine.

For instance, If you have authorized access from ` + "`" + `0.0.0.0/0` + "`" + ` (by default) or
from
` + "`" + `10.0.0.0/8` + "`" + `, you can ask for a token that will only be valid if used from
` + "`" + `10.1.0.0/16` + "`" + `.

Restricting to a network that is not initially authorized by the policy
engine has no effect and may end up making the token unusable.`,
		Exposed: true,
		Name:    "restrictedNetworks",
		SubType: "string",
		Type:    "list",
	},
	"restrictedpermissions": {
		AllowedChoices: []string{},
		ConvertedName:  "RestrictedPermissions",
		Description: `Restricts the permissions of token. This will reduce the existing permissions
that normally apply to the token according to the policy engine.

For instance, if you have administrative role, you can ask for a token that will
tell the policy engine to reduce the permission it would have granted to what is
given defined in the token.

Restricting to some permissions you don't initially have according to the policy
engine has no effect and may end up making the token unusable.`,
		Exposed: true,
		Name:    "restrictedPermissions",
		SubType: "string",
		Type:    "list",
	},
	"token": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Token",
		Description:    `The token to use for the registration.`,
		Exposed:        true,
		Name:           "token",
		ReadOnly:       true,
		Type:           "string",
	},
	"validity": {
		AllowedChoices: []string{},
		ConvertedName:  "Validity",
		DefaultValue:   "24h",
		Description: `Configures the maximum length of validity for a token, using
[Golang duration syntax](https://golang.org/pkg/time/#example_Duration). If it
is
bigger than the configured max validity, it will be capped. Default: ` + "`" + `24h` + "`" + `.`,
		Exposed:   true,
		Name:      "validity",
		Orderable: true,
		Type:      "string",
	},
}

// SparseIssuesList represents a list of SparseIssues
type SparseIssuesList []*SparseIssue

// Identity returns the identity of the objects in the list.
func (o SparseIssuesList) Identity() elemental.Identity {

	return IssueIdentity
}

// Copy returns a pointer to a copy the SparseIssuesList.
func (o SparseIssuesList) Copy() elemental.Identifiables {

	copy := append(SparseIssuesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseIssuesList.
func (o SparseIssuesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseIssuesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseIssue))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseIssuesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseIssuesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseIssuesList converted to IssuesList.
func (o SparseIssuesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseIssuesList) Version() int {

	return 1
}

// SparseIssue represents the sparse version of a issue.
type SparseIssue struct {
	// If given, the issued token will only be valid for the specified namespace.
	// Refer to [JSON Web Token (JWT)RFC
	// 7519](https://tools.ietf.org/html/rfc7519#section-4.1.3).
	// for further information.
	Audience *string `json:"audience,omitempty" msgpack:"audience,omitempty" bson:"-" mapstructure:"audience,omitempty"`

	// The claims in the token. It is only set is the parameter `asCookie` is given.
	Claims *types.MidgardClaims `json:"claims,omitempty" msgpack:"claims,omitempty" bson:"-" mapstructure:"claims,omitempty"`

	// Contains additional data. The value depends on the issuer type.
	Data *string `json:"data,omitempty" msgpack:"data,omitempty" bson:"-" mapstructure:"data,omitempty"`

	// Contains various additional information. Meaning depends on the `realm`.
	Metadata *map[string]interface{} `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"-" mapstructure:"metadata,omitempty"`

	// Opaque data that will be included in the issued token.
	Opaque *map[string]string `json:"opaque,omitempty" msgpack:"opaque,omitempty" bson:"-" mapstructure:"opaque,omitempty"`

	// Restricts the number of times the issued token can be used.
	Quota *int `json:"quota,omitempty" msgpack:"quota,omitempty" bson:"-" mapstructure:"quota,omitempty"`

	// The authentication realm. This will define how to verify
	// credentials from internal or external source of authentication.
	Realm *IssueRealmValue `json:"realm,omitempty" msgpack:"realm,omitempty" bson:"-" mapstructure:"realm,omitempty"`

	// Restricts the namespace where the token can be used.
	//
	// For instance, if you have have access to `/namespace` and below, you can
	// tell the policy engine that it should restrict further more to
	// `/namespace/child`.
	//
	// Restricting to a namespace you don't have initially access according to the
	// policy engine has no effect and may end up making the token unusable.
	RestrictedNamespace *string `json:"restrictedNamespace,omitempty" msgpack:"restrictedNamespace,omitempty" bson:"-" mapstructure:"restrictedNamespace,omitempty"`

	// Restricts the networks from where the token can be used. This will reduce the
	// existing set of authorized networks that normally apply to the token according
	// to the policy engine.
	//
	// For instance, If you have authorized access from `0.0.0.0/0` (by default) or
	// from
	// `10.0.0.0/8`, you can ask for a token that will only be valid if used from
	// `10.1.0.0/16`.
	//
	// Restricting to a network that is not initially authorized by the policy
	// engine has no effect and may end up making the token unusable.
	RestrictedNetworks *[]string `json:"restrictedNetworks,omitempty" msgpack:"restrictedNetworks,omitempty" bson:"-" mapstructure:"restrictedNetworks,omitempty"`

	// Restricts the permissions of token. This will reduce the existing permissions
	// that normally apply to the token according to the policy engine.
	//
	// For instance, if you have administrative role, you can ask for a token that will
	// tell the policy engine to reduce the permission it would have granted to what is
	// given defined in the token.
	//
	// Restricting to some permissions you don't initially have according to the policy
	// engine has no effect and may end up making the token unusable.
	RestrictedPermissions *[]string `json:"restrictedPermissions,omitempty" msgpack:"restrictedPermissions,omitempty" bson:"-" mapstructure:"restrictedPermissions,omitempty"`

	// The token to use for the registration.
	Token *string `json:"token,omitempty" msgpack:"token,omitempty" bson:"-" mapstructure:"token,omitempty"`

	// Configures the maximum length of validity for a token, using
	// [Golang duration syntax](https://golang.org/pkg/time/#example_Duration). If it
	// is
	// bigger than the configured max validity, it will be capped. Default: `24h`.
	Validity *string `json:"validity,omitempty" msgpack:"validity,omitempty" bson:"-" mapstructure:"validity,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseIssue returns a new  SparseIssue.
func NewSparseIssue() *SparseIssue {
	return &SparseIssue{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseIssue) Identity() elemental.Identity {

	return IssueIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseIssue) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseIssue) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseIssue) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseIssue{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseIssue) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseIssue{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseIssue) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseIssue) ToPlain() elemental.PlainIdentifiable {

	out := NewIssue()
	if o.Audience != nil {
		out.Audience = *o.Audience
	}
	if o.Claims != nil {
		out.Claims = o.Claims
	}
	if o.Data != nil {
		out.Data = *o.Data
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Opaque != nil {
		out.Opaque = *o.Opaque
	}
	if o.Quota != nil {
		out.Quota = *o.Quota
	}
	if o.Realm != nil {
		out.Realm = *o.Realm
	}
	if o.RestrictedNamespace != nil {
		out.RestrictedNamespace = *o.RestrictedNamespace
	}
	if o.RestrictedNetworks != nil {
		out.RestrictedNetworks = *o.RestrictedNetworks
	}
	if o.RestrictedPermissions != nil {
		out.RestrictedPermissions = *o.RestrictedPermissions
	}
	if o.Token != nil {
		out.Token = *o.Token
	}
	if o.Validity != nil {
		out.Validity = *o.Validity
	}

	return out
}

// DeepCopy returns a deep copy if the SparseIssue.
func (o *SparseIssue) DeepCopy() *SparseIssue {

	if o == nil {
		return nil
	}

	out := &SparseIssue{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseIssue.
func (o *SparseIssue) DeepCopyInto(out *SparseIssue) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseIssue: %s", err))
	}

	*out = *target.(*SparseIssue)
}

type mongoAttributesIssue struct {
}
type mongoAttributesSparseIssue struct {
}
