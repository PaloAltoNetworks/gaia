package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudManagedNetworkTypeValue represents the possible values for attribute "type".
type CloudManagedNetworkTypeValue string

const (
	// CloudManagedNetworkTypeAWS represents the value AWS.
	CloudManagedNetworkTypeAWS CloudManagedNetworkTypeValue = "AWS"

	// CloudManagedNetworkTypeCustom represents the value Custom.
	CloudManagedNetworkTypeCustom CloudManagedNetworkTypeValue = "Custom"

	// CloudManagedNetworkTypeEnterprise represents the value Enterprise.
	CloudManagedNetworkTypeEnterprise CloudManagedNetworkTypeValue = "Enterprise"

	// CloudManagedNetworkTypeGCP represents the value GCP.
	CloudManagedNetworkTypeGCP CloudManagedNetworkTypeValue = "GCP"
)

// CloudManagedNetworkIdentity represents the Identity of the object.
var CloudManagedNetworkIdentity = elemental.Identity{
	Name:     "cloudmanagednetwork",
	Category: "cloudmanagednetworks",
	Package:  "yeul",
	Private:  false,
}

// CloudManagedNetworksList represents a list of CloudManagedNetworks
type CloudManagedNetworksList []*CloudManagedNetwork

// Identity returns the identity of the objects in the list.
func (o CloudManagedNetworksList) Identity() elemental.Identity {

	return CloudManagedNetworkIdentity
}

// Copy returns a pointer to a copy the CloudManagedNetworksList.
func (o CloudManagedNetworksList) Copy() elemental.Identifiables {

	copy := append(CloudManagedNetworksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CloudManagedNetworksList.
func (o CloudManagedNetworksList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudManagedNetworksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudManagedNetwork))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudManagedNetworksList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudManagedNetworksList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CloudManagedNetworksList converted to SparseCloudManagedNetworksList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudManagedNetworksList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudManagedNetworksList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudManagedNetwork)
	}

	return out
}

// Version returns the version of the content.
func (o CloudManagedNetworksList) Version() int {

	return 1
}

// CloudManagedNetwork represents the model of a cloudmanagednetwork
type CloudManagedNetwork struct {
	// Prisma Cloud API ID (matches Prisma Cloud API ID).
	APIID int `json:"APIID,omitempty" msgpack:"APIID,omitempty" bson:"apiid,omitempty" mapstructure:"APIID,omitempty"`

	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// ID of the host VPC.
	VPCID string `json:"VPCID,omitempty" msgpack:"VPCID,omitempty" bson:"vpcid,omitempty" mapstructure:"VPCID,omitempty"`

	// Cloud account ID associated with the entity (matches Prisma Cloud accountID).
	AccountID string `json:"accountId,omitempty" msgpack:"accountId,omitempty" bson:"accountid,omitempty" mapstructure:"accountId,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// Internal representation of object tags retrieved from the cloud provider.
	CloudTags []string `json:"cloudTags,omitempty" msgpack:"cloudTags,omitempty" bson:"cloudtags,omitempty" mapstructure:"cloudTags,omitempty"`

	// Cloud type of the entity.
	CloudType string `json:"cloudType,omitempty" msgpack:"cloudType,omitempty" bson:"cloudtype,omitempty" mapstructure:"cloudType,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Customer ID as identified by Prisma Cloud.
	CustomerID int `json:"customerID,omitempty" msgpack:"customerID,omitempty" bson:"customerid,omitempty" mapstructure:"customerID,omitempty"`

	// List of CIDRs.
	Entries []string `json:"entries" msgpack:"entries" bson:"entries" mapstructure:"entries,omitempty"`

	// The time that the object was first ingested.
	IngestionTime time.Time `json:"ingestionTime,omitempty" msgpack:"ingestionTime,omitempty" bson:"ingestiontime,omitempty" mapstructure:"ingestionTime,omitempty"`

	// Internal unique key for a resource to guarantee no overlaps at write.
	Key string `json:"-" msgpack:"-" bson:"key" mapstructure:"-,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the object (optional).
	Name string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// ID of the cloud provider object.
	NativeID string `json:"nativeID" msgpack:"nativeID" bson:"nativeid" mapstructure:"nativeID,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// A list of policy references associated with this cloud node.
	PolicyReferences []string `json:"policyReferences,omitempty" msgpack:"policyReferences,omitempty" bson:"policyreferences,omitempty" mapstructure:"policyReferences,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Region name associated with the entity.
	RegionName string `json:"regionName,omitempty" msgpack:"regionName,omitempty" bson:"regionname,omitempty" mapstructure:"regionName,omitempty"`

	// Prisma Cloud Resource ID.
	ResourceID int `json:"resourceID,omitempty" msgpack:"resourceID,omitempty" bson:"resourceid,omitempty" mapstructure:"resourceID,omitempty"`

	// The type of cloud managed network.
	Type CloudManagedNetworkTypeValue `json:"type" msgpack:"type" bson:"type" mapstructure:"type,omitempty"`

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

// NewCloudManagedNetwork returns a new *CloudManagedNetwork
func NewCloudManagedNetwork() *CloudManagedNetwork {

	return &CloudManagedNetwork{
		ModelVersion:     1,
		MigrationsLog:    map[string]string{},
		Annotations:      map[string][]string{},
		AssociatedTags:   []string{},
		CloudTags:        []string{},
		Entries:          []string{},
		PolicyReferences: []string{},
		NormalizedTags:   []string{},
		Type:             CloudManagedNetworkTypeEnterprise,
	}
}

// Identity returns the Identity of the object.
func (o *CloudManagedNetwork) Identity() elemental.Identity {

	return CloudManagedNetworkIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudManagedNetwork) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudManagedNetwork) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudManagedNetwork) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudManagedNetwork{}

	s.APIID = o.APIID
	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.VPCID = o.VPCID
	s.AccountID = o.AccountID
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CloudTags = o.CloudTags
	s.CloudType = o.CloudType
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.CustomerID = o.CustomerID
	s.Entries = o.Entries
	s.IngestionTime = o.IngestionTime
	s.Key = o.Key
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NativeID = o.NativeID
	s.NormalizedTags = o.NormalizedTags
	s.PolicyReferences = o.PolicyReferences
	s.Protected = o.Protected
	s.RegionName = o.RegionName
	s.ResourceID = o.ResourceID
	s.Type = o.Type
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudManagedNetwork) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudManagedNetwork{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.APIID = s.APIID
	o.ID = s.ID.Hex()
	o.VPCID = s.VPCID
	o.AccountID = s.AccountID
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.CloudTags = s.CloudTags
	o.CloudType = s.CloudType
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.CustomerID = s.CustomerID
	o.Entries = s.Entries
	o.IngestionTime = s.IngestionTime
	o.Key = s.Key
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NativeID = s.NativeID
	o.NormalizedTags = s.NormalizedTags
	o.PolicyReferences = s.PolicyReferences
	o.Protected = s.Protected
	o.RegionName = s.RegionName
	o.ResourceID = s.ResourceID
	o.Type = s.Type
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudManagedNetwork) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudManagedNetwork) BleveType() string {

	return "cloudmanagednetwork"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudManagedNetwork) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CloudManagedNetwork) Doc() string {

	return `A cloud managed network represents a set of enterprise subnets that can be used
in policies.`
}

func (o *CloudManagedNetwork) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAPIID returns the APIID of the receiver.
func (o *CloudManagedNetwork) GetAPIID() int {

	return o.APIID
}

// SetAPIID sets the property APIID of the receiver using the given value.
func (o *CloudManagedNetwork) SetAPIID(APIID int) {

	o.APIID = APIID
}

// GetVPCID returns the VPCID of the receiver.
func (o *CloudManagedNetwork) GetVPCID() string {

	return o.VPCID
}

// SetVPCID sets the property VPCID of the receiver using the given value.
func (o *CloudManagedNetwork) SetVPCID(VPCID string) {

	o.VPCID = VPCID
}

// GetAccountID returns the AccountID of the receiver.
func (o *CloudManagedNetwork) GetAccountID() string {

	return o.AccountID
}

// SetAccountID sets the property AccountID of the receiver using the given value.
func (o *CloudManagedNetwork) SetAccountID(accountID string) {

	o.AccountID = accountID
}

// GetAnnotations returns the Annotations of the receiver.
func (o *CloudManagedNetwork) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *CloudManagedNetwork) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *CloudManagedNetwork) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *CloudManagedNetwork) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCloudTags returns the CloudTags of the receiver.
func (o *CloudManagedNetwork) GetCloudTags() []string {

	return o.CloudTags
}

// SetCloudTags sets the property CloudTags of the receiver using the given value.
func (o *CloudManagedNetwork) SetCloudTags(cloudTags []string) {

	o.CloudTags = cloudTags
}

// GetCloudType returns the CloudType of the receiver.
func (o *CloudManagedNetwork) GetCloudType() string {

	return o.CloudType
}

// SetCloudType sets the property CloudType of the receiver using the given value.
func (o *CloudManagedNetwork) SetCloudType(cloudType string) {

	o.CloudType = cloudType
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *CloudManagedNetwork) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *CloudManagedNetwork) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *CloudManagedNetwork) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *CloudManagedNetwork) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetCustomerID returns the CustomerID of the receiver.
func (o *CloudManagedNetwork) GetCustomerID() int {

	return o.CustomerID
}

// SetCustomerID sets the property CustomerID of the receiver using the given value.
func (o *CloudManagedNetwork) SetCustomerID(customerID int) {

	o.CustomerID = customerID
}

// GetIngestionTime returns the IngestionTime of the receiver.
func (o *CloudManagedNetwork) GetIngestionTime() time.Time {

	return o.IngestionTime
}

// SetIngestionTime sets the property IngestionTime of the receiver using the given value.
func (o *CloudManagedNetwork) SetIngestionTime(ingestionTime time.Time) {

	o.IngestionTime = ingestionTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *CloudManagedNetwork) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *CloudManagedNetwork) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *CloudManagedNetwork) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *CloudManagedNetwork) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *CloudManagedNetwork) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *CloudManagedNetwork) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNativeID returns the NativeID of the receiver.
func (o *CloudManagedNetwork) GetNativeID() string {

	return o.NativeID
}

// SetNativeID sets the property NativeID of the receiver using the given value.
func (o *CloudManagedNetwork) SetNativeID(nativeID string) {

	o.NativeID = nativeID
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *CloudManagedNetwork) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *CloudManagedNetwork) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPolicyReferences returns the PolicyReferences of the receiver.
func (o *CloudManagedNetwork) GetPolicyReferences() []string {

	return o.PolicyReferences
}

// SetPolicyReferences sets the property PolicyReferences of the receiver using the given value.
func (o *CloudManagedNetwork) SetPolicyReferences(policyReferences []string) {

	o.PolicyReferences = policyReferences
}

// GetProtected returns the Protected of the receiver.
func (o *CloudManagedNetwork) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *CloudManagedNetwork) SetProtected(protected bool) {

	o.Protected = protected
}

// GetRegionName returns the RegionName of the receiver.
func (o *CloudManagedNetwork) GetRegionName() string {

	return o.RegionName
}

// SetRegionName sets the property RegionName of the receiver using the given value.
func (o *CloudManagedNetwork) SetRegionName(regionName string) {

	o.RegionName = regionName
}

// GetResourceID returns the ResourceID of the receiver.
func (o *CloudManagedNetwork) GetResourceID() int {

	return o.ResourceID
}

// SetResourceID sets the property ResourceID of the receiver using the given value.
func (o *CloudManagedNetwork) SetResourceID(resourceID int) {

	o.ResourceID = resourceID
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *CloudManagedNetwork) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *CloudManagedNetwork) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *CloudManagedNetwork) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *CloudManagedNetwork) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *CloudManagedNetwork) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *CloudManagedNetwork) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *CloudManagedNetwork) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *CloudManagedNetwork) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudManagedNetwork) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudManagedNetwork{
			APIID:                &o.APIID,
			ID:                   &o.ID,
			VPCID:                &o.VPCID,
			AccountID:            &o.AccountID,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			CloudTags:            &o.CloudTags,
			CloudType:            &o.CloudType,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			CustomerID:           &o.CustomerID,
			Entries:              &o.Entries,
			IngestionTime:        &o.IngestionTime,
			Key:                  &o.Key,
			MigrationsLog:        &o.MigrationsLog,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NativeID:             &o.NativeID,
			NormalizedTags:       &o.NormalizedTags,
			PolicyReferences:     &o.PolicyReferences,
			Protected:            &o.Protected,
			RegionName:           &o.RegionName,
			ResourceID:           &o.ResourceID,
			Type:                 &o.Type,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseCloudManagedNetwork{}
	for _, f := range fields {
		switch f {
		case "APIID":
			sp.APIID = &(o.APIID)
		case "ID":
			sp.ID = &(o.ID)
		case "VPCID":
			sp.VPCID = &(o.VPCID)
		case "accountID":
			sp.AccountID = &(o.AccountID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "cloudTags":
			sp.CloudTags = &(o.CloudTags)
		case "cloudType":
			sp.CloudType = &(o.CloudType)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "customerID":
			sp.CustomerID = &(o.CustomerID)
		case "entries":
			sp.Entries = &(o.Entries)
		case "ingestionTime":
			sp.IngestionTime = &(o.IngestionTime)
		case "key":
			sp.Key = &(o.Key)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "nativeID":
			sp.NativeID = &(o.NativeID)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "policyReferences":
			sp.PolicyReferences = &(o.PolicyReferences)
		case "protected":
			sp.Protected = &(o.Protected)
		case "regionName":
			sp.RegionName = &(o.RegionName)
		case "resourceID":
			sp.ResourceID = &(o.ResourceID)
		case "type":
			sp.Type = &(o.Type)
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

// Patch apply the non nil value of a *SparseCloudManagedNetwork to the object.
func (o *CloudManagedNetwork) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudManagedNetwork)
	if so.APIID != nil {
		o.APIID = *so.APIID
	}
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.VPCID != nil {
		o.VPCID = *so.VPCID
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
	if so.CloudTags != nil {
		o.CloudTags = *so.CloudTags
	}
	if so.CloudType != nil {
		o.CloudType = *so.CloudType
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.CustomerID != nil {
		o.CustomerID = *so.CustomerID
	}
	if so.Entries != nil {
		o.Entries = *so.Entries
	}
	if so.IngestionTime != nil {
		o.IngestionTime = *so.IngestionTime
	}
	if so.Key != nil {
		o.Key = *so.Key
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NativeID != nil {
		o.NativeID = *so.NativeID
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.PolicyReferences != nil {
		o.PolicyReferences = *so.PolicyReferences
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.RegionName != nil {
		o.RegionName = *so.RegionName
	}
	if so.ResourceID != nil {
		o.ResourceID = *so.ResourceID
	}
	if so.Type != nil {
		o.Type = *so.Type
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

// DeepCopy returns a deep copy if the CloudManagedNetwork.
func (o *CloudManagedNetwork) DeepCopy() *CloudManagedNetwork {

	if o == nil {
		return nil
	}

	out := &CloudManagedNetwork{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudManagedNetwork.
func (o *CloudManagedNetwork) DeepCopyInto(out *CloudManagedNetwork) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudManagedNetwork: %s", err))
	}

	*out = *target.(*CloudManagedNetwork)
}

// Validate valides the current information stored into the structure.
func (o *CloudManagedNetwork) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateOptionalCIDRList("entries", o.Entries); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("nativeID", o.NativeID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("nativeID", o.NativeID, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateNativeID("nativeID", o.NativeID); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("regionName", o.RegionName, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("type", string(o.Type)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Enterprise", "AWS", "GCP", "Custom"}, false); err != nil {
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
func (*CloudManagedNetwork) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudManagedNetworkAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudManagedNetworkLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudManagedNetwork) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudManagedNetworkAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudManagedNetwork) ValueForAttribute(name string) interface{} {

	switch name {
	case "APIID":
		return o.APIID
	case "ID":
		return o.ID
	case "VPCID":
		return o.VPCID
	case "accountID":
		return o.AccountID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "cloudTags":
		return o.CloudTags
	case "cloudType":
		return o.CloudType
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "customerID":
		return o.CustomerID
	case "entries":
		return o.Entries
	case "ingestionTime":
		return o.IngestionTime
	case "key":
		return o.Key
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "nativeID":
		return o.NativeID
	case "normalizedTags":
		return o.NormalizedTags
	case "policyReferences":
		return o.PolicyReferences
	case "protected":
		return o.Protected
	case "regionName":
		return o.RegionName
	case "resourceID":
		return o.ResourceID
	case "type":
		return o.Type
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

// CloudManagedNetworkAttributesMap represents the map of attribute for CloudManagedNetwork.
var CloudManagedNetworkAttributesMap = map[string]elemental.AttributeSpecification{
	"APIID": {
		AllowedChoices: []string{},
		BSONFieldName:  "apiid",
		ConvertedName:  "APIID",
		Description:    `Prisma Cloud API ID (matches Prisma Cloud API ID).`,
		Exposed:        true,
		Getter:         true,
		Name:           "APIID",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
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
	"VPCID": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcid",
		ConvertedName:  "VPCID",
		Description:    `ID of the host VPC.`,
		Exposed:        true,
		Getter:         true,
		Name:           "VPCID",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"AccountID": {
		AllowedChoices: []string{},
		BSONFieldName:  "accountid",
		ConvertedName:  "AccountID",
		Description:    `Cloud account ID associated with the entity (matches Prisma Cloud accountID).`,
		Exposed:        true,
		Getter:         true,
		Name:           "accountID",
		Setter:         true,
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
	"CloudTags": {
		AllowedChoices: []string{},
		BSONFieldName:  "cloudtags",
		ConvertedName:  "CloudTags",
		Description:    `Internal representation of object tags retrieved from the cloud provider.`,
		Exposed:        true,
		Getter:         true,
		Name:           "cloudTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"CloudType": {
		AllowedChoices: []string{},
		BSONFieldName:  "cloudtype",
		ConvertedName:  "CloudType",
		Description:    `Cloud type of the entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "cloudType",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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
	"CustomerID": {
		AllowedChoices: []string{},
		BSONFieldName:  "customerid",
		ConvertedName:  "CustomerID",
		Description:    `Customer ID as identified by Prisma Cloud.`,
		Exposed:        true,
		Getter:         true,
		Name:           "customerID",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
	"Entries": {
		AllowedChoices: []string{},
		BSONFieldName:  "entries",
		ConvertedName:  "Entries",
		Description:    `List of CIDRs.`,
		Exposed:        true,
		Name:           "entries",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"IngestionTime": {
		AllowedChoices: []string{},
		BSONFieldName:  "ingestiontime",
		ConvertedName:  "IngestionTime",
		Description:    `The time that the object was first ingested.`,
		Exposed:        true,
		Getter:         true,
		Name:           "ingestionTime",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Key": {
		AllowedChoices: []string{},
		BSONFieldName:  "key",
		ConvertedName:  "Key",
		Description:    `Internal unique key for a resource to guarantee no overlaps at write.`,
		Name:           "key",
		Stored:         true,
		Type:           "string",
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
	"Name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `Name of the object (optional).`,
		Exposed:        true,
		Getter:         true,
		Name:           "name",
		Orderable:      true,
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
	"NativeID": {
		AllowedChoices: []string{},
		BSONFieldName:  "nativeid",
		ConvertedName:  "NativeID",
		Description:    `ID of the cloud provider object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      256,
		Name:           "nativeID",
		Orderable:      true,
		Required:       true,
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
	"PolicyReferences": {
		AllowedChoices: []string{},
		BSONFieldName:  "policyreferences",
		ConvertedName:  "PolicyReferences",
		Description:    `A list of policy references associated with this cloud node.`,
		Exposed:        true,
		Getter:         true,
		Name:           "policyReferences",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
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
	"RegionName": {
		AllowedChoices: []string{},
		BSONFieldName:  "regionname",
		ConvertedName:  "RegionName",
		Description:    `Region name associated with the entity.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      256,
		Name:           "regionName",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ResourceID": {
		AllowedChoices: []string{},
		BSONFieldName:  "resourceid",
		ConvertedName:  "ResourceID",
		Description:    `Prisma Cloud Resource ID.`,
		Exposed:        true,
		Getter:         true,
		Name:           "resourceID",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
	"Type": {
		AllowedChoices: []string{"Enterprise", "AWS", "GCP", "Custom"},
		BSONFieldName:  "type",
		ConvertedName:  "Type",
		DefaultValue:   CloudManagedNetworkTypeEnterprise,
		Description:    `The type of cloud managed network.`,
		Exposed:        true,
		Name:           "type",
		Required:       true,
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

// CloudManagedNetworkLowerCaseAttributesMap represents the map of attribute for CloudManagedNetwork.
var CloudManagedNetworkLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"apiid": {
		AllowedChoices: []string{},
		BSONFieldName:  "apiid",
		ConvertedName:  "APIID",
		Description:    `Prisma Cloud API ID (matches Prisma Cloud API ID).`,
		Exposed:        true,
		Getter:         true,
		Name:           "APIID",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
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
	"vpcid": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcid",
		ConvertedName:  "VPCID",
		Description:    `ID of the host VPC.`,
		Exposed:        true,
		Getter:         true,
		Name:           "VPCID",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"accountid": {
		AllowedChoices: []string{},
		BSONFieldName:  "accountid",
		ConvertedName:  "AccountID",
		Description:    `Cloud account ID associated with the entity (matches Prisma Cloud accountID).`,
		Exposed:        true,
		Getter:         true,
		Name:           "accountID",
		Setter:         true,
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
	"cloudtags": {
		AllowedChoices: []string{},
		BSONFieldName:  "cloudtags",
		ConvertedName:  "CloudTags",
		Description:    `Internal representation of object tags retrieved from the cloud provider.`,
		Exposed:        true,
		Getter:         true,
		Name:           "cloudTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"cloudtype": {
		AllowedChoices: []string{},
		BSONFieldName:  "cloudtype",
		ConvertedName:  "CloudType",
		Description:    `Cloud type of the entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "cloudType",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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
	"customerid": {
		AllowedChoices: []string{},
		BSONFieldName:  "customerid",
		ConvertedName:  "CustomerID",
		Description:    `Customer ID as identified by Prisma Cloud.`,
		Exposed:        true,
		Getter:         true,
		Name:           "customerID",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
	"entries": {
		AllowedChoices: []string{},
		BSONFieldName:  "entries",
		ConvertedName:  "Entries",
		Description:    `List of CIDRs.`,
		Exposed:        true,
		Name:           "entries",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"ingestiontime": {
		AllowedChoices: []string{},
		BSONFieldName:  "ingestiontime",
		ConvertedName:  "IngestionTime",
		Description:    `The time that the object was first ingested.`,
		Exposed:        true,
		Getter:         true,
		Name:           "ingestionTime",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"key": {
		AllowedChoices: []string{},
		BSONFieldName:  "key",
		ConvertedName:  "Key",
		Description:    `Internal unique key for a resource to guarantee no overlaps at write.`,
		Name:           "key",
		Stored:         true,
		Type:           "string",
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
	"name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `Name of the object (optional).`,
		Exposed:        true,
		Getter:         true,
		Name:           "name",
		Orderable:      true,
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
	"nativeid": {
		AllowedChoices: []string{},
		BSONFieldName:  "nativeid",
		ConvertedName:  "NativeID",
		Description:    `ID of the cloud provider object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      256,
		Name:           "nativeID",
		Orderable:      true,
		Required:       true,
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
	"policyreferences": {
		AllowedChoices: []string{},
		BSONFieldName:  "policyreferences",
		ConvertedName:  "PolicyReferences",
		Description:    `A list of policy references associated with this cloud node.`,
		Exposed:        true,
		Getter:         true,
		Name:           "policyReferences",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
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
	"regionname": {
		AllowedChoices: []string{},
		BSONFieldName:  "regionname",
		ConvertedName:  "RegionName",
		Description:    `Region name associated with the entity.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      256,
		Name:           "regionName",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"resourceid": {
		AllowedChoices: []string{},
		BSONFieldName:  "resourceid",
		ConvertedName:  "ResourceID",
		Description:    `Prisma Cloud Resource ID.`,
		Exposed:        true,
		Getter:         true,
		Name:           "resourceID",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
	"type": {
		AllowedChoices: []string{"Enterprise", "AWS", "GCP", "Custom"},
		BSONFieldName:  "type",
		ConvertedName:  "Type",
		DefaultValue:   CloudManagedNetworkTypeEnterprise,
		Description:    `The type of cloud managed network.`,
		Exposed:        true,
		Name:           "type",
		Required:       true,
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

// SparseCloudManagedNetworksList represents a list of SparseCloudManagedNetworks
type SparseCloudManagedNetworksList []*SparseCloudManagedNetwork

// Identity returns the identity of the objects in the list.
func (o SparseCloudManagedNetworksList) Identity() elemental.Identity {

	return CloudManagedNetworkIdentity
}

// Copy returns a pointer to a copy the SparseCloudManagedNetworksList.
func (o SparseCloudManagedNetworksList) Copy() elemental.Identifiables {

	copy := append(SparseCloudManagedNetworksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudManagedNetworksList.
func (o SparseCloudManagedNetworksList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudManagedNetworksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudManagedNetwork))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudManagedNetworksList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudManagedNetworksList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCloudManagedNetworksList converted to CloudManagedNetworksList.
func (o SparseCloudManagedNetworksList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudManagedNetworksList) Version() int {

	return 1
}

// SparseCloudManagedNetwork represents the sparse version of a cloudmanagednetwork.
type SparseCloudManagedNetwork struct {
	// Prisma Cloud API ID (matches Prisma Cloud API ID).
	APIID *int `json:"APIID,omitempty" msgpack:"APIID,omitempty" bson:"apiid,omitempty" mapstructure:"APIID,omitempty"`

	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// ID of the host VPC.
	VPCID *string `json:"VPCID,omitempty" msgpack:"VPCID,omitempty" bson:"vpcid,omitempty" mapstructure:"VPCID,omitempty"`

	// Cloud account ID associated with the entity (matches Prisma Cloud accountID).
	AccountID *string `json:"accountId,omitempty" msgpack:"accountId,omitempty" bson:"accountid,omitempty" mapstructure:"accountId,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// Internal representation of object tags retrieved from the cloud provider.
	CloudTags *[]string `json:"cloudTags,omitempty" msgpack:"cloudTags,omitempty" bson:"cloudtags,omitempty" mapstructure:"cloudTags,omitempty"`

	// Cloud type of the entity.
	CloudType *string `json:"cloudType,omitempty" msgpack:"cloudType,omitempty" bson:"cloudtype,omitempty" mapstructure:"cloudType,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Customer ID as identified by Prisma Cloud.
	CustomerID *int `json:"customerID,omitempty" msgpack:"customerID,omitempty" bson:"customerid,omitempty" mapstructure:"customerID,omitempty"`

	// List of CIDRs.
	Entries *[]string `json:"entries,omitempty" msgpack:"entries,omitempty" bson:"entries,omitempty" mapstructure:"entries,omitempty"`

	// The time that the object was first ingested.
	IngestionTime *time.Time `json:"ingestionTime,omitempty" msgpack:"ingestionTime,omitempty" bson:"ingestiontime,omitempty" mapstructure:"ingestionTime,omitempty"`

	// Internal unique key for a resource to guarantee no overlaps at write.
	Key *string `json:"-" msgpack:"-" bson:"key,omitempty" mapstructure:"-,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the object (optional).
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// ID of the cloud provider object.
	NativeID *string `json:"nativeID,omitempty" msgpack:"nativeID,omitempty" bson:"nativeid,omitempty" mapstructure:"nativeID,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// A list of policy references associated with this cloud node.
	PolicyReferences *[]string `json:"policyReferences,omitempty" msgpack:"policyReferences,omitempty" bson:"policyreferences,omitempty" mapstructure:"policyReferences,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Region name associated with the entity.
	RegionName *string `json:"regionName,omitempty" msgpack:"regionName,omitempty" bson:"regionname,omitempty" mapstructure:"regionName,omitempty"`

	// Prisma Cloud Resource ID.
	ResourceID *int `json:"resourceID,omitempty" msgpack:"resourceID,omitempty" bson:"resourceid,omitempty" mapstructure:"resourceID,omitempty"`

	// The type of cloud managed network.
	Type *CloudManagedNetworkTypeValue `json:"type,omitempty" msgpack:"type,omitempty" bson:"type,omitempty" mapstructure:"type,omitempty"`

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

// NewSparseCloudManagedNetwork returns a new  SparseCloudManagedNetwork.
func NewSparseCloudManagedNetwork() *SparseCloudManagedNetwork {
	return &SparseCloudManagedNetwork{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudManagedNetwork) Identity() elemental.Identity {

	return CloudManagedNetworkIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudManagedNetwork) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudManagedNetwork) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudManagedNetwork) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudManagedNetwork{}

	if o.APIID != nil {
		s.APIID = o.APIID
	}
	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.VPCID != nil {
		s.VPCID = o.VPCID
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
	if o.CloudTags != nil {
		s.CloudTags = o.CloudTags
	}
	if o.CloudType != nil {
		s.CloudType = o.CloudType
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.CustomerID != nil {
		s.CustomerID = o.CustomerID
	}
	if o.Entries != nil {
		s.Entries = o.Entries
	}
	if o.IngestionTime != nil {
		s.IngestionTime = o.IngestionTime
	}
	if o.Key != nil {
		s.Key = o.Key
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.NativeID != nil {
		s.NativeID = o.NativeID
	}
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.PolicyReferences != nil {
		s.PolicyReferences = o.PolicyReferences
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.RegionName != nil {
		s.RegionName = o.RegionName
	}
	if o.ResourceID != nil {
		s.ResourceID = o.ResourceID
	}
	if o.Type != nil {
		s.Type = o.Type
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
func (o *SparseCloudManagedNetwork) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudManagedNetwork{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.APIID != nil {
		o.APIID = s.APIID
	}
	id := s.ID.Hex()
	o.ID = &id
	if s.VPCID != nil {
		o.VPCID = s.VPCID
	}
	if s.AccountID != nil {
		o.AccountID = s.AccountID
	}
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.CloudTags != nil {
		o.CloudTags = s.CloudTags
	}
	if s.CloudType != nil {
		o.CloudType = s.CloudType
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.CustomerID != nil {
		o.CustomerID = s.CustomerID
	}
	if s.Entries != nil {
		o.Entries = s.Entries
	}
	if s.IngestionTime != nil {
		o.IngestionTime = s.IngestionTime
	}
	if s.Key != nil {
		o.Key = s.Key
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.NativeID != nil {
		o.NativeID = s.NativeID
	}
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.PolicyReferences != nil {
		o.PolicyReferences = s.PolicyReferences
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.RegionName != nil {
		o.RegionName = s.RegionName
	}
	if s.ResourceID != nil {
		o.ResourceID = s.ResourceID
	}
	if s.Type != nil {
		o.Type = s.Type
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
func (o *SparseCloudManagedNetwork) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudManagedNetwork) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudManagedNetwork()
	if o.APIID != nil {
		out.APIID = *o.APIID
	}
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.VPCID != nil {
		out.VPCID = *o.VPCID
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
	if o.CloudTags != nil {
		out.CloudTags = *o.CloudTags
	}
	if o.CloudType != nil {
		out.CloudType = *o.CloudType
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.CustomerID != nil {
		out.CustomerID = *o.CustomerID
	}
	if o.Entries != nil {
		out.Entries = *o.Entries
	}
	if o.IngestionTime != nil {
		out.IngestionTime = *o.IngestionTime
	}
	if o.Key != nil {
		out.Key = *o.Key
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NativeID != nil {
		out.NativeID = *o.NativeID
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.PolicyReferences != nil {
		out.PolicyReferences = *o.PolicyReferences
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.RegionName != nil {
		out.RegionName = *o.RegionName
	}
	if o.ResourceID != nil {
		out.ResourceID = *o.ResourceID
	}
	if o.Type != nil {
		out.Type = *o.Type
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

// GetAPIID returns the APIID of the receiver.
func (o *SparseCloudManagedNetwork) GetAPIID() (out int) {

	if o.APIID == nil {
		return
	}

	return *o.APIID
}

// SetAPIID sets the property APIID of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetAPIID(APIID int) {

	o.APIID = &APIID
}

// GetVPCID returns the VPCID of the receiver.
func (o *SparseCloudManagedNetwork) GetVPCID() (out string) {

	if o.VPCID == nil {
		return
	}

	return *o.VPCID
}

// SetVPCID sets the property VPCID of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetVPCID(VPCID string) {

	o.VPCID = &VPCID
}

// GetAccountID returns the AccountID of the receiver.
func (o *SparseCloudManagedNetwork) GetAccountID() (out string) {

	if o.AccountID == nil {
		return
	}

	return *o.AccountID
}

// SetAccountID sets the property AccountID of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetAccountID(accountID string) {

	o.AccountID = &accountID
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseCloudManagedNetwork) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseCloudManagedNetwork) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCloudTags returns the CloudTags of the receiver.
func (o *SparseCloudManagedNetwork) GetCloudTags() (out []string) {

	if o.CloudTags == nil {
		return
	}

	return *o.CloudTags
}

// SetCloudTags sets the property CloudTags of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetCloudTags(cloudTags []string) {

	o.CloudTags = &cloudTags
}

// GetCloudType returns the CloudType of the receiver.
func (o *SparseCloudManagedNetwork) GetCloudType() (out string) {

	if o.CloudType == nil {
		return
	}

	return *o.CloudType
}

// SetCloudType sets the property CloudType of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetCloudType(cloudType string) {

	o.CloudType = &cloudType
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseCloudManagedNetwork) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseCloudManagedNetwork) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetCustomerID returns the CustomerID of the receiver.
func (o *SparseCloudManagedNetwork) GetCustomerID() (out int) {

	if o.CustomerID == nil {
		return
	}

	return *o.CustomerID
}

// SetCustomerID sets the property CustomerID of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetCustomerID(customerID int) {

	o.CustomerID = &customerID
}

// GetIngestionTime returns the IngestionTime of the receiver.
func (o *SparseCloudManagedNetwork) GetIngestionTime() (out time.Time) {

	if o.IngestionTime == nil {
		return
	}

	return *o.IngestionTime
}

// SetIngestionTime sets the property IngestionTime of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetIngestionTime(ingestionTime time.Time) {

	o.IngestionTime = &ingestionTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseCloudManagedNetwork) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseCloudManagedNetwork) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseCloudManagedNetwork) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNativeID returns the NativeID of the receiver.
func (o *SparseCloudManagedNetwork) GetNativeID() (out string) {

	if o.NativeID == nil {
		return
	}

	return *o.NativeID
}

// SetNativeID sets the property NativeID of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetNativeID(nativeID string) {

	o.NativeID = &nativeID
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseCloudManagedNetwork) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetPolicyReferences returns the PolicyReferences of the receiver.
func (o *SparseCloudManagedNetwork) GetPolicyReferences() (out []string) {

	if o.PolicyReferences == nil {
		return
	}

	return *o.PolicyReferences
}

// SetPolicyReferences sets the property PolicyReferences of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetPolicyReferences(policyReferences []string) {

	o.PolicyReferences = &policyReferences
}

// GetProtected returns the Protected of the receiver.
func (o *SparseCloudManagedNetwork) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetRegionName returns the RegionName of the receiver.
func (o *SparseCloudManagedNetwork) GetRegionName() (out string) {

	if o.RegionName == nil {
		return
	}

	return *o.RegionName
}

// SetRegionName sets the property RegionName of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetRegionName(regionName string) {

	o.RegionName = &regionName
}

// GetResourceID returns the ResourceID of the receiver.
func (o *SparseCloudManagedNetwork) GetResourceID() (out int) {

	if o.ResourceID == nil {
		return
	}

	return *o.ResourceID
}

// SetResourceID sets the property ResourceID of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetResourceID(resourceID int) {

	o.ResourceID = &resourceID
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseCloudManagedNetwork) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseCloudManagedNetwork) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseCloudManagedNetwork) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseCloudManagedNetwork) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseCloudManagedNetwork) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseCloudManagedNetwork.
func (o *SparseCloudManagedNetwork) DeepCopy() *SparseCloudManagedNetwork {

	if o == nil {
		return nil
	}

	out := &SparseCloudManagedNetwork{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudManagedNetwork.
func (o *SparseCloudManagedNetwork) DeepCopyInto(out *SparseCloudManagedNetwork) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudManagedNetwork: %s", err))
	}

	*out = *target.(*SparseCloudManagedNetwork)
}

type mongoAttributesCloudManagedNetwork struct {
	APIID                int                          `bson:"apiid,omitempty"`
	ID                   bson.ObjectId                `bson:"_id,omitempty"`
	VPCID                string                       `bson:"vpcid,omitempty"`
	AccountID            string                       `bson:"accountid,omitempty"`
	Annotations          map[string][]string          `bson:"annotations"`
	AssociatedTags       []string                     `bson:"associatedtags"`
	CloudTags            []string                     `bson:"cloudtags,omitempty"`
	CloudType            string                       `bson:"cloudtype,omitempty"`
	CreateIdempotencyKey string                       `bson:"createidempotencykey"`
	CreateTime           time.Time                    `bson:"createtime"`
	CustomerID           int                          `bson:"customerid,omitempty"`
	Entries              []string                     `bson:"entries"`
	IngestionTime        time.Time                    `bson:"ingestiontime,omitempty"`
	Key                  string                       `bson:"key"`
	MigrationsLog        map[string]string            `bson:"migrationslog,omitempty"`
	Name                 string                       `bson:"name,omitempty"`
	Namespace            string                       `bson:"namespace"`
	NativeID             string                       `bson:"nativeid"`
	NormalizedTags       []string                     `bson:"normalizedtags"`
	PolicyReferences     []string                     `bson:"policyreferences,omitempty"`
	Protected            bool                         `bson:"protected"`
	RegionName           string                       `bson:"regionname,omitempty"`
	ResourceID           int                          `bson:"resourceid,omitempty"`
	Type                 CloudManagedNetworkTypeValue `bson:"type"`
	UpdateIdempotencyKey string                       `bson:"updateidempotencykey"`
	UpdateTime           time.Time                    `bson:"updatetime"`
	ZHash                int                          `bson:"zhash"`
	Zone                 int                          `bson:"zone"`
}
type mongoAttributesSparseCloudManagedNetwork struct {
	APIID                *int                          `bson:"apiid,omitempty"`
	ID                   bson.ObjectId                 `bson:"_id,omitempty"`
	VPCID                *string                       `bson:"vpcid,omitempty"`
	AccountID            *string                       `bson:"accountid,omitempty"`
	Annotations          *map[string][]string          `bson:"annotations,omitempty"`
	AssociatedTags       *[]string                     `bson:"associatedtags,omitempty"`
	CloudTags            *[]string                     `bson:"cloudtags,omitempty"`
	CloudType            *string                       `bson:"cloudtype,omitempty"`
	CreateIdempotencyKey *string                       `bson:"createidempotencykey,omitempty"`
	CreateTime           *time.Time                    `bson:"createtime,omitempty"`
	CustomerID           *int                          `bson:"customerid,omitempty"`
	Entries              *[]string                     `bson:"entries,omitempty"`
	IngestionTime        *time.Time                    `bson:"ingestiontime,omitempty"`
	Key                  *string                       `bson:"key,omitempty"`
	MigrationsLog        *map[string]string            `bson:"migrationslog,omitempty"`
	Name                 *string                       `bson:"name,omitempty"`
	Namespace            *string                       `bson:"namespace,omitempty"`
	NativeID             *string                       `bson:"nativeid,omitempty"`
	NormalizedTags       *[]string                     `bson:"normalizedtags,omitempty"`
	PolicyReferences     *[]string                     `bson:"policyreferences,omitempty"`
	Protected            *bool                         `bson:"protected,omitempty"`
	RegionName           *string                       `bson:"regionname,omitempty"`
	ResourceID           *int                          `bson:"resourceid,omitempty"`
	Type                 *CloudManagedNetworkTypeValue `bson:"type,omitempty"`
	UpdateIdempotencyKey *string                       `bson:"updateidempotencykey,omitempty"`
	UpdateTime           *time.Time                    `bson:"updatetime,omitempty"`
	ZHash                *int                          `bson:"zhash,omitempty"`
	Zone                 *int                          `bson:"zone,omitempty"`
}