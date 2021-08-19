package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ValidateRQLReqIdentity represents the Identity of the object.
var ValidateRQLReqIdentity = elemental.Identity{
	Name:     "validaterql",
	Category: "validaterql",
	Package:  "larl",
	Private:  false,
}

// ValidateRQLReqsList represents a list of ValidateRQLReqs
type ValidateRQLReqsList []*ValidateRQLReq

// Identity returns the identity of the objects in the list.
func (o ValidateRQLReqsList) Identity() elemental.Identity {

	return ValidateRQLReqIdentity
}

// Copy returns a pointer to a copy the ValidateRQLReqsList.
func (o ValidateRQLReqsList) Copy() elemental.Identifiables {

	copy := append(ValidateRQLReqsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ValidateRQLReqsList.
func (o ValidateRQLReqsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ValidateRQLReqsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ValidateRQLReq))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ValidateRQLReqsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ValidateRQLReqsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ValidateRQLReqsList converted to SparseValidateRQLReqsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ValidateRQLReqsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseValidateRQLReqsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseValidateRQLReq)
	}

	return out
}

// Version returns the version of the content.
func (o ValidateRQLReqsList) Version() int {

	return 1
}

// ValidateRQLReq represents the model of a validaterql
type ValidateRQLReq struct {
	// TODO.
	PolicyType string `json:"policyType" msgpack:"policyType" bson:"-" mapstructure:"policyType,omitempty"`

	// TODO.
	PrismaId int `json:"prismaId" msgpack:"prismaId" bson:"-" mapstructure:"prismaId,omitempty"`

	// TODO.
	Query string `json:"query" msgpack:"query" bson:"-" mapstructure:"query,omitempty"`

	// TODO.
	SearchType string `json:"searchType" msgpack:"searchType" bson:"-" mapstructure:"searchType,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewValidateRQLReq returns a new *ValidateRQLReq
func NewValidateRQLReq() *ValidateRQLReq {

	return &ValidateRQLReq{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *ValidateRQLReq) Identity() elemental.Identity {

	return ValidateRQLReqIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ValidateRQLReq) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ValidateRQLReq) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ValidateRQLReq) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesValidateRQLReq{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ValidateRQLReq) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesValidateRQLReq{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *ValidateRQLReq) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *ValidateRQLReq) BleveType() string {

	return "validaterql"
}

// DefaultOrder returns the list of default ordering fields.
func (o *ValidateRQLReq) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *ValidateRQLReq) Doc() string {

	return `TODO.`
}

func (o *ValidateRQLReq) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ValidateRQLReq) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseValidateRQLReq{
			PolicyType: &o.PolicyType,
			PrismaId:   &o.PrismaId,
			Query:      &o.Query,
			SearchType: &o.SearchType,
		}
	}

	sp := &SparseValidateRQLReq{}
	for _, f := range fields {
		switch f {
		case "policyType":
			sp.PolicyType = &(o.PolicyType)
		case "prismaId":
			sp.PrismaId = &(o.PrismaId)
		case "query":
			sp.Query = &(o.Query)
		case "searchType":
			sp.SearchType = &(o.SearchType)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseValidateRQLReq to the object.
func (o *ValidateRQLReq) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseValidateRQLReq)
	if so.PolicyType != nil {
		o.PolicyType = *so.PolicyType
	}
	if so.PrismaId != nil {
		o.PrismaId = *so.PrismaId
	}
	if so.Query != nil {
		o.Query = *so.Query
	}
	if so.SearchType != nil {
		o.SearchType = *so.SearchType
	}
}

// DeepCopy returns a deep copy if the ValidateRQLReq.
func (o *ValidateRQLReq) DeepCopy() *ValidateRQLReq {

	if o == nil {
		return nil
	}

	out := &ValidateRQLReq{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ValidateRQLReq.
func (o *ValidateRQLReq) DeepCopyInto(out *ValidateRQLReq) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ValidateRQLReq: %s", err))
	}

	*out = *target.(*ValidateRQLReq)
}

// Validate valides the current information stored into the structure.
func (o *ValidateRQLReq) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("policyType", o.PolicyType); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("prismaId", o.PrismaId); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("query", o.Query); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("searchType", o.SearchType); err != nil {
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
func (*ValidateRQLReq) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ValidateRQLReqAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ValidateRQLReqLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ValidateRQLReq) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ValidateRQLReqAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ValidateRQLReq) ValueForAttribute(name string) interface{} {

	switch name {
	case "policyType":
		return o.PolicyType
	case "prismaId":
		return o.PrismaId
	case "query":
		return o.Query
	case "searchType":
		return o.SearchType
	}

	return nil
}

// ValidateRQLReqAttributesMap represents the map of attribute for ValidateRQLReq.
var ValidateRQLReqAttributesMap = map[string]elemental.AttributeSpecification{
	"PolicyType": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyType",
		Description:    `TODO.`,
		Exposed:        true,
		Name:           "policyType",
		Required:       true,
		Type:           "string",
	},
	"PrismaId": {
		AllowedChoices: []string{},
		ConvertedName:  "PrismaId",
		Description:    `TODO.`,
		Exposed:        true,
		Name:           "prismaId",
		Required:       true,
		Type:           "integer",
	},
	"Query": {
		AllowedChoices: []string{},
		ConvertedName:  "Query",
		Description:    `TODO.`,
		Exposed:        true,
		Name:           "query",
		Required:       true,
		Type:           "string",
	},
	"SearchType": {
		AllowedChoices: []string{},
		ConvertedName:  "SearchType",
		Description:    `TODO.`,
		Exposed:        true,
		Name:           "searchType",
		Required:       true,
		Type:           "string",
	},
}

// ValidateRQLReqLowerCaseAttributesMap represents the map of attribute for ValidateRQLReq.
var ValidateRQLReqLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"policytype": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyType",
		Description:    `TODO.`,
		Exposed:        true,
		Name:           "policyType",
		Required:       true,
		Type:           "string",
	},
	"prismaid": {
		AllowedChoices: []string{},
		ConvertedName:  "PrismaId",
		Description:    `TODO.`,
		Exposed:        true,
		Name:           "prismaId",
		Required:       true,
		Type:           "integer",
	},
	"query": {
		AllowedChoices: []string{},
		ConvertedName:  "Query",
		Description:    `TODO.`,
		Exposed:        true,
		Name:           "query",
		Required:       true,
		Type:           "string",
	},
	"searchtype": {
		AllowedChoices: []string{},
		ConvertedName:  "SearchType",
		Description:    `TODO.`,
		Exposed:        true,
		Name:           "searchType",
		Required:       true,
		Type:           "string",
	},
}

// SparseValidateRQLReqsList represents a list of SparseValidateRQLReqs
type SparseValidateRQLReqsList []*SparseValidateRQLReq

// Identity returns the identity of the objects in the list.
func (o SparseValidateRQLReqsList) Identity() elemental.Identity {

	return ValidateRQLReqIdentity
}

// Copy returns a pointer to a copy the SparseValidateRQLReqsList.
func (o SparseValidateRQLReqsList) Copy() elemental.Identifiables {

	copy := append(SparseValidateRQLReqsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseValidateRQLReqsList.
func (o SparseValidateRQLReqsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseValidateRQLReqsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseValidateRQLReq))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseValidateRQLReqsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseValidateRQLReqsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseValidateRQLReqsList converted to ValidateRQLReqsList.
func (o SparseValidateRQLReqsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseValidateRQLReqsList) Version() int {

	return 1
}

// SparseValidateRQLReq represents the sparse version of a validaterql.
type SparseValidateRQLReq struct {
	// TODO.
	PolicyType *string `json:"policyType,omitempty" msgpack:"policyType,omitempty" bson:"-" mapstructure:"policyType,omitempty"`

	// TODO.
	PrismaId *int `json:"prismaId,omitempty" msgpack:"prismaId,omitempty" bson:"-" mapstructure:"prismaId,omitempty"`

	// TODO.
	Query *string `json:"query,omitempty" msgpack:"query,omitempty" bson:"-" mapstructure:"query,omitempty"`

	// TODO.
	SearchType *string `json:"searchType,omitempty" msgpack:"searchType,omitempty" bson:"-" mapstructure:"searchType,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseValidateRQLReq returns a new  SparseValidateRQLReq.
func NewSparseValidateRQLReq() *SparseValidateRQLReq {
	return &SparseValidateRQLReq{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseValidateRQLReq) Identity() elemental.Identity {

	return ValidateRQLReqIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseValidateRQLReq) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseValidateRQLReq) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseValidateRQLReq) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseValidateRQLReq{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseValidateRQLReq) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseValidateRQLReq{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseValidateRQLReq) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseValidateRQLReq) ToPlain() elemental.PlainIdentifiable {

	out := NewValidateRQLReq()
	if o.PolicyType != nil {
		out.PolicyType = *o.PolicyType
	}
	if o.PrismaId != nil {
		out.PrismaId = *o.PrismaId
	}
	if o.Query != nil {
		out.Query = *o.Query
	}
	if o.SearchType != nil {
		out.SearchType = *o.SearchType
	}

	return out
}

// DeepCopy returns a deep copy if the SparseValidateRQLReq.
func (o *SparseValidateRQLReq) DeepCopy() *SparseValidateRQLReq {

	if o == nil {
		return nil
	}

	out := &SparseValidateRQLReq{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseValidateRQLReq.
func (o *SparseValidateRQLReq) DeepCopyInto(out *SparseValidateRQLReq) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseValidateRQLReq: %s", err))
	}

	*out = *target.(*SparseValidateRQLReq)
}

type mongoAttributesValidateRQLReq struct {
}
type mongoAttributesSparseValidateRQLReq struct {
}
