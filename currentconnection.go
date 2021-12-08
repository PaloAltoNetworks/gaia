package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CurrentConnectionIdentity represents the Identity of the object.
var CurrentConnectionIdentity = elemental.Identity{
	Name:     "currentconnection",
	Category: "currentconnections",
	Package:  "guy",
	Private:  false,
}

// CurrentConnectionsList represents a list of CurrentConnections
type CurrentConnectionsList []*CurrentConnection

// Identity returns the identity of the objects in the list.
func (o CurrentConnectionsList) Identity() elemental.Identity {

	return CurrentConnectionIdentity
}

// Copy returns a pointer to a copy the CurrentConnectionsList.
func (o CurrentConnectionsList) Copy() elemental.Identifiables {

	copy := append(CurrentConnectionsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CurrentConnectionsList.
func (o CurrentConnectionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CurrentConnectionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CurrentConnection))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CurrentConnectionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CurrentConnectionsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CurrentConnectionsList converted to SparseCurrentConnectionsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CurrentConnectionsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCurrentConnectionsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCurrentConnection)
	}

	return out
}

// Version returns the version of the content.
func (o CurrentConnectionsList) Version() int {

	return 1
}

// CurrentConnection represents the model of a currentconnection
type CurrentConnection struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// The the duration of the tracked connection.
	Duration string `json:"duration,omitempty" msgpack:"duration,omitempty" bson:"-" mapstructure:"duration,omitempty"`

	// Was the connection existing when the enforcer started.
	Existing bool `json:"existing,omitempty" msgpack:"existing,omitempty" bson:"-" mapstructure:"existing,omitempty"`

	// The flow report for this connection.
	Flow *FlowReport `json:"flow,omitempty" msgpack:"flow,omitempty" bson:"-" mapstructure:"flow,omitempty"`

	// Port of the source.
	SourcePort int `json:"sourcePort,omitempty" msgpack:"sourcePort,omitempty" bson:"-" mapstructure:"sourcePort,omitempty"`

	// The time the enforcer started tracking the connection.
	StartTime time.Time `json:"startTime,omitempty" msgpack:"startTime,omitempty" bson:"-" mapstructure:"startTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCurrentConnection returns a new *CurrentConnection
func NewCurrentConnection() *CurrentConnection {

	return &CurrentConnection{
		ModelVersion: 1,
		Flow:         NewFlowReport(),
	}
}

// Identity returns the Identity of the object.
func (o *CurrentConnection) Identity() elemental.Identity {

	return CurrentConnectionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CurrentConnection) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CurrentConnection) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CurrentConnection) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCurrentConnection{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CurrentConnection) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCurrentConnection{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CurrentConnection) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CurrentConnection) BleveType() string {

	return "currentconnection"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CurrentConnection) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CurrentConnection) Doc() string {

	return `A current connection.`
}

func (o *CurrentConnection) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CurrentConnection) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCurrentConnection{
			ID:         &o.ID,
			Duration:   &o.Duration,
			Existing:   &o.Existing,
			Flow:       o.Flow,
			SourcePort: &o.SourcePort,
			StartTime:  &o.StartTime,
		}
	}

	sp := &SparseCurrentConnection{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "duration":
			sp.Duration = &(o.Duration)
		case "existing":
			sp.Existing = &(o.Existing)
		case "flow":
			sp.Flow = o.Flow
		case "sourcePort":
			sp.SourcePort = &(o.SourcePort)
		case "startTime":
			sp.StartTime = &(o.StartTime)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCurrentConnection to the object.
func (o *CurrentConnection) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCurrentConnection)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Duration != nil {
		o.Duration = *so.Duration
	}
	if so.Existing != nil {
		o.Existing = *so.Existing
	}
	if so.Flow != nil {
		o.Flow = so.Flow
	}
	if so.SourcePort != nil {
		o.SourcePort = *so.SourcePort
	}
	if so.StartTime != nil {
		o.StartTime = *so.StartTime
	}
}

// DeepCopy returns a deep copy if the CurrentConnection.
func (o *CurrentConnection) DeepCopy() *CurrentConnection {

	if o == nil {
		return nil
	}

	out := &CurrentConnection{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CurrentConnection.
func (o *CurrentConnection) DeepCopyInto(out *CurrentConnection) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CurrentConnection: %s", err))
	}

	*out = *target.(*CurrentConnection)
}

// Validate valides the current information stored into the structure.
func (o *CurrentConnection) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTimeDuration("duration", o.Duration); err != nil {
		errors = errors.Append(err)
	}

	if o.Flow != nil {
		elemental.ResetDefaultForZeroValues(o.Flow)
		if err := o.Flow.Validate(); err != nil {
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
func (*CurrentConnection) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CurrentConnectionAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CurrentConnectionLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CurrentConnection) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CurrentConnectionAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CurrentConnection) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "duration":
		return o.Duration
	case "existing":
		return o.Existing
	case "flow":
		return o.Flow
	case "sourcePort":
		return o.SourcePort
	case "startTime":
		return o.StartTime
	}

	return nil
}

// CurrentConnectionAttributesMap represents the map of attribute for CurrentConnection.
var CurrentConnectionAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Duration": {
		AllowedChoices: []string{},
		ConvertedName:  "Duration",
		Description:    `The the duration of the tracked connection.`,
		Exposed:        true,
		Name:           "duration",
		Type:           "string",
	},
	"Existing": {
		AllowedChoices: []string{},
		ConvertedName:  "Existing",
		Description:    `Was the connection existing when the enforcer started.`,
		Exposed:        true,
		Name:           "existing",
		Type:           "boolean",
	},
	"Flow": {
		AllowedChoices: []string{},
		ConvertedName:  "Flow",
		Description:    `The flow report for this connection.`,
		Exposed:        true,
		Name:           "flow",
		SubType:        "flowreport",
		Type:           "ref",
	},
	"SourcePort": {
		AllowedChoices: []string{},
		ConvertedName:  "SourcePort",
		Description:    `Port of the source.`,
		Exposed:        true,
		Name:           "sourcePort",
		Type:           "integer",
	},
	"StartTime": {
		AllowedChoices: []string{},
		ConvertedName:  "StartTime",
		Description:    `The time the enforcer started tracking the connection.`,
		Exposed:        true,
		Name:           "startTime",
		Type:           "time",
	},
}

// CurrentConnectionLowerCaseAttributesMap represents the map of attribute for CurrentConnection.
var CurrentConnectionLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"duration": {
		AllowedChoices: []string{},
		ConvertedName:  "Duration",
		Description:    `The the duration of the tracked connection.`,
		Exposed:        true,
		Name:           "duration",
		Type:           "string",
	},
	"existing": {
		AllowedChoices: []string{},
		ConvertedName:  "Existing",
		Description:    `Was the connection existing when the enforcer started.`,
		Exposed:        true,
		Name:           "existing",
		Type:           "boolean",
	},
	"flow": {
		AllowedChoices: []string{},
		ConvertedName:  "Flow",
		Description:    `The flow report for this connection.`,
		Exposed:        true,
		Name:           "flow",
		SubType:        "flowreport",
		Type:           "ref",
	},
	"sourceport": {
		AllowedChoices: []string{},
		ConvertedName:  "SourcePort",
		Description:    `Port of the source.`,
		Exposed:        true,
		Name:           "sourcePort",
		Type:           "integer",
	},
	"starttime": {
		AllowedChoices: []string{},
		ConvertedName:  "StartTime",
		Description:    `The time the enforcer started tracking the connection.`,
		Exposed:        true,
		Name:           "startTime",
		Type:           "time",
	},
}

// SparseCurrentConnectionsList represents a list of SparseCurrentConnections
type SparseCurrentConnectionsList []*SparseCurrentConnection

// Identity returns the identity of the objects in the list.
func (o SparseCurrentConnectionsList) Identity() elemental.Identity {

	return CurrentConnectionIdentity
}

// Copy returns a pointer to a copy the SparseCurrentConnectionsList.
func (o SparseCurrentConnectionsList) Copy() elemental.Identifiables {

	copy := append(SparseCurrentConnectionsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCurrentConnectionsList.
func (o SparseCurrentConnectionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCurrentConnectionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCurrentConnection))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCurrentConnectionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCurrentConnectionsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCurrentConnectionsList converted to CurrentConnectionsList.
func (o SparseCurrentConnectionsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCurrentConnectionsList) Version() int {

	return 1
}

// SparseCurrentConnection represents the sparse version of a currentconnection.
type SparseCurrentConnection struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// The the duration of the tracked connection.
	Duration *string `json:"duration,omitempty" msgpack:"duration,omitempty" bson:"-" mapstructure:"duration,omitempty"`

	// Was the connection existing when the enforcer started.
	Existing *bool `json:"existing,omitempty" msgpack:"existing,omitempty" bson:"-" mapstructure:"existing,omitempty"`

	// The flow report for this connection.
	Flow *FlowReport `json:"flow,omitempty" msgpack:"flow,omitempty" bson:"-" mapstructure:"flow,omitempty"`

	// Port of the source.
	SourcePort *int `json:"sourcePort,omitempty" msgpack:"sourcePort,omitempty" bson:"-" mapstructure:"sourcePort,omitempty"`

	// The time the enforcer started tracking the connection.
	StartTime *time.Time `json:"startTime,omitempty" msgpack:"startTime,omitempty" bson:"-" mapstructure:"startTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCurrentConnection returns a new  SparseCurrentConnection.
func NewSparseCurrentConnection() *SparseCurrentConnection {
	return &SparseCurrentConnection{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCurrentConnection) Identity() elemental.Identity {

	return CurrentConnectionIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCurrentConnection) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCurrentConnection) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCurrentConnection) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCurrentConnection{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCurrentConnection) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCurrentConnection{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCurrentConnection) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCurrentConnection) ToPlain() elemental.PlainIdentifiable {

	out := NewCurrentConnection()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Duration != nil {
		out.Duration = *o.Duration
	}
	if o.Existing != nil {
		out.Existing = *o.Existing
	}
	if o.Flow != nil {
		out.Flow = o.Flow
	}
	if o.SourcePort != nil {
		out.SourcePort = *o.SourcePort
	}
	if o.StartTime != nil {
		out.StartTime = *o.StartTime
	}

	return out
}

// DeepCopy returns a deep copy if the SparseCurrentConnection.
func (o *SparseCurrentConnection) DeepCopy() *SparseCurrentConnection {

	if o == nil {
		return nil
	}

	out := &SparseCurrentConnection{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCurrentConnection.
func (o *SparseCurrentConnection) DeepCopyInto(out *SparseCurrentConnection) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCurrentConnection: %s", err))
	}

	*out = *target.(*SparseCurrentConnection)
}

type mongoAttributesCurrentConnection struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
}
type mongoAttributesSparseCurrentConnection struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
}
