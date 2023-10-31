// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudAlertsControllerActionValue represents the possible values for attribute "action".
type CloudAlertsControllerActionValue string

const (
	// CloudAlertsControllerActionGenerate represents the value Generate.
	CloudAlertsControllerActionGenerate CloudAlertsControllerActionValue = "Generate"

	// CloudAlertsControllerActionRaise represents the value Raise.
	CloudAlertsControllerActionRaise CloudAlertsControllerActionValue = "Raise"

	// CloudAlertsControllerActionResolve represents the value Resolve.
	CloudAlertsControllerActionResolve CloudAlertsControllerActionValue = "Resolve"
)

// CloudAlertsControllerIdentity represents the Identity of the object.
var CloudAlertsControllerIdentity = elemental.Identity{
	Name:     "cloudalertscontroller",
	Category: "cloudalertscontrollers",
	Package:  "vargid",
	Private:  false,
}

// CloudAlertsControllersList represents a list of CloudAlertsControllers
type CloudAlertsControllersList []*CloudAlertsController

// Identity returns the identity of the objects in the list.
func (o CloudAlertsControllersList) Identity() elemental.Identity {

	return CloudAlertsControllerIdentity
}

// Copy returns a pointer to a copy the CloudAlertsControllersList.
func (o CloudAlertsControllersList) Copy() elemental.Identifiables {

	out := append(CloudAlertsControllersList{}, o...)
	return &out
}

// Append appends the objects to the a new copy of the CloudAlertsControllersList.
func (o CloudAlertsControllersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudAlertsControllersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudAlertsController))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudAlertsControllersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudAlertsControllersList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CloudAlertsControllersList converted to SparseCloudAlertsControllersList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudAlertsControllersList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudAlertsControllersList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudAlertsController)
	}

	return out
}

// Version returns the version of the content.
func (o CloudAlertsControllersList) Version() int {

	return 1
}

// CloudAlertsController represents the model of a cloudalertscontroller
type CloudAlertsController struct {
	// Action type to perform.
	Action CloudAlertsControllerActionValue `json:"action" msgpack:"action" bson:"-" mapstructure:"action,omitempty"`

	// IDs of cloud accounts to scan and generate alerts. When left empty all cloud
	// accounts in the tenant are considered. This attribute is only supported with
	// action 'Generate'.
	CloudAccountIDs []string `json:"cloudAccountIDs" msgpack:"cloudAccountIDs" bson:"-" mapstructure:"cloudAccountIDs,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Prisma Cloud policy IDs to scan and generate alerts. When left empty all
	// policies
	// in the tenant are considered. This attribute is only supported with action
	// 'Generate'.
	PrismaCloudPolicyIDs []string `json:"prismaCloudPolicyIDs" msgpack:"prismaCloudPolicyIDs" bson:"-" mapstructure:"prismaCloudPolicyIDs,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudAlertsController returns a new *CloudAlertsController
func NewCloudAlertsController() *CloudAlertsController {

	return &CloudAlertsController{
		ModelVersion:         1,
		Action:               CloudAlertsControllerActionGenerate,
		CloudAccountIDs:      []string{},
		PrismaCloudPolicyIDs: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *CloudAlertsController) Identity() elemental.Identity {

	return CloudAlertsControllerIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudAlertsController) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudAlertsController) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudAlertsController) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudAlertsController{}

	s.Namespace = o.Namespace

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudAlertsController) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudAlertsController{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Namespace = s.Namespace

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudAlertsController) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudAlertsController) BleveType() string {

	return "cloudalertscontroller"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudAlertsController) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CloudAlertsController) Doc() string {

	return `Control message model to raise/resolve/generate cloud alerts.`
}

func (o *CloudAlertsController) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetNamespace returns the Namespace of the receiver.
func (o *CloudAlertsController) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *CloudAlertsController) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudAlertsController) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudAlertsController{
			Action:               &o.Action,
			CloudAccountIDs:      &o.CloudAccountIDs,
			Namespace:            &o.Namespace,
			PrismaCloudPolicyIDs: &o.PrismaCloudPolicyIDs,
		}
	}

	sp := &SparseCloudAlertsController{}
	for _, f := range fields {
		switch f {
		case "action":
			sp.Action = &(o.Action)
		case "cloudAccountIDs":
			sp.CloudAccountIDs = &(o.CloudAccountIDs)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "prismaCloudPolicyIDs":
			sp.PrismaCloudPolicyIDs = &(o.PrismaCloudPolicyIDs)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCloudAlertsController to the object.
func (o *CloudAlertsController) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudAlertsController)
	if so.Action != nil {
		o.Action = *so.Action
	}
	if so.CloudAccountIDs != nil {
		o.CloudAccountIDs = *so.CloudAccountIDs
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.PrismaCloudPolicyIDs != nil {
		o.PrismaCloudPolicyIDs = *so.PrismaCloudPolicyIDs
	}
}

// DeepCopy returns a deep copy if the CloudAlertsController.
func (o *CloudAlertsController) DeepCopy() *CloudAlertsController {

	if o == nil {
		return nil
	}

	out := &CloudAlertsController{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudAlertsController.
func (o *CloudAlertsController) DeepCopyInto(out *CloudAlertsController) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudAlertsController: %s", err))
	}

	*out = *target.(*CloudAlertsController)
}

// Validate valides the current information stored into the structure.
func (o *CloudAlertsController) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("action", string(o.Action), []string{"Generate", "Raise", "Resolve"}, false); err != nil {
		errors = errors.Append(err)
	}

	// Custom object validation.
	if err := ValidateCloudAlertsControllerEntity(o); err != nil {
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
func (*CloudAlertsController) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudAlertsControllerAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudAlertsControllerLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudAlertsController) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudAlertsControllerAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudAlertsController) ValueForAttribute(name string) any {

	switch name {
	case "action":
		return o.Action
	case "cloudAccountIDs":
		return o.CloudAccountIDs
	case "namespace":
		return o.Namespace
	case "prismaCloudPolicyIDs":
		return o.PrismaCloudPolicyIDs
	}

	return nil
}

// CloudAlertsControllerAttributesMap represents the map of attribute for CloudAlertsController.
var CloudAlertsControllerAttributesMap = map[string]elemental.AttributeSpecification{
	"Action": {
		AllowedChoices: []string{"Generate", "Raise", "Resolve"},
		ConvertedName:  "Action",
		DefaultValue:   CloudAlertsControllerActionGenerate,
		Description:    `Action type to perform.`,
		Exposed:        true,
		Name:           "action",
		Type:           "enum",
	},
	"CloudAccountIDs": {
		AllowedChoices: []string{},
		ConvertedName:  "CloudAccountIDs",
		Description: `IDs of cloud accounts to scan and generate alerts. When left empty all cloud
accounts in the tenant are considered. This attribute is only supported with
action 'Generate'.`,
		Exposed: true,
		Name:    "cloudAccountIDs",
		SubType: "string",
		Type:    "list",
	},
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"PrismaCloudPolicyIDs": {
		AllowedChoices: []string{},
		ConvertedName:  "PrismaCloudPolicyIDs",
		Description: `Prisma Cloud policy IDs to scan and generate alerts. When left empty all
policies
in the tenant are considered. This attribute is only supported with action
'Generate'.`,
		Exposed: true,
		Name:    "prismaCloudPolicyIDs",
		SubType: "string",
		Type:    "list",
	},
}

// CloudAlertsControllerLowerCaseAttributesMap represents the map of attribute for CloudAlertsController.
var CloudAlertsControllerLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"action": {
		AllowedChoices: []string{"Generate", "Raise", "Resolve"},
		ConvertedName:  "Action",
		DefaultValue:   CloudAlertsControllerActionGenerate,
		Description:    `Action type to perform.`,
		Exposed:        true,
		Name:           "action",
		Type:           "enum",
	},
	"cloudaccountids": {
		AllowedChoices: []string{},
		ConvertedName:  "CloudAccountIDs",
		Description: `IDs of cloud accounts to scan and generate alerts. When left empty all cloud
accounts in the tenant are considered. This attribute is only supported with
action 'Generate'.`,
		Exposed: true,
		Name:    "cloudAccountIDs",
		SubType: "string",
		Type:    "list",
	},
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"prismacloudpolicyids": {
		AllowedChoices: []string{},
		ConvertedName:  "PrismaCloudPolicyIDs",
		Description: `Prisma Cloud policy IDs to scan and generate alerts. When left empty all
policies
in the tenant are considered. This attribute is only supported with action
'Generate'.`,
		Exposed: true,
		Name:    "prismaCloudPolicyIDs",
		SubType: "string",
		Type:    "list",
	},
}

// SparseCloudAlertsControllersList represents a list of SparseCloudAlertsControllers
type SparseCloudAlertsControllersList []*SparseCloudAlertsController

// Identity returns the identity of the objects in the list.
func (o SparseCloudAlertsControllersList) Identity() elemental.Identity {

	return CloudAlertsControllerIdentity
}

// Copy returns a pointer to a copy the SparseCloudAlertsControllersList.
func (o SparseCloudAlertsControllersList) Copy() elemental.Identifiables {

	copy := append(SparseCloudAlertsControllersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudAlertsControllersList.
func (o SparseCloudAlertsControllersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudAlertsControllersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudAlertsController))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudAlertsControllersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudAlertsControllersList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCloudAlertsControllersList converted to CloudAlertsControllersList.
func (o SparseCloudAlertsControllersList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudAlertsControllersList) Version() int {

	return 1
}

// SparseCloudAlertsController represents the sparse version of a cloudalertscontroller.
type SparseCloudAlertsController struct {
	// Action type to perform.
	Action *CloudAlertsControllerActionValue `json:"action,omitempty" msgpack:"action,omitempty" bson:"-" mapstructure:"action,omitempty"`

	// IDs of cloud accounts to scan and generate alerts. When left empty all cloud
	// accounts in the tenant are considered. This attribute is only supported with
	// action 'Generate'.
	CloudAccountIDs *[]string `json:"cloudAccountIDs,omitempty" msgpack:"cloudAccountIDs,omitempty" bson:"-" mapstructure:"cloudAccountIDs,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Prisma Cloud policy IDs to scan and generate alerts. When left empty all
	// policies
	// in the tenant are considered. This attribute is only supported with action
	// 'Generate'.
	PrismaCloudPolicyIDs *[]string `json:"prismaCloudPolicyIDs,omitempty" msgpack:"prismaCloudPolicyIDs,omitempty" bson:"-" mapstructure:"prismaCloudPolicyIDs,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCloudAlertsController returns a new  SparseCloudAlertsController.
func NewSparseCloudAlertsController() *SparseCloudAlertsController {
	return &SparseCloudAlertsController{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudAlertsController) Identity() elemental.Identity {

	return CloudAlertsControllerIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudAlertsController) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudAlertsController) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudAlertsController) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudAlertsController{}

	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudAlertsController) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudAlertsController{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCloudAlertsController) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudAlertsController) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudAlertsController()
	if o.Action != nil {
		out.Action = *o.Action
	}
	if o.CloudAccountIDs != nil {
		out.CloudAccountIDs = *o.CloudAccountIDs
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.PrismaCloudPolicyIDs != nil {
		out.PrismaCloudPolicyIDs = *o.PrismaCloudPolicyIDs
	}

	return out
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseCloudAlertsController) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseCloudAlertsController) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// DeepCopy returns a deep copy if the SparseCloudAlertsController.
func (o *SparseCloudAlertsController) DeepCopy() *SparseCloudAlertsController {

	if o == nil {
		return nil
	}

	out := &SparseCloudAlertsController{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudAlertsController.
func (o *SparseCloudAlertsController) DeepCopyInto(out *SparseCloudAlertsController) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudAlertsController: %s", err))
	}

	*out = *target.(*SparseCloudAlertsController)
}

type mongoAttributesCloudAlertsController struct {
	Namespace string `bson:"namespace"`
}
type mongoAttributesSparseCloudAlertsController struct {
	Namespace *string `bson:"namespace,omitempty"`
}
