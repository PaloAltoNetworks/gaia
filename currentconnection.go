package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CurrentConnection represents the model of a currentconnection
type CurrentConnection struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// The duration of the tracked connection.
	Duration string `json:"duration,omitempty" msgpack:"duration,omitempty" bson:"-" mapstructure:"duration,omitempty"`

	// Was the connection existing when the enforcer started.
	Existing bool `json:"existing,omitempty" msgpack:"existing,omitempty" bson:"-" mapstructure:"existing,omitempty"`

	// The flow report for this connection.
	Flow *FlowReport `json:"flow,omitempty" msgpack:"flow,omitempty" bson:"-" mapstructure:"flow,omitempty"`

	// Port of the source.
	SourcePort int `json:"sourcePort,omitempty" msgpack:"sourcePort,omitempty" bson:"-" mapstructure:"sourcePort,omitempty"`

	// The time the enforcer started tracking the connection.
	StartTime time.Time `json:"startTime,omitempty" msgpack:"startTime,omitempty" bson:"-" mapstructure:"startTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCurrentConnection returns a new *CurrentConnection
func NewCurrentConnection() *CurrentConnection {

	return &CurrentConnection{
		ModelVersion: 1,
		Flow:         NewFlowReport(),
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CurrentConnection) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCurrentConnection{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CurrentConnection) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCurrentConnection{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CurrentConnection) BleveType() string {

	return "currentconnection"
}

// DeepCopy returns a deep copy if the CurrentConnection.
func (o *CurrentConnection) DeepCopy() *CurrentConnection {

	if o == nil {
		return nil
	}

	out := &CurrentConnection{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CurrentConnection.
func (o *CurrentConnection) DeepCopyInto(out *CurrentConnection) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CurrentConnection: %s", err))
	}

	*out = *target.(*CurrentConnection)
}

// Validate valides the current information stored into the structure.
func (o *CurrentConnection) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTimeDuration("duration", o.Duration); err != nil {
		errors = errors.Append(err)
	}

	if o.Flow != nil {
		elemental.ResetDefaultForZeroValues(o.Flow)
		if err := o.Flow.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateMaximumInt("sourcePort", o.SourcePort, int(65536), false); err != nil {
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

type mongoAttributesCurrentConnection struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
}
