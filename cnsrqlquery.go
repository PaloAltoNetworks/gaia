package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CNSrqlqueryIdentity represents the Identity of the object.
var CNSrqlqueryIdentity = elemental.Identity{
	Name:     "cnsrqlquery",
	Category: "cnsrqlquery",
	Package:  "vargid",
	Private:  false,
}

// CNSrqlqueriesList represents a list of CNSrqlqueries
type CNSrqlqueriesList []*CNSrqlquery

// Identity returns the identity of the objects in the list.
func (o CNSrqlqueriesList) Identity() elemental.Identity {

	return CNSrqlqueryIdentity
}

// Copy returns a pointer to a copy the CNSrqlqueriesList.
func (o CNSrqlqueriesList) Copy() elemental.Identifiables {

	copy := append(CNSrqlqueriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CNSrqlqueriesList.
func (o CNSrqlqueriesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CNSrqlqueriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CNSrqlquery))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CNSrqlqueriesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CNSrqlqueriesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CNSrqlqueriesList converted to SparseCNSrqlqueriesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CNSrqlqueriesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCNSrqlqueriesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCNSrqlquery)
	}

	return out
}

// Version returns the version of the content.
func (o CNSrqlqueriesList) Version() int {

	return 1
}

// CNSrqlquery represents the model of a cnsrqlquery
type CNSrqlquery struct {
	// ID of the alert request.
	AlertID string `json:"alertID" msgpack:"alertID" bson:"-" mapstructure:"alertID,omitempty"`

	// ID of the alert request.
	PolicyID string `json:"policyID" msgpack:"policyID" bson:"-" mapstructure:"policyID,omitempty"`

	// CNS rql query for alertID.
	Query string `json:"query" msgpack:"query" bson:"-" mapstructure:"query,omitempty"`

	// Whether the response to request is valid.
	Valid bool `json:"valid,omitempty" msgpack:"valid,omitempty" bson:"-" mapstructure:"valid,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCNSrqlquery returns a new *CNSrqlquery
func NewCNSrqlquery() *CNSrqlquery {

	return &CNSrqlquery{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *CNSrqlquery) Identity() elemental.Identity {

	return CNSrqlqueryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CNSrqlquery) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CNSrqlquery) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CNSrqlquery) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCNSrqlquery{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CNSrqlquery) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCNSrqlquery{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CNSrqlquery) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CNSrqlquery) BleveType() string {

	return "cnsrqlquery"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CNSrqlquery) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CNSrqlquery) Doc() string {

	return `A CNS endpoint which will generate the RQL query for the given alert-id.`
}

func (o *CNSrqlquery) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CNSrqlquery) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCNSrqlquery{
			AlertID:  &o.AlertID,
			PolicyID: &o.PolicyID,
			Query:    &o.Query,
			Valid:    &o.Valid,
		}
	}

	sp := &SparseCNSrqlquery{}
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

// Patch apply the non nil value of a *SparseCNSrqlquery to the object.
func (o *CNSrqlquery) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCNSrqlquery)
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

// DeepCopy returns a deep copy if the CNSrqlquery.
func (o *CNSrqlquery) DeepCopy() *CNSrqlquery {

	if o == nil {
		return nil
	}

	out := &CNSrqlquery{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CNSrqlquery.
func (o *CNSrqlquery) DeepCopyInto(out *CNSrqlquery) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CNSrqlquery: %s", err))
	}

	*out = *target.(*CNSrqlquery)
}

// Validate valides the current information stored into the structure.
func (o *CNSrqlquery) Validate() error {

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
func (*CNSrqlquery) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CNSrqlqueryAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CNSrqlqueryLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CNSrqlquery) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CNSrqlqueryAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CNSrqlquery) ValueForAttribute(name string) interface{} {

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

// CNSrqlqueryAttributesMap represents the map of attribute for CNSrqlquery.
var CNSrqlqueryAttributesMap = map[string]elemental.AttributeSpecification{
	"AlertID": {
		AllowedChoices: []string{},
		ConvertedName:  "AlertID",
		Description:    `ID of the alert request.`,
		Exposed:        true,
		Name:           "alertID",
		Type:           "string",
	},
	"PolicyID": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyID",
		Description:    `ID of the alert request.`,
		Exposed:        true,
		Name:           "policyID",
		Type:           "string",
	},
	"Query": {
		AllowedChoices: []string{},
		ConvertedName:  "Query",
		Description:    `CNS rql query for alertID.`,
		Exposed:        true,
		Name:           "query",
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

// CNSrqlqueryLowerCaseAttributesMap represents the map of attribute for CNSrqlquery.
var CNSrqlqueryLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"alertid": {
		AllowedChoices: []string{},
		ConvertedName:  "AlertID",
		Description:    `ID of the alert request.`,
		Exposed:        true,
		Name:           "alertID",
		Type:           "string",
	},
	"policyid": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyID",
		Description:    `ID of the alert request.`,
		Exposed:        true,
		Name:           "policyID",
		Type:           "string",
	},
	"query": {
		AllowedChoices: []string{},
		ConvertedName:  "Query",
		Description:    `CNS rql query for alertID.`,
		Exposed:        true,
		Name:           "query",
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

// SparseCNSrqlqueriesList represents a list of SparseCNSrqlqueries
type SparseCNSrqlqueriesList []*SparseCNSrqlquery

// Identity returns the identity of the objects in the list.
func (o SparseCNSrqlqueriesList) Identity() elemental.Identity {

	return CNSrqlqueryIdentity
}

// Copy returns a pointer to a copy the SparseCNSrqlqueriesList.
func (o SparseCNSrqlqueriesList) Copy() elemental.Identifiables {

	copy := append(SparseCNSrqlqueriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCNSrqlqueriesList.
func (o SparseCNSrqlqueriesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCNSrqlqueriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCNSrqlquery))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCNSrqlqueriesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCNSrqlqueriesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCNSrqlqueriesList converted to CNSrqlqueriesList.
func (o SparseCNSrqlqueriesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCNSrqlqueriesList) Version() int {

	return 1
}

// SparseCNSrqlquery represents the sparse version of a cnsrqlquery.
type SparseCNSrqlquery struct {
	// ID of the alert request.
	AlertID *string `json:"alertID,omitempty" msgpack:"alertID,omitempty" bson:"-" mapstructure:"alertID,omitempty"`

	// ID of the alert request.
	PolicyID *string `json:"policyID,omitempty" msgpack:"policyID,omitempty" bson:"-" mapstructure:"policyID,omitempty"`

	// CNS rql query for alertID.
	Query *string `json:"query,omitempty" msgpack:"query,omitempty" bson:"-" mapstructure:"query,omitempty"`

	// Whether the response to request is valid.
	Valid *bool `json:"valid,omitempty" msgpack:"valid,omitempty" bson:"-" mapstructure:"valid,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCNSrqlquery returns a new  SparseCNSrqlquery.
func NewSparseCNSrqlquery() *SparseCNSrqlquery {
	return &SparseCNSrqlquery{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCNSrqlquery) Identity() elemental.Identity {

	return CNSrqlqueryIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCNSrqlquery) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCNSrqlquery) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCNSrqlquery) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCNSrqlquery{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCNSrqlquery) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCNSrqlquery{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCNSrqlquery) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCNSrqlquery) ToPlain() elemental.PlainIdentifiable {

	out := NewCNSrqlquery()
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

// DeepCopy returns a deep copy if the SparseCNSrqlquery.
func (o *SparseCNSrqlquery) DeepCopy() *SparseCNSrqlquery {

	if o == nil {
		return nil
	}

	out := &SparseCNSrqlquery{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCNSrqlquery.
func (o *SparseCNSrqlquery) DeepCopyInto(out *SparseCNSrqlquery) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCNSrqlquery: %s", err))
	}

	*out = *target.(*SparseCNSrqlquery)
}

type mongoAttributesCNSrqlquery struct {
}
type mongoAttributesSparseCNSrqlquery struct {
}
