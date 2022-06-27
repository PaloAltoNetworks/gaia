package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ReputationIdentity represents the Identity of the object.
var ReputationIdentity = elemental.Identity{
	Name:     "reputation",
	Category: "reputations",
	Package:  "karl",
	Private:  false,
}

// ReputationsList represents a list of Reputations
type ReputationsList []*Reputation

// Identity returns the identity of the objects in the list.
func (o ReputationsList) Identity() elemental.Identity {

	return ReputationIdentity
}

// Copy returns a pointer to a copy the ReputationsList.
func (o ReputationsList) Copy() elemental.Identifiables {

	copy := append(ReputationsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ReputationsList.
func (o ReputationsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ReputationsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Reputation))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ReputationsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ReputationsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ReputationsList converted to SparseReputationsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ReputationsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseReputationsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseReputation)
	}

	return out
}

// Version returns the version of the content.
func (o ReputationsList) Version() int {

	return 1
}

// Reputation represents the model of a reputation
type Reputation struct {
	// List of URL reputations.
	URLReputations []*URLReputation `json:"URLReputations,omitempty" msgpack:"URLReputations,omitempty" bson:"-" mapstructure:"URLReputations,omitempty"`

	// The IP addresses and/or FQDNs to look up.
	URLs []string `json:"URLs" msgpack:"URLs" bson:"-" mapstructure:"URLs,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewReputation returns a new *Reputation
func NewReputation() *Reputation {

	return &Reputation{
		ModelVersion:   1,
		URLReputations: []*URLReputation{},
		URLs:           []string{},
	}
}

// Identity returns the Identity of the object.
func (o *Reputation) Identity() elemental.Identity {

	return ReputationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Reputation) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Reputation) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Reputation) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesReputation{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Reputation) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesReputation{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Reputation) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Reputation) BleveType() string {

	return "reputation"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Reputation) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Reputation) Doc() string {

	return `Represents the reputation information from PANDB.`
}

func (o *Reputation) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Reputation) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseReputation{
			URLReputations: &o.URLReputations,
			URLs:           &o.URLs,
		}
	}

	sp := &SparseReputation{}
	for _, f := range fields {
		switch f {
		case "URLReputations":
			sp.URLReputations = &(o.URLReputations)
		case "URLs":
			sp.URLs = &(o.URLs)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseReputation to the object.
func (o *Reputation) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseReputation)
	if so.URLReputations != nil {
		o.URLReputations = *so.URLReputations
	}
	if so.URLs != nil {
		o.URLs = *so.URLs
	}
}

// DeepCopy returns a deep copy if the Reputation.
func (o *Reputation) DeepCopy() *Reputation {

	if o == nil {
		return nil
	}

	out := &Reputation{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Reputation.
func (o *Reputation) DeepCopyInto(out *Reputation) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Reputation: %s", err))
	}

	*out = *target.(*Reputation)
}

// Validate valides the current information stored into the structure.
func (o *Reputation) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.URLReputations {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateRequiredExternal("URLs", o.URLs); err != nil {
		requiredErrors = requiredErrors.Append(err)
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
func (*Reputation) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ReputationAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ReputationLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Reputation) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ReputationAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Reputation) ValueForAttribute(name string) interface{} {

	switch name {
	case "URLReputations":
		return o.URLReputations
	case "URLs":
		return o.URLs
	}

	return nil
}

// ReputationAttributesMap represents the map of attribute for Reputation.
var ReputationAttributesMap = map[string]elemental.AttributeSpecification{
	"URLReputations": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "URLReputations",
		Description:    `List of URL reputations.`,
		Exposed:        true,
		Name:           "URLReputations",
		ReadOnly:       true,
		SubType:        "urlreputation",
		Type:           "refList",
	},
	"URLs": {
		AllowedChoices: []string{},
		ConvertedName:  "URLs",
		Description:    `The IP addresses and/or FQDNs to look up.`,
		Exposed:        true,
		Name:           "URLs",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
}

// ReputationLowerCaseAttributesMap represents the map of attribute for Reputation.
var ReputationLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"urlreputations": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "URLReputations",
		Description:    `List of URL reputations.`,
		Exposed:        true,
		Name:           "URLReputations",
		ReadOnly:       true,
		SubType:        "urlreputation",
		Type:           "refList",
	},
	"urls": {
		AllowedChoices: []string{},
		ConvertedName:  "URLs",
		Description:    `The IP addresses and/or FQDNs to look up.`,
		Exposed:        true,
		Name:           "URLs",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
}

// SparseReputationsList represents a list of SparseReputations
type SparseReputationsList []*SparseReputation

// Identity returns the identity of the objects in the list.
func (o SparseReputationsList) Identity() elemental.Identity {

	return ReputationIdentity
}

// Copy returns a pointer to a copy the SparseReputationsList.
func (o SparseReputationsList) Copy() elemental.Identifiables {

	copy := append(SparseReputationsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseReputationsList.
func (o SparseReputationsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseReputationsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseReputation))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseReputationsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseReputationsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseReputationsList converted to ReputationsList.
func (o SparseReputationsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseReputationsList) Version() int {

	return 1
}

// SparseReputation represents the sparse version of a reputation.
type SparseReputation struct {
	// List of URL reputations.
	URLReputations *[]*URLReputation `json:"URLReputations,omitempty" msgpack:"URLReputations,omitempty" bson:"-" mapstructure:"URLReputations,omitempty"`

	// The IP addresses and/or FQDNs to look up.
	URLs *[]string `json:"URLs,omitempty" msgpack:"URLs,omitempty" bson:"-" mapstructure:"URLs,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseReputation returns a new  SparseReputation.
func NewSparseReputation() *SparseReputation {
	return &SparseReputation{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseReputation) Identity() elemental.Identity {

	return ReputationIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseReputation) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseReputation) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseReputation) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseReputation{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseReputation) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseReputation{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseReputation) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseReputation) ToPlain() elemental.PlainIdentifiable {

	out := NewReputation()
	if o.URLReputations != nil {
		out.URLReputations = *o.URLReputations
	}
	if o.URLs != nil {
		out.URLs = *o.URLs
	}

	return out
}

// DeepCopy returns a deep copy if the SparseReputation.
func (o *SparseReputation) DeepCopy() *SparseReputation {

	if o == nil {
		return nil
	}

	out := &SparseReputation{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseReputation.
func (o *SparseReputation) DeepCopyInto(out *SparseReputation) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseReputation: %s", err))
	}

	*out = *target.(*SparseReputation)
}

type mongoAttributesReputation struct {
}
type mongoAttributesSparseReputation struct {
}
