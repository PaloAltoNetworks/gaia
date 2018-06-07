package gaia

import (
	"fmt"
	"sync"

	"encoding/json"
	"github.com/aporeto-inc/elemental"
)

// RemoteProcessorModeValue represents the possible values for attribute "mode".
type RemoteProcessorModeValue string

const (
	// RemoteProcessorModePost represents the value Post.
	RemoteProcessorModePost RemoteProcessorModeValue = "Post"

	// RemoteProcessorModePre represents the value Pre.
	RemoteProcessorModePre RemoteProcessorModeValue = "Pre"
)

// RemoteProcessorIdentity represents the Identity of the object.
var RemoteProcessorIdentity = elemental.Identity{
	Name:     "remoteprocessor",
	Category: "remoteprocessors",
	Private:  false,
}

// RemoteProcessorsList represents a list of RemoteProcessors
type RemoteProcessorsList []*RemoteProcessor

// Identity returns the identity of the objects in the list.
func (o RemoteProcessorsList) Identity() elemental.Identity {

	return RemoteProcessorIdentity
}

// Copy returns a pointer to a copy the RemoteProcessorsList.
func (o RemoteProcessorsList) Copy() elemental.Identifiables {

	copy := append(RemoteProcessorsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the RemoteProcessorsList.
func (o RemoteProcessorsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(RemoteProcessorsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*RemoteProcessor))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o RemoteProcessorsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o RemoteProcessorsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o RemoteProcessorsList) Version() int {

	return 1
}

// RemoteProcessor represents the model of a remoteprocessor
type RemoteProcessor struct {
	// Represents the claims of the currently managed object.
	Claims []string `json:"claims" bson:"-" mapstructure:"claims,omitempty"`

	// Represents data received from the service.
	Input json.RawMessage `json:"input" bson:"-" mapstructure:"input,omitempty"`

	// Node defines the type of the hook.
	Mode RemoteProcessorModeValue `json:"mode" bson:"-" mapstructure:"mode,omitempty"`

	// Represents the current namespace.
	Namespace string `json:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	// Define the operation that is currently handled by the service.
	Operation elemental.Operation `json:"operation" bson:"-" mapstructure:"operation,omitempty"`

	// Returns the OutputData filled with the processor information.
	Output elemental.Identifiable `json:"output" bson:"-" mapstructure:"output,omitempty"`

	// RequestID gives the id of the request coming from the main server.
	RequestID string `json:"requestID" bson:"requestid" mapstructure:"requestID,omitempty"`

	// Represents the Identity name of the managed object.
	TargetIdentity string `json:"targetIdentity" bson:"-" mapstructure:"targetIdentity,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewRemoteProcessor returns a new *RemoteProcessor
func NewRemoteProcessor() *RemoteProcessor {

	return &RemoteProcessor{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *RemoteProcessor) Identity() elemental.Identity {

	return RemoteProcessorIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *RemoteProcessor) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *RemoteProcessor) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *RemoteProcessor) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *RemoteProcessor) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *RemoteProcessor) Doc() string {
	return `Hook to integrate an Aporeto service.`
}

func (o *RemoteProcessor) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *RemoteProcessor) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("input", o.Input); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateStringInList("mode", string(o.Mode), []string{"Post", "Pre"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("namespace", o.Namespace); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("operation", o.Operation); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("targetIdentity", o.TargetIdentity); err != nil {
		requiredErrors = append(requiredErrors, err)
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
func (*RemoteProcessor) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := RemoteProcessorAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return RemoteProcessorLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*RemoteProcessor) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RemoteProcessorAttributesMap
}

// RemoteProcessorAttributesMap represents the map of attribute for RemoteProcessor.
var RemoteProcessorAttributesMap = map[string]elemental.AttributeSpecification{
	"Claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `Represents the claims of the currently managed object.`,
		Exposed:        true,
		Name:           "claims",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
	"Input": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Input",
		Description:    `Represents data received from the service.`,
		Exposed:        true,
		Name:           "input",
		Required:       true,
		SubType:        "raw_json",
		Type:           "external",
	},
	"Mode": elemental.AttributeSpecification{
		AllowedChoices: []string{"Post", "Pre"},
		ConvertedName:  "Mode",
		Description:    `Node defines the type of the hook.`,
		Exposed:        true,
		Name:           "mode",
		Type:           "enum",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Represents the current namespace.`,
		Exposed:        true,
		Format:         "free",
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"Operation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Operation",
		Description:    `Define the operation that is currently handled by the service.`,
		Exposed:        true,
		Name:           "operation",
		Required:       true,
		SubType:        "elemental_operation",
		Type:           "external",
	},
	"Output": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Output",
		Description:    `Returns the OutputData filled with the processor information.`,
		Exposed:        true,
		Name:           "output",
		ReadOnly:       true,
		SubType:        "elemental_identitifable",
		Type:           "external",
	},
	"RequestID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RequestID",
		Description:    `RequestID gives the id of the request coming from the main server.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "requestID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"TargetIdentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetIdentity",
		Description:    `Represents the Identity name of the managed object.`,
		Exposed:        true,
		Format:         "free",
		Name:           "targetIdentity",
		Required:       true,
		Type:           "string",
	},
}

// RemoteProcessorLowerCaseAttributesMap represents the map of attribute for RemoteProcessor.
var RemoteProcessorLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `Represents the claims of the currently managed object.`,
		Exposed:        true,
		Name:           "claims",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
	"input": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Input",
		Description:    `Represents data received from the service.`,
		Exposed:        true,
		Name:           "input",
		Required:       true,
		SubType:        "raw_json",
		Type:           "external",
	},
	"mode": elemental.AttributeSpecification{
		AllowedChoices: []string{"Post", "Pre"},
		ConvertedName:  "Mode",
		Description:    `Node defines the type of the hook.`,
		Exposed:        true,
		Name:           "mode",
		Type:           "enum",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Represents the current namespace.`,
		Exposed:        true,
		Format:         "free",
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"operation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Operation",
		Description:    `Define the operation that is currently handled by the service.`,
		Exposed:        true,
		Name:           "operation",
		Required:       true,
		SubType:        "elemental_operation",
		Type:           "external",
	},
	"output": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Output",
		Description:    `Returns the OutputData filled with the processor information.`,
		Exposed:        true,
		Name:           "output",
		ReadOnly:       true,
		SubType:        "elemental_identitifable",
		Type:           "external",
	},
	"requestid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RequestID",
		Description:    `RequestID gives the id of the request coming from the main server.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "requestID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"targetidentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetIdentity",
		Description:    `Represents the Identity name of the managed object.`,
		Exposed:        true,
		Format:         "free",
		Name:           "targetIdentity",
		Required:       true,
		Type:           "string",
	},
}
