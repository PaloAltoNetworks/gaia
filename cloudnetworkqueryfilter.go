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

	// CloudNetworkQueryFilterResourceTypeProcessingUnit represents the value ProcessingUnit.
	CloudNetworkQueryFilterResourceTypeProcessingUnit CloudNetworkQueryFilterResourceTypeValue = "ProcessingUnit"

	// CloudNetworkQueryFilterResourceTypeService represents the value Service.
	CloudNetworkQueryFilterResourceTypeService CloudNetworkQueryFilterResourceTypeValue = "Service"
)

// CloudNetworkQueryFilter represents the model of a cloudnetworkqueryfilter
type CloudNetworkQueryFilter struct {
	// The VPC ID of the target resources.
	VPCIDs []string `json:"VPCIDs,omitempty" msgpack:"VPCIDs,omitempty" bson:"vpcids" mapstructure:"VPCIDs,omitempty"`

	// The accounts that the search must apply to. These are the actually IDs of the
	// account as provided by the cloud provider. One or more IDs can be included.
	AccountIDs []string `json:"accountIDs,omitempty" msgpack:"accountIDs,omitempty" bson:"accountids" mapstructure:"accountIDs,omitempty"`

	// The cloud types that the search must apply to.
	CloudTypes []string `json:"cloudTypes,omitempty" msgpack:"cloudTypes,omitempty" bson:"cloudtypes" mapstructure:"cloudTypes,omitempty"`

	// A list of imageIDs that endpoints can be filtered with. Applies only to
	// resourceType Endpoint.
	ImageIDs []string `json:"imageIDs,omitempty" msgpack:"imageIDs,omitempty" bson:"imageids" mapstructure:"imageIDs,omitempty"`

	// The exact object that the search applies. If ObjectIDs are defined, the rest of
	// the fields are ignored. An object ID can refer to an instance, VPC endpoint, or
	// network interface.
	ObjectIDs []string `json:"objectIDs,omitempty" msgpack:"objectIDs,omitempty" bson:"objectids" mapstructure:"objectIDs,omitempty"`

	// Restricts the query on only endpoints with the given productInfoType.
	ProductInfoType string `json:"productInfoType,omitempty" msgpack:"productInfoType,omitempty" bson:"productinfotype" mapstructure:"productInfoType,omitempty"`

	// Restricts the query to only endpoints with the provided productInfoValue. Does
	// not apply to other resource types.
	ProductInfoValue string `json:"productInfoValue,omitempty" msgpack:"productInfoValue,omitempty" bson:"productinfovalue" mapstructure:"productInfoValue,omitempty"`

	// The region that the search must apply to.
	Regions []string `json:"regions,omitempty" msgpack:"regions,omitempty" bson:"regions" mapstructure:"regions,omitempty"`

	// The type of endpoint resource. The resource type is a mandatory field and a
	// query cannot span multiple resource types.
	ResourceType CloudNetworkQueryFilterResourceTypeValue `json:"resourceType" msgpack:"resourceType" bson:"resourcetype" mapstructure:"resourceType,omitempty"`

	// The list of security tags associated with the targets of the query. Security
	// tags refer to security groups in AWS or network tags in GCP. So they can have
	// different meaning depending on the target cloud.
	SecurityTags []string `json:"securityTags,omitempty" msgpack:"securityTags,omitempty" bson:"securitytags" mapstructure:"securityTags,omitempty"`

	// Identifies a list of service names that should be taken into account. This is
	// only valid with a resource type equal to Service.
	ServiceNames []string `json:"serviceNames,omitempty" msgpack:"serviceNames,omitempty" bson:"servicenames" mapstructure:"serviceNames,omitempty"`

	// Identifies the owner of the service that the resource is attached to. Field is
	// not valid if the resource type is not an interface.
	ServiceOwners []string `json:"serviceOwners,omitempty" msgpack:"serviceOwners,omitempty" bson:"serviceowners" mapstructure:"serviceOwners,omitempty"`

	// Identifies the type of service that the interface is attached to. Field is not
	// valid if the resource type is not an
	// interface.
	ServiceTypes []string `json:"serviceTypes,omitempty" msgpack:"serviceTypes,omitempty" bson:"servicetypes" mapstructure:"serviceTypes,omitempty"`

	// The subnets where the resources must reside. A subnet parameter can only be
	// provided for a network interface resource type.
	Subnets []string `json:"subnets,omitempty" msgpack:"subnets,omitempty" bson:"subnets" mapstructure:"subnets,omitempty"`

	// A list of tags that select the same of endpoints for the query. These tags refer
	// to the tags attached to the resources in the cloud provider definitions.
	Tags []string `json:"tags,omitempty" msgpack:"tags,omitempty" bson:"tags" mapstructure:"tags,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudNetworkQueryFilter returns a new *CloudNetworkQueryFilter
func NewCloudNetworkQueryFilter() *CloudNetworkQueryFilter {

	return &CloudNetworkQueryFilter{
		ModelVersion:  1,
		AccountIDs:    []string{},
		CloudTypes:    []string{},
		ImageIDs:      []string{},
		ObjectIDs:     []string{},
		ServiceOwners: []string{},
		Regions:       []string{},
		ResourceType:  CloudNetworkQueryFilterResourceTypeInstance,
		SecurityTags:  []string{},
		ServiceNames:  []string{},
		ServiceTypes:  []string{},
		Subnets:       []string{},
		Tags:          []string{},
		VPCIDs:        []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudNetworkQueryFilter) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudNetworkQueryFilter{}

	s.VPCIDs = o.VPCIDs
	s.AccountIDs = o.AccountIDs
	s.CloudTypes = o.CloudTypes
	s.ImageIDs = o.ImageIDs
	s.ObjectIDs = o.ObjectIDs
	s.ProductInfoType = o.ProductInfoType
	s.ProductInfoValue = o.ProductInfoValue
	s.Regions = o.Regions
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

	o.VPCIDs = s.VPCIDs
	o.AccountIDs = s.AccountIDs
	o.CloudTypes = s.CloudTypes
	o.ImageIDs = s.ImageIDs
	o.ObjectIDs = s.ObjectIDs
	o.ProductInfoType = s.ProductInfoType
	o.ProductInfoValue = s.ProductInfoValue
	o.Regions = s.Regions
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

	if err := elemental.ValidateStringInList("resourceType", string(o.ResourceType), []string{"Instance", "Interface", "Service", "ProcessingUnit"}, false); err != nil {
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

type mongoAttributesCloudNetworkQueryFilter struct {
	VPCIDs           []string                                 `bson:"vpcids,omitempty"`
	AccountIDs       []string                                 `bson:"accountids,omitempty"`
	CloudTypes       []string                                 `bson:"cloudtypes,omitempty"`
	ImageIDs         []string                                 `bson:"imageids,omitempty"`
	ObjectIDs        []string                                 `bson:"objectids,omitempty"`
	ProductInfoType  string                                   `bson:"productinfotype,omitempty"`
	ProductInfoValue string                                   `bson:"productinfovalue,omitempty"`
	Regions          []string                                 `bson:"regions,omitempty"`
	ResourceType     CloudNetworkQueryFilterResourceTypeValue `bson:"resourcetype"`
	SecurityTags     []string                                 `bson:"securitytags,omitempty"`
	ServiceNames     []string                                 `bson:"servicenames,omitempty"`
	ServiceOwners    []string                                 `bson:"serviceowners,omitempty"`
	ServiceTypes     []string                                 `bson:"servicetypes,omitempty"`
	Subnets          []string                                 `bson:"subnets,omitempty"`
	Tags             []string                                 `bson:"tags,omitempty"`
}
