package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudVPCIdentity represents the Identity of the object.
var CloudVPCIdentity = elemental.Identity{
	Name:     "cloudvpc",
	Category: "cloudvpcs",
	Package:  "yeul",
	Private:  false,
}

// CloudVPCsList represents a list of CloudVPCs
type CloudVPCsList []*CloudVPC

// Identity returns the identity of the objects in the list.
func (o CloudVPCsList) Identity() elemental.Identity {

	return CloudVPCIdentity
}

// Copy returns a pointer to a copy the CloudVPCsList.
func (o CloudVPCsList) Copy() elemental.Identifiables {

	copy := append(CloudVPCsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CloudVPCsList.
func (o CloudVPCsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudVPCsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudVPC))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudVPCsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudVPCsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CloudVPCsList converted to SparseCloudVPCsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudVPCsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudVPCsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudVPC)
	}

	return out
}

// Version returns the version of the content.
func (o CloudVPCsList) Version() int {

	return 1
}

// CloudVPC represents the model of a cloudvpc
type CloudVPC struct {
	// Prisma Cloud API ID (matches Prisma Cloud API ID).
	APIID int `json:"APIID" msgpack:"APIID" bson:"apiid" mapstructure:"APIID,omitempty"`

	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Cloud account ID associated with the entity (matches Prisma Cloud accountID).
	AccountID string `json:"accountId" msgpack:"accountId" bson:"accountid" mapstructure:"accountId,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// Internal representation of object tags retrieved from the cloud provider.
	CloudTags []string `json:"cloudTags" msgpack:"cloudTags" bson:"cloudtags" mapstructure:"cloudTags,omitempty"`

	// Cloud type of the entity.
	CloudType string `json:"cloudType" msgpack:"cloudType" bson:"cloudtype" mapstructure:"cloudType,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Customer ID as identified by Prisma Cloud.
	CustomerID int `json:"customerID" msgpack:"customerID" bson:"customerid" mapstructure:"customerID,omitempty"`

	// The time that the object was first ingested.
	IngestionTime time.Time `json:"ingestionTime" msgpack:"ingestionTime" bson:"ingestiontime" mapstructure:"ingestionTime,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the object (optional).
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// ID of the cloud provider object.
	NativeID string `json:"nativeID" msgpack:"nativeID" bson:"nativeid" mapstructure:"nativeID,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// VPC related parameters.
	Parameters *CloudVPCData `json:"parameters" msgpack:"parameters" bson:"parameters" mapstructure:"parameters,omitempty"`

	// A list of policy references associated with this cloud node.
	PolicyReferences []string `json:"policyReferences" msgpack:"policyReferences" bson:"policyreferences" mapstructure:"policyReferences,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Region name associated with the entity.
	RegionName string `json:"regionName" msgpack:"regionName" bson:"regionname" mapstructure:"regionName,omitempty"`

	// Prisma Cloud Resource ID.
	ResourceID int `json:"resourceID" msgpack:"resourceID" bson:"resourceid" mapstructure:"resourceID,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// The time that the object was updated.
	UpdatedTime time.Time `json:"updatedTime" msgpack:"updatedTime" bson:"updatedtime" mapstructure:"updatedTime,omitempty"`

	// ID of the host VPC.
	VpcID string `json:"vpcID" msgpack:"vpcID" bson:"vpcid" mapstructure:"vpcID,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudVPC returns a new *CloudVPC
func NewCloudVPC() *CloudVPC {

	return &CloudVPC{
		ModelVersion:     1,
		Annotations:      map[string][]string{},
		AssociatedTags:   []string{},
		CloudTags:        []string{},
		NormalizedTags:   []string{},
		Parameters:       NewCloudVPCData(),
		PolicyReferences: []string{},
		MigrationsLog:    map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *CloudVPC) Identity() elemental.Identity {

	return CloudVPCIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudVPC) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudVPC) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudVPC) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudVPC{}

	s.APIID = o.APIID
	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.AccountID = o.AccountID
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CloudTags = o.CloudTags
	s.CloudType = o.CloudType
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
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
	s.UpdatedTime = o.UpdatedTime
	s.VpcID = o.VpcID
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudVPC) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudVPC{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.APIID = s.APIID
	o.ID = s.ID.Hex()
	o.AccountID = s.AccountID
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.CloudTags = s.CloudTags
	o.CloudType = s.CloudType
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
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
	o.UpdatedTime = s.UpdatedTime
	o.VpcID = s.VpcID
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudVPC) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudVPC) BleveType() string {

	return "cloudvpc"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudVPC) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CloudVPC) Doc() string {

	return `A CloudVPC represents a VPC as defined in an cloud provider (AWS/Azure/GCP etc).
The VPC is essentially an L3 routing domain with at least one subnet attached
and it defines an isolated network.`
}

func (o *CloudVPC) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAPIID returns the APIID of the receiver.
func (o *CloudVPC) GetAPIID() int {

	return o.APIID
}

// SetAPIID sets the property APIID of the receiver using the given value.
func (o *CloudVPC) SetAPIID(APIID int) {

	o.APIID = APIID
}

// GetAccountID returns the AccountID of the receiver.
func (o *CloudVPC) GetAccountID() string {

	return o.AccountID
}

// SetAccountID sets the property AccountID of the receiver using the given value.
func (o *CloudVPC) SetAccountID(accountID string) {

	o.AccountID = accountID
}

// GetAnnotations returns the Annotations of the receiver.
func (o *CloudVPC) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *CloudVPC) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *CloudVPC) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *CloudVPC) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCloudTags returns the CloudTags of the receiver.
func (o *CloudVPC) GetCloudTags() []string {

	return o.CloudTags
}

// SetCloudTags sets the property CloudTags of the receiver using the given value.
func (o *CloudVPC) SetCloudTags(cloudTags []string) {

	o.CloudTags = cloudTags
}

// GetCloudType returns the CloudType of the receiver.
func (o *CloudVPC) GetCloudType() string {

	return o.CloudType
}

// SetCloudType sets the property CloudType of the receiver using the given value.
func (o *CloudVPC) SetCloudType(cloudType string) {

	o.CloudType = cloudType
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *CloudVPC) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *CloudVPC) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCustomerID returns the CustomerID of the receiver.
func (o *CloudVPC) GetCustomerID() int {

	return o.CustomerID
}

// SetCustomerID sets the property CustomerID of the receiver using the given value.
func (o *CloudVPC) SetCustomerID(customerID int) {

	o.CustomerID = customerID
}

// GetIngestionTime returns the IngestionTime of the receiver.
func (o *CloudVPC) GetIngestionTime() time.Time {

	return o.IngestionTime
}

// SetIngestionTime sets the property IngestionTime of the receiver using the given value.
func (o *CloudVPC) SetIngestionTime(ingestionTime time.Time) {

	o.IngestionTime = ingestionTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *CloudVPC) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *CloudVPC) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *CloudVPC) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *CloudVPC) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *CloudVPC) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *CloudVPC) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNativeID returns the NativeID of the receiver.
func (o *CloudVPC) GetNativeID() string {

	return o.NativeID
}

// SetNativeID sets the property NativeID of the receiver using the given value.
func (o *CloudVPC) SetNativeID(nativeID string) {

	o.NativeID = nativeID
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *CloudVPC) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *CloudVPC) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPolicyReferences returns the PolicyReferences of the receiver.
func (o *CloudVPC) GetPolicyReferences() []string {

	return o.PolicyReferences
}

// SetPolicyReferences sets the property PolicyReferences of the receiver using the given value.
func (o *CloudVPC) SetPolicyReferences(policyReferences []string) {

	o.PolicyReferences = policyReferences
}

// GetProtected returns the Protected of the receiver.
func (o *CloudVPC) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *CloudVPC) SetProtected(protected bool) {

	o.Protected = protected
}

// GetRegionName returns the RegionName of the receiver.
func (o *CloudVPC) GetRegionName() string {

	return o.RegionName
}

// SetRegionName sets the property RegionName of the receiver using the given value.
func (o *CloudVPC) SetRegionName(regionName string) {

	o.RegionName = regionName
}

// GetResourceID returns the ResourceID of the receiver.
func (o *CloudVPC) GetResourceID() int {

	return o.ResourceID
}

// SetResourceID sets the property ResourceID of the receiver using the given value.
func (o *CloudVPC) SetResourceID(resourceID int) {

	o.ResourceID = resourceID
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *CloudVPC) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *CloudVPC) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdatedTime returns the UpdatedTime of the receiver.
func (o *CloudVPC) GetUpdatedTime() time.Time {

	return o.UpdatedTime
}

// SetUpdatedTime sets the property UpdatedTime of the receiver using the given value.
func (o *CloudVPC) SetUpdatedTime(updatedTime time.Time) {

	o.UpdatedTime = updatedTime
}

// GetVpcID returns the VpcID of the receiver.
func (o *CloudVPC) GetVpcID() string {

	return o.VpcID
}

// SetVpcID sets the property VpcID of the receiver using the given value.
func (o *CloudVPC) SetVpcID(vpcID string) {

	o.VpcID = vpcID
}

// GetZHash returns the ZHash of the receiver.
func (o *CloudVPC) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *CloudVPC) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *CloudVPC) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *CloudVPC) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudVPC) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudVPC{
			APIID:                &o.APIID,
			ID:                   &o.ID,
			AccountID:            &o.AccountID,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			CloudTags:            &o.CloudTags,
			CloudType:            &o.CloudType,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
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
			UpdatedTime:          &o.UpdatedTime,
			VpcID:                &o.VpcID,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseCloudVPC{}
	for _, f := range fields {
		switch f {
		case "APIID":
			sp.APIID = &(o.APIID)
		case "ID":
			sp.ID = &(o.ID)
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
		case "updatedTime":
			sp.UpdatedTime = &(o.UpdatedTime)
		case "vpcID":
			sp.VpcID = &(o.VpcID)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCloudVPC to the object.
func (o *CloudVPC) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudVPC)
	if so.APIID != nil {
		o.APIID = *so.APIID
	}
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
	if so.CloudTags != nil {
		o.CloudTags = *so.CloudTags
	}
	if so.CloudType != nil {
		o.CloudType = *so.CloudType
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
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
	if so.UpdatedTime != nil {
		o.UpdatedTime = *so.UpdatedTime
	}
	if so.VpcID != nil {
		o.VpcID = *so.VpcID
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the CloudVPC.
func (o *CloudVPC) DeepCopy() *CloudVPC {

	if o == nil {
		return nil
	}

	out := &CloudVPC{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudVPC.
func (o *CloudVPC) DeepCopyInto(out *CloudVPC) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudVPC: %s", err))
	}

	*out = *target.(*CloudVPC)
}

// Validate valides the current information stored into the structure.
func (o *CloudVPC) Validate() error {

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
func (*CloudVPC) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudVPCAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudVPCLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudVPC) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudVPCAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudVPC) ValueForAttribute(name string) interface{} {

	switch name {
	case "APIID":
		return o.APIID
	case "ID":
		return o.ID
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
	case "updatedTime":
		return o.UpdatedTime
	case "vpcID":
		return o.VpcID
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// CloudVPCAttributesMap represents the map of attribute for CloudVPC.
var CloudVPCAttributesMap = map[string]elemental.AttributeSpecification{
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
	"AccountID": {
		AllowedChoices: []string{},
		BSONFieldName:  "accountid",
		ConvertedName:  "AccountID",
		Description:    `Cloud account ID associated with the entity (matches Prisma Cloud accountID).`,
		Exposed:        true,
		Filterable:     true,
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
		Filterable:     true,
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
		Filterable:     true,
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
	"CustomerID": {
		AllowedChoices: []string{},
		BSONFieldName:  "customerid",
		ConvertedName:  "CustomerID",
		Description:    `Customer ID as identified by Prisma Cloud.`,
		Exposed:        true,
		Filterable:     true,
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
		Filterable:     true,
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
		Filterable:     true,
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
		Description:    `VPC related parameters.`,
		Exposed:        true,
		Name:           "parameters",
		Stored:         true,
		SubType:        "cloudvpcdata",
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
		Filterable:     true,
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
	"UpdatedTime": {
		AllowedChoices: []string{},
		BSONFieldName:  "updatedtime",
		ConvertedName:  "UpdatedTime",
		Description:    `The time that the object was updated.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updatedTime",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"VpcID": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcid",
		ConvertedName:  "VpcID",
		Description:    `ID of the host VPC.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "vpcID",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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

// CloudVPCLowerCaseAttributesMap represents the map of attribute for CloudVPC.
var CloudVPCLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"accountid": {
		AllowedChoices: []string{},
		BSONFieldName:  "accountid",
		ConvertedName:  "AccountID",
		Description:    `Cloud account ID associated with the entity (matches Prisma Cloud accountID).`,
		Exposed:        true,
		Filterable:     true,
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
		Filterable:     true,
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
		Filterable:     true,
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
	"customerid": {
		AllowedChoices: []string{},
		BSONFieldName:  "customerid",
		ConvertedName:  "CustomerID",
		Description:    `Customer ID as identified by Prisma Cloud.`,
		Exposed:        true,
		Filterable:     true,
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
		Filterable:     true,
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
		Filterable:     true,
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
		Description:    `VPC related parameters.`,
		Exposed:        true,
		Name:           "parameters",
		Stored:         true,
		SubType:        "cloudvpcdata",
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
		Filterable:     true,
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
	"updatedtime": {
		AllowedChoices: []string{},
		BSONFieldName:  "updatedtime",
		ConvertedName:  "UpdatedTime",
		Description:    `The time that the object was updated.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updatedTime",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"vpcid": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcid",
		ConvertedName:  "VpcID",
		Description:    `ID of the host VPC.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "vpcID",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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

// SparseCloudVPCsList represents a list of SparseCloudVPCs
type SparseCloudVPCsList []*SparseCloudVPC

// Identity returns the identity of the objects in the list.
func (o SparseCloudVPCsList) Identity() elemental.Identity {

	return CloudVPCIdentity
}

// Copy returns a pointer to a copy the SparseCloudVPCsList.
func (o SparseCloudVPCsList) Copy() elemental.Identifiables {

	copy := append(SparseCloudVPCsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudVPCsList.
func (o SparseCloudVPCsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudVPCsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudVPC))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudVPCsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudVPCsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCloudVPCsList converted to CloudVPCsList.
func (o SparseCloudVPCsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudVPCsList) Version() int {

	return 1
}

// SparseCloudVPC represents the sparse version of a cloudvpc.
type SparseCloudVPC struct {
	// Prisma Cloud API ID (matches Prisma Cloud API ID).
	APIID *int `json:"APIID,omitempty" msgpack:"APIID,omitempty" bson:"apiid,omitempty" mapstructure:"APIID,omitempty"`

	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

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

	// VPC related parameters.
	Parameters *CloudVPCData `json:"parameters,omitempty" msgpack:"parameters,omitempty" bson:"parameters,omitempty" mapstructure:"parameters,omitempty"`

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

	// The time that the object was updated.
	UpdatedTime *time.Time `json:"updatedTime,omitempty" msgpack:"updatedTime,omitempty" bson:"updatedtime,omitempty" mapstructure:"updatedTime,omitempty"`

	// ID of the host VPC.
	VpcID *string `json:"vpcID,omitempty" msgpack:"vpcID,omitempty" bson:"vpcid,omitempty" mapstructure:"vpcID,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCloudVPC returns a new  SparseCloudVPC.
func NewSparseCloudVPC() *SparseCloudVPC {
	return &SparseCloudVPC{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudVPC) Identity() elemental.Identity {

	return CloudVPCIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudVPC) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudVPC) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudVPC) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudVPC{}

	if o.APIID != nil {
		s.APIID = o.APIID
	}
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
	if o.CloudTags != nil {
		s.CloudTags = o.CloudTags
	}
	if o.CloudType != nil {
		s.CloudType = o.CloudType
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
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
	if o.UpdatedTime != nil {
		s.UpdatedTime = o.UpdatedTime
	}
	if o.VpcID != nil {
		s.VpcID = o.VpcID
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
func (o *SparseCloudVPC) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudVPC{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.APIID != nil {
		o.APIID = s.APIID
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
	if s.CloudTags != nil {
		o.CloudTags = s.CloudTags
	}
	if s.CloudType != nil {
		o.CloudType = s.CloudType
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
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
	if s.UpdatedTime != nil {
		o.UpdatedTime = s.UpdatedTime
	}
	if s.VpcID != nil {
		o.VpcID = s.VpcID
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
func (o *SparseCloudVPC) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudVPC) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudVPC()
	if o.APIID != nil {
		out.APIID = *o.APIID
	}
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
	if o.CloudTags != nil {
		out.CloudTags = *o.CloudTags
	}
	if o.CloudType != nil {
		out.CloudType = *o.CloudType
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
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
	if o.UpdatedTime != nil {
		out.UpdatedTime = *o.UpdatedTime
	}
	if o.VpcID != nil {
		out.VpcID = *o.VpcID
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
func (o *SparseCloudVPC) GetAPIID() (out int) {

	if o.APIID == nil {
		return
	}

	return *o.APIID
}

// SetAPIID sets the property APIID of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetAPIID(APIID int) {

	o.APIID = &APIID
}

// GetAccountID returns the AccountID of the receiver.
func (o *SparseCloudVPC) GetAccountID() (out string) {

	if o.AccountID == nil {
		return
	}

	return *o.AccountID
}

// SetAccountID sets the property AccountID of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetAccountID(accountID string) {

	o.AccountID = &accountID
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseCloudVPC) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseCloudVPC) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCloudTags returns the CloudTags of the receiver.
func (o *SparseCloudVPC) GetCloudTags() (out []string) {

	if o.CloudTags == nil {
		return
	}

	return *o.CloudTags
}

// SetCloudTags sets the property CloudTags of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetCloudTags(cloudTags []string) {

	o.CloudTags = &cloudTags
}

// GetCloudType returns the CloudType of the receiver.
func (o *SparseCloudVPC) GetCloudType() (out string) {

	if o.CloudType == nil {
		return
	}

	return *o.CloudType
}

// SetCloudType sets the property CloudType of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetCloudType(cloudType string) {

	o.CloudType = &cloudType
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseCloudVPC) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCustomerID returns the CustomerID of the receiver.
func (o *SparseCloudVPC) GetCustomerID() (out int) {

	if o.CustomerID == nil {
		return
	}

	return *o.CustomerID
}

// SetCustomerID sets the property CustomerID of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetCustomerID(customerID int) {

	o.CustomerID = &customerID
}

// GetIngestionTime returns the IngestionTime of the receiver.
func (o *SparseCloudVPC) GetIngestionTime() (out time.Time) {

	if o.IngestionTime == nil {
		return
	}

	return *o.IngestionTime
}

// SetIngestionTime sets the property IngestionTime of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetIngestionTime(ingestionTime time.Time) {

	o.IngestionTime = &ingestionTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseCloudVPC) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseCloudVPC) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseCloudVPC) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNativeID returns the NativeID of the receiver.
func (o *SparseCloudVPC) GetNativeID() (out string) {

	if o.NativeID == nil {
		return
	}

	return *o.NativeID
}

// SetNativeID sets the property NativeID of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetNativeID(nativeID string) {

	o.NativeID = &nativeID
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseCloudVPC) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetPolicyReferences returns the PolicyReferences of the receiver.
func (o *SparseCloudVPC) GetPolicyReferences() (out []string) {

	if o.PolicyReferences == nil {
		return
	}

	return *o.PolicyReferences
}

// SetPolicyReferences sets the property PolicyReferences of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetPolicyReferences(policyReferences []string) {

	o.PolicyReferences = &policyReferences
}

// GetProtected returns the Protected of the receiver.
func (o *SparseCloudVPC) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetRegionName returns the RegionName of the receiver.
func (o *SparseCloudVPC) GetRegionName() (out string) {

	if o.RegionName == nil {
		return
	}

	return *o.RegionName
}

// SetRegionName sets the property RegionName of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetRegionName(regionName string) {

	o.RegionName = &regionName
}

// GetResourceID returns the ResourceID of the receiver.
func (o *SparseCloudVPC) GetResourceID() (out int) {

	if o.ResourceID == nil {
		return
	}

	return *o.ResourceID
}

// SetResourceID sets the property ResourceID of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetResourceID(resourceID int) {

	o.ResourceID = &resourceID
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseCloudVPC) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdatedTime returns the UpdatedTime of the receiver.
func (o *SparseCloudVPC) GetUpdatedTime() (out time.Time) {

	if o.UpdatedTime == nil {
		return
	}

	return *o.UpdatedTime
}

// SetUpdatedTime sets the property UpdatedTime of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetUpdatedTime(updatedTime time.Time) {

	o.UpdatedTime = &updatedTime
}

// GetVpcID returns the VpcID of the receiver.
func (o *SparseCloudVPC) GetVpcID() (out string) {

	if o.VpcID == nil {
		return
	}

	return *o.VpcID
}

// SetVpcID sets the property VpcID of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetVpcID(vpcID string) {

	o.VpcID = &vpcID
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseCloudVPC) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseCloudVPC) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseCloudVPC.
func (o *SparseCloudVPC) DeepCopy() *SparseCloudVPC {

	if o == nil {
		return nil
	}

	out := &SparseCloudVPC{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudVPC.
func (o *SparseCloudVPC) DeepCopyInto(out *SparseCloudVPC) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudVPC: %s", err))
	}

	*out = *target.(*SparseCloudVPC)
}

type mongoAttributesCloudVPC struct {
	APIID                int                 `bson:"apiid"`
	ID                   bson.ObjectId       `bson:"_id,omitempty"`
	AccountID            string              `bson:"accountid"`
	Annotations          map[string][]string `bson:"annotations"`
	AssociatedTags       []string            `bson:"associatedtags"`
	CloudTags            []string            `bson:"cloudtags"`
	CloudType            string              `bson:"cloudtype"`
	CreateIdempotencyKey string              `bson:"createidempotencykey"`
	CustomerID           int                 `bson:"customerid"`
	IngestionTime        time.Time           `bson:"ingestiontime"`
	MigrationsLog        map[string]string   `bson:"migrationslog,omitempty"`
	Name                 string              `bson:"name"`
	Namespace            string              `bson:"namespace"`
	NativeID             string              `bson:"nativeid"`
	NormalizedTags       []string            `bson:"normalizedtags"`
	Parameters           *CloudVPCData       `bson:"parameters"`
	PolicyReferences     []string            `bson:"policyreferences"`
	Protected            bool                `bson:"protected"`
	RegionName           string              `bson:"regionname"`
	ResourceID           int                 `bson:"resourceid"`
	UpdateIdempotencyKey string              `bson:"updateidempotencykey"`
	UpdatedTime          time.Time           `bson:"updatedtime"`
	VpcID                string              `bson:"vpcid"`
	ZHash                int                 `bson:"zhash"`
	Zone                 int                 `bson:"zone"`
}
type mongoAttributesSparseCloudVPC struct {
	APIID                *int                 `bson:"apiid,omitempty"`
	ID                   bson.ObjectId        `bson:"_id,omitempty"`
	AccountID            *string              `bson:"accountid,omitempty"`
	Annotations          *map[string][]string `bson:"annotations,omitempty"`
	AssociatedTags       *[]string            `bson:"associatedtags,omitempty"`
	CloudTags            *[]string            `bson:"cloudtags,omitempty"`
	CloudType            *string              `bson:"cloudtype,omitempty"`
	CreateIdempotencyKey *string              `bson:"createidempotencykey,omitempty"`
	CustomerID           *int                 `bson:"customerid,omitempty"`
	IngestionTime        *time.Time           `bson:"ingestiontime,omitempty"`
	MigrationsLog        *map[string]string   `bson:"migrationslog,omitempty"`
	Name                 *string              `bson:"name,omitempty"`
	Namespace            *string              `bson:"namespace,omitempty"`
	NativeID             *string              `bson:"nativeid,omitempty"`
	NormalizedTags       *[]string            `bson:"normalizedtags,omitempty"`
	Parameters           *CloudVPCData        `bson:"parameters,omitempty"`
	PolicyReferences     *[]string            `bson:"policyreferences,omitempty"`
	Protected            *bool                `bson:"protected,omitempty"`
	RegionName           *string              `bson:"regionname,omitempty"`
	ResourceID           *int                 `bson:"resourceid,omitempty"`
	UpdateIdempotencyKey *string              `bson:"updateidempotencykey,omitempty"`
	UpdatedTime          *time.Time           `bson:"updatedtime,omitempty"`
	VpcID                *string              `bson:"vpcid,omitempty"`
	ZHash                *int                 `bson:"zhash,omitempty"`
	Zone                 *int                 `bson:"zone,omitempty"`
}
