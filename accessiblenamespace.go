package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AccessibleNamespaceIdentity represents the Identity of the object.
var AccessibleNamespaceIdentity = elemental.Identity{
	Name:     "accessiblenamespace",
	Category: "accessiblenamespaces",
	Package:  "squall",
	Private:  false,
}

// AccessibleNamespacesList represents a list of AccessibleNamespaces
type AccessibleNamespacesList []*AccessibleNamespace

// Identity returns the identity of the objects in the list.
func (o AccessibleNamespacesList) Identity() elemental.Identity {

	return AccessibleNamespaceIdentity
}

// Copy returns a pointer to a copy the AccessibleNamespacesList.
func (o AccessibleNamespacesList) Copy() elemental.Identifiables {

	copy := append(AccessibleNamespacesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AccessibleNamespacesList.
func (o AccessibleNamespacesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AccessibleNamespacesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AccessibleNamespace))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AccessibleNamespacesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AccessibleNamespacesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the AccessibleNamespacesList converted to SparseAccessibleNamespacesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AccessibleNamespacesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAccessibleNamespacesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAccessibleNamespace)
	}

	return out
}

// Version returns the version of the content.
func (o AccessibleNamespacesList) Version() int {

	return 1
}

// AccessibleNamespace represents the model of a accessiblenamespace
type AccessibleNamespace struct {
	// Identifier of the namespace that is accessible.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Name of the namespace that is accessible.
	Name string `json:"name" msgpack:"name" bson:"-" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAccessibleNamespace returns a new *AccessibleNamespace
func NewAccessibleNamespace() *AccessibleNamespace {

	return &AccessibleNamespace{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *AccessibleNamespace) Identity() elemental.Identity {

	return AccessibleNamespaceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AccessibleNamespace) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AccessibleNamespace) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AccessibleNamespace) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesAccessibleNamespace{}

	s.Namespace = o.Namespace

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AccessibleNamespace) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesAccessibleNamespace{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Namespace = s.Namespace

	return nil
}

// Version returns the hardcoded version of the model.
func (o *AccessibleNamespace) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *AccessibleNamespace) BleveType() string {

	return "accessiblenamespace"
}

// DefaultOrder returns the list of default ordering fields.
func (o *AccessibleNamespace) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *AccessibleNamespace) Doc() string {

	return `An Accessible Namespace represents a namespace that can be accessed by a given
user.`
}

func (o *AccessibleNamespace) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetNamespace returns the Namespace of the receiver.
func (o *AccessibleNamespace) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *AccessibleNamespace) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *AccessibleNamespace) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAccessibleNamespace{
			ID:        &o.ID,
			Name:      &o.Name,
			Namespace: &o.Namespace,
		}
	}

	sp := &SparseAccessibleNamespace{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseAccessibleNamespace to the object.
func (o *AccessibleNamespace) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAccessibleNamespace)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
}

// DeepCopy returns a deep copy if the AccessibleNamespace.
func (o *AccessibleNamespace) DeepCopy() *AccessibleNamespace {

	if o == nil {
		return nil
	}

	out := &AccessibleNamespace{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AccessibleNamespace.
func (o *AccessibleNamespace) DeepCopyInto(out *AccessibleNamespace) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AccessibleNamespace: %s", err))
	}

	*out = *target.(*AccessibleNamespace)
}

// Validate valides the current information stored into the structure.
func (o *AccessibleNamespace) Validate() error {

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
func (*AccessibleNamespace) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AccessibleNamespaceAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AccessibleNamespaceLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AccessibleNamespace) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AccessibleNamespaceAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *AccessibleNamespace) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	}

	return nil
}

// AccessibleNamespaceAttributesMap represents the map of attribute for AccessibleNamespace.
var AccessibleNamespaceAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `Identifier of the namespace that is accessible.`,
		Exposed:        true,
		Name:           "ID",
		ReadOnly:       true,
		Type:           "string",
	},
	"Name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the namespace that is accessible.`,
		Exposed:        true,
		Name:           "name",
		ReadOnly:       true,
		Type:           "string",
	},
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
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

// AccessibleNamespaceLowerCaseAttributesMap represents the map of attribute for AccessibleNamespace.
var AccessibleNamespaceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `Identifier of the namespace that is accessible.`,
		Exposed:        true,
		Name:           "ID",
		ReadOnly:       true,
		Type:           "string",
	},
	"name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the namespace that is accessible.`,
		Exposed:        true,
		Name:           "name",
		ReadOnly:       true,
		Type:           "string",
	},
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
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

// SparseAccessibleNamespacesList represents a list of SparseAccessibleNamespaces
type SparseAccessibleNamespacesList []*SparseAccessibleNamespace

// Identity returns the identity of the objects in the list.
func (o SparseAccessibleNamespacesList) Identity() elemental.Identity {

	return AccessibleNamespaceIdentity
}

// Copy returns a pointer to a copy the SparseAccessibleNamespacesList.
func (o SparseAccessibleNamespacesList) Copy() elemental.Identifiables {

	copy := append(SparseAccessibleNamespacesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAccessibleNamespacesList.
func (o SparseAccessibleNamespacesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAccessibleNamespacesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAccessibleNamespace))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAccessibleNamespacesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAccessibleNamespacesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseAccessibleNamespacesList converted to AccessibleNamespacesList.
func (o SparseAccessibleNamespacesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAccessibleNamespacesList) Version() int {

	return 1
}

// SparseAccessibleNamespace represents the sparse version of a accessiblenamespace.
type SparseAccessibleNamespace struct {
	// Identifier of the namespace that is accessible.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Name of the namespace that is accessible.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"-" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAccessibleNamespace returns a new  SparseAccessibleNamespace.
func NewSparseAccessibleNamespace() *SparseAccessibleNamespace {
	return &SparseAccessibleNamespace{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAccessibleNamespace) Identity() elemental.Identity {

	return AccessibleNamespaceIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAccessibleNamespace) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAccessibleNamespace) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAccessibleNamespace) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseAccessibleNamespace{}

	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAccessibleNamespace) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseAccessibleNamespace{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseAccessibleNamespace) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAccessibleNamespace) ToPlain() elemental.PlainIdentifiable {

	out := NewAccessibleNamespace()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}

	return out
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseAccessibleNamespace) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseAccessibleNamespace) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// DeepCopy returns a deep copy if the SparseAccessibleNamespace.
func (o *SparseAccessibleNamespace) DeepCopy() *SparseAccessibleNamespace {

	if o == nil {
		return nil
	}

	out := &SparseAccessibleNamespace{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAccessibleNamespace.
func (o *SparseAccessibleNamespace) DeepCopyInto(out *SparseAccessibleNamespace) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAccessibleNamespace: %s", err))
	}

	*out = *target.(*SparseAccessibleNamespace)
}

type mongoAttributesAccessibleNamespace struct {
	Namespace string `bson:"namespace"`
}
type mongoAttributesSparseAccessibleNamespace struct {
	Namespace *string `bson:"namespace,omitempty"`
}
