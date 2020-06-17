package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ProcessingUnitRefreshPingModeValue represents the possible values for attribute "pingMode".
type ProcessingUnitRefreshPingModeValue string

const (
	// ProcessingUnitRefreshPingModeAuto represents the value Auto.
	ProcessingUnitRefreshPingModeAuto ProcessingUnitRefreshPingModeValue = "Auto"

	// ProcessingUnitRefreshPingModeL3 represents the value L3.
	ProcessingUnitRefreshPingModeL3 ProcessingUnitRefreshPingModeValue = "L3"

	// ProcessingUnitRefreshPingModeL4 represents the value L4.
	ProcessingUnitRefreshPingModeL4 ProcessingUnitRefreshPingModeValue = "L4"

	// ProcessingUnitRefreshPingModeL7 represents the value L7.
	ProcessingUnitRefreshPingModeL7 ProcessingUnitRefreshPingModeValue = "L7"
)

// ProcessingUnitRefreshIdentity represents the Identity of the object.
var ProcessingUnitRefreshIdentity = elemental.Identity{
	Name:     "processingunitrefresh",
	Category: "processingunitrefreshes",
	Package:  "gaga",
	Private:  false,
}

// ProcessingUnitRefreshsList represents a list of ProcessingUnitRefreshs
type ProcessingUnitRefreshsList []*ProcessingUnitRefresh

// Identity returns the identity of the objects in the list.
func (o ProcessingUnitRefreshsList) Identity() elemental.Identity {

	return ProcessingUnitRefreshIdentity
}

// Copy returns a pointer to a copy the ProcessingUnitRefreshsList.
func (o ProcessingUnitRefreshsList) Copy() elemental.Identifiables {

	copy := append(ProcessingUnitRefreshsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ProcessingUnitRefreshsList.
func (o ProcessingUnitRefreshsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ProcessingUnitRefreshsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ProcessingUnitRefresh))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ProcessingUnitRefreshsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ProcessingUnitRefreshsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ProcessingUnitRefreshsList converted to SparseProcessingUnitRefreshsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ProcessingUnitRefreshsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseProcessingUnitRefreshsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseProcessingUnitRefresh)
	}

	return out
}

// Version returns the version of the content.
func (o ProcessingUnitRefreshsList) Version() int {

	return 1
}

// ProcessingUnitRefresh represents the model of a processingunitrefresh
type ProcessingUnitRefresh struct {
	// Contains the ID of the target processing unit.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// If set to true, start reporting debug information for the target processing
	// unit.
	Debug bool `json:"debug,omitempty" msgpack:"debug,omitempty" bson:"-" mapstructure:"debug,omitempty"`

	// Contains the original namespace of the processing unit.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	// Destination address to run ping.
	PingAddress string `json:"pingAddress,omitempty" msgpack:"pingAddress,omitempty" bson:"-" mapstructure:"pingAddress,omitempty"`

	// If set to true, start ping to the destination.
	PingEnabled bool `json:"pingEnabled,omitempty" msgpack:"pingEnabled,omitempty" bson:"-" mapstructure:"pingEnabled,omitempty"`

	// Number of iterations to run a ping probe.
	PingIterations int `json:"pingIterations,omitempty" msgpack:"pingIterations,omitempty" bson:"-" mapstructure:"pingIterations,omitempty"`

	// Represents the mode of ping to be used.
	PingMode ProcessingUnitRefreshPingModeValue `json:"pingMode" msgpack:"pingMode" bson:"-" mapstructure:"pingMode,omitempty"`

	// Destination port to run ping.
	PingPort int `json:"pingPort,omitempty" msgpack:"pingPort,omitempty" bson:"-" mapstructure:"pingPort,omitempty"`

	// ID unique per purefresh event.
	RefreshID string `json:"refreshID" msgpack:"refreshID" bson:"-" mapstructure:"refreshID,omitempty"`

	// If set to true, the target Processing Unit will refresh its policy immediately.
	RefreshPolicy bool `json:"refreshPolicy,omitempty" msgpack:"refreshPolicy,omitempty" bson:"-" mapstructure:"refreshPolicy,omitempty"`

	// Instructs the defender to send records for all
	// application-initiated connections for the target processing unit.
	TraceApplicationConnections bool `json:"traceApplicationConnections,omitempty" msgpack:"traceApplicationConnections,omitempty" bson:"-" mapstructure:"traceApplicationConnections,omitempty"`

	// Determines the length of the time interval that the trace must be
	// enabled, using [Golang duration
	// syntax](https://golang.org/pkg/time/#example_Duration).
	TraceDuration string `json:"traceDuration,omitempty" msgpack:"traceDuration,omitempty" bson:"-" mapstructure:"traceDuration,omitempty"`

	// Instructs the defenders to provide an iptables trace for the target processing
	// unit.
	TraceIPTables bool `json:"traceIPTables,omitempty" msgpack:"traceIPTables,omitempty" bson:"-" mapstructure:"traceIPTables,omitempty"`

	// Instructs the defender to send records for all
	// network-initiated connections for the target processing unit.
	TraceNetworkConnections bool `json:"traceNetworkConnections,omitempty" msgpack:"traceNetworkConnections,omitempty" bson:"-" mapstructure:"traceNetworkConnections,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewProcessingUnitRefresh returns a new *ProcessingUnitRefresh
func NewProcessingUnitRefresh() *ProcessingUnitRefresh {

	return &ProcessingUnitRefresh{
		ModelVersion:   1,
		PingIterations: 1,
		PingMode:       ProcessingUnitRefreshPingModeAuto,
		TraceDuration:  "10s",
	}
}

// Identity returns the Identity of the object.
func (o *ProcessingUnitRefresh) Identity() elemental.Identity {

	return ProcessingUnitRefreshIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ProcessingUnitRefresh) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ProcessingUnitRefresh) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ProcessingUnitRefresh) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesProcessingUnitRefresh{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ProcessingUnitRefresh) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesProcessingUnitRefresh{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *ProcessingUnitRefresh) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *ProcessingUnitRefresh) BleveType() string {

	return "processingunitrefresh"
}

// DefaultOrder returns the list of default ordering fields.
func (o *ProcessingUnitRefresh) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *ProcessingUnitRefresh) Doc() string {

	return `Sent to client when a poke has been triggered using the
parameter ` + "`" + `?notify=true` + "`" + `. This is used by instances of enforcerd to notify an
external change on the processing unit must be processed.`
}

func (o *ProcessingUnitRefresh) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetID returns the ID of the receiver.
func (o *ProcessingUnitRefresh) GetID() string {

	return o.ID
}

// SetID sets the property ID of the receiver using the given value.
func (o *ProcessingUnitRefresh) SetID(ID string) {

	o.ID = ID
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ProcessingUnitRefresh) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseProcessingUnitRefresh{
			ID:                          &o.ID,
			Debug:                       &o.Debug,
			Namespace:                   &o.Namespace,
			PingAddress:                 &o.PingAddress,
			PingEnabled:                 &o.PingEnabled,
			PingIterations:              &o.PingIterations,
			PingMode:                    &o.PingMode,
			PingPort:                    &o.PingPort,
			RefreshID:                   &o.RefreshID,
			RefreshPolicy:               &o.RefreshPolicy,
			TraceApplicationConnections: &o.TraceApplicationConnections,
			TraceDuration:               &o.TraceDuration,
			TraceIPTables:               &o.TraceIPTables,
			TraceNetworkConnections:     &o.TraceNetworkConnections,
		}
	}

	sp := &SparseProcessingUnitRefresh{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "debug":
			sp.Debug = &(o.Debug)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "pingAddress":
			sp.PingAddress = &(o.PingAddress)
		case "pingEnabled":
			sp.PingEnabled = &(o.PingEnabled)
		case "pingIterations":
			sp.PingIterations = &(o.PingIterations)
		case "pingMode":
			sp.PingMode = &(o.PingMode)
		case "pingPort":
			sp.PingPort = &(o.PingPort)
		case "refreshID":
			sp.RefreshID = &(o.RefreshID)
		case "refreshPolicy":
			sp.RefreshPolicy = &(o.RefreshPolicy)
		case "traceApplicationConnections":
			sp.TraceApplicationConnections = &(o.TraceApplicationConnections)
		case "traceDuration":
			sp.TraceDuration = &(o.TraceDuration)
		case "traceIPTables":
			sp.TraceIPTables = &(o.TraceIPTables)
		case "traceNetworkConnections":
			sp.TraceNetworkConnections = &(o.TraceNetworkConnections)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseProcessingUnitRefresh to the object.
func (o *ProcessingUnitRefresh) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseProcessingUnitRefresh)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Debug != nil {
		o.Debug = *so.Debug
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.PingAddress != nil {
		o.PingAddress = *so.PingAddress
	}
	if so.PingEnabled != nil {
		o.PingEnabled = *so.PingEnabled
	}
	if so.PingIterations != nil {
		o.PingIterations = *so.PingIterations
	}
	if so.PingMode != nil {
		o.PingMode = *so.PingMode
	}
	if so.PingPort != nil {
		o.PingPort = *so.PingPort
	}
	if so.RefreshID != nil {
		o.RefreshID = *so.RefreshID
	}
	if so.RefreshPolicy != nil {
		o.RefreshPolicy = *so.RefreshPolicy
	}
	if so.TraceApplicationConnections != nil {
		o.TraceApplicationConnections = *so.TraceApplicationConnections
	}
	if so.TraceDuration != nil {
		o.TraceDuration = *so.TraceDuration
	}
	if so.TraceIPTables != nil {
		o.TraceIPTables = *so.TraceIPTables
	}
	if so.TraceNetworkConnections != nil {
		o.TraceNetworkConnections = *so.TraceNetworkConnections
	}
}

// DeepCopy returns a deep copy if the ProcessingUnitRefresh.
func (o *ProcessingUnitRefresh) DeepCopy() *ProcessingUnitRefresh {

	if o == nil {
		return nil
	}

	out := &ProcessingUnitRefresh{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ProcessingUnitRefresh.
func (o *ProcessingUnitRefresh) DeepCopyInto(out *ProcessingUnitRefresh) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ProcessingUnitRefresh: %s", err))
	}

	*out = *target.(*ProcessingUnitRefresh)
}

// Validate valides the current information stored into the structure.
func (o *ProcessingUnitRefresh) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateMinimumInt("pingIterations", o.PingIterations, int(1), false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("pingMode", string(o.PingMode), []string{"Auto", "L3", "L4", "L7"}, false); err != nil {
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
func (*ProcessingUnitRefresh) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ProcessingUnitRefreshAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ProcessingUnitRefreshLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ProcessingUnitRefresh) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ProcessingUnitRefreshAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ProcessingUnitRefresh) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "debug":
		return o.Debug
	case "namespace":
		return o.Namespace
	case "pingAddress":
		return o.PingAddress
	case "pingEnabled":
		return o.PingEnabled
	case "pingIterations":
		return o.PingIterations
	case "pingMode":
		return o.PingMode
	case "pingPort":
		return o.PingPort
	case "refreshID":
		return o.RefreshID
	case "refreshPolicy":
		return o.RefreshPolicy
	case "traceApplicationConnections":
		return o.TraceApplicationConnections
	case "traceDuration":
		return o.TraceDuration
	case "traceIPTables":
		return o.TraceIPTables
	case "traceNetworkConnections":
		return o.TraceNetworkConnections
	}

	return nil
}

// ProcessingUnitRefreshAttributesMap represents the map of attribute for ProcessingUnitRefresh.
var ProcessingUnitRefreshAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `Contains the ID of the target processing unit.`,
		Exposed:        true,
		Getter:         true,
		Identifier:     true,
		Name:           "ID",
		ReadOnly:       true,
		Setter:         true,
		Type:           "string",
	},
	"Debug": {
		AllowedChoices: []string{},
		ConvertedName:  "Debug",
		Description: `If set to true, start reporting debug information for the target processing
unit.`,
		Exposed: true,
		Name:    "debug",
		Type:    "boolean",
	},
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Contains the original namespace of the processing unit.`,
		Exposed:        true,
		Name:           "namespace",
		ReadOnly:       true,
		Type:           "string",
	},
	"PingAddress": {
		AllowedChoices: []string{},
		ConvertedName:  "PingAddress",
		Description:    `Destination address to run ping.`,
		Exposed:        true,
		Name:           "pingAddress",
		Type:           "string",
	},
	"PingEnabled": {
		AllowedChoices: []string{},
		ConvertedName:  "PingEnabled",
		Description:    `If set to true, start ping to the destination.`,
		Exposed:        true,
		Name:           "pingEnabled",
		Type:           "boolean",
	},
	"PingIterations": {
		AllowedChoices: []string{},
		ConvertedName:  "PingIterations",
		DefaultValue:   1,
		Description:    `Number of iterations to run a ping probe.`,
		Exposed:        true,
		MinValue:       1,
		Name:           "pingIterations",
		Type:           "integer",
	},
	"PingMode": {
		AllowedChoices: []string{"Auto", "L3", "L4", "L7"},
		ConvertedName:  "PingMode",
		DefaultValue:   ProcessingUnitRefreshPingModeAuto,
		Description:    `Represents the mode of ping to be used.`,
		Exposed:        true,
		Name:           "pingMode",
		Type:           "enum",
	},
	"PingPort": {
		AllowedChoices: []string{},
		ConvertedName:  "PingPort",
		Description:    `Destination port to run ping.`,
		Exposed:        true,
		Name:           "pingPort",
		Type:           "integer",
	},
	"RefreshID": {
		AllowedChoices: []string{},
		ConvertedName:  "RefreshID",
		Description:    `ID unique per purefresh event.`,
		Exposed:        true,
		Name:           "refreshID",
		ReadOnly:       true,
		Type:           "string",
	},
	"RefreshPolicy": {
		AllowedChoices: []string{},
		ConvertedName:  "RefreshPolicy",
		Description:    `If set to true, the target Processing Unit will refresh its policy immediately.`,
		Exposed:        true,
		Name:           "refreshPolicy",
		Type:           "boolean",
	},
	"TraceApplicationConnections": {
		AllowedChoices: []string{},
		ConvertedName:  "TraceApplicationConnections",
		Description: `Instructs the defender to send records for all
application-initiated connections for the target processing unit.`,
		Exposed: true,
		Name:    "traceApplicationConnections",
		Type:    "boolean",
	},
	"TraceDuration": {
		AllowedChoices: []string{},
		ConvertedName:  "TraceDuration",
		DefaultValue:   "10s",
		Description: `Determines the length of the time interval that the trace must be
enabled, using [Golang duration
syntax](https://golang.org/pkg/time/#example_Duration).`,
		Exposed: true,
		Name:    "traceDuration",
		Type:    "string",
	},
	"TraceIPTables": {
		AllowedChoices: []string{},
		ConvertedName:  "TraceIPTables",
		Description: `Instructs the defenders to provide an iptables trace for the target processing
unit.`,
		Exposed: true,
		Name:    "traceIPTables",
		Type:    "boolean",
	},
	"TraceNetworkConnections": {
		AllowedChoices: []string{},
		ConvertedName:  "TraceNetworkConnections",
		Description: `Instructs the defender to send records for all
network-initiated connections for the target processing unit.`,
		Exposed: true,
		Name:    "traceNetworkConnections",
		Type:    "boolean",
	},
}

// ProcessingUnitRefreshLowerCaseAttributesMap represents the map of attribute for ProcessingUnitRefresh.
var ProcessingUnitRefreshLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `Contains the ID of the target processing unit.`,
		Exposed:        true,
		Getter:         true,
		Identifier:     true,
		Name:           "ID",
		ReadOnly:       true,
		Setter:         true,
		Type:           "string",
	},
	"debug": {
		AllowedChoices: []string{},
		ConvertedName:  "Debug",
		Description: `If set to true, start reporting debug information for the target processing
unit.`,
		Exposed: true,
		Name:    "debug",
		Type:    "boolean",
	},
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Contains the original namespace of the processing unit.`,
		Exposed:        true,
		Name:           "namespace",
		ReadOnly:       true,
		Type:           "string",
	},
	"pingaddress": {
		AllowedChoices: []string{},
		ConvertedName:  "PingAddress",
		Description:    `Destination address to run ping.`,
		Exposed:        true,
		Name:           "pingAddress",
		Type:           "string",
	},
	"pingenabled": {
		AllowedChoices: []string{},
		ConvertedName:  "PingEnabled",
		Description:    `If set to true, start ping to the destination.`,
		Exposed:        true,
		Name:           "pingEnabled",
		Type:           "boolean",
	},
	"pingiterations": {
		AllowedChoices: []string{},
		ConvertedName:  "PingIterations",
		DefaultValue:   1,
		Description:    `Number of iterations to run a ping probe.`,
		Exposed:        true,
		MinValue:       1,
		Name:           "pingIterations",
		Type:           "integer",
	},
	"pingmode": {
		AllowedChoices: []string{"Auto", "L3", "L4", "L7"},
		ConvertedName:  "PingMode",
		DefaultValue:   ProcessingUnitRefreshPingModeAuto,
		Description:    `Represents the mode of ping to be used.`,
		Exposed:        true,
		Name:           "pingMode",
		Type:           "enum",
	},
	"pingport": {
		AllowedChoices: []string{},
		ConvertedName:  "PingPort",
		Description:    `Destination port to run ping.`,
		Exposed:        true,
		Name:           "pingPort",
		Type:           "integer",
	},
	"refreshid": {
		AllowedChoices: []string{},
		ConvertedName:  "RefreshID",
		Description:    `ID unique per purefresh event.`,
		Exposed:        true,
		Name:           "refreshID",
		ReadOnly:       true,
		Type:           "string",
	},
	"refreshpolicy": {
		AllowedChoices: []string{},
		ConvertedName:  "RefreshPolicy",
		Description:    `If set to true, the target Processing Unit will refresh its policy immediately.`,
		Exposed:        true,
		Name:           "refreshPolicy",
		Type:           "boolean",
	},
	"traceapplicationconnections": {
		AllowedChoices: []string{},
		ConvertedName:  "TraceApplicationConnections",
		Description: `Instructs the defender to send records for all
application-initiated connections for the target processing unit.`,
		Exposed: true,
		Name:    "traceApplicationConnections",
		Type:    "boolean",
	},
	"traceduration": {
		AllowedChoices: []string{},
		ConvertedName:  "TraceDuration",
		DefaultValue:   "10s",
		Description: `Determines the length of the time interval that the trace must be
enabled, using [Golang duration
syntax](https://golang.org/pkg/time/#example_Duration).`,
		Exposed: true,
		Name:    "traceDuration",
		Type:    "string",
	},
	"traceiptables": {
		AllowedChoices: []string{},
		ConvertedName:  "TraceIPTables",
		Description: `Instructs the defenders to provide an iptables trace for the target processing
unit.`,
		Exposed: true,
		Name:    "traceIPTables",
		Type:    "boolean",
	},
	"tracenetworkconnections": {
		AllowedChoices: []string{},
		ConvertedName:  "TraceNetworkConnections",
		Description: `Instructs the defender to send records for all
network-initiated connections for the target processing unit.`,
		Exposed: true,
		Name:    "traceNetworkConnections",
		Type:    "boolean",
	},
}

// SparseProcessingUnitRefreshsList represents a list of SparseProcessingUnitRefreshs
type SparseProcessingUnitRefreshsList []*SparseProcessingUnitRefresh

// Identity returns the identity of the objects in the list.
func (o SparseProcessingUnitRefreshsList) Identity() elemental.Identity {

	return ProcessingUnitRefreshIdentity
}

// Copy returns a pointer to a copy the SparseProcessingUnitRefreshsList.
func (o SparseProcessingUnitRefreshsList) Copy() elemental.Identifiables {

	copy := append(SparseProcessingUnitRefreshsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseProcessingUnitRefreshsList.
func (o SparseProcessingUnitRefreshsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseProcessingUnitRefreshsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseProcessingUnitRefresh))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseProcessingUnitRefreshsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseProcessingUnitRefreshsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseProcessingUnitRefreshsList converted to ProcessingUnitRefreshsList.
func (o SparseProcessingUnitRefreshsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseProcessingUnitRefreshsList) Version() int {

	return 1
}

// SparseProcessingUnitRefresh represents the sparse version of a processingunitrefresh.
type SparseProcessingUnitRefresh struct {
	// Contains the ID of the target processing unit.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// If set to true, start reporting debug information for the target processing
	// unit.
	Debug *bool `json:"debug,omitempty" msgpack:"debug,omitempty" bson:"-" mapstructure:"debug,omitempty"`

	// Contains the original namespace of the processing unit.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"-" mapstructure:"namespace,omitempty"`

	// Destination address to run ping.
	PingAddress *string `json:"pingAddress,omitempty" msgpack:"pingAddress,omitempty" bson:"-" mapstructure:"pingAddress,omitempty"`

	// If set to true, start ping to the destination.
	PingEnabled *bool `json:"pingEnabled,omitempty" msgpack:"pingEnabled,omitempty" bson:"-" mapstructure:"pingEnabled,omitempty"`

	// Number of iterations to run a ping probe.
	PingIterations *int `json:"pingIterations,omitempty" msgpack:"pingIterations,omitempty" bson:"-" mapstructure:"pingIterations,omitempty"`

	// Represents the mode of ping to be used.
	PingMode *ProcessingUnitRefreshPingModeValue `json:"pingMode,omitempty" msgpack:"pingMode,omitempty" bson:"-" mapstructure:"pingMode,omitempty"`

	// Destination port to run ping.
	PingPort *int `json:"pingPort,omitempty" msgpack:"pingPort,omitempty" bson:"-" mapstructure:"pingPort,omitempty"`

	// ID unique per purefresh event.
	RefreshID *string `json:"refreshID,omitempty" msgpack:"refreshID,omitempty" bson:"-" mapstructure:"refreshID,omitempty"`

	// If set to true, the target Processing Unit will refresh its policy immediately.
	RefreshPolicy *bool `json:"refreshPolicy,omitempty" msgpack:"refreshPolicy,omitempty" bson:"-" mapstructure:"refreshPolicy,omitempty"`

	// Instructs the defender to send records for all
	// application-initiated connections for the target processing unit.
	TraceApplicationConnections *bool `json:"traceApplicationConnections,omitempty" msgpack:"traceApplicationConnections,omitempty" bson:"-" mapstructure:"traceApplicationConnections,omitempty"`

	// Determines the length of the time interval that the trace must be
	// enabled, using [Golang duration
	// syntax](https://golang.org/pkg/time/#example_Duration).
	TraceDuration *string `json:"traceDuration,omitempty" msgpack:"traceDuration,omitempty" bson:"-" mapstructure:"traceDuration,omitempty"`

	// Instructs the defenders to provide an iptables trace for the target processing
	// unit.
	TraceIPTables *bool `json:"traceIPTables,omitempty" msgpack:"traceIPTables,omitempty" bson:"-" mapstructure:"traceIPTables,omitempty"`

	// Instructs the defender to send records for all
	// network-initiated connections for the target processing unit.
	TraceNetworkConnections *bool `json:"traceNetworkConnections,omitempty" msgpack:"traceNetworkConnections,omitempty" bson:"-" mapstructure:"traceNetworkConnections,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseProcessingUnitRefresh returns a new  SparseProcessingUnitRefresh.
func NewSparseProcessingUnitRefresh() *SparseProcessingUnitRefresh {
	return &SparseProcessingUnitRefresh{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseProcessingUnitRefresh) Identity() elemental.Identity {

	return ProcessingUnitRefreshIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseProcessingUnitRefresh) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseProcessingUnitRefresh) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseProcessingUnitRefresh) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseProcessingUnitRefresh{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseProcessingUnitRefresh) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseProcessingUnitRefresh{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseProcessingUnitRefresh) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseProcessingUnitRefresh) ToPlain() elemental.PlainIdentifiable {

	out := NewProcessingUnitRefresh()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Debug != nil {
		out.Debug = *o.Debug
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.PingAddress != nil {
		out.PingAddress = *o.PingAddress
	}
	if o.PingEnabled != nil {
		out.PingEnabled = *o.PingEnabled
	}
	if o.PingIterations != nil {
		out.PingIterations = *o.PingIterations
	}
	if o.PingMode != nil {
		out.PingMode = *o.PingMode
	}
	if o.PingPort != nil {
		out.PingPort = *o.PingPort
	}
	if o.RefreshID != nil {
		out.RefreshID = *o.RefreshID
	}
	if o.RefreshPolicy != nil {
		out.RefreshPolicy = *o.RefreshPolicy
	}
	if o.TraceApplicationConnections != nil {
		out.TraceApplicationConnections = *o.TraceApplicationConnections
	}
	if o.TraceDuration != nil {
		out.TraceDuration = *o.TraceDuration
	}
	if o.TraceIPTables != nil {
		out.TraceIPTables = *o.TraceIPTables
	}
	if o.TraceNetworkConnections != nil {
		out.TraceNetworkConnections = *o.TraceNetworkConnections
	}

	return out
}

// GetID returns the ID of the receiver.
func (o *SparseProcessingUnitRefresh) GetID() (out string) {

	if o.ID == nil {
		return
	}

	return *o.ID
}

// SetID sets the property ID of the receiver using the address of the given value.
func (o *SparseProcessingUnitRefresh) SetID(ID string) {

	o.ID = &ID
}

// DeepCopy returns a deep copy if the SparseProcessingUnitRefresh.
func (o *SparseProcessingUnitRefresh) DeepCopy() *SparseProcessingUnitRefresh {

	if o == nil {
		return nil
	}

	out := &SparseProcessingUnitRefresh{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseProcessingUnitRefresh.
func (o *SparseProcessingUnitRefresh) DeepCopyInto(out *SparseProcessingUnitRefresh) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseProcessingUnitRefresh: %s", err))
	}

	*out = *target.(*SparseProcessingUnitRefresh)
}

type mongoAttributesProcessingUnitRefresh struct {
}
type mongoAttributesSparseProcessingUnitRefresh struct {
}
