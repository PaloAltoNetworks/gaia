package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// LogoutIdentity represents the Identity of the object.
var LogoutIdentity = elemental.Identity{
	Name:     "logout",
	Category: "logout",
	Package:  "midgard",
	Private:  false,
}

// LogoutsList represents a list of Logouts
type LogoutsList []*Logout

// Identity returns the identity of the objects in the list.
func (o LogoutsList) Identity() elemental.Identity {

	return LogoutIdentity
}

// Copy returns a pointer to a copy the LogoutsList.
func (o LogoutsList) Copy() elemental.Identifiables {

	copy := append(LogoutsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the LogoutsList.
func (o LogoutsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(LogoutsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Logout))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o LogoutsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o LogoutsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the LogoutsList converted to SparseLogoutsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o LogoutsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseLogoutsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseLogout)
	}

	return out
}

// Version returns the version of the content.
func (o LogoutsList) Version() int {

	return 1
}

// Logout represents the model of a logout
type Logout struct {
	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewLogout returns a new *Logout
func NewLogout() *Logout {

	return &Logout{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Logout) Identity() elemental.Identity {

	return LogoutIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Logout) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Logout) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Logout) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesLogout{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Logout) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesLogout{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Logout) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Logout) BleveType() string {

	return "logout"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Logout) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Logout) Doc() string {

	return `Perform logout operations. This is only used to unset the secure cookie token
for now.`
}

func (o *Logout) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Logout) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseLogout{}
	}

	sp := &SparseLogout{}
	for _, f := range fields {
		switch f {
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseLogout to the object.
func (o *Logout) Patch(sparse elemental.SparseIdentifiable) {
}

// DeepCopy returns a deep copy if the Logout.
func (o *Logout) DeepCopy() *Logout {

	if o == nil {
		return nil
	}

	out := &Logout{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Logout.
func (o *Logout) DeepCopyInto(out *Logout) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Logout: %s", err))
	}

	*out = *target.(*Logout)
}

// Validate valides the current information stored into the structure.
func (o *Logout) Validate() error {

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
func (*Logout) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := LogoutAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return LogoutLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Logout) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return LogoutAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Logout) ValueForAttribute(name string) interface{} {

	switch name {
	}

	return nil
}

// LogoutAttributesMap represents the map of attribute for Logout.
var LogoutAttributesMap = map[string]elemental.AttributeSpecification{}

// LogoutLowerCaseAttributesMap represents the map of attribute for Logout.
var LogoutLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{}

// SparseLogoutsList represents a list of SparseLogouts
type SparseLogoutsList []*SparseLogout

// Identity returns the identity of the objects in the list.
func (o SparseLogoutsList) Identity() elemental.Identity {

	return LogoutIdentity
}

// Copy returns a pointer to a copy the SparseLogoutsList.
func (o SparseLogoutsList) Copy() elemental.Identifiables {

	copy := append(SparseLogoutsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseLogoutsList.
func (o SparseLogoutsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseLogoutsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseLogout))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseLogoutsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseLogoutsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseLogoutsList converted to LogoutsList.
func (o SparseLogoutsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseLogoutsList) Version() int {

	return 1
}

// SparseLogout represents the sparse version of a logout.
type SparseLogout struct {
	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseLogout returns a new  SparseLogout.
func NewSparseLogout() *SparseLogout {
	return &SparseLogout{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseLogout) Identity() elemental.Identity {

	return LogoutIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseLogout) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseLogout) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseLogout) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseLogout{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseLogout) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseLogout{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseLogout) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseLogout) ToPlain() elemental.PlainIdentifiable {

	out := NewLogout()

	return out
}

// DeepCopy returns a deep copy if the SparseLogout.
func (o *SparseLogout) DeepCopy() *SparseLogout {

	if o == nil {
		return nil
	}

	out := &SparseLogout{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseLogout.
func (o *SparseLogout) DeepCopyInto(out *SparseLogout) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseLogout: %s", err))
	}

	*out = *target.(*SparseLogout)
}

type mongoAttributesLogout struct {
}
type mongoAttributesSparseLogout struct {
}
