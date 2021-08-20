package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CNSRQLQueryIdentity represents the Identity of the object.
var CNSRQLQueryIdentity = elemental.Identity{
	Name:     "cnsrqlquery",
	Category: "cnsrqlquery",
	Package:  "vargid",
	Private:  false,
}

// CNSRQLQueriesList represents a list of CNSRQLQueries
type CNSRQLQueriesList []*CNSRQLQuery

// Identity returns the identity of the objects in the list.
func (o CNSRQLQueriesList) Identity() elemental.Identity {

	return CNSRQLQueryIdentity
}

// Copy returns a pointer to a copy the CNSRQLQueriesList.
func (o CNSRQLQueriesList) Copy() elemental.Identifiables {

	copy := append(CNSRQLQueriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CNSRQLQueriesList.
func (o CNSRQLQueriesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CNSRQLQueriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CNSRQLQuery))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CNSRQLQueriesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CNSRQLQueriesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CNSRQLQueriesList converted to SparseCNSRQLQueriesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CNSRQLQueriesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCNSRQLQueriesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCNSRQLQuery)
	}

	return out
}

// Version returns the version of the content.
func (o CNSRQLQueriesList) Version() int {

	return 1
}

// CNSRQLQuery represents the model of a cnsrqlquery
type CNSRQLQuery struct {
	// ID of the query request.
	AlertID string `json:"alertID" msgpack:"alertID" bson:"-" mapstructure:"alertID,omitempty"`

	// The policy for which the alert was generated.
	PolicyID string `json:"policyID" msgpack:"policyID" bson:"-" mapstructure:"policyID,omitempty"`

	// The RQL query for the alert.
	Query string `json:"query" msgpack:"query" bson:"-" mapstructure:"query,omitempty"`

	// Whether the response to request is valid.
	Valid bool `json:"valid,omitempty" msgpack:"valid,omitempty" bson:"-" mapstructure:"valid,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCNSRQLQuery returns a new *CNSRQLQuery
func NewCNSRQLQuery() *CNSRQLQuery {

	return &CNSRQLQuery{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *CNSRQLQuery) Identity() elemental.Identity {

	return CNSRQLQueryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CNSRQLQuery) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CNSRQLQuery) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CNSRQLQuery) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCNSRQLQuery{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CNSRQLQuery) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCNSRQLQuery{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CNSRQLQuery) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CNSRQLQuery) BleveType() string {

	return "cnsrqlquery"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CNSRQLQuery) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CNSRQLQuery) Doc() string {

	return `A CNS endpoint which will generate the RQL query for the given alert-id.`
}

func (o *CNSRQLQuery) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CNSRQLQuery) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCNSRQLQuery{
			AlertID:  &o.AlertID,
			PolicyID: &o.PolicyID,
			Query:    &o.Query,
			Valid:    &o.Valid,
		}
	}

	sp := &SparseCNSRQLQuery{}
	for _, f := range fields {
		switch f {
		case "alertID":
			sp.AlertID = &(o.AlertID)
		case "policyID":
			sp.PolicyID = &(o.PolicyID)
		case "query":
			sp.Query = &(o.Query)
		case "valid":
			sp.Valid = &(o.Valid)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCNSRQLQuery to the object.
func (o *CNSRQLQuery) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCNSRQLQuery)
	if so.AlertID != nil {
		o.AlertID = *so.AlertID
	}
	if so.PolicyID != nil {
		o.PolicyID = *so.PolicyID
	}
	if so.Query != nil {
		o.Query = *so.Query
	}
	if so.Valid != nil {
		o.Valid = *so.Valid
	}
}

// DeepCopy returns a deep copy if the CNSRQLQuery.
func (o *CNSRQLQuery) DeepCopy() *CNSRQLQuery {

	if o == nil {
		return nil
	}

	out := &CNSRQLQuery{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CNSRQLQuery.
func (o *CNSRQLQuery) DeepCopyInto(out *CNSRQLQuery) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CNSRQLQuery: %s", err))
	}

	*out = *target.(*CNSRQLQuery)
}

// Validate valides the current information stored into the structure.
func (o *CNSRQLQuery) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("alertID", o.AlertID); err != nil {
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
func (*CNSRQLQuery) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CNSRQLQueryAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CNSRQLQueryLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CNSRQLQuery) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CNSRQLQueryAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CNSRQLQuery) ValueForAttribute(name string) interface{} {

	switch name {
	case "alertID":
		return o.AlertID
	case "policyID":
		return o.PolicyID
	case "query":
		return o.Query
	case "valid":
		return o.Valid
	}

	return nil
}

// CNSRQLQueryAttributesMap represents the map of attribute for CNSRQLQuery.
var CNSRQLQueryAttributesMap = map[string]elemental.AttributeSpecification{
	"AlertID": {
		AllowedChoices: []string{},
		ConvertedName:  "AlertID",
		Description:    `ID of the query request.`,
		Exposed:        true,
		Name:           "alertID",
		Required:       true,
		Type:           "string",
	},
	"PolicyID": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyID",
		Description:    `The policy for which the alert was generated.`,
		Exposed:        true,
		Name:           "policyID",
		Type:           "string",
	},
	"Query": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Query",
		Description:    `The RQL query for the alert.`,
		Exposed:        true,
		Name:           "query",
		ReadOnly:       true,
		Type:           "string",
	},
	"Valid": {
		AllowedChoices: []string{},
		ConvertedName:  "Valid",
		Description:    `Whether the response to request is valid.`,
		Exposed:        true,
		Name:           "valid",
		Type:           "boolean",
	},
}

// CNSRQLQueryLowerCaseAttributesMap represents the map of attribute for CNSRQLQuery.
var CNSRQLQueryLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"alertid": {
		AllowedChoices: []string{},
		ConvertedName:  "AlertID",
		Description:    `ID of the query request.`,
		Exposed:        true,
		Name:           "alertID",
		Required:       true,
		Type:           "string",
	},
	"policyid": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyID",
		Description:    `The policy for which the alert was generated.`,
		Exposed:        true,
		Name:           "policyID",
		Type:           "string",
	},
	"query": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Query",
		Description:    `The RQL query for the alert.`,
		Exposed:        true,
		Name:           "query",
		ReadOnly:       true,
		Type:           "string",
	},
	"valid": {
		AllowedChoices: []string{},
		ConvertedName:  "Valid",
		Description:    `Whether the response to request is valid.`,
		Exposed:        true,
		Name:           "valid",
		Type:           "boolean",
	},
}

// SparseCNSRQLQueriesList represents a list of SparseCNSRQLQueries
type SparseCNSRQLQueriesList []*SparseCNSRQLQuery

// Identity returns the identity of the objects in the list.
func (o SparseCNSRQLQueriesList) Identity() elemental.Identity {

	return CNSRQLQueryIdentity
}

// Copy returns a pointer to a copy the SparseCNSRQLQueriesList.
func (o SparseCNSRQLQueriesList) Copy() elemental.Identifiables {

	copy := append(SparseCNSRQLQueriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCNSRQLQueriesList.
func (o SparseCNSRQLQueriesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCNSRQLQueriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCNSRQLQuery))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCNSRQLQueriesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCNSRQLQueriesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCNSRQLQueriesList converted to CNSRQLQueriesList.
func (o SparseCNSRQLQueriesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCNSRQLQueriesList) Version() int {

	return 1
}

// SparseCNSRQLQuery represents the sparse version of a cnsrqlquery.
type SparseCNSRQLQuery struct {
	// ID of the query request.
	AlertID *string `json:"alertID,omitempty" msgpack:"alertID,omitempty" bson:"-" mapstructure:"alertID,omitempty"`

	// The policy for which the alert was generated.
	PolicyID *string `json:"policyID,omitempty" msgpack:"policyID,omitempty" bson:"-" mapstructure:"policyID,omitempty"`

	// The RQL query for the alert.
	Query *string `json:"query,omitempty" msgpack:"query,omitempty" bson:"-" mapstructure:"query,omitempty"`

	// Whether the response to request is valid.
	Valid *bool `json:"valid,omitempty" msgpack:"valid,omitempty" bson:"-" mapstructure:"valid,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCNSRQLQuery returns a new  SparseCNSRQLQuery.
func NewSparseCNSRQLQuery() *SparseCNSRQLQuery {
	return &SparseCNSRQLQuery{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCNSRQLQuery) Identity() elemental.Identity {

	return CNSRQLQueryIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCNSRQLQuery) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCNSRQLQuery) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCNSRQLQuery) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCNSRQLQuery{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCNSRQLQuery) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCNSRQLQuery{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCNSRQLQuery) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCNSRQLQuery) ToPlain() elemental.PlainIdentifiable {

	out := NewCNSRQLQuery()
	if o.AlertID != nil {
		out.AlertID = *o.AlertID
	}
	if o.PolicyID != nil {
		out.PolicyID = *o.PolicyID
	}
	if o.Query != nil {
		out.Query = *o.Query
	}
	if o.Valid != nil {
		out.Valid = *o.Valid
	}

	return out
}

// DeepCopy returns a deep copy if the SparseCNSRQLQuery.
func (o *SparseCNSRQLQuery) DeepCopy() *SparseCNSRQLQuery {

	if o == nil {
		return nil
	}

	out := &SparseCNSRQLQuery{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCNSRQLQuery.
func (o *SparseCNSRQLQuery) DeepCopyInto(out *SparseCNSRQLQuery) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCNSRQLQuery: %s", err))
	}

	*out = *target.(*SparseCNSRQLQuery)
}

type mongoAttributesCNSRQLQuery struct {
}
type mongoAttributesSparseCNSRQLQuery struct {
}
