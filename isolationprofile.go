package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
	"go.aporeto.io/gaia/types"
)

// IsolationProfileIdentity represents the Identity of the object.
var IsolationProfileIdentity = elemental.Identity{
	Name:     "isolationprofile",
	Category: "isolationprofiles",
	Package:  "squall",
	Private:  false,
}

// IsolationProfilesList represents a list of IsolationProfiles
type IsolationProfilesList []*IsolationProfile

// Identity returns the identity of the objects in the list.
func (o IsolationProfilesList) Identity() elemental.Identity {

	return IsolationProfileIdentity
}

// Copy returns a pointer to a copy the IsolationProfilesList.
func (o IsolationProfilesList) Copy() elemental.Identifiables {

	copy := append(IsolationProfilesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the IsolationProfilesList.
func (o IsolationProfilesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(IsolationProfilesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*IsolationProfile))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o IsolationProfilesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o IsolationProfilesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the IsolationProfilesList converted to SparseIsolationProfilesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o IsolationProfilesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseIsolationProfilesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseIsolationProfile)
	}

	return out
}

// Version returns the version of the content.
func (o IsolationProfilesList) Version() int {

	return 1
}

// IsolationProfile represents the model of a isolationprofile
type IsolationProfile struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// The capabilities that should be added to or removed from the processing unit.
	CapabilitiesActions types.CapabilitiesTypeMap `json:"capabilitiesActions" msgpack:"capabilitiesActions" bson:"capabilitiesactions" mapstructure:"capabilitiesActions,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// The default action applied to all system calls of this profile.
	// Default is `Allow`.
	DefaultSyscallAction types.SyscallEnforcementAction `json:"defaultSyscallAction" msgpack:"defaultSyscallAction" bson:"defaultsyscallaction" mapstructure:"defaultSyscallAction,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Propagates the policy to all of its children.
	Propagate bool `json:"propagate" msgpack:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// A list of system call rules that identify actions for particular
	// system calls.
	SyscallRules types.SyscallEnforcementRulesMap `json:"syscallRules" msgpack:"syscallRules" bson:"syscallrules" mapstructure:"syscallRules,omitempty"`

	// The processor architectures that the profile supports. Default `all`.
	TargetArchitectures types.ArchitecturesTypeList `json:"targetArchitectures" msgpack:"targetArchitectures" bson:"targetarchitectures" mapstructure:"targetArchitectures,omitempty"`

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

// NewIsolationProfile returns a new *IsolationProfile
func NewIsolationProfile() *IsolationProfile {

	return &IsolationProfile{
		ModelVersion:        1,
		Annotations:         map[string][]string{},
		AssociatedTags:      []string{},
		CapabilitiesActions: types.CapabilitiesTypeMap{},
		MigrationsLog:       map[string]string{},
		NormalizedTags:      []string{},
		Metadata:            []string{},
		SyscallRules:        types.SyscallEnforcementRulesMap{},
		TargetArchitectures: types.ArchitecturesTypeList{},
	}
}

// Identity returns the Identity of the object.
func (o *IsolationProfile) Identity() elemental.Identity {

	return IsolationProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IsolationProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IsolationProfile) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *IsolationProfile) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesIsolationProfile{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CapabilitiesActions = o.CapabilitiesActions
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.DefaultSyscallAction = o.DefaultSyscallAction
	s.Description = o.Description
	s.Metadata = o.Metadata
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.Propagate = o.Propagate
	s.Protected = o.Protected
	s.SyscallRules = o.SyscallRules
	s.TargetArchitectures = o.TargetArchitectures
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *IsolationProfile) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesIsolationProfile{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.CapabilitiesActions = s.CapabilitiesActions
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.DefaultSyscallAction = s.DefaultSyscallAction
	o.Description = s.Description
	o.Metadata = s.Metadata
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.Propagate = s.Propagate
	o.Protected = s.Protected
	o.SyscallRules = s.SyscallRules
	o.TargetArchitectures = s.TargetArchitectures
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *IsolationProfile) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *IsolationProfile) BleveType() string {

	return "isolationprofile"
}

// DefaultOrder returns the list of default ordering fields.
func (o *IsolationProfile) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *IsolationProfile) Doc() string {

	return `Defines system call rules, system call actions, and other capabilities on a
processing unit.`
}

func (o *IsolationProfile) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *IsolationProfile) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *IsolationProfile) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *IsolationProfile) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *IsolationProfile) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *IsolationProfile) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *IsolationProfile) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *IsolationProfile) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *IsolationProfile) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *IsolationProfile) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *IsolationProfile) SetDescription(description string) {

	o.Description = description
}

// GetMetadata returns the Metadata of the receiver.
func (o *IsolationProfile) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *IsolationProfile) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *IsolationProfile) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *IsolationProfile) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *IsolationProfile) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *IsolationProfile) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *IsolationProfile) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *IsolationProfile) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *IsolationProfile) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *IsolationProfile) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *IsolationProfile) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the given value.
func (o *IsolationProfile) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// GetProtected returns the Protected of the receiver.
func (o *IsolationProfile) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *IsolationProfile) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *IsolationProfile) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *IsolationProfile) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *IsolationProfile) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *IsolationProfile) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *IsolationProfile) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *IsolationProfile) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *IsolationProfile) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *IsolationProfile) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *IsolationProfile) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseIsolationProfile{
			ID:                   &o.ID,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			CapabilitiesActions:  &o.CapabilitiesActions,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			DefaultSyscallAction: &o.DefaultSyscallAction,
			Description:          &o.Description,
			Metadata:             &o.Metadata,
			MigrationsLog:        &o.MigrationsLog,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Propagate:            &o.Propagate,
			Protected:            &o.Protected,
			SyscallRules:         &o.SyscallRules,
			TargetArchitectures:  &o.TargetArchitectures,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseIsolationProfile{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "capabilitiesActions":
			sp.CapabilitiesActions = &(o.CapabilitiesActions)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "defaultSyscallAction":
			sp.DefaultSyscallAction = &(o.DefaultSyscallAction)
		case "description":
			sp.Description = &(o.Description)
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
		case "propagate":
			sp.Propagate = &(o.Propagate)
		case "protected":
			sp.Protected = &(o.Protected)
		case "syscallRules":
			sp.SyscallRules = &(o.SyscallRules)
		case "targetArchitectures":
			sp.TargetArchitectures = &(o.TargetArchitectures)
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

// Patch apply the non nil value of a *SparseIsolationProfile to the object.
func (o *IsolationProfile) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseIsolationProfile)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CapabilitiesActions != nil {
		o.CapabilitiesActions = *so.CapabilitiesActions
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.DefaultSyscallAction != nil {
		o.DefaultSyscallAction = *so.DefaultSyscallAction
	}
	if so.Description != nil {
		o.Description = *so.Description
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
	if so.Propagate != nil {
		o.Propagate = *so.Propagate
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.SyscallRules != nil {
		o.SyscallRules = *so.SyscallRules
	}
	if so.TargetArchitectures != nil {
		o.TargetArchitectures = *so.TargetArchitectures
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

// DeepCopy returns a deep copy if the IsolationProfile.
func (o *IsolationProfile) DeepCopy() *IsolationProfile {

	if o == nil {
		return nil
	}

	out := &IsolationProfile{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *IsolationProfile.
func (o *IsolationProfile) DeepCopyInto(out *IsolationProfile) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy IsolationProfile: %s", err))
	}

	*out = *target.(*IsolationProfile)
}

// Validate valides the current information stored into the structure.
func (o *IsolationProfile) Validate() error {

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
func (*IsolationProfile) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := IsolationProfileAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return IsolationProfileLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*IsolationProfile) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return IsolationProfileAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *IsolationProfile) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "capabilitiesActions":
		return o.CapabilitiesActions
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "defaultSyscallAction":
		return o.DefaultSyscallAction
	case "description":
		return o.Description
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
	case "propagate":
		return o.Propagate
	case "protected":
		return o.Protected
	case "syscallRules":
		return o.SyscallRules
	case "targetArchitectures":
		return o.TargetArchitectures
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

// IsolationProfileAttributesMap represents the map of attribute for IsolationProfile.
var IsolationProfileAttributesMap = map[string]elemental.AttributeSpecification{
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
	"CapabilitiesActions": {
		AllowedChoices: []string{},
		ConvertedName:  "CapabilitiesActions",
		Description:    `The capabilities that should be added to or removed from the processing unit.`,
		Exposed:        true,
		Name:           "capabilitiesActions",
		Orderable:      true,
		Stored:         true,
		SubType:        "_cap_map",
		Type:           "external",
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
	"DefaultSyscallAction": {
		AllowedChoices: []string{},
		ConvertedName:  "DefaultSyscallAction",
		Description: `The default action applied to all system calls of this profile.
Default is ` + "`" + `Allow` + "`" + `.`,
		Exposed: true,
		Name:    "defaultSyscallAction",
		Stored:  true,
		SubType: "_syscall_action",
		Type:    "external",
	},
	"Description": {
		AllowedChoices: []string{},
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
	"Metadata": {
		AllowedChoices: []string{},
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
	"Propagate": {
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagates the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
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
	"SyscallRules": {
		AllowedChoices: []string{},
		ConvertedName:  "SyscallRules",
		Description: `A list of system call rules that identify actions for particular
system calls.`,
		Exposed:   true,
		Name:      "syscallRules",
		Orderable: true,
		Stored:    true,
		SubType:   "_syscall_rules",
		Type:      "external",
	},
	"TargetArchitectures": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetArchitectures",
		Description:    `The processor architectures that the profile supports. Default ` + "`" + `all` + "`" + `.`,
		Exposed:        true,
		Name:           "targetArchitectures",
		Orderable:      true,
		Stored:         true,
		SubType:        "_arch_list",
		Type:           "external",
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

// IsolationProfileLowerCaseAttributesMap represents the map of attribute for IsolationProfile.
var IsolationProfileLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"capabilitiesactions": {
		AllowedChoices: []string{},
		ConvertedName:  "CapabilitiesActions",
		Description:    `The capabilities that should be added to or removed from the processing unit.`,
		Exposed:        true,
		Name:           "capabilitiesActions",
		Orderable:      true,
		Stored:         true,
		SubType:        "_cap_map",
		Type:           "external",
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
	"defaultsyscallaction": {
		AllowedChoices: []string{},
		ConvertedName:  "DefaultSyscallAction",
		Description: `The default action applied to all system calls of this profile.
Default is ` + "`" + `Allow` + "`" + `.`,
		Exposed: true,
		Name:    "defaultSyscallAction",
		Stored:  true,
		SubType: "_syscall_action",
		Type:    "external",
	},
	"description": {
		AllowedChoices: []string{},
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
	"metadata": {
		AllowedChoices: []string{},
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
	"propagate": {
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagates the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
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
	"syscallrules": {
		AllowedChoices: []string{},
		ConvertedName:  "SyscallRules",
		Description: `A list of system call rules that identify actions for particular
system calls.`,
		Exposed:   true,
		Name:      "syscallRules",
		Orderable: true,
		Stored:    true,
		SubType:   "_syscall_rules",
		Type:      "external",
	},
	"targetarchitectures": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetArchitectures",
		Description:    `The processor architectures that the profile supports. Default ` + "`" + `all` + "`" + `.`,
		Exposed:        true,
		Name:           "targetArchitectures",
		Orderable:      true,
		Stored:         true,
		SubType:        "_arch_list",
		Type:           "external",
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

// SparseIsolationProfilesList represents a list of SparseIsolationProfiles
type SparseIsolationProfilesList []*SparseIsolationProfile

// Identity returns the identity of the objects in the list.
func (o SparseIsolationProfilesList) Identity() elemental.Identity {

	return IsolationProfileIdentity
}

// Copy returns a pointer to a copy the SparseIsolationProfilesList.
func (o SparseIsolationProfilesList) Copy() elemental.Identifiables {

	copy := append(SparseIsolationProfilesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseIsolationProfilesList.
func (o SparseIsolationProfilesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseIsolationProfilesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseIsolationProfile))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseIsolationProfilesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseIsolationProfilesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseIsolationProfilesList converted to IsolationProfilesList.
func (o SparseIsolationProfilesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseIsolationProfilesList) Version() int {

	return 1
}

// SparseIsolationProfile represents the sparse version of a isolationprofile.
type SparseIsolationProfile struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// The capabilities that should be added to or removed from the processing unit.
	CapabilitiesActions *types.CapabilitiesTypeMap `json:"capabilitiesActions,omitempty" msgpack:"capabilitiesActions,omitempty" bson:"capabilitiesactions,omitempty" mapstructure:"capabilitiesActions,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// The default action applied to all system calls of this profile.
	// Default is `Allow`.
	DefaultSyscallAction *types.SyscallEnforcementAction `json:"defaultSyscallAction,omitempty" msgpack:"defaultSyscallAction,omitempty" bson:"defaultsyscallaction,omitempty" mapstructure:"defaultSyscallAction,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

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

	// Propagates the policy to all of its children.
	Propagate *bool `json:"propagate,omitempty" msgpack:"propagate,omitempty" bson:"propagate,omitempty" mapstructure:"propagate,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// A list of system call rules that identify actions for particular
	// system calls.
	SyscallRules *types.SyscallEnforcementRulesMap `json:"syscallRules,omitempty" msgpack:"syscallRules,omitempty" bson:"syscallrules,omitempty" mapstructure:"syscallRules,omitempty"`

	// The processor architectures that the profile supports. Default `all`.
	TargetArchitectures *types.ArchitecturesTypeList `json:"targetArchitectures,omitempty" msgpack:"targetArchitectures,omitempty" bson:"targetarchitectures,omitempty" mapstructure:"targetArchitectures,omitempty"`

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

// NewSparseIsolationProfile returns a new  SparseIsolationProfile.
func NewSparseIsolationProfile() *SparseIsolationProfile {
	return &SparseIsolationProfile{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseIsolationProfile) Identity() elemental.Identity {

	return IsolationProfileIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseIsolationProfile) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseIsolationProfile) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseIsolationProfile) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseIsolationProfile{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.CapabilitiesActions != nil {
		s.CapabilitiesActions = o.CapabilitiesActions
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.DefaultSyscallAction != nil {
		s.DefaultSyscallAction = o.DefaultSyscallAction
	}
	if o.Description != nil {
		s.Description = o.Description
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
	if o.Propagate != nil {
		s.Propagate = o.Propagate
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.SyscallRules != nil {
		s.SyscallRules = o.SyscallRules
	}
	if o.TargetArchitectures != nil {
		s.TargetArchitectures = o.TargetArchitectures
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
func (o *SparseIsolationProfile) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseIsolationProfile{}
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
	if s.CapabilitiesActions != nil {
		o.CapabilitiesActions = s.CapabilitiesActions
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.DefaultSyscallAction != nil {
		o.DefaultSyscallAction = s.DefaultSyscallAction
	}
	if s.Description != nil {
		o.Description = s.Description
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
	if s.Propagate != nil {
		o.Propagate = s.Propagate
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.SyscallRules != nil {
		o.SyscallRules = s.SyscallRules
	}
	if s.TargetArchitectures != nil {
		o.TargetArchitectures = s.TargetArchitectures
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
func (o *SparseIsolationProfile) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseIsolationProfile) ToPlain() elemental.PlainIdentifiable {

	out := NewIsolationProfile()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CapabilitiesActions != nil {
		out.CapabilitiesActions = *o.CapabilitiesActions
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.DefaultSyscallAction != nil {
		out.DefaultSyscallAction = *o.DefaultSyscallAction
	}
	if o.Description != nil {
		out.Description = *o.Description
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
	if o.Propagate != nil {
		out.Propagate = *o.Propagate
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.SyscallRules != nil {
		out.SyscallRules = *o.SyscallRules
	}
	if o.TargetArchitectures != nil {
		out.TargetArchitectures = *o.TargetArchitectures
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
func (o *SparseIsolationProfile) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseIsolationProfile) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseIsolationProfile) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseIsolationProfile) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseIsolationProfile) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetDescription(description string) {

	o.Description = &description
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseIsolationProfile) GetMetadata() (out []string) {

	if o.Metadata == nil {
		return
	}

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseIsolationProfile) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseIsolationProfile) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseIsolationProfile) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseIsolationProfile) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *SparseIsolationProfile) GetPropagate() (out bool) {

	if o.Propagate == nil {
		return
	}

	return *o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetPropagate(propagate bool) {

	o.Propagate = &propagate
}

// GetProtected returns the Protected of the receiver.
func (o *SparseIsolationProfile) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseIsolationProfile) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseIsolationProfile) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseIsolationProfile) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseIsolationProfile) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseIsolationProfile.
func (o *SparseIsolationProfile) DeepCopy() *SparseIsolationProfile {

	if o == nil {
		return nil
	}

	out := &SparseIsolationProfile{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseIsolationProfile.
func (o *SparseIsolationProfile) DeepCopyInto(out *SparseIsolationProfile) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseIsolationProfile: %s", err))
	}

	*out = *target.(*SparseIsolationProfile)
}

type mongoAttributesIsolationProfile struct {
	ID                   bson.ObjectId                    `bson:"_id,omitempty"`
	Annotations          map[string][]string              `bson:"annotations"`
	AssociatedTags       []string                         `bson:"associatedtags"`
	CapabilitiesActions  types.CapabilitiesTypeMap        `bson:"capabilitiesactions"`
	CreateIdempotencyKey string                           `bson:"createidempotencykey"`
	CreateTime           time.Time                        `bson:"createtime"`
	DefaultSyscallAction types.SyscallEnforcementAction   `bson:"defaultsyscallaction"`
	Description          string                           `bson:"description"`
	Metadata             []string                         `bson:"metadata"`
	MigrationsLog        map[string]string                `bson:"migrationslog,omitempty"`
	Name                 string                           `bson:"name"`
	Namespace            string                           `bson:"namespace"`
	NormalizedTags       []string                         `bson:"normalizedtags"`
	Propagate            bool                             `bson:"propagate"`
	Protected            bool                             `bson:"protected"`
	SyscallRules         types.SyscallEnforcementRulesMap `bson:"syscallrules"`
	TargetArchitectures  types.ArchitecturesTypeList      `bson:"targetarchitectures"`
	UpdateIdempotencyKey string                           `bson:"updateidempotencykey"`
	UpdateTime           time.Time                        `bson:"updatetime"`
	ZHash                int                              `bson:"zhash"`
	Zone                 int                              `bson:"zone"`
}
type mongoAttributesSparseIsolationProfile struct {
	ID                   bson.ObjectId                     `bson:"_id,omitempty"`
	Annotations          *map[string][]string              `bson:"annotations,omitempty"`
	AssociatedTags       *[]string                         `bson:"associatedtags,omitempty"`
	CapabilitiesActions  *types.CapabilitiesTypeMap        `bson:"capabilitiesactions,omitempty"`
	CreateIdempotencyKey *string                           `bson:"createidempotencykey,omitempty"`
	CreateTime           *time.Time                        `bson:"createtime,omitempty"`
	DefaultSyscallAction *types.SyscallEnforcementAction   `bson:"defaultsyscallaction,omitempty"`
	Description          *string                           `bson:"description,omitempty"`
	Metadata             *[]string                         `bson:"metadata,omitempty"`
	MigrationsLog        *map[string]string                `bson:"migrationslog,omitempty"`
	Name                 *string                           `bson:"name,omitempty"`
	Namespace            *string                           `bson:"namespace,omitempty"`
	NormalizedTags       *[]string                         `bson:"normalizedtags,omitempty"`
	Propagate            *bool                             `bson:"propagate,omitempty"`
	Protected            *bool                             `bson:"protected,omitempty"`
	SyscallRules         *types.SyscallEnforcementRulesMap `bson:"syscallrules,omitempty"`
	TargetArchitectures  *types.ArchitecturesTypeList      `bson:"targetarchitectures,omitempty"`
	UpdateIdempotencyKey *string                           `bson:"updateidempotencykey,omitempty"`
	UpdateTime           *time.Time                        `bson:"updatetime,omitempty"`
	ZHash                *int                              `bson:"zhash,omitempty"`
	Zone                 *int                              `bson:"zone,omitempty"`
}
