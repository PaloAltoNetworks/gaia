package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// MetricsQueryRangeIdentity represents the Identity of the object.
var MetricsQueryRangeIdentity = elemental.Identity{
	Name:     "metricsqueryrange",
	Category: "metricsqueryrange",
	Package:  "jenova",
	Private:  false,
}

// MetricsQueryRangesList represents a list of MetricsQueryRanges
type MetricsQueryRangesList []*MetricsQueryRange

// Identity returns the identity of the objects in the list.
func (o MetricsQueryRangesList) Identity() elemental.Identity {

	return MetricsQueryRangeIdentity
}

// Copy returns a pointer to a copy the MetricsQueryRangesList.
func (o MetricsQueryRangesList) Copy() elemental.Identifiables {

	copy := append(MetricsQueryRangesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the MetricsQueryRangesList.
func (o MetricsQueryRangesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(MetricsQueryRangesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*MetricsQueryRange))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o MetricsQueryRangesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o MetricsQueryRangesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the MetricsQueryRangesList converted to SparseMetricsQueryRangesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o MetricsQueryRangesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseMetricsQueryRangesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseMetricsQueryRange)
	}

	return out
}

// Version returns the version of the content.
func (o MetricsQueryRangesList) Version() int {

	return 1
}

// MetricsQueryRange represents the model of a metricsqueryrange
type MetricsQueryRange struct {
	// End timestamp <rfc3339 | unix_timestamp>.
	End string `json:"end" msgpack:"end" bson:"-" mapstructure:"end,omitempty"`

	// Prometheus expression query string.
	Query string `json:"query" msgpack:"query" bson:"-" mapstructure:"query,omitempty"`

	// Start timestamp <rfc3339 | unix_timestamp>.
	Start string `json:"start" msgpack:"start" bson:"-" mapstructure:"start,omitempty"`

	// Query resolution step width in duration format or float number of seconds.
	Step string `json:"step" msgpack:"step" bson:"-" mapstructure:"step,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewMetricsQueryRange returns a new *MetricsQueryRange
func NewMetricsQueryRange() *MetricsQueryRange {

	return &MetricsQueryRange{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *MetricsQueryRange) Identity() elemental.Identity {

	return MetricsQueryRangeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *MetricsQueryRange) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *MetricsQueryRange) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *MetricsQueryRange) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesMetricsQueryRange{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *MetricsQueryRange) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesMetricsQueryRange{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *MetricsQueryRange) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *MetricsQueryRange) BleveType() string {

	return "metricsqueryrange"
}

// DefaultOrder returns the list of default ordering fields.
func (o *MetricsQueryRange) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *MetricsQueryRange) Doc() string {

	return `Prometheus compatible endpoint to evaluate an expression query over a range of
time. This can be used to retrieve back Aporeto specific metrics for a given
namespace. All queries are protected within the namespace of the caller.`
}

func (o *MetricsQueryRange) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *MetricsQueryRange) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseMetricsQueryRange{
			End:   &o.End,
			Query: &o.Query,
			Start: &o.Start,
			Step:  &o.Step,
		}
	}

	sp := &SparseMetricsQueryRange{}
	for _, f := range fields {
		switch f {
		case "end":
			sp.End = &(o.End)
		case "query":
			sp.Query = &(o.Query)
		case "start":
			sp.Start = &(o.Start)
		case "step":
			sp.Step = &(o.Step)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseMetricsQueryRange to the object.
func (o *MetricsQueryRange) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseMetricsQueryRange)
	if so.End != nil {
		o.End = *so.End
	}
	if so.Query != nil {
		o.Query = *so.Query
	}
	if so.Start != nil {
		o.Start = *so.Start
	}
	if so.Step != nil {
		o.Step = *so.Step
	}
}

// DeepCopy returns a deep copy if the MetricsQueryRange.
func (o *MetricsQueryRange) DeepCopy() *MetricsQueryRange {

	if o == nil {
		return nil
	}

	out := &MetricsQueryRange{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *MetricsQueryRange.
func (o *MetricsQueryRange) DeepCopyInto(out *MetricsQueryRange) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy MetricsQueryRange: %s", err))
	}

	*out = *target.(*MetricsQueryRange)
}

// Validate valides the current information stored into the structure.
func (o *MetricsQueryRange) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("query", o.Query); err != nil {
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
func (*MetricsQueryRange) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := MetricsQueryRangeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return MetricsQueryRangeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*MetricsQueryRange) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return MetricsQueryRangeAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *MetricsQueryRange) ValueForAttribute(name string) interface{} {

	switch name {
	case "end":
		return o.End
	case "query":
		return o.Query
	case "start":
		return o.Start
	case "step":
		return o.Step
	}

	return nil
}

// MetricsQueryRangeAttributesMap represents the map of attribute for MetricsQueryRange.
var MetricsQueryRangeAttributesMap = map[string]elemental.AttributeSpecification{
	"End": {
		AllowedChoices: []string{},
		ConvertedName:  "End",
		Description:    `End timestamp <rfc3339 | unix_timestamp>.`,
		Exposed:        true,
		Name:           "end",
		Type:           "string",
	},
	"Query": {
		AllowedChoices: []string{},
		ConvertedName:  "Query",
		Description:    `Prometheus expression query string.`,
		Exposed:        true,
		Name:           "query",
		Required:       true,
		Type:           "string",
	},
	"Start": {
		AllowedChoices: []string{},
		ConvertedName:  "Start",
		Description:    `Start timestamp <rfc3339 | unix_timestamp>.`,
		Exposed:        true,
		Name:           "start",
		Type:           "string",
	},
	"Step": {
		AllowedChoices: []string{},
		ConvertedName:  "Step",
		Description:    `Query resolution step width in duration format or float number of seconds.`,
		Exposed:        true,
		Name:           "step",
		Type:           "string",
	},
}

// MetricsQueryRangeLowerCaseAttributesMap represents the map of attribute for MetricsQueryRange.
var MetricsQueryRangeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"end": {
		AllowedChoices: []string{},
		ConvertedName:  "End",
		Description:    `End timestamp <rfc3339 | unix_timestamp>.`,
		Exposed:        true,
		Name:           "end",
		Type:           "string",
	},
	"query": {
		AllowedChoices: []string{},
		ConvertedName:  "Query",
		Description:    `Prometheus expression query string.`,
		Exposed:        true,
		Name:           "query",
		Required:       true,
		Type:           "string",
	},
	"start": {
		AllowedChoices: []string{},
		ConvertedName:  "Start",
		Description:    `Start timestamp <rfc3339 | unix_timestamp>.`,
		Exposed:        true,
		Name:           "start",
		Type:           "string",
	},
	"step": {
		AllowedChoices: []string{},
		ConvertedName:  "Step",
		Description:    `Query resolution step width in duration format or float number of seconds.`,
		Exposed:        true,
		Name:           "step",
		Type:           "string",
	},
}

// SparseMetricsQueryRangesList represents a list of SparseMetricsQueryRanges
type SparseMetricsQueryRangesList []*SparseMetricsQueryRange

// Identity returns the identity of the objects in the list.
func (o SparseMetricsQueryRangesList) Identity() elemental.Identity {

	return MetricsQueryRangeIdentity
}

// Copy returns a pointer to a copy the SparseMetricsQueryRangesList.
func (o SparseMetricsQueryRangesList) Copy() elemental.Identifiables {

	copy := append(SparseMetricsQueryRangesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseMetricsQueryRangesList.
func (o SparseMetricsQueryRangesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseMetricsQueryRangesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseMetricsQueryRange))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseMetricsQueryRangesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseMetricsQueryRangesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseMetricsQueryRangesList converted to MetricsQueryRangesList.
func (o SparseMetricsQueryRangesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseMetricsQueryRangesList) Version() int {

	return 1
}

// SparseMetricsQueryRange represents the sparse version of a metricsqueryrange.
type SparseMetricsQueryRange struct {
	// End timestamp <rfc3339 | unix_timestamp>.
	End *string `json:"end,omitempty" msgpack:"end,omitempty" bson:"-" mapstructure:"end,omitempty"`

	// Prometheus expression query string.
	Query *string `json:"query,omitempty" msgpack:"query,omitempty" bson:"-" mapstructure:"query,omitempty"`

	// Start timestamp <rfc3339 | unix_timestamp>.
	Start *string `json:"start,omitempty" msgpack:"start,omitempty" bson:"-" mapstructure:"start,omitempty"`

	// Query resolution step width in duration format or float number of seconds.
	Step *string `json:"step,omitempty" msgpack:"step,omitempty" bson:"-" mapstructure:"step,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseMetricsQueryRange returns a new  SparseMetricsQueryRange.
func NewSparseMetricsQueryRange() *SparseMetricsQueryRange {
	return &SparseMetricsQueryRange{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseMetricsQueryRange) Identity() elemental.Identity {

	return MetricsQueryRangeIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseMetricsQueryRange) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseMetricsQueryRange) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseMetricsQueryRange) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseMetricsQueryRange{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseMetricsQueryRange) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseMetricsQueryRange{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseMetricsQueryRange) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseMetricsQueryRange) ToPlain() elemental.PlainIdentifiable {

	out := NewMetricsQueryRange()
	if o.End != nil {
		out.End = *o.End
	}
	if o.Query != nil {
		out.Query = *o.Query
	}
	if o.Start != nil {
		out.Start = *o.Start
	}
	if o.Step != nil {
		out.Step = *o.Step
	}

	return out
}

// DeepCopy returns a deep copy if the SparseMetricsQueryRange.
func (o *SparseMetricsQueryRange) DeepCopy() *SparseMetricsQueryRange {

	if o == nil {
		return nil
	}

	out := &SparseMetricsQueryRange{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseMetricsQueryRange.
func (o *SparseMetricsQueryRange) DeepCopyInto(out *SparseMetricsQueryRange) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseMetricsQueryRange: %s", err))
	}

	*out = *target.(*SparseMetricsQueryRange)
}

type mongoAttributesMetricsQueryRange struct {
}
type mongoAttributesSparseMetricsQueryRange struct {
}
