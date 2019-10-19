package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// NetworkAccessPolicyActionValue represents the possible values for attribute "action".
type NetworkAccessPolicyActionValue string

const (
	// NetworkAccessPolicyActionAllow represents the value Allow.
	NetworkAccessPolicyActionAllow NetworkAccessPolicyActionValue = "Allow"

	// NetworkAccessPolicyActionContinue represents the value Continue.
	NetworkAccessPolicyActionContinue NetworkAccessPolicyActionValue = "Continue"

	// NetworkAccessPolicyActionReject represents the value Reject.
	NetworkAccessPolicyActionReject NetworkAccessPolicyActionValue = "Reject"
)

// NetworkAccessPolicyApplyPolicyModeValue represents the possible values for attribute "applyPolicyMode".
type NetworkAccessPolicyApplyPolicyModeValue string

const (
	// NetworkAccessPolicyApplyPolicyModeBidirectional represents the value Bidirectional.
	NetworkAccessPolicyApplyPolicyModeBidirectional NetworkAccessPolicyApplyPolicyModeValue = "Bidirectional"

	// NetworkAccessPolicyApplyPolicyModeIncomingTraffic represents the value IncomingTraffic.
	NetworkAccessPolicyApplyPolicyModeIncomingTraffic NetworkAccessPolicyApplyPolicyModeValue = "IncomingTraffic"

	// NetworkAccessPolicyApplyPolicyModeOutgoingTraffic represents the value OutgoingTraffic.
	NetworkAccessPolicyApplyPolicyModeOutgoingTraffic NetworkAccessPolicyApplyPolicyModeValue = "OutgoingTraffic"
)

// NetworkAccessPolicyObservedTrafficActionValue represents the possible values for attribute "observedTrafficAction".
type NetworkAccessPolicyObservedTrafficActionValue string

const (
	// NetworkAccessPolicyObservedTrafficActionApply represents the value Apply.
	NetworkAccessPolicyObservedTrafficActionApply NetworkAccessPolicyObservedTrafficActionValue = "Apply"

	// NetworkAccessPolicyObservedTrafficActionContinue represents the value Continue.
	NetworkAccessPolicyObservedTrafficActionContinue NetworkAccessPolicyObservedTrafficActionValue = "Continue"
)

// NetworkAccessPolicyIdentity represents the Identity of the object.
var NetworkAccessPolicyIdentity = elemental.Identity{
	Name:     "networkaccesspolicy",
	Category: "networkaccesspolicies",
	Package:  "squall",
	Private:  false,
}

// NetworkAccessPoliciesList represents a list of NetworkAccessPolicies
type NetworkAccessPoliciesList []*NetworkAccessPolicy

// Identity returns the identity of the objects in the list.
func (o NetworkAccessPoliciesList) Identity() elemental.Identity {

	return NetworkAccessPolicyIdentity
}

// Copy returns a pointer to a copy the NetworkAccessPoliciesList.
func (o NetworkAccessPoliciesList) Copy() elemental.Identifiables {

	copy := append(NetworkAccessPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the NetworkAccessPoliciesList.
func (o NetworkAccessPoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(NetworkAccessPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*NetworkAccessPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o NetworkAccessPoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o NetworkAccessPoliciesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the NetworkAccessPoliciesList converted to SparseNetworkAccessPoliciesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o NetworkAccessPoliciesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseNetworkAccessPoliciesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseNetworkAccessPolicy)
	}

	return out
}

// Version returns the version of the content.
func (o NetworkAccessPoliciesList) Version() int {

	return 1
}

// NetworkAccessPolicy represents the model of a networkaccesspolicy
type NetworkAccessPolicy struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Defines the action to apply to a flow.
	//
	// - `Allow`: allows the defined traffic.
	// - `Reject`: rejects the defined traffic; useful in conjunction with an allow all policy.
	// - `Continue`: neither allows or rejects the traffic; useful for applying another property to the traffic, such as encryption.
	Action NetworkAccessPolicyActionValue `json:"action" msgpack:"action" bson:"-" mapstructure:"action,omitempty"`

	// Defines for how long the policy will be active according to the
	// `activeSchedule`.
	ActiveDuration string `json:"activeDuration" msgpack:"activeDuration" bson:"activeduration" mapstructure:"activeDuration,omitempty"`

	// Defines when the policy should be active using the cron notation.
	// The policy will be active for the given `activeDuration`.
	ActiveSchedule string `json:"activeSchedule" msgpack:"activeSchedule" bson:"activeschedule" mapstructure:"activeSchedule,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// Sets three different types of policies. `IncomingTraffic`: applies the policy to
	// all
	// processing units that match the `object` and allows them to *accept* connections
	// from
	// processing units or external networks that match the `subject`.
	// `OutgoingTraffic`: applies
	// the policy to all processing units that match the `subject` and allows them to
	// *initiate*
	// connections with processing units or external networks that match the `object`.
	// `Bidirectional` (default): applies the policy to all processing units that match
	// the `object`
	// and allows them to *accept* connections from processing units that match the
	// `subject`.
	// Also applies the policy to all processing units that match the `subject` and
	// allows them
	// to *initiate* connections with processing units that match the `object`.
	ApplyPolicyMode NetworkAccessPolicyApplyPolicyModeValue `json:"applyPolicyMode" msgpack:"applyPolicyMode" bson:"-" mapstructure:"applyPolicyMode,omitempty"`

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

	// Defines if the flow has to be encrypted.
	EncryptionEnabled bool `json:"encryptionEnabled" msgpack:"encryptionEnabled" bson:"-" mapstructure:"encryptionEnabled,omitempty"`

	// If set the policy will be automatically deleted after the given time.
	ExpirationTime time.Time `json:"expirationTime" msgpack:"expirationTime" bson:"expirationtime" mapstructure:"expirationTime,omitempty"`

	// Indicates that this is fallback policy. It will only be
	// applied if no other policies have been resolved. If the policy is also
	// propagated it will become a fallback for children namespaces.
	Fallback bool `json:"fallback" msgpack:"fallback" bson:"fallback" mapstructure:"fallback,omitempty"`

	// If `true`, the relevant flows are logged and available from the Aporeto control
	// plane.
	// Under some advanced scenarios you may wish to set this to `false`, such as to
	// save space or
	// improve performance.
	LogsEnabled bool `json:"logsEnabled" msgpack:"logsEnabled" bson:"-" mapstructure:"logsEnabled,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Setting this to `true` will invert the object to find what is not matching.
	NegateObject bool `json:"negateObject" msgpack:"negateObject" bson:"negateobject" mapstructure:"negateObject,omitempty"`

	// Setting this to `true` will invert the subject to find what is not matching.
	NegateSubject bool `json:"negateSubject" msgpack:"negateSubject" bson:"negatesubject" mapstructure:"negateSubject,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// A tag or tag expression identifying the object of the policy.
	Object [][]string `json:"object" msgpack:"object" bson:"-" mapstructure:"object,omitempty"`

	// If set to `true`, the flow will be in observation mode.
	ObservationEnabled bool `json:"observationEnabled" msgpack:"observationEnabled" bson:"-" mapstructure:"observationEnabled,omitempty"`

	// If `observationEnabled` is set to `true`, this defines the final action taken
	// on the packets: `Apply` or `Continue` (default).
	ObservedTrafficAction NetworkAccessPolicyObservedTrafficActionValue `json:"observedTrafficAction" msgpack:"observedTrafficAction" bson:"-" mapstructure:"observedTrafficAction,omitempty"`

	// Represents the ports and protocols this policy applies to.
	Ports []string `json:"ports" msgpack:"ports" bson:"-" mapstructure:"ports,omitempty"`

	// Propagates the policy to all of its children.
	Propagate bool `json:"propagate" msgpack:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// A tag or tag expression identifying the subject of the policy.
	Subject [][]string `json:"subject" msgpack:"subject" bson:"-" mapstructure:"subject,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewNetworkAccessPolicy returns a new *NetworkAccessPolicy
func NewNetworkAccessPolicy() *NetworkAccessPolicy {

	return &NetworkAccessPolicy{
		ModelVersion:          1,
		Action:                NetworkAccessPolicyActionAllow,
		AssociatedTags:        []string{},
		Annotations:           map[string][]string{},
		ApplyPolicyMode:       NetworkAccessPolicyApplyPolicyModeBidirectional,
		Metadata:              []string{},
		ObservedTrafficAction: NetworkAccessPolicyObservedTrafficActionContinue,
		NormalizedTags:        []string{},
		Object:                [][]string{},
		Ports:                 []string{},
		Subject:               [][]string{},
	}
}

// Identity returns the Identity of the object.
func (o *NetworkAccessPolicy) Identity() elemental.Identity {

	return NetworkAccessPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NetworkAccessPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NetworkAccessPolicy) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NetworkAccessPolicy) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesNetworkAccessPolicy{}

	s.ActiveDuration = o.ActiveDuration
	s.ActiveSchedule = o.ActiveSchedule
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.Description = o.Description
	s.Disabled = o.Disabled
	s.ExpirationTime = o.ExpirationTime
	s.Fallback = o.Fallback
	s.Metadata = o.Metadata
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NegateObject = o.NegateObject
	s.NegateSubject = o.NegateSubject
	s.NormalizedTags = o.NormalizedTags
	s.Propagate = o.Propagate
	s.Protected = o.Protected
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NetworkAccessPolicy) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesNetworkAccessPolicy{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ActiveDuration = s.ActiveDuration
	o.ActiveSchedule = s.ActiveSchedule
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.Description = s.Description
	o.Disabled = s.Disabled
	o.ExpirationTime = s.ExpirationTime
	o.Fallback = s.Fallback
	o.Metadata = s.Metadata
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NegateObject = s.NegateObject
	o.NegateSubject = s.NegateSubject
	o.NormalizedTags = s.NormalizedTags
	o.Propagate = s.Propagate
	o.Protected = s.Protected
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime

	return nil
}

// Version returns the hardcoded version of the model.
func (o *NetworkAccessPolicy) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *NetworkAccessPolicy) BleveType() string {

	return "networkaccesspolicy"
}

// DefaultOrder returns the list of default ordering fields.
func (o *NetworkAccessPolicy) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *NetworkAccessPolicy) Doc() string {

	return `Allows you to define network policies to allow or prevent processing units
identified by their tags to talk to other processing units or external networks
(also identified by their tags).`
}

func (o *NetworkAccessPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetActiveDuration returns the ActiveDuration of the receiver.
func (o *NetworkAccessPolicy) GetActiveDuration() string {

	return o.ActiveDuration
}

// SetActiveDuration sets the property ActiveDuration of the receiver using the given value.
func (o *NetworkAccessPolicy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *NetworkAccessPolicy) GetActiveSchedule() string {

	return o.ActiveSchedule
}

// SetActiveSchedule sets the property ActiveSchedule of the receiver using the given value.
func (o *NetworkAccessPolicy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = activeSchedule
}

// GetAnnotations returns the Annotations of the receiver.
func (o *NetworkAccessPolicy) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *NetworkAccessPolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *NetworkAccessPolicy) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *NetworkAccessPolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *NetworkAccessPolicy) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *NetworkAccessPolicy) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *NetworkAccessPolicy) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *NetworkAccessPolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *NetworkAccessPolicy) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *NetworkAccessPolicy) SetDescription(description string) {

	o.Description = description
}

// GetDisabled returns the Disabled of the receiver.
func (o *NetworkAccessPolicy) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the given value.
func (o *NetworkAccessPolicy) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetExpirationTime returns the ExpirationTime of the receiver.
func (o *NetworkAccessPolicy) GetExpirationTime() time.Time {

	return o.ExpirationTime
}

// SetExpirationTime sets the property ExpirationTime of the receiver using the given value.
func (o *NetworkAccessPolicy) SetExpirationTime(expirationTime time.Time) {

	o.ExpirationTime = expirationTime
}

// GetFallback returns the Fallback of the receiver.
func (o *NetworkAccessPolicy) GetFallback() bool {

	return o.Fallback
}

// SetFallback sets the property Fallback of the receiver using the given value.
func (o *NetworkAccessPolicy) SetFallback(fallback bool) {

	o.Fallback = fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *NetworkAccessPolicy) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *NetworkAccessPolicy) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *NetworkAccessPolicy) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *NetworkAccessPolicy) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *NetworkAccessPolicy) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *NetworkAccessPolicy) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNegateObject returns the NegateObject of the receiver.
func (o *NetworkAccessPolicy) GetNegateObject() bool {

	return o.NegateObject
}

// SetNegateObject sets the property NegateObject of the receiver using the given value.
func (o *NetworkAccessPolicy) SetNegateObject(negateObject bool) {

	o.NegateObject = negateObject
}

// GetNegateSubject returns the NegateSubject of the receiver.
func (o *NetworkAccessPolicy) GetNegateSubject() bool {

	return o.NegateSubject
}

// SetNegateSubject sets the property NegateSubject of the receiver using the given value.
func (o *NetworkAccessPolicy) SetNegateSubject(negateSubject bool) {

	o.NegateSubject = negateSubject
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *NetworkAccessPolicy) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *NetworkAccessPolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *NetworkAccessPolicy) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the given value.
func (o *NetworkAccessPolicy) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// GetProtected returns the Protected of the receiver.
func (o *NetworkAccessPolicy) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *NetworkAccessPolicy) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *NetworkAccessPolicy) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *NetworkAccessPolicy) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *NetworkAccessPolicy) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *NetworkAccessPolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *NetworkAccessPolicy) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseNetworkAccessPolicy{
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
			Ports:                 &o.Ports,
			Propagate:             &o.Propagate,
			Protected:             &o.Protected,
			Subject:               &o.Subject,
			UpdateIdempotencyKey:  &o.UpdateIdempotencyKey,
			UpdateTime:            &o.UpdateTime,
		}
	}

	sp := &SparseNetworkAccessPolicy{}
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
		case "ports":
			sp.Ports = &(o.Ports)
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
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseNetworkAccessPolicy to the object.
func (o *NetworkAccessPolicy) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseNetworkAccessPolicy)
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
	if so.Ports != nil {
		o.Ports = *so.Ports
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
}

// DeepCopy returns a deep copy if the NetworkAccessPolicy.
func (o *NetworkAccessPolicy) DeepCopy() *NetworkAccessPolicy {

	if o == nil {
		return nil
	}

	out := &NetworkAccessPolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *NetworkAccessPolicy.
func (o *NetworkAccessPolicy) DeepCopyInto(out *NetworkAccessPolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy NetworkAccessPolicy: %s", err))
	}

	*out = *target.(*NetworkAccessPolicy)
}

// Validate valides the current information stored into the structure.
func (o *NetworkAccessPolicy) Validate() error {

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

	if err := ValidateProtoPorts("ports", o.Ports); err != nil {
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
func (*NetworkAccessPolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := NetworkAccessPolicyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return NetworkAccessPolicyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*NetworkAccessPolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return NetworkAccessPolicyAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *NetworkAccessPolicy) ValueForAttribute(name string) interface{} {

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
	case "ports":
		return o.Ports
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
	}

	return nil
}

// NetworkAccessPolicyAttributesMap represents the map of attribute for NetworkAccessPolicy.
var NetworkAccessPolicyAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
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
		Type:           "string",
	},
	"Action": elemental.AttributeSpecification{
		AllowedChoices: []string{"Allow", "Reject", "Continue"},
		ConvertedName:  "Action",
		DefaultValue:   NetworkAccessPolicyActionAllow,
		Description: `Defines the action to apply to a flow.

- ` + "`" + `Allow` + "`" + `: allows the defined traffic.
- ` + "`" + `Reject` + "`" + `: rejects the defined traffic; useful in conjunction with an allow all policy.
- ` + "`" + `Continue` + "`" + `: neither allows or rejects the traffic; useful for applying another property to the traffic, such as encryption.`,
		Exposed:   true,
		Name:      "action",
		Orderable: true,
		Type:      "enum",
	},
	"ActiveDuration": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
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
	"ActiveSchedule": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"Annotations": elemental.AttributeSpecification{
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
	"ApplyPolicyMode": elemental.AttributeSpecification{
		AllowedChoices: []string{"OutgoingTraffic", "IncomingTraffic", "Bidirectional"},
		ConvertedName:  "ApplyPolicyMode",
		DefaultValue:   NetworkAccessPolicyApplyPolicyModeBidirectional,
		Description: `Sets three different types of policies. ` + "`" + `IncomingTraffic` + "`" + `: applies the policy to
all
processing units that match the ` + "`" + `object` + "`" + ` and allows them to *accept* connections
from
processing units or external networks that match the ` + "`" + `subject` + "`" + `.
` + "`" + `OutgoingTraffic` + "`" + `: applies
the policy to all processing units that match the ` + "`" + `subject` + "`" + ` and allows them to
*initiate*
connections with processing units or external networks that match the ` + "`" + `object` + "`" + `.
` + "`" + `Bidirectional` + "`" + ` (default): applies the policy to all processing units that match
the ` + "`" + `object` + "`" + `
and allows them to *accept* connections from processing units that match the
` + "`" + `subject` + "`" + `.
Also applies the policy to all processing units that match the ` + "`" + `subject` + "`" + ` and
allows them
to *initiate* connections with processing units that match the ` + "`" + `object` + "`" + `.`,
		Exposed:   true,
		Name:      "applyPolicyMode",
		Orderable: true,
		Type:      "enum",
	},
	"AssociatedTags": elemental.AttributeSpecification{
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
	"Disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"EncryptionEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EncryptionEnabled",
		Description:    `Defines if the flow has to be encrypted.`,
		Exposed:        true,
		Name:           "encryptionEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"ExpirationTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationTime",
		Description:    `If set the policy will be automatically deleted after the given time.`,
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
	"LogsEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LogsEnabled",
		Description: `If ` + "`" + `true` + "`" + `, the relevant flows are logged and available from the Aporeto control
plane.
Under some advanced scenarios you may wish to set this to ` + "`" + `false` + "`" + `, such as to
save space or
improve performance.`,
		Exposed:   true,
		Name:      "logsEnabled",
		Orderable: true,
		Type:      "boolean",
	},
	"Metadata": elemental.AttributeSpecification{
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
	"Name": elemental.AttributeSpecification{
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
	"Namespace": elemental.AttributeSpecification{
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
	"NegateObject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NegateObject",
		Description:    `Setting this to ` + "`" + `true` + "`" + ` will invert the object to find what is not matching.`,
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
		Description:    `Setting this to ` + "`" + `true` + "`" + ` will invert the subject to find what is not matching.`,
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
	"Object": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Object",
		Description:    `A tag or tag expression identifying the object of the policy.`,
		Exposed:        true,
		Name:           "object",
		Orderable:      true,
		SubType:        "[][]string",
		Type:           "external",
	},
	"ObservationEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObservationEnabled",
		Description:    `If set to ` + "`" + `true` + "`" + `, the flow will be in observation mode.`,
		Exposed:        true,
		Name:           "observationEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"ObservedTrafficAction": elemental.AttributeSpecification{
		AllowedChoices: []string{"Apply", "Continue"},
		ConvertedName:  "ObservedTrafficAction",
		DefaultValue:   NetworkAccessPolicyObservedTrafficActionContinue,
		Description: `If ` + "`" + `observationEnabled` + "`" + ` is set to ` + "`" + `true` + "`" + `, this defines the final action taken
on the packets: ` + "`" + `Apply` + "`" + ` or ` + "`" + `Continue` + "`" + ` (default).`,
		Exposed:   true,
		Name:      "observedTrafficAction",
		Orderable: true,
		Type:      "enum",
	},
	"Ports": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Ports",
		Description:    `Represents the ports and protocols this policy applies to.`,
		Exposed:        true,
		Name:           "ports",
		Orderable:      true,
		SubType:        "string",
		Type:           "list",
	},
	"Propagate": elemental.AttributeSpecification{
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
	"Protected": elemental.AttributeSpecification{
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
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description:    `A tag or tag expression identifying the subject of the policy.`,
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
}

// NetworkAccessPolicyLowerCaseAttributesMap represents the map of attribute for NetworkAccessPolicy.
var NetworkAccessPolicyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
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
		Type:           "string",
	},
	"action": elemental.AttributeSpecification{
		AllowedChoices: []string{"Allow", "Reject", "Continue"},
		ConvertedName:  "Action",
		DefaultValue:   NetworkAccessPolicyActionAllow,
		Description: `Defines the action to apply to a flow.

- ` + "`" + `Allow` + "`" + `: allows the defined traffic.
- ` + "`" + `Reject` + "`" + `: rejects the defined traffic; useful in conjunction with an allow all policy.
- ` + "`" + `Continue` + "`" + `: neither allows or rejects the traffic; useful for applying another property to the traffic, such as encryption.`,
		Exposed:   true,
		Name:      "action",
		Orderable: true,
		Type:      "enum",
	},
	"activeduration": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
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
	"activeschedule": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"annotations": elemental.AttributeSpecification{
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
	"applypolicymode": elemental.AttributeSpecification{
		AllowedChoices: []string{"OutgoingTraffic", "IncomingTraffic", "Bidirectional"},
		ConvertedName:  "ApplyPolicyMode",
		DefaultValue:   NetworkAccessPolicyApplyPolicyModeBidirectional,
		Description: `Sets three different types of policies. ` + "`" + `IncomingTraffic` + "`" + `: applies the policy to
all
processing units that match the ` + "`" + `object` + "`" + ` and allows them to *accept* connections
from
processing units or external networks that match the ` + "`" + `subject` + "`" + `.
` + "`" + `OutgoingTraffic` + "`" + `: applies
the policy to all processing units that match the ` + "`" + `subject` + "`" + ` and allows them to
*initiate*
connections with processing units or external networks that match the ` + "`" + `object` + "`" + `.
` + "`" + `Bidirectional` + "`" + ` (default): applies the policy to all processing units that match
the ` + "`" + `object` + "`" + `
and allows them to *accept* connections from processing units that match the
` + "`" + `subject` + "`" + `.
Also applies the policy to all processing units that match the ` + "`" + `subject` + "`" + ` and
allows them
to *initiate* connections with processing units that match the ` + "`" + `object` + "`" + `.`,
		Exposed:   true,
		Name:      "applyPolicyMode",
		Orderable: true,
		Type:      "enum",
	},
	"associatedtags": elemental.AttributeSpecification{
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
	"disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"encryptionenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EncryptionEnabled",
		Description:    `Defines if the flow has to be encrypted.`,
		Exposed:        true,
		Name:           "encryptionEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"expirationtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationTime",
		Description:    `If set the policy will be automatically deleted after the given time.`,
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
	"logsenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LogsEnabled",
		Description: `If ` + "`" + `true` + "`" + `, the relevant flows are logged and available from the Aporeto control
plane.
Under some advanced scenarios you may wish to set this to ` + "`" + `false` + "`" + `, such as to
save space or
improve performance.`,
		Exposed:   true,
		Name:      "logsEnabled",
		Orderable: true,
		Type:      "boolean",
	},
	"metadata": elemental.AttributeSpecification{
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
	"name": elemental.AttributeSpecification{
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
	"namespace": elemental.AttributeSpecification{
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
	"negateobject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NegateObject",
		Description:    `Setting this to ` + "`" + `true` + "`" + ` will invert the object to find what is not matching.`,
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
		Description:    `Setting this to ` + "`" + `true` + "`" + ` will invert the subject to find what is not matching.`,
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
	"object": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Object",
		Description:    `A tag or tag expression identifying the object of the policy.`,
		Exposed:        true,
		Name:           "object",
		Orderable:      true,
		SubType:        "[][]string",
		Type:           "external",
	},
	"observationenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObservationEnabled",
		Description:    `If set to ` + "`" + `true` + "`" + `, the flow will be in observation mode.`,
		Exposed:        true,
		Name:           "observationEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"observedtrafficaction": elemental.AttributeSpecification{
		AllowedChoices: []string{"Apply", "Continue"},
		ConvertedName:  "ObservedTrafficAction",
		DefaultValue:   NetworkAccessPolicyObservedTrafficActionContinue,
		Description: `If ` + "`" + `observationEnabled` + "`" + ` is set to ` + "`" + `true` + "`" + `, this defines the final action taken
on the packets: ` + "`" + `Apply` + "`" + ` or ` + "`" + `Continue` + "`" + ` (default).`,
		Exposed:   true,
		Name:      "observedTrafficAction",
		Orderable: true,
		Type:      "enum",
	},
	"ports": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Ports",
		Description:    `Represents the ports and protocols this policy applies to.`,
		Exposed:        true,
		Name:           "ports",
		Orderable:      true,
		SubType:        "string",
		Type:           "list",
	},
	"propagate": elemental.AttributeSpecification{
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
	"protected": elemental.AttributeSpecification{
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
	"subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description:    `A tag or tag expression identifying the subject of the policy.`,
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
}

// SparseNetworkAccessPoliciesList represents a list of SparseNetworkAccessPolicies
type SparseNetworkAccessPoliciesList []*SparseNetworkAccessPolicy

// Identity returns the identity of the objects in the list.
func (o SparseNetworkAccessPoliciesList) Identity() elemental.Identity {

	return NetworkAccessPolicyIdentity
}

// Copy returns a pointer to a copy the SparseNetworkAccessPoliciesList.
func (o SparseNetworkAccessPoliciesList) Copy() elemental.Identifiables {

	copy := append(SparseNetworkAccessPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseNetworkAccessPoliciesList.
func (o SparseNetworkAccessPoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseNetworkAccessPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseNetworkAccessPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseNetworkAccessPoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseNetworkAccessPoliciesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseNetworkAccessPoliciesList converted to NetworkAccessPoliciesList.
func (o SparseNetworkAccessPoliciesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseNetworkAccessPoliciesList) Version() int {

	return 1
}

// SparseNetworkAccessPolicy represents the sparse version of a networkaccesspolicy.
type SparseNetworkAccessPolicy struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Defines the action to apply to a flow.
	//
	// - `Allow`: allows the defined traffic.
	// - `Reject`: rejects the defined traffic; useful in conjunction with an allow all policy.
	// - `Continue`: neither allows or rejects the traffic; useful for applying another property to the traffic, such as encryption.
	Action *NetworkAccessPolicyActionValue `json:"action,omitempty" msgpack:"action,omitempty" bson:"-" mapstructure:"action,omitempty"`

	// Defines for how long the policy will be active according to the
	// `activeSchedule`.
	ActiveDuration *string `json:"activeDuration,omitempty" msgpack:"activeDuration,omitempty" bson:"activeduration,omitempty" mapstructure:"activeDuration,omitempty"`

	// Defines when the policy should be active using the cron notation.
	// The policy will be active for the given `activeDuration`.
	ActiveSchedule *string `json:"activeSchedule,omitempty" msgpack:"activeSchedule,omitempty" bson:"activeschedule,omitempty" mapstructure:"activeSchedule,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// Sets three different types of policies. `IncomingTraffic`: applies the policy to
	// all
	// processing units that match the `object` and allows them to *accept* connections
	// from
	// processing units or external networks that match the `subject`.
	// `OutgoingTraffic`: applies
	// the policy to all processing units that match the `subject` and allows them to
	// *initiate*
	// connections with processing units or external networks that match the `object`.
	// `Bidirectional` (default): applies the policy to all processing units that match
	// the `object`
	// and allows them to *accept* connections from processing units that match the
	// `subject`.
	// Also applies the policy to all processing units that match the `subject` and
	// allows them
	// to *initiate* connections with processing units that match the `object`.
	ApplyPolicyMode *NetworkAccessPolicyApplyPolicyModeValue `json:"applyPolicyMode,omitempty" msgpack:"applyPolicyMode,omitempty" bson:"-" mapstructure:"applyPolicyMode,omitempty"`

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

	// Defines if the flow has to be encrypted.
	EncryptionEnabled *bool `json:"encryptionEnabled,omitempty" msgpack:"encryptionEnabled,omitempty" bson:"-" mapstructure:"encryptionEnabled,omitempty"`

	// If set the policy will be automatically deleted after the given time.
	ExpirationTime *time.Time `json:"expirationTime,omitempty" msgpack:"expirationTime,omitempty" bson:"expirationtime,omitempty" mapstructure:"expirationTime,omitempty"`

	// Indicates that this is fallback policy. It will only be
	// applied if no other policies have been resolved. If the policy is also
	// propagated it will become a fallback for children namespaces.
	Fallback *bool `json:"fallback,omitempty" msgpack:"fallback,omitempty" bson:"fallback,omitempty" mapstructure:"fallback,omitempty"`

	// If `true`, the relevant flows are logged and available from the Aporeto control
	// plane.
	// Under some advanced scenarios you may wish to set this to `false`, such as to
	// save space or
	// improve performance.
	LogsEnabled *bool `json:"logsEnabled,omitempty" msgpack:"logsEnabled,omitempty" bson:"-" mapstructure:"logsEnabled,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Setting this to `true` will invert the object to find what is not matching.
	NegateObject *bool `json:"negateObject,omitempty" msgpack:"negateObject,omitempty" bson:"negateobject,omitempty" mapstructure:"negateObject,omitempty"`

	// Setting this to `true` will invert the subject to find what is not matching.
	NegateSubject *bool `json:"negateSubject,omitempty" msgpack:"negateSubject,omitempty" bson:"negatesubject,omitempty" mapstructure:"negateSubject,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// A tag or tag expression identifying the object of the policy.
	Object *[][]string `json:"object,omitempty" msgpack:"object,omitempty" bson:"-" mapstructure:"object,omitempty"`

	// If set to `true`, the flow will be in observation mode.
	ObservationEnabled *bool `json:"observationEnabled,omitempty" msgpack:"observationEnabled,omitempty" bson:"-" mapstructure:"observationEnabled,omitempty"`

	// If `observationEnabled` is set to `true`, this defines the final action taken
	// on the packets: `Apply` or `Continue` (default).
	ObservedTrafficAction *NetworkAccessPolicyObservedTrafficActionValue `json:"observedTrafficAction,omitempty" msgpack:"observedTrafficAction,omitempty" bson:"-" mapstructure:"observedTrafficAction,omitempty"`

	// Represents the ports and protocols this policy applies to.
	Ports *[]string `json:"ports,omitempty" msgpack:"ports,omitempty" bson:"-" mapstructure:"ports,omitempty"`

	// Propagates the policy to all of its children.
	Propagate *bool `json:"propagate,omitempty" msgpack:"propagate,omitempty" bson:"propagate,omitempty" mapstructure:"propagate,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// A tag or tag expression identifying the subject of the policy.
	Subject *[][]string `json:"subject,omitempty" msgpack:"subject,omitempty" bson:"-" mapstructure:"subject,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseNetworkAccessPolicy returns a new  SparseNetworkAccessPolicy.
func NewSparseNetworkAccessPolicy() *SparseNetworkAccessPolicy {
	return &SparseNetworkAccessPolicy{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseNetworkAccessPolicy) Identity() elemental.Identity {

	return NetworkAccessPolicyIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseNetworkAccessPolicy) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseNetworkAccessPolicy) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseNetworkAccessPolicy) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseNetworkAccessPolicy{}

	if o.ActiveDuration != nil {
		s.ActiveDuration = o.ActiveDuration
	}
	if o.ActiveSchedule != nil {
		s.ActiveSchedule = o.ActiveSchedule
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
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.NegateObject != nil {
		s.NegateObject = o.NegateObject
	}
	if o.NegateSubject != nil {
		s.NegateSubject = o.NegateSubject
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
	if o.UpdateIdempotencyKey != nil {
		s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		s.UpdateTime = o.UpdateTime
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseNetworkAccessPolicy) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseNetworkAccessPolicy{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.ActiveDuration != nil {
		o.ActiveDuration = s.ActiveDuration
	}
	if s.ActiveSchedule != nil {
		o.ActiveSchedule = s.ActiveSchedule
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
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.NegateObject != nil {
		o.NegateObject = s.NegateObject
	}
	if s.NegateSubject != nil {
		o.NegateSubject = s.NegateSubject
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
	if s.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	}
	if s.UpdateTime != nil {
		o.UpdateTime = s.UpdateTime
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseNetworkAccessPolicy) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseNetworkAccessPolicy) ToPlain() elemental.PlainIdentifiable {

	out := NewNetworkAccessPolicy()
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
	if o.Ports != nil {
		out.Ports = *o.Ports
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

	return out
}

// GetActiveDuration returns the ActiveDuration of the receiver.
func (o *SparseNetworkAccessPolicy) GetActiveDuration() string {

	return *o.ActiveDuration
}

// SetActiveDuration sets the property ActiveDuration of the receiver using the address of the given value.
func (o *SparseNetworkAccessPolicy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = &activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *SparseNetworkAccessPolicy) GetActiveSchedule() string {

	return *o.ActiveSchedule
}

// SetActiveSchedule sets the property ActiveSchedule of the receiver using the address of the given value.
func (o *SparseNetworkAccessPolicy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = &activeSchedule
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseNetworkAccessPolicy) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseNetworkAccessPolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseNetworkAccessPolicy) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseNetworkAccessPolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseNetworkAccessPolicy) GetCreateIdempotencyKey() string {

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseNetworkAccessPolicy) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseNetworkAccessPolicy) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseNetworkAccessPolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseNetworkAccessPolicy) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseNetworkAccessPolicy) SetDescription(description string) {

	o.Description = &description
}

// GetDisabled returns the Disabled of the receiver.
func (o *SparseNetworkAccessPolicy) GetDisabled() bool {

	return *o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the address of the given value.
func (o *SparseNetworkAccessPolicy) SetDisabled(disabled bool) {

	o.Disabled = &disabled
}

// GetExpirationTime returns the ExpirationTime of the receiver.
func (o *SparseNetworkAccessPolicy) GetExpirationTime() time.Time {

	return *o.ExpirationTime
}

// SetExpirationTime sets the property ExpirationTime of the receiver using the address of the given value.
func (o *SparseNetworkAccessPolicy) SetExpirationTime(expirationTime time.Time) {

	o.ExpirationTime = &expirationTime
}

// GetFallback returns the Fallback of the receiver.
func (o *SparseNetworkAccessPolicy) GetFallback() bool {

	return *o.Fallback
}

// SetFallback sets the property Fallback of the receiver using the address of the given value.
func (o *SparseNetworkAccessPolicy) SetFallback(fallback bool) {

	o.Fallback = &fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseNetworkAccessPolicy) GetMetadata() []string {

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseNetworkAccessPolicy) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetName returns the Name of the receiver.
func (o *SparseNetworkAccessPolicy) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseNetworkAccessPolicy) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseNetworkAccessPolicy) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseNetworkAccessPolicy) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNegateObject returns the NegateObject of the receiver.
func (o *SparseNetworkAccessPolicy) GetNegateObject() bool {

	return *o.NegateObject
}

// SetNegateObject sets the property NegateObject of the receiver using the address of the given value.
func (o *SparseNetworkAccessPolicy) SetNegateObject(negateObject bool) {

	o.NegateObject = &negateObject
}

// GetNegateSubject returns the NegateSubject of the receiver.
func (o *SparseNetworkAccessPolicy) GetNegateSubject() bool {

	return *o.NegateSubject
}

// SetNegateSubject sets the property NegateSubject of the receiver using the address of the given value.
func (o *SparseNetworkAccessPolicy) SetNegateSubject(negateSubject bool) {

	o.NegateSubject = &negateSubject
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseNetworkAccessPolicy) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseNetworkAccessPolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *SparseNetworkAccessPolicy) GetPropagate() bool {

	return *o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the address of the given value.
func (o *SparseNetworkAccessPolicy) SetPropagate(propagate bool) {

	o.Propagate = &propagate
}

// GetProtected returns the Protected of the receiver.
func (o *SparseNetworkAccessPolicy) GetProtected() bool {

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseNetworkAccessPolicy) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseNetworkAccessPolicy) GetUpdateIdempotencyKey() string {

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseNetworkAccessPolicy) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseNetworkAccessPolicy) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseNetworkAccessPolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// DeepCopy returns a deep copy if the SparseNetworkAccessPolicy.
func (o *SparseNetworkAccessPolicy) DeepCopy() *SparseNetworkAccessPolicy {

	if o == nil {
		return nil
	}

	out := &SparseNetworkAccessPolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseNetworkAccessPolicy.
func (o *SparseNetworkAccessPolicy) DeepCopyInto(out *SparseNetworkAccessPolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseNetworkAccessPolicy: %s", err))
	}

	*out = *target.(*SparseNetworkAccessPolicy)
}

type mongoAttributesNetworkAccessPolicy struct {
	ActiveDuration       string              `bson:"activeduration"`
	ActiveSchedule       string              `bson:"activeschedule"`
	Annotations          map[string][]string `bson:"annotations"`
	AssociatedTags       []string            `bson:"associatedtags"`
	CreateIdempotencyKey string              `bson:"createidempotencykey"`
	CreateTime           time.Time           `bson:"createtime"`
	Description          string              `bson:"description"`
	Disabled             bool                `bson:"disabled"`
	ExpirationTime       time.Time           `bson:"expirationtime"`
	Fallback             bool                `bson:"fallback"`
	Metadata             []string            `bson:"metadata"`
	Name                 string              `bson:"name"`
	Namespace            string              `bson:"namespace"`
	NegateObject         bool                `bson:"negateobject"`
	NegateSubject        bool                `bson:"negatesubject"`
	NormalizedTags       []string            `bson:"normalizedtags"`
	Propagate            bool                `bson:"propagate"`
	Protected            bool                `bson:"protected"`
	UpdateIdempotencyKey string              `bson:"updateidempotencykey"`
	UpdateTime           time.Time           `bson:"updatetime"`
}
type mongoAttributesSparseNetworkAccessPolicy struct {
	ActiveDuration       *string              `bson:"activeduration,omitempty"`
	ActiveSchedule       *string              `bson:"activeschedule,omitempty"`
	Annotations          *map[string][]string `bson:"annotations,omitempty"`
	AssociatedTags       *[]string            `bson:"associatedtags,omitempty"`
	CreateIdempotencyKey *string              `bson:"createidempotencykey,omitempty"`
	CreateTime           *time.Time           `bson:"createtime,omitempty"`
	Description          *string              `bson:"description,omitempty"`
	Disabled             *bool                `bson:"disabled,omitempty"`
	ExpirationTime       *time.Time           `bson:"expirationtime,omitempty"`
	Fallback             *bool                `bson:"fallback,omitempty"`
	Metadata             *[]string            `bson:"metadata,omitempty"`
	Name                 *string              `bson:"name,omitempty"`
	Namespace            *string              `bson:"namespace,omitempty"`
	NegateObject         *bool                `bson:"negateobject,omitempty"`
	NegateSubject        *bool                `bson:"negatesubject,omitempty"`
	NormalizedTags       *[]string            `bson:"normalizedtags,omitempty"`
	Propagate            *bool                `bson:"propagate,omitempty"`
	Protected            *bool                `bson:"protected,omitempty"`
	UpdateIdempotencyKey *string              `bson:"updateidempotencykey,omitempty"`
	UpdateTime           *time.Time           `bson:"updatetime,omitempty"`
}
