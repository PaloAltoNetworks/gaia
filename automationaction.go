package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AutomationActionIdentity represents the Identity of the object.
var AutomationActionIdentity = elemental.Identity{
	Name:     "automationaction",
	Category: "automationactions",
	Package:  "sephiroth",
	Private:  false,
}

// AutomationActionsList represents a list of AutomationActions
type AutomationActionsList []*AutomationAction

// Identity returns the identity of the objects in the list.
func (o AutomationActionsList) Identity() elemental.Identity {

	return AutomationActionIdentity
}

// Copy returns a pointer to a copy the AutomationActionsList.
func (o AutomationActionsList) Copy() elemental.Identifiables {

	copy := append(AutomationActionsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AutomationActionsList.
func (o AutomationActionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AutomationActionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AutomationAction))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AutomationActionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AutomationActionsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the AutomationActionsList converted to SparseAutomationActionsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AutomationActionsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAutomationActionsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAutomationAction)
	}

	return out
}

// Version returns the version of the content.
func (o AutomationActionsList) Version() int {

	return 1
}

// AutomationAction represents the model of a automationaction
type AutomationAction struct {
	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Contains the entitlements needed for executing the function.
	Entitlements map[string][]elemental.Operation `json:"entitlements" msgpack:"entitlements" bson:"-" mapstructure:"entitlements,omitempty"`

	// Function contains the code.
	Function string `json:"function" msgpack:"function" bson:"-" mapstructure:"function,omitempty"`

	// Contains the unique identifier key for the action.
	Key string `json:"key" msgpack:"key" bson:"-" mapstructure:"key,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Contains the computed parameters.
	Parameters map[string]interface{} `json:"parameters" msgpack:"parameters" bson:"-" mapstructure:"parameters,omitempty"`

	// Contains all the steps with parameters.
	Steps []*UIStep `json:"steps" msgpack:"steps" bson:"-" mapstructure:"steps,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAutomationAction returns a new *AutomationAction
func NewAutomationAction() *AutomationAction {

	return &AutomationAction{
		ModelVersion: 1,
		Entitlements: map[string][]elemental.Operation{},
		Parameters:   map[string]interface{}{},
		Steps:        []*UIStep{},
	}
}

// Identity returns the Identity of the object.
func (o *AutomationAction) Identity() elemental.Identity {

	return AutomationActionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AutomationAction) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AutomationAction) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AutomationAction) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesAutomationAction{}

	s.Description = o.Description
	s.Name = o.Name

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AutomationAction) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesAutomationAction{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Description = s.Description
	o.Name = s.Name

	return nil
}

// Version returns the hardcoded version of the model.
func (o *AutomationAction) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *AutomationAction) BleveType() string {

	return "automationaction"
}

// DefaultOrder returns the list of default ordering fields.
func (o *AutomationAction) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *AutomationAction) Doc() string {

	return `Action that can be used in automations.`
}

func (o *AutomationAction) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetDescription returns the Description of the receiver.
func (o *AutomationAction) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *AutomationAction) SetDescription(description string) {

	o.Description = description
}

// GetName returns the Name of the receiver.
func (o *AutomationAction) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *AutomationAction) SetName(name string) {

	o.Name = name
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *AutomationAction) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAutomationAction{
			Description:  &o.Description,
			Entitlements: &o.Entitlements,
			Function:     &o.Function,
			Key:          &o.Key,
			Name:         &o.Name,
			Parameters:   &o.Parameters,
			Steps:        &o.Steps,
		}
	}

	sp := &SparseAutomationAction{}
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

// Patch apply the non nil value of a *SparseAutomationAction to the object.
func (o *AutomationAction) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAutomationAction)
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

// DeepCopy returns a deep copy if the AutomationAction.
func (o *AutomationAction) DeepCopy() *AutomationAction {

	if o == nil {
		return nil
	}

	out := &AutomationAction{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AutomationAction.
func (o *AutomationAction) DeepCopyInto(out *AutomationAction) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AutomationAction: %s", err))
	}

	*out = *target.(*AutomationAction)
}

// Validate valides the current information stored into the structure.
func (o *AutomationAction) Validate() error {

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
func (*AutomationAction) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AutomationActionAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AutomationActionLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AutomationAction) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AutomationActionAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *AutomationAction) ValueForAttribute(name string) interface{} {

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

// AutomationActionAttributesMap represents the map of attribute for AutomationAction.
var AutomationActionAttributesMap = map[string]elemental.AttributeSpecification{
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
		Description:    `Contains the unique identifier key for the action.`,
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

// AutomationActionLowerCaseAttributesMap represents the map of attribute for AutomationAction.
var AutomationActionLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		Description:    `Contains the unique identifier key for the action.`,
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

// SparseAutomationActionsList represents a list of SparseAutomationActions
type SparseAutomationActionsList []*SparseAutomationAction

// Identity returns the identity of the objects in the list.
func (o SparseAutomationActionsList) Identity() elemental.Identity {

	return AutomationActionIdentity
}

// Copy returns a pointer to a copy the SparseAutomationActionsList.
func (o SparseAutomationActionsList) Copy() elemental.Identifiables {

	copy := append(SparseAutomationActionsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAutomationActionsList.
func (o SparseAutomationActionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAutomationActionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAutomationAction))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAutomationActionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAutomationActionsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseAutomationActionsList converted to AutomationActionsList.
func (o SparseAutomationActionsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAutomationActionsList) Version() int {

	return 1
}

// SparseAutomationAction represents the sparse version of a automationaction.
type SparseAutomationAction struct {
	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Contains the entitlements needed for executing the function.
	Entitlements *map[string][]elemental.Operation `json:"entitlements,omitempty" msgpack:"entitlements,omitempty" bson:"-" mapstructure:"entitlements,omitempty"`

	// Function contains the code.
	Function *string `json:"function,omitempty" msgpack:"function,omitempty" bson:"-" mapstructure:"function,omitempty"`

	// Contains the unique identifier key for the action.
	Key *string `json:"key,omitempty" msgpack:"key,omitempty" bson:"-" mapstructure:"key,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Contains the computed parameters.
	Parameters *map[string]interface{} `json:"parameters,omitempty" msgpack:"parameters,omitempty" bson:"-" mapstructure:"parameters,omitempty"`

	// Contains all the steps with parameters.
	Steps *[]*UIStep `json:"steps,omitempty" msgpack:"steps,omitempty" bson:"-" mapstructure:"steps,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAutomationAction returns a new  SparseAutomationAction.
func NewSparseAutomationAction() *SparseAutomationAction {
	return &SparseAutomationAction{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAutomationAction) Identity() elemental.Identity {

	return AutomationActionIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAutomationAction) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAutomationAction) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAutomationAction) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseAutomationAction{}

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
func (o *SparseAutomationAction) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseAutomationAction{}
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
func (o *SparseAutomationAction) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAutomationAction) ToPlain() elemental.PlainIdentifiable {

	out := NewAutomationAction()
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
func (o *SparseAutomationAction) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseAutomationAction) SetDescription(description string) {

	o.Description = &description
}

// GetName returns the Name of the receiver.
func (o *SparseAutomationAction) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseAutomationAction) SetName(name string) {

	o.Name = &name
}

// DeepCopy returns a deep copy if the SparseAutomationAction.
func (o *SparseAutomationAction) DeepCopy() *SparseAutomationAction {

	if o == nil {
		return nil
	}

	out := &SparseAutomationAction{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAutomationAction.
func (o *SparseAutomationAction) DeepCopyInto(out *SparseAutomationAction) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAutomationAction: %s", err))
	}

	*out = *target.(*SparseAutomationAction)
}

type mongoAttributesAutomationAction struct {
	Description string `bson:"description"`
	Name        string `bson:"name"`
}
type mongoAttributesSparseAutomationAction struct {
	Description *string `bson:"description,omitempty"`
	Name        *string `bson:"name,omitempty"`
}
