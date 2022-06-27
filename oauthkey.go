package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// OAUTHKeyIdentity represents the Identity of the object.
var OAUTHKeyIdentity = elemental.Identity{
	Name:     "oauthkey",
	Category: "oauthkeys",
	Package:  "cactuar",
	Private:  false,
}

// OAUTHKeysList represents a list of OAUTHKeys
type OAUTHKeysList []*OAUTHKey

// Identity returns the identity of the objects in the list.
func (o OAUTHKeysList) Identity() elemental.Identity {

	return OAUTHKeyIdentity
}

// Copy returns a pointer to a copy the OAUTHKeysList.
func (o OAUTHKeysList) Copy() elemental.Identifiables {

	copy := append(OAUTHKeysList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the OAUTHKeysList.
func (o OAUTHKeysList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(OAUTHKeysList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*OAUTHKey))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o OAUTHKeysList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o OAUTHKeysList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the OAUTHKeysList converted to SparseOAUTHKeysList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o OAUTHKeysList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseOAUTHKeysList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseOAUTHKey)
	}

	return out
}

// Version returns the version of the content.
func (o OAUTHKeysList) Version() int {

	return 1
}

// OAUTHKey represents the model of a oauthkey
type OAUTHKey struct {
	// KeyString is the JWKS key response for an OAUTH verifier. It provides the OAUTH
	// compatible signing keys.
	KeyString string `json:"keyString" msgpack:"keyString" bson:"-" mapstructure:"keyString,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewOAUTHKey returns a new *OAUTHKey
func NewOAUTHKey() *OAUTHKey {

	return &OAUTHKey{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *OAUTHKey) Identity() elemental.Identity {

	return OAUTHKeyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *OAUTHKey) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *OAUTHKey) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *OAUTHKey) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesOAUTHKey{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *OAUTHKey) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesOAUTHKey{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *OAUTHKey) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *OAUTHKey) BleveType() string {

	return "oauthkey"
}

// DefaultOrder returns the list of default ordering fields.
func (o *OAUTHKey) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *OAUTHKey) Doc() string {

	return `OAUTHInfo provides the information for an OAUTH server to retrieve the secrets
that can validate a JWT token issued by us.`
}

func (o *OAUTHKey) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *OAUTHKey) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseOAUTHKey{
			KeyString: &o.KeyString,
		}
	}

	sp := &SparseOAUTHKey{}
	for _, f := range fields {
		switch f {
		case "keyString":
			sp.KeyString = &(o.KeyString)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseOAUTHKey to the object.
func (o *OAUTHKey) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseOAUTHKey)
	if so.KeyString != nil {
		o.KeyString = *so.KeyString
	}
}

// DeepCopy returns a deep copy if the OAUTHKey.
func (o *OAUTHKey) DeepCopy() *OAUTHKey {

	if o == nil {
		return nil
	}

	out := &OAUTHKey{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *OAUTHKey.
func (o *OAUTHKey) DeepCopyInto(out *OAUTHKey) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy OAUTHKey: %s", err))
	}

	*out = *target.(*OAUTHKey)
}

// Validate valides the current information stored into the structure.
func (o *OAUTHKey) Validate() error {

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
func (*OAUTHKey) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := OAUTHKeyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return OAUTHKeyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*OAUTHKey) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return OAUTHKeyAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *OAUTHKey) ValueForAttribute(name string) interface{} {

	switch name {
	case "keyString":
		return o.KeyString
	}

	return nil
}

// OAUTHKeyAttributesMap represents the map of attribute for OAUTHKey.
var OAUTHKeyAttributesMap = map[string]elemental.AttributeSpecification{
	"KeyString": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "KeyString",
		Description: `KeyString is the JWKS key response for an OAUTH verifier. It provides the OAUTH
compatible signing keys.`,
		Exposed:  true,
		Name:     "keyString",
		ReadOnly: true,
		Type:     "string",
	},
}

// OAUTHKeyLowerCaseAttributesMap represents the map of attribute for OAUTHKey.
var OAUTHKeyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"keystring": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "KeyString",
		Description: `KeyString is the JWKS key response for an OAUTH verifier. It provides the OAUTH
compatible signing keys.`,
		Exposed:  true,
		Name:     "keyString",
		ReadOnly: true,
		Type:     "string",
	},
}

// SparseOAUTHKeysList represents a list of SparseOAUTHKeys
type SparseOAUTHKeysList []*SparseOAUTHKey

// Identity returns the identity of the objects in the list.
func (o SparseOAUTHKeysList) Identity() elemental.Identity {

	return OAUTHKeyIdentity
}

// Copy returns a pointer to a copy the SparseOAUTHKeysList.
func (o SparseOAUTHKeysList) Copy() elemental.Identifiables {

	copy := append(SparseOAUTHKeysList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseOAUTHKeysList.
func (o SparseOAUTHKeysList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseOAUTHKeysList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseOAUTHKey))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseOAUTHKeysList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseOAUTHKeysList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseOAUTHKeysList converted to OAUTHKeysList.
func (o SparseOAUTHKeysList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseOAUTHKeysList) Version() int {

	return 1
}

// SparseOAUTHKey represents the sparse version of a oauthkey.
type SparseOAUTHKey struct {
	// KeyString is the JWKS key response for an OAUTH verifier. It provides the OAUTH
	// compatible signing keys.
	KeyString *string `json:"keyString,omitempty" msgpack:"keyString,omitempty" bson:"-" mapstructure:"keyString,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseOAUTHKey returns a new  SparseOAUTHKey.
func NewSparseOAUTHKey() *SparseOAUTHKey {
	return &SparseOAUTHKey{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseOAUTHKey) Identity() elemental.Identity {

	return OAUTHKeyIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseOAUTHKey) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseOAUTHKey) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseOAUTHKey) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseOAUTHKey{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseOAUTHKey) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseOAUTHKey{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseOAUTHKey) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseOAUTHKey) ToPlain() elemental.PlainIdentifiable {

	out := NewOAUTHKey()
	if o.KeyString != nil {
		out.KeyString = *o.KeyString
	}

	return out
}

// DeepCopy returns a deep copy if the SparseOAUTHKey.
func (o *SparseOAUTHKey) DeepCopy() *SparseOAUTHKey {

	if o == nil {
		return nil
	}

	out := &SparseOAUTHKey{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseOAUTHKey.
func (o *SparseOAUTHKey) DeepCopyInto(out *SparseOAUTHKey) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseOAUTHKey: %s", err))
	}

	*out = *target.(*SparseOAUTHKey)
}

type mongoAttributesOAUTHKey struct {
}
type mongoAttributesSparseOAUTHKey struct {
}
