// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudNetworkQueryFilterResourceTypeValue represents the possible values for attribute "resourceType".
type CloudNetworkQueryFilterResourceTypeValue string

const (
	// CloudNetworkQueryFilterResourceTypeInstance represents the value Instance.
	CloudNetworkQueryFilterResourceTypeInstance CloudNetworkQueryFilterResourceTypeValue = "Instance"

	// CloudNetworkQueryFilterResourceTypeInterface represents the value Interface.
	CloudNetworkQueryFilterResourceTypeInterface CloudNetworkQueryFilterResourceTypeValue = "Interface"

	// CloudNetworkQueryFilterResourceTypeK8sService represents the value K8sService.
	CloudNetworkQueryFilterResourceTypeK8sService CloudNetworkQueryFilterResourceTypeValue = "K8sService"

	// CloudNetworkQueryFilterResourceTypePaaS represents the value PaaS.
	CloudNetworkQueryFilterResourceTypePaaS CloudNetworkQueryFilterResourceTypeValue = "PaaS"

	// CloudNetworkQueryFilterResourceTypeProcessingUnit represents the value ProcessingUnit.
	CloudNetworkQueryFilterResourceTypeProcessingUnit CloudNetworkQueryFilterResourceTypeValue = "ProcessingUnit"

	// CloudNetworkQueryFilterResourceTypeService represents the value Service.
	CloudNetworkQueryFilterResourceTypeService CloudNetworkQueryFilterResourceTypeValue = "Service"
)

// CloudNetworkQueryFilter represents the model of a cloudnetworkqueryfilter
type CloudNetworkQueryFilter struct {
	// The cluster name of the target K8s resources. Applies only to resourceType K8s.
	K8sClusterNames []string `json:"K8sClusterNames,omitempty" msgpack:"K8sClusterNames,omitempty" bson:"k8sclusternames,omitempty" mapstructure:"K8sClusterNames,omitempty"`

	// A list of K8s images that resources can be filtered with. Applies only to
	// resourceType K8s.
	K8sContainerImages []string `json:"K8sContainerImages,omitempty" msgpack:"K8sContainerImages,omitempty" bson:"k8scontainerimages,omitempty" mapstructure:"K8sContainerImages,omitempty"`

	// A list of labels that apply to the queried resource. Applies only to
	// resourceType K8s.
	K8sLabels []string `json:"K8sLabels,omitempty" msgpack:"K8sLabels,omitempty" bson:"k8slabels,omitempty" mapstructure:"K8sLabels,omitempty"`

	// The namespace of the target K8s resources. Applies only to resourceType K8s.
	K8sNamespaces []string `json:"K8sNamespaces,omitempty" msgpack:"K8sNamespaces,omitempty" bson:"k8snamespaces,omitempty" mapstructure:"K8sNamespaces,omitempty"`

	// The service name of the target K8s resources. Applies only to resourceType K8s.
	K8sServiceNames []string `json:"K8sServiceNames,omitempty" msgpack:"K8sServiceNames,omitempty" bson:"k8sservicenames,omitempty" mapstructure:"K8sServiceNames,omitempty"`

	// Identifies a list of K8s Service types. Applies only to resourceType K8s.
	K8sServiceTypes []string `json:"K8sServiceTypes,omitempty" msgpack:"K8sServiceTypes,omitempty" bson:"k8sservicetypes,omitempty" mapstructure:"K8sServiceTypes,omitempty"`

	// The VPC ID of the target resources.
	VPCIDs []string `json:"VPCIDs,omitempty" msgpack:"VPCIDs,omitempty" bson:"vpcids,omitempty" mapstructure:"VPCIDs,omitempty"`

	// The accounts that the search must apply to. These are the actually IDs of the
	// account as provided by the cloud provider. One or more IDs can be included.
	AccountIDs []string `json:"accountIDs,omitempty" msgpack:"accountIDs,omitempty" bson:"accountids,omitempty" mapstructure:"accountIDs,omitempty"`

	// The cloud types that the search must apply to.
	CloudTypes []string `json:"cloudTypes,omitempty" msgpack:"cloudTypes,omitempty" bson:"cloudtypes,omitempty" mapstructure:"cloudTypes,omitempty"`

	// A list of imageIDs that endpoints can be filtered with. Applies only to
	// resourceType Endpoint.
	ImageIDs []string `json:"imageIDs,omitempty" msgpack:"imageIDs,omitempty" bson:"imageids,omitempty" mapstructure:"imageIDs,omitempty"`

	// A list of tags that exclude the matching endpoints for the query. These tags
	// refer to the tags attached to the resources in the cloud provider definitions.
	NotTags []string `json:"notTags,omitempty" msgpack:"notTags,omitempty" bson:"nottags,omitempty" mapstructure:"notTags,omitempty"`

	// If set to true, the VPC IDs in `VPCIDs` will be excluded rather than included.
	NotVPCIDs bool `json:"notVPCIDs" msgpack:"notVPCIDs" bson:"-" mapstructure:"notVPCIDs,omitempty"`

	// The exact object that the search applies. If ObjectIDs are defined, the rest of
	// the fields are ignored. An object ID can refer to an instance, VPC endpoint, or
	// network interface.
	ObjectIDs []string `json:"objectIDs,omitempty" msgpack:"objectIDs,omitempty" bson:"objectids,omitempty" mapstructure:"objectIDs,omitempty"`

	// Identifies a list of Platform as a Service types.
	PaasTypes []string `json:"paasTypes,omitempty" msgpack:"paasTypes,omitempty" bson:"paastypes,omitempty" mapstructure:"paasTypes,omitempty"`

	// Restricts the query on only endpoints with the given productInfoType.
	ProductInfoType string `json:"productInfoType,omitempty" msgpack:"productInfoType,omitempty" bson:"productinfotype,omitempty" mapstructure:"productInfoType,omitempty"`

	// Restricts the query to only endpoints with the provided productInfoValue. Does
	// not apply to other resource types.
	ProductInfoValue string `json:"productInfoValue,omitempty" msgpack:"productInfoValue,omitempty" bson:"productinfovalue,omitempty" mapstructure:"productInfoValue,omitempty"`

	// The region that the search must apply to.
	Regions []string `json:"regions,omitempty" msgpack:"regions,omitempty" bson:"regions,omitempty" mapstructure:"regions,omitempty"`

	// The status of the resource.
	ResourceStatus string `json:"resourceStatus,omitempty" msgpack:"resourceStatus,omitempty" bson:"resourcestatus,omitempty" mapstructure:"resourceStatus,omitempty"`

	// The type of endpoint resource. The resource type is a mandatory field and a
	// query cannot span multiple resource types.
	ResourceType CloudNetworkQueryFilterResourceTypeValue `json:"resourceType" msgpack:"resourceType" bson:"resourcetype" mapstructure:"resourceType,omitempty"`

	// The list of security tags associated with the targets of the query. Security
	// tags refer to security groups in AWS or network tags in GCP. So they can have
	// different meaning depending on the target cloud.
	SecurityTags []string `json:"securityTags,omitempty" msgpack:"securityTags,omitempty" bson:"securitytags,omitempty" mapstructure:"securityTags,omitempty"`

	// Identifies a list of service names that should be taken into account. This is
	// only valid with a resource type equal to Service.
	ServiceNames []string `json:"serviceNames,omitempty" msgpack:"serviceNames,omitempty" bson:"servicenames,omitempty" mapstructure:"serviceNames,omitempty"`

	// Identifies the owner of the service that the resource is attached to. Field is
	// not valid if the resource type is not an interface.
	ServiceOwners []string `json:"serviceOwners,omitempty" msgpack:"serviceOwners,omitempty" bson:"serviceowners,omitempty" mapstructure:"serviceOwners,omitempty"`

	// Identifies the type of service that the interface is attached to. Field is not
	// valid if the resource type is not an
	// interface.
	ServiceTypes []string `json:"serviceTypes,omitempty" msgpack:"serviceTypes,omitempty" bson:"servicetypes,omitempty" mapstructure:"serviceTypes,omitempty"`

	// The subnets where the resources must reside. A subnet parameter can only be
	// provided for a network interface resource type.
	Subnets []string `json:"subnets,omitempty" msgpack:"subnets,omitempty" bson:"subnets,omitempty" mapstructure:"subnets,omitempty"`

	// A list of tags that select the same of endpoints for the query. These tags refer
	// to the tags attached to the resources in the cloud provider definitions.
	Tags []string `json:"tags,omitempty" msgpack:"tags,omitempty" bson:"tags,omitempty" mapstructure:"tags,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudNetworkQueryFilter returns a new *CloudNetworkQueryFilter
func NewCloudNetworkQueryFilter() *CloudNetworkQueryFilter {

	return &CloudNetworkQueryFilter{
		ModelVersion:       1,
		K8sClusterNames:    []string{},
		K8sContainerImages: []string{},
		K8sLabels:          []string{},
		K8sNamespaces:      []string{},
		K8sServiceNames:    []string{},
		K8sServiceTypes:    []string{},
		VPCIDs:             []string{},
		AccountIDs:         []string{},
		CloudTypes:         []string{},
		ImageIDs:           []string{},
		NotTags:            []string{},
		ObjectIDs:          []string{},
		PaasTypes:          []string{},
		Regions:            []string{},
		ResourceType:       CloudNetworkQueryFilterResourceTypeInstance,
		SecurityTags:       []string{},
		ServiceNames:       []string{},
		ServiceOwners:      []string{},
		ServiceTypes:       []string{},
		Subnets:            []string{},
		Tags:               []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudNetworkQueryFilter) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudNetworkQueryFilter{}

	s.K8sClusterNames = o.K8sClusterNames
	s.K8sContainerImages = o.K8sContainerImages
	s.K8sLabels = o.K8sLabels
	s.K8sNamespaces = o.K8sNamespaces
	s.K8sServiceNames = o.K8sServiceNames
	s.K8sServiceTypes = o.K8sServiceTypes
	s.VPCIDs = o.VPCIDs
	s.AccountIDs = o.AccountIDs
	s.CloudTypes = o.CloudTypes
	s.ImageIDs = o.ImageIDs
	s.NotTags = o.NotTags
	s.ObjectIDs = o.ObjectIDs
	s.PaasTypes = o.PaasTypes
	s.ProductInfoType = o.ProductInfoType
	s.ProductInfoValue = o.ProductInfoValue
	s.Regions = o.Regions
	s.ResourceStatus = o.ResourceStatus
	s.ResourceType = o.ResourceType
	s.SecurityTags = o.SecurityTags
	s.ServiceNames = o.ServiceNames
	s.ServiceOwners = o.ServiceOwners
	s.ServiceTypes = o.ServiceTypes
	s.Subnets = o.Subnets
	s.Tags = o.Tags

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudNetworkQueryFilter) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudNetworkQueryFilter{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.K8sClusterNames = s.K8sClusterNames
	o.K8sContainerImages = s.K8sContainerImages
	o.K8sLabels = s.K8sLabels
	o.K8sNamespaces = s.K8sNamespaces
	o.K8sServiceNames = s.K8sServiceNames
	o.K8sServiceTypes = s.K8sServiceTypes
	o.VPCIDs = s.VPCIDs
	o.AccountIDs = s.AccountIDs
	o.CloudTypes = s.CloudTypes
	o.ImageIDs = s.ImageIDs
	o.NotTags = s.NotTags
	o.ObjectIDs = s.ObjectIDs
	o.PaasTypes = s.PaasTypes
	o.ProductInfoType = s.ProductInfoType
	o.ProductInfoValue = s.ProductInfoValue
	o.Regions = s.Regions
	o.ResourceStatus = s.ResourceStatus
	o.ResourceType = s.ResourceType
	o.SecurityTags = s.SecurityTags
	o.ServiceNames = s.ServiceNames
	o.ServiceOwners = s.ServiceOwners
	o.ServiceTypes = s.ServiceTypes
	o.Subnets = s.Subnets
	o.Tags = s.Tags

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudNetworkQueryFilter) BleveType() string {

	return "cloudnetworkqueryfilter"
}

// DeepCopy returns a deep copy if the CloudNetworkQueryFilter.
func (o *CloudNetworkQueryFilter) DeepCopy() *CloudNetworkQueryFilter {

	if o == nil {
		return nil
	}

	out := &CloudNetworkQueryFilter{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudNetworkQueryFilter.
func (o *CloudNetworkQueryFilter) DeepCopyInto(out *CloudNetworkQueryFilter) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudNetworkQueryFilter: %s", err))
	}

	*out = *target.(*CloudNetworkQueryFilter)
}

// Validate valides the current information stored into the structure.
func (o *CloudNetworkQueryFilter) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("resourceType", string(o.ResourceType)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("resourceType", string(o.ResourceType), []string{"Instance", "Interface", "Service", "ProcessingUnit", "PaaS", "K8sService"}, false); err != nil {
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
func (*CloudNetworkQueryFilter) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudNetworkQueryFilterAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudNetworkQueryFilterLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudNetworkQueryFilter) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudNetworkQueryFilterAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudNetworkQueryFilter) ValueForAttribute(name string) any {

	switch name {
	case "K8sClusterNames":
		return o.K8sClusterNames
	case "K8sContainerImages":
		return o.K8sContainerImages
	case "K8sLabels":
		return o.K8sLabels
	case "K8sNamespaces":
		return o.K8sNamespaces
	case "K8sServiceNames":
		return o.K8sServiceNames
	case "K8sServiceTypes":
		return o.K8sServiceTypes
	case "VPCIDs":
		return o.VPCIDs
	case "accountIDs":
		return o.AccountIDs
	case "cloudTypes":
		return o.CloudTypes
	case "imageIDs":
		return o.ImageIDs
	case "notTags":
		return o.NotTags
	case "notVPCIDs":
		return o.NotVPCIDs
	case "objectIDs":
		return o.ObjectIDs
	case "paasTypes":
		return o.PaasTypes
	case "productInfoType":
		return o.ProductInfoType
	case "productInfoValue":
		return o.ProductInfoValue
	case "regions":
		return o.Regions
	case "resourceStatus":
		return o.ResourceStatus
	case "resourceType":
		return o.ResourceType
	case "securityTags":
		return o.SecurityTags
	case "serviceNames":
		return o.ServiceNames
	case "serviceOwners":
		return o.ServiceOwners
	case "serviceTypes":
		return o.ServiceTypes
	case "subnets":
		return o.Subnets
	case "tags":
		return o.Tags
	}

	return nil
}

// CloudNetworkQueryFilterAttributesMap represents the map of attribute for CloudNetworkQueryFilter.
var CloudNetworkQueryFilterAttributesMap = map[string]elemental.AttributeSpecification{
	"K8sClusterNames": {
		AllowedChoices: []string{},
		BSONFieldName:  "k8sclusternames",
		ConvertedName:  "K8sClusterNames",
		Description:    `The cluster name of the target K8s resources. Applies only to resourceType K8s.`,
		Exposed:        true,
		Name:           "K8sClusterNames",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"K8sContainerImages": {
		AllowedChoices: []string{},
		BSONFieldName:  "k8scontainerimages",
		ConvertedName:  "K8sContainerImages",
		Description: `A list of K8s images that resources can be filtered with. Applies only to
resourceType K8s.`,
		Exposed: true,
		Name:    "K8sContainerImages",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"K8sLabels": {
		AllowedChoices: []string{},
		BSONFieldName:  "k8slabels",
		ConvertedName:  "K8sLabels",
		Description: `A list of labels that apply to the queried resource. Applies only to
resourceType K8s.`,
		Exposed: true,
		Name:    "K8sLabels",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"K8sNamespaces": {
		AllowedChoices: []string{},
		BSONFieldName:  "k8snamespaces",
		ConvertedName:  "K8sNamespaces",
		Description:    `The namespace of the target K8s resources. Applies only to resourceType K8s.`,
		Exposed:        true,
		Name:           "K8sNamespaces",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"K8sServiceNames": {
		AllowedChoices: []string{},
		BSONFieldName:  "k8sservicenames",
		ConvertedName:  "K8sServiceNames",
		Description:    `The service name of the target K8s resources. Applies only to resourceType K8s.`,
		Exposed:        true,
		Name:           "K8sServiceNames",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"K8sServiceTypes": {
		AllowedChoices: []string{},
		BSONFieldName:  "k8sservicetypes",
		ConvertedName:  "K8sServiceTypes",
		Description:    `Identifies a list of K8s Service types. Applies only to resourceType K8s.`,
		Exposed:        true,
		Name:           "K8sServiceTypes",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"VPCIDs": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcids",
		ConvertedName:  "VPCIDs",
		Description:    `The VPC ID of the target resources.`,
		Exposed:        true,
		Name:           "VPCIDs",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"AccountIDs": {
		AllowedChoices: []string{},
		BSONFieldName:  "accountids",
		ConvertedName:  "AccountIDs",
		Description: `The accounts that the search must apply to. These are the actually IDs of the
account as provided by the cloud provider. One or more IDs can be included.`,
		Exposed: true,
		Name:    "accountIDs",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"CloudTypes": {
		AllowedChoices: []string{},
		BSONFieldName:  "cloudtypes",
		ConvertedName:  "CloudTypes",
		Description:    `The cloud types that the search must apply to.`,
		Exposed:        true,
		Name:           "cloudTypes",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"ImageIDs": {
		AllowedChoices: []string{},
		BSONFieldName:  "imageids",
		ConvertedName:  "ImageIDs",
		Description: `A list of imageIDs that endpoints can be filtered with. Applies only to
resourceType Endpoint.`,
		Exposed: true,
		Name:    "imageIDs",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"NotTags": {
		AllowedChoices: []string{},
		BSONFieldName:  "nottags",
		ConvertedName:  "NotTags",
		Description: `A list of tags that exclude the matching endpoints for the query. These tags
refer to the tags attached to the resources in the cloud provider definitions.`,
		Exposed: true,
		Name:    "notTags",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"NotVPCIDs": {
		AllowedChoices: []string{},
		ConvertedName:  "NotVPCIDs",
		Description:    `If set to true, the VPC IDs in ` + "`" + `VPCIDs` + "`" + ` will be excluded rather than included.`,
		Exposed:        true,
		Name:           "notVPCIDs",
		Type:           "boolean",
	},
	"ObjectIDs": {
		AllowedChoices: []string{},
		BSONFieldName:  "objectids",
		ConvertedName:  "ObjectIDs",
		Description: `The exact object that the search applies. If ObjectIDs are defined, the rest of
the fields are ignored. An object ID can refer to an instance, VPC endpoint, or
network interface.`,
		Exposed: true,
		Name:    "objectIDs",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"PaasTypes": {
		AllowedChoices: []string{},
		BSONFieldName:  "paastypes",
		ConvertedName:  "PaasTypes",
		Description:    `Identifies a list of Platform as a Service types.`,
		Exposed:        true,
		Name:           "paasTypes",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"ProductInfoType": {
		AllowedChoices: []string{},
		BSONFieldName:  "productinfotype",
		ConvertedName:  "ProductInfoType",
		Description:    `Restricts the query on only endpoints with the given productInfoType.`,
		Exposed:        true,
		Name:           "productInfoType",
		Stored:         true,
		Type:           "string",
	},
	"ProductInfoValue": {
		AllowedChoices: []string{},
		BSONFieldName:  "productinfovalue",
		ConvertedName:  "ProductInfoValue",
		Description: `Restricts the query to only endpoints with the provided productInfoValue. Does
not apply to other resource types.`,
		Exposed: true,
		Name:    "productInfoValue",
		Stored:  true,
		Type:    "string",
	},
	"Regions": {
		AllowedChoices: []string{},
		BSONFieldName:  "regions",
		ConvertedName:  "Regions",
		Description:    `The region that the search must apply to.`,
		Exposed:        true,
		Name:           "regions",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"ResourceStatus": {
		AllowedChoices: []string{},
		BSONFieldName:  "resourcestatus",
		ConvertedName:  "ResourceStatus",
		Description:    `The status of the resource.`,
		Exposed:        true,
		Name:           "resourceStatus",
		Stored:         true,
		Type:           "string",
	},
	"ResourceType": {
		AllowedChoices: []string{"Instance", "Interface", "Service", "ProcessingUnit", "PaaS", "K8sService"},
		BSONFieldName:  "resourcetype",
		ConvertedName:  "ResourceType",
		DefaultValue:   CloudNetworkQueryFilterResourceTypeInstance,
		Description: `The type of endpoint resource. The resource type is a mandatory field and a
query cannot span multiple resource types.`,
		Exposed:  true,
		Name:     "resourceType",
		Required: true,
		Stored:   true,
		Type:     "enum",
	},
	"SecurityTags": {
		AllowedChoices: []string{},
		BSONFieldName:  "securitytags",
		ConvertedName:  "SecurityTags",
		Description: `The list of security tags associated with the targets of the query. Security
tags refer to security groups in AWS or network tags in GCP. So they can have
different meaning depending on the target cloud.`,
		Exposed: true,
		Name:    "securityTags",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"ServiceNames": {
		AllowedChoices: []string{},
		BSONFieldName:  "servicenames",
		ConvertedName:  "ServiceNames",
		Description: `Identifies a list of service names that should be taken into account. This is
only valid with a resource type equal to Service.`,
		Exposed: true,
		Name:    "serviceNames",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"ServiceOwners": {
		AllowedChoices: []string{},
		BSONFieldName:  "serviceowners",
		ConvertedName:  "ServiceOwners",
		Description: `Identifies the owner of the service that the resource is attached to. Field is
not valid if the resource type is not an interface.`,
		Exposed: true,
		Name:    "serviceOwners",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"ServiceTypes": {
		AllowedChoices: []string{},
		BSONFieldName:  "servicetypes",
		ConvertedName:  "ServiceTypes",
		Description: `Identifies the type of service that the interface is attached to. Field is not
valid if the resource type is not an
interface.`,
		Exposed: true,
		Name:    "serviceTypes",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"Subnets": {
		AllowedChoices: []string{},
		BSONFieldName:  "subnets",
		ConvertedName:  "Subnets",
		Description: `The subnets where the resources must reside. A subnet parameter can only be
provided for a network interface resource type.`,
		Exposed: true,
		Name:    "subnets",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"Tags": {
		AllowedChoices: []string{},
		BSONFieldName:  "tags",
		ConvertedName:  "Tags",
		Description: `A list of tags that select the same of endpoints for the query. These tags refer
to the tags attached to the resources in the cloud provider definitions.`,
		Exposed: true,
		Name:    "tags",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
}

// CloudNetworkQueryFilterLowerCaseAttributesMap represents the map of attribute for CloudNetworkQueryFilter.
var CloudNetworkQueryFilterLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"k8sclusternames": {
		AllowedChoices: []string{},
		BSONFieldName:  "k8sclusternames",
		ConvertedName:  "K8sClusterNames",
		Description:    `The cluster name of the target K8s resources. Applies only to resourceType K8s.`,
		Exposed:        true,
		Name:           "K8sClusterNames",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"k8scontainerimages": {
		AllowedChoices: []string{},
		BSONFieldName:  "k8scontainerimages",
		ConvertedName:  "K8sContainerImages",
		Description: `A list of K8s images that resources can be filtered with. Applies only to
resourceType K8s.`,
		Exposed: true,
		Name:    "K8sContainerImages",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"k8slabels": {
		AllowedChoices: []string{},
		BSONFieldName:  "k8slabels",
		ConvertedName:  "K8sLabels",
		Description: `A list of labels that apply to the queried resource. Applies only to
resourceType K8s.`,
		Exposed: true,
		Name:    "K8sLabels",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"k8snamespaces": {
		AllowedChoices: []string{},
		BSONFieldName:  "k8snamespaces",
		ConvertedName:  "K8sNamespaces",
		Description:    `The namespace of the target K8s resources. Applies only to resourceType K8s.`,
		Exposed:        true,
		Name:           "K8sNamespaces",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"k8sservicenames": {
		AllowedChoices: []string{},
		BSONFieldName:  "k8sservicenames",
		ConvertedName:  "K8sServiceNames",
		Description:    `The service name of the target K8s resources. Applies only to resourceType K8s.`,
		Exposed:        true,
		Name:           "K8sServiceNames",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"k8sservicetypes": {
		AllowedChoices: []string{},
		BSONFieldName:  "k8sservicetypes",
		ConvertedName:  "K8sServiceTypes",
		Description:    `Identifies a list of K8s Service types. Applies only to resourceType K8s.`,
		Exposed:        true,
		Name:           "K8sServiceTypes",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"vpcids": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcids",
		ConvertedName:  "VPCIDs",
		Description:    `The VPC ID of the target resources.`,
		Exposed:        true,
		Name:           "VPCIDs",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"accountids": {
		AllowedChoices: []string{},
		BSONFieldName:  "accountids",
		ConvertedName:  "AccountIDs",
		Description: `The accounts that the search must apply to. These are the actually IDs of the
account as provided by the cloud provider. One or more IDs can be included.`,
		Exposed: true,
		Name:    "accountIDs",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"cloudtypes": {
		AllowedChoices: []string{},
		BSONFieldName:  "cloudtypes",
		ConvertedName:  "CloudTypes",
		Description:    `The cloud types that the search must apply to.`,
		Exposed:        true,
		Name:           "cloudTypes",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"imageids": {
		AllowedChoices: []string{},
		BSONFieldName:  "imageids",
		ConvertedName:  "ImageIDs",
		Description: `A list of imageIDs that endpoints can be filtered with. Applies only to
resourceType Endpoint.`,
		Exposed: true,
		Name:    "imageIDs",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"nottags": {
		AllowedChoices: []string{},
		BSONFieldName:  "nottags",
		ConvertedName:  "NotTags",
		Description: `A list of tags that exclude the matching endpoints for the query. These tags
refer to the tags attached to the resources in the cloud provider definitions.`,
		Exposed: true,
		Name:    "notTags",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"notvpcids": {
		AllowedChoices: []string{},
		ConvertedName:  "NotVPCIDs",
		Description:    `If set to true, the VPC IDs in ` + "`" + `VPCIDs` + "`" + ` will be excluded rather than included.`,
		Exposed:        true,
		Name:           "notVPCIDs",
		Type:           "boolean",
	},
	"objectids": {
		AllowedChoices: []string{},
		BSONFieldName:  "objectids",
		ConvertedName:  "ObjectIDs",
		Description: `The exact object that the search applies. If ObjectIDs are defined, the rest of
the fields are ignored. An object ID can refer to an instance, VPC endpoint, or
network interface.`,
		Exposed: true,
		Name:    "objectIDs",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"paastypes": {
		AllowedChoices: []string{},
		BSONFieldName:  "paastypes",
		ConvertedName:  "PaasTypes",
		Description:    `Identifies a list of Platform as a Service types.`,
		Exposed:        true,
		Name:           "paasTypes",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"productinfotype": {
		AllowedChoices: []string{},
		BSONFieldName:  "productinfotype",
		ConvertedName:  "ProductInfoType",
		Description:    `Restricts the query on only endpoints with the given productInfoType.`,
		Exposed:        true,
		Name:           "productInfoType",
		Stored:         true,
		Type:           "string",
	},
	"productinfovalue": {
		AllowedChoices: []string{},
		BSONFieldName:  "productinfovalue",
		ConvertedName:  "ProductInfoValue",
		Description: `Restricts the query to only endpoints with the provided productInfoValue. Does
not apply to other resource types.`,
		Exposed: true,
		Name:    "productInfoValue",
		Stored:  true,
		Type:    "string",
	},
	"regions": {
		AllowedChoices: []string{},
		BSONFieldName:  "regions",
		ConvertedName:  "Regions",
		Description:    `The region that the search must apply to.`,
		Exposed:        true,
		Name:           "regions",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"resourcestatus": {
		AllowedChoices: []string{},
		BSONFieldName:  "resourcestatus",
		ConvertedName:  "ResourceStatus",
		Description:    `The status of the resource.`,
		Exposed:        true,
		Name:           "resourceStatus",
		Stored:         true,
		Type:           "string",
	},
	"resourcetype": {
		AllowedChoices: []string{"Instance", "Interface", "Service", "ProcessingUnit", "PaaS", "K8sService"},
		BSONFieldName:  "resourcetype",
		ConvertedName:  "ResourceType",
		DefaultValue:   CloudNetworkQueryFilterResourceTypeInstance,
		Description: `The type of endpoint resource. The resource type is a mandatory field and a
query cannot span multiple resource types.`,
		Exposed:  true,
		Name:     "resourceType",
		Required: true,
		Stored:   true,
		Type:     "enum",
	},
	"securitytags": {
		AllowedChoices: []string{},
		BSONFieldName:  "securitytags",
		ConvertedName:  "SecurityTags",
		Description: `The list of security tags associated with the targets of the query. Security
tags refer to security groups in AWS or network tags in GCP. So they can have
different meaning depending on the target cloud.`,
		Exposed: true,
		Name:    "securityTags",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"servicenames": {
		AllowedChoices: []string{},
		BSONFieldName:  "servicenames",
		ConvertedName:  "ServiceNames",
		Description: `Identifies a list of service names that should be taken into account. This is
only valid with a resource type equal to Service.`,
		Exposed: true,
		Name:    "serviceNames",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"serviceowners": {
		AllowedChoices: []string{},
		BSONFieldName:  "serviceowners",
		ConvertedName:  "ServiceOwners",
		Description: `Identifies the owner of the service that the resource is attached to. Field is
not valid if the resource type is not an interface.`,
		Exposed: true,
		Name:    "serviceOwners",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"servicetypes": {
		AllowedChoices: []string{},
		BSONFieldName:  "servicetypes",
		ConvertedName:  "ServiceTypes",
		Description: `Identifies the type of service that the interface is attached to. Field is not
valid if the resource type is not an
interface.`,
		Exposed: true,
		Name:    "serviceTypes",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"subnets": {
		AllowedChoices: []string{},
		BSONFieldName:  "subnets",
		ConvertedName:  "Subnets",
		Description: `The subnets where the resources must reside. A subnet parameter can only be
provided for a network interface resource type.`,
		Exposed: true,
		Name:    "subnets",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"tags": {
		AllowedChoices: []string{},
		BSONFieldName:  "tags",
		ConvertedName:  "Tags",
		Description: `A list of tags that select the same of endpoints for the query. These tags refer
to the tags attached to the resources in the cloud provider definitions.`,
		Exposed: true,
		Name:    "tags",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
}

type mongoAttributesCloudNetworkQueryFilter struct {
	K8sClusterNames    []string                                 `bson:"k8sclusternames,omitempty"`
	K8sContainerImages []string                                 `bson:"k8scontainerimages,omitempty"`
	K8sLabels          []string                                 `bson:"k8slabels,omitempty"`
	K8sNamespaces      []string                                 `bson:"k8snamespaces,omitempty"`
	K8sServiceNames    []string                                 `bson:"k8sservicenames,omitempty"`
	K8sServiceTypes    []string                                 `bson:"k8sservicetypes,omitempty"`
	VPCIDs             []string                                 `bson:"vpcids,omitempty"`
	AccountIDs         []string                                 `bson:"accountids,omitempty"`
	CloudTypes         []string                                 `bson:"cloudtypes,omitempty"`
	ImageIDs           []string                                 `bson:"imageids,omitempty"`
	NotTags            []string                                 `bson:"nottags,omitempty"`
	ObjectIDs          []string                                 `bson:"objectids,omitempty"`
	PaasTypes          []string                                 `bson:"paastypes,omitempty"`
	ProductInfoType    string                                   `bson:"productinfotype,omitempty"`
	ProductInfoValue   string                                   `bson:"productinfovalue,omitempty"`
	Regions            []string                                 `bson:"regions,omitempty"`
	ResourceStatus     string                                   `bson:"resourcestatus,omitempty"`
	ResourceType       CloudNetworkQueryFilterResourceTypeValue `bson:"resourcetype"`
	SecurityTags       []string                                 `bson:"securitytags,omitempty"`
	ServiceNames       []string                                 `bson:"servicenames,omitempty"`
	ServiceOwners      []string                                 `bson:"serviceowners,omitempty"`
	ServiceTypes       []string                                 `bson:"servicetypes,omitempty"`
	Subnets            []string                                 `bson:"subnets,omitempty"`
	Tags               []string                                 `bson:"tags,omitempty"`
}
