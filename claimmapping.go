package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ClaimMapping represents the model of a claimmapping
type ClaimMapping struct {
	// The name of the claim to map to the HTTP header. header.
	ClaimName string `json:"claimName" msgpack:"claimName" bson:"claimname" mapstructure:"claimName,omitempty"`

	// The HTTP header that will be the destination of the mapped claim.
	TargetHTTPHeader string `json:"targetHTTPHeader" msgpack:"targetHTTPHeader" bson:"targethttpheader" mapstructure:"targetHTTPHeader,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewClaimMapping returns a new *ClaimMapping
func NewClaimMapping() *ClaimMapping {

	return &ClaimMapping{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ClaimMapping) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesClaimMapping{}

	s.ClaimName = o.ClaimName
	s.TargetHTTPHeader = o.TargetHTTPHeader

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ClaimMapping) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesClaimMapping{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ClaimName = s.ClaimName
	o.TargetHTTPHeader = s.TargetHTTPHeader

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *ClaimMapping) BleveType() string {

	return "claimmapping"
}

// DeepCopy returns a deep copy if the ClaimMapping.
func (o *ClaimMapping) DeepCopy() *ClaimMapping {

	if o == nil {
		return nil
	}

	out := &ClaimMapping{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ClaimMapping.
func (o *ClaimMapping) DeepCopyInto(out *ClaimMapping) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ClaimMapping: %s", err))
	}

	*out = *target.(*ClaimMapping)
}

// Validate valides the current information stored into the structure.
func (o *ClaimMapping) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("claimName", o.ClaimName); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidatePattern("claimName", o.ClaimName, `^[a-zA-Z0-9-_/*#&@\+\$~:]+$`, `must be an alpha numerical character or '-', '_', '/', '*', '#', '&', '@', '_', '$' ~ or ':'`, true); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("targetHTTPHeader", o.TargetHTTPHeader); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidatePattern("targetHTTPHeader", o.TargetHTTPHeader, `^[a-zA-Z0-9-_/*#&@\+\$~:]+$`, `must be an alpha numerical character or '-', '_', '/', '*', '#', '&', '@', '_', '$' ~ or ':'`, true); err != nil {
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

type mongoAttributesClaimMapping struct {
	ClaimName        string `bson:"claimname"`
	TargetHTTPHeader string `bson:"targethttpheader"`
}
