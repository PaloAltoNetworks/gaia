package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudNodeTypeValue represents the possible values for attribute "type".
type CloudNodeTypeValue string

const (
	// CloudNodeTypeEndpoint represents the value Endpoint.
	CloudNodeTypeEndpoint CloudNodeTypeValue = "Endpoint"

	// CloudNodeTypeInterface represents the value Interface.
	CloudNodeTypeInterface CloudNodeTypeValue = "Interface"

	// CloudNodeTypeLoadBalancer represents the value LoadBalancer.
	CloudNodeTypeLoadBalancer CloudNodeTypeValue = "LoadBalancer"

	// CloudNodeTypeLoadBalancerRoute represents the value LoadBalancerRoute.
	CloudNodeTypeLoadBalancerRoute CloudNodeTypeValue = "LoadBalancerRoute"

	// CloudNodeTypeNetworkRuleSet represents the value NetworkRuleSet.
	CloudNodeTypeNetworkRuleSet CloudNodeTypeValue = "NetworkRuleSet"

	// CloudNodeTypeRouteTable represents the value RouteTable.
	CloudNodeTypeRouteTable CloudNodeTypeValue = "RouteTable"

	// CloudNodeTypeSubnet represents the value Subnet.
	CloudNodeTypeSubnet CloudNodeTypeValue = "Subnet"

	// CloudNodeTypeVPC represents the value VPC.
	CloudNodeTypeVPC CloudNodeTypeValue = "VPC"
)

// CloudNodeIdentity represents the Identity of the object.
var CloudNodeIdentity = elemental.Identity{
	Name:     "cloudnode",
	Category: "cloudnodes",
	Package:  "yeul",
	Private:  false,
}

// CloudNodesList represents a list of CloudNodes
type CloudNodesList []*CloudNode

// Identity returns the identity of the objects in the list.
func (o CloudNodesList) Identity() elemental.Identity {

	return CloudNodeIdentity
}

// Copy returns a pointer to a copy the CloudNodesList.
func (o CloudNodesList) Copy() elemental.Identifiables {

	copy := append(CloudNodesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CloudNodesList.
func (o CloudNodesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudNodesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudNode))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudNodesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudNodesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CloudNodesList converted to SparseCloudNodesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudNodesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudNodesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudNode)
	}

	return out
}

// Version returns the version of the content.
func (o CloudNodesList) Version() int {

	return 1
}

// CloudNode represents the model of a cloudnode
type CloudNode struct {
	// Prisma Cloud API ID (matches Prisma Cloud API ID).
	APIID int `json:"APIID,omitempty" msgpack:"APIID,omitempty" bson:"apiid" mapstructure:"APIID,omitempty"`

	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// ID of the host VPC.
	VPCID string `json:"VPCID,omitempty" msgpack:"VPCID,omitempty" bson:"vpcid" mapstructure:"VPCID,omitempty"`

	// Cloud account ID associated with the entity (matches Prisma Cloud accountID).
	AccountID string `json:"accountId,omitempty" msgpack:"accountId,omitempty" bson:"accountid" mapstructure:"accountId,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// The list of attachments for this node.
	Attachments []string `json:"attachments,omitempty" msgpack:"attachments,omitempty" bson:"attachments" mapstructure:"attachments,omitempty"`

	// Internal representation of object tags retrieved from the cloud provider.
	CloudTags []string `json:"cloudTags,omitempty" msgpack:"cloudTags,omitempty" bson:"cloudtags" mapstructure:"cloudTags,omitempty"`

	// Cloud type of the entity.
	CloudType string `json:"cloudType,omitempty" msgpack:"cloudType,omitempty" bson:"cloudtype" mapstructure:"cloudType,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Customer ID as identified by Prisma Cloud.
	CustomerID int `json:"customerID,omitempty" msgpack:"customerID,omitempty" bson:"customerid" mapstructure:"customerID,omitempty"`

	// The time that the object was first ingested.
	IngestionTime time.Time `json:"ingestionTime,omitempty" msgpack:"ingestionTime,omitempty" bson:"ingestiontime" mapstructure:"ingestionTime,omitempty"`

	// Internal unique key for a resource to guarantee no overlaps at write.
	Key string `json:"-" msgpack:"-" bson:"key" mapstructure:"-,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Name of the object (optional).
	Name string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// ID of the cloud provider object.
	NativeID string `json:"nativeID" msgpack:"nativeID" bson:"nativeid" mapstructure:"nativeID,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// The cloud attributes of the object.
	Parameters map[string]interface{} `json:"parameters,omitempty" msgpack:"parameters,omitempty" bson:"parameters" mapstructure:"parameters,omitempty"`

	// A list of policy references associated with this cloud node.
	PolicyReferences []string `json:"policyReferences,omitempty" msgpack:"policyReferences,omitempty" bson:"policyreferences" mapstructure:"policyReferences,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Region name associated with the entity.
	RegionName string `json:"regionName,omitempty" msgpack:"regionName,omitempty" bson:"regionname" mapstructure:"regionName,omitempty"`

	// A reference to a related object.
	RelatedObjectID string `json:"relatedObjectID,omitempty" msgpack:"relatedObjectID,omitempty" bson:"relatedobjectid" mapstructure:"relatedObjectID,omitempty"`

	// Prisma Cloud Resource ID.
	ResourceID int `json:"resourceID,omitempty" msgpack:"resourceID,omitempty" bson:"resourceid" mapstructure:"resourceID,omitempty"`

	// List of security tags associated with the node.
	SecurityTags []string `json:"securityTags,omitempty" msgpack:"securityTags,omitempty" bson:"securitytags" mapstructure:"securityTags,omitempty"`

	// The sub-type of the object as found in the parameters. Used for indexing.
	SubType string `json:"subType,omitempty" msgpack:"subType,omitempty" bson:"subtype" mapstructure:"subType,omitempty"`

	// Type of the endpoint.
	Type CloudNodeTypeValue `json:"type" msgpack:"type" bson:"type" mapstructure:"type,omitempty"`

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

// NewCloudNode returns a new *CloudNode
func NewCloudNode() *CloudNode {

	return &CloudNode{
		ModelVersion:     1,
		CloudTags:        []string{},
		Annotations:      map[string][]string{},
		AssociatedTags:   []string{},
		Attachments:      []string{},
		MigrationsLog:    map[string]string{},
		SecurityTags:     []string{},
		NormalizedTags:   []string{},
		PolicyReferences: []string{},
		Parameters:       map[string]interface{}{},
	}
}

// Identity returns the Identity of the object.
func (o *CloudNode) Identity() elemental.Identity {

	return CloudNodeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudNode) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudNode) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudNode) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudNode{}

	s.APIID = o.APIID
	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.VPCID = o.VPCID
	s.AccountID = o.AccountID
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.Attachments = o.Attachments
	s.CloudTags = o.CloudTags
	s.CloudType = o.CloudType
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.CustomerID = o.CustomerID
	s.IngestionTime = o.IngestionTime
	s.Key = o.Key
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NativeID = o.NativeID
	s.NormalizedTags = o.NormalizedTags
	s.Parameters = o.Parameters
	s.PolicyReferences = o.PolicyReferences
	s.Protected = o.Protected
	s.RegionName = o.RegionName
	s.RelatedObjectID = o.RelatedObjectID
	s.ResourceID = o.ResourceID
	s.SecurityTags = o.SecurityTags
	s.SubType = o.SubType
	s.Type = o.Type
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudNode) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudNode{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.APIID = s.APIID
	o.ID = s.ID.Hex()
	o.VPCID = s.VPCID
	o.AccountID = s.AccountID
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.Attachments = s.Attachments
	o.CloudTags = s.CloudTags
	o.CloudType = s.CloudType
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.CustomerID = s.CustomerID
	o.IngestionTime = s.IngestionTime
	o.Key = s.Key
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NativeID = s.NativeID
	o.NormalizedTags = s.NormalizedTags
	o.Parameters = s.Parameters
	o.PolicyReferences = s.PolicyReferences
	o.Protected = s.Protected
	o.RegionName = s.RegionName
	o.RelatedObjectID = s.RelatedObjectID
	o.ResourceID = s.ResourceID
	o.SecurityTags = s.SecurityTags
	o.SubType = s.SubType
	o.Type = s.Type
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudNode) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudNode) BleveType() string {

	return "cloudnode"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudNode) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CloudNode) Doc() string {

	return `Manages the list of cloud nodes available in a cloud deployment.`
}

func (o *CloudNode) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAPIID returns the APIID of the receiver.
func (o *CloudNode) GetAPIID() int {

	return o.APIID
}

// SetAPIID sets the property APIID of the receiver using the given value.
func (o *CloudNode) SetAPIID(APIID int) {

	o.APIID = APIID
}

// GetVPCID returns the VPCID of the receiver.
func (o *CloudNode) GetVPCID() string {

	return o.VPCID
}

// SetVPCID sets the property VPCID of the receiver using the given value.
func (o *CloudNode) SetVPCID(VPCID string) {

	o.VPCID = VPCID
}

// GetAccountID returns the AccountID of the receiver.
func (o *CloudNode) GetAccountID() string {

	return o.AccountID
}

// SetAccountID sets the property AccountID of the receiver using the given value.
func (o *CloudNode) SetAccountID(accountID string) {

	o.AccountID = accountID
}

// GetAnnotations returns the Annotations of the receiver.
func (o *CloudNode) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *CloudNode) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *CloudNode) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *CloudNode) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetAttachments returns the Attachments of the receiver.
func (o *CloudNode) GetAttachments() []string {

	return o.Attachments
}

// SetAttachments sets the property Attachments of the receiver using the given value.
func (o *CloudNode) SetAttachments(attachments []string) {

	o.Attachments = attachments
}

// GetCloudTags returns the CloudTags of the receiver.
func (o *CloudNode) GetCloudTags() []string {

	return o.CloudTags
}

// SetCloudTags sets the property CloudTags of the receiver using the given value.
func (o *CloudNode) SetCloudTags(cloudTags []string) {

	o.CloudTags = cloudTags
}

// GetCloudType returns the CloudType of the receiver.
func (o *CloudNode) GetCloudType() string {

	return o.CloudType
}

// SetCloudType sets the property CloudType of the receiver using the given value.
func (o *CloudNode) SetCloudType(cloudType string) {

	o.CloudType = cloudType
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *CloudNode) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *CloudNode) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *CloudNode) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *CloudNode) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetCustomerID returns the CustomerID of the receiver.
func (o *CloudNode) GetCustomerID() int {

	return o.CustomerID
}

// SetCustomerID sets the property CustomerID of the receiver using the given value.
func (o *CloudNode) SetCustomerID(customerID int) {

	o.CustomerID = customerID
}

// GetIngestionTime returns the IngestionTime of the receiver.
func (o *CloudNode) GetIngestionTime() time.Time {

	return o.IngestionTime
}

// SetIngestionTime sets the property IngestionTime of the receiver using the given value.
func (o *CloudNode) SetIngestionTime(ingestionTime time.Time) {

	o.IngestionTime = ingestionTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *CloudNode) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *CloudNode) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *CloudNode) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *CloudNode) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *CloudNode) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *CloudNode) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNativeID returns the NativeID of the receiver.
func (o *CloudNode) GetNativeID() string {

	return o.NativeID
}

// SetNativeID sets the property NativeID of the receiver using the given value.
func (o *CloudNode) SetNativeID(nativeID string) {

	o.NativeID = nativeID
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *CloudNode) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *CloudNode) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetParameters returns the Parameters of the receiver.
func (o *CloudNode) GetParameters() map[string]interface{} {

	return o.Parameters
}

// SetParameters sets the property Parameters of the receiver using the given value.
func (o *CloudNode) SetParameters(parameters map[string]interface{}) {

	o.Parameters = parameters
}

// GetPolicyReferences returns the PolicyReferences of the receiver.
func (o *CloudNode) GetPolicyReferences() []string {

	return o.PolicyReferences
}

// SetPolicyReferences sets the property PolicyReferences of the receiver using the given value.
func (o *CloudNode) SetPolicyReferences(policyReferences []string) {

	o.PolicyReferences = policyReferences
}

// GetProtected returns the Protected of the receiver.
func (o *CloudNode) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *CloudNode) SetProtected(protected bool) {

	o.Protected = protected
}

// GetRegionName returns the RegionName of the receiver.
func (o *CloudNode) GetRegionName() string {

	return o.RegionName
}

// SetRegionName sets the property RegionName of the receiver using the given value.
func (o *CloudNode) SetRegionName(regionName string) {

	o.RegionName = regionName
}

// GetResourceID returns the ResourceID of the receiver.
func (o *CloudNode) GetResourceID() int {

	return o.ResourceID
}

// SetResourceID sets the property ResourceID of the receiver using the given value.
func (o *CloudNode) SetResourceID(resourceID int) {

	o.ResourceID = resourceID
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *CloudNode) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *CloudNode) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *CloudNode) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *CloudNode) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *CloudNode) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *CloudNode) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *CloudNode) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *CloudNode) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudNode) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudNode{
			APIID:                &o.APIID,
			ID:                   &o.ID,
			VPCID:                &o.VPCID,
			AccountID:            &o.AccountID,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			Attachments:          &o.Attachments,
			CloudTags:            &o.CloudTags,
			CloudType:            &o.CloudType,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			CustomerID:           &o.CustomerID,
			IngestionTime:        &o.IngestionTime,
			Key:                  &o.Key,
			MigrationsLog:        &o.MigrationsLog,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NativeID:             &o.NativeID,
			NormalizedTags:       &o.NormalizedTags,
			Parameters:           &o.Parameters,
			PolicyReferences:     &o.PolicyReferences,
			Protected:            &o.Protected,
			RegionName:           &o.RegionName,
			RelatedObjectID:      &o.RelatedObjectID,
			ResourceID:           &o.ResourceID,
			SecurityTags:         &o.SecurityTags,
			SubType:              &o.SubType,
			Type:                 &o.Type,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseCloudNode{}
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
		case "attachments":
			sp.Attachments = &(o.Attachments)
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
		case "parameters":
			sp.Parameters = &(o.Parameters)
		case "policyReferences":
			sp.PolicyReferences = &(o.PolicyReferences)
		case "protected":
			sp.Protected = &(o.Protected)
		case "regionName":
			sp.RegionName = &(o.RegionName)
		case "relatedObjectID":
			sp.RelatedObjectID = &(o.RelatedObjectID)
		case "resourceID":
			sp.ResourceID = &(o.ResourceID)
		case "securityTags":
			sp.SecurityTags = &(o.SecurityTags)
		case "subType":
			sp.SubType = &(o.SubType)
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

// Patch apply the non nil value of a *SparseCloudNode to the object.
func (o *CloudNode) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudNode)
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
	if so.Attachments != nil {
		o.Attachments = *so.Attachments
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
	if so.Parameters != nil {
		o.Parameters = *so.Parameters
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
	if so.RelatedObjectID != nil {
		o.RelatedObjectID = *so.RelatedObjectID
	}
	if so.ResourceID != nil {
		o.ResourceID = *so.ResourceID
	}
	if so.SecurityTags != nil {
		o.SecurityTags = *so.SecurityTags
	}
	if so.SubType != nil {
		o.SubType = *so.SubType
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

// DeepCopy returns a deep copy if the CloudNode.
func (o *CloudNode) DeepCopy() *CloudNode {

	if o == nil {
		return nil
	}

	out := &CloudNode{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudNode.
func (o *CloudNode) DeepCopyInto(out *CloudNode) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudNode: %s", err))
	}

	*out = *target.(*CloudNode)
}

// Validate valides the current information stored into the structure.
func (o *CloudNode) Validate() error {

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

	if err := elemental.ValidateMaximumLength("regionName", o.RegionName, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("type", string(o.Type)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Endpoint", "Subnet", "VPC", "Interface", "RouteTable", "NetworkRuleSet", "LoadBalancer", "LoadBalancerRoute"}, false); err != nil {
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
func (*CloudNode) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudNodeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudNodeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudNode) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudNodeAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudNode) ValueForAttribute(name string) interface{} {

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
	case "attachments":
		return o.Attachments
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
	case "parameters":
		return o.Parameters
	case "policyReferences":
		return o.PolicyReferences
	case "protected":
		return o.Protected
	case "regionName":
		return o.RegionName
	case "relatedObjectID":
		return o.RelatedObjectID
	case "resourceID":
		return o.ResourceID
	case "securityTags":
		return o.SecurityTags
	case "subType":
		return o.SubType
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

// CloudNodeAttributesMap represents the map of attribute for CloudNode.
var CloudNodeAttributesMap = map[string]elemental.AttributeSpecification{
	"APIID": {
		AllowedChoices: []string{},
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
	"Attachments": {
		AllowedChoices: []string{},
		ConvertedName:  "Attachments",
		Description:    `The list of attachments for this node.`,
		Exposed:        true,
		Getter:         true,
		Name:           "attachments",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"CloudTags": {
		AllowedChoices: []string{},
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
		ConvertedName:  "Key",
		Description:    `Internal unique key for a resource to guarantee no overlaps at write.`,
		Name:           "key",
		Stored:         true,
		Type:           "string",
	},
	"MigrationsLog": {
		AllowedChoices: []string{},
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
		ConvertedName:  "Parameters",
		Description:    `The cloud attributes of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "parameters",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"PolicyReferences": {
		AllowedChoices: []string{},
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
	"RelatedObjectID": {
		AllowedChoices: []string{},
		ConvertedName:  "RelatedObjectID",
		Description:    `A reference to a related object.`,
		Exposed:        true,
		Name:           "relatedObjectID",
		Stored:         true,
		Type:           "string",
	},
	"ResourceID": {
		AllowedChoices: []string{},
		ConvertedName:  "ResourceID",
		Description:    `Prisma Cloud Resource ID.`,
		Exposed:        true,
		Getter:         true,
		Name:           "resourceID",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
	"SecurityTags": {
		AllowedChoices: []string{},
		ConvertedName:  "SecurityTags",
		Description:    `List of security tags associated with the node.`,
		Exposed:        true,
		Name:           "securityTags",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"SubType": {
		AllowedChoices: []string{},
		ConvertedName:  "SubType",
		Description:    `The sub-type of the object as found in the parameters. Used for indexing.`,
		Exposed:        true,
		Name:           "subType",
		Stored:         true,
		Type:           "string",
	},
	"Type": {
		AllowedChoices: []string{"Endpoint", "Subnet", "VPC", "Interface", "RouteTable", "NetworkRuleSet", "LoadBalancer", "LoadBalancerRoute"},
		ConvertedName:  "Type",
		Description:    `Type of the endpoint.`,
		Exposed:        true,
		Name:           "type",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"UpdateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
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

// CloudNodeLowerCaseAttributesMap represents the map of attribute for CloudNode.
var CloudNodeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"apiid": {
		AllowedChoices: []string{},
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
	"attachments": {
		AllowedChoices: []string{},
		ConvertedName:  "Attachments",
		Description:    `The list of attachments for this node.`,
		Exposed:        true,
		Getter:         true,
		Name:           "attachments",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"cloudtags": {
		AllowedChoices: []string{},
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
		ConvertedName:  "Key",
		Description:    `Internal unique key for a resource to guarantee no overlaps at write.`,
		Name:           "key",
		Stored:         true,
		Type:           "string",
	},
	"migrationslog": {
		AllowedChoices: []string{},
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
		ConvertedName:  "Parameters",
		Description:    `The cloud attributes of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "parameters",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"policyreferences": {
		AllowedChoices: []string{},
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
	"relatedobjectid": {
		AllowedChoices: []string{},
		ConvertedName:  "RelatedObjectID",
		Description:    `A reference to a related object.`,
		Exposed:        true,
		Name:           "relatedObjectID",
		Stored:         true,
		Type:           "string",
	},
	"resourceid": {
		AllowedChoices: []string{},
		ConvertedName:  "ResourceID",
		Description:    `Prisma Cloud Resource ID.`,
		Exposed:        true,
		Getter:         true,
		Name:           "resourceID",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
	"securitytags": {
		AllowedChoices: []string{},
		ConvertedName:  "SecurityTags",
		Description:    `List of security tags associated with the node.`,
		Exposed:        true,
		Name:           "securityTags",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"subtype": {
		AllowedChoices: []string{},
		ConvertedName:  "SubType",
		Description:    `The sub-type of the object as found in the parameters. Used for indexing.`,
		Exposed:        true,
		Name:           "subType",
		Stored:         true,
		Type:           "string",
	},
	"type": {
		AllowedChoices: []string{"Endpoint", "Subnet", "VPC", "Interface", "RouteTable", "NetworkRuleSet", "LoadBalancer", "LoadBalancerRoute"},
		ConvertedName:  "Type",
		Description:    `Type of the endpoint.`,
		Exposed:        true,
		Name:           "type",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"updateidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
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

// SparseCloudNodesList represents a list of SparseCloudNodes
type SparseCloudNodesList []*SparseCloudNode

// Identity returns the identity of the objects in the list.
func (o SparseCloudNodesList) Identity() elemental.Identity {

	return CloudNodeIdentity
}

// Copy returns a pointer to a copy the SparseCloudNodesList.
func (o SparseCloudNodesList) Copy() elemental.Identifiables {

	copy := append(SparseCloudNodesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudNodesList.
func (o SparseCloudNodesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudNodesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudNode))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudNodesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudNodesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCloudNodesList converted to CloudNodesList.
func (o SparseCloudNodesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudNodesList) Version() int {

	return 1
}

// SparseCloudNode represents the sparse version of a cloudnode.
type SparseCloudNode struct {
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

	// The list of attachments for this node.
	Attachments *[]string `json:"attachments,omitempty" msgpack:"attachments,omitempty" bson:"attachments,omitempty" mapstructure:"attachments,omitempty"`

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

	// The cloud attributes of the object.
	Parameters *map[string]interface{} `json:"parameters,omitempty" msgpack:"parameters,omitempty" bson:"parameters,omitempty" mapstructure:"parameters,omitempty"`

	// A list of policy references associated with this cloud node.
	PolicyReferences *[]string `json:"policyReferences,omitempty" msgpack:"policyReferences,omitempty" bson:"policyreferences,omitempty" mapstructure:"policyReferences,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Region name associated with the entity.
	RegionName *string `json:"regionName,omitempty" msgpack:"regionName,omitempty" bson:"regionname,omitempty" mapstructure:"regionName,omitempty"`

	// A reference to a related object.
	RelatedObjectID *string `json:"relatedObjectID,omitempty" msgpack:"relatedObjectID,omitempty" bson:"relatedobjectid,omitempty" mapstructure:"relatedObjectID,omitempty"`

	// Prisma Cloud Resource ID.
	ResourceID *int `json:"resourceID,omitempty" msgpack:"resourceID,omitempty" bson:"resourceid,omitempty" mapstructure:"resourceID,omitempty"`

	// List of security tags associated with the node.
	SecurityTags *[]string `json:"securityTags,omitempty" msgpack:"securityTags,omitempty" bson:"securitytags,omitempty" mapstructure:"securityTags,omitempty"`

	// The sub-type of the object as found in the parameters. Used for indexing.
	SubType *string `json:"subType,omitempty" msgpack:"subType,omitempty" bson:"subtype,omitempty" mapstructure:"subType,omitempty"`

	// Type of the endpoint.
	Type *CloudNodeTypeValue `json:"type,omitempty" msgpack:"type,omitempty" bson:"type,omitempty" mapstructure:"type,omitempty"`

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

// NewSparseCloudNode returns a new  SparseCloudNode.
func NewSparseCloudNode() *SparseCloudNode {
	return &SparseCloudNode{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudNode) Identity() elemental.Identity {

	return CloudNodeIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudNode) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudNode) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudNode) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudNode{}

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
	if o.Attachments != nil {
		s.Attachments = o.Attachments
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
	if o.RelatedObjectID != nil {
		s.RelatedObjectID = o.RelatedObjectID
	}
	if o.ResourceID != nil {
		s.ResourceID = o.ResourceID
	}
	if o.SecurityTags != nil {
		s.SecurityTags = o.SecurityTags
	}
	if o.SubType != nil {
		s.SubType = o.SubType
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
func (o *SparseCloudNode) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudNode{}
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
	if s.Attachments != nil {
		o.Attachments = s.Attachments
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
	if s.RelatedObjectID != nil {
		o.RelatedObjectID = s.RelatedObjectID
	}
	if s.ResourceID != nil {
		o.ResourceID = s.ResourceID
	}
	if s.SecurityTags != nil {
		o.SecurityTags = s.SecurityTags
	}
	if s.SubType != nil {
		o.SubType = s.SubType
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
func (o *SparseCloudNode) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudNode) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudNode()
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
	if o.Attachments != nil {
		out.Attachments = *o.Attachments
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
	if o.Parameters != nil {
		out.Parameters = *o.Parameters
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
	if o.RelatedObjectID != nil {
		out.RelatedObjectID = *o.RelatedObjectID
	}
	if o.ResourceID != nil {
		out.ResourceID = *o.ResourceID
	}
	if o.SecurityTags != nil {
		out.SecurityTags = *o.SecurityTags
	}
	if o.SubType != nil {
		out.SubType = *o.SubType
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
func (o *SparseCloudNode) GetAPIID() (out int) {

	if o.APIID == nil {
		return
	}

	return *o.APIID
}

// SetAPIID sets the property APIID of the receiver using the address of the given value.
func (o *SparseCloudNode) SetAPIID(APIID int) {

	o.APIID = &APIID
}

// GetVPCID returns the VPCID of the receiver.
func (o *SparseCloudNode) GetVPCID() (out string) {

	if o.VPCID == nil {
		return
	}

	return *o.VPCID
}

// SetVPCID sets the property VPCID of the receiver using the address of the given value.
func (o *SparseCloudNode) SetVPCID(VPCID string) {

	o.VPCID = &VPCID
}

// GetAccountID returns the AccountID of the receiver.
func (o *SparseCloudNode) GetAccountID() (out string) {

	if o.AccountID == nil {
		return
	}

	return *o.AccountID
}

// SetAccountID sets the property AccountID of the receiver using the address of the given value.
func (o *SparseCloudNode) SetAccountID(accountID string) {

	o.AccountID = &accountID
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseCloudNode) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseCloudNode) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseCloudNode) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseCloudNode) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetAttachments returns the Attachments of the receiver.
func (o *SparseCloudNode) GetAttachments() (out []string) {

	if o.Attachments == nil {
		return
	}

	return *o.Attachments
}

// SetAttachments sets the property Attachments of the receiver using the address of the given value.
func (o *SparseCloudNode) SetAttachments(attachments []string) {

	o.Attachments = &attachments
}

// GetCloudTags returns the CloudTags of the receiver.
func (o *SparseCloudNode) GetCloudTags() (out []string) {

	if o.CloudTags == nil {
		return
	}

	return *o.CloudTags
}

// SetCloudTags sets the property CloudTags of the receiver using the address of the given value.
func (o *SparseCloudNode) SetCloudTags(cloudTags []string) {

	o.CloudTags = &cloudTags
}

// GetCloudType returns the CloudType of the receiver.
func (o *SparseCloudNode) GetCloudType() (out string) {

	if o.CloudType == nil {
		return
	}

	return *o.CloudType
}

// SetCloudType sets the property CloudType of the receiver using the address of the given value.
func (o *SparseCloudNode) SetCloudType(cloudType string) {

	o.CloudType = &cloudType
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseCloudNode) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudNode) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseCloudNode) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseCloudNode) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetCustomerID returns the CustomerID of the receiver.
func (o *SparseCloudNode) GetCustomerID() (out int) {

	if o.CustomerID == nil {
		return
	}

	return *o.CustomerID
}

// SetCustomerID sets the property CustomerID of the receiver using the address of the given value.
func (o *SparseCloudNode) SetCustomerID(customerID int) {

	o.CustomerID = &customerID
}

// GetIngestionTime returns the IngestionTime of the receiver.
func (o *SparseCloudNode) GetIngestionTime() (out time.Time) {

	if o.IngestionTime == nil {
		return
	}

	return *o.IngestionTime
}

// SetIngestionTime sets the property IngestionTime of the receiver using the address of the given value.
func (o *SparseCloudNode) SetIngestionTime(ingestionTime time.Time) {

	o.IngestionTime = &ingestionTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseCloudNode) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseCloudNode) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseCloudNode) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseCloudNode) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseCloudNode) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseCloudNode) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNativeID returns the NativeID of the receiver.
func (o *SparseCloudNode) GetNativeID() (out string) {

	if o.NativeID == nil {
		return
	}

	return *o.NativeID
}

// SetNativeID sets the property NativeID of the receiver using the address of the given value.
func (o *SparseCloudNode) SetNativeID(nativeID string) {

	o.NativeID = &nativeID
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseCloudNode) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseCloudNode) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetParameters returns the Parameters of the receiver.
func (o *SparseCloudNode) GetParameters() (out map[string]interface{}) {

	if o.Parameters == nil {
		return
	}

	return *o.Parameters
}

// SetParameters sets the property Parameters of the receiver using the address of the given value.
func (o *SparseCloudNode) SetParameters(parameters map[string]interface{}) {

	o.Parameters = &parameters
}

// GetPolicyReferences returns the PolicyReferences of the receiver.
func (o *SparseCloudNode) GetPolicyReferences() (out []string) {

	if o.PolicyReferences == nil {
		return
	}

	return *o.PolicyReferences
}

// SetPolicyReferences sets the property PolicyReferences of the receiver using the address of the given value.
func (o *SparseCloudNode) SetPolicyReferences(policyReferences []string) {

	o.PolicyReferences = &policyReferences
}

// GetProtected returns the Protected of the receiver.
func (o *SparseCloudNode) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseCloudNode) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetRegionName returns the RegionName of the receiver.
func (o *SparseCloudNode) GetRegionName() (out string) {

	if o.RegionName == nil {
		return
	}

	return *o.RegionName
}

// SetRegionName sets the property RegionName of the receiver using the address of the given value.
func (o *SparseCloudNode) SetRegionName(regionName string) {

	o.RegionName = &regionName
}

// GetResourceID returns the ResourceID of the receiver.
func (o *SparseCloudNode) GetResourceID() (out int) {

	if o.ResourceID == nil {
		return
	}

	return *o.ResourceID
}

// SetResourceID sets the property ResourceID of the receiver using the address of the given value.
func (o *SparseCloudNode) SetResourceID(resourceID int) {

	o.ResourceID = &resourceID
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseCloudNode) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudNode) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseCloudNode) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseCloudNode) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseCloudNode) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseCloudNode) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseCloudNode) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseCloudNode) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseCloudNode.
func (o *SparseCloudNode) DeepCopy() *SparseCloudNode {

	if o == nil {
		return nil
	}

	out := &SparseCloudNode{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudNode.
func (o *SparseCloudNode) DeepCopyInto(out *SparseCloudNode) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudNode: %s", err))
	}

	*out = *target.(*SparseCloudNode)
}

type mongoAttributesCloudNode struct {
	APIID                int                    `bson:"apiid,omitempty"`
	ID                   bson.ObjectId          `bson:"_id,omitempty"`
	VPCID                string                 `bson:"vpcid,omitempty"`
	AccountID            string                 `bson:"accountid,omitempty"`
	Annotations          map[string][]string    `bson:"annotations"`
	AssociatedTags       []string               `bson:"associatedtags"`
	Attachments          []string               `bson:"attachments,omitempty"`
	CloudTags            []string               `bson:"cloudtags,omitempty"`
	CloudType            string                 `bson:"cloudtype,omitempty"`
	CreateIdempotencyKey string                 `bson:"createidempotencykey"`
	CreateTime           time.Time              `bson:"createtime"`
	CustomerID           int                    `bson:"customerid,omitempty"`
	IngestionTime        time.Time              `bson:"ingestiontime,omitempty"`
	Key                  string                 `bson:"key"`
	MigrationsLog        map[string]string      `bson:"migrationslog,omitempty"`
	Name                 string                 `bson:"name,omitempty"`
	Namespace            string                 `bson:"namespace"`
	NativeID             string                 `bson:"nativeid"`
	NormalizedTags       []string               `bson:"normalizedtags"`
	Parameters           map[string]interface{} `bson:"parameters,omitempty"`
	PolicyReferences     []string               `bson:"policyreferences,omitempty"`
	Protected            bool                   `bson:"protected"`
	RegionName           string                 `bson:"regionname,omitempty"`
	RelatedObjectID      string                 `bson:"relatedobjectid,omitempty"`
	ResourceID           int                    `bson:"resourceid,omitempty"`
	SecurityTags         []string               `bson:"securitytags,omitempty"`
	SubType              string                 `bson:"subtype,omitempty"`
	Type                 CloudNodeTypeValue     `bson:"type"`
	UpdateIdempotencyKey string                 `bson:"updateidempotencykey"`
	UpdateTime           time.Time              `bson:"updatetime"`
	ZHash                int                    `bson:"zhash"`
	Zone                 int                    `bson:"zone"`
}
type mongoAttributesSparseCloudNode struct {
	APIID                *int                    `bson:"apiid,omitempty"`
	ID                   bson.ObjectId           `bson:"_id,omitempty"`
	VPCID                *string                 `bson:"vpcid,omitempty"`
	AccountID            *string                 `bson:"accountid,omitempty"`
	Annotations          *map[string][]string    `bson:"annotations,omitempty"`
	AssociatedTags       *[]string               `bson:"associatedtags,omitempty"`
	Attachments          *[]string               `bson:"attachments,omitempty"`
	CloudTags            *[]string               `bson:"cloudtags,omitempty"`
	CloudType            *string                 `bson:"cloudtype,omitempty"`
	CreateIdempotencyKey *string                 `bson:"createidempotencykey,omitempty"`
	CreateTime           *time.Time              `bson:"createtime,omitempty"`
	CustomerID           *int                    `bson:"customerid,omitempty"`
	IngestionTime        *time.Time              `bson:"ingestiontime,omitempty"`
	Key                  *string                 `bson:"key,omitempty"`
	MigrationsLog        *map[string]string      `bson:"migrationslog,omitempty"`
	Name                 *string                 `bson:"name,omitempty"`
	Namespace            *string                 `bson:"namespace,omitempty"`
	NativeID             *string                 `bson:"nativeid,omitempty"`
	NormalizedTags       *[]string               `bson:"normalizedtags,omitempty"`
	Parameters           *map[string]interface{} `bson:"parameters,omitempty"`
	PolicyReferences     *[]string               `bson:"policyreferences,omitempty"`
	Protected            *bool                   `bson:"protected,omitempty"`
	RegionName           *string                 `bson:"regionname,omitempty"`
	RelatedObjectID      *string                 `bson:"relatedobjectid,omitempty"`
	ResourceID           *int                    `bson:"resourceid,omitempty"`
	SecurityTags         *[]string               `bson:"securitytags,omitempty"`
	SubType              *string                 `bson:"subtype,omitempty"`
	Type                 *CloudNodeTypeValue     `bson:"type,omitempty"`
	UpdateIdempotencyKey *string                 `bson:"updateidempotencykey,omitempty"`
	UpdateTime           *time.Time              `bson:"updatetime,omitempty"`
	ZHash                *int                    `bson:"zhash,omitempty"`
	Zone                 *int                    `bson:"zone,omitempty"`
}
