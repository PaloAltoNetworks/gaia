package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// SuspiciousActivityRiskValue represents the possible values for attribute "risk".
type SuspiciousActivityRiskValue string

const (
	// SuspiciousActivityRiskHigh represents the value High.
	SuspiciousActivityRiskHigh SuspiciousActivityRiskValue = "High"

	// SuspiciousActivityRiskLow represents the value Low.
	SuspiciousActivityRiskLow SuspiciousActivityRiskValue = "Low"

	// SuspiciousActivityRiskMedium represents the value Medium.
	SuspiciousActivityRiskMedium SuspiciousActivityRiskValue = "Medium"

	// SuspiciousActivityRiskUnknown represents the value Unknown.
	SuspiciousActivityRiskUnknown SuspiciousActivityRiskValue = "Unknown"
)

// SuspiciousActivityIdentity represents the Identity of the object.
var SuspiciousActivityIdentity = elemental.Identity{
	Name:     "suspiciousactivity",
	Category: "suspiciousactivities",
	Package:  "karl",
	Private:  false,
}

// SuspiciousActivitiesList represents a list of SuspiciousActivities
type SuspiciousActivitiesList []*SuspiciousActivity

// Identity returns the identity of the objects in the list.
func (o SuspiciousActivitiesList) Identity() elemental.Identity {

	return SuspiciousActivityIdentity
}

// Copy returns a pointer to a copy the SuspiciousActivitiesList.
func (o SuspiciousActivitiesList) Copy() elemental.Identifiables {

	copy := append(SuspiciousActivitiesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SuspiciousActivitiesList.
func (o SuspiciousActivitiesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SuspiciousActivitiesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SuspiciousActivity))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SuspiciousActivitiesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SuspiciousActivitiesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the SuspiciousActivitiesList converted to SparseSuspiciousActivitiesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o SuspiciousActivitiesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseSuspiciousActivitiesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseSuspiciousActivity)
	}

	return out
}

// Version returns the version of the content.
func (o SuspiciousActivitiesList) Version() int {

	return 1
}

// SuspiciousActivity represents the model of a suspiciousactivity
type SuspiciousActivity struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// The list of category names.
	Categories []int `json:"categories" msgpack:"categories" bson:"categories" mapstructure:"categories,omitempty"`

	// Number of times this suspicious activity is found.
	Counter int `json:"counter" msgpack:"counter" bson:"counter" mapstructure:"counter,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// The timestamp when the suspicious activity first occurred.
	FirstOccurrenceTime time.Time `json:"firstOccurrenceTime" msgpack:"firstOccurrenceTime" bson:"firstoccurrencetime" mapstructure:"firstOccurrenceTime,omitempty"`

	// The timestamp when the suspicious activity last occurred.
	LastOccurrenceTime time.Time `json:"lastOccurrenceTime" msgpack:"lastOccurrenceTime" bson:"lastoccurrencetime" mapstructure:"lastOccurrenceTime,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// The level of risk.
	Risk SuspiciousActivityRiskValue `json:"risk" msgpack:"risk" bson:"risk" mapstructure:"risk,omitempty"`

	// The identifier of the source.
	SourceID string `json:"sourceID" msgpack:"sourceID" bson:"sourceid" mapstructure:"sourceID,omitempty"`

	// The name of the source.
	SourceName string `json:"sourceName" msgpack:"sourceName" bson:"sourcename" mapstructure:"sourceName,omitempty"`

	// The namespace of the source.
	SourceNamespace string `json:"sourceNamespace" msgpack:"sourceNamespace" bson:"sourcenamespace" mapstructure:"sourceNamespace,omitempty"`

	// The type of the source.
	SourceType string `json:"sourceType" msgpack:"sourceType" bson:"sourcetype" mapstructure:"sourceType,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSuspiciousActivity returns a new *SuspiciousActivity
func NewSuspiciousActivity() *SuspiciousActivity {

	return &SuspiciousActivity{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		Categories:     []int{},
		NormalizedTags: []string{},
		Risk:           SuspiciousActivityRiskUnknown,
		MigrationsLog:  map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *SuspiciousActivity) Identity() elemental.Identity {

	return SuspiciousActivityIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SuspiciousActivity) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SuspiciousActivity) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SuspiciousActivity) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSuspiciousActivity{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.Categories = o.Categories
	s.Counter = o.Counter
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.FirstOccurrenceTime = o.FirstOccurrenceTime
	s.LastOccurrenceTime = o.LastOccurrenceTime
	s.MigrationsLog = o.MigrationsLog
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.Protected = o.Protected
	s.Risk = o.Risk
	s.SourceID = o.SourceID
	s.SourceName = o.SourceName
	s.SourceNamespace = o.SourceNamespace
	s.SourceType = o.SourceType
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SuspiciousActivity) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSuspiciousActivity{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.Categories = s.Categories
	o.Counter = s.Counter
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.FirstOccurrenceTime = s.FirstOccurrenceTime
	o.LastOccurrenceTime = s.LastOccurrenceTime
	o.MigrationsLog = s.MigrationsLog
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.Protected = s.Protected
	o.Risk = s.Risk
	o.SourceID = s.SourceID
	o.SourceName = s.SourceName
	o.SourceNamespace = s.SourceNamespace
	o.SourceType = s.SourceType
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SuspiciousActivity) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *SuspiciousActivity) BleveType() string {

	return "suspiciousactivity"
}

// DefaultOrder returns the list of default ordering fields.
func (o *SuspiciousActivity) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *SuspiciousActivity) Doc() string {

	return `Represents a suspicious activity found on the platform.`
}

func (o *SuspiciousActivity) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SuspiciousActivity) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *SuspiciousActivity) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SuspiciousActivity) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *SuspiciousActivity) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SuspiciousActivity) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *SuspiciousActivity) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SuspiciousActivity) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *SuspiciousActivity) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SuspiciousActivity) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *SuspiciousActivity) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *SuspiciousActivity) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *SuspiciousActivity) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SuspiciousActivity) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *SuspiciousActivity) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SuspiciousActivity) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *SuspiciousActivity) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SuspiciousActivity) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *SuspiciousActivity) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SuspiciousActivity) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *SuspiciousActivity) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SuspiciousActivity) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *SuspiciousActivity) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *SuspiciousActivity) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *SuspiciousActivity) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *SuspiciousActivity) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseSuspiciousActivity{
			ID:                   &o.ID,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			Categories:           &o.Categories,
			Counter:              &o.Counter,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			FirstOccurrenceTime:  &o.FirstOccurrenceTime,
			LastOccurrenceTime:   &o.LastOccurrenceTime,
			MigrationsLog:        &o.MigrationsLog,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Protected:            &o.Protected,
			Risk:                 &o.Risk,
			SourceID:             &o.SourceID,
			SourceName:           &o.SourceName,
			SourceNamespace:      &o.SourceNamespace,
			SourceType:           &o.SourceType,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseSuspiciousActivity{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "categories":
			sp.Categories = &(o.Categories)
		case "counter":
			sp.Counter = &(o.Counter)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "firstOccurrenceTime":
			sp.FirstOccurrenceTime = &(o.FirstOccurrenceTime)
		case "lastOccurrenceTime":
			sp.LastOccurrenceTime = &(o.LastOccurrenceTime)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "protected":
			sp.Protected = &(o.Protected)
		case "risk":
			sp.Risk = &(o.Risk)
		case "sourceID":
			sp.SourceID = &(o.SourceID)
		case "sourceName":
			sp.SourceName = &(o.SourceName)
		case "sourceNamespace":
			sp.SourceNamespace = &(o.SourceNamespace)
		case "sourceType":
			sp.SourceType = &(o.SourceType)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
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

// Patch apply the non nil value of a *SparseSuspiciousActivity to the object.
func (o *SuspiciousActivity) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseSuspiciousActivity)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.Categories != nil {
		o.Categories = *so.Categories
	}
	if so.Counter != nil {
		o.Counter = *so.Counter
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.FirstOccurrenceTime != nil {
		o.FirstOccurrenceTime = *so.FirstOccurrenceTime
	}
	if so.LastOccurrenceTime != nil {
		o.LastOccurrenceTime = *so.LastOccurrenceTime
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Risk != nil {
		o.Risk = *so.Risk
	}
	if so.SourceID != nil {
		o.SourceID = *so.SourceID
	}
	if so.SourceName != nil {
		o.SourceName = *so.SourceName
	}
	if so.SourceNamespace != nil {
		o.SourceNamespace = *so.SourceNamespace
	}
	if so.SourceType != nil {
		o.SourceType = *so.SourceType
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
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

// DeepCopy returns a deep copy if the SuspiciousActivity.
func (o *SuspiciousActivity) DeepCopy() *SuspiciousActivity {

	if o == nil {
		return nil
	}

	out := &SuspiciousActivity{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SuspiciousActivity.
func (o *SuspiciousActivity) DeepCopyInto(out *SuspiciousActivity) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SuspiciousActivity: %s", err))
	}

	*out = *target.(*SuspiciousActivity)
}

// Validate valides the current information stored into the structure.
func (o *SuspiciousActivity) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredExternal("categories", o.Categories); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredTime("firstOccurrenceTime", o.FirstOccurrenceTime); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredTime("lastOccurrenceTime", o.LastOccurrenceTime); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("risk", string(o.Risk), []string{"Low", "Medium", "High", "Unknown"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("sourceID", o.SourceID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("sourceName", o.SourceName); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("sourceNamespace", o.SourceNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("sourceType", o.SourceType); err != nil {
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
func (*SuspiciousActivity) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := SuspiciousActivityAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return SuspiciousActivityLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*SuspiciousActivity) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return SuspiciousActivityAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *SuspiciousActivity) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "categories":
		return o.Categories
	case "counter":
		return o.Counter
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "firstOccurrenceTime":
		return o.FirstOccurrenceTime
	case "lastOccurrenceTime":
		return o.LastOccurrenceTime
	case "migrationsLog":
		return o.MigrationsLog
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "protected":
		return o.Protected
	case "risk":
		return o.Risk
	case "sourceID":
		return o.SourceID
	case "sourceName":
		return o.SourceName
	case "sourceNamespace":
		return o.SourceNamespace
	case "sourceType":
		return o.SourceType
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// SuspiciousActivityAttributesMap represents the map of attribute for SuspiciousActivity.
var SuspiciousActivityAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Annotations": {
		AllowedChoices: []string{},
		BSONFieldName:  "annotations",
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"AssociatedTags": {
		AllowedChoices: []string{},
		BSONFieldName:  "associatedtags",
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Categories": {
		AllowedChoices: []string{},
		BSONFieldName:  "categories",
		ConvertedName:  "Categories",
		Description:    `The list of category names.`,
		Exposed:        true,
		Name:           "categories",
		Required:       true,
		Stored:         true,
		SubType:        "integer",
		Type:           "list",
	},
	"Counter": {
		AllowedChoices: []string{},
		BSONFieldName:  "counter",
		ConvertedName:  "Counter",
		Description:    `Number of times this suspicious activity is found.`,
		Exposed:        true,
		Name:           "counter",
		Stored:         true,
		Type:           "integer",
	},
	"CreateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createidempotencykey",
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
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
	"FirstOccurrenceTime": {
		AllowedChoices: []string{},
		BSONFieldName:  "firstoccurrencetime",
		ConvertedName:  "FirstOccurrenceTime",
		Description:    `The timestamp when the suspicious activity first occurred.`,
		Exposed:        true,
		Name:           "firstOccurrenceTime",
		Required:       true,
		Stored:         true,
		Type:           "time",
	},
	"LastOccurrenceTime": {
		AllowedChoices: []string{},
		BSONFieldName:  "lastoccurrencetime",
		ConvertedName:  "LastOccurrenceTime",
		Description:    `The timestamp when the suspicious activity last occurred.`,
		Exposed:        true,
		Name:           "lastOccurrenceTime",
		Required:       true,
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
	"NormalizedTags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "normalizedtags",
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"Protected": {
		AllowedChoices: []string{},
		BSONFieldName:  "protected",
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Risk": {
		AllowedChoices: []string{"Low", "Medium", "High", "Unknown"},
		BSONFieldName:  "risk",
		ConvertedName:  "Risk",
		DefaultValue:   SuspiciousActivityRiskUnknown,
		Description:    `The level of risk.`,
		Exposed:        true,
		Name:           "risk",
		Stored:         true,
		Type:           "enum",
	},
	"SourceID": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourceid",
		ConvertedName:  "SourceID",
		CreationOnly:   true,
		Description:    `The identifier of the source.`,
		Exposed:        true,
		Name:           "sourceID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"SourceName": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourcename",
		ConvertedName:  "SourceName",
		CreationOnly:   true,
		Description:    `The name of the source.`,
		Exposed:        true,
		Name:           "sourceName",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"SourceNamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourcenamespace",
		ConvertedName:  "SourceNamespace",
		CreationOnly:   true,
		Description:    `The namespace of the source.`,
		Exposed:        true,
		Name:           "sourceNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"SourceType": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourcetype",
		ConvertedName:  "SourceType",
		CreationOnly:   true,
		Description:    `The type of the source.`,
		Exposed:        true,
		Name:           "sourceType",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"UpdateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updateidempotencykey",
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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

// SuspiciousActivityLowerCaseAttributesMap represents the map of attribute for SuspiciousActivity.
var SuspiciousActivityLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"annotations": {
		AllowedChoices: []string{},
		BSONFieldName:  "annotations",
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"associatedtags": {
		AllowedChoices: []string{},
		BSONFieldName:  "associatedtags",
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"categories": {
		AllowedChoices: []string{},
		BSONFieldName:  "categories",
		ConvertedName:  "Categories",
		Description:    `The list of category names.`,
		Exposed:        true,
		Name:           "categories",
		Required:       true,
		Stored:         true,
		SubType:        "integer",
		Type:           "list",
	},
	"counter": {
		AllowedChoices: []string{},
		BSONFieldName:  "counter",
		ConvertedName:  "Counter",
		Description:    `Number of times this suspicious activity is found.`,
		Exposed:        true,
		Name:           "counter",
		Stored:         true,
		Type:           "integer",
	},
	"createidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createidempotencykey",
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
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
	"firstoccurrencetime": {
		AllowedChoices: []string{},
		BSONFieldName:  "firstoccurrencetime",
		ConvertedName:  "FirstOccurrenceTime",
		Description:    `The timestamp when the suspicious activity first occurred.`,
		Exposed:        true,
		Name:           "firstOccurrenceTime",
		Required:       true,
		Stored:         true,
		Type:           "time",
	},
	"lastoccurrencetime": {
		AllowedChoices: []string{},
		BSONFieldName:  "lastoccurrencetime",
		ConvertedName:  "LastOccurrenceTime",
		Description:    `The timestamp when the suspicious activity last occurred.`,
		Exposed:        true,
		Name:           "lastOccurrenceTime",
		Required:       true,
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
	"normalizedtags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "normalizedtags",
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"protected": {
		AllowedChoices: []string{},
		BSONFieldName:  "protected",
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"risk": {
		AllowedChoices: []string{"Low", "Medium", "High", "Unknown"},
		BSONFieldName:  "risk",
		ConvertedName:  "Risk",
		DefaultValue:   SuspiciousActivityRiskUnknown,
		Description:    `The level of risk.`,
		Exposed:        true,
		Name:           "risk",
		Stored:         true,
		Type:           "enum",
	},
	"sourceid": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourceid",
		ConvertedName:  "SourceID",
		CreationOnly:   true,
		Description:    `The identifier of the source.`,
		Exposed:        true,
		Name:           "sourceID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"sourcename": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourcename",
		ConvertedName:  "SourceName",
		CreationOnly:   true,
		Description:    `The name of the source.`,
		Exposed:        true,
		Name:           "sourceName",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"sourcenamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourcenamespace",
		ConvertedName:  "SourceNamespace",
		CreationOnly:   true,
		Description:    `The namespace of the source.`,
		Exposed:        true,
		Name:           "sourceNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"sourcetype": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourcetype",
		ConvertedName:  "SourceType",
		CreationOnly:   true,
		Description:    `The type of the source.`,
		Exposed:        true,
		Name:           "sourceType",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"updateidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updateidempotencykey",
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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

// SparseSuspiciousActivitiesList represents a list of SparseSuspiciousActivities
type SparseSuspiciousActivitiesList []*SparseSuspiciousActivity

// Identity returns the identity of the objects in the list.
func (o SparseSuspiciousActivitiesList) Identity() elemental.Identity {

	return SuspiciousActivityIdentity
}

// Copy returns a pointer to a copy the SparseSuspiciousActivitiesList.
func (o SparseSuspiciousActivitiesList) Copy() elemental.Identifiables {

	copy := append(SparseSuspiciousActivitiesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseSuspiciousActivitiesList.
func (o SparseSuspiciousActivitiesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseSuspiciousActivitiesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseSuspiciousActivity))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseSuspiciousActivitiesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseSuspiciousActivitiesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseSuspiciousActivitiesList converted to SuspiciousActivitiesList.
func (o SparseSuspiciousActivitiesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseSuspiciousActivitiesList) Version() int {

	return 1
}

// SparseSuspiciousActivity represents the sparse version of a suspiciousactivity.
type SparseSuspiciousActivity struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// The list of category names.
	Categories *[]int `json:"categories,omitempty" msgpack:"categories,omitempty" bson:"categories,omitempty" mapstructure:"categories,omitempty"`

	// Number of times this suspicious activity is found.
	Counter *int `json:"counter,omitempty" msgpack:"counter,omitempty" bson:"counter,omitempty" mapstructure:"counter,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// The timestamp when the suspicious activity first occurred.
	FirstOccurrenceTime *time.Time `json:"firstOccurrenceTime,omitempty" msgpack:"firstOccurrenceTime,omitempty" bson:"firstoccurrencetime,omitempty" mapstructure:"firstOccurrenceTime,omitempty"`

	// The timestamp when the suspicious activity last occurred.
	LastOccurrenceTime *time.Time `json:"lastOccurrenceTime,omitempty" msgpack:"lastOccurrenceTime,omitempty" bson:"lastoccurrencetime,omitempty" mapstructure:"lastOccurrenceTime,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// The level of risk.
	Risk *SuspiciousActivityRiskValue `json:"risk,omitempty" msgpack:"risk,omitempty" bson:"risk,omitempty" mapstructure:"risk,omitempty"`

	// The identifier of the source.
	SourceID *string `json:"sourceID,omitempty" msgpack:"sourceID,omitempty" bson:"sourceid,omitempty" mapstructure:"sourceID,omitempty"`

	// The name of the source.
	SourceName *string `json:"sourceName,omitempty" msgpack:"sourceName,omitempty" bson:"sourcename,omitempty" mapstructure:"sourceName,omitempty"`

	// The namespace of the source.
	SourceNamespace *string `json:"sourceNamespace,omitempty" msgpack:"sourceNamespace,omitempty" bson:"sourcenamespace,omitempty" mapstructure:"sourceNamespace,omitempty"`

	// The type of the source.
	SourceType *string `json:"sourceType,omitempty" msgpack:"sourceType,omitempty" bson:"sourcetype,omitempty" mapstructure:"sourceType,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseSuspiciousActivity returns a new  SparseSuspiciousActivity.
func NewSparseSuspiciousActivity() *SparseSuspiciousActivity {
	return &SparseSuspiciousActivity{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseSuspiciousActivity) Identity() elemental.Identity {

	return SuspiciousActivityIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseSuspiciousActivity) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseSuspiciousActivity) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseSuspiciousActivity) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseSuspiciousActivity{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.Categories != nil {
		s.Categories = o.Categories
	}
	if o.Counter != nil {
		s.Counter = o.Counter
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.FirstOccurrenceTime != nil {
		s.FirstOccurrenceTime = o.FirstOccurrenceTime
	}
	if o.LastOccurrenceTime != nil {
		s.LastOccurrenceTime = o.LastOccurrenceTime
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.Risk != nil {
		s.Risk = o.Risk
	}
	if o.SourceID != nil {
		s.SourceID = o.SourceID
	}
	if o.SourceName != nil {
		s.SourceName = o.SourceName
	}
	if o.SourceNamespace != nil {
		s.SourceNamespace = o.SourceNamespace
	}
	if o.SourceType != nil {
		s.SourceType = o.SourceType
	}
	if o.UpdateIdempotencyKey != nil {
		s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
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
func (o *SparseSuspiciousActivity) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseSuspiciousActivity{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.Categories != nil {
		o.Categories = s.Categories
	}
	if s.Counter != nil {
		o.Counter = s.Counter
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.FirstOccurrenceTime != nil {
		o.FirstOccurrenceTime = s.FirstOccurrenceTime
	}
	if s.LastOccurrenceTime != nil {
		o.LastOccurrenceTime = s.LastOccurrenceTime
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.Risk != nil {
		o.Risk = s.Risk
	}
	if s.SourceID != nil {
		o.SourceID = s.SourceID
	}
	if s.SourceName != nil {
		o.SourceName = s.SourceName
	}
	if s.SourceNamespace != nil {
		o.SourceNamespace = s.SourceNamespace
	}
	if s.SourceType != nil {
		o.SourceType = s.SourceType
	}
	if s.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
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
func (o *SparseSuspiciousActivity) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseSuspiciousActivity) ToPlain() elemental.PlainIdentifiable {

	out := NewSuspiciousActivity()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.Categories != nil {
		out.Categories = *o.Categories
	}
	if o.Counter != nil {
		out.Counter = *o.Counter
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.FirstOccurrenceTime != nil {
		out.FirstOccurrenceTime = *o.FirstOccurrenceTime
	}
	if o.LastOccurrenceTime != nil {
		out.LastOccurrenceTime = *o.LastOccurrenceTime
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Risk != nil {
		out.Risk = *o.Risk
	}
	if o.SourceID != nil {
		out.SourceID = *o.SourceID
	}
	if o.SourceName != nil {
		out.SourceName = *o.SourceName
	}
	if o.SourceNamespace != nil {
		out.SourceNamespace = *o.SourceNamespace
	}
	if o.SourceType != nil {
		out.SourceType = *o.SourceType
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
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

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseSuspiciousActivity) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseSuspiciousActivity) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseSuspiciousActivity) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseSuspiciousActivity) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseSuspiciousActivity) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseSuspiciousActivity) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseSuspiciousActivity) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseSuspiciousActivity) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseSuspiciousActivity) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseSuspiciousActivity) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseSuspiciousActivity) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseSuspiciousActivity) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseSuspiciousActivity) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseSuspiciousActivity) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseSuspiciousActivity) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseSuspiciousActivity) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseSuspiciousActivity) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseSuspiciousActivity) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseSuspiciousActivity) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseSuspiciousActivity) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseSuspiciousActivity) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseSuspiciousActivity) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseSuspiciousActivity) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseSuspiciousActivity) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseSuspiciousActivity.
func (o *SparseSuspiciousActivity) DeepCopy() *SparseSuspiciousActivity {

	if o == nil {
		return nil
	}

	out := &SparseSuspiciousActivity{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseSuspiciousActivity.
func (o *SparseSuspiciousActivity) DeepCopyInto(out *SparseSuspiciousActivity) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseSuspiciousActivity: %s", err))
	}

	*out = *target.(*SparseSuspiciousActivity)
}

type mongoAttributesSuspiciousActivity struct {
	ID                   bson.ObjectId               `bson:"_id,omitempty"`
	Annotations          map[string][]string         `bson:"annotations"`
	AssociatedTags       []string                    `bson:"associatedtags"`
	Categories           []int                       `bson:"categories"`
	Counter              int                         `bson:"counter"`
	CreateIdempotencyKey string                      `bson:"createidempotencykey"`
	CreateTime           time.Time                   `bson:"createtime"`
	FirstOccurrenceTime  time.Time                   `bson:"firstoccurrencetime"`
	LastOccurrenceTime   time.Time                   `bson:"lastoccurrencetime"`
	MigrationsLog        map[string]string           `bson:"migrationslog,omitempty"`
	Namespace            string                      `bson:"namespace"`
	NormalizedTags       []string                    `bson:"normalizedtags"`
	Protected            bool                        `bson:"protected"`
	Risk                 SuspiciousActivityRiskValue `bson:"risk"`
	SourceID             string                      `bson:"sourceid"`
	SourceName           string                      `bson:"sourcename"`
	SourceNamespace      string                      `bson:"sourcenamespace"`
	SourceType           string                      `bson:"sourcetype"`
	UpdateIdempotencyKey string                      `bson:"updateidempotencykey"`
	UpdateTime           time.Time                   `bson:"updatetime"`
	ZHash                int                         `bson:"zhash"`
	Zone                 int                         `bson:"zone"`
}
type mongoAttributesSparseSuspiciousActivity struct {
	ID                   bson.ObjectId                `bson:"_id,omitempty"`
	Annotations          *map[string][]string         `bson:"annotations,omitempty"`
	AssociatedTags       *[]string                    `bson:"associatedtags,omitempty"`
	Categories           *[]int                       `bson:"categories,omitempty"`
	Counter              *int                         `bson:"counter,omitempty"`
	CreateIdempotencyKey *string                      `bson:"createidempotencykey,omitempty"`
	CreateTime           *time.Time                   `bson:"createtime,omitempty"`
	FirstOccurrenceTime  *time.Time                   `bson:"firstoccurrencetime,omitempty"`
	LastOccurrenceTime   *time.Time                   `bson:"lastoccurrencetime,omitempty"`
	MigrationsLog        *map[string]string           `bson:"migrationslog,omitempty"`
	Namespace            *string                      `bson:"namespace,omitempty"`
	NormalizedTags       *[]string                    `bson:"normalizedtags,omitempty"`
	Protected            *bool                        `bson:"protected,omitempty"`
	Risk                 *SuspiciousActivityRiskValue `bson:"risk,omitempty"`
	SourceID             *string                      `bson:"sourceid,omitempty"`
	SourceName           *string                      `bson:"sourcename,omitempty"`
	SourceNamespace      *string                      `bson:"sourcenamespace,omitempty"`
	SourceType           *string                      `bson:"sourcetype,omitempty"`
	UpdateIdempotencyKey *string                      `bson:"updateidempotencykey,omitempty"`
	UpdateTime           *time.Time                   `bson:"updatetime,omitempty"`
	ZHash                *int                         `bson:"zhash,omitempty"`
	Zone                 *int                         `bson:"zone,omitempty"`
}
