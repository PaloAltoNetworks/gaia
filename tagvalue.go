package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TagValueIdentity represents the Identity of the object.
var TagValueIdentity = elemental.Identity{
	Name:     "tagvalue",
	Category: "tagvalues",
	Package:  "tagle",
	Private:  false,
}

// TagValuesList represents a list of TagValues
type TagValuesList []*TagValue

// Identity returns the identity of the objects in the list.
func (o TagValuesList) Identity() elemental.Identity {

	return TagValueIdentity
}

// Copy returns a pointer to a copy the TagValuesList.
func (o TagValuesList) Copy() elemental.Identifiables {

	copy := append(TagValuesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TagValuesList.
func (o TagValuesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(TagValuesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*TagValue))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TagValuesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TagValuesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the TagValuesList converted to SparseTagValuesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o TagValuesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseTagValuesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseTagValue)
	}

	return out
}

// Version returns the version of the content.
func (o TagValuesList) Version() int {

	return 1
}

// TagValue represents the model of a tagvalue
type TagValue struct {
	// The requested key.
	Key string `json:"key" msgpack:"key" bson:"-" mapstructure:"key,omitempty"`

	// List of all values.
	Values []string `json:"values" msgpack:"values" bson:"-" mapstructure:"values,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewTagValue returns a new *TagValue
func NewTagValue() *TagValue {

	return &TagValue{
		ModelVersion: 1,
		Values:       []string{},
	}
}

// Identity returns the Identity of the object.
func (o *TagValue) Identity() elemental.Identity {

	return TagValueIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *TagValue) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *TagValue) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TagValue) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesTagValue{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TagValue) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesTagValue{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *TagValue) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *TagValue) BleveType() string {

	return "tagvalue"
}

// DefaultOrder returns the list of default ordering fields.
func (o *TagValue) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *TagValue) Doc() string {

	return `Represents all values associated to a tag key.`
}

func (o *TagValue) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *TagValue) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseTagValue{
			Key:    &o.Key,
			Values: &o.Values,
		}
	}

	sp := &SparseTagValue{}
	for _, f := range fields {
		switch f {
		case "key":
			sp.Key = &(o.Key)
		case "values":
			sp.Values = &(o.Values)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseTagValue to the object.
func (o *TagValue) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseTagValue)
	if so.Key != nil {
		o.Key = *so.Key
	}
	if so.Values != nil {
		o.Values = *so.Values
	}
}

// DeepCopy returns a deep copy if the TagValue.
func (o *TagValue) DeepCopy() *TagValue {

	if o == nil {
		return nil
	}

	out := &TagValue{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *TagValue.
func (o *TagValue) DeepCopyInto(out *TagValue) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy TagValue: %s", err))
	}

	*out = *target.(*TagValue)
}

// Validate valides the current information stored into the structure.
func (o *TagValue) Validate() error {

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
func (*TagValue) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TagValueAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TagValueLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*TagValue) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TagValueAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *TagValue) ValueForAttribute(name string) interface{} {

	switch name {
	case "key":
		return o.Key
	case "values":
		return o.Values
	}

	return nil
}

// TagValueAttributesMap represents the map of attribute for TagValue.
var TagValueAttributesMap = map[string]elemental.AttributeSpecification{
	"Key": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Key",
		Description:    `The requested key.`,
		Exposed:        true,
		Name:           "key",
		ReadOnly:       true,
		Type:           "string",
	},
	"Values": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Values",
		Description:    `List of all values.`,
		Exposed:        true,
		Name:           "values",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
}

// TagValueLowerCaseAttributesMap represents the map of attribute for TagValue.
var TagValueLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"key": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Key",
		Description:    `The requested key.`,
		Exposed:        true,
		Name:           "key",
		ReadOnly:       true,
		Type:           "string",
	},
	"values": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Values",
		Description:    `List of all values.`,
		Exposed:        true,
		Name:           "values",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
}

// SparseTagValuesList represents a list of SparseTagValues
type SparseTagValuesList []*SparseTagValue

// Identity returns the identity of the objects in the list.
func (o SparseTagValuesList) Identity() elemental.Identity {

	return TagValueIdentity
}

// Copy returns a pointer to a copy the SparseTagValuesList.
func (o SparseTagValuesList) Copy() elemental.Identifiables {

	copy := append(SparseTagValuesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseTagValuesList.
func (o SparseTagValuesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseTagValuesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseTagValue))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseTagValuesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseTagValuesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseTagValuesList converted to TagValuesList.
func (o SparseTagValuesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseTagValuesList) Version() int {

	return 1
}

// SparseTagValue represents the sparse version of a tagvalue.
type SparseTagValue struct {
	// The requested key.
	Key *string `json:"key,omitempty" msgpack:"key,omitempty" bson:"-" mapstructure:"key,omitempty"`

	// List of all values.
	Values *[]string `json:"values,omitempty" msgpack:"values,omitempty" bson:"-" mapstructure:"values,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseTagValue returns a new  SparseTagValue.
func NewSparseTagValue() *SparseTagValue {
	return &SparseTagValue{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseTagValue) Identity() elemental.Identity {

	return TagValueIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseTagValue) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseTagValue) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTagValue) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseTagValue{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTagValue) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseTagValue{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseTagValue) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseTagValue) ToPlain() elemental.PlainIdentifiable {

	out := NewTagValue()
	if o.Key != nil {
		out.Key = *o.Key
	}
	if o.Values != nil {
		out.Values = *o.Values
	}

	return out
}

// DeepCopy returns a deep copy if the SparseTagValue.
func (o *SparseTagValue) DeepCopy() *SparseTagValue {

	if o == nil {
		return nil
	}

	out := &SparseTagValue{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseTagValue.
func (o *SparseTagValue) DeepCopyInto(out *SparseTagValue) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseTagValue: %s", err))
	}

	*out = *target.(*SparseTagValue)
}

type mongoAttributesTagValue struct {
}
type mongoAttributesSparseTagValue struct {
}
