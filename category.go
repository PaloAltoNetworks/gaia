package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CategoryIdentity represents the Identity of the object.
var CategoryIdentity = elemental.Identity{
	Name:     "category",
	Category: "categories",
	Package:  "highwind",
	Private:  false,
}

// CategoriesList represents a list of Categories
type CategoriesList []*Category

// Identity returns the identity of the objects in the list.
func (o CategoriesList) Identity() elemental.Identity {

	return CategoryIdentity
}

// Copy returns a pointer to a copy the CategoriesList.
func (o CategoriesList) Copy() elemental.Identifiables {

	copy := append(CategoriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CategoriesList.
func (o CategoriesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CategoriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Category))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CategoriesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CategoriesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the CategoriesList converted to SparseCategoriesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CategoriesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCategoriesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCategory)
	}

	return out
}

// Version returns the version of the content.
func (o CategoriesList) Version() int {

	return 1
}

// Category represents the model of a category
type Category struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCategory returns a new *Category
func NewCategory() *Category {

	return &Category{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Category) Identity() elemental.Identity {

	return CategoryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Category) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Category) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Category) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCategory{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Description = o.Description
	s.Name = o.Name

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Category) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCategory{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Description = s.Description
	o.Name = s.Name

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Category) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Category) BleveType() string {

	return "category"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Category) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *Category) Doc() string {

	return `Allows you to categorize services.`
}

func (o *Category) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetDescription returns the Description of the receiver.
func (o *Category) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *Category) SetDescription(description string) {

	o.Description = description
}

// GetName returns the Name of the receiver.
func (o *Category) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *Category) SetName(name string) {

	o.Name = name
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Category) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCategory{
			ID:          &o.ID,
			Description: &o.Description,
			Name:        &o.Name,
		}
	}

	sp := &SparseCategory{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "description":
			sp.Description = &(o.Description)
		case "name":
			sp.Name = &(o.Name)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCategory to the object.
func (o *Category) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCategory)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
}

// DeepCopy returns a deep copy if the Category.
func (o *Category) DeepCopy() *Category {

	if o == nil {
		return nil
	}

	out := &Category{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Category.
func (o *Category) DeepCopyInto(out *Category) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Category: %s", err))
	}

	*out = *target.(*Category)
}

// Validate valides the current information stored into the structure.
func (o *Category) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
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
func (*Category) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CategoryAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CategoryLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Category) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CategoryAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Category) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "description":
		return o.Description
	case "name":
		return o.Name
	}

	return nil
}

// CategoryAttributesMap represents the map of attribute for Category.
var CategoryAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
}

// CategoryLowerCaseAttributesMap represents the map of attribute for Category.
var CategoryLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
}

// SparseCategoriesList represents a list of SparseCategories
type SparseCategoriesList []*SparseCategory

// Identity returns the identity of the objects in the list.
func (o SparseCategoriesList) Identity() elemental.Identity {

	return CategoryIdentity
}

// Copy returns a pointer to a copy the SparseCategoriesList.
func (o SparseCategoriesList) Copy() elemental.Identifiables {

	copy := append(SparseCategoriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCategoriesList.
func (o SparseCategoriesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCategoriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCategory))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCategoriesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCategoriesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseCategoriesList converted to CategoriesList.
func (o SparseCategoriesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCategoriesList) Version() int {

	return 1
}

// SparseCategory represents the sparse version of a category.
type SparseCategory struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCategory returns a new  SparseCategory.
func NewSparseCategory() *SparseCategory {
	return &SparseCategory{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCategory) Identity() elemental.Identity {

	return CategoryIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCategory) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCategory) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCategory) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCategory{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.Name != nil {
		s.Name = o.Name
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCategory) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCategory{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Name != nil {
		o.Name = s.Name
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCategory) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCategory) ToPlain() elemental.PlainIdentifiable {

	out := NewCategory()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Name != nil {
		out.Name = *o.Name
	}

	return out
}

// GetDescription returns the Description of the receiver.
func (o *SparseCategory) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseCategory) SetDescription(description string) {

	o.Description = &description
}

// GetName returns the Name of the receiver.
func (o *SparseCategory) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseCategory) SetName(name string) {

	o.Name = &name
}

// DeepCopy returns a deep copy if the SparseCategory.
func (o *SparseCategory) DeepCopy() *SparseCategory {

	if o == nil {
		return nil
	}

	out := &SparseCategory{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCategory.
func (o *SparseCategory) DeepCopyInto(out *SparseCategory) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCategory: %s", err))
	}

	*out = *target.(*SparseCategory)
}

type mongoAttributesCategory struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Description string        `bson:"description"`
	Name        string        `bson:"name"`
}
type mongoAttributesSparseCategory struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Description *string       `bson:"description,omitempty"`
	Name        *string       `bson:"name,omitempty"`
}
