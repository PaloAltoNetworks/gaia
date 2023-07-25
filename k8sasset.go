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

// K8sAssetKindValue represents the possible values for attribute "kind".
type K8sAssetKindValue string

const (
	// K8sAssetKindCluster represents the value Cluster.
	K8sAssetKindCluster K8sAssetKindValue = "Cluster"

	// K8sAssetKindDaemonSets represents the value DaemonSets.
	K8sAssetKindDaemonSets K8sAssetKindValue = "DaemonSets"

	// K8sAssetKindDeployment represents the value Deployment.
	K8sAssetKindDeployment K8sAssetKindValue = "Deployment"

	// K8sAssetKindEndpoints represents the value Endpoints.
	K8sAssetKindEndpoints K8sAssetKindValue = "Endpoints"

	// K8sAssetKindNamespace represents the value Namespace.
	K8sAssetKindNamespace K8sAssetKindValue = "Namespace"

	// K8sAssetKindNetworkPolicy represents the value NetworkPolicy.
	K8sAssetKindNetworkPolicy K8sAssetKindValue = "NetworkPolicy"

	// K8sAssetKindPending represents the value Pending.
	K8sAssetKindPending K8sAssetKindValue = "Pending"

	// K8sAssetKindPods represents the value Pods.
	K8sAssetKindPods K8sAssetKindValue = "Pods"

	// K8sAssetKindReplicaSets represents the value ReplicaSets.
	K8sAssetKindReplicaSets K8sAssetKindValue = "ReplicaSets"

	// K8sAssetKindService represents the value Service.
	K8sAssetKindService K8sAssetKindValue = "Service"

	// K8sAssetKindStatefulSets represents the value StatefulSets.
	K8sAssetKindStatefulSets K8sAssetKindValue = "StatefulSets"
)

// K8sAssetIdentity represents the Identity of the object.
var K8sAssetIdentity = elemental.Identity{
	Name:     "k8sasset",
	Category: "k8sassets",
	Package:  "pandemona",
	Private:  false,
}

// K8sAssetsList represents a list of K8sAssets
type K8sAssetsList []*K8sAsset

// Identity returns the identity of the objects in the list.
func (o K8sAssetsList) Identity() elemental.Identity {

	return K8sAssetIdentity
}

// Copy returns a pointer to a copy the K8sAssetsList.
func (o K8sAssetsList) Copy() elemental.Identifiables {

	out := append(K8sAssetsList{}, o...)
	return &out
}

// Append appends the objects to the a new copy of the K8sAssetsList.
func (o K8sAssetsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(K8sAssetsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*K8sAsset))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o K8sAssetsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o K8sAssetsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the K8sAssetsList converted to SparseK8sAssetsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o K8sAssetsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseK8sAssetsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseK8sAsset)
	}

	return out
}

// Version returns the version of the content.
func (o K8sAssetsList) Version() int {

	return 1
}

// K8sAsset represents the model of a k8sasset
type K8sAsset struct {
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
	Kind K8sAssetKindValue `json:"kind" msgpack:"kind" bson:"kind" mapstructure:"kind,omitempty"`

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

// NewK8sAsset returns a new *K8sAsset
func NewK8sAsset() *K8sAsset {

	return &K8sAsset{
		ModelVersion:   1,
		Data:           []byte{},
		DenormedFields: []string{},
		Kind:           K8sAssetKindPending,
		MigrationsLog:  map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *K8sAsset) Identity() elemental.Identity {

	return K8sAssetIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *K8sAsset) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *K8sAsset) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *K8sAsset) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesK8sAsset{}

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
func (o *K8sAsset) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesK8sAsset{}
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
func (o *K8sAsset) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *K8sAsset) BleveType() string {

	return "k8sasset"
}

// DefaultOrder returns the list of default ordering fields.
func (o *K8sAsset) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *K8sAsset) Doc() string {

	return `Represents a read-only K8s resource such as a service.`
}

func (o *K8sAsset) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *K8sAsset) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *K8sAsset) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *K8sAsset) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *K8sAsset) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *K8sAsset) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *K8sAsset) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *K8sAsset) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *K8sAsset) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *K8sAsset) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *K8sAsset) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *K8sAsset) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *K8sAsset) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *K8sAsset) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseK8sAsset{
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

	sp := &SparseK8sAsset{}
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

// Patch apply the non nil value of a *SparseK8sAsset to the object.
func (o *K8sAsset) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseK8sAsset)
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

// DeepCopy returns a deep copy if the K8sAsset.
func (o *K8sAsset) DeepCopy() *K8sAsset {

	if o == nil {
		return nil
	}

	out := &K8sAsset{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *K8sAsset.
func (o *K8sAsset) DeepCopyInto(out *K8sAsset) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy K8sAsset: %s", err))
	}

	*out = *target.(*K8sAsset)
}

// Validate valides the current information stored into the structure.
func (o *K8sAsset) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("clusterID", o.ClusterID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredExternal("data", o.Data); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("kind", string(o.Kind), []string{"Pending", "Pods", "Deployment", "ReplicaSets", "StatefulSets", "DaemonSets", "Namespace", "NetworkPolicy", "Cluster", "Service", "Endpoints"}, true); err != nil {
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
func (*K8sAsset) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := K8sAssetAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return K8sAssetLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*K8sAsset) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return K8sAssetAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *K8sAsset) ValueForAttribute(name string) any {

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

// K8sAssetAttributesMap represents the map of attribute for K8sAsset.
var K8sAssetAttributesMap = map[string]elemental.AttributeSpecification{
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
		AllowedChoices: []string{"Pending", "Pods", "Deployment", "ReplicaSets", "StatefulSets", "DaemonSets", "Namespace", "NetworkPolicy", "Cluster", "Service", "Endpoints"},
		Autogenerated:  true,
		BSONFieldName:  "kind",
		ConvertedName:  "Kind",
		DefaultValue:   K8sAssetKindPending,
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

// K8sAssetLowerCaseAttributesMap represents the map of attribute for K8sAsset.
var K8sAssetLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		AllowedChoices: []string{"Pending", "Pods", "Deployment", "ReplicaSets", "StatefulSets", "DaemonSets", "Namespace", "NetworkPolicy", "Cluster", "Service", "Endpoints"},
		Autogenerated:  true,
		BSONFieldName:  "kind",
		ConvertedName:  "Kind",
		DefaultValue:   K8sAssetKindPending,
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

// SparseK8sAssetsList represents a list of SparseK8sAssets
type SparseK8sAssetsList []*SparseK8sAsset

// Identity returns the identity of the objects in the list.
func (o SparseK8sAssetsList) Identity() elemental.Identity {

	return K8sAssetIdentity
}

// Copy returns a pointer to a copy the SparseK8sAssetsList.
func (o SparseK8sAssetsList) Copy() elemental.Identifiables {

	copy := append(SparseK8sAssetsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseK8sAssetsList.
func (o SparseK8sAssetsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseK8sAssetsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseK8sAsset))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseK8sAssetsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseK8sAssetsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseK8sAssetsList converted to K8sAssetsList.
func (o SparseK8sAssetsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseK8sAssetsList) Version() int {

	return 1
}

// SparseK8sAsset represents the sparse version of a k8sasset.
type SparseK8sAsset struct {
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
	Kind *K8sAssetKindValue `json:"kind,omitempty" msgpack:"kind,omitempty" bson:"kind,omitempty" mapstructure:"kind,omitempty"`

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

// NewSparseK8sAsset returns a new  SparseK8sAsset.
func NewSparseK8sAsset() *SparseK8sAsset {
	return &SparseK8sAsset{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseK8sAsset) Identity() elemental.Identity {

	return K8sAssetIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseK8sAsset) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseK8sAsset) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseK8sAsset) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseK8sAsset{}

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
func (o *SparseK8sAsset) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseK8sAsset{}
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
func (o *SparseK8sAsset) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseK8sAsset) ToPlain() elemental.PlainIdentifiable {

	out := NewK8sAsset()
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
func (o *SparseK8sAsset) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseK8sAsset) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseK8sAsset) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseK8sAsset) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseK8sAsset) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseK8sAsset) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseK8sAsset) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseK8sAsset) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseK8sAsset) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseK8sAsset) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseK8sAsset) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseK8sAsset) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseK8sAsset.
func (o *SparseK8sAsset) DeepCopy() *SparseK8sAsset {

	if o == nil {
		return nil
	}

	out := &SparseK8sAsset{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseK8sAsset.
func (o *SparseK8sAsset) DeepCopyInto(out *SparseK8sAsset) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseK8sAsset: %s", err))
	}

	*out = *target.(*SparseK8sAsset)
}

type mongoAttributesK8sAsset struct {
	ID                   bson.ObjectId     `bson:"_id,omitempty"`
	ClusterID            string            `bson:"clusterid"`
	CreateTime           time.Time         `bson:"createtime"`
	Data                 []byte            `bson:"data"`
	DenormedFields       []string          `bson:"denormedfields"`
	Kind                 K8sAssetKindValue `bson:"kind"`
	MigrationsLog        map[string]string `bson:"migrationslog,omitempty"`
	Name                 string            `bson:"name"`
	Namespace            string            `bson:"namespace"`
	PrismaUnifiedAssetID string            `bson:"prismaunifiedassetid"`
	UpdateTime           time.Time         `bson:"updatetime"`
	ZHash                int               `bson:"zhash"`
	Zone                 int               `bson:"zone"`
}
type mongoAttributesSparseK8sAsset struct {
	ID                   bson.ObjectId      `bson:"_id,omitempty"`
	ClusterID            *string            `bson:"clusterid,omitempty"`
	CreateTime           *time.Time         `bson:"createtime,omitempty"`
	Data                 *[]byte            `bson:"data,omitempty"`
	DenormedFields       *[]string          `bson:"denormedfields,omitempty"`
	Kind                 *K8sAssetKindValue `bson:"kind,omitempty"`
	MigrationsLog        *map[string]string `bson:"migrationslog,omitempty"`
	Name                 *string            `bson:"name,omitempty"`
	Namespace            *string            `bson:"namespace,omitempty"`
	PrismaUnifiedAssetID *string            `bson:"prismaunifiedassetid,omitempty"`
	UpdateTime           *time.Time         `bson:"updatetime,omitempty"`
	ZHash                *int               `bson:"zhash,omitempty"`
	Zone                 *int               `bson:"zone,omitempty"`
}
