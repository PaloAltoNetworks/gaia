package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// NamespaceRendererIdentity represents the Identity of the object.
var NamespaceRendererIdentity = elemental.Identity{
	Name:     "namespacerenderer",
	Category: "namespacerenderers",
	Package:  "squall",
	Private:  false,
}

// NamespaceRenderersList represents a list of NamespaceRenderers
type NamespaceRenderersList []*NamespaceRenderer

// Identity returns the identity of the objects in the list.
func (o NamespaceRenderersList) Identity() elemental.Identity {

	return NamespaceRendererIdentity
}

// Copy returns a pointer to a copy the NamespaceRenderersList.
func (o NamespaceRenderersList) Copy() elemental.Identifiables {

	copy := append(NamespaceRenderersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the NamespaceRenderersList.
func (o NamespaceRenderersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(NamespaceRenderersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*NamespaceRenderer))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o NamespaceRenderersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o NamespaceRenderersList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the NamespaceRenderersList converted to SparseNamespaceRenderersList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o NamespaceRenderersList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseNamespaceRenderersList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseNamespaceRenderer)
	}

	return out
}

// Version returns the version of the content.
func (o NamespaceRenderersList) Version() int {

	return 1
}

// NamespaceRenderer represents the model of a namespacerenderer
type NamespaceRenderer struct {
	// The namespace where the object should reside in.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	// List of tags of the object to render the namespace for.
	Tags []string `json:"tags" msgpack:"tags" bson:"-" mapstructure:"tags,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewNamespaceRenderer returns a new *NamespaceRenderer
func NewNamespaceRenderer() *NamespaceRenderer {

	return &NamespaceRenderer{
		ModelVersion: 1,
		Tags:         []string{},
	}
}

// Identity returns the Identity of the object.
func (o *NamespaceRenderer) Identity() elemental.Identity {

	return NamespaceRendererIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NamespaceRenderer) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NamespaceRenderer) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NamespaceRenderer) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesNamespaceRenderer{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NamespaceRenderer) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesNamespaceRenderer{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *NamespaceRenderer) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *NamespaceRenderer) BleveType() string {

	return "namespacerenderer"
}

// DefaultOrder returns the list of default ordering fields.
func (o *NamespaceRenderer) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *NamespaceRenderer) Doc() string {

	return `This object allows you to determine which namespace an object should reside in
based on the tags provided.`
}

func (o *NamespaceRenderer) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *NamespaceRenderer) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseNamespaceRenderer{
			Namespace: &o.Namespace,
			Tags:      &o.Tags,
		}
	}

	sp := &SparseNamespaceRenderer{}
	for _, f := range fields {
		switch f {
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "tags":
			sp.Tags = &(o.Tags)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseNamespaceRenderer to the object.
func (o *NamespaceRenderer) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseNamespaceRenderer)
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.Tags != nil {
		o.Tags = *so.Tags
	}
}

// DeepCopy returns a deep copy if the NamespaceRenderer.
func (o *NamespaceRenderer) DeepCopy() *NamespaceRenderer {

	if o == nil {
		return nil
	}

	out := &NamespaceRenderer{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *NamespaceRenderer.
func (o *NamespaceRenderer) DeepCopyInto(out *NamespaceRenderer) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy NamespaceRenderer: %s", err))
	}

	*out = *target.(*NamespaceRenderer)
}

// Validate valides the current information stored into the structure.
func (o *NamespaceRenderer) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("tags", o.Tags); err != nil {
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
func (*NamespaceRenderer) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := NamespaceRendererAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return NamespaceRendererLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*NamespaceRenderer) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return NamespaceRendererAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *NamespaceRenderer) ValueForAttribute(name string) interface{} {

	switch name {
	case "namespace":
		return o.Namespace
	case "tags":
		return o.Tags
	}

	return nil
}

// NamespaceRendererAttributesMap represents the map of attribute for NamespaceRenderer.
var NamespaceRendererAttributesMap = map[string]elemental.AttributeSpecification{
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `The namespace where the object should reside in.`,
		Exposed:        true,
		Name:           "namespace",
		ReadOnly:       true,
		Type:           "string",
	},
	"Tags": {
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		Description:    `List of tags of the object to render the namespace for.`,
		Exposed:        true,
		Name:           "tags",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
}

// NamespaceRendererLowerCaseAttributesMap represents the map of attribute for NamespaceRenderer.
var NamespaceRendererLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `The namespace where the object should reside in.`,
		Exposed:        true,
		Name:           "namespace",
		ReadOnly:       true,
		Type:           "string",
	},
	"tags": {
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		Description:    `List of tags of the object to render the namespace for.`,
		Exposed:        true,
		Name:           "tags",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
}

// SparseNamespaceRenderersList represents a list of SparseNamespaceRenderers
type SparseNamespaceRenderersList []*SparseNamespaceRenderer

// Identity returns the identity of the objects in the list.
func (o SparseNamespaceRenderersList) Identity() elemental.Identity {

	return NamespaceRendererIdentity
}

// Copy returns a pointer to a copy the SparseNamespaceRenderersList.
func (o SparseNamespaceRenderersList) Copy() elemental.Identifiables {

	copy := append(SparseNamespaceRenderersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseNamespaceRenderersList.
func (o SparseNamespaceRenderersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseNamespaceRenderersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseNamespaceRenderer))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseNamespaceRenderersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseNamespaceRenderersList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseNamespaceRenderersList converted to NamespaceRenderersList.
func (o SparseNamespaceRenderersList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseNamespaceRenderersList) Version() int {

	return 1
}

// SparseNamespaceRenderer represents the sparse version of a namespacerenderer.
type SparseNamespaceRenderer struct {
	// The namespace where the object should reside in.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"-" mapstructure:"namespace,omitempty"`

	// List of tags of the object to render the namespace for.
	Tags *[]string `json:"tags,omitempty" msgpack:"tags,omitempty" bson:"-" mapstructure:"tags,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseNamespaceRenderer returns a new  SparseNamespaceRenderer.
func NewSparseNamespaceRenderer() *SparseNamespaceRenderer {
	return &SparseNamespaceRenderer{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseNamespaceRenderer) Identity() elemental.Identity {

	return NamespaceRendererIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseNamespaceRenderer) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseNamespaceRenderer) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseNamespaceRenderer) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseNamespaceRenderer{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseNamespaceRenderer) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseNamespaceRenderer{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseNamespaceRenderer) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseNamespaceRenderer) ToPlain() elemental.PlainIdentifiable {

	out := NewNamespaceRenderer()
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.Tags != nil {
		out.Tags = *o.Tags
	}

	return out
}

// DeepCopy returns a deep copy if the SparseNamespaceRenderer.
func (o *SparseNamespaceRenderer) DeepCopy() *SparseNamespaceRenderer {

	if o == nil {
		return nil
	}

	out := &SparseNamespaceRenderer{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseNamespaceRenderer.
func (o *SparseNamespaceRenderer) DeepCopyInto(out *SparseNamespaceRenderer) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseNamespaceRenderer: %s", err))
	}

	*out = *target.(*SparseNamespaceRenderer)
}

type mongoAttributesNamespaceRenderer struct {
}
type mongoAttributesSparseNamespaceRenderer struct {
}
