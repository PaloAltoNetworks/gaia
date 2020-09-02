package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TenantIdentity represents the Identity of the object.
var TenantIdentity = elemental.Identity{
	Name:     "tenant",
	Category: "tenants",
	Package:  "karl",
	Private:  false,
}

// TenantsList represents a list of Tenants
type TenantsList []*Tenant

// Identity returns the identity of the objects in the list.
func (o TenantsList) Identity() elemental.Identity {

	return TenantIdentity
}

// Copy returns a pointer to a copy the TenantsList.
func (o TenantsList) Copy() elemental.Identifiables {

	copy := append(TenantsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TenantsList.
func (o TenantsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(TenantsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Tenant))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TenantsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TenantsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the TenantsList converted to SparseTenantsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o TenantsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseTenantsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseTenant)
	}

	return out
}

// Version returns the version of the content.
func (o TenantsList) Version() int {

	return 1
}

// Tenant represents the model of a tenant
type Tenant struct {
	// The external ID of the tenant.
	ExternalID string `json:"externalID" msgpack:"externalID" bson:"-" mapstructure:"externalID,omitempty"`

	// The name of the tenant.
	Name string `json:"name" msgpack:"name" bson:"-" mapstructure:"name,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewTenant returns a new *Tenant
func NewTenant() *Tenant {

	return &Tenant{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Tenant) Identity() elemental.Identity {

	return TenantIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Tenant) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Tenant) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Tenant) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesTenant{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Tenant) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesTenant{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Tenant) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Tenant) BleveType() string {

	return "tenant"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Tenant) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Tenant) Doc() string {

	return `Can be used to create a tenant's namespace and API authorization policy to grant
access.`
}

func (o *Tenant) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Tenant) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseTenant{
			ExternalID: &o.ExternalID,
			Name:       &o.Name,
		}
	}

	sp := &SparseTenant{}
	for _, f := range fields {
		switch f {
		case "externalID":
			sp.ExternalID = &(o.ExternalID)
		case "name":
			sp.Name = &(o.Name)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseTenant to the object.
func (o *Tenant) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseTenant)
	if so.ExternalID != nil {
		o.ExternalID = *so.ExternalID
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
}

// DeepCopy returns a deep copy if the Tenant.
func (o *Tenant) DeepCopy() *Tenant {

	if o == nil {
		return nil
	}

	out := &Tenant{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Tenant.
func (o *Tenant) DeepCopyInto(out *Tenant) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Tenant: %s", err))
	}

	*out = *target.(*Tenant)
}

// Validate valides the current information stored into the structure.
func (o *Tenant) Validate() error {

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
func (*Tenant) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TenantAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TenantLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Tenant) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TenantAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Tenant) ValueForAttribute(name string) interface{} {

	switch name {
	case "externalID":
		return o.ExternalID
	case "name":
		return o.Name
	}

	return nil
}

// TenantAttributesMap represents the map of attribute for Tenant.
var TenantAttributesMap = map[string]elemental.AttributeSpecification{
	"ExternalID": {
		AllowedChoices: []string{},
		ConvertedName:  "ExternalID",
		Description:    `The external ID of the tenant.`,
		Exposed:        true,
		Name:           "externalID",
		Required:       true,
		Transient:      true,
		Type:           "string",
	},
	"Name": {
		AllowedChars:   `^[a-zA-Z0-9-_/]+$`,
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `The name of the tenant.`,
		Exposed:        true,
		Name:           "name",
		Required:       true,
		Transient:      true,
		Type:           "string",
	},
}

// TenantLowerCaseAttributesMap represents the map of attribute for Tenant.
var TenantLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"externalid": {
		AllowedChoices: []string{},
		ConvertedName:  "ExternalID",
		Description:    `The external ID of the tenant.`,
		Exposed:        true,
		Name:           "externalID",
		Required:       true,
		Transient:      true,
		Type:           "string",
	},
	"name": {
		AllowedChars:   `^[a-zA-Z0-9-_/]+$`,
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `The name of the tenant.`,
		Exposed:        true,
		Name:           "name",
		Required:       true,
		Transient:      true,
		Type:           "string",
	},
}

// SparseTenantsList represents a list of SparseTenants
type SparseTenantsList []*SparseTenant

// Identity returns the identity of the objects in the list.
func (o SparseTenantsList) Identity() elemental.Identity {

	return TenantIdentity
}

// Copy returns a pointer to a copy the SparseTenantsList.
func (o SparseTenantsList) Copy() elemental.Identifiables {

	copy := append(SparseTenantsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseTenantsList.
func (o SparseTenantsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseTenantsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseTenant))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseTenantsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseTenantsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseTenantsList converted to TenantsList.
func (o SparseTenantsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseTenantsList) Version() int {

	return 1
}

// SparseTenant represents the sparse version of a tenant.
type SparseTenant struct {
	// The external ID of the tenant.
	ExternalID *string `json:"externalID,omitempty" msgpack:"externalID,omitempty" bson:"-" mapstructure:"externalID,omitempty"`

	// The name of the tenant.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"-" mapstructure:"name,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseTenant returns a new  SparseTenant.
func NewSparseTenant() *SparseTenant {
	return &SparseTenant{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseTenant) Identity() elemental.Identity {

	return TenantIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseTenant) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseTenant) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTenant) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseTenant{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTenant) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseTenant{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseTenant) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseTenant) ToPlain() elemental.PlainIdentifiable {

	out := NewTenant()
	if o.ExternalID != nil {
		out.ExternalID = *o.ExternalID
	}
	if o.Name != nil {
		out.Name = *o.Name
	}

	return out
}

// DeepCopy returns a deep copy if the SparseTenant.
func (o *SparseTenant) DeepCopy() *SparseTenant {

	if o == nil {
		return nil
	}

	out := &SparseTenant{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseTenant.
func (o *SparseTenant) DeepCopyInto(out *SparseTenant) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseTenant: %s", err))
	}

	*out = *target.(*SparseTenant)
}

type mongoAttributesTenant struct {
}
type mongoAttributesSparseTenant struct {
}
