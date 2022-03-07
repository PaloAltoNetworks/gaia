package types

import jwt "github.com/golang-jwt/jwt/v4"

// MidgardClaimsRestrictions represents permissions restrictions
// declared in the token.
type MidgardClaimsRestrictions struct {
	Permissions []string `msgpack:"perms,omitempty" json:"perms,omitempty"`
	Namespace   string   `msgpack:"namespace,omitempty" json:"namespace,omitempty"`
	Networks    []string `msgpack:"networks,omitempty" json:"networks,omitempty"`
}

// MidgardClaims is a struct to represeting the data some a Midgard issued claims.
type MidgardClaims struct {
	Realm        string                     `msgpack:"realm" json:"realm"`
	Quota        int                        `msgpack:"quota,omitempty" json:"quota,omitempty"`
	Data         map[string]string          `msgpack:"data" json:"data"`
	Opaque       map[string]string          `msgpack:"opaque,omitempty" json:"opaque,omitempty"`
	Restrictions *MidgardClaimsRestrictions `msgpack:"restrictions,omitempty" json:"restrictions,omitempty"`

	jwt.RegisteredClaims
}

// NewMidgardClaims returns a new Claims.
func NewMidgardClaims() *MidgardClaims {
	return &MidgardClaims{
		Data:             map[string]string{},
		RegisteredClaims: jwt.RegisteredClaims{},
	}
}

// ServiceToken is a struct to represent the service tokens issued by the system.
type ServiceToken struct {
	User    map[string]interface{} `json:"user,omitempty"`
	Service map[string]interface{} `json:"service,omitempty"`

	jwt.RegisteredClaims
}
