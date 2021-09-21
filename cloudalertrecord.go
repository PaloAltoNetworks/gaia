package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudAlertRecordResourceTypeValue represents the possible values for attribute "resourceType".
type CloudAlertRecordResourceTypeValue string

const (
	// CloudAlertRecordResourceTypeInstance represents the value Instance.
	CloudAlertRecordResourceTypeInstance CloudAlertRecordResourceTypeValue = "Instance"

	// CloudAlertRecordResourceTypeInterface represents the value Interface.
	CloudAlertRecordResourceTypeInterface CloudAlertRecordResourceTypeValue = "Interface"

	// CloudAlertRecordResourceTypeSubnet represents the value Subnet.
	CloudAlertRecordResourceTypeSubnet CloudAlertRecordResourceTypeValue = "Subnet"

	// CloudAlertRecordResourceTypeVPC represents the value VPC.
	CloudAlertRecordResourceTypeVPC CloudAlertRecordResourceTypeValue = "VPC"
)

// CloudAlertRecordIdentity represents the Identity of the object.
var CloudAlertRecordIdentity = elemental.Identity{
	Name:     "cloudalertrecord",
	Category: "cloudalertrecords",
	Package:  "vargid",
	Private:  false,
}

// CloudAlertRecordsList represents a list of CloudAlertRecords
type CloudAlertRecordsList []*CloudAlertRecord

// Identity returns the identity of the objects in the list.
func (o CloudAlertRecordsList) Identity() elemental.Identity {

	return CloudAlertRecordIdentity
}

// Copy returns a pointer to a copy the CloudAlertRecordsList.
func (o CloudAlertRecordsList) Copy() elemental.Identifiables {

	copy := append(CloudAlertRecordsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CloudAlertRecordsList.
func (o CloudAlertRecordsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudAlertRecordsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudAlertRecord))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudAlertRecordsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudAlertRecordsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the CloudAlertRecordsList converted to SparseCloudAlertRecordsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudAlertRecordsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudAlertRecordsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudAlertRecord)
	}

	return out
}

// Version returns the version of the content.
func (o CloudAlertRecordsList) Version() int {

	return 1
}

// CloudAlertRecord represents the model of a cloudalertrecord
type CloudAlertRecord struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Account ID of the resource for which the Alert Record is
	// raised.
	AccountID string `json:"accountID" msgpack:"accountID" bson:"accountid" mapstructure:"accountID,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Result of the last execution timestamp.
	LastExecutionTimestamp time.Time `json:"lastExecutionTimestamp" msgpack:"lastExecutionTimestamp" bson:"lastexecutiontimestamp" mapstructure:"lastExecutionTimestamp,omitempty"`

	// Sum of FNV-32a hashes of all the instances or interfaces grouped under the
	// resource.
	MetadataHash int `json:"metadataHash" msgpack:"metadataHash" bson:"metadatahash" mapstructure:"metadataHash,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Prisma Cloud Alert Rule which generated the Alert Record.
	PrismaCloudAlertRuleID string `json:"prismaCloudAlertRuleID" msgpack:"prismaCloudAlertRuleID" bson:"prismacloudalertruleid" mapstructure:"prismaCloudAlertRuleID,omitempty"`

	// Policy ID which generated the Alert Record.
	PrismaCloudPolicyID string `json:"prismaCloudPolicyID" msgpack:"prismaCloudPolicyID" bson:"prismacloudpolicyid" mapstructure:"prismaCloudPolicyID,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Indicates if the alert is published to PC.
	Published bool `json:"published" msgpack:"published" bson:"published" mapstructure:"published,omitempty"`

	// Region of the resource for which the Alert Record is
	// raised.
	Region string `json:"region" msgpack:"region" bson:"region" mapstructure:"region,omitempty"`

	// Number of interfaces/instances for which the alert is raised.
	ResourceCount int `json:"resourceCount" msgpack:"resourceCount" bson:"resourcecount" mapstructure:"resourceCount,omitempty"`

	// Resource ID of the resource for which the Alert Record is raised.
	ResourceID string `json:"resourceID" msgpack:"resourceID" bson:"resourceid" mapstructure:"resourceID,omitempty"`

	// Returns the type of the resource on which alert was raised.
	ResourceType CloudAlertRecordResourceTypeValue `json:"resourceType" msgpack:"resourceType" bson:"resourcetype" mapstructure:"resourceType,omitempty"`

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

// NewCloudAlertRecord returns a new *CloudAlertRecord
func NewCloudAlertRecord() *CloudAlertRecord {

	return &CloudAlertRecord{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		Published:      false,
		NormalizedTags: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *CloudAlertRecord) Identity() elemental.Identity {

	return CloudAlertRecordIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudAlertRecord) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudAlertRecord) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudAlertRecord) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudAlertRecord{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.AccountID = o.AccountID
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.LastExecutionTimestamp = o.LastExecutionTimestamp
	s.MetadataHash = o.MetadataHash
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.PrismaCloudAlertRuleID = o.PrismaCloudAlertRuleID
	s.PrismaCloudPolicyID = o.PrismaCloudPolicyID
	s.Protected = o.Protected
	s.Published = o.Published
	s.Region = o.Region
	s.ResourceCount = o.ResourceCount
	s.ResourceID = o.ResourceID
	s.ResourceType = o.ResourceType
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudAlertRecord) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudAlertRecord{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.AccountID = s.AccountID
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.LastExecutionTimestamp = s.LastExecutionTimestamp
	o.MetadataHash = s.MetadataHash
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.PrismaCloudAlertRuleID = s.PrismaCloudAlertRuleID
	o.PrismaCloudPolicyID = s.PrismaCloudPolicyID
	o.Protected = s.Protected
	o.Published = s.Published
	o.Region = s.Region
	o.ResourceCount = s.ResourceCount
	o.ResourceID = s.ResourceID
	o.ResourceType = s.ResourceType
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudAlertRecord) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudAlertRecord) BleveType() string {

	return "cloudalertrecord"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudAlertRecord) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *CloudAlertRecord) Doc() string {

	return `Represents an Alert that is raised against a resource based on a RQL (policy)
associated to an alert rule.`
}

func (o *CloudAlertRecord) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *CloudAlertRecord) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *CloudAlertRecord) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *CloudAlertRecord) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *CloudAlertRecord) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *CloudAlertRecord) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *CloudAlertRecord) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *CloudAlertRecord) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *CloudAlertRecord) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetName returns the Name of the receiver.
func (o *CloudAlertRecord) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *CloudAlertRecord) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *CloudAlertRecord) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *CloudAlertRecord) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *CloudAlertRecord) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *CloudAlertRecord) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *CloudAlertRecord) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *CloudAlertRecord) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *CloudAlertRecord) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *CloudAlertRecord) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *CloudAlertRecord) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *CloudAlertRecord) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *CloudAlertRecord) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *CloudAlertRecord) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *CloudAlertRecord) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *CloudAlertRecord) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudAlertRecord) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudAlertRecord{
			ID:                     &o.ID,
			AccountID:              &o.AccountID,
			Annotations:            &o.Annotations,
			AssociatedTags:         &o.AssociatedTags,
			CreateIdempotencyKey:   &o.CreateIdempotencyKey,
			CreateTime:             &o.CreateTime,
			LastExecutionTimestamp: &o.LastExecutionTimestamp,
			MetadataHash:           &o.MetadataHash,
			Name:                   &o.Name,
			Namespace:              &o.Namespace,
			NormalizedTags:         &o.NormalizedTags,
			PrismaCloudAlertRuleID: &o.PrismaCloudAlertRuleID,
			PrismaCloudPolicyID:    &o.PrismaCloudPolicyID,
			Protected:              &o.Protected,
			Published:              &o.Published,
			Region:                 &o.Region,
			ResourceCount:          &o.ResourceCount,
			ResourceID:             &o.ResourceID,
			ResourceType:           &o.ResourceType,
			UpdateIdempotencyKey:   &o.UpdateIdempotencyKey,
			UpdateTime:             &o.UpdateTime,
			ZHash:                  &o.ZHash,
			Zone:                   &o.Zone,
		}
	}

	sp := &SparseCloudAlertRecord{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "accountID":
			sp.AccountID = &(o.AccountID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "lastExecutionTimestamp":
			sp.LastExecutionTimestamp = &(o.LastExecutionTimestamp)
		case "metadataHash":
			sp.MetadataHash = &(o.MetadataHash)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "prismaCloudAlertRuleID":
			sp.PrismaCloudAlertRuleID = &(o.PrismaCloudAlertRuleID)
		case "prismaCloudPolicyID":
			sp.PrismaCloudPolicyID = &(o.PrismaCloudPolicyID)
		case "protected":
			sp.Protected = &(o.Protected)
		case "published":
			sp.Published = &(o.Published)
		case "region":
			sp.Region = &(o.Region)
		case "resourceCount":
			sp.ResourceCount = &(o.ResourceCount)
		case "resourceID":
			sp.ResourceID = &(o.ResourceID)
		case "resourceType":
			sp.ResourceType = &(o.ResourceType)
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

// Patch apply the non nil value of a *SparseCloudAlertRecord to the object.
func (o *CloudAlertRecord) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudAlertRecord)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.AccountID != nil {
		o.AccountID = *so.AccountID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.LastExecutionTimestamp != nil {
		o.LastExecutionTimestamp = *so.LastExecutionTimestamp
	}
	if so.MetadataHash != nil {
		o.MetadataHash = *so.MetadataHash
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.PrismaCloudAlertRuleID != nil {
		o.PrismaCloudAlertRuleID = *so.PrismaCloudAlertRuleID
	}
	if so.PrismaCloudPolicyID != nil {
		o.PrismaCloudPolicyID = *so.PrismaCloudPolicyID
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Published != nil {
		o.Published = *so.Published
	}
	if so.Region != nil {
		o.Region = *so.Region
	}
	if so.ResourceCount != nil {
		o.ResourceCount = *so.ResourceCount
	}
	if so.ResourceID != nil {
		o.ResourceID = *so.ResourceID
	}
	if so.ResourceType != nil {
		o.ResourceType = *so.ResourceType
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

// DeepCopy returns a deep copy if the CloudAlertRecord.
func (o *CloudAlertRecord) DeepCopy() *CloudAlertRecord {

	if o == nil {
		return nil
	}

	out := &CloudAlertRecord{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudAlertRecord.
func (o *CloudAlertRecord) DeepCopyInto(out *CloudAlertRecord) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudAlertRecord: %s", err))
	}

	*out = *target.(*CloudAlertRecord)
}

// Validate valides the current information stored into the structure.
func (o *CloudAlertRecord) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("resourceType", string(o.ResourceType), []string{"Interface", "Instance", "VPC", "Subnet"}, true); err != nil {
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
func (*CloudAlertRecord) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudAlertRecordAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudAlertRecordLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudAlertRecord) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudAlertRecordAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudAlertRecord) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "accountID":
		return o.AccountID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "lastExecutionTimestamp":
		return o.LastExecutionTimestamp
	case "metadataHash":
		return o.MetadataHash
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "prismaCloudAlertRuleID":
		return o.PrismaCloudAlertRuleID
	case "prismaCloudPolicyID":
		return o.PrismaCloudPolicyID
	case "protected":
		return o.Protected
	case "published":
		return o.Published
	case "region":
		return o.Region
	case "resourceCount":
		return o.ResourceCount
	case "resourceID":
		return o.ResourceID
	case "resourceType":
		return o.ResourceType
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

// CloudAlertRecordAttributesMap represents the map of attribute for CloudAlertRecord.
var CloudAlertRecordAttributesMap = map[string]elemental.AttributeSpecification{
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
	"AccountID": {
		AllowedChoices: []string{},
		BSONFieldName:  "accountid",
		ConvertedName:  "AccountID",
		Description: `Account ID of the resource for which the Alert Record is
raised.`,
		Exposed: true,
		Name:    "accountID",
		Stored:  true,
		SubType: "string",
		Type:    "string",
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
	"LastExecutionTimestamp": {
		AllowedChoices: []string{},
		BSONFieldName:  "lastexecutiontimestamp",
		ConvertedName:  "LastExecutionTimestamp",
		Description:    `Result of the last execution timestamp.`,
		Exposed:        true,
		Name:           "lastExecutionTimestamp",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"MetadataHash": {
		AllowedChoices: []string{},
		BSONFieldName:  "metadatahash",
		ConvertedName:  "MetadataHash",
		Description: `Sum of FNV-32a hashes of all the instances or interfaces grouped under the
resource.`,
		Exposed: true,
		Name:    "metadataHash",
		Stored:  true,
		SubType: "integer",
		Type:    "integer",
	},
	"Name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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
	"PrismaCloudAlertRuleID": {
		AllowedChoices: []string{},
		BSONFieldName:  "prismacloudalertruleid",
		ConvertedName:  "PrismaCloudAlertRuleID",
		Description:    `Prisma Cloud Alert Rule which generated the Alert Record.`,
		Exposed:        true,
		Name:           "prismaCloudAlertRuleID",
		Stored:         true,
		SubType:        "string",
		Type:           "string",
	},
	"PrismaCloudPolicyID": {
		AllowedChoices: []string{},
		BSONFieldName:  "prismacloudpolicyid",
		ConvertedName:  "PrismaCloudPolicyID",
		Description:    `Policy ID which generated the Alert Record.`,
		Exposed:        true,
		Name:           "prismaCloudPolicyID",
		Stored:         true,
		SubType:        "string",
		Type:           "string",
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
	"Published": {
		AllowedChoices: []string{},
		BSONFieldName:  "published",
		ConvertedName:  "Published",
		Description:    `Indicates if the alert is published to PC.`,
		Exposed:        true,
		Name:           "published",
		Stored:         true,
		Type:           "boolean",
	},
	"Region": {
		AllowedChoices: []string{},
		BSONFieldName:  "region",
		ConvertedName:  "Region",
		Description: `Region of the resource for which the Alert Record is
raised.`,
		Exposed: true,
		Name:    "region",
		Stored:  true,
		SubType: "string",
		Type:    "string",
	},
	"ResourceCount": {
		AllowedChoices: []string{},
		BSONFieldName:  "resourcecount",
		ConvertedName:  "ResourceCount",
		Description:    `Number of interfaces/instances for which the alert is raised.`,
		Exposed:        true,
		Name:           "resourceCount",
		Stored:         true,
		SubType:        "integer",
		Type:           "integer",
	},
	"ResourceID": {
		AllowedChoices: []string{},
		BSONFieldName:  "resourceid",
		ConvertedName:  "ResourceID",
		Description:    `Resource ID of the resource for which the Alert Record is raised.`,
		Exposed:        true,
		Name:           "resourceID",
		Stored:         true,
		SubType:        "string",
		Type:           "string",
	},
	"ResourceType": {
		AllowedChoices: []string{"Interface", "Instance", "VPC", "Subnet"},
		Autogenerated:  true,
		BSONFieldName:  "resourcetype",
		ConvertedName:  "ResourceType",
		Description:    `Returns the type of the resource on which alert was raised.`,
		Exposed:        true,
		Name:           "resourceType",
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
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

// CloudAlertRecordLowerCaseAttributesMap represents the map of attribute for CloudAlertRecord.
var CloudAlertRecordLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"accountid": {
		AllowedChoices: []string{},
		BSONFieldName:  "accountid",
		ConvertedName:  "AccountID",
		Description: `Account ID of the resource for which the Alert Record is
raised.`,
		Exposed: true,
		Name:    "accountID",
		Stored:  true,
		SubType: "string",
		Type:    "string",
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
	"lastexecutiontimestamp": {
		AllowedChoices: []string{},
		BSONFieldName:  "lastexecutiontimestamp",
		ConvertedName:  "LastExecutionTimestamp",
		Description:    `Result of the last execution timestamp.`,
		Exposed:        true,
		Name:           "lastExecutionTimestamp",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"metadatahash": {
		AllowedChoices: []string{},
		BSONFieldName:  "metadatahash",
		ConvertedName:  "MetadataHash",
		Description: `Sum of FNV-32a hashes of all the instances or interfaces grouped under the
resource.`,
		Exposed: true,
		Name:    "metadataHash",
		Stored:  true,
		SubType: "integer",
		Type:    "integer",
	},
	"name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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
	"prismacloudalertruleid": {
		AllowedChoices: []string{},
		BSONFieldName:  "prismacloudalertruleid",
		ConvertedName:  "PrismaCloudAlertRuleID",
		Description:    `Prisma Cloud Alert Rule which generated the Alert Record.`,
		Exposed:        true,
		Name:           "prismaCloudAlertRuleID",
		Stored:         true,
		SubType:        "string",
		Type:           "string",
	},
	"prismacloudpolicyid": {
		AllowedChoices: []string{},
		BSONFieldName:  "prismacloudpolicyid",
		ConvertedName:  "PrismaCloudPolicyID",
		Description:    `Policy ID which generated the Alert Record.`,
		Exposed:        true,
		Name:           "prismaCloudPolicyID",
		Stored:         true,
		SubType:        "string",
		Type:           "string",
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
	"published": {
		AllowedChoices: []string{},
		BSONFieldName:  "published",
		ConvertedName:  "Published",
		Description:    `Indicates if the alert is published to PC.`,
		Exposed:        true,
		Name:           "published",
		Stored:         true,
		Type:           "boolean",
	},
	"region": {
		AllowedChoices: []string{},
		BSONFieldName:  "region",
		ConvertedName:  "Region",
		Description: `Region of the resource for which the Alert Record is
raised.`,
		Exposed: true,
		Name:    "region",
		Stored:  true,
		SubType: "string",
		Type:    "string",
	},
	"resourcecount": {
		AllowedChoices: []string{},
		BSONFieldName:  "resourcecount",
		ConvertedName:  "ResourceCount",
		Description:    `Number of interfaces/instances for which the alert is raised.`,
		Exposed:        true,
		Name:           "resourceCount",
		Stored:         true,
		SubType:        "integer",
		Type:           "integer",
	},
	"resourceid": {
		AllowedChoices: []string{},
		BSONFieldName:  "resourceid",
		ConvertedName:  "ResourceID",
		Description:    `Resource ID of the resource for which the Alert Record is raised.`,
		Exposed:        true,
		Name:           "resourceID",
		Stored:         true,
		SubType:        "string",
		Type:           "string",
	},
	"resourcetype": {
		AllowedChoices: []string{"Interface", "Instance", "VPC", "Subnet"},
		Autogenerated:  true,
		BSONFieldName:  "resourcetype",
		ConvertedName:  "ResourceType",
		Description:    `Returns the type of the resource on which alert was raised.`,
		Exposed:        true,
		Name:           "resourceType",
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
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

// SparseCloudAlertRecordsList represents a list of SparseCloudAlertRecords
type SparseCloudAlertRecordsList []*SparseCloudAlertRecord

// Identity returns the identity of the objects in the list.
func (o SparseCloudAlertRecordsList) Identity() elemental.Identity {

	return CloudAlertRecordIdentity
}

// Copy returns a pointer to a copy the SparseCloudAlertRecordsList.
func (o SparseCloudAlertRecordsList) Copy() elemental.Identifiables {

	copy := append(SparseCloudAlertRecordsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudAlertRecordsList.
func (o SparseCloudAlertRecordsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudAlertRecordsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudAlertRecord))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudAlertRecordsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudAlertRecordsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseCloudAlertRecordsList converted to CloudAlertRecordsList.
func (o SparseCloudAlertRecordsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudAlertRecordsList) Version() int {

	return 1
}

// SparseCloudAlertRecord represents the sparse version of a cloudalertrecord.
type SparseCloudAlertRecord struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Account ID of the resource for which the Alert Record is
	// raised.
	AccountID *string `json:"accountID,omitempty" msgpack:"accountID,omitempty" bson:"accountid,omitempty" mapstructure:"accountID,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Result of the last execution timestamp.
	LastExecutionTimestamp *time.Time `json:"lastExecutionTimestamp,omitempty" msgpack:"lastExecutionTimestamp,omitempty" bson:"lastexecutiontimestamp,omitempty" mapstructure:"lastExecutionTimestamp,omitempty"`

	// Sum of FNV-32a hashes of all the instances or interfaces grouped under the
	// resource.
	MetadataHash *int `json:"metadataHash,omitempty" msgpack:"metadataHash,omitempty" bson:"metadatahash,omitempty" mapstructure:"metadataHash,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Prisma Cloud Alert Rule which generated the Alert Record.
	PrismaCloudAlertRuleID *string `json:"prismaCloudAlertRuleID,omitempty" msgpack:"prismaCloudAlertRuleID,omitempty" bson:"prismacloudalertruleid,omitempty" mapstructure:"prismaCloudAlertRuleID,omitempty"`

	// Policy ID which generated the Alert Record.
	PrismaCloudPolicyID *string `json:"prismaCloudPolicyID,omitempty" msgpack:"prismaCloudPolicyID,omitempty" bson:"prismacloudpolicyid,omitempty" mapstructure:"prismaCloudPolicyID,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Indicates if the alert is published to PC.
	Published *bool `json:"published,omitempty" msgpack:"published,omitempty" bson:"published,omitempty" mapstructure:"published,omitempty"`

	// Region of the resource for which the Alert Record is
	// raised.
	Region *string `json:"region,omitempty" msgpack:"region,omitempty" bson:"region,omitempty" mapstructure:"region,omitempty"`

	// Number of interfaces/instances for which the alert is raised.
	ResourceCount *int `json:"resourceCount,omitempty" msgpack:"resourceCount,omitempty" bson:"resourcecount,omitempty" mapstructure:"resourceCount,omitempty"`

	// Resource ID of the resource for which the Alert Record is raised.
	ResourceID *string `json:"resourceID,omitempty" msgpack:"resourceID,omitempty" bson:"resourceid,omitempty" mapstructure:"resourceID,omitempty"`

	// Returns the type of the resource on which alert was raised.
	ResourceType *CloudAlertRecordResourceTypeValue `json:"resourceType,omitempty" msgpack:"resourceType,omitempty" bson:"resourcetype,omitempty" mapstructure:"resourceType,omitempty"`

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

// NewSparseCloudAlertRecord returns a new  SparseCloudAlertRecord.
func NewSparseCloudAlertRecord() *SparseCloudAlertRecord {
	return &SparseCloudAlertRecord{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudAlertRecord) Identity() elemental.Identity {

	return CloudAlertRecordIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudAlertRecord) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudAlertRecord) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudAlertRecord) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudAlertRecord{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.AccountID != nil {
		s.AccountID = o.AccountID
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.LastExecutionTimestamp != nil {
		s.LastExecutionTimestamp = o.LastExecutionTimestamp
	}
	if o.MetadataHash != nil {
		s.MetadataHash = o.MetadataHash
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.PrismaCloudAlertRuleID != nil {
		s.PrismaCloudAlertRuleID = o.PrismaCloudAlertRuleID
	}
	if o.PrismaCloudPolicyID != nil {
		s.PrismaCloudPolicyID = o.PrismaCloudPolicyID
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.Published != nil {
		s.Published = o.Published
	}
	if o.Region != nil {
		s.Region = o.Region
	}
	if o.ResourceCount != nil {
		s.ResourceCount = o.ResourceCount
	}
	if o.ResourceID != nil {
		s.ResourceID = o.ResourceID
	}
	if o.ResourceType != nil {
		s.ResourceType = o.ResourceType
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
func (o *SparseCloudAlertRecord) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudAlertRecord{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.AccountID != nil {
		o.AccountID = s.AccountID
	}
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.LastExecutionTimestamp != nil {
		o.LastExecutionTimestamp = s.LastExecutionTimestamp
	}
	if s.MetadataHash != nil {
		o.MetadataHash = s.MetadataHash
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.PrismaCloudAlertRuleID != nil {
		o.PrismaCloudAlertRuleID = s.PrismaCloudAlertRuleID
	}
	if s.PrismaCloudPolicyID != nil {
		o.PrismaCloudPolicyID = s.PrismaCloudPolicyID
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.Published != nil {
		o.Published = s.Published
	}
	if s.Region != nil {
		o.Region = s.Region
	}
	if s.ResourceCount != nil {
		o.ResourceCount = s.ResourceCount
	}
	if s.ResourceID != nil {
		o.ResourceID = s.ResourceID
	}
	if s.ResourceType != nil {
		o.ResourceType = s.ResourceType
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
func (o *SparseCloudAlertRecord) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudAlertRecord) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudAlertRecord()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.AccountID != nil {
		out.AccountID = *o.AccountID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.LastExecutionTimestamp != nil {
		out.LastExecutionTimestamp = *o.LastExecutionTimestamp
	}
	if o.MetadataHash != nil {
		out.MetadataHash = *o.MetadataHash
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.PrismaCloudAlertRuleID != nil {
		out.PrismaCloudAlertRuleID = *o.PrismaCloudAlertRuleID
	}
	if o.PrismaCloudPolicyID != nil {
		out.PrismaCloudPolicyID = *o.PrismaCloudPolicyID
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Published != nil {
		out.Published = *o.Published
	}
	if o.Region != nil {
		out.Region = *o.Region
	}
	if o.ResourceCount != nil {
		out.ResourceCount = *o.ResourceCount
	}
	if o.ResourceID != nil {
		out.ResourceID = *o.ResourceID
	}
	if o.ResourceType != nil {
		out.ResourceType = *o.ResourceType
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
func (o *SparseCloudAlertRecord) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseCloudAlertRecord) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseCloudAlertRecord) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseCloudAlertRecord) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseCloudAlertRecord) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudAlertRecord) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseCloudAlertRecord) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseCloudAlertRecord) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetName returns the Name of the receiver.
func (o *SparseCloudAlertRecord) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseCloudAlertRecord) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseCloudAlertRecord) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseCloudAlertRecord) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseCloudAlertRecord) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseCloudAlertRecord) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseCloudAlertRecord) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseCloudAlertRecord) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseCloudAlertRecord) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudAlertRecord) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseCloudAlertRecord) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseCloudAlertRecord) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseCloudAlertRecord) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseCloudAlertRecord) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseCloudAlertRecord) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseCloudAlertRecord) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseCloudAlertRecord.
func (o *SparseCloudAlertRecord) DeepCopy() *SparseCloudAlertRecord {

	if o == nil {
		return nil
	}

	out := &SparseCloudAlertRecord{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudAlertRecord.
func (o *SparseCloudAlertRecord) DeepCopyInto(out *SparseCloudAlertRecord) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudAlertRecord: %s", err))
	}

	*out = *target.(*SparseCloudAlertRecord)
}

type mongoAttributesCloudAlertRecord struct {
	ID                     bson.ObjectId                     `bson:"_id,omitempty"`
	AccountID              string                            `bson:"accountid"`
	Annotations            map[string][]string               `bson:"annotations"`
	AssociatedTags         []string                          `bson:"associatedtags"`
	CreateIdempotencyKey   string                            `bson:"createidempotencykey"`
	CreateTime             time.Time                         `bson:"createtime"`
	LastExecutionTimestamp time.Time                         `bson:"lastexecutiontimestamp"`
	MetadataHash           int                               `bson:"metadatahash"`
	Name                   string                            `bson:"name"`
	Namespace              string                            `bson:"namespace"`
	NormalizedTags         []string                          `bson:"normalizedtags"`
	PrismaCloudAlertRuleID string                            `bson:"prismacloudalertruleid"`
	PrismaCloudPolicyID    string                            `bson:"prismacloudpolicyid"`
	Protected              bool                              `bson:"protected"`
	Published              bool                              `bson:"published"`
	Region                 string                            `bson:"region"`
	ResourceCount          int                               `bson:"resourcecount"`
	ResourceID             string                            `bson:"resourceid"`
	ResourceType           CloudAlertRecordResourceTypeValue `bson:"resourcetype"`
	UpdateIdempotencyKey   string                            `bson:"updateidempotencykey"`
	UpdateTime             time.Time                         `bson:"updatetime"`
	ZHash                  int                               `bson:"zhash"`
	Zone                   int                               `bson:"zone"`
}
type mongoAttributesSparseCloudAlertRecord struct {
	ID                     bson.ObjectId                      `bson:"_id,omitempty"`
	AccountID              *string                            `bson:"accountid,omitempty"`
	Annotations            *map[string][]string               `bson:"annotations,omitempty"`
	AssociatedTags         *[]string                          `bson:"associatedtags,omitempty"`
	CreateIdempotencyKey   *string                            `bson:"createidempotencykey,omitempty"`
	CreateTime             *time.Time                         `bson:"createtime,omitempty"`
	LastExecutionTimestamp *time.Time                         `bson:"lastexecutiontimestamp,omitempty"`
	MetadataHash           *int                               `bson:"metadatahash,omitempty"`
	Name                   *string                            `bson:"name,omitempty"`
	Namespace              *string                            `bson:"namespace,omitempty"`
	NormalizedTags         *[]string                          `bson:"normalizedtags,omitempty"`
	PrismaCloudAlertRuleID *string                            `bson:"prismacloudalertruleid,omitempty"`
	PrismaCloudPolicyID    *string                            `bson:"prismacloudpolicyid,omitempty"`
	Protected              *bool                              `bson:"protected,omitempty"`
	Published              *bool                              `bson:"published,omitempty"`
	Region                 *string                            `bson:"region,omitempty"`
	ResourceCount          *int                               `bson:"resourcecount,omitempty"`
	ResourceID             *string                            `bson:"resourceid,omitempty"`
	ResourceType           *CloudAlertRecordResourceTypeValue `bson:"resourcetype,omitempty"`
	UpdateIdempotencyKey   *string                            `bson:"updateidempotencykey,omitempty"`
	UpdateTime             *time.Time                         `bson:"updatetime,omitempty"`
	ZHash                  *int                               `bson:"zhash,omitempty"`
	Zone                   *int                               `bson:"zone,omitempty"`
}