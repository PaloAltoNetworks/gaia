package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudPublicIpAddressIdentity represents the Identity of the object.
var CloudPublicIpAddressIdentity = elemental.Identity{
	Name:     "cloudpublicipaddress",
	Category: "cloudpublicipaddresses",
	Package:  "yeul",
	Private:  false,
}

// CloudPublicIpAddressList represents a list of CloudPublicIpAddress
type CloudPublicIpAddressList []*CloudPublicIpAddress

// Identity returns the identity of the objects in the list.
func (o CloudPublicIpAddressList) Identity() elemental.Identity {

	return CloudPublicIpAddressIdentity
}

// Copy returns a pointer to a copy the CloudPublicIpAddressList.
func (o CloudPublicIpAddressList) Copy() elemental.Identifiables {

	copy := append(CloudPublicIpAddressList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CloudPublicIpAddressList.
func (o CloudPublicIpAddressList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudPublicIpAddressList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudPublicIpAddress))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudPublicIpAddressList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudPublicIpAddressList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CloudPublicIpAddressList converted to SparseCloudPublicIpAddressList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudPublicIpAddressList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudPublicIpAddressList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudPublicIpAddress)
	}

	return out
}

// Version returns the version of the content.
func (o CloudPublicIpAddressList) Version() int {

	return 1
}

// CloudPublicIpAddress represents the model of a cloudpublicipaddress
type CloudPublicIpAddress struct {
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

	// The time that the object was first ingested.
	IngestionTime time.Time `json:"ingestionTime,omitempty" msgpack:"ingestionTime,omitempty" bson:"ingestiontime,omitempty" mapstructure:"ingestionTime,omitempty"`

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

	// PublicIpAddress related parameters.
	Parameters *CloudAddress `json:"parameters" msgpack:"parameters" bson:"parameters" mapstructure:"parameters,omitempty"`

	// A list of policy references associated with this cloud node.
	PolicyReferences []string `json:"policyReferences,omitempty" msgpack:"policyReferences,omitempty" bson:"policyreferences,omitempty" mapstructure:"policyReferences,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Region name associated with the entity.
	RegionName string `json:"regionName,omitempty" msgpack:"regionName,omitempty" bson:"regionname,omitempty" mapstructure:"regionName,omitempty"`

	// Prisma Cloud Resource ID.
	ResourceID int `json:"resourceID,omitempty" msgpack:"resourceID,omitempty" bson:"resourceid,omitempty" mapstructure:"resourceID,omitempty"`

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

// NewCloudPublicIpAddress returns a new *CloudPublicIpAddress
func NewCloudPublicIpAddress() *CloudPublicIpAddress {

	return &CloudPublicIpAddress{
		ModelVersion:     1,
		Annotations:      map[string][]string{},
		AssociatedTags:   []string{},
		CloudTags:        []string{},
		PolicyReferences: []string{},
		MigrationsLog:    map[string]string{},
		NormalizedTags:   []string{},
		Parameters:       NewCloudAddress(),
	}
}

// Identity returns the Identity of the object.
func (o *CloudPublicIpAddress) Identity() elemental.Identity {

	return CloudPublicIpAddressIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudPublicIpAddress) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudPublicIpAddress) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudPublicIpAddress) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudPublicIpAddress{}

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
	s.IngestionTime = o.IngestionTime
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NativeID = o.NativeID
	s.NormalizedTags = o.NormalizedTags
	s.Parameters = o.Parameters
	s.PolicyReferences = o.PolicyReferences
	s.Protected = o.Protected
	s.RegionName = o.RegionName
	s.ResourceID = o.ResourceID
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudPublicIpAddress) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudPublicIpAddress{}
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
	o.IngestionTime = s.IngestionTime
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NativeID = s.NativeID
	o.NormalizedTags = s.NormalizedTags
	o.Parameters = s.Parameters
	o.PolicyReferences = s.PolicyReferences
	o.Protected = s.Protected
	o.RegionName = s.RegionName
	o.ResourceID = s.ResourceID
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudPublicIpAddress) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudPublicIpAddress) BleveType() string {

	return "cloudpublicipaddress"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudPublicIpAddress) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CloudPublicIpAddress) Doc() string {

	return `A CloudPublicIpAddress represents a PublicIpAdress as defined in an Azure cloud
provider.`
}

func (o *CloudPublicIpAddress) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAPIID returns the APIID of the receiver.
func (o *CloudPublicIpAddress) GetAPIID() int {

	return o.APIID
}

// SetAPIID sets the property APIID of the receiver using the given value.
func (o *CloudPublicIpAddress) SetAPIID(APIID int) {

	o.APIID = APIID
}

// GetVPCID returns the VPCID of the receiver.
func (o *CloudPublicIpAddress) GetVPCID() string {

	return o.VPCID
}

// SetVPCID sets the property VPCID of the receiver using the given value.
func (o *CloudPublicIpAddress) SetVPCID(VPCID string) {

	o.VPCID = VPCID
}

// GetAccountID returns the AccountID of the receiver.
func (o *CloudPublicIpAddress) GetAccountID() string {

	return o.AccountID
}

// SetAccountID sets the property AccountID of the receiver using the given value.
func (o *CloudPublicIpAddress) SetAccountID(accountID string) {

	o.AccountID = accountID
}

// GetAnnotations returns the Annotations of the receiver.
func (o *CloudPublicIpAddress) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *CloudPublicIpAddress) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *CloudPublicIpAddress) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *CloudPublicIpAddress) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCloudTags returns the CloudTags of the receiver.
func (o *CloudPublicIpAddress) GetCloudTags() []string {

	return o.CloudTags
}

// SetCloudTags sets the property CloudTags of the receiver using the given value.
func (o *CloudPublicIpAddress) SetCloudTags(cloudTags []string) {

	o.CloudTags = cloudTags
}

// GetCloudType returns the CloudType of the receiver.
func (o *CloudPublicIpAddress) GetCloudType() string {

	return o.CloudType
}

// SetCloudType sets the property CloudType of the receiver using the given value.
func (o *CloudPublicIpAddress) SetCloudType(cloudType string) {

	o.CloudType = cloudType
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *CloudPublicIpAddress) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *CloudPublicIpAddress) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *CloudPublicIpAddress) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *CloudPublicIpAddress) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetCustomerID returns the CustomerID of the receiver.
func (o *CloudPublicIpAddress) GetCustomerID() int {

	return o.CustomerID
}

// SetCustomerID sets the property CustomerID of the receiver using the given value.
func (o *CloudPublicIpAddress) SetCustomerID(customerID int) {

	o.CustomerID = customerID
}

// GetIngestionTime returns the IngestionTime of the receiver.
func (o *CloudPublicIpAddress) GetIngestionTime() time.Time {

	return o.IngestionTime
}

// SetIngestionTime sets the property IngestionTime of the receiver using the given value.
func (o *CloudPublicIpAddress) SetIngestionTime(ingestionTime time.Time) {

	o.IngestionTime = ingestionTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *CloudPublicIpAddress) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *CloudPublicIpAddress) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *CloudPublicIpAddress) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *CloudPublicIpAddress) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *CloudPublicIpAddress) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *CloudPublicIpAddress) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNativeID returns the NativeID of the receiver.
func (o *CloudPublicIpAddress) GetNativeID() string {

	return o.NativeID
}

// SetNativeID sets the property NativeID of the receiver using the given value.
func (o *CloudPublicIpAddress) SetNativeID(nativeID string) {

	o.NativeID = nativeID
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *CloudPublicIpAddress) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *CloudPublicIpAddress) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPolicyReferences returns the PolicyReferences of the receiver.
func (o *CloudPublicIpAddress) GetPolicyReferences() []string {

	return o.PolicyReferences
}

// SetPolicyReferences sets the property PolicyReferences of the receiver using the given value.
func (o *CloudPublicIpAddress) SetPolicyReferences(policyReferences []string) {

	o.PolicyReferences = policyReferences
}

// GetProtected returns the Protected of the receiver.
func (o *CloudPublicIpAddress) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *CloudPublicIpAddress) SetProtected(protected bool) {

	o.Protected = protected
}

// GetRegionName returns the RegionName of the receiver.
func (o *CloudPublicIpAddress) GetRegionName() string {

	return o.RegionName
}

// SetRegionName sets the property RegionName of the receiver using the given value.
func (o *CloudPublicIpAddress) SetRegionName(regionName string) {

	o.RegionName = regionName
}

// GetResourceID returns the ResourceID of the receiver.
func (o *CloudPublicIpAddress) GetResourceID() int {

	return o.ResourceID
}

// SetResourceID sets the property ResourceID of the receiver using the given value.
func (o *CloudPublicIpAddress) SetResourceID(resourceID int) {

	o.ResourceID = resourceID
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *CloudPublicIpAddress) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *CloudPublicIpAddress) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *CloudPublicIpAddress) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *CloudPublicIpAddress) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *CloudPublicIpAddress) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *CloudPublicIpAddress) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *CloudPublicIpAddress) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *CloudPublicIpAddress) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudPublicIpAddress) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudPublicIpAddress{
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
			IngestionTime:        &o.IngestionTime,
			MigrationsLog:        &o.MigrationsLog,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NativeID:             &o.NativeID,
			NormalizedTags:       &o.NormalizedTags,
			Parameters:           o.Parameters,
			PolicyReferences:     &o.PolicyReferences,
			Protected:            &o.Protected,
			RegionName:           &o.RegionName,
			ResourceID:           &o.ResourceID,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseCloudPublicIpAddress{}
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
		case "ingestionTime":
			sp.IngestionTime = &(o.IngestionTime)
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
		case "parameters":
			sp.Parameters = o.Parameters
		case "policyReferences":
			sp.PolicyReferences = &(o.PolicyReferences)
		case "protected":
			sp.Protected = &(o.Protected)
		case "regionName":
			sp.RegionName = &(o.RegionName)
		case "resourceID":
			sp.ResourceID = &(o.ResourceID)
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

// Patch apply the non nil value of a *SparseCloudPublicIpAddress to the object.
func (o *CloudPublicIpAddress) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudPublicIpAddress)
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
	if so.IngestionTime != nil {
		o.IngestionTime = *so.IngestionTime
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
	if so.Parameters != nil {
		o.Parameters = so.Parameters
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

// DeepCopy returns a deep copy if the CloudPublicIpAddress.
func (o *CloudPublicIpAddress) DeepCopy() *CloudPublicIpAddress {

	if o == nil {
		return nil
	}

	out := &CloudPublicIpAddress{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudPublicIpAddress.
func (o *CloudPublicIpAddress) DeepCopyInto(out *CloudPublicIpAddress) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudPublicIpAddress: %s", err))
	}

	*out = *target.(*CloudPublicIpAddress)
}

// Validate valides the current information stored into the structure.
func (o *CloudPublicIpAddress) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
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

	if o.Parameters != nil {
		elemental.ResetDefaultForZeroValues(o.Parameters)
		if err := o.Parameters.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateMaximumLength("regionName", o.RegionName, 256, false); err != nil {
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
func (*CloudPublicIpAddress) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudPublicIpAddressAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudPublicIpAddressLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudPublicIpAddress) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudPublicIpAddressAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudPublicIpAddress) ValueForAttribute(name string) interface{} {

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
	case "ingestionTime":
		return o.IngestionTime
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
	case "parameters":
		return o.Parameters
	case "policyReferences":
		return o.PolicyReferences
	case "protected":
		return o.Protected
	case "regionName":
		return o.RegionName
	case "resourceID":
		return o.ResourceID
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

// CloudPublicIpAddressAttributesMap represents the map of attribute for CloudPublicIpAddress.
var CloudPublicIpAddressAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Parameters": {
		AllowedChoices: []string{},
		BSONFieldName:  "parameters",
		ConvertedName:  "Parameters",
		Description:    `PublicIpAddress related parameters.`,
		Exposed:        true,
		Name:           "parameters",
		Stored:         true,
		SubType:        "cloudaddress",
		Type:           "ref",
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

// CloudPublicIpAddressLowerCaseAttributesMap represents the map of attribute for CloudPublicIpAddress.
var CloudPublicIpAddressLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"parameters": {
		AllowedChoices: []string{},
		BSONFieldName:  "parameters",
		ConvertedName:  "Parameters",
		Description:    `PublicIpAddress related parameters.`,
		Exposed:        true,
		Name:           "parameters",
		Stored:         true,
		SubType:        "cloudaddress",
		Type:           "ref",
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

// SparseCloudPublicIpAddressList represents a list of SparseCloudPublicIpAddress
type SparseCloudPublicIpAddressList []*SparseCloudPublicIpAddress

// Identity returns the identity of the objects in the list.
func (o SparseCloudPublicIpAddressList) Identity() elemental.Identity {

	return CloudPublicIpAddressIdentity
}

// Copy returns a pointer to a copy the SparseCloudPublicIpAddressList.
func (o SparseCloudPublicIpAddressList) Copy() elemental.Identifiables {

	copy := append(SparseCloudPublicIpAddressList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudPublicIpAddressList.
func (o SparseCloudPublicIpAddressList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudPublicIpAddressList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudPublicIpAddress))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudPublicIpAddressList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudPublicIpAddressList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCloudPublicIpAddressList converted to CloudPublicIpAddressList.
func (o SparseCloudPublicIpAddressList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudPublicIpAddressList) Version() int {

	return 1
}

// SparseCloudPublicIpAddress represents the sparse version of a cloudpublicipaddress.
type SparseCloudPublicIpAddress struct {
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

	// The time that the object was first ingested.
	IngestionTime *time.Time `json:"ingestionTime,omitempty" msgpack:"ingestionTime,omitempty" bson:"ingestiontime,omitempty" mapstructure:"ingestionTime,omitempty"`

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

	// PublicIpAddress related parameters.
	Parameters *CloudAddress `json:"parameters,omitempty" msgpack:"parameters,omitempty" bson:"parameters,omitempty" mapstructure:"parameters,omitempty"`

	// A list of policy references associated with this cloud node.
	PolicyReferences *[]string `json:"policyReferences,omitempty" msgpack:"policyReferences,omitempty" bson:"policyreferences,omitempty" mapstructure:"policyReferences,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Region name associated with the entity.
	RegionName *string `json:"regionName,omitempty" msgpack:"regionName,omitempty" bson:"regionname,omitempty" mapstructure:"regionName,omitempty"`

	// Prisma Cloud Resource ID.
	ResourceID *int `json:"resourceID,omitempty" msgpack:"resourceID,omitempty" bson:"resourceid,omitempty" mapstructure:"resourceID,omitempty"`

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

// NewSparseCloudPublicIpAddress returns a new  SparseCloudPublicIpAddress.
func NewSparseCloudPublicIpAddress() *SparseCloudPublicIpAddress {
	return &SparseCloudPublicIpAddress{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudPublicIpAddress) Identity() elemental.Identity {

	return CloudPublicIpAddressIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudPublicIpAddress) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudPublicIpAddress) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudPublicIpAddress) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudPublicIpAddress{}

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
	if o.IngestionTime != nil {
		s.IngestionTime = o.IngestionTime
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
	if o.Parameters != nil {
		s.Parameters = o.Parameters
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
func (o *SparseCloudPublicIpAddress) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudPublicIpAddress{}
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
	if s.IngestionTime != nil {
		o.IngestionTime = s.IngestionTime
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
	if s.Parameters != nil {
		o.Parameters = s.Parameters
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
func (o *SparseCloudPublicIpAddress) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudPublicIpAddress) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudPublicIpAddress()
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
	if o.IngestionTime != nil {
		out.IngestionTime = *o.IngestionTime
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
	if o.Parameters != nil {
		out.Parameters = o.Parameters
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
func (o *SparseCloudPublicIpAddress) GetAPIID() (out int) {

	if o.APIID == nil {
		return
	}

	return *o.APIID
}

// SetAPIID sets the property APIID of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetAPIID(APIID int) {

	o.APIID = &APIID
}

// GetVPCID returns the VPCID of the receiver.
func (o *SparseCloudPublicIpAddress) GetVPCID() (out string) {

	if o.VPCID == nil {
		return
	}

	return *o.VPCID
}

// SetVPCID sets the property VPCID of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetVPCID(VPCID string) {

	o.VPCID = &VPCID
}

// GetAccountID returns the AccountID of the receiver.
func (o *SparseCloudPublicIpAddress) GetAccountID() (out string) {

	if o.AccountID == nil {
		return
	}

	return *o.AccountID
}

// SetAccountID sets the property AccountID of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetAccountID(accountID string) {

	o.AccountID = &accountID
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseCloudPublicIpAddress) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseCloudPublicIpAddress) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCloudTags returns the CloudTags of the receiver.
func (o *SparseCloudPublicIpAddress) GetCloudTags() (out []string) {

	if o.CloudTags == nil {
		return
	}

	return *o.CloudTags
}

// SetCloudTags sets the property CloudTags of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetCloudTags(cloudTags []string) {

	o.CloudTags = &cloudTags
}

// GetCloudType returns the CloudType of the receiver.
func (o *SparseCloudPublicIpAddress) GetCloudType() (out string) {

	if o.CloudType == nil {
		return
	}

	return *o.CloudType
}

// SetCloudType sets the property CloudType of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetCloudType(cloudType string) {

	o.CloudType = &cloudType
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseCloudPublicIpAddress) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseCloudPublicIpAddress) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetCustomerID returns the CustomerID of the receiver.
func (o *SparseCloudPublicIpAddress) GetCustomerID() (out int) {

	if o.CustomerID == nil {
		return
	}

	return *o.CustomerID
}

// SetCustomerID sets the property CustomerID of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetCustomerID(customerID int) {

	o.CustomerID = &customerID
}

// GetIngestionTime returns the IngestionTime of the receiver.
func (o *SparseCloudPublicIpAddress) GetIngestionTime() (out time.Time) {

	if o.IngestionTime == nil {
		return
	}

	return *o.IngestionTime
}

// SetIngestionTime sets the property IngestionTime of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetIngestionTime(ingestionTime time.Time) {

	o.IngestionTime = &ingestionTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseCloudPublicIpAddress) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseCloudPublicIpAddress) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseCloudPublicIpAddress) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNativeID returns the NativeID of the receiver.
func (o *SparseCloudPublicIpAddress) GetNativeID() (out string) {

	if o.NativeID == nil {
		return
	}

	return *o.NativeID
}

// SetNativeID sets the property NativeID of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetNativeID(nativeID string) {

	o.NativeID = &nativeID
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseCloudPublicIpAddress) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetPolicyReferences returns the PolicyReferences of the receiver.
func (o *SparseCloudPublicIpAddress) GetPolicyReferences() (out []string) {

	if o.PolicyReferences == nil {
		return
	}

	return *o.PolicyReferences
}

// SetPolicyReferences sets the property PolicyReferences of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetPolicyReferences(policyReferences []string) {

	o.PolicyReferences = &policyReferences
}

// GetProtected returns the Protected of the receiver.
func (o *SparseCloudPublicIpAddress) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetRegionName returns the RegionName of the receiver.
func (o *SparseCloudPublicIpAddress) GetRegionName() (out string) {

	if o.RegionName == nil {
		return
	}

	return *o.RegionName
}

// SetRegionName sets the property RegionName of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetRegionName(regionName string) {

	o.RegionName = &regionName
}

// GetResourceID returns the ResourceID of the receiver.
func (o *SparseCloudPublicIpAddress) GetResourceID() (out int) {

	if o.ResourceID == nil {
		return
	}

	return *o.ResourceID
}

// SetResourceID sets the property ResourceID of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetResourceID(resourceID int) {

	o.ResourceID = &resourceID
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseCloudPublicIpAddress) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseCloudPublicIpAddress) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseCloudPublicIpAddress) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseCloudPublicIpAddress) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseCloudPublicIpAddress) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseCloudPublicIpAddress.
func (o *SparseCloudPublicIpAddress) DeepCopy() *SparseCloudPublicIpAddress {

	if o == nil {
		return nil
	}

	out := &SparseCloudPublicIpAddress{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudPublicIpAddress.
func (o *SparseCloudPublicIpAddress) DeepCopyInto(out *SparseCloudPublicIpAddress) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudPublicIpAddress: %s", err))
	}

	*out = *target.(*SparseCloudPublicIpAddress)
}

type mongoAttributesCloudPublicIpAddress struct {
	APIID                int                 `bson:"apiid,omitempty"`
	ID                   bson.ObjectId       `bson:"_id,omitempty"`
	VPCID                string              `bson:"vpcid,omitempty"`
	AccountID            string              `bson:"accountid,omitempty"`
	Annotations          map[string][]string `bson:"annotations"`
	AssociatedTags       []string            `bson:"associatedtags"`
	CloudTags            []string            `bson:"cloudtags,omitempty"`
	CloudType            string              `bson:"cloudtype,omitempty"`
	CreateIdempotencyKey string              `bson:"createidempotencykey"`
	CreateTime           time.Time           `bson:"createtime"`
	CustomerID           int                 `bson:"customerid,omitempty"`
	IngestionTime        time.Time           `bson:"ingestiontime,omitempty"`
	MigrationsLog        map[string]string   `bson:"migrationslog,omitempty"`
	Name                 string              `bson:"name,omitempty"`
	Namespace            string              `bson:"namespace"`
	NativeID             string              `bson:"nativeid"`
	NormalizedTags       []string            `bson:"normalizedtags"`
	Parameters           *CloudAddress       `bson:"parameters"`
	PolicyReferences     []string            `bson:"policyreferences,omitempty"`
	Protected            bool                `bson:"protected"`
	RegionName           string              `bson:"regionname,omitempty"`
	ResourceID           int                 `bson:"resourceid,omitempty"`
	UpdateIdempotencyKey string              `bson:"updateidempotencykey"`
	UpdateTime           time.Time           `bson:"updatetime"`
	ZHash                int                 `bson:"zhash"`
	Zone                 int                 `bson:"zone"`
}
type mongoAttributesSparseCloudPublicIpAddress struct {
	APIID                *int                 `bson:"apiid,omitempty"`
	ID                   bson.ObjectId        `bson:"_id,omitempty"`
	VPCID                *string              `bson:"vpcid,omitempty"`
	AccountID            *string              `bson:"accountid,omitempty"`
	Annotations          *map[string][]string `bson:"annotations,omitempty"`
	AssociatedTags       *[]string            `bson:"associatedtags,omitempty"`
	CloudTags            *[]string            `bson:"cloudtags,omitempty"`
	CloudType            *string              `bson:"cloudtype,omitempty"`
	CreateIdempotencyKey *string              `bson:"createidempotencykey,omitempty"`
	CreateTime           *time.Time           `bson:"createtime,omitempty"`
	CustomerID           *int                 `bson:"customerid,omitempty"`
	IngestionTime        *time.Time           `bson:"ingestiontime,omitempty"`
	MigrationsLog        *map[string]string   `bson:"migrationslog,omitempty"`
	Name                 *string              `bson:"name,omitempty"`
	Namespace            *string              `bson:"namespace,omitempty"`
	NativeID             *string              `bson:"nativeid,omitempty"`
	NormalizedTags       *[]string            `bson:"normalizedtags,omitempty"`
	Parameters           *CloudAddress        `bson:"parameters,omitempty"`
	PolicyReferences     *[]string            `bson:"policyreferences,omitempty"`
	Protected            *bool                `bson:"protected,omitempty"`
	RegionName           *string              `bson:"regionname,omitempty"`
	ResourceID           *int                 `bson:"resourceid,omitempty"`
	UpdateIdempotencyKey *string              `bson:"updateidempotencykey,omitempty"`
	UpdateTime           *time.Time           `bson:"updatetime,omitempty"`
	ZHash                *int                 `bson:"zhash,omitempty"`
	Zone                 *int                 `bson:"zone,omitempty"`
}
