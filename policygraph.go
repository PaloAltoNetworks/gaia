// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PolicyGraphPolicyTypeValue represents the possible values for attribute "policyType".
type PolicyGraphPolicyTypeValue string

const (
	// PolicyGraphPolicyTypeAuthorization represents the value Authorization.
	PolicyGraphPolicyTypeAuthorization PolicyGraphPolicyTypeValue = "Authorization"

	// PolicyGraphPolicyTypeCombined represents the value Combined.
	PolicyGraphPolicyTypeCombined PolicyGraphPolicyTypeValue = "Combined"

	// PolicyGraphPolicyTypeInfrastructure represents the value Infrastructure.
	PolicyGraphPolicyTypeInfrastructure PolicyGraphPolicyTypeValue = "Infrastructure"
)

// PolicyGraphIdentity represents the Identity of the object.
var PolicyGraphIdentity = elemental.Identity{
	Name:     "policygraph",
	Category: "policygraphs",
	Package:  "yeul",
	Private:  false,
}

// PolicyGraphsList represents a list of PolicyGraphs
type PolicyGraphsList []*PolicyGraph

// Identity returns the identity of the objects in the list.
func (o PolicyGraphsList) Identity() elemental.Identity {

	return PolicyGraphIdentity
}

// Copy returns a pointer to a copy the PolicyGraphsList.
func (o PolicyGraphsList) Copy() elemental.Identifiables {

	copy := append(PolicyGraphsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PolicyGraphsList.
func (o PolicyGraphsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PolicyGraphsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PolicyGraph))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PolicyGraphsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PolicyGraphsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PolicyGraphsList converted to SparsePolicyGraphsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PolicyGraphsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePolicyGraphsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePolicyGraph)
	}

	return out
}

// Version returns the version of the content.
func (o PolicyGraphsList) Version() int {

	return 1
}

// PolicyGraph represents the model of a policygraph
type PolicyGraph struct {
	// The set of tags that a future-activated processing unit will have for which
	// the user wants to evaluate policies and understand its connectivity options.
	PUIdentity []string `json:"PUIdentity" msgpack:"PUIdentity" bson:"-" mapstructure:"PUIdentity,omitempty"`

	// Contains the output of the policy evaluation. It is the same type of
	// dependency map as created by other APIs.
	DependencyMap *DependencyMap `json:"dependencyMap" msgpack:"dependencyMap" bson:"-" mapstructure:"dependencyMap,omitempty"`

	// Identifies the type of policy that should be analyzed: `Authorization`(default),
	// `Infrastructure`, or `Combined`.
	PolicyType PolicyGraphPolicyTypeValue `json:"policyType" msgpack:"policyType" bson:"-" mapstructure:"policyType,omitempty"`

	// Contains the tag expression that a processing unit must match in order
	// to evaluate policy for it.
	Selectors [][]string `json:"selectors" msgpack:"selectors" bson:"-" mapstructure:"selectors,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPolicyGraph returns a new *PolicyGraph
func NewPolicyGraph() *PolicyGraph {

	return &PolicyGraph{
		ModelVersion:  1,
		PUIdentity:    []string{},
		DependencyMap: NewDependencyMap(),
		PolicyType:    PolicyGraphPolicyTypeAuthorization,
		Selectors:     [][]string{},
	}
}

// Identity returns the Identity of the object.
func (o *PolicyGraph) Identity() elemental.Identity {

	return PolicyGraphIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PolicyGraph) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PolicyGraph) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PolicyGraph) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPolicyGraph{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PolicyGraph) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPolicyGraph{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *PolicyGraph) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *PolicyGraph) BleveType() string {

	return "policygraph"
}

// DefaultOrder returns the list of default ordering fields.
func (o *PolicyGraph) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PolicyGraph) Doc() string {

	return `Returns a data structure representing the policy graph of all selected
processing units and their possible connectivity based on the current policies
associated with the namespace. Users can define a selector of processing units
in which they are interested or define the identity tags of a virtual processing
unit that is not yet activated.`
}

func (o *PolicyGraph) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PolicyGraph) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePolicyGraph{
			PUIdentity:    &o.PUIdentity,
			DependencyMap: o.DependencyMap,
			PolicyType:    &o.PolicyType,
			Selectors:     &o.Selectors,
		}
	}

	sp := &SparsePolicyGraph{}
	for _, f := range fields {
		switch f {
		case "PUIdentity":
			sp.PUIdentity = &(o.PUIdentity)
		case "dependencyMap":
			sp.DependencyMap = o.DependencyMap
		case "policyType":
			sp.PolicyType = &(o.PolicyType)
		case "selectors":
			sp.Selectors = &(o.Selectors)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePolicyGraph to the object.
func (o *PolicyGraph) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePolicyGraph)
	if so.PUIdentity != nil {
		o.PUIdentity = *so.PUIdentity
	}
	if so.DependencyMap != nil {
		o.DependencyMap = so.DependencyMap
	}
	if so.PolicyType != nil {
		o.PolicyType = *so.PolicyType
	}
	if so.Selectors != nil {
		o.Selectors = *so.Selectors
	}
}

// DeepCopy returns a deep copy if the PolicyGraph.
func (o *PolicyGraph) DeepCopy() *PolicyGraph {

	if o == nil {
		return nil
	}

	out := &PolicyGraph{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PolicyGraph.
func (o *PolicyGraph) DeepCopyInto(out *PolicyGraph) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PolicyGraph: %s", err))
	}

	*out = *target.(*PolicyGraph)
}

// Validate valides the current information stored into the structure.
func (o *PolicyGraph) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if o.DependencyMap != nil {
		elemental.ResetDefaultForZeroValues(o.DependencyMap)
		if err := o.DependencyMap.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateStringInList("policyType", string(o.PolicyType), []string{"Authorization", "Infrastructure", "Combined"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsExpression("selectors", o.Selectors); err != nil {
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
func (*PolicyGraph) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PolicyGraphAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PolicyGraphLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PolicyGraph) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PolicyGraphAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PolicyGraph) ValueForAttribute(name string) any {

	switch name {
	case "PUIdentity":
		return o.PUIdentity
	case "dependencyMap":
		return o.DependencyMap
	case "policyType":
		return o.PolicyType
	case "selectors":
		return o.Selectors
	}

	return nil
}

// PolicyGraphAttributesMap represents the map of attribute for PolicyGraph.
var PolicyGraphAttributesMap = map[string]elemental.AttributeSpecification{
	"PUIdentity": {
		AllowedChoices: []string{},
		ConvertedName:  "PUIdentity",
		Description: `The set of tags that a future-activated processing unit will have for which
the user wants to evaluate policies and understand its connectivity options.`,
		Exposed: true,
		Name:    "PUIdentity",
		SubType: "string",
		Type:    "list",
	},
	"DependencyMap": {
		AllowedChoices: []string{},
		ConvertedName:  "DependencyMap",
		Description: `Contains the output of the policy evaluation. It is the same type of
dependency map as created by other APIs.`,
		Exposed: true,
		Name:    "dependencyMap",
		SubType: "dependencymap",
		Type:    "ref",
	},
	"PolicyType": {
		AllowedChoices: []string{"Authorization", "Infrastructure", "Combined"},
		ConvertedName:  "PolicyType",
		DefaultValue:   PolicyGraphPolicyTypeAuthorization,
		Description: `Identifies the type of policy that should be analyzed: ` + "`" + `Authorization` + "`" + `(default),
` + "`" + `Infrastructure` + "`" + `, or ` + "`" + `Combined` + "`" + `.`,
		Exposed: true,
		Name:    "policyType",
		Type:    "enum",
	},
	"Selectors": {
		AllowedChoices: []string{},
		ConvertedName:  "Selectors",
		Description: `Contains the tag expression that a processing unit must match in order
to evaluate policy for it.`,
		Exposed: true,
		Name:    "selectors",
		SubType: "[][]string",
		Type:    "external",
	},
}

// PolicyGraphLowerCaseAttributesMap represents the map of attribute for PolicyGraph.
var PolicyGraphLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"puidentity": {
		AllowedChoices: []string{},
		ConvertedName:  "PUIdentity",
		Description: `The set of tags that a future-activated processing unit will have for which
the user wants to evaluate policies and understand its connectivity options.`,
		Exposed: true,
		Name:    "PUIdentity",
		SubType: "string",
		Type:    "list",
	},
	"dependencymap": {
		AllowedChoices: []string{},
		ConvertedName:  "DependencyMap",
		Description: `Contains the output of the policy evaluation. It is the same type of
dependency map as created by other APIs.`,
		Exposed: true,
		Name:    "dependencyMap",
		SubType: "dependencymap",
		Type:    "ref",
	},
	"policytype": {
		AllowedChoices: []string{"Authorization", "Infrastructure", "Combined"},
		ConvertedName:  "PolicyType",
		DefaultValue:   PolicyGraphPolicyTypeAuthorization,
		Description: `Identifies the type of policy that should be analyzed: ` + "`" + `Authorization` + "`" + `(default),
` + "`" + `Infrastructure` + "`" + `, or ` + "`" + `Combined` + "`" + `.`,
		Exposed: true,
		Name:    "policyType",
		Type:    "enum",
	},
	"selectors": {
		AllowedChoices: []string{},
		ConvertedName:  "Selectors",
		Description: `Contains the tag expression that a processing unit must match in order
to evaluate policy for it.`,
		Exposed: true,
		Name:    "selectors",
		SubType: "[][]string",
		Type:    "external",
	},
}

// SparsePolicyGraphsList represents a list of SparsePolicyGraphs
type SparsePolicyGraphsList []*SparsePolicyGraph

// Identity returns the identity of the objects in the list.
func (o SparsePolicyGraphsList) Identity() elemental.Identity {

	return PolicyGraphIdentity
}

// Copy returns a pointer to a copy the SparsePolicyGraphsList.
func (o SparsePolicyGraphsList) Copy() elemental.Identifiables {

	copy := append(SparsePolicyGraphsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePolicyGraphsList.
func (o SparsePolicyGraphsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePolicyGraphsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePolicyGraph))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePolicyGraphsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePolicyGraphsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePolicyGraphsList converted to PolicyGraphsList.
func (o SparsePolicyGraphsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePolicyGraphsList) Version() int {

	return 1
}

// SparsePolicyGraph represents the sparse version of a policygraph.
type SparsePolicyGraph struct {
	// The set of tags that a future-activated processing unit will have for which
	// the user wants to evaluate policies and understand its connectivity options.
	PUIdentity *[]string `json:"PUIdentity,omitempty" msgpack:"PUIdentity,omitempty" bson:"-" mapstructure:"PUIdentity,omitempty"`

	// Contains the output of the policy evaluation. It is the same type of
	// dependency map as created by other APIs.
	DependencyMap *DependencyMap `json:"dependencyMap,omitempty" msgpack:"dependencyMap,omitempty" bson:"-" mapstructure:"dependencyMap,omitempty"`

	// Identifies the type of policy that should be analyzed: `Authorization`(default),
	// `Infrastructure`, or `Combined`.
	PolicyType *PolicyGraphPolicyTypeValue `json:"policyType,omitempty" msgpack:"policyType,omitempty" bson:"-" mapstructure:"policyType,omitempty"`

	// Contains the tag expression that a processing unit must match in order
	// to evaluate policy for it.
	Selectors *[][]string `json:"selectors,omitempty" msgpack:"selectors,omitempty" bson:"-" mapstructure:"selectors,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePolicyGraph returns a new  SparsePolicyGraph.
func NewSparsePolicyGraph() *SparsePolicyGraph {
	return &SparsePolicyGraph{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePolicyGraph) Identity() elemental.Identity {

	return PolicyGraphIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePolicyGraph) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePolicyGraph) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePolicyGraph) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePolicyGraph{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePolicyGraph) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePolicyGraph{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparsePolicyGraph) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePolicyGraph) ToPlain() elemental.PlainIdentifiable {

	out := NewPolicyGraph()
	if o.PUIdentity != nil {
		out.PUIdentity = *o.PUIdentity
	}
	if o.DependencyMap != nil {
		out.DependencyMap = o.DependencyMap
	}
	if o.PolicyType != nil {
		out.PolicyType = *o.PolicyType
	}
	if o.Selectors != nil {
		out.Selectors = *o.Selectors
	}

	return out
}

// DeepCopy returns a deep copy if the SparsePolicyGraph.
func (o *SparsePolicyGraph) DeepCopy() *SparsePolicyGraph {

	if o == nil {
		return nil
	}

	out := &SparsePolicyGraph{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePolicyGraph.
func (o *SparsePolicyGraph) DeepCopyInto(out *SparsePolicyGraph) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePolicyGraph: %s", err))
	}

	*out = *target.(*SparsePolicyGraph)
}

type mongoAttributesPolicyGraph struct {
}
type mongoAttributesSparsePolicyGraph struct {
}
