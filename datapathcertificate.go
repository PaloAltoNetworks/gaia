package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// DataPathCertificateTypeValue represents the possible values for attribute "type".
type DataPathCertificateTypeValue string

const (
	// DataPathCertificateTypeEnforcer represents the value Enforcer.
	DataPathCertificateTypeEnforcer DataPathCertificateTypeValue = "Enforcer"

	// DataPathCertificateTypeK8SAPIServer represents the value K8SAPIServer.
	DataPathCertificateTypeK8SAPIServer DataPathCertificateTypeValue = "K8SAPIServer"

	// DataPathCertificateTypeService represents the value Service.
	DataPathCertificateTypeService DataPathCertificateTypeValue = "Service"

	// DataPathCertificateTypeServicePing represents the value ServicePing.
	DataPathCertificateTypeServicePing DataPathCertificateTypeValue = "ServicePing"
)

// DataPathCertificateIdentity represents the Identity of the object.
var DataPathCertificateIdentity = elemental.Identity{
	Name:     "datapathcertificate",
	Category: "datapathcertificates",
	Package:  "squall",
	Private:  false,
}

// DataPathCertificatesList represents a list of DataPathCertificates
type DataPathCertificatesList []*DataPathCertificate

// Identity returns the identity of the objects in the list.
func (o DataPathCertificatesList) Identity() elemental.Identity {

	return DataPathCertificateIdentity
}

// Copy returns a pointer to a copy the DataPathCertificatesList.
func (o DataPathCertificatesList) Copy() elemental.Identifiables {

	copy := append(DataPathCertificatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the DataPathCertificatesList.
func (o DataPathCertificatesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(DataPathCertificatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*DataPathCertificate))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o DataPathCertificatesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o DataPathCertificatesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the DataPathCertificatesList converted to SparseDataPathCertificatesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o DataPathCertificatesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseDataPathCertificatesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseDataPathCertificate)
	}

	return out
}

// Version returns the version of the content.
func (o DataPathCertificatesList) Version() int {

	return 1
}

// DataPathCertificate represents the model of a datapathcertificate
type DataPathCertificate struct {
	// Contains a certificate signing request (CSR) from the enforcer. Depending on the
	// certificate there will be various requirements for the Microsegmentation Console
	// to accept the CSR.
	CSR string `json:"CSR" msgpack:"CSR" bson:"-" mapstructure:"CSR,omitempty"`

	// The certificate.
	Certificate string `json:"certificate" msgpack:"certificate" bson:"-" mapstructure:"certificate,omitempty"`

	// ID of the object you want to issue a certificate for.
	ObjectID string `json:"objectID" msgpack:"objectID" bson:"-" mapstructure:"objectID,omitempty"`

	// Provides the session ID of the enforcer when retrieving a datapath certificate.
	SessionID string `json:"sessionID" msgpack:"sessionID" bson:"-" mapstructure:"sessionID,omitempty"`

	// Contains the CA that signed the delivered certificate.
	Signer string `json:"signer" msgpack:"signer" bson:"-" mapstructure:"signer,omitempty"`

	// Contains a cryptographic token.
	Token string `json:"token" msgpack:"token" bson:"-" mapstructure:"token,omitempty"`

	// Type of certificate.
	Type DataPathCertificateTypeValue `json:"type" msgpack:"type" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewDataPathCertificate returns a new *DataPathCertificate
func NewDataPathCertificate() *DataPathCertificate {

	return &DataPathCertificate{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *DataPathCertificate) Identity() elemental.Identity {

	return DataPathCertificateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DataPathCertificate) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DataPathCertificate) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *DataPathCertificate) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesDataPathCertificate{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *DataPathCertificate) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesDataPathCertificate{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *DataPathCertificate) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *DataPathCertificate) BleveType() string {

	return "datapathcertificate"
}

// DefaultOrder returns the list of default ordering fields.
func (o *DataPathCertificate) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *DataPathCertificate) Doc() string {

	return `Used by enforcer instances to retrieve various certificates used
for the datapath.`
}

func (o *DataPathCertificate) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *DataPathCertificate) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseDataPathCertificate{
			CSR:         &o.CSR,
			Certificate: &o.Certificate,
			ObjectID:    &o.ObjectID,
			SessionID:   &o.SessionID,
			Signer:      &o.Signer,
			Token:       &o.Token,
			Type:        &o.Type,
		}
	}

	sp := &SparseDataPathCertificate{}
	for _, f := range fields {
		switch f {
		case "CSR":
			sp.CSR = &(o.CSR)
		case "certificate":
			sp.Certificate = &(o.Certificate)
		case "objectID":
			sp.ObjectID = &(o.ObjectID)
		case "sessionID":
			sp.SessionID = &(o.SessionID)
		case "signer":
			sp.Signer = &(o.Signer)
		case "token":
			sp.Token = &(o.Token)
		case "type":
			sp.Type = &(o.Type)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseDataPathCertificate to the object.
func (o *DataPathCertificate) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseDataPathCertificate)
	if so.CSR != nil {
		o.CSR = *so.CSR
	}
	if so.Certificate != nil {
		o.Certificate = *so.Certificate
	}
	if so.ObjectID != nil {
		o.ObjectID = *so.ObjectID
	}
	if so.SessionID != nil {
		o.SessionID = *so.SessionID
	}
	if so.Signer != nil {
		o.Signer = *so.Signer
	}
	if so.Token != nil {
		o.Token = *so.Token
	}
	if so.Type != nil {
		o.Type = *so.Type
	}
}

// DeepCopy returns a deep copy if the DataPathCertificate.
func (o *DataPathCertificate) DeepCopy() *DataPathCertificate {

	if o == nil {
		return nil
	}

	out := &DataPathCertificate{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *DataPathCertificate.
func (o *DataPathCertificate) DeepCopyInto(out *DataPathCertificate) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy DataPathCertificate: %s", err))
	}

	*out = *target.(*DataPathCertificate)
}

// Validate valides the current information stored into the structure.
func (o *DataPathCertificate) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("CSR", o.CSR); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("objectID", o.ObjectID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Enforcer", "Service", "ServicePing", "K8SAPIServer"}, false); err != nil {
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
func (*DataPathCertificate) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := DataPathCertificateAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return DataPathCertificateLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*DataPathCertificate) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return DataPathCertificateAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *DataPathCertificate) ValueForAttribute(name string) interface{} {

	switch name {
	case "CSR":
		return o.CSR
	case "certificate":
		return o.Certificate
	case "objectID":
		return o.ObjectID
	case "sessionID":
		return o.SessionID
	case "signer":
		return o.Signer
	case "token":
		return o.Token
	case "type":
		return o.Type
	}

	return nil
}

// DataPathCertificateAttributesMap represents the map of attribute for DataPathCertificate.
var DataPathCertificateAttributesMap = map[string]elemental.AttributeSpecification{
	"CSR": {
		AllowedChoices: []string{},
		ConvertedName:  "CSR",
		Description: `Contains a certificate signing request (CSR) from the enforcer. Depending on the
certificate there will be various requirements for the Microsegmentation Console
to accept the CSR.`,
		Exposed:  true,
		Name:     "CSR",
		Required: true,
		Type:     "string",
	},
	"Certificate": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Certificate",
		Description:    `The certificate.`,
		Exposed:        true,
		Name:           "certificate",
		ReadOnly:       true,
		Type:           "string",
	},
	"ObjectID": {
		AllowedChoices: []string{},
		ConvertedName:  "ObjectID",
		Description:    `ID of the object you want to issue a certificate for.`,
		Exposed:        true,
		Name:           "objectID",
		Required:       true,
		Type:           "string",
	},
	"SessionID": {
		AllowedChoices: []string{},
		ConvertedName:  "SessionID",
		Description:    `Provides the session ID of the enforcer when retrieving a datapath certificate.`,
		Exposed:        true,
		Name:           "sessionID",
		Type:           "string",
	},
	"Signer": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Signer",
		Description:    `Contains the CA that signed the delivered certificate.`,
		Exposed:        true,
		Name:           "signer",
		ReadOnly:       true,
		Type:           "string",
	},
	"Token": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Token",
		Description:    `Contains a cryptographic token.`,
		Exposed:        true,
		Name:           "token",
		ReadOnly:       true,
		Type:           "string",
	},
	"Type": {
		AllowedChoices: []string{"Enforcer", "Service", "ServicePing", "K8SAPIServer"},
		ConvertedName:  "Type",
		Description:    `Type of certificate.`,
		Exposed:        true,
		Name:           "type",
		Type:           "enum",
	},
}

// DataPathCertificateLowerCaseAttributesMap represents the map of attribute for DataPathCertificate.
var DataPathCertificateLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"csr": {
		AllowedChoices: []string{},
		ConvertedName:  "CSR",
		Description: `Contains a certificate signing request (CSR) from the enforcer. Depending on the
certificate there will be various requirements for the Microsegmentation Console
to accept the CSR.`,
		Exposed:  true,
		Name:     "CSR",
		Required: true,
		Type:     "string",
	},
	"certificate": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Certificate",
		Description:    `The certificate.`,
		Exposed:        true,
		Name:           "certificate",
		ReadOnly:       true,
		Type:           "string",
	},
	"objectid": {
		AllowedChoices: []string{},
		ConvertedName:  "ObjectID",
		Description:    `ID of the object you want to issue a certificate for.`,
		Exposed:        true,
		Name:           "objectID",
		Required:       true,
		Type:           "string",
	},
	"sessionid": {
		AllowedChoices: []string{},
		ConvertedName:  "SessionID",
		Description:    `Provides the session ID of the enforcer when retrieving a datapath certificate.`,
		Exposed:        true,
		Name:           "sessionID",
		Type:           "string",
	},
	"signer": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Signer",
		Description:    `Contains the CA that signed the delivered certificate.`,
		Exposed:        true,
		Name:           "signer",
		ReadOnly:       true,
		Type:           "string",
	},
	"token": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Token",
		Description:    `Contains a cryptographic token.`,
		Exposed:        true,
		Name:           "token",
		ReadOnly:       true,
		Type:           "string",
	},
	"type": {
		AllowedChoices: []string{"Enforcer", "Service", "ServicePing", "K8SAPIServer"},
		ConvertedName:  "Type",
		Description:    `Type of certificate.`,
		Exposed:        true,
		Name:           "type",
		Type:           "enum",
	},
}

// SparseDataPathCertificatesList represents a list of SparseDataPathCertificates
type SparseDataPathCertificatesList []*SparseDataPathCertificate

// Identity returns the identity of the objects in the list.
func (o SparseDataPathCertificatesList) Identity() elemental.Identity {

	return DataPathCertificateIdentity
}

// Copy returns a pointer to a copy the SparseDataPathCertificatesList.
func (o SparseDataPathCertificatesList) Copy() elemental.Identifiables {

	copy := append(SparseDataPathCertificatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseDataPathCertificatesList.
func (o SparseDataPathCertificatesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseDataPathCertificatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseDataPathCertificate))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseDataPathCertificatesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseDataPathCertificatesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseDataPathCertificatesList converted to DataPathCertificatesList.
func (o SparseDataPathCertificatesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseDataPathCertificatesList) Version() int {

	return 1
}

// SparseDataPathCertificate represents the sparse version of a datapathcertificate.
type SparseDataPathCertificate struct {
	// Contains a certificate signing request (CSR) from the enforcer. Depending on the
	// certificate there will be various requirements for the Microsegmentation Console
	// to accept the CSR.
	CSR *string `json:"CSR,omitempty" msgpack:"CSR,omitempty" bson:"-" mapstructure:"CSR,omitempty"`

	// The certificate.
	Certificate *string `json:"certificate,omitempty" msgpack:"certificate,omitempty" bson:"-" mapstructure:"certificate,omitempty"`

	// ID of the object you want to issue a certificate for.
	ObjectID *string `json:"objectID,omitempty" msgpack:"objectID,omitempty" bson:"-" mapstructure:"objectID,omitempty"`

	// Provides the session ID of the enforcer when retrieving a datapath certificate.
	SessionID *string `json:"sessionID,omitempty" msgpack:"sessionID,omitempty" bson:"-" mapstructure:"sessionID,omitempty"`

	// Contains the CA that signed the delivered certificate.
	Signer *string `json:"signer,omitempty" msgpack:"signer,omitempty" bson:"-" mapstructure:"signer,omitempty"`

	// Contains a cryptographic token.
	Token *string `json:"token,omitempty" msgpack:"token,omitempty" bson:"-" mapstructure:"token,omitempty"`

	// Type of certificate.
	Type *DataPathCertificateTypeValue `json:"type,omitempty" msgpack:"type,omitempty" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseDataPathCertificate returns a new  SparseDataPathCertificate.
func NewSparseDataPathCertificate() *SparseDataPathCertificate {
	return &SparseDataPathCertificate{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseDataPathCertificate) Identity() elemental.Identity {

	return DataPathCertificateIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseDataPathCertificate) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseDataPathCertificate) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseDataPathCertificate) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseDataPathCertificate{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseDataPathCertificate) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseDataPathCertificate{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseDataPathCertificate) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseDataPathCertificate) ToPlain() elemental.PlainIdentifiable {

	out := NewDataPathCertificate()
	if o.CSR != nil {
		out.CSR = *o.CSR
	}
	if o.Certificate != nil {
		out.Certificate = *o.Certificate
	}
	if o.ObjectID != nil {
		out.ObjectID = *o.ObjectID
	}
	if o.SessionID != nil {
		out.SessionID = *o.SessionID
	}
	if o.Signer != nil {
		out.Signer = *o.Signer
	}
	if o.Token != nil {
		out.Token = *o.Token
	}
	if o.Type != nil {
		out.Type = *o.Type
	}

	return out
}

// DeepCopy returns a deep copy if the SparseDataPathCertificate.
func (o *SparseDataPathCertificate) DeepCopy() *SparseDataPathCertificate {

	if o == nil {
		return nil
	}

	out := &SparseDataPathCertificate{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseDataPathCertificate.
func (o *SparseDataPathCertificate) DeepCopyInto(out *SparseDataPathCertificate) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseDataPathCertificate: %s", err))
	}

	*out = *target.(*SparseDataPathCertificate)
}

type mongoAttributesDataPathCertificate struct {
}
type mongoAttributesSparseDataPathCertificate struct {
}
