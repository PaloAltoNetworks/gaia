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

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*JWKS) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := JWKSAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return JWKSLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*JWKS) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return JWKSAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *JWKS) ValueForAttribute(name string) interface{} {

	switch name {
	case "alg":
		return o.Alg
	case "curve":
		return o.Curve
	case "e":
		return o.E
	case "kid":
		return o.Kid
	case "kty":
		return o.Kty
	case "n":
		return o.N
	case "use":
		return o.Use
	case "x":
		return o.X
	case "y":
		return o.Y
	}

	return nil
}

// JWKSAttributesMap represents the map of attribute for JWKS.
var JWKSAttributesMap = map[string]elemental.AttributeSpecification{
	"Alg": {
		AllowedChoices: []string{"RS256", "ES256"},
		Autogenerated:  true,
		ConvertedName:  "Alg",
		Description: `Alg defines the algorithm used for signing as per the JWKS specification (RFC
7518 section 7).`,
		Exposed:  true,
		Name:     "alg",
		ReadOnly: true,
		Type:     "enum",
	},
	"Curve": {
		AllowedChoices: []string{"P256"},
		Autogenerated:  true,
		ConvertedName:  "Curve",
		Description: `Curve is the curve type used for signing. Valid only when signing method is EC
(RFC 7518 section 6).`,
		Exposed:  true,
		Name:     "curve",
		ReadOnly: true,
		Type:     "enum",
	},
	"E": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "E",
		Description: `E is the exponent value of an RSA public key. Valid only when the signing
method is RSA (RFC 7518 section 6).`,
		Exposed:  true,
		Name:     "e",
		ReadOnly: true,
		Type:     "string",
	},
	"Kid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Kid",
		Description:    `kid is the key ID as per the JWKS specification.`,
		Exposed:        true,
		Name:           "kid",
		ReadOnly:       true,
		Type:           "string",
	},
	"Kty": {
		AllowedChoices: []string{"RSA", "EC"},
		Autogenerated:  true,
		ConvertedName:  "Kty",
		Description:    `kty defines the key type as per the JWKS specification.`,
		Exposed:        true,
		Name:           "kty",
		ReadOnly:       true,
		Type:           "enum",
	},
	"N": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "N",
		Description: `N is the modulos value of an RSA public key. Valid only when the signing
method is RSA (RFC 7518 section 6).`,
		Exposed:  true,
		Name:     "n",
		ReadOnly: true,
		Type:     "string",
	},
	"Use": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Use",
		Description:    `Use defines the use of the signing key as per the JWKS specification.`,
		Exposed:        true,
		Name:           "use",
		ReadOnly:       true,
		Type:           "string",
	},
	"X": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "X",
		Description: `X is the X value of the public key. Valid only when signing method is EC (RFC
7518 section 6).`,
		Exposed:  true,
		Name:     "x",
		ReadOnly: true,
		Type:     "string",
	},
	"Y": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Y",
		Description: `Y is the Y value of the public key. Valid only when signing method is EC (RFC
7518 section 6).`,
		Exposed:  true,
		Name:     "y",
		ReadOnly: true,
		Type:     "string",
	},
}

// JWKSLowerCaseAttributesMap represents the map of attribute for JWKS.
var JWKSLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"alg": {
		AllowedChoices: []string{"RS256", "ES256"},
		Autogenerated:  true,
		ConvertedName:  "Alg",
		Description: `Alg defines the algorithm used for signing as per the JWKS specification (RFC
7518 section 7).`,
		Exposed:  true,
		Name:     "alg",
		ReadOnly: true,
		Type:     "enum",
	},
	"curve": {
		AllowedChoices: []string{"P256"},
		Autogenerated:  true,
		ConvertedName:  "Curve",
		Description: `Curve is the curve type used for signing. Valid only when signing method is EC
(RFC 7518 section 6).`,
		Exposed:  true,
		Name:     "curve",
		ReadOnly: true,
		Type:     "enum",
	},
	"e": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "E",
		Description: `E is the exponent value of an RSA public key. Valid only when the signing
method is RSA (RFC 7518 section 6).`,
		Exposed:  true,
		Name:     "e",
		ReadOnly: true,
		Type:     "string",
	},
	"kid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Kid",
		Description:    `kid is the key ID as per the JWKS specification.`,
		Exposed:        true,
		Name:           "kid",
		ReadOnly:       true,
		Type:           "string",
	},
	"kty": {
		AllowedChoices: []string{"RSA", "EC"},
		Autogenerated:  true,
		ConvertedName:  "Kty",
		Description:    `kty defines the key type as per the JWKS specification.`,
		Exposed:        true,
		Name:           "kty",
		ReadOnly:       true,
		Type:           "enum",
	},
	"n": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "N",
		Description: `N is the modulos value of an RSA public key. Valid only when the signing
method is RSA (RFC 7518 section 6).`,
		Exposed:  true,
		Name:     "n",
		ReadOnly: true,
		Type:     "string",
	},
	"use": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Use",
		Description:    `Use defines the use of the signing key as per the JWKS specification.`,
		Exposed:        true,
		Name:           "use",
		ReadOnly:       true,
		Type:           "string",
	},
	"x": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "X",
		Description: `X is the X value of the public key. Valid only when signing method is EC (RFC
7518 section 6).`,
		Exposed:  true,
		Name:     "x",
		ReadOnly: true,
		Type:     "string",
	},
	"y": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Y",
		Description: `Y is the Y value of the public key. Valid only when signing method is EC (RFC
7518 section 6).`,
		Exposed:  true,
		Name:     "y",
		ReadOnly: true,
		Type:     "string",
	},
}

type mongoAttributesJWKS struct {
}
