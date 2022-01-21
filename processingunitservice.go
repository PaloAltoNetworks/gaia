package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ProcessingUnitService represents the model of a processingunitservice
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

type mongoAttributesProcessingUnitService struct {
	Ports       string   `bson:"ports"`
	Protocol    int      `bson:"protocol"`
	TargetPorts []string `bson:"targetports"`
}
