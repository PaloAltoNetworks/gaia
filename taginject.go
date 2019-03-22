package gaia

import (
	"fmt"
	"sync"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TagInjectIdentity represents the Identity of the object.
var TagInjectIdentity = elemental.Identity{
	Name:     "taginject",
	Category: "taginjects",
	Package:  "tagle",
	Private:  true,
}

// TagInjectsList represents a list of TagInjects
type TagInjectsList []*TagInject

// Identity returns the identity of the objects in the list.
func (o TagInjectsList) Identity() elemental.Identity {

	return TagInjectIdentity
}

// Copy returns a pointer to a copy the TagInjectsList.
func (o TagInjectsList) Copy() elemental.Identifiables {

	copy := append(TagInjectsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TagInjectsList.
func (o TagInjectsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(TagInjectsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*TagInject))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TagInjectsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TagInjectsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the TagInjectsList converted to SparseTagInjectsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o TagInjectsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o TagInjectsList) Version() int {

	return 1
}

// TagInject represents the model of a taginject
type TagInject struct {
	// List of tags to be added.
	AddedTags map[string]int `json:"addedTags" bson:"-" mapstructure:"addedTags,omitempty"`

	// List of tags to be removed.
	RemovedTags map[string]int `json:"removedTags" bson:"-" mapstructure:"removedTags,omitempty"`

	// List of tags to inject.
	TargetNamespace string `json:"targetNamespace" bson:"-" mapstructure:"targetNamespace,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewTagInject returns a new *TagInject
func NewTagInject() *TagInject {

	return &TagInject{
		ModelVersion: 1,
		AddedTags:    map[string]int{},
		RemovedTags:  map[string]int{},
	}
}

// Identity returns the Identity of the object.
func (o *TagInject) Identity() elemental.Identity {

	return TagInjectIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *TagInject) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *TagInject) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *TagInject) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *TagInject) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *TagInject) Doc() string {
	return `This internal api is used to inject a new tag in the database.`
}

func (o *TagInject) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *TagInject) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseTagInject{
			AddedTags:       &o.AddedTags,
			RemovedTags:     &o.RemovedTags,
			TargetNamespace: &o.TargetNamespace,
		}
	}

	sp := &SparseTagInject{}
	for _, f := range fields {
		switch f {
		case "addedTags":
			sp.AddedTags = &(o.AddedTags)
		case "removedTags":
			sp.RemovedTags = &(o.RemovedTags)
		case "targetNamespace":
			sp.TargetNamespace = &(o.TargetNamespace)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseTagInject to the object.
func (o *TagInject) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseTagInject)
	if so.AddedTags != nil {
		o.AddedTags = *so.AddedTags
	}
	if so.RemovedTags != nil {
		o.RemovedTags = *so.RemovedTags
	}
	if so.TargetNamespace != nil {
		o.TargetNamespace = *so.TargetNamespace
	}
}

// DeepCopy returns a deep copy if the TagInject.
func (o *TagInject) DeepCopy() *TagInject {

	if o == nil {
		return nil
	}

	out := &TagInject{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *TagInject.
func (o *TagInject) DeepCopyInto(out *TagInject) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy TagInject: %s", err))
	}

	*out = *target.(*TagInject)
}

// Validate valides the current information stored into the structure.
func (o *TagInject) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("targetNamespace", o.TargetNamespace); err != nil {
		requiredErrors = append(requiredErrors, err)
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
func (*TagInject) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TagInjectAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TagInjectLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*TagInject) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TagInjectAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *TagInject) ValueForAttribute(name string) interface{} {

	switch name {
	case "addedTags":
		return o.AddedTags
	case "removedTags":
		return o.RemovedTags
	case "targetNamespace":
		return o.TargetNamespace
	}

	return nil
}

// TagInjectAttributesMap represents the map of attribute for TagInject.
var TagInjectAttributesMap = map[string]elemental.AttributeSpecification{
	"AddedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AddedTags",
		Description:    `List of tags to be added.`,
		Exposed:        true,
		Name:           "addedTags",
		SubType:        "map[string]int",
		Type:           "external",
	},
	"RemovedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RemovedTags",
		Description:    `List of tags to be removed.`,
		Exposed:        true,
		Name:           "removedTags",
		SubType:        "map[string]int",
		Type:           "external",
	},
	"TargetNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNamespace",
		Description:    `List of tags to inject.`,
		Exposed:        true,
		Name:           "targetNamespace",
		Required:       true,
		Type:           "string",
	},
}

// TagInjectLowerCaseAttributesMap represents the map of attribute for TagInject.
var TagInjectLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"addedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AddedTags",
		Description:    `List of tags to be added.`,
		Exposed:        true,
		Name:           "addedTags",
		SubType:        "map[string]int",
		Type:           "external",
	},
	"removedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RemovedTags",
		Description:    `List of tags to be removed.`,
		Exposed:        true,
		Name:           "removedTags",
		SubType:        "map[string]int",
		Type:           "external",
	},
	"targetnamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNamespace",
		Description:    `List of tags to inject.`,
		Exposed:        true,
		Name:           "targetNamespace",
		Required:       true,
		Type:           "string",
	},
}

// SparseTagInjectsList represents a list of SparseTagInjects
type SparseTagInjectsList []*SparseTagInject

// Identity returns the identity of the objects in the list.
func (o SparseTagInjectsList) Identity() elemental.Identity {

	return TagInjectIdentity
}

// Copy returns a pointer to a copy the SparseTagInjectsList.
func (o SparseTagInjectsList) Copy() elemental.Identifiables {

	copy := append(SparseTagInjectsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseTagInjectsList.
func (o SparseTagInjectsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseTagInjectsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseTagInject))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseTagInjectsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseTagInjectsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseTagInjectsList converted to TagInjectsList.
func (o SparseTagInjectsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseTagInjectsList) Version() int {

	return 1
}

// SparseTagInject represents the sparse version of a taginject.
type SparseTagInject struct {
	// List of tags to be added.
	AddedTags *map[string]int `json:"addedTags,omitempty" bson:"-" mapstructure:"addedTags,omitempty"`

	// List of tags to be removed.
	RemovedTags *map[string]int `json:"removedTags,omitempty" bson:"-" mapstructure:"removedTags,omitempty"`

	// List of tags to inject.
	TargetNamespace *string `json:"targetNamespace,omitempty" bson:"-" mapstructure:"targetNamespace,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseTagInject returns a new  SparseTagInject.
func NewSparseTagInject() *SparseTagInject {
	return &SparseTagInject{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseTagInject) Identity() elemental.Identity {

	return TagInjectIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseTagInject) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseTagInject) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseTagInject) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseTagInject) ToPlain() elemental.PlainIdentifiable {

	out := NewTagInject()
	if o.AddedTags != nil {
		out.AddedTags = *o.AddedTags
	}
	if o.RemovedTags != nil {
		out.RemovedTags = *o.RemovedTags
	}
	if o.TargetNamespace != nil {
		out.TargetNamespace = *o.TargetNamespace
	}

	return out
}

// DeepCopy returns a deep copy if the SparseTagInject.
func (o *SparseTagInject) DeepCopy() *SparseTagInject {

	if o == nil {
		return nil
	}

	out := &SparseTagInject{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseTagInject.
func (o *SparseTagInject) DeepCopyInto(out *SparseTagInject) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseTagInject: %s", err))
	}

	*out = *target.(*SparseTagInject)
}
