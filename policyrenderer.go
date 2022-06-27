package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PolicyRendererProcessModeValue represents the possible values for attribute "processMode".
type PolicyRendererProcessModeValue string

const (
	// PolicyRendererProcessModeObject represents the value Object.
	PolicyRendererProcessModeObject PolicyRendererProcessModeValue = "Object"

	// PolicyRendererProcessModeSubject represents the value Subject.
	PolicyRendererProcessModeSubject PolicyRendererProcessModeValue = "Subject"
)

// PolicyRendererTypeValue represents the possible values for attribute "type".
type PolicyRendererTypeValue string

const (
	// PolicyRendererTypeAPIAuthorization represents the value APIAuthorization.
	PolicyRendererTypeAPIAuthorization PolicyRendererTypeValue = "APIAuthorization"

	// PolicyRendererTypeEnforcerProfile represents the value EnforcerProfile.
	PolicyRendererTypeEnforcerProfile PolicyRendererTypeValue = "EnforcerProfile"

	// PolicyRendererTypeFile represents the value File.
	PolicyRendererTypeFile PolicyRendererTypeValue = "File"

	// PolicyRendererTypeHook represents the value Hook.
	PolicyRendererTypeHook PolicyRendererTypeValue = "Hook"

	// PolicyRendererTypeInfrastructure represents the value Infrastructure.
	PolicyRendererTypeInfrastructure PolicyRendererTypeValue = "Infrastructure"

	// PolicyRendererTypeNamespaceMapping represents the value NamespaceMapping.
	PolicyRendererTypeNamespaceMapping PolicyRendererTypeValue = "NamespaceMapping"

	// PolicyRendererTypeNetwork represents the value Network.
	PolicyRendererTypeNetwork PolicyRendererTypeValue = "Network"

	// PolicyRendererTypeProcessingUnit represents the value ProcessingUnit.
	PolicyRendererTypeProcessingUnit PolicyRendererTypeValue = "ProcessingUnit"

	// PolicyRendererTypeQuota represents the value Quota.
	PolicyRendererTypeQuota PolicyRendererTypeValue = "Quota"

	// PolicyRendererTypeSSHAuthorization represents the value SSHAuthorization.
	PolicyRendererTypeSSHAuthorization PolicyRendererTypeValue = "SSHAuthorization"

	// PolicyRendererTypeSyscall represents the value Syscall.
	PolicyRendererTypeSyscall PolicyRendererTypeValue = "Syscall"

	// PolicyRendererTypeTokenScope represents the value TokenScope.
	PolicyRendererTypeTokenScope PolicyRendererTypeValue = "TokenScope"

	// PolicyRendererTypeUserAccess represents the value UserAccess.
	PolicyRendererTypeUserAccess PolicyRendererTypeValue = "UserAccess"
)

// PolicyRendererIdentity represents the Identity of the object.
var PolicyRendererIdentity = elemental.Identity{
	Name:     "policyrenderer",
	Category: "policyrenderers",
	Package:  "squall",
	Private:  false,
}

// PolicyRenderersList represents a list of PolicyRenderers
type PolicyRenderersList []*PolicyRenderer

// Identity returns the identity of the objects in the list.
func (o PolicyRenderersList) Identity() elemental.Identity {

	return PolicyRendererIdentity
}

// Copy returns a pointer to a copy the PolicyRenderersList.
func (o PolicyRenderersList) Copy() elemental.Identifiables {

	copy := append(PolicyRenderersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PolicyRenderersList.
func (o PolicyRenderersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PolicyRenderersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PolicyRenderer))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PolicyRenderersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PolicyRenderersList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PolicyRenderersList converted to SparsePolicyRenderersList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PolicyRenderersList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePolicyRenderersList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePolicyRenderer)
	}

	return out
}

// Version returns the version of the content.
func (o PolicyRenderersList) Version() int {

	return 1
}

// PolicyRenderer represents the model of a policyrenderer
type PolicyRenderer struct {
	// List of policies rendered for the given set of tags.
	Policies PolicyRulesList `json:"policies" msgpack:"policies" bson:"-" mapstructure:"policies,omitempty"`

	// `Subject` (default): Set if the `processMode` should use the subject. `Object`:
	// Set if
	// the `processMode` should use the object. This only has effect when rendering an
	// SSH
	// authorization for now.
	ProcessMode PolicyRendererProcessModeValue `json:"processMode" msgpack:"processMode" bson:"-" mapstructure:"processMode,omitempty"`

	// List of tags of the object to render the hook for.
	Tags []string `json:"tags" msgpack:"tags" bson:"-" mapstructure:"tags,omitempty"`

	// Type of policy to render.
	Type PolicyRendererTypeValue `json:"type" msgpack:"type" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPolicyRenderer returns a new *PolicyRenderer
func NewPolicyRenderer() *PolicyRenderer {

	return &PolicyRenderer{
		ModelVersion: 1,
		Policies:     PolicyRulesList{},
		ProcessMode:  PolicyRendererProcessModeSubject,
		Tags:         []string{},
	}
}

// Identity returns the Identity of the object.
func (o *PolicyRenderer) Identity() elemental.Identity {

	return PolicyRendererIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PolicyRenderer) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PolicyRenderer) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PolicyRenderer) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPolicyRenderer{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PolicyRenderer) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPolicyRenderer{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *PolicyRenderer) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *PolicyRenderer) BleveType() string {

	return "policyrenderer"
}

// DefaultOrder returns the list of default ordering fields.
func (o *PolicyRenderer) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PolicyRenderer) Doc() string {

	return `Allows you to render policies of a given type for a given set of tags.`
}

func (o *PolicyRenderer) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PolicyRenderer) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePolicyRenderer{
			Policies:    &o.Policies,
			ProcessMode: &o.ProcessMode,
			Tags:        &o.Tags,
			Type:        &o.Type,
		}
	}

	sp := &SparsePolicyRenderer{}
	for _, f := range fields {
		switch f {
		case "policies":
			sp.Policies = &(o.Policies)
		case "processMode":
			sp.ProcessMode = &(o.ProcessMode)
		case "tags":
			sp.Tags = &(o.Tags)
		case "type":
			sp.Type = &(o.Type)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePolicyRenderer to the object.
func (o *PolicyRenderer) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePolicyRenderer)
	if so.Policies != nil {
		o.Policies = *so.Policies
	}
	if so.ProcessMode != nil {
		o.ProcessMode = *so.ProcessMode
	}
	if so.Tags != nil {
		o.Tags = *so.Tags
	}
	if so.Type != nil {
		o.Type = *so.Type
	}
}

// DeepCopy returns a deep copy if the PolicyRenderer.
func (o *PolicyRenderer) DeepCopy() *PolicyRenderer {

	if o == nil {
		return nil
	}

	out := &PolicyRenderer{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PolicyRenderer.
func (o *PolicyRenderer) DeepCopyInto(out *PolicyRenderer) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PolicyRenderer: %s", err))
	}

	*out = *target.(*PolicyRenderer)
}

// Validate valides the current information stored into the structure.
func (o *PolicyRenderer) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.Policies {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateStringInList("processMode", string(o.ProcessMode), []string{"Subject", "Object"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredExternal("tags", o.Tags); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("type", string(o.Type)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"APIAuthorization", "EnforcerProfile", "File", "Hook", "Infrastructure", "NamespaceMapping", "Network", "ProcessingUnit", "Quota", "Syscall", "TokenScope", "SSHAuthorization", "UserAccess"}, false); err != nil {
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
func (*PolicyRenderer) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PolicyRendererAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PolicyRendererLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PolicyRenderer) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PolicyRendererAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PolicyRenderer) ValueForAttribute(name string) interface{} {

	switch name {
	case "policies":
		return o.Policies
	case "processMode":
		return o.ProcessMode
	case "tags":
		return o.Tags
	case "type":
		return o.Type
	}

	return nil
}

// PolicyRendererAttributesMap represents the map of attribute for PolicyRenderer.
var PolicyRendererAttributesMap = map[string]elemental.AttributeSpecification{
	"Policies": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Policies",
		Description:    `List of policies rendered for the given set of tags.`,
		Exposed:        true,
		Name:           "policies",
		ReadOnly:       true,
		SubType:        "policyrule",
		Type:           "refList",
	},
	"ProcessMode": {
		AllowedChoices: []string{"Subject", "Object"},
		ConvertedName:  "ProcessMode",
		DefaultValue:   PolicyRendererProcessModeSubject,
		Description: `` + "`" + `Subject` + "`" + ` (default): Set if the ` + "`" + `processMode` + "`" + ` should use the subject. ` + "`" + `Object` + "`" + `:
Set if
the ` + "`" + `processMode` + "`" + ` should use the object. This only has effect when rendering an
SSH
authorization for now.`,
		Exposed: true,
		Name:    "processMode",
		Type:    "enum",
	},
	"Tags": {
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		Description:    `List of tags of the object to render the hook for.`,
		Exposed:        true,
		Name:           "tags",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
	"Type": {
		AllowedChoices: []string{"APIAuthorization", "EnforcerProfile", "File", "Hook", "Infrastructure", "NamespaceMapping", "Network", "ProcessingUnit", "Quota", "Syscall", "TokenScope", "SSHAuthorization", "UserAccess"},
		ConvertedName:  "Type",
		Description:    `Type of policy to render.`,
		Exposed:        true,
		Name:           "type",
		Required:       true,
		Type:           "enum",
	},
}

// PolicyRendererLowerCaseAttributesMap represents the map of attribute for PolicyRenderer.
var PolicyRendererLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"policies": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Policies",
		Description:    `List of policies rendered for the given set of tags.`,
		Exposed:        true,
		Name:           "policies",
		ReadOnly:       true,
		SubType:        "policyrule",
		Type:           "refList",
	},
	"processmode": {
		AllowedChoices: []string{"Subject", "Object"},
		ConvertedName:  "ProcessMode",
		DefaultValue:   PolicyRendererProcessModeSubject,
		Description: `` + "`" + `Subject` + "`" + ` (default): Set if the ` + "`" + `processMode` + "`" + ` should use the subject. ` + "`" + `Object` + "`" + `:
Set if
the ` + "`" + `processMode` + "`" + ` should use the object. This only has effect when rendering an
SSH
authorization for now.`,
		Exposed: true,
		Name:    "processMode",
		Type:    "enum",
	},
	"tags": {
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		Description:    `List of tags of the object to render the hook for.`,
		Exposed:        true,
		Name:           "tags",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
	"type": {
		AllowedChoices: []string{"APIAuthorization", "EnforcerProfile", "File", "Hook", "Infrastructure", "NamespaceMapping", "Network", "ProcessingUnit", "Quota", "Syscall", "TokenScope", "SSHAuthorization", "UserAccess"},
		ConvertedName:  "Type",
		Description:    `Type of policy to render.`,
		Exposed:        true,
		Name:           "type",
		Required:       true,
		Type:           "enum",
	},
}

// SparsePolicyRenderersList represents a list of SparsePolicyRenderers
type SparsePolicyRenderersList []*SparsePolicyRenderer

// Identity returns the identity of the objects in the list.
func (o SparsePolicyRenderersList) Identity() elemental.Identity {

	return PolicyRendererIdentity
}

// Copy returns a pointer to a copy the SparsePolicyRenderersList.
func (o SparsePolicyRenderersList) Copy() elemental.Identifiables {

	copy := append(SparsePolicyRenderersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePolicyRenderersList.
func (o SparsePolicyRenderersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePolicyRenderersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePolicyRenderer))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePolicyRenderersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePolicyRenderersList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePolicyRenderersList converted to PolicyRenderersList.
func (o SparsePolicyRenderersList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePolicyRenderersList) Version() int {

	return 1
}

// SparsePolicyRenderer represents the sparse version of a policyrenderer.
type SparsePolicyRenderer struct {
	// List of policies rendered for the given set of tags.
	Policies *PolicyRulesList `json:"policies,omitempty" msgpack:"policies,omitempty" bson:"-" mapstructure:"policies,omitempty"`

	// `Subject` (default): Set if the `processMode` should use the subject. `Object`:
	// Set if
	// the `processMode` should use the object. This only has effect when rendering an
	// SSH
	// authorization for now.
	ProcessMode *PolicyRendererProcessModeValue `json:"processMode,omitempty" msgpack:"processMode,omitempty" bson:"-" mapstructure:"processMode,omitempty"`

	// List of tags of the object to render the hook for.
	Tags *[]string `json:"tags,omitempty" msgpack:"tags,omitempty" bson:"-" mapstructure:"tags,omitempty"`

	// Type of policy to render.
	Type *PolicyRendererTypeValue `json:"type,omitempty" msgpack:"type,omitempty" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePolicyRenderer returns a new  SparsePolicyRenderer.
func NewSparsePolicyRenderer() *SparsePolicyRenderer {
	return &SparsePolicyRenderer{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePolicyRenderer) Identity() elemental.Identity {

	return PolicyRendererIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePolicyRenderer) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePolicyRenderer) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePolicyRenderer) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePolicyRenderer{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePolicyRenderer) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePolicyRenderer{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparsePolicyRenderer) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePolicyRenderer) ToPlain() elemental.PlainIdentifiable {

	out := NewPolicyRenderer()
	if o.Policies != nil {
		out.Policies = *o.Policies
	}
	if o.ProcessMode != nil {
		out.ProcessMode = *o.ProcessMode
	}
	if o.Tags != nil {
		out.Tags = *o.Tags
	}
	if o.Type != nil {
		out.Type = *o.Type
	}

	return out
}

// DeepCopy returns a deep copy if the SparsePolicyRenderer.
func (o *SparsePolicyRenderer) DeepCopy() *SparsePolicyRenderer {

	if o == nil {
		return nil
	}

	out := &SparsePolicyRenderer{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePolicyRenderer.
func (o *SparsePolicyRenderer) DeepCopyInto(out *SparsePolicyRenderer) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePolicyRenderer: %s", err))
	}

	*out = *target.(*SparsePolicyRenderer)
}

type mongoAttributesPolicyRenderer struct {
}
type mongoAttributesSparsePolicyRenderer struct {
}
