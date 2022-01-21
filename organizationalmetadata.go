package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// OrganizationalMetadataIdentity represents the Identity of the object.
var OrganizationalMetadataIdentity = elemental.Identity{
	Name:     "organizationalmetadata",
	Category: "organizationalmetadata",
	Package:  "squall",
	Private:  false,
}

// OrganizationalMetadatasList represents a list of OrganizationalMetadatas
type OrganizationalMetadatasList []*OrganizationalMetadata

// Identity returns the identity of the objects in the list.
func (o OrganizationalMetadatasList) Identity() elemental.Identity {

	return OrganizationalMetadataIdentity
}

// Copy returns a pointer to a copy the OrganizationalMetadatasList.
func (o OrganizationalMetadatasList) Copy() elemental.Identifiables {

	copy := append(OrganizationalMetadatasList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the OrganizationalMetadatasList.
func (o OrganizationalMetadatasList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(OrganizationalMetadatasList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*OrganizationalMetadata))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o OrganizationalMetadatasList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o OrganizationalMetadatasList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the OrganizationalMetadatasList converted to SparseOrganizationalMetadatasList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o OrganizationalMetadatasList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseOrganizationalMetadatasList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseOrganizationalMetadata)
	}

	return out
}

// Version returns the version of the content.
func (o OrganizationalMetadatasList) Version() int {

	return 1
}

// OrganizationalMetadata represents the model of a organizationalmetadata
type OrganizationalMetadata struct {
	// List of organizational metadata for the namespace.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewOrganizationalMetadata returns a new *OrganizationalMetadata
func NewOrganizationalMetadata() *OrganizationalMetadata {

	return &OrganizationalMetadata{
		ModelVersion: 1,
		Metadata:     []string{},
	}
}

// Identity returns the Identity of the object.
func (o *OrganizationalMetadata) Identity() elemental.Identity {

	return OrganizationalMetadataIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *OrganizationalMetadata) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *OrganizationalMetadata) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *OrganizationalMetadata) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesOrganizationalMetadata{}

	s.Metadata = o.Metadata
	s.Namespace = o.Namespace

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *OrganizationalMetadata) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesOrganizationalMetadata{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Metadata = s.Metadata
	o.Namespace = s.Namespace

	return nil
}

// Version returns the hardcoded version of the model.
func (o *OrganizationalMetadata) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *OrganizationalMetadata) BleveType() string {

	return "organizationalmetadata"
}

// DefaultOrder returns the list of default ordering fields.
func (o *OrganizationalMetadata) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *OrganizationalMetadata) Doc() string {

	return `Can be used to retrieve the organizational metadata of the namespace.`
}

func (o *OrganizationalMetadata) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetNamespace returns the Namespace of the receiver.
func (o *OrganizationalMetadata) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *OrganizationalMetadata) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *OrganizationalMetadata) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseOrganizationalMetadata{
			Metadata:  &o.Metadata,
			Namespace: &o.Namespace,
		}
	}

	sp := &SparseOrganizationalMetadata{}
	for _, f := range fields {
		switch f {
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseOrganizationalMetadata to the object.
func (o *OrganizationalMetadata) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseOrganizationalMetadata)
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
}

// DeepCopy returns a deep copy if the OrganizationalMetadata.
func (o *OrganizationalMetadata) DeepCopy() *OrganizationalMetadata {

	if o == nil {
		return nil
	}

	out := &OrganizationalMetadata{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *OrganizationalMetadata.
func (o *OrganizationalMetadata) DeepCopyInto(out *OrganizationalMetadata) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy OrganizationalMetadata: %s", err))
	}

	*out = *target.(*OrganizationalMetadata)
}

// Validate valides the current information stored into the structure.
func (o *OrganizationalMetadata) Validate() error {

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
func (*OrganizationalMetadata) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := OrganizationalMetadataAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return OrganizationalMetadataLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*OrganizationalMetadata) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return OrganizationalMetadataAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *OrganizationalMetadata) ValueForAttribute(name string) interface{} {

	switch name {
	case "metadata":
		return o.Metadata
	case "namespace":
		return o.Namespace
	}

	return nil
}

// OrganizationalMetadataAttributesMap represents the map of attribute for OrganizationalMetadata.
var OrganizationalMetadataAttributesMap = map[string]elemental.AttributeSpecification{
	"Metadata": {
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		Description:    `List of organizational metadata for the namespace.`,
		Exposed:        true,
		Name:           "metadata",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
}

// OrganizationalMetadataLowerCaseAttributesMap represents the map of attribute for OrganizationalMetadata.
var OrganizationalMetadataLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"metadata": {
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		Description:    `List of organizational metadata for the namespace.`,
		Exposed:        true,
		Name:           "metadata",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
}

// SparseOrganizationalMetadatasList represents a list of SparseOrganizationalMetadatas
type SparseOrganizationalMetadatasList []*SparseOrganizationalMetadata

// Identity returns the identity of the objects in the list.
func (o SparseOrganizationalMetadatasList) Identity() elemental.Identity {

	return OrganizationalMetadataIdentity
}

// Copy returns a pointer to a copy the SparseOrganizationalMetadatasList.
func (o SparseOrganizationalMetadatasList) Copy() elemental.Identifiables {

	copy := append(SparseOrganizationalMetadatasList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseOrganizationalMetadatasList.
func (o SparseOrganizationalMetadatasList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseOrganizationalMetadatasList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseOrganizationalMetadata))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseOrganizationalMetadatasList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseOrganizationalMetadatasList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseOrganizationalMetadatasList converted to OrganizationalMetadatasList.
func (o SparseOrganizationalMetadatasList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseOrganizationalMetadatasList) Version() int {

	return 1
}

// SparseOrganizationalMetadata represents the sparse version of a organizationalmetadata.
type SparseOrganizationalMetadata struct {
	// List of organizational metadata for the namespace.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseOrganizationalMetadata returns a new  SparseOrganizationalMetadata.
func NewSparseOrganizationalMetadata() *SparseOrganizationalMetadata {
	return &SparseOrganizationalMetadata{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseOrganizationalMetadata) Identity() elemental.Identity {

	return OrganizationalMetadataIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseOrganizationalMetadata) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseOrganizationalMetadata) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseOrganizationalMetadata) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseOrganizationalMetadata{}

	if o.Metadata != nil {
		s.Metadata = o.Metadata
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseOrganizationalMetadata) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseOrganizationalMetadata{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Metadata != nil {
		o.Metadata = s.Metadata
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseOrganizationalMetadata) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseOrganizationalMetadata) ToPlain() elemental.PlainIdentifiable {

	out := NewOrganizationalMetadata()
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}

	return out
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseOrganizationalMetadata) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseOrganizationalMetadata) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// DeepCopy returns a deep copy if the SparseOrganizationalMetadata.
func (o *SparseOrganizationalMetadata) DeepCopy() *SparseOrganizationalMetadata {

	if o == nil {
		return nil
	}

	out := &SparseOrganizationalMetadata{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseOrganizationalMetadata.
func (o *SparseOrganizationalMetadata) DeepCopyInto(out *SparseOrganizationalMetadata) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseOrganizationalMetadata: %s", err))
	}

	*out = *target.(*SparseOrganizationalMetadata)
}

type mongoAttributesOrganizationalMetadata struct {
	Metadata  []string `bson:"metadata"`
	Namespace string   `bson:"namespace"`
}
type mongoAttributesSparseOrganizationalMetadata struct {
	Metadata  *[]string `bson:"metadata,omitempty"`
	Namespace *string   `bson:"namespace,omitempty"`
}
