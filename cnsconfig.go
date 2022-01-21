package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CNSConfigIdentity represents the Identity of the object.
var CNSConfigIdentity = elemental.Identity{
	Name:     "cnsconfig",
	Category: "cnsconfigs",
	Package:  "karl",
	Private:  false,
}

// CNSConfigsList represents a list of CNSConfigs
type CNSConfigsList []*CNSConfig

// Identity returns the identity of the objects in the list.
func (o CNSConfigsList) Identity() elemental.Identity {

	return CNSConfigIdentity
}

// Copy returns a pointer to a copy the CNSConfigsList.
func (o CNSConfigsList) Copy() elemental.Identifiables {

	copy := append(CNSConfigsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CNSConfigsList.
func (o CNSConfigsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CNSConfigsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CNSConfig))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CNSConfigsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CNSConfigsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CNSConfigsList converted to SparseCNSConfigsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CNSConfigsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCNSConfigsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCNSConfig)
	}

	return out
}

// Version returns the version of the content.
func (o CNSConfigsList) Version() int {

	return 1
}

// CNSConfig represents the model of a cnsconfig
type CNSConfig struct {
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

	// If `true` net effective permissions feature is enabled.
	EnableNetEffectivePermissions bool `json:"enableNetEffectivePermissions" msgpack:"enableNetEffectivePermissions" bson:"enableneteffectivepermissions" mapstructure:"enableNetEffectivePermissions,omitempty"`

	// If `true` network security feature is enabled.
	EnableNetworkSecurity bool `json:"enableNetworkSecurity" msgpack:"enableNetworkSecurity" bson:"enablenetworksecurity" mapstructure:"enableNetworkSecurity,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Unique Prisma ID identifying the CNS configuration.
	PrismaID string `json:"prismaID" msgpack:"prismaID" bson:"prismaid" mapstructure:"prismaID,omitempty"`

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

// NewCNSConfig returns a new *CNSConfig
func NewCNSConfig() *CNSConfig {

	return &CNSConfig{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		MigrationsLog:  map[string]string{},
		NormalizedTags: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *CNSConfig) Identity() elemental.Identity {

	return CNSConfigIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CNSConfig) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CNSConfig) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CNSConfig) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCNSConfig{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.EnableNetEffectivePermissions = o.EnableNetEffectivePermissions
	s.EnableNetworkSecurity = o.EnableNetworkSecurity
	s.MigrationsLog = o.MigrationsLog
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.PrismaID = o.PrismaID
	s.Protected = o.Protected
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CNSConfig) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCNSConfig{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.EnableNetEffectivePermissions = s.EnableNetEffectivePermissions
	o.EnableNetworkSecurity = s.EnableNetworkSecurity
	o.MigrationsLog = s.MigrationsLog
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.PrismaID = s.PrismaID
	o.Protected = s.Protected
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CNSConfig) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CNSConfig) BleveType() string {

	return "cnsconfig"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CNSConfig) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CNSConfig) Doc() string {

	return `Holds the CNS configuration for a namespace.`
}

func (o *CNSConfig) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *CNSConfig) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *CNSConfig) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *CNSConfig) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *CNSConfig) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *CNSConfig) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *CNSConfig) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *CNSConfig) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *CNSConfig) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *CNSConfig) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *CNSConfig) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *CNSConfig) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *CNSConfig) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *CNSConfig) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *CNSConfig) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *CNSConfig) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *CNSConfig) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *CNSConfig) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *CNSConfig) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *CNSConfig) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *CNSConfig) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *CNSConfig) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *CNSConfig) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *CNSConfig) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *CNSConfig) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CNSConfig) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCNSConfig{
			ID:                            &o.ID,
			Annotations:                   &o.Annotations,
			AssociatedTags:                &o.AssociatedTags,
			CreateIdempotencyKey:          &o.CreateIdempotencyKey,
			CreateTime:                    &o.CreateTime,
			EnableNetEffectivePermissions: &o.EnableNetEffectivePermissions,
			EnableNetworkSecurity:         &o.EnableNetworkSecurity,
			MigrationsLog:                 &o.MigrationsLog,
			Namespace:                     &o.Namespace,
			NormalizedTags:                &o.NormalizedTags,
			PrismaID:                      &o.PrismaID,
			Protected:                     &o.Protected,
			UpdateIdempotencyKey:          &o.UpdateIdempotencyKey,
			UpdateTime:                    &o.UpdateTime,
			ZHash:                         &o.ZHash,
			Zone:                          &o.Zone,
		}
	}

	sp := &SparseCNSConfig{}
	for _, f := range fields {
		switch f {
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
		case "enableNetEffectivePermissions":
			sp.EnableNetEffectivePermissions = &(o.EnableNetEffectivePermissions)
		case "enableNetworkSecurity":
			sp.EnableNetworkSecurity = &(o.EnableNetworkSecurity)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "prismaID":
			sp.PrismaID = &(o.PrismaID)
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

// Patch apply the non nil value of a *SparseCNSConfig to the object.
func (o *CNSConfig) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCNSConfig)
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
	if so.EnableNetEffectivePermissions != nil {
		o.EnableNetEffectivePermissions = *so.EnableNetEffectivePermissions
	}
	if so.EnableNetworkSecurity != nil {
		o.EnableNetworkSecurity = *so.EnableNetworkSecurity
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
	if so.PrismaID != nil {
		o.PrismaID = *so.PrismaID
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

// DeepCopy returns a deep copy if the CNSConfig.
func (o *CNSConfig) DeepCopy() *CNSConfig {

	if o == nil {
		return nil
	}

	out := &CNSConfig{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CNSConfig.
func (o *CNSConfig) DeepCopyInto(out *CNSConfig) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CNSConfig: %s", err))
	}

	*out = *target.(*CNSConfig)
}

// Validate valides the current information stored into the structure.
func (o *CNSConfig) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
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
func (*CNSConfig) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CNSConfigAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CNSConfigLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CNSConfig) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CNSConfigAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CNSConfig) ValueForAttribute(name string) interface{} {

	switch name {
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
	case "enableNetEffectivePermissions":
		return o.EnableNetEffectivePermissions
	case "enableNetworkSecurity":
		return o.EnableNetworkSecurity
	case "migrationsLog":
		return o.MigrationsLog
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "prismaID":
		return o.PrismaID
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

// CNSConfigAttributesMap represents the map of attribute for CNSConfig.
var CNSConfigAttributesMap = map[string]elemental.AttributeSpecification{
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
	"EnableNetEffectivePermissions": {
		AllowedChoices: []string{},
		ConvertedName:  "EnableNetEffectivePermissions",
		Description:    `If ` + "`" + `true` + "`" + ` net effective permissions feature is enabled.`,
		Exposed:        true,
		Name:           "enableNetEffectivePermissions",
		Stored:         true,
		Type:           "boolean",
	},
	"EnableNetworkSecurity": {
		AllowedChoices: []string{},
		ConvertedName:  "EnableNetworkSecurity",
		Description:    `If ` + "`" + `true` + "`" + ` network security feature is enabled.`,
		Exposed:        true,
		Name:           "enableNetworkSecurity",
		Stored:         true,
		Type:           "boolean",
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
	"PrismaID": {
		AllowedChoices: []string{},
		ConvertedName:  "PrismaID",
		CreationOnly:   true,
		Description:    `Unique Prisma ID identifying the CNS configuration.`,
		Exposed:        true,
		Name:           "prismaID",
		Stored:         true,
		Type:           "string",
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

// CNSConfigLowerCaseAttributesMap represents the map of attribute for CNSConfig.
var CNSConfigLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"enableneteffectivepermissions": {
		AllowedChoices: []string{},
		ConvertedName:  "EnableNetEffectivePermissions",
		Description:    `If ` + "`" + `true` + "`" + ` net effective permissions feature is enabled.`,
		Exposed:        true,
		Name:           "enableNetEffectivePermissions",
		Stored:         true,
		Type:           "boolean",
	},
	"enablenetworksecurity": {
		AllowedChoices: []string{},
		ConvertedName:  "EnableNetworkSecurity",
		Description:    `If ` + "`" + `true` + "`" + ` network security feature is enabled.`,
		Exposed:        true,
		Name:           "enableNetworkSecurity",
		Stored:         true,
		Type:           "boolean",
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
	"prismaid": {
		AllowedChoices: []string{},
		ConvertedName:  "PrismaID",
		CreationOnly:   true,
		Description:    `Unique Prisma ID identifying the CNS configuration.`,
		Exposed:        true,
		Name:           "prismaID",
		Stored:         true,
		Type:           "string",
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

// SparseCNSConfigsList represents a list of SparseCNSConfigs
type SparseCNSConfigsList []*SparseCNSConfig

// Identity returns the identity of the objects in the list.
func (o SparseCNSConfigsList) Identity() elemental.Identity {

	return CNSConfigIdentity
}

// Copy returns a pointer to a copy the SparseCNSConfigsList.
func (o SparseCNSConfigsList) Copy() elemental.Identifiables {

	copy := append(SparseCNSConfigsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCNSConfigsList.
func (o SparseCNSConfigsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCNSConfigsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCNSConfig))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCNSConfigsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCNSConfigsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCNSConfigsList converted to CNSConfigsList.
func (o SparseCNSConfigsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCNSConfigsList) Version() int {

	return 1
}

// SparseCNSConfig represents the sparse version of a cnsconfig.
type SparseCNSConfig struct {
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

	// If `true` net effective permissions feature is enabled.
	EnableNetEffectivePermissions *bool `json:"enableNetEffectivePermissions,omitempty" msgpack:"enableNetEffectivePermissions,omitempty" bson:"enableneteffectivepermissions,omitempty" mapstructure:"enableNetEffectivePermissions,omitempty"`

	// If `true` network security feature is enabled.
	EnableNetworkSecurity *bool `json:"enableNetworkSecurity,omitempty" msgpack:"enableNetworkSecurity,omitempty" bson:"enablenetworksecurity,omitempty" mapstructure:"enableNetworkSecurity,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Unique Prisma ID identifying the CNS configuration.
	PrismaID *string `json:"prismaID,omitempty" msgpack:"prismaID,omitempty" bson:"prismaid,omitempty" mapstructure:"prismaID,omitempty"`

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

// NewSparseCNSConfig returns a new  SparseCNSConfig.
func NewSparseCNSConfig() *SparseCNSConfig {
	return &SparseCNSConfig{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCNSConfig) Identity() elemental.Identity {

	return CNSConfigIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCNSConfig) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCNSConfig) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCNSConfig) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCNSConfig{}

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
	if o.EnableNetEffectivePermissions != nil {
		s.EnableNetEffectivePermissions = o.EnableNetEffectivePermissions
	}
	if o.EnableNetworkSecurity != nil {
		s.EnableNetworkSecurity = o.EnableNetworkSecurity
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
	if o.PrismaID != nil {
		s.PrismaID = o.PrismaID
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
func (o *SparseCNSConfig) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCNSConfig{}
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
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.EnableNetEffectivePermissions != nil {
		o.EnableNetEffectivePermissions = s.EnableNetEffectivePermissions
	}
	if s.EnableNetworkSecurity != nil {
		o.EnableNetworkSecurity = s.EnableNetworkSecurity
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
	if s.PrismaID != nil {
		o.PrismaID = s.PrismaID
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
func (o *SparseCNSConfig) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCNSConfig) ToPlain() elemental.PlainIdentifiable {

	out := NewCNSConfig()
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
	if o.EnableNetEffectivePermissions != nil {
		out.EnableNetEffectivePermissions = *o.EnableNetEffectivePermissions
	}
	if o.EnableNetworkSecurity != nil {
		out.EnableNetworkSecurity = *o.EnableNetworkSecurity
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
	if o.PrismaID != nil {
		out.PrismaID = *o.PrismaID
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
func (o *SparseCNSConfig) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseCNSConfig) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseCNSConfig) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseCNSConfig) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseCNSConfig) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCNSConfig) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseCNSConfig) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseCNSConfig) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseCNSConfig) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseCNSConfig) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseCNSConfig) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseCNSConfig) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseCNSConfig) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseCNSConfig) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseCNSConfig) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseCNSConfig) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseCNSConfig) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCNSConfig) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseCNSConfig) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseCNSConfig) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseCNSConfig) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseCNSConfig) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseCNSConfig) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseCNSConfig) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseCNSConfig.
func (o *SparseCNSConfig) DeepCopy() *SparseCNSConfig {

	if o == nil {
		return nil
	}

	out := &SparseCNSConfig{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCNSConfig.
func (o *SparseCNSConfig) DeepCopyInto(out *SparseCNSConfig) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCNSConfig: %s", err))
	}

	*out = *target.(*SparseCNSConfig)
}

type mongoAttributesCNSConfig struct {
	ID                            bson.ObjectId       `bson:"_id,omitempty"`
	Annotations                   map[string][]string `bson:"annotations"`
	AssociatedTags                []string            `bson:"associatedtags"`
	CreateIdempotencyKey          string              `bson:"createidempotencykey"`
	CreateTime                    time.Time           `bson:"createtime"`
	EnableNetEffectivePermissions bool                `bson:"enableneteffectivepermissions"`
	EnableNetworkSecurity         bool                `bson:"enablenetworksecurity"`
	MigrationsLog                 map[string]string   `bson:"migrationslog,omitempty"`
	Namespace                     string              `bson:"namespace"`
	NormalizedTags                []string            `bson:"normalizedtags"`
	PrismaID                      string              `bson:"prismaid"`
	Protected                     bool                `bson:"protected"`
	UpdateIdempotencyKey          string              `bson:"updateidempotencykey"`
	UpdateTime                    time.Time           `bson:"updatetime"`
	ZHash                         int                 `bson:"zhash"`
	Zone                          int                 `bson:"zone"`
}
type mongoAttributesSparseCNSConfig struct {
	ID                            bson.ObjectId        `bson:"_id,omitempty"`
	Annotations                   *map[string][]string `bson:"annotations,omitempty"`
	AssociatedTags                *[]string            `bson:"associatedtags,omitempty"`
	CreateIdempotencyKey          *string              `bson:"createidempotencykey,omitempty"`
	CreateTime                    *time.Time           `bson:"createtime,omitempty"`
	EnableNetEffectivePermissions *bool                `bson:"enableneteffectivepermissions,omitempty"`
	EnableNetworkSecurity         *bool                `bson:"enablenetworksecurity,omitempty"`
	MigrationsLog                 *map[string]string   `bson:"migrationslog,omitempty"`
	Namespace                     *string              `bson:"namespace,omitempty"`
	NormalizedTags                *[]string            `bson:"normalizedtags,omitempty"`
	PrismaID                      *string              `bson:"prismaid,omitempty"`
	Protected                     *bool                `bson:"protected,omitempty"`
	UpdateIdempotencyKey          *string              `bson:"updateidempotencykey,omitempty"`
	UpdateTime                    *time.Time           `bson:"updatetime,omitempty"`
	ZHash                         *int                 `bson:"zhash,omitempty"`
	Zone                          *int                 `bson:"zone,omitempty"`
}
