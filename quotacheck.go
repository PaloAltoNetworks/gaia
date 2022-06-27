package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// QuotaCheckIdentity represents the Identity of the object.
var QuotaCheckIdentity = elemental.Identity{
	Name:     "quotacheck",
	Category: "quotacheck",
	Package:  "squall",
	Private:  false,
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

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o QuotaChecksList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the QuotaChecksList converted to SparseQuotaChecksList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o QuotaChecksList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseQuotaChecksList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseQuotaCheck)
	}

	return out
}

// Version returns the version of the content.
func (o QuotaChecksList) Version() int {

	return 1
}

// QuotaCheck represents the model of a quotacheck
type QuotaCheck struct {
	// Contains the maximum number of matching entities that can be created.
	Quota int `json:"quota" msgpack:"quota" bson:"-" mapstructure:"quota,omitempty"`

	// If the parameter `remaining=true` is passed, this value will be populated with
	// the number of remaining objects in the quota.
	Remaining int `json:"remaining" msgpack:"remaining" bson:"-" mapstructure:"remaining,omitempty"`

	// The identity name of the object you want to check the quota on.
	TargetIdentity string `json:"targetIdentity" msgpack:"targetIdentity" bson:"-" mapstructure:"targetIdentity,omitempty"`

	// The namespace from which you want to check the quota on.
	TargetNamespace string `json:"targetNamespace" msgpack:"targetNamespace" bson:"-" mapstructure:"targetNamespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewQuotaCheck returns a new *QuotaCheck
func NewQuotaCheck() *QuotaCheck {

	return &QuotaCheck{
		ModelVersion: 1,
		Remaining:    -1,
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

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *QuotaCheck) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesQuotaCheck{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *QuotaCheck) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesQuotaCheck{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *QuotaCheck) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *QuotaCheck) BleveType() string {

	return "quotacheck"
}

// DefaultOrder returns the list of default ordering fields.
func (o *QuotaCheck) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *QuotaCheck) Doc() string {

	return `Allows you to verify the quota for a given identity in a given namespace
with the given tags.`
}

func (o *QuotaCheck) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *QuotaCheck) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseQuotaCheck{
			Quota:           &o.Quota,
			Remaining:       &o.Remaining,
			TargetIdentity:  &o.TargetIdentity,
			TargetNamespace: &o.TargetNamespace,
		}
	}

	sp := &SparseQuotaCheck{}
	for _, f := range fields {
		switch f {
		case "quota":
			sp.Quota = &(o.Quota)
		case "remaining":
			sp.Remaining = &(o.Remaining)
		case "targetIdentity":
			sp.TargetIdentity = &(o.TargetIdentity)
		case "targetNamespace":
			sp.TargetNamespace = &(o.TargetNamespace)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseQuotaCheck to the object.
func (o *QuotaCheck) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseQuotaCheck)
	if so.Quota != nil {
		o.Quota = *so.Quota
	}
	if so.Remaining != nil {
		o.Remaining = *so.Remaining
	}
	if so.TargetIdentity != nil {
		o.TargetIdentity = *so.TargetIdentity
	}
	if so.TargetNamespace != nil {
		o.TargetNamespace = *so.TargetNamespace
	}
}

// DeepCopy returns a deep copy if the QuotaCheck.
func (o *QuotaCheck) DeepCopy() *QuotaCheck {

	if o == nil {
		return nil
	}

	out := &QuotaCheck{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *QuotaCheck.
func (o *QuotaCheck) DeepCopyInto(out *QuotaCheck) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy QuotaCheck: %s", err))
	}

	*out = *target.(*QuotaCheck)
}

// Validate valides the current information stored into the structure.
func (o *QuotaCheck) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("targetIdentity", o.TargetIdentity); err != nil {
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

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *QuotaCheck) ValueForAttribute(name string) interface{} {

	switch name {
	case "quota":
		return o.Quota
	case "remaining":
		return o.Remaining
	case "targetIdentity":
		return o.TargetIdentity
	case "targetNamespace":
		return o.TargetNamespace
	}

	return nil
}

// QuotaCheckAttributesMap represents the map of attribute for QuotaCheck.
var QuotaCheckAttributesMap = map[string]elemental.AttributeSpecification{
	"Quota": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Quota",
		Description:    `Contains the maximum number of matching entities that can be created.`,
		Exposed:        true,
		Name:           "quota",
		ReadOnly:       true,
		Type:           "integer",
	},
	"Remaining": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Remaining",
		DefaultValue:   -1,
		Description: `If the parameter ` + "`" + `remaining=true` + "`" + ` is passed, this value will be populated with
the number of remaining objects in the quota.`,
		Exposed:  true,
		Name:     "remaining",
		ReadOnly: true,
		Type:     "integer",
	},
	"TargetIdentity": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetIdentity",
		Description:    `The identity name of the object you want to check the quota on.`,
		Exposed:        true,
		Name:           "targetIdentity",
		Required:       true,
		Type:           "string",
	},
	"TargetNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetNamespace",
		Description:    `The namespace from which you want to check the quota on.`,
		Exposed:        true,
		Name:           "targetNamespace",
		Type:           "string",
	},
}

// QuotaCheckLowerCaseAttributesMap represents the map of attribute for QuotaCheck.
var QuotaCheckLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"quota": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Quota",
		Description:    `Contains the maximum number of matching entities that can be created.`,
		Exposed:        true,
		Name:           "quota",
		ReadOnly:       true,
		Type:           "integer",
	},
	"remaining": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Remaining",
		DefaultValue:   -1,
		Description: `If the parameter ` + "`" + `remaining=true` + "`" + ` is passed, this value will be populated with
the number of remaining objects in the quota.`,
		Exposed:  true,
		Name:     "remaining",
		ReadOnly: true,
		Type:     "integer",
	},
	"targetidentity": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetIdentity",
		Description:    `The identity name of the object you want to check the quota on.`,
		Exposed:        true,
		Name:           "targetIdentity",
		Required:       true,
		Type:           "string",
	},
	"targetnamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetNamespace",
		Description:    `The namespace from which you want to check the quota on.`,
		Exposed:        true,
		Name:           "targetNamespace",
		Type:           "string",
	},
}

// SparseQuotaChecksList represents a list of SparseQuotaChecks
type SparseQuotaChecksList []*SparseQuotaCheck

// Identity returns the identity of the objects in the list.
func (o SparseQuotaChecksList) Identity() elemental.Identity {

	return QuotaCheckIdentity
}

// Copy returns a pointer to a copy the SparseQuotaChecksList.
func (o SparseQuotaChecksList) Copy() elemental.Identifiables {

	copy := append(SparseQuotaChecksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseQuotaChecksList.
func (o SparseQuotaChecksList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseQuotaChecksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseQuotaCheck))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseQuotaChecksList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseQuotaChecksList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseQuotaChecksList converted to QuotaChecksList.
func (o SparseQuotaChecksList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseQuotaChecksList) Version() int {

	return 1
}

// SparseQuotaCheck represents the sparse version of a quotacheck.
type SparseQuotaCheck struct {
	// Contains the maximum number of matching entities that can be created.
	Quota *int `json:"quota,omitempty" msgpack:"quota,omitempty" bson:"-" mapstructure:"quota,omitempty"`

	// If the parameter `remaining=true` is passed, this value will be populated with
	// the number of remaining objects in the quota.
	Remaining *int `json:"remaining,omitempty" msgpack:"remaining,omitempty" bson:"-" mapstructure:"remaining,omitempty"`

	// The identity name of the object you want to check the quota on.
	TargetIdentity *string `json:"targetIdentity,omitempty" msgpack:"targetIdentity,omitempty" bson:"-" mapstructure:"targetIdentity,omitempty"`

	// The namespace from which you want to check the quota on.
	TargetNamespace *string `json:"targetNamespace,omitempty" msgpack:"targetNamespace,omitempty" bson:"-" mapstructure:"targetNamespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseQuotaCheck returns a new  SparseQuotaCheck.
func NewSparseQuotaCheck() *SparseQuotaCheck {
	return &SparseQuotaCheck{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseQuotaCheck) Identity() elemental.Identity {

	return QuotaCheckIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseQuotaCheck) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseQuotaCheck) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseQuotaCheck) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseQuotaCheck{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseQuotaCheck) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseQuotaCheck{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseQuotaCheck) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseQuotaCheck) ToPlain() elemental.PlainIdentifiable {

	out := NewQuotaCheck()
	if o.Quota != nil {
		out.Quota = *o.Quota
	}
	if o.Remaining != nil {
		out.Remaining = *o.Remaining
	}
	if o.TargetIdentity != nil {
		out.TargetIdentity = *o.TargetIdentity
	}
	if o.TargetNamespace != nil {
		out.TargetNamespace = *o.TargetNamespace
	}

	return out
}

// DeepCopy returns a deep copy if the SparseQuotaCheck.
func (o *SparseQuotaCheck) DeepCopy() *SparseQuotaCheck {

	if o == nil {
		return nil
	}

	out := &SparseQuotaCheck{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseQuotaCheck.
func (o *SparseQuotaCheck) DeepCopyInto(out *SparseQuotaCheck) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseQuotaCheck: %s", err))
	}

	*out = *target.(*SparseQuotaCheck)
}

type mongoAttributesQuotaCheck struct {
}
type mongoAttributesSparseQuotaCheck struct {
}
