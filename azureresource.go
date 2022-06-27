package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AzureResourceKindValue represents the possible values for attribute "kind".
type AzureResourceKindValue string

const (
	// AzureResourceKindBackendAddressPool represents the value BackendAddressPool.
	AzureResourceKindBackendAddressPool AzureResourceKindValue = "BackendAddressPool"

	// AzureResourceKindIPConfiguration represents the value IPConfiguration.
	AzureResourceKindIPConfiguration AzureResourceKindValue = "IPConfiguration"

	// AzureResourceKindLoadBalancer represents the value LoadBalancer.
	AzureResourceKindLoadBalancer AzureResourceKindValue = "LoadBalancer"

	// AzureResourceKindNATGateway represents the value NATGateway.
	AzureResourceKindNATGateway AzureResourceKindValue = "NATGateway"

	// AzureResourceKindNetworkInterface represents the value NetworkInterface.
	AzureResourceKindNetworkInterface AzureResourceKindValue = "NetworkInterface"

	// AzureResourceKindNetworkSecurityGroup represents the value NetworkSecurityGroup.
	AzureResourceKindNetworkSecurityGroup AzureResourceKindValue = "NetworkSecurityGroup"

	// AzureResourceKindPublicIPAddress represents the value PublicIPAddress.
	AzureResourceKindPublicIPAddress AzureResourceKindValue = "PublicIPAddress"

	// AzureResourceKindPublicIPPrefix represents the value PublicIPPrefix.
	AzureResourceKindPublicIPPrefix AzureResourceKindValue = "PublicIPPrefix"

	// AzureResourceKindSubnet represents the value Subnet.
	AzureResourceKindSubnet AzureResourceKindValue = "Subnet"

	// AzureResourceKindVirtualMachine represents the value VirtualMachine.
	AzureResourceKindVirtualMachine AzureResourceKindValue = "VirtualMachine"

	// AzureResourceKindVirtualMachineScaleSet represents the value VirtualMachineScaleSet.
	AzureResourceKindVirtualMachineScaleSet AzureResourceKindValue = "VirtualMachineScaleSet"

	// AzureResourceKindVirtualMachineScaleSetVM represents the value VirtualMachineScaleSetVM.
	AzureResourceKindVirtualMachineScaleSetVM AzureResourceKindValue = "VirtualMachineScaleSetVM"

	// AzureResourceKindVirtualNetwork represents the value VirtualNetwork.
	AzureResourceKindVirtualNetwork AzureResourceKindValue = "VirtualNetwork"
)

// AzureResourceProviderValue represents the possible values for attribute "provider".
type AzureResourceProviderValue string

const (
	// AzureResourceProviderMicrosoftCompute represents the value MicrosoftCompute.
	AzureResourceProviderMicrosoftCompute AzureResourceProviderValue = "MicrosoftCompute"

	// AzureResourceProviderMicrosoftNetwork represents the value MicrosoftNetwork.
	AzureResourceProviderMicrosoftNetwork AzureResourceProviderValue = "MicrosoftNetwork"
)

// AzureResourceIdentity represents the Identity of the object.
var AzureResourceIdentity = elemental.Identity{
	Name:     "azureresource",
	Category: "azureresources",
	Package:  "pandemona",
	Private:  true,
}

// AzureResourcesList represents a list of AzureResources
type AzureResourcesList []*AzureResource

// Identity returns the identity of the objects in the list.
func (o AzureResourcesList) Identity() elemental.Identity {

	return AzureResourceIdentity
}

// Copy returns a pointer to a copy the AzureResourcesList.
func (o AzureResourcesList) Copy() elemental.Identifiables {

	copy := append(AzureResourcesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AzureResourcesList.
func (o AzureResourcesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AzureResourcesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AzureResource))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AzureResourcesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AzureResourcesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the AzureResourcesList converted to SparseAzureResourcesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AzureResourcesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAzureResourcesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAzureResource)
	}

	return out
}

// Version returns the version of the content.
func (o AzureResourcesList) Version() int {

	return 1
}

// AzureResource represents the model of a azureresource
type AzureResource struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// The json-encoded data that represents the resource.
	Data []byte `json:"data" msgpack:"data" bson:"data" mapstructure:"data,omitempty"`

	// The specific kind of the resource.
	Kind AzureResourceKindValue `json:"kind" msgpack:"kind" bson:"kind" mapstructure:"kind,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// The name of this resource.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// The major type of the resource.
	Provider AzureResourceProviderValue `json:"provider" msgpack:"provider" bson:"provider" mapstructure:"provider,omitempty"`

	// The name of the logical subcontainer of cloud resources.
	ResourceGroup string `json:"resourceGroup" msgpack:"resourceGroup" bson:"resourcegroup" mapstructure:"resourceGroup,omitempty"`

	// The identifier of the resource as presented by Azure, which is a path.
	ResourceID string `json:"resourceID" msgpack:"resourceID" bson:"resourceid" mapstructure:"resourceID,omitempty"`

	// The logical ID of the container which contains the cloud resources.
	SubscriptionID string `json:"subscriptionID" msgpack:"subscriptionID" bson:"subscriptionid" mapstructure:"subscriptionID,omitempty"`

	// Contextual values that can be used to narrow searching of resources if the
	// resourceID is not known. For instance, it could be used to store a resource's
	// location or public IP addresses to support cross-cloud analysis.
	Tags []string `json:"tags" msgpack:"tags" bson:"tags" mapstructure:"tags,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAzureResource returns a new *AzureResource
func NewAzureResource() *AzureResource {

	return &AzureResource{
		ModelVersion:  1,
		Data:          []byte{},
		MigrationsLog: map[string]string{},
		Tags:          []string{},
	}
}

// Identity returns the Identity of the object.
func (o *AzureResource) Identity() elemental.Identity {

	return AzureResourceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AzureResource) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AzureResource) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AzureResource) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesAzureResource{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Data = o.Data
	s.Kind = o.Kind
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.Provider = o.Provider
	s.ResourceGroup = o.ResourceGroup
	s.ResourceID = o.ResourceID
	s.SubscriptionID = o.SubscriptionID
	s.Tags = o.Tags
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AzureResource) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesAzureResource{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Data = s.Data
	o.Kind = s.Kind
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.Provider = s.Provider
	o.ResourceGroup = s.ResourceGroup
	o.ResourceID = s.ResourceID
	o.SubscriptionID = s.SubscriptionID
	o.Tags = s.Tags
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *AzureResource) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *AzureResource) BleveType() string {

	return "azureresource"
}

// DefaultOrder returns the list of default ordering fields.
func (o *AzureResource) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *AzureResource) Doc() string {

	return `Represents an Azure cloud resource such as virtualMachines and subnets.`
}

func (o *AzureResource) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *AzureResource) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *AzureResource) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *AzureResource) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *AzureResource) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetZHash returns the ZHash of the receiver.
func (o *AzureResource) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *AzureResource) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *AzureResource) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *AzureResource) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *AzureResource) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAzureResource{
			ID:             &o.ID,
			Data:           &o.Data,
			Kind:           &o.Kind,
			MigrationsLog:  &o.MigrationsLog,
			Name:           &o.Name,
			Namespace:      &o.Namespace,
			Provider:       &o.Provider,
			ResourceGroup:  &o.ResourceGroup,
			ResourceID:     &o.ResourceID,
			SubscriptionID: &o.SubscriptionID,
			Tags:           &o.Tags,
			ZHash:          &o.ZHash,
			Zone:           &o.Zone,
		}
	}

	sp := &SparseAzureResource{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "data":
			sp.Data = &(o.Data)
		case "kind":
			sp.Kind = &(o.Kind)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "provider":
			sp.Provider = &(o.Provider)
		case "resourceGroup":
			sp.ResourceGroup = &(o.ResourceGroup)
		case "resourceID":
			sp.ResourceID = &(o.ResourceID)
		case "subscriptionID":
			sp.SubscriptionID = &(o.SubscriptionID)
		case "tags":
			sp.Tags = &(o.Tags)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseAzureResource to the object.
func (o *AzureResource) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAzureResource)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Data != nil {
		o.Data = *so.Data
	}
	if so.Kind != nil {
		o.Kind = *so.Kind
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.Provider != nil {
		o.Provider = *so.Provider
	}
	if so.ResourceGroup != nil {
		o.ResourceGroup = *so.ResourceGroup
	}
	if so.ResourceID != nil {
		o.ResourceID = *so.ResourceID
	}
	if so.SubscriptionID != nil {
		o.SubscriptionID = *so.SubscriptionID
	}
	if so.Tags != nil {
		o.Tags = *so.Tags
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the AzureResource.
func (o *AzureResource) DeepCopy() *AzureResource {

	if o == nil {
		return nil
	}

	out := &AzureResource{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AzureResource.
func (o *AzureResource) DeepCopyInto(out *AzureResource) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AzureResource: %s", err))
	}

	*out = *target.(*AzureResource)
}

// Validate valides the current information stored into the structure.
func (o *AzureResource) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("data", o.Data); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("kind", string(o.Kind)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("kind", string(o.Kind), []string{"VirtualMachine", "NetworkInterface", "Subnet", "IPConfiguration", "VirtualNetwork", "NetworkSecurityGroup", "NATGateway", "PublicIPAddress", "PublicIPPrefix", "VirtualMachineScaleSet", "VirtualMachineScaleSetVM", "LoadBalancer", "BackendAddressPool"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("provider", string(o.Provider)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("provider", string(o.Provider), []string{"MicrosoftCompute", "MicrosoftNetwork"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("resourceGroup", o.ResourceGroup); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("resourceID", o.ResourceID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("subscriptionID", o.SubscriptionID); err != nil {
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
func (*AzureResource) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AzureResourceAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AzureResourceLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AzureResource) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AzureResourceAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *AzureResource) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "data":
		return o.Data
	case "kind":
		return o.Kind
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "provider":
		return o.Provider
	case "resourceGroup":
		return o.ResourceGroup
	case "resourceID":
		return o.ResourceID
	case "subscriptionID":
		return o.SubscriptionID
	case "tags":
		return o.Tags
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// AzureResourceAttributesMap represents the map of attribute for AzureResource.
var AzureResourceAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Data": {
		AllowedChoices: []string{},
		BSONFieldName:  "data",
		ConvertedName:  "Data",
		Description:    `The json-encoded data that represents the resource.`,
		Exposed:        true,
		Name:           "data",
		Required:       true,
		Stored:         true,
		SubType:        "[]byte",
		Type:           "external",
	},
	"Kind": {
		AllowedChoices: []string{"VirtualMachine", "NetworkInterface", "Subnet", "IPConfiguration", "VirtualNetwork", "NetworkSecurityGroup", "NATGateway", "PublicIPAddress", "PublicIPPrefix", "VirtualMachineScaleSet", "VirtualMachineScaleSetVM", "LoadBalancer", "BackendAddressPool"},
		BSONFieldName:  "kind",
		ConvertedName:  "Kind",
		Description:    `The specific kind of the resource.`,
		Exposed:        true,
		Name:           "kind",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"MigrationsLog": {
		AllowedChoices: []string{},
		BSONFieldName:  "migrationslog",
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `The name of this resource.`,
		Exposed:        true,
		Name:           "name",
		Required:       true,
		Stored:         true,
		Type:           "string",
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
	"Provider": {
		AllowedChoices: []string{"MicrosoftCompute", "MicrosoftNetwork"},
		BSONFieldName:  "provider",
		ConvertedName:  "Provider",
		Description:    `The major type of the resource.`,
		Exposed:        true,
		Name:           "provider",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"ResourceGroup": {
		AllowedChoices: []string{},
		BSONFieldName:  "resourcegroup",
		ConvertedName:  "ResourceGroup",
		Description:    `The name of the logical subcontainer of cloud resources.`,
		Exposed:        true,
		Name:           "resourceGroup",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"ResourceID": {
		AllowedChoices: []string{},
		BSONFieldName:  "resourceid",
		ConvertedName:  "ResourceID",
		Description:    `The identifier of the resource as presented by Azure, which is a path.`,
		Exposed:        true,
		Name:           "resourceID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"SubscriptionID": {
		AllowedChoices: []string{},
		BSONFieldName:  "subscriptionid",
		ConvertedName:  "SubscriptionID",
		Description:    `The logical ID of the container which contains the cloud resources.`,
		Exposed:        true,
		Name:           "subscriptionID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Tags": {
		AllowedChoices: []string{},
		BSONFieldName:  "tags",
		ConvertedName:  "Tags",
		Description: `Contextual values that can be used to narrow searching of resources if the
resourceID is not known. For instance, it could be used to store a resource's
location or public IP addresses to support cross-cloud analysis.`,
		Exposed: true,
		Name:    "tags",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"ZHash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zhash",
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zone",
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// AzureResourceLowerCaseAttributesMap represents the map of attribute for AzureResource.
var AzureResourceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"data": {
		AllowedChoices: []string{},
		BSONFieldName:  "data",
		ConvertedName:  "Data",
		Description:    `The json-encoded data that represents the resource.`,
		Exposed:        true,
		Name:           "data",
		Required:       true,
		Stored:         true,
		SubType:        "[]byte",
		Type:           "external",
	},
	"kind": {
		AllowedChoices: []string{"VirtualMachine", "NetworkInterface", "Subnet", "IPConfiguration", "VirtualNetwork", "NetworkSecurityGroup", "NATGateway", "PublicIPAddress", "PublicIPPrefix", "VirtualMachineScaleSet", "VirtualMachineScaleSetVM", "LoadBalancer", "BackendAddressPool"},
		BSONFieldName:  "kind",
		ConvertedName:  "Kind",
		Description:    `The specific kind of the resource.`,
		Exposed:        true,
		Name:           "kind",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"migrationslog": {
		AllowedChoices: []string{},
		BSONFieldName:  "migrationslog",
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `The name of this resource.`,
		Exposed:        true,
		Name:           "name",
		Required:       true,
		Stored:         true,
		Type:           "string",
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
	"provider": {
		AllowedChoices: []string{"MicrosoftCompute", "MicrosoftNetwork"},
		BSONFieldName:  "provider",
		ConvertedName:  "Provider",
		Description:    `The major type of the resource.`,
		Exposed:        true,
		Name:           "provider",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"resourcegroup": {
		AllowedChoices: []string{},
		BSONFieldName:  "resourcegroup",
		ConvertedName:  "ResourceGroup",
		Description:    `The name of the logical subcontainer of cloud resources.`,
		Exposed:        true,
		Name:           "resourceGroup",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"resourceid": {
		AllowedChoices: []string{},
		BSONFieldName:  "resourceid",
		ConvertedName:  "ResourceID",
		Description:    `The identifier of the resource as presented by Azure, which is a path.`,
		Exposed:        true,
		Name:           "resourceID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"subscriptionid": {
		AllowedChoices: []string{},
		BSONFieldName:  "subscriptionid",
		ConvertedName:  "SubscriptionID",
		Description:    `The logical ID of the container which contains the cloud resources.`,
		Exposed:        true,
		Name:           "subscriptionID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"tags": {
		AllowedChoices: []string{},
		BSONFieldName:  "tags",
		ConvertedName:  "Tags",
		Description: `Contextual values that can be used to narrow searching of resources if the
resourceID is not known. For instance, it could be used to store a resource's
location or public IP addresses to support cross-cloud analysis.`,
		Exposed: true,
		Name:    "tags",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"zhash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zhash",
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zone",
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SparseAzureResourcesList represents a list of SparseAzureResources
type SparseAzureResourcesList []*SparseAzureResource

// Identity returns the identity of the objects in the list.
func (o SparseAzureResourcesList) Identity() elemental.Identity {

	return AzureResourceIdentity
}

// Copy returns a pointer to a copy the SparseAzureResourcesList.
func (o SparseAzureResourcesList) Copy() elemental.Identifiables {

	copy := append(SparseAzureResourcesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAzureResourcesList.
func (o SparseAzureResourcesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAzureResourcesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAzureResource))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAzureResourcesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAzureResourcesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseAzureResourcesList converted to AzureResourcesList.
func (o SparseAzureResourcesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAzureResourcesList) Version() int {

	return 1
}

// SparseAzureResource represents the sparse version of a azureresource.
type SparseAzureResource struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// The json-encoded data that represents the resource.
	Data *[]byte `json:"data,omitempty" msgpack:"data,omitempty" bson:"data,omitempty" mapstructure:"data,omitempty"`

	// The specific kind of the resource.
	Kind *AzureResourceKindValue `json:"kind,omitempty" msgpack:"kind,omitempty" bson:"kind,omitempty" mapstructure:"kind,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// The name of this resource.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// The major type of the resource.
	Provider *AzureResourceProviderValue `json:"provider,omitempty" msgpack:"provider,omitempty" bson:"provider,omitempty" mapstructure:"provider,omitempty"`

	// The name of the logical subcontainer of cloud resources.
	ResourceGroup *string `json:"resourceGroup,omitempty" msgpack:"resourceGroup,omitempty" bson:"resourcegroup,omitempty" mapstructure:"resourceGroup,omitempty"`

	// The identifier of the resource as presented by Azure, which is a path.
	ResourceID *string `json:"resourceID,omitempty" msgpack:"resourceID,omitempty" bson:"resourceid,omitempty" mapstructure:"resourceID,omitempty"`

	// The logical ID of the container which contains the cloud resources.
	SubscriptionID *string `json:"subscriptionID,omitempty" msgpack:"subscriptionID,omitempty" bson:"subscriptionid,omitempty" mapstructure:"subscriptionID,omitempty"`

	// Contextual values that can be used to narrow searching of resources if the
	// resourceID is not known. For instance, it could be used to store a resource's
	// location or public IP addresses to support cross-cloud analysis.
	Tags *[]string `json:"tags,omitempty" msgpack:"tags,omitempty" bson:"tags,omitempty" mapstructure:"tags,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAzureResource returns a new  SparseAzureResource.
func NewSparseAzureResource() *SparseAzureResource {
	return &SparseAzureResource{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAzureResource) Identity() elemental.Identity {

	return AzureResourceIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAzureResource) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAzureResource) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAzureResource) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseAzureResource{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Data != nil {
		s.Data = o.Data
	}
	if o.Kind != nil {
		s.Kind = o.Kind
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.Provider != nil {
		s.Provider = o.Provider
	}
	if o.ResourceGroup != nil {
		s.ResourceGroup = o.ResourceGroup
	}
	if o.ResourceID != nil {
		s.ResourceID = o.ResourceID
	}
	if o.SubscriptionID != nil {
		s.SubscriptionID = o.SubscriptionID
	}
	if o.Tags != nil {
		s.Tags = o.Tags
	}
	if o.ZHash != nil {
		s.ZHash = o.ZHash
	}
	if o.Zone != nil {
		s.Zone = o.Zone
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAzureResource) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseAzureResource{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Data != nil {
		o.Data = s.Data
	}
	if s.Kind != nil {
		o.Kind = s.Kind
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.Provider != nil {
		o.Provider = s.Provider
	}
	if s.ResourceGroup != nil {
		o.ResourceGroup = s.ResourceGroup
	}
	if s.ResourceID != nil {
		o.ResourceID = s.ResourceID
	}
	if s.SubscriptionID != nil {
		o.SubscriptionID = s.SubscriptionID
	}
	if s.Tags != nil {
		o.Tags = s.Tags
	}
	if s.ZHash != nil {
		o.ZHash = s.ZHash
	}
	if s.Zone != nil {
		o.Zone = s.Zone
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseAzureResource) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAzureResource) ToPlain() elemental.PlainIdentifiable {

	out := NewAzureResource()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Data != nil {
		out.Data = *o.Data
	}
	if o.Kind != nil {
		out.Kind = *o.Kind
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.Provider != nil {
		out.Provider = *o.Provider
	}
	if o.ResourceGroup != nil {
		out.ResourceGroup = *o.ResourceGroup
	}
	if o.ResourceID != nil {
		out.ResourceID = *o.ResourceID
	}
	if o.SubscriptionID != nil {
		out.SubscriptionID = *o.SubscriptionID
	}
	if o.Tags != nil {
		out.Tags = *o.Tags
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseAzureResource) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseAzureResource) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseAzureResource) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseAzureResource) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseAzureResource) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseAzureResource) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseAzureResource) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseAzureResource) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseAzureResource.
func (o *SparseAzureResource) DeepCopy() *SparseAzureResource {

	if o == nil {
		return nil
	}

	out := &SparseAzureResource{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAzureResource.
func (o *SparseAzureResource) DeepCopyInto(out *SparseAzureResource) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAzureResource: %s", err))
	}

	*out = *target.(*SparseAzureResource)
}

type mongoAttributesAzureResource struct {
	ID             bson.ObjectId              `bson:"_id,omitempty"`
	Data           []byte                     `bson:"data"`
	Kind           AzureResourceKindValue     `bson:"kind"`
	MigrationsLog  map[string]string          `bson:"migrationslog,omitempty"`
	Name           string                     `bson:"name"`
	Namespace      string                     `bson:"namespace"`
	Provider       AzureResourceProviderValue `bson:"provider"`
	ResourceGroup  string                     `bson:"resourcegroup"`
	ResourceID     string                     `bson:"resourceid"`
	SubscriptionID string                     `bson:"subscriptionid"`
	Tags           []string                   `bson:"tags"`
	ZHash          int                        `bson:"zhash"`
	Zone           int                        `bson:"zone"`
}
type mongoAttributesSparseAzureResource struct {
	ID             bson.ObjectId               `bson:"_id,omitempty"`
	Data           *[]byte                     `bson:"data,omitempty"`
	Kind           *AzureResourceKindValue     `bson:"kind,omitempty"`
	MigrationsLog  *map[string]string          `bson:"migrationslog,omitempty"`
	Name           *string                     `bson:"name,omitempty"`
	Namespace      *string                     `bson:"namespace,omitempty"`
	Provider       *AzureResourceProviderValue `bson:"provider,omitempty"`
	ResourceGroup  *string                     `bson:"resourcegroup,omitempty"`
	ResourceID     *string                     `bson:"resourceid,omitempty"`
	SubscriptionID *string                     `bson:"subscriptionid,omitempty"`
	Tags           *[]string                   `bson:"tags,omitempty"`
	ZHash          *int                        `bson:"zhash,omitempty"`
	Zone           *int                        `bson:"zone,omitempty"`
}
