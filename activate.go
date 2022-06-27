package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ActivateIdentity represents the Identity of the object.
var ActivateIdentity = elemental.Identity{
	Name:     "activate",
	Category: "activate",
	Package:  "vince",
	Private:  false,
}

// ActivatesList represents a list of Activates
type ActivatesList []*Activate

// Identity returns the identity of the objects in the list.
func (o ActivatesList) Identity() elemental.Identity {

	return ActivateIdentity
}

// Copy returns a pointer to a copy the ActivatesList.
func (o ActivatesList) Copy() elemental.Identifiables {

	copy := append(ActivatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ActivatesList.
func (o ActivatesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ActivatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Activate))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ActivatesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ActivatesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ActivatesList converted to SparseActivatesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ActivatesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseActivatesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseActivate)
	}

	return out
}

// Version returns the version of the content.
func (o ActivatesList) Version() int {

	return 1
}

// Activate represents the model of a activate
type Activate struct {
	// Contains the activation token.
	Token string `json:"token" msgpack:"token" bson:"-" mapstructure:"token,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewActivate returns a new *Activate
func NewActivate() *Activate {

	return &Activate{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Activate) Identity() elemental.Identity {

	return ActivateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Activate) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Activate) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Activate) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesActivate{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Activate) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesActivate{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Activate) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Activate) BleveType() string {

	return "activate"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Activate) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Activate) Doc() string {

	return `Used to activate a pending account.`
}

func (o *Activate) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Activate) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseActivate{
			Token: &o.Token,
		}
	}

	sp := &SparseActivate{}
	for _, f := range fields {
		switch f {
		case "token":
			sp.Token = &(o.Token)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseActivate to the object.
func (o *Activate) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseActivate)
	if so.Token != nil {
		o.Token = *so.Token
	}
}

// DeepCopy returns a deep copy if the Activate.
func (o *Activate) DeepCopy() *Activate {

	if o == nil {
		return nil
	}

	out := &Activate{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Activate.
func (o *Activate) DeepCopyInto(out *Activate) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Activate: %s", err))
	}

	*out = *target.(*Activate)
}

// Validate valides the current information stored into the structure.
func (o *Activate) Validate() error {

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
func (*Activate) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ActivateAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ActivateLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Activate) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ActivateAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Activate) ValueForAttribute(name string) interface{} {

	switch name {
	case "token":
		return o.Token
	}

	return nil
}

// ActivateAttributesMap represents the map of attribute for Activate.
var ActivateAttributesMap = map[string]elemental.AttributeSpecification{
	"Token": {
		AllowedChoices: []string{},
		ConvertedName:  "Token",
		CreationOnly:   true,
		Description:    `Contains the activation token.`,
		Exposed:        true,
		Name:           "token",
		Type:           "string",
	},
}

// ActivateLowerCaseAttributesMap represents the map of attribute for Activate.
var ActivateLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"token": {
		AllowedChoices: []string{},
		ConvertedName:  "Token",
		CreationOnly:   true,
		Description:    `Contains the activation token.`,
		Exposed:        true,
		Name:           "token",
		Type:           "string",
	},
}

// SparseActivatesList represents a list of SparseActivates
type SparseActivatesList []*SparseActivate

// Identity returns the identity of the objects in the list.
func (o SparseActivatesList) Identity() elemental.Identity {

	return ActivateIdentity
}

// Copy returns a pointer to a copy the SparseActivatesList.
func (o SparseActivatesList) Copy() elemental.Identifiables {

	copy := append(SparseActivatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseActivatesList.
func (o SparseActivatesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseActivatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseActivate))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseActivatesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseActivatesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseActivatesList converted to ActivatesList.
func (o SparseActivatesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseActivatesList) Version() int {

	return 1
}

// SparseActivate represents the sparse version of a activate.
type SparseActivate struct {
	// Contains the activation token.
	Token *string `json:"token,omitempty" msgpack:"token,omitempty" bson:"-" mapstructure:"token,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseActivate returns a new  SparseActivate.
func NewSparseActivate() *SparseActivate {
	return &SparseActivate{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseActivate) Identity() elemental.Identity {

	return ActivateIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseActivate) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseActivate) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseActivate) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseActivate{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseActivate) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseActivate{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseActivate) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseActivate) ToPlain() elemental.PlainIdentifiable {

	out := NewActivate()
	if o.Token != nil {
		out.Token = *o.Token
	}

	return out
}

// DeepCopy returns a deep copy if the SparseActivate.
func (o *SparseActivate) DeepCopy() *SparseActivate {

	if o == nil {
		return nil
	}

	out := &SparseActivate{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseActivate.
func (o *SparseActivate) DeepCopyInto(out *SparseActivate) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseActivate: %s", err))
	}

	*out = *target.(*SparseActivate)
}

type mongoAttributesActivate struct {
}
type mongoAttributesSparseActivate struct {
}
