package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// IssueServiceTokenIdentity represents the Identity of the object.
var IssueServiceTokenIdentity = elemental.Identity{
	Name:     "issueservicetoken",
	Category: "issueservicetokens",
	Package:  "barret",
	Private:  true,
}

// IssueServiceTokensList represents a list of IssueServiceTokens
type IssueServiceTokensList []*IssueServiceToken

// Identity returns the identity of the objects in the list.
func (o IssueServiceTokensList) Identity() elemental.Identity {

	return IssueServiceTokenIdentity
}

// Copy returns a pointer to a copy the IssueServiceTokensList.
func (o IssueServiceTokensList) Copy() elemental.Identifiables {

	copy := append(IssueServiceTokensList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the IssueServiceTokensList.
func (o IssueServiceTokensList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(IssueServiceTokensList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*IssueServiceToken))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o IssueServiceTokensList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o IssueServiceTokensList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the IssueServiceTokensList converted to SparseIssueServiceTokensList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o IssueServiceTokensList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseIssueServiceTokensList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseIssueServiceToken)
	}

	return out
}

// Version returns the version of the content.
func (o IssueServiceTokensList) Version() int {

	return 1
}

// IssueServiceToken represents the model of a issueservicetoken
type IssueServiceToken struct {
	// Audience is the valid audience for this token.
	Audience string `json:"audience" msgpack:"audience" bson:"-" mapstructure:"audience,omitempty"`

	// The ID of the corresponding namespace.
	NamespaceID string `json:"namespaceID" msgpack:"namespaceID" bson:"-" mapstructure:"namespaceID,omitempty"`

	// ServiceClaims is a list of service claims that have been validated provided as
	// key/value pairs. If the same key is provided multiple times it will be converted
	// to an array. The claims  will appear under the Data section of the token.
	ServiceClaims []string `json:"serviceClaims" msgpack:"serviceClaims" bson:"-" mapstructure:"serviceClaims,omitempty"`

	// SigningKeyID holds the ID of the private certificate to use to sign the token.
	SigningKeyID string `json:"signingKeyID" msgpack:"signingKeyID" bson:"signingkeyid" mapstructure:"signingKeyID,omitempty"`

	// The subject claims of the token.
	Subject string `json:"subject" msgpack:"subject" bson:"-" mapstructure:"subject,omitempty"`

	// Token contains the generated token.
	Token string `json:"token" msgpack:"token" bson:"-" mapstructure:"token,omitempty"`

	// UserClaims is a list of user claims that have been validated provided as
	// key/value pairs. If the same key is provided multiple times it will be converted
	// to an array. The claims  will appear under the Data section of the token.
	UserClaims []string `json:"userClaims" msgpack:"userClaims" bson:"-" mapstructure:"userClaims,omitempty"`

	// Validity contains the token validity duration.
	Validity string `json:"validity" msgpack:"validity" bson:"-" mapstructure:"validity,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewIssueServiceToken returns a new *IssueServiceToken
func NewIssueServiceToken() *IssueServiceToken {

	return &IssueServiceToken{
		ModelVersion:  1,
		ServiceClaims: []string{},
		UserClaims:    []string{},
		Validity:      "15m",
	}
}

// Identity returns the Identity of the object.
func (o *IssueServiceToken) Identity() elemental.Identity {

	return IssueServiceTokenIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IssueServiceToken) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IssueServiceToken) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *IssueServiceToken) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesIssueServiceToken{}

	s.SigningKeyID = o.SigningKeyID

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *IssueServiceToken) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesIssueServiceToken{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.SigningKeyID = s.SigningKeyID

	return nil
}

// Version returns the hardcoded version of the model.
func (o *IssueServiceToken) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *IssueServiceToken) BleveType() string {

	return "issueservicetoken"
}

// DefaultOrder returns the list of default ordering fields.
func (o *IssueServiceToken) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *IssueServiceToken) Doc() string {

	return `This is an internal API. Services can call this API to issue service tokens on
behalf of users.`
}

func (o *IssueServiceToken) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *IssueServiceToken) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseIssueServiceToken{
			Audience:      &o.Audience,
			NamespaceID:   &o.NamespaceID,
			ServiceClaims: &o.ServiceClaims,
			SigningKeyID:  &o.SigningKeyID,
			Subject:       &o.Subject,
			Token:         &o.Token,
			UserClaims:    &o.UserClaims,
			Validity:      &o.Validity,
		}
	}

	sp := &SparseIssueServiceToken{}
	for _, f := range fields {
		switch f {
		case "audience":
			sp.Audience = &(o.Audience)
		case "namespaceID":
			sp.NamespaceID = &(o.NamespaceID)
		case "serviceClaims":
			sp.ServiceClaims = &(o.ServiceClaims)
		case "signingKeyID":
			sp.SigningKeyID = &(o.SigningKeyID)
		case "subject":
			sp.Subject = &(o.Subject)
		case "token":
			sp.Token = &(o.Token)
		case "userClaims":
			sp.UserClaims = &(o.UserClaims)
		case "validity":
			sp.Validity = &(o.Validity)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseIssueServiceToken to the object.
func (o *IssueServiceToken) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseIssueServiceToken)
	if so.Audience != nil {
		o.Audience = *so.Audience
	}
	if so.NamespaceID != nil {
		o.NamespaceID = *so.NamespaceID
	}
	if so.ServiceClaims != nil {
		o.ServiceClaims = *so.ServiceClaims
	}
	if so.SigningKeyID != nil {
		o.SigningKeyID = *so.SigningKeyID
	}
	if so.Subject != nil {
		o.Subject = *so.Subject
	}
	if so.Token != nil {
		o.Token = *so.Token
	}
	if so.UserClaims != nil {
		o.UserClaims = *so.UserClaims
	}
	if so.Validity != nil {
		o.Validity = *so.Validity
	}
}

// DeepCopy returns a deep copy if the IssueServiceToken.
func (o *IssueServiceToken) DeepCopy() *IssueServiceToken {

	if o == nil {
		return nil
	}

	out := &IssueServiceToken{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *IssueServiceToken.
func (o *IssueServiceToken) DeepCopyInto(out *IssueServiceToken) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy IssueServiceToken: %s", err))
	}

	*out = *target.(*IssueServiceToken)
}

// Validate valides the current information stored into the structure.
func (o *IssueServiceToken) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("audience", o.Audience); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("signingKeyID", o.SigningKeyID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("subject", o.Subject); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("validity", o.Validity); err != nil {
		requiredErrors = requiredErrors.Append(err)
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
func (*IssueServiceToken) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := IssueServiceTokenAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return IssueServiceTokenLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*IssueServiceToken) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return IssueServiceTokenAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *IssueServiceToken) ValueForAttribute(name string) interface{} {

	switch name {
	case "audience":
		return o.Audience
	case "namespaceID":
		return o.NamespaceID
	case "serviceClaims":
		return o.ServiceClaims
	case "signingKeyID":
		return o.SigningKeyID
	case "subject":
		return o.Subject
	case "token":
		return o.Token
	case "userClaims":
		return o.UserClaims
	case "validity":
		return o.Validity
	}

	return nil
}

// IssueServiceTokenAttributesMap represents the map of attribute for IssueServiceToken.
var IssueServiceTokenAttributesMap = map[string]elemental.AttributeSpecification{
	"Audience": {
		AllowedChoices: []string{},
		ConvertedName:  "Audience",
		CreationOnly:   true,
		Description:    `Audience is the valid audience for this token.`,
		Exposed:        true,
		Name:           "audience",
		Required:       true,
		Type:           "string",
	},
	"NamespaceID": {
		AllowedChoices: []string{},
		ConvertedName:  "NamespaceID",
		CreationOnly:   true,
		Description:    `The ID of the corresponding namespace.`,
		Exposed:        true,
		Name:           "namespaceID",
		Type:           "string",
	},
	"ServiceClaims": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceClaims",
		CreationOnly:   true,
		Description: `ServiceClaims is a list of service claims that have been validated provided as
key/value pairs. If the same key is provided multiple times it will be converted
to an array. The claims  will appear under the Data section of the token.`,
		Exposed: true,
		Name:    "serviceClaims",
		SubType: "string",
		Type:    "list",
	},
	"SigningKeyID": {
		AllowedChoices: []string{},
		ConvertedName:  "SigningKeyID",
		Description:    `SigningKeyID holds the ID of the private certificate to use to sign the token.`,
		Exposed:        true,
		Name:           "signingKeyID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Subject": {
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		CreationOnly:   true,
		Description:    `The subject claims of the token.`,
		Exposed:        true,
		Name:           "subject",
		Required:       true,
		Type:           "string",
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
	"UserClaims": {
		AllowedChoices: []string{},
		ConvertedName:  "UserClaims",
		CreationOnly:   true,
		Description: `UserClaims is a list of user claims that have been validated provided as
key/value pairs. If the same key is provided multiple times it will be converted
to an array. The claims  will appear under the Data section of the token.`,
		Exposed: true,
		Name:    "userClaims",
		SubType: "string",
		Type:    "list",
	},
	"Validity": {
		AllowedChoices: []string{},
		ConvertedName:  "Validity",
		CreationOnly:   true,
		DefaultValue:   "15m",
		Description:    `Validity contains the token validity duration.`,
		Exposed:        true,
		Name:           "validity",
		Required:       true,
		Type:           "string",
	},
}

// IssueServiceTokenLowerCaseAttributesMap represents the map of attribute for IssueServiceToken.
var IssueServiceTokenLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"audience": {
		AllowedChoices: []string{},
		ConvertedName:  "Audience",
		CreationOnly:   true,
		Description:    `Audience is the valid audience for this token.`,
		Exposed:        true,
		Name:           "audience",
		Required:       true,
		Type:           "string",
	},
	"namespaceid": {
		AllowedChoices: []string{},
		ConvertedName:  "NamespaceID",
		CreationOnly:   true,
		Description:    `The ID of the corresponding namespace.`,
		Exposed:        true,
		Name:           "namespaceID",
		Type:           "string",
	},
	"serviceclaims": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceClaims",
		CreationOnly:   true,
		Description: `ServiceClaims is a list of service claims that have been validated provided as
key/value pairs. If the same key is provided multiple times it will be converted
to an array. The claims  will appear under the Data section of the token.`,
		Exposed: true,
		Name:    "serviceClaims",
		SubType: "string",
		Type:    "list",
	},
	"signingkeyid": {
		AllowedChoices: []string{},
		ConvertedName:  "SigningKeyID",
		Description:    `SigningKeyID holds the ID of the private certificate to use to sign the token.`,
		Exposed:        true,
		Name:           "signingKeyID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"subject": {
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		CreationOnly:   true,
		Description:    `The subject claims of the token.`,
		Exposed:        true,
		Name:           "subject",
		Required:       true,
		Type:           "string",
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
	"userclaims": {
		AllowedChoices: []string{},
		ConvertedName:  "UserClaims",
		CreationOnly:   true,
		Description: `UserClaims is a list of user claims that have been validated provided as
key/value pairs. If the same key is provided multiple times it will be converted
to an array. The claims  will appear under the Data section of the token.`,
		Exposed: true,
		Name:    "userClaims",
		SubType: "string",
		Type:    "list",
	},
	"validity": {
		AllowedChoices: []string{},
		ConvertedName:  "Validity",
		CreationOnly:   true,
		DefaultValue:   "15m",
		Description:    `Validity contains the token validity duration.`,
		Exposed:        true,
		Name:           "validity",
		Required:       true,
		Type:           "string",
	},
}

// SparseIssueServiceTokensList represents a list of SparseIssueServiceTokens
type SparseIssueServiceTokensList []*SparseIssueServiceToken

// Identity returns the identity of the objects in the list.
func (o SparseIssueServiceTokensList) Identity() elemental.Identity {

	return IssueServiceTokenIdentity
}

// Copy returns a pointer to a copy the SparseIssueServiceTokensList.
func (o SparseIssueServiceTokensList) Copy() elemental.Identifiables {

	copy := append(SparseIssueServiceTokensList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseIssueServiceTokensList.
func (o SparseIssueServiceTokensList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseIssueServiceTokensList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseIssueServiceToken))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseIssueServiceTokensList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseIssueServiceTokensList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseIssueServiceTokensList converted to IssueServiceTokensList.
func (o SparseIssueServiceTokensList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseIssueServiceTokensList) Version() int {

	return 1
}

// SparseIssueServiceToken represents the sparse version of a issueservicetoken.
type SparseIssueServiceToken struct {
	// Audience is the valid audience for this token.
	Audience *string `json:"audience,omitempty" msgpack:"audience,omitempty" bson:"-" mapstructure:"audience,omitempty"`

	// The ID of the corresponding namespace.
	NamespaceID *string `json:"namespaceID,omitempty" msgpack:"namespaceID,omitempty" bson:"-" mapstructure:"namespaceID,omitempty"`

	// ServiceClaims is a list of service claims that have been validated provided as
	// key/value pairs. If the same key is provided multiple times it will be converted
	// to an array. The claims  will appear under the Data section of the token.
	ServiceClaims *[]string `json:"serviceClaims,omitempty" msgpack:"serviceClaims,omitempty" bson:"-" mapstructure:"serviceClaims,omitempty"`

	// SigningKeyID holds the ID of the private certificate to use to sign the token.
	SigningKeyID *string `json:"signingKeyID,omitempty" msgpack:"signingKeyID,omitempty" bson:"signingkeyid,omitempty" mapstructure:"signingKeyID,omitempty"`

	// The subject claims of the token.
	Subject *string `json:"subject,omitempty" msgpack:"subject,omitempty" bson:"-" mapstructure:"subject,omitempty"`

	// Token contains the generated token.
	Token *string `json:"token,omitempty" msgpack:"token,omitempty" bson:"-" mapstructure:"token,omitempty"`

	// UserClaims is a list of user claims that have been validated provided as
	// key/value pairs. If the same key is provided multiple times it will be converted
	// to an array. The claims  will appear under the Data section of the token.
	UserClaims *[]string `json:"userClaims,omitempty" msgpack:"userClaims,omitempty" bson:"-" mapstructure:"userClaims,omitempty"`

	// Validity contains the token validity duration.
	Validity *string `json:"validity,omitempty" msgpack:"validity,omitempty" bson:"-" mapstructure:"validity,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseIssueServiceToken returns a new  SparseIssueServiceToken.
func NewSparseIssueServiceToken() *SparseIssueServiceToken {
	return &SparseIssueServiceToken{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseIssueServiceToken) Identity() elemental.Identity {

	return IssueServiceTokenIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseIssueServiceToken) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseIssueServiceToken) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseIssueServiceToken) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseIssueServiceToken{}

	if o.SigningKeyID != nil {
		s.SigningKeyID = o.SigningKeyID
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseIssueServiceToken) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseIssueServiceToken{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.SigningKeyID != nil {
		o.SigningKeyID = s.SigningKeyID
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseIssueServiceToken) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseIssueServiceToken) ToPlain() elemental.PlainIdentifiable {

	out := NewIssueServiceToken()
	if o.Audience != nil {
		out.Audience = *o.Audience
	}
	if o.NamespaceID != nil {
		out.NamespaceID = *o.NamespaceID
	}
	if o.ServiceClaims != nil {
		out.ServiceClaims = *o.ServiceClaims
	}
	if o.SigningKeyID != nil {
		out.SigningKeyID = *o.SigningKeyID
	}
	if o.Subject != nil {
		out.Subject = *o.Subject
	}
	if o.Token != nil {
		out.Token = *o.Token
	}
	if o.UserClaims != nil {
		out.UserClaims = *o.UserClaims
	}
	if o.Validity != nil {
		out.Validity = *o.Validity
	}

	return out
}

// DeepCopy returns a deep copy if the SparseIssueServiceToken.
func (o *SparseIssueServiceToken) DeepCopy() *SparseIssueServiceToken {

	if o == nil {
		return nil
	}

	out := &SparseIssueServiceToken{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseIssueServiceToken.
func (o *SparseIssueServiceToken) DeepCopyInto(out *SparseIssueServiceToken) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseIssueServiceToken: %s", err))
	}

	*out = *target.(*SparseIssueServiceToken)
}

type mongoAttributesIssueServiceToken struct {
	SigningKeyID string `bson:"signingkeyid"`
}
type mongoAttributesSparseIssueServiceToken struct {
	SigningKeyID *string `bson:"signingkeyid,omitempty"`
}
