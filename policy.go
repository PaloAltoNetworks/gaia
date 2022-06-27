package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PolicyTypeValue represents the possible values for attribute "type".
type PolicyTypeValue string

const (
	// PolicyTypeAPIAuthorization represents the value APIAuthorization.
	PolicyTypeAPIAuthorization PolicyTypeValue = "APIAuthorization"

	// PolicyTypeAuditProfileMapping represents the value AuditProfileMapping.
	PolicyTypeAuditProfileMapping PolicyTypeValue = "AuditProfileMapping"

	// PolicyTypeEnforcerProfile represents the value EnforcerProfile.
	PolicyTypeEnforcerProfile PolicyTypeValue = "EnforcerProfile"

	// PolicyTypeFile represents the value File.
	PolicyTypeFile PolicyTypeValue = "File"

	// PolicyTypeHook represents the value Hook.
	PolicyTypeHook PolicyTypeValue = "Hook"

	// PolicyTypeHostServiceMapping represents the value HostServiceMapping.
	PolicyTypeHostServiceMapping PolicyTypeValue = "HostServiceMapping"

	// PolicyTypeInfrastructure represents the value Infrastructure.
	PolicyTypeInfrastructure PolicyTypeValue = "Infrastructure"

	// PolicyTypeNamespaceMapping represents the value NamespaceMapping.
	PolicyTypeNamespaceMapping PolicyTypeValue = "NamespaceMapping"

	// PolicyTypeNetwork represents the value Network.
	PolicyTypeNetwork PolicyTypeValue = "Network"

	// PolicyTypeNetworkRuleSet represents the value NetworkRuleSet.
	PolicyTypeNetworkRuleSet PolicyTypeValue = "NetworkRuleSet"

	// PolicyTypeProcessingUnit represents the value ProcessingUnit.
	PolicyTypeProcessingUnit PolicyTypeValue = "ProcessingUnit"

	// PolicyTypeQuota represents the value Quota.
	PolicyTypeQuota PolicyTypeValue = "Quota"

	// PolicyTypeSSHAuthorization represents the value SSHAuthorization.
	PolicyTypeSSHAuthorization PolicyTypeValue = "SSHAuthorization"

	// PolicyTypeService represents the value Service.
	PolicyTypeService PolicyTypeValue = "Service"

	// PolicyTypeServiceDependency represents the value ServiceDependency.
	PolicyTypeServiceDependency PolicyTypeValue = "ServiceDependency"

	// PolicyTypeSyscall represents the value Syscall.
	PolicyTypeSyscall PolicyTypeValue = "Syscall"

	// PolicyTypeTokenScope represents the value TokenScope.
	PolicyTypeTokenScope PolicyTypeValue = "TokenScope"

	// PolicyTypeUserAccess represents the value UserAccess.
	PolicyTypeUserAccess PolicyTypeValue = "UserAccess"
)

// PolicyIdentity represents the Identity of the object.
var PolicyIdentity = elemental.Identity{
	Name:     "policy",
	Category: "policies",
	Package:  "squall",
	Private:  false,
}

// PoliciesList represents a list of Policies
type PoliciesList []*Policy

// Identity returns the identity of the objects in the list.
func (o PoliciesList) Identity() elemental.Identity {

	return PolicyIdentity
}

// Copy returns a pointer to a copy the PoliciesList.
func (o PoliciesList) Copy() elemental.Identifiables {

	copy := append(PoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PoliciesList.
func (o PoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Policy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PoliciesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the PoliciesList converted to SparsePoliciesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PoliciesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePoliciesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePolicy)
	}

	return out
}

// Version returns the version of the content.
func (o PoliciesList) Version() int {

	return 1
}

// Policy represents the model of a policy
type Policy struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Defines a set of actions that must be enforced when a dependency is met.
	Action map[string]map[string]interface{} `json:"action" msgpack:"action" bson:"action" mapstructure:"action,omitempty"`

	// Defines for how long the policy will be active according to the
	// `activeSchedule`.
	ActiveDuration string `json:"activeDuration" msgpack:"activeDuration" bson:"activeduration" mapstructure:"activeDuration,omitempty"`

	// Defines when the policy should be active using the cron notation.
	// The policy will be active for the given `activeDuration`.
	ActiveSchedule string `json:"activeSchedule" msgpack:"activeSchedule" bson:"activeschedule" mapstructure:"activeSchedule,omitempty"`

	// This is a set of all object tags for matching in the DB.
	AllObjectTags []string `json:"-" msgpack:"-" bson:"allobjecttags" mapstructure:"-,omitempty"`

	// This is a set of all subject tags for matching in the DB.
	AllSubjectTags []string `json:"-" msgpack:"-" bson:"allsubjecttags" mapstructure:"-,omitempty"`

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

	// Defines if the property is disabled.
	Disabled bool `json:"disabled" msgpack:"disabled" bson:"disabled" mapstructure:"disabled,omitempty"`

	// If set the policy will be automatically deleted at the given time.
	ExpirationTime time.Time `json:"expirationTime" msgpack:"expirationTime" bson:"expirationtime" mapstructure:"expirationTime,omitempty"`

	// Indicates that this is fallback policy. It will only be
	// applied if no other policies have been resolved. If the policy is also
	// propagated it will become a fallback for children namespaces.
	Fallback bool `json:"fallback" msgpack:"fallback" bson:"fallback" mapstructure:"fallback,omitempty"`

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

	// Represents set of entities that another entity depends on. As subjects,
	// objects are identified as logical operations on tags when a policy is defined.
	Object [][]string `json:"object" msgpack:"object" bson:"object" mapstructure:"object,omitempty"`

	// Propagates the policy to all of its children.
	Propagate bool `json:"propagate" msgpack:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	// If set to `true` while the policy is propagating, it won't be visible to
	// children
	// namespace, but still used for policy resolution.
	PropagationHidden bool `json:"propagationHidden" msgpack:"propagationHidden" bson:"propagationhidden" mapstructure:"propagationHidden,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Describes the required operation to be performed between subjects and
	// objects.
	Relation []string `json:"relation" msgpack:"relation" bson:"relation" mapstructure:"relation,omitempty"`

	// Represents sets of entities that will have a dependency other entities.
	// Subjects are defined as logical operations on tags. Logical operations can
	// include `AND` and `OR`.
	Subject [][]string `json:"subject" msgpack:"subject" bson:"subject" mapstructure:"subject,omitempty"`

	// Type of the policy.
	Type PolicyTypeValue `json:"type" msgpack:"type" bson:"type" mapstructure:"type,omitempty"`

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

// NewPolicy returns a new *Policy
func NewPolicy() *Policy {

	return &Policy{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AllObjectTags:  []string{},
		AllSubjectTags: []string{},
		AssociatedTags: []string{},
		MigrationsLog:  map[string]string{},
		Metadata:       []string{},
		Relation:       []string{},
		NormalizedTags: []string{},
		Object:         [][]string{},
		Subject:        [][]string{},
	}
}

// Identity returns the Identity of the object.
func (o *Policy) Identity() elemental.Identity {

	return PolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Policy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Policy) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Policy) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPolicy{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Action = o.Action
	s.ActiveDuration = o.ActiveDuration
	s.ActiveSchedule = o.ActiveSchedule
	s.AllObjectTags = o.AllObjectTags
	s.AllSubjectTags = o.AllSubjectTags
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.Description = o.Description
	s.Disabled = o.Disabled
	s.ExpirationTime = o.ExpirationTime
	s.Fallback = o.Fallback
	s.Metadata = o.Metadata
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.Object = o.Object
	s.Propagate = o.Propagate
	s.PropagationHidden = o.PropagationHidden
	s.Protected = o.Protected
	s.Relation = o.Relation
	s.Subject = o.Subject
	s.Type = o.Type
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Policy) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPolicy{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Action = s.Action
	o.ActiveDuration = s.ActiveDuration
	o.ActiveSchedule = s.ActiveSchedule
	o.AllObjectTags = s.AllObjectTags
	o.AllSubjectTags = s.AllSubjectTags
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.Description = s.Description
	o.Disabled = s.Disabled
	o.ExpirationTime = s.ExpirationTime
	o.Fallback = s.Fallback
	o.Metadata = s.Metadata
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.Object = s.Object
	o.Propagate = s.Propagate
	o.PropagationHidden = s.PropagationHidden
	o.Protected = s.Protected
	o.Relation = s.Relation
	o.Subject = s.Subject
	o.Type = s.Type
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Policy) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Policy) BleveType() string {

	return "policy"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Policy) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *Policy) Doc() string {

	return `Represents the policy primitive used by all Microsegmentation policies.`
}

func (o *Policy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetActiveDuration returns the ActiveDuration of the receiver.
func (o *Policy) GetActiveDuration() string {

	return o.ActiveDuration
}

// SetActiveDuration sets the property ActiveDuration of the receiver using the given value.
func (o *Policy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *Policy) GetActiveSchedule() string {

	return o.ActiveSchedule
}

// SetActiveSchedule sets the property ActiveSchedule of the receiver using the given value.
func (o *Policy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = activeSchedule
}

// GetAnnotations returns the Annotations of the receiver.
func (o *Policy) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *Policy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *Policy) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *Policy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *Policy) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *Policy) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Policy) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *Policy) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *Policy) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *Policy) SetDescription(description string) {

	o.Description = description
}

// GetDisabled returns the Disabled of the receiver.
func (o *Policy) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the given value.
func (o *Policy) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetExpirationTime returns the ExpirationTime of the receiver.
func (o *Policy) GetExpirationTime() time.Time {

	return o.ExpirationTime
}

// SetExpirationTime sets the property ExpirationTime of the receiver using the given value.
func (o *Policy) SetExpirationTime(expirationTime time.Time) {

	o.ExpirationTime = expirationTime
}

// GetFallback returns the Fallback of the receiver.
func (o *Policy) GetFallback() bool {

	return o.Fallback
}

// SetFallback sets the property Fallback of the receiver using the given value.
func (o *Policy) SetFallback(fallback bool) {

	o.Fallback = fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *Policy) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *Policy) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *Policy) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *Policy) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *Policy) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *Policy) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *Policy) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *Policy) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *Policy) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *Policy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetObject returns the Object of the receiver.
func (o *Policy) GetObject() [][]string {

	return o.Object
}

// SetObject sets the property Object of the receiver using the given value.
func (o *Policy) SetObject(object [][]string) {

	o.Object = object
}

// GetPropagate returns the Propagate of the receiver.
func (o *Policy) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the given value.
func (o *Policy) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// GetPropagationHidden returns the PropagationHidden of the receiver.
func (o *Policy) GetPropagationHidden() bool {

	return o.PropagationHidden
}

// SetPropagationHidden sets the property PropagationHidden of the receiver using the given value.
func (o *Policy) SetPropagationHidden(propagationHidden bool) {

	o.PropagationHidden = propagationHidden
}

// GetProtected returns the Protected of the receiver.
func (o *Policy) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *Policy) SetProtected(protected bool) {

	o.Protected = protected
}

// GetSubject returns the Subject of the receiver.
func (o *Policy) GetSubject() [][]string {

	return o.Subject
}

// SetSubject sets the property Subject of the receiver using the given value.
func (o *Policy) SetSubject(subject [][]string) {

	o.Subject = subject
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *Policy) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *Policy) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *Policy) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *Policy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *Policy) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *Policy) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *Policy) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *Policy) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Policy) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePolicy{
			ID:                   &o.ID,
			Action:               &o.Action,
			ActiveDuration:       &o.ActiveDuration,
			ActiveSchedule:       &o.ActiveSchedule,
			AllObjectTags:        &o.AllObjectTags,
			AllSubjectTags:       &o.AllSubjectTags,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			Description:          &o.Description,
			Disabled:             &o.Disabled,
			ExpirationTime:       &o.ExpirationTime,
			Fallback:             &o.Fallback,
			Metadata:             &o.Metadata,
			MigrationsLog:        &o.MigrationsLog,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Object:               &o.Object,
			Propagate:            &o.Propagate,
			PropagationHidden:    &o.PropagationHidden,
			Protected:            &o.Protected,
			Relation:             &o.Relation,
			Subject:              &o.Subject,
			Type:                 &o.Type,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparsePolicy{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "action":
			sp.Action = &(o.Action)
		case "activeDuration":
			sp.ActiveDuration = &(o.ActiveDuration)
		case "activeSchedule":
			sp.ActiveSchedule = &(o.ActiveSchedule)
		case "allObjectTags":
			sp.AllObjectTags = &(o.AllObjectTags)
		case "allSubjectTags":
			sp.AllSubjectTags = &(o.AllSubjectTags)
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
		case "disabled":
			sp.Disabled = &(o.Disabled)
		case "expirationTime":
			sp.ExpirationTime = &(o.ExpirationTime)
		case "fallback":
			sp.Fallback = &(o.Fallback)
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
		case "object":
			sp.Object = &(o.Object)
		case "propagate":
			sp.Propagate = &(o.Propagate)
		case "propagationHidden":
			sp.PropagationHidden = &(o.PropagationHidden)
		case "protected":
			sp.Protected = &(o.Protected)
		case "relation":
			sp.Relation = &(o.Relation)
		case "subject":
			sp.Subject = &(o.Subject)
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

// Patch apply the non nil value of a *SparsePolicy to the object.
func (o *Policy) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePolicy)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Action != nil {
		o.Action = *so.Action
	}
	if so.ActiveDuration != nil {
		o.ActiveDuration = *so.ActiveDuration
	}
	if so.ActiveSchedule != nil {
		o.ActiveSchedule = *so.ActiveSchedule
	}
	if so.AllObjectTags != nil {
		o.AllObjectTags = *so.AllObjectTags
	}
	if so.AllSubjectTags != nil {
		o.AllSubjectTags = *so.AllSubjectTags
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
	if so.Disabled != nil {
		o.Disabled = *so.Disabled
	}
	if so.ExpirationTime != nil {
		o.ExpirationTime = *so.ExpirationTime
	}
	if so.Fallback != nil {
		o.Fallback = *so.Fallback
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
	if so.Object != nil {
		o.Object = *so.Object
	}
	if so.Propagate != nil {
		o.Propagate = *so.Propagate
	}
	if so.PropagationHidden != nil {
		o.PropagationHidden = *so.PropagationHidden
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Relation != nil {
		o.Relation = *so.Relation
	}
	if so.Subject != nil {
		o.Subject = *so.Subject
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

// DeepCopy returns a deep copy if the Policy.
func (o *Policy) DeepCopy() *Policy {

	if o == nil {
		return nil
	}

	out := &Policy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Policy.
func (o *Policy) DeepCopyInto(out *Policy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Policy: %s", err))
	}

	*out = *target.(*Policy)
}

// Validate valides the current information stored into the structure.
func (o *Policy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidatePattern("activeDuration", o.ActiveDuration, `^[0-9]+[smh]$`, `must be a valid duration like <n>s or <n>s or <n>h`, false); err != nil {
		errors = errors.Append(err)
	}

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

	if err := ValidateTagsExpression("object", o.Object); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsExpression("subject", o.Subject); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"APIAuthorization", "AuditProfileMapping", "EnforcerProfile", "File", "Hook", "HostServiceMapping", "Infrastructure", "NamespaceMapping", "Network", "NetworkRuleSet", "ProcessingUnit", "Quota", "Service", "ServiceDependency", "Syscall", "TokenScope", "SSHAuthorization", "UserAccess"}, false); err != nil {
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
func (*Policy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PolicyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PolicyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Policy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PolicyAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Policy) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "action":
		return o.Action
	case "activeDuration":
		return o.ActiveDuration
	case "activeSchedule":
		return o.ActiveSchedule
	case "allObjectTags":
		return o.AllObjectTags
	case "allSubjectTags":
		return o.AllSubjectTags
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
	case "disabled":
		return o.Disabled
	case "expirationTime":
		return o.ExpirationTime
	case "fallback":
		return o.Fallback
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
	case "object":
		return o.Object
	case "propagate":
		return o.Propagate
	case "propagationHidden":
		return o.PropagationHidden
	case "protected":
		return o.Protected
	case "relation":
		return o.Relation
	case "subject":
		return o.Subject
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

// PolicyAttributesMap represents the map of attribute for Policy.
var PolicyAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Action": {
		AllowedChoices: []string{},
		BSONFieldName:  "action",
		ConvertedName:  "Action",
		Description:    `Defines a set of actions that must be enforced when a dependency is met.`,
		Exposed:        true,
		Name:           "action",
		Stored:         true,
		SubType:        "map[string]map[string]interface{}",
		Type:           "external",
	},
	"ActiveDuration": {
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		BSONFieldName:  "activeduration",
		ConvertedName:  "ActiveDuration",
		Description: `Defines for how long the policy will be active according to the
` + "`" + `activeSchedule` + "`" + `.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeDuration",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"ActiveSchedule": {
		AllowedChoices: []string{},
		BSONFieldName:  "activeschedule",
		ConvertedName:  "ActiveSchedule",
		Description: `Defines when the policy should be active using the cron notation.
The policy will be active for the given ` + "`" + `activeDuration` + "`" + `.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeSchedule",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"AllObjectTags": {
		AllowedChoices: []string{},
		BSONFieldName:  "allobjecttags",
		ConvertedName:  "AllObjectTags",
		Description:    `This is a set of all object tags for matching in the DB.`,
		Name:           "allObjectTags",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"AllSubjectTags": {
		AllowedChoices: []string{},
		BSONFieldName:  "allsubjecttags",
		ConvertedName:  "AllSubjectTags",
		Description:    `This is a set of all subject tags for matching in the DB.`,
		Name:           "allSubjectTags",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"Disabled": {
		AllowedChoices: []string{},
		BSONFieldName:  "disabled",
		ConvertedName:  "Disabled",
		Description:    `Defines if the property is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"ExpirationTime": {
		AllowedChoices: []string{},
		BSONFieldName:  "expirationtime",
		ConvertedName:  "ExpirationTime",
		Description:    `If set the policy will be automatically deleted at the given time.`,
		Exposed:        true,
		Getter:         true,
		Name:           "expirationTime",
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Fallback": {
		AllowedChoices: []string{},
		BSONFieldName:  "fallback",
		ConvertedName:  "Fallback",
		Description: `Indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.`,
		Exposed:   true,
		Getter:    true,
		Name:      "fallback",
		Orderable: true,
		Setter:    true,
		Stored:    true,
		Type:      "boolean",
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
	"Object": {
		AllowedChoices: []string{},
		BSONFieldName:  "object",
		ConvertedName:  "Object",
		Description: `Represents set of entities that another entity depends on. As subjects,
objects are identified as logical operations on tags when a policy is defined.`,
		Exposed: true,
		Getter:  true,
		Name:    "object",
		Setter:  true,
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"Propagate": {
		AllowedChoices: []string{},
		BSONFieldName:  "propagate",
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
	"PropagationHidden": {
		AllowedChoices: []string{},
		BSONFieldName:  "propagationhidden",
		ConvertedName:  "PropagationHidden",
		Description: `If set to ` + "`" + `true` + "`" + ` while the policy is propagating, it won't be visible to
children
namespace, but still used for policy resolution.`,
		Exposed:   true,
		Getter:    true,
		Name:      "propagationHidden",
		Orderable: true,
		Setter:    true,
		Stored:    true,
		Type:      "boolean",
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
	"Relation": {
		AllowedChoices: []string{},
		BSONFieldName:  "relation",
		ConvertedName:  "Relation",
		Description: `Describes the required operation to be performed between subjects and
objects.`,
		Exposed: true,
		Name:    "relation",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"Subject": {
		AllowedChoices: []string{},
		BSONFieldName:  "subject",
		ConvertedName:  "Subject",
		Description: `Represents sets of entities that will have a dependency other entities.
Subjects are defined as logical operations on tags. Logical operations can
include ` + "`" + `AND` + "`" + ` and ` + "`" + `OR` + "`" + `.`,
		Exposed: true,
		Getter:  true,
		Name:    "subject",
		Setter:  true,
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"Type": {
		AllowedChoices: []string{"APIAuthorization", "AuditProfileMapping", "EnforcerProfile", "File", "Hook", "HostServiceMapping", "Infrastructure", "NamespaceMapping", "Network", "NetworkRuleSet", "ProcessingUnit", "Quota", "Service", "ServiceDependency", "Syscall", "TokenScope", "SSHAuthorization", "UserAccess"},
		BSONFieldName:  "type",
		ConvertedName:  "Type",
		CreationOnly:   true,
		Description:    `Type of the policy.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "type",
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

// PolicyLowerCaseAttributesMap represents the map of attribute for Policy.
var PolicyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"action": {
		AllowedChoices: []string{},
		BSONFieldName:  "action",
		ConvertedName:  "Action",
		Description:    `Defines a set of actions that must be enforced when a dependency is met.`,
		Exposed:        true,
		Name:           "action",
		Stored:         true,
		SubType:        "map[string]map[string]interface{}",
		Type:           "external",
	},
	"activeduration": {
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		BSONFieldName:  "activeduration",
		ConvertedName:  "ActiveDuration",
		Description: `Defines for how long the policy will be active according to the
` + "`" + `activeSchedule` + "`" + `.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeDuration",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"activeschedule": {
		AllowedChoices: []string{},
		BSONFieldName:  "activeschedule",
		ConvertedName:  "ActiveSchedule",
		Description: `Defines when the policy should be active using the cron notation.
The policy will be active for the given ` + "`" + `activeDuration` + "`" + `.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeSchedule",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"allobjecttags": {
		AllowedChoices: []string{},
		BSONFieldName:  "allobjecttags",
		ConvertedName:  "AllObjectTags",
		Description:    `This is a set of all object tags for matching in the DB.`,
		Name:           "allObjectTags",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"allsubjecttags": {
		AllowedChoices: []string{},
		BSONFieldName:  "allsubjecttags",
		ConvertedName:  "AllSubjectTags",
		Description:    `This is a set of all subject tags for matching in the DB.`,
		Name:           "allSubjectTags",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"disabled": {
		AllowedChoices: []string{},
		BSONFieldName:  "disabled",
		ConvertedName:  "Disabled",
		Description:    `Defines if the property is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"expirationtime": {
		AllowedChoices: []string{},
		BSONFieldName:  "expirationtime",
		ConvertedName:  "ExpirationTime",
		Description:    `If set the policy will be automatically deleted at the given time.`,
		Exposed:        true,
		Getter:         true,
		Name:           "expirationTime",
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"fallback": {
		AllowedChoices: []string{},
		BSONFieldName:  "fallback",
		ConvertedName:  "Fallback",
		Description: `Indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.`,
		Exposed:   true,
		Getter:    true,
		Name:      "fallback",
		Orderable: true,
		Setter:    true,
		Stored:    true,
		Type:      "boolean",
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
	"object": {
		AllowedChoices: []string{},
		BSONFieldName:  "object",
		ConvertedName:  "Object",
		Description: `Represents set of entities that another entity depends on. As subjects,
objects are identified as logical operations on tags when a policy is defined.`,
		Exposed: true,
		Getter:  true,
		Name:    "object",
		Setter:  true,
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"propagate": {
		AllowedChoices: []string{},
		BSONFieldName:  "propagate",
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
	"propagationhidden": {
		AllowedChoices: []string{},
		BSONFieldName:  "propagationhidden",
		ConvertedName:  "PropagationHidden",
		Description: `If set to ` + "`" + `true` + "`" + ` while the policy is propagating, it won't be visible to
children
namespace, but still used for policy resolution.`,
		Exposed:   true,
		Getter:    true,
		Name:      "propagationHidden",
		Orderable: true,
		Setter:    true,
		Stored:    true,
		Type:      "boolean",
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
	"relation": {
		AllowedChoices: []string{},
		BSONFieldName:  "relation",
		ConvertedName:  "Relation",
		Description: `Describes the required operation to be performed between subjects and
objects.`,
		Exposed: true,
		Name:    "relation",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"subject": {
		AllowedChoices: []string{},
		BSONFieldName:  "subject",
		ConvertedName:  "Subject",
		Description: `Represents sets of entities that will have a dependency other entities.
Subjects are defined as logical operations on tags. Logical operations can
include ` + "`" + `AND` + "`" + ` and ` + "`" + `OR` + "`" + `.`,
		Exposed: true,
		Getter:  true,
		Name:    "subject",
		Setter:  true,
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"type": {
		AllowedChoices: []string{"APIAuthorization", "AuditProfileMapping", "EnforcerProfile", "File", "Hook", "HostServiceMapping", "Infrastructure", "NamespaceMapping", "Network", "NetworkRuleSet", "ProcessingUnit", "Quota", "Service", "ServiceDependency", "Syscall", "TokenScope", "SSHAuthorization", "UserAccess"},
		BSONFieldName:  "type",
		ConvertedName:  "Type",
		CreationOnly:   true,
		Description:    `Type of the policy.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "type",
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

// SparsePoliciesList represents a list of SparsePolicies
type SparsePoliciesList []*SparsePolicy

// Identity returns the identity of the objects in the list.
func (o SparsePoliciesList) Identity() elemental.Identity {

	return PolicyIdentity
}

// Copy returns a pointer to a copy the SparsePoliciesList.
func (o SparsePoliciesList) Copy() elemental.Identifiables {

	copy := append(SparsePoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePoliciesList.
func (o SparsePoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePoliciesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparsePoliciesList converted to PoliciesList.
func (o SparsePoliciesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePoliciesList) Version() int {

	return 1
}

// SparsePolicy represents the sparse version of a policy.
type SparsePolicy struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Defines a set of actions that must be enforced when a dependency is met.
	Action *map[string]map[string]interface{} `json:"action,omitempty" msgpack:"action,omitempty" bson:"action,omitempty" mapstructure:"action,omitempty"`

	// Defines for how long the policy will be active according to the
	// `activeSchedule`.
	ActiveDuration *string `json:"activeDuration,omitempty" msgpack:"activeDuration,omitempty" bson:"activeduration,omitempty" mapstructure:"activeDuration,omitempty"`

	// Defines when the policy should be active using the cron notation.
	// The policy will be active for the given `activeDuration`.
	ActiveSchedule *string `json:"activeSchedule,omitempty" msgpack:"activeSchedule,omitempty" bson:"activeschedule,omitempty" mapstructure:"activeSchedule,omitempty"`

	// This is a set of all object tags for matching in the DB.
	AllObjectTags *[]string `json:"-" msgpack:"-" bson:"allobjecttags,omitempty" mapstructure:"-,omitempty"`

	// This is a set of all subject tags for matching in the DB.
	AllSubjectTags *[]string `json:"-" msgpack:"-" bson:"allsubjecttags,omitempty" mapstructure:"-,omitempty"`

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

	// Defines if the property is disabled.
	Disabled *bool `json:"disabled,omitempty" msgpack:"disabled,omitempty" bson:"disabled,omitempty" mapstructure:"disabled,omitempty"`

	// If set the policy will be automatically deleted at the given time.
	ExpirationTime *time.Time `json:"expirationTime,omitempty" msgpack:"expirationTime,omitempty" bson:"expirationtime,omitempty" mapstructure:"expirationTime,omitempty"`

	// Indicates that this is fallback policy. It will only be
	// applied if no other policies have been resolved. If the policy is also
	// propagated it will become a fallback for children namespaces.
	Fallback *bool `json:"fallback,omitempty" msgpack:"fallback,omitempty" bson:"fallback,omitempty" mapstructure:"fallback,omitempty"`

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

	// Represents set of entities that another entity depends on. As subjects,
	// objects are identified as logical operations on tags when a policy is defined.
	Object *[][]string `json:"object,omitempty" msgpack:"object,omitempty" bson:"object,omitempty" mapstructure:"object,omitempty"`

	// Propagates the policy to all of its children.
	Propagate *bool `json:"propagate,omitempty" msgpack:"propagate,omitempty" bson:"propagate,omitempty" mapstructure:"propagate,omitempty"`

	// If set to `true` while the policy is propagating, it won't be visible to
	// children
	// namespace, but still used for policy resolution.
	PropagationHidden *bool `json:"propagationHidden,omitempty" msgpack:"propagationHidden,omitempty" bson:"propagationhidden,omitempty" mapstructure:"propagationHidden,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Describes the required operation to be performed between subjects and
	// objects.
	Relation *[]string `json:"relation,omitempty" msgpack:"relation,omitempty" bson:"relation,omitempty" mapstructure:"relation,omitempty"`

	// Represents sets of entities that will have a dependency other entities.
	// Subjects are defined as logical operations on tags. Logical operations can
	// include `AND` and `OR`.
	Subject *[][]string `json:"subject,omitempty" msgpack:"subject,omitempty" bson:"subject,omitempty" mapstructure:"subject,omitempty"`

	// Type of the policy.
	Type *PolicyTypeValue `json:"type,omitempty" msgpack:"type,omitempty" bson:"type,omitempty" mapstructure:"type,omitempty"`

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

// NewSparsePolicy returns a new  SparsePolicy.
func NewSparsePolicy() *SparsePolicy {
	return &SparsePolicy{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePolicy) Identity() elemental.Identity {

	return PolicyIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePolicy) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePolicy) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePolicy) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePolicy{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Action != nil {
		s.Action = o.Action
	}
	if o.ActiveDuration != nil {
		s.ActiveDuration = o.ActiveDuration
	}
	if o.ActiveSchedule != nil {
		s.ActiveSchedule = o.ActiveSchedule
	}
	if o.AllObjectTags != nil {
		s.AllObjectTags = o.AllObjectTags
	}
	if o.AllSubjectTags != nil {
		s.AllSubjectTags = o.AllSubjectTags
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
	if o.Disabled != nil {
		s.Disabled = o.Disabled
	}
	if o.ExpirationTime != nil {
		s.ExpirationTime = o.ExpirationTime
	}
	if o.Fallback != nil {
		s.Fallback = o.Fallback
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
	if o.Object != nil {
		s.Object = o.Object
	}
	if o.Propagate != nil {
		s.Propagate = o.Propagate
	}
	if o.PropagationHidden != nil {
		s.PropagationHidden = o.PropagationHidden
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.Relation != nil {
		s.Relation = o.Relation
	}
	if o.Subject != nil {
		s.Subject = o.Subject
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
func (o *SparsePolicy) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePolicy{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Action != nil {
		o.Action = s.Action
	}
	if s.ActiveDuration != nil {
		o.ActiveDuration = s.ActiveDuration
	}
	if s.ActiveSchedule != nil {
		o.ActiveSchedule = s.ActiveSchedule
	}
	if s.AllObjectTags != nil {
		o.AllObjectTags = s.AllObjectTags
	}
	if s.AllSubjectTags != nil {
		o.AllSubjectTags = s.AllSubjectTags
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
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Disabled != nil {
		o.Disabled = s.Disabled
	}
	if s.ExpirationTime != nil {
		o.ExpirationTime = s.ExpirationTime
	}
	if s.Fallback != nil {
		o.Fallback = s.Fallback
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
	if s.Object != nil {
		o.Object = s.Object
	}
	if s.Propagate != nil {
		o.Propagate = s.Propagate
	}
	if s.PropagationHidden != nil {
		o.PropagationHidden = s.PropagationHidden
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.Relation != nil {
		o.Relation = s.Relation
	}
	if s.Subject != nil {
		o.Subject = s.Subject
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
func (o *SparsePolicy) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePolicy) ToPlain() elemental.PlainIdentifiable {

	out := NewPolicy()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Action != nil {
		out.Action = *o.Action
	}
	if o.ActiveDuration != nil {
		out.ActiveDuration = *o.ActiveDuration
	}
	if o.ActiveSchedule != nil {
		out.ActiveSchedule = *o.ActiveSchedule
	}
	if o.AllObjectTags != nil {
		out.AllObjectTags = *o.AllObjectTags
	}
	if o.AllSubjectTags != nil {
		out.AllSubjectTags = *o.AllSubjectTags
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
	if o.Disabled != nil {
		out.Disabled = *o.Disabled
	}
	if o.ExpirationTime != nil {
		out.ExpirationTime = *o.ExpirationTime
	}
	if o.Fallback != nil {
		out.Fallback = *o.Fallback
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
	if o.Object != nil {
		out.Object = *o.Object
	}
	if o.Propagate != nil {
		out.Propagate = *o.Propagate
	}
	if o.PropagationHidden != nil {
		out.PropagationHidden = *o.PropagationHidden
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Relation != nil {
		out.Relation = *o.Relation
	}
	if o.Subject != nil {
		out.Subject = *o.Subject
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

// GetActiveDuration returns the ActiveDuration of the receiver.
func (o *SparsePolicy) GetActiveDuration() (out string) {

	if o.ActiveDuration == nil {
		return
	}

	return *o.ActiveDuration
}

// SetActiveDuration sets the property ActiveDuration of the receiver using the address of the given value.
func (o *SparsePolicy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = &activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *SparsePolicy) GetActiveSchedule() (out string) {

	if o.ActiveSchedule == nil {
		return
	}

	return *o.ActiveSchedule
}

// SetActiveSchedule sets the property ActiveSchedule of the receiver using the address of the given value.
func (o *SparsePolicy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = &activeSchedule
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparsePolicy) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparsePolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparsePolicy) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparsePolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparsePolicy) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparsePolicy) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparsePolicy) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparsePolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparsePolicy) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparsePolicy) SetDescription(description string) {

	o.Description = &description
}

// GetDisabled returns the Disabled of the receiver.
func (o *SparsePolicy) GetDisabled() (out bool) {

	if o.Disabled == nil {
		return
	}

	return *o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the address of the given value.
func (o *SparsePolicy) SetDisabled(disabled bool) {

	o.Disabled = &disabled
}

// GetExpirationTime returns the ExpirationTime of the receiver.
func (o *SparsePolicy) GetExpirationTime() (out time.Time) {

	if o.ExpirationTime == nil {
		return
	}

	return *o.ExpirationTime
}

// SetExpirationTime sets the property ExpirationTime of the receiver using the address of the given value.
func (o *SparsePolicy) SetExpirationTime(expirationTime time.Time) {

	o.ExpirationTime = &expirationTime
}

// GetFallback returns the Fallback of the receiver.
func (o *SparsePolicy) GetFallback() (out bool) {

	if o.Fallback == nil {
		return
	}

	return *o.Fallback
}

// SetFallback sets the property Fallback of the receiver using the address of the given value.
func (o *SparsePolicy) SetFallback(fallback bool) {

	o.Fallback = &fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparsePolicy) GetMetadata() (out []string) {

	if o.Metadata == nil {
		return
	}

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparsePolicy) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparsePolicy) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparsePolicy) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparsePolicy) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparsePolicy) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparsePolicy) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparsePolicy) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparsePolicy) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparsePolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetObject returns the Object of the receiver.
func (o *SparsePolicy) GetObject() (out [][]string) {

	if o.Object == nil {
		return
	}

	return *o.Object
}

// SetObject sets the property Object of the receiver using the address of the given value.
func (o *SparsePolicy) SetObject(object [][]string) {

	o.Object = &object
}

// GetPropagate returns the Propagate of the receiver.
func (o *SparsePolicy) GetPropagate() (out bool) {

	if o.Propagate == nil {
		return
	}

	return *o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the address of the given value.
func (o *SparsePolicy) SetPropagate(propagate bool) {

	o.Propagate = &propagate
}

// GetPropagationHidden returns the PropagationHidden of the receiver.
func (o *SparsePolicy) GetPropagationHidden() (out bool) {

	if o.PropagationHidden == nil {
		return
	}

	return *o.PropagationHidden
}

// SetPropagationHidden sets the property PropagationHidden of the receiver using the address of the given value.
func (o *SparsePolicy) SetPropagationHidden(propagationHidden bool) {

	o.PropagationHidden = &propagationHidden
}

// GetProtected returns the Protected of the receiver.
func (o *SparsePolicy) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparsePolicy) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetSubject returns the Subject of the receiver.
func (o *SparsePolicy) GetSubject() (out [][]string) {

	if o.Subject == nil {
		return
	}

	return *o.Subject
}

// SetSubject sets the property Subject of the receiver using the address of the given value.
func (o *SparsePolicy) SetSubject(subject [][]string) {

	o.Subject = &subject
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparsePolicy) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparsePolicy) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparsePolicy) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparsePolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparsePolicy) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparsePolicy) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparsePolicy) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparsePolicy) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparsePolicy.
func (o *SparsePolicy) DeepCopy() *SparsePolicy {

	if o == nil {
		return nil
	}

	out := &SparsePolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePolicy.
func (o *SparsePolicy) DeepCopyInto(out *SparsePolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePolicy: %s", err))
	}

	*out = *target.(*SparsePolicy)
}

type mongoAttributesPolicy struct {
	ID                   bson.ObjectId                     `bson:"_id,omitempty"`
	Action               map[string]map[string]interface{} `bson:"action"`
	ActiveDuration       string                            `bson:"activeduration"`
	ActiveSchedule       string                            `bson:"activeschedule"`
	AllObjectTags        []string                          `bson:"allobjecttags"`
	AllSubjectTags       []string                          `bson:"allsubjecttags"`
	Annotations          map[string][]string               `bson:"annotations"`
	AssociatedTags       []string                          `bson:"associatedtags"`
	CreateIdempotencyKey string                            `bson:"createidempotencykey"`
	CreateTime           time.Time                         `bson:"createtime"`
	Description          string                            `bson:"description"`
	Disabled             bool                              `bson:"disabled"`
	ExpirationTime       time.Time                         `bson:"expirationtime"`
	Fallback             bool                              `bson:"fallback"`
	Metadata             []string                          `bson:"metadata"`
	MigrationsLog        map[string]string                 `bson:"migrationslog,omitempty"`
	Name                 string                            `bson:"name"`
	Namespace            string                            `bson:"namespace"`
	NormalizedTags       []string                          `bson:"normalizedtags"`
	Object               [][]string                        `bson:"object"`
	Propagate            bool                              `bson:"propagate"`
	PropagationHidden    bool                              `bson:"propagationhidden"`
	Protected            bool                              `bson:"protected"`
	Relation             []string                          `bson:"relation"`
	Subject              [][]string                        `bson:"subject"`
	Type                 PolicyTypeValue                   `bson:"type"`
	UpdateIdempotencyKey string                            `bson:"updateidempotencykey"`
	UpdateTime           time.Time                         `bson:"updatetime"`
	ZHash                int                               `bson:"zhash"`
	Zone                 int                               `bson:"zone"`
}
type mongoAttributesSparsePolicy struct {
	ID                   bson.ObjectId                      `bson:"_id,omitempty"`
	Action               *map[string]map[string]interface{} `bson:"action,omitempty"`
	ActiveDuration       *string                            `bson:"activeduration,omitempty"`
	ActiveSchedule       *string                            `bson:"activeschedule,omitempty"`
	AllObjectTags        *[]string                          `bson:"allobjecttags,omitempty"`
	AllSubjectTags       *[]string                          `bson:"allsubjecttags,omitempty"`
	Annotations          *map[string][]string               `bson:"annotations,omitempty"`
	AssociatedTags       *[]string                          `bson:"associatedtags,omitempty"`
	CreateIdempotencyKey *string                            `bson:"createidempotencykey,omitempty"`
	CreateTime           *time.Time                         `bson:"createtime,omitempty"`
	Description          *string                            `bson:"description,omitempty"`
	Disabled             *bool                              `bson:"disabled,omitempty"`
	ExpirationTime       *time.Time                         `bson:"expirationtime,omitempty"`
	Fallback             *bool                              `bson:"fallback,omitempty"`
	Metadata             *[]string                          `bson:"metadata,omitempty"`
	MigrationsLog        *map[string]string                 `bson:"migrationslog,omitempty"`
	Name                 *string                            `bson:"name,omitempty"`
	Namespace            *string                            `bson:"namespace,omitempty"`
	NormalizedTags       *[]string                          `bson:"normalizedtags,omitempty"`
	Object               *[][]string                        `bson:"object,omitempty"`
	Propagate            *bool                              `bson:"propagate,omitempty"`
	PropagationHidden    *bool                              `bson:"propagationhidden,omitempty"`
	Protected            *bool                              `bson:"protected,omitempty"`
	Relation             *[]string                          `bson:"relation,omitempty"`
	Subject              *[][]string                        `bson:"subject,omitempty"`
	Type                 *PolicyTypeValue                   `bson:"type,omitempty"`
	UpdateIdempotencyKey *string                            `bson:"updateidempotencykey,omitempty"`
	UpdateTime           *time.Time                         `bson:"updatetime,omitempty"`
	ZHash                *int                               `bson:"zhash,omitempty"`
	Zone                 *int                               `bson:"zone,omitempty"`
}
