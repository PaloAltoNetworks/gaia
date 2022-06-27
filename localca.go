package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// LocalCAIdentity represents the Identity of the object.
var LocalCAIdentity = elemental.Identity{
	Name:     "localca",
	Category: "localcas",
	Package:  "squall",
	Private:  false,
}

// LocalCAsList represents a list of LocalCAs
type LocalCAsList []*LocalCA

// Identity returns the identity of the objects in the list.
func (o LocalCAsList) Identity() elemental.Identity {

	return LocalCAIdentity
}

// Copy returns a pointer to a copy the LocalCAsList.
func (o LocalCAsList) Copy() elemental.Identifiables {

	copy := append(LocalCAsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the LocalCAsList.
func (o LocalCAsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(LocalCAsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*LocalCA))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o LocalCAsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o LocalCAsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the LocalCAsList converted to SparseLocalCAsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o LocalCAsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseLocalCAsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseLocalCA)
	}

	return out
}

// Version returns the version of the content.
func (o LocalCAsList) Version() int {

	return 1
}

// LocalCA represents the model of a localca
type LocalCA struct {
	// The SSH certificate authority used by the namespace.
	SSHCertificate string `json:"SSHCertificate" msgpack:"SSHCertificate" bson:"-" mapstructure:"SSHCertificate,omitempty"`

	// Set to `true` to renew the SSH certificate authority of the namespace.
	SSHCertificateRenew bool `json:"SSHCertificateRenew" msgpack:"SSHCertificateRenew" bson:"-" mapstructure:"SSHCertificateRenew,omitempty"`

	// The certificate authority used by the namespace.
	Certificate string `json:"certificate" msgpack:"certificate" bson:"-" mapstructure:"certificate,omitempty"`

	// Set to `true` to renew the certificate authority of the namespace.
	CertificateRenew bool `json:"certificateRenew" msgpack:"certificateRenew" bson:"-" mapstructure:"certificateRenew,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewLocalCA returns a new *LocalCA
func NewLocalCA() *LocalCA {

	return &LocalCA{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *LocalCA) Identity() elemental.Identity {

	return LocalCAIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *LocalCA) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *LocalCA) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *LocalCA) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesLocalCA{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *LocalCA) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesLocalCA{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *LocalCA) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *LocalCA) BleveType() string {

	return "localca"
}

// DefaultOrder returns the list of default ordering fields.
func (o *LocalCA) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *LocalCA) Doc() string {

	return `Can be used to retrieve or renew the local and SSH certificate authorities of
the namespace.`
}

func (o *LocalCA) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *LocalCA) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseLocalCA{
			SSHCertificate:      &o.SSHCertificate,
			SSHCertificateRenew: &o.SSHCertificateRenew,
			Certificate:         &o.Certificate,
			CertificateRenew:    &o.CertificateRenew,
		}
	}

	sp := &SparseLocalCA{}
	for _, f := range fields {
		switch f {
		case "SSHCertificate":
			sp.SSHCertificate = &(o.SSHCertificate)
		case "SSHCertificateRenew":
			sp.SSHCertificateRenew = &(o.SSHCertificateRenew)
		case "certificate":
			sp.Certificate = &(o.Certificate)
		case "certificateRenew":
			sp.CertificateRenew = &(o.CertificateRenew)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseLocalCA to the object.
func (o *LocalCA) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseLocalCA)
	if so.SSHCertificate != nil {
		o.SSHCertificate = *so.SSHCertificate
	}
	if so.SSHCertificateRenew != nil {
		o.SSHCertificateRenew = *so.SSHCertificateRenew
	}
	if so.Certificate != nil {
		o.Certificate = *so.Certificate
	}
	if so.CertificateRenew != nil {
		o.CertificateRenew = *so.CertificateRenew
	}
}

// DeepCopy returns a deep copy if the LocalCA.
func (o *LocalCA) DeepCopy() *LocalCA {

	if o == nil {
		return nil
	}

	out := &LocalCA{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *LocalCA.
func (o *LocalCA) DeepCopyInto(out *LocalCA) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy LocalCA: %s", err))
	}

	*out = *target.(*LocalCA)
}

// Validate valides the current information stored into the structure.
func (o *LocalCA) Validate() error {

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
func (*LocalCA) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := LocalCAAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return LocalCALowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*LocalCA) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return LocalCAAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *LocalCA) ValueForAttribute(name string) interface{} {

	switch name {
	case "SSHCertificate":
		return o.SSHCertificate
	case "SSHCertificateRenew":
		return o.SSHCertificateRenew
	case "certificate":
		return o.Certificate
	case "certificateRenew":
		return o.CertificateRenew
	}

	return nil
}

// LocalCAAttributesMap represents the map of attribute for LocalCA.
var LocalCAAttributesMap = map[string]elemental.AttributeSpecification{
	"SSHCertificate": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SSHCertificate",
		Description:    `The SSH certificate authority used by the namespace.`,
		Exposed:        true,
		Name:           "SSHCertificate",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"SSHCertificateRenew": {
		AllowedChoices: []string{},
		ConvertedName:  "SSHCertificateRenew",
		Description:    `Set to ` + "`" + `true` + "`" + ` to renew the SSH certificate authority of the namespace.`,
		Exposed:        true,
		Name:           "SSHCertificateRenew",
		Transient:      true,
		Type:           "boolean",
	},
	"Certificate": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Certificate",
		Description:    `The certificate authority used by the namespace.`,
		Exposed:        true,
		Name:           "certificate",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"CertificateRenew": {
		AllowedChoices: []string{},
		ConvertedName:  "CertificateRenew",
		Description:    `Set to ` + "`" + `true` + "`" + ` to renew the certificate authority of the namespace.`,
		Exposed:        true,
		Name:           "certificateRenew",
		Transient:      true,
		Type:           "boolean",
	},
}

// LocalCALowerCaseAttributesMap represents the map of attribute for LocalCA.
var LocalCALowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"sshcertificate": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SSHCertificate",
		Description:    `The SSH certificate authority used by the namespace.`,
		Exposed:        true,
		Name:           "SSHCertificate",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"sshcertificaterenew": {
		AllowedChoices: []string{},
		ConvertedName:  "SSHCertificateRenew",
		Description:    `Set to ` + "`" + `true` + "`" + ` to renew the SSH certificate authority of the namespace.`,
		Exposed:        true,
		Name:           "SSHCertificateRenew",
		Transient:      true,
		Type:           "boolean",
	},
	"certificate": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Certificate",
		Description:    `The certificate authority used by the namespace.`,
		Exposed:        true,
		Name:           "certificate",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"certificaterenew": {
		AllowedChoices: []string{},
		ConvertedName:  "CertificateRenew",
		Description:    `Set to ` + "`" + `true` + "`" + ` to renew the certificate authority of the namespace.`,
		Exposed:        true,
		Name:           "certificateRenew",
		Transient:      true,
		Type:           "boolean",
	},
}

// SparseLocalCAsList represents a list of SparseLocalCAs
type SparseLocalCAsList []*SparseLocalCA

// Identity returns the identity of the objects in the list.
func (o SparseLocalCAsList) Identity() elemental.Identity {

	return LocalCAIdentity
}

// Copy returns a pointer to a copy the SparseLocalCAsList.
func (o SparseLocalCAsList) Copy() elemental.Identifiables {

	copy := append(SparseLocalCAsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseLocalCAsList.
func (o SparseLocalCAsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseLocalCAsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseLocalCA))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseLocalCAsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseLocalCAsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseLocalCAsList converted to LocalCAsList.
func (o SparseLocalCAsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseLocalCAsList) Version() int {

	return 1
}

// SparseLocalCA represents the sparse version of a localca.
type SparseLocalCA struct {
	// The SSH certificate authority used by the namespace.
	SSHCertificate *string `json:"SSHCertificate,omitempty" msgpack:"SSHCertificate,omitempty" bson:"-" mapstructure:"SSHCertificate,omitempty"`

	// Set to `true` to renew the SSH certificate authority of the namespace.
	SSHCertificateRenew *bool `json:"SSHCertificateRenew,omitempty" msgpack:"SSHCertificateRenew,omitempty" bson:"-" mapstructure:"SSHCertificateRenew,omitempty"`

	// The certificate authority used by the namespace.
	Certificate *string `json:"certificate,omitempty" msgpack:"certificate,omitempty" bson:"-" mapstructure:"certificate,omitempty"`

	// Set to `true` to renew the certificate authority of the namespace.
	CertificateRenew *bool `json:"certificateRenew,omitempty" msgpack:"certificateRenew,omitempty" bson:"-" mapstructure:"certificateRenew,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseLocalCA returns a new  SparseLocalCA.
func NewSparseLocalCA() *SparseLocalCA {
	return &SparseLocalCA{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseLocalCA) Identity() elemental.Identity {

	return LocalCAIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseLocalCA) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseLocalCA) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseLocalCA) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseLocalCA{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseLocalCA) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseLocalCA{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseLocalCA) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseLocalCA) ToPlain() elemental.PlainIdentifiable {

	out := NewLocalCA()
	if o.SSHCertificate != nil {
		out.SSHCertificate = *o.SSHCertificate
	}
	if o.SSHCertificateRenew != nil {
		out.SSHCertificateRenew = *o.SSHCertificateRenew
	}
	if o.Certificate != nil {
		out.Certificate = *o.Certificate
	}
	if o.CertificateRenew != nil {
		out.CertificateRenew = *o.CertificateRenew
	}

	return out
}

// DeepCopy returns a deep copy if the SparseLocalCA.
func (o *SparseLocalCA) DeepCopy() *SparseLocalCA {

	if o == nil {
		return nil
	}

	out := &SparseLocalCA{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseLocalCA.
func (o *SparseLocalCA) DeepCopyInto(out *SparseLocalCA) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseLocalCA: %s", err))
	}

	*out = *target.(*SparseLocalCA)
}

type mongoAttributesLocalCA struct {
}
type mongoAttributesSparseLocalCA struct {
}
