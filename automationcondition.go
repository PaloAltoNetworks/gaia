package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AutomationConditionIdentity represents the Identity of the object.
var AutomationConditionIdentity = elemental.Identity{
	Name:     "automationcondition",
	Category: "automationconditions",
	Package:  "sephiroth",
	Private:  false,
}

// AutomationConditionsList represents a list of AutomationConditions
type AutomationConditionsList []*AutomationCondition

// Identity returns the identity of the objects in the list.
func (o AutomationConditionsList) Identity() elemental.Identity {

	return AutomationConditionIdentity
}

// Copy returns a pointer to a copy the AutomationConditionsList.
func (o AutomationConditionsList) Copy() elemental.Identifiables {

	copy := append(AutomationConditionsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AutomationConditionsList.
func (o AutomationConditionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AutomationConditionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AutomationCondition))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AutomationConditionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AutomationConditionsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the AutomationConditionsList converted to SparseAutomationConditionsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AutomationConditionsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAutomationConditionsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAutomationCondition)
	}

	return out
}

// Version returns the version of the content.
func (o AutomationConditionsList) Version() int {

	return 1
}

// AutomationCondition represents the model of a automationcondition
type AutomationCondition struct {
	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Contains the entitlements needed for executing the function.
	Entitlements map[string][]elemental.Operation `json:"entitlements" msgpack:"entitlements" bson:"-" mapstructure:"entitlements,omitempty"`

	// Function contains the code.
	Function string `json:"function" msgpack:"function" bson:"-" mapstructure:"function,omitempty"`

	// Contains the unique identifier key for the condition.
	Key string `json:"key" msgpack:"key" bson:"-" mapstructure:"key,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Contains the computed parameters.
	Parameters map[string]interface{} `json:"parameters" msgpack:"parameters" bson:"-" mapstructure:"parameters,omitempty"`

	// Contains all the steps with parameters.
	Steps []*UIStep `json:"steps" msgpack:"steps" bson:"-" mapstructure:"steps,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAutomationCondition returns a new *AutomationCondition
func NewAutomationCondition() *AutomationCondition {

	return &AutomationCondition{
		ModelVersion: 1,
		Entitlements: map[string][]elemental.Operation{},
		Parameters:   map[string]interface{}{},
		Steps:        []*UIStep{},
	}
}

// Identity returns the Identity of the object.
func (o *AutomationCondition) Identity() elemental.Identity {

	return AutomationConditionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AutomationCondition) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AutomationCondition) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AutomationCondition) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesAutomationCondition{}

	s.Description = o.Description
	s.Name = o.Name

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AutomationCondition) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesAutomationCondition{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Description = s.Description
	o.Name = s.Name

	return nil
}

// Version returns the hardcoded version of the model.
func (o *AutomationCondition) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *AutomationCondition) BleveType() string {

	return "automationcondition"
}

// DefaultOrder returns the list of default ordering fields.
func (o *AutomationCondition) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *AutomationCondition) Doc() string {

	return `Condition that can be used in automations.`
}

func (o *AutomationCondition) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetDescription returns the Description of the receiver.
func (o *AutomationCondition) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *AutomationCondition) SetDescription(description string) {

	o.Description = description
}

// GetName returns the Name of the receiver.
func (o *AutomationCondition) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *AutomationCondition) SetName(name string) {

	o.Name = name
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *AutomationCondition) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAutomationCondition{
			Description:  &o.Description,
			Entitlements: &o.Entitlements,
			Function:     &o.Function,
			Key:          &o.Key,
			Name:         &o.Name,
			Parameters:   &o.Parameters,
			Steps:        &o.Steps,
		}
	}

	sp := &SparseAutomationCondition{}
	for _, f := range fields {
		switch f {
		case "description":
			sp.Description = &(o.Description)
		case "entitlements":
			sp.Entitlements = &(o.Entitlements)
		case "function":
			sp.Function = &(o.Function)
		case "key":
			sp.Key = &(o.Key)
		case "name":
			sp.Name = &(o.Name)
		case "parameters":
			sp.Parameters = &(o.Parameters)
		case "steps":
			sp.Steps = &(o.Steps)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseAutomationCondition to the object.
func (o *AutomationCondition) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAutomationCondition)
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Entitlements != nil {
		o.Entitlements = *so.Entitlements
	}
	if so.Function != nil {
		o.Function = *so.Function
	}
	if so.Key != nil {
		o.Key = *so.Key
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Parameters != nil {
		o.Parameters = *so.Parameters
	}
	if so.Steps != nil {
		o.Steps = *so.Steps
	}
}

// DeepCopy returns a deep copy if the AutomationCondition.
func (o *AutomationCondition) DeepCopy() *AutomationCondition {

	if o == nil {
		return nil
	}

	out := &AutomationCondition{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AutomationCondition.
func (o *AutomationCondition) DeepCopyInto(out *AutomationCondition) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AutomationCondition: %s", err))
	}

	*out = *target.(*AutomationCondition)
}

// Validate valides the current information stored into the structure.
func (o *AutomationCondition) Validate() error {

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

	for _, sub := range o.Steps {
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
func (*AutomationCondition) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AutomationConditionAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AutomationConditionLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AutomationCondition) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AutomationConditionAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *AutomationCondition) ValueForAttribute(name string) interface{} {

	switch name {
	case "description":
		return o.Description
	case "entitlements":
		return o.Entitlements
	case "function":
		return o.Function
	case "key":
		return o.Key
	case "name":
		return o.Name
	case "parameters":
		return o.Parameters
	case "steps":
		return o.Steps
	}

	return nil
}

// AutomationConditionAttributesMap represents the map of attribute for AutomationCondition.
var AutomationConditionAttributesMap = map[string]elemental.AttributeSpecification{
	"Description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
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
	"Entitlements": {
		AllowedChoices: []string{},
		ConvertedName:  "Entitlements",
		Description:    `Contains the entitlements needed for executing the function.`,
		Exposed:        true,
		Name:           "entitlements",
		SubType:        "_automation_entitlements",
		Type:           "external",
	},
	"Function": {
		AllowedChoices: []string{},
		ConvertedName:  "Function",
		Description:    `Function contains the code.`,
		Exposed:        true,
		Name:           "function",
		Type:           "string",
	},
	"Key": {
		AllowedChoices: []string{},
		ConvertedName:  "Key",
		Description:    `Contains the unique identifier key for the condition.`,
		Exposed:        true,
		Name:           "key",
		Type:           "string",
	},
	"Name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
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
	"Parameters": {
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Contains the computed parameters.`,
		Exposed:        true,
		Name:           "parameters",
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"Steps": {
		AllowedChoices: []string{},
		ConvertedName:  "Steps",
		Description:    `Contains all the steps with parameters.`,
		Exposed:        true,
		Name:           "steps",
		SubType:        "uistep",
		Type:           "refList",
	},
}

// AutomationConditionLowerCaseAttributesMap represents the map of attribute for AutomationCondition.
var AutomationConditionLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
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
	"entitlements": {
		AllowedChoices: []string{},
		ConvertedName:  "Entitlements",
		Description:    `Contains the entitlements needed for executing the function.`,
		Exposed:        true,
		Name:           "entitlements",
		SubType:        "_automation_entitlements",
		Type:           "external",
	},
	"function": {
		AllowedChoices: []string{},
		ConvertedName:  "Function",
		Description:    `Function contains the code.`,
		Exposed:        true,
		Name:           "function",
		Type:           "string",
	},
	"key": {
		AllowedChoices: []string{},
		ConvertedName:  "Key",
		Description:    `Contains the unique identifier key for the condition.`,
		Exposed:        true,
		Name:           "key",
		Type:           "string",
	},
	"name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
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
	"parameters": {
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Contains the computed parameters.`,
		Exposed:        true,
		Name:           "parameters",
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"steps": {
		AllowedChoices: []string{},
		ConvertedName:  "Steps",
		Description:    `Contains all the steps with parameters.`,
		Exposed:        true,
		Name:           "steps",
		SubType:        "uistep",
		Type:           "refList",
	},
}

// SparseAutomationConditionsList represents a list of SparseAutomationConditions
type SparseAutomationConditionsList []*SparseAutomationCondition

// Identity returns the identity of the objects in the list.
func (o SparseAutomationConditionsList) Identity() elemental.Identity {

	return AutomationConditionIdentity
}

// Copy returns a pointer to a copy the SparseAutomationConditionsList.
func (o SparseAutomationConditionsList) Copy() elemental.Identifiables {

	copy := append(SparseAutomationConditionsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAutomationConditionsList.
func (o SparseAutomationConditionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAutomationConditionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAutomationCondition))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAutomationConditionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAutomationConditionsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseAutomationConditionsList converted to AutomationConditionsList.
func (o SparseAutomationConditionsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAutomationConditionsList) Version() int {

	return 1
}

// SparseAutomationCondition represents the sparse version of a automationcondition.
type SparseAutomationCondition struct {
	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Contains the entitlements needed for executing the function.
	Entitlements *map[string][]elemental.Operation `json:"entitlements,omitempty" msgpack:"entitlements,omitempty" bson:"-" mapstructure:"entitlements,omitempty"`

	// Function contains the code.
	Function *string `json:"function,omitempty" msgpack:"function,omitempty" bson:"-" mapstructure:"function,omitempty"`

	// Contains the unique identifier key for the condition.
	Key *string `json:"key,omitempty" msgpack:"key,omitempty" bson:"-" mapstructure:"key,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Contains the computed parameters.
	Parameters *map[string]interface{} `json:"parameters,omitempty" msgpack:"parameters,omitempty" bson:"-" mapstructure:"parameters,omitempty"`

	// Contains all the steps with parameters.
	Steps *[]*UIStep `json:"steps,omitempty" msgpack:"steps,omitempty" bson:"-" mapstructure:"steps,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAutomationCondition returns a new  SparseAutomationCondition.
func NewSparseAutomationCondition() *SparseAutomationCondition {
	return &SparseAutomationCondition{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAutomationCondition) Identity() elemental.Identity {

	return AutomationConditionIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAutomationCondition) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAutomationCondition) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAutomationCondition) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseAutomationCondition{}

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
func (o *SparseAutomationCondition) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseAutomationCondition{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Name != nil {
		o.Name = s.Name
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseAutomationCondition) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAutomationCondition) ToPlain() elemental.PlainIdentifiable {

	out := NewAutomationCondition()
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Entitlements != nil {
		out.Entitlements = *o.Entitlements
	}
	if o.Function != nil {
		out.Function = *o.Function
	}
	if o.Key != nil {
		out.Key = *o.Key
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Parameters != nil {
		out.Parameters = *o.Parameters
	}
	if o.Steps != nil {
		out.Steps = *o.Steps
	}

	return out
}

// GetDescription returns the Description of the receiver.
func (o *SparseAutomationCondition) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseAutomationCondition) SetDescription(description string) {

	o.Description = &description
}

// GetName returns the Name of the receiver.
func (o *SparseAutomationCondition) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseAutomationCondition) SetName(name string) {

	o.Name = &name
}

// DeepCopy returns a deep copy if the SparseAutomationCondition.
func (o *SparseAutomationCondition) DeepCopy() *SparseAutomationCondition {

	if o == nil {
		return nil
	}

	out := &SparseAutomationCondition{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAutomationCondition.
func (o *SparseAutomationCondition) DeepCopyInto(out *SparseAutomationCondition) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAutomationCondition: %s", err))
	}

	*out = *target.(*SparseAutomationCondition)
}

type mongoAttributesAutomationCondition struct {
	Description string `bson:"description"`
	Name        string `bson:"name"`
}
type mongoAttributesSparseAutomationCondition struct {
	Description *string `bson:"description,omitempty"`
	Name        *string `bson:"name,omitempty"`
}
