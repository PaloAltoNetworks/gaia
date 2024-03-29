// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// FloodRPCIdentity represents the Identity of the object.
var FloodRPCIdentity = elemental.Identity{
	Name:     "floodrpc",
	Category: "floodrpcs",
	Package:  "yeul",
	Private:  true,
}

// FloodRPCsList represents a list of FloodRPCs
type FloodRPCsList []*FloodRPC

// Identity returns the identity of the objects in the list.
func (o FloodRPCsList) Identity() elemental.Identity {

	return FloodRPCIdentity
}

// Copy returns a pointer to a copy the FloodRPCsList.
func (o FloodRPCsList) Copy() elemental.Identifiables {

	out := append(FloodRPCsList{}, o...)
	return &out
}

// Append appends the objects to the a new copy of the FloodRPCsList.
func (o FloodRPCsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(FloodRPCsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*FloodRPC))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o FloodRPCsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o FloodRPCsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the FloodRPCsList converted to SparseFloodRPCsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o FloodRPCsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseFloodRPCsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseFloodRPC)
	}

	return out
}

// Version returns the version of the content.
func (o FloodRPCsList) Version() int {

	return 1
}

// FloodRPC represents the model of a floodrpc
type FloodRPC struct {
	// The parameters needed to create and start a flooder.
	FloodParams []*FloodParam `json:"floodParams" msgpack:"floodParams" bson:"-" mapstructure:"floodParams,omitempty"`

	// The options needed to create nodemakers that are registered with a cached mux.
	NodeMakerConfigs *FloodNodeMakerConfig `json:"nodeMakerConfigs" msgpack:"nodeMakerConfigs" bson:"-" mapstructure:"nodeMakerConfigs,omitempty"`

	// If set, trails will be omitted from the results.
	OptionResultOmitTrails bool `json:"optionResultOmitTrails" msgpack:"optionResultOmitTrails" bson:"-" mapstructure:"optionResultOmitTrails,omitempty"`

	// a unique random string that is used to associate a cached mux nodemaker with
	// successive requests.
	SessionID string `json:"sessionID" msgpack:"sessionID" bson:"-" mapstructure:"sessionID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewFloodRPC returns a new *FloodRPC
func NewFloodRPC() *FloodRPC {

	return &FloodRPC{
		ModelVersion:     1,
		FloodParams:      []*FloodParam{},
		NodeMakerConfigs: NewFloodNodeMakerConfig(),
	}
}

// Identity returns the Identity of the object.
func (o *FloodRPC) Identity() elemental.Identity {

	return FloodRPCIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FloodRPC) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FloodRPC) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FloodRPC) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesFloodRPC{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FloodRPC) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesFloodRPC{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *FloodRPC) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *FloodRPC) BleveType() string {

	return "floodrpc"
}

// DefaultOrder returns the list of default ordering fields.
func (o *FloodRPC) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *FloodRPC) Doc() string {

	return `Starts a flood remote procedural call for a given source/destination/payload
triplet.`
}

func (o *FloodRPC) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *FloodRPC) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseFloodRPC{
			FloodParams:            &o.FloodParams,
			NodeMakerConfigs:       o.NodeMakerConfigs,
			OptionResultOmitTrails: &o.OptionResultOmitTrails,
			SessionID:              &o.SessionID,
		}
	}

	sp := &SparseFloodRPC{}
	for _, f := range fields {
		switch f {
		case "floodParams":
			sp.FloodParams = &(o.FloodParams)
		case "nodeMakerConfigs":
			sp.NodeMakerConfigs = o.NodeMakerConfigs
		case "optionResultOmitTrails":
			sp.OptionResultOmitTrails = &(o.OptionResultOmitTrails)
		case "sessionID":
			sp.SessionID = &(o.SessionID)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseFloodRPC to the object.
func (o *FloodRPC) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseFloodRPC)
	if so.FloodParams != nil {
		o.FloodParams = *so.FloodParams
	}
	if so.NodeMakerConfigs != nil {
		o.NodeMakerConfigs = so.NodeMakerConfigs
	}
	if so.OptionResultOmitTrails != nil {
		o.OptionResultOmitTrails = *so.OptionResultOmitTrails
	}
	if so.SessionID != nil {
		o.SessionID = *so.SessionID
	}
}

// DeepCopy returns a deep copy if the FloodRPC.
func (o *FloodRPC) DeepCopy() *FloodRPC {

	if o == nil {
		return nil
	}

	out := &FloodRPC{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *FloodRPC.
func (o *FloodRPC) DeepCopyInto(out *FloodRPC) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy FloodRPC: %s", err))
	}

	*out = *target.(*FloodRPC)
}

// Validate valides the current information stored into the structure.
func (o *FloodRPC) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if o.NodeMakerConfigs != nil {
		elemental.ResetDefaultForZeroValues(o.NodeMakerConfigs)
		if err := o.NodeMakerConfigs.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateRequiredString("sessionID", o.SessionID); err != nil {
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
func (*FloodRPC) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := FloodRPCAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return FloodRPCLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*FloodRPC) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return FloodRPCAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *FloodRPC) ValueForAttribute(name string) any {

	switch name {
	case "floodParams":
		return o.FloodParams
	case "nodeMakerConfigs":
		return o.NodeMakerConfigs
	case "optionResultOmitTrails":
		return o.OptionResultOmitTrails
	case "sessionID":
		return o.SessionID
	}

	return nil
}

// FloodRPCAttributesMap represents the map of attribute for FloodRPC.
var FloodRPCAttributesMap = map[string]elemental.AttributeSpecification{
	"FloodParams": {
		AllowedChoices: []string{},
		ConvertedName:  "FloodParams",
		Description:    `The parameters needed to create and start a flooder.`,
		Exposed:        true,
		Name:           "floodParams",
		SubType:        "[]floodparam",
		Type:           "external",
	},
	"NodeMakerConfigs": {
		AllowedChoices: []string{},
		ConvertedName:  "NodeMakerConfigs",
		Description:    `The options needed to create nodemakers that are registered with a cached mux.`,
		Exposed:        true,
		Name:           "nodeMakerConfigs",
		SubType:        "floodnodemakerconfig",
		Type:           "ref",
	},
	"OptionResultOmitTrails": {
		AllowedChoices: []string{},
		ConvertedName:  "OptionResultOmitTrails",
		Description:    `If set, trails will be omitted from the results.`,
		Exposed:        true,
		Name:           "optionResultOmitTrails",
		Type:           "boolean",
	},
	"SessionID": {
		AllowedChoices: []string{},
		ConvertedName:  "SessionID",
		Description: `a unique random string that is used to associate a cached mux nodemaker with
successive requests.`,
		Exposed:  true,
		Name:     "sessionID",
		Required: true,
		Type:     "string",
	},
}

// FloodRPCLowerCaseAttributesMap represents the map of attribute for FloodRPC.
var FloodRPCLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"floodparams": {
		AllowedChoices: []string{},
		ConvertedName:  "FloodParams",
		Description:    `The parameters needed to create and start a flooder.`,
		Exposed:        true,
		Name:           "floodParams",
		SubType:        "[]floodparam",
		Type:           "external",
	},
	"nodemakerconfigs": {
		AllowedChoices: []string{},
		ConvertedName:  "NodeMakerConfigs",
		Description:    `The options needed to create nodemakers that are registered with a cached mux.`,
		Exposed:        true,
		Name:           "nodeMakerConfigs",
		SubType:        "floodnodemakerconfig",
		Type:           "ref",
	},
	"optionresultomittrails": {
		AllowedChoices: []string{},
		ConvertedName:  "OptionResultOmitTrails",
		Description:    `If set, trails will be omitted from the results.`,
		Exposed:        true,
		Name:           "optionResultOmitTrails",
		Type:           "boolean",
	},
	"sessionid": {
		AllowedChoices: []string{},
		ConvertedName:  "SessionID",
		Description: `a unique random string that is used to associate a cached mux nodemaker with
successive requests.`,
		Exposed:  true,
		Name:     "sessionID",
		Required: true,
		Type:     "string",
	},
}

// SparseFloodRPCsList represents a list of SparseFloodRPCs
type SparseFloodRPCsList []*SparseFloodRPC

// Identity returns the identity of the objects in the list.
func (o SparseFloodRPCsList) Identity() elemental.Identity {

	return FloodRPCIdentity
}

// Copy returns a pointer to a copy the SparseFloodRPCsList.
func (o SparseFloodRPCsList) Copy() elemental.Identifiables {

	copy := append(SparseFloodRPCsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseFloodRPCsList.
func (o SparseFloodRPCsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseFloodRPCsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseFloodRPC))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseFloodRPCsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseFloodRPCsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseFloodRPCsList converted to FloodRPCsList.
func (o SparseFloodRPCsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseFloodRPCsList) Version() int {

	return 1
}

// SparseFloodRPC represents the sparse version of a floodrpc.
type SparseFloodRPC struct {
	// The parameters needed to create and start a flooder.
	FloodParams *[]*FloodParam `json:"floodParams,omitempty" msgpack:"floodParams,omitempty" bson:"-" mapstructure:"floodParams,omitempty"`

	// The options needed to create nodemakers that are registered with a cached mux.
	NodeMakerConfigs *FloodNodeMakerConfig `json:"nodeMakerConfigs,omitempty" msgpack:"nodeMakerConfigs,omitempty" bson:"-" mapstructure:"nodeMakerConfigs,omitempty"`

	// If set, trails will be omitted from the results.
	OptionResultOmitTrails *bool `json:"optionResultOmitTrails,omitempty" msgpack:"optionResultOmitTrails,omitempty" bson:"-" mapstructure:"optionResultOmitTrails,omitempty"`

	// a unique random string that is used to associate a cached mux nodemaker with
	// successive requests.
	SessionID *string `json:"sessionID,omitempty" msgpack:"sessionID,omitempty" bson:"-" mapstructure:"sessionID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseFloodRPC returns a new  SparseFloodRPC.
func NewSparseFloodRPC() *SparseFloodRPC {
	return &SparseFloodRPC{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseFloodRPC) Identity() elemental.Identity {

	return FloodRPCIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseFloodRPC) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseFloodRPC) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseFloodRPC) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseFloodRPC{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseFloodRPC) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseFloodRPC{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseFloodRPC) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseFloodRPC) ToPlain() elemental.PlainIdentifiable {

	out := NewFloodRPC()
	if o.FloodParams != nil {
		out.FloodParams = *o.FloodParams
	}
	if o.NodeMakerConfigs != nil {
		out.NodeMakerConfigs = o.NodeMakerConfigs
	}
	if o.OptionResultOmitTrails != nil {
		out.OptionResultOmitTrails = *o.OptionResultOmitTrails
	}
	if o.SessionID != nil {
		out.SessionID = *o.SessionID
	}

	return out
}

// DeepCopy returns a deep copy if the SparseFloodRPC.
func (o *SparseFloodRPC) DeepCopy() *SparseFloodRPC {

	if o == nil {
		return nil
	}

	out := &SparseFloodRPC{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseFloodRPC.
func (o *SparseFloodRPC) DeepCopyInto(out *SparseFloodRPC) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseFloodRPC: %s", err))
	}

	*out = *target.(*SparseFloodRPC)
}

type mongoAttributesFloodRPC struct {
}
type mongoAttributesSparseFloodRPC struct {
}
