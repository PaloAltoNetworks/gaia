package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// HitIdentity represents the Identity of the object.
var HitIdentity = elemental.Identity{
	Name:     "hit",
	Category: "hits",
	Package:  "minwu",
	Private:  false,
}

// HitsList represents a list of Hits
type HitsList []*Hit

// Identity returns the identity of the objects in the list.
func (o HitsList) Identity() elemental.Identity {

	return HitIdentity
}

// Copy returns a pointer to a copy the HitsList.
func (o HitsList) Copy() elemental.Identifiables {

	copy := append(HitsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the HitsList.
func (o HitsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(HitsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Hit))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o HitsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o HitsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the HitsList converted to SparseHitsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o HitsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseHitsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseHit)
	}

	return out
}

// Version returns the version of the content.
func (o HitsList) Version() int {

	return 1
}

// Hit represents the model of a hit
type Hit struct {
	// name of the counter.
	Name string `json:"name" msgpack:"name" bson:"-" mapstructure:"name,omitempty"`

	// The ID of the referenced object..
	TargetID string `json:"targetID" msgpack:"targetID" bson:"-" mapstructure:"targetID,omitempty"`

	// The identity of the referenced object.
	TargetIdentity string `json:"targetIdentity" msgpack:"targetIdentity" bson:"targetidentity" mapstructure:"targetIdentity,omitempty"`

	// The value of the hit.
	Value int `json:"value" msgpack:"value" bson:"-" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewHit returns a new *Hit
func NewHit() *Hit {

	return &Hit{
		ModelVersion: 1,
		Name:         "counter",
	}
}

// Identity returns the Identity of the object.
func (o *Hit) Identity() elemental.Identity {

	return HitIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Hit) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Hit) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Hit) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesHit{}

	s.TargetIdentity = o.TargetIdentity

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Hit) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesHit{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.TargetIdentity = s.TargetIdentity

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Hit) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Hit) BleveType() string {

	return "hit"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Hit) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Hit) Doc() string {

	return `This API allows to retrieve a generic hit counter for a given object.`
}

func (o *Hit) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Hit) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseHit{
			Name:           &o.Name,
			TargetID:       &o.TargetID,
			TargetIdentity: &o.TargetIdentity,
			Value:          &o.Value,
		}
	}

	sp := &SparseHit{}
	for _, f := range fields {
		switch f {
		case "name":
			sp.Name = &(o.Name)
		case "targetID":
			sp.TargetID = &(o.TargetID)
		case "targetIdentity":
			sp.TargetIdentity = &(o.TargetIdentity)
		case "value":
			sp.Value = &(o.Value)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseHit to the object.
func (o *Hit) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseHit)
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.TargetID != nil {
		o.TargetID = *so.TargetID
	}
	if so.TargetIdentity != nil {
		o.TargetIdentity = *so.TargetIdentity
	}
	if so.Value != nil {
		o.Value = *so.Value
	}
}

// DeepCopy returns a deep copy if the Hit.
func (o *Hit) DeepCopy() *Hit {

	if o == nil {
		return nil
	}

	out := &Hit{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Hit.
func (o *Hit) DeepCopyInto(out *Hit) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Hit: %s", err))
	}

	*out = *target.(*Hit)
}

// Validate valides the current information stored into the structure.
func (o *Hit) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("targetIdentity", o.TargetIdentity); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := ValidateIdentity("targetIdentity", o.TargetIdentity); err != nil {
		errors = errors.Append(err)
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
func (*Hit) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := HitAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return HitLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Hit) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return HitAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Hit) ValueForAttribute(name string) interface{} {

	switch name {
	case "name":
		return o.Name
	case "targetID":
		return o.TargetID
	case "targetIdentity":
		return o.TargetIdentity
	case "value":
		return o.Value
	}

	return nil
}

// HitAttributesMap represents the map of attribute for Hit.
var HitAttributesMap = map[string]elemental.AttributeSpecification{
	"Name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultValue:   "counter",
		Description:    `name of the counter.`,
		Exposed:        true,
		Name:           "name",
		Required:       true,
		Type:           "string",
	},
	"TargetID": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetID",
		Description:    `The ID of the referenced object..`,
		Exposed:        true,
		Name:           "targetID",
		Type:           "string",
	},
	"TargetIdentity": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetIdentity",
		Description:    `The identity of the referenced object.`,
		Exposed:        true,
		Name:           "targetIdentity",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Value": {
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `The value of the hit.`,
		Exposed:        true,
		Name:           "value",
		ReadOnly:       true,
		Type:           "integer",
	},
}

// HitLowerCaseAttributesMap represents the map of attribute for Hit.
var HitLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultValue:   "counter",
		Description:    `name of the counter.`,
		Exposed:        true,
		Name:           "name",
		Required:       true,
		Type:           "string",
	},
	"targetid": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetID",
		Description:    `The ID of the referenced object..`,
		Exposed:        true,
		Name:           "targetID",
		Type:           "string",
	},
	"targetidentity": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetIdentity",
		Description:    `The identity of the referenced object.`,
		Exposed:        true,
		Name:           "targetIdentity",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"value": {
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `The value of the hit.`,
		Exposed:        true,
		Name:           "value",
		ReadOnly:       true,
		Type:           "integer",
	},
}

// SparseHitsList represents a list of SparseHits
type SparseHitsList []*SparseHit

// Identity returns the identity of the objects in the list.
func (o SparseHitsList) Identity() elemental.Identity {

	return HitIdentity
}

// Copy returns a pointer to a copy the SparseHitsList.
func (o SparseHitsList) Copy() elemental.Identifiables {

	copy := append(SparseHitsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseHitsList.
func (o SparseHitsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseHitsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseHit))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseHitsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseHitsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseHitsList converted to HitsList.
func (o SparseHitsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseHitsList) Version() int {

	return 1
}

// SparseHit represents the sparse version of a hit.
type SparseHit struct {
	// name of the counter.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"-" mapstructure:"name,omitempty"`

	// The ID of the referenced object..
	TargetID *string `json:"targetID,omitempty" msgpack:"targetID,omitempty" bson:"-" mapstructure:"targetID,omitempty"`

	// The identity of the referenced object.
	TargetIdentity *string `json:"targetIdentity,omitempty" msgpack:"targetIdentity,omitempty" bson:"targetidentity,omitempty" mapstructure:"targetIdentity,omitempty"`

	// The value of the hit.
	Value *int `json:"value,omitempty" msgpack:"value,omitempty" bson:"-" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseHit returns a new  SparseHit.
func NewSparseHit() *SparseHit {
	return &SparseHit{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseHit) Identity() elemental.Identity {

	return HitIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseHit) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseHit) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseHit) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseHit{}

	if o.TargetIdentity != nil {
		s.TargetIdentity = o.TargetIdentity
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseHit) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseHit{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.TargetIdentity != nil {
		o.TargetIdentity = s.TargetIdentity
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseHit) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseHit) ToPlain() elemental.PlainIdentifiable {

	out := NewHit()
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.TargetID != nil {
		out.TargetID = *o.TargetID
	}
	if o.TargetIdentity != nil {
		out.TargetIdentity = *o.TargetIdentity
	}
	if o.Value != nil {
		out.Value = *o.Value
	}

	return out
}

// DeepCopy returns a deep copy if the SparseHit.
func (o *SparseHit) DeepCopy() *SparseHit {

	if o == nil {
		return nil
	}

	out := &SparseHit{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseHit.
func (o *SparseHit) DeepCopyInto(out *SparseHit) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseHit: %s", err))
	}

	*out = *target.(*SparseHit)
}

type mongoAttributesHit struct {
	TargetIdentity string `bson:"targetidentity"`
}
type mongoAttributesSparseHit struct {
	TargetIdentity *string `bson:"targetidentity,omitempty"`
}
