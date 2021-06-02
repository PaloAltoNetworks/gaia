package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TagPrefixIdentity represents the Identity of the object.
var TagPrefixIdentity = elemental.Identity{
	Name:     "tagprefix",
	Category: "tagprefixes",
	Package:  "squall",
	Private:  false,
}

// TagPrefixsList represents a list of TagPrefixs
type TagPrefixsList []*TagPrefix

// Identity returns the identity of the objects in the list.
func (o TagPrefixsList) Identity() elemental.Identity {

	return TagPrefixIdentity
}

// Copy returns a pointer to a copy the TagPrefixsList.
func (o TagPrefixsList) Copy() elemental.Identifiables {

	copy := append(TagPrefixsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TagPrefixsList.
func (o TagPrefixsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(TagPrefixsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*TagPrefix))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TagPrefixsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TagPrefixsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the TagPrefixsList converted to SparseTagPrefixsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o TagPrefixsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseTagPrefixsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseTagPrefix)
	}

	return out
}

// Version returns the version of the content.
func (o TagPrefixsList) Version() int {

	return 1
}

// TagPrefix represents the model of a tagprefix
type TagPrefix struct {
	// The namespace ID of the current namespace.
	NamespaceID string `json:"namespaceID" msgpack:"namespaceID" bson:"-" mapstructure:"namespaceID,omitempty"`

	// List of tag prefixes that will be used to suggest policies. Only these tags will
	// be transmitted on the wire.
	Prefixes []string `json:"prefixes" msgpack:"prefixes" bson:"prefixes" mapstructure:"prefixes,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewTagPrefix returns a new *TagPrefix
func NewTagPrefix() *TagPrefix {

	return &TagPrefix{
		ModelVersion: 1,
		Prefixes:     []string{},
	}
}

// Identity returns the Identity of the object.
func (o *TagPrefix) Identity() elemental.Identity {

	return TagPrefixIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *TagPrefix) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *TagPrefix) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TagPrefix) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesTagPrefix{}

	s.Prefixes = o.Prefixes

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TagPrefix) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesTagPrefix{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Prefixes = s.Prefixes

	return nil
}

// Version returns the hardcoded version of the model.
func (o *TagPrefix) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *TagPrefix) BleveType() string {

	return "tagprefix"
}

// DefaultOrder returns the list of default ordering fields.
func (o *TagPrefix) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *TagPrefix) Doc() string {

	return `Returns the tag prefixes of the specified namespace.`
}

func (o *TagPrefix) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *TagPrefix) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseTagPrefix{
			NamespaceID: &o.NamespaceID,
			Prefixes:    &o.Prefixes,
		}
	}

	sp := &SparseTagPrefix{}
	for _, f := range fields {
		switch f {
		case "namespaceID":
			sp.NamespaceID = &(o.NamespaceID)
		case "prefixes":
			sp.Prefixes = &(o.Prefixes)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseTagPrefix to the object.
func (o *TagPrefix) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseTagPrefix)
	if so.NamespaceID != nil {
		o.NamespaceID = *so.NamespaceID
	}
	if so.Prefixes != nil {
		o.Prefixes = *so.Prefixes
	}
}

// DeepCopy returns a deep copy if the TagPrefix.
func (o *TagPrefix) DeepCopy() *TagPrefix {

	if o == nil {
		return nil
	}

	out := &TagPrefix{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *TagPrefix.
func (o *TagPrefix) DeepCopyInto(out *TagPrefix) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy TagPrefix: %s", err))
	}

	*out = *target.(*TagPrefix)
}

// Validate valides the current information stored into the structure.
func (o *TagPrefix) Validate() error {

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
func (*TagPrefix) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TagPrefixAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TagPrefixLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*TagPrefix) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TagPrefixAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *TagPrefix) ValueForAttribute(name string) interface{} {

	switch name {
	case "namespaceID":
		return o.NamespaceID
	case "prefixes":
		return o.Prefixes
	}

	return nil
}

// TagPrefixAttributesMap represents the map of attribute for TagPrefix.
var TagPrefixAttributesMap = map[string]elemental.AttributeSpecification{
	"NamespaceID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NamespaceID",
		Description:    `The namespace ID of the current namespace.`,
		Exposed:        true,
		Name:           "namespaceID",
		ReadOnly:       true,
		Type:           "string",
	},
	"Prefixes": {
		AllowedChoices: []string{},
		BSONFieldName:  "prefixes",
		ConvertedName:  "Prefixes",
		Description: `List of tag prefixes that will be used to suggest policies. Only these tags will
be transmitted on the wire.`,
		Exposed: true,
		Name:    "prefixes",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
}

// TagPrefixLowerCaseAttributesMap represents the map of attribute for TagPrefix.
var TagPrefixLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"namespaceid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NamespaceID",
		Description:    `The namespace ID of the current namespace.`,
		Exposed:        true,
		Name:           "namespaceID",
		ReadOnly:       true,
		Type:           "string",
	},
	"prefixes": {
		AllowedChoices: []string{},
		BSONFieldName:  "prefixes",
		ConvertedName:  "Prefixes",
		Description: `List of tag prefixes that will be used to suggest policies. Only these tags will
be transmitted on the wire.`,
		Exposed: true,
		Name:    "prefixes",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
}

// SparseTagPrefixsList represents a list of SparseTagPrefixs
type SparseTagPrefixsList []*SparseTagPrefix

// Identity returns the identity of the objects in the list.
func (o SparseTagPrefixsList) Identity() elemental.Identity {

	return TagPrefixIdentity
}

// Copy returns a pointer to a copy the SparseTagPrefixsList.
func (o SparseTagPrefixsList) Copy() elemental.Identifiables {

	copy := append(SparseTagPrefixsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseTagPrefixsList.
func (o SparseTagPrefixsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseTagPrefixsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseTagPrefix))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseTagPrefixsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseTagPrefixsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseTagPrefixsList converted to TagPrefixsList.
func (o SparseTagPrefixsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseTagPrefixsList) Version() int {

	return 1
}

// SparseTagPrefix represents the sparse version of a tagprefix.
type SparseTagPrefix struct {
	// The namespace ID of the current namespace.
	NamespaceID *string `json:"namespaceID,omitempty" msgpack:"namespaceID,omitempty" bson:"-" mapstructure:"namespaceID,omitempty"`

	// List of tag prefixes that will be used to suggest policies. Only these tags will
	// be transmitted on the wire.
	Prefixes *[]string `json:"prefixes,omitempty" msgpack:"prefixes,omitempty" bson:"prefixes,omitempty" mapstructure:"prefixes,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseTagPrefix returns a new  SparseTagPrefix.
func NewSparseTagPrefix() *SparseTagPrefix {
	return &SparseTagPrefix{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseTagPrefix) Identity() elemental.Identity {

	return TagPrefixIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseTagPrefix) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseTagPrefix) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTagPrefix) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseTagPrefix{}

	if o.Prefixes != nil {
		s.Prefixes = o.Prefixes
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTagPrefix) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseTagPrefix{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Prefixes != nil {
		o.Prefixes = s.Prefixes
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseTagPrefix) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseTagPrefix) ToPlain() elemental.PlainIdentifiable {

	out := NewTagPrefix()
	if o.NamespaceID != nil {
		out.NamespaceID = *o.NamespaceID
	}
	if o.Prefixes != nil {
		out.Prefixes = *o.Prefixes
	}

	return out
}

// DeepCopy returns a deep copy if the SparseTagPrefix.
func (o *SparseTagPrefix) DeepCopy() *SparseTagPrefix {

	if o == nil {
		return nil
	}

	out := &SparseTagPrefix{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseTagPrefix.
func (o *SparseTagPrefix) DeepCopyInto(out *SparseTagPrefix) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseTagPrefix: %s", err))
	}

	*out = *target.(*SparseTagPrefix)
}

type mongoAttributesTagPrefix struct {
	Prefixes []string `bson:"prefixes"`
}
type mongoAttributesSparseTagPrefix struct {
	Prefixes *[]string `bson:"prefixes,omitempty"`
}
