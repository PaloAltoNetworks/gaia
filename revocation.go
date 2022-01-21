package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// RevocationIdentity represents the Identity of the object.
var RevocationIdentity = elemental.Identity{
	Name:     "revocation",
	Category: "revocations",
	Package:  "barret",
	Private:  true,
}

// RevocationsList represents a list of Revocations
type RevocationsList []*Revocation

// Identity returns the identity of the objects in the list.
func (o RevocationsList) Identity() elemental.Identity {

	return RevocationIdentity
}

// Copy returns a pointer to a copy the RevocationsList.
func (o RevocationsList) Copy() elemental.Identifiables {

	copy := append(RevocationsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the RevocationsList.
func (o RevocationsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(RevocationsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Revocation))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o RevocationsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o RevocationsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the RevocationsList converted to SparseRevocationsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o RevocationsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseRevocationsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseRevocation)
	}

	return out
}

// Version returns the version of the content.
func (o RevocationsList) Version() int {

	return 1
}

// Revocation represents the model of a revocation
type Revocation struct {
	// ID contains the ID of the revocation.
	ID string `json:"-" msgpack:"-" bson:"-" mapstructure:"-,omitempty"`

	// Contains the certificate expiration date. This will be used to clean up revoked
	// certificates that have expired.
	ExpirationDate time.Time `json:"expirationDate" msgpack:"expirationDate" bson:"expirationdate" mapstructure:"expirationDate,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Set time from when the certificate will be revoked.
	RevokeDate time.Time `json:"revokeDate" msgpack:"revokeDate" bson:"revokedate" mapstructure:"revokeDate,omitempty"`

	// SerialNumber of the revoked certificate.
	SerialNumber string `json:"serialNumber" msgpack:"serialNumber" bson:"serialnumber" mapstructure:"serialNumber,omitempty"`

	// Subject of the certificate related to the revocation.
	Subject string `json:"subject" msgpack:"subject" bson:"subject" mapstructure:"subject,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewRevocation returns a new *Revocation
func NewRevocation() *Revocation {

	return &Revocation{
		ModelVersion:  1,
		MigrationsLog: map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *Revocation) Identity() elemental.Identity {

	return RevocationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Revocation) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Revocation) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Revocation) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesRevocation{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.ExpirationDate = o.ExpirationDate
	s.MigrationsLog = o.MigrationsLog
	s.RevokeDate = o.RevokeDate
	s.SerialNumber = o.SerialNumber
	s.Subject = o.Subject
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Revocation) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesRevocation{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.ExpirationDate = s.ExpirationDate
	o.MigrationsLog = s.MigrationsLog
	o.RevokeDate = s.RevokeDate
	o.SerialNumber = s.SerialNumber
	o.Subject = s.Subject
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Revocation) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Revocation) BleveType() string {

	return "revocation"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Revocation) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Revocation) Doc() string {

	return `Used to revoke a certificate.`
}

func (o *Revocation) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *Revocation) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *Revocation) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetZHash returns the ZHash of the receiver.
func (o *Revocation) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *Revocation) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *Revocation) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *Revocation) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Revocation) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseRevocation{
			ID:             &o.ID,
			ExpirationDate: &o.ExpirationDate,
			MigrationsLog:  &o.MigrationsLog,
			RevokeDate:     &o.RevokeDate,
			SerialNumber:   &o.SerialNumber,
			Subject:        &o.Subject,
			ZHash:          &o.ZHash,
			Zone:           &o.Zone,
		}
	}

	sp := &SparseRevocation{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "expirationDate":
			sp.ExpirationDate = &(o.ExpirationDate)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "revokeDate":
			sp.RevokeDate = &(o.RevokeDate)
		case "serialNumber":
			sp.SerialNumber = &(o.SerialNumber)
		case "subject":
			sp.Subject = &(o.Subject)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseRevocation to the object.
func (o *Revocation) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseRevocation)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.ExpirationDate != nil {
		o.ExpirationDate = *so.ExpirationDate
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.RevokeDate != nil {
		o.RevokeDate = *so.RevokeDate
	}
	if so.SerialNumber != nil {
		o.SerialNumber = *so.SerialNumber
	}
	if so.Subject != nil {
		o.Subject = *so.Subject
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the Revocation.
func (o *Revocation) DeepCopy() *Revocation {

	if o == nil {
		return nil
	}

	out := &Revocation{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Revocation.
func (o *Revocation) DeepCopyInto(out *Revocation) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Revocation: %s", err))
	}

	*out = *target.(*Revocation)
}

// Validate valides the current information stored into the structure.
func (o *Revocation) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*Revocation) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := RevocationAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return RevocationLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Revocation) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RevocationAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Revocation) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "expirationDate":
		return o.ExpirationDate
	case "migrationsLog":
		return o.MigrationsLog
	case "revokeDate":
		return o.RevokeDate
	case "serialNumber":
		return o.SerialNumber
	case "subject":
		return o.Subject
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// RevocationAttributesMap represents the map of attribute for Revocation.
var RevocationAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID contains the ID of the revocation.`,
		Identifier:     true,
		Name:           "ID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ExpirationDate": {
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationDate",
		CreationOnly:   true,
		Description: `Contains the certificate expiration date. This will be used to clean up revoked
certificates that have expired.`,
		Exposed: true,
		Name:    "expirationDate",
		Stored:  true,
		Type:    "time",
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
	"RevokeDate": {
		AllowedChoices: []string{},
		ConvertedName:  "RevokeDate",
		Description:    `Set time from when the certificate will be revoked.`,
		Exposed:        true,
		Name:           "revokeDate",
		Stored:         true,
		Type:           "time",
	},
	"SerialNumber": {
		AllowedChoices: []string{},
		ConvertedName:  "SerialNumber",
		CreationOnly:   true,
		Description:    `SerialNumber of the revoked certificate.`,
		Exposed:        true,
		Name:           "serialNumber",
		Stored:         true,
		Type:           "string",
	},
	"Subject": {
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		CreationOnly:   true,
		Description:    `Subject of the certificate related to the revocation.`,
		Exposed:        true,
		Name:           "subject",
		Stored:         true,
		Type:           "string",
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

// RevocationLowerCaseAttributesMap represents the map of attribute for Revocation.
var RevocationLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID contains the ID of the revocation.`,
		Identifier:     true,
		Name:           "ID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"expirationdate": {
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationDate",
		CreationOnly:   true,
		Description: `Contains the certificate expiration date. This will be used to clean up revoked
certificates that have expired.`,
		Exposed: true,
		Name:    "expirationDate",
		Stored:  true,
		Type:    "time",
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
	"revokedate": {
		AllowedChoices: []string{},
		ConvertedName:  "RevokeDate",
		Description:    `Set time from when the certificate will be revoked.`,
		Exposed:        true,
		Name:           "revokeDate",
		Stored:         true,
		Type:           "time",
	},
	"serialnumber": {
		AllowedChoices: []string{},
		ConvertedName:  "SerialNumber",
		CreationOnly:   true,
		Description:    `SerialNumber of the revoked certificate.`,
		Exposed:        true,
		Name:           "serialNumber",
		Stored:         true,
		Type:           "string",
	},
	"subject": {
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		CreationOnly:   true,
		Description:    `Subject of the certificate related to the revocation.`,
		Exposed:        true,
		Name:           "subject",
		Stored:         true,
		Type:           "string",
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

// SparseRevocationsList represents a list of SparseRevocations
type SparseRevocationsList []*SparseRevocation

// Identity returns the identity of the objects in the list.
func (o SparseRevocationsList) Identity() elemental.Identity {

	return RevocationIdentity
}

// Copy returns a pointer to a copy the SparseRevocationsList.
func (o SparseRevocationsList) Copy() elemental.Identifiables {

	copy := append(SparseRevocationsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseRevocationsList.
func (o SparseRevocationsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseRevocationsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseRevocation))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseRevocationsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseRevocationsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseRevocationsList converted to RevocationsList.
func (o SparseRevocationsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseRevocationsList) Version() int {

	return 1
}

// SparseRevocation represents the sparse version of a revocation.
type SparseRevocation struct {
	// ID contains the ID of the revocation.
	ID *string `json:"-" msgpack:"-" bson:"-" mapstructure:"-,omitempty"`

	// Contains the certificate expiration date. This will be used to clean up revoked
	// certificates that have expired.
	ExpirationDate *time.Time `json:"expirationDate,omitempty" msgpack:"expirationDate,omitempty" bson:"expirationdate,omitempty" mapstructure:"expirationDate,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Set time from when the certificate will be revoked.
	RevokeDate *time.Time `json:"revokeDate,omitempty" msgpack:"revokeDate,omitempty" bson:"revokedate,omitempty" mapstructure:"revokeDate,omitempty"`

	// SerialNumber of the revoked certificate.
	SerialNumber *string `json:"serialNumber,omitempty" msgpack:"serialNumber,omitempty" bson:"serialnumber,omitempty" mapstructure:"serialNumber,omitempty"`

	// Subject of the certificate related to the revocation.
	Subject *string `json:"subject,omitempty" msgpack:"subject,omitempty" bson:"subject,omitempty" mapstructure:"subject,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseRevocation returns a new  SparseRevocation.
func NewSparseRevocation() *SparseRevocation {
	return &SparseRevocation{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseRevocation) Identity() elemental.Identity {

	return RevocationIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseRevocation) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseRevocation) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseRevocation) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseRevocation{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.ExpirationDate != nil {
		s.ExpirationDate = o.ExpirationDate
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.RevokeDate != nil {
		s.RevokeDate = o.RevokeDate
	}
	if o.SerialNumber != nil {
		s.SerialNumber = o.SerialNumber
	}
	if o.Subject != nil {
		s.Subject = o.Subject
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
func (o *SparseRevocation) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseRevocation{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.ExpirationDate != nil {
		o.ExpirationDate = s.ExpirationDate
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.RevokeDate != nil {
		o.RevokeDate = s.RevokeDate
	}
	if s.SerialNumber != nil {
		o.SerialNumber = s.SerialNumber
	}
	if s.Subject != nil {
		o.Subject = s.Subject
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
func (o *SparseRevocation) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseRevocation) ToPlain() elemental.PlainIdentifiable {

	out := NewRevocation()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.ExpirationDate != nil {
		out.ExpirationDate = *o.ExpirationDate
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.RevokeDate != nil {
		out.RevokeDate = *o.RevokeDate
	}
	if o.SerialNumber != nil {
		out.SerialNumber = *o.SerialNumber
	}
	if o.Subject != nil {
		out.Subject = *o.Subject
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseRevocation) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseRevocation) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseRevocation) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseRevocation) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseRevocation) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseRevocation) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseRevocation.
func (o *SparseRevocation) DeepCopy() *SparseRevocation {

	if o == nil {
		return nil
	}

	out := &SparseRevocation{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseRevocation.
func (o *SparseRevocation) DeepCopyInto(out *SparseRevocation) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseRevocation: %s", err))
	}

	*out = *target.(*SparseRevocation)
}

type mongoAttributesRevocation struct {
	ID             bson.ObjectId     `bson:"_id,omitempty"`
	ExpirationDate time.Time         `bson:"expirationdate"`
	MigrationsLog  map[string]string `bson:"migrationslog,omitempty"`
	RevokeDate     time.Time         `bson:"revokedate"`
	SerialNumber   string            `bson:"serialnumber"`
	Subject        string            `bson:"subject"`
	ZHash          int               `bson:"zhash"`
	Zone           int               `bson:"zone"`
}
type mongoAttributesSparseRevocation struct {
	ID             bson.ObjectId      `bson:"_id,omitempty"`
	ExpirationDate *time.Time         `bson:"expirationdate,omitempty"`
	MigrationsLog  *map[string]string `bson:"migrationslog,omitempty"`
	RevokeDate     *time.Time         `bson:"revokedate,omitempty"`
	SerialNumber   *string            `bson:"serialnumber,omitempty"`
	Subject        *string            `bson:"subject,omitempty"`
	ZHash          *int               `bson:"zhash,omitempty"`
	Zone           *int               `bson:"zone,omitempty"`
}
