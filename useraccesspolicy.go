package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// UserAccessPolicyIdentity represents the Identity of the object.
var UserAccessPolicyIdentity = elemental.Identity{
	Name:     "useraccesspolicy",
	Category: "useraccesspolicies",
	Package:  "squall",
	Private:  false,
}

// UserAccessPoliciesList represents a list of UserAccessPolicies
type UserAccessPoliciesList []*UserAccessPolicy

// Identity returns the identity of the objects in the list.
func (o UserAccessPoliciesList) Identity() elemental.Identity {

	return UserAccessPolicyIdentity
}

// Copy returns a pointer to a copy the UserAccessPoliciesList.
func (o UserAccessPoliciesList) Copy() elemental.Identifiables {

	copy := append(UserAccessPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the UserAccessPoliciesList.
func (o UserAccessPoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(UserAccessPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*UserAccessPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o UserAccessPoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o UserAccessPoliciesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the UserAccessPoliciesList converted to SparseUserAccessPoliciesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o UserAccessPoliciesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseUserAccessPoliciesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseUserAccessPolicy)
	}

	return out
}

// Version returns the version of the content.
func (o UserAccessPoliciesList) Version() int {

	return 1
}

// UserAccessPolicy represents the model of a useraccesspolicy
type UserAccessPolicy struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Defines for how long the policy will be active according to the
	// `activeSchedule`.
	ActiveDuration string `json:"activeDuration" msgpack:"activeDuration" bson:"activeduration" mapstructure:"activeDuration,omitempty"`

	// Defines when the policy should be active using the cron notation.
	// The policy will be active for the given `activeDuration`.
	ActiveSchedule string `json:"activeSchedule" msgpack:"activeSchedule" bson:"activeschedule" mapstructure:"activeSchedule,omitempty"`

	// Indicates the list of user who can use `sudo` commands.
	AllowedSudoUsers []string `json:"allowedSudoUsers" msgpack:"allowedSudoUsers" bson:"-" mapstructure:"allowedSudoUsers,omitempty"`

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

	// If set the policy will be automatically deleted after the given time.
	ExpirationTime time.Time `json:"expirationTime" msgpack:"expirationTime" bson:"expirationtime" mapstructure:"expirationTime,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Contains the tag expression matching the enforcers the subject is allowed
	// to connect to.
	Object [][]string `json:"object" msgpack:"object" bson:"-" mapstructure:"object,omitempty"`

	// Propagates the policy to all of its children.
	Propagate bool `json:"propagate" msgpack:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Contains the tag expression the tags need to match for the policy to
	// apply.
	Subject [][]string `json:"subject" msgpack:"subject" bson:"-" mapstructure:"subject,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewUserAccessPolicy returns a new *UserAccessPolicy
func NewUserAccessPolicy() *UserAccessPolicy {

	return &UserAccessPolicy{
		ModelVersion:     1,
		AllowedSudoUsers: []string{},
		Annotations:      map[string][]string{},
		AssociatedTags:   []string{},
		Metadata:         []string{},
		NormalizedTags:   []string{},
		Object:           [][]string{},
		Subject:          [][]string{},
	}
}

// Identity returns the Identity of the object.
func (o *UserAccessPolicy) Identity() elemental.Identity {

	return UserAccessPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *UserAccessPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *UserAccessPolicy) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *UserAccessPolicy) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesUserAccessPolicy{}

	s.ActiveDuration = o.ActiveDuration
	s.ActiveSchedule = o.ActiveSchedule
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.Description = o.Description
	s.Disabled = o.Disabled
	s.ExpirationTime = o.ExpirationTime
	s.Metadata = o.Metadata
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.Propagate = o.Propagate
	s.Protected = o.Protected
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *UserAccessPolicy) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesUserAccessPolicy{}
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
	o.Metadata = s.Metadata
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.Propagate = s.Propagate
	o.Protected = s.Protected
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime

	return nil
}

// Version returns the hardcoded version of the model.
func (o *UserAccessPolicy) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *UserAccessPolicy) BleveType() string {

	return "useraccesspolicy"
}

// DefaultOrder returns the list of default ordering fields.
func (o *UserAccessPolicy) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *UserAccessPolicy) Doc() string {

	return `The enforcer policy that controls user access.`
}

func (o *UserAccessPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetActiveDuration returns the ActiveDuration of the receiver.
func (o *UserAccessPolicy) GetActiveDuration() string {

	return o.ActiveDuration
}

// SetActiveDuration sets the property ActiveDuration of the receiver using the given value.
func (o *UserAccessPolicy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *UserAccessPolicy) GetActiveSchedule() string {

	return o.ActiveSchedule
}

// SetActiveSchedule sets the property ActiveSchedule of the receiver using the given value.
func (o *UserAccessPolicy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = activeSchedule
}

// GetAnnotations returns the Annotations of the receiver.
func (o *UserAccessPolicy) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *UserAccessPolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *UserAccessPolicy) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *UserAccessPolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *UserAccessPolicy) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *UserAccessPolicy) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *UserAccessPolicy) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *UserAccessPolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *UserAccessPolicy) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *UserAccessPolicy) SetDescription(description string) {

	o.Description = description
}

// GetDisabled returns the Disabled of the receiver.
func (o *UserAccessPolicy) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the given value.
func (o *UserAccessPolicy) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetExpirationTime returns the ExpirationTime of the receiver.
func (o *UserAccessPolicy) GetExpirationTime() time.Time {

	return o.ExpirationTime
}

// SetExpirationTime sets the property ExpirationTime of the receiver using the given value.
func (o *UserAccessPolicy) SetExpirationTime(expirationTime time.Time) {

	o.ExpirationTime = expirationTime
}

// GetMetadata returns the Metadata of the receiver.
func (o *UserAccessPolicy) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *UserAccessPolicy) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *UserAccessPolicy) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *UserAccessPolicy) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *UserAccessPolicy) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *UserAccessPolicy) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *UserAccessPolicy) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *UserAccessPolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *UserAccessPolicy) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the given value.
func (o *UserAccessPolicy) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// GetProtected returns the Protected of the receiver.
func (o *UserAccessPolicy) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *UserAccessPolicy) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *UserAccessPolicy) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *UserAccessPolicy) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *UserAccessPolicy) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *UserAccessPolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *UserAccessPolicy) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseUserAccessPolicy{
			ID:                   &o.ID,
			ActiveDuration:       &o.ActiveDuration,
			ActiveSchedule:       &o.ActiveSchedule,
			AllowedSudoUsers:     &o.AllowedSudoUsers,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			Description:          &o.Description,
			Disabled:             &o.Disabled,
			ExpirationTime:       &o.ExpirationTime,
			Metadata:             &o.Metadata,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Object:               &o.Object,
			Propagate:            &o.Propagate,
			Protected:            &o.Protected,
			Subject:              &o.Subject,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
		}
	}

	sp := &SparseUserAccessPolicy{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "activeDuration":
			sp.ActiveDuration = &(o.ActiveDuration)
		case "activeSchedule":
			sp.ActiveSchedule = &(o.ActiveSchedule)
		case "allowedSudoUsers":
			sp.AllowedSudoUsers = &(o.AllowedSudoUsers)
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

// Patch apply the non nil value of a *SparseUserAccessPolicy to the object.
func (o *UserAccessPolicy) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseUserAccessPolicy)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.ActiveDuration != nil {
		o.ActiveDuration = *so.ActiveDuration
	}
	if so.ActiveSchedule != nil {
		o.ActiveSchedule = *so.ActiveSchedule
	}
	if so.AllowedSudoUsers != nil {
		o.AllowedSudoUsers = *so.AllowedSudoUsers
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

// DeepCopy returns a deep copy if the UserAccessPolicy.
func (o *UserAccessPolicy) DeepCopy() *UserAccessPolicy {

	if o == nil {
		return nil
	}

	out := &UserAccessPolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *UserAccessPolicy.
func (o *UserAccessPolicy) DeepCopyInto(out *UserAccessPolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy UserAccessPolicy: %s", err))
	}

	*out = *target.(*UserAccessPolicy)
}

// Validate valides the current information stored into the structure.
func (o *UserAccessPolicy) Validate() error {

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

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*UserAccessPolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := UserAccessPolicyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return UserAccessPolicyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*UserAccessPolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return UserAccessPolicyAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *UserAccessPolicy) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "activeDuration":
		return o.ActiveDuration
	case "activeSchedule":
		return o.ActiveSchedule
	case "allowedSudoUsers":
		return o.AllowedSudoUsers
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

// UserAccessPolicyAttributesMap represents the map of attribute for UserAccessPolicy.
var UserAccessPolicyAttributesMap = map[string]elemental.AttributeSpecification{
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
		Type:           "string",
	},
	"ActiveDuration": {
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
	"ActiveSchedule": {
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
	"AllowedSudoUsers": {
		AllowedChoices: []string{},
		ConvertedName:  "AllowedSudoUsers",
		Description:    `Indicates the list of user who can use ` + "`" + `sudo` + "`" + ` commands.`,
		Exposed:        true,
		Name:           "allowedSudoUsers",
		SubType:        "string",
		Type:           "list",
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
	"Disabled": {
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
	"ExpirationTime": {
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
	"Object": {
		AllowedChoices: []string{},
		ConvertedName:  "Object",
		Description: `Contains the tag expression matching the enforcers the subject is allowed
to connect to.`,
		Exposed:   true,
		Name:      "object",
		Orderable: true,
		SubType:   "[][]string",
		Type:      "external",
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
	"Subject": {
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description: `Contains the tag expression the tags need to match for the policy to
apply.`,
		Exposed:   true,
		Name:      "subject",
		Orderable: true,
		SubType:   "[][]string",
		Type:      "external",
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
}

// UserAccessPolicyLowerCaseAttributesMap represents the map of attribute for UserAccessPolicy.
var UserAccessPolicyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		Type:           "string",
	},
	"activeduration": {
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
	"activeschedule": {
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
	"allowedsudousers": {
		AllowedChoices: []string{},
		ConvertedName:  "AllowedSudoUsers",
		Description:    `Indicates the list of user who can use ` + "`" + `sudo` + "`" + ` commands.`,
		Exposed:        true,
		Name:           "allowedSudoUsers",
		SubType:        "string",
		Type:           "list",
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
	"disabled": {
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
	"expirationtime": {
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
	"object": {
		AllowedChoices: []string{},
		ConvertedName:  "Object",
		Description: `Contains the tag expression matching the enforcers the subject is allowed
to connect to.`,
		Exposed:   true,
		Name:      "object",
		Orderable: true,
		SubType:   "[][]string",
		Type:      "external",
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
	"subject": {
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description: `Contains the tag expression the tags need to match for the policy to
apply.`,
		Exposed:   true,
		Name:      "subject",
		Orderable: true,
		SubType:   "[][]string",
		Type:      "external",
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
}

// SparseUserAccessPoliciesList represents a list of SparseUserAccessPolicies
type SparseUserAccessPoliciesList []*SparseUserAccessPolicy

// Identity returns the identity of the objects in the list.
func (o SparseUserAccessPoliciesList) Identity() elemental.Identity {

	return UserAccessPolicyIdentity
}

// Copy returns a pointer to a copy the SparseUserAccessPoliciesList.
func (o SparseUserAccessPoliciesList) Copy() elemental.Identifiables {

	copy := append(SparseUserAccessPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseUserAccessPoliciesList.
func (o SparseUserAccessPoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseUserAccessPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseUserAccessPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseUserAccessPoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseUserAccessPoliciesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseUserAccessPoliciesList converted to UserAccessPoliciesList.
func (o SparseUserAccessPoliciesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseUserAccessPoliciesList) Version() int {

	return 1
}

// SparseUserAccessPolicy represents the sparse version of a useraccesspolicy.
type SparseUserAccessPolicy struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Defines for how long the policy will be active according to the
	// `activeSchedule`.
	ActiveDuration *string `json:"activeDuration,omitempty" msgpack:"activeDuration,omitempty" bson:"activeduration,omitempty" mapstructure:"activeDuration,omitempty"`

	// Defines when the policy should be active using the cron notation.
	// The policy will be active for the given `activeDuration`.
	ActiveSchedule *string `json:"activeSchedule,omitempty" msgpack:"activeSchedule,omitempty" bson:"activeschedule,omitempty" mapstructure:"activeSchedule,omitempty"`

	// Indicates the list of user who can use `sudo` commands.
	AllowedSudoUsers *[]string `json:"allowedSudoUsers,omitempty" msgpack:"allowedSudoUsers,omitempty" bson:"-" mapstructure:"allowedSudoUsers,omitempty"`

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

	// If set the policy will be automatically deleted after the given time.
	ExpirationTime *time.Time `json:"expirationTime,omitempty" msgpack:"expirationTime,omitempty" bson:"expirationtime,omitempty" mapstructure:"expirationTime,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Contains the tag expression matching the enforcers the subject is allowed
	// to connect to.
	Object *[][]string `json:"object,omitempty" msgpack:"object,omitempty" bson:"-" mapstructure:"object,omitempty"`

	// Propagates the policy to all of its children.
	Propagate *bool `json:"propagate,omitempty" msgpack:"propagate,omitempty" bson:"propagate,omitempty" mapstructure:"propagate,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Contains the tag expression the tags need to match for the policy to
	// apply.
	Subject *[][]string `json:"subject,omitempty" msgpack:"subject,omitempty" bson:"-" mapstructure:"subject,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseUserAccessPolicy returns a new  SparseUserAccessPolicy.
func NewSparseUserAccessPolicy() *SparseUserAccessPolicy {
	return &SparseUserAccessPolicy{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseUserAccessPolicy) Identity() elemental.Identity {

	return UserAccessPolicyIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseUserAccessPolicy) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseUserAccessPolicy) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseUserAccessPolicy) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseUserAccessPolicy{}

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
	if o.Metadata != nil {
		s.Metadata = o.Metadata
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
func (o *SparseUserAccessPolicy) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseUserAccessPolicy{}
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
	if s.Metadata != nil {
		o.Metadata = s.Metadata
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
	if s.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	}
	if s.UpdateTime != nil {
		o.UpdateTime = s.UpdateTime
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseUserAccessPolicy) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseUserAccessPolicy) ToPlain() elemental.PlainIdentifiable {

	out := NewUserAccessPolicy()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.ActiveDuration != nil {
		out.ActiveDuration = *o.ActiveDuration
	}
	if o.ActiveSchedule != nil {
		out.ActiveSchedule = *o.ActiveSchedule
	}
	if o.AllowedSudoUsers != nil {
		out.AllowedSudoUsers = *o.AllowedSudoUsers
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
func (o *SparseUserAccessPolicy) GetActiveDuration() (out string) {

	if o.ActiveDuration == nil {
		return
	}

	return *o.ActiveDuration
}

// SetActiveDuration sets the property ActiveDuration of the receiver using the address of the given value.
func (o *SparseUserAccessPolicy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = &activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *SparseUserAccessPolicy) GetActiveSchedule() (out string) {

	if o.ActiveSchedule == nil {
		return
	}

	return *o.ActiveSchedule
}

// SetActiveSchedule sets the property ActiveSchedule of the receiver using the address of the given value.
func (o *SparseUserAccessPolicy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = &activeSchedule
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseUserAccessPolicy) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseUserAccessPolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseUserAccessPolicy) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseUserAccessPolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseUserAccessPolicy) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseUserAccessPolicy) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseUserAccessPolicy) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseUserAccessPolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseUserAccessPolicy) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseUserAccessPolicy) SetDescription(description string) {

	o.Description = &description
}

// GetDisabled returns the Disabled of the receiver.
func (o *SparseUserAccessPolicy) GetDisabled() (out bool) {

	if o.Disabled == nil {
		return
	}

	return *o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the address of the given value.
func (o *SparseUserAccessPolicy) SetDisabled(disabled bool) {

	o.Disabled = &disabled
}

// GetExpirationTime returns the ExpirationTime of the receiver.
func (o *SparseUserAccessPolicy) GetExpirationTime() (out time.Time) {

	if o.ExpirationTime == nil {
		return
	}

	return *o.ExpirationTime
}

// SetExpirationTime sets the property ExpirationTime of the receiver using the address of the given value.
func (o *SparseUserAccessPolicy) SetExpirationTime(expirationTime time.Time) {

	o.ExpirationTime = &expirationTime
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseUserAccessPolicy) GetMetadata() (out []string) {

	if o.Metadata == nil {
		return
	}

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseUserAccessPolicy) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetName returns the Name of the receiver.
func (o *SparseUserAccessPolicy) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseUserAccessPolicy) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseUserAccessPolicy) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseUserAccessPolicy) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseUserAccessPolicy) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseUserAccessPolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *SparseUserAccessPolicy) GetPropagate() (out bool) {

	if o.Propagate == nil {
		return
	}

	return *o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the address of the given value.
func (o *SparseUserAccessPolicy) SetPropagate(propagate bool) {

	o.Propagate = &propagate
}

// GetProtected returns the Protected of the receiver.
func (o *SparseUserAccessPolicy) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseUserAccessPolicy) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseUserAccessPolicy) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseUserAccessPolicy) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseUserAccessPolicy) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseUserAccessPolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// DeepCopy returns a deep copy if the SparseUserAccessPolicy.
func (o *SparseUserAccessPolicy) DeepCopy() *SparseUserAccessPolicy {

	if o == nil {
		return nil
	}

	out := &SparseUserAccessPolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseUserAccessPolicy.
func (o *SparseUserAccessPolicy) DeepCopyInto(out *SparseUserAccessPolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseUserAccessPolicy: %s", err))
	}

	*out = *target.(*SparseUserAccessPolicy)
}

type mongoAttributesUserAccessPolicy struct {
	ActiveDuration       string              `bson:"activeduration"`
	ActiveSchedule       string              `bson:"activeschedule"`
	Annotations          map[string][]string `bson:"annotations"`
	AssociatedTags       []string            `bson:"associatedtags"`
	CreateIdempotencyKey string              `bson:"createidempotencykey"`
	CreateTime           time.Time           `bson:"createtime"`
	Description          string              `bson:"description"`
	Disabled             bool                `bson:"disabled"`
	ExpirationTime       time.Time           `bson:"expirationtime"`
	Metadata             []string            `bson:"metadata"`
	Name                 string              `bson:"name"`
	Namespace            string              `bson:"namespace"`
	NormalizedTags       []string            `bson:"normalizedtags"`
	Propagate            bool                `bson:"propagate"`
	Protected            bool                `bson:"protected"`
	UpdateIdempotencyKey string              `bson:"updateidempotencykey"`
	UpdateTime           time.Time           `bson:"updatetime"`
}
type mongoAttributesSparseUserAccessPolicy struct {
	ActiveDuration       *string              `bson:"activeduration,omitempty"`
	ActiveSchedule       *string              `bson:"activeschedule,omitempty"`
	Annotations          *map[string][]string `bson:"annotations,omitempty"`
	AssociatedTags       *[]string            `bson:"associatedtags,omitempty"`
	CreateIdempotencyKey *string              `bson:"createidempotencykey,omitempty"`
	CreateTime           *time.Time           `bson:"createtime,omitempty"`
	Description          *string              `bson:"description,omitempty"`
	Disabled             *bool                `bson:"disabled,omitempty"`
	ExpirationTime       *time.Time           `bson:"expirationtime,omitempty"`
	Metadata             *[]string            `bson:"metadata,omitempty"`
	Name                 *string              `bson:"name,omitempty"`
	Namespace            *string              `bson:"namespace,omitempty"`
	NormalizedTags       *[]string            `bson:"normalizedtags,omitempty"`
	Propagate            *bool                `bson:"propagate,omitempty"`
	Protected            *bool                `bson:"protected,omitempty"`
	UpdateIdempotencyKey *string              `bson:"updateidempotencykey,omitempty"`
	UpdateTime           *time.Time           `bson:"updatetime,omitempty"`
}
