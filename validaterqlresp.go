package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ValidateRQLRespIdentity represents the Identity of the object.
var ValidateRQLRespIdentity = elemental.Identity{
	Name:     "validaterqlresp",
	Category: "validaterqlresp",
	Package:  "larl",
	Private:  false,
}

// ValidateRQLRespsList represents a list of ValidateRQLResps
type ValidateRQLRespsList []*ValidateRQLResp

// Identity returns the identity of the objects in the list.
func (o ValidateRQLRespsList) Identity() elemental.Identity {

	return ValidateRQLRespIdentity
}

// Copy returns a pointer to a copy the ValidateRQLRespsList.
func (o ValidateRQLRespsList) Copy() elemental.Identifiables {

	copy := append(ValidateRQLRespsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ValidateRQLRespsList.
func (o ValidateRQLRespsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ValidateRQLRespsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ValidateRQLResp))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ValidateRQLRespsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ValidateRQLRespsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ValidateRQLRespsList converted to SparseValidateRQLRespsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ValidateRQLRespsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseValidateRQLRespsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseValidateRQLResp)
	}

	return out
}

// Version returns the version of the content.
func (o ValidateRQLRespsList) Version() int {

	return 1
}

// ValidateRQLResp represents the model of a validaterqlresp
type ValidateRQLResp struct {
	// TODO.
	Error string `json:"error" msgpack:"error" bson:"-" mapstructure:"error,omitempty"`

	// TODO.
	Status int `json:"status" msgpack:"status" bson:"-" mapstructure:"status,omitempty"`

	// TODO.
	Timestamp int `json:"timestamp" msgpack:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewValidateRQLResp returns a new *ValidateRQLResp
func NewValidateRQLResp() *ValidateRQLResp {

	return &ValidateRQLResp{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *ValidateRQLResp) Identity() elemental.Identity {

	return ValidateRQLRespIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ValidateRQLResp) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ValidateRQLResp) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ValidateRQLResp) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesValidateRQLResp{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ValidateRQLResp) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesValidateRQLResp{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *ValidateRQLResp) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *ValidateRQLResp) BleveType() string {

	return "validaterqlresp"
}

// DefaultOrder returns the list of default ordering fields.
func (o *ValidateRQLResp) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *ValidateRQLResp) Doc() string {

	return `TODO.`
}

func (o *ValidateRQLResp) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ValidateRQLResp) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseValidateRQLResp{
			Error:     &o.Error,
			Status:    &o.Status,
			Timestamp: &o.Timestamp,
		}
	}

	sp := &SparseValidateRQLResp{}
	for _, f := range fields {
		switch f {
		case "error":
			sp.Error = &(o.Error)
		case "status":
			sp.Status = &(o.Status)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseValidateRQLResp to the object.
func (o *ValidateRQLResp) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseValidateRQLResp)
	if so.Error != nil {
		o.Error = *so.Error
	}
	if so.Status != nil {
		o.Status = *so.Status
	}
	if so.Timestamp != nil {
		o.Timestamp = *so.Timestamp
	}
}

// DeepCopy returns a deep copy if the ValidateRQLResp.
func (o *ValidateRQLResp) DeepCopy() *ValidateRQLResp {

	if o == nil {
		return nil
	}

	out := &ValidateRQLResp{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ValidateRQLResp.
func (o *ValidateRQLResp) DeepCopyInto(out *ValidateRQLResp) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ValidateRQLResp: %s", err))
	}

	*out = *target.(*ValidateRQLResp)
}

// Validate valides the current information stored into the structure.
func (o *ValidateRQLResp) Validate() error {

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
func (*ValidateRQLResp) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ValidateRQLRespAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ValidateRQLRespLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ValidateRQLResp) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ValidateRQLRespAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ValidateRQLResp) ValueForAttribute(name string) interface{} {

	switch name {
	case "error":
		return o.Error
	case "status":
		return o.Status
	case "timestamp":
		return o.Timestamp
	}

	return nil
}

// ValidateRQLRespAttributesMap represents the map of attribute for ValidateRQLResp.
var ValidateRQLRespAttributesMap = map[string]elemental.AttributeSpecification{
	"Error": {
		AllowedChoices: []string{},
		ConvertedName:  "Error",
		Description:    `TODO.`,
		Exposed:        true,
		Name:           "error",
		Type:           "string",
	},
	"Status": {
		AllowedChoices: []string{},
		ConvertedName:  "Status",
		Description:    `TODO.`,
		Exposed:        true,
		Name:           "status",
		Type:           "integer",
	},
	"Timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `TODO.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "integer",
	},
}

// ValidateRQLRespLowerCaseAttributesMap represents the map of attribute for ValidateRQLResp.
var ValidateRQLRespLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"error": {
		AllowedChoices: []string{},
		ConvertedName:  "Error",
		Description:    `TODO.`,
		Exposed:        true,
		Name:           "error",
		Type:           "string",
	},
	"status": {
		AllowedChoices: []string{},
		ConvertedName:  "Status",
		Description:    `TODO.`,
		Exposed:        true,
		Name:           "status",
		Type:           "integer",
	},
	"timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `TODO.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "integer",
	},
}

// SparseValidateRQLRespsList represents a list of SparseValidateRQLResps
type SparseValidateRQLRespsList []*SparseValidateRQLResp

// Identity returns the identity of the objects in the list.
func (o SparseValidateRQLRespsList) Identity() elemental.Identity {

	return ValidateRQLRespIdentity
}

// Copy returns a pointer to a copy the SparseValidateRQLRespsList.
func (o SparseValidateRQLRespsList) Copy() elemental.Identifiables {

	copy := append(SparseValidateRQLRespsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseValidateRQLRespsList.
func (o SparseValidateRQLRespsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseValidateRQLRespsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseValidateRQLResp))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseValidateRQLRespsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseValidateRQLRespsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseValidateRQLRespsList converted to ValidateRQLRespsList.
func (o SparseValidateRQLRespsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseValidateRQLRespsList) Version() int {

	return 1
}

// SparseValidateRQLResp represents the sparse version of a validaterqlresp.
type SparseValidateRQLResp struct {
	// TODO.
	Error *string `json:"error,omitempty" msgpack:"error,omitempty" bson:"-" mapstructure:"error,omitempty"`

	// TODO.
	Status *int `json:"status,omitempty" msgpack:"status,omitempty" bson:"-" mapstructure:"status,omitempty"`

	// TODO.
	Timestamp *int `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseValidateRQLResp returns a new  SparseValidateRQLResp.
func NewSparseValidateRQLResp() *SparseValidateRQLResp {
	return &SparseValidateRQLResp{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseValidateRQLResp) Identity() elemental.Identity {

	return ValidateRQLRespIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseValidateRQLResp) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseValidateRQLResp) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseValidateRQLResp) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseValidateRQLResp{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseValidateRQLResp) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseValidateRQLResp{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseValidateRQLResp) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseValidateRQLResp) ToPlain() elemental.PlainIdentifiable {

	out := NewValidateRQLResp()
	if o.Error != nil {
		out.Error = *o.Error
	}
	if o.Status != nil {
		out.Status = *o.Status
	}
	if o.Timestamp != nil {
		out.Timestamp = *o.Timestamp
	}

	return out
}

// DeepCopy returns a deep copy if the SparseValidateRQLResp.
func (o *SparseValidateRQLResp) DeepCopy() *SparseValidateRQLResp {

	if o == nil {
		return nil
	}

	out := &SparseValidateRQLResp{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseValidateRQLResp.
func (o *SparseValidateRQLResp) DeepCopyInto(out *SparseValidateRQLResp) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseValidateRQLResp: %s", err))
	}

	*out = *target.(*SparseValidateRQLResp)
}

type mongoAttributesValidateRQLResp struct {
}
type mongoAttributesSparseValidateRQLResp struct {
}
