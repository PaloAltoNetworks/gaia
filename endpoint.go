package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// Endpoint represents the model of a endpoint
type Endpoint struct {
	// URI of the exposed API.
	URI string `json:"URI" msgpack:"URI" bson:"uri" mapstructure:"URI,omitempty"`

	// The scopes authorized to access the API.
	AllowedScopes [][]string `json:"allowedScopes" msgpack:"allowedScopes" bson:"allowedscopes" mapstructure:"allowedScopes,omitempty"`

	// Methods exposed to access the API.
	Methods []string `json:"methods" msgpack:"methods" bson:"methods" mapstructure:"methods,omitempty"`

	// If `true`, the API is public.
	Public bool `json:"public" msgpack:"public" bson:"public" mapstructure:"public,omitempty"`

	// Use `allowedScopes`.
	Scopes []string `json:"scopes" msgpack:"scopes" bson:"-" mapstructure:"scopes,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewEndpoint returns a new *Endpoint
func NewEndpoint() *Endpoint {

	return &Endpoint{
		ModelVersion:  1,
		AllowedScopes: [][]string{},
		Methods:       []string{},
		Scopes:        []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Endpoint) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesEndpoint{}

	s.URI = o.URI
	s.AllowedScopes = o.AllowedScopes
	s.Methods = o.Methods
	s.Public = o.Public

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Endpoint) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesEndpoint{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.URI = s.URI
	o.AllowedScopes = s.AllowedScopes
	o.Methods = s.Methods
	o.Public = s.Public

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *Endpoint) BleveType() string {

	return "endpoint"
}

// DeepCopy returns a deep copy if the Endpoint.
func (o *Endpoint) DeepCopy() *Endpoint {

	if o == nil {
		return nil
	}

	out := &Endpoint{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Endpoint.
func (o *Endpoint) DeepCopyInto(out *Endpoint) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Endpoint: %s", err))
	}

	*out = *target.(*Endpoint)
}

// Validate valides the current information stored into the structure.
func (o *Endpoint) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateHTTPMethods("methods", o.Methods); err != nil {
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

type mongoAttributesEndpoint struct {
	URI           string     `bson:"uri"`
	AllowedScopes [][]string `bson:"allowedscopes"`
	Methods       []string   `bson:"methods"`
	Public        bool       `bson:"public"`
}
