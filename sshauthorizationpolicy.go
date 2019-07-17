package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// SSHAuthorizationPolicyIdentity represents the Identity of the object.
var SSHAuthorizationPolicyIdentity = elemental.Identity{
	Name:     "sshauthorizationpolicy",
	Category: "sshauthorizationpolicies",
	Package:  "squall",
	Private:  false,
}

// SSHAuthorizationPoliciesList represents a list of SSHAuthorizationPolicies
type SSHAuthorizationPoliciesList []*SSHAuthorizationPolicy

// Identity returns the identity of the objects in the list.
func (o SSHAuthorizationPoliciesList) Identity() elemental.Identity {

	return SSHAuthorizationPolicyIdentity
}

// Copy returns a pointer to a copy the SSHAuthorizationPoliciesList.
func (o SSHAuthorizationPoliciesList) Copy() elemental.Identifiables {

	copy := append(SSHAuthorizationPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SSHAuthorizationPoliciesList.
func (o SSHAuthorizationPoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SSHAuthorizationPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SSHAuthorizationPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SSHAuthorizationPoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SSHAuthorizationPoliciesList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToSparse returns the SSHAuthorizationPoliciesList converted to SparseSSHAuthorizationPoliciesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o SSHAuthorizationPoliciesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseSSHAuthorizationPoliciesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseSSHAuthorizationPolicy)
	}

	return out
}

// Version returns the version of the content.
func (o SSHAuthorizationPoliciesList) Version() int {

	return 1
}

// SSHAuthorizationPolicy represents the model of a sshauthorizationpolicy
type SSHAuthorizationPolicy struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Defines for how long the policy will be active according to the
	// `activeSchedule`.
	ActiveDuration string `json:"activeDuration" msgpack:"activeDuration" bson:"activeduration" mapstructure:"activeDuration,omitempty"`

	// Defines when the policy should be active using the cron notation.
	// The policy will be active for the given `activeDuration`.
	ActiveSchedule string `json:"activeSchedule" msgpack:"activeSchedule" bson:"activeschedule" mapstructure:"activeSchedule,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// If set, the SSH authorization will only be valid if the request comes from one
	// the declared subnets.
	AuthorizedSubnets []string `json:"authorizedSubnets" msgpack:"authorizedSubnets" bson:"-" mapstructure:"authorizedSubnets,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Defines if the property is disabled.
	Disabled bool `json:"disabled" msgpack:"disabled" bson:"disabled" mapstructure:"disabled,omitempty"`

	// If set the SSH authorization will be automatically deleted after the given time.
	ExpirationTime time.Time `json:"expirationTime" msgpack:"expirationTime" bson:"expirationtime" mapstructure:"expirationTime,omitempty"`

	// The list of permissions to apply to the OpenSSH certificate. You can check the list of
	// standard extensions at <https://github.com/openssh/openssh-portable/blob/38e83e4f219c752ebb1560633b73f06f0392018b/PROTOCOL.certkeys#L281>.
	Extensions []string `json:"extensions" msgpack:"extensions" bson:"-" mapstructure:"extensions,omitempty"`

	// Indicates that this is fallback policy. It will only be
	// applied if no other policies have been resolved. If the policy is also
	// propagated it will become a fallback for children namespaces.
	Fallback bool `json:"fallback" msgpack:"fallback" bson:"fallback" mapstructure:"fallback,omitempty"`

	// Specify a single command that the user can issue on the remote host. This can be useful
	// for issuing single-purpose certificates; ensuring that users stay in their home directories
	// (`internal-sftp`); and restricting users to a bash shell (`/bin/bash`), preventing them
	// from running arbitrary and unlogged commands such as `scp`, `rsync`, `-essh`, and `sftp`.
	// Refer to the [FreeBSD documentation](https://www.freebsd.org/cgi/man.cgi?sshd_config(5))
	// for more information.
	ForceCommand string `json:"forceCommand" msgpack:"forceCommand" bson:"-" mapstructure:"forceCommand,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Contains the tag expression identifying the enforcers on the hosts the `subject` is
	// allowed to access.
	Object [][]string `json:"object" msgpack:"object" bson:"-" mapstructure:"object,omitempty"`

	// On systems without the Aporeto enforcer, you must provide the name of the Linux user.
	// Otherwise, Aporeto will automatically populate this field and adding a value here is
	// optional and not used during the authorization. However, the value becomes a tag
	// associated with the SSH processing unit, which could be useful.
	Principals []string `json:"principals" msgpack:"principals" bson:"-" mapstructure:"principals,omitempty"`

	// Propagates the policy to all of its children.
	Propagate bool `json:"propagate" msgpack:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Contains the tag expression that identifies the user or group of users that should be
	// allowed to access the remote hosts. If the user authenticates against an OIDC provider,
	// these tags correspond to claims in the ID token.
	Subject [][]string `json:"subject" msgpack:"subject" bson:"-" mapstructure:"subject,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// Set the validity of the delivered SSH certificate.
	Validity string `json:"validity" msgpack:"validity" bson:"-" mapstructure:"validity,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSSHAuthorizationPolicy returns a new *SSHAuthorizationPolicy
func NewSSHAuthorizationPolicy() *SSHAuthorizationPolicy {

	return &SSHAuthorizationPolicy{
		ModelVersion:      1,
		Annotations:       map[string][]string{},
		AssociatedTags:    []string{},
		AuthorizedSubnets: []string{},
		Extensions:        []string{},
		Metadata:          []string{},
		NormalizedTags:    []string{},
		Object:            [][]string{},
		Principals:        []string{},
		Subject:           [][]string{},
		Validity:          "1h",
	}
}

// Identity returns the Identity of the object.
func (o *SSHAuthorizationPolicy) Identity() elemental.Identity {

	return SSHAuthorizationPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SSHAuthorizationPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SSHAuthorizationPolicy) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *SSHAuthorizationPolicy) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *SSHAuthorizationPolicy) BleveType() string {

	return "sshauthorizationpolicy"
}

// DefaultOrder returns the list of default ordering fields.
func (o *SSHAuthorizationPolicy) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// Doc returns the documentation for the object
func (o *SSHAuthorizationPolicy) Doc() string {

	return `An SSH authorization allows you to define the permissions for the owner
of a OpenSSH certificate issued by an Aporeto certificate authority. You can 
define if a user with some claims can connect to an ` + "`" + `sshd` + "`" + ` server managed by 
an instance of ` + "`" + `enforcerd` + "`" + ` according to its tags, what permissions he has and 
for how long delivered certificates are valid.`
}

func (o *SSHAuthorizationPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetActiveDuration returns the ActiveDuration of the receiver.
func (o *SSHAuthorizationPolicy) GetActiveDuration() string {

	return o.ActiveDuration
}

// SetActiveDuration sets the property ActiveDuration of the receiver using the given value.
func (o *SSHAuthorizationPolicy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *SSHAuthorizationPolicy) GetActiveSchedule() string {

	return o.ActiveSchedule
}

// SetActiveSchedule sets the property ActiveSchedule of the receiver using the given value.
func (o *SSHAuthorizationPolicy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = activeSchedule
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SSHAuthorizationPolicy) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *SSHAuthorizationPolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SSHAuthorizationPolicy) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *SSHAuthorizationPolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SSHAuthorizationPolicy) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *SSHAuthorizationPolicy) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SSHAuthorizationPolicy) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *SSHAuthorizationPolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *SSHAuthorizationPolicy) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *SSHAuthorizationPolicy) SetDescription(description string) {

	o.Description = description
}

// GetDisabled returns the Disabled of the receiver.
func (o *SSHAuthorizationPolicy) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the given value.
func (o *SSHAuthorizationPolicy) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetExpirationTime returns the ExpirationTime of the receiver.
func (o *SSHAuthorizationPolicy) GetExpirationTime() time.Time {

	return o.ExpirationTime
}

// SetExpirationTime sets the property ExpirationTime of the receiver using the given value.
func (o *SSHAuthorizationPolicy) SetExpirationTime(expirationTime time.Time) {

	o.ExpirationTime = expirationTime
}

// GetFallback returns the Fallback of the receiver.
func (o *SSHAuthorizationPolicy) GetFallback() bool {

	return o.Fallback
}

// SetFallback sets the property Fallback of the receiver using the given value.
func (o *SSHAuthorizationPolicy) SetFallback(fallback bool) {

	o.Fallback = fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *SSHAuthorizationPolicy) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *SSHAuthorizationPolicy) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *SSHAuthorizationPolicy) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *SSHAuthorizationPolicy) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SSHAuthorizationPolicy) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *SSHAuthorizationPolicy) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SSHAuthorizationPolicy) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *SSHAuthorizationPolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *SSHAuthorizationPolicy) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the given value.
func (o *SSHAuthorizationPolicy) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// GetProtected returns the Protected of the receiver.
func (o *SSHAuthorizationPolicy) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *SSHAuthorizationPolicy) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SSHAuthorizationPolicy) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *SSHAuthorizationPolicy) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SSHAuthorizationPolicy) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *SSHAuthorizationPolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *SSHAuthorizationPolicy) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseSSHAuthorizationPolicy{
			ID:                   &o.ID,
			ActiveDuration:       &o.ActiveDuration,
			ActiveSchedule:       &o.ActiveSchedule,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			AuthorizedSubnets:    &o.AuthorizedSubnets,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			Description:          &o.Description,
			Disabled:             &o.Disabled,
			ExpirationTime:       &o.ExpirationTime,
			Extensions:           &o.Extensions,
			Fallback:             &o.Fallback,
			ForceCommand:         &o.ForceCommand,
			Metadata:             &o.Metadata,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Object:               &o.Object,
			Principals:           &o.Principals,
			Propagate:            &o.Propagate,
			Protected:            &o.Protected,
			Subject:              &o.Subject,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
			Validity:             &o.Validity,
		}
	}

	sp := &SparseSSHAuthorizationPolicy{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "activeDuration":
			sp.ActiveDuration = &(o.ActiveDuration)
		case "activeSchedule":
			sp.ActiveSchedule = &(o.ActiveSchedule)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "authorizedSubnets":
			sp.AuthorizedSubnets = &(o.AuthorizedSubnets)
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
		case "extensions":
			sp.Extensions = &(o.Extensions)
		case "fallback":
			sp.Fallback = &(o.Fallback)
		case "forceCommand":
			sp.ForceCommand = &(o.ForceCommand)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "object":
			sp.Object = &(o.Object)
		case "principals":
			sp.Principals = &(o.Principals)
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
		case "validity":
			sp.Validity = &(o.Validity)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseSSHAuthorizationPolicy to the object.
func (o *SSHAuthorizationPolicy) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseSSHAuthorizationPolicy)
	if so.ID != nil {
		o.ID = *so.ID
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
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.AuthorizedSubnets != nil {
		o.AuthorizedSubnets = *so.AuthorizedSubnets
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
	if so.Extensions != nil {
		o.Extensions = *so.Extensions
	}
	if so.Fallback != nil {
		o.Fallback = *so.Fallback
	}
	if so.ForceCommand != nil {
		o.ForceCommand = *so.ForceCommand
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
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Object != nil {
		o.Object = *so.Object
	}
	if so.Principals != nil {
		o.Principals = *so.Principals
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
	if so.Validity != nil {
		o.Validity = *so.Validity
	}
}

// DeepCopy returns a deep copy if the SSHAuthorizationPolicy.
func (o *SSHAuthorizationPolicy) DeepCopy() *SSHAuthorizationPolicy {

	if o == nil {
		return nil
	}

	out := &SSHAuthorizationPolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SSHAuthorizationPolicy.
func (o *SSHAuthorizationPolicy) DeepCopyInto(out *SSHAuthorizationPolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SSHAuthorizationPolicy: %s", err))
	}

	*out = *target.(*SSHAuthorizationPolicy)
}

// Validate valides the current information stored into the structure.
func (o *SSHAuthorizationPolicy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidatePattern("activeDuration", o.ActiveDuration, `^[0-9]+[smh]$`, `must be a valid duration like <n>s or <n>s or <n>h`, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateOptionalNetworkList("authorizedSubnets", o.AuthorizedSubnets); err != nil {
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

	if err := ValidateTimeDuration("validity", o.Validity); err != nil {
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
func (*SSHAuthorizationPolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := SSHAuthorizationPolicyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return SSHAuthorizationPolicyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*SSHAuthorizationPolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return SSHAuthorizationPolicyAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *SSHAuthorizationPolicy) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "activeDuration":
		return o.ActiveDuration
	case "activeSchedule":
		return o.ActiveSchedule
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "authorizedSubnets":
		return o.AuthorizedSubnets
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
	case "extensions":
		return o.Extensions
	case "fallback":
		return o.Fallback
	case "forceCommand":
		return o.ForceCommand
	case "metadata":
		return o.Metadata
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "object":
		return o.Object
	case "principals":
		return o.Principals
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
	case "validity":
		return o.Validity
	}

	return nil
}

// SSHAuthorizationPolicyAttributesMap represents the map of attribute for SSHAuthorizationPolicy.
var SSHAuthorizationPolicyAttributesMap = map[string]elemental.AttributeSpecification{
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
	"AuthorizedSubnets": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuthorizedSubnets",
		Description: `If set, the SSH authorization will only be valid if the request comes from one
the declared subnets.`,
		Exposed: true,
		Name:    "authorizedSubnets",
		SubType: "string",
		Type:    "list",
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
	"ExpirationTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationTime",
		Description:    `If set the SSH authorization will be automatically deleted after the given time.`,
		Exposed:        true,
		Getter:         true,
		Name:           "expirationTime",
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Extensions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Extensions",
		Description: `The list of permissions to apply to the OpenSSH certificate. You can check the list of 
standard extensions at <https://github.com/openssh/openssh-portable/blob/38e83e4f219c752ebb1560633b73f06f0392018b/PROTOCOL.certkeys#L281>.`,
		Exposed: true,
		Name:    "extensions",
		SubType: "string",
		Type:    "list",
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
	"ForceCommand": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ForceCommand",
		Description: `Specify a single command that the user can issue on the remote host. This can be useful 
for issuing single-purpose certificates; ensuring that users stay in their home directories 
(` + "`" + `internal-sftp` + "`" + `); and restricting users to a bash shell (` + "`" + `/bin/bash` + "`" + `), preventing them 
from running arbitrary and unlogged commands such as ` + "`" + `scp` + "`" + `, ` + "`" + `rsync` + "`" + `, ` + "`" + `-essh` + "`" + `, and ` + "`" + `sftp` + "`" + `. 
Refer to the [FreeBSD documentation](https://www.freebsd.org/cgi/man.cgi?sshd_config(5)) 
for more information.`,
		Exposed: true,
		Name:    "forceCommand",
		Type:    "string",
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
		DefaultOrder:   true,
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
		DefaultOrder:   true,
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
		Description: `Contains the tag expression identifying the enforcers on the hosts the ` + "`" + `subject` + "`" + ` is 
allowed to access.`,
		Exposed:   true,
		Name:      "object",
		Orderable: true,
		SubType:   "[][]string",
		Type:      "external",
	},
	"Principals": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Principals",
		Description: `On systems without the Aporeto enforcer, you must provide the name of the Linux user. 
Otherwise, Aporeto will automatically populate this field and adding a value here is 
optional and not used during the authorization. However, the value becomes a tag 
associated with the SSH processing unit, which could be useful.`,
		Exposed: true,
		Name:    "principals",
		SubType: "string",
		Type:    "list",
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
		Description: `Contains the tag expression that identifies the user or group of users that should be 
allowed to access the remote hosts. If the user authenticates against an OIDC provider, 
these tags correspond to claims in the ID token.`,
		Exposed:   true,
		Name:      "subject",
		Orderable: true,
		SubType:   "[][]string",
		Type:      "external",
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
	"Validity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Validity",
		DefaultValue:   "1h",
		Description:    `Set the validity of the delivered SSH certificate.`,
		Exposed:        true,
		Name:           "validity",
		Type:           "string",
	},
}

// SSHAuthorizationPolicyLowerCaseAttributesMap represents the map of attribute for SSHAuthorizationPolicy.
var SSHAuthorizationPolicyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"authorizedsubnets": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuthorizedSubnets",
		Description: `If set, the SSH authorization will only be valid if the request comes from one
the declared subnets.`,
		Exposed: true,
		Name:    "authorizedSubnets",
		SubType: "string",
		Type:    "list",
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
	"expirationtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationTime",
		Description:    `If set the SSH authorization will be automatically deleted after the given time.`,
		Exposed:        true,
		Getter:         true,
		Name:           "expirationTime",
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"extensions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Extensions",
		Description: `The list of permissions to apply to the OpenSSH certificate. You can check the list of 
standard extensions at <https://github.com/openssh/openssh-portable/blob/38e83e4f219c752ebb1560633b73f06f0392018b/PROTOCOL.certkeys#L281>.`,
		Exposed: true,
		Name:    "extensions",
		SubType: "string",
		Type:    "list",
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
	"forcecommand": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ForceCommand",
		Description: `Specify a single command that the user can issue on the remote host. This can be useful 
for issuing single-purpose certificates; ensuring that users stay in their home directories 
(` + "`" + `internal-sftp` + "`" + `); and restricting users to a bash shell (` + "`" + `/bin/bash` + "`" + `), preventing them 
from running arbitrary and unlogged commands such as ` + "`" + `scp` + "`" + `, ` + "`" + `rsync` + "`" + `, ` + "`" + `-essh` + "`" + `, and ` + "`" + `sftp` + "`" + `. 
Refer to the [FreeBSD documentation](https://www.freebsd.org/cgi/man.cgi?sshd_config(5)) 
for more information.`,
		Exposed: true,
		Name:    "forceCommand",
		Type:    "string",
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
		DefaultOrder:   true,
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
		DefaultOrder:   true,
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
		Description: `Contains the tag expression identifying the enforcers on the hosts the ` + "`" + `subject` + "`" + ` is 
allowed to access.`,
		Exposed:   true,
		Name:      "object",
		Orderable: true,
		SubType:   "[][]string",
		Type:      "external",
	},
	"principals": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Principals",
		Description: `On systems without the Aporeto enforcer, you must provide the name of the Linux user. 
Otherwise, Aporeto will automatically populate this field and adding a value here is 
optional and not used during the authorization. However, the value becomes a tag 
associated with the SSH processing unit, which could be useful.`,
		Exposed: true,
		Name:    "principals",
		SubType: "string",
		Type:    "list",
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
		Description: `Contains the tag expression that identifies the user or group of users that should be 
allowed to access the remote hosts. If the user authenticates against an OIDC provider, 
these tags correspond to claims in the ID token.`,
		Exposed:   true,
		Name:      "subject",
		Orderable: true,
		SubType:   "[][]string",
		Type:      "external",
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
	"validity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Validity",
		DefaultValue:   "1h",
		Description:    `Set the validity of the delivered SSH certificate.`,
		Exposed:        true,
		Name:           "validity",
		Type:           "string",
	},
}

// SparseSSHAuthorizationPoliciesList represents a list of SparseSSHAuthorizationPolicies
type SparseSSHAuthorizationPoliciesList []*SparseSSHAuthorizationPolicy

// Identity returns the identity of the objects in the list.
func (o SparseSSHAuthorizationPoliciesList) Identity() elemental.Identity {

	return SSHAuthorizationPolicyIdentity
}

// Copy returns a pointer to a copy the SparseSSHAuthorizationPoliciesList.
func (o SparseSSHAuthorizationPoliciesList) Copy() elemental.Identifiables {

	copy := append(SparseSSHAuthorizationPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseSSHAuthorizationPoliciesList.
func (o SparseSSHAuthorizationPoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseSSHAuthorizationPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseSSHAuthorizationPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseSSHAuthorizationPoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseSSHAuthorizationPoliciesList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToPlain returns the SparseSSHAuthorizationPoliciesList converted to SSHAuthorizationPoliciesList.
func (o SparseSSHAuthorizationPoliciesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseSSHAuthorizationPoliciesList) Version() int {

	return 1
}

// SparseSSHAuthorizationPolicy represents the sparse version of a sshauthorizationpolicy.
type SparseSSHAuthorizationPolicy struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Defines for how long the policy will be active according to the
	// `activeSchedule`.
	ActiveDuration *string `json:"activeDuration,omitempty" msgpack:"activeDuration,omitempty" bson:"activeduration,omitempty" mapstructure:"activeDuration,omitempty"`

	// Defines when the policy should be active using the cron notation.
	// The policy will be active for the given `activeDuration`.
	ActiveSchedule *string `json:"activeSchedule,omitempty" msgpack:"activeSchedule,omitempty" bson:"activeschedule,omitempty" mapstructure:"activeSchedule,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// If set, the SSH authorization will only be valid if the request comes from one
	// the declared subnets.
	AuthorizedSubnets *[]string `json:"authorizedSubnets,omitempty" msgpack:"authorizedSubnets,omitempty" bson:"-" mapstructure:"authorizedSubnets,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Defines if the property is disabled.
	Disabled *bool `json:"disabled,omitempty" msgpack:"disabled,omitempty" bson:"disabled,omitempty" mapstructure:"disabled,omitempty"`

	// If set the SSH authorization will be automatically deleted after the given time.
	ExpirationTime *time.Time `json:"expirationTime,omitempty" msgpack:"expirationTime,omitempty" bson:"expirationtime,omitempty" mapstructure:"expirationTime,omitempty"`

	// The list of permissions to apply to the OpenSSH certificate. You can check the list of
	// standard extensions at <https://github.com/openssh/openssh-portable/blob/38e83e4f219c752ebb1560633b73f06f0392018b/PROTOCOL.certkeys#L281>.
	Extensions *[]string `json:"extensions,omitempty" msgpack:"extensions,omitempty" bson:"-" mapstructure:"extensions,omitempty"`

	// Indicates that this is fallback policy. It will only be
	// applied if no other policies have been resolved. If the policy is also
	// propagated it will become a fallback for children namespaces.
	Fallback *bool `json:"fallback,omitempty" msgpack:"fallback,omitempty" bson:"fallback,omitempty" mapstructure:"fallback,omitempty"`

	// Specify a single command that the user can issue on the remote host. This can be useful
	// for issuing single-purpose certificates; ensuring that users stay in their home directories
	// (`internal-sftp`); and restricting users to a bash shell (`/bin/bash`), preventing them
	// from running arbitrary and unlogged commands such as `scp`, `rsync`, `-essh`, and `sftp`.
	// Refer to the [FreeBSD documentation](https://www.freebsd.org/cgi/man.cgi?sshd_config(5))
	// for more information.
	ForceCommand *string `json:"forceCommand,omitempty" msgpack:"forceCommand,omitempty" bson:"-" mapstructure:"forceCommand,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Contains the tag expression identifying the enforcers on the hosts the `subject` is
	// allowed to access.
	Object *[][]string `json:"object,omitempty" msgpack:"object,omitempty" bson:"-" mapstructure:"object,omitempty"`

	// On systems without the Aporeto enforcer, you must provide the name of the Linux user.
	// Otherwise, Aporeto will automatically populate this field and adding a value here is
	// optional and not used during the authorization. However, the value becomes a tag
	// associated with the SSH processing unit, which could be useful.
	Principals *[]string `json:"principals,omitempty" msgpack:"principals,omitempty" bson:"-" mapstructure:"principals,omitempty"`

	// Propagates the policy to all of its children.
	Propagate *bool `json:"propagate,omitempty" msgpack:"propagate,omitempty" bson:"propagate,omitempty" mapstructure:"propagate,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Contains the tag expression that identifies the user or group of users that should be
	// allowed to access the remote hosts. If the user authenticates against an OIDC provider,
	// these tags correspond to claims in the ID token.
	Subject *[][]string `json:"subject,omitempty" msgpack:"subject,omitempty" bson:"-" mapstructure:"subject,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// Set the validity of the delivered SSH certificate.
	Validity *string `json:"validity,omitempty" msgpack:"validity,omitempty" bson:"-" mapstructure:"validity,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseSSHAuthorizationPolicy returns a new  SparseSSHAuthorizationPolicy.
func NewSparseSSHAuthorizationPolicy() *SparseSSHAuthorizationPolicy {
	return &SparseSSHAuthorizationPolicy{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseSSHAuthorizationPolicy) Identity() elemental.Identity {

	return SSHAuthorizationPolicyIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseSSHAuthorizationPolicy) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseSSHAuthorizationPolicy) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseSSHAuthorizationPolicy) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseSSHAuthorizationPolicy) ToPlain() elemental.PlainIdentifiable {

	out := NewSSHAuthorizationPolicy()
	if o.ID != nil {
		out.ID = *o.ID
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
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.AuthorizedSubnets != nil {
		out.AuthorizedSubnets = *o.AuthorizedSubnets
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
	if o.Extensions != nil {
		out.Extensions = *o.Extensions
	}
	if o.Fallback != nil {
		out.Fallback = *o.Fallback
	}
	if o.ForceCommand != nil {
		out.ForceCommand = *o.ForceCommand
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
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Object != nil {
		out.Object = *o.Object
	}
	if o.Principals != nil {
		out.Principals = *o.Principals
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
	if o.Validity != nil {
		out.Validity = *o.Validity
	}

	return out
}

// GetActiveDuration returns the ActiveDuration of the receiver.
func (o *SparseSSHAuthorizationPolicy) GetActiveDuration() string {

	return *o.ActiveDuration
}

// SetActiveDuration sets the property ActiveDuration of the receiver using the address of the given value.
func (o *SparseSSHAuthorizationPolicy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = &activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *SparseSSHAuthorizationPolicy) GetActiveSchedule() string {

	return *o.ActiveSchedule
}

// SetActiveSchedule sets the property ActiveSchedule of the receiver using the address of the given value.
func (o *SparseSSHAuthorizationPolicy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = &activeSchedule
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseSSHAuthorizationPolicy) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseSSHAuthorizationPolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseSSHAuthorizationPolicy) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseSSHAuthorizationPolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseSSHAuthorizationPolicy) GetCreateIdempotencyKey() string {

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseSSHAuthorizationPolicy) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseSSHAuthorizationPolicy) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseSSHAuthorizationPolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseSSHAuthorizationPolicy) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseSSHAuthorizationPolicy) SetDescription(description string) {

	o.Description = &description
}

// GetDisabled returns the Disabled of the receiver.
func (o *SparseSSHAuthorizationPolicy) GetDisabled() bool {

	return *o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the address of the given value.
func (o *SparseSSHAuthorizationPolicy) SetDisabled(disabled bool) {

	o.Disabled = &disabled
}

// GetExpirationTime returns the ExpirationTime of the receiver.
func (o *SparseSSHAuthorizationPolicy) GetExpirationTime() time.Time {

	return *o.ExpirationTime
}

// SetExpirationTime sets the property ExpirationTime of the receiver using the address of the given value.
func (o *SparseSSHAuthorizationPolicy) SetExpirationTime(expirationTime time.Time) {

	o.ExpirationTime = &expirationTime
}

// GetFallback returns the Fallback of the receiver.
func (o *SparseSSHAuthorizationPolicy) GetFallback() bool {

	return *o.Fallback
}

// SetFallback sets the property Fallback of the receiver using the address of the given value.
func (o *SparseSSHAuthorizationPolicy) SetFallback(fallback bool) {

	o.Fallback = &fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseSSHAuthorizationPolicy) GetMetadata() []string {

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseSSHAuthorizationPolicy) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetName returns the Name of the receiver.
func (o *SparseSSHAuthorizationPolicy) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseSSHAuthorizationPolicy) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseSSHAuthorizationPolicy) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseSSHAuthorizationPolicy) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseSSHAuthorizationPolicy) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseSSHAuthorizationPolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *SparseSSHAuthorizationPolicy) GetPropagate() bool {

	return *o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the address of the given value.
func (o *SparseSSHAuthorizationPolicy) SetPropagate(propagate bool) {

	o.Propagate = &propagate
}

// GetProtected returns the Protected of the receiver.
func (o *SparseSSHAuthorizationPolicy) GetProtected() bool {

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseSSHAuthorizationPolicy) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseSSHAuthorizationPolicy) GetUpdateIdempotencyKey() string {

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseSSHAuthorizationPolicy) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseSSHAuthorizationPolicy) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseSSHAuthorizationPolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// DeepCopy returns a deep copy if the SparseSSHAuthorizationPolicy.
func (o *SparseSSHAuthorizationPolicy) DeepCopy() *SparseSSHAuthorizationPolicy {

	if o == nil {
		return nil
	}

	out := &SparseSSHAuthorizationPolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseSSHAuthorizationPolicy.
func (o *SparseSSHAuthorizationPolicy) DeepCopyInto(out *SparseSSHAuthorizationPolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseSSHAuthorizationPolicy: %s", err))
	}

	*out = *target.(*SparseSSHAuthorizationPolicy)
}
