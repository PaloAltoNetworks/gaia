package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// SearchIdentity represents the Identity of the object.
var SearchIdentity = elemental.Identity{
	Name:     "search",
	Category: "search",
	Package:  "gogole",
	Private:  false,
}

// SearchesList represents a list of Searches
type SearchesList []*Search

// Identity returns the identity of the objects in the list.
func (o SearchesList) Identity() elemental.Identity {

	return SearchIdentity
}

// Copy returns a pointer to a copy the SearchesList.
func (o SearchesList) Copy() elemental.Identifiables {

	copy := append(SearchesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SearchesList.
func (o SearchesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SearchesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Search))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SearchesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SearchesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the SearchesList converted to SparseSearchesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o SearchesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseSearchesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseSearch)
	}

	return out
}

// Version returns the version of the content.
func (o SearchesList) Version() int {

	return 1
}

// Search represents the model of a search
type Search struct {
	// Contains the matched object.
	Object interface{} `json:"object" msgpack:"object" bson:"-" mapstructure:"object,omitempty"`

	// Contains the ID of the match.
	ObjectID string `json:"objectID" msgpack:"objectID" bson:"-" mapstructure:"objectID,omitempty"`

	// Contains the identity of the match.
	ObjectIdentity string `json:"objectIdentity" msgpack:"objectIdentity" bson:"-" mapstructure:"objectIdentity,omitempty"`

	// Contains the namespace of the match.
	ObjectNamespace string `json:"objectNamespace" msgpack:"objectNamespace" bson:"-" mapstructure:"objectNamespace,omitempty"`

	// Contains the score of the match.
	Score float64 `json:"score" msgpack:"score" bson:"-" mapstructure:"score,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSearch returns a new *Search
func NewSearch() *Search {

	return &Search{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Search) Identity() elemental.Identity {

	return SearchIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Search) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Search) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Search) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSearch{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Search) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSearch{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Search) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Search) BleveType() string {

	return "search"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Search) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Search) Doc() string {

	return `Perform a full text search on the database.`
}

func (o *Search) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Search) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseSearch{
			Object:          &o.Object,
			ObjectID:        &o.ObjectID,
			ObjectIdentity:  &o.ObjectIdentity,
			ObjectNamespace: &o.ObjectNamespace,
			Score:           &o.Score,
		}
	}

	sp := &SparseSearch{}
	for _, f := range fields {
		switch f {
		case "object":
			sp.Object = &(o.Object)
		case "objectID":
			sp.ObjectID = &(o.ObjectID)
		case "objectIdentity":
			sp.ObjectIdentity = &(o.ObjectIdentity)
		case "objectNamespace":
			sp.ObjectNamespace = &(o.ObjectNamespace)
		case "score":
			sp.Score = &(o.Score)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseSearch to the object.
func (o *Search) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseSearch)
	if so.Object != nil {
		o.Object = *so.Object
	}
	if so.ObjectID != nil {
		o.ObjectID = *so.ObjectID
	}
	if so.ObjectIdentity != nil {
		o.ObjectIdentity = *so.ObjectIdentity
	}
	if so.ObjectNamespace != nil {
		o.ObjectNamespace = *so.ObjectNamespace
	}
	if so.Score != nil {
		o.Score = *so.Score
	}
}

// DeepCopy returns a deep copy if the Search.
func (o *Search) DeepCopy() *Search {

	if o == nil {
		return nil
	}

	out := &Search{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Search.
func (o *Search) DeepCopyInto(out *Search) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Search: %s", err))
	}

	*out = *target.(*Search)
}

// Validate valides the current information stored into the structure.
func (o *Search) Validate() error {

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
func (*Search) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := SearchAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return SearchLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Search) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return SearchAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Search) ValueForAttribute(name string) interface{} {

	switch name {
	case "object":
		return o.Object
	case "objectID":
		return o.ObjectID
	case "objectIdentity":
		return o.ObjectIdentity
	case "objectNamespace":
		return o.ObjectNamespace
	case "score":
		return o.Score
	}

	return nil
}

// SearchAttributesMap represents the map of attribute for Search.
var SearchAttributesMap = map[string]elemental.AttributeSpecification{
	"Object": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Object",
		Description:    `Contains the matched object.`,
		Exposed:        true,
		Name:           "object",
		ReadOnly:       true,
		Type:           "object",
	},
	"ObjectID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ObjectID",
		Description:    `Contains the ID of the match.`,
		Exposed:        true,
		Name:           "objectID",
		ReadOnly:       true,
		Type:           "string",
	},
	"ObjectIdentity": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ObjectIdentity",
		Description:    `Contains the identity of the match.`,
		Exposed:        true,
		Name:           "objectIdentity",
		ReadOnly:       true,
		Type:           "string",
	},
	"ObjectNamespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ObjectNamespace",
		Description:    `Contains the namespace of the match.`,
		Exposed:        true,
		Name:           "objectNamespace",
		ReadOnly:       true,
		Type:           "string",
	},
	"Score": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Score",
		Description:    `Contains the score of the match.`,
		Exposed:        true,
		Name:           "score",
		ReadOnly:       true,
		Type:           "float",
	},
}

// SearchLowerCaseAttributesMap represents the map of attribute for Search.
var SearchLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"object": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Object",
		Description:    `Contains the matched object.`,
		Exposed:        true,
		Name:           "object",
		ReadOnly:       true,
		Type:           "object",
	},
	"objectid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ObjectID",
		Description:    `Contains the ID of the match.`,
		Exposed:        true,
		Name:           "objectID",
		ReadOnly:       true,
		Type:           "string",
	},
	"objectidentity": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ObjectIdentity",
		Description:    `Contains the identity of the match.`,
		Exposed:        true,
		Name:           "objectIdentity",
		ReadOnly:       true,
		Type:           "string",
	},
	"objectnamespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ObjectNamespace",
		Description:    `Contains the namespace of the match.`,
		Exposed:        true,
		Name:           "objectNamespace",
		ReadOnly:       true,
		Type:           "string",
	},
	"score": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Score",
		Description:    `Contains the score of the match.`,
		Exposed:        true,
		Name:           "score",
		ReadOnly:       true,
		Type:           "float",
	},
}

// SparseSearchesList represents a list of SparseSearches
type SparseSearchesList []*SparseSearch

// Identity returns the identity of the objects in the list.
func (o SparseSearchesList) Identity() elemental.Identity {

	return SearchIdentity
}

// Copy returns a pointer to a copy the SparseSearchesList.
func (o SparseSearchesList) Copy() elemental.Identifiables {

	copy := append(SparseSearchesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseSearchesList.
func (o SparseSearchesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseSearchesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseSearch))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseSearchesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseSearchesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseSearchesList converted to SearchesList.
func (o SparseSearchesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseSearchesList) Version() int {

	return 1
}

// SparseSearch represents the sparse version of a search.
type SparseSearch struct {
	// Contains the matched object.
	Object *interface{} `json:"object,omitempty" msgpack:"object,omitempty" bson:"-" mapstructure:"object,omitempty"`

	// Contains the ID of the match.
	ObjectID *string `json:"objectID,omitempty" msgpack:"objectID,omitempty" bson:"-" mapstructure:"objectID,omitempty"`

	// Contains the identity of the match.
	ObjectIdentity *string `json:"objectIdentity,omitempty" msgpack:"objectIdentity,omitempty" bson:"-" mapstructure:"objectIdentity,omitempty"`

	// Contains the namespace of the match.
	ObjectNamespace *string `json:"objectNamespace,omitempty" msgpack:"objectNamespace,omitempty" bson:"-" mapstructure:"objectNamespace,omitempty"`

	// Contains the score of the match.
	Score *float64 `json:"score,omitempty" msgpack:"score,omitempty" bson:"-" mapstructure:"score,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseSearch returns a new  SparseSearch.
func NewSparseSearch() *SparseSearch {
	return &SparseSearch{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseSearch) Identity() elemental.Identity {

	return SearchIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseSearch) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseSearch) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseSearch) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseSearch{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseSearch) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseSearch{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseSearch) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseSearch) ToPlain() elemental.PlainIdentifiable {

	out := NewSearch()
	if o.Object != nil {
		out.Object = *o.Object
	}
	if o.ObjectID != nil {
		out.ObjectID = *o.ObjectID
	}
	if o.ObjectIdentity != nil {
		out.ObjectIdentity = *o.ObjectIdentity
	}
	if o.ObjectNamespace != nil {
		out.ObjectNamespace = *o.ObjectNamespace
	}
	if o.Score != nil {
		out.Score = *o.Score
	}

	return out
}

// DeepCopy returns a deep copy if the SparseSearch.
func (o *SparseSearch) DeepCopy() *SparseSearch {

	if o == nil {
		return nil
	}

	out := &SparseSearch{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseSearch.
func (o *SparseSearch) DeepCopyInto(out *SparseSearch) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseSearch: %s", err))
	}

	*out = *target.(*SparseSearch)
}

type mongoAttributesSearch struct {
}
type mongoAttributesSparseSearch struct {
}
