package gaia

import (
	"fmt"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PasswordResetIdentity represents the Identity of the object.
var PasswordResetIdentity = elemental.Identity{
	Name:     "passwordreset",
	Category: "passwordreset",
	Package:  "vince",
	Private:  false,
}

// PasswordResetsList represents a list of PasswordResets
type PasswordResetsList []*PasswordReset

// Identity returns the identity of the objects in the list.
func (o PasswordResetsList) Identity() elemental.Identity {

	return PasswordResetIdentity
}

// Copy returns a pointer to a copy the PasswordResetsList.
func (o PasswordResetsList) Copy() elemental.Identifiables {

	copy := append(PasswordResetsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PasswordResetsList.
func (o PasswordResetsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PasswordResetsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PasswordReset))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PasswordResetsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PasswordResetsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PasswordResetsList converted to SparsePasswordResetsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PasswordResetsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePasswordResetsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePasswordReset)
	}

	return out
}

// Version returns the version of the content.
func (o PasswordResetsList) Version() int {

	return 1
}

// PasswordReset represents the model of a passwordreset
type PasswordReset struct {
	// Password contains the new password.
	Password string `json:"password" msgpack:"password" bson:"-" mapstructure:"password,omitempty"`

	// Token contains the reset password token.
	Token string `json:"token" msgpack:"token" bson:"-" mapstructure:"token,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPasswordReset returns a new *PasswordReset
func NewPasswordReset() *PasswordReset {

	return &PasswordReset{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *PasswordReset) Identity() elemental.Identity {

	return PasswordResetIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PasswordReset) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PasswordReset) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *PasswordReset) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *PasswordReset) BleveType() string {

	return "passwordreset"
}

// DefaultOrder returns the list of default ordering fields.
func (o *PasswordReset) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PasswordReset) Doc() string {

	return `Used to reset an account password.`
}

func (o *PasswordReset) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PasswordReset) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePasswordReset{
			Password: &o.Password,
			Token:    &o.Token,
		}
	}

	sp := &SparsePasswordReset{}
	for _, f := range fields {
		switch f {
		case "password":
			sp.Password = &(o.Password)
		case "token":
			sp.Token = &(o.Token)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePasswordReset to the object.
func (o *PasswordReset) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePasswordReset)
	if so.Password != nil {
		o.Password = *so.Password
	}
	if so.Token != nil {
		o.Token = *so.Token
	}
}

// DeepCopy returns a deep copy if the PasswordReset.
func (o *PasswordReset) DeepCopy() *PasswordReset {

	if o == nil {
		return nil
	}

	out := &PasswordReset{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PasswordReset.
func (o *PasswordReset) DeepCopyInto(out *PasswordReset) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PasswordReset: %s", err))
	}

	*out = *target.(*PasswordReset)
}

// Validate valides the current information stored into the structure.
func (o *PasswordReset) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("password", o.Password); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("token", o.Token); err != nil {
		requiredErrors = requiredErrors.Append(err)
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
func (*PasswordReset) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PasswordResetAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PasswordResetLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PasswordReset) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PasswordResetAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PasswordReset) ValueForAttribute(name string) interface{} {

	switch name {
	case "password":
		return o.Password
	case "token":
		return o.Token
	}

	return nil
}

// PasswordResetAttributesMap represents the map of attribute for PasswordReset.
var PasswordResetAttributesMap = map[string]elemental.AttributeSpecification{
	"Password": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Password",
		Description:    `Password contains the new password.`,
		Exposed:        true,
		Name:           "password",
		Required:       true,
		Type:           "string",
	},
	"Token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Token",
		Description:    `Token contains the reset password token.`,
		Exposed:        true,
		Name:           "token",
		Required:       true,
		Type:           "string",
	},
}

// PasswordResetLowerCaseAttributesMap represents the map of attribute for PasswordReset.
var PasswordResetLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"password": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Password",
		Description:    `Password contains the new password.`,
		Exposed:        true,
		Name:           "password",
		Required:       true,
		Type:           "string",
	},
	"token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Token",
		Description:    `Token contains the reset password token.`,
		Exposed:        true,
		Name:           "token",
		Required:       true,
		Type:           "string",
	},
}

// SparsePasswordResetsList represents a list of SparsePasswordResets
type SparsePasswordResetsList []*SparsePasswordReset

// Identity returns the identity of the objects in the list.
func (o SparsePasswordResetsList) Identity() elemental.Identity {

	return PasswordResetIdentity
}

// Copy returns a pointer to a copy the SparsePasswordResetsList.
func (o SparsePasswordResetsList) Copy() elemental.Identifiables {

	copy := append(SparsePasswordResetsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePasswordResetsList.
func (o SparsePasswordResetsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePasswordResetsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePasswordReset))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePasswordResetsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePasswordResetsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePasswordResetsList converted to PasswordResetsList.
func (o SparsePasswordResetsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePasswordResetsList) Version() int {

	return 1
}

// SparsePasswordReset represents the sparse version of a passwordreset.
type SparsePasswordReset struct {
	// Password contains the new password.
	Password *string `json:"password,omitempty" msgpack:"password,omitempty" bson:"-" mapstructure:"password,omitempty"`

	// Token contains the reset password token.
	Token *string `json:"token,omitempty" msgpack:"token,omitempty" bson:"-" mapstructure:"token,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePasswordReset returns a new  SparsePasswordReset.
func NewSparsePasswordReset() *SparsePasswordReset {
	return &SparsePasswordReset{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePasswordReset) Identity() elemental.Identity {

	return PasswordResetIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePasswordReset) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePasswordReset) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparsePasswordReset) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePasswordReset) ToPlain() elemental.PlainIdentifiable {

	out := NewPasswordReset()
	if o.Password != nil {
		out.Password = *o.Password
	}
	if o.Token != nil {
		out.Token = *o.Token
	}

	return out
}

// DeepCopy returns a deep copy if the SparsePasswordReset.
func (o *SparsePasswordReset) DeepCopy() *SparsePasswordReset {

	if o == nil {
		return nil
	}

	out := &SparsePasswordReset{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePasswordReset.
func (o *SparsePasswordReset) DeepCopyInto(out *SparsePasswordReset) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePasswordReset: %s", err))
	}

	*out = *target.(*SparsePasswordReset)
}
