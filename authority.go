package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AuthorityAlgorithmValue represents the possible values for attribute "algorithm".
type AuthorityAlgorithmValue string

const (
	// AuthorityAlgorithmECDSA represents the value ECDSA.
	AuthorityAlgorithmECDSA AuthorityAlgorithmValue = "ECDSA"

	// AuthorityAlgorithmRSA represents the value RSA.
	AuthorityAlgorithmRSA AuthorityAlgorithmValue = "RSA"
)

// AuthorityTypeValue represents the possible values for attribute "type".
type AuthorityTypeValue string

const (
	// AuthorityTypeCA represents the value CA.
	AuthorityTypeCA AuthorityTypeValue = "CA"

	// AuthorityTypeTokenSigning represents the value TokenSigning.
	AuthorityTypeTokenSigning AuthorityTypeValue = "TokenSigning"
)

// AuthorityIdentity represents the Identity of the object.
var AuthorityIdentity = elemental.Identity{
	Name:     "authority",
	Category: "authorities",
	Package:  "barret",
	Private:  true,
}

// AuthoritiesList represents a list of Authorities
type AuthoritiesList []*Authority

// Identity returns the identity of the objects in the list.
func (o AuthoritiesList) Identity() elemental.Identity {

	return AuthorityIdentity
}

// Copy returns a pointer to a copy the AuthoritiesList.
func (o AuthoritiesList) Copy() elemental.Identifiables {

	copy := append(AuthoritiesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AuthoritiesList.
func (o AuthoritiesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AuthoritiesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Authority))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AuthoritiesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AuthoritiesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the AuthoritiesList converted to SparseAuthoritiesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AuthoritiesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAuthoritiesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAuthority)
	}

	return out
}

// Version returns the version of the content.
func (o AuthoritiesList) Version() int {

	return 1
}

// Authority represents the model of a authority
type Authority struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Algorithm defines the signing algorithm to be used.
	Algorithm AuthorityAlgorithmValue `json:"algorithm" msgpack:"algorithm" bson:"algorithm" mapstructure:"algorithm,omitempty"`

	// PEM encoded certificate data.
	Certificate string `json:"certificate" msgpack:"certificate" bson:"certificate" mapstructure:"certificate,omitempty"`

	// CommonName contains the common name of the certificate.
	CommonName string `json:"commonName" msgpack:"commonName" bson:"commonname" mapstructure:"commonName,omitempty"`

	// Date of expiration of the issued certificate.
	ExpirationDate time.Time `json:"expirationDate" msgpack:"expirationDate" bson:"expirationdate" mapstructure:"expirationDate,omitempty"`

	// Encrypted private key of the Authority.
	Key string `json:"-" msgpack:"-" bson:"key" mapstructure:"-,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Organization field of the CA.
	Organization string `json:"organization" msgpack:"organization" bson:"organization" mapstructure:"organization,omitempty"`

	// serialNumber of the certificate.
	SerialNumber string `json:"serialNumber" msgpack:"serialNumber" bson:"serialnumber" mapstructure:"serialNumber,omitempty"`

	// Type of signing authority can be a CA or a JWT signing certificate.
	Type AuthorityTypeValue `json:"type" msgpack:"type" bson:"type" mapstructure:"type,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAuthority returns a new *Authority
func NewAuthority() *Authority {

	return &Authority{
		ModelVersion:  1,
		Algorithm:     AuthorityAlgorithmECDSA,
		MigrationsLog: map[string]string{},
		Type:          AuthorityTypeCA,
	}
}

// Identity returns the Identity of the object.
func (o *Authority) Identity() elemental.Identity {

	return AuthorityIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Authority) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Authority) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Authority) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesAuthority{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Algorithm = o.Algorithm
	s.Certificate = o.Certificate
	s.CommonName = o.CommonName
	s.ExpirationDate = o.ExpirationDate
	s.Key = o.Key
	s.MigrationsLog = o.MigrationsLog
	s.Organization = o.Organization
	s.SerialNumber = o.SerialNumber
	s.Type = o.Type
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Authority) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesAuthority{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Algorithm = s.Algorithm
	o.Certificate = s.Certificate
	o.CommonName = s.CommonName
	o.ExpirationDate = s.ExpirationDate
	o.Key = s.Key
	o.MigrationsLog = s.MigrationsLog
	o.Organization = s.Organization
	o.SerialNumber = s.SerialNumber
	o.Type = s.Type
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Authority) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Authority) BleveType() string {

	return "authority"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Authority) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Authority) Doc() string {

	return `Authority represents a certificate authority.`
}

func (o *Authority) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *Authority) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *Authority) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetZHash returns the ZHash of the receiver.
func (o *Authority) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *Authority) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *Authority) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *Authority) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Authority) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAuthority{
			ID:             &o.ID,
			Algorithm:      &o.Algorithm,
			Certificate:    &o.Certificate,
			CommonName:     &o.CommonName,
			ExpirationDate: &o.ExpirationDate,
			Key:            &o.Key,
			MigrationsLog:  &o.MigrationsLog,
			Organization:   &o.Organization,
			SerialNumber:   &o.SerialNumber,
			Type:           &o.Type,
			ZHash:          &o.ZHash,
			Zone:           &o.Zone,
		}
	}

	sp := &SparseAuthority{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "algorithm":
			sp.Algorithm = &(o.Algorithm)
		case "certificate":
			sp.Certificate = &(o.Certificate)
		case "commonName":
			sp.CommonName = &(o.CommonName)
		case "expirationDate":
			sp.ExpirationDate = &(o.ExpirationDate)
		case "key":
			sp.Key = &(o.Key)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "organization":
			sp.Organization = &(o.Organization)
		case "serialNumber":
			sp.SerialNumber = &(o.SerialNumber)
		case "type":
			sp.Type = &(o.Type)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// EncryptAttributes encrypts the attributes marked as `encrypted` using the given encrypter.
func (o *Authority) EncryptAttributes(encrypter elemental.AttributeEncrypter) (err error) {

	if o.Key, err = encrypter.EncryptString(o.Key); err != nil {
		return fmt.Errorf("unable to encrypt attribute 'Key' for 'Authority' (%s): %s", o.Identifier(), err)
	}

	return nil
}

// DecryptAttributes decrypts the attributes marked as `encrypted` using the given decrypter.
func (o *Authority) DecryptAttributes(encrypter elemental.AttributeEncrypter) (err error) {

	if o.Key, err = encrypter.DecryptString(o.Key); err != nil {
		return fmt.Errorf("unable to decrypt attribute 'Key' for 'Authority' (%s): %s", o.Identifier(), err)
	}

	return nil
}

// Patch apply the non nil value of a *SparseAuthority to the object.
func (o *Authority) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAuthority)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Algorithm != nil {
		o.Algorithm = *so.Algorithm
	}
	if so.Certificate != nil {
		o.Certificate = *so.Certificate
	}
	if so.CommonName != nil {
		o.CommonName = *so.CommonName
	}
	if so.ExpirationDate != nil {
		o.ExpirationDate = *so.ExpirationDate
	}
	if so.Key != nil {
		o.Key = *so.Key
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Organization != nil {
		o.Organization = *so.Organization
	}
	if so.SerialNumber != nil {
		o.SerialNumber = *so.SerialNumber
	}
	if so.Type != nil {
		o.Type = *so.Type
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the Authority.
func (o *Authority) DeepCopy() *Authority {

	if o == nil {
		return nil
	}

	out := &Authority{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Authority.
func (o *Authority) DeepCopyInto(out *Authority) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Authority: %s", err))
	}

	*out = *target.(*Authority)
}

// Validate valides the current information stored into the structure.
func (o *Authority) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("algorithm", string(o.Algorithm), []string{"ECDSA", "RSA"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("commonName", o.CommonName); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"CA", "TokenSigning"}, true); err != nil {
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
func (*Authority) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AuthorityAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AuthorityLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Authority) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AuthorityAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Authority) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "algorithm":
		return o.Algorithm
	case "certificate":
		return o.Certificate
	case "commonName":
		return o.CommonName
	case "expirationDate":
		return o.ExpirationDate
	case "key":
		return o.Key
	case "migrationsLog":
		return o.MigrationsLog
	case "organization":
		return o.Organization
	case "serialNumber":
		return o.SerialNumber
	case "type":
		return o.Type
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// AuthorityAttributesMap represents the map of attribute for Authority.
var AuthorityAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Algorithm": {
		AllowedChoices: []string{"ECDSA", "RSA"},
		ConvertedName:  "Algorithm",
		DefaultValue:   AuthorityAlgorithmECDSA,
		Description:    `Algorithm defines the signing algorithm to be used.`,
		Exposed:        true,
		Name:           "algorithm",
		Stored:         true,
		Type:           "enum",
	},
	"Certificate": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Certificate",
		Description:    `PEM encoded certificate data.`,
		Exposed:        true,
		Name:           "certificate",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"CommonName": {
		AllowedChoices: []string{},
		ConvertedName:  "CommonName",
		CreationOnly:   true,
		Description:    `CommonName contains the common name of the certificate.`,
		Exposed:        true,
		Name:           "commonName",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"ExpirationDate": {
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationDate",
		CreationOnly:   true,
		Description:    `Date of expiration of the issued certificate.`,
		Exposed:        true,
		Name:           "expirationDate",
		Stored:         true,
		Type:           "time",
	},
	"Key": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Key",
		Description:    `Encrypted private key of the Authority.`,
		Encrypted:      true,
		Name:           "key",
		Stored:         true,
		Type:           "string",
	},
	"MigrationsLog": {
		AllowedChoices: []string{},
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Organization": {
		AllowedChoices: []string{},
		ConvertedName:  "Organization",
		CreationOnly:   true,
		Description:    `Organization field of the CA.`,
		Exposed:        true,
		Name:           "organization",
		Stored:         true,
		Type:           "string",
	},
	"SerialNumber": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SerialNumber",
		Description:    `serialNumber of the certificate.`,
		Exposed:        true,
		Name:           "serialNumber",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Type": {
		AllowedChoices: []string{"CA", "TokenSigning"},
		Autogenerated:  true,
		ConvertedName:  "Type",
		DefaultValue:   AuthorityTypeCA,
		Description:    `Type of signing authority can be a CA or a JWT signing certificate.`,
		Exposed:        true,
		Name:           "type",
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
	},
	"ZHash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// AuthorityLowerCaseAttributesMap represents the map of attribute for Authority.
var AuthorityLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"algorithm": {
		AllowedChoices: []string{"ECDSA", "RSA"},
		ConvertedName:  "Algorithm",
		DefaultValue:   AuthorityAlgorithmECDSA,
		Description:    `Algorithm defines the signing algorithm to be used.`,
		Exposed:        true,
		Name:           "algorithm",
		Stored:         true,
		Type:           "enum",
	},
	"certificate": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Certificate",
		Description:    `PEM encoded certificate data.`,
		Exposed:        true,
		Name:           "certificate",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"commonname": {
		AllowedChoices: []string{},
		ConvertedName:  "CommonName",
		CreationOnly:   true,
		Description:    `CommonName contains the common name of the certificate.`,
		Exposed:        true,
		Name:           "commonName",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"expirationdate": {
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationDate",
		CreationOnly:   true,
		Description:    `Date of expiration of the issued certificate.`,
		Exposed:        true,
		Name:           "expirationDate",
		Stored:         true,
		Type:           "time",
	},
	"key": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Key",
		Description:    `Encrypted private key of the Authority.`,
		Encrypted:      true,
		Name:           "key",
		Stored:         true,
		Type:           "string",
	},
	"migrationslog": {
		AllowedChoices: []string{},
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"organization": {
		AllowedChoices: []string{},
		ConvertedName:  "Organization",
		CreationOnly:   true,
		Description:    `Organization field of the CA.`,
		Exposed:        true,
		Name:           "organization",
		Stored:         true,
		Type:           "string",
	},
	"serialnumber": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SerialNumber",
		Description:    `serialNumber of the certificate.`,
		Exposed:        true,
		Name:           "serialNumber",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"type": {
		AllowedChoices: []string{"CA", "TokenSigning"},
		Autogenerated:  true,
		ConvertedName:  "Type",
		DefaultValue:   AuthorityTypeCA,
		Description:    `Type of signing authority can be a CA or a JWT signing certificate.`,
		Exposed:        true,
		Name:           "type",
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
	},
	"zhash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SparseAuthoritiesList represents a list of SparseAuthorities
type SparseAuthoritiesList []*SparseAuthority

// Identity returns the identity of the objects in the list.
func (o SparseAuthoritiesList) Identity() elemental.Identity {

	return AuthorityIdentity
}

// Copy returns a pointer to a copy the SparseAuthoritiesList.
func (o SparseAuthoritiesList) Copy() elemental.Identifiables {

	copy := append(SparseAuthoritiesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAuthoritiesList.
func (o SparseAuthoritiesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAuthoritiesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAuthority))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAuthoritiesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAuthoritiesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseAuthoritiesList converted to AuthoritiesList.
func (o SparseAuthoritiesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAuthoritiesList) Version() int {

	return 1
}

// SparseAuthority represents the sparse version of a authority.
type SparseAuthority struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Algorithm defines the signing algorithm to be used.
	Algorithm *AuthorityAlgorithmValue `json:"algorithm,omitempty" msgpack:"algorithm,omitempty" bson:"algorithm,omitempty" mapstructure:"algorithm,omitempty"`

	// PEM encoded certificate data.
	Certificate *string `json:"certificate,omitempty" msgpack:"certificate,omitempty" bson:"certificate,omitempty" mapstructure:"certificate,omitempty"`

	// CommonName contains the common name of the certificate.
	CommonName *string `json:"commonName,omitempty" msgpack:"commonName,omitempty" bson:"commonname,omitempty" mapstructure:"commonName,omitempty"`

	// Date of expiration of the issued certificate.
	ExpirationDate *time.Time `json:"expirationDate,omitempty" msgpack:"expirationDate,omitempty" bson:"expirationdate,omitempty" mapstructure:"expirationDate,omitempty"`

	// Encrypted private key of the Authority.
	Key *string `json:"-" msgpack:"-" bson:"key,omitempty" mapstructure:"-,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Organization field of the CA.
	Organization *string `json:"organization,omitempty" msgpack:"organization,omitempty" bson:"organization,omitempty" mapstructure:"organization,omitempty"`

	// serialNumber of the certificate.
	SerialNumber *string `json:"serialNumber,omitempty" msgpack:"serialNumber,omitempty" bson:"serialnumber,omitempty" mapstructure:"serialNumber,omitempty"`

	// Type of signing authority can be a CA or a JWT signing certificate.
	Type *AuthorityTypeValue `json:"type,omitempty" msgpack:"type,omitempty" bson:"type,omitempty" mapstructure:"type,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAuthority returns a new  SparseAuthority.
func NewSparseAuthority() *SparseAuthority {
	return &SparseAuthority{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAuthority) Identity() elemental.Identity {

	return AuthorityIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAuthority) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAuthority) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAuthority) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseAuthority{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Algorithm != nil {
		s.Algorithm = o.Algorithm
	}
	if o.Certificate != nil {
		s.Certificate = o.Certificate
	}
	if o.CommonName != nil {
		s.CommonName = o.CommonName
	}
	if o.ExpirationDate != nil {
		s.ExpirationDate = o.ExpirationDate
	}
	if o.Key != nil {
		s.Key = o.Key
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Organization != nil {
		s.Organization = o.Organization
	}
	if o.SerialNumber != nil {
		s.SerialNumber = o.SerialNumber
	}
	if o.Type != nil {
		s.Type = o.Type
	}
	if o.ZHash != nil {
		s.ZHash = o.ZHash
	}
	if o.Zone != nil {
		s.Zone = o.Zone
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAuthority) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseAuthority{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Algorithm != nil {
		o.Algorithm = s.Algorithm
	}
	if s.Certificate != nil {
		o.Certificate = s.Certificate
	}
	if s.CommonName != nil {
		o.CommonName = s.CommonName
	}
	if s.ExpirationDate != nil {
		o.ExpirationDate = s.ExpirationDate
	}
	if s.Key != nil {
		o.Key = s.Key
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Organization != nil {
		o.Organization = s.Organization
	}
	if s.SerialNumber != nil {
		o.SerialNumber = s.SerialNumber
	}
	if s.Type != nil {
		o.Type = s.Type
	}
	if s.ZHash != nil {
		o.ZHash = s.ZHash
	}
	if s.Zone != nil {
		o.Zone = s.Zone
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseAuthority) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAuthority) ToPlain() elemental.PlainIdentifiable {

	out := NewAuthority()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Algorithm != nil {
		out.Algorithm = *o.Algorithm
	}
	if o.Certificate != nil {
		out.Certificate = *o.Certificate
	}
	if o.CommonName != nil {
		out.CommonName = *o.CommonName
	}
	if o.ExpirationDate != nil {
		out.ExpirationDate = *o.ExpirationDate
	}
	if o.Key != nil {
		out.Key = *o.Key
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Organization != nil {
		out.Organization = *o.Organization
	}
	if o.SerialNumber != nil {
		out.SerialNumber = *o.SerialNumber
	}
	if o.Type != nil {
		out.Type = *o.Type
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// EncryptAttributes encrypts the attributes marked as `encrypted` using the given encrypter.
func (o *SparseAuthority) EncryptAttributes(encrypter elemental.AttributeEncrypter) (err error) {

	if *o.Key, err = encrypter.EncryptString(*o.Key); err != nil {
		return fmt.Errorf("unable to encrypt attribute 'Key' for 'SparseAuthority' (%s): %s", o.Identifier(), err)
	}

	return nil
}

// DecryptAttributes decrypts the attributes marked as `encrypted` using the given decrypter.
func (o *SparseAuthority) DecryptAttributes(encrypter elemental.AttributeEncrypter) (err error) {

	if *o.Key, err = encrypter.DecryptString(*o.Key); err != nil {
		return fmt.Errorf("unable to decrypt attribute 'Key' for 'SparseAuthority' (%s): %s", o.Identifier(), err)
	}

	return nil
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseAuthority) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseAuthority) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseAuthority) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseAuthority) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseAuthority) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseAuthority) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseAuthority.
func (o *SparseAuthority) DeepCopy() *SparseAuthority {

	if o == nil {
		return nil
	}

	out := &SparseAuthority{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAuthority.
func (o *SparseAuthority) DeepCopyInto(out *SparseAuthority) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAuthority: %s", err))
	}

	*out = *target.(*SparseAuthority)
}

type mongoAttributesAuthority struct {
	ID             bson.ObjectId           `bson:"_id,omitempty"`
	Algorithm      AuthorityAlgorithmValue `bson:"algorithm"`
	Certificate    string                  `bson:"certificate"`
	CommonName     string                  `bson:"commonname"`
	ExpirationDate time.Time               `bson:"expirationdate"`
	Key            string                  `bson:"key"`
	MigrationsLog  map[string]string       `bson:"migrationslog,omitempty"`
	Organization   string                  `bson:"organization"`
	SerialNumber   string                  `bson:"serialnumber"`
	Type           AuthorityTypeValue      `bson:"type"`
	ZHash          int                     `bson:"zhash"`
	Zone           int                     `bson:"zone"`
}
type mongoAttributesSparseAuthority struct {
	ID             bson.ObjectId            `bson:"_id,omitempty"`
	Algorithm      *AuthorityAlgorithmValue `bson:"algorithm,omitempty"`
	Certificate    *string                  `bson:"certificate,omitempty"`
	CommonName     *string                  `bson:"commonname,omitempty"`
	ExpirationDate *time.Time               `bson:"expirationdate,omitempty"`
	Key            *string                  `bson:"key,omitempty"`
	MigrationsLog  *map[string]string       `bson:"migrationslog,omitempty"`
	Organization   *string                  `bson:"organization,omitempty"`
	SerialNumber   *string                  `bson:"serialnumber,omitempty"`
	Type           *AuthorityTypeValue      `bson:"type,omitempty"`
	ZHash          *int                     `bson:"zhash,omitempty"`
	Zone           *int                     `bson:"zone,omitempty"`
}
