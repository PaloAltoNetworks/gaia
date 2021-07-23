package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// NamespaceDetailPUIncomingTrafficActionValue represents the possible values for attribute "PUIncomingTrafficAction".
type NamespaceDetailPUIncomingTrafficActionValue string

const (
	// NamespaceDetailPUIncomingTrafficActionAllow represents the value Allow.
	NamespaceDetailPUIncomingTrafficActionAllow NamespaceDetailPUIncomingTrafficActionValue = "Allow"

	// NamespaceDetailPUIncomingTrafficActionInherit represents the value Inherit.
	NamespaceDetailPUIncomingTrafficActionInherit NamespaceDetailPUIncomingTrafficActionValue = "Inherit"

	// NamespaceDetailPUIncomingTrafficActionReject represents the value Reject.
	NamespaceDetailPUIncomingTrafficActionReject NamespaceDetailPUIncomingTrafficActionValue = "Reject"
)

// NamespaceDetailPUOutgoingTrafficActionValue represents the possible values for attribute "PUOutgoingTrafficAction".
type NamespaceDetailPUOutgoingTrafficActionValue string

const (
	// NamespaceDetailPUOutgoingTrafficActionAllow represents the value Allow.
	NamespaceDetailPUOutgoingTrafficActionAllow NamespaceDetailPUOutgoingTrafficActionValue = "Allow"

	// NamespaceDetailPUOutgoingTrafficActionInherit represents the value Inherit.
	NamespaceDetailPUOutgoingTrafficActionInherit NamespaceDetailPUOutgoingTrafficActionValue = "Inherit"

	// NamespaceDetailPUOutgoingTrafficActionReject represents the value Reject.
	NamespaceDetailPUOutgoingTrafficActionReject NamespaceDetailPUOutgoingTrafficActionValue = "Reject"
)

// NamespaceDetailIdentity represents the Identity of the object.
var NamespaceDetailIdentity = elemental.Identity{
	Name:     "namespacedetail",
	Category: "namespacedetails",
	Package:  "squall",
	Private:  false,
}

// NamespaceDetailsList represents a list of NamespaceDetails
type NamespaceDetailsList []*NamespaceDetail

// Identity returns the identity of the objects in the list.
func (o NamespaceDetailsList) Identity() elemental.Identity {

	return NamespaceDetailIdentity
}

// Copy returns a pointer to a copy the NamespaceDetailsList.
func (o NamespaceDetailsList) Copy() elemental.Identifiables {

	copy := append(NamespaceDetailsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the NamespaceDetailsList.
func (o NamespaceDetailsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(NamespaceDetailsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*NamespaceDetail))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o NamespaceDetailsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o NamespaceDetailsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the NamespaceDetailsList converted to SparseNamespaceDetailsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o NamespaceDetailsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseNamespaceDetailsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseNamespaceDetail)
	}

	return out
}

// Version returns the version of the content.
func (o NamespaceDetailsList) Version() int {

	return 1
}

// NamespaceDetail represents the model of a namespacedetail
type NamespaceDetail struct {
	// The processing unit action for incoming traffic for the namespace.
	PUIncomingTrafficAction NamespaceDetailPUIncomingTrafficActionValue `json:"PUIncomingTrafficAction" msgpack:"PUIncomingTrafficAction" bson:"-" mapstructure:"PUIncomingTrafficAction,omitempty"`

	// The processing unit action for outgoing traffic for the namespace.
	PUOutgoingTrafficAction NamespaceDetailPUOutgoingTrafficActionValue `json:"PUOutgoingTrafficAction" msgpack:"PUOutgoingTrafficAction" bson:"-" mapstructure:"PUOutgoingTrafficAction,omitempty"`

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

// NewNamespaceDetail returns a new *NamespaceDetail
func NewNamespaceDetail() *NamespaceDetail {

	return &NamespaceDetail{
		ModelVersion: 1,
		Prefixes:     []string{},
	}
}

// Identity returns the Identity of the object.
func (o *NamespaceDetail) Identity() elemental.Identity {

	return NamespaceDetailIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NamespaceDetail) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NamespaceDetail) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NamespaceDetail) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesNamespaceDetail{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NamespaceDetail) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesNamespaceDetail{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *NamespaceDetail) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *NamespaceDetail) BleveType() string {

	return "namespacedetail"
}

// DefaultOrder returns the list of default ordering fields.
func (o *NamespaceDetail) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *NamespaceDetail) Doc() string {

	return `Returns the information of the specified namespace.`
}

func (o *NamespaceDetail) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *NamespaceDetail) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseNamespaceDetail{
			PUIncomingTrafficAction: &o.PUIncomingTrafficAction,
			PUOutgoingTrafficAction: &o.PUOutgoingTrafficAction,
			Description:             &o.Description,
			Name:                    &o.Name,
			Prefixes:                &o.Prefixes,
			Protected:               &o.Protected,
		}
	}

	sp := &SparseNamespaceDetail{}
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

// Patch apply the non nil value of a *SparseNamespaceDetail to the object.
func (o *NamespaceDetail) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseNamespaceDetail)
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

// DeepCopy returns a deep copy if the NamespaceDetail.
func (o *NamespaceDetail) DeepCopy() *NamespaceDetail {

	if o == nil {
		return nil
	}

	out := &NamespaceDetail{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *NamespaceDetail.
func (o *NamespaceDetail) DeepCopyInto(out *NamespaceDetail) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy NamespaceDetail: %s", err))
	}

	*out = *target.(*NamespaceDetail)
}

// Validate valides the current information stored into the structure.
func (o *NamespaceDetail) Validate() error {

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
func (*NamespaceDetail) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := NamespaceDetailAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return NamespaceDetailLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*NamespaceDetail) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return NamespaceDetailAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *NamespaceDetail) ValueForAttribute(name string) interface{} {

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

// NamespaceDetailAttributesMap represents the map of attribute for NamespaceDetail.
var NamespaceDetailAttributesMap = map[string]elemental.AttributeSpecification{
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

// NamespaceDetailLowerCaseAttributesMap represents the map of attribute for NamespaceDetail.
var NamespaceDetailLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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

// SparseNamespaceDetailsList represents a list of SparseNamespaceDetails
type SparseNamespaceDetailsList []*SparseNamespaceDetail

// Identity returns the identity of the objects in the list.
func (o SparseNamespaceDetailsList) Identity() elemental.Identity {

	return NamespaceDetailIdentity
}

// Copy returns a pointer to a copy the SparseNamespaceDetailsList.
func (o SparseNamespaceDetailsList) Copy() elemental.Identifiables {

	copy := append(SparseNamespaceDetailsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseNamespaceDetailsList.
func (o SparseNamespaceDetailsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseNamespaceDetailsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseNamespaceDetail))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseNamespaceDetailsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseNamespaceDetailsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseNamespaceDetailsList converted to NamespaceDetailsList.
func (o SparseNamespaceDetailsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseNamespaceDetailsList) Version() int {

	return 1
}

// SparseNamespaceDetail represents the sparse version of a namespacedetail.
type SparseNamespaceDetail struct {
	// The processing unit action for incoming traffic for the namespace.
	PUIncomingTrafficAction *NamespaceDetailPUIncomingTrafficActionValue `json:"PUIncomingTrafficAction,omitempty" msgpack:"PUIncomingTrafficAction,omitempty" bson:"-" mapstructure:"PUIncomingTrafficAction,omitempty"`

	// The processing unit action for outgoing traffic for the namespace.
	PUOutgoingTrafficAction *NamespaceDetailPUOutgoingTrafficActionValue `json:"PUOutgoingTrafficAction,omitempty" msgpack:"PUOutgoingTrafficAction,omitempty" bson:"-" mapstructure:"PUOutgoingTrafficAction,omitempty"`

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

// NewSparseNamespaceDetail returns a new  SparseNamespaceDetail.
func NewSparseNamespaceDetail() *SparseNamespaceDetail {
	return &SparseNamespaceDetail{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseNamespaceDetail) Identity() elemental.Identity {

	return NamespaceDetailIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseNamespaceDetail) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseNamespaceDetail) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseNamespaceDetail) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseNamespaceDetail{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseNamespaceDetail) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseNamespaceDetail{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseNamespaceDetail) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseNamespaceDetail) ToPlain() elemental.PlainIdentifiable {

	out := NewNamespaceDetail()
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

// DeepCopy returns a deep copy if the SparseNamespaceDetail.
func (o *SparseNamespaceDetail) DeepCopy() *SparseNamespaceDetail {

	if o == nil {
		return nil
	}

	out := &SparseNamespaceDetail{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseNamespaceDetail.
func (o *SparseNamespaceDetail) DeepCopyInto(out *SparseNamespaceDetail) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseNamespaceDetail: %s", err))
	}

	*out = *target.(*SparseNamespaceDetail)
}

type mongoAttributesNamespaceDetail struct {
}
type mongoAttributesSparseNamespaceDetail struct {
}
