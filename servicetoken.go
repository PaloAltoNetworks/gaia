package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ServiceTokenTypeValue represents the possible values for attribute "type".
type ServiceTokenTypeValue string

const (
	// ServiceTokenTypeProcessingUnit represents the value ProcessingUnit.
	ServiceTokenTypeProcessingUnit ServiceTokenTypeValue = "ProcessingUnit"

	// ServiceTokenTypeService represents the value Service.
	ServiceTokenTypeService ServiceTokenTypeValue = "Service"
)

// ServiceTokenIdentity represents the Identity of the object.
var ServiceTokenIdentity = elemental.Identity{
	Name:     "servicetoken",
	Category: "servicetoken",
	Package:  "cactuar",
	Private:  false,
}

// ServiceTokensList represents a list of ServiceTokens
type ServiceTokensList []*ServiceToken

// Identity returns the identity of the objects in the list.
func (o ServiceTokensList) Identity() elemental.Identity {

	return ServiceTokenIdentity
}

// Copy returns a pointer to a copy the ServiceTokensList.
func (o ServiceTokensList) Copy() elemental.Identifiables {

	copy := append(ServiceTokensList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ServiceTokensList.
func (o ServiceTokensList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ServiceTokensList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ServiceToken))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ServiceTokensList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ServiceTokensList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ServiceTokensList converted to SparseServiceTokensList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ServiceTokensList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseServiceTokensList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseServiceToken)
	}

	return out
}

// Version returns the version of the content.
func (o ServiceTokensList) Version() int {

	return 1
}

// ServiceToken represents the model of a servicetoken
type ServiceToken struct {
	// If given, the issued token will only be valid for the audience provided. If
	// empty, the audience will be resolved from the policies. If no audience can be
	// resolved, the request will be rejected with an error.
	Audience string `json:"audience" msgpack:"audience" bson:"-" mapstructure:"audience,omitempty"`

	// ID of the object you want to issue a token for.
	ObjectID string `json:"objectID" msgpack:"objectID" bson:"-" mapstructure:"objectID,omitempty"`

	// Provides the session ID of the enforcer when retrieving a datapath certificate.
	SessionID string `json:"sessionID" msgpack:"sessionID" bson:"-" mapstructure:"sessionID,omitempty"`

	// Token is the signed JWT service token.
	Token string `json:"token" msgpack:"token" bson:"-" mapstructure:"token,omitempty"`

	// Type of token request.
	Type ServiceTokenTypeValue `json:"type" msgpack:"type" bson:"-" mapstructure:"type,omitempty"`

	// Validity configures the max validity time for a token. If it is bigger than the
	// configured max validity, it will be capped.
	Validity string `json:"validity" msgpack:"validity" bson:"-" mapstructure:"validity,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewServiceToken returns a new *ServiceToken
func NewServiceToken() *ServiceToken {

	return &ServiceToken{
		ModelVersion: 1,
		Type:         ServiceTokenTypeService,
		Validity:     "15m",
	}
}

// Identity returns the Identity of the object.
func (o *ServiceToken) Identity() elemental.Identity {

	return ServiceTokenIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ServiceToken) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ServiceToken) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ServiceToken) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesServiceToken{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ServiceToken) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesServiceToken{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *ServiceToken) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *ServiceToken) BleveType() string {

	return "servicetoken"
}

// DefaultOrder returns the list of default ordering fields.
func (o *ServiceToken) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *ServiceToken) Doc() string {

	return `This API issues a new service token using the namespace certificate that can be
used by third party applications.`
}

func (o *ServiceToken) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ServiceToken) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseServiceToken{
			Audience:  &o.Audience,
			ObjectID:  &o.ObjectID,
			SessionID: &o.SessionID,
			Token:     &o.Token,
			Type:      &o.Type,
			Validity:  &o.Validity,
		}
	}

	sp := &SparseServiceToken{}
	for _, f := range fields {
		switch f {
		case "audience":
			sp.Audience = &(o.Audience)
		case "objectID":
			sp.ObjectID = &(o.ObjectID)
		case "sessionID":
			sp.SessionID = &(o.SessionID)
		case "token":
			sp.Token = &(o.Token)
		case "type":
			sp.Type = &(o.Type)
		case "validity":
			sp.Validity = &(o.Validity)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseServiceToken to the object.
func (o *ServiceToken) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseServiceToken)
	if so.Audience != nil {
		o.Audience = *so.Audience
	}
	if so.ObjectID != nil {
		o.ObjectID = *so.ObjectID
	}
	if so.SessionID != nil {
		o.SessionID = *so.SessionID
	}
	if so.Token != nil {
		o.Token = *so.Token
	}
	if so.Type != nil {
		o.Type = *so.Type
	}
	if so.Validity != nil {
		o.Validity = *so.Validity
	}
}

// DeepCopy returns a deep copy if the ServiceToken.
func (o *ServiceToken) DeepCopy() *ServiceToken {

	if o == nil {
		return nil
	}

	out := &ServiceToken{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ServiceToken.
func (o *ServiceToken) DeepCopyInto(out *ServiceToken) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ServiceToken: %s", err))
	}

	*out = *target.(*ServiceToken)
}

// Validate valides the current information stored into the structure.
func (o *ServiceToken) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateAudience("audience", o.Audience); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"ProcessingUnit", "Service"}, false); err != nil {
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
func (*ServiceToken) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ServiceTokenAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ServiceTokenLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ServiceToken) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ServiceTokenAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ServiceToken) ValueForAttribute(name string) interface{} {

	switch name {
	case "audience":
		return o.Audience
	case "objectID":
		return o.ObjectID
	case "sessionID":
		return o.SessionID
	case "token":
		return o.Token
	case "type":
		return o.Type
	case "validity":
		return o.Validity
	}

	return nil
}

// ServiceTokenAttributesMap represents the map of attribute for ServiceToken.
var ServiceTokenAttributesMap = map[string]elemental.AttributeSpecification{
	"Audience": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Audience",
		Description: `If given, the issued token will only be valid for the audience provided. If
empty, the audience will be resolved from the policies. If no audience can be
resolved, the request will be rejected with an error.`,
		Exposed: true,
		Name:    "audience",
		Type:    "string",
	},
	"ObjectID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObjectID",
		Description:    `ID of the object you want to issue a token for.`,
		Exposed:        true,
		Name:           "objectID",
		Type:           "string",
	},
	"SessionID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SessionID",
		Description:    `Provides the session ID of the enforcer when retrieving a datapath certificate.`,
		Exposed:        true,
		Name:           "sessionID",
		Type:           "string",
	},
	"Token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Token",
		Description:    `Token is the signed JWT service token.`,
		Exposed:        true,
		Name:           "token",
		ReadOnly:       true,
		Type:           "string",
	},
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{"ProcessingUnit", "Service"},
		ConvertedName:  "Type",
		DefaultValue:   ServiceTokenTypeService,
		Description:    `Type of token request.`,
		Exposed:        true,
		Name:           "type",
		Type:           "enum",
	},
	"Validity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Validity",
		DefaultValue:   "15m",
		Description: `Validity configures the max validity time for a token. If it is bigger than the
configured max validity, it will be capped.`,
		Exposed:   true,
		Name:      "validity",
		Orderable: true,
		Type:      "string",
	},
}

// ServiceTokenLowerCaseAttributesMap represents the map of attribute for ServiceToken.
var ServiceTokenLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"audience": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Audience",
		Description: `If given, the issued token will only be valid for the audience provided. If
empty, the audience will be resolved from the policies. If no audience can be
resolved, the request will be rejected with an error.`,
		Exposed: true,
		Name:    "audience",
		Type:    "string",
	},
	"objectid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObjectID",
		Description:    `ID of the object you want to issue a token for.`,
		Exposed:        true,
		Name:           "objectID",
		Type:           "string",
	},
	"sessionid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SessionID",
		Description:    `Provides the session ID of the enforcer when retrieving a datapath certificate.`,
		Exposed:        true,
		Name:           "sessionID",
		Type:           "string",
	},
	"token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Token",
		Description:    `Token is the signed JWT service token.`,
		Exposed:        true,
		Name:           "token",
		ReadOnly:       true,
		Type:           "string",
	},
	"type": elemental.AttributeSpecification{
		AllowedChoices: []string{"ProcessingUnit", "Service"},
		ConvertedName:  "Type",
		DefaultValue:   ServiceTokenTypeService,
		Description:    `Type of token request.`,
		Exposed:        true,
		Name:           "type",
		Type:           "enum",
	},
	"validity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Validity",
		DefaultValue:   "15m",
		Description: `Validity configures the max validity time for a token. If it is bigger than the
configured max validity, it will be capped.`,
		Exposed:   true,
		Name:      "validity",
		Orderable: true,
		Type:      "string",
	},
}

// SparseServiceTokensList represents a list of SparseServiceTokens
type SparseServiceTokensList []*SparseServiceToken

// Identity returns the identity of the objects in the list.
func (o SparseServiceTokensList) Identity() elemental.Identity {

	return ServiceTokenIdentity
}

// Copy returns a pointer to a copy the SparseServiceTokensList.
func (o SparseServiceTokensList) Copy() elemental.Identifiables {

	copy := append(SparseServiceTokensList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseServiceTokensList.
func (o SparseServiceTokensList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseServiceTokensList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseServiceToken))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseServiceTokensList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseServiceTokensList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseServiceTokensList converted to ServiceTokensList.
func (o SparseServiceTokensList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseServiceTokensList) Version() int {

	return 1
}

// SparseServiceToken represents the sparse version of a servicetoken.
type SparseServiceToken struct {
	// If given, the issued token will only be valid for the audience provided. If
	// empty, the audience will be resolved from the policies. If no audience can be
	// resolved, the request will be rejected with an error.
	Audience *string `json:"audience,omitempty" msgpack:"audience,omitempty" bson:"-" mapstructure:"audience,omitempty"`

	// ID of the object you want to issue a token for.
	ObjectID *string `json:"objectID,omitempty" msgpack:"objectID,omitempty" bson:"-" mapstructure:"objectID,omitempty"`

	// Provides the session ID of the enforcer when retrieving a datapath certificate.
	SessionID *string `json:"sessionID,omitempty" msgpack:"sessionID,omitempty" bson:"-" mapstructure:"sessionID,omitempty"`

	// Token is the signed JWT service token.
	Token *string `json:"token,omitempty" msgpack:"token,omitempty" bson:"-" mapstructure:"token,omitempty"`

	// Type of token request.
	Type *ServiceTokenTypeValue `json:"type,omitempty" msgpack:"type,omitempty" bson:"-" mapstructure:"type,omitempty"`

	// Validity configures the max validity time for a token. If it is bigger than the
	// configured max validity, it will be capped.
	Validity *string `json:"validity,omitempty" msgpack:"validity,omitempty" bson:"-" mapstructure:"validity,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseServiceToken returns a new  SparseServiceToken.
func NewSparseServiceToken() *SparseServiceToken {
	return &SparseServiceToken{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseServiceToken) Identity() elemental.Identity {

	return ServiceTokenIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseServiceToken) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseServiceToken) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseServiceToken) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseServiceToken{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseServiceToken) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseServiceToken{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseServiceToken) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseServiceToken) ToPlain() elemental.PlainIdentifiable {

	out := NewServiceToken()
	if o.Audience != nil {
		out.Audience = *o.Audience
	}
	if o.ObjectID != nil {
		out.ObjectID = *o.ObjectID
	}
	if o.SessionID != nil {
		out.SessionID = *o.SessionID
	}
	if o.Token != nil {
		out.Token = *o.Token
	}
	if o.Type != nil {
		out.Type = *o.Type
	}
	if o.Validity != nil {
		out.Validity = *o.Validity
	}

	return out
}

// DeepCopy returns a deep copy if the SparseServiceToken.
func (o *SparseServiceToken) DeepCopy() *SparseServiceToken {

	if o == nil {
		return nil
	}

	out := &SparseServiceToken{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseServiceToken.
func (o *SparseServiceToken) DeepCopyInto(out *SparseServiceToken) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseServiceToken: %s", err))
	}

	*out = *target.(*SparseServiceToken)
}

type mongoAttributesServiceToken struct {
}
type mongoAttributesSparseServiceToken struct {
}
