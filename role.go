package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// RoleIdentity represents the Identity of the object.
var RoleIdentity = elemental.Identity{
	Name:     "role",
	Category: "roles",
	Package:  "cid",
	Private:  false,
}

// RolesList represents a list of Roles
type RolesList []*Role

// Identity returns the identity of the objects in the list.
func (o RolesList) Identity() elemental.Identity {

	return RoleIdentity
}

// Copy returns a pointer to a copy the RolesList.
func (o RolesList) Copy() elemental.Identifiables {

	copy := append(RolesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the RolesList.
func (o RolesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(RolesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Role))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o RolesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o RolesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the RolesList converted to SparseRolesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o RolesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseRolesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseRole)
	}

	return out
}

// Version returns the version of the content.
func (o RolesList) Version() int {

	return 1
}

// Role represents the model of a role
type Role struct {
	// Authorizations of the role.
	Authorizations map[string][]string `json:"authorizations" msgpack:"authorizations" bson:"-" mapstructure:"authorizations,omitempty"`

	// Description of the role.
	Description string `json:"description" msgpack:"description" bson:"-" mapstructure:"description,omitempty"`

	// Key of the role.
	Key string `json:"key" msgpack:"key" bson:"-" mapstructure:"key,omitempty"`

	// Name of the role.
	Name string `json:"name" msgpack:"name" bson:"-" mapstructure:"name,omitempty"`

	// Set to `true` to make the role private and hidden from the UI.
	Private bool `json:"private" msgpack:"private" bson:"-" mapstructure:"private,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewRole returns a new *Role
func NewRole() *Role {

	return &Role{
		ModelVersion:   1,
		Authorizations: map[string][]string{},
	}
}

// Identity returns the Identity of the object.
func (o *Role) Identity() elemental.Identity {

	return RoleIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Role) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Role) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Role) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesRole{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Role) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesRole{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Role) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Role) BleveType() string {

	return "role"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Role) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Role) Doc() string {

	return `Returns the available roles that can be used with API authorizations.`
}

func (o *Role) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Role) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseRole{
			Authorizations: &o.Authorizations,
			Description:    &o.Description,
			Key:            &o.Key,
			Name:           &o.Name,
			Private:        &o.Private,
		}
	}

	sp := &SparseRole{}
	for _, f := range fields {
		switch f {
		case "authorizations":
			sp.Authorizations = &(o.Authorizations)
		case "description":
			sp.Description = &(o.Description)
		case "key":
			sp.Key = &(o.Key)
		case "name":
			sp.Name = &(o.Name)
		case "private":
			sp.Private = &(o.Private)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseRole to the object.
func (o *Role) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseRole)
	if so.Authorizations != nil {
		o.Authorizations = *so.Authorizations
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Key != nil {
		o.Key = *so.Key
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Private != nil {
		o.Private = *so.Private
	}
}

// DeepCopy returns a deep copy if the Role.
func (o *Role) DeepCopy() *Role {

	if o == nil {
		return nil
	}

	out := &Role{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Role.
func (o *Role) DeepCopyInto(out *Role) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Role: %s", err))
	}

	*out = *target.(*Role)
}

// Validate valides the current information stored into the structure.
func (o *Role) Validate() error {

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
func (*Role) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := RoleAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return RoleLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Role) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RoleAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Role) ValueForAttribute(name string) interface{} {

	switch name {
	case "authorizations":
		return o.Authorizations
	case "description":
		return o.Description
	case "key":
		return o.Key
	case "name":
		return o.Name
	case "private":
		return o.Private
	}

	return nil
}

// RoleAttributesMap represents the map of attribute for Role.
var RoleAttributesMap = map[string]elemental.AttributeSpecification{
	"Authorizations": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Authorizations",
		Description:    `Authorizations of the role.`,
		Exposed:        true,
		Name:           "authorizations",
		ReadOnly:       true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"Description": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Description",
		Description:    `Description of the role.`,
		Exposed:        true,
		Name:           "description",
		ReadOnly:       true,
		Type:           "string",
	},
	"Key": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Key",
		Description:    `Key of the role.`,
		Exposed:        true,
		Name:           "key",
		ReadOnly:       true,
		Type:           "string",
	},
	"Name": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Name",
		Description:    `Name of the role.`,
		Exposed:        true,
		Name:           "name",
		ReadOnly:       true,
		Type:           "string",
	},
	"Private": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Private",
		Description:    `Set to ` + "`" + `true` + "`" + ` to make the role private and hidden from the UI.`,
		Exposed:        true,
		Name:           "private",
		ReadOnly:       true,
		Type:           "boolean",
	},
}

// RoleLowerCaseAttributesMap represents the map of attribute for Role.
var RoleLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"authorizations": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Authorizations",
		Description:    `Authorizations of the role.`,
		Exposed:        true,
		Name:           "authorizations",
		ReadOnly:       true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"description": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Description",
		Description:    `Description of the role.`,
		Exposed:        true,
		Name:           "description",
		ReadOnly:       true,
		Type:           "string",
	},
	"key": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Key",
		Description:    `Key of the role.`,
		Exposed:        true,
		Name:           "key",
		ReadOnly:       true,
		Type:           "string",
	},
	"name": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Name",
		Description:    `Name of the role.`,
		Exposed:        true,
		Name:           "name",
		ReadOnly:       true,
		Type:           "string",
	},
	"private": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Private",
		Description:    `Set to ` + "`" + `true` + "`" + ` to make the role private and hidden from the UI.`,
		Exposed:        true,
		Name:           "private",
		ReadOnly:       true,
		Type:           "boolean",
	},
}

// SparseRolesList represents a list of SparseRoles
type SparseRolesList []*SparseRole

// Identity returns the identity of the objects in the list.
func (o SparseRolesList) Identity() elemental.Identity {

	return RoleIdentity
}

// Copy returns a pointer to a copy the SparseRolesList.
func (o SparseRolesList) Copy() elemental.Identifiables {

	copy := append(SparseRolesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseRolesList.
func (o SparseRolesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseRolesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseRole))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseRolesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseRolesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseRolesList converted to RolesList.
func (o SparseRolesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseRolesList) Version() int {

	return 1
}

// SparseRole represents the sparse version of a role.
type SparseRole struct {
	// Authorizations of the role.
	Authorizations *map[string][]string `json:"authorizations,omitempty" msgpack:"authorizations,omitempty" bson:"-" mapstructure:"authorizations,omitempty"`

	// Description of the role.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"-" mapstructure:"description,omitempty"`

	// Key of the role.
	Key *string `json:"key,omitempty" msgpack:"key,omitempty" bson:"-" mapstructure:"key,omitempty"`

	// Name of the role.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"-" mapstructure:"name,omitempty"`

	// Set to `true` to make the role private and hidden from the UI.
	Private *bool `json:"private,omitempty" msgpack:"private,omitempty" bson:"-" mapstructure:"private,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseRole returns a new  SparseRole.
func NewSparseRole() *SparseRole {
	return &SparseRole{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseRole) Identity() elemental.Identity {

	return RoleIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseRole) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseRole) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseRole) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseRole{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseRole) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseRole{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseRole) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseRole) ToPlain() elemental.PlainIdentifiable {

	out := NewRole()
	if o.Authorizations != nil {
		out.Authorizations = *o.Authorizations
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Key != nil {
		out.Key = *o.Key
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Private != nil {
		out.Private = *o.Private
	}

	return out
}

// DeepCopy returns a deep copy if the SparseRole.
func (o *SparseRole) DeepCopy() *SparseRole {

	if o == nil {
		return nil
	}

	out := &SparseRole{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseRole.
func (o *SparseRole) DeepCopyInto(out *SparseRole) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseRole: %s", err))
	}

	*out = *target.(*SparseRole)
}

type mongoAttributesRole struct {
}
type mongoAttributesSparseRole struct {
}
