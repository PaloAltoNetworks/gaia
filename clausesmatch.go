package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ClauseMatchIdentity represents the Identity of the object.
var ClauseMatchIdentity = elemental.Identity{
	Name:     "clausesmatch",
	Category: "clausesmatches",
	Package:  "squall",
	Private:  false,
}

// ClauseMatchesList represents a list of ClauseMatches
type ClauseMatchesList []*ClauseMatch

// Identity returns the identity of the objects in the list.
func (o ClauseMatchesList) Identity() elemental.Identity {

	return ClauseMatchIdentity
}

// Copy returns a pointer to a copy the ClauseMatchesList.
func (o ClauseMatchesList) Copy() elemental.Identifiables {

	copy := append(ClauseMatchesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ClauseMatchesList.
func (o ClauseMatchesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ClauseMatchesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ClauseMatch))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ClauseMatchesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ClauseMatchesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ClauseMatchesList converted to SparseClauseMatchesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ClauseMatchesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseClauseMatchesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseClauseMatch)
	}

	return out
}

// Version returns the version of the content.
func (o ClauseMatchesList) Version() int {

	return 1
}

// ClauseMatch represents the model of a clausesmatch
type ClauseMatch struct {
	// The tag clause to resolve.
	Clauses [][]string `json:"clauses" msgpack:"clauses" bson:"clauses" mapstructure:"clauses,omitempty"`

	// Contains the matched objects.
	Match []map[string]interface{} `json:"match" msgpack:"match" bson:"-" mapstructure:"match,omitempty"`

	// The identity to render the clauses from.
	TargetIdentity string `json:"targetIdentity" msgpack:"targetIdentity" bson:"-" mapstructure:"targetIdentity,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewClauseMatch returns a new *ClauseMatch
func NewClauseMatch() *ClauseMatch {

	return &ClauseMatch{
		ModelVersion: 1,
		Clauses:      [][]string{},
		Match:        []map[string]interface{}{},
	}
}

// Identity returns the Identity of the object.
func (o *ClauseMatch) Identity() elemental.Identity {

	return ClauseMatchIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ClauseMatch) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ClauseMatch) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ClauseMatch) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesClauseMatch{}

	s.Clauses = o.Clauses

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ClauseMatch) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesClauseMatch{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Clauses = s.Clauses

	return nil
}

// Version returns the hardcoded version of the model.
func (o *ClauseMatch) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *ClauseMatch) BleveType() string {

	return "clausesmatch"
}

// DefaultOrder returns the list of default ordering fields.
func (o *ClauseMatch) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *ClauseMatch) Doc() string {

	return `This API allows to pass a set of tags and find the objects that would match the
clause in a policy resolution.`
}

func (o *ClauseMatch) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ClauseMatch) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseClauseMatch{
			Clauses:        &o.Clauses,
			Match:          &o.Match,
			TargetIdentity: &o.TargetIdentity,
		}
	}

	sp := &SparseClauseMatch{}
	for _, f := range fields {
		switch f {
		case "clauses":
			sp.Clauses = &(o.Clauses)
		case "match":
			sp.Match = &(o.Match)
		case "targetIdentity":
			sp.TargetIdentity = &(o.TargetIdentity)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseClauseMatch to the object.
func (o *ClauseMatch) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseClauseMatch)
	if so.Clauses != nil {
		o.Clauses = *so.Clauses
	}
	if so.Match != nil {
		o.Match = *so.Match
	}
	if so.TargetIdentity != nil {
		o.TargetIdentity = *so.TargetIdentity
	}
}

// DeepCopy returns a deep copy if the ClauseMatch.
func (o *ClauseMatch) DeepCopy() *ClauseMatch {

	if o == nil {
		return nil
	}

	out := &ClauseMatch{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ClauseMatch.
func (o *ClauseMatch) DeepCopyInto(out *ClauseMatch) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ClauseMatch: %s", err))
	}

	*out = *target.(*ClauseMatch)
}

// Validate valides the current information stored into the structure.
func (o *ClauseMatch) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("clauses", o.Clauses); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := ValidateTagsExpression("clauses", o.Clauses); err != nil {
		errors = errors.Append(err)
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
func (*ClauseMatch) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ClauseMatchAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ClauseMatchLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ClauseMatch) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ClauseMatchAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ClauseMatch) ValueForAttribute(name string) interface{} {

	switch name {
	case "clauses":
		return o.Clauses
	case "match":
		return o.Match
	case "targetIdentity":
		return o.TargetIdentity
	}

	return nil
}

// ClauseMatchAttributesMap represents the map of attribute for ClauseMatch.
var ClauseMatchAttributesMap = map[string]elemental.AttributeSpecification{
	"Clauses": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Clauses",
		Description:    `The tag clause to resolve.`,
		Exposed:        true,
		Name:           "clauses",
		Required:       true,
		Stored:         true,
		SubType:        "[][]string",
		Type:           "external",
	},
	"Match": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Match",
		Description:    `Contains the matched objects.`,
		Exposed:        true,
		Name:           "match",
		ReadOnly:       true,
		SubType:        "[]map[string]interface{}",
		Type:           "external",
	},
	"TargetIdentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetIdentity",
		Description:    `The identity to render the clauses from.`,
		Exposed:        true,
		Name:           "targetIdentity",
		Required:       true,
		Type:           "string",
	},
}

// ClauseMatchLowerCaseAttributesMap represents the map of attribute for ClauseMatch.
var ClauseMatchLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"clauses": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Clauses",
		Description:    `The tag clause to resolve.`,
		Exposed:        true,
		Name:           "clauses",
		Required:       true,
		Stored:         true,
		SubType:        "[][]string",
		Type:           "external",
	},
	"match": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Match",
		Description:    `Contains the matched objects.`,
		Exposed:        true,
		Name:           "match",
		ReadOnly:       true,
		SubType:        "[]map[string]interface{}",
		Type:           "external",
	},
	"targetidentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetIdentity",
		Description:    `The identity to render the clauses from.`,
		Exposed:        true,
		Name:           "targetIdentity",
		Required:       true,
		Type:           "string",
	},
}

// SparseClauseMatchesList represents a list of SparseClauseMatches
type SparseClauseMatchesList []*SparseClauseMatch

// Identity returns the identity of the objects in the list.
func (o SparseClauseMatchesList) Identity() elemental.Identity {

	return ClauseMatchIdentity
}

// Copy returns a pointer to a copy the SparseClauseMatchesList.
func (o SparseClauseMatchesList) Copy() elemental.Identifiables {

	copy := append(SparseClauseMatchesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseClauseMatchesList.
func (o SparseClauseMatchesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseClauseMatchesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseClauseMatch))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseClauseMatchesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseClauseMatchesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseClauseMatchesList converted to ClauseMatchesList.
func (o SparseClauseMatchesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseClauseMatchesList) Version() int {

	return 1
}

// SparseClauseMatch represents the sparse version of a clausesmatch.
type SparseClauseMatch struct {
	// The tag clause to resolve.
	Clauses *[][]string `json:"clauses,omitempty" msgpack:"clauses,omitempty" bson:"clauses,omitempty" mapstructure:"clauses,omitempty"`

	// Contains the matched objects.
	Match *[]map[string]interface{} `json:"match,omitempty" msgpack:"match,omitempty" bson:"-" mapstructure:"match,omitempty"`

	// The identity to render the clauses from.
	TargetIdentity *string `json:"targetIdentity,omitempty" msgpack:"targetIdentity,omitempty" bson:"-" mapstructure:"targetIdentity,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseClauseMatch returns a new  SparseClauseMatch.
func NewSparseClauseMatch() *SparseClauseMatch {
	return &SparseClauseMatch{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseClauseMatch) Identity() elemental.Identity {

	return ClauseMatchIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseClauseMatch) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseClauseMatch) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseClauseMatch) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseClauseMatch{}

	if o.Clauses != nil {
		s.Clauses = o.Clauses
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseClauseMatch) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseClauseMatch{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Clauses != nil {
		o.Clauses = s.Clauses
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseClauseMatch) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseClauseMatch) ToPlain() elemental.PlainIdentifiable {

	out := NewClauseMatch()
	if o.Clauses != nil {
		out.Clauses = *o.Clauses
	}
	if o.Match != nil {
		out.Match = *o.Match
	}
	if o.TargetIdentity != nil {
		out.TargetIdentity = *o.TargetIdentity
	}

	return out
}

// DeepCopy returns a deep copy if the SparseClauseMatch.
func (o *SparseClauseMatch) DeepCopy() *SparseClauseMatch {

	if o == nil {
		return nil
	}

	out := &SparseClauseMatch{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseClauseMatch.
func (o *SparseClauseMatch) DeepCopyInto(out *SparseClauseMatch) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseClauseMatch: %s", err))
	}

	*out = *target.(*SparseClauseMatch)
}

type mongoAttributesClauseMatch struct {
	Clauses [][]string `bson:"clauses"`
}
type mongoAttributesSparseClauseMatch struct {
	Clauses *[][]string `bson:"clauses,omitempty"`
}
