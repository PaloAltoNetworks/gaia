package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// InfrastructurePolicyActionValue represents the possible values for attribute "action".
type InfrastructurePolicyActionValue string

const (
	// InfrastructurePolicyActionAllow represents the value Allow.
	InfrastructurePolicyActionAllow InfrastructurePolicyActionValue = "Allow"

	// InfrastructurePolicyActionContinue represents the value Continue.
	InfrastructurePolicyActionContinue InfrastructurePolicyActionValue = "Continue"

	// InfrastructurePolicyActionReject represents the value Reject.
	InfrastructurePolicyActionReject InfrastructurePolicyActionValue = "Reject"
)

// InfrastructurePolicyApplyPolicyModeValue represents the possible values for attribute "applyPolicyMode".
type InfrastructurePolicyApplyPolicyModeValue string

const (
	// InfrastructurePolicyApplyPolicyModeBidirectional represents the value Bidirectional.
	InfrastructurePolicyApplyPolicyModeBidirectional InfrastructurePolicyApplyPolicyModeValue = "Bidirectional"

	// InfrastructurePolicyApplyPolicyModeIncomingTraffic represents the value IncomingTraffic.
	InfrastructurePolicyApplyPolicyModeIncomingTraffic InfrastructurePolicyApplyPolicyModeValue = "IncomingTraffic"

	// InfrastructurePolicyApplyPolicyModeOutgoingTraffic represents the value OutgoingTraffic.
	InfrastructurePolicyApplyPolicyModeOutgoingTraffic InfrastructurePolicyApplyPolicyModeValue = "OutgoingTraffic"
)

// InfrastructurePolicyObservedTrafficActionValue represents the possible values for attribute "observedTrafficAction".
type InfrastructurePolicyObservedTrafficActionValue string

const (
	// InfrastructurePolicyObservedTrafficActionApply represents the value Apply.
	InfrastructurePolicyObservedTrafficActionApply InfrastructurePolicyObservedTrafficActionValue = "Apply"

	// InfrastructurePolicyObservedTrafficActionContinue represents the value Continue.
	InfrastructurePolicyObservedTrafficActionContinue InfrastructurePolicyObservedTrafficActionValue = "Continue"
)

// InfrastructurePolicyIdentity represents the Identity of the object.
var InfrastructurePolicyIdentity = elemental.Identity{
	Name:     "infrastructurepolicy",
	Category: "infrastructurepolicies",
	Package:  "squall",
	Private:  false,
}

// InfrastructurePoliciesList represents a list of InfrastructurePolicies
type InfrastructurePoliciesList []*InfrastructurePolicy

// Identity returns the identity of the objects in the list.
func (o InfrastructurePoliciesList) Identity() elemental.Identity {

	return InfrastructurePolicyIdentity
}

// Copy returns a pointer to a copy the InfrastructurePoliciesList.
func (o InfrastructurePoliciesList) Copy() elemental.Identifiables {

	copy := append(InfrastructurePoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the InfrastructurePoliciesList.
func (o InfrastructurePoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(InfrastructurePoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*InfrastructurePolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o InfrastructurePoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o InfrastructurePoliciesList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToSparse returns the InfrastructurePoliciesList converted to SparseInfrastructurePoliciesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o InfrastructurePoliciesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseInfrastructurePoliciesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseInfrastructurePolicy)
	}

	return out
}

// Version returns the version of the content.
func (o InfrastructurePoliciesList) Version() int {

	return 1
}

// InfrastructurePolicy represents the model of a infrastructurepolicy
type InfrastructurePolicy struct {
	// ID is the identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Action defines the action to apply to a flow.
	Action InfrastructurePolicyActionValue `json:"action" msgpack:"action" bson:"-" mapstructure:"action,omitempty"`

	// ActiveDuration defines for how long the policy will be active according to the
	// activeSchedule.
	ActiveDuration string `json:"activeDuration" msgpack:"activeDuration" bson:"activeduration" mapstructure:"activeDuration,omitempty"`

	// ActiveSchedule defines when the policy should be active using the cron notation.
	// The policy will be active for the given activeDuration.
	ActiveSchedule string `json:"activeSchedule" msgpack:"activeSchedule" bson:"activeschedule" mapstructure:"activeSchedule,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// applyPolicyMode determines if the policy has to be applied to the
	// outgoing traffic of a PU or the incoming traffic of a PU or in both directions.
	// Default is both directions.
	ApplyPolicyMode InfrastructurePolicyApplyPolicyModeValue `json:"applyPolicyMode" msgpack:"applyPolicyMode" bson:"-" mapstructure:"applyPolicyMode,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Disabled defines if the propert is disabled.
	Disabled bool `json:"disabled" msgpack:"disabled" bson:"disabled" mapstructure:"disabled,omitempty"`

	// EncryptionEnabled defines if the flow has to be encrypted.
	EncryptionEnabled bool `json:"encryptionEnabled" msgpack:"encryptionEnabled" bson:"-" mapstructure:"encryptionEnabled,omitempty"`

	// If set the policy will be auto deleted after the given time.
	ExpirationTime time.Time `json:"expirationTime" msgpack:"expirationTime" bson:"expirationtime" mapstructure:"expirationTime,omitempty"`

	// Fallback indicates that this is fallback policy. It will only be
	// applied if no other policies have been resolved. If the policy is also
	// propagated it will become a fallback for children namespaces.
	Fallback bool `json:"fallback" msgpack:"fallback" bson:"fallback" mapstructure:"fallback,omitempty"`

	// LogsEnabled defines if the flow has to be logged.
	LogsEnabled bool `json:"logsEnabled" msgpack:"logsEnabled" bson:"-" mapstructure:"logsEnabled,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Setting this to true will invert the object to find what is not matching.
	NegateObject bool `json:"negateObject" msgpack:"negateObject" bson:"negateobject" mapstructure:"negateObject,omitempty"`

	// Setting this to true will invert the subject to find what is not matching.
	NegateSubject bool `json:"negateSubject" msgpack:"negateSubject" bson:"negatesubject" mapstructure:"negateSubject,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Object of the policy.
	Object [][]string `json:"object" msgpack:"object" bson:"-" mapstructure:"object,omitempty"`

	// If set to true, the flow will be in observation mode.
	ObservationEnabled bool `json:"observationEnabled" msgpack:"observationEnabled" bson:"-" mapstructure:"observationEnabled,omitempty"`

	// If observationEnabled is set to true, this will be the final action taken on the
	// packets.
	ObservedTrafficAction InfrastructurePolicyObservedTrafficActionValue `json:"observedTrafficAction" msgpack:"observedTrafficAction" bson:"-" mapstructure:"observedTrafficAction,omitempty"`

	// Propagate will propagate the policy to all of its children.
	Propagate bool `json:"propagate" msgpack:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Subject of the policy.
	Subject [][]string `json:"subject" msgpack:"subject" bson:"-" mapstructure:"subject,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewInfrastructurePolicy returns a new *InfrastructurePolicy
func NewInfrastructurePolicy() *InfrastructurePolicy {

	return &InfrastructurePolicy{
		ModelVersion:          1,
		Action:                InfrastructurePolicyActionAllow,
		AssociatedTags:        []string{},
		Annotations:           map[string][]string{},
		ApplyPolicyMode:       InfrastructurePolicyApplyPolicyModeBidirectional,
		Metadata:              []string{},
		ObservedTrafficAction: InfrastructurePolicyObservedTrafficActionContinue,
		NormalizedTags:        []string{},
		Object:                [][]string{},
		Subject:               [][]string{},
	}
}

// Identity returns the Identity of the object.
func (o *InfrastructurePolicy) Identity() elemental.Identity {

	return InfrastructurePolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *InfrastructurePolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *InfrastructurePolicy) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *InfrastructurePolicy) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *InfrastructurePolicy) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// Doc returns the documentation for the object
func (o *InfrastructurePolicy) Doc() string {

	return `Allows defining infrastructure policies to allow or prevent processing units
identified by their tags to talk to other processing units or external services
(also identified by their tags).`
}

func (o *InfrastructurePolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetActiveDuration returns the ActiveDuration of the receiver.
func (o *InfrastructurePolicy) GetActiveDuration() string {

	return o.ActiveDuration
}

// SetActiveDuration sets the property ActiveDuration of the receiver using the given value.
func (o *InfrastructurePolicy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *InfrastructurePolicy) GetActiveSchedule() string {

	return o.ActiveSchedule
}

// SetActiveSchedule sets the property ActiveSchedule of the receiver using the given value.
func (o *InfrastructurePolicy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = activeSchedule
}

// GetAnnotations returns the Annotations of the receiver.
func (o *InfrastructurePolicy) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *InfrastructurePolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *InfrastructurePolicy) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *InfrastructurePolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *InfrastructurePolicy) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *InfrastructurePolicy) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *InfrastructurePolicy) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *InfrastructurePolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *InfrastructurePolicy) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *InfrastructurePolicy) SetDescription(description string) {

	o.Description = description
}

// GetDisabled returns the Disabled of the receiver.
func (o *InfrastructurePolicy) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the given value.
func (o *InfrastructurePolicy) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetExpirationTime returns the ExpirationTime of the receiver.
func (o *InfrastructurePolicy) GetExpirationTime() time.Time {

	return o.ExpirationTime
}

// SetExpirationTime sets the property ExpirationTime of the receiver using the given value.
func (o *InfrastructurePolicy) SetExpirationTime(expirationTime time.Time) {

	o.ExpirationTime = expirationTime
}

// GetFallback returns the Fallback of the receiver.
func (o *InfrastructurePolicy) GetFallback() bool {

	return o.Fallback
}

// SetFallback sets the property Fallback of the receiver using the given value.
func (o *InfrastructurePolicy) SetFallback(fallback bool) {

	o.Fallback = fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *InfrastructurePolicy) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *InfrastructurePolicy) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *InfrastructurePolicy) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *InfrastructurePolicy) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *InfrastructurePolicy) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *InfrastructurePolicy) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNegateObject returns the NegateObject of the receiver.
func (o *InfrastructurePolicy) GetNegateObject() bool {

	return o.NegateObject
}

// SetNegateObject sets the property NegateObject of the receiver using the given value.
func (o *InfrastructurePolicy) SetNegateObject(negateObject bool) {

	o.NegateObject = negateObject
}

// GetNegateSubject returns the NegateSubject of the receiver.
func (o *InfrastructurePolicy) GetNegateSubject() bool {

	return o.NegateSubject
}

// SetNegateSubject sets the property NegateSubject of the receiver using the given value.
func (o *InfrastructurePolicy) SetNegateSubject(negateSubject bool) {

	o.NegateSubject = negateSubject
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *InfrastructurePolicy) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *InfrastructurePolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *InfrastructurePolicy) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the given value.
func (o *InfrastructurePolicy) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// GetProtected returns the Protected of the receiver.
func (o *InfrastructurePolicy) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *InfrastructurePolicy) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *InfrastructurePolicy) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *InfrastructurePolicy) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *InfrastructurePolicy) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *InfrastructurePolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *InfrastructurePolicy) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *InfrastructurePolicy) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *InfrastructurePolicy) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *InfrastructurePolicy) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *InfrastructurePolicy) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseInfrastructurePolicy{
			ID:                    &o.ID,
			Action:                &o.Action,
			ActiveDuration:        &o.ActiveDuration,
			ActiveSchedule:        &o.ActiveSchedule,
			Annotations:           &o.Annotations,
			ApplyPolicyMode:       &o.ApplyPolicyMode,
			AssociatedTags:        &o.AssociatedTags,
			CreateIdempotencyKey:  &o.CreateIdempotencyKey,
			CreateTime:            &o.CreateTime,
			Description:           &o.Description,
			Disabled:              &o.Disabled,
			EncryptionEnabled:     &o.EncryptionEnabled,
			ExpirationTime:        &o.ExpirationTime,
			Fallback:              &o.Fallback,
			LogsEnabled:           &o.LogsEnabled,
			Metadata:              &o.Metadata,
			Name:                  &o.Name,
			Namespace:             &o.Namespace,
			NegateObject:          &o.NegateObject,
			NegateSubject:         &o.NegateSubject,
			NormalizedTags:        &o.NormalizedTags,
			Object:                &o.Object,
			ObservationEnabled:    &o.ObservationEnabled,
			ObservedTrafficAction: &o.ObservedTrafficAction,
			Propagate:             &o.Propagate,
			Protected:             &o.Protected,
			Subject:               &o.Subject,
			UpdateIdempotencyKey:  &o.UpdateIdempotencyKey,
			UpdateTime:            &o.UpdateTime,
			ZHash:                 &o.ZHash,
			Zone:                  &o.Zone,
		}
	}

	sp := &SparseInfrastructurePolicy{}
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
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "applyPolicyMode":
			sp.ApplyPolicyMode = &(o.ApplyPolicyMode)
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
		case "encryptionEnabled":
			sp.EncryptionEnabled = &(o.EncryptionEnabled)
		case "expirationTime":
			sp.ExpirationTime = &(o.ExpirationTime)
		case "fallback":
			sp.Fallback = &(o.Fallback)
		case "logsEnabled":
			sp.LogsEnabled = &(o.LogsEnabled)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "negateObject":
			sp.NegateObject = &(o.NegateObject)
		case "negateSubject":
			sp.NegateSubject = &(o.NegateSubject)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "object":
			sp.Object = &(o.Object)
		case "observationEnabled":
			sp.ObservationEnabled = &(o.ObservationEnabled)
		case "observedTrafficAction":
			sp.ObservedTrafficAction = &(o.ObservedTrafficAction)
		case "propagate":
			sp.Propagate = &(o.Propagate)
		case "protected":
			sp.Protected = &(o.Protected)
		case "subject":
			sp.Subject = &(o.Subject)
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

// Patch apply the non nil value of a *SparseInfrastructurePolicy to the object.
func (o *InfrastructurePolicy) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseInfrastructurePolicy)
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
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.ApplyPolicyMode != nil {
		o.ApplyPolicyMode = *so.ApplyPolicyMode
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
	if so.EncryptionEnabled != nil {
		o.EncryptionEnabled = *so.EncryptionEnabled
	}
	if so.ExpirationTime != nil {
		o.ExpirationTime = *so.ExpirationTime
	}
	if so.Fallback != nil {
		o.Fallback = *so.Fallback
	}
	if so.LogsEnabled != nil {
		o.LogsEnabled = *so.LogsEnabled
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NegateObject != nil {
		o.NegateObject = *so.NegateObject
	}
	if so.NegateSubject != nil {
		o.NegateSubject = *so.NegateSubject
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Object != nil {
		o.Object = *so.Object
	}
	if so.ObservationEnabled != nil {
		o.ObservationEnabled = *so.ObservationEnabled
	}
	if so.ObservedTrafficAction != nil {
		o.ObservedTrafficAction = *so.ObservedTrafficAction
	}
	if so.Propagate != nil {
		o.Propagate = *so.Propagate
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Subject != nil {
		o.Subject = *so.Subject
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

// DeepCopy returns a deep copy if the InfrastructurePolicy.
func (o *InfrastructurePolicy) DeepCopy() *InfrastructurePolicy {

	if o == nil {
		return nil
	}

	out := &InfrastructurePolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *InfrastructurePolicy.
func (o *InfrastructurePolicy) DeepCopyInto(out *InfrastructurePolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy InfrastructurePolicy: %s", err))
	}

	*out = *target.(*InfrastructurePolicy)
}

// Validate valides the current information stored into the structure.
func (o *InfrastructurePolicy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("action", string(o.Action), []string{"Allow", "Reject", "Continue"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidatePattern("activeDuration", o.ActiveDuration, `^[0-9]+[smh]$`, `must be a valid duration like <n>s or <n>s or <n>h`, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("applyPolicyMode", string(o.ApplyPolicyMode), []string{"OutgoingTraffic", "IncomingTraffic", "Bidirectional"}, false); err != nil {
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

	if err := elemental.ValidateStringInList("observedTrafficAction", string(o.ObservedTrafficAction), []string{"Apply", "Continue"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsExpression("subject", o.Subject); err != nil {
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
func (*InfrastructurePolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := InfrastructurePolicyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return InfrastructurePolicyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*InfrastructurePolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return InfrastructurePolicyAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *InfrastructurePolicy) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "action":
		return o.Action
	case "activeDuration":
		return o.ActiveDuration
	case "activeSchedule":
		return o.ActiveSchedule
	case "annotations":
		return o.Annotations
	case "applyPolicyMode":
		return o.ApplyPolicyMode
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
	case "encryptionEnabled":
		return o.EncryptionEnabled
	case "expirationTime":
		return o.ExpirationTime
	case "fallback":
		return o.Fallback
	case "logsEnabled":
		return o.LogsEnabled
	case "metadata":
		return o.Metadata
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "negateObject":
		return o.NegateObject
	case "negateSubject":
		return o.NegateSubject
	case "normalizedTags":
		return o.NormalizedTags
	case "object":
		return o.Object
	case "observationEnabled":
		return o.ObservationEnabled
	case "observedTrafficAction":
		return o.ObservedTrafficAction
	case "propagate":
		return o.Propagate
	case "protected":
		return o.Protected
	case "subject":
		return o.Subject
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

// InfrastructurePolicyAttributesMap represents the map of attribute for InfrastructurePolicy.
var InfrastructurePolicyAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"Action": elemental.AttributeSpecification{
		AllowedChoices: []string{"Allow", "Reject", "Continue"},
		ConvertedName:  "Action",
		DefaultValue:   InfrastructurePolicyActionAllow,
		Description:    `Action defines the action to apply to a flow.`,
		Exposed:        true,
		Name:           "action",
		Orderable:      true,
		Type:           "enum",
	},
	"ActiveDuration": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "ActiveDuration",
		Description: `ActiveDuration defines for how long the policy will be active according to the
activeSchedule.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeDuration",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"ActiveSchedule": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ActiveSchedule",
		Description: `ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeSchedule",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"Annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"ApplyPolicyMode": elemental.AttributeSpecification{
		AllowedChoices: []string{"OutgoingTraffic", "IncomingTraffic", "Bidirectional"},
		ConvertedName:  "ApplyPolicyMode",
		DefaultValue:   InfrastructurePolicyApplyPolicyModeBidirectional,
		Description: `applyPolicyMode determines if the policy has to be applied to the
outgoing traffic of a PU or the incoming traffic of a PU or in both directions.
Default is both directions.`,
		Exposed:   true,
		Name:      "applyPolicyMode",
		Orderable: true,
		Type:      "enum",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"CreateIdempotencyKey": elemental.AttributeSpecification{
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
	"CreateTime": elemental.AttributeSpecification{
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
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"EncryptionEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EncryptionEnabled",
		Description:    `EncryptionEnabled defines if the flow has to be encrypted.`,
		Exposed:        true,
		Name:           "encryptionEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"ExpirationTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationTime",
		Description:    `If set the policy will be auto deleted after the given time.`,
		Exposed:        true,
		Getter:         true,
		Name:           "expirationTime",
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Fallback": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fallback",
		Description: `Fallback indicates that this is fallback policy. It will only be
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
	"LogsEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LogsEnabled",
		Description:    `LogsEnabled defines if the flow has to be logged.`,
		Exposed:        true,
		Name:           "logsEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Metadata contains tags that can only be set during creation. They must all start
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
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
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
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		DefaultOrder:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NegateObject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NegateObject",
		Description:    `Setting this to true will invert the object to find what is not matching.`,
		Exposed:        true,
		Getter:         true,
		Name:           "negateObject",
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"NegateSubject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NegateSubject",
		Description:    `Setting this to true will invert the subject to find what is not matching.`,
		Exposed:        true,
		Getter:         true,
		Name:           "negateSubject",
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
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
	"Object": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Object",
		Description:    `Object of the policy.`,
		Exposed:        true,
		Name:           "object",
		Orderable:      true,
		SubType:        "[][]string",
		Type:           "external",
	},
	"ObservationEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObservationEnabled",
		Description:    `If set to true, the flow will be in observation mode.`,
		Exposed:        true,
		Name:           "observationEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"ObservedTrafficAction": elemental.AttributeSpecification{
		AllowedChoices: []string{"Apply", "Continue"},
		ConvertedName:  "ObservedTrafficAction",
		DefaultValue:   InfrastructurePolicyObservedTrafficActionContinue,
		Description: `If observationEnabled is set to true, this will be the final action taken on the
packets.`,
		Exposed:   true,
		Name:      "observedTrafficAction",
		Orderable: true,
		Type:      "enum",
	},
	"Propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagate will propagate the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description:    `Subject of the policy.`,
		Exposed:        true,
		Name:           "subject",
		Orderable:      true,
		SubType:        "[][]string",
		Type:           "external",
	},
	"UpdateIdempotencyKey": elemental.AttributeSpecification{
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
	"UpdateTime": elemental.AttributeSpecification{
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
	"ZHash": elemental.AttributeSpecification{
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
	"Zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
}

// InfrastructurePolicyLowerCaseAttributesMap represents the map of attribute for InfrastructurePolicy.
var InfrastructurePolicyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"action": elemental.AttributeSpecification{
		AllowedChoices: []string{"Allow", "Reject", "Continue"},
		ConvertedName:  "Action",
		DefaultValue:   InfrastructurePolicyActionAllow,
		Description:    `Action defines the action to apply to a flow.`,
		Exposed:        true,
		Name:           "action",
		Orderable:      true,
		Type:           "enum",
	},
	"activeduration": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "ActiveDuration",
		Description: `ActiveDuration defines for how long the policy will be active according to the
activeSchedule.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeDuration",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"activeschedule": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ActiveSchedule",
		Description: `ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeSchedule",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"applypolicymode": elemental.AttributeSpecification{
		AllowedChoices: []string{"OutgoingTraffic", "IncomingTraffic", "Bidirectional"},
		ConvertedName:  "ApplyPolicyMode",
		DefaultValue:   InfrastructurePolicyApplyPolicyModeBidirectional,
		Description: `applyPolicyMode determines if the policy has to be applied to the
outgoing traffic of a PU or the incoming traffic of a PU or in both directions.
Default is both directions.`,
		Exposed:   true,
		Name:      "applyPolicyMode",
		Orderable: true,
		Type:      "enum",
	},
	"associatedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"createidempotencykey": elemental.AttributeSpecification{
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
	"createtime": elemental.AttributeSpecification{
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
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"encryptionenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EncryptionEnabled",
		Description:    `EncryptionEnabled defines if the flow has to be encrypted.`,
		Exposed:        true,
		Name:           "encryptionEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"expirationtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationTime",
		Description:    `If set the policy will be auto deleted after the given time.`,
		Exposed:        true,
		Getter:         true,
		Name:           "expirationTime",
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"fallback": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fallback",
		Description: `Fallback indicates that this is fallback policy. It will only be
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
	"logsenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LogsEnabled",
		Description:    `LogsEnabled defines if the flow has to be logged.`,
		Exposed:        true,
		Name:           "logsEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Metadata contains tags that can only be set during creation. They must all start
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
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
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
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		DefaultOrder:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"negateobject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NegateObject",
		Description:    `Setting this to true will invert the object to find what is not matching.`,
		Exposed:        true,
		Getter:         true,
		Name:           "negateObject",
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"negatesubject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NegateSubject",
		Description:    `Setting this to true will invert the subject to find what is not matching.`,
		Exposed:        true,
		Getter:         true,
		Name:           "negateSubject",
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"normalizedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
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
	"object": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Object",
		Description:    `Object of the policy.`,
		Exposed:        true,
		Name:           "object",
		Orderable:      true,
		SubType:        "[][]string",
		Type:           "external",
	},
	"observationenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObservationEnabled",
		Description:    `If set to true, the flow will be in observation mode.`,
		Exposed:        true,
		Name:           "observationEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"observedtrafficaction": elemental.AttributeSpecification{
		AllowedChoices: []string{"Apply", "Continue"},
		ConvertedName:  "ObservedTrafficAction",
		DefaultValue:   InfrastructurePolicyObservedTrafficActionContinue,
		Description: `If observationEnabled is set to true, this will be the final action taken on the
packets.`,
		Exposed:   true,
		Name:      "observedTrafficAction",
		Orderable: true,
		Type:      "enum",
	},
	"propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagate will propagate the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description:    `Subject of the policy.`,
		Exposed:        true,
		Name:           "subject",
		Orderable:      true,
		SubType:        "[][]string",
		Type:           "external",
	},
	"updateidempotencykey": elemental.AttributeSpecification{
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
	"updatetime": elemental.AttributeSpecification{
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
	"zhash": elemental.AttributeSpecification{
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
	"zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
}

// SparseInfrastructurePoliciesList represents a list of SparseInfrastructurePolicies
type SparseInfrastructurePoliciesList []*SparseInfrastructurePolicy

// Identity returns the identity of the objects in the list.
func (o SparseInfrastructurePoliciesList) Identity() elemental.Identity {

	return InfrastructurePolicyIdentity
}

// Copy returns a pointer to a copy the SparseInfrastructurePoliciesList.
func (o SparseInfrastructurePoliciesList) Copy() elemental.Identifiables {

	copy := append(SparseInfrastructurePoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseInfrastructurePoliciesList.
func (o SparseInfrastructurePoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseInfrastructurePoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseInfrastructurePolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseInfrastructurePoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseInfrastructurePoliciesList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToPlain returns the SparseInfrastructurePoliciesList converted to InfrastructurePoliciesList.
func (o SparseInfrastructurePoliciesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseInfrastructurePoliciesList) Version() int {

	return 1
}

// SparseInfrastructurePolicy represents the sparse version of a infrastructurepolicy.
type SparseInfrastructurePolicy struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Action defines the action to apply to a flow.
	Action *InfrastructurePolicyActionValue `json:"action,omitempty" msgpack:"action,omitempty" bson:"-" mapstructure:"action,omitempty"`

	// ActiveDuration defines for how long the policy will be active according to the
	// activeSchedule.
	ActiveDuration *string `json:"activeDuration,omitempty" msgpack:"activeDuration,omitempty" bson:"activeduration,omitempty" mapstructure:"activeDuration,omitempty"`

	// ActiveSchedule defines when the policy should be active using the cron notation.
	// The policy will be active for the given activeDuration.
	ActiveSchedule *string `json:"activeSchedule,omitempty" msgpack:"activeSchedule,omitempty" bson:"activeschedule,omitempty" mapstructure:"activeSchedule,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// applyPolicyMode determines if the policy has to be applied to the
	// outgoing traffic of a PU or the incoming traffic of a PU or in both directions.
	// Default is both directions.
	ApplyPolicyMode *InfrastructurePolicyApplyPolicyModeValue `json:"applyPolicyMode,omitempty" msgpack:"applyPolicyMode,omitempty" bson:"-" mapstructure:"applyPolicyMode,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Disabled defines if the propert is disabled.
	Disabled *bool `json:"disabled,omitempty" msgpack:"disabled,omitempty" bson:"disabled,omitempty" mapstructure:"disabled,omitempty"`

	// EncryptionEnabled defines if the flow has to be encrypted.
	EncryptionEnabled *bool `json:"encryptionEnabled,omitempty" msgpack:"encryptionEnabled,omitempty" bson:"-" mapstructure:"encryptionEnabled,omitempty"`

	// If set the policy will be auto deleted after the given time.
	ExpirationTime *time.Time `json:"expirationTime,omitempty" msgpack:"expirationTime,omitempty" bson:"expirationtime,omitempty" mapstructure:"expirationTime,omitempty"`

	// Fallback indicates that this is fallback policy. It will only be
	// applied if no other policies have been resolved. If the policy is also
	// propagated it will become a fallback for children namespaces.
	Fallback *bool `json:"fallback,omitempty" msgpack:"fallback,omitempty" bson:"fallback,omitempty" mapstructure:"fallback,omitempty"`

	// LogsEnabled defines if the flow has to be logged.
	LogsEnabled *bool `json:"logsEnabled,omitempty" msgpack:"logsEnabled,omitempty" bson:"-" mapstructure:"logsEnabled,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Setting this to true will invert the object to find what is not matching.
	NegateObject *bool `json:"negateObject,omitempty" msgpack:"negateObject,omitempty" bson:"negateobject,omitempty" mapstructure:"negateObject,omitempty"`

	// Setting this to true will invert the subject to find what is not matching.
	NegateSubject *bool `json:"negateSubject,omitempty" msgpack:"negateSubject,omitempty" bson:"negatesubject,omitempty" mapstructure:"negateSubject,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Object of the policy.
	Object *[][]string `json:"object,omitempty" msgpack:"object,omitempty" bson:"-" mapstructure:"object,omitempty"`

	// If set to true, the flow will be in observation mode.
	ObservationEnabled *bool `json:"observationEnabled,omitempty" msgpack:"observationEnabled,omitempty" bson:"-" mapstructure:"observationEnabled,omitempty"`

	// If observationEnabled is set to true, this will be the final action taken on the
	// packets.
	ObservedTrafficAction *InfrastructurePolicyObservedTrafficActionValue `json:"observedTrafficAction,omitempty" msgpack:"observedTrafficAction,omitempty" bson:"-" mapstructure:"observedTrafficAction,omitempty"`

	// Propagate will propagate the policy to all of its children.
	Propagate *bool `json:"propagate,omitempty" msgpack:"propagate,omitempty" bson:"propagate,omitempty" mapstructure:"propagate,omitempty"`

	// Protected defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Subject of the policy.
	Subject *[][]string `json:"subject,omitempty" msgpack:"subject,omitempty" bson:"-" mapstructure:"subject,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseInfrastructurePolicy returns a new  SparseInfrastructurePolicy.
func NewSparseInfrastructurePolicy() *SparseInfrastructurePolicy {
	return &SparseInfrastructurePolicy{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseInfrastructurePolicy) Identity() elemental.Identity {

	return InfrastructurePolicyIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseInfrastructurePolicy) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseInfrastructurePolicy) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseInfrastructurePolicy) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseInfrastructurePolicy) ToPlain() elemental.PlainIdentifiable {

	out := NewInfrastructurePolicy()
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
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.ApplyPolicyMode != nil {
		out.ApplyPolicyMode = *o.ApplyPolicyMode
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
	if o.EncryptionEnabled != nil {
		out.EncryptionEnabled = *o.EncryptionEnabled
	}
	if o.ExpirationTime != nil {
		out.ExpirationTime = *o.ExpirationTime
	}
	if o.Fallback != nil {
		out.Fallback = *o.Fallback
	}
	if o.LogsEnabled != nil {
		out.LogsEnabled = *o.LogsEnabled
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NegateObject != nil {
		out.NegateObject = *o.NegateObject
	}
	if o.NegateSubject != nil {
		out.NegateSubject = *o.NegateSubject
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Object != nil {
		out.Object = *o.Object
	}
	if o.ObservationEnabled != nil {
		out.ObservationEnabled = *o.ObservationEnabled
	}
	if o.ObservedTrafficAction != nil {
		out.ObservedTrafficAction = *o.ObservedTrafficAction
	}
	if o.Propagate != nil {
		out.Propagate = *o.Propagate
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Subject != nil {
		out.Subject = *o.Subject
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
func (o *SparseInfrastructurePolicy) GetActiveDuration() string {

	return *o.ActiveDuration
}

// SetActiveDuration sets the property ActiveDuration of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = &activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *SparseInfrastructurePolicy) GetActiveSchedule() string {

	return *o.ActiveSchedule
}

// SetActiveSchedule sets the property ActiveSchedule of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = &activeSchedule
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseInfrastructurePolicy) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseInfrastructurePolicy) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseInfrastructurePolicy) GetCreateIdempotencyKey() string {

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseInfrastructurePolicy) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseInfrastructurePolicy) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetDescription(description string) {

	o.Description = &description
}

// GetDisabled returns the Disabled of the receiver.
func (o *SparseInfrastructurePolicy) GetDisabled() bool {

	return *o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetDisabled(disabled bool) {

	o.Disabled = &disabled
}

// GetExpirationTime returns the ExpirationTime of the receiver.
func (o *SparseInfrastructurePolicy) GetExpirationTime() time.Time {

	return *o.ExpirationTime
}

// SetExpirationTime sets the property ExpirationTime of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetExpirationTime(expirationTime time.Time) {

	o.ExpirationTime = &expirationTime
}

// GetFallback returns the Fallback of the receiver.
func (o *SparseInfrastructurePolicy) GetFallback() bool {

	return *o.Fallback
}

// SetFallback sets the property Fallback of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetFallback(fallback bool) {

	o.Fallback = &fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseInfrastructurePolicy) GetMetadata() []string {

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetName returns the Name of the receiver.
func (o *SparseInfrastructurePolicy) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseInfrastructurePolicy) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNegateObject returns the NegateObject of the receiver.
func (o *SparseInfrastructurePolicy) GetNegateObject() bool {

	return *o.NegateObject
}

// SetNegateObject sets the property NegateObject of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetNegateObject(negateObject bool) {

	o.NegateObject = &negateObject
}

// GetNegateSubject returns the NegateSubject of the receiver.
func (o *SparseInfrastructurePolicy) GetNegateSubject() bool {

	return *o.NegateSubject
}

// SetNegateSubject sets the property NegateSubject of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetNegateSubject(negateSubject bool) {

	o.NegateSubject = &negateSubject
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseInfrastructurePolicy) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *SparseInfrastructurePolicy) GetPropagate() bool {

	return *o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetPropagate(propagate bool) {

	o.Propagate = &propagate
}

// GetProtected returns the Protected of the receiver.
func (o *SparseInfrastructurePolicy) GetProtected() bool {

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseInfrastructurePolicy) GetUpdateIdempotencyKey() string {

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseInfrastructurePolicy) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseInfrastructurePolicy) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseInfrastructurePolicy) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseInfrastructurePolicy) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseInfrastructurePolicy.
func (o *SparseInfrastructurePolicy) DeepCopy() *SparseInfrastructurePolicy {

	if o == nil {
		return nil
	}

	out := &SparseInfrastructurePolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseInfrastructurePolicy.
func (o *SparseInfrastructurePolicy) DeepCopyInto(out *SparseInfrastructurePolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseInfrastructurePolicy: %s", err))
	}

	*out = *target.(*SparseInfrastructurePolicy)
}
