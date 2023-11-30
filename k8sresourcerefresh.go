// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// K8sResourceRefreshIdentity represents the Identity of the object.
var K8sResourceRefreshIdentity = elemental.Identity{
	Name:     "k8sresourcerefresh",
	Category: "k8sresourcerefreshes",
	Package:  "pandemona",
	Private:  false,
}

// K8sResourceRefreshsList represents a list of K8sResourceRefreshs
type K8sResourceRefreshsList []*K8sResourceRefresh

// Identity returns the identity of the objects in the list.
func (o K8sResourceRefreshsList) Identity() elemental.Identity {

	return K8sResourceRefreshIdentity
}

// Copy returns a pointer to a copy the K8sResourceRefreshsList.
func (o K8sResourceRefreshsList) Copy() elemental.Identifiables {

	out := append(K8sResourceRefreshsList{}, o...)
	return &out
}

// Append appends the objects to the a new copy of the K8sResourceRefreshsList.
func (o K8sResourceRefreshsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(K8sResourceRefreshsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*K8sResourceRefresh))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o K8sResourceRefreshsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o K8sResourceRefreshsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the K8sResourceRefreshsList converted to SparseK8sResourceRefreshsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o K8sResourceRefreshsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseK8sResourceRefreshsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseK8sResourceRefresh)
	}

	return out
}

// Version returns the version of the content.
func (o K8sResourceRefreshsList) Version() int {

	return 1
}

// K8sResourceRefresh represents the model of a k8sresourcerefresh
type K8sResourceRefresh struct {
	// Set to `true` to make the request run in the background.
	Background bool `json:"background" msgpack:"background" bson:"-" mapstructure:"background,omitempty"`

	// The amount of time the refresh has before timing out.
	Timeout string `json:"timeout" msgpack:"timeout" bson:"-" mapstructure:"timeout,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewK8sResourceRefresh returns a new *K8sResourceRefresh
func NewK8sResourceRefresh() *K8sResourceRefresh {

	return &K8sResourceRefresh{
		ModelVersion: 1,
		Timeout:      "60m",
	}
}

// Identity returns the Identity of the object.
func (o *K8sResourceRefresh) Identity() elemental.Identity {

	return K8sResourceRefreshIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *K8sResourceRefresh) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *K8sResourceRefresh) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *K8sResourceRefresh) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesK8sResourceRefresh{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *K8sResourceRefresh) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesK8sResourceRefresh{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *K8sResourceRefresh) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *K8sResourceRefresh) BleveType() string {

	return "k8sresourcerefresh"
}

// DefaultOrder returns the list of default ordering fields.
func (o *K8sResourceRefresh) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *K8sResourceRefresh) Doc() string {

	return `When requested, k8s resources will start the process of pulling down all
Kubernetes resources from PrismaCloud.`
}

func (o *K8sResourceRefresh) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *K8sResourceRefresh) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseK8sResourceRefresh{
			Background: &o.Background,
			Timeout:    &o.Timeout,
		}
	}

	sp := &SparseK8sResourceRefresh{}
	for _, f := range fields {
		switch f {
		case "background":
			sp.Background = &(o.Background)
		case "timeout":
			sp.Timeout = &(o.Timeout)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseK8sResourceRefresh to the object.
func (o *K8sResourceRefresh) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseK8sResourceRefresh)
	if so.Background != nil {
		o.Background = *so.Background
	}
	if so.Timeout != nil {
		o.Timeout = *so.Timeout
	}
}

// DeepCopy returns a deep copy if the K8sResourceRefresh.
func (o *K8sResourceRefresh) DeepCopy() *K8sResourceRefresh {

	if o == nil {
		return nil
	}

	out := &K8sResourceRefresh{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *K8sResourceRefresh.
func (o *K8sResourceRefresh) DeepCopyInto(out *K8sResourceRefresh) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy K8sResourceRefresh: %s", err))
	}

	*out = *target.(*K8sResourceRefresh)
}

// Validate valides the current information stored into the structure.
func (o *K8sResourceRefresh) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTimeDuration("timeout", o.Timeout); err != nil {
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
func (*K8sResourceRefresh) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := K8sResourceRefreshAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return K8sResourceRefreshLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*K8sResourceRefresh) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return K8sResourceRefreshAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *K8sResourceRefresh) ValueForAttribute(name string) any {

	switch name {
	case "background":
		return o.Background
	case "timeout":
		return o.Timeout
	}

	return nil
}

// K8sResourceRefreshAttributesMap represents the map of attribute for K8sResourceRefresh.
var K8sResourceRefreshAttributesMap = map[string]elemental.AttributeSpecification{
	"Background": {
		AllowedChoices: []string{},
		ConvertedName:  "Background",
		Description:    `Set to ` + "`" + `true` + "`" + ` to make the request run in the background.`,
		Exposed:        true,
		Name:           "background",
		Type:           "boolean",
	},
	"Timeout": {
		AllowedChoices: []string{},
		ConvertedName:  "Timeout",
		DefaultValue:   "60m",
		Description:    `The amount of time the refresh has before timing out.`,
		Exposed:        true,
		Name:           "timeout",
		Type:           "string",
	},
}

// K8sResourceRefreshLowerCaseAttributesMap represents the map of attribute for K8sResourceRefresh.
var K8sResourceRefreshLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"background": {
		AllowedChoices: []string{},
		ConvertedName:  "Background",
		Description:    `Set to ` + "`" + `true` + "`" + ` to make the request run in the background.`,
		Exposed:        true,
		Name:           "background",
		Type:           "boolean",
	},
	"timeout": {
		AllowedChoices: []string{},
		ConvertedName:  "Timeout",
		DefaultValue:   "60m",
		Description:    `The amount of time the refresh has before timing out.`,
		Exposed:        true,
		Name:           "timeout",
		Type:           "string",
	},
}

// SparseK8sResourceRefreshsList represents a list of SparseK8sResourceRefreshs
type SparseK8sResourceRefreshsList []*SparseK8sResourceRefresh

// Identity returns the identity of the objects in the list.
func (o SparseK8sResourceRefreshsList) Identity() elemental.Identity {

	return K8sResourceRefreshIdentity
}

// Copy returns a pointer to a copy the SparseK8sResourceRefreshsList.
func (o SparseK8sResourceRefreshsList) Copy() elemental.Identifiables {

	copy := append(SparseK8sResourceRefreshsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseK8sResourceRefreshsList.
func (o SparseK8sResourceRefreshsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseK8sResourceRefreshsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseK8sResourceRefresh))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseK8sResourceRefreshsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseK8sResourceRefreshsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseK8sResourceRefreshsList converted to K8sResourceRefreshsList.
func (o SparseK8sResourceRefreshsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseK8sResourceRefreshsList) Version() int {

	return 1
}

// SparseK8sResourceRefresh represents the sparse version of a k8sresourcerefresh.
type SparseK8sResourceRefresh struct {
	// Set to `true` to make the request run in the background.
	Background *bool `json:"background,omitempty" msgpack:"background,omitempty" bson:"-" mapstructure:"background,omitempty"`

	// The amount of time the refresh has before timing out.
	Timeout *string `json:"timeout,omitempty" msgpack:"timeout,omitempty" bson:"-" mapstructure:"timeout,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseK8sResourceRefresh returns a new  SparseK8sResourceRefresh.
func NewSparseK8sResourceRefresh() *SparseK8sResourceRefresh {
	return &SparseK8sResourceRefresh{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseK8sResourceRefresh) Identity() elemental.Identity {

	return K8sResourceRefreshIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseK8sResourceRefresh) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseK8sResourceRefresh) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseK8sResourceRefresh) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseK8sResourceRefresh{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseK8sResourceRefresh) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseK8sResourceRefresh{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseK8sResourceRefresh) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseK8sResourceRefresh) ToPlain() elemental.PlainIdentifiable {

	out := NewK8sResourceRefresh()
	if o.Background != nil {
		out.Background = *o.Background
	}
	if o.Timeout != nil {
		out.Timeout = *o.Timeout
	}

	return out
}

// DeepCopy returns a deep copy if the SparseK8sResourceRefresh.
func (o *SparseK8sResourceRefresh) DeepCopy() *SparseK8sResourceRefresh {

	if o == nil {
		return nil
	}

	out := &SparseK8sResourceRefresh{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseK8sResourceRefresh.
func (o *SparseK8sResourceRefresh) DeepCopyInto(out *SparseK8sResourceRefresh) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseK8sResourceRefresh: %s", err))
	}

	*out = *target.(*SparseK8sResourceRefresh)
}

type mongoAttributesK8sResourceRefresh struct {
}
type mongoAttributesSparseK8sResourceRefresh struct {
}
