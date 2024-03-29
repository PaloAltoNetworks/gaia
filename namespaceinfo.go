// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// NamespaceInfoPUIncomingTrafficActionValue represents the possible values for attribute "PUIncomingTrafficAction".
type NamespaceInfoPUIncomingTrafficActionValue string

const (
	// NamespaceInfoPUIncomingTrafficActionAllow represents the value Allow.
	NamespaceInfoPUIncomingTrafficActionAllow NamespaceInfoPUIncomingTrafficActionValue = "Allow"

	// NamespaceInfoPUIncomingTrafficActionInherit represents the value Inherit.
	NamespaceInfoPUIncomingTrafficActionInherit NamespaceInfoPUIncomingTrafficActionValue = "Inherit"

	// NamespaceInfoPUIncomingTrafficActionReject represents the value Reject.
	NamespaceInfoPUIncomingTrafficActionReject NamespaceInfoPUIncomingTrafficActionValue = "Reject"
)

// NamespaceInfoPUOutgoingTrafficActionValue represents the possible values for attribute "PUOutgoingTrafficAction".
type NamespaceInfoPUOutgoingTrafficActionValue string

const (
	// NamespaceInfoPUOutgoingTrafficActionAllow represents the value Allow.
	NamespaceInfoPUOutgoingTrafficActionAllow NamespaceInfoPUOutgoingTrafficActionValue = "Allow"

	// NamespaceInfoPUOutgoingTrafficActionInherit represents the value Inherit.
	NamespaceInfoPUOutgoingTrafficActionInherit NamespaceInfoPUOutgoingTrafficActionValue = "Inherit"

	// NamespaceInfoPUOutgoingTrafficActionReject represents the value Reject.
	NamespaceInfoPUOutgoingTrafficActionReject NamespaceInfoPUOutgoingTrafficActionValue = "Reject"
)

// NamespaceInfoIdentity represents the Identity of the object.
var NamespaceInfoIdentity = elemental.Identity{
	Name:     "namespaceinfo",
	Category: "namespaceinfo",
	Package:  "squall",
	Private:  false,
}

// NamespaceInfosList represents a list of NamespaceInfos
type NamespaceInfosList []*NamespaceInfo

// Identity returns the identity of the objects in the list.
func (o NamespaceInfosList) Identity() elemental.Identity {

	return NamespaceInfoIdentity
}

// Copy returns a pointer to a copy the NamespaceInfosList.
func (o NamespaceInfosList) Copy() elemental.Identifiables {

	out := append(NamespaceInfosList{}, o...)
	return &out
}

// Append appends the objects to the a new copy of the NamespaceInfosList.
func (o NamespaceInfosList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(NamespaceInfosList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*NamespaceInfo))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o NamespaceInfosList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o NamespaceInfosList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the NamespaceInfosList converted to SparseNamespaceInfosList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o NamespaceInfosList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseNamespaceInfosList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseNamespaceInfo)
	}

	return out
}

// Version returns the version of the content.
func (o NamespaceInfosList) Version() int {

	return 1
}

// NamespaceInfo represents the model of a namespaceinfo
type NamespaceInfo struct {
	// The processing unit action for incoming traffic for the namespace.
	PUIncomingTrafficAction NamespaceInfoPUIncomingTrafficActionValue `json:"PUIncomingTrafficAction" msgpack:"PUIncomingTrafficAction" bson:"-" mapstructure:"PUIncomingTrafficAction,omitempty"`

	// The processing unit action for outgoing traffic for the namespace.
	PUOutgoingTrafficAction NamespaceInfoPUOutgoingTrafficActionValue `json:"PUOutgoingTrafficAction" msgpack:"PUOutgoingTrafficAction" bson:"-" mapstructure:"PUOutgoingTrafficAction,omitempty"`

	// Description of the namespace.
	Description string `json:"description" msgpack:"description" bson:"-" mapstructure:"description,omitempty"`

	// Name of the namespace.
	Name string `json:"name" msgpack:"name" bson:"-" mapstructure:"name,omitempty"`

	// List of tag prefixes that will be used to suggest policies.
	Prefixes []string `json:"prefixes" msgpack:"prefixes" bson:"-" mapstructure:"prefixes,omitempty"`

	// Defines if the namespace is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"-" mapstructure:"protected,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewNamespaceInfo returns a new *NamespaceInfo
func NewNamespaceInfo() *NamespaceInfo {

	return &NamespaceInfo{
		ModelVersion: 1,
		Prefixes:     []string{},
	}
}

// Identity returns the Identity of the object.
func (o *NamespaceInfo) Identity() elemental.Identity {

	return NamespaceInfoIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NamespaceInfo) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NamespaceInfo) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NamespaceInfo) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesNamespaceInfo{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NamespaceInfo) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesNamespaceInfo{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *NamespaceInfo) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *NamespaceInfo) BleveType() string {

	return "namespaceinfo"
}

// DefaultOrder returns the list of default ordering fields.
func (o *NamespaceInfo) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *NamespaceInfo) Doc() string {

	return `Returns the information of the specified namespace.`
}

func (o *NamespaceInfo) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *NamespaceInfo) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseNamespaceInfo{
			PUIncomingTrafficAction: &o.PUIncomingTrafficAction,
			PUOutgoingTrafficAction: &o.PUOutgoingTrafficAction,
			Description:             &o.Description,
			Name:                    &o.Name,
			Prefixes:                &o.Prefixes,
			Protected:               &o.Protected,
		}
	}

	sp := &SparseNamespaceInfo{}
	for _, f := range fields {
		switch f {
		case "PUIncomingTrafficAction":
			sp.PUIncomingTrafficAction = &(o.PUIncomingTrafficAction)
		case "PUOutgoingTrafficAction":
			sp.PUOutgoingTrafficAction = &(o.PUOutgoingTrafficAction)
		case "description":
			sp.Description = &(o.Description)
		case "name":
			sp.Name = &(o.Name)
		case "prefixes":
			sp.Prefixes = &(o.Prefixes)
		case "protected":
			sp.Protected = &(o.Protected)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseNamespaceInfo to the object.
func (o *NamespaceInfo) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseNamespaceInfo)
	if so.PUIncomingTrafficAction != nil {
		o.PUIncomingTrafficAction = *so.PUIncomingTrafficAction
	}
	if so.PUOutgoingTrafficAction != nil {
		o.PUOutgoingTrafficAction = *so.PUOutgoingTrafficAction
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Prefixes != nil {
		o.Prefixes = *so.Prefixes
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
}

// DeepCopy returns a deep copy if the NamespaceInfo.
func (o *NamespaceInfo) DeepCopy() *NamespaceInfo {

	if o == nil {
		return nil
	}

	out := &NamespaceInfo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *NamespaceInfo.
func (o *NamespaceInfo) DeepCopyInto(out *NamespaceInfo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy NamespaceInfo: %s", err))
	}

	*out = *target.(*NamespaceInfo)
}

// Validate valides the current information stored into the structure.
func (o *NamespaceInfo) Validate() error {

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
func (*NamespaceInfo) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := NamespaceInfoAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return NamespaceInfoLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*NamespaceInfo) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return NamespaceInfoAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *NamespaceInfo) ValueForAttribute(name string) any {

	switch name {
	case "PUIncomingTrafficAction":
		return o.PUIncomingTrafficAction
	case "PUOutgoingTrafficAction":
		return o.PUOutgoingTrafficAction
	case "description":
		return o.Description
	case "name":
		return o.Name
	case "prefixes":
		return o.Prefixes
	case "protected":
		return o.Protected
	}

	return nil
}

// NamespaceInfoAttributesMap represents the map of attribute for NamespaceInfo.
var NamespaceInfoAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the namespace.`,
		Exposed:        true,
		Name:           "description",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"Name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the namespace.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "name",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
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
	"Protected": {
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Defines if the namespace is protected.`,
		Exposed:        true,
		Name:           "protected",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "boolean",
	},
}

// NamespaceInfoLowerCaseAttributesMap represents the map of attribute for NamespaceInfo.
var NamespaceInfoLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the namespace.`,
		Exposed:        true,
		Name:           "description",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the namespace.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "name",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
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
	"protected": {
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Defines if the namespace is protected.`,
		Exposed:        true,
		Name:           "protected",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "boolean",
	},
}

// SparseNamespaceInfosList represents a list of SparseNamespaceInfos
type SparseNamespaceInfosList []*SparseNamespaceInfo

// Identity returns the identity of the objects in the list.
func (o SparseNamespaceInfosList) Identity() elemental.Identity {

	return NamespaceInfoIdentity
}

// Copy returns a pointer to a copy the SparseNamespaceInfosList.
func (o SparseNamespaceInfosList) Copy() elemental.Identifiables {

	copy := append(SparseNamespaceInfosList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseNamespaceInfosList.
func (o SparseNamespaceInfosList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseNamespaceInfosList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseNamespaceInfo))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseNamespaceInfosList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseNamespaceInfosList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseNamespaceInfosList converted to NamespaceInfosList.
func (o SparseNamespaceInfosList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseNamespaceInfosList) Version() int {

	return 1
}

// SparseNamespaceInfo represents the sparse version of a namespaceinfo.
type SparseNamespaceInfo struct {
	// The processing unit action for incoming traffic for the namespace.
	PUIncomingTrafficAction *NamespaceInfoPUIncomingTrafficActionValue `json:"PUIncomingTrafficAction,omitempty" msgpack:"PUIncomingTrafficAction,omitempty" bson:"-" mapstructure:"PUIncomingTrafficAction,omitempty"`

	// The processing unit action for outgoing traffic for the namespace.
	PUOutgoingTrafficAction *NamespaceInfoPUOutgoingTrafficActionValue `json:"PUOutgoingTrafficAction,omitempty" msgpack:"PUOutgoingTrafficAction,omitempty" bson:"-" mapstructure:"PUOutgoingTrafficAction,omitempty"`

	// Description of the namespace.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"-" mapstructure:"description,omitempty"`

	// Name of the namespace.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"-" mapstructure:"name,omitempty"`

	// List of tag prefixes that will be used to suggest policies.
	Prefixes *[]string `json:"prefixes,omitempty" msgpack:"prefixes,omitempty" bson:"-" mapstructure:"prefixes,omitempty"`

	// Defines if the namespace is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"-" mapstructure:"protected,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseNamespaceInfo returns a new  SparseNamespaceInfo.
func NewSparseNamespaceInfo() *SparseNamespaceInfo {
	return &SparseNamespaceInfo{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseNamespaceInfo) Identity() elemental.Identity {

	return NamespaceInfoIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseNamespaceInfo) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseNamespaceInfo) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseNamespaceInfo) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseNamespaceInfo{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseNamespaceInfo) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseNamespaceInfo{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseNamespaceInfo) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseNamespaceInfo) ToPlain() elemental.PlainIdentifiable {

	out := NewNamespaceInfo()
	if o.PUIncomingTrafficAction != nil {
		out.PUIncomingTrafficAction = *o.PUIncomingTrafficAction
	}
	if o.PUOutgoingTrafficAction != nil {
		out.PUOutgoingTrafficAction = *o.PUOutgoingTrafficAction
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Prefixes != nil {
		out.Prefixes = *o.Prefixes
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}

	return out
}

// DeepCopy returns a deep copy if the SparseNamespaceInfo.
func (o *SparseNamespaceInfo) DeepCopy() *SparseNamespaceInfo {

	if o == nil {
		return nil
	}

	out := &SparseNamespaceInfo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseNamespaceInfo.
func (o *SparseNamespaceInfo) DeepCopyInto(out *SparseNamespaceInfo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseNamespaceInfo: %s", err))
	}

	*out = *target.(*SparseNamespaceInfo)
}

type mongoAttributesNamespaceInfo struct {
}
type mongoAttributesSparseNamespaceInfo struct {
}
