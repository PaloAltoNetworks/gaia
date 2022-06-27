package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// OAUTHInfoIdentity represents the Identity of the object.
var OAUTHInfoIdentity = elemental.Identity{
	Name:     "oauthinfo",
	Category: "oauthinfo",
	Package:  "cactuar",
	Private:  false,
}

// OAUTHInfosList represents a list of OAUTHInfos
type OAUTHInfosList []*OAUTHInfo

// Identity returns the identity of the objects in the list.
func (o OAUTHInfosList) Identity() elemental.Identity {

	return OAUTHInfoIdentity
}

// Copy returns a pointer to a copy the OAUTHInfosList.
func (o OAUTHInfosList) Copy() elemental.Identifiables {

	copy := append(OAUTHInfosList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the OAUTHInfosList.
func (o OAUTHInfosList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(OAUTHInfosList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*OAUTHInfo))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o OAUTHInfosList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o OAUTHInfosList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the OAUTHInfosList converted to SparseOAUTHInfosList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o OAUTHInfosList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseOAUTHInfosList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseOAUTHInfo)
	}

	return out
}

// Version returns the version of the content.
func (o OAUTHInfosList) Version() int {

	return 1
}

// OAUTHInfo represents the model of a oauthinfo
type OAUTHInfo struct {
	// IDTokenSigningAlgValuesSupported is corresponding attribute of the OIDC
	// spec.
	IDTokenSigningAlgValuesSupported []string `json:"id_token_signing_alg_values_supported" msgpack:"id_token_signing_alg_values_supported" bson:"-" mapstructure:"id_token_signing_alg_values_supported,omitempty"`

	// JWKSURI is the URI that can be used to retrieve the public keys that will
	// verify a JWT.
	JWKSURI string `json:"jwks_uri" msgpack:"jwks_uri" bson:"-" mapstructure:"jwks_uri,omitempty"`

	// AuhorizationEndpoint is the authorization endpoint.
	AuhorizationEndpoint string `json:"auhorization_endpoint" msgpack:"auhorization_endpoint" bson:"-" mapstructure:"auhorization_endpoint,omitempty"`

	// ClaimsSupported is corresponding attribute of the OIDC spec.
	ClaimsSupported []string `json:"claims_supported" msgpack:"claims_supported" bson:"-" mapstructure:"claims_supported,omitempty"`

	// Issuer is the the URL pointing to the issuer of the token.
	Issuer string `json:"issuer" msgpack:"issuer" bson:"-" mapstructure:"issuer,omitempty"`

	// ResponseTypesSupported is corresponding attribute of the OIDC spec.
	ResponseTypesSupported []string `json:"response_types_supported" msgpack:"response_types_supported" bson:"-" mapstructure:"response_types_supported,omitempty"`

	// ScopesSupported is corresponding attribute of the OIDC spec.
	ScopesSupported []string `json:"scopes_supported" msgpack:"scopes_supported" bson:"-" mapstructure:"scopes_supported,omitempty"`

	// SubjectTypesSupported is corresponding attribute of the OIDC spec.
	SubjectTypesSupported []string `json:"subject_types_supported" msgpack:"subject_types_supported" bson:"-" mapstructure:"subject_types_supported,omitempty"`

	// TokenEndpointAuthMethodsSupported is corresponding attribute of the OIDC
	// spec.
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported" msgpack:"token_endpoint_auth_methods_supported" bson:"-" mapstructure:"token_endpoint_auth_methods_supported,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewOAUTHInfo returns a new *OAUTHInfo
func NewOAUTHInfo() *OAUTHInfo {

	return &OAUTHInfo{
		ModelVersion:                      1,
		IDTokenSigningAlgValuesSupported:  []string{},
		ClaimsSupported:                   []string{},
		ResponseTypesSupported:            []string{},
		ScopesSupported:                   []string{},
		SubjectTypesSupported:             []string{},
		TokenEndpointAuthMethodsSupported: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *OAUTHInfo) Identity() elemental.Identity {

	return OAUTHInfoIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *OAUTHInfo) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *OAUTHInfo) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *OAUTHInfo) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesOAUTHInfo{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *OAUTHInfo) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesOAUTHInfo{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *OAUTHInfo) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *OAUTHInfo) BleveType() string {

	return "oauthinfo"
}

// DefaultOrder returns the list of default ordering fields.
func (o *OAUTHInfo) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *OAUTHInfo) Doc() string {

	return `OAUTHInfo provides the information for an OAUTH server to retrieve the secrets
that can validate a JWT token issued by us.`
}

func (o *OAUTHInfo) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *OAUTHInfo) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseOAUTHInfo{
			IDTokenSigningAlgValuesSupported:  &o.IDTokenSigningAlgValuesSupported,
			JWKSURI:                           &o.JWKSURI,
			AuhorizationEndpoint:              &o.AuhorizationEndpoint,
			ClaimsSupported:                   &o.ClaimsSupported,
			Issuer:                            &o.Issuer,
			ResponseTypesSupported:            &o.ResponseTypesSupported,
			ScopesSupported:                   &o.ScopesSupported,
			SubjectTypesSupported:             &o.SubjectTypesSupported,
			TokenEndpointAuthMethodsSupported: &o.TokenEndpointAuthMethodsSupported,
		}
	}

	sp := &SparseOAUTHInfo{}
	for _, f := range fields {
		switch f {
		case "IDTokenSigningAlgValuesSupported":
			sp.IDTokenSigningAlgValuesSupported = &(o.IDTokenSigningAlgValuesSupported)
		case "JWKSURI":
			sp.JWKSURI = &(o.JWKSURI)
		case "auhorizationEndpoint":
			sp.AuhorizationEndpoint = &(o.AuhorizationEndpoint)
		case "claimsSupported":
			sp.ClaimsSupported = &(o.ClaimsSupported)
		case "issuer":
			sp.Issuer = &(o.Issuer)
		case "responseTypesSupported":
			sp.ResponseTypesSupported = &(o.ResponseTypesSupported)
		case "scopesSupported":
			sp.ScopesSupported = &(o.ScopesSupported)
		case "subjectTypesSupported":
			sp.SubjectTypesSupported = &(o.SubjectTypesSupported)
		case "tokenEndpointAuthMethodsSupported":
			sp.TokenEndpointAuthMethodsSupported = &(o.TokenEndpointAuthMethodsSupported)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseOAUTHInfo to the object.
func (o *OAUTHInfo) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseOAUTHInfo)
	if so.IDTokenSigningAlgValuesSupported != nil {
		o.IDTokenSigningAlgValuesSupported = *so.IDTokenSigningAlgValuesSupported
	}
	if so.JWKSURI != nil {
		o.JWKSURI = *so.JWKSURI
	}
	if so.AuhorizationEndpoint != nil {
		o.AuhorizationEndpoint = *so.AuhorizationEndpoint
	}
	if so.ClaimsSupported != nil {
		o.ClaimsSupported = *so.ClaimsSupported
	}
	if so.Issuer != nil {
		o.Issuer = *so.Issuer
	}
	if so.ResponseTypesSupported != nil {
		o.ResponseTypesSupported = *so.ResponseTypesSupported
	}
	if so.ScopesSupported != nil {
		o.ScopesSupported = *so.ScopesSupported
	}
	if so.SubjectTypesSupported != nil {
		o.SubjectTypesSupported = *so.SubjectTypesSupported
	}
	if so.TokenEndpointAuthMethodsSupported != nil {
		o.TokenEndpointAuthMethodsSupported = *so.TokenEndpointAuthMethodsSupported
	}
}

// DeepCopy returns a deep copy if the OAUTHInfo.
func (o *OAUTHInfo) DeepCopy() *OAUTHInfo {

	if o == nil {
		return nil
	}

	out := &OAUTHInfo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *OAUTHInfo.
func (o *OAUTHInfo) DeepCopyInto(out *OAUTHInfo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy OAUTHInfo: %s", err))
	}

	*out = *target.(*OAUTHInfo)
}

// Validate valides the current information stored into the structure.
func (o *OAUTHInfo) Validate() error {

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
func (*OAUTHInfo) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := OAUTHInfoAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return OAUTHInfoLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*OAUTHInfo) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return OAUTHInfoAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *OAUTHInfo) ValueForAttribute(name string) interface{} {

	switch name {
	case "IDTokenSigningAlgValuesSupported":
		return o.IDTokenSigningAlgValuesSupported
	case "JWKSURI":
		return o.JWKSURI
	case "auhorizationEndpoint":
		return o.AuhorizationEndpoint
	case "claimsSupported":
		return o.ClaimsSupported
	case "issuer":
		return o.Issuer
	case "responseTypesSupported":
		return o.ResponseTypesSupported
	case "scopesSupported":
		return o.ScopesSupported
	case "subjectTypesSupported":
		return o.SubjectTypesSupported
	case "tokenEndpointAuthMethodsSupported":
		return o.TokenEndpointAuthMethodsSupported
	}

	return nil
}

// OAUTHInfoAttributesMap represents the map of attribute for OAUTHInfo.
var OAUTHInfoAttributesMap = map[string]elemental.AttributeSpecification{
	"IDTokenSigningAlgValuesSupported": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "IDTokenSigningAlgValuesSupported",
		Description: `IDTokenSigningAlgValuesSupported is corresponding attribute of the OIDC
spec.`,
		Exposed:  true,
		Name:     "IDTokenSigningAlgValuesSupported",
		ReadOnly: true,
		SubType:  "string",
		Type:     "list",
	},
	"JWKSURI": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "JWKSURI",
		Description: `JWKSURI is the URI that can be used to retrieve the public keys that will
verify a JWT.`,
		Exposed:  true,
		Name:     "JWKSURI",
		ReadOnly: true,
		Type:     "string",
	},
	"AuhorizationEndpoint": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "AuhorizationEndpoint",
		Description:    `AuhorizationEndpoint is the authorization endpoint.`,
		Exposed:        true,
		Name:           "auhorizationEndpoint",
		ReadOnly:       true,
		Type:           "string",
	},
	"ClaimsSupported": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ClaimsSupported",
		Description:    `ClaimsSupported is corresponding attribute of the OIDC spec.`,
		Exposed:        true,
		Name:           "claimsSupported",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"Issuer": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Issuer",
		Description:    `Issuer is the the URL pointing to the issuer of the token.`,
		Exposed:        true,
		Name:           "issuer",
		ReadOnly:       true,
		Type:           "string",
	},
	"ResponseTypesSupported": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ResponseTypesSupported",
		Description:    `ResponseTypesSupported is corresponding attribute of the OIDC spec.`,
		Exposed:        true,
		Name:           "responseTypesSupported",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"ScopesSupported": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ScopesSupported",
		Description:    `ScopesSupported is corresponding attribute of the OIDC spec.`,
		Exposed:        true,
		Name:           "scopesSupported",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"SubjectTypesSupported": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SubjectTypesSupported",
		Description:    `SubjectTypesSupported is corresponding attribute of the OIDC spec.`,
		Exposed:        true,
		Name:           "subjectTypesSupported",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"TokenEndpointAuthMethodsSupported": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "TokenEndpointAuthMethodsSupported",
		Description: `TokenEndpointAuthMethodsSupported is corresponding attribute of the OIDC
spec.`,
		Exposed:  true,
		Name:     "tokenEndpointAuthMethodsSupported",
		ReadOnly: true,
		SubType:  "string",
		Type:     "list",
	},
}

// OAUTHInfoLowerCaseAttributesMap represents the map of attribute for OAUTHInfo.
var OAUTHInfoLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"idtokensigningalgvaluessupported": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "IDTokenSigningAlgValuesSupported",
		Description: `IDTokenSigningAlgValuesSupported is corresponding attribute of the OIDC
spec.`,
		Exposed:  true,
		Name:     "IDTokenSigningAlgValuesSupported",
		ReadOnly: true,
		SubType:  "string",
		Type:     "list",
	},
	"jwksuri": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "JWKSURI",
		Description: `JWKSURI is the URI that can be used to retrieve the public keys that will
verify a JWT.`,
		Exposed:  true,
		Name:     "JWKSURI",
		ReadOnly: true,
		Type:     "string",
	},
	"auhorizationendpoint": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "AuhorizationEndpoint",
		Description:    `AuhorizationEndpoint is the authorization endpoint.`,
		Exposed:        true,
		Name:           "auhorizationEndpoint",
		ReadOnly:       true,
		Type:           "string",
	},
	"claimssupported": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ClaimsSupported",
		Description:    `ClaimsSupported is corresponding attribute of the OIDC spec.`,
		Exposed:        true,
		Name:           "claimsSupported",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"issuer": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Issuer",
		Description:    `Issuer is the the URL pointing to the issuer of the token.`,
		Exposed:        true,
		Name:           "issuer",
		ReadOnly:       true,
		Type:           "string",
	},
	"responsetypessupported": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ResponseTypesSupported",
		Description:    `ResponseTypesSupported is corresponding attribute of the OIDC spec.`,
		Exposed:        true,
		Name:           "responseTypesSupported",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"scopessupported": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ScopesSupported",
		Description:    `ScopesSupported is corresponding attribute of the OIDC spec.`,
		Exposed:        true,
		Name:           "scopesSupported",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"subjecttypessupported": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SubjectTypesSupported",
		Description:    `SubjectTypesSupported is corresponding attribute of the OIDC spec.`,
		Exposed:        true,
		Name:           "subjectTypesSupported",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"tokenendpointauthmethodssupported": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "TokenEndpointAuthMethodsSupported",
		Description: `TokenEndpointAuthMethodsSupported is corresponding attribute of the OIDC
spec.`,
		Exposed:  true,
		Name:     "tokenEndpointAuthMethodsSupported",
		ReadOnly: true,
		SubType:  "string",
		Type:     "list",
	},
}

// SparseOAUTHInfosList represents a list of SparseOAUTHInfos
type SparseOAUTHInfosList []*SparseOAUTHInfo

// Identity returns the identity of the objects in the list.
func (o SparseOAUTHInfosList) Identity() elemental.Identity {

	return OAUTHInfoIdentity
}

// Copy returns a pointer to a copy the SparseOAUTHInfosList.
func (o SparseOAUTHInfosList) Copy() elemental.Identifiables {

	copy := append(SparseOAUTHInfosList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseOAUTHInfosList.
func (o SparseOAUTHInfosList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseOAUTHInfosList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseOAUTHInfo))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseOAUTHInfosList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseOAUTHInfosList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseOAUTHInfosList converted to OAUTHInfosList.
func (o SparseOAUTHInfosList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseOAUTHInfosList) Version() int {

	return 1
}

// SparseOAUTHInfo represents the sparse version of a oauthinfo.
type SparseOAUTHInfo struct {
	// IDTokenSigningAlgValuesSupported is corresponding attribute of the OIDC
	// spec.
	IDTokenSigningAlgValuesSupported *[]string `json:"id_token_signing_alg_values_supported,omitempty" msgpack:"id_token_signing_alg_values_supported,omitempty" bson:"-" mapstructure:"id_token_signing_alg_values_supported,omitempty"`

	// JWKSURI is the URI that can be used to retrieve the public keys that will
	// verify a JWT.
	JWKSURI *string `json:"jwks_uri,omitempty" msgpack:"jwks_uri,omitempty" bson:"-" mapstructure:"jwks_uri,omitempty"`

	// AuhorizationEndpoint is the authorization endpoint.
	AuhorizationEndpoint *string `json:"auhorization_endpoint,omitempty" msgpack:"auhorization_endpoint,omitempty" bson:"-" mapstructure:"auhorization_endpoint,omitempty"`

	// ClaimsSupported is corresponding attribute of the OIDC spec.
	ClaimsSupported *[]string `json:"claims_supported,omitempty" msgpack:"claims_supported,omitempty" bson:"-" mapstructure:"claims_supported,omitempty"`

	// Issuer is the the URL pointing to the issuer of the token.
	Issuer *string `json:"issuer,omitempty" msgpack:"issuer,omitempty" bson:"-" mapstructure:"issuer,omitempty"`

	// ResponseTypesSupported is corresponding attribute of the OIDC spec.
	ResponseTypesSupported *[]string `json:"response_types_supported,omitempty" msgpack:"response_types_supported,omitempty" bson:"-" mapstructure:"response_types_supported,omitempty"`

	// ScopesSupported is corresponding attribute of the OIDC spec.
	ScopesSupported *[]string `json:"scopes_supported,omitempty" msgpack:"scopes_supported,omitempty" bson:"-" mapstructure:"scopes_supported,omitempty"`

	// SubjectTypesSupported is corresponding attribute of the OIDC spec.
	SubjectTypesSupported *[]string `json:"subject_types_supported,omitempty" msgpack:"subject_types_supported,omitempty" bson:"-" mapstructure:"subject_types_supported,omitempty"`

	// TokenEndpointAuthMethodsSupported is corresponding attribute of the OIDC
	// spec.
	TokenEndpointAuthMethodsSupported *[]string `json:"token_endpoint_auth_methods_supported,omitempty" msgpack:"token_endpoint_auth_methods_supported,omitempty" bson:"-" mapstructure:"token_endpoint_auth_methods_supported,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseOAUTHInfo returns a new  SparseOAUTHInfo.
func NewSparseOAUTHInfo() *SparseOAUTHInfo {
	return &SparseOAUTHInfo{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseOAUTHInfo) Identity() elemental.Identity {

	return OAUTHInfoIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseOAUTHInfo) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseOAUTHInfo) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseOAUTHInfo) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseOAUTHInfo{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseOAUTHInfo) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseOAUTHInfo{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseOAUTHInfo) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseOAUTHInfo) ToPlain() elemental.PlainIdentifiable {

	out := NewOAUTHInfo()
	if o.IDTokenSigningAlgValuesSupported != nil {
		out.IDTokenSigningAlgValuesSupported = *o.IDTokenSigningAlgValuesSupported
	}
	if o.JWKSURI != nil {
		out.JWKSURI = *o.JWKSURI
	}
	if o.AuhorizationEndpoint != nil {
		out.AuhorizationEndpoint = *o.AuhorizationEndpoint
	}
	if o.ClaimsSupported != nil {
		out.ClaimsSupported = *o.ClaimsSupported
	}
	if o.Issuer != nil {
		out.Issuer = *o.Issuer
	}
	if o.ResponseTypesSupported != nil {
		out.ResponseTypesSupported = *o.ResponseTypesSupported
	}
	if o.ScopesSupported != nil {
		out.ScopesSupported = *o.ScopesSupported
	}
	if o.SubjectTypesSupported != nil {
		out.SubjectTypesSupported = *o.SubjectTypesSupported
	}
	if o.TokenEndpointAuthMethodsSupported != nil {
		out.TokenEndpointAuthMethodsSupported = *o.TokenEndpointAuthMethodsSupported
	}

	return out
}

// DeepCopy returns a deep copy if the SparseOAUTHInfo.
func (o *SparseOAUTHInfo) DeepCopy() *SparseOAUTHInfo {

	if o == nil {
		return nil
	}

	out := &SparseOAUTHInfo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseOAUTHInfo.
func (o *SparseOAUTHInfo) DeepCopyInto(out *SparseOAUTHInfo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseOAUTHInfo: %s", err))
	}

	*out = *target.(*SparseOAUTHInfo)
}

type mongoAttributesOAUTHInfo struct {
}
type mongoAttributesSparseOAUTHInfo struct {
}
