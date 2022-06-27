package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// IPInfoIdentity represents the Identity of the object.
var IPInfoIdentity = elemental.Identity{
	Name:     "ipinfo",
	Category: "ipinfos",
	Package:  "canyon",
	Private:  false,
}

// IPInfosList represents a list of IPInfos
type IPInfosList []*IPInfo

// Identity returns the identity of the objects in the list.
func (o IPInfosList) Identity() elemental.Identity {

	return IPInfoIdentity
}

// Copy returns a pointer to a copy the IPInfosList.
func (o IPInfosList) Copy() elemental.Identifiables {

	copy := append(IPInfosList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the IPInfosList.
func (o IPInfosList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(IPInfosList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*IPInfo))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o IPInfosList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o IPInfosList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the IPInfosList converted to SparseIPInfosList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o IPInfosList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseIPInfosList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseIPInfo)
	}

	return out
}

// Version returns the version of the content.
func (o IPInfosList) Version() int {

	return 1
}

// IPInfo represents the model of a ipinfo
type IPInfo struct {
	// The IP address.
	IP string `json:"IP" msgpack:"IP" bson:"-" mapstructure:"IP,omitempty"`

	// Error that occurred during resolution.
	Error string `json:"error" msgpack:"error" bson:"-" mapstructure:"error,omitempty"`

	// List of DNS records associated with the IP address.
	Records map[string]string `json:"records" msgpack:"records" bson:"-" mapstructure:"records,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewIPInfo returns a new *IPInfo
func NewIPInfo() *IPInfo {

	return &IPInfo{
		ModelVersion: 1,
		Records:      map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *IPInfo) Identity() elemental.Identity {

	return IPInfoIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IPInfo) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IPInfo) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *IPInfo) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesIPInfo{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *IPInfo) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesIPInfo{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *IPInfo) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *IPInfo) BleveType() string {

	return "ipinfo"
}

// DefaultOrder returns the list of default ordering fields.
func (o *IPInfo) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *IPInfo) Doc() string {

	return `Provides information about IP address resolution.`
}

func (o *IPInfo) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *IPInfo) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseIPInfo{
			IP:      &o.IP,
			Error:   &o.Error,
			Records: &o.Records,
		}
	}

	sp := &SparseIPInfo{}
	for _, f := range fields {
		switch f {
		case "IP":
			sp.IP = &(o.IP)
		case "error":
			sp.Error = &(o.Error)
		case "records":
			sp.Records = &(o.Records)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseIPInfo to the object.
func (o *IPInfo) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseIPInfo)
	if so.IP != nil {
		o.IP = *so.IP
	}
	if so.Error != nil {
		o.Error = *so.Error
	}
	if so.Records != nil {
		o.Records = *so.Records
	}
}

// DeepCopy returns a deep copy if the IPInfo.
func (o *IPInfo) DeepCopy() *IPInfo {

	if o == nil {
		return nil
	}

	out := &IPInfo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *IPInfo.
func (o *IPInfo) DeepCopyInto(out *IPInfo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy IPInfo: %s", err))
	}

	*out = *target.(*IPInfo)
}

// Validate valides the current information stored into the structure.
func (o *IPInfo) Validate() error {

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
func (*IPInfo) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := IPInfoAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return IPInfoLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*IPInfo) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return IPInfoAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *IPInfo) ValueForAttribute(name string) interface{} {

	switch name {
	case "IP":
		return o.IP
	case "error":
		return o.Error
	case "records":
		return o.Records
	}

	return nil
}

// IPInfoAttributesMap represents the map of attribute for IPInfo.
var IPInfoAttributesMap = map[string]elemental.AttributeSpecification{
	"IP": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "IP",
		Description:    `The IP address.`,
		Exposed:        true,
		Name:           "IP",
		ReadOnly:       true,
		Type:           "string",
	},
	"Error": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Error",
		Description:    `Error that occurred during resolution.`,
		Exposed:        true,
		Name:           "error",
		ReadOnly:       true,
		Type:           "string",
	},
	"Records": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Records",
		Description:    `List of DNS records associated with the IP address.`,
		Exposed:        true,
		Name:           "records",
		ReadOnly:       true,
		SubType:        "map[string]string",
		Type:           "external",
	},
}

// IPInfoLowerCaseAttributesMap represents the map of attribute for IPInfo.
var IPInfoLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"ip": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "IP",
		Description:    `The IP address.`,
		Exposed:        true,
		Name:           "IP",
		ReadOnly:       true,
		Type:           "string",
	},
	"error": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Error",
		Description:    `Error that occurred during resolution.`,
		Exposed:        true,
		Name:           "error",
		ReadOnly:       true,
		Type:           "string",
	},
	"records": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Records",
		Description:    `List of DNS records associated with the IP address.`,
		Exposed:        true,
		Name:           "records",
		ReadOnly:       true,
		SubType:        "map[string]string",
		Type:           "external",
	},
}

// SparseIPInfosList represents a list of SparseIPInfos
type SparseIPInfosList []*SparseIPInfo

// Identity returns the identity of the objects in the list.
func (o SparseIPInfosList) Identity() elemental.Identity {

	return IPInfoIdentity
}

// Copy returns a pointer to a copy the SparseIPInfosList.
func (o SparseIPInfosList) Copy() elemental.Identifiables {

	copy := append(SparseIPInfosList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseIPInfosList.
func (o SparseIPInfosList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseIPInfosList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseIPInfo))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseIPInfosList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseIPInfosList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseIPInfosList converted to IPInfosList.
func (o SparseIPInfosList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseIPInfosList) Version() int {

	return 1
}

// SparseIPInfo represents the sparse version of a ipinfo.
type SparseIPInfo struct {
	// The IP address.
	IP *string `json:"IP,omitempty" msgpack:"IP,omitempty" bson:"-" mapstructure:"IP,omitempty"`

	// Error that occurred during resolution.
	Error *string `json:"error,omitempty" msgpack:"error,omitempty" bson:"-" mapstructure:"error,omitempty"`

	// List of DNS records associated with the IP address.
	Records *map[string]string `json:"records,omitempty" msgpack:"records,omitempty" bson:"-" mapstructure:"records,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseIPInfo returns a new  SparseIPInfo.
func NewSparseIPInfo() *SparseIPInfo {
	return &SparseIPInfo{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseIPInfo) Identity() elemental.Identity {

	return IPInfoIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseIPInfo) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseIPInfo) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseIPInfo) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseIPInfo{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseIPInfo) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseIPInfo{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseIPInfo) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseIPInfo) ToPlain() elemental.PlainIdentifiable {

	out := NewIPInfo()
	if o.IP != nil {
		out.IP = *o.IP
	}
	if o.Error != nil {
		out.Error = *o.Error
	}
	if o.Records != nil {
		out.Records = *o.Records
	}

	return out
}

// DeepCopy returns a deep copy if the SparseIPInfo.
func (o *SparseIPInfo) DeepCopy() *SparseIPInfo {

	if o == nil {
		return nil
	}

	out := &SparseIPInfo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseIPInfo.
func (o *SparseIPInfo) DeepCopyInto(out *SparseIPInfo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseIPInfo: %s", err))
	}

	*out = *target.(*SparseIPInfo)
}

type mongoAttributesIPInfo struct {
}
type mongoAttributesSparseIPInfo struct {
}
