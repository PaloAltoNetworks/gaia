package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// SquallTagIdentity represents the Identity of the object.
var SquallTagIdentity = elemental.Identity{
	Name:     "squalltag",
	Category: "squalltags",
	Package:  "squall",
	Private:  true,
}

// SquallTagsList represents a list of SquallTags
type SquallTagsList []*SquallTag

// Identity returns the identity of the objects in the list.
func (o SquallTagsList) Identity() elemental.Identity {

	return SquallTagIdentity
}

// Copy returns a pointer to a copy the SquallTagsList.
func (o SquallTagsList) Copy() elemental.Identifiables {

	copy := append(SquallTagsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SquallTagsList.
func (o SquallTagsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SquallTagsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SquallTag))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SquallTagsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SquallTagsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the SquallTagsList converted to SparseSquallTagsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o SquallTagsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseSquallTagsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseSquallTag)
	}

	return out
}

// Version returns the version of the content.
func (o SquallTagsList) Version() int {

	return 1
}

// SquallTag represents the model of a squalltag
type SquallTag struct {
	// Number of time this tag is used.
	Count int `json:"count" msgpack:"count" bson:"count" mapstructure:"count,omitempty"`

	// namespace containing these tags.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Value of the tag.
	Value string `json:"value" msgpack:"value" bson:"value" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSquallTag returns a new *SquallTag
func NewSquallTag() *SquallTag {

	return &SquallTag{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *SquallTag) Identity() elemental.Identity {

	return SquallTagIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SquallTag) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SquallTag) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SquallTag) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSquallTag{}

	s.Count = o.Count
	s.Namespace = o.Namespace
	s.Value = o.Value

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SquallTag) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSquallTag{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Count = s.Count
	o.Namespace = s.Namespace
	o.Value = s.Value

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SquallTag) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *SquallTag) BleveType() string {

	return "squalltag"
}

// DefaultOrder returns the list of default ordering fields.
func (o *SquallTag) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *SquallTag) Doc() string {

	return `Internal api that retrieve all tags from squall and their count for given
` + "`" + `?identity=<identity>` + "`" + ` parameter with their counts.`
}

func (o *SquallTag) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *SquallTag) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseSquallTag{
			Count:     &o.Count,
			Namespace: &o.Namespace,
			Value:     &o.Value,
		}
	}

	sp := &SparseSquallTag{}
	for _, f := range fields {
		switch f {
		case "count":
			sp.Count = &(o.Count)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "value":
			sp.Value = &(o.Value)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseSquallTag to the object.
func (o *SquallTag) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseSquallTag)
	if so.Count != nil {
		o.Count = *so.Count
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.Value != nil {
		o.Value = *so.Value
	}
}

// DeepCopy returns a deep copy if the SquallTag.
func (o *SquallTag) DeepCopy() *SquallTag {

	if o == nil {
		return nil
	}

	out := &SquallTag{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SquallTag.
func (o *SquallTag) DeepCopyInto(out *SquallTag) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SquallTag: %s", err))
	}

	*out = *target.(*SquallTag)
}

// Validate valides the current information stored into the structure.
func (o *SquallTag) Validate() error {

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
func (*SquallTag) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := SquallTagAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return SquallTagLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*SquallTag) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return SquallTagAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *SquallTag) ValueForAttribute(name string) interface{} {

	switch name {
	case "count":
		return o.Count
	case "namespace":
		return o.Namespace
	case "value":
		return o.Value
	}

	return nil
}

// SquallTagAttributesMap represents the map of attribute for SquallTag.
var SquallTagAttributesMap = map[string]elemental.AttributeSpecification{
	"Count": {
		AllowedChoices: []string{},
		ConvertedName:  "Count",
		Description:    `Number of time this tag is used.`,
		Exposed:        true,
		Name:           "count",
		Stored:         true,
		Type:           "integer",
	},
	"Namespace": {
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `namespace containing these tags.`,
		Exposed:        true,
		Name:           "namespace",
		Stored:         true,
		Type:           "string",
	},
	"Value": {
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Value of the tag.`,
		Exposed:        true,
		Name:           "value",
		Stored:         true,
		Type:           "string",
	},
}

// SquallTagLowerCaseAttributesMap represents the map of attribute for SquallTag.
var SquallTagLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"count": {
		AllowedChoices: []string{},
		ConvertedName:  "Count",
		Description:    `Number of time this tag is used.`,
		Exposed:        true,
		Name:           "count",
		Stored:         true,
		Type:           "integer",
	},
	"namespace": {
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `namespace containing these tags.`,
		Exposed:        true,
		Name:           "namespace",
		Stored:         true,
		Type:           "string",
	},
	"value": {
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Value of the tag.`,
		Exposed:        true,
		Name:           "value",
		Stored:         true,
		Type:           "string",
	},
}

// SparseSquallTagsList represents a list of SparseSquallTags
type SparseSquallTagsList []*SparseSquallTag

// Identity returns the identity of the objects in the list.
func (o SparseSquallTagsList) Identity() elemental.Identity {

	return SquallTagIdentity
}

// Copy returns a pointer to a copy the SparseSquallTagsList.
func (o SparseSquallTagsList) Copy() elemental.Identifiables {

	copy := append(SparseSquallTagsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseSquallTagsList.
func (o SparseSquallTagsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseSquallTagsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseSquallTag))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseSquallTagsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseSquallTagsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseSquallTagsList converted to SquallTagsList.
func (o SparseSquallTagsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseSquallTagsList) Version() int {

	return 1
}

// SparseSquallTag represents the sparse version of a squalltag.
type SparseSquallTag struct {
	// Number of time this tag is used.
	Count *int `json:"count,omitempty" msgpack:"count,omitempty" bson:"count,omitempty" mapstructure:"count,omitempty"`

	// namespace containing these tags.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Value of the tag.
	Value *string `json:"value,omitempty" msgpack:"value,omitempty" bson:"value,omitempty" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseSquallTag returns a new  SparseSquallTag.
func NewSparseSquallTag() *SparseSquallTag {
	return &SparseSquallTag{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseSquallTag) Identity() elemental.Identity {

	return SquallTagIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseSquallTag) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseSquallTag) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseSquallTag) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseSquallTag{}

	if o.Count != nil {
		s.Count = o.Count
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.Value != nil {
		s.Value = o.Value
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseSquallTag) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseSquallTag{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Count != nil {
		o.Count = s.Count
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.Value != nil {
		o.Value = s.Value
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseSquallTag) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseSquallTag) ToPlain() elemental.PlainIdentifiable {

	out := NewSquallTag()
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.Value != nil {
		out.Value = *o.Value
	}

	return out
}

// DeepCopy returns a deep copy if the SparseSquallTag.
func (o *SparseSquallTag) DeepCopy() *SparseSquallTag {

	if o == nil {
		return nil
	}

	out := &SparseSquallTag{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseSquallTag.
func (o *SparseSquallTag) DeepCopyInto(out *SparseSquallTag) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseSquallTag: %s", err))
	}

	*out = *target.(*SparseSquallTag)
}

type mongoAttributesSquallTag struct {
	Count     int    `bson:"count"`
	Namespace string `bson:"namespace"`
	Value     string `bson:"value"`
}
type mongoAttributesSparseSquallTag struct {
	Count     *int    `bson:"count,omitempty"`
	Namespace *string `bson:"namespace,omitempty"`
	Value     *string `bson:"value,omitempty"`
}
