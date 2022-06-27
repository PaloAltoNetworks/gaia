package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TrustedCATypeValue represents the possible values for attribute "type".
type TrustedCATypeValue string

const (
	// TrustedCATypeJWT represents the value JWT.
	TrustedCATypeJWT TrustedCATypeValue = "JWT"

	// TrustedCATypeSSH represents the value SSH.
	TrustedCATypeSSH TrustedCATypeValue = "SSH"

	// TrustedCATypeX509 represents the value X509.
	TrustedCATypeX509 TrustedCATypeValue = "X509"
)

// TrustedCAIdentity represents the Identity of the object.
var TrustedCAIdentity = elemental.Identity{
	Name:     "trustedca",
	Category: "trustedcas",
	Package:  "squall",
	Private:  false,
}

// TrustedCAsList represents a list of TrustedCAs
type TrustedCAsList []*TrustedCA

// Identity returns the identity of the objects in the list.
func (o TrustedCAsList) Identity() elemental.Identity {

	return TrustedCAIdentity
}

// Copy returns a pointer to a copy the TrustedCAsList.
func (o TrustedCAsList) Copy() elemental.Identifiables {

	copy := append(TrustedCAsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TrustedCAsList.
func (o TrustedCAsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(TrustedCAsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*TrustedCA))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TrustedCAsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TrustedCAsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the TrustedCAsList converted to SparseTrustedCAsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o TrustedCAsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseTrustedCAsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseTrustedCA)
	}

	return out
}

// Version returns the version of the content.
func (o TrustedCAsList) Version() int {

	return 1
}

// TrustedCA represents the model of a trustedca
type TrustedCA struct {
	// The private certificate of the corresponding type associated with this
	// namespace.
	Certificate string `json:"certificate" msgpack:"certificate" bson:"-" mapstructure:"certificate,omitempty"`

	// The controller that this certificate or CA was issued from.
	Controller string `json:"controller,omitempty" msgpack:"controller,omitempty" bson:"-" mapstructure:"controller,omitempty"`

	// The namespace that this certificate or CA was defined at.
	Namespace string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"-" mapstructure:"namespace,omitempty"`

	// The ID of namespace that this certificate or CA was defined at.
	NamespaceID string `json:"namespaceID,omitempty" msgpack:"namespaceID,omitempty" bson:"-" mapstructure:"namespaceID,omitempty"`

	// SerialNumber is the serial number of the certificate.
	Serialnumber string `json:"serialnumber,omitempty" msgpack:"serialnumber,omitempty" bson:"-" mapstructure:"serialnumber,omitempty"`

	// Type of the certificate.
	Type TrustedCATypeValue `json:"type" msgpack:"type" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewTrustedCA returns a new *TrustedCA
func NewTrustedCA() *TrustedCA {

	return &TrustedCA{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *TrustedCA) Identity() elemental.Identity {

	return TrustedCAIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *TrustedCA) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *TrustedCA) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TrustedCA) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesTrustedCA{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TrustedCA) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesTrustedCA{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *TrustedCA) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *TrustedCA) BleveType() string {

	return "trustedca"
}

// DefaultOrder returns the list of default ordering fields.
func (o *TrustedCA) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *TrustedCA) Doc() string {

	return `Represents a trusted certificate authority (CA).`
}

func (o *TrustedCA) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *TrustedCA) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseTrustedCA{
			Certificate:  &o.Certificate,
			Controller:   &o.Controller,
			Namespace:    &o.Namespace,
			NamespaceID:  &o.NamespaceID,
			Serialnumber: &o.Serialnumber,
			Type:         &o.Type,
		}
	}

	sp := &SparseTrustedCA{}
	for _, f := range fields {
		switch f {
		case "certificate":
			sp.Certificate = &(o.Certificate)
		case "controller":
			sp.Controller = &(o.Controller)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "namespaceID":
			sp.NamespaceID = &(o.NamespaceID)
		case "serialnumber":
			sp.Serialnumber = &(o.Serialnumber)
		case "type":
			sp.Type = &(o.Type)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseTrustedCA to the object.
func (o *TrustedCA) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseTrustedCA)
	if so.Certificate != nil {
		o.Certificate = *so.Certificate
	}
	if so.Controller != nil {
		o.Controller = *so.Controller
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NamespaceID != nil {
		o.NamespaceID = *so.NamespaceID
	}
	if so.Serialnumber != nil {
		o.Serialnumber = *so.Serialnumber
	}
	if so.Type != nil {
		o.Type = *so.Type
	}
}

// DeepCopy returns a deep copy if the TrustedCA.
func (o *TrustedCA) DeepCopy() *TrustedCA {

	if o == nil {
		return nil
	}

	out := &TrustedCA{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *TrustedCA.
func (o *TrustedCA) DeepCopyInto(out *TrustedCA) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy TrustedCA: %s", err))
	}

	*out = *target.(*TrustedCA)
}

// Validate valides the current information stored into the structure.
func (o *TrustedCA) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"X509", "SSH", "JWT"}, true); err != nil {
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
func (*TrustedCA) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TrustedCAAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TrustedCALowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*TrustedCA) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TrustedCAAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *TrustedCA) ValueForAttribute(name string) interface{} {

	switch name {
	case "certificate":
		return o.Certificate
	case "controller":
		return o.Controller
	case "namespace":
		return o.Namespace
	case "namespaceID":
		return o.NamespaceID
	case "serialnumber":
		return o.Serialnumber
	case "type":
		return o.Type
	}

	return nil
}

// TrustedCAAttributesMap represents the map of attribute for TrustedCA.
var TrustedCAAttributesMap = map[string]elemental.AttributeSpecification{
	"Certificate": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Certificate",
		Description: `The private certificate of the corresponding type associated with this
namespace.`,
		Exposed:  true,
		Name:     "certificate",
		ReadOnly: true,
		Type:     "string",
	},
	"Controller": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Controller",
		Description:    `The controller that this certificate or CA was issued from.`,
		Exposed:        true,
		Name:           "controller",
		ReadOnly:       true,
		Type:           "string",
	},
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `The namespace that this certificate or CA was defined at.`,
		Exposed:        true,
		Name:           "namespace",
		ReadOnly:       true,
		Type:           "string",
	},
	"NamespaceID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NamespaceID",
		Description:    `The ID of namespace that this certificate or CA was defined at.`,
		Exposed:        true,
		Name:           "namespaceID",
		ReadOnly:       true,
		Type:           "string",
	},
	"Serialnumber": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Serialnumber",
		Description:    `SerialNumber is the serial number of the certificate.`,
		Exposed:        true,
		Name:           "serialnumber",
		ReadOnly:       true,
		Type:           "string",
	},
	"Type": {
		AllowedChoices: []string{"X509", "SSH", "JWT"},
		Autogenerated:  true,
		ConvertedName:  "Type",
		Description:    `Type of the certificate.`,
		Exposed:        true,
		Name:           "type",
		ReadOnly:       true,
		Type:           "enum",
	},
}

// TrustedCALowerCaseAttributesMap represents the map of attribute for TrustedCA.
var TrustedCALowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"certificate": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Certificate",
		Description: `The private certificate of the corresponding type associated with this
namespace.`,
		Exposed:  true,
		Name:     "certificate",
		ReadOnly: true,
		Type:     "string",
	},
	"controller": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Controller",
		Description:    `The controller that this certificate or CA was issued from.`,
		Exposed:        true,
		Name:           "controller",
		ReadOnly:       true,
		Type:           "string",
	},
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `The namespace that this certificate or CA was defined at.`,
		Exposed:        true,
		Name:           "namespace",
		ReadOnly:       true,
		Type:           "string",
	},
	"namespaceid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NamespaceID",
		Description:    `The ID of namespace that this certificate or CA was defined at.`,
		Exposed:        true,
		Name:           "namespaceID",
		ReadOnly:       true,
		Type:           "string",
	},
	"serialnumber": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Serialnumber",
		Description:    `SerialNumber is the serial number of the certificate.`,
		Exposed:        true,
		Name:           "serialnumber",
		ReadOnly:       true,
		Type:           "string",
	},
	"type": {
		AllowedChoices: []string{"X509", "SSH", "JWT"},
		Autogenerated:  true,
		ConvertedName:  "Type",
		Description:    `Type of the certificate.`,
		Exposed:        true,
		Name:           "type",
		ReadOnly:       true,
		Type:           "enum",
	},
}

// SparseTrustedCAsList represents a list of SparseTrustedCAs
type SparseTrustedCAsList []*SparseTrustedCA

// Identity returns the identity of the objects in the list.
func (o SparseTrustedCAsList) Identity() elemental.Identity {

	return TrustedCAIdentity
}

// Copy returns a pointer to a copy the SparseTrustedCAsList.
func (o SparseTrustedCAsList) Copy() elemental.Identifiables {

	copy := append(SparseTrustedCAsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseTrustedCAsList.
func (o SparseTrustedCAsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseTrustedCAsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseTrustedCA))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseTrustedCAsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseTrustedCAsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseTrustedCAsList converted to TrustedCAsList.
func (o SparseTrustedCAsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseTrustedCAsList) Version() int {

	return 1
}

// SparseTrustedCA represents the sparse version of a trustedca.
type SparseTrustedCA struct {
	// The private certificate of the corresponding type associated with this
	// namespace.
	Certificate *string `json:"certificate,omitempty" msgpack:"certificate,omitempty" bson:"-" mapstructure:"certificate,omitempty"`

	// The controller that this certificate or CA was issued from.
	Controller *string `json:"controller,omitempty" msgpack:"controller,omitempty" bson:"-" mapstructure:"controller,omitempty"`

	// The namespace that this certificate or CA was defined at.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"-" mapstructure:"namespace,omitempty"`

	// The ID of namespace that this certificate or CA was defined at.
	NamespaceID *string `json:"namespaceID,omitempty" msgpack:"namespaceID,omitempty" bson:"-" mapstructure:"namespaceID,omitempty"`

	// SerialNumber is the serial number of the certificate.
	Serialnumber *string `json:"serialnumber,omitempty" msgpack:"serialnumber,omitempty" bson:"-" mapstructure:"serialnumber,omitempty"`

	// Type of the certificate.
	Type *TrustedCATypeValue `json:"type,omitempty" msgpack:"type,omitempty" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseTrustedCA returns a new  SparseTrustedCA.
func NewSparseTrustedCA() *SparseTrustedCA {
	return &SparseTrustedCA{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseTrustedCA) Identity() elemental.Identity {

	return TrustedCAIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseTrustedCA) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseTrustedCA) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTrustedCA) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseTrustedCA{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTrustedCA) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseTrustedCA{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseTrustedCA) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseTrustedCA) ToPlain() elemental.PlainIdentifiable {

	out := NewTrustedCA()
	if o.Certificate != nil {
		out.Certificate = *o.Certificate
	}
	if o.Controller != nil {
		out.Controller = *o.Controller
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NamespaceID != nil {
		out.NamespaceID = *o.NamespaceID
	}
	if o.Serialnumber != nil {
		out.Serialnumber = *o.Serialnumber
	}
	if o.Type != nil {
		out.Type = *o.Type
	}

	return out
}

// DeepCopy returns a deep copy if the SparseTrustedCA.
func (o *SparseTrustedCA) DeepCopy() *SparseTrustedCA {

	if o == nil {
		return nil
	}

	out := &SparseTrustedCA{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseTrustedCA.
func (o *SparseTrustedCA) DeepCopyInto(out *SparseTrustedCA) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseTrustedCA: %s", err))
	}

	*out = *target.(*SparseTrustedCA)
}

type mongoAttributesTrustedCA struct {
}
type mongoAttributesSparseTrustedCA struct {
}
