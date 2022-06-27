package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// RenderTemplateIdentity represents the Identity of the object.
var RenderTemplateIdentity = elemental.Identity{
	Name:     "rendertemplate",
	Category: "rendertemplates",
	Package:  "ignis",
	Private:  false,
}

// RenderTemplatesList represents a list of RenderTemplates
type RenderTemplatesList []*RenderTemplate

// Identity returns the identity of the objects in the list.
func (o RenderTemplatesList) Identity() elemental.Identity {

	return RenderTemplateIdentity
}

// Copy returns a pointer to a copy the RenderTemplatesList.
func (o RenderTemplatesList) Copy() elemental.Identifiables {

	copy := append(RenderTemplatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the RenderTemplatesList.
func (o RenderTemplatesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(RenderTemplatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*RenderTemplate))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o RenderTemplatesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o RenderTemplatesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the RenderTemplatesList converted to SparseRenderTemplatesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o RenderTemplatesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseRenderTemplatesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseRenderTemplate)
	}

	return out
}

// Version returns the version of the content.
func (o RenderTemplatesList) Version() int {

	return 1
}

// RenderTemplate represents the model of a rendertemplate
type RenderTemplate struct {
	// Holds the rendered template.
	Output string `json:"output" msgpack:"output" bson:"-" mapstructure:"output,omitempty"`

	// Contains the computed parameters.
	Parameters map[string]interface{} `json:"parameters" msgpack:"parameters" bson:"-" mapstructure:"parameters,omitempty"`

	// Template of the recipe.
	Template string `json:"template" msgpack:"template" bson:"-" mapstructure:"template,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewRenderTemplate returns a new *RenderTemplate
func NewRenderTemplate() *RenderTemplate {

	return &RenderTemplate{
		ModelVersion: 1,
		Parameters:   map[string]interface{}{},
	}
}

// Identity returns the Identity of the object.
func (o *RenderTemplate) Identity() elemental.Identity {

	return RenderTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *RenderTemplate) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *RenderTemplate) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *RenderTemplate) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesRenderTemplate{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *RenderTemplate) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesRenderTemplate{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *RenderTemplate) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *RenderTemplate) BleveType() string {

	return "rendertemplate"
}

// DefaultOrder returns the list of default ordering fields.
func (o *RenderTemplate) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *RenderTemplate) Doc() string {

	return `Cooks a template based on some parameters.`
}

func (o *RenderTemplate) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *RenderTemplate) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseRenderTemplate{
			Output:     &o.Output,
			Parameters: &o.Parameters,
			Template:   &o.Template,
		}
	}

	sp := &SparseRenderTemplate{}
	for _, f := range fields {
		switch f {
		case "output":
			sp.Output = &(o.Output)
		case "parameters":
			sp.Parameters = &(o.Parameters)
		case "template":
			sp.Template = &(o.Template)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseRenderTemplate to the object.
func (o *RenderTemplate) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseRenderTemplate)
	if so.Output != nil {
		o.Output = *so.Output
	}
	if so.Parameters != nil {
		o.Parameters = *so.Parameters
	}
	if so.Template != nil {
		o.Template = *so.Template
	}
}

// DeepCopy returns a deep copy if the RenderTemplate.
func (o *RenderTemplate) DeepCopy() *RenderTemplate {

	if o == nil {
		return nil
	}

	out := &RenderTemplate{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *RenderTemplate.
func (o *RenderTemplate) DeepCopyInto(out *RenderTemplate) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy RenderTemplate: %s", err))
	}

	*out = *target.(*RenderTemplate)
}

// Validate valides the current information stored into the structure.
func (o *RenderTemplate) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*RenderTemplate) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := RenderTemplateAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return RenderTemplateLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*RenderTemplate) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RenderTemplateAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *RenderTemplate) ValueForAttribute(name string) interface{} {

	switch name {
	case "output":
		return o.Output
	case "parameters":
		return o.Parameters
	case "template":
		return o.Template
	}

	return nil
}

// RenderTemplateAttributesMap represents the map of attribute for RenderTemplate.
var RenderTemplateAttributesMap = map[string]elemental.AttributeSpecification{
	"Output": {
		AllowedChoices: []string{},
		ConvertedName:  "Output",
		Description:    `Holds the rendered template.`,
		Exposed:        true,
		Name:           "output",
		Type:           "string",
	},
	"Parameters": {
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Contains the computed parameters.`,
		Exposed:        true,
		Name:           "parameters",
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"Template": {
		AllowedChoices: []string{},
		ConvertedName:  "Template",
		Description:    `Template of the recipe.`,
		Exposed:        true,
		Name:           "template",
		Type:           "string",
	},
}

// RenderTemplateLowerCaseAttributesMap represents the map of attribute for RenderTemplate.
var RenderTemplateLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"output": {
		AllowedChoices: []string{},
		ConvertedName:  "Output",
		Description:    `Holds the rendered template.`,
		Exposed:        true,
		Name:           "output",
		Type:           "string",
	},
	"parameters": {
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Contains the computed parameters.`,
		Exposed:        true,
		Name:           "parameters",
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"template": {
		AllowedChoices: []string{},
		ConvertedName:  "Template",
		Description:    `Template of the recipe.`,
		Exposed:        true,
		Name:           "template",
		Type:           "string",
	},
}

// SparseRenderTemplatesList represents a list of SparseRenderTemplates
type SparseRenderTemplatesList []*SparseRenderTemplate

// Identity returns the identity of the objects in the list.
func (o SparseRenderTemplatesList) Identity() elemental.Identity {

	return RenderTemplateIdentity
}

// Copy returns a pointer to a copy the SparseRenderTemplatesList.
func (o SparseRenderTemplatesList) Copy() elemental.Identifiables {

	copy := append(SparseRenderTemplatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseRenderTemplatesList.
func (o SparseRenderTemplatesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseRenderTemplatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseRenderTemplate))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseRenderTemplatesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseRenderTemplatesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseRenderTemplatesList converted to RenderTemplatesList.
func (o SparseRenderTemplatesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseRenderTemplatesList) Version() int {

	return 1
}

// SparseRenderTemplate represents the sparse version of a rendertemplate.
type SparseRenderTemplate struct {
	// Holds the rendered template.
	Output *string `json:"output,omitempty" msgpack:"output,omitempty" bson:"-" mapstructure:"output,omitempty"`

	// Contains the computed parameters.
	Parameters *map[string]interface{} `json:"parameters,omitempty" msgpack:"parameters,omitempty" bson:"-" mapstructure:"parameters,omitempty"`

	// Template of the recipe.
	Template *string `json:"template,omitempty" msgpack:"template,omitempty" bson:"-" mapstructure:"template,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseRenderTemplate returns a new  SparseRenderTemplate.
func NewSparseRenderTemplate() *SparseRenderTemplate {
	return &SparseRenderTemplate{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseRenderTemplate) Identity() elemental.Identity {

	return RenderTemplateIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseRenderTemplate) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseRenderTemplate) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseRenderTemplate) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseRenderTemplate{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseRenderTemplate) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseRenderTemplate{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseRenderTemplate) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseRenderTemplate) ToPlain() elemental.PlainIdentifiable {

	out := NewRenderTemplate()
	if o.Output != nil {
		out.Output = *o.Output
	}
	if o.Parameters != nil {
		out.Parameters = *o.Parameters
	}
	if o.Template != nil {
		out.Template = *o.Template
	}

	return out
}

// DeepCopy returns a deep copy if the SparseRenderTemplate.
func (o *SparseRenderTemplate) DeepCopy() *SparseRenderTemplate {

	if o == nil {
		return nil
	}

	out := &SparseRenderTemplate{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseRenderTemplate.
func (o *SparseRenderTemplate) DeepCopyInto(out *SparseRenderTemplate) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseRenderTemplate: %s", err))
	}

	*out = *target.(*SparseRenderTemplate)
}

type mongoAttributesRenderTemplate struct {
}
type mongoAttributesSparseRenderTemplate struct {
}
