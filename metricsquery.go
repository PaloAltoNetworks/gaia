package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// MetricsQueryIdentity represents the Identity of the object.
var MetricsQueryIdentity = elemental.Identity{
	Name:     "metricsquery",
	Category: "metricsquery",
	Package:  "jenova",
	Private:  false,
}

// MetricsQueriesList represents a list of MetricsQueries
type MetricsQueriesList []*MetricsQuery

// Identity returns the identity of the objects in the list.
func (o MetricsQueriesList) Identity() elemental.Identity {

	return MetricsQueryIdentity
}

// Copy returns a pointer to a copy the MetricsQueriesList.
func (o MetricsQueriesList) Copy() elemental.Identifiables {

	copy := append(MetricsQueriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the MetricsQueriesList.
func (o MetricsQueriesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(MetricsQueriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*MetricsQuery))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o MetricsQueriesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o MetricsQueriesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the MetricsQueriesList converted to SparseMetricsQueriesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o MetricsQueriesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseMetricsQueriesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseMetricsQuery)
	}

	return out
}

// Version returns the version of the content.
func (o MetricsQueriesList) Version() int {

	return 1
}

// MetricsQuery represents the model of a metricsquery
type MetricsQuery struct {
	// Prometheus expression query string.
	Query string `json:"query" msgpack:"query" bson:"-" mapstructure:"query,omitempty"`

	// Contains the result of the query.
	Result *MetricsQueryResult `json:"result" msgpack:"result" bson:"-" mapstructure:"result,omitempty"`

	// Evaluation timestamp <rfc3339 | unix_timestamp>.
	Time string `json:"time" msgpack:"time" bson:"-" mapstructure:"time,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewMetricsQuery returns a new *MetricsQuery
func NewMetricsQuery() *MetricsQuery {

	return &MetricsQuery{
		ModelVersion: 1,
		Result:       NewMetricsQueryResult(),
	}
}

// Identity returns the Identity of the object.
func (o *MetricsQuery) Identity() elemental.Identity {

	return MetricsQueryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *MetricsQuery) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *MetricsQuery) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *MetricsQuery) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesMetricsQuery{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *MetricsQuery) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesMetricsQuery{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *MetricsQuery) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *MetricsQuery) BleveType() string {

	return "metricsquery"
}

// DefaultOrder returns the list of default ordering fields.
func (o *MetricsQuery) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *MetricsQuery) Doc() string {

	return `Prometheus compatible endpoint to evaluate a Prometheus query expression at a
single instant or over a range of time. This can be used to retrieve back
Aporeto specific metrics for a given namespace. All queries are protected within
the namespace of the caller.`
}

func (o *MetricsQuery) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *MetricsQuery) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseMetricsQuery{
			Query:  &o.Query,
			Result: o.Result,
			Time:   &o.Time,
		}
	}

	sp := &SparseMetricsQuery{}
	for _, f := range fields {
		switch f {
		case "query":
			sp.Query = &(o.Query)
		case "result":
			sp.Result = o.Result
		case "time":
			sp.Time = &(o.Time)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseMetricsQuery to the object.
func (o *MetricsQuery) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseMetricsQuery)
	if so.Query != nil {
		o.Query = *so.Query
	}
	if so.Result != nil {
		o.Result = so.Result
	}
	if so.Time != nil {
		o.Time = *so.Time
	}
}

// DeepCopy returns a deep copy if the MetricsQuery.
func (o *MetricsQuery) DeepCopy() *MetricsQuery {

	if o == nil {
		return nil
	}

	out := &MetricsQuery{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *MetricsQuery.
func (o *MetricsQuery) DeepCopyInto(out *MetricsQuery) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy MetricsQuery: %s", err))
	}

	*out = *target.(*MetricsQuery)
}

// Validate valides the current information stored into the structure.
func (o *MetricsQuery) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("query", o.Query); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if o.Result != nil {
		elemental.ResetDefaultForZeroValues(o.Result)
		if err := o.Result.Validate(); err != nil {
			errors = errors.Append(err)
		}
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
func (*MetricsQuery) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := MetricsQueryAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return MetricsQueryLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*MetricsQuery) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return MetricsQueryAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *MetricsQuery) ValueForAttribute(name string) interface{} {

	switch name {
	case "query":
		return o.Query
	case "result":
		return o.Result
	case "time":
		return o.Time
	}

	return nil
}

// MetricsQueryAttributesMap represents the map of attribute for MetricsQuery.
var MetricsQueryAttributesMap = map[string]elemental.AttributeSpecification{
	"Query": {
		AllowedChoices: []string{},
		ConvertedName:  "Query",
		Description:    `Prometheus expression query string.`,
		Exposed:        true,
		Name:           "query",
		Required:       true,
		Type:           "string",
	},
	"Result": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Result",
		Description:    `Contains the result of the query.`,
		Exposed:        true,
		Name:           "result",
		ReadOnly:       true,
		SubType:        "metricsqueryresult",
		Type:           "ref",
	},
	"Time": {
		AllowedChoices: []string{},
		ConvertedName:  "Time",
		Description:    `Evaluation timestamp <rfc3339 | unix_timestamp>.`,
		Exposed:        true,
		Name:           "time",
		Type:           "string",
	},
}

// MetricsQueryLowerCaseAttributesMap represents the map of attribute for MetricsQuery.
var MetricsQueryLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"query": {
		AllowedChoices: []string{},
		ConvertedName:  "Query",
		Description:    `Prometheus expression query string.`,
		Exposed:        true,
		Name:           "query",
		Required:       true,
		Type:           "string",
	},
	"result": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Result",
		Description:    `Contains the result of the query.`,
		Exposed:        true,
		Name:           "result",
		ReadOnly:       true,
		SubType:        "metricsqueryresult",
		Type:           "ref",
	},
	"time": {
		AllowedChoices: []string{},
		ConvertedName:  "Time",
		Description:    `Evaluation timestamp <rfc3339 | unix_timestamp>.`,
		Exposed:        true,
		Name:           "time",
		Type:           "string",
	},
}

// SparseMetricsQueriesList represents a list of SparseMetricsQueries
type SparseMetricsQueriesList []*SparseMetricsQuery

// Identity returns the identity of the objects in the list.
func (o SparseMetricsQueriesList) Identity() elemental.Identity {

	return MetricsQueryIdentity
}

// Copy returns a pointer to a copy the SparseMetricsQueriesList.
func (o SparseMetricsQueriesList) Copy() elemental.Identifiables {

	copy := append(SparseMetricsQueriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseMetricsQueriesList.
func (o SparseMetricsQueriesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseMetricsQueriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseMetricsQuery))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseMetricsQueriesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseMetricsQueriesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseMetricsQueriesList converted to MetricsQueriesList.
func (o SparseMetricsQueriesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseMetricsQueriesList) Version() int {

	return 1
}

// SparseMetricsQuery represents the sparse version of a metricsquery.
type SparseMetricsQuery struct {
	// Prometheus expression query string.
	Query *string `json:"query,omitempty" msgpack:"query,omitempty" bson:"-" mapstructure:"query,omitempty"`

	// Contains the result of the query.
	Result *MetricsQueryResult `json:"result,omitempty" msgpack:"result,omitempty" bson:"-" mapstructure:"result,omitempty"`

	// Evaluation timestamp <rfc3339 | unix_timestamp>.
	Time *string `json:"time,omitempty" msgpack:"time,omitempty" bson:"-" mapstructure:"time,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseMetricsQuery returns a new  SparseMetricsQuery.
func NewSparseMetricsQuery() *SparseMetricsQuery {
	return &SparseMetricsQuery{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseMetricsQuery) Identity() elemental.Identity {

	return MetricsQueryIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseMetricsQuery) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseMetricsQuery) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseMetricsQuery) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseMetricsQuery{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseMetricsQuery) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseMetricsQuery{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseMetricsQuery) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseMetricsQuery) ToPlain() elemental.PlainIdentifiable {

	out := NewMetricsQuery()
	if o.Query != nil {
		out.Query = *o.Query
	}
	if o.Result != nil {
		out.Result = o.Result
	}
	if o.Time != nil {
		out.Time = *o.Time
	}

	return out
}

// DeepCopy returns a deep copy if the SparseMetricsQuery.
func (o *SparseMetricsQuery) DeepCopy() *SparseMetricsQuery {

	if o == nil {
		return nil
	}

	out := &SparseMetricsQuery{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseMetricsQuery.
func (o *SparseMetricsQuery) DeepCopyInto(out *SparseMetricsQuery) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseMetricsQuery: %s", err))
	}

	*out = *target.(*SparseMetricsQuery)
}

type mongoAttributesMetricsQuery struct {
}
type mongoAttributesSparseMetricsQuery struct {
}
