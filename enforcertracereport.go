package gaia

import (
	"fmt"
	"sync"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// EnforcerTraceReportIdentity represents the Identity of the object.
var EnforcerTraceReportIdentity = elemental.Identity{
	Name:     "enforcertracereport",
	Category: "enforcertracereports",
	Package:  "zack",
	Private:  false,
}

// EnforcerTraceReportsList represents a list of EnforcerTraceReports
type EnforcerTraceReportsList []*EnforcerTraceReport

// Identity returns the identity of the objects in the list.
func (o EnforcerTraceReportsList) Identity() elemental.Identity {

	return EnforcerTraceReportIdentity
}

// Copy returns a pointer to a copy the EnforcerTraceReportsList.
func (o EnforcerTraceReportsList) Copy() elemental.Identifiables {

	copy := append(EnforcerTraceReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the EnforcerTraceReportsList.
func (o EnforcerTraceReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(EnforcerTraceReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*EnforcerTraceReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o EnforcerTraceReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EnforcerTraceReportsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the EnforcerTraceReportsList converted to SparseEnforcerTraceReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o EnforcerTraceReportsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o EnforcerTraceReportsList) Version() int {

	return 1
}

// EnforcerTraceReport represents the model of a enforcertracereport
type EnforcerTraceReport struct {
	// EnforcerID of the enforcer where the trace was collected.
	EnforcerID string `json:"enforcerID" bson:"enforcerid" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer where the trace was collected.
	EnforcerNamespace string `json:"enforcerNamespace" bson:"enforcernamespace" mapstructure:"enforcerNamespace,omitempty"`

	// Namespace of the PU where the trace was collected.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// ID of the pu where the trace was collected.
	PuID string `json:"puID" bson:"puid" mapstructure:"puID,omitempty"`

	// List of iptables trace records collected.
	Records []*TraceRecord `json:"-" bson:"records" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewEnforcerTraceReport returns a new *EnforcerTraceReport
func NewEnforcerTraceReport() *EnforcerTraceReport {

	return &EnforcerTraceReport{
		ModelVersion: 1,
		Mutex:        &sync.Mutex{},
		Records:      []*TraceRecord{},
	}
}

// Identity returns the Identity of the object.
func (o *EnforcerTraceReport) Identity() elemental.Identity {

	return EnforcerTraceReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EnforcerTraceReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EnforcerTraceReport) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *EnforcerTraceReport) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *EnforcerTraceReport) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *EnforcerTraceReport) Doc() string {

	return `Post a new enforcer trace that determines how packets are.`
}

func (o *EnforcerTraceReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *EnforcerTraceReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseEnforcerTraceReport{
			EnforcerID:        &o.EnforcerID,
			EnforcerNamespace: &o.EnforcerNamespace,
			Namespace:         &o.Namespace,
			PuID:              &o.PuID,
			Records:           &o.Records,
		}
	}

	sp := &SparseEnforcerTraceReport{}
	for _, f := range fields {
		switch f {
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "enforcerNamespace":
			sp.EnforcerNamespace = &(o.EnforcerNamespace)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "puID":
			sp.PuID = &(o.PuID)
		case "records":
			sp.Records = &(o.Records)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseEnforcerTraceReport to the object.
func (o *EnforcerTraceReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseEnforcerTraceReport)
	if so.EnforcerID != nil {
		o.EnforcerID = *so.EnforcerID
	}
	if so.EnforcerNamespace != nil {
		o.EnforcerNamespace = *so.EnforcerNamespace
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.PuID != nil {
		o.PuID = *so.PuID
	}
	if so.Records != nil {
		o.Records = *so.Records
	}
}

// DeepCopy returns a deep copy if the EnforcerTraceReport.
func (o *EnforcerTraceReport) DeepCopy() *EnforcerTraceReport {

	if o == nil {
		return nil
	}

	out := &EnforcerTraceReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *EnforcerTraceReport.
func (o *EnforcerTraceReport) DeepCopyInto(out *EnforcerTraceReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy EnforcerTraceReport: %s", err))
	}

	*out = *target.(*EnforcerTraceReport)
}

// Validate valides the current information stored into the structure.
func (o *EnforcerTraceReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("enforcerID", o.EnforcerID); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("enforcerNamespace", o.EnforcerNamespace); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("namespace", o.Namespace); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("puID", o.PuID); err != nil {
		requiredErrors = append(requiredErrors, err)
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
func (*EnforcerTraceReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := EnforcerTraceReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return EnforcerTraceReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*EnforcerTraceReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return EnforcerTraceReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *EnforcerTraceReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "enforcerID":
		return o.EnforcerID
	case "enforcerNamespace":
		return o.EnforcerNamespace
	case "namespace":
		return o.Namespace
	case "puID":
		return o.PuID
	case "records":
		return o.Records
	}

	return nil
}

// EnforcerTraceReportAttributesMap represents the map of attribute for EnforcerTraceReport.
var EnforcerTraceReportAttributesMap = map[string]elemental.AttributeSpecification{
	"EnforcerID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `EnforcerID of the enforcer where the trace was collected.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"EnforcerNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the enforcer where the trace was collected.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the PU where the trace was collected.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"PuID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PuID",
		Description:    `ID of the pu where the trace was collected.`,
		Exposed:        true,
		Name:           "puID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Records": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Records",
		Description:    `List of iptables trace records collected.`,
		Name:           "records",
		Stored:         true,
		SubType:        "tracerecord",
		Type:           "refList",
	},
}

// EnforcerTraceReportLowerCaseAttributesMap represents the map of attribute for EnforcerTraceReport.
var EnforcerTraceReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"enforcerid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `EnforcerID of the enforcer where the trace was collected.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"enforcernamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the enforcer where the trace was collected.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the PU where the trace was collected.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"puid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PuID",
		Description:    `ID of the pu where the trace was collected.`,
		Exposed:        true,
		Name:           "puID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"records": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Records",
		Description:    `List of iptables trace records collected.`,
		Name:           "records",
		Stored:         true,
		SubType:        "tracerecord",
		Type:           "refList",
	},
}

// SparseEnforcerTraceReportsList represents a list of SparseEnforcerTraceReports
type SparseEnforcerTraceReportsList []*SparseEnforcerTraceReport

// Identity returns the identity of the objects in the list.
func (o SparseEnforcerTraceReportsList) Identity() elemental.Identity {

	return EnforcerTraceReportIdentity
}

// Copy returns a pointer to a copy the SparseEnforcerTraceReportsList.
func (o SparseEnforcerTraceReportsList) Copy() elemental.Identifiables {

	copy := append(SparseEnforcerTraceReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseEnforcerTraceReportsList.
func (o SparseEnforcerTraceReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseEnforcerTraceReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseEnforcerTraceReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseEnforcerTraceReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseEnforcerTraceReportsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseEnforcerTraceReportsList converted to EnforcerTraceReportsList.
func (o SparseEnforcerTraceReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseEnforcerTraceReportsList) Version() int {

	return 1
}

// SparseEnforcerTraceReport represents the sparse version of a enforcertracereport.
type SparseEnforcerTraceReport struct {
	// EnforcerID of the enforcer where the trace was collected.
	EnforcerID *string `json:"enforcerID,omitempty" bson:"enforcerid,omitempty" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer where the trace was collected.
	EnforcerNamespace *string `json:"enforcerNamespace,omitempty" bson:"enforcernamespace,omitempty" mapstructure:"enforcerNamespace,omitempty"`

	// Namespace of the PU where the trace was collected.
	Namespace *string `json:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// ID of the pu where the trace was collected.
	PuID *string `json:"puID,omitempty" bson:"puid,omitempty" mapstructure:"puID,omitempty"`

	// List of iptables trace records collected.
	Records *[]*TraceRecord `json:"-" bson:"records,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewSparseEnforcerTraceReport returns a new  SparseEnforcerTraceReport.
func NewSparseEnforcerTraceReport() *SparseEnforcerTraceReport {
	return &SparseEnforcerTraceReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseEnforcerTraceReport) Identity() elemental.Identity {

	return EnforcerTraceReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseEnforcerTraceReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseEnforcerTraceReport) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseEnforcerTraceReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseEnforcerTraceReport) ToPlain() elemental.PlainIdentifiable {

	out := NewEnforcerTraceReport()
	if o.EnforcerID != nil {
		out.EnforcerID = *o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		out.EnforcerNamespace = *o.EnforcerNamespace
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.PuID != nil {
		out.PuID = *o.PuID
	}
	if o.Records != nil {
		out.Records = *o.Records
	}

	return out
}

// DeepCopy returns a deep copy if the SparseEnforcerTraceReport.
func (o *SparseEnforcerTraceReport) DeepCopy() *SparseEnforcerTraceReport {

	if o == nil {
		return nil
	}

	out := &SparseEnforcerTraceReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseEnforcerTraceReport.
func (o *SparseEnforcerTraceReport) DeepCopyInto(out *SparseEnforcerTraceReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseEnforcerTraceReport: %s", err))
	}

	*out = *target.(*SparseEnforcerTraceReport)
}
