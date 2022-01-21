package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PlanIdentity represents the Identity of the object.
var PlanIdentity = elemental.Identity{
	Name:     "plan",
	Category: "plans",
	Package:  "vince",
	Private:  false,
}

// PlansList represents a list of Plans
type PlansList []*Plan

// Identity returns the identity of the objects in the list.
func (o PlansList) Identity() elemental.Identity {

	return PlanIdentity
}

// Copy returns a pointer to a copy the PlansList.
func (o PlansList) Copy() elemental.Identifiables {

	copy := append(PlansList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PlansList.
func (o PlansList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PlansList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Plan))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PlansList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PlansList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PlansList converted to SparsePlansList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PlansList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePlansList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePlan)
	}

	return out
}

// Version returns the version of the content.
func (o PlansList) Version() int {

	return 1
}

// Plan represents the model of a plan
type Plan struct {
	// Contains the description of the plan.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Contains the key identifier of the plan.
	Key string `json:"key" msgpack:"key" bson:"key" mapstructure:"key,omitempty"`

	// Contains the name of the plan.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Defines the quota for each identity.
	Quotas map[string]int `json:"-" msgpack:"-" bson:"-" mapstructure:"-,omitempty"`

	// RequireAdminValidation indicates if the plan requires an admin approval.
	RequireAdminValidation bool `json:"-" msgpack:"-" bson:"requireadminvalidation" mapstructure:"-,omitempty"`

	// Roles defines the roles given to the account.
	Roles []string `json:"-" msgpack:"-" bson:"-" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPlan returns a new *Plan
func NewPlan() *Plan {

	return &Plan{
		ModelVersion: 1,
		Quotas:       map[string]int{},
		Roles:        []string{},
	}
}

// Identity returns the Identity of the object.
func (o *Plan) Identity() elemental.Identity {

	return PlanIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Plan) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Plan) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Plan) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPlan{}

	s.Description = o.Description
	s.Key = o.Key
	s.Name = o.Name
	s.RequireAdminValidation = o.RequireAdminValidation

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Plan) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPlan{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Description = s.Description
	o.Key = s.Key
	o.Name = s.Name
	o.RequireAdminValidation = s.RequireAdminValidation

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Plan) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Plan) BleveType() string {

	return "plan"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Plan) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Plan) Doc() string {

	return `Contains the various billing plans available.`
}

func (o *Plan) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Plan) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePlan{
			Description:            &o.Description,
			Key:                    &o.Key,
			Name:                   &o.Name,
			Quotas:                 &o.Quotas,
			RequireAdminValidation: &o.RequireAdminValidation,
			Roles:                  &o.Roles,
		}
	}

	sp := &SparsePlan{}
	for _, f := range fields {
		switch f {
		case "description":
			sp.Description = &(o.Description)
		case "key":
			sp.Key = &(o.Key)
		case "name":
			sp.Name = &(o.Name)
		case "quotas":
			sp.Quotas = &(o.Quotas)
		case "requireAdminValidation":
			sp.RequireAdminValidation = &(o.RequireAdminValidation)
		case "roles":
			sp.Roles = &(o.Roles)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePlan to the object.
func (o *Plan) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePlan)
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Key != nil {
		o.Key = *so.Key
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Quotas != nil {
		o.Quotas = *so.Quotas
	}
	if so.RequireAdminValidation != nil {
		o.RequireAdminValidation = *so.RequireAdminValidation
	}
	if so.Roles != nil {
		o.Roles = *so.Roles
	}
}

// DeepCopy returns a deep copy if the Plan.
func (o *Plan) DeepCopy() *Plan {

	if o == nil {
		return nil
	}

	out := &Plan{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Plan.
func (o *Plan) DeepCopyInto(out *Plan) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Plan: %s", err))
	}

	*out = *target.(*Plan)
}

// Validate valides the current information stored into the structure.
func (o *Plan) Validate() error {

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
func (*Plan) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PlanAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PlanLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Plan) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PlanAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Plan) ValueForAttribute(name string) interface{} {

	switch name {
	case "description":
		return o.Description
	case "key":
		return o.Key
	case "name":
		return o.Name
	case "quotas":
		return o.Quotas
	case "requireAdminValidation":
		return o.RequireAdminValidation
	case "roles":
		return o.Roles
	}

	return nil
}

// PlanAttributesMap represents the map of attribute for Plan.
var PlanAttributesMap = map[string]elemental.AttributeSpecification{
	"Description": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Description",
		Description:    `Contains the description of the plan.`,
		Exposed:        true,
		Name:           "description",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Key": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Key",
		Description:    `Contains the key identifier of the plan.`,
		Exposed:        true,
		Name:           "key",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Name": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Name",
		Description:    `Contains the name of the plan.`,
		Exposed:        true,
		Name:           "name",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Quotas": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Quotas",
		Description:    `Defines the quota for each identity.`,
		Name:           "quotas",
		ReadOnly:       true,
		SubType:        "map[string]int",
		Type:           "external",
	},
	"RequireAdminValidation": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "RequireAdminValidation",
		Description:    `RequireAdminValidation indicates if the plan requires an admin approval.`,
		Name:           "requireAdminValidation",
		ReadOnly:       true,
		Stored:         true,
		Type:           "boolean",
	},
	"Roles": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Roles",
		Description:    `Roles defines the roles given to the account.`,
		Name:           "roles",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
}

// PlanLowerCaseAttributesMap represents the map of attribute for Plan.
var PlanLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"description": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Description",
		Description:    `Contains the description of the plan.`,
		Exposed:        true,
		Name:           "description",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"key": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Key",
		Description:    `Contains the key identifier of the plan.`,
		Exposed:        true,
		Name:           "key",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"name": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Name",
		Description:    `Contains the name of the plan.`,
		Exposed:        true,
		Name:           "name",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"quotas": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Quotas",
		Description:    `Defines the quota for each identity.`,
		Name:           "quotas",
		ReadOnly:       true,
		SubType:        "map[string]int",
		Type:           "external",
	},
	"requireadminvalidation": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "RequireAdminValidation",
		Description:    `RequireAdminValidation indicates if the plan requires an admin approval.`,
		Name:           "requireAdminValidation",
		ReadOnly:       true,
		Stored:         true,
		Type:           "boolean",
	},
	"roles": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Roles",
		Description:    `Roles defines the roles given to the account.`,
		Name:           "roles",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
}

// SparsePlansList represents a list of SparsePlans
type SparsePlansList []*SparsePlan

// Identity returns the identity of the objects in the list.
func (o SparsePlansList) Identity() elemental.Identity {

	return PlanIdentity
}

// Copy returns a pointer to a copy the SparsePlansList.
func (o SparsePlansList) Copy() elemental.Identifiables {

	copy := append(SparsePlansList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePlansList.
func (o SparsePlansList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePlansList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePlan))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePlansList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePlansList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePlansList converted to PlansList.
func (o SparsePlansList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePlansList) Version() int {

	return 1
}

// SparsePlan represents the sparse version of a plan.
type SparsePlan struct {
	// Contains the description of the plan.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Contains the key identifier of the plan.
	Key *string `json:"key,omitempty" msgpack:"key,omitempty" bson:"key,omitempty" mapstructure:"key,omitempty"`

	// Contains the name of the plan.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Defines the quota for each identity.
	Quotas *map[string]int `json:"-" msgpack:"-" bson:"-" mapstructure:"-,omitempty"`

	// RequireAdminValidation indicates if the plan requires an admin approval.
	RequireAdminValidation *bool `json:"-" msgpack:"-" bson:"requireadminvalidation,omitempty" mapstructure:"-,omitempty"`

	// Roles defines the roles given to the account.
	Roles *[]string `json:"-" msgpack:"-" bson:"-" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePlan returns a new  SparsePlan.
func NewSparsePlan() *SparsePlan {
	return &SparsePlan{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePlan) Identity() elemental.Identity {

	return PlanIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePlan) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePlan) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePlan) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePlan{}

	if o.Description != nil {
		s.Description = o.Description
	}
	if o.Key != nil {
		s.Key = o.Key
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.RequireAdminValidation != nil {
		s.RequireAdminValidation = o.RequireAdminValidation
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePlan) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePlan{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Key != nil {
		o.Key = s.Key
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.RequireAdminValidation != nil {
		o.RequireAdminValidation = s.RequireAdminValidation
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparsePlan) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePlan) ToPlain() elemental.PlainIdentifiable {

	out := NewPlan()
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Key != nil {
		out.Key = *o.Key
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Quotas != nil {
		out.Quotas = *o.Quotas
	}
	if o.RequireAdminValidation != nil {
		out.RequireAdminValidation = *o.RequireAdminValidation
	}
	if o.Roles != nil {
		out.Roles = *o.Roles
	}

	return out
}

// DeepCopy returns a deep copy if the SparsePlan.
func (o *SparsePlan) DeepCopy() *SparsePlan {

	if o == nil {
		return nil
	}

	out := &SparsePlan{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePlan.
func (o *SparsePlan) DeepCopyInto(out *SparsePlan) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePlan: %s", err))
	}

	*out = *target.(*SparsePlan)
}

type mongoAttributesPlan struct {
	Description            string `bson:"description"`
	Key                    string `bson:"key"`
	Name                   string `bson:"name"`
	RequireAdminValidation bool   `bson:"requireadminvalidation"`
}
type mongoAttributesSparsePlan struct {
	Description            *string `bson:"description,omitempty"`
	Key                    *string `bson:"key,omitempty"`
	Name                   *string `bson:"name,omitempty"`
	RequireAdminValidation *bool   `bson:"requireadminvalidation,omitempty"`
}
