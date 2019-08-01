package gaia

import (
	"fmt"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// StatsInfoMeasurementValue represents the possible values for attribute "measurement".
type StatsInfoMeasurementValue string

const (
	// StatsInfoMeasurementAccesses represents the value Accesses.
	StatsInfoMeasurementAccesses StatsInfoMeasurementValue = "Accesses"

	// StatsInfoMeasurementAudit represents the value Audit.
	StatsInfoMeasurementAudit StatsInfoMeasurementValue = "Audit"

	// StatsInfoMeasurementCounters represents the value Counters.
	StatsInfoMeasurementCounters StatsInfoMeasurementValue = "Counters"

	// StatsInfoMeasurementEnforcers represents the value Enforcers.
	StatsInfoMeasurementEnforcers StatsInfoMeasurementValue = "Enforcers"

	// StatsInfoMeasurementEventLogs represents the value EventLogs.
	StatsInfoMeasurementEventLogs StatsInfoMeasurementValue = "EventLogs"

	// StatsInfoMeasurementFiles represents the value Files.
	StatsInfoMeasurementFiles StatsInfoMeasurementValue = "Files"

	// StatsInfoMeasurementFlows represents the value Flows.
	StatsInfoMeasurementFlows StatsInfoMeasurementValue = "Flows"

	// StatsInfoMeasurementPackets represents the value Packets.
	StatsInfoMeasurementPackets StatsInfoMeasurementValue = "Packets"
)

// StatsInfoIdentity represents the Identity of the object.
var StatsInfoIdentity = elemental.Identity{
	Name:     "statsinfo",
	Category: "statsinfo",
	Package:  "jenova",
	Private:  false,
}

// StatsInfosList represents a list of StatsInfos
type StatsInfosList []*StatsInfo

// Identity returns the identity of the objects in the list.
func (o StatsInfosList) Identity() elemental.Identity {

	return StatsInfoIdentity
}

// Copy returns a pointer to a copy the StatsInfosList.
func (o StatsInfosList) Copy() elemental.Identifiables {

	copy := append(StatsInfosList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the StatsInfosList.
func (o StatsInfosList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(StatsInfosList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*StatsInfo))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o StatsInfosList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o StatsInfosList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the StatsInfosList converted to SparseStatsInfosList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o StatsInfosList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseStatsInfosList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseStatsInfo)
	}

	return out
}

// Version returns the version of the content.
func (o StatsInfosList) Version() int {

	return 1
}

// StatsInfo represents the model of a statsinfo
type StatsInfo struct {
	// Contains the list of fields. You cannot group by these fields.
	Fields map[string]string `json:"fields" msgpack:"fields" bson:"-" mapstructure:"fields,omitempty"`

	// Name of the measurement.
	Measurement StatsInfoMeasurementValue `json:"measurement" msgpack:"measurement" bson:"-" mapstructure:"measurement,omitempty"`

	// Contains the list of tags. You can group by these tags.
	Tags []string `json:"tags" msgpack:"tags" bson:"-" mapstructure:"tags,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewStatsInfo returns a new *StatsInfo
func NewStatsInfo() *StatsInfo {

	return &StatsInfo{
		ModelVersion: 1,
		Fields:       map[string]string{},
		Measurement:  StatsInfoMeasurementFlows,
		Tags:         []string{},
	}
}

// Identity returns the Identity of the object.
func (o *StatsInfo) Identity() elemental.Identity {

	return StatsInfoIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *StatsInfo) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *StatsInfo) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *StatsInfo) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *StatsInfo) BleveType() string {

	return "statsinfo"
}

// DefaultOrder returns the list of default ordering fields.
func (o *StatsInfo) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *StatsInfo) Doc() string {

	return `Lists the fields and tags available in a statistics measurement.`
}

func (o *StatsInfo) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *StatsInfo) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseStatsInfo{
			Fields:      &o.Fields,
			Measurement: &o.Measurement,
			Tags:        &o.Tags,
		}
	}

	sp := &SparseStatsInfo{}
	for _, f := range fields {
		switch f {
		case "fields":
			sp.Fields = &(o.Fields)
		case "measurement":
			sp.Measurement = &(o.Measurement)
		case "tags":
			sp.Tags = &(o.Tags)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseStatsInfo to the object.
func (o *StatsInfo) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseStatsInfo)
	if so.Fields != nil {
		o.Fields = *so.Fields
	}
	if so.Measurement != nil {
		o.Measurement = *so.Measurement
	}
	if so.Tags != nil {
		o.Tags = *so.Tags
	}
}

// DeepCopy returns a deep copy if the StatsInfo.
func (o *StatsInfo) DeepCopy() *StatsInfo {

	if o == nil {
		return nil
	}

	out := &StatsInfo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *StatsInfo.
func (o *StatsInfo) DeepCopyInto(out *StatsInfo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy StatsInfo: %s", err))
	}

	*out = *target.(*StatsInfo)
}

// Validate valides the current information stored into the structure.
func (o *StatsInfo) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("measurement", string(o.Measurement), []string{"Flows", "Audit", "Enforcers", "Files", "EventLogs", "Counters", "Accesses", "Packets"}, false); err != nil {
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
func (*StatsInfo) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := StatsInfoAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return StatsInfoLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*StatsInfo) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return StatsInfoAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *StatsInfo) ValueForAttribute(name string) interface{} {

	switch name {
	case "fields":
		return o.Fields
	case "measurement":
		return o.Measurement
	case "tags":
		return o.Tags
	}

	return nil
}

// StatsInfoAttributesMap represents the map of attribute for StatsInfo.
var StatsInfoAttributesMap = map[string]elemental.AttributeSpecification{
	"Fields": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Fields",
		Description:    `Contains the list of fields. You cannot group by these fields.`,
		Exposed:        true,
		Name:           "fields",
		ReadOnly:       true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Measurement": elemental.AttributeSpecification{
		AllowedChoices: []string{"Flows", "Audit", "Enforcers", "Files", "EventLogs", "Counters", "Accesses", "Packets"},
		ConvertedName:  "Measurement",
		DefaultValue:   StatsInfoMeasurementFlows,
		Description:    `Name of the measurement.`,
		Exposed:        true,
		Name:           "measurement",
		Type:           "enum",
	},
	"Tags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Tags",
		Description:    `Contains the list of tags. You can group by these tags.`,
		Exposed:        true,
		Name:           "tags",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
}

// StatsInfoLowerCaseAttributesMap represents the map of attribute for StatsInfo.
var StatsInfoLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"fields": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Fields",
		Description:    `Contains the list of fields. You cannot group by these fields.`,
		Exposed:        true,
		Name:           "fields",
		ReadOnly:       true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"measurement": elemental.AttributeSpecification{
		AllowedChoices: []string{"Flows", "Audit", "Enforcers", "Files", "EventLogs", "Counters", "Accesses", "Packets"},
		ConvertedName:  "Measurement",
		DefaultValue:   StatsInfoMeasurementFlows,
		Description:    `Name of the measurement.`,
		Exposed:        true,
		Name:           "measurement",
		Type:           "enum",
	},
	"tags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Tags",
		Description:    `Contains the list of tags. You can group by these tags.`,
		Exposed:        true,
		Name:           "tags",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
}

// SparseStatsInfosList represents a list of SparseStatsInfos
type SparseStatsInfosList []*SparseStatsInfo

// Identity returns the identity of the objects in the list.
func (o SparseStatsInfosList) Identity() elemental.Identity {

	return StatsInfoIdentity
}

// Copy returns a pointer to a copy the SparseStatsInfosList.
func (o SparseStatsInfosList) Copy() elemental.Identifiables {

	copy := append(SparseStatsInfosList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseStatsInfosList.
func (o SparseStatsInfosList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseStatsInfosList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseStatsInfo))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseStatsInfosList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseStatsInfosList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseStatsInfosList converted to StatsInfosList.
func (o SparseStatsInfosList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseStatsInfosList) Version() int {

	return 1
}

// SparseStatsInfo represents the sparse version of a statsinfo.
type SparseStatsInfo struct {
	// Contains the list of fields. You cannot group by these fields.
	Fields *map[string]string `json:"fields,omitempty" msgpack:"fields,omitempty" bson:"-" mapstructure:"fields,omitempty"`

	// Name of the measurement.
	Measurement *StatsInfoMeasurementValue `json:"measurement,omitempty" msgpack:"measurement,omitempty" bson:"-" mapstructure:"measurement,omitempty"`

	// Contains the list of tags. You can group by these tags.
	Tags *[]string `json:"tags,omitempty" msgpack:"tags,omitempty" bson:"-" mapstructure:"tags,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseStatsInfo returns a new  SparseStatsInfo.
func NewSparseStatsInfo() *SparseStatsInfo {
	return &SparseStatsInfo{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseStatsInfo) Identity() elemental.Identity {

	return StatsInfoIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseStatsInfo) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseStatsInfo) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseStatsInfo) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseStatsInfo) ToPlain() elemental.PlainIdentifiable {

	out := NewStatsInfo()
	if o.Fields != nil {
		out.Fields = *o.Fields
	}
	if o.Measurement != nil {
		out.Measurement = *o.Measurement
	}
	if o.Tags != nil {
		out.Tags = *o.Tags
	}

	return out
}

// DeepCopy returns a deep copy if the SparseStatsInfo.
func (o *SparseStatsInfo) DeepCopy() *SparseStatsInfo {

	if o == nil {
		return nil
	}

	out := &SparseStatsInfo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseStatsInfo.
func (o *SparseStatsInfo) DeepCopyInto(out *SparseStatsInfo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseStatsInfo: %s", err))
	}

	*out = *target.(*SparseStatsInfo)
}
