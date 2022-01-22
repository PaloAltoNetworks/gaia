package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PingPair represents the model of a pingpair
type PingPair struct {
	// Contains the request probe information.
	Request *PingProbe `json:"request" msgpack:"request" bson:"request" mapstructure:"request,omitempty"`

	// Contains the response probe information.
	Response *PingProbe `json:"response" msgpack:"response" bson:"response" mapstructure:"response,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPingPair returns a new *PingPair
func NewPingPair() *PingPair {

	return &PingPair{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PingPair) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPingPair{}

	s.Request = o.Request
	s.Response = o.Response

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PingPair) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPingPair{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Request = s.Request
	o.Response = s.Response

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *PingPair) BleveType() string {

	return "pingpair"
}

// DeepCopy returns a deep copy if the PingPair.
func (o *PingPair) DeepCopy() *PingPair {

	if o == nil {
		return nil
	}

	out := &PingPair{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PingPair.
func (o *PingPair) DeepCopyInto(out *PingPair) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PingPair: %s", err))
	}

	*out = *target.(*PingPair)
}

// Validate valides the current information stored into the structure.
func (o *PingPair) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if o.Request != nil {
		elemental.ResetDefaultForZeroValues(o.Request)
		if err := o.Request.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if o.Response != nil {
		elemental.ResetDefaultForZeroValues(o.Response)
		if err := o.Response.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

type mongoAttributesPingPair struct {
	Request  *PingProbe `bson:"request"`
	Response *PingProbe `bson:"response"`
}
