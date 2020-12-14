package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PCTimeRangeIdentity represents the Identity of the object.
var PCTimeRangeIdentity = elemental.Identity{
	Name:     "pctimerange",
	Category: "pctimeranges",
	Package:  "karl",
	Private:  false,
}

// PCTimeRangesList represents a list of PCTimeRanges
type PCTimeRangesList []*PCTimeRange

// Identity returns the identity of the objects in the list.
func (o PCTimeRangesList) Identity() elemental.Identity {

	return PCTimeRangeIdentity
}

// Copy returns a pointer to a copy the PCTimeRangesList.
func (o PCTimeRangesList) Copy() elemental.Identifiables {

	copy := append(PCTimeRangesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PCTimeRangesList.
func (o PCTimeRangesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PCTimeRangesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PCTimeRange))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PCTimeRangesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PCTimeRangesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PCTimeRangesList converted to SparsePCTimeRangesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PCTimeRangesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePCTimeRangesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePCTimeRange)
	}

	return out
}

// Version returns the version of the content.
func (o PCTimeRangesList) Version() int {

	return 1
}

// PCTimeRange represents the model of a pctimerange
type PCTimeRange struct {
	// Type of the time range.
	Type string `json:"type" msgpack:"type" bson:"-" mapstructure:"type,omitempty"`

	// Value of the time range.
	Value *PCTimeRangeValue `json:"value" msgpack:"value" bson:"-" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPCTimeRange returns a new *PCTimeRange
func NewPCTimeRange() *PCTimeRange {

	return &PCTimeRange{
		ModelVersion: 1,
		Value:        NewPCTimeRangeValue(),
	}
}

// Identity returns the Identity of the object.
func (o *PCTimeRange) Identity() elemental.Identity {

	return PCTimeRangeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PCTimeRange) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PCTimeRange) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PCTimeRange) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPCTimeRange{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PCTimeRange) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPCTimeRange{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *PCTimeRange) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *PCTimeRange) BleveType() string {

	return "pctimerange"
}

// DefaultOrder returns the list of default ordering fields.
func (o *PCTimeRange) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PCTimeRange) Doc() string {

	return `Represents the time range option of rql search request.`
}

func (o *PCTimeRange) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PCTimeRange) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePCTimeRange{
			Type:  &o.Type,
			Value: o.Value,
		}
	}

	sp := &SparsePCTimeRange{}
	for _, f := range fields {
		switch f {
		case "type":
			sp.Type = &(o.Type)
		case "value":
			sp.Value = o.Value
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePCTimeRange to the object.
func (o *PCTimeRange) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePCTimeRange)
	if so.Type != nil {
		o.Type = *so.Type
	}
	if so.Value != nil {
		o.Value = so.Value
	}
}

// DeepCopy returns a deep copy if the PCTimeRange.
func (o *PCTimeRange) DeepCopy() *PCTimeRange {

	if o == nil {
		return nil
	}

	out := &PCTimeRange{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PCTimeRange.
func (o *PCTimeRange) DeepCopyInto(out *PCTimeRange) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PCTimeRange: %s", err))
	}

	*out = *target.(*PCTimeRange)
}

// Validate valides the current information stored into the structure.
func (o *PCTimeRange) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if o.Value != nil {
		elemental.ResetDefaultForZeroValues(o.Value)
		if err := o.Value.Validate(); err != nil {
			errors = errors.Append(err)
		}
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
func (*PCTimeRange) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PCTimeRangeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PCTimeRangeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PCTimeRange) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PCTimeRangeAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PCTimeRange) ValueForAttribute(name string) interface{} {

	switch name {
	case "type":
		return o.Type
	case "value":
		return o.Value
	}

	return nil
}

// PCTimeRangeAttributesMap represents the map of attribute for PCTimeRange.
var PCTimeRangeAttributesMap = map[string]elemental.AttributeSpecification{
	"Type": {
		AllowedChoices: []string{},
		ConvertedName:  "Type",
		Description:    `Type of the time range.`,
		Exposed:        true,
		Name:           "type",
		Type:           "string",
	},
	"Value": {
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Value of the time range.`,
		Exposed:        true,
		Name:           "value",
		SubType:        "pctimerangevalue",
		Type:           "ref",
	},
}

// PCTimeRangeLowerCaseAttributesMap represents the map of attribute for PCTimeRange.
var PCTimeRangeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"type": {
		AllowedChoices: []string{},
		ConvertedName:  "Type",
		Description:    `Type of the time range.`,
		Exposed:        true,
		Name:           "type",
		Type:           "string",
	},
	"value": {
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Value of the time range.`,
		Exposed:        true,
		Name:           "value",
		SubType:        "pctimerangevalue",
		Type:           "ref",
	},
}

// SparsePCTimeRangesList represents a list of SparsePCTimeRanges
type SparsePCTimeRangesList []*SparsePCTimeRange

// Identity returns the identity of the objects in the list.
func (o SparsePCTimeRangesList) Identity() elemental.Identity {

	return PCTimeRangeIdentity
}

// Copy returns a pointer to a copy the SparsePCTimeRangesList.
func (o SparsePCTimeRangesList) Copy() elemental.Identifiables {

	copy := append(SparsePCTimeRangesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePCTimeRangesList.
func (o SparsePCTimeRangesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePCTimeRangesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePCTimeRange))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePCTimeRangesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePCTimeRangesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePCTimeRangesList converted to PCTimeRangesList.
func (o SparsePCTimeRangesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePCTimeRangesList) Version() int {

	return 1
}

// SparsePCTimeRange represents the sparse version of a pctimerange.
type SparsePCTimeRange struct {
	// Type of the time range.
	Type *string `json:"type,omitempty" msgpack:"type,omitempty" bson:"-" mapstructure:"type,omitempty"`

	// Value of the time range.
	Value *PCTimeRangeValue `json:"value,omitempty" msgpack:"value,omitempty" bson:"-" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePCTimeRange returns a new  SparsePCTimeRange.
func NewSparsePCTimeRange() *SparsePCTimeRange {
	return &SparsePCTimeRange{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePCTimeRange) Identity() elemental.Identity {

	return PCTimeRangeIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePCTimeRange) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePCTimeRange) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePCTimeRange) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePCTimeRange{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePCTimeRange) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePCTimeRange{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparsePCTimeRange) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePCTimeRange) ToPlain() elemental.PlainIdentifiable {

	out := NewPCTimeRange()
	if o.Type != nil {
		out.Type = *o.Type
	}
	if o.Value != nil {
		out.Value = o.Value
	}

	return out
}

// DeepCopy returns a deep copy if the SparsePCTimeRange.
func (o *SparsePCTimeRange) DeepCopy() *SparsePCTimeRange {

	if o == nil {
		return nil
	}

	out := &SparsePCTimeRange{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePCTimeRange.
func (o *SparsePCTimeRange) DeepCopyInto(out *SparsePCTimeRange) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePCTimeRange: %s", err))
	}

	*out = *target.(*SparsePCTimeRange)
}

type mongoAttributesPCTimeRange struct {
}
type mongoAttributesSparsePCTimeRange struct {
}
