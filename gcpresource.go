// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// GCPResourceKindValue represents the possible values for attribute "kind".
type GCPResourceKindValue string

const (
	// GCPResourceKindComputeInstance represents the value ComputeInstance.
	GCPResourceKindComputeInstance GCPResourceKindValue = "ComputeInstance"

	// GCPResourceKindComputeNetwork represents the value ComputeNetwork.
	GCPResourceKindComputeNetwork GCPResourceKindValue = "ComputeNetwork"

	// GCPResourceKindComputeSubnetwork represents the value ComputeSubnetwork.
	GCPResourceKindComputeSubnetwork GCPResourceKindValue = "ComputeSubnetwork"

	// GCPResourceKindPending represents the value Pending.
	GCPResourceKindPending GCPResourceKindValue = "Pending"
)

// GCPResourceIdentity represents the Identity of the object.
var GCPResourceIdentity = elemental.Identity{
	Name:     "gcpresource",
	Category: "gcpresources",
	Package:  "pandemona",
	Private:  false,
}

// GCPResourcesList represents a list of GCPResources
type GCPResourcesList []*GCPResource

// Identity returns the identity of the objects in the list.
func (o GCPResourcesList) Identity() elemental.Identity {

	return GCPResourceIdentity
}

// Copy returns a pointer to a copy the GCPResourcesList.
func (o GCPResourcesList) Copy() elemental.Identifiables {

	copy := append(GCPResourcesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the GCPResourcesList.
func (o GCPResourcesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(GCPResourcesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*GCPResource))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o GCPResourcesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o GCPResourcesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the GCPResourcesList converted to SparseGCPResourcesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o GCPResourcesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseGCPResourcesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseGCPResource)
	}

	return out
}

// Version returns the version of the content.
func (o GCPResourcesList) Version() int {

	return 1
}

// GCPResource represents the model of a gcpresource
type GCPResource struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// The JSON-encoded data that represents the resource.
	Data []byte `json:"data" msgpack:"data" bson:"data" mapstructure:"data,omitempty"`

	// Contextual values that can be used to narrow searching of resources if the
	// numericID or selflink are not known. For instance, it could be used to store
	// a resource's location or public IP addresses to support cross-cloud analysis.
	DenormedFields []string `json:"denormedFields" msgpack:"denormedFields" bson:"denormedfields" mapstructure:"denormedFields,omitempty"`

	// The specific kind of the resource.
	Kind GCPResourceKindValue `json:"kind" msgpack:"kind" bson:"kind" mapstructure:"kind,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// The name of the resource.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// A numeric resource ID that will mainly be used in RQL queries.
	NumericID string `json:"numericID" msgpack:"numericID" bson:"numericid" mapstructure:"numericID,omitempty"`

	// The identifier used by Prisma Cloud to locate the same resource.
	PrismaCloudRRN string `json:"prismaCloudRRN,omitempty" msgpack:"prismaCloudRRN,omitempty" bson:"prismacloudrrn,omitempty" mapstructure:"prismaCloudRRN,omitempty"`

	// The identifier of the resource as presented by GCP, which is a URL.
	Selflink string `json:"selflink" msgpack:"selflink" bson:"selflink" mapstructure:"selflink,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewGCPResource returns a new *GCPResource
func NewGCPResource() *GCPResource {

	return &GCPResource{
		ModelVersion:   1,
		Data:           []byte{},
		DenormedFields: []string{},
		Kind:           GCPResourceKindPending,
		MigrationsLog:  map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *GCPResource) Identity() elemental.Identity {

	return GCPResourceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *GCPResource) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *GCPResource) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *GCPResource) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesGCPResource{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Data = o.Data
	s.DenormedFields = o.DenormedFields
	s.Kind = o.Kind
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NumericID = o.NumericID
	s.PrismaCloudRRN = o.PrismaCloudRRN
	s.Selflink = o.Selflink
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *GCPResource) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesGCPResource{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Data = s.Data
	o.DenormedFields = s.DenormedFields
	o.Kind = s.Kind
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NumericID = s.NumericID
	o.PrismaCloudRRN = s.PrismaCloudRRN
	o.Selflink = s.Selflink
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *GCPResource) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *GCPResource) BleveType() string {

	return "gcpresource"
}

// DefaultOrder returns the list of default ordering fields.
func (o *GCPResource) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *GCPResource) Doc() string {

	return `Represents a GCP cloud resource such as a virtual machine.`
}

func (o *GCPResource) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *GCPResource) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *GCPResource) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *GCPResource) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *GCPResource) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetZHash returns the ZHash of the receiver.
func (o *GCPResource) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *GCPResource) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *GCPResource) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *GCPResource) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *GCPResource) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseGCPResource{
			ID:             &o.ID,
			Data:           &o.Data,
			DenormedFields: &o.DenormedFields,
			Kind:           &o.Kind,
			MigrationsLog:  &o.MigrationsLog,
			Name:           &o.Name,
			Namespace:      &o.Namespace,
			NumericID:      &o.NumericID,
			PrismaCloudRRN: &o.PrismaCloudRRN,
			Selflink:       &o.Selflink,
			ZHash:          &o.ZHash,
			Zone:           &o.Zone,
		}
	}

	sp := &SparseGCPResource{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "data":
			sp.Data = &(o.Data)
		case "denormedFields":
			sp.DenormedFields = &(o.DenormedFields)
		case "kind":
			sp.Kind = &(o.Kind)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "numericID":
			sp.NumericID = &(o.NumericID)
		case "prismaCloudRRN":
			sp.PrismaCloudRRN = &(o.PrismaCloudRRN)
		case "selflink":
			sp.Selflink = &(o.Selflink)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseGCPResource to the object.
func (o *GCPResource) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseGCPResource)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Data != nil {
		o.Data = *so.Data
	}
	if so.DenormedFields != nil {
		o.DenormedFields = *so.DenormedFields
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
	if so.NumericID != nil {
		o.NumericID = *so.NumericID
	}
	if so.PrismaCloudRRN != nil {
		o.PrismaCloudRRN = *so.PrismaCloudRRN
	}
	if so.Selflink != nil {
		o.Selflink = *so.Selflink
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the GCPResource.
func (o *GCPResource) DeepCopy() *GCPResource {

	if o == nil {
		return nil
	}

	out := &GCPResource{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *GCPResource.
func (o *GCPResource) DeepCopyInto(out *GCPResource) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy GCPResource: %s", err))
	}

	*out = *target.(*GCPResource)
}

// Validate valides the current information stored into the structure.
func (o *GCPResource) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("data", o.Data); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("kind", string(o.Kind), []string{"ComputeInstance", "ComputeSubnetwork", "ComputeNetwork", "Pending"}, false); err != nil {
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
func (*GCPResource) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := GCPResourceAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return GCPResourceLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*GCPResource) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return GCPResourceAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *GCPResource) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "data":
		return o.Data
	case "denormedFields":
		return o.DenormedFields
	case "kind":
		return o.Kind
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "numericID":
		return o.NumericID
	case "prismaCloudRRN":
		return o.PrismaCloudRRN
	case "selflink":
		return o.Selflink
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// GCPResourceAttributesMap represents the map of attribute for GCPResource.
var GCPResourceAttributesMap = map[string]elemental.AttributeSpecification{
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
		Description:    `The JSON-encoded data that represents the resource.`,
		Exposed:        true,
		Name:           "data",
		Required:       true,
		Stored:         true,
		SubType:        "[]byte",
		Type:           "external",
	},
	"DenormedFields": {
		AllowedChoices: []string{},
		BSONFieldName:  "denormedfields",
		ConvertedName:  "DenormedFields",
		Description: `Contextual values that can be used to narrow searching of resources if the
numericID or selflink are not known. For instance, it could be used to store
a resource's location or public IP addresses to support cross-cloud analysis.`,
		Exposed: true,
		Name:    "denormedFields",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"Kind": {
		AllowedChoices: []string{"ComputeInstance", "ComputeSubnetwork", "ComputeNetwork", "Pending"},
		BSONFieldName:  "kind",
		ConvertedName:  "Kind",
		DefaultValue:   GCPResourceKindPending,
		Description:    `The specific kind of the resource.`,
		Exposed:        true,
		Name:           "kind",
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
		Description:    `The name of the resource.`,
		Exposed:        true,
		Name:           "name",
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
	"NumericID": {
		AllowedChoices: []string{},
		BSONFieldName:  "numericid",
		ConvertedName:  "NumericID",
		Description:    `A numeric resource ID that will mainly be used in RQL queries.`,
		Exposed:        true,
		Name:           "numericID",
		Stored:         true,
		Type:           "string",
	},
	"PrismaCloudRRN": {
		AllowedChoices: []string{},
		BSONFieldName:  "prismacloudrrn",
		ConvertedName:  "PrismaCloudRRN",
		Description:    `The identifier used by Prisma Cloud to locate the same resource.`,
		Exposed:        true,
		Name:           "prismaCloudRRN",
		Stored:         true,
		Type:           "string",
	},
	"Selflink": {
		AllowedChoices: []string{},
		BSONFieldName:  "selflink",
		ConvertedName:  "Selflink",
		Description:    `The identifier of the resource as presented by GCP, which is a URL.`,
		Exposed:        true,
		Name:           "selflink",
		Stored:         true,
		Type:           "string",
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

// GCPResourceLowerCaseAttributesMap represents the map of attribute for GCPResource.
var GCPResourceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		Description:    `The JSON-encoded data that represents the resource.`,
		Exposed:        true,
		Name:           "data",
		Required:       true,
		Stored:         true,
		SubType:        "[]byte",
		Type:           "external",
	},
	"denormedfields": {
		AllowedChoices: []string{},
		BSONFieldName:  "denormedfields",
		ConvertedName:  "DenormedFields",
		Description: `Contextual values that can be used to narrow searching of resources if the
numericID or selflink are not known. For instance, it could be used to store
a resource's location or public IP addresses to support cross-cloud analysis.`,
		Exposed: true,
		Name:    "denormedFields",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"kind": {
		AllowedChoices: []string{"ComputeInstance", "ComputeSubnetwork", "ComputeNetwork", "Pending"},
		BSONFieldName:  "kind",
		ConvertedName:  "Kind",
		DefaultValue:   GCPResourceKindPending,
		Description:    `The specific kind of the resource.`,
		Exposed:        true,
		Name:           "kind",
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
		Description:    `The name of the resource.`,
		Exposed:        true,
		Name:           "name",
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
	"numericid": {
		AllowedChoices: []string{},
		BSONFieldName:  "numericid",
		ConvertedName:  "NumericID",
		Description:    `A numeric resource ID that will mainly be used in RQL queries.`,
		Exposed:        true,
		Name:           "numericID",
		Stored:         true,
		Type:           "string",
	},
	"prismacloudrrn": {
		AllowedChoices: []string{},
		BSONFieldName:  "prismacloudrrn",
		ConvertedName:  "PrismaCloudRRN",
		Description:    `The identifier used by Prisma Cloud to locate the same resource.`,
		Exposed:        true,
		Name:           "prismaCloudRRN",
		Stored:         true,
		Type:           "string",
	},
	"selflink": {
		AllowedChoices: []string{},
		BSONFieldName:  "selflink",
		ConvertedName:  "Selflink",
		Description:    `The identifier of the resource as presented by GCP, which is a URL.`,
		Exposed:        true,
		Name:           "selflink",
		Stored:         true,
		Type:           "string",
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

// SparseGCPResourcesList represents a list of SparseGCPResources
type SparseGCPResourcesList []*SparseGCPResource

// Identity returns the identity of the objects in the list.
func (o SparseGCPResourcesList) Identity() elemental.Identity {

	return GCPResourceIdentity
}

// Copy returns a pointer to a copy the SparseGCPResourcesList.
func (o SparseGCPResourcesList) Copy() elemental.Identifiables {

	copy := append(SparseGCPResourcesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseGCPResourcesList.
func (o SparseGCPResourcesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseGCPResourcesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseGCPResource))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseGCPResourcesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseGCPResourcesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseGCPResourcesList converted to GCPResourcesList.
func (o SparseGCPResourcesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseGCPResourcesList) Version() int {

	return 1
}

// SparseGCPResource represents the sparse version of a gcpresource.
type SparseGCPResource struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// The JSON-encoded data that represents the resource.
	Data *[]byte `json:"data,omitempty" msgpack:"data,omitempty" bson:"data,omitempty" mapstructure:"data,omitempty"`

	// Contextual values that can be used to narrow searching of resources if the
	// numericID or selflink are not known. For instance, it could be used to store
	// a resource's location or public IP addresses to support cross-cloud analysis.
	DenormedFields *[]string `json:"denormedFields,omitempty" msgpack:"denormedFields,omitempty" bson:"denormedfields,omitempty" mapstructure:"denormedFields,omitempty"`

	// The specific kind of the resource.
	Kind *GCPResourceKindValue `json:"kind,omitempty" msgpack:"kind,omitempty" bson:"kind,omitempty" mapstructure:"kind,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// A numeric resource ID that will mainly be used in RQL queries.
	NumericID *string `json:"numericID,omitempty" msgpack:"numericID,omitempty" bson:"numericid,omitempty" mapstructure:"numericID,omitempty"`

	// The identifier used by Prisma Cloud to locate the same resource.
	PrismaCloudRRN *string `json:"prismaCloudRRN,omitempty" msgpack:"prismaCloudRRN,omitempty" bson:"prismacloudrrn,omitempty" mapstructure:"prismaCloudRRN,omitempty"`

	// The identifier of the resource as presented by GCP, which is a URL.
	Selflink *string `json:"selflink,omitempty" msgpack:"selflink,omitempty" bson:"selflink,omitempty" mapstructure:"selflink,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseGCPResource returns a new  SparseGCPResource.
func NewSparseGCPResource() *SparseGCPResource {
	return &SparseGCPResource{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseGCPResource) Identity() elemental.Identity {

	return GCPResourceIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseGCPResource) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseGCPResource) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseGCPResource) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseGCPResource{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Data != nil {
		s.Data = o.Data
	}
	if o.DenormedFields != nil {
		s.DenormedFields = o.DenormedFields
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
	if o.NumericID != nil {
		s.NumericID = o.NumericID
	}
	if o.PrismaCloudRRN != nil {
		s.PrismaCloudRRN = o.PrismaCloudRRN
	}
	if o.Selflink != nil {
		s.Selflink = o.Selflink
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
func (o *SparseGCPResource) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseGCPResource{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Data != nil {
		o.Data = s.Data
	}
	if s.DenormedFields != nil {
		o.DenormedFields = s.DenormedFields
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
	if s.NumericID != nil {
		o.NumericID = s.NumericID
	}
	if s.PrismaCloudRRN != nil {
		o.PrismaCloudRRN = s.PrismaCloudRRN
	}
	if s.Selflink != nil {
		o.Selflink = s.Selflink
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
func (o *SparseGCPResource) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseGCPResource) ToPlain() elemental.PlainIdentifiable {

	out := NewGCPResource()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Data != nil {
		out.Data = *o.Data
	}
	if o.DenormedFields != nil {
		out.DenormedFields = *o.DenormedFields
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
	if o.NumericID != nil {
		out.NumericID = *o.NumericID
	}
	if o.PrismaCloudRRN != nil {
		out.PrismaCloudRRN = *o.PrismaCloudRRN
	}
	if o.Selflink != nil {
		out.Selflink = *o.Selflink
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
func (o *SparseGCPResource) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseGCPResource) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseGCPResource) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseGCPResource) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseGCPResource) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseGCPResource) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseGCPResource) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseGCPResource) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseGCPResource.
func (o *SparseGCPResource) DeepCopy() *SparseGCPResource {

	if o == nil {
		return nil
	}

	out := &SparseGCPResource{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseGCPResource.
func (o *SparseGCPResource) DeepCopyInto(out *SparseGCPResource) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseGCPResource: %s", err))
	}

	*out = *target.(*SparseGCPResource)
}

type mongoAttributesGCPResource struct {
	ID             bson.ObjectId        `bson:"_id,omitempty"`
	Data           []byte               `bson:"data"`
	DenormedFields []string             `bson:"denormedfields"`
	Kind           GCPResourceKindValue `bson:"kind"`
	MigrationsLog  map[string]string    `bson:"migrationslog,omitempty"`
	Name           string               `bson:"name"`
	Namespace      string               `bson:"namespace"`
	NumericID      string               `bson:"numericid"`
	PrismaCloudRRN string               `bson:"prismacloudrrn,omitempty"`
	Selflink       string               `bson:"selflink"`
	ZHash          int                  `bson:"zhash"`
	Zone           int                  `bson:"zone"`
}
type mongoAttributesSparseGCPResource struct {
	ID             bson.ObjectId         `bson:"_id,omitempty"`
	Data           *[]byte               `bson:"data,omitempty"`
	DenormedFields *[]string             `bson:"denormedfields,omitempty"`
	Kind           *GCPResourceKindValue `bson:"kind,omitempty"`
	MigrationsLog  *map[string]string    `bson:"migrationslog,omitempty"`
	Name           *string               `bson:"name,omitempty"`
	Namespace      *string               `bson:"namespace,omitempty"`
	NumericID      *string               `bson:"numericid,omitempty"`
	PrismaCloudRRN *string               `bson:"prismacloudrrn,omitempty"`
	Selflink       *string               `bson:"selflink,omitempty"`
	ZHash          *int                  `bson:"zhash,omitempty"`
	Zone           *int                  `bson:"zone,omitempty"`
}
