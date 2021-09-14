package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// KubernetesClusterIdentity represents the Identity of the object.
var KubernetesClusterIdentity = elemental.Identity{
	Name:     "kubernetescluster",
	Category: "kubernetesclusters",
	Package:  "squall",
	Private:  false,
}

// KubernetesClustersList represents a list of KubernetesClusters
type KubernetesClustersList []*KubernetesCluster

// Identity returns the identity of the objects in the list.
func (o KubernetesClustersList) Identity() elemental.Identity {

	return KubernetesClusterIdentity
}

// Copy returns a pointer to a copy the KubernetesClustersList.
func (o KubernetesClustersList) Copy() elemental.Identifiables {

	copy := append(KubernetesClustersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the KubernetesClustersList.
func (o KubernetesClustersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(KubernetesClustersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*KubernetesCluster))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o KubernetesClustersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o KubernetesClustersList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the KubernetesClustersList converted to SparseKubernetesClustersList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o KubernetesClustersList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseKubernetesClustersList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseKubernetesCluster)
	}

	return out
}

// Version returns the version of the content.
func (o KubernetesClustersList) Version() int {

	return 1
}

// KubernetesCluster represents the model of a kubernetescluster
type KubernetesCluster struct {
	// It contains the fqdns that will be set in the certificate SANS field.
	APIServerServiceFQDNs []string `json:"APIServerServiceFQDNs" msgpack:"APIServerServiceFQDNs" bson:"apiserverservicefqdns" mapstructure:"APIServerServiceFQDNs,omitempty"`

	// It contains the ips that will be set in the certificate SANS field.
	APIServerServiceIPs []string `json:"APIServerServiceIPs" msgpack:"APIServerServiceIPs" bson:"apiserverserviceips" mapstructure:"APIServerServiceIPs,omitempty"`

	// Kubernetes service name in the format <service name>.<service name
	// namespace>.svc will be set in the certificate CommonName field.
	APIServerServiceName string `json:"APIServerServiceName" msgpack:"APIServerServiceName" bson:"apiserverservicename" mapstructure:"APIServerServiceName,omitempty"`

	// API versions supported by the API server.
	APIServerVersions []string `json:"APIServerVersions" msgpack:"APIServerVersions" bson:"apiserverversions" mapstructure:"APIServerVersions,omitempty"`

	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Cluster external IP address.
	ExternalIP string `json:"externalIP" msgpack:"externalIP" bson:"externalip" mapstructure:"externalIP,omitempty"`

	// Cluster internal IP address.
	InternalIP string `json:"internalIP" msgpack:"internalIP" bson:"internalip" mapstructure:"internalIP,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

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

// NewKubernetesCluster returns a new *KubernetesCluster
func NewKubernetesCluster() *KubernetesCluster {

	return &KubernetesCluster{
		ModelVersion:          1,
		APIServerServiceFQDNs: []string{},
		Metadata:              []string{},
		Annotations:           map[string][]string{},
		AssociatedTags:        []string{},
		MigrationsLog:         map[string]string{},
		APIServerServiceIPs:   []string{},
		APIServerVersions:     []string{},
		NormalizedTags:        []string{},
	}
}

// Identity returns the Identity of the object.
func (o *KubernetesCluster) Identity() elemental.Identity {

	return KubernetesClusterIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *KubernetesCluster) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *KubernetesCluster) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *KubernetesCluster) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesKubernetesCluster{}

	s.APIServerServiceFQDNs = o.APIServerServiceFQDNs
	s.APIServerServiceIPs = o.APIServerServiceIPs
	s.APIServerServiceName = o.APIServerServiceName
	s.APIServerVersions = o.APIServerVersions
	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.Description = o.Description
	s.ExternalIP = o.ExternalIP
	s.InternalIP = o.InternalIP
	s.Metadata = o.Metadata
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.Protected = o.Protected
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *KubernetesCluster) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesKubernetesCluster{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.APIServerServiceFQDNs = s.APIServerServiceFQDNs
	o.APIServerServiceIPs = s.APIServerServiceIPs
	o.APIServerServiceName = s.APIServerServiceName
	o.APIServerVersions = s.APIServerVersions
	o.ID = s.ID.Hex()
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.Description = s.Description
	o.ExternalIP = s.ExternalIP
	o.InternalIP = s.InternalIP
	o.Metadata = s.Metadata
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.Protected = s.Protected
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *KubernetesCluster) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *KubernetesCluster) BleveType() string {

	return "kubernetescluster"
}

// DefaultOrder returns the list of default ordering fields.
func (o *KubernetesCluster) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *KubernetesCluster) Doc() string {

	return `Used to represent an instance of a Kubernetes API server.`
}

func (o *KubernetesCluster) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *KubernetesCluster) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *KubernetesCluster) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *KubernetesCluster) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *KubernetesCluster) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *KubernetesCluster) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *KubernetesCluster) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *KubernetesCluster) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *KubernetesCluster) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *KubernetesCluster) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *KubernetesCluster) SetDescription(description string) {

	o.Description = description
}

// GetMetadata returns the Metadata of the receiver.
func (o *KubernetesCluster) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *KubernetesCluster) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *KubernetesCluster) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *KubernetesCluster) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *KubernetesCluster) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *KubernetesCluster) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *KubernetesCluster) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *KubernetesCluster) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *KubernetesCluster) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *KubernetesCluster) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *KubernetesCluster) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *KubernetesCluster) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *KubernetesCluster) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *KubernetesCluster) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *KubernetesCluster) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *KubernetesCluster) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *KubernetesCluster) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *KubernetesCluster) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *KubernetesCluster) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *KubernetesCluster) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *KubernetesCluster) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseKubernetesCluster{
			APIServerServiceFQDNs: &o.APIServerServiceFQDNs,
			APIServerServiceIPs:   &o.APIServerServiceIPs,
			APIServerServiceName:  &o.APIServerServiceName,
			APIServerVersions:     &o.APIServerVersions,
			ID:                    &o.ID,
			Annotations:           &o.Annotations,
			AssociatedTags:        &o.AssociatedTags,
			CreateIdempotencyKey:  &o.CreateIdempotencyKey,
			CreateTime:            &o.CreateTime,
			Description:           &o.Description,
			ExternalIP:            &o.ExternalIP,
			InternalIP:            &o.InternalIP,
			Metadata:              &o.Metadata,
			MigrationsLog:         &o.MigrationsLog,
			Name:                  &o.Name,
			Namespace:             &o.Namespace,
			NormalizedTags:        &o.NormalizedTags,
			Protected:             &o.Protected,
			UpdateIdempotencyKey:  &o.UpdateIdempotencyKey,
			UpdateTime:            &o.UpdateTime,
			ZHash:                 &o.ZHash,
			Zone:                  &o.Zone,
		}
	}

	sp := &SparseKubernetesCluster{}
	for _, f := range fields {
		switch f {
		case "APIServerServiceFQDNs":
			sp.APIServerServiceFQDNs = &(o.APIServerServiceFQDNs)
		case "APIServerServiceIPs":
			sp.APIServerServiceIPs = &(o.APIServerServiceIPs)
		case "APIServerServiceName":
			sp.APIServerServiceName = &(o.APIServerServiceName)
		case "APIServerVersions":
			sp.APIServerVersions = &(o.APIServerVersions)
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "externalIP":
			sp.ExternalIP = &(o.ExternalIP)
		case "internalIP":
			sp.InternalIP = &(o.InternalIP)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "protected":
			sp.Protected = &(o.Protected)
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

// Patch apply the non nil value of a *SparseKubernetesCluster to the object.
func (o *KubernetesCluster) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseKubernetesCluster)
	if so.APIServerServiceFQDNs != nil {
		o.APIServerServiceFQDNs = *so.APIServerServiceFQDNs
	}
	if so.APIServerServiceIPs != nil {
		o.APIServerServiceIPs = *so.APIServerServiceIPs
	}
	if so.APIServerServiceName != nil {
		o.APIServerServiceName = *so.APIServerServiceName
	}
	if so.APIServerVersions != nil {
		o.APIServerVersions = *so.APIServerVersions
	}
	if so.ID != nil {
		o.ID = *so.ID
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
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.ExternalIP != nil {
		o.ExternalIP = *so.ExternalIP
	}
	if so.InternalIP != nil {
		o.InternalIP = *so.InternalIP
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
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
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
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

// DeepCopy returns a deep copy if the KubernetesCluster.
func (o *KubernetesCluster) DeepCopy() *KubernetesCluster {

	if o == nil {
		return nil
	}

	out := &KubernetesCluster{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *KubernetesCluster.
func (o *KubernetesCluster) DeepCopyInto(out *KubernetesCluster) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy KubernetesCluster: %s", err))
	}

	*out = *target.(*KubernetesCluster)
}

// Validate valides the current information stored into the structure.
func (o *KubernetesCluster) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateMetadata("metadata", o.Metadata); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
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
func (*KubernetesCluster) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := KubernetesClusterAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return KubernetesClusterLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*KubernetesCluster) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return KubernetesClusterAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *KubernetesCluster) ValueForAttribute(name string) interface{} {

	switch name {
	case "APIServerServiceFQDNs":
		return o.APIServerServiceFQDNs
	case "APIServerServiceIPs":
		return o.APIServerServiceIPs
	case "APIServerServiceName":
		return o.APIServerServiceName
	case "APIServerVersions":
		return o.APIServerVersions
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "externalIP":
		return o.ExternalIP
	case "internalIP":
		return o.InternalIP
	case "metadata":
		return o.Metadata
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "protected":
		return o.Protected
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

// KubernetesClusterAttributesMap represents the map of attribute for KubernetesCluster.
var KubernetesClusterAttributesMap = map[string]elemental.AttributeSpecification{
	"APIServerServiceFQDNs": {
		AllowedChoices: []string{},
		BSONFieldName:  "apiserverservicefqdns",
		ConvertedName:  "APIServerServiceFQDNs",
		Description:    `It contains the fqdns that will be set in the certificate SANS field.`,
		Exposed:        true,
		Name:           "APIServerServiceFQDNs",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"APIServerServiceIPs": {
		AllowedChoices: []string{},
		BSONFieldName:  "apiserverserviceips",
		ConvertedName:  "APIServerServiceIPs",
		Description:    `It contains the ips that will be set in the certificate SANS field.`,
		Exposed:        true,
		Name:           "APIServerServiceIPs",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"APIServerServiceName": {
		AllowedChoices: []string{},
		BSONFieldName:  "apiserverservicename",
		ConvertedName:  "APIServerServiceName",
		Description: `Kubernetes service name in the format <service name>.<service name
namespace>.svc will be set in the certificate CommonName field.`,
		Exposed: true,
		Name:    "APIServerServiceName",
		Stored:  true,
		Type:    "string",
	},
	"APIServerVersions": {
		AllowedChoices: []string{},
		BSONFieldName:  "apiserverversions",
		ConvertedName:  "APIServerVersions",
		Description:    `API versions supported by the API server.`,
		Exposed:        true,
		Name:           "APIServerVersions",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"Description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ExternalIP": {
		AllowedChoices: []string{},
		BSONFieldName:  "externalip",
		ConvertedName:  "ExternalIP",
		Description:    `Cluster external IP address.`,
		Exposed:        true,
		Name:           "externalIP",
		Stored:         true,
		Type:           "string",
	},
	"InternalIP": {
		AllowedChoices: []string{},
		BSONFieldName:  "internalip",
		ConvertedName:  "InternalIP",
		Description:    `Cluster internal IP address.`,
		Exposed:        true,
		Name:           "internalIP",
		Stored:         true,
		Type:           "string",
	},
	"Metadata": {
		AllowedChoices: []string{},
		BSONFieldName:  "metadata",
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
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

// KubernetesClusterLowerCaseAttributesMap represents the map of attribute for KubernetesCluster.
var KubernetesClusterLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"apiserverservicefqdns": {
		AllowedChoices: []string{},
		BSONFieldName:  "apiserverservicefqdns",
		ConvertedName:  "APIServerServiceFQDNs",
		Description:    `It contains the fqdns that will be set in the certificate SANS field.`,
		Exposed:        true,
		Name:           "APIServerServiceFQDNs",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"apiserverserviceips": {
		AllowedChoices: []string{},
		BSONFieldName:  "apiserverserviceips",
		ConvertedName:  "APIServerServiceIPs",
		Description:    `It contains the ips that will be set in the certificate SANS field.`,
		Exposed:        true,
		Name:           "APIServerServiceIPs",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"apiserverservicename": {
		AllowedChoices: []string{},
		BSONFieldName:  "apiserverservicename",
		ConvertedName:  "APIServerServiceName",
		Description: `Kubernetes service name in the format <service name>.<service name
namespace>.svc will be set in the certificate CommonName field.`,
		Exposed: true,
		Name:    "APIServerServiceName",
		Stored:  true,
		Type:    "string",
	},
	"apiserverversions": {
		AllowedChoices: []string{},
		BSONFieldName:  "apiserverversions",
		ConvertedName:  "APIServerVersions",
		Description:    `API versions supported by the API server.`,
		Exposed:        true,
		Name:           "APIServerVersions",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"externalip": {
		AllowedChoices: []string{},
		BSONFieldName:  "externalip",
		ConvertedName:  "ExternalIP",
		Description:    `Cluster external IP address.`,
		Exposed:        true,
		Name:           "externalIP",
		Stored:         true,
		Type:           "string",
	},
	"internalip": {
		AllowedChoices: []string{},
		BSONFieldName:  "internalip",
		ConvertedName:  "InternalIP",
		Description:    `Cluster internal IP address.`,
		Exposed:        true,
		Name:           "internalIP",
		Stored:         true,
		Type:           "string",
	},
	"metadata": {
		AllowedChoices: []string{},
		BSONFieldName:  "metadata",
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
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

// SparseKubernetesClustersList represents a list of SparseKubernetesClusters
type SparseKubernetesClustersList []*SparseKubernetesCluster

// Identity returns the identity of the objects in the list.
func (o SparseKubernetesClustersList) Identity() elemental.Identity {

	return KubernetesClusterIdentity
}

// Copy returns a pointer to a copy the SparseKubernetesClustersList.
func (o SparseKubernetesClustersList) Copy() elemental.Identifiables {

	copy := append(SparseKubernetesClustersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseKubernetesClustersList.
func (o SparseKubernetesClustersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseKubernetesClustersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseKubernetesCluster))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseKubernetesClustersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseKubernetesClustersList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseKubernetesClustersList converted to KubernetesClustersList.
func (o SparseKubernetesClustersList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseKubernetesClustersList) Version() int {

	return 1
}

// SparseKubernetesCluster represents the sparse version of a kubernetescluster.
type SparseKubernetesCluster struct {
	// It contains the fqdns that will be set in the certificate SANS field.
	APIServerServiceFQDNs *[]string `json:"APIServerServiceFQDNs,omitempty" msgpack:"APIServerServiceFQDNs,omitempty" bson:"apiserverservicefqdns,omitempty" mapstructure:"APIServerServiceFQDNs,omitempty"`

	// It contains the ips that will be set in the certificate SANS field.
	APIServerServiceIPs *[]string `json:"APIServerServiceIPs,omitempty" msgpack:"APIServerServiceIPs,omitempty" bson:"apiserverserviceips,omitempty" mapstructure:"APIServerServiceIPs,omitempty"`

	// Kubernetes service name in the format <service name>.<service name
	// namespace>.svc will be set in the certificate CommonName field.
	APIServerServiceName *string `json:"APIServerServiceName,omitempty" msgpack:"APIServerServiceName,omitempty" bson:"apiserverservicename,omitempty" mapstructure:"APIServerServiceName,omitempty"`

	// API versions supported by the API server.
	APIServerVersions *[]string `json:"APIServerVersions,omitempty" msgpack:"APIServerVersions,omitempty" bson:"apiserverversions,omitempty" mapstructure:"APIServerVersions,omitempty"`

	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Cluster external IP address.
	ExternalIP *string `json:"externalIP,omitempty" msgpack:"externalIP,omitempty" bson:"externalip,omitempty" mapstructure:"externalIP,omitempty"`

	// Cluster internal IP address.
	InternalIP *string `json:"internalIP,omitempty" msgpack:"internalIP,omitempty" bson:"internalip,omitempty" mapstructure:"internalIP,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

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

// NewSparseKubernetesCluster returns a new  SparseKubernetesCluster.
func NewSparseKubernetesCluster() *SparseKubernetesCluster {
	return &SparseKubernetesCluster{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseKubernetesCluster) Identity() elemental.Identity {

	return KubernetesClusterIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseKubernetesCluster) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseKubernetesCluster) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseKubernetesCluster) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseKubernetesCluster{}

	if o.APIServerServiceFQDNs != nil {
		s.APIServerServiceFQDNs = o.APIServerServiceFQDNs
	}
	if o.APIServerServiceIPs != nil {
		s.APIServerServiceIPs = o.APIServerServiceIPs
	}
	if o.APIServerServiceName != nil {
		s.APIServerServiceName = o.APIServerServiceName
	}
	if o.APIServerVersions != nil {
		s.APIServerVersions = o.APIServerVersions
	}
	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
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
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.ExternalIP != nil {
		s.ExternalIP = o.ExternalIP
	}
	if o.InternalIP != nil {
		s.InternalIP = o.InternalIP
	}
	if o.Metadata != nil {
		s.Metadata = o.Metadata
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
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.Protected != nil {
		s.Protected = o.Protected
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
func (o *SparseKubernetesCluster) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseKubernetesCluster{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.APIServerServiceFQDNs != nil {
		o.APIServerServiceFQDNs = s.APIServerServiceFQDNs
	}
	if s.APIServerServiceIPs != nil {
		o.APIServerServiceIPs = s.APIServerServiceIPs
	}
	if s.APIServerServiceName != nil {
		o.APIServerServiceName = s.APIServerServiceName
	}
	if s.APIServerVersions != nil {
		o.APIServerVersions = s.APIServerVersions
	}
	id := s.ID.Hex()
	o.ID = &id
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
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.ExternalIP != nil {
		o.ExternalIP = s.ExternalIP
	}
	if s.InternalIP != nil {
		o.InternalIP = s.InternalIP
	}
	if s.Metadata != nil {
		o.Metadata = s.Metadata
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
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.Protected != nil {
		o.Protected = s.Protected
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
func (o *SparseKubernetesCluster) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseKubernetesCluster) ToPlain() elemental.PlainIdentifiable {

	out := NewKubernetesCluster()
	if o.APIServerServiceFQDNs != nil {
		out.APIServerServiceFQDNs = *o.APIServerServiceFQDNs
	}
	if o.APIServerServiceIPs != nil {
		out.APIServerServiceIPs = *o.APIServerServiceIPs
	}
	if o.APIServerServiceName != nil {
		out.APIServerServiceName = *o.APIServerServiceName
	}
	if o.APIServerVersions != nil {
		out.APIServerVersions = *o.APIServerVersions
	}
	if o.ID != nil {
		out.ID = *o.ID
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
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.ExternalIP != nil {
		out.ExternalIP = *o.ExternalIP
	}
	if o.InternalIP != nil {
		out.InternalIP = *o.InternalIP
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
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
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
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
func (o *SparseKubernetesCluster) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseKubernetesCluster) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseKubernetesCluster) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseKubernetesCluster) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseKubernetesCluster) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseKubernetesCluster) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseKubernetesCluster) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseKubernetesCluster) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseKubernetesCluster) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseKubernetesCluster) SetDescription(description string) {

	o.Description = &description
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseKubernetesCluster) GetMetadata() (out []string) {

	if o.Metadata == nil {
		return
	}

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseKubernetesCluster) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseKubernetesCluster) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseKubernetesCluster) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseKubernetesCluster) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseKubernetesCluster) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseKubernetesCluster) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseKubernetesCluster) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseKubernetesCluster) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseKubernetesCluster) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseKubernetesCluster) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseKubernetesCluster) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseKubernetesCluster) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseKubernetesCluster) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseKubernetesCluster) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseKubernetesCluster) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseKubernetesCluster) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseKubernetesCluster) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseKubernetesCluster) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseKubernetesCluster) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseKubernetesCluster.
func (o *SparseKubernetesCluster) DeepCopy() *SparseKubernetesCluster {

	if o == nil {
		return nil
	}

	out := &SparseKubernetesCluster{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseKubernetesCluster.
func (o *SparseKubernetesCluster) DeepCopyInto(out *SparseKubernetesCluster) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseKubernetesCluster: %s", err))
	}

	*out = *target.(*SparseKubernetesCluster)
}

type mongoAttributesKubernetesCluster struct {
	APIServerServiceFQDNs []string            `bson:"apiserverservicefqdns"`
	APIServerServiceIPs   []string            `bson:"apiserverserviceips"`
	APIServerServiceName  string              `bson:"apiserverservicename"`
	APIServerVersions     []string            `bson:"apiserverversions"`
	ID                    bson.ObjectId       `bson:"_id,omitempty"`
	Annotations           map[string][]string `bson:"annotations"`
	AssociatedTags        []string            `bson:"associatedtags"`
	CreateIdempotencyKey  string              `bson:"createidempotencykey"`
	CreateTime            time.Time           `bson:"createtime"`
	Description           string              `bson:"description"`
	ExternalIP            string              `bson:"externalip"`
	InternalIP            string              `bson:"internalip"`
	Metadata              []string            `bson:"metadata"`
	MigrationsLog         map[string]string   `bson:"migrationslog,omitempty"`
	Name                  string              `bson:"name"`
	Namespace             string              `bson:"namespace"`
	NormalizedTags        []string            `bson:"normalizedtags"`
	Protected             bool                `bson:"protected"`
	UpdateIdempotencyKey  string              `bson:"updateidempotencykey"`
	UpdateTime            time.Time           `bson:"updatetime"`
	ZHash                 int                 `bson:"zhash"`
	Zone                  int                 `bson:"zone"`
}
type mongoAttributesSparseKubernetesCluster struct {
	APIServerServiceFQDNs *[]string            `bson:"apiserverservicefqdns,omitempty"`
	APIServerServiceIPs   *[]string            `bson:"apiserverserviceips,omitempty"`
	APIServerServiceName  *string              `bson:"apiserverservicename,omitempty"`
	APIServerVersions     *[]string            `bson:"apiserverversions,omitempty"`
	ID                    bson.ObjectId        `bson:"_id,omitempty"`
	Annotations           *map[string][]string `bson:"annotations,omitempty"`
	AssociatedTags        *[]string            `bson:"associatedtags,omitempty"`
	CreateIdempotencyKey  *string              `bson:"createidempotencykey,omitempty"`
	CreateTime            *time.Time           `bson:"createtime,omitempty"`
	Description           *string              `bson:"description,omitempty"`
	ExternalIP            *string              `bson:"externalip,omitempty"`
	InternalIP            *string              `bson:"internalip,omitempty"`
	Metadata              *[]string            `bson:"metadata,omitempty"`
	MigrationsLog         *map[string]string   `bson:"migrationslog,omitempty"`
	Name                  *string              `bson:"name,omitempty"`
	Namespace             *string              `bson:"namespace,omitempty"`
	NormalizedTags        *[]string            `bson:"normalizedtags,omitempty"`
	Protected             *bool                `bson:"protected,omitempty"`
	UpdateIdempotencyKey  *string              `bson:"updateidempotencykey,omitempty"`
	UpdateTime            *time.Time           `bson:"updatetime,omitempty"`
	ZHash                 *int                 `bson:"zhash,omitempty"`
	Zone                  *int                 `bson:"zone,omitempty"`
}
