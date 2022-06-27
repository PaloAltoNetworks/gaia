package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// LogIdentity represents the Identity of the object.
var LogIdentity = elemental.Identity{
	Name:     "log",
	Category: "logs",
	Package:  "highwind",
	Private:  false,
}

// LogsList represents a list of Logs
type LogsList []*Log

// Identity returns the identity of the objects in the list.
func (o LogsList) Identity() elemental.Identity {

	return LogIdentity
}

// Copy returns a pointer to a copy the LogsList.
func (o LogsList) Copy() elemental.Identifiables {

	copy := append(LogsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the LogsList.
func (o LogsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(LogsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Log))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o LogsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o LogsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the LogsList converted to SparseLogsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o LogsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseLogsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseLog)
	}

	return out
}

// Version returns the version of the content.
func (o LogsList) Version() int {

	return 1
}

// Log represents the model of a log
type Log struct {
	// Contains all log data.
	Data map[string]string `json:"data" msgpack:"data" bson:"-" mapstructure:"data,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewLog returns a new *Log
func NewLog() *Log {

	return &Log{
		ModelVersion: 1,
		Data:         map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *Log) Identity() elemental.Identity {

	return LogIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Log) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Log) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Log) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesLog{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Log) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesLog{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Log) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Log) BleveType() string {

	return "log"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Log) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Log) Doc() string {

	return `Retrieves the logs of a deployed application.`
}

func (o *Log) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Log) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseLog{
			Data: &o.Data,
		}
	}

	sp := &SparseLog{}
	for _, f := range fields {
		switch f {
		case "data":
			sp.Data = &(o.Data)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseLog to the object.
func (o *Log) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseLog)
	if so.Data != nil {
		o.Data = *so.Data
	}
}

// DeepCopy returns a deep copy if the Log.
func (o *Log) DeepCopy() *Log {

	if o == nil {
		return nil
	}

	out := &Log{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Log.
func (o *Log) DeepCopyInto(out *Log) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Log: %s", err))
	}

	*out = *target.(*Log)
}

// Validate valides the current information stored into the structure.
func (o *Log) Validate() error {

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
func (*Log) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := LogAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return LogLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Log) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return LogAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Log) ValueForAttribute(name string) interface{} {

	switch name {
	case "data":
		return o.Data
	}

	return nil
}

// LogAttributesMap represents the map of attribute for Log.
var LogAttributesMap = map[string]elemental.AttributeSpecification{
	"Data": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Data",
		Description:    `Contains all log data.`,
		Exposed:        true,
		Name:           "data",
		ReadOnly:       true,
		SubType:        "map[string]string",
		Type:           "external",
	},
}

// LogLowerCaseAttributesMap represents the map of attribute for Log.
var LogLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"data": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Data",
		Description:    `Contains all log data.`,
		Exposed:        true,
		Name:           "data",
		ReadOnly:       true,
		SubType:        "map[string]string",
		Type:           "external",
	},
}

// SparseLogsList represents a list of SparseLogs
type SparseLogsList []*SparseLog

// Identity returns the identity of the objects in the list.
func (o SparseLogsList) Identity() elemental.Identity {

	return LogIdentity
}

// Copy returns a pointer to a copy the SparseLogsList.
func (o SparseLogsList) Copy() elemental.Identifiables {

	copy := append(SparseLogsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseLogsList.
func (o SparseLogsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseLogsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseLog))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseLogsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseLogsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseLogsList converted to LogsList.
func (o SparseLogsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseLogsList) Version() int {

	return 1
}

// SparseLog represents the sparse version of a log.
type SparseLog struct {
	// Contains all log data.
	Data *map[string]string `json:"data,omitempty" msgpack:"data,omitempty" bson:"-" mapstructure:"data,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseLog returns a new  SparseLog.
func NewSparseLog() *SparseLog {
	return &SparseLog{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseLog) Identity() elemental.Identity {

	return LogIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseLog) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseLog) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseLog) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseLog{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseLog) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseLog{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseLog) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseLog) ToPlain() elemental.PlainIdentifiable {

	out := NewLog()
	if o.Data != nil {
		out.Data = *o.Data
	}

	return out
}

// DeepCopy returns a deep copy if the SparseLog.
func (o *SparseLog) DeepCopy() *SparseLog {

	if o == nil {
		return nil
	}

	out := &SparseLog{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseLog.
func (o *SparseLog) DeepCopyInto(out *SparseLog) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseLog: %s", err))
	}

	*out = *target.(*SparseLog)
}

type mongoAttributesLog struct {
}
type mongoAttributesSparseLog struct {
}
