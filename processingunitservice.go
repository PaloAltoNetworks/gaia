package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ProcessingUnitService represents the model of a processingunitservice
// +k8s:openapi-gen=true
type ProcessingUnitService struct {
	// Contains the list of allowed ports and ranges.
	Ports string `json:"ports" msgpack:"ports" bson:"ports" mapstructure:"ports,omitempty"`

	// Protocol used by the service.
	Protocol int `json:"protocol" msgpack:"protocol" bson:"protocol" mapstructure:"protocol,omitempty"`

	// List of single ports or range (xx:yy).
	TargetPorts []string `json:"targetPorts" msgpack:"targetPorts" bson:"targetports" mapstructure:"targetPorts,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewProcessingUnitService returns a new *ProcessingUnitService
func NewProcessingUnitService() *ProcessingUnitService {

	return &ProcessingUnitService{
		ModelVersion: 1,
		TargetPorts:  []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ProcessingUnitService) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesProcessingUnitService{}

	s.Ports = o.Ports
	s.Protocol = o.Protocol
	s.TargetPorts = o.TargetPorts

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ProcessingUnitService) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesProcessingUnitService{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Ports = s.Ports
	o.Protocol = s.Protocol
	o.TargetPorts = s.TargetPorts

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *ProcessingUnitService) BleveType() string {

	return "processingunitservice"
}

// DeepCopy returns a deep copy if the ProcessingUnitService.
func (o *ProcessingUnitService) DeepCopy() *ProcessingUnitService {

	if o == nil {
		return nil
	}

	out := &ProcessingUnitService{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ProcessingUnitService.
func (o *ProcessingUnitService) DeepCopyInto(out *ProcessingUnitService) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ProcessingUnitService: %s", err))
	}

	*out = *target.(*ProcessingUnitService)
}

// Validate valides the current information stored into the structure.
func (o *ProcessingUnitService) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidatePortStringList("targetPorts", o.TargetPorts); err != nil {
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
func (*ProcessingUnitService) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ProcessingUnitServiceAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ProcessingUnitServiceLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ProcessingUnitService) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ProcessingUnitServiceAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ProcessingUnitService) ValueForAttribute(name string) interface{} {

	switch name {
	case "ports":
		return o.Ports
	case "protocol":
		return o.Protocol
	case "targetPorts":
		return o.TargetPorts
	}

	return nil
}

// ProcessingUnitServiceAttributesMap represents the map of attribute for ProcessingUnitService.
var ProcessingUnitServiceAttributesMap = map[string]elemental.AttributeSpecification{
	"Ports": {
		AllowedChoices: []string{},
		BSONFieldName:  "ports",
		ConvertedName:  "Ports",
		Deprecated:     true,
		Description:    `Contains the list of allowed ports and ranges.`,
		Exposed:        true,
		Name:           "ports",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Protocol": {
		AllowedChoices: []string{},
		BSONFieldName:  "protocol",
		ConvertedName:  "Protocol",
		Description:    `Protocol used by the service.`,
		Exposed:        true,
		Name:           "protocol",
		Stored:         true,
		Type:           "integer",
	},
	"TargetPorts": {
		AllowedChoices: []string{},
		BSONFieldName:  "targetports",
		ConvertedName:  "TargetPorts",
		Description:    `List of single ports or range (xx:yy).`,
		Exposed:        true,
		Name:           "targetPorts",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
}

// ProcessingUnitServiceLowerCaseAttributesMap represents the map of attribute for ProcessingUnitService.
var ProcessingUnitServiceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"ports": {
		AllowedChoices: []string{},
		BSONFieldName:  "ports",
		ConvertedName:  "Ports",
		Deprecated:     true,
		Description:    `Contains the list of allowed ports and ranges.`,
		Exposed:        true,
		Name:           "ports",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"protocol": {
		AllowedChoices: []string{},
		BSONFieldName:  "protocol",
		ConvertedName:  "Protocol",
		Description:    `Protocol used by the service.`,
		Exposed:        true,
		Name:           "protocol",
		Stored:         true,
		Type:           "integer",
	},
	"targetports": {
		AllowedChoices: []string{},
		BSONFieldName:  "targetports",
		ConvertedName:  "TargetPorts",
		Description:    `List of single ports or range (xx:yy).`,
		Exposed:        true,
		Name:           "targetPorts",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
}

type mongoAttributesProcessingUnitService struct {
	Ports       string   `bson:"ports"`
	Protocol    int      `bson:"protocol"`
	TargetPorts []string `bson:"targetports"`
}
