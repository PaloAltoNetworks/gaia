package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TraceMode represents the model of a tracemode
// +k8s:openapi-gen=true
type TraceMode struct {
	// Instructs the enforcers to provide an iptables trace for a processing unit.
	IPTables bool `json:"IPTables" msgpack:"IPTables" bson:"iptables" mapstructure:"IPTables,omitempty"`

	// Instructs the enforcer to send records for all
	// application-initiated connections.
	ApplicationConnections bool `json:"applicationConnections" msgpack:"applicationConnections" bson:"applicationconnections" mapstructure:"applicationConnections,omitempty"`

	// Determines the length of the time interval that the trace must be
	// enabled, using [Golang duration
	// syntax](https://golang.org/pkg/time/#example_Duration).
	Interval string `json:"interval" msgpack:"interval" bson:"interval" mapstructure:"interval,omitempty"`

	// Instructs the enforcer to send records for all
	// network-initiated connections.
	NetworkConnections bool `json:"networkConnections" msgpack:"networkConnections" bson:"networkconnections" mapstructure:"networkConnections,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewTraceMode returns a new *TraceMode
func NewTraceMode() *TraceMode {

	return &TraceMode{
		ModelVersion: 1,
		Interval:     "10s",
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TraceMode) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesTraceMode{}

	s.IPTables = o.IPTables
	s.ApplicationConnections = o.ApplicationConnections
	s.Interval = o.Interval
	s.NetworkConnections = o.NetworkConnections

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TraceMode) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesTraceMode{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.IPTables = s.IPTables
	o.ApplicationConnections = s.ApplicationConnections
	o.Interval = s.Interval
	o.NetworkConnections = s.NetworkConnections

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *TraceMode) BleveType() string {

	return "tracemode"
}

// DeepCopy returns a deep copy if the TraceMode.
func (o *TraceMode) DeepCopy() *TraceMode {

	if o == nil {
		return nil
	}

	out := &TraceMode{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *TraceMode.
func (o *TraceMode) DeepCopyInto(out *TraceMode) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy TraceMode: %s", err))
	}

	*out = *target.(*TraceMode)
}

// Validate valides the current information stored into the structure.
func (o *TraceMode) Validate() error {

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
func (*TraceMode) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TraceModeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TraceModeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*TraceMode) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TraceModeAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *TraceMode) ValueForAttribute(name string) interface{} {

	switch name {
	case "IPTables":
		return o.IPTables
	case "applicationConnections":
		return o.ApplicationConnections
	case "interval":
		return o.Interval
	case "networkConnections":
		return o.NetworkConnections
	}

	return nil
}

// TraceModeAttributesMap represents the map of attribute for TraceMode.
var TraceModeAttributesMap = map[string]elemental.AttributeSpecification{
	"IPTables": {
		AllowedChoices: []string{},
		BSONFieldName:  "iptables",
		ConvertedName:  "IPTables",
		Description:    `Instructs the enforcers to provide an iptables trace for a processing unit.`,
		Exposed:        true,
		Name:           "IPTables",
		Stored:         true,
		Type:           "boolean",
	},
	"ApplicationConnections": {
		AllowedChoices: []string{},
		BSONFieldName:  "applicationconnections",
		ConvertedName:  "ApplicationConnections",
		Description: `Instructs the enforcer to send records for all
application-initiated connections.`,
		Exposed: true,
		Name:    "applicationConnections",
		Stored:  true,
		Type:    "boolean",
	},
	"Interval": {
		AllowedChoices: []string{},
		BSONFieldName:  "interval",
		ConvertedName:  "Interval",
		DefaultValue:   "10s",
		Description: `Determines the length of the time interval that the trace must be
enabled, using [Golang duration
syntax](https://golang.org/pkg/time/#example_Duration).`,
		Exposed: true,
		Name:    "interval",
		Stored:  true,
		Type:    "string",
	},
	"NetworkConnections": {
		AllowedChoices: []string{},
		BSONFieldName:  "networkconnections",
		ConvertedName:  "NetworkConnections",
		Description: `Instructs the enforcer to send records for all
network-initiated connections.`,
		Exposed: true,
		Name:    "networkConnections",
		Stored:  true,
		Type:    "boolean",
	},
}

// TraceModeLowerCaseAttributesMap represents the map of attribute for TraceMode.
var TraceModeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"iptables": {
		AllowedChoices: []string{},
		BSONFieldName:  "iptables",
		ConvertedName:  "IPTables",
		Description:    `Instructs the enforcers to provide an iptables trace for a processing unit.`,
		Exposed:        true,
		Name:           "IPTables",
		Stored:         true,
		Type:           "boolean",
	},
	"applicationconnections": {
		AllowedChoices: []string{},
		BSONFieldName:  "applicationconnections",
		ConvertedName:  "ApplicationConnections",
		Description: `Instructs the enforcer to send records for all
application-initiated connections.`,
		Exposed: true,
		Name:    "applicationConnections",
		Stored:  true,
		Type:    "boolean",
	},
	"interval": {
		AllowedChoices: []string{},
		BSONFieldName:  "interval",
		ConvertedName:  "Interval",
		DefaultValue:   "10s",
		Description: `Determines the length of the time interval that the trace must be
enabled, using [Golang duration
syntax](https://golang.org/pkg/time/#example_Duration).`,
		Exposed: true,
		Name:    "interval",
		Stored:  true,
		Type:    "string",
	},
	"networkconnections": {
		AllowedChoices: []string{},
		BSONFieldName:  "networkconnections",
		ConvertedName:  "NetworkConnections",
		Description: `Instructs the enforcer to send records for all
network-initiated connections.`,
		Exposed: true,
		Name:    "networkConnections",
		Stored:  true,
		Type:    "boolean",
	},
}

type mongoAttributesTraceMode struct {
	IPTables               bool   `bson:"iptables"`
	ApplicationConnections bool   `bson:"applicationconnections"`
	Interval               string `bson:"interval"`
	NetworkConnections     bool   `bson:"networkconnections"`
}
