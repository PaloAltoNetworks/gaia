package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TriggerIdentity represents the Identity of the object.
var TriggerIdentity = elemental.Identity{
	Name:     "trigger",
	Category: "triggers",
	Package:  "sephiroth",
	Private:  false,
}

// TriggersList represents a list of Triggers
type TriggersList []*Trigger

// Identity returns the identity of the objects in the list.
func (o TriggersList) Identity() elemental.Identity {

	return TriggerIdentity
}

// Copy returns a pointer to a copy the TriggersList.
func (o TriggersList) Copy() elemental.Identifiables {

	copy := append(TriggersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TriggersList.
func (o TriggersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(TriggersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Trigger))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TriggersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TriggersList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the TriggersList converted to SparseTriggersList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o TriggersList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseTriggersList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseTrigger)
	}

	return out
}

// Version returns the version of the content.
func (o TriggersList) Version() int {

	return 1
}

// Trigger represents the model of a trigger
type Trigger struct {
	// Contains the remote `POST` payload.
	Payload string `json:"-" msgpack:"-" bson:"-" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewTrigger returns a new *Trigger
func NewTrigger() *Trigger {

	return &Trigger{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Trigger) Identity() elemental.Identity {

	return TriggerIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Trigger) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Trigger) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Trigger) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesTrigger{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Trigger) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesTrigger{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Trigger) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Trigger) BleveType() string {

	return "trigger"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Trigger) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Trigger) Doc() string {

	return `Can be used to remotely trigger an automation.`
}

func (o *Trigger) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Trigger) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseTrigger{
			Payload: &o.Payload,
		}
	}

	sp := &SparseTrigger{}
	for _, f := range fields {
		switch f {
		case "payload":
			sp.Payload = &(o.Payload)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseTrigger to the object.
func (o *Trigger) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseTrigger)
	if so.Payload != nil {
		o.Payload = *so.Payload
	}
}

// DeepCopy returns a deep copy if the Trigger.
func (o *Trigger) DeepCopy() *Trigger {

	if o == nil {
		return nil
	}

	out := &Trigger{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Trigger.
func (o *Trigger) DeepCopyInto(out *Trigger) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Trigger: %s", err))
	}

	*out = *target.(*Trigger)
}

// Validate valides the current information stored into the structure.
func (o *Trigger) Validate() error {

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
func (*Trigger) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TriggerAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TriggerLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Trigger) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TriggerAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Trigger) ValueForAttribute(name string) interface{} {

	switch name {
	case "payload":
		return o.Payload
	}

	return nil
}

// TriggerAttributesMap represents the map of attribute for Trigger.
var TriggerAttributesMap = map[string]elemental.AttributeSpecification{
	"Payload": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Payload",
		Description:    `Contains the remote ` + "`" + `POST` + "`" + ` payload.`,
		Name:           "payload",
		Type:           "string",
	},
}

// TriggerLowerCaseAttributesMap represents the map of attribute for Trigger.
var TriggerLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"payload": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Payload",
		Description:    `Contains the remote ` + "`" + `POST` + "`" + ` payload.`,
		Name:           "payload",
		Type:           "string",
	},
}

// SparseTriggersList represents a list of SparseTriggers
type SparseTriggersList []*SparseTrigger

// Identity returns the identity of the objects in the list.
func (o SparseTriggersList) Identity() elemental.Identity {

	return TriggerIdentity
}

// Copy returns a pointer to a copy the SparseTriggersList.
func (o SparseTriggersList) Copy() elemental.Identifiables {

	copy := append(SparseTriggersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseTriggersList.
func (o SparseTriggersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseTriggersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseTrigger))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseTriggersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseTriggersList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseTriggersList converted to TriggersList.
func (o SparseTriggersList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseTriggersList) Version() int {

	return 1
}

// SparseTrigger represents the sparse version of a trigger.
type SparseTrigger struct {
	// Contains the remote `POST` payload.
	Payload *string `json:"-" msgpack:"-" bson:"-" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseTrigger returns a new  SparseTrigger.
func NewSparseTrigger() *SparseTrigger {
	return &SparseTrigger{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseTrigger) Identity() elemental.Identity {

	return TriggerIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseTrigger) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseTrigger) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTrigger) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseTrigger{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTrigger) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseTrigger{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseTrigger) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseTrigger) ToPlain() elemental.PlainIdentifiable {

	out := NewTrigger()
	if o.Payload != nil {
		out.Payload = *o.Payload
	}

	return out
}

// DeepCopy returns a deep copy if the SparseTrigger.
func (o *SparseTrigger) DeepCopy() *SparseTrigger {

	if o == nil {
		return nil
	}

	out := &SparseTrigger{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseTrigger.
func (o *SparseTrigger) DeepCopyInto(out *SparseTrigger) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseTrigger: %s", err))
	}

	*out = *target.(*SparseTrigger)
}

type mongoAttributesTrigger struct {
}
type mongoAttributesSparseTrigger struct {
}
