package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CNSSuggestionIdentity represents the Identity of the object.
var CNSSuggestionIdentity = elemental.Identity{
	Name:     "cnssuggestion",
	Category: "cnssuggestions",
	Package:  "karl",
	Private:  false,
}

// CNSSuggestionsList represents a list of CNSSuggestions
type CNSSuggestionsList []*CNSSuggestion

// Identity returns the identity of the objects in the list.
func (o CNSSuggestionsList) Identity() elemental.Identity {

	return CNSSuggestionIdentity
}

// Copy returns a pointer to a copy the CNSSuggestionsList.
func (o CNSSuggestionsList) Copy() elemental.Identifiables {

	copy := append(CNSSuggestionsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CNSSuggestionsList.
func (o CNSSuggestionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CNSSuggestionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CNSSuggestion))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CNSSuggestionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CNSSuggestionsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CNSSuggestionsList converted to SparseCNSSuggestionsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CNSSuggestionsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCNSSuggestionsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCNSSuggestion)
	}

	return out
}

// Version returns the version of the content.
func (o CNSSuggestionsList) Version() int {

	return 1
}

// CNSSuggestion represents the model of a cnssuggestion
type CNSSuggestion struct {
	// Required by Prisma Cloud. Always set to true.
	NeedsOffsetUpdate bool `json:"needsOffsetUpdate" msgpack:"needsOffsetUpdate" bson:"-" mapstructure:"needsOffsetUpdate,omitempty"`

	// The length of the RQL query part that is valid.
	Offset int `json:"offset" msgpack:"offset" bson:"-" mapstructure:"offset,omitempty"`

	// Prisma Cloud's RQL query.
	Query string `json:"query,omitempty" msgpack:"query,omitempty" bson:"-" mapstructure:"query,omitempty"`

	// List of query suggestions.
	Suggestions []string `json:"suggestions" msgpack:"suggestions" bson:"-" mapstructure:"suggestions,omitempty"`

	// Required by Prisma Cloud. Always set to false.
	Translate bool `json:"translate" msgpack:"translate" bson:"-" mapstructure:"translate,omitempty"`

	// The validity of the RQL query.
	Valid bool `json:"valid" msgpack:"valid" bson:"-" mapstructure:"valid,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCNSSuggestion returns a new *CNSSuggestion
func NewCNSSuggestion() *CNSSuggestion {

	return &CNSSuggestion{
		ModelVersion:      1,
		NeedsOffsetUpdate: true,
		Offset:            0,
		Suggestions:       []string{},
		Translate:         false,
		Valid:             false,
	}
}

// Identity returns the Identity of the object.
func (o *CNSSuggestion) Identity() elemental.Identity {

	return CNSSuggestionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CNSSuggestion) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CNSSuggestion) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CNSSuggestion) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCNSSuggestion{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CNSSuggestion) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCNSSuggestion{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CNSSuggestion) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CNSSuggestion) BleveType() string {

	return "cnssuggestion"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CNSSuggestion) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CNSSuggestion) Doc() string {

	return `Provides query suggestions for Prisma Cloud's investigate page.`
}

func (o *CNSSuggestion) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CNSSuggestion) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCNSSuggestion{
			NeedsOffsetUpdate: &o.NeedsOffsetUpdate,
			Offset:            &o.Offset,
			Query:             &o.Query,
			Suggestions:       &o.Suggestions,
			Translate:         &o.Translate,
			Valid:             &o.Valid,
		}
	}

	sp := &SparseCNSSuggestion{}
	for _, f := range fields {
		switch f {
		case "needsOffsetUpdate":
			sp.NeedsOffsetUpdate = &(o.NeedsOffsetUpdate)
		case "offset":
			sp.Offset = &(o.Offset)
		case "query":
			sp.Query = &(o.Query)
		case "suggestions":
			sp.Suggestions = &(o.Suggestions)
		case "translate":
			sp.Translate = &(o.Translate)
		case "valid":
			sp.Valid = &(o.Valid)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCNSSuggestion to the object.
func (o *CNSSuggestion) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCNSSuggestion)
	if so.NeedsOffsetUpdate != nil {
		o.NeedsOffsetUpdate = *so.NeedsOffsetUpdate
	}
	if so.Offset != nil {
		o.Offset = *so.Offset
	}
	if so.Query != nil {
		o.Query = *so.Query
	}
	if so.Suggestions != nil {
		o.Suggestions = *so.Suggestions
	}
	if so.Translate != nil {
		o.Translate = *so.Translate
	}
	if so.Valid != nil {
		o.Valid = *so.Valid
	}
}

// DeepCopy returns a deep copy if the CNSSuggestion.
func (o *CNSSuggestion) DeepCopy() *CNSSuggestion {

	if o == nil {
		return nil
	}

	out := &CNSSuggestion{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CNSSuggestion.
func (o *CNSSuggestion) DeepCopyInto(out *CNSSuggestion) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CNSSuggestion: %s", err))
	}

	*out = *target.(*CNSSuggestion)
}

// Validate valides the current information stored into the structure.
func (o *CNSSuggestion) Validate() error {

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
func (*CNSSuggestion) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CNSSuggestionAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CNSSuggestionLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CNSSuggestion) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CNSSuggestionAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CNSSuggestion) ValueForAttribute(name string) interface{} {

	switch name {
	case "needsOffsetUpdate":
		return o.NeedsOffsetUpdate
	case "offset":
		return o.Offset
	case "query":
		return o.Query
	case "suggestions":
		return o.Suggestions
	case "translate":
		return o.Translate
	case "valid":
		return o.Valid
	}

	return nil
}

// CNSSuggestionAttributesMap represents the map of attribute for CNSSuggestion.
var CNSSuggestionAttributesMap = map[string]elemental.AttributeSpecification{
	"NeedsOffsetUpdate": {
		AllowedChoices: []string{},
		ConvertedName:  "NeedsOffsetUpdate",
		DefaultValue:   true,
		Description:    `Required by Prisma Cloud. Always set to true.`,
		Exposed:        true,
		Name:           "needsOffsetUpdate",
		Type:           "boolean",
	},
	"Offset": {
		AllowedChoices: []string{},
		ConvertedName:  "Offset",
		Description:    `The length of the RQL query part that is valid.`,
		Exposed:        true,
		Name:           "offset",
		Type:           "integer",
	},
	"Query": {
		AllowedChoices: []string{},
		ConvertedName:  "Query",
		Description:    `Prisma Cloud's RQL query.`,
		Exposed:        true,
		Name:           "query",
		ReadOnly:       true,
		Type:           "string",
	},
	"Suggestions": {
		AllowedChoices: []string{},
		ConvertedName:  "Suggestions",
		Description:    `List of query suggestions.`,
		Exposed:        true,
		Name:           "suggestions",
		SubType:        "string",
		Type:           "list",
	},
	"Translate": {
		AllowedChoices: []string{},
		ConvertedName:  "Translate",
		Description:    `Required by Prisma Cloud. Always set to false.`,
		Exposed:        true,
		Name:           "translate",
		Type:           "boolean",
	},
	"Valid": {
		AllowedChoices: []string{},
		ConvertedName:  "Valid",
		Description:    `The validity of the RQL query.`,
		Exposed:        true,
		Name:           "valid",
		Type:           "boolean",
	},
}

// CNSSuggestionLowerCaseAttributesMap represents the map of attribute for CNSSuggestion.
var CNSSuggestionLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"needsoffsetupdate": {
		AllowedChoices: []string{},
		ConvertedName:  "NeedsOffsetUpdate",
		DefaultValue:   true,
		Description:    `Required by Prisma Cloud. Always set to true.`,
		Exposed:        true,
		Name:           "needsOffsetUpdate",
		Type:           "boolean",
	},
	"offset": {
		AllowedChoices: []string{},
		ConvertedName:  "Offset",
		Description:    `The length of the RQL query part that is valid.`,
		Exposed:        true,
		Name:           "offset",
		Type:           "integer",
	},
	"query": {
		AllowedChoices: []string{},
		ConvertedName:  "Query",
		Description:    `Prisma Cloud's RQL query.`,
		Exposed:        true,
		Name:           "query",
		ReadOnly:       true,
		Type:           "string",
	},
	"suggestions": {
		AllowedChoices: []string{},
		ConvertedName:  "Suggestions",
		Description:    `List of query suggestions.`,
		Exposed:        true,
		Name:           "suggestions",
		SubType:        "string",
		Type:           "list",
	},
	"translate": {
		AllowedChoices: []string{},
		ConvertedName:  "Translate",
		Description:    `Required by Prisma Cloud. Always set to false.`,
		Exposed:        true,
		Name:           "translate",
		Type:           "boolean",
	},
	"valid": {
		AllowedChoices: []string{},
		ConvertedName:  "Valid",
		Description:    `The validity of the RQL query.`,
		Exposed:        true,
		Name:           "valid",
		Type:           "boolean",
	},
}

// SparseCNSSuggestionsList represents a list of SparseCNSSuggestions
type SparseCNSSuggestionsList []*SparseCNSSuggestion

// Identity returns the identity of the objects in the list.
func (o SparseCNSSuggestionsList) Identity() elemental.Identity {

	return CNSSuggestionIdentity
}

// Copy returns a pointer to a copy the SparseCNSSuggestionsList.
func (o SparseCNSSuggestionsList) Copy() elemental.Identifiables {

	copy := append(SparseCNSSuggestionsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCNSSuggestionsList.
func (o SparseCNSSuggestionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCNSSuggestionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCNSSuggestion))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCNSSuggestionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCNSSuggestionsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCNSSuggestionsList converted to CNSSuggestionsList.
func (o SparseCNSSuggestionsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCNSSuggestionsList) Version() int {

	return 1
}

// SparseCNSSuggestion represents the sparse version of a cnssuggestion.
type SparseCNSSuggestion struct {
	// Required by Prisma Cloud. Always set to true.
	NeedsOffsetUpdate *bool `json:"needsOffsetUpdate,omitempty" msgpack:"needsOffsetUpdate,omitempty" bson:"-" mapstructure:"needsOffsetUpdate,omitempty"`

	// The length of the RQL query part that is valid.
	Offset *int `json:"offset,omitempty" msgpack:"offset,omitempty" bson:"-" mapstructure:"offset,omitempty"`

	// Prisma Cloud's RQL query.
	Query *string `json:"query,omitempty" msgpack:"query,omitempty" bson:"-" mapstructure:"query,omitempty"`

	// List of query suggestions.
	Suggestions *[]string `json:"suggestions,omitempty" msgpack:"suggestions,omitempty" bson:"-" mapstructure:"suggestions,omitempty"`

	// Required by Prisma Cloud. Always set to false.
	Translate *bool `json:"translate,omitempty" msgpack:"translate,omitempty" bson:"-" mapstructure:"translate,omitempty"`

	// The validity of the RQL query.
	Valid *bool `json:"valid,omitempty" msgpack:"valid,omitempty" bson:"-" mapstructure:"valid,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCNSSuggestion returns a new  SparseCNSSuggestion.
func NewSparseCNSSuggestion() *SparseCNSSuggestion {
	return &SparseCNSSuggestion{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCNSSuggestion) Identity() elemental.Identity {

	return CNSSuggestionIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCNSSuggestion) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCNSSuggestion) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCNSSuggestion) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCNSSuggestion{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCNSSuggestion) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCNSSuggestion{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCNSSuggestion) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCNSSuggestion) ToPlain() elemental.PlainIdentifiable {

	out := NewCNSSuggestion()
	if o.NeedsOffsetUpdate != nil {
		out.NeedsOffsetUpdate = *o.NeedsOffsetUpdate
	}
	if o.Offset != nil {
		out.Offset = *o.Offset
	}
	if o.Query != nil {
		out.Query = *o.Query
	}
	if o.Suggestions != nil {
		out.Suggestions = *o.Suggestions
	}
	if o.Translate != nil {
		out.Translate = *o.Translate
	}
	if o.Valid != nil {
		out.Valid = *o.Valid
	}

	return out
}

// DeepCopy returns a deep copy if the SparseCNSSuggestion.
func (o *SparseCNSSuggestion) DeepCopy() *SparseCNSSuggestion {

	if o == nil {
		return nil
	}

	out := &SparseCNSSuggestion{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCNSSuggestion.
func (o *SparseCNSSuggestion) DeepCopyInto(out *SparseCNSSuggestion) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCNSSuggestion: %s", err))
	}

	*out = *target.(*SparseCNSSuggestion)
}

type mongoAttributesCNSSuggestion struct {
}
type mongoAttributesSparseCNSSuggestion struct {
}
