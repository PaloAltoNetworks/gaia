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

// K8sResourceKindValue represents the possible values for attribute "kind".
type K8sResourceKindValue string

const (
	// K8sResourceKindCluster represents the value Cluster.
	K8sResourceKindCluster K8sResourceKindValue = "Cluster"

	// K8sResourceKindDaemonSet represents the value DaemonSet.
	K8sResourceKindDaemonSet K8sResourceKindValue = "DaemonSet"

	// K8sResourceKindDeployment represents the value Deployment.
	K8sResourceKindDeployment K8sResourceKindValue = "Deployment"

	// K8sResourceKindEndpointSlice represents the value EndpointSlice.
	K8sResourceKindEndpointSlice K8sResourceKindValue = "EndpointSlice"

	// K8sResourceKindNamespace represents the value Namespace.
	K8sResourceKindNamespace K8sResourceKindValue = "Namespace"

	// K8sResourceKindNetworkPolicy represents the value NetworkPolicy.
	K8sResourceKindNetworkPolicy K8sResourceKindValue = "NetworkPolicy"

	// K8sResourceKindPending represents the value Pending.
	K8sResourceKindPending K8sResourceKindValue = "Pending"

	// K8sResourceKindPod represents the value Pod.
	K8sResourceKindPod K8sResourceKindValue = "Pod"

	// K8sResourceKindReplicaSet represents the value ReplicaSet.
	K8sResourceKindReplicaSet K8sResourceKindValue = "ReplicaSet"

	// K8sResourceKindService represents the value Service.
	K8sResourceKindService K8sResourceKindValue = "Service"

	// K8sResourceKindStatefulSet represents the value StatefulSet.
	K8sResourceKindStatefulSet K8sResourceKindValue = "StatefulSet"
)

// K8sResourceIdentity represents the Identity of the object.
var K8sResourceIdentity = elemental.Identity{
	Name:     "k8sresource",
	Category: "k8sresources",
	Package:  "pandemona",
	Private:  true,
}

// K8sResourcesList represents a list of K8sResources
type K8sResourcesList []*K8sResource

// Identity returns the identity of the objects in the list.
func (o K8sResourcesList) Identity() elemental.Identity {

	return K8sResourceIdentity
}

// Copy returns a pointer to a copy the K8sResourcesList.
func (o K8sResourcesList) Copy() elemental.Identifiables {

	out := append(K8sResourcesList{}, o...)
	return &out
}

// Append appends the objects to the a new copy of the K8sResourcesList.
func (o K8sResourcesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(K8sResourcesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*K8sResource))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o K8sResourcesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o K8sResourcesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the K8sResourcesList converted to SparseK8sResourcesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o K8sResourcesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseK8sResourcesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseK8sResource)
	}

	return out
}

// Version returns the version of the content.
func (o K8sResourcesList) Version() int {

	return 1
}

// K8sResource represents the model of a k8sresource
type K8sResource struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// The ID of the cloud cluster the resource belongs.
	ClusterID string `json:"clusterID" msgpack:"clusterID" bson:"clusterid" mapstructure:"clusterID,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// The JSON-encoded data that represents the resource.
	Data []byte `json:"data" msgpack:"data" bson:"data" mapstructure:"data,omitempty"`

	// Contextual values that can be used to narrow searching of resources.
	DenormedFields []string `json:"denormedFields" msgpack:"denormedFields" bson:"denormedfields" mapstructure:"denormedFields,omitempty"`

	// The specific kind of the k8s resource.
	Kind K8sResourceKindValue `json:"kind" msgpack:"kind" bson:"kind" mapstructure:"kind,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// The name of the resource.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// The Prisma Cloud Unified Asset ID.
	PrismaUnifiedAssetID string `json:"prismaUnifiedAssetID" msgpack:"prismaUnifiedAssetID" bson:"prismaunifiedassetid" mapstructure:"prismaUnifiedAssetID,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewK8sResource returns a new *K8sResource
func NewK8sResource() *K8sResource {

	return &K8sResource{
		ModelVersion:   1,
		Data:           []byte{},
		DenormedFields: []string{},
		Kind:           K8sResourceKindPending,
		MigrationsLog:  map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *K8sResource) Identity() elemental.Identity {

	return K8sResourceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *K8sResource) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *K8sResource) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *K8sResource) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesK8sResource{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.ClusterID = o.ClusterID
	s.CreateTime = o.CreateTime
	s.Data = o.Data
	s.DenormedFields = o.DenormedFields
	s.Kind = o.Kind
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.PrismaUnifiedAssetID = o.PrismaUnifiedAssetID
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *K8sResource) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesK8sResource{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.ClusterID = s.ClusterID
	o.CreateTime = s.CreateTime
	o.Data = s.Data
	o.DenormedFields = s.DenormedFields
	o.Kind = s.Kind
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.PrismaUnifiedAssetID = s.PrismaUnifiedAssetID
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *K8sResource) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *K8sResource) BleveType() string {

	return "k8sresource"
}

// DefaultOrder returns the list of default ordering fields.
func (o *K8sResource) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *K8sResource) Doc() string {

	return `Represents a K8s resource such as a service.`
}

func (o *K8sResource) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *K8sResource) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *K8sResource) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *K8sResource) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *K8sResource) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *K8sResource) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *K8sResource) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *K8sResource) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *K8sResource) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *K8sResource) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *K8sResource) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *K8sResource) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *K8sResource) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *K8sResource) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseK8sResource{
			ID:                   &o.ID,
			ClusterID:            &o.ClusterID,
			CreateTime:           &o.CreateTime,
			Data:                 &o.Data,
			DenormedFields:       &o.DenormedFields,
			Kind:                 &o.Kind,
			MigrationsLog:        &o.MigrationsLog,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			PrismaUnifiedAssetID: &o.PrismaUnifiedAssetID,
			UpdateTime:           &o.UpdateTime,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseK8sResource{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "clusterID":
			sp.ClusterID = &(o.ClusterID)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "data":
			sp.Data = &(o.Data)
		case "denormedFields":
			sp.DenormedFields = &(o.DenormedFields)
		case "kind":
			sp.Kind = &(o.Kind)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "prismaUnifiedAssetID":
			sp.PrismaUnifiedAssetID = &(o.PrismaUnifiedAssetID)
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

// Patch apply the non nil value of a *SparseK8sResource to the object.
func (o *K8sResource) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseK8sResource)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.ClusterID != nil {
		o.ClusterID = *so.ClusterID
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Data != nil {
		o.Data = *so.Data
	}
	if so.DenormedFields != nil {
		o.DenormedFields = *so.DenormedFields
	}
	if so.Kind != nil {
		o.Kind = *so.Kind
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
	if so.PrismaUnifiedAssetID != nil {
		o.PrismaUnifiedAssetID = *so.PrismaUnifiedAssetID
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

// DeepCopy returns a deep copy if the K8sResource.
func (o *K8sResource) DeepCopy() *K8sResource {

	if o == nil {
		return nil
	}

	out := &K8sResource{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *K8sResource.
func (o *K8sResource) DeepCopyInto(out *K8sResource) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy K8sResource: %s", err))
	}

	*out = *target.(*K8sResource)
}

// Validate valides the current information stored into the structure.
func (o *K8sResource) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("clusterID", o.ClusterID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredExternal("data", o.Data); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("kind", string(o.Kind), []string{"Pending", "Pod", "Deployment", "ReplicaSet", "StatefulSet", "DaemonSet", "Namespace", "NetworkPolicy", "Cluster", "Service", "EndpointSlice"}, true); err != nil {
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
func (*K8sResource) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := K8sResourceAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return K8sResourceLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*K8sResource) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return K8sResourceAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *K8sResource) ValueForAttribute(name string) any {

	switch name {
	case "ID":
		return o.ID
	case "clusterID":
		return o.ClusterID
	case "createTime":
		return o.CreateTime
	case "data":
		return o.Data
	case "denormedFields":
		return o.DenormedFields
	case "kind":
		return o.Kind
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "prismaUnifiedAssetID":
		return o.PrismaUnifiedAssetID
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// K8sResourceAttributesMap represents the map of attribute for K8sResource.
var K8sResourceAttributesMap = map[string]elemental.AttributeSpecification{
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
	"ClusterID": {
		AllowedChoices: []string{},
		BSONFieldName:  "clusterid",
		ConvertedName:  "ClusterID",
		Description:    `The ID of the cloud cluster the resource belongs.`,
		Exposed:        true,
		Name:           "clusterID",
		Required:       true,
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
	"Data": {
		AllowedChoices: []string{},
		BSONFieldName:  "data",
		ConvertedName:  "Data",
		Description:    `The JSON-encoded data that represents the resource.`,
		Exposed:        true,
		Name:           "data",
		Required:       true,
		Stored:         true,
		SubType:        "[]byte",
		Type:           "external",
	},
	"DenormedFields": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "denormedfields",
		ConvertedName:  "DenormedFields",
		Description:    `Contextual values that can be used to narrow searching of resources.`,
		Exposed:        true,
		Name:           "denormedFields",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Kind": {
		AllowedChoices: []string{"Pending", "Pod", "Deployment", "ReplicaSet", "StatefulSet", "DaemonSet", "Namespace", "NetworkPolicy", "Cluster", "Service", "EndpointSlice"},
		Autogenerated:  true,
		BSONFieldName:  "kind",
		ConvertedName:  "Kind",
		DefaultValue:   K8sResourceKindPending,
		Description:    `The specific kind of the k8s resource.`,
		Exposed:        true,
		Name:           "kind",
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
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
		Description:    `The name of the resource.`,
		Exposed:        true,
		Name:           "name",
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
	"PrismaUnifiedAssetID": {
		AllowedChoices: []string{},
		BSONFieldName:  "prismaunifiedassetid",
		ConvertedName:  "PrismaUnifiedAssetID",
		Description:    `The Prisma Cloud Unified Asset ID.`,
		Exposed:        true,
		Name:           "prismaUnifiedAssetID",
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

// K8sResourceLowerCaseAttributesMap represents the map of attribute for K8sResource.
var K8sResourceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"clusterid": {
		AllowedChoices: []string{},
		BSONFieldName:  "clusterid",
		ConvertedName:  "ClusterID",
		Description:    `The ID of the cloud cluster the resource belongs.`,
		Exposed:        true,
		Name:           "clusterID",
		Required:       true,
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
	"data": {
		AllowedChoices: []string{},
		BSONFieldName:  "data",
		ConvertedName:  "Data",
		Description:    `The JSON-encoded data that represents the resource.`,
		Exposed:        true,
		Name:           "data",
		Required:       true,
		Stored:         true,
		SubType:        "[]byte",
		Type:           "external",
	},
	"denormedfields": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "denormedfields",
		ConvertedName:  "DenormedFields",
		Description:    `Contextual values that can be used to narrow searching of resources.`,
		Exposed:        true,
		Name:           "denormedFields",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"kind": {
		AllowedChoices: []string{"Pending", "Pod", "Deployment", "ReplicaSet", "StatefulSet", "DaemonSet", "Namespace", "NetworkPolicy", "Cluster", "Service", "EndpointSlice"},
		Autogenerated:  true,
		BSONFieldName:  "kind",
		ConvertedName:  "Kind",
		DefaultValue:   K8sResourceKindPending,
		Description:    `The specific kind of the k8s resource.`,
		Exposed:        true,
		Name:           "kind",
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
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
		Description:    `The name of the resource.`,
		Exposed:        true,
		Name:           "name",
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
	"prismaunifiedassetid": {
		AllowedChoices: []string{},
		BSONFieldName:  "prismaunifiedassetid",
		ConvertedName:  "PrismaUnifiedAssetID",
		Description:    `The Prisma Cloud Unified Asset ID.`,
		Exposed:        true,
		Name:           "prismaUnifiedAssetID",
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

// SparseK8sResourcesList represents a list of SparseK8sResources
type SparseK8sResourcesList []*SparseK8sResource

// Identity returns the identity of the objects in the list.
func (o SparseK8sResourcesList) Identity() elemental.Identity {

	return K8sResourceIdentity
}

// Copy returns a pointer to a copy the SparseK8sResourcesList.
func (o SparseK8sResourcesList) Copy() elemental.Identifiables {

	copy := append(SparseK8sResourcesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseK8sResourcesList.
func (o SparseK8sResourcesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseK8sResourcesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseK8sResource))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseK8sResourcesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseK8sResourcesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseK8sResourcesList converted to K8sResourcesList.
func (o SparseK8sResourcesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseK8sResourcesList) Version() int {

	return 1
}

// SparseK8sResource represents the sparse version of a k8sresource.
type SparseK8sResource struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// The ID of the cloud cluster the resource belongs.
	ClusterID *string `json:"clusterID,omitempty" msgpack:"clusterID,omitempty" bson:"clusterid,omitempty" mapstructure:"clusterID,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// The JSON-encoded data that represents the resource.
	Data *[]byte `json:"data,omitempty" msgpack:"data,omitempty" bson:"data,omitempty" mapstructure:"data,omitempty"`

	// Contextual values that can be used to narrow searching of resources.
	DenormedFields *[]string `json:"denormedFields,omitempty" msgpack:"denormedFields,omitempty" bson:"denormedfields,omitempty" mapstructure:"denormedFields,omitempty"`

	// The specific kind of the k8s resource.
	Kind *K8sResourceKindValue `json:"kind,omitempty" msgpack:"kind,omitempty" bson:"kind,omitempty" mapstructure:"kind,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// The Prisma Cloud Unified Asset ID.
	PrismaUnifiedAssetID *string `json:"prismaUnifiedAssetID,omitempty" msgpack:"prismaUnifiedAssetID,omitempty" bson:"prismaunifiedassetid,omitempty" mapstructure:"prismaUnifiedAssetID,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseK8sResource returns a new  SparseK8sResource.
func NewSparseK8sResource() *SparseK8sResource {
	return &SparseK8sResource{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseK8sResource) Identity() elemental.Identity {

	return K8sResourceIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseK8sResource) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseK8sResource) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseK8sResource) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseK8sResource{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.ClusterID != nil {
		s.ClusterID = o.ClusterID
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.Data != nil {
		s.Data = o.Data
	}
	if o.DenormedFields != nil {
		s.DenormedFields = o.DenormedFields
	}
	if o.Kind != nil {
		s.Kind = o.Kind
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
	if o.PrismaUnifiedAssetID != nil {
		s.PrismaUnifiedAssetID = o.PrismaUnifiedAssetID
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
func (o *SparseK8sResource) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseK8sResource{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.ClusterID != nil {
		o.ClusterID = s.ClusterID
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.Data != nil {
		o.Data = s.Data
	}
	if s.DenormedFields != nil {
		o.DenormedFields = s.DenormedFields
	}
	if s.Kind != nil {
		o.Kind = s.Kind
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
	if s.PrismaUnifiedAssetID != nil {
		o.PrismaUnifiedAssetID = s.PrismaUnifiedAssetID
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
func (o *SparseK8sResource) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseK8sResource) ToPlain() elemental.PlainIdentifiable {

	out := NewK8sResource()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.ClusterID != nil {
		out.ClusterID = *o.ClusterID
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Data != nil {
		out.Data = *o.Data
	}
	if o.DenormedFields != nil {
		out.DenormedFields = *o.DenormedFields
	}
	if o.Kind != nil {
		out.Kind = *o.Kind
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
	if o.PrismaUnifiedAssetID != nil {
		out.PrismaUnifiedAssetID = *o.PrismaUnifiedAssetID
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
func (o *SparseK8sResource) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseK8sResource) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseK8sResource) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseK8sResource) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseK8sResource) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseK8sResource) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseK8sResource) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseK8sResource) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseK8sResource) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseK8sResource) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseK8sResource) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseK8sResource) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseK8sResource.
func (o *SparseK8sResource) DeepCopy() *SparseK8sResource {

	if o == nil {
		return nil
	}

	out := &SparseK8sResource{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseK8sResource.
func (o *SparseK8sResource) DeepCopyInto(out *SparseK8sResource) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseK8sResource: %s", err))
	}

	*out = *target.(*SparseK8sResource)
}

type mongoAttributesK8sResource struct {
	ID                   bson.ObjectId        `bson:"_id,omitempty"`
	ClusterID            string               `bson:"clusterid"`
	CreateTime           time.Time            `bson:"createtime"`
	Data                 []byte               `bson:"data"`
	DenormedFields       []string             `bson:"denormedfields"`
	Kind                 K8sResourceKindValue `bson:"kind"`
	MigrationsLog        map[string]string    `bson:"migrationslog,omitempty"`
	Name                 string               `bson:"name"`
	Namespace            string               `bson:"namespace"`
	PrismaUnifiedAssetID string               `bson:"prismaunifiedassetid"`
	UpdateTime           time.Time            `bson:"updatetime"`
	ZHash                int                  `bson:"zhash"`
	Zone                 int                  `bson:"zone"`
}
type mongoAttributesSparseK8sResource struct {
	ID                   bson.ObjectId         `bson:"_id,omitempty"`
	ClusterID            *string               `bson:"clusterid,omitempty"`
	CreateTime           *time.Time            `bson:"createtime,omitempty"`
	Data                 *[]byte               `bson:"data,omitempty"`
	DenormedFields       *[]string             `bson:"denormedfields,omitempty"`
	Kind                 *K8sResourceKindValue `bson:"kind,omitempty"`
	MigrationsLog        *map[string]string    `bson:"migrationslog,omitempty"`
	Name                 *string               `bson:"name,omitempty"`
	Namespace            *string               `bson:"namespace,omitempty"`
	PrismaUnifiedAssetID *string               `bson:"prismaunifiedassetid,omitempty"`
	UpdateTime           *time.Time            `bson:"updatetime,omitempty"`
	ZHash                *int                  `bson:"zhash,omitempty"`
	Zone                 *int                  `bson:"zone,omitempty"`
}
