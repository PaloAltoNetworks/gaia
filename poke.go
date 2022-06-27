package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PokeIdentity represents the Identity of the object.
var PokeIdentity = elemental.Identity{
	Name:     "poke",
	Category: "poke",
	Package:  "squall",
	Private:  false,
}

// PokesList represents a list of Pokes
type PokesList []*Poke

// Identity returns the identity of the objects in the list.
func (o PokesList) Identity() elemental.Identity {

	return PokeIdentity
}

// Copy returns a pointer to a copy the PokesList.
func (o PokesList) Copy() elemental.Identifiables {

	copy := append(PokesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PokesList.
func (o PokesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PokesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Poke))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PokesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PokesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PokesList converted to SparsePokesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PokesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePokesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePoke)
	}

	return out
}

// Version returns the version of the content.
func (o PokesList) Version() int {

	return 1
}

// Poke represents the model of a poke
type Poke struct {
	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPoke returns a new *Poke
func NewPoke() *Poke {

	return &Poke{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Poke) Identity() elemental.Identity {

	return PokeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Poke) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Poke) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Poke) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPoke{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Poke) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPoke{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Poke) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Poke) BleveType() string {

	return "poke"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Poke) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Poke) Doc() string {

	return `When available, poke can be used to update various information about the parent.
For instance, for enforcers, poke will be used as the heartbeat.`
}

func (o *Poke) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Poke) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePoke{}
	}

	sp := &SparsePoke{}
	for _, f := range fields {
		switch f {
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePoke to the object.
func (o *Poke) Patch(sparse elemental.SparseIdentifiable) {
}

// DeepCopy returns a deep copy if the Poke.
func (o *Poke) DeepCopy() *Poke {

	if o == nil {
		return nil
	}

	out := &Poke{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Poke.
func (o *Poke) DeepCopyInto(out *Poke) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Poke: %s", err))
	}

	*out = *target.(*Poke)
}

// Validate valides the current information stored into the structure.
func (o *Poke) Validate() error {

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
func (*Poke) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PokeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PokeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Poke) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PokeAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Poke) ValueForAttribute(name string) interface{} {

	switch name {
	}

	return nil
}

// PokeAttributesMap represents the map of attribute for Poke.
var PokeAttributesMap = map[string]elemental.AttributeSpecification{}

// PokeLowerCaseAttributesMap represents the map of attribute for Poke.
var PokeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{}

// SparsePokesList represents a list of SparsePokes
type SparsePokesList []*SparsePoke

// Identity returns the identity of the objects in the list.
func (o SparsePokesList) Identity() elemental.Identity {

	return PokeIdentity
}

// Copy returns a pointer to a copy the SparsePokesList.
func (o SparsePokesList) Copy() elemental.Identifiables {

	copy := append(SparsePokesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePokesList.
func (o SparsePokesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePokesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePoke))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePokesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePokesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePokesList converted to PokesList.
func (o SparsePokesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePokesList) Version() int {

	return 1
}

// SparsePoke represents the sparse version of a poke.
type SparsePoke struct {
	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePoke returns a new  SparsePoke.
func NewSparsePoke() *SparsePoke {
	return &SparsePoke{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePoke) Identity() elemental.Identity {

	return PokeIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePoke) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePoke) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePoke) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePoke{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePoke) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePoke{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparsePoke) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePoke) ToPlain() elemental.PlainIdentifiable {

	out := NewPoke()

	return out
}

// DeepCopy returns a deep copy if the SparsePoke.
func (o *SparsePoke) DeepCopy() *SparsePoke {

	if o == nil {
		return nil
	}

	out := &SparsePoke{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePoke.
func (o *SparsePoke) DeepCopyInto(out *SparsePoke) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePoke: %s", err))
	}

	*out = *target.(*SparsePoke)
}

type mongoAttributesPoke struct {
}
type mongoAttributesSparsePoke struct {
}
