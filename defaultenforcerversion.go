package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// DefaultEnforcerVersionIdentity represents the Identity of the object.
var DefaultEnforcerVersionIdentity = elemental.Identity{
	Name:     "defaultenforcerversion",
	Category: "defaultenforcerversion",
	Package:  "squall",
	Private:  false,
}

// DefaultEnforcerVersionsList represents a list of DefaultEnforcerVersions
type DefaultEnforcerVersionsList []*DefaultEnforcerVersion

// Identity returns the identity of the objects in the list.
func (o DefaultEnforcerVersionsList) Identity() elemental.Identity {

	return DefaultEnforcerVersionIdentity
}

// Copy returns a pointer to a copy the DefaultEnforcerVersionsList.
func (o DefaultEnforcerVersionsList) Copy() elemental.Identifiables {

	copy := append(DefaultEnforcerVersionsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the DefaultEnforcerVersionsList.
func (o DefaultEnforcerVersionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(DefaultEnforcerVersionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*DefaultEnforcerVersion))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o DefaultEnforcerVersionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o DefaultEnforcerVersionsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the DefaultEnforcerVersionsList converted to SparseDefaultEnforcerVersionsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o DefaultEnforcerVersionsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseDefaultEnforcerVersionsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseDefaultEnforcerVersion)
	}

	return out
}

// Version returns the version of the content.
func (o DefaultEnforcerVersionsList) Version() int {

	return 1
}

// DefaultEnforcerVersion represents the model of a defaultenforcerversion
type DefaultEnforcerVersion struct {
	// The default enforcer version for the namespace.
	DefaultVersion string `json:"defaultVersion" msgpack:"defaultVersion" bson:"-" mapstructure:"defaultVersion,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewDefaultEnforcerVersion returns a new *DefaultEnforcerVersion
func NewDefaultEnforcerVersion() *DefaultEnforcerVersion {

	return &DefaultEnforcerVersion{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *DefaultEnforcerVersion) Identity() elemental.Identity {

	return DefaultEnforcerVersionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DefaultEnforcerVersion) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DefaultEnforcerVersion) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *DefaultEnforcerVersion) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesDefaultEnforcerVersion{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *DefaultEnforcerVersion) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesDefaultEnforcerVersion{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *DefaultEnforcerVersion) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *DefaultEnforcerVersion) BleveType() string {

	return "defaultenforcerversion"
}

// DefaultOrder returns the list of default ordering fields.
func (o *DefaultEnforcerVersion) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *DefaultEnforcerVersion) Doc() string {

	return `Returns the default enforcer version of the specified namespace.`
}

func (o *DefaultEnforcerVersion) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *DefaultEnforcerVersion) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseDefaultEnforcerVersion{
			DefaultVersion: &o.DefaultVersion,
		}
	}

	sp := &SparseDefaultEnforcerVersion{}
	for _, f := range fields {
		switch f {
		case "defaultVersion":
			sp.DefaultVersion = &(o.DefaultVersion)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseDefaultEnforcerVersion to the object.
func (o *DefaultEnforcerVersion) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseDefaultEnforcerVersion)
	if so.DefaultVersion != nil {
		o.DefaultVersion = *so.DefaultVersion
	}
}

// DeepCopy returns a deep copy if the DefaultEnforcerVersion.
func (o *DefaultEnforcerVersion) DeepCopy() *DefaultEnforcerVersion {

	if o == nil {
		return nil
	}

	out := &DefaultEnforcerVersion{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *DefaultEnforcerVersion.
func (o *DefaultEnforcerVersion) DeepCopyInto(out *DefaultEnforcerVersion) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy DefaultEnforcerVersion: %s", err))
	}

	*out = *target.(*DefaultEnforcerVersion)
}

// Validate valides the current information stored into the structure.
func (o *DefaultEnforcerVersion) Validate() error {

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
func (*DefaultEnforcerVersion) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := DefaultEnforcerVersionAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return DefaultEnforcerVersionLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*DefaultEnforcerVersion) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return DefaultEnforcerVersionAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *DefaultEnforcerVersion) ValueForAttribute(name string) interface{} {

	switch name {
	case "defaultVersion":
		return o.DefaultVersion
	}

	return nil
}

// DefaultEnforcerVersionAttributesMap represents the map of attribute for DefaultEnforcerVersion.
var DefaultEnforcerVersionAttributesMap = map[string]elemental.AttributeSpecification{
	"DefaultVersion": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "DefaultVersion",
		Description:    `The default enforcer version for the namespace.`,
		Exposed:        true,
		Name:           "defaultVersion",
		ReadOnly:       true,
		Type:           "string",
	},
}

// DefaultEnforcerVersionLowerCaseAttributesMap represents the map of attribute for DefaultEnforcerVersion.
var DefaultEnforcerVersionLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"defaultversion": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "DefaultVersion",
		Description:    `The default enforcer version for the namespace.`,
		Exposed:        true,
		Name:           "defaultVersion",
		ReadOnly:       true,
		Type:           "string",
	},
}

// SparseDefaultEnforcerVersionsList represents a list of SparseDefaultEnforcerVersions
type SparseDefaultEnforcerVersionsList []*SparseDefaultEnforcerVersion

// Identity returns the identity of the objects in the list.
func (o SparseDefaultEnforcerVersionsList) Identity() elemental.Identity {

	return DefaultEnforcerVersionIdentity
}

// Copy returns a pointer to a copy the SparseDefaultEnforcerVersionsList.
func (o SparseDefaultEnforcerVersionsList) Copy() elemental.Identifiables {

	copy := append(SparseDefaultEnforcerVersionsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseDefaultEnforcerVersionsList.
func (o SparseDefaultEnforcerVersionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseDefaultEnforcerVersionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseDefaultEnforcerVersion))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseDefaultEnforcerVersionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseDefaultEnforcerVersionsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseDefaultEnforcerVersionsList converted to DefaultEnforcerVersionsList.
func (o SparseDefaultEnforcerVersionsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseDefaultEnforcerVersionsList) Version() int {

	return 1
}

// SparseDefaultEnforcerVersion represents the sparse version of a defaultenforcerversion.
type SparseDefaultEnforcerVersion struct {
	// The default enforcer version for the namespace.
	DefaultVersion *string `json:"defaultVersion,omitempty" msgpack:"defaultVersion,omitempty" bson:"-" mapstructure:"defaultVersion,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseDefaultEnforcerVersion returns a new  SparseDefaultEnforcerVersion.
func NewSparseDefaultEnforcerVersion() *SparseDefaultEnforcerVersion {
	return &SparseDefaultEnforcerVersion{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseDefaultEnforcerVersion) Identity() elemental.Identity {

	return DefaultEnforcerVersionIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseDefaultEnforcerVersion) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseDefaultEnforcerVersion) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseDefaultEnforcerVersion) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseDefaultEnforcerVersion{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseDefaultEnforcerVersion) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseDefaultEnforcerVersion{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseDefaultEnforcerVersion) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseDefaultEnforcerVersion) ToPlain() elemental.PlainIdentifiable {

	out := NewDefaultEnforcerVersion()
	if o.DefaultVersion != nil {
		out.DefaultVersion = *o.DefaultVersion
	}

	return out
}

// DeepCopy returns a deep copy if the SparseDefaultEnforcerVersion.
func (o *SparseDefaultEnforcerVersion) DeepCopy() *SparseDefaultEnforcerVersion {

	if o == nil {
		return nil
	}

	out := &SparseDefaultEnforcerVersion{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseDefaultEnforcerVersion.
func (o *SparseDefaultEnforcerVersion) DeepCopyInto(out *SparseDefaultEnforcerVersion) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseDefaultEnforcerVersion: %s", err))
	}

	*out = *target.(*SparseDefaultEnforcerVersion)
}

type mongoAttributesDefaultEnforcerVersion struct {
}
type mongoAttributesSparseDefaultEnforcerVersion struct {
}
