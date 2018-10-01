package gaia

import (
	"fmt"
	"sync"

	"go.aporeto.io/elemental"
)

// QuotaCheckIdentity represents the Identity of the object.
var QuotaCheckIdentity = elemental.Identity{
	Name:     "quotacheck",
	Category: "quotacheck",
	Package:  "squall",
	Private:  true,
}

// QuotaChecksList represents a list of QuotaChecks
type QuotaChecksList []*QuotaCheck

// Identity returns the identity of the objects in the list.
func (o QuotaChecksList) Identity() elemental.Identity {

	return QuotaCheckIdentity
}

// Copy returns a pointer to a copy the QuotaChecksList.
func (o QuotaChecksList) Copy() elemental.Identifiables {

	copy := append(QuotaChecksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the QuotaChecksList.
func (o QuotaChecksList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(QuotaChecksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*QuotaCheck))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o QuotaChecksList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o QuotaChecksList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o QuotaChecksList) Version() int {

	return 1
}

// QuotaCheck represents the model of a quotacheck
type QuotaCheck struct {
	// Contains the maximum number of matching entities that can be created.
	Quota int `json:"quota" bson:"-" mapstructure:"quota,omitempty"`

	// The identity name of the object you want to check the quota on.
	TargetIdentity string `json:"targetIdentity" bson:"-" mapstructure:"targetIdentity,omitempty"`

	// The namespace from which you want to check the quota on.
	TargetNamespace string `json:"targetNamespace" bson:"-" mapstructure:"targetNamespace,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewQuotaCheck returns a new *QuotaCheck
func NewQuotaCheck() *QuotaCheck {

	return &QuotaCheck{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *QuotaCheck) Identity() elemental.Identity {

	return QuotaCheckIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *QuotaCheck) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *QuotaCheck) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *QuotaCheck) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *QuotaCheck) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *QuotaCheck) Doc() string {
	return `This api allows to verify the quota for a given identity in a given namespace
with the given tags.`
}

func (o *QuotaCheck) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *QuotaCheck) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("targetIdentity", o.TargetIdentity); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("targetNamespace", o.TargetNamespace); err != nil {
		requiredErrors = append(requiredErrors, err)
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
func (*QuotaCheck) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := QuotaCheckAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return QuotaCheckLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*QuotaCheck) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return QuotaCheckAttributesMap
}

// QuotaCheckAttributesMap represents the map of attribute for QuotaCheck.
var QuotaCheckAttributesMap = map[string]elemental.AttributeSpecification{
	"Quota": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Quota",
		Description:    `Contains the maximum number of matching entities that can be created.`,
		Exposed:        true,
		Name:           "quota",
		ReadOnly:       true,
		Type:           "integer",
	},
	"TargetIdentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetIdentity",
		Description:    `The identity name of the object you want to check the quota on.`,
		Exposed:        true,
		Name:           "targetIdentity",
		Required:       true,
		Type:           "string",
	},
	"TargetNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNamespace",
		Description:    `The namespace from which you want to check the quota on.`,
		Exposed:        true,
		Name:           "targetNamespace",
		Required:       true,
		Type:           "string",
	},
}

// QuotaCheckLowerCaseAttributesMap represents the map of attribute for QuotaCheck.
var QuotaCheckLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"quota": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Quota",
		Description:    `Contains the maximum number of matching entities that can be created.`,
		Exposed:        true,
		Name:           "quota",
		ReadOnly:       true,
		Type:           "integer",
	},
	"targetidentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetIdentity",
		Description:    `The identity name of the object you want to check the quota on.`,
		Exposed:        true,
		Name:           "targetIdentity",
		Required:       true,
		Type:           "string",
	},
	"targetnamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNamespace",
		Description:    `The namespace from which you want to check the quota on.`,
		Exposed:        true,
		Name:           "targetNamespace",
		Required:       true,
		Type:           "string",
	},
}
