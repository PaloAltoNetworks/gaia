package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AppCredentialIdentity represents the Identity of the object.
var AppCredentialIdentity = elemental.Identity{
	Name:     "appcredential",
	Category: "appcredentials",
	Package:  "cactuar",
	Private:  false,
}

// AppCredentialsList represents a list of AppCredentials
type AppCredentialsList []*AppCredential

// Identity returns the identity of the objects in the list.
func (o AppCredentialsList) Identity() elemental.Identity {

	return AppCredentialIdentity
}

// Copy returns a pointer to a copy the AppCredentialsList.
func (o AppCredentialsList) Copy() elemental.Identifiables {

	copy := append(AppCredentialsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AppCredentialsList.
func (o AppCredentialsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AppCredentialsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AppCredential))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AppCredentialsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AppCredentialsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the AppCredentialsList converted to SparseAppCredentialsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AppCredentialsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAppCredentialsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAppCredential)
	}

	return out
}

// Version returns the version of the content.
func (o AppCredentialsList) Version() int {

	return 1
}

// AppCredential represents the model of a appcredential
type AppCredential struct {
	// Contains a PEM-encoded certificate signing request (CSR). It can
	// only be set during a renew.
	//
	// - The CN **MUST** be `app:credential:<appcred-id>:<appcred-name>`
	// - The O **MUST** be the namespace of the app credential
	//
	// If you send anything else, the signing request will be rejected.
	CSR string `json:"CSR" msgpack:"CSR" bson:"-" mapstructure:"CSR,omitempty"`

	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// If set, the app credential will only be valid if the request comes from one
	// the declared subnets.
	AuthorizedSubnets []string `json:"authorizedSubnets" msgpack:"authorizedSubnets" bson:"authorizedsubnets" mapstructure:"authorizedSubnets,omitempty"`

	// The string representation of the certificate used by the app credential.
	Certificate string `json:"certificate" msgpack:"certificate" bson:"certificate" mapstructure:"certificate,omitempty"`

	// Link to the certificate created for this app credential.
	CertificateSN string `json:"-" msgpack:"-" bson:"certificatesn" mapstructure:"-,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// The app credential data.
	Credentials *Credential `json:"credentials" msgpack:"credentials" bson:"-" mapstructure:"credentials,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Defines if the property is disabled.
	Disabled bool `json:"disabled" msgpack:"disabled" bson:"disabled" mapstructure:"disabled,omitempty"`

	// The email address that will receive a copy of the app credential.
	Email string `json:"email" msgpack:"email" bson:"email" mapstructure:"email,omitempty"`

	// If set, this will limit the maximum validity of the token issued from this app
	// credential. This information will be embedded into the delivered certificate and
	// cannot be changed once set. In order to change it, you need to renew the
	// certificate.
	MaxIssuedTokenValidity string `json:"maxIssuedTokenValidity" msgpack:"maxIssuedTokenValidity" bson:"maxissuedtokenvalidity" mapstructure:"maxIssuedTokenValidity,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Contains the ID of the parent app credential if this is a derived app
	// credential.
	ParentIDs []string `json:"parentIDs" msgpack:"parentIDs" bson:"parentids" mapstructure:"parentIDs,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// List of roles to give the app credential.
	Roles []string `json:"roles" msgpack:"roles" bson:"roles" mapstructure:"roles,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAppCredential returns a new *AppCredential
func NewAppCredential() *AppCredential {

	return &AppCredential{
		ModelVersion:      1,
		AssociatedTags:    []string{},
		Annotations:       map[string][]string{},
		AuthorizedSubnets: []string{},
		MigrationsLog:     map[string]string{},
		NormalizedTags:    []string{},
		ParentIDs:         []string{},
		Roles:             []string{},
		Metadata:          []string{},
	}
}

// Identity returns the Identity of the object.
func (o *AppCredential) Identity() elemental.Identity {

	return AppCredentialIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AppCredential) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AppCredential) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AppCredential) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesAppCredential{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.AuthorizedSubnets = o.AuthorizedSubnets
	s.Certificate = o.Certificate
	s.CertificateSN = o.CertificateSN
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.Description = o.Description
	s.Disabled = o.Disabled
	s.Email = o.Email
	s.MaxIssuedTokenValidity = o.MaxIssuedTokenValidity
	s.Metadata = o.Metadata
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.ParentIDs = o.ParentIDs
	s.Protected = o.Protected
	s.Roles = o.Roles
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AppCredential) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesAppCredential{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.AuthorizedSubnets = s.AuthorizedSubnets
	o.Certificate = s.Certificate
	o.CertificateSN = s.CertificateSN
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.Description = s.Description
	o.Disabled = s.Disabled
	o.Email = s.Email
	o.MaxIssuedTokenValidity = s.MaxIssuedTokenValidity
	o.Metadata = s.Metadata
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.ParentIDs = s.ParentIDs
	o.Protected = s.Protected
	o.Roles = s.Roles
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *AppCredential) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *AppCredential) BleveType() string {

	return "appcredential"
}

// DefaultOrder returns the list of default ordering fields.
func (o *AppCredential) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *AppCredential) Doc() string {

	return `Create an app credential.`
}

func (o *AppCredential) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *AppCredential) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *AppCredential) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *AppCredential) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *AppCredential) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *AppCredential) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *AppCredential) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *AppCredential) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *AppCredential) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *AppCredential) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *AppCredential) SetDescription(description string) {

	o.Description = description
}

// GetDisabled returns the Disabled of the receiver.
func (o *AppCredential) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the given value.
func (o *AppCredential) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetMetadata returns the Metadata of the receiver.
func (o *AppCredential) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *AppCredential) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *AppCredential) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *AppCredential) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *AppCredential) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *AppCredential) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *AppCredential) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *AppCredential) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *AppCredential) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *AppCredential) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *AppCredential) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *AppCredential) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *AppCredential) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *AppCredential) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *AppCredential) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *AppCredential) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *AppCredential) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *AppCredential) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *AppCredential) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *AppCredential) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *AppCredential) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAppCredential{
			CSR:                    &o.CSR,
			ID:                     &o.ID,
			Annotations:            &o.Annotations,
			AssociatedTags:         &o.AssociatedTags,
			AuthorizedSubnets:      &o.AuthorizedSubnets,
			Certificate:            &o.Certificate,
			CertificateSN:          &o.CertificateSN,
			CreateIdempotencyKey:   &o.CreateIdempotencyKey,
			CreateTime:             &o.CreateTime,
			Credentials:            o.Credentials,
			Description:            &o.Description,
			Disabled:               &o.Disabled,
			Email:                  &o.Email,
			MaxIssuedTokenValidity: &o.MaxIssuedTokenValidity,
			Metadata:               &o.Metadata,
			MigrationsLog:          &o.MigrationsLog,
			Name:                   &o.Name,
			Namespace:              &o.Namespace,
			NormalizedTags:         &o.NormalizedTags,
			ParentIDs:              &o.ParentIDs,
			Protected:              &o.Protected,
			Roles:                  &o.Roles,
			UpdateIdempotencyKey:   &o.UpdateIdempotencyKey,
			UpdateTime:             &o.UpdateTime,
			ZHash:                  &o.ZHash,
			Zone:                   &o.Zone,
		}
	}

	sp := &SparseAppCredential{}
	for _, f := range fields {
		switch f {
		case "CSR":
			sp.CSR = &(o.CSR)
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "authorizedSubnets":
			sp.AuthorizedSubnets = &(o.AuthorizedSubnets)
		case "certificate":
			sp.Certificate = &(o.Certificate)
		case "certificateSN":
			sp.CertificateSN = &(o.CertificateSN)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "credentials":
			sp.Credentials = o.Credentials
		case "description":
			sp.Description = &(o.Description)
		case "disabled":
			sp.Disabled = &(o.Disabled)
		case "email":
			sp.Email = &(o.Email)
		case "maxIssuedTokenValidity":
			sp.MaxIssuedTokenValidity = &(o.MaxIssuedTokenValidity)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "parentIDs":
			sp.ParentIDs = &(o.ParentIDs)
		case "protected":
			sp.Protected = &(o.Protected)
		case "roles":
			sp.Roles = &(o.Roles)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseAppCredential to the object.
func (o *AppCredential) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAppCredential)
	if so.CSR != nil {
		o.CSR = *so.CSR
	}
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.AuthorizedSubnets != nil {
		o.AuthorizedSubnets = *so.AuthorizedSubnets
	}
	if so.Certificate != nil {
		o.Certificate = *so.Certificate
	}
	if so.CertificateSN != nil {
		o.CertificateSN = *so.CertificateSN
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Credentials != nil {
		o.Credentials = so.Credentials
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Disabled != nil {
		o.Disabled = *so.Disabled
	}
	if so.Email != nil {
		o.Email = *so.Email
	}
	if so.MaxIssuedTokenValidity != nil {
		o.MaxIssuedTokenValidity = *so.MaxIssuedTokenValidity
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.ParentIDs != nil {
		o.ParentIDs = *so.ParentIDs
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Roles != nil {
		o.Roles = *so.Roles
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the AppCredential.
func (o *AppCredential) DeepCopy() *AppCredential {

	if o == nil {
		return nil
	}

	out := &AppCredential{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AppCredential.
func (o *AppCredential) DeepCopyInto(out *AppCredential) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AppCredential: %s", err))
	}

	*out = *target.(*AppCredential)
}

// Validate valides the current information stored into the structure.
func (o *AppCredential) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateOptionalCIDRList("authorizedSubnets", o.AuthorizedSubnets); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateOptionalTimeDuration("maxIssuedTokenValidity", o.MaxIssuedTokenValidity); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateMetadata("metadata", o.Metadata); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredExternal("roles", o.Roles); err != nil {
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
func (*AppCredential) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AppCredentialAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AppCredentialLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AppCredential) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AppCredentialAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *AppCredential) ValueForAttribute(name string) interface{} {

	switch name {
	case "CSR":
		return o.CSR
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "authorizedSubnets":
		return o.AuthorizedSubnets
	case "certificate":
		return o.Certificate
	case "certificateSN":
		return o.CertificateSN
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "credentials":
		return o.Credentials
	case "description":
		return o.Description
	case "disabled":
		return o.Disabled
	case "email":
		return o.Email
	case "maxIssuedTokenValidity":
		return o.MaxIssuedTokenValidity
	case "metadata":
		return o.Metadata
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "parentIDs":
		return o.ParentIDs
	case "protected":
		return o.Protected
	case "roles":
		return o.Roles
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// AppCredentialAttributesMap represents the map of attribute for AppCredential.
var AppCredentialAttributesMap = map[string]elemental.AttributeSpecification{
	"CSR": {
		AllowedChoices: []string{},
		ConvertedName:  "CSR",
		Description: `Contains a PEM-encoded certificate signing request (CSR). It can
only be set during a renew.

- The CN **MUST** be ` + "`" + `app:credential:<appcred-id>:<appcred-name>` + "`" + `
- The O **MUST** be the namespace of the app credential

If you send anything else, the signing request will be rejected.`,
		Exposed: true,
		Name:    "CSR",
		Type:    "string",
	},
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
	"Annotations": {
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"AssociatedTags": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"AuthorizedSubnets": {
		AllowedChoices: []string{},
		ConvertedName:  "AuthorizedSubnets",
		Description: `If set, the app credential will only be valid if the request comes from one
the declared subnets.`,
		Exposed: true,
		Name:    "authorizedSubnets",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"Certificate": {
		AllowedChoices: []string{},
		ConvertedName:  "Certificate",
		Description:    `The string representation of the certificate used by the app credential.`,
		Exposed:        true,
		Name:           "certificate",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"CertificateSN": {
		AllowedChoices: []string{},
		ConvertedName:  "CertificateSN",
		Description:    `Link to the certificate created for this app credential.`,
		Name:           "certificateSN",
		Stored:         true,
		Type:           "string",
	},
	"CreateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Credentials": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Credentials",
		Description:    `The app credential data.`,
		Exposed:        true,
		Name:           "credentials",
		Orderable:      true,
		ReadOnly:       true,
		SubType:        "credential",
		Transient:      true,
		Type:           "ref",
	},
	"Description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Disabled": {
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Defines if the property is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Email": {
		AllowedChoices: []string{},
		ConvertedName:  "Email",
		Description:    `The email address that will receive a copy of the app credential.`,
		Exposed:        true,
		Name:           "email",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"MaxIssuedTokenValidity": {
		AllowedChoices: []string{},
		ConvertedName:  "MaxIssuedTokenValidity",
		Description: `If set, this will limit the maximum validity of the token issued from this app
credential. This information will be embedded into the delivered certificate and
cannot be changed once set. In order to change it, you need to renew the
certificate.`,
		Exposed: true,
		Name:    "maxIssuedTokenValidity",
		Stored:  true,
		Type:    "string",
	},
	"Metadata": {
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
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
	"Name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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
	"NormalizedTags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"ParentIDs": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentIDs",
		Description: `Contains the ID of the parent app credential if this is a derived app
credential.`,
		Exposed:  true,
		Name:     "parentIDs",
		ReadOnly: true,
		Stored:   true,
		SubType:  "string",
		Type:     "list",
	},
	"Protected": {
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Roles": {
		AllowedChoices: []string{},
		ConvertedName:  "Roles",
		Description:    `List of roles to give the app credential.`,
		Exposed:        true,
		Name:           "roles",
		Required:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"UpdateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"UpdateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
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

// AppCredentialLowerCaseAttributesMap represents the map of attribute for AppCredential.
var AppCredentialLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"csr": {
		AllowedChoices: []string{},
		ConvertedName:  "CSR",
		Description: `Contains a PEM-encoded certificate signing request (CSR). It can
only be set during a renew.

- The CN **MUST** be ` + "`" + `app:credential:<appcred-id>:<appcred-name>` + "`" + `
- The O **MUST** be the namespace of the app credential

If you send anything else, the signing request will be rejected.`,
		Exposed: true,
		Name:    "CSR",
		Type:    "string",
	},
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
	"annotations": {
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"associatedtags": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"authorizedsubnets": {
		AllowedChoices: []string{},
		ConvertedName:  "AuthorizedSubnets",
		Description: `If set, the app credential will only be valid if the request comes from one
the declared subnets.`,
		Exposed: true,
		Name:    "authorizedSubnets",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"certificate": {
		AllowedChoices: []string{},
		ConvertedName:  "Certificate",
		Description:    `The string representation of the certificate used by the app credential.`,
		Exposed:        true,
		Name:           "certificate",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"certificatesn": {
		AllowedChoices: []string{},
		ConvertedName:  "CertificateSN",
		Description:    `Link to the certificate created for this app credential.`,
		Name:           "certificateSN",
		Stored:         true,
		Type:           "string",
	},
	"createidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"credentials": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Credentials",
		Description:    `The app credential data.`,
		Exposed:        true,
		Name:           "credentials",
		Orderable:      true,
		ReadOnly:       true,
		SubType:        "credential",
		Transient:      true,
		Type:           "ref",
	},
	"description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"disabled": {
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Defines if the property is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"email": {
		AllowedChoices: []string{},
		ConvertedName:  "Email",
		Description:    `The email address that will receive a copy of the app credential.`,
		Exposed:        true,
		Name:           "email",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"maxissuedtokenvalidity": {
		AllowedChoices: []string{},
		ConvertedName:  "MaxIssuedTokenValidity",
		Description: `If set, this will limit the maximum validity of the token issued from this app
credential. This information will be embedded into the delivered certificate and
cannot be changed once set. In order to change it, you need to renew the
certificate.`,
		Exposed: true,
		Name:    "maxIssuedTokenValidity",
		Stored:  true,
		Type:    "string",
	},
	"metadata": {
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
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
	"name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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
	"normalizedtags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"parentids": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentIDs",
		Description: `Contains the ID of the parent app credential if this is a derived app
credential.`,
		Exposed:  true,
		Name:     "parentIDs",
		ReadOnly: true,
		Stored:   true,
		SubType:  "string",
		Type:     "list",
	},
	"protected": {
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"roles": {
		AllowedChoices: []string{},
		ConvertedName:  "Roles",
		Description:    `List of roles to give the app credential.`,
		Exposed:        true,
		Name:           "roles",
		Required:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"updateidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"updatetime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
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

// SparseAppCredentialsList represents a list of SparseAppCredentials
type SparseAppCredentialsList []*SparseAppCredential

// Identity returns the identity of the objects in the list.
func (o SparseAppCredentialsList) Identity() elemental.Identity {

	return AppCredentialIdentity
}

// Copy returns a pointer to a copy the SparseAppCredentialsList.
func (o SparseAppCredentialsList) Copy() elemental.Identifiables {

	copy := append(SparseAppCredentialsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAppCredentialsList.
func (o SparseAppCredentialsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAppCredentialsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAppCredential))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAppCredentialsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAppCredentialsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseAppCredentialsList converted to AppCredentialsList.
func (o SparseAppCredentialsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAppCredentialsList) Version() int {

	return 1
}

// SparseAppCredential represents the sparse version of a appcredential.
type SparseAppCredential struct {
	// Contains a PEM-encoded certificate signing request (CSR). It can
	// only be set during a renew.
	//
	// - The CN **MUST** be `app:credential:<appcred-id>:<appcred-name>`
	// - The O **MUST** be the namespace of the app credential
	//
	// If you send anything else, the signing request will be rejected.
	CSR *string `json:"CSR,omitempty" msgpack:"CSR,omitempty" bson:"-" mapstructure:"CSR,omitempty"`

	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// If set, the app credential will only be valid if the request comes from one
	// the declared subnets.
	AuthorizedSubnets *[]string `json:"authorizedSubnets,omitempty" msgpack:"authorizedSubnets,omitempty" bson:"authorizedsubnets,omitempty" mapstructure:"authorizedSubnets,omitempty"`

	// The string representation of the certificate used by the app credential.
	Certificate *string `json:"certificate,omitempty" msgpack:"certificate,omitempty" bson:"certificate,omitempty" mapstructure:"certificate,omitempty"`

	// Link to the certificate created for this app credential.
	CertificateSN *string `json:"-" msgpack:"-" bson:"certificatesn,omitempty" mapstructure:"-,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// The app credential data.
	Credentials *Credential `json:"credentials,omitempty" msgpack:"credentials,omitempty" bson:"-" mapstructure:"credentials,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Defines if the property is disabled.
	Disabled *bool `json:"disabled,omitempty" msgpack:"disabled,omitempty" bson:"disabled,omitempty" mapstructure:"disabled,omitempty"`

	// The email address that will receive a copy of the app credential.
	Email *string `json:"email,omitempty" msgpack:"email,omitempty" bson:"email,omitempty" mapstructure:"email,omitempty"`

	// If set, this will limit the maximum validity of the token issued from this app
	// credential. This information will be embedded into the delivered certificate and
	// cannot be changed once set. In order to change it, you need to renew the
	// certificate.
	MaxIssuedTokenValidity *string `json:"maxIssuedTokenValidity,omitempty" msgpack:"maxIssuedTokenValidity,omitempty" bson:"maxissuedtokenvalidity,omitempty" mapstructure:"maxIssuedTokenValidity,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Contains the ID of the parent app credential if this is a derived app
	// credential.
	ParentIDs *[]string `json:"parentIDs,omitempty" msgpack:"parentIDs,omitempty" bson:"parentids,omitempty" mapstructure:"parentIDs,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// List of roles to give the app credential.
	Roles *[]string `json:"roles,omitempty" msgpack:"roles,omitempty" bson:"roles,omitempty" mapstructure:"roles,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAppCredential returns a new  SparseAppCredential.
func NewSparseAppCredential() *SparseAppCredential {
	return &SparseAppCredential{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAppCredential) Identity() elemental.Identity {

	return AppCredentialIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAppCredential) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAppCredential) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAppCredential) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseAppCredential{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.AuthorizedSubnets != nil {
		s.AuthorizedSubnets = o.AuthorizedSubnets
	}
	if o.Certificate != nil {
		s.Certificate = o.Certificate
	}
	if o.CertificateSN != nil {
		s.CertificateSN = o.CertificateSN
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.Disabled != nil {
		s.Disabled = o.Disabled
	}
	if o.Email != nil {
		s.Email = o.Email
	}
	if o.MaxIssuedTokenValidity != nil {
		s.MaxIssuedTokenValidity = o.MaxIssuedTokenValidity
	}
	if o.Metadata != nil {
		s.Metadata = o.Metadata
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.ParentIDs != nil {
		s.ParentIDs = o.ParentIDs
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.Roles != nil {
		s.Roles = o.Roles
	}
	if o.UpdateIdempotencyKey != nil {
		s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		s.UpdateTime = o.UpdateTime
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
func (o *SparseAppCredential) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseAppCredential{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.AuthorizedSubnets != nil {
		o.AuthorizedSubnets = s.AuthorizedSubnets
	}
	if s.Certificate != nil {
		o.Certificate = s.Certificate
	}
	if s.CertificateSN != nil {
		o.CertificateSN = s.CertificateSN
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Disabled != nil {
		o.Disabled = s.Disabled
	}
	if s.Email != nil {
		o.Email = s.Email
	}
	if s.MaxIssuedTokenValidity != nil {
		o.MaxIssuedTokenValidity = s.MaxIssuedTokenValidity
	}
	if s.Metadata != nil {
		o.Metadata = s.Metadata
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.ParentIDs != nil {
		o.ParentIDs = s.ParentIDs
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.Roles != nil {
		o.Roles = s.Roles
	}
	if s.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	}
	if s.UpdateTime != nil {
		o.UpdateTime = s.UpdateTime
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
func (o *SparseAppCredential) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAppCredential) ToPlain() elemental.PlainIdentifiable {

	out := NewAppCredential()
	if o.CSR != nil {
		out.CSR = *o.CSR
	}
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.AuthorizedSubnets != nil {
		out.AuthorizedSubnets = *o.AuthorizedSubnets
	}
	if o.Certificate != nil {
		out.Certificate = *o.Certificate
	}
	if o.CertificateSN != nil {
		out.CertificateSN = *o.CertificateSN
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Credentials != nil {
		out.Credentials = o.Credentials
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Disabled != nil {
		out.Disabled = *o.Disabled
	}
	if o.Email != nil {
		out.Email = *o.Email
	}
	if o.MaxIssuedTokenValidity != nil {
		out.MaxIssuedTokenValidity = *o.MaxIssuedTokenValidity
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.ParentIDs != nil {
		out.ParentIDs = *o.ParentIDs
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Roles != nil {
		out.Roles = *o.Roles
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseAppCredential) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseAppCredential) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseAppCredential) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseAppCredential) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseAppCredential) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseAppCredential) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseAppCredential) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseAppCredential) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseAppCredential) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseAppCredential) SetDescription(description string) {

	o.Description = &description
}

// GetDisabled returns the Disabled of the receiver.
func (o *SparseAppCredential) GetDisabled() (out bool) {

	if o.Disabled == nil {
		return
	}

	return *o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the address of the given value.
func (o *SparseAppCredential) SetDisabled(disabled bool) {

	o.Disabled = &disabled
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseAppCredential) GetMetadata() (out []string) {

	if o.Metadata == nil {
		return
	}

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseAppCredential) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseAppCredential) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseAppCredential) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseAppCredential) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseAppCredential) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseAppCredential) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseAppCredential) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseAppCredential) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseAppCredential) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseAppCredential) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseAppCredential) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseAppCredential) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseAppCredential) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseAppCredential) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseAppCredential) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseAppCredential) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseAppCredential) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseAppCredential) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseAppCredential) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseAppCredential.
func (o *SparseAppCredential) DeepCopy() *SparseAppCredential {

	if o == nil {
		return nil
	}

	out := &SparseAppCredential{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAppCredential.
func (o *SparseAppCredential) DeepCopyInto(out *SparseAppCredential) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAppCredential: %s", err))
	}

	*out = *target.(*SparseAppCredential)
}

type mongoAttributesAppCredential struct {
	ID                     bson.ObjectId       `bson:"_id,omitempty"`
	Annotations            map[string][]string `bson:"annotations"`
	AssociatedTags         []string            `bson:"associatedtags"`
	AuthorizedSubnets      []string            `bson:"authorizedsubnets"`
	Certificate            string              `bson:"certificate"`
	CertificateSN          string              `bson:"certificatesn"`
	CreateIdempotencyKey   string              `bson:"createidempotencykey"`
	CreateTime             time.Time           `bson:"createtime"`
	Description            string              `bson:"description"`
	Disabled               bool                `bson:"disabled"`
	Email                  string              `bson:"email"`
	MaxIssuedTokenValidity string              `bson:"maxissuedtokenvalidity"`
	Metadata               []string            `bson:"metadata"`
	MigrationsLog          map[string]string   `bson:"migrationslog,omitempty"`
	Name                   string              `bson:"name"`
	Namespace              string              `bson:"namespace"`
	NormalizedTags         []string            `bson:"normalizedtags"`
	ParentIDs              []string            `bson:"parentids"`
	Protected              bool                `bson:"protected"`
	Roles                  []string            `bson:"roles"`
	UpdateIdempotencyKey   string              `bson:"updateidempotencykey"`
	UpdateTime             time.Time           `bson:"updatetime"`
	ZHash                  int                 `bson:"zhash"`
	Zone                   int                 `bson:"zone"`
}
type mongoAttributesSparseAppCredential struct {
	ID                     bson.ObjectId        `bson:"_id,omitempty"`
	Annotations            *map[string][]string `bson:"annotations,omitempty"`
	AssociatedTags         *[]string            `bson:"associatedtags,omitempty"`
	AuthorizedSubnets      *[]string            `bson:"authorizedsubnets,omitempty"`
	Certificate            *string              `bson:"certificate,omitempty"`
	CertificateSN          *string              `bson:"certificatesn,omitempty"`
	CreateIdempotencyKey   *string              `bson:"createidempotencykey,omitempty"`
	CreateTime             *time.Time           `bson:"createtime,omitempty"`
	Description            *string              `bson:"description,omitempty"`
	Disabled               *bool                `bson:"disabled,omitempty"`
	Email                  *string              `bson:"email,omitempty"`
	MaxIssuedTokenValidity *string              `bson:"maxissuedtokenvalidity,omitempty"`
	Metadata               *[]string            `bson:"metadata,omitempty"`
	MigrationsLog          *map[string]string   `bson:"migrationslog,omitempty"`
	Name                   *string              `bson:"name,omitempty"`
	Namespace              *string              `bson:"namespace,omitempty"`
	NormalizedTags         *[]string            `bson:"normalizedtags,omitempty"`
	ParentIDs              *[]string            `bson:"parentids,omitempty"`
	Protected              *bool                `bson:"protected,omitempty"`
	Roles                  *[]string            `bson:"roles,omitempty"`
	UpdateIdempotencyKey   *string              `bson:"updateidempotencykey,omitempty"`
	UpdateTime             *time.Time           `bson:"updatetime,omitempty"`
	ZHash                  *int                 `bson:"zhash,omitempty"`
	Zone                   *int                 `bson:"zone,omitempty"`
}
