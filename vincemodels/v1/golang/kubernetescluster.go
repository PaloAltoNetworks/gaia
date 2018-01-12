package vincemodels

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"time"
)

// KubernetesClusterIdentity represents the Identity of the object.
var KubernetesClusterIdentity = elemental.Identity{
	Name:     "kubernetescluster",
	Category: "kubernetesclusters",
}

// KubernetesClustersList represents a list of KubernetesClusters
type KubernetesClustersList []*KubernetesCluster

// ContentIdentity returns the identity of the objects in the list.
func (o KubernetesClustersList) ContentIdentity() elemental.Identity {

	return KubernetesClusterIdentity
}

// Copy returns a pointer to a copy the KubernetesClustersList.
func (o KubernetesClustersList) Copy() elemental.ContentIdentifiable {

	copy := append(KubernetesClustersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the KubernetesClustersList.
func (o KubernetesClustersList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

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
	APIAuthorizationPolicyID string `json:"-" bson:"apiauthorizationpolicyid"`

	// Link to the certificate created in Vince for this cluster
	CertificateID string `json:"-" bson:"certificateid"`

	// base64 of the .tar.gz file that contains all the .YAMLs files needed to create the aporeto side on your kubernetes Cluster
	KubernetesDefinitions string `json:"kubernetesDefinitions" bson:"-"`

	// The name of your cluster
	Name string `json:"name" bson:"name"`

	// Link to your namespace
	NamespaceID string `json:"-" bson:"namespaceid"`

	// ID of the parent account.
	ParentID string `json:"parentID" bson:"parentid"`

	// Regenerates the k8s files and certificates.
	Regenerate bool `json:"regenerate" bson:"-"`

	// The namespace in which the Kubernetes specific namespace will be created. By default your account namespace.
	TargetNamespace string `json:"targetNamespace" bson:"targetnamespace"`

	// List of target networks [deprecated]
	TargetNetworks []string `json:"targetNetworks" bson:"targetnetworks"`

	// ID of the object.
	ID string `json:"ID" bson:"_id"`

	// createdAt represents the creation date of the object.
	CreateTime time.Time `json:"createTime" bson:"createtime"`

	// UpdateTime represents the last update date of the objct.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewKubernetesCluster returns a new *KubernetesCluster
func NewKubernetesCluster() *KubernetesCluster {

	return &KubernetesCluster{
		ModelVersion: 1,
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
		Description:    `Link to the API authorization policy.`,
		Format:         "free",
		Name:           "APIAuthorizationPolicyID",
		Stored:         true,
		Type:           "string",
	},
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"CertificateID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Link to the certificate created in Vince for this cluster `,
		Format:         "free",
		Name:           "certificateID",
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `createdAt represents the creation date of the object.`,
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
		Description:    `base64 of the .tar.gz file that contains all the .YAMLs files needed to create the aporeto side on your kubernetes Cluster`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "kubernetesDefinitions",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `The name of your cluster`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"NamespaceID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Link to your namespace `,
		Format:         "free",
		Name:           "namespaceID",
		Stored:         true,
		Type:           "string",
	},
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `ID of the parent account.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Regenerate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Regenerates the k8s files and certificates.`,
		Exposed:        true,
		Name:           "regenerate",
		Type:           "boolean",
	},
	"TargetNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `The namespace in which the Kubernetes specific namespace will be created. By default your account namespace.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "targetNamespace",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"TargetNetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `List of target networks [deprecated]`,
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
		Description:    `UpdateTime represents the last update date of the objct.`,
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
		Description:    `Link to the API authorization policy.`,
		Format:         "free",
		Name:           "APIAuthorizationPolicyID",
		Stored:         true,
		Type:           "string",
	},
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"certificateid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Link to the certificate created in Vince for this cluster `,
		Format:         "free",
		Name:           "certificateID",
		Stored:         true,
		Type:           "string",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `createdAt represents the creation date of the object.`,
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
		Description:    `base64 of the .tar.gz file that contains all the .YAMLs files needed to create the aporeto side on your kubernetes Cluster`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "kubernetesDefinitions",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `The name of your cluster`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"namespaceid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Link to your namespace `,
		Format:         "free",
		Name:           "namespaceID",
		Stored:         true,
		Type:           "string",
	},
	"parentid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `ID of the parent account.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"regenerate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Regenerates the k8s files and certificates.`,
		Exposed:        true,
		Name:           "regenerate",
		Type:           "boolean",
	},
	"targetnamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `The namespace in which the Kubernetes specific namespace will be created. By default your account namespace.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "targetNamespace",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"targetnetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `List of target networks [deprecated]`,
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
		Description:    `UpdateTime represents the last update date of the objct.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}
