package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ReportKindValue represents the possible values for attribute "kind".
type ReportKindValue string

const (
	// ReportKindAudit represents the value Audit.
	ReportKindAudit ReportKindValue = "Audit"

	// ReportKindClaims represents the value Claims.
	ReportKindClaims ReportKindValue = "Claims"

	// ReportKindEnforcer represents the value Enforcer.
	ReportKindEnforcer ReportKindValue = "Enforcer"

	// ReportKindFileAccess represents the value FileAccess.
	ReportKindFileAccess ReportKindValue = "FileAccess"

	// ReportKindFlow represents the value Flow.
	ReportKindFlow ReportKindValue = "Flow"

	// ReportKindProcessingUnit represents the value ProcessingUnit.
	ReportKindProcessingUnit ReportKindValue = "ProcessingUnit"

	// ReportKindSyscall represents the value Syscall.
	ReportKindSyscall ReportKindValue = "Syscall"
)

// ReportIdentity represents the Identity of the object.
var ReportIdentity = elemental.Identity{
	Name:     "report",
	Category: "reports",
	Package:  "zack",
	Private:  false,
}

// ReportsList represents a list of Reports
type ReportsList []*Report

// Identity returns the identity of the objects in the list.
func (o ReportsList) Identity() elemental.Identity {

	return ReportIdentity
}

// Copy returns a pointer to a copy the ReportsList.
func (o ReportsList) Copy() elemental.Identifiables {

	copy := append(ReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ReportsList.
func (o ReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Report))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ReportsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ReportsList converted to SparseReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ReportsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseReportsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseReport)
	}

	return out
}

// Version returns the version of the content.
func (o ReportsList) Version() int {

	return 1
}

// Report represents the model of a report
type Report struct {
	// TSDB Fields to set for the report.
	Fields map[string]interface{} `json:"fields" msgpack:"fields" bson:"-" mapstructure:"fields,omitempty"`

	// Kind contains the kind of report.
	Kind ReportKindValue `json:"kind" msgpack:"kind" bson:"-" mapstructure:"kind,omitempty"`

	// Tags contains the tags associated to the data point.
	Tags map[string]string `json:"tags" msgpack:"tags" bson:"-" mapstructure:"tags,omitempty"`

	// Timestamp contains the time for the report.
	Timestamp time.Time `json:"timestamp" msgpack:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	// Value contains the value for the report.
	Value float64 `json:"value" msgpack:"value" bson:"-" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewReport returns a new *Report
func NewReport() *Report {

	return &Report{
		ModelVersion: 1,
		Fields:       map[string]interface{}{},
		Tags:         map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *Report) Identity() elemental.Identity {

	return ReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Report) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Report) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Report) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesReport{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Report) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Report) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Report) BleveType() string {

	return "report"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Report) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Report) Doc() string {

	return `Post a new statistics report.`
}

func (o *Report) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Report) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseReport{
			Fields:    &o.Fields,
			Kind:      &o.Kind,
			Tags:      &o.Tags,
			Timestamp: &o.Timestamp,
			Value:     &o.Value,
		}
	}

	sp := &SparseReport{}
	for _, f := range fields {
		switch f {
		case "fields":
			sp.Fields = &(o.Fields)
		case "kind":
			sp.Kind = &(o.Kind)
		case "tags":
			sp.Tags = &(o.Tags)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		case "value":
			sp.Value = &(o.Value)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseReport to the object.
func (o *Report) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseReport)
	if so.Fields != nil {
		o.Fields = *so.Fields
	}
	if so.Kind != nil {
		o.Kind = *so.Kind
	}
	if so.Tags != nil {
		o.Tags = *so.Tags
	}
	if so.Timestamp != nil {
		o.Timestamp = *so.Timestamp
	}
	if so.Value != nil {
		o.Value = *so.Value
	}
}

// DeepCopy returns a deep copy if the Report.
func (o *Report) DeepCopy() *Report {

	if o == nil {
		return nil
	}

	out := &Report{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Report.
func (o *Report) DeepCopyInto(out *Report) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Report: %s", err))
	}

	*out = *target.(*Report)
}

// Validate valides the current information stored into the structure.
func (o *Report) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("kind", string(o.Kind), []string{"Audit", "Enforcer", "FileAccess", "Flow", "ProcessingUnit", "Syscall", "Claims"}, false); err != nil {
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
func (*Report) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Report) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Report) ValueForAttribute(name string) interface{} {

	switch name {
	case "fields":
		return o.Fields
	case "kind":
		return o.Kind
	case "tags":
		return o.Tags
	case "timestamp":
		return o.Timestamp
	case "value":
		return o.Value
	}

	return nil
}

// ReportAttributesMap represents the map of attribute for Report.
var ReportAttributesMap = map[string]elemental.AttributeSpecification{
	"Fields": {
		AllowedChoices: []string{},
		ConvertedName:  "Fields",
		Description:    `TSDB Fields to set for the report.`,
		Exposed:        true,
		Name:           "fields",
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"Kind": {
		AllowedChoices: []string{"Audit", "Enforcer", "FileAccess", "Flow", "ProcessingUnit", "Syscall", "Claims"},
		ConvertedName:  "Kind",
		Description:    `Kind contains the kind of report.`,
		Exposed:        true,
		Name:           "kind",
		Type:           "enum",
	},
	"Tags": {
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		Description:    `Tags contains the tags associated to the data point.`,
		Exposed:        true,
		Name:           "tags",
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Timestamp contains the time for the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
	"Value": {
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Value contains the value for the report.`,
		Exposed:        true,
		Name:           "value",
		Type:           "float",
	},
}

// ReportLowerCaseAttributesMap represents the map of attribute for Report.
var ReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"fields": {
		AllowedChoices: []string{},
		ConvertedName:  "Fields",
		Description:    `TSDB Fields to set for the report.`,
		Exposed:        true,
		Name:           "fields",
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"kind": {
		AllowedChoices: []string{"Audit", "Enforcer", "FileAccess", "Flow", "ProcessingUnit", "Syscall", "Claims"},
		ConvertedName:  "Kind",
		Description:    `Kind contains the kind of report.`,
		Exposed:        true,
		Name:           "kind",
		Type:           "enum",
	},
	"tags": {
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		Description:    `Tags contains the tags associated to the data point.`,
		Exposed:        true,
		Name:           "tags",
		SubType:        "map[string]string",
		Type:           "external",
	},
	"timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Timestamp contains the time for the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
	"value": {
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Value contains the value for the report.`,
		Exposed:        true,
		Name:           "value",
		Type:           "float",
	},
}

// SparseReportsList represents a list of SparseReports
type SparseReportsList []*SparseReport

// Identity returns the identity of the objects in the list.
func (o SparseReportsList) Identity() elemental.Identity {

	return ReportIdentity
}

// Copy returns a pointer to a copy the SparseReportsList.
func (o SparseReportsList) Copy() elemental.Identifiables {

	copy := append(SparseReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseReportsList.
func (o SparseReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseReportsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseReportsList converted to ReportsList.
func (o SparseReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseReportsList) Version() int {

	return 1
}

// SparseReport represents the sparse version of a report.
type SparseReport struct {
	// TSDB Fields to set for the report.
	Fields *map[string]interface{} `json:"fields,omitempty" msgpack:"fields,omitempty" bson:"-" mapstructure:"fields,omitempty"`

	// Kind contains the kind of report.
	Kind *ReportKindValue `json:"kind,omitempty" msgpack:"kind,omitempty" bson:"-" mapstructure:"kind,omitempty"`

	// Tags contains the tags associated to the data point.
	Tags *map[string]string `json:"tags,omitempty" msgpack:"tags,omitempty" bson:"-" mapstructure:"tags,omitempty"`

	// Timestamp contains the time for the report.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"-" mapstructure:"timestamp,omitempty"`

	// Value contains the value for the report.
	Value *float64 `json:"value,omitempty" msgpack:"value,omitempty" bson:"-" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseReport returns a new  SparseReport.
func NewSparseReport() *SparseReport {
	return &SparseReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseReport) Identity() elemental.Identity {

	return ReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseReport) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseReport{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseReport) ToPlain() elemental.PlainIdentifiable {

	out := NewReport()
	if o.Fields != nil {
		out.Fields = *o.Fields
	}
	if o.Kind != nil {
		out.Kind = *o.Kind
	}
	if o.Tags != nil {
		out.Tags = *o.Tags
	}
	if o.Timestamp != nil {
		out.Timestamp = *o.Timestamp
	}
	if o.Value != nil {
		out.Value = *o.Value
	}

	return out
}

// DeepCopy returns a deep copy if the SparseReport.
func (o *SparseReport) DeepCopy() *SparseReport {

	if o == nil {
		return nil
	}

	out := &SparseReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseReport.
func (o *SparseReport) DeepCopyInto(out *SparseReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseReport: %s", err))
	}

	*out = *target.(*SparseReport)
}

type mongoAttributesReport struct {
}
type mongoAttributesSparseReport struct {
}
