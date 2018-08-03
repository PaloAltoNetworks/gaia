package gaia

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
)

// KubernetesClusterIndexes lists the attribute compound indexes.
var KubernetesClusterIndexes = [][]string{}

// KubernetesClusterActivationTypeValue represents the possible values for attribute "activationType".
type KubernetesClusterActivationTypeValue string

const (
	// KubernetesClusterActivationTypeKubeSquall represents the value KubeSquall.
	KubernetesClusterActivationTypeKubeSquall KubernetesClusterActivationTypeValue = "KubeSquall"

	// KubernetesClusterActivationTypePodAtomic represents the value PodAtomic.
	KubernetesClusterActivationTypePodAtomic KubernetesClusterActivationTypeValue = "PodAtomic"

	// KubernetesClusterActivationTypePodContainers represents the value PodContainers.
	KubernetesClusterActivationTypePodContainers KubernetesClusterActivationTypeValue = "PodContainers"
)

// KubernetesClusterIdentity represents the Identity of the object.
var KubernetesClusterIdentity = elemental.Identity{
	Name:     "kubernetescluster",
	Category: "kubernetesclusters",
	Private:  false,
}

// KubernetesClustersList represents a list of KubernetesClusters
type KubernetesClustersList []*KubernetesCluster

// Identity returns the identity of the objects in the list.
func (o KubernetesClustersList) Identity() elemental.Identity {

	return KubernetesClusterIdentity
}

// Copy returns a pointer to a copy the KubernetesClustersList.
func (o KubernetesClustersList) Copy() elemental.Identifiables {

	copy := append(KubernetesClustersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the KubernetesClustersList.
func (o KubernetesClustersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(KubernetesClustersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*KubernetesCluster))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o KubernetesClustersList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o KubernetesClustersList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o KubernetesClustersList) Version() int {

	return 1
}

// KubernetesCluster represents the model of a kubernetescluster
type KubernetesCluster struct {
	// Link to the API authorization policy.
	APIAuthorizationPolicyID string `json:"-" bson:"apiauthorizationpolicyid" mapstructure:"-,omitempty"`

	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Defines the mode of activation on the KubernetesCluster.
	ActivationType KubernetesClusterActivationTypeValue `json:"activationType" bson:"activationtype" mapstructure:"activationType,omitempty"`

	// Link to the certificate created in Vince for this cluster.
	CertificateID string `json:"-" bson:"certificateid" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// base64 of the .tar.gz file that contains all the .YAMLs files needed to create
	// the aporeto side on your kubernetes Cluster.
	KubernetesDefinitions string `json:"kubernetesDefinitions" bson:"-" mapstructure:"kubernetesDefinitions,omitempty"`

	// The name of your cluster.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Link to your namespace.
	NamespaceID string `json:"-" bson:"namespaceid" mapstructure:"-,omitempty"`

	// ID of the parent account.
	ParentID string `json:"parentID" bson:"parentid" mapstructure:"parentID,omitempty"`

	// Regenerates the k8s files and certificates.
	Regenerate bool `json:"regenerate" bson:"-" mapstructure:"regenerate,omitempty"`

	// The namespace in which the Kubernetes specific namespace will be created. By
	// default your account namespace.
	TargetNamespace string `json:"targetNamespace" bson:"targetnamespace" mapstructure:"targetNamespace,omitempty"`

	// List of target networks.
	TargetNetworks []string `json:"targetNetworks" bson:"targetnetworks" mapstructure:"targetNetworks,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewKubernetesCluster returns a new *KubernetesCluster
func NewKubernetesCluster() *KubernetesCluster {

	return &KubernetesCluster{
		ModelVersion:   1,
		ActivationType: KubernetesClusterActivationTypeKubeSquall,
	}
}

// Identity returns the Identity of the object.
func (o *KubernetesCluster) Identity() elemental.Identity {

	return KubernetesClusterIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *KubernetesCluster) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *KubernetesCluster) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *KubernetesCluster) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *KubernetesCluster) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *KubernetesCluster) Doc() string {
	return `Create a remote Kubernetes Cluster integration.`
}

func (o *KubernetesCluster) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *KubernetesCluster) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("activationType", string(o.ActivationType), []string{"KubeSquall", "PodAtomic", "PodContainers"}, false); err != nil {
		errors = append(errors, err)
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
func (*KubernetesCluster) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := KubernetesClusterAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return KubernetesClusterLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*KubernetesCluster) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return KubernetesClusterAttributesMap
}

// KubernetesClusterAttributesMap represents the map of attribute for KubernetesCluster.
var KubernetesClusterAttributesMap = map[string]elemental.AttributeSpecification{
	"APIAuthorizationPolicyID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "APIAuthorizationPolicyID",
		Description:    `Link to the API authorization policy.`,
		Name:           "APIAuthorizationPolicyID",
		Stored:         true,
		Type:           "string",
	},
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ActivationType": elemental.AttributeSpecification{
		AllowedChoices: []string{"KubeSquall", "PodAtomic", "PodContainers"},
		ConvertedName:  "ActivationType",
		DefaultValue:   KubernetesClusterActivationTypeKubeSquall,
		Description:    `Defines the mode of activation on the KubernetesCluster.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "activationType",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"CertificateID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateID",
		Description:    `Link to the certificate created in Vince for this cluster.`,
		Name:           "certificateID",
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"KubernetesDefinitions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "KubernetesDefinitions",
		Description: `base64 of the .tar.gz file that contains all the .YAMLs files needed to create
the aporeto side on your kubernetes Cluster.`,
		Exposed:    true,
		Filterable: true,
		Name:       "kubernetesDefinitions",
		Orderable:  true,
		ReadOnly:   true,
		Type:       "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `The name of your cluster.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "name",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"NamespaceID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NamespaceID",
		Description:    `Link to your namespace.`,
		Name:           "namespaceID",
		Stored:         true,
		Type:           "string",
	},
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ParentID",
		Description:    `ID of the parent account.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Regenerate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Regenerate",
		Description:    `Regenerates the k8s files and certificates.`,
		Exposed:        true,
		Name:           "regenerate",
		Type:           "boolean",
	},
	"TargetNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNamespace",
		Description: `The namespace in which the Kubernetes specific namespace will be created. By
default your account namespace.`,
		Exposed:    true,
		Filterable: true,
		Name:       "targetNamespace",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
	"TargetNetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNetworks",
		Deprecated:     true,
		Description:    `List of target networks.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "targetNetworks",
		Orderable:      true,
		Stored:         true,
		SubType:        "target_networks_list",
		Type:           "external",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}

// KubernetesClusterLowerCaseAttributesMap represents the map of attribute for KubernetesCluster.
var KubernetesClusterLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"apiauthorizationpolicyid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "APIAuthorizationPolicyID",
		Description:    `Link to the API authorization policy.`,
		Name:           "APIAuthorizationPolicyID",
		Stored:         true,
		Type:           "string",
	},
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"activationtype": elemental.AttributeSpecification{
		AllowedChoices: []string{"KubeSquall", "PodAtomic", "PodContainers"},
		ConvertedName:  "ActivationType",
		DefaultValue:   KubernetesClusterActivationTypeKubeSquall,
		Description:    `Defines the mode of activation on the KubernetesCluster.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "activationType",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"certificateid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateID",
		Description:    `Link to the certificate created in Vince for this cluster.`,
		Name:           "certificateID",
		Stored:         true,
		Type:           "string",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"kubernetesdefinitions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "KubernetesDefinitions",
		Description: `base64 of the .tar.gz file that contains all the .YAMLs files needed to create
the aporeto side on your kubernetes Cluster.`,
		Exposed:    true,
		Filterable: true,
		Name:       "kubernetesDefinitions",
		Orderable:  true,
		ReadOnly:   true,
		Type:       "string",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `The name of your cluster.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "name",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"namespaceid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NamespaceID",
		Description:    `Link to your namespace.`,
		Name:           "namespaceID",
		Stored:         true,
		Type:           "string",
	},
	"parentid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ParentID",
		Description:    `ID of the parent account.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"regenerate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Regenerate",
		Description:    `Regenerates the k8s files and certificates.`,
		Exposed:        true,
		Name:           "regenerate",
		Type:           "boolean",
	},
	"targetnamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNamespace",
		Description: `The namespace in which the Kubernetes specific namespace will be created. By
default your account namespace.`,
		Exposed:    true,
		Filterable: true,
		Name:       "targetNamespace",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
	"targetnetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNetworks",
		Deprecated:     true,
		Description:    `List of target networks.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "targetNetworks",
		Orderable:      true,
		Stored:         true,
		SubType:        "target_networks_list",
		Type:           "external",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}
