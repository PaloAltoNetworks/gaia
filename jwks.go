package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// JWKSAlgValue represents the possible values for attribute "alg".
type JWKSAlgValue string

const (
	// JWKSAlgES256 represents the value ES256.
	JWKSAlgES256 JWKSAlgValue = "ES256"

	// JWKSAlgRS256 represents the value RS256.
	JWKSAlgRS256 JWKSAlgValue = "RS256"
)

// JWKSCurveValue represents the possible values for attribute "curve".
type JWKSCurveValue string

const (
	// JWKSCurveP256 represents the value P256.
	JWKSCurveP256 JWKSCurveValue = "P256"
)

// JWKSKtyValue represents the possible values for attribute "kty".
type JWKSKtyValue string

const (
	// JWKSKtyEC represents the value EC.
	JWKSKtyEC JWKSKtyValue = "EC"

	// JWKSKtyRSA represents the value RSA.
	JWKSKtyRSA JWKSKtyValue = "RSA"
)

// JWKS represents the model of a jwks
type JWKS struct {
	// Alg defines the algorithm used for signing as per the JWKS specification (RFC
	// 7518 section 7).
	Alg JWKSAlgValue `json:"alg" msgpack:"alg" bson:"-" mapstructure:"alg,omitempty"`

	// Curve is the curve type used for signing. Valid only when signing method is EC
	// (RFC 7518 section 6).
	Curve JWKSCurveValue `json:"curve,omitempty" msgpack:"curve,omitempty" bson:"-" mapstructure:"curve,omitempty"`

	// E is the exponent value of an RSA public key. Valid only when the signing
	// method is RSA (RFC 7518 section 6).
	E string `json:"e,omitempty" msgpack:"e,omitempty" bson:"-" mapstructure:"e,omitempty"`

	// kid is the key ID as per the JWKS specification.
	Kid string `json:"kid" msgpack:"kid" bson:"-" mapstructure:"kid,omitempty"`

	// kty defines the key type as per the JWKS specification.
	Kty JWKSKtyValue `json:"kty" msgpack:"kty" bson:"-" mapstructure:"kty,omitempty"`

	// N is the modulos value of an RSA public key. Valid only when the signing
	// method is RSA (RFC 7518 section 6).
	N string `json:"n,omitempty" msgpack:"n,omitempty" bson:"-" mapstructure:"n,omitempty"`

	// Use defines the use of the signing key as per the JWKS specification.
	Use string `json:"use" msgpack:"use" bson:"-" mapstructure:"use,omitempty"`

	// X is the X value of the public key. Valid only when signing method is EC (RFC
	// 7518 section 6).
	X string `json:"x,omitempty" msgpack:"x,omitempty" bson:"-" mapstructure:"x,omitempty"`

	// Y is the Y value of the public key. Valid only when signing method is EC (RFC
	// 7518 section 6).
	Y string `json:"y,omitempty" msgpack:"y,omitempty" bson:"-" mapstructure:"y,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewJWKS returns a new *JWKS
func NewJWKS() *JWKS {

	return &JWKS{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *JWKS) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesJWKS{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *JWKS) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesJWKS{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *JWKS) BleveType() string {

	return "jwks"
}

// DeepCopy returns a deep copy if the JWKS.
func (o *JWKS) DeepCopy() *JWKS {

	if o == nil {
		return nil
	}

	out := &JWKS{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *JWKS.
func (o *JWKS) DeepCopyInto(out *JWKS) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy JWKS: %s", err))
	}

	*out = *target.(*JWKS)
}

// Validate valides the current information stored into the structure.
func (o *JWKS) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("alg", string(o.Alg), []string{"RS256", "ES256"}, true); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("curve", string(o.Curve), []string{"P256"}, true); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("kty", string(o.Kty), []string{"RSA", "EC"}, true); err != nil {
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

type mongoAttributesJWKS struct {
}
