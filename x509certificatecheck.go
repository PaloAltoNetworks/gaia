package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// X509CertificateCheckIdentity represents the Identity of the object.
var X509CertificateCheckIdentity = elemental.Identity{
	Name:     "x509certificatecheck",
	Category: "x509certificatechecks",
	Package:  "barret",
	Private:  true,
}

// X509CertificateChecksList represents a list of X509CertificateChecks
type X509CertificateChecksList []*X509CertificateCheck

// Identity returns the identity of the objects in the list.
func (o X509CertificateChecksList) Identity() elemental.Identity {

	return X509CertificateCheckIdentity
}

// Copy returns a pointer to a copy the X509CertificateChecksList.
func (o X509CertificateChecksList) Copy() elemental.Identifiables {

	copy := append(X509CertificateChecksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the X509CertificateChecksList.
func (o X509CertificateChecksList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(X509CertificateChecksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*X509CertificateCheck))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o X509CertificateChecksList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o X509CertificateChecksList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the X509CertificateChecksList converted to SparseX509CertificateChecksList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o X509CertificateChecksList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseX509CertificateChecksList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseX509CertificateCheck)
	}

	return out
}

// Version returns the version of the content.
func (o X509CertificateChecksList) Version() int {

	return 1
}

// X509CertificateCheck represents the model of a x509certificatecheck
type X509CertificateCheck struct {
	// ID contains the certificate serialNumber.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewX509CertificateCheck returns a new *X509CertificateCheck
func NewX509CertificateCheck() *X509CertificateCheck {

	return &X509CertificateCheck{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *X509CertificateCheck) Identity() elemental.Identity {

	return X509CertificateCheckIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *X509CertificateCheck) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *X509CertificateCheck) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *X509CertificateCheck) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesX509CertificateCheck{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *X509CertificateCheck) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesX509CertificateCheck{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *X509CertificateCheck) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *X509CertificateCheck) BleveType() string {

	return "x509certificatecheck"
}

// DefaultOrder returns the list of default ordering fields.
func (o *X509CertificateCheck) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *X509CertificateCheck) Doc() string {

	return `Verifies a certificate has not been revoked.`
}

func (o *X509CertificateCheck) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *X509CertificateCheck) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseX509CertificateCheck{
			ID: &o.ID,
		}
	}

	sp := &SparseX509CertificateCheck{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseX509CertificateCheck to the object.
func (o *X509CertificateCheck) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseX509CertificateCheck)
	if so.ID != nil {
		o.ID = *so.ID
	}
}

// DeepCopy returns a deep copy if the X509CertificateCheck.
func (o *X509CertificateCheck) DeepCopy() *X509CertificateCheck {

	if o == nil {
		return nil
	}

	out := &X509CertificateCheck{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *X509CertificateCheck.
func (o *X509CertificateCheck) DeepCopyInto(out *X509CertificateCheck) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy X509CertificateCheck: %s", err))
	}

	*out = *target.(*X509CertificateCheck)
}

// Validate valides the current information stored into the structure.
func (o *X509CertificateCheck) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("ID", o.ID); err != nil {
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
func (*X509CertificateCheck) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := X509CertificateCheckAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return X509CertificateCheckLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*X509CertificateCheck) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return X509CertificateCheckAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *X509CertificateCheck) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	}

	return nil
}

// X509CertificateCheckAttributesMap represents the map of attribute for X509CertificateCheck.
var X509CertificateCheckAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID contains the certificate serialNumber.`,
		Exposed:        true,
		Identifier:     true,
		Name:           "ID",
		Required:       true,
		Type:           "string",
	},
}

// X509CertificateCheckLowerCaseAttributesMap represents the map of attribute for X509CertificateCheck.
var X509CertificateCheckLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID contains the certificate serialNumber.`,
		Exposed:        true,
		Identifier:     true,
		Name:           "ID",
		Required:       true,
		Type:           "string",
	},
}

// SparseX509CertificateChecksList represents a list of SparseX509CertificateChecks
type SparseX509CertificateChecksList []*SparseX509CertificateCheck

// Identity returns the identity of the objects in the list.
func (o SparseX509CertificateChecksList) Identity() elemental.Identity {

	return X509CertificateCheckIdentity
}

// Copy returns a pointer to a copy the SparseX509CertificateChecksList.
func (o SparseX509CertificateChecksList) Copy() elemental.Identifiables {

	copy := append(SparseX509CertificateChecksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseX509CertificateChecksList.
func (o SparseX509CertificateChecksList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseX509CertificateChecksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseX509CertificateCheck))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseX509CertificateChecksList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseX509CertificateChecksList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseX509CertificateChecksList converted to X509CertificateChecksList.
func (o SparseX509CertificateChecksList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseX509CertificateChecksList) Version() int {

	return 1
}

// SparseX509CertificateCheck represents the sparse version of a x509certificatecheck.
type SparseX509CertificateCheck struct {
	// ID contains the certificate serialNumber.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseX509CertificateCheck returns a new  SparseX509CertificateCheck.
func NewSparseX509CertificateCheck() *SparseX509CertificateCheck {
	return &SparseX509CertificateCheck{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseX509CertificateCheck) Identity() elemental.Identity {

	return X509CertificateCheckIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseX509CertificateCheck) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseX509CertificateCheck) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseX509CertificateCheck) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseX509CertificateCheck{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseX509CertificateCheck) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseX509CertificateCheck{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseX509CertificateCheck) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseX509CertificateCheck) ToPlain() elemental.PlainIdentifiable {

	out := NewX509CertificateCheck()
	if o.ID != nil {
		out.ID = *o.ID
	}

	return out
}

// DeepCopy returns a deep copy if the SparseX509CertificateCheck.
func (o *SparseX509CertificateCheck) DeepCopy() *SparseX509CertificateCheck {

	if o == nil {
		return nil
	}

	out := &SparseX509CertificateCheck{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseX509CertificateCheck.
func (o *SparseX509CertificateCheck) DeepCopyInto(out *SparseX509CertificateCheck) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseX509CertificateCheck: %s", err))
	}

	*out = *target.(*SparseX509CertificateCheck)
}

type mongoAttributesX509CertificateCheck struct {
}
type mongoAttributesSparseX509CertificateCheck struct {
}
