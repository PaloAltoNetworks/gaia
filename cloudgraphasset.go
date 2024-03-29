// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudGraphAssetDirectionValue represents the possible values for attribute "direction".
type CloudGraphAssetDirectionValue string

const (
	// CloudGraphAssetDirectionInbound represents the value Inbound.
	CloudGraphAssetDirectionInbound CloudGraphAssetDirectionValue = "Inbound"

	// CloudGraphAssetDirectionOutbound represents the value Outbound.
	CloudGraphAssetDirectionOutbound CloudGraphAssetDirectionValue = "Outbound"
)

// CloudGraphAssetIdentity represents the Identity of the object.
var CloudGraphAssetIdentity = elemental.Identity{
	Name:     "cloudgraphasset",
	Category: "cloudgraphassets",
	Package:  "yeul",
	Private:  false,
}

// CloudGraphAssetsList represents a list of CloudGraphAssets
type CloudGraphAssetsList []*CloudGraphAsset

// Identity returns the identity of the objects in the list.
func (o CloudGraphAssetsList) Identity() elemental.Identity {

	return CloudGraphAssetIdentity
}

// Copy returns a pointer to a copy the CloudGraphAssetsList.
func (o CloudGraphAssetsList) Copy() elemental.Identifiables {

	out := append(CloudGraphAssetsList{}, o...)
	return &out
}

// Append appends the objects to the a new copy of the CloudGraphAssetsList.
func (o CloudGraphAssetsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudGraphAssetsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudGraphAsset))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudGraphAssetsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudGraphAssetsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CloudGraphAssetsList converted to SparseCloudGraphAssetsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudGraphAssetsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudGraphAssetsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudGraphAsset)
	}

	return out
}

// Version returns the version of the content.
func (o CloudGraphAssetsList) Version() int {

	return 1
}

// CloudGraphAsset represents the model of a cloudgraphasset
type CloudGraphAsset struct {
	// The set of nodes and paths from the given source to the given destination.
	CloudGraphs map[string]*CloudGraph `json:"cloudGraphs" msgpack:"cloudGraphs" bson:"-" mapstructure:"cloudGraphs,omitempty"`

	// Direction of the network path.
	Direction CloudGraphAssetDirectionValue `json:"direction" msgpack:"direction" bson:"-" mapstructure:"direction,omitempty"`

	// The error message if encountered any.
	Errors []string `json:"errors,omitempty" msgpack:"errors,omitempty" bson:"-" mapstructure:"errors,omitempty"`

	// Result of the graph execution timestamp.
	ExecutionTimestamp time.Time `json:"executionTimestamp" msgpack:"executionTimestamp" bson:"-" mapstructure:"executionTimestamp,omitempty"`

	// Prisma Cloud Policy ID.
	PrismaCloudPolicyID string `json:"prismaCloudPolicyID,omitempty" msgpack:"prismaCloudPolicyID,omitempty" bson:"-" mapstructure:"prismaCloudPolicyID,omitempty"`

	// Prisma Cloud Unified Asset IDs.
	UnifiedAssetIDs []string `json:"unifiedAssetIDs,omitempty" msgpack:"unifiedAssetIDs,omitempty" bson:"-" mapstructure:"unifiedAssetIDs,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudGraphAsset returns a new *CloudGraphAsset
func NewCloudGraphAsset() *CloudGraphAsset {

	return &CloudGraphAsset{
		ModelVersion:    1,
		CloudGraphs:     map[string]*CloudGraph{},
		Direction:       CloudGraphAssetDirectionOutbound,
		Errors:          []string{},
		UnifiedAssetIDs: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *CloudGraphAsset) Identity() elemental.Identity {

	return CloudGraphAssetIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudGraphAsset) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudGraphAsset) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudGraphAsset) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudGraphAsset{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudGraphAsset) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudGraphAsset{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudGraphAsset) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudGraphAsset) BleveType() string {

	return "cloudgraphasset"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudGraphAsset) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CloudGraphAsset) Doc() string {

	return `Returns a data structure representing the graph of all cloud nodes and their
connections for the given source unified asset IDs.`
}

func (o *CloudGraphAsset) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudGraphAsset) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudGraphAsset{
			CloudGraphs:         &o.CloudGraphs,
			Direction:           &o.Direction,
			Errors:              &o.Errors,
			ExecutionTimestamp:  &o.ExecutionTimestamp,
			PrismaCloudPolicyID: &o.PrismaCloudPolicyID,
			UnifiedAssetIDs:     &o.UnifiedAssetIDs,
		}
	}

	sp := &SparseCloudGraphAsset{}
	for _, f := range fields {
		switch f {
		case "cloudGraphs":
			sp.CloudGraphs = &(o.CloudGraphs)
		case "direction":
			sp.Direction = &(o.Direction)
		case "errors":
			sp.Errors = &(o.Errors)
		case "executionTimestamp":
			sp.ExecutionTimestamp = &(o.ExecutionTimestamp)
		case "prismaCloudPolicyID":
			sp.PrismaCloudPolicyID = &(o.PrismaCloudPolicyID)
		case "unifiedAssetIDs":
			sp.UnifiedAssetIDs = &(o.UnifiedAssetIDs)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCloudGraphAsset to the object.
func (o *CloudGraphAsset) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudGraphAsset)
	if so.CloudGraphs != nil {
		o.CloudGraphs = *so.CloudGraphs
	}
	if so.Direction != nil {
		o.Direction = *so.Direction
	}
	if so.Errors != nil {
		o.Errors = *so.Errors
	}
	if so.ExecutionTimestamp != nil {
		o.ExecutionTimestamp = *so.ExecutionTimestamp
	}
	if so.PrismaCloudPolicyID != nil {
		o.PrismaCloudPolicyID = *so.PrismaCloudPolicyID
	}
	if so.UnifiedAssetIDs != nil {
		o.UnifiedAssetIDs = *so.UnifiedAssetIDs
	}
}

// DeepCopy returns a deep copy if the CloudGraphAsset.
func (o *CloudGraphAsset) DeepCopy() *CloudGraphAsset {

	if o == nil {
		return nil
	}

	out := &CloudGraphAsset{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudGraphAsset.
func (o *CloudGraphAsset) DeepCopyInto(out *CloudGraphAsset) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudGraphAsset: %s", err))
	}

	*out = *target.(*CloudGraphAsset)
}

// Validate valides the current information stored into the structure.
func (o *CloudGraphAsset) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.CloudGraphs {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateStringInList("direction", string(o.Direction), []string{"Inbound", "Outbound"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("prismaCloudPolicyID", o.PrismaCloudPolicyID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredExternal("unifiedAssetIDs", o.UnifiedAssetIDs); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := ValidateMaxUnifiedAssetIDs("unifiedAssetIDs", o.UnifiedAssetIDs); err != nil {
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
func (*CloudGraphAsset) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudGraphAssetAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudGraphAssetLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudGraphAsset) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudGraphAssetAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudGraphAsset) ValueForAttribute(name string) any {

	switch name {
	case "cloudGraphs":
		return o.CloudGraphs
	case "direction":
		return o.Direction
	case "errors":
		return o.Errors
	case "executionTimestamp":
		return o.ExecutionTimestamp
	case "prismaCloudPolicyID":
		return o.PrismaCloudPolicyID
	case "unifiedAssetIDs":
		return o.UnifiedAssetIDs
	}

	return nil
}

// CloudGraphAssetAttributesMap represents the map of attribute for CloudGraphAsset.
var CloudGraphAssetAttributesMap = map[string]elemental.AttributeSpecification{
	"CloudGraphs": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CloudGraphs",
		Description:    `The set of nodes and paths from the given source to the given destination.`,
		Exposed:        true,
		Name:           "cloudGraphs",
		ReadOnly:       true,
		SubType:        "cloudgraph",
		Type:           "refMap",
	},
	"Direction": {
		AllowedChoices: []string{"Inbound", "Outbound"},
		ConvertedName:  "Direction",
		DefaultValue:   CloudGraphAssetDirectionOutbound,
		Description:    `Direction of the network path.`,
		Exposed:        true,
		Name:           "direction",
		Type:           "enum",
	},
	"Errors": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Errors",
		Description:    `The error message if encountered any.`,
		Exposed:        true,
		Name:           "errors",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"ExecutionTimestamp": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ExecutionTimestamp",
		Description:    `Result of the graph execution timestamp.`,
		Exposed:        true,
		Name:           "executionTimestamp",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "time",
	},
	"PrismaCloudPolicyID": {
		AllowedChoices: []string{},
		ConvertedName:  "PrismaCloudPolicyID",
		Description:    `Prisma Cloud Policy ID.`,
		Exposed:        true,
		Name:           "prismaCloudPolicyID",
		Required:       true,
		Type:           "string",
	},
	"UnifiedAssetIDs": {
		AllowedChoices: []string{},
		ConvertedName:  "UnifiedAssetIDs",
		Description:    `Prisma Cloud Unified Asset IDs.`,
		Exposed:        true,
		Name:           "unifiedAssetIDs",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
}

// CloudGraphAssetLowerCaseAttributesMap represents the map of attribute for CloudGraphAsset.
var CloudGraphAssetLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"cloudgraphs": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CloudGraphs",
		Description:    `The set of nodes and paths from the given source to the given destination.`,
		Exposed:        true,
		Name:           "cloudGraphs",
		ReadOnly:       true,
		SubType:        "cloudgraph",
		Type:           "refMap",
	},
	"direction": {
		AllowedChoices: []string{"Inbound", "Outbound"},
		ConvertedName:  "Direction",
		DefaultValue:   CloudGraphAssetDirectionOutbound,
		Description:    `Direction of the network path.`,
		Exposed:        true,
		Name:           "direction",
		Type:           "enum",
	},
	"errors": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Errors",
		Description:    `The error message if encountered any.`,
		Exposed:        true,
		Name:           "errors",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"executiontimestamp": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ExecutionTimestamp",
		Description:    `Result of the graph execution timestamp.`,
		Exposed:        true,
		Name:           "executionTimestamp",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "time",
	},
	"prismacloudpolicyid": {
		AllowedChoices: []string{},
		ConvertedName:  "PrismaCloudPolicyID",
		Description:    `Prisma Cloud Policy ID.`,
		Exposed:        true,
		Name:           "prismaCloudPolicyID",
		Required:       true,
		Type:           "string",
	},
	"unifiedassetids": {
		AllowedChoices: []string{},
		ConvertedName:  "UnifiedAssetIDs",
		Description:    `Prisma Cloud Unified Asset IDs.`,
		Exposed:        true,
		Name:           "unifiedAssetIDs",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
}

// SparseCloudGraphAssetsList represents a list of SparseCloudGraphAssets
type SparseCloudGraphAssetsList []*SparseCloudGraphAsset

// Identity returns the identity of the objects in the list.
func (o SparseCloudGraphAssetsList) Identity() elemental.Identity {

	return CloudGraphAssetIdentity
}

// Copy returns a pointer to a copy the SparseCloudGraphAssetsList.
func (o SparseCloudGraphAssetsList) Copy() elemental.Identifiables {

	copy := append(SparseCloudGraphAssetsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudGraphAssetsList.
func (o SparseCloudGraphAssetsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudGraphAssetsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudGraphAsset))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudGraphAssetsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudGraphAssetsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCloudGraphAssetsList converted to CloudGraphAssetsList.
func (o SparseCloudGraphAssetsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudGraphAssetsList) Version() int {

	return 1
}

// SparseCloudGraphAsset represents the sparse version of a cloudgraphasset.
type SparseCloudGraphAsset struct {
	// The set of nodes and paths from the given source to the given destination.
	CloudGraphs *map[string]*CloudGraph `json:"cloudGraphs,omitempty" msgpack:"cloudGraphs,omitempty" bson:"-" mapstructure:"cloudGraphs,omitempty"`

	// Direction of the network path.
	Direction *CloudGraphAssetDirectionValue `json:"direction,omitempty" msgpack:"direction,omitempty" bson:"-" mapstructure:"direction,omitempty"`

	// The error message if encountered any.
	Errors *[]string `json:"errors,omitempty" msgpack:"errors,omitempty" bson:"-" mapstructure:"errors,omitempty"`

	// Result of the graph execution timestamp.
	ExecutionTimestamp *time.Time `json:"executionTimestamp,omitempty" msgpack:"executionTimestamp,omitempty" bson:"-" mapstructure:"executionTimestamp,omitempty"`

	// Prisma Cloud Policy ID.
	PrismaCloudPolicyID *string `json:"prismaCloudPolicyID,omitempty" msgpack:"prismaCloudPolicyID,omitempty" bson:"-" mapstructure:"prismaCloudPolicyID,omitempty"`

	// Prisma Cloud Unified Asset IDs.
	UnifiedAssetIDs *[]string `json:"unifiedAssetIDs,omitempty" msgpack:"unifiedAssetIDs,omitempty" bson:"-" mapstructure:"unifiedAssetIDs,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCloudGraphAsset returns a new  SparseCloudGraphAsset.
func NewSparseCloudGraphAsset() *SparseCloudGraphAsset {
	return &SparseCloudGraphAsset{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudGraphAsset) Identity() elemental.Identity {

	return CloudGraphAssetIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudGraphAsset) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudGraphAsset) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudGraphAsset) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudGraphAsset{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudGraphAsset) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudGraphAsset{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCloudGraphAsset) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudGraphAsset) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudGraphAsset()
	if o.CloudGraphs != nil {
		out.CloudGraphs = *o.CloudGraphs
	}
	if o.Direction != nil {
		out.Direction = *o.Direction
	}
	if o.Errors != nil {
		out.Errors = *o.Errors
	}
	if o.ExecutionTimestamp != nil {
		out.ExecutionTimestamp = *o.ExecutionTimestamp
	}
	if o.PrismaCloudPolicyID != nil {
		out.PrismaCloudPolicyID = *o.PrismaCloudPolicyID
	}
	if o.UnifiedAssetIDs != nil {
		out.UnifiedAssetIDs = *o.UnifiedAssetIDs
	}

	return out
}

// DeepCopy returns a deep copy if the SparseCloudGraphAsset.
func (o *SparseCloudGraphAsset) DeepCopy() *SparseCloudGraphAsset {

	if o == nil {
		return nil
	}

	out := &SparseCloudGraphAsset{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudGraphAsset.
func (o *SparseCloudGraphAsset) DeepCopyInto(out *SparseCloudGraphAsset) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudGraphAsset: %s", err))
	}

	*out = *target.(*SparseCloudGraphAsset)
}

type mongoAttributesCloudGraphAsset struct {
}
type mongoAttributesSparseCloudGraphAsset struct {
}
