package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// NamespacePolicyInfoPUIncomingTrafficActionValue represents the possible values for attribute "PUIncomingTrafficAction".
type NamespacePolicyInfoPUIncomingTrafficActionValue string

const (
	// NamespacePolicyInfoPUIncomingTrafficActionAllow represents the value Allow.
	NamespacePolicyInfoPUIncomingTrafficActionAllow NamespacePolicyInfoPUIncomingTrafficActionValue = "Allow"

	// NamespacePolicyInfoPUIncomingTrafficActionInherit represents the value Inherit.
	NamespacePolicyInfoPUIncomingTrafficActionInherit NamespacePolicyInfoPUIncomingTrafficActionValue = "Inherit"

	// NamespacePolicyInfoPUIncomingTrafficActionReject represents the value Reject.
	NamespacePolicyInfoPUIncomingTrafficActionReject NamespacePolicyInfoPUIncomingTrafficActionValue = "Reject"
)

// NamespacePolicyInfoPUOutgoingTrafficActionValue represents the possible values for attribute "PUOutgoingTrafficAction".
type NamespacePolicyInfoPUOutgoingTrafficActionValue string

const (
	// NamespacePolicyInfoPUOutgoingTrafficActionAllow represents the value Allow.
	NamespacePolicyInfoPUOutgoingTrafficActionAllow NamespacePolicyInfoPUOutgoingTrafficActionValue = "Allow"

	// NamespacePolicyInfoPUOutgoingTrafficActionInherit represents the value Inherit.
	NamespacePolicyInfoPUOutgoingTrafficActionInherit NamespacePolicyInfoPUOutgoingTrafficActionValue = "Inherit"

	// NamespacePolicyInfoPUOutgoingTrafficActionReject represents the value Reject.
	NamespacePolicyInfoPUOutgoingTrafficActionReject NamespacePolicyInfoPUOutgoingTrafficActionValue = "Reject"
)

// NamespacePolicyInfoIdentity represents the Identity of the object.
var NamespacePolicyInfoIdentity = elemental.Identity{
	Name:     "namespacepolicyinfo",
	Category: "namespacepolicyinfo",
	Package:  "squall",
	Private:  false,
}

// NamespacePolicyInfosList represents a list of NamespacePolicyInfos
type NamespacePolicyInfosList []*NamespacePolicyInfo

// Identity returns the identity of the objects in the list.
func (o NamespacePolicyInfosList) Identity() elemental.Identity {

	return NamespacePolicyInfoIdentity
}

// Copy returns a pointer to a copy the NamespacePolicyInfosList.
func (o NamespacePolicyInfosList) Copy() elemental.Identifiables {

	copy := append(NamespacePolicyInfosList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the NamespacePolicyInfosList.
func (o NamespacePolicyInfosList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(NamespacePolicyInfosList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*NamespacePolicyInfo))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o NamespacePolicyInfosList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o NamespacePolicyInfosList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the NamespacePolicyInfosList converted to SparseNamespacePolicyInfosList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o NamespacePolicyInfosList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseNamespacePolicyInfosList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseNamespacePolicyInfo)
	}

	return out
}

// Version returns the version of the content.
func (o NamespacePolicyInfosList) Version() int {

	return 1
}

// NamespacePolicyInfo represents the model of a namespacepolicyinfo
type NamespacePolicyInfo struct {
	// The processing unit action for incoming traffic for the namespace.
	PUIncomingTrafficAction NamespacePolicyInfoPUIncomingTrafficActionValue `json:"PUIncomingTrafficAction" msgpack:"PUIncomingTrafficAction" bson:"-" mapstructure:"PUIncomingTrafficAction,omitempty"`

	// The processing unit action for outgoing traffic for the namespace.
	PUOutgoingTrafficAction NamespacePolicyInfoPUOutgoingTrafficActionValue `json:"PUOutgoingTrafficAction" msgpack:"PUOutgoingTrafficAction" bson:"-" mapstructure:"PUOutgoingTrafficAction,omitempty"`

	// List of tag prefixes that will be used to suggest policies.
	Prefixes []string `json:"prefixes" msgpack:"prefixes" bson:"-" mapstructure:"prefixes,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewNamespacePolicyInfo returns a new *NamespacePolicyInfo
func NewNamespacePolicyInfo() *NamespacePolicyInfo {

	return &NamespacePolicyInfo{
		ModelVersion: 1,
		Prefixes:     []string{},
	}
}

// Identity returns the Identity of the object.
func (o *NamespacePolicyInfo) Identity() elemental.Identity {

	return NamespacePolicyInfoIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NamespacePolicyInfo) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NamespacePolicyInfo) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NamespacePolicyInfo) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesNamespacePolicyInfo{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NamespacePolicyInfo) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesNamespacePolicyInfo{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *NamespacePolicyInfo) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *NamespacePolicyInfo) BleveType() string {

	return "namespacepolicyinfo"
}

// DefaultOrder returns the list of default ordering fields.
func (o *NamespacePolicyInfo) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *NamespacePolicyInfo) Doc() string {

	return `Returns the policy info of the specified namespace.`
}

func (o *NamespacePolicyInfo) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *NamespacePolicyInfo) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseNamespacePolicyInfo{
			PUIncomingTrafficAction: &o.PUIncomingTrafficAction,
			PUOutgoingTrafficAction: &o.PUOutgoingTrafficAction,
			Prefixes:                &o.Prefixes,
		}
	}

	sp := &SparseNamespacePolicyInfo{}
	for _, f := range fields {
		switch f {
		case "PUIncomingTrafficAction":
			sp.PUIncomingTrafficAction = &(o.PUIncomingTrafficAction)
		case "PUOutgoingTrafficAction":
			sp.PUOutgoingTrafficAction = &(o.PUOutgoingTrafficAction)
		case "prefixes":
			sp.Prefixes = &(o.Prefixes)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseNamespacePolicyInfo to the object.
func (o *NamespacePolicyInfo) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseNamespacePolicyInfo)
	if so.PUIncomingTrafficAction != nil {
		o.PUIncomingTrafficAction = *so.PUIncomingTrafficAction
	}
	if so.PUOutgoingTrafficAction != nil {
		o.PUOutgoingTrafficAction = *so.PUOutgoingTrafficAction
	}
	if so.Prefixes != nil {
		o.Prefixes = *so.Prefixes
	}
}

// DeepCopy returns a deep copy if the NamespacePolicyInfo.
func (o *NamespacePolicyInfo) DeepCopy() *NamespacePolicyInfo {

	if o == nil {
		return nil
	}

	out := &NamespacePolicyInfo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *NamespacePolicyInfo.
func (o *NamespacePolicyInfo) DeepCopyInto(out *NamespacePolicyInfo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy NamespacePolicyInfo: %s", err))
	}

	*out = *target.(*NamespacePolicyInfo)
}

// Validate valides the current information stored into the structure.
func (o *NamespacePolicyInfo) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("PUIncomingTrafficAction", string(o.PUIncomingTrafficAction), []string{"Allow", "Reject", "Inherit"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("PUOutgoingTrafficAction", string(o.PUOutgoingTrafficAction), []string{"Allow", "Reject", "Inherit"}, false); err != nil {
		errors = errors.Append(err)
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
func (*NamespacePolicyInfo) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := NamespacePolicyInfoAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return NamespacePolicyInfoLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*NamespacePolicyInfo) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return NamespacePolicyInfoAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *NamespacePolicyInfo) ValueForAttribute(name string) interface{} {

	switch name {
	case "PUIncomingTrafficAction":
		return o.PUIncomingTrafficAction
	case "PUOutgoingTrafficAction":
		return o.PUOutgoingTrafficAction
	case "prefixes":
		return o.Prefixes
	}

	return nil
}

// NamespacePolicyInfoAttributesMap represents the map of attribute for NamespacePolicyInfo.
var NamespacePolicyInfoAttributesMap = map[string]elemental.AttributeSpecification{
	"PUIncomingTrafficAction": {
		AllowedChoices: []string{"Allow", "Reject", "Inherit"},
		ConvertedName:  "PUIncomingTrafficAction",
		Description:    `The processing unit action for incoming traffic for the namespace.`,
		Exposed:        true,
		Name:           "PUIncomingTrafficAction",
		ReadOnly:       true,
		Type:           "enum",
	},
	"PUOutgoingTrafficAction": {
		AllowedChoices: []string{"Allow", "Reject", "Inherit"},
		ConvertedName:  "PUOutgoingTrafficAction",
		Description:    `The processing unit action for outgoing traffic for the namespace.`,
		Exposed:        true,
		Name:           "PUOutgoingTrafficAction",
		ReadOnly:       true,
		Type:           "enum",
	},
	"Prefixes": {
		AllowedChoices: []string{},
		ConvertedName:  "Prefixes",
		Description:    `List of tag prefixes that will be used to suggest policies.`,
		Exposed:        true,
		Name:           "prefixes",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
}

// NamespacePolicyInfoLowerCaseAttributesMap represents the map of attribute for NamespacePolicyInfo.
var NamespacePolicyInfoLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"puincomingtrafficaction": {
		AllowedChoices: []string{"Allow", "Reject", "Inherit"},
		ConvertedName:  "PUIncomingTrafficAction",
		Description:    `The processing unit action for incoming traffic for the namespace.`,
		Exposed:        true,
		Name:           "PUIncomingTrafficAction",
		ReadOnly:       true,
		Type:           "enum",
	},
	"puoutgoingtrafficaction": {
		AllowedChoices: []string{"Allow", "Reject", "Inherit"},
		ConvertedName:  "PUOutgoingTrafficAction",
		Description:    `The processing unit action for outgoing traffic for the namespace.`,
		Exposed:        true,
		Name:           "PUOutgoingTrafficAction",
		ReadOnly:       true,
		Type:           "enum",
	},
	"prefixes": {
		AllowedChoices: []string{},
		ConvertedName:  "Prefixes",
		Description:    `List of tag prefixes that will be used to suggest policies.`,
		Exposed:        true,
		Name:           "prefixes",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
}

// SparseNamespacePolicyInfosList represents a list of SparseNamespacePolicyInfos
type SparseNamespacePolicyInfosList []*SparseNamespacePolicyInfo

// Identity returns the identity of the objects in the list.
func (o SparseNamespacePolicyInfosList) Identity() elemental.Identity {

	return NamespacePolicyInfoIdentity
}

// Copy returns a pointer to a copy the SparseNamespacePolicyInfosList.
func (o SparseNamespacePolicyInfosList) Copy() elemental.Identifiables {

	copy := append(SparseNamespacePolicyInfosList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseNamespacePolicyInfosList.
func (o SparseNamespacePolicyInfosList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseNamespacePolicyInfosList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseNamespacePolicyInfo))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseNamespacePolicyInfosList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseNamespacePolicyInfosList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseNamespacePolicyInfosList converted to NamespacePolicyInfosList.
func (o SparseNamespacePolicyInfosList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseNamespacePolicyInfosList) Version() int {

	return 1
}

// SparseNamespacePolicyInfo represents the sparse version of a namespacepolicyinfo.
type SparseNamespacePolicyInfo struct {
	// The processing unit action for incoming traffic for the namespace.
	PUIncomingTrafficAction *NamespacePolicyInfoPUIncomingTrafficActionValue `json:"PUIncomingTrafficAction,omitempty" msgpack:"PUIncomingTrafficAction,omitempty" bson:"-" mapstructure:"PUIncomingTrafficAction,omitempty"`

	// The processing unit action for outgoing traffic for the namespace.
	PUOutgoingTrafficAction *NamespacePolicyInfoPUOutgoingTrafficActionValue `json:"PUOutgoingTrafficAction,omitempty" msgpack:"PUOutgoingTrafficAction,omitempty" bson:"-" mapstructure:"PUOutgoingTrafficAction,omitempty"`

	// List of tag prefixes that will be used to suggest policies.
	Prefixes *[]string `json:"prefixes,omitempty" msgpack:"prefixes,omitempty" bson:"-" mapstructure:"prefixes,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseNamespacePolicyInfo returns a new  SparseNamespacePolicyInfo.
func NewSparseNamespacePolicyInfo() *SparseNamespacePolicyInfo {
	return &SparseNamespacePolicyInfo{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseNamespacePolicyInfo) Identity() elemental.Identity {

	return NamespacePolicyInfoIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseNamespacePolicyInfo) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseNamespacePolicyInfo) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseNamespacePolicyInfo) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseNamespacePolicyInfo{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseNamespacePolicyInfo) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseNamespacePolicyInfo{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseNamespacePolicyInfo) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseNamespacePolicyInfo) ToPlain() elemental.PlainIdentifiable {

	out := NewNamespacePolicyInfo()
	if o.PUIncomingTrafficAction != nil {
		out.PUIncomingTrafficAction = *o.PUIncomingTrafficAction
	}
	if o.PUOutgoingTrafficAction != nil {
		out.PUOutgoingTrafficAction = *o.PUOutgoingTrafficAction
	}
	if o.Prefixes != nil {
		out.Prefixes = *o.Prefixes
	}

	return out
}

// DeepCopy returns a deep copy if the SparseNamespacePolicyInfo.
func (o *SparseNamespacePolicyInfo) DeepCopy() *SparseNamespacePolicyInfo {

	if o == nil {
		return nil
	}

	out := &SparseNamespacePolicyInfo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseNamespacePolicyInfo.
func (o *SparseNamespacePolicyInfo) DeepCopyInto(out *SparseNamespacePolicyInfo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseNamespacePolicyInfo: %s", err))
	}

	*out = *target.(*SparseNamespacePolicyInfo)
}

type mongoAttributesNamespacePolicyInfo struct {
}
type mongoAttributesSparseNamespacePolicyInfo struct {
}
