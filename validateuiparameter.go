package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ValidateUIParameterIdentity represents the Identity of the object.
var ValidateUIParameterIdentity = elemental.Identity{
	Name:     "validateuiparameter",
	Category: "validateuiparameters",
	Package:  "ignis",
	Private:  false,
}

// ValidateUIParametersList represents a list of ValidateUIParameters
type ValidateUIParametersList []*ValidateUIParameter

// Identity returns the identity of the objects in the list.
func (o ValidateUIParametersList) Identity() elemental.Identity {

	return ValidateUIParameterIdentity
}

// Copy returns a pointer to a copy the ValidateUIParametersList.
func (o ValidateUIParametersList) Copy() elemental.Identifiables {

	copy := append(ValidateUIParametersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ValidateUIParametersList.
func (o ValidateUIParametersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ValidateUIParametersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ValidateUIParameter))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ValidateUIParametersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ValidateUIParametersList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ValidateUIParametersList converted to SparseValidateUIParametersList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ValidateUIParametersList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseValidateUIParametersList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseValidateUIParameter)
	}

	return out
}

// Version returns the version of the content.
func (o ValidateUIParametersList) Version() int {

	return 1
}

// ValidateUIParameter represents the model of a validateuiparameter
type ValidateUIParameter struct {
	// Contains the list of errors.
	Errors map[string]string `json:"errors" msgpack:"errors" bson:"-" mapstructure:"errors,omitempty"`

	// List of parameters to validate.
	Parameters []*UIParameter `json:"parameters" msgpack:"parameters" bson:"parameters" mapstructure:"parameters,omitempty"`

	// Contains the computed values.
	Values map[string]interface{} `json:"values" msgpack:"values" bson:"-" mapstructure:"values,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewValidateUIParameter returns a new *ValidateUIParameter
func NewValidateUIParameter() *ValidateUIParameter {

	return &ValidateUIParameter{
		ModelVersion: 1,
		Errors:       map[string]string{},
		Parameters:   []*UIParameter{},
		Values:       map[string]interface{}{},
	}
}

// Identity returns the Identity of the object.
func (o *ValidateUIParameter) Identity() elemental.Identity {

	return ValidateUIParameterIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ValidateUIParameter) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ValidateUIParameter) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ValidateUIParameter) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesValidateUIParameter{}

	s.Parameters = o.Parameters

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ValidateUIParameter) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesValidateUIParameter{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Parameters = s.Parameters

	return nil
}

// Version returns the hardcoded version of the model.
func (o *ValidateUIParameter) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *ValidateUIParameter) BleveType() string {

	return "validateuiparameter"
}

// DefaultOrder returns the list of default ordering fields.
func (o *ValidateUIParameter) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *ValidateUIParameter) Doc() string {

	return `Validates a list of [UIParameter](#uiparameter) parameters.`
}

func (o *ValidateUIParameter) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ValidateUIParameter) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseValidateUIParameter{
			Errors:     &o.Errors,
			Parameters: &o.Parameters,
			Values:     &o.Values,
		}
	}

	sp := &SparseValidateUIParameter{}
	for _, f := range fields {
		switch f {
		case "errors":
			sp.Errors = &(o.Errors)
		case "parameters":
			sp.Parameters = &(o.Parameters)
		case "values":
			sp.Values = &(o.Values)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseValidateUIParameter to the object.
func (o *ValidateUIParameter) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseValidateUIParameter)
	if so.Errors != nil {
		o.Errors = *so.Errors
	}
	if so.Parameters != nil {
		o.Parameters = *so.Parameters
	}
	if so.Values != nil {
		o.Values = *so.Values
	}
}

// DeepCopy returns a deep copy if the ValidateUIParameter.
func (o *ValidateUIParameter) DeepCopy() *ValidateUIParameter {

	if o == nil {
		return nil
	}

	out := &ValidateUIParameter{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ValidateUIParameter.
func (o *ValidateUIParameter) DeepCopyInto(out *ValidateUIParameter) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ValidateUIParameter: %s", err))
	}

	*out = *target.(*ValidateUIParameter)
}

// Validate valides the current information stored into the structure.
func (o *ValidateUIParameter) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.Parameters {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
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
func (*ValidateUIParameter) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ValidateUIParameterAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ValidateUIParameterLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ValidateUIParameter) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ValidateUIParameterAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ValidateUIParameter) ValueForAttribute(name string) interface{} {

	switch name {
	case "errors":
		return o.Errors
	case "parameters":
		return o.Parameters
	case "values":
		return o.Values
	}

	return nil
}

// ValidateUIParameterAttributesMap represents the map of attribute for ValidateUIParameter.
var ValidateUIParameterAttributesMap = map[string]elemental.AttributeSpecification{
	"Errors": {
		AllowedChoices: []string{},
		ConvertedName:  "Errors",
		Description:    `Contains the list of errors.`,
		Exposed:        true,
		Name:           "errors",
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Parameters": {
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `List of parameters to validate.`,
		Exposed:        true,
		Name:           "parameters",
		Stored:         true,
		SubType:        "uiparameter",
		Type:           "refList",
	},
	"Values": {
		AllowedChoices: []string{},
		ConvertedName:  "Values",
		Description:    `Contains the computed values.`,
		Exposed:        true,
		Name:           "values",
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
}

// ValidateUIParameterLowerCaseAttributesMap represents the map of attribute for ValidateUIParameter.
var ValidateUIParameterLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"errors": {
		AllowedChoices: []string{},
		ConvertedName:  "Errors",
		Description:    `Contains the list of errors.`,
		Exposed:        true,
		Name:           "errors",
		SubType:        "map[string]string",
		Type:           "external",
	},
	"parameters": {
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `List of parameters to validate.`,
		Exposed:        true,
		Name:           "parameters",
		Stored:         true,
		SubType:        "uiparameter",
		Type:           "refList",
	},
	"values": {
		AllowedChoices: []string{},
		ConvertedName:  "Values",
		Description:    `Contains the computed values.`,
		Exposed:        true,
		Name:           "values",
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
}

// SparseValidateUIParametersList represents a list of SparseValidateUIParameters
type SparseValidateUIParametersList []*SparseValidateUIParameter

// Identity returns the identity of the objects in the list.
func (o SparseValidateUIParametersList) Identity() elemental.Identity {

	return ValidateUIParameterIdentity
}

// Copy returns a pointer to a copy the SparseValidateUIParametersList.
func (o SparseValidateUIParametersList) Copy() elemental.Identifiables {

	copy := append(SparseValidateUIParametersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseValidateUIParametersList.
func (o SparseValidateUIParametersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseValidateUIParametersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseValidateUIParameter))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseValidateUIParametersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseValidateUIParametersList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseValidateUIParametersList converted to ValidateUIParametersList.
func (o SparseValidateUIParametersList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseValidateUIParametersList) Version() int {

	return 1
}

// SparseValidateUIParameter represents the sparse version of a validateuiparameter.
type SparseValidateUIParameter struct {
	// Contains the list of errors.
	Errors *map[string]string `json:"errors,omitempty" msgpack:"errors,omitempty" bson:"-" mapstructure:"errors,omitempty"`

	// List of parameters to validate.
	Parameters *[]*UIParameter `json:"parameters,omitempty" msgpack:"parameters,omitempty" bson:"parameters,omitempty" mapstructure:"parameters,omitempty"`

	// Contains the computed values.
	Values *map[string]interface{} `json:"values,omitempty" msgpack:"values,omitempty" bson:"-" mapstructure:"values,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseValidateUIParameter returns a new  SparseValidateUIParameter.
func NewSparseValidateUIParameter() *SparseValidateUIParameter {
	return &SparseValidateUIParameter{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseValidateUIParameter) Identity() elemental.Identity {

	return ValidateUIParameterIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseValidateUIParameter) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseValidateUIParameter) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseValidateUIParameter) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseValidateUIParameter{}

	if o.Parameters != nil {
		s.Parameters = o.Parameters
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseValidateUIParameter) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseValidateUIParameter{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Parameters != nil {
		o.Parameters = s.Parameters
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseValidateUIParameter) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseValidateUIParameter) ToPlain() elemental.PlainIdentifiable {

	out := NewValidateUIParameter()
	if o.Errors != nil {
		out.Errors = *o.Errors
	}
	if o.Parameters != nil {
		out.Parameters = *o.Parameters
	}
	if o.Values != nil {
		out.Values = *o.Values
	}

	return out
}

// DeepCopy returns a deep copy if the SparseValidateUIParameter.
func (o *SparseValidateUIParameter) DeepCopy() *SparseValidateUIParameter {

	if o == nil {
		return nil
	}

	out := &SparseValidateUIParameter{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseValidateUIParameter.
func (o *SparseValidateUIParameter) DeepCopyInto(out *SparseValidateUIParameter) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseValidateUIParameter: %s", err))
	}

	*out = *target.(*SparseValidateUIParameter)
}

type mongoAttributesValidateUIParameter struct {
	Parameters []*UIParameter `bson:"parameters"`
}
type mongoAttributesSparseValidateUIParameter struct {
	Parameters *[]*UIParameter `bson:"parameters,omitempty"`
}
