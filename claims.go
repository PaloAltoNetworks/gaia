package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ClaimsIdentity represents the Identity of the object.
var ClaimsIdentity = elemental.Identity{
	Name:     "claims",
	Category: "claims",
	Package:  "guy",
	Private:  false,
}

// ClaimsList represents a list of Claims
type ClaimsList []*Claims

// Identity returns the identity of the objects in the list.
func (o ClaimsList) Identity() elemental.Identity {

	return ClaimsIdentity
}

// Copy returns a pointer to a copy the ClaimsList.
func (o ClaimsList) Copy() elemental.Identifiables {

	copy := append(ClaimsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ClaimsList.
func (o ClaimsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ClaimsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Claims))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ClaimsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ClaimsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ClaimsList converted to SparseClaimsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ClaimsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseClaimsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseClaims)
	}

	return out
}

// Version returns the version of the content.
func (o ClaimsList) Version() int {

	return 1
}

// Claims represents the model of a claims
type Claims struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Contains the raw JSON web token (JWT) claims.
	Content map[string]string `json:"content,omitempty" msgpack:"content,omitempty" bson:"content" mapstructure:"content,omitempty"`

	// Contains the date of the first appearance of the claims.
	FirstSeen time.Time `json:"-" msgpack:"-" bson:"firstseen" mapstructure:"-,omitempty"`

	// XXH64 hash of the claims content. It will be used as ID. To compute a correct
	// hash,
	// you must first clob `content` as an string array in the form `key=value`, sort
	// it
	// then apply the XXH64 function.
	Hash string `json:"hash" msgpack:"hash" bson:"hash" mapstructure:"hash,omitempty"`

	// Contains the date of the last appearance of the claims.
	LastSeen time.Time `json:"-" msgpack:"-" bson:"lastseen" mapstructure:"-,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewClaims returns a new *Claims
func NewClaims() *Claims {

	return &Claims{
		ModelVersion:  1,
		Content:       map[string]string{},
		MigrationsLog: map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *Claims) Identity() elemental.Identity {

	return ClaimsIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Claims) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Claims) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Claims) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesClaims{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Content = o.Content
	s.FirstSeen = o.FirstSeen
	s.Hash = o.Hash
	s.LastSeen = o.LastSeen
	s.MigrationsLog = o.MigrationsLog
	s.Namespace = o.Namespace
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Claims) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesClaims{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Content = s.Content
	o.FirstSeen = s.FirstSeen
	o.Hash = s.Hash
	o.LastSeen = s.LastSeen
	o.MigrationsLog = s.MigrationsLog
	o.Namespace = s.Namespace
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Claims) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Claims) BleveType() string {

	return "claims"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Claims) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Claims) Doc() string {

	return `Represents the claims in the token used to access a service.`
}

func (o *Claims) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *Claims) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *Claims) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *Claims) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *Claims) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetZHash returns the ZHash of the receiver.
func (o *Claims) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *Claims) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *Claims) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *Claims) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Claims) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseClaims{
			ID:            &o.ID,
			Content:       &o.Content,
			FirstSeen:     &o.FirstSeen,
			Hash:          &o.Hash,
			LastSeen:      &o.LastSeen,
			MigrationsLog: &o.MigrationsLog,
			Namespace:     &o.Namespace,
			ZHash:         &o.ZHash,
			Zone:          &o.Zone,
		}
	}

	sp := &SparseClaims{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "content":
			sp.Content = &(o.Content)
		case "firstSeen":
			sp.FirstSeen = &(o.FirstSeen)
		case "hash":
			sp.Hash = &(o.Hash)
		case "lastSeen":
			sp.LastSeen = &(o.LastSeen)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseClaims to the object.
func (o *Claims) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseClaims)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Content != nil {
		o.Content = *so.Content
	}
	if so.FirstSeen != nil {
		o.FirstSeen = *so.FirstSeen
	}
	if so.Hash != nil {
		o.Hash = *so.Hash
	}
	if so.LastSeen != nil {
		o.LastSeen = *so.LastSeen
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the Claims.
func (o *Claims) DeepCopy() *Claims {

	if o == nil {
		return nil
	}

	out := &Claims{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Claims.
func (o *Claims) DeepCopyInto(out *Claims) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Claims: %s", err))
	}

	*out = *target.(*Claims)
}

// Validate valides the current information stored into the structure.
func (o *Claims) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("hash", o.Hash); err != nil {
		requiredErrors = requiredErrors.Append(err)
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
func (*Claims) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ClaimsAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ClaimsLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Claims) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ClaimsAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Claims) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "content":
		return o.Content
	case "firstSeen":
		return o.FirstSeen
	case "hash":
		return o.Hash
	case "lastSeen":
		return o.LastSeen
	case "migrationsLog":
		return o.MigrationsLog
	case "namespace":
		return o.Namespace
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// ClaimsAttributesMap represents the map of attribute for Claims.
var ClaimsAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Content": {
		AllowedChoices: []string{},
		ConvertedName:  "Content",
		CreationOnly:   true,
		Description:    `Contains the raw JSON web token (JWT) claims.`,
		Exposed:        true,
		Name:           "content",
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"FirstSeen": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "FirstSeen",
		Description:    `Contains the date of the first appearance of the claims.`,
		Name:           "firstSeen",
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"Hash": {
		AllowedChoices: []string{},
		ConvertedName:  "Hash",
		Description: `XXH64 hash of the claims content. It will be used as ID. To compute a correct
hash,
you must first clob ` + "`" + `content` + "`" + ` as an string array in the form ` + "`" + `key=value` + "`" + `, sort
it
then apply the XXH64 function.`,
		Exposed:  true,
		Name:     "hash",
		Required: true,
		Stored:   true,
		Type:     "string",
	},
	"LastSeen": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LastSeen",
		Description:    `Contains the date of the last appearance of the claims.`,
		Name:           "lastSeen",
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
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
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
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

// ClaimsLowerCaseAttributesMap represents the map of attribute for Claims.
var ClaimsLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"content": {
		AllowedChoices: []string{},
		ConvertedName:  "Content",
		CreationOnly:   true,
		Description:    `Contains the raw JSON web token (JWT) claims.`,
		Exposed:        true,
		Name:           "content",
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"firstseen": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "FirstSeen",
		Description:    `Contains the date of the first appearance of the claims.`,
		Name:           "firstSeen",
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"hash": {
		AllowedChoices: []string{},
		ConvertedName:  "Hash",
		Description: `XXH64 hash of the claims content. It will be used as ID. To compute a correct
hash,
you must first clob ` + "`" + `content` + "`" + ` as an string array in the form ` + "`" + `key=value` + "`" + `, sort
it
then apply the XXH64 function.`,
		Exposed:  true,
		Name:     "hash",
		Required: true,
		Stored:   true,
		Type:     "string",
	},
	"lastseen": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LastSeen",
		Description:    `Contains the date of the last appearance of the claims.`,
		Name:           "lastSeen",
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
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
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
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

// SparseClaimsList represents a list of SparseClaims
type SparseClaimsList []*SparseClaims

// Identity returns the identity of the objects in the list.
func (o SparseClaimsList) Identity() elemental.Identity {

	return ClaimsIdentity
}

// Copy returns a pointer to a copy the SparseClaimsList.
func (o SparseClaimsList) Copy() elemental.Identifiables {

	copy := append(SparseClaimsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseClaimsList.
func (o SparseClaimsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseClaimsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseClaims))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseClaimsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseClaimsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseClaimsList converted to ClaimsList.
func (o SparseClaimsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseClaimsList) Version() int {

	return 1
}

// SparseClaims represents the sparse version of a claims.
type SparseClaims struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Contains the raw JSON web token (JWT) claims.
	Content *map[string]string `json:"content,omitempty" msgpack:"content,omitempty" bson:"content,omitempty" mapstructure:"content,omitempty"`

	// Contains the date of the first appearance of the claims.
	FirstSeen *time.Time `json:"-" msgpack:"-" bson:"firstseen,omitempty" mapstructure:"-,omitempty"`

	// XXH64 hash of the claims content. It will be used as ID. To compute a correct
	// hash,
	// you must first clob `content` as an string array in the form `key=value`, sort
	// it
	// then apply the XXH64 function.
	Hash *string `json:"hash,omitempty" msgpack:"hash,omitempty" bson:"hash,omitempty" mapstructure:"hash,omitempty"`

	// Contains the date of the last appearance of the claims.
	LastSeen *time.Time `json:"-" msgpack:"-" bson:"lastseen,omitempty" mapstructure:"-,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseClaims returns a new  SparseClaims.
func NewSparseClaims() *SparseClaims {
	return &SparseClaims{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseClaims) Identity() elemental.Identity {

	return ClaimsIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseClaims) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseClaims) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseClaims) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseClaims{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Content != nil {
		s.Content = o.Content
	}
	if o.FirstSeen != nil {
		s.FirstSeen = o.FirstSeen
	}
	if o.Hash != nil {
		s.Hash = o.Hash
	}
	if o.LastSeen != nil {
		s.LastSeen = o.LastSeen
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
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
func (o *SparseClaims) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseClaims{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Content != nil {
		o.Content = s.Content
	}
	if s.FirstSeen != nil {
		o.FirstSeen = s.FirstSeen
	}
	if s.Hash != nil {
		o.Hash = s.Hash
	}
	if s.LastSeen != nil {
		o.LastSeen = s.LastSeen
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
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
func (o *SparseClaims) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseClaims) ToPlain() elemental.PlainIdentifiable {

	out := NewClaims()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Content != nil {
		out.Content = *o.Content
	}
	if o.FirstSeen != nil {
		out.FirstSeen = *o.FirstSeen
	}
	if o.Hash != nil {
		out.Hash = *o.Hash
	}
	if o.LastSeen != nil {
		out.LastSeen = *o.LastSeen
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
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
func (o *SparseClaims) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseClaims) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseClaims) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseClaims) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseClaims) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseClaims) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseClaims) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseClaims) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseClaims.
func (o *SparseClaims) DeepCopy() *SparseClaims {

	if o == nil {
		return nil
	}

	out := &SparseClaims{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseClaims.
func (o *SparseClaims) DeepCopyInto(out *SparseClaims) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseClaims: %s", err))
	}

	*out = *target.(*SparseClaims)
}

type mongoAttributesClaims struct {
	ID            bson.ObjectId     `bson:"_id,omitempty"`
	Content       map[string]string `bson:"content,omitempty"`
	FirstSeen     time.Time         `bson:"firstseen,omitempty"`
	Hash          string            `bson:"hash"`
	LastSeen      time.Time         `bson:"lastseen"`
	MigrationsLog map[string]string `bson:"migrationslog,omitempty"`
	Namespace     string            `bson:"namespace"`
	ZHash         int               `bson:"zhash"`
	Zone          int               `bson:"zone"`
}
type mongoAttributesSparseClaims struct {
	ID            bson.ObjectId      `bson:"_id,omitempty"`
	Content       *map[string]string `bson:"content,omitempty"`
	FirstSeen     *time.Time         `bson:"firstseen,omitempty"`
	Hash          *string            `bson:"hash,omitempty"`
	LastSeen      *time.Time         `bson:"lastseen,omitempty"`
	MigrationsLog *map[string]string `bson:"migrationslog,omitempty"`
	Namespace     *string            `bson:"namespace,omitempty"`
	ZHash         *int               `bson:"zhash,omitempty"`
	Zone          *int               `bson:"zone,omitempty"`
}
