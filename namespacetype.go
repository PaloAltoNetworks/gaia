package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// NamespaceTypeIdentity represents the Identity of the object.
var NamespaceTypeIdentity = elemental.Identity{
	Name:     "namespacetype",
	Category: "namespacetypes",
	Package:  "squall",
	Private:  false,
}

// NamespaceTypesList represents a list of NamespaceTypes
type NamespaceTypesList []*NamespaceType

// Identity returns the identity of the objects in the list.
func (o NamespaceTypesList) Identity() elemental.Identity {

	return NamespaceTypeIdentity
}

// Copy returns a pointer to a copy the NamespaceTypesList.
func (o NamespaceTypesList) Copy() elemental.Identifiables {

	copy := append(NamespaceTypesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the NamespaceTypesList.
func (o NamespaceTypesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(NamespaceTypesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*NamespaceType))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o NamespaceTypesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o NamespaceTypesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the NamespaceTypesList converted to SparseNamespaceTypesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o NamespaceTypesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseNamespaceTypesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseNamespaceType)
	}

	return out
}

// Version returns the version of the content.
func (o NamespaceTypesList) Version() int {

	return 1
}

// NamespaceType represents the model of a namespacetype
type NamespaceType struct {
	// the namespace type for the current namespace.
	Type string `json:"type" msgpack:"type" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewNamespaceType returns a new *NamespaceType
func NewNamespaceType() *NamespaceType {

	return &NamespaceType{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *NamespaceType) Identity() elemental.Identity {

	return NamespaceTypeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NamespaceType) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NamespaceType) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NamespaceType) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesNamespaceType{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NamespaceType) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesNamespaceType{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *NamespaceType) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *NamespaceType) BleveType() string {

	return "namespacetype"
}

// DefaultOrder returns the list of default ordering fields.
func (o *NamespaceType) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *NamespaceType) Doc() string {

	return `Returns the type of the specified namespace.`
}

func (o *NamespaceType) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *NamespaceType) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseNamespaceType{
			Type: &o.Type,
		}
	}

	sp := &SparseNamespaceType{}
	for _, f := range fields {
		switch f {
		case "type":
			sp.Type = &(o.Type)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseNamespaceType to the object.
func (o *NamespaceType) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseNamespaceType)
	if so.Type != nil {
		o.Type = *so.Type
	}
}

// DeepCopy returns a deep copy if the NamespaceType.
func (o *NamespaceType) DeepCopy() *NamespaceType {

	if o == nil {
		return nil
	}

	out := &NamespaceType{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *NamespaceType.
func (o *NamespaceType) DeepCopyInto(out *NamespaceType) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy NamespaceType: %s", err))
	}

	*out = *target.(*NamespaceType)
}

// Validate valides the current information stored into the structure.
func (o *NamespaceType) Validate() error {

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
func (*NamespaceType) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := NamespaceTypeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return NamespaceTypeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*NamespaceType) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return NamespaceTypeAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *NamespaceType) ValueForAttribute(name string) interface{} {

	switch name {
	case "type":
		return o.Type
	}

	return nil
}

// NamespaceTypeAttributesMap represents the map of attribute for NamespaceType.
var NamespaceTypeAttributesMap = map[string]elemental.AttributeSpecification{
	"Type": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Type",
		Description:    `the namespace type for the current namespace.`,
		Exposed:        true,
		Name:           "type",
		ReadOnly:       true,
		Type:           "string",
	},
}

// NamespaceTypeLowerCaseAttributesMap represents the map of attribute for NamespaceType.
var NamespaceTypeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"type": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Type",
		Description:    `the namespace type for the current namespace.`,
		Exposed:        true,
		Name:           "type",
		ReadOnly:       true,
		Type:           "string",
	},
}

// SparseNamespaceTypesList represents a list of SparseNamespaceTypes
type SparseNamespaceTypesList []*SparseNamespaceType

// Identity returns the identity of the objects in the list.
func (o SparseNamespaceTypesList) Identity() elemental.Identity {

	return NamespaceTypeIdentity
}

// Copy returns a pointer to a copy the SparseNamespaceTypesList.
func (o SparseNamespaceTypesList) Copy() elemental.Identifiables {

	copy := append(SparseNamespaceTypesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseNamespaceTypesList.
func (o SparseNamespaceTypesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseNamespaceTypesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseNamespaceType))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseNamespaceTypesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseNamespaceTypesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseNamespaceTypesList converted to NamespaceTypesList.
func (o SparseNamespaceTypesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseNamespaceTypesList) Version() int {

	return 1
}

// SparseNamespaceType represents the sparse version of a namespacetype.
type SparseNamespaceType struct {
	// the namespace type for the current namespace.
	Type *string `json:"type,omitempty" msgpack:"type,omitempty" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseNamespaceType returns a new  SparseNamespaceType.
func NewSparseNamespaceType() *SparseNamespaceType {
	return &SparseNamespaceType{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseNamespaceType) Identity() elemental.Identity {

	return NamespaceTypeIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseNamespaceType) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseNamespaceType) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseNamespaceType) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseNamespaceType{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseNamespaceType) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseNamespaceType{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseNamespaceType) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseNamespaceType) ToPlain() elemental.PlainIdentifiable {

	out := NewNamespaceType()
	if o.Type != nil {
		out.Type = *o.Type
	}

	return out
}

// DeepCopy returns a deep copy if the SparseNamespaceType.
func (o *SparseNamespaceType) DeepCopy() *SparseNamespaceType {

	if o == nil {
		return nil
	}

	out := &SparseNamespaceType{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseNamespaceType.
func (o *SparseNamespaceType) DeepCopyInto(out *SparseNamespaceType) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseNamespaceType: %s", err))
	}

	*out = *target.(*SparseNamespaceType)
}

type mongoAttributesNamespaceType struct {
}
type mongoAttributesSparseNamespaceType struct {
}
