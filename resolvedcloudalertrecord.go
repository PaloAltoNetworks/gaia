// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ResolvedCloudAlertRecordIdentity represents the Identity of the object.
var ResolvedCloudAlertRecordIdentity = elemental.Identity{
	Name:     "resolvedcloudalertrecord",
	Category: "resolvedcloudalertrecords",
	Package:  "vargid",
	Private:  false,
}

// ResolvedCloudAlertRecordsList represents a list of ResolvedCloudAlertRecords
type ResolvedCloudAlertRecordsList []*ResolvedCloudAlertRecord

// Identity returns the identity of the objects in the list.
func (o ResolvedCloudAlertRecordsList) Identity() elemental.Identity {

	return ResolvedCloudAlertRecordIdentity
}

// Copy returns a pointer to a copy the ResolvedCloudAlertRecordsList.
func (o ResolvedCloudAlertRecordsList) Copy() elemental.Identifiables {

	out := append(ResolvedCloudAlertRecordsList{}, o...)
	return &out
}

// Append appends the objects to the a new copy of the ResolvedCloudAlertRecordsList.
func (o ResolvedCloudAlertRecordsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ResolvedCloudAlertRecordsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ResolvedCloudAlertRecord))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ResolvedCloudAlertRecordsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ResolvedCloudAlertRecordsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ResolvedCloudAlertRecordsList converted to SparseResolvedCloudAlertRecordsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ResolvedCloudAlertRecordsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseResolvedCloudAlertRecordsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseResolvedCloudAlertRecord)
	}

	return out
}

// Version returns the version of the content.
func (o ResolvedCloudAlertRecordsList) Version() int {

	return 1
}

// ResolvedCloudAlertRecord represents the model of a resolvedcloudalertrecord
type ResolvedCloudAlertRecord struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// The record that was raised and later resolved.
	Record *CloudAlertRecord `json:"record" msgpack:"record" bson:"record" mapstructure:"record,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewResolvedCloudAlertRecord returns a new *ResolvedCloudAlertRecord
func NewResolvedCloudAlertRecord() *ResolvedCloudAlertRecord {

	return &ResolvedCloudAlertRecord{
		ModelVersion: 1,
		Record:       NewCloudAlertRecord(),
	}
}

// Identity returns the Identity of the object.
func (o *ResolvedCloudAlertRecord) Identity() elemental.Identity {

	return ResolvedCloudAlertRecordIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ResolvedCloudAlertRecord) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ResolvedCloudAlertRecord) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ResolvedCloudAlertRecord) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesResolvedCloudAlertRecord{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.CreateTime = o.CreateTime
	s.Namespace = o.Namespace
	s.Record = o.Record
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ResolvedCloudAlertRecord) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesResolvedCloudAlertRecord{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.CreateTime = s.CreateTime
	o.Namespace = s.Namespace
	o.Record = s.Record
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *ResolvedCloudAlertRecord) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *ResolvedCloudAlertRecord) BleveType() string {

	return "resolvedcloudalertrecord"
}

// DefaultOrder returns the list of default ordering fields.
func (o *ResolvedCloudAlertRecord) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *ResolvedCloudAlertRecord) Doc() string {

	return `Represents an Alert that was raised and later resolved.`
}

func (o *ResolvedCloudAlertRecord) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *ResolvedCloudAlertRecord) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *ResolvedCloudAlertRecord) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetNamespace returns the Namespace of the receiver.
func (o *ResolvedCloudAlertRecord) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *ResolvedCloudAlertRecord) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *ResolvedCloudAlertRecord) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *ResolvedCloudAlertRecord) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *ResolvedCloudAlertRecord) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *ResolvedCloudAlertRecord) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *ResolvedCloudAlertRecord) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *ResolvedCloudAlertRecord) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ResolvedCloudAlertRecord) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseResolvedCloudAlertRecord{
			ID:         &o.ID,
			CreateTime: &o.CreateTime,
			Namespace:  &o.Namespace,
			Record:     o.Record,
			UpdateTime: &o.UpdateTime,
			ZHash:      &o.ZHash,
			Zone:       &o.Zone,
		}
	}

	sp := &SparseResolvedCloudAlertRecord{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "record":
			sp.Record = o.Record
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseResolvedCloudAlertRecord to the object.
func (o *ResolvedCloudAlertRecord) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseResolvedCloudAlertRecord)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.Record != nil {
		o.Record = so.Record
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the ResolvedCloudAlertRecord.
func (o *ResolvedCloudAlertRecord) DeepCopy() *ResolvedCloudAlertRecord {

	if o == nil {
		return nil
	}

	out := &ResolvedCloudAlertRecord{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ResolvedCloudAlertRecord.
func (o *ResolvedCloudAlertRecord) DeepCopyInto(out *ResolvedCloudAlertRecord) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ResolvedCloudAlertRecord: %s", err))
	}

	*out = *target.(*ResolvedCloudAlertRecord)
}

// Validate valides the current information stored into the structure.
func (o *ResolvedCloudAlertRecord) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if o.Record != nil {
		elemental.ResetDefaultForZeroValues(o.Record)
		if err := o.Record.Validate(); err != nil {
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
func (*ResolvedCloudAlertRecord) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ResolvedCloudAlertRecordAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ResolvedCloudAlertRecordLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ResolvedCloudAlertRecord) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ResolvedCloudAlertRecordAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ResolvedCloudAlertRecord) ValueForAttribute(name string) any {

	switch name {
	case "ID":
		return o.ID
	case "createTime":
		return o.CreateTime
	case "namespace":
		return o.Namespace
	case "record":
		return o.Record
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// ResolvedCloudAlertRecordAttributesMap represents the map of attribute for ResolvedCloudAlertRecord.
var ResolvedCloudAlertRecordAttributesMap = map[string]elemental.AttributeSpecification{
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
	"CreateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createtime",
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Record": {
		AllowedChoices: []string{},
		BSONFieldName:  "record",
		ConvertedName:  "Record",
		Description:    `The record that was raised and later resolved.`,
		Exposed:        true,
		Name:           "record",
		Stored:         true,
		SubType:        "cloudalertrecord",
		Type:           "ref",
	},
	"UpdateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updatetime",
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"ZHash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zhash",
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zone",
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// ResolvedCloudAlertRecordLowerCaseAttributesMap represents the map of attribute for ResolvedCloudAlertRecord.
var ResolvedCloudAlertRecordLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"createtime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createtime",
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"record": {
		AllowedChoices: []string{},
		BSONFieldName:  "record",
		ConvertedName:  "Record",
		Description:    `The record that was raised and later resolved.`,
		Exposed:        true,
		Name:           "record",
		Stored:         true,
		SubType:        "cloudalertrecord",
		Type:           "ref",
	},
	"updatetime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updatetime",
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"zhash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zhash",
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zone",
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SparseResolvedCloudAlertRecordsList represents a list of SparseResolvedCloudAlertRecords
type SparseResolvedCloudAlertRecordsList []*SparseResolvedCloudAlertRecord

// Identity returns the identity of the objects in the list.
func (o SparseResolvedCloudAlertRecordsList) Identity() elemental.Identity {

	return ResolvedCloudAlertRecordIdentity
}

// Copy returns a pointer to a copy the SparseResolvedCloudAlertRecordsList.
func (o SparseResolvedCloudAlertRecordsList) Copy() elemental.Identifiables {

	copy := append(SparseResolvedCloudAlertRecordsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseResolvedCloudAlertRecordsList.
func (o SparseResolvedCloudAlertRecordsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseResolvedCloudAlertRecordsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseResolvedCloudAlertRecord))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseResolvedCloudAlertRecordsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseResolvedCloudAlertRecordsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseResolvedCloudAlertRecordsList converted to ResolvedCloudAlertRecordsList.
func (o SparseResolvedCloudAlertRecordsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseResolvedCloudAlertRecordsList) Version() int {

	return 1
}

// SparseResolvedCloudAlertRecord represents the sparse version of a resolvedcloudalertrecord.
type SparseResolvedCloudAlertRecord struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// The record that was raised and later resolved.
	Record *CloudAlertRecord `json:"record,omitempty" msgpack:"record,omitempty" bson:"record,omitempty" mapstructure:"record,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseResolvedCloudAlertRecord returns a new  SparseResolvedCloudAlertRecord.
func NewSparseResolvedCloudAlertRecord() *SparseResolvedCloudAlertRecord {
	return &SparseResolvedCloudAlertRecord{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseResolvedCloudAlertRecord) Identity() elemental.Identity {

	return ResolvedCloudAlertRecordIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseResolvedCloudAlertRecord) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseResolvedCloudAlertRecord) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseResolvedCloudAlertRecord) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseResolvedCloudAlertRecord{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.Record != nil {
		s.Record = o.Record
	}
	if o.UpdateTime != nil {
		s.UpdateTime = o.UpdateTime
	}
	if o.ZHash != nil {
		s.ZHash = o.ZHash
	}
	if o.Zone != nil {
		s.Zone = o.Zone
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseResolvedCloudAlertRecord) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseResolvedCloudAlertRecord{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.Record != nil {
		o.Record = s.Record
	}
	if s.UpdateTime != nil {
		o.UpdateTime = s.UpdateTime
	}
	if s.ZHash != nil {
		o.ZHash = s.ZHash
	}
	if s.Zone != nil {
		o.Zone = s.Zone
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseResolvedCloudAlertRecord) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseResolvedCloudAlertRecord) ToPlain() elemental.PlainIdentifiable {

	out := NewResolvedCloudAlertRecord()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.Record != nil {
		out.Record = o.Record
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseResolvedCloudAlertRecord) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseResolvedCloudAlertRecord) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseResolvedCloudAlertRecord) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseResolvedCloudAlertRecord) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseResolvedCloudAlertRecord) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseResolvedCloudAlertRecord) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseResolvedCloudAlertRecord) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseResolvedCloudAlertRecord) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseResolvedCloudAlertRecord) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseResolvedCloudAlertRecord) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseResolvedCloudAlertRecord.
func (o *SparseResolvedCloudAlertRecord) DeepCopy() *SparseResolvedCloudAlertRecord {

	if o == nil {
		return nil
	}

	out := &SparseResolvedCloudAlertRecord{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseResolvedCloudAlertRecord.
func (o *SparseResolvedCloudAlertRecord) DeepCopyInto(out *SparseResolvedCloudAlertRecord) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseResolvedCloudAlertRecord: %s", err))
	}

	*out = *target.(*SparseResolvedCloudAlertRecord)
}

type mongoAttributesResolvedCloudAlertRecord struct {
	ID         bson.ObjectId     `bson:"_id,omitempty"`
	CreateTime time.Time         `bson:"createtime"`
	Namespace  string            `bson:"namespace"`
	Record     *CloudAlertRecord `bson:"record"`
	UpdateTime time.Time         `bson:"updatetime"`
	ZHash      int               `bson:"zhash"`
	Zone       int               `bson:"zone"`
}
type mongoAttributesSparseResolvedCloudAlertRecord struct {
	ID         bson.ObjectId     `bson:"_id,omitempty"`
	CreateTime *time.Time        `bson:"createtime,omitempty"`
	Namespace  *string           `bson:"namespace,omitempty"`
	Record     *CloudAlertRecord `bson:"record,omitempty"`
	UpdateTime *time.Time        `bson:"updatetime,omitempty"`
	ZHash      *int              `bson:"zhash,omitempty"`
	Zone       *int              `bson:"zone,omitempty"`
}
