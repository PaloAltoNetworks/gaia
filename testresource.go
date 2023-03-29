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

// TestResourceIdentity represents the Identity of the object.
var TestResourceIdentity = elemental.Identity{
	Name:     "testresource",
	Category: "testresources",
	Package:  "basefetcher",
	Private:  true,
}

// TestResourcesList represents a list of TestResources
type TestResourcesList []*TestResource

// Identity returns the identity of the objects in the list.
func (o TestResourcesList) Identity() elemental.Identity {

	return TestResourceIdentity
}

// Copy returns a pointer to a copy the TestResourcesList.
func (o TestResourcesList) Copy() elemental.Identifiables {

	out := append(TestResourcesList{}, o...)
	return &out
}

// Append appends the objects to the a new copy of the TestResourcesList.
func (o TestResourcesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(TestResourcesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*TestResource))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TestResourcesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TestResourcesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the TestResourcesList converted to SparseTestResourcesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o TestResourcesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseTestResourcesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseTestResource)
	}

	return out
}

// Version returns the version of the content.
func (o TestResourcesList) Version() int {

	return 1
}

// TestResource represents the model of a testresource
type TestResource struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// The unique identifier of the resource.
	UID string `json:"UID" msgpack:"UID" bson:"uid" mapstructure:"UID,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// The raw data that represents the resource.
	RawData []byte `json:"rawData" msgpack:"rawData" bson:"rawdata" mapstructure:"rawData,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewTestResource returns a new *TestResource
func NewTestResource() *TestResource {

	return &TestResource{
		ModelVersion:  1,
		MigrationsLog: map[string]string{},
		RawData:       []byte{},
	}
}

// Identity returns the Identity of the object.
func (o *TestResource) Identity() elemental.Identity {

	return TestResourceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *TestResource) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *TestResource) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TestResource) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesTestResource{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.UID = o.UID
	s.CreateTime = o.CreateTime
	s.MigrationsLog = o.MigrationsLog
	s.Namespace = o.Namespace
	s.RawData = o.RawData
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TestResource) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesTestResource{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.UID = s.UID
	o.CreateTime = s.CreateTime
	o.MigrationsLog = s.MigrationsLog
	o.Namespace = s.Namespace
	o.RawData = s.RawData
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *TestResource) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *TestResource) BleveType() string {

	return "testresource"
}

// DefaultOrder returns the list of default ordering fields.
func (o *TestResource) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *TestResource) Doc() string {

	return `Represents a resource meant to test neocna basefetcher.`
}

func (o *TestResource) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *TestResource) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *TestResource) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *TestResource) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *TestResource) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *TestResource) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *TestResource) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *TestResource) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *TestResource) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *TestResource) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *TestResource) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *TestResource) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *TestResource) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *TestResource) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseTestResource{
			ID:            &o.ID,
			UID:           &o.UID,
			CreateTime:    &o.CreateTime,
			MigrationsLog: &o.MigrationsLog,
			Namespace:     &o.Namespace,
			RawData:       &o.RawData,
			UpdateTime:    &o.UpdateTime,
			ZHash:         &o.ZHash,
			Zone:          &o.Zone,
		}
	}

	sp := &SparseTestResource{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "UID":
			sp.UID = &(o.UID)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "rawData":
			sp.RawData = &(o.RawData)
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

// Patch apply the non nil value of a *SparseTestResource to the object.
func (o *TestResource) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseTestResource)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.UID != nil {
		o.UID = *so.UID
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.RawData != nil {
		o.RawData = *so.RawData
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

// DeepCopy returns a deep copy if the TestResource.
func (o *TestResource) DeepCopy() *TestResource {

	if o == nil {
		return nil
	}

	out := &TestResource{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *TestResource.
func (o *TestResource) DeepCopyInto(out *TestResource) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy TestResource: %s", err))
	}

	*out = *target.(*TestResource)
}

// Validate valides the current information stored into the structure.
func (o *TestResource) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("rawData", o.RawData); err != nil {
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
func (*TestResource) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TestResourceAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TestResourceLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*TestResource) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TestResourceAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *TestResource) ValueForAttribute(name string) any {

	switch name {
	case "ID":
		return o.ID
	case "UID":
		return o.UID
	case "createTime":
		return o.CreateTime
	case "migrationsLog":
		return o.MigrationsLog
	case "namespace":
		return o.Namespace
	case "rawData":
		return o.RawData
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// TestResourceAttributesMap represents the map of attribute for TestResource.
var TestResourceAttributesMap = map[string]elemental.AttributeSpecification{
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
	"UID": {
		AllowedChoices: []string{},
		BSONFieldName:  "uid",
		ConvertedName:  "UID",
		Description:    `The unique identifier of the resource.`,
		Exposed:        true,
		Name:           "UID",
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
	"MigrationsLog": {
		AllowedChoices: []string{},
		BSONFieldName:  "migrationslog",
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
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
	"RawData": {
		AllowedChoices: []string{},
		BSONFieldName:  "rawdata",
		ConvertedName:  "RawData",
		Description:    `The raw data that represents the resource.`,
		Exposed:        true,
		Name:           "rawData",
		Required:       true,
		Stored:         true,
		SubType:        "[]byte",
		Type:           "external",
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

// TestResourceLowerCaseAttributesMap represents the map of attribute for TestResource.
var TestResourceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"uid": {
		AllowedChoices: []string{},
		BSONFieldName:  "uid",
		ConvertedName:  "UID",
		Description:    `The unique identifier of the resource.`,
		Exposed:        true,
		Name:           "UID",
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
	"migrationslog": {
		AllowedChoices: []string{},
		BSONFieldName:  "migrationslog",
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
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
	"rawdata": {
		AllowedChoices: []string{},
		BSONFieldName:  "rawdata",
		ConvertedName:  "RawData",
		Description:    `The raw data that represents the resource.`,
		Exposed:        true,
		Name:           "rawData",
		Required:       true,
		Stored:         true,
		SubType:        "[]byte",
		Type:           "external",
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

// SparseTestResourcesList represents a list of SparseTestResources
type SparseTestResourcesList []*SparseTestResource

// Identity returns the identity of the objects in the list.
func (o SparseTestResourcesList) Identity() elemental.Identity {

	return TestResourceIdentity
}

// Copy returns a pointer to a copy the SparseTestResourcesList.
func (o SparseTestResourcesList) Copy() elemental.Identifiables {

	copy := append(SparseTestResourcesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseTestResourcesList.
func (o SparseTestResourcesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseTestResourcesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseTestResource))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseTestResourcesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseTestResourcesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseTestResourcesList converted to TestResourcesList.
func (o SparseTestResourcesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseTestResourcesList) Version() int {

	return 1
}

// SparseTestResource represents the sparse version of a testresource.
type SparseTestResource struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// The unique identifier of the resource.
	UID *string `json:"UID,omitempty" msgpack:"UID,omitempty" bson:"uid,omitempty" mapstructure:"UID,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// The raw data that represents the resource.
	RawData *[]byte `json:"rawData,omitempty" msgpack:"rawData,omitempty" bson:"rawdata,omitempty" mapstructure:"rawData,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseTestResource returns a new  SparseTestResource.
func NewSparseTestResource() *SparseTestResource {
	return &SparseTestResource{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseTestResource) Identity() elemental.Identity {

	return TestResourceIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseTestResource) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseTestResource) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTestResource) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseTestResource{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.UID != nil {
		s.UID = o.UID
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.RawData != nil {
		s.RawData = o.RawData
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
func (o *SparseTestResource) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseTestResource{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.UID != nil {
		o.UID = s.UID
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.RawData != nil {
		o.RawData = s.RawData
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
func (o *SparseTestResource) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseTestResource) ToPlain() elemental.PlainIdentifiable {

	out := NewTestResource()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.UID != nil {
		out.UID = *o.UID
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.RawData != nil {
		out.RawData = *o.RawData
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
func (o *SparseTestResource) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseTestResource) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseTestResource) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseTestResource) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseTestResource) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseTestResource) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseTestResource) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseTestResource) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseTestResource) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseTestResource) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseTestResource) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseTestResource) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseTestResource.
func (o *SparseTestResource) DeepCopy() *SparseTestResource {

	if o == nil {
		return nil
	}

	out := &SparseTestResource{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseTestResource.
func (o *SparseTestResource) DeepCopyInto(out *SparseTestResource) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseTestResource: %s", err))
	}

	*out = *target.(*SparseTestResource)
}

type mongoAttributesTestResource struct {
	ID            bson.ObjectId     `bson:"_id,omitempty"`
	UID           string            `bson:"uid"`
	CreateTime    time.Time         `bson:"createtime"`
	MigrationsLog map[string]string `bson:"migrationslog,omitempty"`
	Namespace     string            `bson:"namespace"`
	RawData       []byte            `bson:"rawdata"`
	UpdateTime    time.Time         `bson:"updatetime"`
	ZHash         int               `bson:"zhash"`
	Zone          int               `bson:"zone"`
}
type mongoAttributesSparseTestResource struct {
	ID            bson.ObjectId      `bson:"_id,omitempty"`
	UID           *string            `bson:"uid,omitempty"`
	CreateTime    *time.Time         `bson:"createtime,omitempty"`
	MigrationsLog *map[string]string `bson:"migrationslog,omitempty"`
	Namespace     *string            `bson:"namespace,omitempty"`
	RawData       *[]byte            `bson:"rawdata,omitempty"`
	UpdateTime    *time.Time         `bson:"updatetime,omitempty"`
	ZHash         *int               `bson:"zhash,omitempty"`
	Zone          *int               `bson:"zone,omitempty"`
}
