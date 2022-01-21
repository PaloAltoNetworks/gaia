package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// RecipeDeploymentModeValue represents the possible values for attribute "deploymentMode".
type RecipeDeploymentModeValue string

const (
	// RecipeDeploymentModeNamespaceUnique represents the value NamespaceUnique.
	RecipeDeploymentModeNamespaceUnique RecipeDeploymentModeValue = "NamespaceUnique"

	// RecipeDeploymentModeUnique represents the value Unique.
	RecipeDeploymentModeUnique RecipeDeploymentModeValue = "Unique"

	// RecipeDeploymentModeUnrestricted represents the value Unrestricted.
	RecipeDeploymentModeUnrestricted RecipeDeploymentModeValue = "Unrestricted"
)

// RecipeIdentity represents the Identity of the object.
var RecipeIdentity = elemental.Identity{
	Name:     "recipe",
	Category: "recipes",
	Package:  "ignis",
	Private:  false,
}

// RecipesList represents a list of Recipes
type RecipesList []*Recipe

// Identity returns the identity of the objects in the list.
func (o RecipesList) Identity() elemental.Identity {

	return RecipeIdentity
}

// Copy returns a pointer to a copy the RecipesList.
func (o RecipesList) Copy() elemental.Identifiables {

	copy := append(RecipesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the RecipesList.
func (o RecipesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(RecipesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Recipe))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o RecipesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o RecipesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the RecipesList converted to SparseRecipesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o RecipesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseRecipesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseRecipe)
	}

	return out
}

// Version returns the version of the content.
func (o RecipesList) Version() int {

	return 1
}

// Recipe represents the model of a recipe
type Recipe struct {
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

	// Defines the deployment mode of the recipe. If `Unrestricted`, the recipe can be
	// deployed multiple times in the current namespace and below. If `Unique`,only one
	// deployment is allowed in the current namespace and its child namespaces. If
	// `NamespaceUnique`, only one deployment is allowed in the current namespace.
	DeploymentMode RecipeDeploymentModeValue `json:"deploymentMode" msgpack:"deploymentMode" bson:"deploymentmode" mapstructure:"deploymentMode,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Contains a base64-encoded image for the recipe.
	Icon string `json:"icon" msgpack:"icon" bson:"icon" mapstructure:"icon,omitempty"`

	// The unique key of the recipe.
	Key string `json:"key" msgpack:"key" bson:"key" mapstructure:"key,omitempty"`

	// Defines the recipe.
	Label string `json:"label" msgpack:"label" bson:"label" mapstructure:"label,omitempty"`

	// Provides a long description of the recipe.
	LongDescription string `json:"longDescription" msgpack:"longDescription" bson:"longdescription" mapstructure:"longDescription,omitempty"`

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

	// Options of the recipe.
	Options *RecipeOptions `json:"options" msgpack:"options" bson:"options" mapstructure:"options,omitempty"`

	// Propagates the policy to all of its children.
	Propagate bool `json:"propagate" msgpack:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Contains all the steps with parameters to follow for the recipe.
	Steps []*UIStep `json:"steps" msgpack:"steps" bson:"steps" mapstructure:"steps,omitempty"`

	// A string message presented upon success (optional).
	SuccessfullMessage string `json:"successfullMessage" msgpack:"successfullMessage" bson:"successfullmessage" mapstructure:"successfullMessage,omitempty"`

	// Contains the list of identities the recipes will try to create.
	TargetIdentities []string `json:"targetIdentities" msgpack:"targetIdentities" bson:"targetidentities" mapstructure:"targetIdentities,omitempty"`

	// Template of the recipe to import.
	Template string `json:"template" msgpack:"template" bson:"template" mapstructure:"template,omitempty"`

	// A hash of the template.
	TemplateHash string `json:"templateHash" msgpack:"templateHash" bson:"templatehash" mapstructure:"templateHash,omitempty"`

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

// NewRecipe returns a new *Recipe
func NewRecipe() *Recipe {

	return &Recipe{
		ModelVersion:     1,
		Annotations:      map[string][]string{},
		AssociatedTags:   []string{},
		DeploymentMode:   RecipeDeploymentModeUnrestricted,
		Steps:            []*UIStep{},
		MigrationsLog:    map[string]string{},
		TargetIdentities: []string{},
		Options:          NewRecipeOptions(),
		Label:            "magicpanda",
		Metadata:         []string{},
		NormalizedTags:   []string{},
	}
}

// Identity returns the Identity of the object.
func (o *Recipe) Identity() elemental.Identity {

	return RecipeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Recipe) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Recipe) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Recipe) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesRecipe{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.DeploymentMode = o.DeploymentMode
	s.Description = o.Description
	s.Icon = o.Icon
	s.Key = o.Key
	s.Label = o.Label
	s.LongDescription = o.LongDescription
	s.Metadata = o.Metadata
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.Options = o.Options
	s.Propagate = o.Propagate
	s.Protected = o.Protected
	s.Steps = o.Steps
	s.SuccessfullMessage = o.SuccessfullMessage
	s.TargetIdentities = o.TargetIdentities
	s.Template = o.Template
	s.TemplateHash = o.TemplateHash
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Recipe) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesRecipe{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.DeploymentMode = s.DeploymentMode
	o.Description = s.Description
	o.Icon = s.Icon
	o.Key = s.Key
	o.Label = s.Label
	o.LongDescription = s.LongDescription
	o.Metadata = s.Metadata
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.Options = s.Options
	o.Propagate = s.Propagate
	o.Protected = s.Protected
	o.Steps = s.Steps
	o.SuccessfullMessage = s.SuccessfullMessage
	o.TargetIdentities = s.TargetIdentities
	o.Template = s.Template
	o.TemplateHash = s.TemplateHash
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Recipe) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Recipe) BleveType() string {

	return "recipe"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Recipe) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *Recipe) Doc() string {

	return `Defines a list of steps that make up a workflow.`
}

func (o *Recipe) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *Recipe) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *Recipe) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *Recipe) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *Recipe) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *Recipe) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *Recipe) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Recipe) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *Recipe) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *Recipe) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *Recipe) SetDescription(description string) {

	o.Description = description
}

// GetMetadata returns the Metadata of the receiver.
func (o *Recipe) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *Recipe) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *Recipe) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *Recipe) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *Recipe) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *Recipe) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *Recipe) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *Recipe) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *Recipe) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *Recipe) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *Recipe) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the given value.
func (o *Recipe) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// GetProtected returns the Protected of the receiver.
func (o *Recipe) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *Recipe) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *Recipe) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *Recipe) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *Recipe) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *Recipe) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *Recipe) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *Recipe) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *Recipe) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *Recipe) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Recipe) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseRecipe{
			ID:                   &o.ID,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			DeploymentMode:       &o.DeploymentMode,
			Description:          &o.Description,
			Icon:                 &o.Icon,
			Key:                  &o.Key,
			Label:                &o.Label,
			LongDescription:      &o.LongDescription,
			Metadata:             &o.Metadata,
			MigrationsLog:        &o.MigrationsLog,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Options:              o.Options,
			Propagate:            &o.Propagate,
			Protected:            &o.Protected,
			Steps:                &o.Steps,
			SuccessfullMessage:   &o.SuccessfullMessage,
			TargetIdentities:     &o.TargetIdentities,
			Template:             &o.Template,
			TemplateHash:         &o.TemplateHash,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseRecipe{}
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
		case "deploymentMode":
			sp.DeploymentMode = &(o.DeploymentMode)
		case "description":
			sp.Description = &(o.Description)
		case "icon":
			sp.Icon = &(o.Icon)
		case "key":
			sp.Key = &(o.Key)
		case "label":
			sp.Label = &(o.Label)
		case "longDescription":
			sp.LongDescription = &(o.LongDescription)
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
		case "options":
			sp.Options = o.Options
		case "propagate":
			sp.Propagate = &(o.Propagate)
		case "protected":
			sp.Protected = &(o.Protected)
		case "steps":
			sp.Steps = &(o.Steps)
		case "successfullMessage":
			sp.SuccessfullMessage = &(o.SuccessfullMessage)
		case "targetIdentities":
			sp.TargetIdentities = &(o.TargetIdentities)
		case "template":
			sp.Template = &(o.Template)
		case "templateHash":
			sp.TemplateHash = &(o.TemplateHash)
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

// Patch apply the non nil value of a *SparseRecipe to the object.
func (o *Recipe) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseRecipe)
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
	if so.DeploymentMode != nil {
		o.DeploymentMode = *so.DeploymentMode
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Icon != nil {
		o.Icon = *so.Icon
	}
	if so.Key != nil {
		o.Key = *so.Key
	}
	if so.Label != nil {
		o.Label = *so.Label
	}
	if so.LongDescription != nil {
		o.LongDescription = *so.LongDescription
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
	if so.Options != nil {
		o.Options = so.Options
	}
	if so.Propagate != nil {
		o.Propagate = *so.Propagate
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Steps != nil {
		o.Steps = *so.Steps
	}
	if so.SuccessfullMessage != nil {
		o.SuccessfullMessage = *so.SuccessfullMessage
	}
	if so.TargetIdentities != nil {
		o.TargetIdentities = *so.TargetIdentities
	}
	if so.Template != nil {
		o.Template = *so.Template
	}
	if so.TemplateHash != nil {
		o.TemplateHash = *so.TemplateHash
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

// DeepCopy returns a deep copy if the Recipe.
func (o *Recipe) DeepCopy() *Recipe {

	if o == nil {
		return nil
	}

	out := &Recipe{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Recipe.
func (o *Recipe) DeepCopyInto(out *Recipe) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Recipe: %s", err))
	}

	*out = *target.(*Recipe)
}

// Validate valides the current information stored into the structure.
func (o *Recipe) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("deploymentMode", string(o.DeploymentMode), []string{"Unrestricted", "Unique", "NamespaceUnique"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("label", o.Label); err != nil {
		requiredErrors = requiredErrors.Append(err)
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

	if o.Options != nil {
		elemental.ResetDefaultForZeroValues(o.Options)
		if err := o.Options.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	for _, sub := range o.Steps {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateRequiredExternal("targetIdentities", o.TargetIdentities); err != nil {
		requiredErrors = requiredErrors.Append(err)
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
func (*Recipe) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := RecipeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return RecipeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Recipe) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RecipeAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Recipe) ValueForAttribute(name string) interface{} {

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
	case "deploymentMode":
		return o.DeploymentMode
	case "description":
		return o.Description
	case "icon":
		return o.Icon
	case "key":
		return o.Key
	case "label":
		return o.Label
	case "longDescription":
		return o.LongDescription
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
	case "options":
		return o.Options
	case "propagate":
		return o.Propagate
	case "protected":
		return o.Protected
	case "steps":
		return o.Steps
	case "successfullMessage":
		return o.SuccessfullMessage
	case "targetIdentities":
		return o.TargetIdentities
	case "template":
		return o.Template
	case "templateHash":
		return o.TemplateHash
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

// RecipeAttributesMap represents the map of attribute for Recipe.
var RecipeAttributesMap = map[string]elemental.AttributeSpecification{
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
	"DeploymentMode": {
		AllowedChoices: []string{"Unrestricted", "Unique", "NamespaceUnique"},
		BSONFieldName:  "deploymentmode",
		ConvertedName:  "DeploymentMode",
		DefaultValue:   RecipeDeploymentModeUnrestricted,
		Description: `Defines the deployment mode of the recipe. If ` + "`" + `Unrestricted` + "`" + `, the recipe can be
deployed multiple times in the current namespace and below. If ` + "`" + `Unique` + "`" + `,only one
deployment is allowed in the current namespace and its child namespaces. If
` + "`" + `NamespaceUnique` + "`" + `, only one deployment is allowed in the current namespace.`,
		Exposed: true,
		Name:    "deploymentMode",
		Stored:  true,
		Type:    "enum",
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
	"Icon": {
		AllowedChoices: []string{},
		BSONFieldName:  "icon",
		ConvertedName:  "Icon",
		Description:    `Contains a base64-encoded image for the recipe.`,
		Exposed:        true,
		Name:           "icon",
		Stored:         true,
		Type:           "string",
	},
	"Key": {
		AllowedChoices: []string{},
		BSONFieldName:  "key",
		ConvertedName:  "Key",
		Description:    `The unique key of the recipe.`,
		Exposed:        true,
		Name:           "key",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Label": {
		AllowedChoices: []string{},
		BSONFieldName:  "label",
		ConvertedName:  "Label",
		CreationOnly:   true,
		DefaultValue:   "magicpanda",
		Description:    `Defines the recipe.`,
		Exposed:        true,
		Name:           "label",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"LongDescription": {
		AllowedChoices: []string{},
		BSONFieldName:  "longdescription",
		ConvertedName:  "LongDescription",
		Description:    `Provides a long description of the recipe.`,
		Exposed:        true,
		Name:           "longDescription",
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
	"Options": {
		AllowedChoices: []string{},
		BSONFieldName:  "options",
		ConvertedName:  "Options",
		Description:    `Options of the recipe.`,
		Exposed:        true,
		Name:           "options",
		Stored:         true,
		SubType:        "recipeoptions",
		Type:           "ref",
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
	"Steps": {
		AllowedChoices: []string{},
		BSONFieldName:  "steps",
		ConvertedName:  "Steps",
		Description:    `Contains all the steps with parameters to follow for the recipe.`,
		Exposed:        true,
		Name:           "steps",
		Stored:         true,
		SubType:        "uistep",
		Type:           "refList",
	},
	"SuccessfullMessage": {
		AllowedChoices: []string{},
		BSONFieldName:  "successfullmessage",
		ConvertedName:  "SuccessfullMessage",
		Description:    `A string message presented upon success (optional).`,
		Exposed:        true,
		Name:           "successfullMessage",
		Stored:         true,
		Type:           "string",
	},
	"TargetIdentities": {
		AllowedChoices: []string{},
		BSONFieldName:  "targetidentities",
		ConvertedName:  "TargetIdentities",
		Description:    `Contains the list of identities the recipes will try to create.`,
		Exposed:        true,
		Name:           "targetIdentities",
		Required:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Template": {
		AllowedChoices: []string{},
		BSONFieldName:  "template",
		ConvertedName:  "Template",
		Description:    `Template of the recipe to import.`,
		Exposed:        true,
		Name:           "template",
		Stored:         true,
		Type:           "string",
	},
	"TemplateHash": {
		AllowedChoices: []string{},
		BSONFieldName:  "templatehash",
		ConvertedName:  "TemplateHash",
		Description:    `A hash of the template.`,
		Exposed:        true,
		Name:           "templateHash",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
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

// RecipeLowerCaseAttributesMap represents the map of attribute for Recipe.
var RecipeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"deploymentmode": {
		AllowedChoices: []string{"Unrestricted", "Unique", "NamespaceUnique"},
		BSONFieldName:  "deploymentmode",
		ConvertedName:  "DeploymentMode",
		DefaultValue:   RecipeDeploymentModeUnrestricted,
		Description: `Defines the deployment mode of the recipe. If ` + "`" + `Unrestricted` + "`" + `, the recipe can be
deployed multiple times in the current namespace and below. If ` + "`" + `Unique` + "`" + `,only one
deployment is allowed in the current namespace and its child namespaces. If
` + "`" + `NamespaceUnique` + "`" + `, only one deployment is allowed in the current namespace.`,
		Exposed: true,
		Name:    "deploymentMode",
		Stored:  true,
		Type:    "enum",
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
	"icon": {
		AllowedChoices: []string{},
		BSONFieldName:  "icon",
		ConvertedName:  "Icon",
		Description:    `Contains a base64-encoded image for the recipe.`,
		Exposed:        true,
		Name:           "icon",
		Stored:         true,
		Type:           "string",
	},
	"key": {
		AllowedChoices: []string{},
		BSONFieldName:  "key",
		ConvertedName:  "Key",
		Description:    `The unique key of the recipe.`,
		Exposed:        true,
		Name:           "key",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"label": {
		AllowedChoices: []string{},
		BSONFieldName:  "label",
		ConvertedName:  "Label",
		CreationOnly:   true,
		DefaultValue:   "magicpanda",
		Description:    `Defines the recipe.`,
		Exposed:        true,
		Name:           "label",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"longdescription": {
		AllowedChoices: []string{},
		BSONFieldName:  "longdescription",
		ConvertedName:  "LongDescription",
		Description:    `Provides a long description of the recipe.`,
		Exposed:        true,
		Name:           "longDescription",
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
	"options": {
		AllowedChoices: []string{},
		BSONFieldName:  "options",
		ConvertedName:  "Options",
		Description:    `Options of the recipe.`,
		Exposed:        true,
		Name:           "options",
		Stored:         true,
		SubType:        "recipeoptions",
		Type:           "ref",
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
	"steps": {
		AllowedChoices: []string{},
		BSONFieldName:  "steps",
		ConvertedName:  "Steps",
		Description:    `Contains all the steps with parameters to follow for the recipe.`,
		Exposed:        true,
		Name:           "steps",
		Stored:         true,
		SubType:        "uistep",
		Type:           "refList",
	},
	"successfullmessage": {
		AllowedChoices: []string{},
		BSONFieldName:  "successfullmessage",
		ConvertedName:  "SuccessfullMessage",
		Description:    `A string message presented upon success (optional).`,
		Exposed:        true,
		Name:           "successfullMessage",
		Stored:         true,
		Type:           "string",
	},
	"targetidentities": {
		AllowedChoices: []string{},
		BSONFieldName:  "targetidentities",
		ConvertedName:  "TargetIdentities",
		Description:    `Contains the list of identities the recipes will try to create.`,
		Exposed:        true,
		Name:           "targetIdentities",
		Required:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"template": {
		AllowedChoices: []string{},
		BSONFieldName:  "template",
		ConvertedName:  "Template",
		Description:    `Template of the recipe to import.`,
		Exposed:        true,
		Name:           "template",
		Stored:         true,
		Type:           "string",
	},
	"templatehash": {
		AllowedChoices: []string{},
		BSONFieldName:  "templatehash",
		ConvertedName:  "TemplateHash",
		Description:    `A hash of the template.`,
		Exposed:        true,
		Name:           "templateHash",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
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

// SparseRecipesList represents a list of SparseRecipes
type SparseRecipesList []*SparseRecipe

// Identity returns the identity of the objects in the list.
func (o SparseRecipesList) Identity() elemental.Identity {

	return RecipeIdentity
}

// Copy returns a pointer to a copy the SparseRecipesList.
func (o SparseRecipesList) Copy() elemental.Identifiables {

	copy := append(SparseRecipesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseRecipesList.
func (o SparseRecipesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseRecipesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseRecipe))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseRecipesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseRecipesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseRecipesList converted to RecipesList.
func (o SparseRecipesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseRecipesList) Version() int {

	return 1
}

// SparseRecipe represents the sparse version of a recipe.
type SparseRecipe struct {
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

	// Defines the deployment mode of the recipe. If `Unrestricted`, the recipe can be
	// deployed multiple times in the current namespace and below. If `Unique`,only one
	// deployment is allowed in the current namespace and its child namespaces. If
	// `NamespaceUnique`, only one deployment is allowed in the current namespace.
	DeploymentMode *RecipeDeploymentModeValue `json:"deploymentMode,omitempty" msgpack:"deploymentMode,omitempty" bson:"deploymentmode,omitempty" mapstructure:"deploymentMode,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Contains a base64-encoded image for the recipe.
	Icon *string `json:"icon,omitempty" msgpack:"icon,omitempty" bson:"icon,omitempty" mapstructure:"icon,omitempty"`

	// The unique key of the recipe.
	Key *string `json:"key,omitempty" msgpack:"key,omitempty" bson:"key,omitempty" mapstructure:"key,omitempty"`

	// Defines the recipe.
	Label *string `json:"label,omitempty" msgpack:"label,omitempty" bson:"label,omitempty" mapstructure:"label,omitempty"`

	// Provides a long description of the recipe.
	LongDescription *string `json:"longDescription,omitempty" msgpack:"longDescription,omitempty" bson:"longdescription,omitempty" mapstructure:"longDescription,omitempty"`

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

	// Options of the recipe.
	Options *RecipeOptions `json:"options,omitempty" msgpack:"options,omitempty" bson:"options,omitempty" mapstructure:"options,omitempty"`

	// Propagates the policy to all of its children.
	Propagate *bool `json:"propagate,omitempty" msgpack:"propagate,omitempty" bson:"propagate,omitempty" mapstructure:"propagate,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Contains all the steps with parameters to follow for the recipe.
	Steps *[]*UIStep `json:"steps,omitempty" msgpack:"steps,omitempty" bson:"steps,omitempty" mapstructure:"steps,omitempty"`

	// A string message presented upon success (optional).
	SuccessfullMessage *string `json:"successfullMessage,omitempty" msgpack:"successfullMessage,omitempty" bson:"successfullmessage,omitempty" mapstructure:"successfullMessage,omitempty"`

	// Contains the list of identities the recipes will try to create.
	TargetIdentities *[]string `json:"targetIdentities,omitempty" msgpack:"targetIdentities,omitempty" bson:"targetidentities,omitempty" mapstructure:"targetIdentities,omitempty"`

	// Template of the recipe to import.
	Template *string `json:"template,omitempty" msgpack:"template,omitempty" bson:"template,omitempty" mapstructure:"template,omitempty"`

	// A hash of the template.
	TemplateHash *string `json:"templateHash,omitempty" msgpack:"templateHash,omitempty" bson:"templatehash,omitempty" mapstructure:"templateHash,omitempty"`

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

// NewSparseRecipe returns a new  SparseRecipe.
func NewSparseRecipe() *SparseRecipe {
	return &SparseRecipe{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseRecipe) Identity() elemental.Identity {

	return RecipeIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseRecipe) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseRecipe) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseRecipe) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseRecipe{}

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
	if o.DeploymentMode != nil {
		s.DeploymentMode = o.DeploymentMode
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.Icon != nil {
		s.Icon = o.Icon
	}
	if o.Key != nil {
		s.Key = o.Key
	}
	if o.Label != nil {
		s.Label = o.Label
	}
	if o.LongDescription != nil {
		s.LongDescription = o.LongDescription
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
	if o.Options != nil {
		s.Options = o.Options
	}
	if o.Propagate != nil {
		s.Propagate = o.Propagate
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.Steps != nil {
		s.Steps = o.Steps
	}
	if o.SuccessfullMessage != nil {
		s.SuccessfullMessage = o.SuccessfullMessage
	}
	if o.TargetIdentities != nil {
		s.TargetIdentities = o.TargetIdentities
	}
	if o.Template != nil {
		s.Template = o.Template
	}
	if o.TemplateHash != nil {
		s.TemplateHash = o.TemplateHash
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
func (o *SparseRecipe) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseRecipe{}
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
	if s.DeploymentMode != nil {
		o.DeploymentMode = s.DeploymentMode
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Icon != nil {
		o.Icon = s.Icon
	}
	if s.Key != nil {
		o.Key = s.Key
	}
	if s.Label != nil {
		o.Label = s.Label
	}
	if s.LongDescription != nil {
		o.LongDescription = s.LongDescription
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
	if s.Options != nil {
		o.Options = s.Options
	}
	if s.Propagate != nil {
		o.Propagate = s.Propagate
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.Steps != nil {
		o.Steps = s.Steps
	}
	if s.SuccessfullMessage != nil {
		o.SuccessfullMessage = s.SuccessfullMessage
	}
	if s.TargetIdentities != nil {
		o.TargetIdentities = s.TargetIdentities
	}
	if s.Template != nil {
		o.Template = s.Template
	}
	if s.TemplateHash != nil {
		o.TemplateHash = s.TemplateHash
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
func (o *SparseRecipe) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseRecipe) ToPlain() elemental.PlainIdentifiable {

	out := NewRecipe()
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
	if o.DeploymentMode != nil {
		out.DeploymentMode = *o.DeploymentMode
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Icon != nil {
		out.Icon = *o.Icon
	}
	if o.Key != nil {
		out.Key = *o.Key
	}
	if o.Label != nil {
		out.Label = *o.Label
	}
	if o.LongDescription != nil {
		out.LongDescription = *o.LongDescription
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
	if o.Options != nil {
		out.Options = o.Options
	}
	if o.Propagate != nil {
		out.Propagate = *o.Propagate
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Steps != nil {
		out.Steps = *o.Steps
	}
	if o.SuccessfullMessage != nil {
		out.SuccessfullMessage = *o.SuccessfullMessage
	}
	if o.TargetIdentities != nil {
		out.TargetIdentities = *o.TargetIdentities
	}
	if o.Template != nil {
		out.Template = *o.Template
	}
	if o.TemplateHash != nil {
		out.TemplateHash = *o.TemplateHash
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
func (o *SparseRecipe) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseRecipe) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseRecipe) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseRecipe) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseRecipe) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseRecipe) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseRecipe) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseRecipe) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseRecipe) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseRecipe) SetDescription(description string) {

	o.Description = &description
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseRecipe) GetMetadata() (out []string) {

	if o.Metadata == nil {
		return
	}

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseRecipe) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseRecipe) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseRecipe) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseRecipe) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseRecipe) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseRecipe) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseRecipe) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseRecipe) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseRecipe) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *SparseRecipe) GetPropagate() (out bool) {

	if o.Propagate == nil {
		return
	}

	return *o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the address of the given value.
func (o *SparseRecipe) SetPropagate(propagate bool) {

	o.Propagate = &propagate
}

// GetProtected returns the Protected of the receiver.
func (o *SparseRecipe) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseRecipe) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseRecipe) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseRecipe) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseRecipe) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseRecipe) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseRecipe) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseRecipe) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseRecipe) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseRecipe) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseRecipe.
func (o *SparseRecipe) DeepCopy() *SparseRecipe {

	if o == nil {
		return nil
	}

	out := &SparseRecipe{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseRecipe.
func (o *SparseRecipe) DeepCopyInto(out *SparseRecipe) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseRecipe: %s", err))
	}

	*out = *target.(*SparseRecipe)
}

type mongoAttributesRecipe struct {
	ID                   bson.ObjectId             `bson:"_id,omitempty"`
	Annotations          map[string][]string       `bson:"annotations"`
	AssociatedTags       []string                  `bson:"associatedtags"`
	CreateIdempotencyKey string                    `bson:"createidempotencykey"`
	CreateTime           time.Time                 `bson:"createtime"`
	DeploymentMode       RecipeDeploymentModeValue `bson:"deploymentmode"`
	Description          string                    `bson:"description"`
	Icon                 string                    `bson:"icon"`
	Key                  string                    `bson:"key"`
	Label                string                    `bson:"label"`
	LongDescription      string                    `bson:"longdescription"`
	Metadata             []string                  `bson:"metadata"`
	MigrationsLog        map[string]string         `bson:"migrationslog,omitempty"`
	Name                 string                    `bson:"name"`
	Namespace            string                    `bson:"namespace"`
	NormalizedTags       []string                  `bson:"normalizedtags"`
	Options              *RecipeOptions            `bson:"options"`
	Propagate            bool                      `bson:"propagate"`
	Protected            bool                      `bson:"protected"`
	Steps                []*UIStep                 `bson:"steps"`
	SuccessfullMessage   string                    `bson:"successfullmessage"`
	TargetIdentities     []string                  `bson:"targetidentities"`
	Template             string                    `bson:"template"`
	TemplateHash         string                    `bson:"templatehash"`
	UpdateIdempotencyKey string                    `bson:"updateidempotencykey"`
	UpdateTime           time.Time                 `bson:"updatetime"`
	ZHash                int                       `bson:"zhash"`
	Zone                 int                       `bson:"zone"`
}
type mongoAttributesSparseRecipe struct {
	ID                   bson.ObjectId              `bson:"_id,omitempty"`
	Annotations          *map[string][]string       `bson:"annotations,omitempty"`
	AssociatedTags       *[]string                  `bson:"associatedtags,omitempty"`
	CreateIdempotencyKey *string                    `bson:"createidempotencykey,omitempty"`
	CreateTime           *time.Time                 `bson:"createtime,omitempty"`
	DeploymentMode       *RecipeDeploymentModeValue `bson:"deploymentmode,omitempty"`
	Description          *string                    `bson:"description,omitempty"`
	Icon                 *string                    `bson:"icon,omitempty"`
	Key                  *string                    `bson:"key,omitempty"`
	Label                *string                    `bson:"label,omitempty"`
	LongDescription      *string                    `bson:"longdescription,omitempty"`
	Metadata             *[]string                  `bson:"metadata,omitempty"`
	MigrationsLog        *map[string]string         `bson:"migrationslog,omitempty"`
	Name                 *string                    `bson:"name,omitempty"`
	Namespace            *string                    `bson:"namespace,omitempty"`
	NormalizedTags       *[]string                  `bson:"normalizedtags,omitempty"`
	Options              *RecipeOptions             `bson:"options,omitempty"`
	Propagate            *bool                      `bson:"propagate,omitempty"`
	Protected            *bool                      `bson:"protected,omitempty"`
	Steps                *[]*UIStep                 `bson:"steps,omitempty"`
	SuccessfullMessage   *string                    `bson:"successfullmessage,omitempty"`
	TargetIdentities     *[]string                  `bson:"targetidentities,omitempty"`
	Template             *string                    `bson:"template,omitempty"`
	TemplateHash         *string                    `bson:"templatehash,omitempty"`
	UpdateIdempotencyKey *string                    `bson:"updateidempotencykey,omitempty"`
	UpdateTime           *time.Time                 `bson:"updatetime,omitempty"`
	ZHash                *int                       `bson:"zhash,omitempty"`
	Zone                 *int                       `bson:"zone,omitempty"`
}
