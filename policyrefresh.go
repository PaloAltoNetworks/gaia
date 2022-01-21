package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PolicyRefreshIdentity represents the Identity of the object.
var PolicyRefreshIdentity = elemental.Identity{
	Name:     "policyrefresh",
	Category: "policyrefreshs",
	Package:  "squall",
	Private:  false,
}

// PolicyRefreshsList represents a list of PolicyRefreshs
type PolicyRefreshsList []*PolicyRefresh

// Identity returns the identity of the objects in the list.
func (o PolicyRefreshsList) Identity() elemental.Identity {

	return PolicyRefreshIdentity
}

// Copy returns a pointer to a copy the PolicyRefreshsList.
func (o PolicyRefreshsList) Copy() elemental.Identifiables {

	copy := append(PolicyRefreshsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PolicyRefreshsList.
func (o PolicyRefreshsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PolicyRefreshsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PolicyRefresh))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PolicyRefreshsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PolicyRefreshsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PolicyRefreshsList converted to SparsePolicyRefreshsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PolicyRefreshsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePolicyRefreshsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePolicyRefresh)
	}

	return out
}

// Version returns the version of the content.
func (o PolicyRefreshsList) Version() int {

	return 1
}

// PolicyRefresh represents the model of a policyrefresh
type PolicyRefresh struct {
	// Contains the original ID of the updated object.
	SourceID string `json:"sourceID" msgpack:"sourceID" bson:"sourceid" mapstructure:"sourceID,omitempty"`

	// Contains the original namespace of the updated object.
	SourceNamespace string `json:"sourceNamespace" msgpack:"sourceNamespace" bson:"sourcenamespace" mapstructure:"sourceNamespace,omitempty"`

	// Contains the policy type that is affected.
	Type string `json:"type" msgpack:"type" bson:"type" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPolicyRefresh returns a new *PolicyRefresh
func NewPolicyRefresh() *PolicyRefresh {

	return &PolicyRefresh{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *PolicyRefresh) Identity() elemental.Identity {

	return PolicyRefreshIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PolicyRefresh) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PolicyRefresh) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PolicyRefresh) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPolicyRefresh{}

	s.SourceID = o.SourceID
	s.SourceNamespace = o.SourceNamespace
	s.Type = o.Type

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PolicyRefresh) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPolicyRefresh{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.SourceID = s.SourceID
	o.SourceNamespace = s.SourceNamespace
	o.Type = s.Type

	return nil
}

// Version returns the hardcoded version of the model.
func (o *PolicyRefresh) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *PolicyRefresh) BleveType() string {

	return "policyrefresh"
}

// DefaultOrder returns the list of default ordering fields.
func (o *PolicyRefresh) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PolicyRefresh) Doc() string {

	return `Sent to a client as a push event when a policy refresh is needed on their side.`
}

func (o *PolicyRefresh) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PolicyRefresh) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePolicyRefresh{
			SourceID:        &o.SourceID,
			SourceNamespace: &o.SourceNamespace,
			Type:            &o.Type,
		}
	}

	sp := &SparsePolicyRefresh{}
	for _, f := range fields {
		switch f {
		case "sourceID":
			sp.SourceID = &(o.SourceID)
		case "sourceNamespace":
			sp.SourceNamespace = &(o.SourceNamespace)
		case "type":
			sp.Type = &(o.Type)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePolicyRefresh to the object.
func (o *PolicyRefresh) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePolicyRefresh)
	if so.SourceID != nil {
		o.SourceID = *so.SourceID
	}
	if so.SourceNamespace != nil {
		o.SourceNamespace = *so.SourceNamespace
	}
	if so.Type != nil {
		o.Type = *so.Type
	}
}

// DeepCopy returns a deep copy if the PolicyRefresh.
func (o *PolicyRefresh) DeepCopy() *PolicyRefresh {

	if o == nil {
		return nil
	}

	out := &PolicyRefresh{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PolicyRefresh.
func (o *PolicyRefresh) DeepCopyInto(out *PolicyRefresh) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PolicyRefresh: %s", err))
	}

	*out = *target.(*PolicyRefresh)
}

// Validate valides the current information stored into the structure.
func (o *PolicyRefresh) Validate() error {

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
func (*PolicyRefresh) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PolicyRefreshAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PolicyRefreshLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PolicyRefresh) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PolicyRefreshAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PolicyRefresh) ValueForAttribute(name string) interface{} {

	switch name {
	case "sourceID":
		return o.SourceID
	case "sourceNamespace":
		return o.SourceNamespace
	case "type":
		return o.Type
	}

	return nil
}

// PolicyRefreshAttributesMap represents the map of attribute for PolicyRefresh.
var PolicyRefreshAttributesMap = map[string]elemental.AttributeSpecification{
	"SourceID": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceID",
		Description:    `Contains the original ID of the updated object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "sourceID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"SourceNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceNamespace",
		Description:    `Contains the original namespace of the updated object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "sourceNamespace",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Type": {
		AllowedChoices: []string{},
		ConvertedName:  "Type",
		Description:    `Contains the policy type that is affected.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "type",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}

// PolicyRefreshLowerCaseAttributesMap represents the map of attribute for PolicyRefresh.
var PolicyRefreshLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"sourceid": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceID",
		Description:    `Contains the original ID of the updated object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "sourceID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"sourcenamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceNamespace",
		Description:    `Contains the original namespace of the updated object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "sourceNamespace",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"type": {
		AllowedChoices: []string{},
		ConvertedName:  "Type",
		Description:    `Contains the policy type that is affected.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "type",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}

// SparsePolicyRefreshsList represents a list of SparsePolicyRefreshs
type SparsePolicyRefreshsList []*SparsePolicyRefresh

// Identity returns the identity of the objects in the list.
func (o SparsePolicyRefreshsList) Identity() elemental.Identity {

	return PolicyRefreshIdentity
}

// Copy returns a pointer to a copy the SparsePolicyRefreshsList.
func (o SparsePolicyRefreshsList) Copy() elemental.Identifiables {

	copy := append(SparsePolicyRefreshsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePolicyRefreshsList.
func (o SparsePolicyRefreshsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePolicyRefreshsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePolicyRefresh))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePolicyRefreshsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePolicyRefreshsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePolicyRefreshsList converted to PolicyRefreshsList.
func (o SparsePolicyRefreshsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePolicyRefreshsList) Version() int {

	return 1
}

// SparsePolicyRefresh represents the sparse version of a policyrefresh.
type SparsePolicyRefresh struct {
	// Contains the original ID of the updated object.
	SourceID *string `json:"sourceID,omitempty" msgpack:"sourceID,omitempty" bson:"sourceid,omitempty" mapstructure:"sourceID,omitempty"`

	// Contains the original namespace of the updated object.
	SourceNamespace *string `json:"sourceNamespace,omitempty" msgpack:"sourceNamespace,omitempty" bson:"sourcenamespace,omitempty" mapstructure:"sourceNamespace,omitempty"`

	// Contains the policy type that is affected.
	Type *string `json:"type,omitempty" msgpack:"type,omitempty" bson:"type,omitempty" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePolicyRefresh returns a new  SparsePolicyRefresh.
func NewSparsePolicyRefresh() *SparsePolicyRefresh {
	return &SparsePolicyRefresh{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePolicyRefresh) Identity() elemental.Identity {

	return PolicyRefreshIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePolicyRefresh) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePolicyRefresh) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePolicyRefresh) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePolicyRefresh{}

	if o.SourceID != nil {
		s.SourceID = o.SourceID
	}
	if o.SourceNamespace != nil {
		s.SourceNamespace = o.SourceNamespace
	}
	if o.Type != nil {
		s.Type = o.Type
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePolicyRefresh) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePolicyRefresh{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.SourceID != nil {
		o.SourceID = s.SourceID
	}
	if s.SourceNamespace != nil {
		o.SourceNamespace = s.SourceNamespace
	}
	if s.Type != nil {
		o.Type = s.Type
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparsePolicyRefresh) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePolicyRefresh) ToPlain() elemental.PlainIdentifiable {

	out := NewPolicyRefresh()
	if o.SourceID != nil {
		out.SourceID = *o.SourceID
	}
	if o.SourceNamespace != nil {
		out.SourceNamespace = *o.SourceNamespace
	}
	if o.Type != nil {
		out.Type = *o.Type
	}

	return out
}

// DeepCopy returns a deep copy if the SparsePolicyRefresh.
func (o *SparsePolicyRefresh) DeepCopy() *SparsePolicyRefresh {

	if o == nil {
		return nil
	}

	out := &SparsePolicyRefresh{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePolicyRefresh.
func (o *SparsePolicyRefresh) DeepCopyInto(out *SparsePolicyRefresh) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePolicyRefresh: %s", err))
	}

	*out = *target.(*SparsePolicyRefresh)
}

type mongoAttributesPolicyRefresh struct {
	SourceID        string `bson:"sourceid"`
	SourceNamespace string `bson:"sourcenamespace"`
	Type            string `bson:"type"`
}
type mongoAttributesSparsePolicyRefresh struct {
	SourceID        *string `bson:"sourceid,omitempty"`
	SourceNamespace *string `bson:"sourcenamespace,omitempty"`
	Type            *string `bson:"type,omitempty"`
}
