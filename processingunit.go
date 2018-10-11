package gaia

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
	"go.aporeto.io/gaia/types"
)

// ProcessingUnitEnforcementStatusValue represents the possible values for attribute "enforcementStatus".
type ProcessingUnitEnforcementStatusValue string

const (
	// ProcessingUnitEnforcementStatusFailed represents the value Failed.
	ProcessingUnitEnforcementStatusFailed ProcessingUnitEnforcementStatusValue = "Failed"

	// ProcessingUnitEnforcementStatusInactive represents the value Inactive.
	ProcessingUnitEnforcementStatusInactive ProcessingUnitEnforcementStatusValue = "Inactive"

	// ProcessingUnitEnforcementStatusProtected represents the value Protected.
	ProcessingUnitEnforcementStatusProtected ProcessingUnitEnforcementStatusValue = "Protected"
)

// ProcessingUnitOperationalStatusValue represents the possible values for attribute "operationalStatus".
type ProcessingUnitOperationalStatusValue string

const (
	// ProcessingUnitOperationalStatusInitialized represents the value Initialized.
	ProcessingUnitOperationalStatusInitialized ProcessingUnitOperationalStatusValue = "Initialized"

	// ProcessingUnitOperationalStatusPaused represents the value Paused.
	ProcessingUnitOperationalStatusPaused ProcessingUnitOperationalStatusValue = "Paused"

	// ProcessingUnitOperationalStatusRunning represents the value Running.
	ProcessingUnitOperationalStatusRunning ProcessingUnitOperationalStatusValue = "Running"

	// ProcessingUnitOperationalStatusStopped represents the value Stopped.
	ProcessingUnitOperationalStatusStopped ProcessingUnitOperationalStatusValue = "Stopped"

	// ProcessingUnitOperationalStatusTerminated represents the value Terminated.
	ProcessingUnitOperationalStatusTerminated ProcessingUnitOperationalStatusValue = "Terminated"
)

// ProcessingUnitTypeValue represents the possible values for attribute "type".
type ProcessingUnitTypeValue string

const (
	// ProcessingUnitTypeAPIGateway represents the value APIGateway.
	ProcessingUnitTypeAPIGateway ProcessingUnitTypeValue = "APIGateway"

	// ProcessingUnitTypeDocker represents the value Docker.
	ProcessingUnitTypeDocker ProcessingUnitTypeValue = "Docker"

	// ProcessingUnitTypeLinuxService represents the value LinuxService.
	ProcessingUnitTypeLinuxService ProcessingUnitTypeValue = "LinuxService"

	// ProcessingUnitTypeRKT represents the value RKT.
	ProcessingUnitTypeRKT ProcessingUnitTypeValue = "RKT"

	// ProcessingUnitTypeUser represents the value User.
	ProcessingUnitTypeUser ProcessingUnitTypeValue = "User"
)

// ProcessingUnitIdentity represents the Identity of the object.
var ProcessingUnitIdentity = elemental.Identity{
	Name:     "processingunit",
	Category: "processingunits",
	Package:  "squall",
	Private:  false,
}

// ProcessingUnitsList represents a list of ProcessingUnits
type ProcessingUnitsList []*ProcessingUnit

// Identity returns the identity of the objects in the list.
func (o ProcessingUnitsList) Identity() elemental.Identity {

	return ProcessingUnitIdentity
}

// Copy returns a pointer to a copy the ProcessingUnitsList.
func (o ProcessingUnitsList) Copy() elemental.Identifiables {

	copy := append(ProcessingUnitsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ProcessingUnitsList.
func (o ProcessingUnitsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ProcessingUnitsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ProcessingUnit))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ProcessingUnitsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ProcessingUnitsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o ProcessingUnitsList) Version() int {

	return 1
}

// ProcessingUnit represents the model of a processingunit
type ProcessingUnit struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// Archived defines if the object is archived.
	Archived bool `json:"-" bson:"archived" mapstructure:"-,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// EnforcementStatus communicates the state of the enforcer for that PU.
	EnforcementStatus ProcessingUnitEnforcementStatusValue `json:"enforcementStatus" bson:"enforcementstatus" mapstructure:"enforcementStatus,omitempty"`

	// EnforcerID is the ID of the enforcer associated with the processing unit.
	EnforcerID string `json:"enforcerID" bson:"enforcerid" mapstructure:"enforcerID,omitempty"`

	// Docker image, or path to executable.
	Image string `json:"image" bson:"image" mapstructure:"image,omitempty"`

	// Last poke is the time when the pu got last poked.
	LastPokeTime time.Time `json:"-" bson:"lastpoketime" mapstructure:"-,omitempty"`

	// LastSyncTime is the time when the policy was last resolved.
	LastSyncTime time.Time `json:"lastSyncTime" bson:"lastsynctime" mapstructure:"lastSyncTime,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NativeContextID is the Docker UUID or service PID.
	NativeContextID string `json:"nativeContextID" bson:"nativecontextid" mapstructure:"nativeContextID,omitempty"`

	// NetworkServices is the list of services that this processing unit has declared
	// that it will be listening to. This can happen either with an activation command
	// or by exposing the ports in a container manifest.
	NetworkServices types.ProcessingUnitServicesList `json:"networkServices" bson:"networkservices" mapstructure:"networkServices,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// OperationalStatus of the processing unit.
	OperationalStatus ProcessingUnitOperationalStatusValue `json:"operationalStatus" bson:"operationalstatus" mapstructure:"operationalStatus,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Type of the container ecosystem.
	Type ProcessingUnitTypeValue `json:"type" bson:"type" mapstructure:"type,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewProcessingUnit returns a new *ProcessingUnit
func NewProcessingUnit() *ProcessingUnit {

	return &ProcessingUnit{
		ModelVersion:      1,
		Annotations:       map[string][]string{},
		EnforcementStatus: ProcessingUnitEnforcementStatusInactive,
		AssociatedTags:    []string{},
		Metadata:          []string{},
		NetworkServices:   types.ProcessingUnitServicesList{},
		NormalizedTags:    []string{},
		OperationalStatus: ProcessingUnitOperationalStatusInitialized,
	}
}

// Identity returns the Identity of the object.
func (o *ProcessingUnit) Identity() elemental.Identity {

	return ProcessingUnitIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ProcessingUnit) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ProcessingUnit) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *ProcessingUnit) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *ProcessingUnit) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *ProcessingUnit) Doc() string {
	return `A Processing Unit reprents anything that can compute. It can be a Docker
container, or a simple Unix process. They are created, updated and deleted by
the system as they come and go. You can only modify its tags.  Processing Units
use Network Access Policies to define which other Processing Units or External
Services they can communicate with and File Access Policies to define what File
Paths they can use.`
}

func (o *ProcessingUnit) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *ProcessingUnit) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *ProcessingUnit) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetArchived returns the Archived of the receiver.
func (o *ProcessingUnit) GetArchived() bool {

	return o.Archived
}

// SetArchived sets the given Archived of the receiver.
func (o *ProcessingUnit) SetArchived(archived bool) {

	o.Archived = archived
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *ProcessingUnit) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *ProcessingUnit) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *ProcessingUnit) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *ProcessingUnit) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMetadata returns the Metadata of the receiver.
func (o *ProcessingUnit) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the given Metadata of the receiver.
func (o *ProcessingUnit) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *ProcessingUnit) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *ProcessingUnit) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *ProcessingUnit) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *ProcessingUnit) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *ProcessingUnit) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *ProcessingUnit) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *ProcessingUnit) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *ProcessingUnit) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *ProcessingUnit) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// Validate valides the current information stored into the structure.
func (o *ProcessingUnit) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("enforcementStatus", string(o.EnforcementStatus), []string{"Protected", "Failed", "Inactive"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("operationalStatus", string(o.OperationalStatus), []string{"Initialized", "Paused", "Running", "Stopped", "Terminated"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Docker", "LinuxService", "RKT", "User", "APIGateway"}, false); err != nil {
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
func (*ProcessingUnit) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ProcessingUnitAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ProcessingUnitLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ProcessingUnit) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ProcessingUnitAttributesMap
}

// ProcessingUnitAttributesMap represents the map of attribute for ProcessingUnit.
var ProcessingUnitAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "annotations",
		Type:           "external",
	},
	"Archived": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Archived",
		Description:    `Archived defines if the object is archived.`,
		Getter:         true,
		Name:           "archived",
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `CreatedTime is the time at which the object was created.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"EnforcementStatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"Protected", "Failed", "Inactive"},
		ConvertedName:  "EnforcementStatus",
		DefaultValue:   ProcessingUnitEnforcementStatusInactive,
		Description:    `EnforcementStatus communicates the state of the enforcer for that PU.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "enforcementStatus",
		Stored:         true,
		Type:           "enum",
	},
	"EnforcerID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `EnforcerID is the ID of the enforcer associated with the processing unit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "enforcerID",
		Stored:         true,
		Type:           "string",
	},
	"Image": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Image",
		Description:    `Docker image, or path to executable.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "image",
		Stored:         true,
		Type:           "string",
	},
	"LastPokeTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastPokeTime",
		Description:    `Last poke is the time when the pu got last poked.`,
		Name:           "lastPokeTime",
		Stored:         true,
		Type:           "time",
	},
	"LastSyncTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LastSyncTime",
		Description:    `LastSyncTime is the time when the policy was last resolved.`,
		Exposed:        true,
		Name:           "lastSyncTime",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "metadata_list",
		Type:       "external",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NativeContextID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NativeContextID",
		Description:    `NativeContextID is the Docker UUID or service PID.`,
		Exposed:        true,
		Name:           "nativeContextID",
		Stored:         true,
		Type:           "string",
	},
	"NetworkServices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NetworkServices",
		Description: `NetworkServices is the list of services that this processing unit has declared
that it will be listening to. This can happen either with an activation command
or by exposing the ports in a container manifest.`,
		Exposed:   true,
		Name:      "networkServices",
		Orderable: true,
		Stored:    true,
		SubType:   "processing_unit_services_list",
		Type:      "external",
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Transient:      true,
		Type:           "external",
	},
	"OperationalStatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"Initialized", "Paused", "Running", "Stopped", "Terminated"},
		ConvertedName:  "OperationalStatus",
		DefaultValue:   ProcessingUnitOperationalStatusInitialized,
		Description:    `OperationalStatus of the processing unit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operationalStatus",
		Stored:         true,
		Type:           "enum",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{"Docker", "LinuxService", "RKT", "User", "APIGateway"},
		ConvertedName:  "Type",
		CreationOnly:   true,
		Description:    `Type of the container ecosystem.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "type",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `UpdateTime is the time at which an entity was updated.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}

// ProcessingUnitLowerCaseAttributesMap represents the map of attribute for ProcessingUnit.
var ProcessingUnitLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "annotations",
		Type:           "external",
	},
	"archived": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Archived",
		Description:    `Archived defines if the object is archived.`,
		Getter:         true,
		Name:           "archived",
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"associatedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `CreatedTime is the time at which the object was created.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"enforcementstatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"Protected", "Failed", "Inactive"},
		ConvertedName:  "EnforcementStatus",
		DefaultValue:   ProcessingUnitEnforcementStatusInactive,
		Description:    `EnforcementStatus communicates the state of the enforcer for that PU.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "enforcementStatus",
		Stored:         true,
		Type:           "enum",
	},
	"enforcerid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `EnforcerID is the ID of the enforcer associated with the processing unit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "enforcerID",
		Stored:         true,
		Type:           "string",
	},
	"image": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Image",
		Description:    `Docker image, or path to executable.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "image",
		Stored:         true,
		Type:           "string",
	},
	"lastpoketime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastPokeTime",
		Description:    `Last poke is the time when the pu got last poked.`,
		Name:           "lastPokeTime",
		Stored:         true,
		Type:           "time",
	},
	"lastsynctime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LastSyncTime",
		Description:    `LastSyncTime is the time when the policy was last resolved.`,
		Exposed:        true,
		Name:           "lastSyncTime",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "metadata_list",
		Type:       "external",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"nativecontextid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NativeContextID",
		Description:    `NativeContextID is the Docker UUID or service PID.`,
		Exposed:        true,
		Name:           "nativeContextID",
		Stored:         true,
		Type:           "string",
	},
	"networkservices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NetworkServices",
		Description: `NetworkServices is the list of services that this processing unit has declared
that it will be listening to. This can happen either with an activation command
or by exposing the ports in a container manifest.`,
		Exposed:   true,
		Name:      "networkServices",
		Orderable: true,
		Stored:    true,
		SubType:   "processing_unit_services_list",
		Type:      "external",
	},
	"normalizedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Transient:      true,
		Type:           "external",
	},
	"operationalstatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"Initialized", "Paused", "Running", "Stopped", "Terminated"},
		ConvertedName:  "OperationalStatus",
		DefaultValue:   ProcessingUnitOperationalStatusInitialized,
		Description:    `OperationalStatus of the processing unit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operationalStatus",
		Stored:         true,
		Type:           "enum",
	},
	"protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"type": elemental.AttributeSpecification{
		AllowedChoices: []string{"Docker", "LinuxService", "RKT", "User", "APIGateway"},
		ConvertedName:  "Type",
		CreationOnly:   true,
		Description:    `Type of the container ecosystem.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "type",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `UpdateTime is the time at which an entity was updated.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}
