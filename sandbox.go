package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// SandboxIdentity represents the Identity of the object.
var SandboxIdentity = elemental.Identity{
	Name:     "sandbox",
	Category: "sandboxes",
	Package:  "service",
	Private:  false,
}

// SandboxsList represents a list of Sandboxs
type SandboxsList []*Sandbox

// Identity returns the identity of the objects in the list.
func (o SandboxsList) Identity() elemental.Identity {

	return SandboxIdentity
}

// Copy returns a pointer to a copy the SandboxsList.
func (o SandboxsList) Copy() elemental.Identifiables {

	copy := append(SandboxsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SandboxsList.
func (o SandboxsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SandboxsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Sandbox))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SandboxsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SandboxsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the SandboxsList converted to SparseSandboxsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o SandboxsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseSandboxsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseSandbox)
	}

	return out
}

// Version returns the version of the content.
func (o SandboxsList) Version() int {

	return 1
}

// Sandbox represents the model of a sandbox
type Sandbox struct {
	// Contains a link to directly connect the UI to your API sandbox.
	URL string `json:"URL" msgpack:"URL" bson:"-" mapstructure:"URL,omitempty"`

	// The app credential data.
	Credentials *Credential `json:"credentials" msgpack:"credentials" bson:"-" mapstructure:"credentials,omitempty"`

	// Contains the lifetime of the sandbox namespace.
	Lifetime string `json:"lifetime" msgpack:"lifetime" bson:"-" mapstructure:"lifetime,omitempty"`

	// Contains the name of the sandbox namespace that has been created.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSandbox returns a new *Sandbox
func NewSandbox() *Sandbox {

	return &Sandbox{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Sandbox) Identity() elemental.Identity {

	return SandboxIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Sandbox) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Sandbox) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Sandbox) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSandbox{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Sandbox) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSandbox{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Sandbox) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Sandbox) BleveType() string {

	return "sandbox"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Sandbox) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Sandbox) Doc() string {

	return `This APIs allows to create a temporary namespace to experiment with
Microsegmentation.
This API is not authenticated, and contains small quotas. After one hour,
everything will be deleted.`
}

func (o *Sandbox) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Sandbox) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseSandbox{
			URL:         &o.URL,
			Credentials: o.Credentials,
			Lifetime:    &o.Lifetime,
			Namespace:   &o.Namespace,
		}
	}

	sp := &SparseSandbox{}
	for _, f := range fields {
		switch f {
		case "URL":
			sp.URL = &(o.URL)
		case "credentials":
			sp.Credentials = o.Credentials
		case "lifetime":
			sp.Lifetime = &(o.Lifetime)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseSandbox to the object.
func (o *Sandbox) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseSandbox)
	if so.URL != nil {
		o.URL = *so.URL
	}
	if so.Credentials != nil {
		o.Credentials = so.Credentials
	}
	if so.Lifetime != nil {
		o.Lifetime = *so.Lifetime
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
}

// DeepCopy returns a deep copy if the Sandbox.
func (o *Sandbox) DeepCopy() *Sandbox {

	if o == nil {
		return nil
	}

	out := &Sandbox{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Sandbox.
func (o *Sandbox) DeepCopyInto(out *Sandbox) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Sandbox: %s", err))
	}

	*out = *target.(*Sandbox)
}

// Validate valides the current information stored into the structure.
func (o *Sandbox) Validate() error {

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
func (*Sandbox) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := SandboxAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return SandboxLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Sandbox) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return SandboxAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Sandbox) ValueForAttribute(name string) interface{} {

	switch name {
	case "URL":
		return o.URL
	case "credentials":
		return o.Credentials
	case "lifetime":
		return o.Lifetime
	case "namespace":
		return o.Namespace
	}

	return nil
}

// SandboxAttributesMap represents the map of attribute for Sandbox.
var SandboxAttributesMap = map[string]elemental.AttributeSpecification{
	"URL": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "URL",
		Description:    `Contains a link to directly connect the UI to your API sandbox.`,
		Exposed:        true,
		Name:           "URL",
		ReadOnly:       true,
		Type:           "string",
	},
	"Credentials": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Credentials",
		Description:    `The app credential data.`,
		Exposed:        true,
		Name:           "credentials",
		Orderable:      true,
		ReadOnly:       true,
		SubType:        "credential",
		Transient:      true,
		Type:           "ref",
	},
	"Lifetime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Lifetime",
		Description:    `Contains the lifetime of the sandbox namespace.`,
		Exposed:        true,
		Name:           "lifetime",
		ReadOnly:       true,
		Type:           "string",
	},
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Contains the name of the sandbox namespace that has been created.`,
		Exposed:        true,
		Name:           "namespace",
		ReadOnly:       true,
		Type:           "string",
	},
}

// SandboxLowerCaseAttributesMap represents the map of attribute for Sandbox.
var SandboxLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"url": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "URL",
		Description:    `Contains a link to directly connect the UI to your API sandbox.`,
		Exposed:        true,
		Name:           "URL",
		ReadOnly:       true,
		Type:           "string",
	},
	"credentials": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Credentials",
		Description:    `The app credential data.`,
		Exposed:        true,
		Name:           "credentials",
		Orderable:      true,
		ReadOnly:       true,
		SubType:        "credential",
		Transient:      true,
		Type:           "ref",
	},
	"lifetime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Lifetime",
		Description:    `Contains the lifetime of the sandbox namespace.`,
		Exposed:        true,
		Name:           "lifetime",
		ReadOnly:       true,
		Type:           "string",
	},
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Contains the name of the sandbox namespace that has been created.`,
		Exposed:        true,
		Name:           "namespace",
		ReadOnly:       true,
		Type:           "string",
	},
}

// SparseSandboxsList represents a list of SparseSandboxs
type SparseSandboxsList []*SparseSandbox

// Identity returns the identity of the objects in the list.
func (o SparseSandboxsList) Identity() elemental.Identity {

	return SandboxIdentity
}

// Copy returns a pointer to a copy the SparseSandboxsList.
func (o SparseSandboxsList) Copy() elemental.Identifiables {

	copy := append(SparseSandboxsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseSandboxsList.
func (o SparseSandboxsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseSandboxsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseSandbox))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseSandboxsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseSandboxsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseSandboxsList converted to SandboxsList.
func (o SparseSandboxsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseSandboxsList) Version() int {

	return 1
}

// SparseSandbox represents the sparse version of a sandbox.
type SparseSandbox struct {
	// Contains a link to directly connect the UI to your API sandbox.
	URL *string `json:"URL,omitempty" msgpack:"URL,omitempty" bson:"-" mapstructure:"URL,omitempty"`

	// The app credential data.
	Credentials *Credential `json:"credentials,omitempty" msgpack:"credentials,omitempty" bson:"-" mapstructure:"credentials,omitempty"`

	// Contains the lifetime of the sandbox namespace.
	Lifetime *string `json:"lifetime,omitempty" msgpack:"lifetime,omitempty" bson:"-" mapstructure:"lifetime,omitempty"`

	// Contains the name of the sandbox namespace that has been created.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"-" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseSandbox returns a new  SparseSandbox.
func NewSparseSandbox() *SparseSandbox {
	return &SparseSandbox{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseSandbox) Identity() elemental.Identity {

	return SandboxIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseSandbox) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseSandbox) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseSandbox) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseSandbox{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseSandbox) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseSandbox{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseSandbox) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseSandbox) ToPlain() elemental.PlainIdentifiable {

	out := NewSandbox()
	if o.URL != nil {
		out.URL = *o.URL
	}
	if o.Credentials != nil {
		out.Credentials = o.Credentials
	}
	if o.Lifetime != nil {
		out.Lifetime = *o.Lifetime
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}

	return out
}

// DeepCopy returns a deep copy if the SparseSandbox.
func (o *SparseSandbox) DeepCopy() *SparseSandbox {

	if o == nil {
		return nil
	}

	out := &SparseSandbox{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseSandbox.
func (o *SparseSandbox) DeepCopyInto(out *SparseSandbox) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseSandbox: %s", err))
	}

	*out = *target.(*SparseSandbox)
}

type mongoAttributesSandbox struct {
}
type mongoAttributesSparseSandbox struct {
}
