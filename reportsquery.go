package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ReportsQueryReportValue represents the possible values for attribute "report".
type ReportsQueryReportValue string

const (
	// ReportsQueryReportCounters represents the value Counters.
	ReportsQueryReportCounters ReportsQueryReportValue = "Counters"

	// ReportsQueryReportDNSLookups represents the value DNSLookups.
	ReportsQueryReportDNSLookups ReportsQueryReportValue = "DNSLookups"

	// ReportsQueryReportEnforcers represents the value Enforcers.
	ReportsQueryReportEnforcers ReportsQueryReportValue = "Enforcers"

	// ReportsQueryReportEventLogs represents the value EventLogs.
	ReportsQueryReportEventLogs ReportsQueryReportValue = "EventLogs"

	// ReportsQueryReportFlows represents the value Flows.
	ReportsQueryReportFlows ReportsQueryReportValue = "Flows"

	// ReportsQueryReportPackets represents the value Packets.
	ReportsQueryReportPackets ReportsQueryReportValue = "Packets"
)

// ReportsQueryIdentity represents the Identity of the object.
var ReportsQueryIdentity = elemental.Identity{
	Name:     "reportsquery",
	Category: "reportsqueries",
	Package:  "jenova",
	Private:  false,
}

// ReportsQueriesList represents a list of ReportsQueries
type ReportsQueriesList []*ReportsQuery

// Identity returns the identity of the objects in the list.
func (o ReportsQueriesList) Identity() elemental.Identity {

	return ReportsQueryIdentity
}

// Copy returns a pointer to a copy the ReportsQueriesList.
func (o ReportsQueriesList) Copy() elemental.Identifiables {

	copy := append(ReportsQueriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ReportsQueriesList.
func (o ReportsQueriesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ReportsQueriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ReportsQuery))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ReportsQueriesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ReportsQueriesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ReportsQueriesList converted to SparseReportsQueriesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ReportsQueriesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseReportsQueriesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseReportsQuery)
	}

	return out
}

// Version returns the version of the content.
func (o ReportsQueriesList) Version() int {

	return 1
}

// ReportsQuery represents the model of a reportsquery
type ReportsQuery struct {
	// List of DNSLookupReports.
	DNSLookupReports DNSLookupReportsList `json:"DNSLookupReports,omitempty" msgpack:"DNSLookupReports,omitempty" bson:"-" mapstructure:"DNSLookupReports,omitempty"`

	// List of CounterReports.
	CounterReports CounterReportsList `json:"counterReports,omitempty" msgpack:"counterReports,omitempty" bson:"-" mapstructure:"counterReports,omitempty"`

	// List of EnforcerReports.
	EnforcerReports EnforcerReportsList `json:"enforcerReports,omitempty" msgpack:"enforcerReports,omitempty" bson:"-" mapstructure:"enforcerReports,omitempty"`

	// List of EventLogs.
	EventLogs EventLogsList `json:"eventLogs,omitempty" msgpack:"eventLogs,omitempty" bson:"-" mapstructure:"eventLogs,omitempty"`

	// List of FlowReports.
	FlowReports FlowReportsList `json:"flowReports,omitempty" msgpack:"flowReports,omitempty" bson:"-" mapstructure:"flowReports,omitempty"`

	// List of PacketReports.
	PacketReports PacketReportsList `json:"packetReports,omitempty" msgpack:"packetReports,omitempty" bson:"-" mapstructure:"packetReports,omitempty"`

	// Name of the report type to query.
	Report ReportsQueryReportValue `json:"report" msgpack:"report" bson:"-" mapstructure:"report,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewReportsQuery returns a new *ReportsQuery
func NewReportsQuery() *ReportsQuery {

	return &ReportsQuery{
		ModelVersion:     1,
		CounterReports:   CounterReportsList{},
		DNSLookupReports: DNSLookupReportsList{},
		EnforcerReports:  EnforcerReportsList{},
		EventLogs:        EventLogsList{},
		FlowReports:      FlowReportsList{},
		PacketReports:    PacketReportsList{},
		Report:           ReportsQueryReportFlows,
	}
}

// Identity returns the Identity of the object.
func (o *ReportsQuery) Identity() elemental.Identity {

	return ReportsQueryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ReportsQuery) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ReportsQuery) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ReportsQuery) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesReportsQuery{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ReportsQuery) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesReportsQuery{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *ReportsQuery) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *ReportsQuery) BleveType() string {

	return "reportsquery"
}

// DefaultOrder returns the list of default ordering fields.
func (o *ReportsQuery) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *ReportsQuery) Doc() string {

	return `Supports querying Aporeto reports. All queries are protected within the
namespace of the user.`
}

func (o *ReportsQuery) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ReportsQuery) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseReportsQuery{
			DNSLookupReports: &o.DNSLookupReports,
			CounterReports:   &o.CounterReports,
			EnforcerReports:  &o.EnforcerReports,
			EventLogs:        &o.EventLogs,
			FlowReports:      &o.FlowReports,
			PacketReports:    &o.PacketReports,
			Report:           &o.Report,
		}
	}

	sp := &SparseReportsQuery{}
	for _, f := range fields {
		switch f {
		case "DNSLookupReports":
			sp.DNSLookupReports = &(o.DNSLookupReports)
		case "counterReports":
			sp.CounterReports = &(o.CounterReports)
		case "enforcerReports":
			sp.EnforcerReports = &(o.EnforcerReports)
		case "eventLogs":
			sp.EventLogs = &(o.EventLogs)
		case "flowReports":
			sp.FlowReports = &(o.FlowReports)
		case "packetReports":
			sp.PacketReports = &(o.PacketReports)
		case "report":
			sp.Report = &(o.Report)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseReportsQuery to the object.
func (o *ReportsQuery) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseReportsQuery)
	if so.DNSLookupReports != nil {
		o.DNSLookupReports = *so.DNSLookupReports
	}
	if so.CounterReports != nil {
		o.CounterReports = *so.CounterReports
	}
	if so.EnforcerReports != nil {
		o.EnforcerReports = *so.EnforcerReports
	}
	if so.EventLogs != nil {
		o.EventLogs = *so.EventLogs
	}
	if so.FlowReports != nil {
		o.FlowReports = *so.FlowReports
	}
	if so.PacketReports != nil {
		o.PacketReports = *so.PacketReports
	}
	if so.Report != nil {
		o.Report = *so.Report
	}
}

// DeepCopy returns a deep copy if the ReportsQuery.
func (o *ReportsQuery) DeepCopy() *ReportsQuery {

	if o == nil {
		return nil
	}

	out := &ReportsQuery{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ReportsQuery.
func (o *ReportsQuery) DeepCopyInto(out *ReportsQuery) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ReportsQuery: %s", err))
	}

	*out = *target.(*ReportsQuery)
}

// Validate valides the current information stored into the structure.
func (o *ReportsQuery) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.DNSLookupReports {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	for _, sub := range o.CounterReports {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	for _, sub := range o.EnforcerReports {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	for _, sub := range o.EventLogs {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	for _, sub := range o.FlowReports {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	for _, sub := range o.PacketReports {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateStringInList("report", string(o.Report), []string{"Flows", "Enforcers", "EventLogs", "Packets", "Counters", "DNSLookups"}, false); err != nil {
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
func (*ReportsQuery) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ReportsQueryAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ReportsQueryLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ReportsQuery) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ReportsQueryAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ReportsQuery) ValueForAttribute(name string) interface{} {

	switch name {
	case "DNSLookupReports":
		return o.DNSLookupReports
	case "counterReports":
		return o.CounterReports
	case "enforcerReports":
		return o.EnforcerReports
	case "eventLogs":
		return o.EventLogs
	case "flowReports":
		return o.FlowReports
	case "packetReports":
		return o.PacketReports
	case "report":
		return o.Report
	}

	return nil
}

// ReportsQueryAttributesMap represents the map of attribute for ReportsQuery.
var ReportsQueryAttributesMap = map[string]elemental.AttributeSpecification{
	"DNSLookupReports": {
		AllowedChoices: []string{},
		ConvertedName:  "DNSLookupReports",
		Description:    `List of DNSLookupReports.`,
		Exposed:        true,
		Name:           "DNSLookupReports",
		SubType:        "dnslookupreport",
		Type:           "refList",
	},
	"CounterReports": {
		AllowedChoices: []string{},
		ConvertedName:  "CounterReports",
		Description:    `List of CounterReports.`,
		Exposed:        true,
		Name:           "counterReports",
		SubType:        "counterreport",
		Type:           "refList",
	},
	"EnforcerReports": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerReports",
		Description:    `List of EnforcerReports.`,
		Exposed:        true,
		Name:           "enforcerReports",
		SubType:        "enforcerreport",
		Type:           "refList",
	},
	"EventLogs": {
		AllowedChoices: []string{},
		ConvertedName:  "EventLogs",
		Description:    `List of EventLogs.`,
		Exposed:        true,
		Name:           "eventLogs",
		SubType:        "eventlog",
		Type:           "refList",
	},
	"FlowReports": {
		AllowedChoices: []string{},
		ConvertedName:  "FlowReports",
		Description:    `List of FlowReports.`,
		Exposed:        true,
		Name:           "flowReports",
		SubType:        "flowreport",
		Type:           "refList",
	},
	"PacketReports": {
		AllowedChoices: []string{},
		ConvertedName:  "PacketReports",
		Description:    `List of PacketReports.`,
		Exposed:        true,
		Name:           "packetReports",
		SubType:        "packetreport",
		Type:           "refList",
	},
	"Report": {
		AllowedChoices: []string{"Flows", "Enforcers", "EventLogs", "Packets", "Counters", "DNSLookups"},
		ConvertedName:  "Report",
		DefaultValue:   ReportsQueryReportFlows,
		Description:    `Name of the report type to query.`,
		Exposed:        true,
		Name:           "report",
		Type:           "enum",
	},
}

// ReportsQueryLowerCaseAttributesMap represents the map of attribute for ReportsQuery.
var ReportsQueryLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"dnslookupreports": {
		AllowedChoices: []string{},
		ConvertedName:  "DNSLookupReports",
		Description:    `List of DNSLookupReports.`,
		Exposed:        true,
		Name:           "DNSLookupReports",
		SubType:        "dnslookupreport",
		Type:           "refList",
	},
	"counterreports": {
		AllowedChoices: []string{},
		ConvertedName:  "CounterReports",
		Description:    `List of CounterReports.`,
		Exposed:        true,
		Name:           "counterReports",
		SubType:        "counterreport",
		Type:           "refList",
	},
	"enforcerreports": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerReports",
		Description:    `List of EnforcerReports.`,
		Exposed:        true,
		Name:           "enforcerReports",
		SubType:        "enforcerreport",
		Type:           "refList",
	},
	"eventlogs": {
		AllowedChoices: []string{},
		ConvertedName:  "EventLogs",
		Description:    `List of EventLogs.`,
		Exposed:        true,
		Name:           "eventLogs",
		SubType:        "eventlog",
		Type:           "refList",
	},
	"flowreports": {
		AllowedChoices: []string{},
		ConvertedName:  "FlowReports",
		Description:    `List of FlowReports.`,
		Exposed:        true,
		Name:           "flowReports",
		SubType:        "flowreport",
		Type:           "refList",
	},
	"packetreports": {
		AllowedChoices: []string{},
		ConvertedName:  "PacketReports",
		Description:    `List of PacketReports.`,
		Exposed:        true,
		Name:           "packetReports",
		SubType:        "packetreport",
		Type:           "refList",
	},
	"report": {
		AllowedChoices: []string{"Flows", "Enforcers", "EventLogs", "Packets", "Counters", "DNSLookups"},
		ConvertedName:  "Report",
		DefaultValue:   ReportsQueryReportFlows,
		Description:    `Name of the report type to query.`,
		Exposed:        true,
		Name:           "report",
		Type:           "enum",
	},
}

// SparseReportsQueriesList represents a list of SparseReportsQueries
type SparseReportsQueriesList []*SparseReportsQuery

// Identity returns the identity of the objects in the list.
func (o SparseReportsQueriesList) Identity() elemental.Identity {

	return ReportsQueryIdentity
}

// Copy returns a pointer to a copy the SparseReportsQueriesList.
func (o SparseReportsQueriesList) Copy() elemental.Identifiables {

	copy := append(SparseReportsQueriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseReportsQueriesList.
func (o SparseReportsQueriesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseReportsQueriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseReportsQuery))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseReportsQueriesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseReportsQueriesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseReportsQueriesList converted to ReportsQueriesList.
func (o SparseReportsQueriesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseReportsQueriesList) Version() int {

	return 1
}

// SparseReportsQuery represents the sparse version of a reportsquery.
type SparseReportsQuery struct {
	// List of DNSLookupReports.
	DNSLookupReports *DNSLookupReportsList `json:"DNSLookupReports,omitempty" msgpack:"DNSLookupReports,omitempty" bson:"-" mapstructure:"DNSLookupReports,omitempty"`

	// List of CounterReports.
	CounterReports *CounterReportsList `json:"counterReports,omitempty" msgpack:"counterReports,omitempty" bson:"-" mapstructure:"counterReports,omitempty"`

	// List of EnforcerReports.
	EnforcerReports *EnforcerReportsList `json:"enforcerReports,omitempty" msgpack:"enforcerReports,omitempty" bson:"-" mapstructure:"enforcerReports,omitempty"`

	// List of EventLogs.
	EventLogs *EventLogsList `json:"eventLogs,omitempty" msgpack:"eventLogs,omitempty" bson:"-" mapstructure:"eventLogs,omitempty"`

	// List of FlowReports.
	FlowReports *FlowReportsList `json:"flowReports,omitempty" msgpack:"flowReports,omitempty" bson:"-" mapstructure:"flowReports,omitempty"`

	// List of PacketReports.
	PacketReports *PacketReportsList `json:"packetReports,omitempty" msgpack:"packetReports,omitempty" bson:"-" mapstructure:"packetReports,omitempty"`

	// Name of the report type to query.
	Report *ReportsQueryReportValue `json:"report,omitempty" msgpack:"report,omitempty" bson:"-" mapstructure:"report,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseReportsQuery returns a new  SparseReportsQuery.
func NewSparseReportsQuery() *SparseReportsQuery {
	return &SparseReportsQuery{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseReportsQuery) Identity() elemental.Identity {

	return ReportsQueryIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseReportsQuery) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseReportsQuery) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseReportsQuery) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseReportsQuery{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseReportsQuery) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseReportsQuery{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseReportsQuery) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseReportsQuery) ToPlain() elemental.PlainIdentifiable {

	out := NewReportsQuery()
	if o.DNSLookupReports != nil {
		out.DNSLookupReports = *o.DNSLookupReports
	}
	if o.CounterReports != nil {
		out.CounterReports = *o.CounterReports
	}
	if o.EnforcerReports != nil {
		out.EnforcerReports = *o.EnforcerReports
	}
	if o.EventLogs != nil {
		out.EventLogs = *o.EventLogs
	}
	if o.FlowReports != nil {
		out.FlowReports = *o.FlowReports
	}
	if o.PacketReports != nil {
		out.PacketReports = *o.PacketReports
	}
	if o.Report != nil {
		out.Report = *o.Report
	}

	return out
}

// DeepCopy returns a deep copy if the SparseReportsQuery.
func (o *SparseReportsQuery) DeepCopy() *SparseReportsQuery {

	if o == nil {
		return nil
	}

	out := &SparseReportsQuery{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseReportsQuery.
func (o *SparseReportsQuery) DeepCopyInto(out *SparseReportsQuery) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseReportsQuery: %s", err))
	}

	*out = *target.(*SparseReportsQuery)
}

type mongoAttributesReportsQuery struct {
}
type mongoAttributesSparseReportsQuery struct {
}
