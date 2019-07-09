package gaia

import (
	"fmt"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// StatsQueryMeasurementValue represents the possible values for attribute "measurement".
type StatsQueryMeasurementValue string

const (
	// StatsQueryMeasurementAccesses represents the value Accesses.
	StatsQueryMeasurementAccesses StatsQueryMeasurementValue = "Accesses"

	// StatsQueryMeasurementAudit represents the value Audit.
	StatsQueryMeasurementAudit StatsQueryMeasurementValue = "Audit"

	// StatsQueryMeasurementEnforcerTraces represents the value EnforcerTraces.
	StatsQueryMeasurementEnforcerTraces StatsQueryMeasurementValue = "EnforcerTraces"

	// StatsQueryMeasurementEnforcers represents the value Enforcers.
	StatsQueryMeasurementEnforcers StatsQueryMeasurementValue = "Enforcers"

	// StatsQueryMeasurementEventLogs represents the value EventLogs.
	StatsQueryMeasurementEventLogs StatsQueryMeasurementValue = "EventLogs"

	// StatsQueryMeasurementFiles represents the value Files.
	StatsQueryMeasurementFiles StatsQueryMeasurementValue = "Files"

	// StatsQueryMeasurementFlows represents the value Flows.
	StatsQueryMeasurementFlows StatsQueryMeasurementValue = "Flows"

	// StatsQueryMeasurementPackets represents the value Packets.
	StatsQueryMeasurementPackets StatsQueryMeasurementValue = "Packets"
)

// StatsQueryIdentity represents the Identity of the object.
var StatsQueryIdentity = elemental.Identity{
	Name:     "statsquery",
	Category: "statsqueries",
	Package:  "jenova",
	Private:  false,
}

// StatsQueriesList represents a list of StatsQueries
type StatsQueriesList []*StatsQuery

// Identity returns the identity of the objects in the list.
func (o StatsQueriesList) Identity() elemental.Identity {

	return StatsQueryIdentity
}

// Copy returns a pointer to a copy the StatsQueriesList.
func (o StatsQueriesList) Copy() elemental.Identifiables {

	copy := append(StatsQueriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the StatsQueriesList.
func (o StatsQueriesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(StatsQueriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*StatsQuery))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o StatsQueriesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o StatsQueriesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the StatsQueriesList converted to SparseStatsQueriesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o StatsQueriesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseStatsQueriesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseStatsQuery)
	}

	return out
}

// Version returns the version of the content.
func (o StatsQueriesList) Version() int {

	return 1
}

// StatsQuery represents the model of a statsquery
type StatsQuery struct {
	// If set, the results will be order by time from the most recent to the oldest.
	Descending bool `json:"descending" msgpack:"descending" bson:"-" mapstructure:"descending,omitempty"`

	// List of fields to extract. If you don't pass anything, all available fields will
	// be returned. It is also possible to use a function like `sum(value)`.
	Fields []string `json:"fields" msgpack:"fields" bson:"-" mapstructure:"fields,omitempty"`

	// Apply a filter to the query.
	Filter string `json:"filter" msgpack:"filter" bson:"-" mapstructure:"filter,omitempty"`

	// Group results by the provided values. Note that not all fields can be used to
	// group the results.
	Groups []string `json:"groups" msgpack:"groups" bson:"-" mapstructure:"groups,omitempty"`

	// Limits the number of results. `-1` means no limit.
	Limit int `json:"limit" msgpack:"limit" bson:"-" mapstructure:"limit,omitempty"`

	// Name of the measurement.
	Measurement StatsQueryMeasurementValue `json:"measurement" msgpack:"measurement" bson:"-" mapstructure:"measurement,omitempty"`

	// Offsets the results. -1 means no offset.
	Offset int `json:"offset" msgpack:"offset" bson:"-" mapstructure:"offset,omitempty"`

	// Contains the result of the query.
	Results []*TimeSeriesQueryResults `json:"results" msgpack:"results" bson:"-" mapstructure:"results,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewStatsQuery returns a new *StatsQuery
func NewStatsQuery() *StatsQuery {

	return &StatsQuery{
		ModelVersion: 1,
		Fields:       []string{},
		Groups:       []string{},
		Limit:        -1,
		Measurement:  StatsQueryMeasurementFlows,
		Offset:       -1,
		Results:      []*TimeSeriesQueryResults{},
	}
}

// Identity returns the Identity of the object.
func (o *StatsQuery) Identity() elemental.Identity {

	return StatsQueryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *StatsQuery) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *StatsQuery) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *StatsQuery) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *StatsQuery) BleveType() string {

	return "statsquery"
}

// DefaultOrder returns the list of default ordering fields.
func (o *StatsQuery) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *StatsQuery) Doc() string {

	return `Retrieves time-series data stored by the Aporeto
system. Allows different types of queries that are all protected within
the namespace of the user.`
}

func (o *StatsQuery) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *StatsQuery) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseStatsQuery{
			Descending:  &o.Descending,
			Fields:      &o.Fields,
			Filter:      &o.Filter,
			Groups:      &o.Groups,
			Limit:       &o.Limit,
			Measurement: &o.Measurement,
			Offset:      &o.Offset,
			Results:     &o.Results,
		}
	}

	sp := &SparseStatsQuery{}
	for _, f := range fields {
		switch f {
		case "descending":
			sp.Descending = &(o.Descending)
		case "fields":
			sp.Fields = &(o.Fields)
		case "filter":
			sp.Filter = &(o.Filter)
		case "groups":
			sp.Groups = &(o.Groups)
		case "limit":
			sp.Limit = &(o.Limit)
		case "measurement":
			sp.Measurement = &(o.Measurement)
		case "offset":
			sp.Offset = &(o.Offset)
		case "results":
			sp.Results = &(o.Results)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseStatsQuery to the object.
func (o *StatsQuery) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseStatsQuery)
	if so.Descending != nil {
		o.Descending = *so.Descending
	}
	if so.Fields != nil {
		o.Fields = *so.Fields
	}
	if so.Filter != nil {
		o.Filter = *so.Filter
	}
	if so.Groups != nil {
		o.Groups = *so.Groups
	}
	if so.Limit != nil {
		o.Limit = *so.Limit
	}
	if so.Measurement != nil {
		o.Measurement = *so.Measurement
	}
	if so.Offset != nil {
		o.Offset = *so.Offset
	}
	if so.Results != nil {
		o.Results = *so.Results
	}
}

// DeepCopy returns a deep copy if the StatsQuery.
func (o *StatsQuery) DeepCopy() *StatsQuery {

	if o == nil {
		return nil
	}

	out := &StatsQuery{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *StatsQuery.
func (o *StatsQuery) DeepCopyInto(out *StatsQuery) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy StatsQuery: %s", err))
	}

	*out = *target.(*StatsQuery)
}

// Validate valides the current information stored into the structure.
func (o *StatsQuery) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("measurement", string(o.Measurement), []string{"Flows", "Audit", "Enforcers", "Files", "EventLogs", "Packets", "EnforcerTraces", "Accesses"}, false); err != nil {
		errors = errors.Append(err)
	}

	for _, sub := range o.Results {
		if sub == nil {
			continue
		}
		if err := sub.Validate(); err != nil {
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
func (*StatsQuery) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := StatsQueryAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return StatsQueryLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*StatsQuery) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return StatsQueryAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *StatsQuery) ValueForAttribute(name string) interface{} {

	switch name {
	case "descending":
		return o.Descending
	case "fields":
		return o.Fields
	case "filter":
		return o.Filter
	case "groups":
		return o.Groups
	case "limit":
		return o.Limit
	case "measurement":
		return o.Measurement
	case "offset":
		return o.Offset
	case "results":
		return o.Results
	}

	return nil
}

// StatsQueryAttributesMap represents the map of attribute for StatsQuery.
var StatsQueryAttributesMap = map[string]elemental.AttributeSpecification{
	"Descending": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Descending",
		Description:    `If set, the results will be order by time from the most recent to the oldest.`,
		Exposed:        true,
		Name:           "descending",
		Type:           "boolean",
	},
	"Fields": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fields",
		Description: `List of fields to extract. If you don't pass anything, all available fields will
be returned. It is also possible to use a function like ` + "`" + `sum(value)` + "`" + `.`,
		Exposed: true,
		Name:    "fields",
		SubType: "string",
		Type:    "list",
	},
	"Filter": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Filter",
		Description:    `Apply a filter to the query.`,
		Exposed:        true,
		Name:           "filter",
		Type:           "string",
	},
	"Groups": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Groups",
		Description: `Group results by the provided values. Note that not all fields can be used to
group the results.`,
		Exposed: true,
		Name:    "groups",
		SubType: "string",
		Type:    "list",
	},
	"Limit": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Limit",
		DefaultValue:   -1,
		Description:    `Limits the number of results. ` + "`" + `-1` + "`" + ` means no limit.`,
		Exposed:        true,
		Name:           "limit",
		Type:           "integer",
	},
	"Measurement": elemental.AttributeSpecification{
		AllowedChoices: []string{"Flows", "Audit", "Enforcers", "Files", "EventLogs", "Packets", "EnforcerTraces", "Accesses"},
		ConvertedName:  "Measurement",
		DefaultValue:   StatsQueryMeasurementFlows,
		Description:    `Name of the measurement.`,
		Exposed:        true,
		Name:           "measurement",
		Type:           "enum",
	},
	"Offset": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Offset",
		DefaultValue:   -1,
		Description:    `Offsets the results. -1 means no offset.`,
		Exposed:        true,
		Name:           "offset",
		Type:           "integer",
	},
	"Results": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Results",
		Description:    `Contains the result of the query.`,
		Exposed:        true,
		Name:           "results",
		ReadOnly:       true,
		SubType:        "timeseriesqueryresults",
		Type:           "refList",
	},
}

// StatsQueryLowerCaseAttributesMap represents the map of attribute for StatsQuery.
var StatsQueryLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"descending": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Descending",
		Description:    `If set, the results will be order by time from the most recent to the oldest.`,
		Exposed:        true,
		Name:           "descending",
		Type:           "boolean",
	},
	"fields": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fields",
		Description: `List of fields to extract. If you don't pass anything, all available fields will
be returned. It is also possible to use a function like ` + "`" + `sum(value)` + "`" + `.`,
		Exposed: true,
		Name:    "fields",
		SubType: "string",
		Type:    "list",
	},
	"filter": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Filter",
		Description:    `Apply a filter to the query.`,
		Exposed:        true,
		Name:           "filter",
		Type:           "string",
	},
	"groups": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Groups",
		Description: `Group results by the provided values. Note that not all fields can be used to
group the results.`,
		Exposed: true,
		Name:    "groups",
		SubType: "string",
		Type:    "list",
	},
	"limit": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Limit",
		DefaultValue:   -1,
		Description:    `Limits the number of results. ` + "`" + `-1` + "`" + ` means no limit.`,
		Exposed:        true,
		Name:           "limit",
		Type:           "integer",
	},
	"measurement": elemental.AttributeSpecification{
		AllowedChoices: []string{"Flows", "Audit", "Enforcers", "Files", "EventLogs", "Packets", "EnforcerTraces", "Accesses"},
		ConvertedName:  "Measurement",
		DefaultValue:   StatsQueryMeasurementFlows,
		Description:    `Name of the measurement.`,
		Exposed:        true,
		Name:           "measurement",
		Type:           "enum",
	},
	"offset": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Offset",
		DefaultValue:   -1,
		Description:    `Offsets the results. -1 means no offset.`,
		Exposed:        true,
		Name:           "offset",
		Type:           "integer",
	},
	"results": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Results",
		Description:    `Contains the result of the query.`,
		Exposed:        true,
		Name:           "results",
		ReadOnly:       true,
		SubType:        "timeseriesqueryresults",
		Type:           "refList",
	},
}

// SparseStatsQueriesList represents a list of SparseStatsQueries
type SparseStatsQueriesList []*SparseStatsQuery

// Identity returns the identity of the objects in the list.
func (o SparseStatsQueriesList) Identity() elemental.Identity {

	return StatsQueryIdentity
}

// Copy returns a pointer to a copy the SparseStatsQueriesList.
func (o SparseStatsQueriesList) Copy() elemental.Identifiables {

	copy := append(SparseStatsQueriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseStatsQueriesList.
func (o SparseStatsQueriesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseStatsQueriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseStatsQuery))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseStatsQueriesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseStatsQueriesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseStatsQueriesList converted to StatsQueriesList.
func (o SparseStatsQueriesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseStatsQueriesList) Version() int {

	return 1
}

// SparseStatsQuery represents the sparse version of a statsquery.
type SparseStatsQuery struct {
	// If set, the results will be order by time from the most recent to the oldest.
	Descending *bool `json:"descending,omitempty" msgpack:"descending,omitempty" bson:"-" mapstructure:"descending,omitempty"`

	// List of fields to extract. If you don't pass anything, all available fields will
	// be returned. It is also possible to use a function like `sum(value)`.
	Fields *[]string `json:"fields,omitempty" msgpack:"fields,omitempty" bson:"-" mapstructure:"fields,omitempty"`

	// Apply a filter to the query.
	Filter *string `json:"filter,omitempty" msgpack:"filter,omitempty" bson:"-" mapstructure:"filter,omitempty"`

	// Group results by the provided values. Note that not all fields can be used to
	// group the results.
	Groups *[]string `json:"groups,omitempty" msgpack:"groups,omitempty" bson:"-" mapstructure:"groups,omitempty"`

	// Limits the number of results. `-1` means no limit.
	Limit *int `json:"limit,omitempty" msgpack:"limit,omitempty" bson:"-" mapstructure:"limit,omitempty"`

	// Name of the measurement.
	Measurement *StatsQueryMeasurementValue `json:"measurement,omitempty" msgpack:"measurement,omitempty" bson:"-" mapstructure:"measurement,omitempty"`

	// Offsets the results. -1 means no offset.
	Offset *int `json:"offset,omitempty" msgpack:"offset,omitempty" bson:"-" mapstructure:"offset,omitempty"`

	// Contains the result of the query.
	Results *[]*TimeSeriesQueryResults `json:"results,omitempty" msgpack:"results,omitempty" bson:"-" mapstructure:"results,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseStatsQuery returns a new  SparseStatsQuery.
func NewSparseStatsQuery() *SparseStatsQuery {
	return &SparseStatsQuery{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseStatsQuery) Identity() elemental.Identity {

	return StatsQueryIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseStatsQuery) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseStatsQuery) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseStatsQuery) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseStatsQuery) ToPlain() elemental.PlainIdentifiable {

	out := NewStatsQuery()
	if o.Descending != nil {
		out.Descending = *o.Descending
	}
	if o.Fields != nil {
		out.Fields = *o.Fields
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Groups != nil {
		out.Groups = *o.Groups
	}
	if o.Limit != nil {
		out.Limit = *o.Limit
	}
	if o.Measurement != nil {
		out.Measurement = *o.Measurement
	}
	if o.Offset != nil {
		out.Offset = *o.Offset
	}
	if o.Results != nil {
		out.Results = *o.Results
	}

	return out
}

// DeepCopy returns a deep copy if the SparseStatsQuery.
func (o *SparseStatsQuery) DeepCopy() *SparseStatsQuery {

	if o == nil {
		return nil
	}

	out := &SparseStatsQuery{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseStatsQuery.
func (o *SparseStatsQuery) DeepCopyInto(out *SparseStatsQuery) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseStatsQuery: %s", err))
	}

	*out = *target.(*SparseStatsQuery)
}
